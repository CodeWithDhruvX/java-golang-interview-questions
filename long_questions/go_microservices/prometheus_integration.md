# Prometheus Integration with Go Microservices

## Overview

Prometheus is a leading monitoring and alerting toolkit that integrates seamlessly with Go applications. This guide covers Prometheus integration patterns specifically for Go microservices.

## Core Concepts

### 1. Prometheus Client Library for Go

The official Prometheus Go client library provides:
- Metrics collection and exposition
- HTTP handlers for metrics endpoints
- Custom metric types (Counter, Gauge, Histogram, Summary)

### 2. Metric Types

#### Counter
```go
// Counter for tracking total requests
var httpRequestsTotal = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total number of HTTP requests",
    },
    []string{"method", "endpoint", "status"},
)

func incrementCounter(method, endpoint, status string) {
    httpRequestsTotal.WithLabelValues(method, endpoint, status).Inc()
}
```

#### Gauge
```go
// Gauge for tracking current connections
var activeConnections = prometheus.NewGauge(
    prometheus.GaugeOpts{
        Name: "active_connections",
        Help: "Number of active database connections",
    },
)

func setActiveConnections(count float64) {
    activeConnections.Set(count)
}
```

#### Histogram
```go
// Histogram for request duration
var requestDuration = prometheus.NewHistogramVec(
    prometheus.HistogramOpts{
        Name: "http_request_duration_seconds",
        Help: "HTTP request duration in seconds",
        Buckets: prometheus.DefBuckets,
    },
    []string{"method", "endpoint"},
)

func observeDuration(method, endpoint string, duration time.Duration) {
    requestDuration.WithLabelValues(method, endpoint).Observe(duration.Seconds())
}
```

## Implementation Patterns

### 1. HTTP Middleware for Metrics

```go
func PrometheusMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Wrap response writer to capture status code
        wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
        
        next.ServeHTTP(wrapped, r)
        
        duration := time.Since(start)
        
        // Record metrics
        incrementCounter(r.Method, r.URL.Path, strconv.Itoa(wrapped.statusCode))
        observeDuration(r.Method, r.URL.Path, duration)
    })
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}
```

### 2. Database Connection Pool Metrics

```go
type DBMetrics struct {
    poolConnections    prometheus.Gauge
    poolIdleConnections prometheus.Gauge
    poolOpenConnections prometheus.Gauge
    queryDuration     *prometheus.HistogramVec
    queryErrors       *prometheus.CounterVec
}

func NewDBMetrics() *DBMetrics {
    return &DBMetrics{
        poolConnections: prometheus.NewGauge(
            prometheus.GaugeOpts{
                Name: "db_pool_connections",
                Help: "Total number of database connections in pool",
            },
        ),
        poolIdleConnections: prometheus.NewGauge(
            prometheus.GaugeOpts{
                Name: "db_pool_idle_connections",
                Help: "Number of idle database connections",
            },
        ),
        poolOpenConnections: prometheus.NewGauge(
            prometheus.GaugeOpts{
                Name: "db_pool_open_connections",
                Help: "Number of open database connections",
            },
        ),
        queryDuration: prometheus.NewHistogramVec(
            prometheus.HistogramOpts{
                Name: "db_query_duration_seconds",
                Help: "Database query duration in seconds",
                Buckets: []float64{0.001, 0.01, 0.1, 1, 5, 10},
            },
            []string{"operation", "table"},
        ),
        queryErrors: prometheus.NewCounterVec(
            prometheus.CounterOpts{
                Name: "db_query_errors_total",
                Help: "Total number of database query errors",
            },
            []string{"operation", "table", "error_type"},
        ),
    }
}

func (dm *DBMetrics) RecordQuery(operation, table string, duration time.Duration, err error) {
    dm.queryDuration.WithLabelValues(operation, table).Observe(duration.Seconds())
    if err != nil {
        errorType := "unknown"
        if errors.Is(err, sql.ErrNoRows) {
            errorType = "no_rows"
        } else if errors.Is(err, sql.ErrTxDone) {
            errorType = "tx_done"
        } else if errors.Is(err, sql.ErrConnDone) {
            errorType = "conn_done"
        }
        dm.queryErrors.WithLabelValues(operation, table, errorType).Inc()
    }
}
```

