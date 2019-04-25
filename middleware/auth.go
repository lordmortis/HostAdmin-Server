package middleware

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/gofrs/uuid"
	"math"
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
}

var (
	expiryDuration time.Duration
)

func init() {
	expiryDuration = time.Minute
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
		redisClient := ctx.MustGet("redisConnection").(*redis.Client)
		redisPrefix := ctx.MustGet("redisPrefix").(string)

		sessionID, err := ctx.Cookie("sessionID")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

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

		redisClient.Set(redisPrefix + "session:" + sessionID, data, expiryDuration)

		maxAge := int(math.Round(expiryDuration.Seconds()))
		domain := strings.Split(ctx.Request.Host, ":")[0]
		ctx.SetCookie("sessionID", sessionID, maxAge, "/", domain, false, true)

		ctx.Set("AuthModelType", redisSession.Model)
		ctx.Set("AuthModelID", redisSession.ModelID)

		_ = dbCon
	}
}

func AuthCreateSession(ctx *gin.Context, model string, modelID []byte) (*Session, error) {
	redisClient := ctx.MustGet("redisConnection").(*redis.Client)
	redisPrefix := ctx.MustGet("redisPrefix").(string)

	sessionID, _ := uuid.NewV4()
	base64ID := base64.StdEncoding.EncodeToString(sessionID.Bytes())

	redisSession := redisSession{model, modelID}
	encodedRedisSession, err := json.Marshal(redisSession)
	if err != nil {
		return nil, errors.Wrap(err, "Could not serialize session")
	}

	redisClient.Set(redisPrefix + "session:" + base64ID, encodedRedisSession, expiryDuration)
	maxAge := int(math.Round(expiryDuration.Seconds()))
	domain := strings.Split(ctx.Request.Host, ":")[0]
	ctx.SetCookie("sessionID", base64ID, maxAge, "/", domain, false, true)

	session := Session{sessionID, base64ID, time.Now().Add(expiryDuration),  model, modelID}
	return &session, nil
}