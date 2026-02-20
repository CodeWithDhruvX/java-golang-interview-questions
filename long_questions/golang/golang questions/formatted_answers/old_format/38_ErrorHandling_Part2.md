# 📦 **741–760: Error Handling & Observability (Part 2)**

### 741. How do you implement a custom error type in Go?
"I define a struct implementing the `Error() string` interface.
`type ValidationError struct { Field, Reason string }`.
`func (e *ValidationError) Error() string { return e.Field + ": " + e.Reason }`.
I usually add an `Unwrap() error` method so `errors.Is` works on the underlying cause."

#### Indepth
**Pointer Receivers**. Always define error methods on the *pointer* receiver (`*ValidationError`), not the value. If you use value receiver, `errors.As(err, &target)` might fail or panic because it relies on reflection to detecting if `*target` implements `error`. standard practice: `func (e *MyErr) Error() string`.

---

### 742. How do you wrap errors in Go?
"I use `%w` in `fmt.Errorf`.
`return fmt.Errorf("failed to open file: %w", err)`.
This creates a wrapped error.
I can then use `errors.Unwrap(err)` to access the original error, or `errors.Is` to check for specific root causes (like `io.EOF`) despite the wrapping."

#### Indepth
**Opaque Errors**. Wrapping potentially exposes implementation details (like "sql: no rows"). If you want to *hide* the details from the caller (forcing them to handle only "UserNotFound"), don't wrap. Just return the sentinel. Wrapping is for *adding context* ("failed to get user: [cause]"), not just passing the buck.

---

### 743. What is `errors.Is()` and `errors.As()` used for?
"**`errors.Is(err, target)`**: Checks if `err` matches a specific sentinel value (like comparing by value).
**`errors.As(err, &target)`**: Checks if `err` matches a specific *type* and assigns it to `target`.
Example: `var vErr *ValidationError; if errors.As(err, &vErr) { // access vErr.Field }`."

#### Indepth
**Error Chains**. Both functions traverse the *entire* chain of wrapped errors (the tree). `errors.Is` is generally faster (value comparison). `errors.As` uses reflection and is slower. Prefer `errors.Is` for control flow (sentinels) and `errors.As` only when you need to extract data properties from the error.

---

### 744. How do you categorize errors in large Go applications?
"I define **Sentinel Errors** in my domain package.
`var ErrNotFound = errors.New("not found")`.
`var ErrUnauthorized = errors.New("unauthorized")`.
The implementation layer (DB) wraps internal SQL errors into these domain errors.
The HTTP layer switches on these domain errors to return 404 or 401, keeping layers decoupled."

#### Indepth
**Behavior Interface**. Instead of switching on types (Coupling), define interfaces. `type NotFounder interface { NotFound() bool }`. The HTTP layer checks `if e, ok := err.(NotFounder); ok && e.NotFound() { 404 }`. This allows any package to define a "Not Found" error without importing a central "Errors" package.

---

### 745. How do you log structured errors in Go?
"I use `slog` or `zap`.
`logger.Error("payment failed", "amount", 100, "error", err)`.
The output `json` contains `{"level":"error", "msg":"payment failed", "amount":100, "error":"timeout"}`.
This allows me to query logs by field (`amount > 50`) in Datadog."

#### Indepth
**Sampling**. In high-throughput systems, logging *every* error might kill your IO. Use **Sampling**. Log 100% of "Critical" errors, but only 1% of "Debug" logs. `zap` supports sampling configuration out of the box. This keeps costs down while still providing statistical visibility.

---

### 746. How do you use Sentry/Bugsnag with Go?
"I initialize the SDK in `main()`.
I create a deferred recovery middleware.
`defer func() { if r := recover(); r != nil { sentry.CaptureMessage(fmt.Sprint(r)) } }()`.
For handled errors, I manually call `sentry.CaptureException(err)` if it's something unexpected (like 500 Internal Server Error)."

#### Indepth
**Source Maps**. Compiled Go binaries don't look like code. Sentry needs your source code to show helpful context. You can embed source code in the binary (Go 1.18+) or upload the source code to Sentry during the build process to get clickable stack traces in the UI.

---

