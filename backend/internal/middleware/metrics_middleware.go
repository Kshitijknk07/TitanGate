package middleware

import (
    "strconv"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/Kshitijknk07/TitanGate/backend/internal/metrics"
)

func MetricsMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        startTime := time.Now()
        path := c.Route().Path // Using Route().Path instead of Path() for consistent metrics
        method := c.Method()

        metrics.ActiveRequests.Inc()

        // Execute the next handler
        err := c.Next()

        metrics.ActiveRequests.Dec()
        
        // Record metrics after response
        status := strconv.Itoa(c.Response().StatusCode())
        duration := time.Since(startTime).Seconds()

        metrics.RequestCounter.WithLabelValues(path, method, status).Inc()
        metrics.ResponseTime.WithLabelValues(path, method).Observe(duration)

        return err
    }
}