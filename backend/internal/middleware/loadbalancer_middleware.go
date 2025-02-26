package middleware

import (
	"fmt"
	"github.com/Kshitijknk07/TitanGate/backend/internal/loadbalancer"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func LoadBalancerMiddleware(lb *loadbalancer.LoadBalancer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		backend := lb.NextBackend()
		if backend == nil {
			return c.Status(503).SendString("No available backends")
		}

		targetURL := backend.URL + c.Path()
		if err := proxy.Do(c, targetURL); err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Error: %v", err))
		}

		return nil
	}
}