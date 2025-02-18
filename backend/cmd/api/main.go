// cmd/api/main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Kshitijknk07/TitanGate/backend/internal/config"
	"github.com/Kshitijknk07/TitanGate/backend/internal/middleware"
	"github.com/Kshitijknk07/TitanGate/backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize Redis
	services.InitRedis()

	port := os.Getenv("PORT")
	appName := os.Getenv("APP_NAME")

	app := fiber.New()

	// Apply rate limiting middleware globally
	app.Use(middleware.RateLimit)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("%s Backend is Running!", appName))
	})

	log.Fatal(app.Listen(":" + port))
}
