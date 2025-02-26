package middleware

import (
    "strconv"
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/Kshitijknk07/TitanGate/backend/internal/metrics"
)

func MetricsMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        start := time.Now()
        err := c.Next()
        duration := time.Since(start)

        status := strconv.Itoa(c.Response().StatusCode())
        metrics.RequestCounter.WithLabelValues(c.Path(), c.Method(), status).Inc()

        return err
    }
}