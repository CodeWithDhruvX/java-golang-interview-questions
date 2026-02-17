# 🟢 Go Theory Questions: 741–760 Error Handling & Observability

## 741. How do you implement a custom error type in Go?

**Answer:**
We define a struct that implements the `error` interface (`Error() string`).

```go
type ValidationError struct {
    Field string
    Reason string
}
func (e *ValidationError) Error() string {
    return fmt.Sprintf("invalid %s: %s", e.Field, e.Reason)
}
```
This allows us to add context (fields, retry-ability, error codes) that a simple string cannot convey, enabling the caller to type-switch and handle specific errors intelligently.

---

## 742. How do you wrap errors in Go?

**Answer:**
We use `%w` with `fmt.Errorf`.
`return fmt.Errorf("query failed: %w", err)`.

This wraps the original error inside the new one.
It preserves the type hierarchy. We can later use `errors.Is(err, sql.ErrNoRows)` and it will return true even if `ErrNoRows` is wrapped 5 layers deep.
Before Go 1.13, we used `github.com/pkg/errors` (`errors.Wrap`), but now the standard library handles it natively.

---

## 743. What is `errors.Is()` and `errors.As()` used for?

**Answer:**
**`errors.Is(err, target)`**: Checks if `err` wraps `target`. It replaces `err == target`.
**`errors.As(err, &target)`**: Checks if `err` wraps a specific **Type** and assigns it. It replaces type assertions (`err.(*MyError)`).

Example:
`var netErr net.Error`
`if errors.As(err, &netErr) && netErr.Temporary() { retry() }`.
These functions unwrap the error chain automatically to find the match.

---

## 744. How do you categorize errors in large Go applications?

**Answer:**
We use **Sentinel Errors** or **Error Codes**.
1.  **Sentinels**: Global vars `var ErrNotFound = errors.New("...")`. Good for simple comparisons.
2.  **ErrorCode pattern**:
```go
type AppError struct {
    Code    int // 404, 500
    Message string
    Err     error
}
```
The central HTTP middleware checks the `Code` field to determine the HTTP Status (400 vs 500) to return to the user.

---

## 745. How do you log structured errors in Go?

**Answer:**
We use **Slog** (Go 1.21+) or **Zap**.
We don't just log strings; we log Key-Value pairs.

`slog.Error("login failed", "user_id", u.ID, "error", err)`
Output: `{"level":"error", "msg":"login failed", "user_id":101, "error":"password mismatch"}`.
This machine-readable format allows log aggregators (Datadog/Elastic) to index the `user_id` field effectively, making debugging separate from parsing.

---

## 746. How do you use sentry/bugsnag with Go?

**Answer:**
We wrap the main handler with a `Recovery` middleware.
In `defer func()`:
`if r := recover(); r != nil { sentry.CaptureMessage(r) }`.

For explicit errors: `sentry.CaptureException(err)`.
We always attach **Tags** (`release_version`, `environment`) and **Breadcrumbs** ("User clicked Button A") to give context to the crash report.

---

## 747. How do you implement centralized error logging?

**Answer:**
We strictly separate **Propagating** from **Logging**.
Rule: **Log OR Return, never both**.
If a low-level DB function fails, it returns the error (wrapped).
It bubbles up to the **Top-Level Handler**.
The Handler logs it *once* (centrally) and sends a response.
This prevents "Log Noise" where the same error is printed 10 times in the stack trace.

---

## 748. What is the role of stack traces in debugging Go apps?

