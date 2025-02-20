package routes

import (
	"github.com/Kshitijknk07/TitanGate/backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(app *fiber.App) {
	app.Get("/health", handlers.HealthCheck)
}
