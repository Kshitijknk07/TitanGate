package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/Kshitijknk07/TitanGate/backend/internal/cache"
	"github.com/Kshitijknk07/TitanGate/backend/internal/loadbalancer"
	"github.com/Kshitijknk07/TitanGate/backend/internal/metrics"
	"github.com/Kshitijknk07/TitanGate/backend/internal/middleware"
	"github.com/Kshitijknk07/TitanGate/backend/internal/services"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "TitanGate API Gateway",
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	backends := []loadbalancer.Backend{
		{URL: "http://localhost:8081", Weight: 1, Active: true},
		{URL: "http://localhost:8082", Weight: 2, Active: true},
		{URL: "http://localhost:8083", Weight: 1, Active: true},
	}

	weights := make([]int, len(backends))
	for i, b := range backends {
		weights[i] = b.Weight
	}

	algorithm := loadbalancer.NewWeightedRoundRobin(weights)
	lb := loadbalancer.NewLoadBalancer(backends, algorithm)
	healthChecker := loadbalancer.NewHealthChecker(lb, 10*time.Second)
	healthChecker.Start()

	app.Use(func(c *fiber.Ctx) error {
		nextBackend := lb.NextBackend()
		if !backends[nextBackend].Active {
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error": "No healthy backends available",
			})
		}

		proxy := fiber.New()
		proxy.All("/*", func(c *fiber.Ctx) error {
			url := backends[nextBackend].URL + c.Path()
			return c.Redirect(url, fiber.StatusTemporaryRedirect)
		})

		return proxy.Handler()(c)
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
			"time":   time.Now(),
		})
	})

	app.Get("/metrics", metrics.Handler)

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "TitanGate API Gateway is running",
		})
	})

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatal(err)
	}
}
