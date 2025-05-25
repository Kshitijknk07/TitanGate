package loadbalancer

import (
	"sync/atomic"
	"time"
)

type Backend struct {
	URL          string
	Weight       int
	Active       bool
	Health       float64
	LastCheck    time.Time
	ActiveConns  int32
	ResponseTime time.Duration
	FailureCount int32
	SuccessCount int32
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