**Answer:**
Go errors don't have stack traces by default.
We use `pkg/errors` (or custom wrapper) to attach a stack trace at the point of creation.
`errors.WithStack(err)`.
When logging centrally, we print `%+v` to see the stack. Use stack traces for **Panics** and **500 Internal Errors**, but not for **400 Bad Requests** (Validation errors don't need stacks).

---

## 749. How do you implement panic recovery with context?

**Answer:**
Standard `recover()` loses context.
We write a middleware that clones the logger with RequestID.

```go
defer func() {
    if r := recover() {
        logger.With("req_id", ctx.Value("id")).Error("PANIC", "panic", r)
        debug.PrintStack()
        http.Error(w, "Internal Error", 500)
    }
}()
```
This ensures the panic log line is tied to the specific user request that caused it.

---

## 750. How do you differentiate retryable vs fatal errors?

**Answer:**
We define an interface `Retryable` or inspect the error type.

`net.Error` has a `Temporary()` method.
For custom errors:
`type TemporaryError struct { ... }`.
In the retry loop:
```go
if errors.As(err, &tempErr) {
    time.Sleep(backoff)
    continue
}
return err // Fatal
```
Fatal errors (JSON syntax error) will never succeed on retry; Temporary ones (Timeout) might.

---

## 751. How do you expose Prometheus metrics in Go?

**Answer:**
`import "github.com/prometheus/client_golang/prometheus"`.
1.  Define Metric: `var reqCount = prometheus.NewCounter(...)`.
2.  Register: `prometheus.MustRegister(reqCount)`.
3.  Instrument Code: `reqCount.Inc()`.
4.  Expose Handler: `http.Handle("/metrics", promhttp.Handler())`.
Prometheus pulls (scrapes) this endpoint every 15s.

---

## 752. How do you set up OpenTelemetry in Go?

**Answer:**
1.  **Provider**: Configure the TracerProvider (usually exporting to Jaeger/OTLP).
2.  **Propagator**: Set `otel.SetTextMapPropagator` to handle headers (`traceparent`).
3.  **Context**: Use `ctx, span := tracer.Start(ctx, "operation")`.
4.  **Defer**: `defer span.End()`.
This creates spans that are automatically linked if the `ctx` is passed down the call stack.

---

## 753. How do you trace gRPC requests in Go?

**Answer:**
Use the `otelgrpc` interceptor.
`server := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))`.
`conn, _ := grpc.Dial(..., grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))`.
The interceptors automatically extract the Trace Context from gRPC Metadata, create a Child Span, and inject it back into outgoing calls, forming a complete distributed trace with zero manual code.

---

## 754. How do you record and export application traces?

**Answer:**
We use a **Span Exporter**.
Common exporters: **OTLP** (OpenTelemetry Protocol) -> Collector -> Datadog/Jaeger.
In development: `stdout` exporter.
Only sample a percentage (e.g., 1%) in production to save costs, unless it's an Error trace (tail-based sampling), which we always want to keep.

---

## 755. How do you handle slow endpoints in production Go apps?

**Answer:**
1.  **Observability**: Look at pprof and Trace span durations.
2.  **Timeout**: Ensure `http.Server.WriteTimeout` is set (e.g., 30s) so they don't hang forever.
3.  **Optimization**: Is it DB? (Add index). Is it CPU? (Optimize algorithm). Is it blocking upstream? (Add concurrency).

---

## 756. How do you add custom labels/tags to logs?

**Answer:**
In Zap/Slog, we create a child logger with context.
`reqLogger := logger.With("request_id", id, "user_role", role)`.
Pass `reqLogger` down to service functions.
Any log line written by `reqLogger.Info("done")` will automatically include `{request_id: ..., user_role: ...}`. this technique is called **Contextual Logging**.

---

## 757. How do you redact sensitive data in logs?

**Answer:**
1.  **Custom Marshaler**: Implement `MarshalLogObject` for structs. For the `Password` field, print `*****`.
2.  **Middleware scrubbing**: Search/Replace patterns (Credit Card Regex) in the log output buffer (slow, brittle).
3.  **Strict Types**: Use a special `SecretString` type that prints `***` when formatted with `%s` or JSON marshaled.

---

## 758. How do you detect memory leaks using Go tools?

**Answer:**
`go tool pprof`.
Compare two heap profiles (diff):
`go tool pprof -base base.prof current.prof`.
It shows the **delta**.
If `inuse_objects` is +50,000 for `NewSession` and -0 for `CloseSession`, you are leaking sessions.

---

## 759. How do you instrument performance counters in Go?

**Answer:**
Use Prometheus **Histogram** and **Summary**.
**Counter**: Total requests (Goes up).
**Gauge**: Queue depth (Up/Down).
**Histogram**: Request Duration (Buckets: 0.1s, 0.5s, 1s).
We wrap the handler:
`timer := prometheus.NewTimer(durationHist); defer timer.ObserveDuration()`.

---

## 760. How do you implement a tracing middleware?

**Answer:**
```go
func Tracer(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w, r) {
        ctx, span := tracer.Start(r.Context(), r.URL.Path)
        defer span.End()
        
        // Inject TraceID into Response Header for debugging
        w.Header().Set("X-Trace-ID", span.SpanContext().TraceID().String())
        
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```
This ensures every HTTP request starts a span and propagates the context to children.
