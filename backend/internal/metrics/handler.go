package metrics

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Handler(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(promhttp.Handler())(c)
} 