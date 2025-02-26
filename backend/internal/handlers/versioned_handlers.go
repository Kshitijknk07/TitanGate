package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// HealthCheckV2 is an example of a v2 endpoint with enhanced response
func HealthCheckV2(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":    "OK",
		"message":   "Backend is running smoothly!",
		"version":   "2.0",
		"timestamp": time.Now(),
		"metrics": fiber.Map{
			"uptime":    "24h",
			"requests":  1000,
			"memory":    "512MB",
		},
	})
}

// GetUserHandlerV2 is an example of a v2 user endpoint
func GetUserHandlerV2(c *fiber.Ctx) error {
	userID := c.Query("id")
	cacheKey := "user:v2:" + userID

	
	userData := fiber.Map{
		"id":           userID,
		"name":        "John Doe",
		"email":       "john@example.com",
		"created_at":  time.Now().Add(-24 * time.Hour),
		"updated_at":  time.Now(),
		"preferences": fiber.Map{
			"theme":     "dark",
			"language": "en",
		},
	}

	return c.JSON(userData)
}