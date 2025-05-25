package cache

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru/v2"
)

type LRUCache struct {
	cache *lru.Cache[string, string]
}

func NewLRUCache(size int) *LRUCache {
	l, err := lru.New[string, string](size)
	if err != nil {
		fmt.Println("Failed to create LRU Cache:", err)
		return nil
	}
	return &LRUCache{cache: l}
}

func (l *LRUCache) Set(key string, value string) {
	l.cache.Add(key, value)
}

func (l *LRUCache) Get(key string) (string, bool) {
	value, ok := l.cache.Get(key)
	return value, ok
}

func (l *LRUCache) Remove(key string) {
	l.cache.Remove(key)
}
