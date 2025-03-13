package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/Kshitijknk07/TitanGate/backend/internal/services"
)

func RateLimit(c *fiber.Ctx) error {
	ip := c.IP()
	key := fmt.Sprintf("ratelimit:%s", ip)

	count, err := services.RedisClient.Incr(services.Ctx, key).Result()
	if err != nil {
		return c.Status(500).SendString("Rate limit error")
	}

	if count == 1 {
		services.RedisClient.Expire(services.Ctx, key, time.Hour)
	}

	if count > 100 {
		return c.Status(429).SendString("Too many requests")
	}

	return c.Next()
}
