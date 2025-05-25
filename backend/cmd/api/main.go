package loadbalancer

import (
	"sync"
)

type Backend struct {
	URL    string
	Weight int
	Active bool
}

type WeightedRoundRobin struct {
	weights       []int
	currentIndex  int
	currentWeight int
	maxWeight     int
	gcdWeight     int
	mu            sync.Mutex
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func getGCD(weights []int) int {
	g := weights[0]
	for _, w := range weights[1:] {
		g = gcd(g, w)
	}
	return g
}

func max(weights []int) int {
	m := weights[0]
	for _, w := range weights[1:] {
		if w > m {
			m = w
		}
	}
	return m
}

func NewWeightedRoundRobin(weights []int) *WeightedRoundRobin {
	return &WeightedRoundRobin{
		weights:       weights,
		currentIndex:  -1,
		currentWeight: 0,
		maxWeight:     max(weights),
		gcdWeight:     getGCD(weights),
	}
}

func (w *WeightedRoundRobin) Next() int {
	w.mu.Lock()
	defer w.mu.Unlock()

	for {
		w.currentIndex = (w.currentIndex + 1) % len(w.weights)
		if w.currentIndex == 0 {
			w.currentWeight -= w.gcdWeight
			if w.currentWeight <= 0 {
				w.currentWeight = w.maxWeight
				if w.currentWeight == 0 {
					return -1
				}
			}
		}
		if w.weights[w.currentIndex] >= w.currentWeight {
			return w.currentIndex
		}
	}
}
