package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type VersionConfig struct {
	DefaultVersion string
	HeaderName    string
}

func NewVersionConfig() VersionConfig {
	return VersionConfig{
		DefaultVersion: "v1",
		HeaderName:    "Accept-Version",
	}
}

func APIVersionMiddleware(config VersionConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var version string

		version = c.Get(config.HeaderName)

		if version == "" {
			path := c.Path()
			if strings.HasPrefix(path, "/api/") {
				parts := strings.Split(path, "/")
				if len(parts) > 2 && strings.HasPrefix(parts[2], "v") {
					version = parts[2]
				}
			}
		}

		if version == "" {
			version = config.DefaultVersion
		}

		c.Locals("version", version)
		c.Set("X-API-Version", version)

		return c.Next()
	}
}