### 747. How do you implement centralized error logging?
"I create a global `ErrorHandler` middleware.
Every handler returns `error`.
The middleware catches it.
1.  Logs it to Sentry/Stdout with Request ID.
2.  Determines HTTP Status Code.
3.  Writes JSON response to user.
This ensures no error is ever silently swallowed."

#### Indepth
**panic(http.Abort)**. Some frameworks (Gin) allow you to `panic(err)` and catch it in middleware. **Don't do this**. It destroys the stack trace of where the error *actually* occurred (replaced by the panic location). Always return errors explicitly up the stack to the middleware.

---

### 748. What is the role of stack traces in debugging Go apps?
"Go errors don't have stack traces by default.
I use `github.com/pkg/errors` to wrap them: `errors.Wrap(err, "context")`.
Or standard `errors` + a logger that prints stack traces.
When I see a log, I need the *path* (Controller -> Service -> Repo) to reproduce the bug."

#### Indepth
**Cost of Stack Traces**. Generating a stack trace (`runtime.Stack`) is expensive (stops execution, walks stack). Don't add stack traces to *every* error (like "User not found"). Only add them for "Unexpected/System" errors (e.g., DB connection died). `pkg/errors` adds the stack trace at the point of `Wait` or `New`.

---

### 749. How do you implement panic recovery with context?
"In my recovery middleware, I extract the Request Context.
`defer func() { if r := recover(); r != nil { log.Error("Panic", "path", r.URL.Path, "user", userFromCtx(ctx)) } }()`.
Knowing *who* triggered the panic and *which* endpoint is often enough to find the bug immediately."

#### Indepth
**Named Return Parameters**. If you want to recover from a panic and *return an error* to the caller (instead of crashing), you MUST use named return parameters. `func Safe() (err error) { defer func() { if r := recover(); r != nil { err = fmt.Errorf("panic: %v", r) } }() }`. This modifies the return value `err` even after the function "stopped".

---

### 750. How do you differentiate retryable vs fatal errors?
"I implement an interface `Retryable { Temporary() bool }`.
My network errors implement this.
If `err.Temporary()`, I retry with backoff.
If not (e.g., JSON Syntax Error), I fail immediately.
If I'm unsure, I distinguish by HTTP codes: 50x (Retry), 40x (Fail)."

#### Indepth
**Idempotency Key**. Retrying non-idempotent operations (like "Charge $10") is dangerous. Always include an `Idempotency-Key` header (UUID) in the request. The server checks redis: "Did I already process Key X?". If yes, return previous result. This makes retries safe.

---

### 751. How do you expose Prometheus metrics in Go?
"I use `prometheus/client_golang`.
`http.Handle("/metrics", promhttp.Handler())`.
I define globals: `var reqCount = promauto.NewCounter(...)`.
I increment them in business logic.
Prometheus scrapes the endpoint every 15s. It’s the standard for cloud-native Go apps."

#### Indepth
**Push vs Pull**. Prometheus is "Pull" (scraper calls you). If running a batch job that runs for 5 seconds and dies, the scraper might miss it. For short-lived jobs, use a **Pushgateway**. The Go job pushes metrics to the Gateway, and Prometheus scrapes the Gateway.

---

### 752. How do you set up OpenTelemetry in Go?
"1.  Init `TracerProvider` (sending to Jaeger/Otlp).
2.  Use `otelhttp` middleware to trace HTTP requests.
3.  In code: `ctx, span := tracer.Start(ctx, "op")`.
It unifies Traces, Metrics, and Logs under one standard SDK."

#### Indepth
**Propagators**. How does trace context jump from Service A to Service B? HTTP Headers (`traceparent`). You must configure the `TextMapPropagator` in OTel. Without this, Service B starts a *new* trace instead of continuing Service A's trace, breaking the distributed view.

---

### 753. How do you trace gRPC requests in Go?
"I add the `otelgrpc.UnaryServerInterceptor` to my gRPC server options.
It automatically creates a Span for each RPC call.
It extracts the `traceparent` metadata from headers, linking the client's trace to the server's trace (Distributed Tracing)."

#### Indepth
**Baggage**. OTel allows carrying "Baggage" (KV pairs) alongside the trace. e.g., `UserID=123`. This propagates to *all* downstream services automatically. Service D can log "UserID=123" without Service A explicitly passing it in the gRPC body. Use carefully (header size limits).

