package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "titangate_http_requests_total",
			Help: "Total number of HTTP requests by endpoint and method",
		},
		[]string{"endpoint", "method", "status"},
	)

	ResponseTime = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "titangate_http_response_time_seconds",
			Help:    "Response time distribution by endpoint",
			Buckets: []float64{0.1, 0.5, 1, 2, 5},
		},
		[]string{"endpoint", "method"},
	)

	ActiveRequests = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "titangate_active_requests",
			Help: "Number of requests currently being processed",
		},
	)

	CacheHits = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "titangate_cache_hits_total",
			Help: "Total number of cache hits",
		},
	)

	CacheMisses = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "titangate_cache_misses_total",
			Help: "Total number of cache misses",
		},
	)
)