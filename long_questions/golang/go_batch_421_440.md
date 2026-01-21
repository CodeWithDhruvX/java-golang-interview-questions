## 🟣 Error Handling & Observability (Questions 421-440)

### Question 421: How do you create custom error types in Go?

**Answer:**
Define a struct that implements the `error` interface (specifically the `Error() string` method).

```go
type ValidationError struct {
    Field string
    Msg   string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("invalid %s: %s", e.Field, e.Msg)
}

func main() {
    err := &ValidationError{Field: "email", Msg: "missing @ symbol"}
    fmt.Println(err)
}
```

---

### Question 422: How does Go 1.20+ `errors.Join` and `errors.Is` work?

**Answer:**
- **`errors.Join`**: Combines multiple errors into a single error that wraps them all. Useful for validation where you want to report multiple failures.
- **`errors.Is`**: Checks if an error wraps a specific target error (unwrapping recursively).

```go
err1 := errors.New("err1")
err2 := errors.New("err2")
joined := errors.Join(err1, err2)

if errors.Is(joined, err1) {
    fmt.Println("joined contains err1")
}
```

---

### Question 423: How do you implement error wrapping and unwrapping?

**Answer:**
Use the `%w` verb in `fmt.Errorf` to wrap an error.
Use `errors.Unwrap` to get the underlying error (though usually `errors.Is` or `errors.As` is preferred).

```go
baseErr := errors.New("db connection failed")
wrappedErr := fmt.Errorf("query failed: %w", baseErr)

fmt.Println(wrappedErr) // "query failed: db connection failed"

// Unwrapping
origin := errors.Unwrap(wrappedErr)
fmt.Println(origin == baseErr) // true
```

---

### Question 424: What are best practices for error categorization?

**Answer:**
1.  **Sentinel Errors (Values):** Global variables like `io.EOF`. Good for simple comparisons.
2.  **Error Types:** Custom structs. Good when you need extra data (e.g., Status Code, Field Name).
3.  **Behavior Interfaces:** Define an interface (e.g., `Temporary() bool`) to check error properties without binding to specific types.

---

### Question 425: How do you handle critical vs recoverable errors?

**Answer:**
- **Recoverable:** Return as a normal `error` value. The caller decides (retry, log, ignore).
- **Critical (Unrecoverable):** Use `panic` only for truly exceptional states (startup config missing, programmer error).
- **Graceful Degradation:** If a non-essential service fails (e.g., recommendation engine), log the error but return the main content (default list).

---

### Question 426: How do you recover from panics in goroutines?

**Answer:**
Use `defer` with `recover()`. **Critical:** Go panics do strictly confined to the goroutine where they happen. A panic in a spawned goroutine crashes the whole app if not recovered *inside* that goroutine.

```go
func safeGo(f func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Recovered:", r)
            }
        }()
        f()
    }()
}
```

---

### Question 427: How to capture stack traces on error?

**Answer:**
The standard `error` doesn't carry stack traces. Use `pkg/errors` (common legacy) or modern replacements.

```go
import "github.com/pkg/errors"

func do() error {
    return errors.New("something went wrong")
}

func main() {
    err := do()
    fmt.Printf("%+v", err) // Prints full stack trace
}
```

---

### Question 428: How do you notify Sentry/Bugsnag from Go?

**Answer:**
Use their official SDKs. Typically, you initialize the client in `main` and defer a flush.

```go
import "github.com/getsentry/sentry-go"

func main() {
    sentry.Init(sentry.ClientOptions{Dsn: "your-dsn"})
    defer sentry.Flush(2 * time.Second)

    _, err := db.Query(...)
    if err != nil {
        sentry.CaptureException(err)
    }
}
```

---

### Question 429: How do you do structured error reporting in Go?

**Answer:**
Use specialized logging libraries (like `slog` or `zap`) instead of just string manipulation errors. Create error types that serialize to JSON nicely.

```go
logger.Error("login failed",
    "user_id", u.ID,
    "error", err,
    "ip", r.RemoteAddr,
)
```

---

### Question 430: How do you correlate logs, errors, and traces together?

**Answer:**
Use a **Correlation ID** (Trace ID).
1.  Generate a UUID at the entry point (Middleware).
2.  Store it in `context.Context`.
3.  Pass `ctx` to all functions.
4.  Include Trace ID in every log line and error report.

