package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GetUsers handles GET /api/v1/users
func GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "List of users",
		"data":    []string{},
	})
}

// GetUser handles GET /api/v1/users/:id
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"message": "User details",
		"id":      id,
	})
}

// CreateUser handles POST /api/v1/users
func CreateUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

// UpdateUser handles PUT /api/v1/users/:id
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"id":      id,
	})
}

// DeleteUser handles DELETE /api/v1/users/:id
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
		"id":      id,
	})
}
