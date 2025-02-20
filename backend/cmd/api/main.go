package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Kshitijknk07/TitanGate/backend/internal/config"
	"github.com/Kshitijknk07/TitanGate/backend/internal/middleware"
	"github.com/Kshitijknk07/TitanGate/backend/internal/routes"
	"github.com/Kshitijknk07/TitanGate/backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	services.InitRedis()
	services.InitCache()

	port := os.Getenv("PORT")
	appName := os.Getenv("APP_NAME")

	app := fiber.New()

	app.Use(middleware.RateLimit)
	app.Use(middleware.CacheMiddleware)

	router.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("%s Backend is Running!", appName))
	})

	log.Fatal(app.Listen(":" + port))
}
