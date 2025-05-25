package routes

import (
	"github.com/Kshitijknk07/TitanGate/backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, vRouter *VersionedRouter) {

	v1 := vRouter.Group("/v1")
	v1.Get("/user", handlers.GetUserHandler)
	v1.Get("/health", handlers.HealthCheck)

	v2 := vRouter.Group("/v2")
	v2.Get("/user", handlers.GetUserHandler)
	v2.Get("/health", handlers.HealthCheck)
}
