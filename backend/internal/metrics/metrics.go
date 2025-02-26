package metrics

import (
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
)