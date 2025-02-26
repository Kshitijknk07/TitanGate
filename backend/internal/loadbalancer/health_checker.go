package loadbalancer

import (
	"net/http"
	"time"
)

type HealthChecker struct {
	client  *http.Client
	lb      *LoadBalancer
	timeout time.Duration
}

func NewHealthChecker(lb *LoadBalancer, timeout time.Duration) *HealthChecker {
	return &HealthChecker{
		client:  &http.Client{Timeout: timeout},
		lb:      lb,
		timeout: timeout,
	}
}

func (hc *HealthChecker) Start() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			hc.checkHealth()
		}
	}()
}

func (hc *HealthChecker) checkHealth() {
	for _, backend := range hc.lb.backends {
		resp, err := hc.client.Get(backend.URL + "/health")
		backend.Active = err == nil && resp != nil && resp.StatusCode == http.StatusOK
		if resp != nil {
			resp.Body.Close()
		}
	}
}