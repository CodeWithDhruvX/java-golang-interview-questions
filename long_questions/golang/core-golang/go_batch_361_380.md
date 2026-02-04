## 🟣 System Design and Observability (Questions 361-380)

### Question 361: How do you design a rate limiter in Go?

**Answer:**
A rate limiter controls the number of requests a user can make in a given timeframe.

**Token Bucket Algorithm (Memory-based):**
Using `golang.org/x/time/rate`.

```go
package main

import (
    "golang.org/x/time/rate"
    "net/http"
    "time"
)

var limiter = rate.NewLimiter(1, 5) // 1 req/sec, burst of 5

func limit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Too many requests", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello"))
    })
    http.ListenAndServe(":8080", limit(mux))
}
```

**Distributed Rate Limiter (Redis):**
For multiple instances, use Redis (Fixed Window or Sliding Window).
- **Key:** `rate_limit:{user_id}`
- **Value:** Count
- **TTL:** 1 second

Increment key; if > limit, block.

---

### Question 362: How do you implement distributed tracing in Go?

**Answer:**
Use **OpenTelemetry (OTel)** to propagate context (Trace ID, Span ID) across services.

**Setup with Jaeger:**

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/trace"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() *sdktrace.TracerProvider {
    exporter, _ := jaeger.New(jaeger.WithCollectorEndpoint())
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String("my-service"),
        )),
    )
    otel.SetTracerProvider(tp)
    return tp
}

func handler(w http.ResponseWriter, r *http.Request) {
    ctx, span := otel.Tracer("my-service").Start(r.Context(), "handler-span")
    defer span.End()
    
    // Pass ctx to downstream functions/services
    doWork(ctx)
}
```

---

### Question 363: How do you handle distributed transactions?

**Answer:**
Distributed transactions (across microservices) cannot use ACID properties of a single DB.

**Patterns:**
1. **Saga Pattern (Choreography):**
   - Service A publishes `OrderCreated`.
   - Service B listens, reserves inventory, publishes `InventoryReserved`.
   - If Service B fails, it publishes `InventoryFailed`.
   - Service A listens to fail event and executes **Compensating Transaction** (undo).

2. **Saga Pattern (Orchestration):**
   - A central "Orchestrator" service tells A, then B, then C what to do.
   - If any fail, the orchestrator triggers rollbacks.

3. **Two-Phase Commit (2PC):** (Avoid in microservices - blocking & slow).

---

### Question 364: What is the Circuit Breaker pattern and how to implement it?

**Answer:**
Prevents cascading failures when a downstream service is down.

**States:**
- **Closed:** Normal operation.
- **Open:** Fails immediately (after N errors).
- **Half-Open:** Allows 1 test request. If success -> Closed; else -> Open.

**Library:** `github.com/sony/gobreaker`

```go
var cb *gobreaker.CircuitBreaker

func init() {
    st := gobreaker.Settings{
        Name:        "HTTP_GET",
        MaxRequests: 1,
        Timeout:     5 * time.Second,
        ReadyToTrip: func(counts gobreaker.Counts) bool {
             return counts.ConsecutiveFailures > 3
        },
    }
    cb = gobreaker.NewCircuitBreaker(st)
}

func callService() ([]byte, error) {
    body, err := cb.Execute(func() (interface{}, error) {
        resp, err := http.Get("http://example.com")
        if err != nil { return nil, err }
        return io.ReadAll(resp.Body)
    })
    return body.([]byte), err
}
```

---

### Question 365: How do you design a notification system in Go?

**Answer:**
A system to send Email, SMS, Push notifications.

**Architecture:**
1. **API Gateway:** Receives `POST /notify`.
2. **Message Queue (Kafka/RabbitMQ):** Decouples API from senders.
   - Topics: `notify-email`, `notify-sms`.
3. **Workers:** 
   - `EmailWorker`: Consumes `notify-email`, calls SendGrid/SES.
   - `SMSWorker`: Consumes `notify-sms`, calls Twilio.
4. **Retry Logic:** Exponential backoff for failed external calls.
5. **Rate Limiting:** Throttle sends per user.

**Go Interface:**
```go
type Notifier interface {
    Send(user User, msg string) error
}

