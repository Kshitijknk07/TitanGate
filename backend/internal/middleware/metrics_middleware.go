package middleware

import (
	"time"

	"github.com/Kshitijknk07/TitanGate/backend/internal/metrics"
	"github.com/gofiber/fiber/v2"
)

func MetricsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start).Seconds()

		metrics.RequestDuration.WithLabelValues(c.Method(), c.Path()).Observe(duration)
		metrics.RequestCounter.WithLabelValues(c.Method(), c.Path()).Inc()

		return err
	}
}
