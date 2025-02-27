package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Kshitijknk07/TitanGate/backend/internal/config"
	"github.com/Kshitijknk07/TitanGate/backend/internal/middleware"
	"github.com/Kshitijknk07/TitanGate/backend/internal/routes"
	"github.com/Kshitijknk07/TitanGate/backend/internal/services"
	"github.com/Kshitijknk07/TitanGate/backend/internal/loadbalancer"  // Updated this line
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/adaptor/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config.LoadEnv()
	services.InitRedis()
	services.InitCache()
	
	port := os.Getenv("PORT")
	appName := os.Getenv("APP_NAME")
	
	app := fiber.New(fiber.Config{
        EnablePrintRoutes: true,
    })
    
    
    app.Use(middleware.MetricsMiddleware())
    
    
    app.Use(middleware.RateLimit)
    app.Use(middleware.CacheMiddleware)

    
    metricsHandler := adaptor.HTTPHandler(promhttp.Handler())
    app.Get("/metrics", func(c *fiber.Ctx) error {
        return metricsHandler(c)
    })

	versionConfig := middleware.NewVersionConfig()
	app.Use(middleware.APIVersionMiddleware(versionConfig))
	
	backends := []loadbalancer.Backend{
		{URL: "http://localhost:3001", Weight: 1, Active: true},
		{URL: "http://localhost:3002", Weight: 1, Active: true},
		{URL: "http://localhost:3003", Weight: 1, Active: true},
	}
	
	lb := loadbalancer.NewLoadBalancer(backends)
	healthChecker := loadbalancer.NewHealthChecker(lb, 5*time.Second)
	healthChecker.Start()
	app.Static("/", "./static")
	app.Use("/api", middleware.LoadBalancerMiddleware(lb))
	vRouter := routes.NewVersionedRouter(app)
	routes.SetupRoutes(app, vRouter)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("%s Backend is Running!", appName))
	})
	log.Fatal(app.Listen(":" + port))
}
