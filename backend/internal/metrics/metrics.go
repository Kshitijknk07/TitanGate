package metrics

import (
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "titangate_requests_total",
			Help: "Total requests processed",
		},
		[]string{"path", "method", "status"},
	)

	RequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "titangate_request_duration_seconds",
			Help:    "Request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)

	CacheHits = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "titangate_cache_hits_total",
			Help: "Total cache hits",
		},
	)

	CacheMisses = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "titangate_cache_misses_total",
			Help: "Total cache misses",
		},
	)

	RateLimitHits = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "titangate_ratelimit_hits_total",
			Help: "Total rate limit hits",
		},
		[]string{"ip"},
	)

	BackendHealth = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "titangate_backend_health",
			Help: "Backend health status (1=healthy, 0=unhealthy)",
		},
		[]string{"backend"},
	)

	BackendLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "titangate_backend_latency_seconds",
			Help:    "Backend response latency in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"backend"},
	)

	ActiveConnections = promauto.NewGauge(
    RequestCounter = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "titangate_requests_total",
            Help: "Total requests processed",
        },
        []string{"path", "method", "status"},
    )

    RequestDuration = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "titangate_request_duration_seconds",
            Help:    "Request duration in seconds",
            Buckets: prometheus.DefBuckets,
        },
        []string{"path", "method"},
    )

    CacheHits = promauto.NewCounter(
        prometheus.CounterOpts{
            Name: "titangate_cache_hits_total",
            Help: "Total cache hits",
        },
    )

    CacheMisses = promauto.NewCounter(
        prometheus.CounterOpts{
            Name: "titangate_cache_misses_total",
            Help: "Total cache misses",
        },
    )

    RateLimitHits = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "titangate_ratelimit_hits_total",
            Help: "Total rate limit hits",
        },
        []string{"ip"},
    )

    BackendHealth = promauto.NewGaugeVec(
        prometheus.GaugeOpts{
            Name: "titangate_backend_health",
            Help: "Backend health status (1=healthy, 0=unhealthy)",
        },
        []string{"backend"},
    )

    BackendLatency = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "titangate_backend_latency_seconds",
            Help:    "Backend response latency in seconds",
            Buckets: prometheus.DefBuckets,
        },
        []string{"backend"},
    )

    ActiveConnections = promauto.NewGauge(
        prometheus.GaugeOpts{
            Name: "titangate_active_connections",
            Help: "Number of active connections",
        },
    )

    ErrorCounter = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "titangate_errors_total",
            Help: "Total number of errors",
        },
        []string{"type"},
    )
)