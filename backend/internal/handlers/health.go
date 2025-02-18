package handlers

import "github.com/gofiber/fiber/v2"

// HealthCheck returns the status of the API
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Backend is running smoothly!",
	})
}
