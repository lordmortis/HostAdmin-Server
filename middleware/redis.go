package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/lordmortis/HostAdmin-Server/config"
	"strconv"
)

func Redis(config config.RedisConfig) (gin.HandlerFunc, error) {
	options := redis.Options{
		Addr: config.Hostname + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB: config.Database,

	}

	client := redis.NewClient(&options)
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	handler := func(ctx *gin.Context) {
		ctx.Set("redisConnection", client)
		if len(config.Namespace) > 0  {
			ctx.Set("redisPrefix", config.Namespace + ":")
		} else {
			ctx.Set("redisPrefix", "")
		}
	}
	return handler, nil
}