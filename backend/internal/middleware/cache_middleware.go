// internal/middleware/cache_middleware.go
package middleware

import (
	"context"
	"net/http"

	"github.com/golang/groupcache"
)

const cacheSize = 64 << 20 // 64MB Cache

var TitanCache = groupcache.NewGroup("TitanGateCache", cacheSize, groupcache.GetterFunc(
	func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
		// This function is called when cache miss occurs
		// Implement how to fetch and store the data if not in cache
		return nil
	},
))

// CacheMiddleware checks if response exists in cache before processing request
func CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.RequestURI // Use request URI as cache key
		if data, err := GetFromCache(key); err == nil {
			w.Write(data) // Serve cached response
			return
		}

		// Capture response and store in cache
		rec := &responseRecorder{ResponseWriter: w}
		next.ServeHTTP(rec, r)

		SetCache(key, rec.body) // Store response in cache
	})
}

// GetFromCache retrieves data from cache
func GetFromCache(key string) ([]byte, error) {
	var data []byte
	err := TitanCache.Get(context.TODO(), key, groupcache.AllocatingByteSliceSink(&data))
	if err != nil {
		return nil, err
	}
	return data, nil
}

// SetCache stores data in cache manually
func SetCache(key string, value []byte) {
	// Populate the cache using a GetterFunc
	TitanCache.Get(context.TODO(), key, groupcache.AllocatingByteSliceSink(&value))
}

// responseRecorder captures response for caching
type responseRecorder struct {
	http.ResponseWriter
	body []byte
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body = append(r.body, b...)
	return r.ResponseWriter.Write(b)
}
