package middleware

import (
	"time"

	"github.com/Kshitijknk07/TitanGate/backend/internal/metrics"
	"github.com/gofiber/fiber/v2"
)

func MetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		path := c.Path()
		method := c.Method()

		// Track active requests
		metrics.ActiveRequests.Inc()
		defer metrics.ActiveRequests.Dec()

		// Process request
		err := c.Next()

		// Record metrics
		status := c.Response().StatusCode()
		duration := time.Since(start).Seconds()

		metrics.RequestCounter.WithLabelValues(path, method, string(status)).Inc()
		metrics.ResponseTime.WithLabelValues(path, method).Observe(duration)

		return err
	}
}