type EmailNotifier struct {} // impl...
type SMSNotifier struct {}   // impl...
```

---

### Question 366: How do you handle configuration hot-reloading?

**Answer:**
Reload config without restarting the app.

**Using `viper`:**
```go
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
    fmt.Println("Config file changed:", e.Name)
    // Re-read or update global config struct
    viper.Unmarshal(&AppConfig)
    updateConnectionPool() // if needed
})
```

**Kubernetes:**
- Update `ConfigMap`.
- K8s updates mounted file.
- Viper/fsnotify detects change -> Reload.

---

### Question 367: How do you implement health checks for microservices?

**Answer:**
Expose `/health` endpoints.

1. **Liveness Probe (Am I running?):**
   - Returns 200 OK if process is up.
   - If 500/Timeout, K8s restarts pod.

2. **Readiness Probe (Can I serve traffic?):**
   - Checks dependencies (DB, Cache).
   - Returns 200 OK only if connected.
   - If 500, K8s stops sending traffic to this pod.

**Implementation:**
```go
func health(w http.ResponseWriter, r *http.Request) {
    if err := db.Ping(); err != nil {
        w.WriteHeader(503)
        return
    }
    w.WriteHeader(200)
    w.Write([]byte("OK"))
}
```

---

### Question 368: How do you design a URL shortener in Go?

**Answer:**
**components:**
- **API:** `POST /shorten`, `GET /{short_code}`.
- **DB:** SQL (Postgres) or NoSQL (DynamoDB/Redis).
- **ID Generation:**
  - Base62 encoding of an auto-increment integer ID.
  - Or K-Sortable Unique ID (KSUID/Snowflake).

**Code Logic:**
1. `POST`: Insert URL to DB -> Get ID (1001) -> Base62(1001) = "g7" -> Return `site.com/g7`.
2. `GET /g7`: Base62Decode("g7") -> 1001 -> Select from DB -> `http.Redirect(301)`.

**Concurrency:**
- DB handles collision (Unique constraint).
- Read-heavy -> Cache in Redis (`g7` -> `full_url`).

---

### Question 369: How do you debug high CPU usage in a Go app?

**Answer:**
Use `pprof`.

1. **Enable pprof:**
   ```go
   import _ "net/http/pprof"
   go func() { log.Println(http.ListenAndServe("localhost:6060", nil)) }()
   ```

2. **Capture Profile:**
   ```bash
   go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
   ```

3. **Analyze:**
   - `top`: Shows functions using most CPU.
   - `list <func>`: Shows source code with CPU time per line.
   - `web`: Visual graph in browser.

**Common Culprits:**
- Tight loops.
- Excessive Garbage Collection (check allocations).
- Serialization (JSON) in hot paths.

---

### Question 370: How do you debug a memory leak in Go?

**Answer:**
Memory leaks in Go are usually:
1. **Goroutine Leaks:** Goroutines blocked forever (waiting on nil channel, unclosed body).
2. **Retained References:** Global map growing forever, Slice with small window into large array.

**Debugging:**
1. **Capture Heap Profile:**
   ```bash
   go tool pprof http://localhost:6060/debug/pprof/heap
   ```
2. **Compare Profiles (diff):**
   Capture heap at T1 and T2, then `pprof -base heap1 heap2`.
3. **Analyze:** Look for objects with high `inuse_space` or `inuse_objects`.

---

### Question 371: What is the Sidecar pattern and how is it used with Go?

**Answer:**
Deploying a helper container alongside the main Go application container in the same Pod.

**Use Cases:**
- **Logging:** Sidecar reads log files, forwards to Splunk.
- **Proxy:** Envoy/Istio sidecar handles mTLS, circuit breaking, tracing (Service Mesh).
- **Config:** Sidecar watches remote config and updates local file.

**Go Impact:**
- Go app talks to `localhost` for these services (e.g., `localhost:5432` proxy to DB).
- Simplifies Go code (offloads infrastructure concerns).

---

### Question 372: How do you design a job scheduler in Go?

**Answer:**
**Simple (In-Memory):**
- Ticker + Goroutine.
```go
ticker := time.NewTicker(1 * time.Hour)
for range ticker.C {
    go cleanup()
}
```

**Distributed (Robust):**
- Use a library like `gocron` or `robfig/cron`.
- **Leader Election:** Only one instance runs the job (using Redis lock or Etcd).
- **Persistent Queue:** If job fails, retry later (patterns like RabbitMQ Delayed Exchange).

**Leader Election Pattern:**
```go
// Redis SetNX (Set if Not Exists) with TTL
success, _ := redis.SetNX(ctx, "leader_lock", "my-id", 10*time.Second).Result()
if success {
    runJob()
    // Refresh lock periodically (Heartbeat)
}
```

---

### Question 373: How do you implement API versioning?

**Answer:**
1. **URL Path (Most Common):** `/api/v1/users`
   - Easy to see and route.
   - Using `http.NewServeMux`:
     ```go
     mux.Handle("/api/v1/", v1Handler)
     mux.Handle("/api/v2/", v2Handler)
     ```

2. **Header:** `Accept: application/vnd.myapi.v1+json`
   - Cleaner URLs.
   - Harder to test in browser.

3. **Query Param:** `/users?v=1`

**Best Practice:**
Keep logic separate (`package v1`, `package v2`) to avoid `if v1 {...} else {...}` spaghetti code.

---

### Question 374: What is Context Propagation?

**Answer:**
Passing request-scoped values (User ID, Trace ID, Auth Token) through the call chain using `context.Context`.

**Example:**
1. Middleware extracts `Trace-ID` header -> puts in `ctx`.
2. Handler calls Database -> passes `ctx`.
3. DB Driver (if instrumented) uses `ctx` to log Trace ID or handle cancellation.

**Custom Context Value:**
```go
type key int
const userKey key = 0

