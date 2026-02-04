# Advanced Go Combination Skills: Cloud-Native & Distributed Systems

This document covers high-value "combination skills" that bridge the gap between core Golang knowledge and Senior/Lead Engineer requirements. It focuses on how Go interacts with the modern cloud-native ecosystem.

---

## 1. Cloud-Native Go (Go + Kubernetes)

### 1.1 Graceful Shutdowns
In Kubernetes, when a pod is terminated, it receives a `SIGTERM` signal. Your application must handle this signal to stop accepting new requests, finish current requests, and close connections (DB, Redis) cleanly.

**Implementation Pattern:**
```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{Addr: ":8080"}

	// run server in a goroutine so it doesn't block the main thread
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
```

### 1.2 Liveness & Readiness Probes
*   **Liveness**: "Am I broken?" (If yes, restart me). Checks if the app is deadlocked or permanently broken.
*   **Readiness**: "Can I take traffic?" (If no, remove from load balancer). Checks if dependencies (like DB) are connected.

**Code Example:**
```go
func healthz(w http.ResponseWriter, r *http.Request) {
    // Simple liveness probe
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
}

func readyz(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Readiness probe: Check DB connection
        if err := db.Ping(); err != nil {
            http.Error(w, "Functionality not available", http.StatusServiceUnavailable)
            return
        }
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ready"))
    }
}
```

### 1.3 Kubernetes Operators & Controllers
Operators extend Kubernetes by using Custom Resource Definitions (CRDs). They act like a SRE codified in software.

*   **Key Concept**: The "Reconcile Loop". The controller watches the *desired state* (from YAML) and compares it to the *actual state* (in the cluster), then takes action to make them match.
*   **Tools**: `Kubebuilder` or `Operator SDK` are the standard tools to scaffold these projects.

---

## 2. High-Performance Communication (Go + gRPC)

### 2.1 gRPC vs REST
| Feature | REST (JSON/HTTP1.1) | gRPC (Protobuf/HTTP2) |
| :--- | :--- | :--- |
| **Payload** | Text (JSON), larger size | Binary (Protobuf), smaller size |
| **Protocol** | HTTP/1.1 (Request/Response) | HTTP/2 (Multiplexing, Streaming) |
| **Contract** | Loose (OpenAPI/Swagger) | Strict (.proto files) |
| **Use Case** | Public APIs, Browser Clients | Internal Microservices, Low Latency |

### 2.2 gRPC Interceptors
Interceptors are middleware for gRPC. They are used for logging, authentication, metrics, and tracing.

**Unary Interceptor Example (Logging):**
```go
func UnaryServerInterceptor(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (interface{}, error) {
    log.Printf("Method: %s", info.FullMethod)
    
    // Call the handler to complete the normal request
    resp, err := handler(ctx, req)
    
    // Logic after handling the request
    return resp, err
}
```

---

## 3. Distributed Data Systems

### 3.1 Kafka in Go
Use **Sarama** or **Confluent-Kafka-Go** libraries.
*   **Consumer Groups**: Crucial for scalability. Allows multiple instances of your service to share the load of processing messages from a topic.
*   **At-Least-Once Delivery**: Handle potential duplicate messages (idempotency).

**Consumer Group Snippet (Sarama):**
```go
// Setup config
config := sarama.NewConfig()
config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
config.Consumer.Offsets.Initial = sarama.OffsetOldest

// Create consumer group
group, _ := sarama.NewConsumerGroup([]string{"localhost:9092"}, "my-group", config)

// Consume loop
for {
    group.Consume(ctx, []string{"my-topic"}, &myHandler{})
}
```

### 3.2 Redis Pipelines & Pub/Sub
*   **Pipelining**: Send multiple commands to Redis in a single network round-trip. Drastically reduces latency for batch operations.
*   **Pub/Sub**: Real-time messaging where publishers send messages to channels and subscribers listen.

### 3.3 Database Tuning (`database/sql`)
The default `sql.Open` does not constrain connections, which can crash your DB under load.
**ALWAYS** tune these settings:
```go
db.SetMaxOpenConns(25) // Limit total open connections
db.SetMaxIdleConns(25) // Limit idle connections in pool
db.SetConnMaxLifetime(5 * time.Minute) // Recycle connections to prevent stale timeouts
```

---

## 4. Observability & Reliability

### 4.1 Prometheus Metrics
Don't just rely on logs. Expose metrics.
*   **Counter**: Things that go up (e.g., `requests_total`).
*   **Gauge**: Things that fluctuate (e.g., `goroutines_active`, `memory_usage`).
*   **Histogram**: Distributions (e.g., `request_duration_seconds` buckets).

**Using `prometheus/client_golang`:**
```go
var (
    opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
        Name: "myapp_processed_ops_total",
        Help: "The total number of processed events",
    })
)

func recordMetrics() {
    go func() {
        for {
            opsProcessed.Inc()
            time.Sleep(2 * time.Second)
        }
    }()
}

func main() {
    recordMetrics()
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":2112", nil)
}
```

### 4.2 Circuit Breakers
Prevent cascading failures. If a downstream service is failing, stop calling it for a while to let it recover.
*   **Library**: `github.com/sony/gobreaker` is a popular choice.

---

## 5. Advanced Testing

### 5.1 Testcontainers
Stop mocking everything. Use **Testcontainers** to spin up real Infrastructure (Postgres, Redis) in Docker for integration tests.

```go
func TestWithRealDB(t *testing.T) {
    ctx := context.Background()
    postgresContainer, err := postgres.RunContainer(ctx, 
        testcontainers.WithImage("postgres:15-alpine"),
        postgres.WithDatabase("testdb"),
        postgres.WithUsername("user"),
        postgres.WithPassword("password"),
    )
    if err != nil {
        t.Fatal(err)
    }
    defer postgresContainer.Terminate(ctx)
    
    // Connect to the containerized DB and run tests...
}
```

### 5.2 Fuzz Testing (Go 1.18+)
Find edge cases by feeding random data to your functions.

```go
func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc) // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```
