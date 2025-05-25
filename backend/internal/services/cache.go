package services

import (
	"log"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
)

var cache *lru.Cache[string, []byte]

func InitCache() {
	var err error
	cache, err = lru.New[string, []byte](1000)
	if err != nil {
		log.Fatal("Failed to initialize cache:", err)
	}
}

func SetCache(key string, value []byte, expiration time.Duration) {
	cache.Add(key, value)
}

func GetCache(key string) ([]byte, bool) {
	if value, ok := cache.Get(key); ok {
		return value, true
	}
	return nil, false
}
