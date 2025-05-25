package loadbalancer

import "sync"

type LoadBalancer struct {
	backends  []*Backend
	algorithm Algorithm
	mu        sync.RWMutex
}

func NewLoadBalancer(backends []*Backend, algorithm Algorithm) *LoadBalancer {
	return &LoadBalancer{
		backends:  backends,
		algorithm: algorithm,
	}
}

func (lb *LoadBalancer) NextBackend() *Backend {
	lb.mu.RLock()
	defer lb.mu.RUnlock()
	return lb.algorithm.NextBackend()
}
