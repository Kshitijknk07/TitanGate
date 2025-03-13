package loadbalancer

import (
	"math/rand"
	"sync"
	"sync/atomic"
)

type RoundRobin struct {
	current uint64
	mu      sync.Mutex
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{}
}

func (rr *RoundRobin) NextBackend() int {
	next := atomic.AddUint64(&rr.current, 1)
	return int(next - 1)
}

type Random struct {
	mu sync.Mutex
}

func NewRandom() *Random {
	return &Random{}
}

func (r *Random) NextBackend() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return rand.Intn(100)
}

type WeightedRoundRobin struct {
	current uint64
	weights []int
	mu      sync.Mutex
}

func NewWeightedRoundRobin(weights []int) *WeightedRoundRobin {
	return &WeightedRoundRobin{
		weights: weights,
	}
}

func (wrr *WeightedRoundRobin) NextBackend() int {
	wrr.mu.Lock()
	defer wrr.mu.Unlock()

	total := 0
	for _, weight := range wrr.weights {
		total += weight
	}

	current := atomic.AddUint64(&wrr.current, 1)
	value := int(current) % total

	for i, weight := range wrr.weights {
		if value < weight {
			return i
		}
		value -= weight
	}

	return 0
} 