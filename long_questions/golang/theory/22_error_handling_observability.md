# 🟢 Go Theory Questions: 421–440 Error Handling & Observability

## 421. How do you create custom error types in Go?

**Answer:**
We define a struct that implements the `error` interface (`Error() string`).

`type NotFoundError struct { ID string }`
`func (e *NotFoundError) Error() string { return "not found: " + e.ID }`

This allows us to carry **Structured Data** (like the ID or UserID) inside the error, rather than just a string. This is vital for upper layers to make decisions (e.g., "If it's a NotFoundError, return 404; if it's a DBError, return 500").

---

## 422. How does Go 1.20+ `errors.Join` and `errors.Is` work?

**Answer:**
`errors.Join(err1, err2)` creates a **Multi-Error**. It bundles multiple errors into one parent error.
`errors.Is(err, target)` recursively unwraps the error tree to see if *any* error in the chain matches the target.

This is huge for parallel tasks. If you launch 10 goroutines and 3 fail, you can use `errors.Join` to return all 3 errors to the caller, and the caller can check `errors.Is(err, sql.ErrNoRows)` to detect specific failure modes within the aggregate.

---

## 423. How do you implement error wrapping and unwrapping?

**Answer:**
We use `%w` in `fmt.Errorf`.

`return fmt.Errorf("query failed: %w", err)`
This wraps the original error. The result preserves the underlying type.

To unwrap, we use `errors.Unwrap(err)`, but usually we use `errors.As(err, &target)` to find a specific custom error type deep in the chain. This preserves the "Stack Trace" of context ("query failed") while keeping the root cause (DB disconnected) inspectionable.

---

## 424. What are best practices for error categorization?

**Answer:**
We categorize errors by **Behavior**, not just Type.

We define interfaces or helper methods:
`IsTemporary(err) bool` -> Retry.
`IsInput(err) bool` -> Return 400.
`IsSystem(err) bool` -> Return 500 and Alert Op.

We often use a central `apperrors` package that maps low-level errors (postgres duplicate key) to high-level domains (UserAlreadyExists), decoupling the HTTP layer from the DB layer.

---

## 425. How do you handle critical vs recoverable errors?

**Answer:**
**Recoverable** errors (File missing, Network timeout) are returned as values (`error`). The caller decides to retry or log.

**Critical** errors (Programmer bugs, Nil pointer dereference, Memory corruption) cause a `panic`.
We generally catch panics at the top level (Middleware) to prevent crashing the whole server, logging the stack trace, and returning a 500. We *never* use panic for normal control flow.

---

## 426. How do you recover from panics in goroutines?

**Answer:**
A panic in a goroutine kills the **entire application**, not just that goroutine.

To prevent this, every long-running goroutine must have a `defer recover()` block.
```go
go func() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Recovered: %v", r)
        }
    }()
    doWork()
}()
```
We often wrap this in a `SafeGo(fn)` utility function to DRY up this pattern.

---

## 427. How to capture stack traces on error?

**Answer:**
Standard `fmt.Errorf` doesn't capture stack traces.
We use `github.com/pkg/errors` (or a modern equivalent like `cockroachdb/errors`).

`errors.New("boom")` captures the file and line number.
`fmt.Printf("%+v", err)` prints the full stack trace.
In Go 1.20+, while `errors.Join` exists, standard errors still lack stack traces, so 3rd party libraries or custom wrappers are still the standard for production observability.

---

## 428. How do you notify Sentry/Bugsnag from Go?

**Answer:**
We use a **Hook** mechanism in our Logger or a dedicated Middleware.

When an error is logged at level `Error` or `Fatal`, the hook fires.
`sentry.CaptureException(err)`.
Crucially, we attach `Context` to the event, including `User ID`, `IP Address`, and `Version`. This allows us to debug: "This crash only happens for User 123 on Version 2.0."

---

## 429. How do you do structured error reporting in Go?

**Answer:**
We avoid returning raw strings like "failed to connect".
We return a struct (or JSON response) with codes.

`{ "code": "PAYMENT_FAILED", "message": "Card declined", "trace_id": "abc-123" }`

In Go, we define an interface `ReportableError` with a `Code() string` method. The HTTP middleware checks this interface. If implemented, it uses that code; otherwise, it defaults to `INTERNAL_SERVER_ERROR`.

---

## 430. How do you correlate logs, errors, and traces together?

**Answer:**
**Trace ID** is the glue.

1.  **Logs**: Every log line has `trace_id`.
2.  **Errors**: Sentry report includes `trace_id` tag.
3.  **Traces**: Jaeger stores the timing data under `trace_id`.

When we get a Sentry alert, we copy the Trace ID, paste it into Jaeger to see *what* was slow, and look it up in Kibana/Loki to see the raw logs, giving a 360-degree view of the failure.

---

## 431. How would you add distributed tracing to an existing Go service?