// Set
ctx = context.WithValue(ctx, userKey, "Alice")

// Get
user, ok := ctx.Value(userKey).(string)
```
*Note: Only use context for request-scoped data, not for optional parameters.*

---

### Question 375: How to handle 10 million concurrent WebSocket connections?

**Answer:**
The "C10M" problem.

1. **Kernel Tuning:**
   - Increase Max Open Files (`ulimit -n`).
   - Tune TCP buffer sizes (`net.ipv4.tcp_rmem`, `wmem`).

2. **Go Optimization:**
   - **Goroutines:** 10M goroutines ≈ 20-40GB RAM (2KB each). Doable on large implementation.
   - **Epoll (Advanced):** Instead of 1 goroutine per conn, use `syscall.Epoll` (on Linux) to manage connections event-driven (Library: `gnet` or `evio`).

3. **Architecture:**
   - **User Level Sharding:** Load balancer distributes users to different server nodes.
   - **State:** Keep connection state minimal.

---

### Question 376: How do you secure internal gRPC services?

**Answer:**
1. **mTLS (Mutual TLS):**
   - Both Client and Server authenticate each other using certificates.
   - Best for zero-trust internal networks.

2. **Token Authentication (Interceptors):**
   - Client sends JWT in Metadata (`authorization` header).
   - Server Interceptor validates JWT.

**Code (Unary Interceptor):**
```go
func authInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
    md, _ := metadata.FromIncomingContext(ctx)
    token := md["authorization"][0]
    
    if !valid(token) {
        return nil, status.Error(codes.Unauthenticated, "invalid token")
    }
    return handler(ctx, req)
}
```

---

### Question 377: Explain the Outbox Pattern.

**Answer:**
Ensures **Atomicity** when writing to DB and publishing an Event.

**Problem:** 
- `DB.Insert(Order)` succeeds.
- `Kafka.Publish(OrderCreated)` fails.
- Inconsistency!

**Solution (Outbox):**
1. Start Transaction.
2. `DB.Insert(Order)`.
3. `DB.Insert(Outbox, {Event: OrderCreated})`.
4. Commit Transaction. (Atomic).
5. **Separate Process (CDC/Poller):** Reads `Outbox` table and publishes to Kafka.
6. Delete from `Outbox` after publish.

---

### Question 378: How do you implement checking for race conditions in CI?

**Answer:**
Use the Go Race Detector.

**Command:**
```bash
go test -race ./...
```

**In CI (GitHub Actions):**
```yaml
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
      - run: go test -race -v ./...
```
**Caveat:** Race detector slows down execution (5-10x) and uses more memory. Run it on integration tests or a subset of tests if full suite is too slow.

---

### Question 379: How do you optimize Go garbage collection for low latency?

**Answer:**
1. **Reduce Allocations:**
   - Use `sync.Pool`.
   - Pre-allocate slices/maps (`make([]int, 0, 1000)`).
   - Use value types instead of pointers where appropriate (less work for scanner).

2. **GOGC Tuning:**
   - `GOGC=200`: GC runs less often (uses more RAM).
   - `GOGC=off`: Disable manual GC (dangerous).

3. **Memory Ballast (Legacy trick):**
   - Allocate a huge byte array (e.g., 1GB) on startup.
   - Increases heap size, so GC triggers less frequently (since it triggers based on % growth).
   - *Note: Go 1.19 `SetMemoryLimit` (Soft Limit) is the modern way.*

---

### Question 380: What is Semantic Versioning in Go modules?

**Answer:**
Go modules strictly follow SemVer (`vMajor.Minor.Patch`).

- **v1.x.x:** Public API is stable.
- **v2.0.0:** Breaking changes.
  - **Import Path changes:** `github.com/user/lib/v2`.
  - Directory structure usually involves a `v2` folder or `go.mod` change.
  
**Pseudo-version:**
Using a specific commit hash: `v0.0.0-20230101-abcdef123456`.

**Direct/Indirect:**
- `// indirect` in `go.mod`: Dependency of a dependency.

---
