
package services

import (
	"context"
	"log"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx = context.Background()
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB: 0,
	})
	if err := RedisClient.Ping(Ctx).Err(); err != nil {
		log.Fatal("Redis connection failed:", err)
	}
}
