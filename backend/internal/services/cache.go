package services

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)


var RedisCacheClient *redis.Client


func InitCache() {
	RedisCacheClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", 
		Password: "",               
		DB:       0,                
	})
}


func GetCache(key string) ([]byte, error) {
	val, err := RedisCacheClient.Get(Ctx, key).Bytes()
	if err == redis.Nil {
		return nil, fmt.Errorf("cache miss")
	}
	return val, err
}


func SetCache(key string, value []byte, expiration time.Duration) error {
	return RedisCacheClient.Set(Ctx, key, value, expiration).Err()
}
