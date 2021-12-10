package config

import (
	"github.com/go-redis/redis/v8"
	"os"
)

func getRedisHost() string {
	if os.Getenv("MODE") == "production" {
		return "redis:6379"
	}
	return "localhost:6379"
}

var RedisClient = redis.NewClient(&redis.Options{
	Addr: getRedisHost(),
})
