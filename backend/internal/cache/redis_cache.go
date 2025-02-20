package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisCache struct {
	Client *redis.Client
}


func NewRedisCache() *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0,
	})
	return &RedisCache{Client: client}
}

func (r *RedisCache) Set(key string, value string, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}


func (r *RedisCache) Delete(key string) error {
	return r.Client.Del(ctx, key).Err()
}
