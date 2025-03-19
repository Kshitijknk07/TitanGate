package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/Kshitijknk07/TitanGate/backend/internal/config"
	"github.com/Kshitijknk07/TitanGate/backend/internal/handlers"
	"github.com/Kshitijknk07/TitanGate/backend/internal/loadbalancer"
	"github.com/Kshitijknk07/TitanGate/backend/internal/metrics"
	"github.com/Kshitijknk07/TitanGate/backend/internal/middleware"
	"github.com/Kshitijknk07/TitanGate/backend/internal/routes"
	"github.com/Kshitijknk07/TitanGate/backend/internal/services"
)

type Backend struct {
	URL    string
	Weight int
	Active bool
}

type WeightedRoundRobin struct {
	weights       []int
	currentIndex  int
	currentWeight int
	maxWeight     int
	gcdWeight     int
	mu            sync.Mutex
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func getGCD(weights []int) int {
	g := weights[0]
	for _, w := range weights[1:] {
		g = gcd(g, w)
	}
	return g
}

func max(weights []int) int {
	m := weights[0]
	for _, w := range weights[1:] {
		if w > m {
			m = w
		}
	}
	return m
}

func NewWeightedRoundRobin(weights []int) *WeightedRoundRobin {
	return &WeightedRoundRobin{
		weights:       weights,
		currentIndex:  -1,
		currentWeight: 0,
		maxWeight:     max(weights),
		gcdWeight:     getGCD(weights),
	}
}

func (w *WeightedRoundRobin) Next() int {
	w.mu.Lock()
	defer w.mu.Unlock()

	for {
		w.currentIndex = (w.currentIndex + 1) % len(w.weights)
		if w.currentIndex == 0 {
			w.currentWeight -= w.gcdWeight
			if w.currentWeight <= 0 {
				w.currentWeight = w.maxWeight
				if w.currentWeight == 0 {
					return -1
				}
			}
		}
		if w.weights[w.currentIndex] >= w.currentWeight {
			return w.currentIndex
		}
	}
}

func main() {
	config.LoadEnv()
	services.InitRedis()
	defer services.CloseRedis()
	services.InitCache()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(middleware.LoggerMiddleware())
	app.Use(middleware.MetricsMiddleware())

	authConfig := middleware.NewAuthConfig()
	app.Use(middleware.AuthMiddleware(authConfig))

	app.Use(middleware.RateLimit)

	backends := []*loadbalancer.Backend{
		{URL: "http://localhost:8081", Weight: 5, Active: true},
		{URL: "http://localhost:8082", Weight: 3, Active: true},
		{URL: "http://localhost:8083", Weight: 2, Active: true},
	}
	algorithm := loadbalancer.NewRoundRobin(backends)
	lb := loadbalancer.NewLoadBalancer(backends, algorithm)

	healthChecker := loadbalancer.NewHealthChecker(lb, 10*time.Second)
	healthChecker.Start()

	app.Use(middleware.CircuitBreakerMiddleware(5, 30*time.Second))

	routes.SetupRoutes(app)

	app.Use("/api/*", middleware.LoadBalancerMiddleware(lb))

	app.Get("/health", handlers.HealthCheck)
	app.Get("/metrics", metrics.Handler)

	app.Static("/", "./static")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Fiber server error: %v", err)
		}
	}()
	log.Printf("TitanGate API Gateway running on :%s", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gracefully...")
	_ = app.Shutdown()
}