---

### 754. How do you record and export application traces?
"I configure an **Exporter**.
`exporter, _ := jaeger.New(...)`.
I register it with the Trace Provider.
My app buffers spans and batches them via UDP/HTTP to the collector.
I verify to call `provider.Shutdown(ctx)` on exit to flush buffered traces."

#### Indepth
**Sampling Strategy**. You can't trace 100% of requests in production (too much data). Use `TraceIDRatioBased(0.01)` (1%). Or `ParentBased`, which respects the incoming sampling decision (if the caller traced it, I trace it). This ensures you get full traces for the 1% of sampled requests.

---

### 755. How do you handle slow endpoints in production Go apps?
"I use **Profiling** and **Tracing**.
Tracing shows *which* part is slow (DB? External API?).
Profiling (`pprof`) shows *why* (CPU loop? Lock contention?).
I also use `net/http/pprof` specifically on the live pod to take a 30s sample."

#### Indepth
**Block Profile**. Often the CPU is low, but the app is slow. This means it's waiting on locks or IO. `runtime.SetBlockProfileRate(1)`. Then look at `/debug/pprof/block`. It shows exactly where goroutines are waiting. Essential for debugging mutex contention.

---

### 756. How do you add custom labels/tags to logs?
"I use `slog.With`.
`logger = logger.With("service", "billing", "env", "prod")`.
Every log line from this logger instance will have those tags.
I pass this logger down to sub-components."

#### Indepth
**Context Logger**. Passing `logger` explicitly to every function is hideous. Storing it in `context.Context` is controversial but common. `log := logger.FromContext(ctx)`. Using `slog.Default()` is nicer, but you lose the "request-scoped" fields (request_id) unless you use a context-aware handler.

---

### 757. How do you redact sensitive data in logs?
"I implement the `LogValuer` interface for sensitive structs.
`func (u User) LogValue() slog.Value { return slog.StringValue("User{ID=" + u.ID + ", Pass=*** }") }`.
Or I use a middleware that scans for keys like `password` and replaces values with `[REDACTED]`."

#### Indepth
**PII Scanning**. Hard filters (`password`) miss things. Better: Don't log struct dumps. Explicitly log fields: `log.Info("user login", "user_id", u.ID)`. Never `log.Info("user", u)`. Whitelisting fields is safer than blacklisting.

---

### 758. How do you detect memory leaks using Go tools?
"**Heap Profile**: `go tool pprof -sample_index=inuse_space heap.out`.
I compare two profiles: `pprof -base base.out current.out`.
If `inuse_space` is growing for a specific function, I check if it's retaining pointers in a global map or leaking goroutines."

#### Indepth
**Goroutine Leaks**. The #1 cause of memory leaks in Go is not memory, but *stuck goroutines*. Each consumes 2KB+ stack. Use `goleak` in your tests to assert that every test finishes with 0 helper goroutines running. A leaked listener goroutine will prevent the entire server struct from being GC'd.

---

### 759. How do you instrument performance counters in Go?
"I use atomic integers for high-speed unchecked counters.
`atomic.AddInt64(&ops, 1)`.
For monitoring, I prefer Prometheus **Histograms** to track latency distribution (p95, p99), not just averages."

#### Indepth
**High Cardinality**. Counters are cheap. Histograms are expensive (they create 10+ time series per metric bucket). Do not put "UserID" or "IP" in histogram labels. Use `Summary` if you need exact client-side quantiles, but `Histogram` is better for aggregation across multiple pods.

---

### 760. How do you implement a tracing middleware?
"1. Start Span.
2. Add TraceID to Response Header.
3. `next.ServeHTTP`.
4. Record Status Code and Duration in Span.
5. End Span.
This gives me visibility into every incoming HTTP request's duration and result."

#### Indepth
**Response Capture**. Standard `http.ResponseWriter` is write-only. You cannot read the status code back after writing it. You must wrap it: `type loggingResponseWriter struct { http.ResponseWriter; statusCode int }`. Overload `WriteHeader` to capture the int. Pass this wrapper to `next.ServeHTTP`.
