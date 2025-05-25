package loadbalancer

import (
	"net/http"
	"sync"
	"time"
)

type HealthChecker struct {
	lb       *LoadBalancer
	interval time.Duration
	client   *http.Client
	mu       sync.RWMutex
}

func NewHealthChecker(lb *LoadBalancer, interval time.Duration) *HealthChecker {
	return &HealthChecker{
		lb:       lb,
		interval: interval,
		client:   &http.Client{Timeout: 5 * time.Second},
	}
}

func (hc *HealthChecker) Start() {
	go hc.checkHealth()
}

func (hc *HealthChecker) checkHealth() {
	ticker := time.NewTicker(hc.interval)
	defer ticker.Stop()

	for range ticker.C {
		hc.mu.RLock()
		backends := hc.lb.backends
		hc.mu.RUnlock()

		for i := range backends {
			go func(i int) {
				backend := backends[i]
				resp, err := hc.client.Get(backend.URL + "/health")
				if err != nil {
					backend.Active = false
					return
				}
				defer resp.Body.Close()

				backend.Active = resp.StatusCode == http.StatusOK
			}(i)
		}
	}
}
