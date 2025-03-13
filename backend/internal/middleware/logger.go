package middleware

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/Kshitijknk07/TitanGate/backend/internal/metrics"
)

type LogEntry struct {
	Timestamp   time.Time `json:"timestamp"`
	Method      string    `json:"method"`
	Path        string    `json:"path"`
	Status      int       `json:"status"`
	Duration    float64   `json:"duration"`
	IP          string    `json:"ip"`
	UserAgent   string    `json:"user_agent"`
	RequestID   string    `json:"request_id"`
	RequestBody string    `json:"request_body,omitempty"`
}

func LoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		requestID := generateRequestID()

		body := c.Body()
		if len(body) > 0 {
			c.Request().SetBody(body)
		}

		err := c.Next()
		if err != nil {
			return err
		}

		duration := time.Since(start).Seconds()
		metrics.RequestDuration.WithLabelValues(c.Method(), c.Path()).Observe(duration)
		metrics.RequestCounter.WithLabelValues(c.Method(), c.Path()).Inc()

		logEntry := LogEntry{
			Timestamp:   time.Now(),
			Method:      c.Method(),
			Path:        c.Path(),
			Status:      c.Response().StatusCode(),
			Duration:    duration,
			IP:          c.IP(),
			UserAgent:   c.Get("User-Agent"),
			RequestID:   requestID,
		}

		if c.Method() != "GET" && len(body) > 0 {
			logEntry.RequestBody = string(body)
		}

		logJSON, _ := json.Marshal(logEntry)
		os.Stdout.Write(append(logJSON, '\n'))

		return nil
	}
}

func generateRequestID() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
} 