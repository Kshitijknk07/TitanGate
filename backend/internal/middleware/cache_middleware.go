package middleware

import (
    "bytes"
    "github.com/gofiber/fiber/v2"
    "github.com/Kshitijknk07/TitanGate/backend/internal/cache"
    "github.com/Kshitijknk07/TitanGate/backend/internal/metrics"
    "time"
)

var redisCache = cache.NewRedisCache()

func CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cacheKey := "resp:" + r.URL.Path

		if cachedResponse, err := redisCache.Get(cacheKey); err == nil {
			metrics.CacheHits.Inc()
			w.Write([]byte(cachedResponse))
			return
		}
		metrics.CacheMisses.Inc()

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
func CacheMiddleware(c *fiber.Ctx) error {
    if c.Method() != "GET" {
        return c.Next()
    }
    key := "cache:" + c.Path()
    
    if cached, err := services.RedisClient.Get(services.Ctx, key).Result(); err == nil {
        metrics.CacheHits.Inc()
        return c.Type("json").Send([]byte(cached))
    }
    metrics.CacheMisses.Inc()
    return c.Next()
}
