package routes

import (
	"os"

	"github.com/Kshitijknk07/TitanGate/backend/internal/handlers"
	"github.com/Kshitijknk07/TitanGate/backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func SetupRoutes(app *fiber.App) {
	transformConfig := middleware.NewTransformConfig()
	transformConfig.InputFormat = "json"
	transformConfig.OutputFormat = "json"
	transformConfig.HeaderTransform = map[string]string{
		"X-API-Key": "Authorization",
	}
	transformConfig.QueryTransform = map[string]string{
		"api_key": "token",
	}

	if os.Getenv("ENV") == "development" {
		transformConfig.DebugMode = true
		transformConfig.MaxBodySize = 5 * 1024 * 1024
	} else {
		transformConfig.MaxBodySize = 1024 * 1024
		transformConfig.CacheEnabled = true
	}

	app.Use(middleware.LoggerMiddleware())
	app.Use(middleware.MetricsMiddleware())
	app.Use(middleware.RateLimit)
	app.Use(middleware.TransformerMiddleware(transformConfig))

	api := app.Group("/api")
	api.Use(middleware.AuthMiddleware(middleware.NewAuthConfig()))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	promHandler := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
	app.Get("/metrics", func(c *fiber.Ctx) error {
		promHandler(c.Context())
		return nil
	})

	v1 := api.Group("/v1")
	{
		users := v1.Group("/users")
		users.Get("/", handlers.GetUsers)
		users.Get("/:id", handlers.GetUser)
		users.Post("/", handlers.CreateUser)
		users.Put("/:id", handlers.UpdateUser)
		users.Delete("/:id", handlers.DeleteUser)
	}
}