---

### Question 431: How would you add distributed tracing to an existing Go service?

**Answer:**
Use **OpenTelemetry (OTel)**.
1.  Initialize a TracerProvider (exporter to Jaeger/Tempo).
2.  Instrument HTTP/gRPC handlers (`otelhttp`, `otelgrpc`).
3.  Use `tracer.Start(ctx, "spanName")` manually for internal logic.

---

### Question 432: What are tags, attributes, and spans in tracing?

**Answer:**
- **Span:** Represents a single unit of work (e.g., "DB Query", "HTTP Request"). Has a start/end time.
- **Trace:** A tree of Spans representing an entire lifecycle of a request across services.
- **Attributes (Tags):** Key-value pairs attached to a Span (e.g., `http.method="GET"`, `db.statement="SELECT..."`) to verify/filter data later.

---

### Question 433: What is a traceparent header?

**Answer:**
A W3C standard HTTP header used to propagate trace context between services.
Format: `version-trace_id-parent_id-flags`
Example: `00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01`
It tells the downstream service: "Here is the Transaction ID, and here is the ID of the caller span."

---

### Question 434: How do you send custom metrics to Prometheus?

**Answer:**
Use the `prometheus/client_golang` library.

```go
var requestCount = promauto.NewCounter(prometheus.CounterOpts{
    Name: "http_requests_total",
    Help: "Total number of HTTP requests",
})

func handler(w http.ResponseWriter, r *http.Request) {
    requestCount.Inc()
    // handle request
}

func main() {
    http.Handle("/metrics", promhttp.Handler()) // Expose endpoint
    http.ListenAndServe(":8080", nil)
}
```

---

### Question 435: What is RED metrics model and how do you apply it?

**Answer:**
A standard for monitoring services:
- **R**ate: Number of requests per second (Counter).
- **E**rrors: Number of failed requests per second (Counter).
- **D**uration: How long requests take (Histogram).

Apply by middleware that measures these three for every HTTP endpoint.

---

### Question 436: How do you expose application health and readiness probes?

**Answer:**
Create 2 endpoints:
1.  `/healthz` (Liveness): Returns 200 OK if the app builds/runs. K8s restarts pod if this fails.
2.  `/readyz` (Readiness): Returns 200 OK only if DB/Cache connections are active. K8s stops sending traffic if this fails.

```go
func readiness(w http.ResponseWriter, r *http.Request) {
    if db.Ping() != nil {
        w.WriteHeader(503)
        return
    }
    w.WriteHeader(200)
}
```

---

### Question 437: What’s the difference between logs, metrics, and traces?

**Answer:**
- **Logs:** "What happened?" (Unstructured/Structured text events). *High volume.*
- **Metrics:** "Is it healthy?" (Aggregated numbers: counts, gauges). *Low storage cost.*
- **Traces:** "Where did it happen?" (Request lifecycle across services). *Debugging latency.*

---

### Question 438: How do you benchmark error impact on performance?

**Answer:**
Write a benchmark comparing a "Happy Path" vs "Error Path". Creating error strings and formatting stack traces is expensive.

```go
func BenchmarkError(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = fmt.Errorf("error %d", i) // allocation heavy
    }
}
```
**Optimization:** Use pre-defined variables (`var ErrNotFound = errors.New("...")`) to avoid allocation.

---

### Question 439: What’s the tradeoff between verbose and silent error handling?

**Answer:**
- **Verbose:** Easier debugging, but fills logs (noise), slower, and risks leaking info to users (`sql: table not found`).
- **Silent (or Generic):** Better security and UX, but hard to debug production issues.
- **Solution:** Log Verbose errors internally (to files/Sentry), return Generic errors to the API Client.

---

### Question 440: How would you enforce observability in a Go microservice?

**Answer:**
1.  **Middleware:** Attach Logging, Metrics (RED), and Tracing automatically to every request.
2.  **Standards:** Enforce JSON logging format across all teams.
3.  **Context Propagation:** Ensure `ctx` is passed to DB, HTTP clients, and queues to maintain Trace ID.
4.  **Dashboards:** Setup Grafana to visualize the exported Prometheus metrics.

---
