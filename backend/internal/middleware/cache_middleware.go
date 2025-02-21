package middleware

import (
	"bytes"
	"net/http"
	"time"

	"titangate/internal/cache"
)

var redisCache = cache.NewRedisCache()

func CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cacheKey := "resp:" + r.URL.Path

		if cachedResponse, err := redisCache.Get(cacheKey); err == nil {
			w.Write([]byte(cachedResponse))
			return
		}

		rec := &responseRecorder{ResponseWriter: w, body: new(bytes.Buffer)}
		next.ServeHTTP(rec, r)

		redisCache.Set(cacheKey, rec.body.String(), 10*time.Minute)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (r *responseRecorder) Write(p []byte) (int, error) {
	r.body.Write(p)
	return r.ResponseWriter.Write(p)
}
