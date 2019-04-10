package middleware

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"time"
)

type Session struct {
	ID uuid.UUID
	Expiry time.Time
	Model string
	ModelID []byte
}

type redisSession struct {
	Model string
	ModelID []byte
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func AuthCreateSession(redisClient *redis.Client, redisPrefix string, model string, modelID []byte) (*Session, error) {
	sessionID := uuid.NewV4()
	base64ID := base64.StdEncoding.EncodeToString(sessionID.Bytes())
	expiryDuration := time.Minute

	redisSession := redisSession{model, modelID}
	encodedRedisSession, err := json.Marshal(redisSession)

	if err != nil {
		return nil, errors.Wrap(err, "Could not serialize session")
	}

	redisClient.Set(redisPrefix + "session:" + base64ID, encodedRedisSession, expiryDuration)

	session := Session{sessionID, time.Now().Add(expiryDuration),  model, modelID}
	return &session, nil
}