### 3. Business Metrics

```go
type BusinessMetrics struct {
    ordersProcessed     *prometheus.CounterVec
    revenueGenerated    *prometheus.GaugeVec
    userActivity       *prometheus.CounterVec
    cacheHitRatio      *prometheus.GaugeVec
}

func NewBusinessMetrics() *BusinessMetrics {
    return &BusinessMetrics{
        ordersProcessed: prometheus.NewCounterVec(
            prometheus.CounterOpts{
                Name: "orders_processed_total",
                Help: "Total number of orders processed",
            },
            []string{"status", "payment_method"},
        ),
        revenueGenerated: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: "revenue_generated",
                Help: "Total revenue generated",
            },
            []string{"currency", "region"},
        ),
        userActivity: prometheus.NewCounterVec(
            prometheus.CounterOpts{
                Name: "user_activity_total",
                Help: "Total user activities",
            },
            []string{"activity_type", "user_tier"},
        ),
        cacheHitRatio: prometheus.NewGaugeVec(
            prometheus.GaugeOpts{
                Name: "cache_hit_ratio",
                Help: "Cache hit ratio",
            },
            []string{"cache_type", "operation"},
        ),
    }
}
```

### 4. Custom Collectors

```go
type SystemMetrics struct {
    cpuUsage    prometheus.Gauge
    memoryUsage prometheus.Gauge
    goroutineCount prometheus.Gauge
}

func (sm *SystemMetrics) Collect(ch chan<- prometheus.Metric) {
    // CPU Usage
    if cpuPercent, err := getCPUUsage(); err == nil {
        sm.cpuUsage.Set(cpuPercent)
    }
    
    // Memory Usage
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    sm.memoryUsage.Set(float64(m.Alloc))
    
    // Goroutine Count
    sm.goroutineCount.Set(float64(runtime.NumGoroutine()))
    
    ch <- sm.cpuUsage
    ch <- sm.memoryUsage
    ch <- sm.goroutineCount
}

func (sm *SystemMetrics) Describe(ch chan<- *prometheus.Desc) {
    ch <- sm.cpuUsage.Desc()
    ch <- sm.memoryUsage.Desc()
    ch <- sm.goroutineCount.Desc()
}
```

## Service Integration

### 1. Metrics Registration

```go
func RegisterMetrics() {
    // Register HTTP metrics
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(requestDuration)
    
    // Register database metrics
    dbMetrics := NewDBMetrics()
    prometheus.MustRegister(dbMetrics.poolConnections)
    prometheus.MustRegister(dbMetrics.poolIdleConnections)
    prometheus.MustRegister(dbMetrics.queryDuration)
    prometheus.MustRegister(dbMetrics.queryErrors)
    
    // Register business metrics
    businessMetrics := NewBusinessMetrics()
    prometheus.MustRegister(businessMetrics.ordersProcessed)
    prometheus.MustRegister(businessMetrics.revenueGenerated)
    
    // Register custom collector
    systemMetrics := &SystemMetrics{
        cpuUsage: prometheus.NewGauge(prometheus.GaugeOpts{
            Name: "system_cpu_usage_percent",
            Help: "System CPU usage percentage",
        }),
        memoryUsage: prometheus.NewGauge(prometheus.GaugeOpts{
            Name: "system_memory_usage_bytes",
            Help: "System memory usage in bytes",
        }),
        goroutineCount: prometheus.NewGauge(prometheus.GaugeOpts{
            Name: "go_goroutines",
            Help: "Number of goroutines",
        }),
    }
    prometheus.MustRegister(systemMetrics)
}
```

### 2. Metrics Endpoint

