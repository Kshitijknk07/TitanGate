package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/Kshitijknk07/TitanGate/backend/internal/metrics"
)

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

type CircuitBreaker struct {
	state           CircuitState
	failures        int
	lastFailureTime time.Time
	threshold       int
	resetTimeout    time.Duration
	mu              sync.RWMutex
}

func NewCircuitBreaker(threshold int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:        StateClosed,
		threshold:    threshold,
		resetTimeout: resetTimeout,
	}
}

func (cb *CircuitBreaker) Execute(c *fiber.Ctx, handler fiber.Handler) error {
	if !cb.canExecute() {
		metrics.ErrorCounter.WithLabelValues("circuit_breaker").Inc()
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "Circuit breaker is open",
		})
	}

	err := handler(c)
	if err != nil {
		cb.recordFailure()
		return err
	}

	cb.recordSuccess()
	return nil
}

func (cb *CircuitBreaker) canExecute() bool {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	switch cb.state {
	case StateClosed:
		return true
	case StateOpen:
		if time.Since(cb.lastFailureTime) > cb.resetTimeout {
			cb.mu.RUnlock()
			cb.mu.Lock()
			cb.state = StateHalfOpen
			cb.mu.Unlock()
			return true
		}
		return false
	case StateHalfOpen:
		return true
	default:
		return false
	}
}

func (cb *CircuitBreaker) recordFailure() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.failures++
	cb.lastFailureTime = time.Now()

	if cb.failures >= cb.threshold {
		cb.state = StateOpen
		metrics.ErrorCounter.WithLabelValues("circuit_breaker_open").Inc()
	}
}

func (cb *CircuitBreaker) recordSuccess() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.state == StateHalfOpen {
		cb.state = StateClosed
		cb.failures = 0
	}
}

func CircuitBreakerMiddleware(threshold int, resetTimeout time.Duration) fiber.Handler {
	cb := NewCircuitBreaker(threshold, resetTimeout)

	return func(c *fiber.Ctx) error {
		return cb.Execute(c, func(c *fiber.Ctx) error {
			return c.Next()
		})
	}
} 