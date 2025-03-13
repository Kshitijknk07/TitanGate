package loadbalancer

import (
	"sync"
	"sync/atomic"
	"time"
)

type Algorithm string

const (
	RoundRobin  Algorithm = "round-robin"
	LeastConn   Algorithm = "least-connections"
	WeightedRR  Algorithm = "weighted-round-robin"
)

type Backend struct {
	URL           string
	Weight        int
	Active        bool
	Health        float64
	LastCheck     time.Time
	ActiveConns   int32
	ResponseTime  time.Duration
	FailureCount  int32
	SuccessCount  int32
}

type Algorithm interface {
	NextBackend() *Backend
}

type LoadBalancer struct {
	backends  []Backend
	algorithm Algorithm
	mu        sync.RWMutex
}

func NewLoadBalancer(backends []Backend, algorithm Algorithm) *LoadBalancer {
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

func (lb *LoadBalancer) getLeastConnectedBackend(backends []*Backend) *Backend {
	var selected *Backend
	minConn := int32(^uint32(0) >> 1)

	for _, b := range backends {
		if conns := atomic.LoadInt32(&b.ActiveConns); conns < minConn {
			minConn = conns
			selected = b
		}
	}
	return selected
}

func (lb *LoadBalancer) getWeightedRoundRobinBackend(backends []*Backend) *Backend {
	totalWeight := 0
	for _, b := range backends {
		totalWeight += b.Weight
	}
	
	if totalWeight == 0 {
		return lb.getRoundRobinBackend(backends)
	}

	current := atomic.AddUint64(&lb.current, 1)
	pick := int(current % uint64(totalWeight))
	
	for _, b := range backends {
		pick -= b.Weight
		if pick < 0 {
			return b
		}
	}
	
	return backends[0]
}

func (lb *LoadBalancer) getRoundRobinBackend(backends []*Backend) *Backend {
	idx := atomic.AddUint64(&lb.current, 1) % uint64(len(backends))
	return backends[idx]
}

func (lb *LoadBalancer) getActiveBackends() []*Backend {
	active := make([]*Backend, 0)
	for _, b := range lb.backends {
		if b.Active && b.Health > 0.5 && atomic.LoadInt32(&b.FailureCount) < 3 {
			active = append(active, b)
		}
	}
	return active
}

func (b *Backend) IncrementConnections() {
	atomic.AddInt32(&b.ActiveConns, 1)
}

func (b *Backend) DecrementConnections() {
	atomic.AddInt32(&b.ActiveConns, -1)
}

func (b *Backend) RecordResponse(duration time.Duration, success bool) {
	b.ResponseTime = duration
	if success {
		atomic.AddInt32(&b.SuccessCount, 1)
		atomic.StoreInt32(&b.FailureCount, 0)
	} else {
		atomic.AddInt32(&b.FailureCount, 1)
	}
}