```go
func SetupMetricsEndpoint() {
    http.Handle("/metrics", promhttp.Handler())
    
    // Add custom metrics endpoint with additional info
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        status := map[string]interface{}{
            "status": "healthy",
            "timestamp": time.Now().Unix(),
            "version": "1.0.0",
            "goroutines": runtime.NumGoroutine(),
        }
        json.NewEncoder(w).Encode(status)
    })
}
```

## Advanced Patterns

### 1. Distributed Tracing Integration

```go
type TracingMetrics struct {
    traceDuration *prometheus.HistogramVec
    traceErrors   *prometheus.CounterVec
}

func (tm *TracingMetrics) RecordTrace(traceID, operation string, duration time.Duration, err error) {
    tm.traceDuration.WithLabelValues(operation).Observe(duration.Seconds())
    if err != nil {
        tm.traceErrors.WithLabelValues(operation, "error").Inc()
    }
}
```

### 2. Circuit Breaker Metrics

```go
type CircuitBreakerMetrics struct {
    state           prometheus.Gauge
    failures        prometheus.Counter
    successes       prometheus.Counter
    timeouts        prometheus.Counter
    requestsTotal   prometheus.Counter
}

func (cbm *CircuitBreakerMetrics) RecordRequest(success bool, timeout bool, failure bool) {
    cbm.requestsTotal.Inc()
    if success {
        cbm.successes.Inc()
    }
    if timeout {
        cbm.timeouts.Inc()
    }
    if failure {
        cbm.failures.Inc()
    }
}
```

### 3. Cache Metrics

```go
type CacheMetrics struct {
    hits           prometheus.Counter
    misses         prometheus.Counter
    sets           prometheus.Counter
    evictions      prometheus.Counter
    size           prometheus.Gauge
}

func (cm *CacheMetrics) RecordOperation(operation string) {
    switch operation {
    case "hit":
        cm.hits.Inc()
    case "miss":
        cm.misses.Inc()
    case "set":
        cm.sets.Inc()
    case "eviction":
        cm.evictions.Inc()
    }
}
```

## Configuration Best Practices

### 1. Metric Naming Conventions

```go
// Good naming
const (
    metricNamespace = "myapp"
    metricSubsystem = "http"
)

var (
    requestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Namespace: metricNamespace,
            Subsystem: metricSubsystem,
            Name:      "requests_total",
            Help:      "Total number of HTTP requests",
        },
        []string{"method", "endpoint", "status"},
    ),
)
```

### 2. Label Management

```go
// Use consistent labels across metrics
var commonLabels = []string{"service", "version", "environment"}

// Avoid high cardinality labels
// BAD: user_id, request_id (too many unique values)
// GOOD: method, endpoint, status_code (limited values)
```

## Interview Questions

### Technical Questions
1. How do you implement custom metrics in a Go microservice?
2. What are the different types of Prometheus metrics and when would you use each?
3. How do you instrument database connection pools with Prometheus?
4. Explain how to create middleware for HTTP request metrics.
5. How do you handle metric aggregation across multiple microservice instances?

### Design Questions
1. Design a metrics collection strategy for a microservices architecture.
2. How would you implement business metrics for an e-commerce platform?
3. Explain how to monitor circuit breaker patterns with Prometheus.
4. Design a dashboard strategy for Go microservices monitoring.
5. How do you ensure metrics don't impact application performance?

### Practical Coding
1. Write a Go middleware that tracks HTTP request metrics.
2. Implement a custom Prometheus collector for system metrics.
3. Create metrics for database query performance monitoring.
4. Design a metrics structure for a payment processing service.
5. Implement cache performance metrics with Prometheus.

## Best Practices

1. **Keep metrics lightweight** - Avoid expensive operations in metric collection
2. **Use appropriate metric types** - Counter for counting, Gauge for values, Histogram for distributions
3. **Limit label cardinality** - Avoid unlimited unique label values
4. **Document metrics** - Clear help text and naming conventions
5. **Test metrics** - Verify metric collection in unit tests
6. **Monitor metrics** - Set up alerts for critical metrics
7. **Version metrics** - Handle metric changes gracefully across versions
