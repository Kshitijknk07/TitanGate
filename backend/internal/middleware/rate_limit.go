package middleware

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Kshitijknk07/TitanGate/backend/internal/services"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)


var routeLimits = map[string]int{
	"/login":  10,  // Max 10 requests per hour
	"/signup": 5,   // Max 5 requests per hour
	"/api":    100, // General API limit
	"*":       50,  // Default limit
}

func RateLimit(c *fiber.Ctx) error {
	clientIP := c.IP()
	route := c.Path()
	limit, exists := routeLimits[route]
	if !exists {
		limit = routeLimits["*"]
	}

	limitKey := fmt.Sprintf("rate_limit:%s:%s", route, clientIP)
	now := time.Now().Unix()
	window := int64(3600)

	
	pipe := services.RedisClient.TxPipeline()

	
	pipe.ZRemRangeByScore(services.Ctx, limitKey, "0", fmt.Sprintf("%d", now-window))

	
	countCmd := pipe.ZCard(services.Ctx, limitKey)

	
	pipe.ZAdd(services.Ctx, limitKey, &redis.Z{
		Score:  float64(now),
		Member: now,
	})

	
	pipe.Expire(services.Ctx, limitKey, time.Duration(window)*time.Second)

	
	_, err := pipe.Exec(services.Ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	
	count, err := countCmd.Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}


	if count > int64(limit) {
		return c.Status(fiber.StatusTooManyRequests).SendString("Rate limit exceeded")
	}

	return c.Next()
}
