package routes

import (
	"github.com/Kshitijknk07/TitanGate/backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, vRouter *VersionedRouter) {
	
	v1 := vRouter.Group("v1")
	setupV1Routes(v1)

	
	v2 := vRouter.Group("v2")
	setupV2Routes(v2)

	
	app.Get("/api", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"versions": vRouter.GetSupportedVersions(),
			"latest":   vRouter.GetLatestVersion(),
			"current": c.Locals("version"),
		})
	})
}

func setupV1Routes(router fiber.Router) {
	router.Get("/health", handlers.HealthCheckV1)
	router.Get("/user", handlers.GetUserHandlerV1)
}

func setupV2Routes(router fiber.Router) {
	router.Get("/health", handlers.HealthCheckV2)
	router.Get("/user", handlers.GetUserHandlerV2)
}