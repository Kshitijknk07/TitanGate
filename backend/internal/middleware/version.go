package middleware

import "github.com/gofiber/fiber/v2"

func VersionMiddleware(version string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("API-Version", version)
		return c.Next()
	}
}