**Answer:**
We use **OpenTelemetry Auto-Instrumentation** where possible, or manual hooks.

1.  Initialize OTel Exporter (sending to Jaeger/Tempo).
2.  Wrap HTTP Handler: `otelhttp.NewHandler(handler, "api")`.
3.  Wrap HTTP Client: `http.Client{Transport: otelhttp.NewTransport(...) }`.
4.  Pass `context.Context` EVERYWHERE.

This automatically creates spans for every incoming request and outgoing call, linking them up.

---

## 432. What are tags, attributes, and spans in tracing?

**Answer:**
**Span**: A unit of work (e.g., "SQL Query", "HTTP Request"). It has a start/end time.
**Attribute (Tag)**: Key-value metadata on a span. `db.statement="SELECT * FROM users"`, `http.status_code=200`.
**Event**: A timestamped log inside a span (e.g., "Retrying connection").

We query traces by attributes: "Show me all spans where `http.status_code=500` and `service=billing`".

---

## 433. What is a traceparent header?

**Answer:**
It's the W3C standard HTTP header for Trace Context propagation.
Format: `00-{trace-id}-{parent-id}-{flags}`.

When Service A calls Service B, it injects `traceparent`.
Service B reads it and starts its own span with `parent-id` set to Service A's span.
This links the two services into a single directed graph (Trace) without them needing to know about each other specifically.

---

## 434. How do you send custom metrics to Prometheus?

**Answer:**
1.  Define a metric global: `var opsProcessed = promauto.NewCounter(prometheus.CounterOpts{Name: "myapp_ops_total"})`.
2.  Instrument code: `opsProcessed.Inc()`.
3.  Expose `/metrics` endpoint.

We prefer **Counter** (cumulative total) and **Histogram** (distribution of latency). We avoid **Gauge** for things that change rapidly unless they are state snapshots (like "Memory Used" or "Active Goroutines").

---

## 435. What is RED metrics model and how do you apply it?

**Answer:**
RED stands for **Rate, Errors, Duration**. It is the golden standard for monitoring microservices.

*   **Rate**: Request per second (Counter).
*   **Errors**: Failed requests per second (Counter).
*   **Duration**: How long requests take (Histogram).

For every API endpoint in Go, our middleware automatically emits these three metrics. This allows us to build a standardized Grafana dashboard that works for *any* service.

---

## 436. How do you expose application health and readiness probes?

**Answer:**
**Liveness (`/healthz`)**: "Am I running?" (Returns 200 OK). Use for restarting dead pods.
**Readiness (`/readyz`)**: "Can I serve traffic?". Checks DB connection, Redis, etc.

In Go, we perform a lightweight "Ping" to dependencies. If DB is down, `/readyz` returns 503. K8s removes the pod from the Load Balancer but leaves it running (to recover). If we fail `/healthz`, K8s kills the pod.

---

## 437. What’s the difference between logs, metrics, and traces?

**Answer:**
*   **Logs**: "What happened?" (High fidelity, expensive). "User 123 bought Item X".
*   **Metrics**: "What is happening?" (Aggregated, cheap). "Sales/min = 50".
*   **Traces**: "Where did it happen?" (Causal, sampled). "Checkout took 500ms (DB=400ms)".

We use Metrics for **Alerting** (Is site down?).
We use Traces for **Debugging Latency**.
We use Logs for **Root Cause Analysis**.

---

## 438. How do you benchmark error impact on performance?

**Answer:**
Creating an `error` in Go involves allocations (if using formatted strings).
Stack traces (runtime.Callers) are very expensive.

We write a `Benchmark` test in `_test.go`:
`func BenchmarkError(b *testing.B) { for i:=0; i<b.N; i++ { _ = fmt.Errorf("bad") } }`
We compare this against a sentinel error (`return ErrBad`). Sentinel errors are zero-allocation. If a hot loop returns errors frequently (e.g., "EOF"), use Sentinels.

---

## 439. What’s the tradeoff between verbose and silent error handling?

**Answer:**
**Verbose** (Returning every error detail): Great for debugging, but risks **Information Leakage** (exposing table names/IPs to clients).
**Silent** ("Something went wrong"): Secure, but impossible to debug.

**Balance**: Log Verbose, Return Silent.
We log the full error with cause on the server. We return a generic UUID ("Error Ref: 5A23B") to the client. The client gives support the UUID, and we look up the verbose log.

---

## 440. How would you enforce observability in a Go microservice?

**Answer:**
We use **Shared Libraries** (Chassis pattern) or **Service Mesh**.

We don't trust developers to remember to log.
We write a `platform-go` library that provides a standard `NewServer()` function.
This server comes pre-wired with:
1.  Request Logging Middleware.
2.  Prometheus Exporter.
3.  OpenTelemetry Tracing.
4.  Panic Recovery.

Developers just write business handlers; the observability scaffolding is mandatory and uniform.
