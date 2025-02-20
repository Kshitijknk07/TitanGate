package handlers

import "github.com/gofiber/fiber/v2"


func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "OK",
		"message": "Backend is running smoothly!",
	})
}
