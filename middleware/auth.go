package middleware

import (
	cryptoRand "crypto/rand"
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Session struct {
	ID uuid.UUID
	Base64ID string
	Expiry time.Time
	Model string
	ModelID []byte
}

type redisSession struct {
	Model string
	ModelID []byte
	Secret []byte
}

var (
	expiryDuration time.Duration
)

func init() {
	expiryDuration = time.Minute
	reader := cryptoRand.Reader
	seedBytes := []byte{0,0,0,0,0,0,0,0}
	num, err := reader.Read(seedBytes)
	if num < 8 || err != nil {
		panic("Unable to generate randomized seed!")
	}
	seed := binary.BigEndian.Uint64(seedBytes)
	rand.Seed(int64(seed))
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
		redisClient := ctx.MustGet("redisConnection").(*redis.Client)
		redisPrefix := ctx.MustGet("redisPrefix").(string)

		header := ctx.GetHeader("Authorization")
		if len(header) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 && headerParts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		secretString, err := ctx.Cookie("sessionSecret")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		secret, err := base64.StdEncoding.DecodeString(secretString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		sessionID := headerParts[1]

		data, err := redisClient.Get(redisPrefix + "session:" + sessionID).Result()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		redisSession := redisSession{}
		err = json.Unmarshal([]byte(data), &redisSession)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "internal server error"})
			return
		}

		if subtle.ConstantTimeCompare(redisSession.Secret, secret) != 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		expiryTime := time.Now().Add(expiryDuration)
		redisClient.Set(redisPrefix + "session:" + sessionID, data, expiryDuration)

		maxAge := int(math.Round(expiryDuration.Seconds()))
		domain := strings.Split(ctx.Request.Host, ":")[0]
		ctx.SetCookie("sessionSecret", secretString, maxAge, "/", domain, false, true)

		ctx.Set("AuthModelType", redisSession.Model)
		ctx.Set("AuthModelID", redisSession.ModelID)
		ctx.Set("SessionID", sessionID)
		ctx.Set("ExpiryTime", expiryTime)

		_ = dbCon
	}
}

func AuthCreateSession(ctx *gin.Context, model string, modelID []byte) (*Session, error) {
	redisClient := ctx.MustGet("redisConnection").(*redis.Client)
	redisPrefix := ctx.MustGet("redisPrefix").(string)

	sessionID, _ := uuid.NewV4()
	base64ID := base64.StdEncoding.EncodeToString(sessionID.Bytes())
	secret := []byte{0,0,0,0,0,0,0,0}
	readLength, err := rand.Read(secret)
	if readLength != len(secret) {
		return nil, errors.New("unable to generate session secret")
	} else if err != nil {
		return nil, err
	}

	base64Secret := base64.StdEncoding.EncodeToString(secret)
	redisSession := redisSession{model, modelID, secret}
	encodedRedisSession, err := json.Marshal(redisSession)
	if err != nil {
		return nil, errors.Wrap(err, "Could not serialize session")
	}

	redisClient.Set(redisPrefix + "session:" + base64ID, encodedRedisSession, expiryDuration)
	maxAge := int(math.Round(expiryDuration.Seconds()))
	domain := strings.Split(ctx.Request.Host, ":")[0]
	ctx.SetCookie("sessionSecret", base64Secret, maxAge, "/", domain, false, true)

	session := Session{sessionID, base64ID, time.Now().Add(expiryDuration),  model, modelID}
	return &session, nil
}