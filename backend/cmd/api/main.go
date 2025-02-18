package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Kshitijknk07/TitanGate/backend/internal/config"
	"github.com/Kshitijknk07/TitanGate/backend/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	port := os.Getenv("PORT")
	appName := os.Getenv("APP_NAME")

	app := fiber.New()

	// Set up routes
	routes.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("%s Backend is Running!", appName))
	})

	log.Fatal(app.Listen(":" + port))
}
