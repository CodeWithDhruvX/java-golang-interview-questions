## 📦 Error Handling & Observability (Questions 741-760)

### Question 741: How do you implement a custom error type in Go?

**Answer:**
Define a struct and implement the `error` interface method `Error() string`.

```go
type APIError struct {
    Code int
    Msg  string
}

func (e *APIError) Error() string {
    return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}
```

---

### Question 742: How do you wrap errors in Go?

**Answer:**
Use `%w` verb in `fmt.Errorf`.

```go
if err != nil {
    return fmt.Errorf("failed to open config: %w", err)
}
```
This preserves the original error chain, allowing `errors.Is` and `errors.As` to inspect the cause.

---

### Question 743: What is `errors.Is()` and `errors.As()` used for?

**Answer:**
- **`errors.Is(err, target)`**: Checks if `err` (or any error in its wrapper chain) matches a specific sentinel value (e.g., `io.EOF`).
- **`errors.As(err, &target)`**: Checks if `err` chain contains a specific **type** (e.g., `*net.OpError`) and assigns it to `target` so you can access fields.

---

### Question 744: How do you categorize errors in large Go applications?

**Answer:**
Use Error Types or Codes.
Define a base `AppError` with an `Op` (Operation) and `Kind` (Permission, NotFound, Internal).
This decouples the messiness of low-level errors (SQL/Network) from the HTTP Status Code logic.

---

### Question 745: How do you log structured errors in Go?

**Answer:**
Use `slog` (Go 1.21+) or `zap`.

```go
logger.Error("login failed",
    slog.String("user", "bob"),
    slog.Any("error", err),
)
```
Outputs JSON: `{"msg":"login failed", "user":"bob", "error":"invalid password"}`.

---

### Question 746: How do you use Sentry/Bugsnag with Go?

**Answer:**
Initialize client -> Defer Flush -> Capture.
Use `sentry-go` SDK.
**Middleware:** `sentryhttp` middleware automatically captures panics and sends them to Sentry.

---

### Question 747: How do you implement centralized error logging?

**Answer:**
Log to `stdout` locally.
In Production, the container runtime (Docker/K8s) captures stdout.
**Fluentd / Filebeat** runs as a daemon, tails the logs, and pushes them to **Elasticsearch / Loki**.

---

### Question 748: What is the role of stack traces in debugging Go apps?

**Answer:**
Essential for panics.
Standard `error` does not have stack straces.
Use `github.com/pkg/errors` (Wrap) so that when you log, you can print `%+v` to see the exact line number where the error originated, not just "file not found".

---

### Question 749: How do you implement panic recovery with context?

**Answer:**
Recovered panics lose context (request ID, user).
In your recovery middleware:
1.  Read values from `r.Context()`.
2.  Include them in the log message when `recover()` catches a crash.
3.  Responds with 500 Internal Server Error.

---

### Question 750: How do you differentiate retryable vs fatal errors?

**Answer:**
Define a `Temporary() bool` interface.
If `net.Error.Temporary() == true` (timeout), retry.
If DB returns "Unique Constraint Violation", do not retry (Fatal).

---

### Question 751: How do you expose Prometheus metrics in Go?

**Answer:**
Use `github.com/prometheus/client_golang`.
1.  Define `var opsProcessed = promauto.NewCounter(...)`.
2.  Increment `opsProcessed.Inc()`.
3.  Serve `http.Handle("/metrics", promhttp.Handler())`.

---

### Question 752: How do you set up OpenTelemetry in Go?

**Answer:**
1.  Create a TraceProvider (Exporter -> Jaeger/OTLP).
2.  Use `otel.SetTracerProvider`.
3.  In code: `tracer := otel.Tracer("name")`.
4.  `ctx, span := tracer.Start(ctx, "funcName")`.
5.  `defer span.End()`.

---

### Question 753: How do you trace gRPC requests in Go?

**Answer:**
Use middleware: `go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc`.
Client: `grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor())`.
Server: `grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor())`.
Matches trace headers automatically across services.

---

### Question 754: How do you record and export application traces?

**Answer:**
Spans are recorded in memory (buffered).
The Exporter pushes them via HTTP/gRPC to a Collector (Jaeger/Tempo) immediately or periodically (BatchSpanProcessor is more efficient).

---

### Question 755: How do you handle slow endpoints in production Go apps?

**Answer:**
1.  **Timeout:** `http.TimeoutHandler`.
2.  **Instrumentation:** Log duration > 500ms.
3.  **Profiling:** Capture CPU profile during slowness.
4.  **Tracing:** Identify if it's DB, 3rd Party API, or CPU calculation.

---

### Question 756: How do you add custom labels/tags to logs?

**Answer:**
With `slog` or `zap`, you can create a child logger with context.

```go
reqLogger := logger.With("request_id", id, "user_id", uid)
reqLogger.Info("step 1")
reqLogger.Info("step 2")
```
Both logs will have the tags automatically attached.

---

### Question 757: How do you redact sensitive data in logs?

**Answer:**
1.  **Custom Type:** `type Password string` with a `String()` method that returns "***".
2.  **Filter:** Implement a `slog.Handler` that checks keys (like "password", "token") and replaces values before writing.

---

### Question 758: How do you detect memory leaks using Go tools?

**Answer:**
pprof Heap Profile.
Look for high `inuse_objects` counts that constantly increase and never drop after GC runs.

---

### Question 759: How do you instrument performance counters in Go?

**Answer:**
`expvar` package (simple JSON Map of counters at `/debug/vars`).
Or Prometheus Counters/Gauges/Histograms.

---

### Question 760: How do you implement a tracing middleware?

**Answer:**
```go
func TraceMid(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w, r) {
        // Start Span
        ctx, span := tracer.Start(r.Context(), r.URL.Path)
        defer span.End()
        
        r = r.WithContext(ctx)
        next.ServeHTTP(w, r)
        
        // Record Status Code
    })
}
```

---
