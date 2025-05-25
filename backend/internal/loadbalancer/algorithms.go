package loadbalancer

import (
	"sync"
)

type Algorithm interface {
	NextBackend() *Backend
}

type RoundRobin struct {
	backends []*Backend
	current  uint64
	mu       sync.Mutex
}

func NewRoundRobin(backends []*Backend) *RoundRobin {
	return &RoundRobin{
		backends: backends,
	}
}

func (rr *RoundRobin) NextBackend() *Backend {
	rr.mu.Lock()
	defer rr.mu.Unlock()
	if len(rr.backends) == 0 {
		return nil
	}
	idx := rr.current % uint64(len(rr.backends))
	rr.current++
	return rr.backends[idx]
}
