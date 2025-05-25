package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckV2(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":    "OK",
		"message":   "Backend is running smoothly!",
		"version":   "2.0",
		"timestamp": time.Now(),
		"metrics": fiber.Map{
			"uptime":   "24h",
			"requests": 1000,
			"memory":   "512MB",
		},
	})
}

func GetUserHandlerV2(c *fiber.Ctx) error {
	userID := c.Query("id")

	userData := fiber.Map{
		"id":         userID,
		"name":       "John Doe",
		"email":      "john@example.com",
		"created_at": time.Now().Add(-24 * time.Hour),
		"updated_at": time.Now(),
		"preferences": fiber.Map{
			"theme":    "dark",
			"language": "en",
		},
	}

	return c.JSON(userData)
}
