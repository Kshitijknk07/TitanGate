// internal/middleware/rate_limit.go
package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Kshitijknk07/TitanGate/backend/internal/services"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// RateLimit is the rate limiting middleware
func RateLimit(c *fiber.Ctx) error {
	clientIP := c.IP()
	limitKey := fmt.Sprintf("rate_limit:%s", clientIP)
	limit := 100            // Max requests per time window (e.g., 100 requests)
	window := 1 * time.Hour // Time window for rate limiting (1 hour)

	// Check current count from Redis
	count, err := services.RedisClient.Get(services.Ctx, limitKey).Result()
	if err != nil && err != redis.Nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if count != "" && atoi(count) >= limit {
		return c.Status(fiber.StatusTooManyRequests).SendString("Rate limit exceeded")
	}

	// Set/Increment the request counter
	err = services.RedisClient.Incr(services.Ctx, limitKey).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Set the expiration time for the rate limit key
	err = services.RedisClient.Expire(services.Ctx, limitKey, window).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Next()
}

// Helper function to convert string to int
func atoi(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}
