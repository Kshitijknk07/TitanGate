package services

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCacheClient is the Redis client instance for caching
var RedisCacheClient *redis.Client

// Use the existing Ctx from the other file

// InitCache initializes Redis for caching
func InitCache() {
	RedisCacheClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Update this if needed
		Password: "",               // Set password if required
		DB:       0,                // Default DB
	})
}

// GetCache retrieves a value from Redis cache
func GetCache(key string) ([]byte, error) {
	val, err := RedisCacheClient.Get(Ctx, key).Bytes()
	if err == redis.Nil {
		return nil, fmt.Errorf("cache miss")
	}
	return val, err
}

// SetCache stores a value in Redis cache with expiration
func SetCache(key string, value []byte, expiration time.Duration) error {
	return RedisCacheClient.Set(Ctx, key, value, expiration).Err()
}
