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
	
	
	versionConfig := middleware.NewVersionConfig()
	app.Use(middleware.APIVersionMiddleware(versionConfig))
	
	
	app.Use(middleware.RateLimit)
	app.Use(middleware.CacheMiddleware)
	
	
	vRouter := routes.NewVersionedRouter(app)
	routes.SetupRoutes(app, vRouter)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("%s Backend is Running!", appName))
	})

	log.Fatal(app.Listen(":" + port))
}
