# 🟣 **421–440: Error Handling & Observability**

### 421. How do you create custom error types in Go?
"I define a struct that implements `error`.

`type MyError struct { Msg string; Code int }`.
`func (e *MyError) Error() string { return fmt.Sprintf("%d: %s", e.Code, e.Msg) }`.
This lets me attach metadata (like HTTP Status 404) to the error, which my middleware can extract later using `errors.As`."

#### Indepth
`errors.As` is the safe alternative to type assertion errors. Using `err.(*MyError)` panics if the error is nil or of a different type (and not wrapped). `errors.As(err, &target)` handles unwrap logic recursively and safe-guards against panics. Always use `As` for inspecting custom error fields.

---

### 422. How does Go 1.20+ `errors.Join` and `errors.Is` work?
"**errors.Join**: Combines multiple errors into one.
`err := errors.Join(err1, err2)`.
This is great for validation (returning 5 missing fields at once).

**errors.Is**: Checks if *any* error in the chain matches my target.
`if errors.Is(err, fs.ErrNotExist)`.
It unwraps the error tree automatically, so I don't need to manually check `err.Unwrap()`."

#### Indepth
`errors.Join` simply returns a type that implements `Unwrap() []error`. This triggers specific behavior in `errors.Is`: it checks *all* errors in the slice. Be careful: standard string formatting of joined errors uses newlines (`\n`), which might break single-line log parsers if you aren't careful.

---

### 423. How do you implement error wrapping and unwrapping?
"I use `%w` in `fmt.Errorf`.

`return fmt.Errorf("query failed: %w", err)`.
This wraps the original error.
To inspect the cause, I use `errors.Unwrap(err)` or `errors.Is`.
This preserves the *Root Cause* (e.g., 'DB Connection Lost') while adding *Context* (e.g., 'Could not fetch user'), so I can debug *why* it failed while telling the user *what* failed."

#### Indepth
Be careful with `fmt.Errorf("... %w", err)` inside a tight loop. It creates a linked list of error objects. If you wrap 1000 times, GC overhead increases. For deeply nested stack traces in hot paths, consider if you really need to wrap *every* step or just the boundaries (Access Layer -> Domain Layer).

---

### 424. What are best practices for error categorization?
"I define **Sentinel Errors** for broad categories.

`var ErrNotFound = errors.New("not found")`.
`var ErrPermission = errors.New("permission denied")`.
My Service Layer returns these.
My HTTP Handler checks them:
`if errors.Is(err, ErrNotFound) { return 404 }`.
This keeps my HTTP logic clean and decoupled from my Database logic."

#### Indepth
Don't use `errors.New` for dynamic errors! Sentinel errors should be immutable. If you need dynamic data (like "User 123 not found"), use a custom error type. Mixing the two (`return fmt.Errorf("%w: user %d", ErrNotFound, id)`) allows `errors.Is(err, ErrNotFound)` to still work while preserving context.

---

### 425. How do you handle critical vs recoverable errors?
"**Recoverable**: Transient issues (Network glitch, DB busy). I retry or degrade gracefully.
**Critical**: Invariant violations (Config missing, OOM). I `panic` or `log.Fatal` during startup.
I fail fast at boot, but I *never* crash a running request handler unless it's a catastrophic memory corruption."

#### Indepth
Go panics are not Exceptions. They are for **Programmer Errors** (index out of bounds, nil pointer dereference). Don't use panic for "File Not Found". The only valid use of `panic` in business logic is during `init()` (e.g., config parsing failed) where the app *cannot* start safely.

---

### 426. How do you recover from panics in goroutines?
"A panic in a goroutine kills the whole process.

So I wrap my goroutines:
`go func() { defer func() { if r := recover(); r != nil { log.Error("Panic!", r) } }() ; doWork() }()`.
Frameworks like **Gin** or **Echo** have `Recovery()` middleware that does this automatically for every HTTP request."

#### Indepth
The default stack trace from `recover()` is just text. If you want structured logging, you need to parse the stack trace (using `runtime.Callers`). Most automated frameworks do this. Also unexpected: `recover()` returns `nil` if there was no panic, so `if r := recover(); r != nil` is the only safe way to checks.

---

### 427. How to capture stack traces on error?
"Standard `error` has no stack trace.

I use `github.com/pkg/errors` (or standard with a helper).
`errors.WithStack(err)`.
When logging, I use `%+v`.
This prints the full trace: `main.go:42 -> service.go:10 -> db.go:5`.
Without this, debugging 'database error' is a guessing game."

#### Indepth
`pkg/errors` is deprecated (archived). The community standard is shifting to standard lib `errors` + some stack trace helper, or libraries like `gitlab.com/tozd/go/errors`. However, for legacy apps, `pkg/errors` is still rock solid. Just don't mix it blindly with `errors.Join`.

---

### 428. How do you notify Sentry/Bugsnag from Go?
"I use **Middleware hooks**.

In my HTTP `Recovery` middleware:
`if r := recover(); r != nil { sentry.CaptureException(r) }`.
I also use a custom `slog.Handler`.
Any log with level `ERROR` is automatically sent to Sentry. This ensures I don't miss any error, even if I forget to call `sentry.Capture` explicitly."

#### Indepth
Sentry grouping relies on the "Fingerprint". If you just send `err.Error()`, and the error contains a timestamp or ID (`"Duplicate entry 123"`), Sentry will create a *new issue* for every error! Always strip dynamic data / use a static error message for the fingerprint, or use structured logging fields.

---

### 429. How do you do structured error reporting in Go?
"I avoid string concatenation.

Bad: `log.Error("failed to update user " + id)`.
Good: `log.Error("failed to update user", "user_id", id, "error", err)`.
This outputs JSON.
In Kibana/Datadog, I can then query `error.user_id = "123"` regardless of the error message text."

#### Indepth
Go 1.21 introduced `log/slog`. It is faster than zap in many cases and standardizes the interface. Use `slog.Group("user", "id", 1, "role", "admin")` to nest JSON fields cleanly. This makes logs infinitely more queryable in backend systems like Loki or Elastic.

---

### 430. How do you correlate logs, errors, and traces together?
"**Trace ID** is the glue.

1.  Extract TraceID from Context.
2.  Add it to every Log line (`"trace_id": "abc"`).
3.  Add it to the Sentry Event tag.
4.  Add it to the OpenTelemetry Span.
This allows me to click a button in Sentry and jump instantly to the Jaeger trace showing *why* that error happened."

#### Indepth
Propagation is handled by `otel.GetTextMapPropagator().Inject(ctx, header)`. If you are calling a downstream service (even via HTTP), you *must* inject these headers manually if you aren't using an auto-instrumented client. Otherwise, the trace breaks at the service boundary.

---

### 431. How would you add distributed tracing to an existing Go service?
"I start at the edges.

1.  Add **Otel Middleware** to the HTTP Router.
2.  Add **Otel Interceptors** to `http.Client` and `grpc.Client`.
This gives me 90% of the value (Service Map and Latency) with zero code changes to the business logic.
Only then do I manually add `tracer.Start(ctx, "complex_calc")` to critical internal functions."

#### Indepth
Manual instrumentation: `ctx, span := tracer.Start(ctx, "op_name")` followed by `defer span.End()`. Always check `span.IsRecording()` before doing expensive work (like dumping a huge payload) to attach to the span attributes, to avoid overhead when tracing is sampled out.

---

### 432. What are tags, attributes, and spans in tracing?
"**Span**: A unit of work (e.g., 'DB Query'). Has a start and end time.
**Attribute (Tag)**: Metadata (`db.statement="SELECT..."`, `http.status=200`).
**Trace**: A tree of Spans.

I use attributes to filter: 'Show me all traces where `user_type=admin` and `duration > 500ms`'."

#### Indepth
Semantic Conventions! Don't make up attribute names. Use `semconv` packages (`go.opentelemetry.io/otel/semconv/v1.17.0`). Use `db.system` instead of `database_type`. This allows UI tools (Jaeger/Datadog) to auto-render fancy icons and categorize traffic correctly.

---

### 433. What is a traceparent header?
"It’s the W3C standard for trace propagation.

`traceparent: 00-{trace-id}-{span-id}-{flags}`.
My Go service reads this header to know 'I am part of Trace X, and my parent is Span Y'.
It ensures the trace continues unbroken as the request jumps from my Load Balancer -> Go -> Python -> Database."

#### Indepth
There is also `baggage` header. It carries key-value pairs (`userid=123`) across the *entire* trace (not just parent-child). Use it sparingly! If you put 1KB of data in baggage, you are sending 1KB extra header on *every* internal microservice call. Limits are usually strict (4KB/8KB).

---

### 434. How do you send custom metrics to Prometheus?
"I define specific collectors.

`var activeUsers = promauto.NewGaugeVec(...)`.
In my code: `activeUsers.WithLabelValues("US").Inc()`.
Prometheus scrapes my `/metrics` endpoint.
**Trap**: I carefully manage **Cardinality**. I never use 'UserID' or 'Email' as a label, or I'll explode my Prometheus memory usage."

#### Indepth
The "Metric Explosion" problem. `http_requests_total{path="/users/123"}` -> 1 million metrics for 1 million users. Prometheus creates a new time series for every unique label combination. Always normalize: `path="/users/:id"`. Use logs for high-cardinality details, metrics for aggregates.

---

### 435. What is RED metrics model and how do you apply it?
"It’s the Golden Signal set for Microservices.

**R**ate: Requests per second (`http_requests_total`).
**E**rrors: Failed requests per second (`http_requests_total{status=5xx}`).
**D**uration: Latency (`http_request_duration_seconds` Histogram).
I ensure every service exposes these three. If Error Rate spikes or Duration P99 goes up, I page the on-call engineer."

#### Indepth
Also consider **Saturation** (the 4th Golden Signal). How "full" is my service? Thread pool usage, Memory usage, File Descriptor usage. RED tells you if you are failing; Saturation tells you *if you are about to fail*. Monitor `conn_pool_open_connections` vs `max_connections`.

---

### 436. How do you expose application health and readiness probes?
"I use two endpoints.

`/live`: Returns 200 OK immediately (Is the process running?).
`/ready`: Returns 200 OK only if waiting for dependency checks (DB connected, Cache warm).
In K8s, I use `readinessProbe` to prevent traffic from hitting a pod that is technically 'up' but not yet ready to serve traffic."

#### Indepth
`readinessProbe` failures remove the pod from the Load Balancer. `livenessProbe` failures **Restart** the pod. Do *not* check the Database in your Liveness probe! If the DB goes down, all your pods will fail liveness and restart simultaneously, causing a crash loop and potentially hammering the recovering DB. Keep liveness simple (e.g., "Main thread not dead").

---

### 437. What’s the difference between logs, metrics, and traces?
"**Logs**: Detailed events ('User 123 clicked Buy'). Expensive to store.
**Metrics**: Aggregated numbers ('Order Count = 50'). Cheap to store, great for alerts.
**Traces**: Latency analysis ('DB took 4s'). Great for debugging slowness.

I need all three. Metrics tell me *something* is wrong. Traces tell me *where*. Logs tell me *why*."

#### Indepth
**Exemplars** (OpenMetrics) link Metrics to Traces. In a histogram bucket "Latency 1s-2s", Prometheus can store a "TraceID" of a specific request that fell into that bucket. This is the holy grail: spotting a spike in a graph -> clicking a dot -> seeing the exact trace. Go's Prometheus client supports this.

---

### 438. How do you benchmark error impact on performance?
"I write a Benchmark.

`Algorithm A` returns `nil`.
`Algorithm B` constructs and returns `fmt.Errorf(...)`.
I'll find that *creating* errors with stack traces allows is slow (allocations).
So I avoid using errors for **Control Flow** (like 'End of Loop'). Errors should be exceptional."

#### Indepth
Stack traces are the heavy part. Creating a simple error `errors.New("fail")` is cheap (just an allocation). `pkg/errors.New("fail")` captures the PC (Program Counter) for every frame. In a tight inner loop (parsing a million lines), avoid errors with stacks. Return bools or specialized error codes.

---

### 439. What’s the tradeoff between verbose and silent error handling?
"**Verbose**: logs everything. Risk: Disk full, signal noise.
**Silent**: ignores them. Risk: Flying blind.

**Balance**: I only log errors at the **Edge** (HTTP Handler) or when I *handle* them (swallow them).
If I return an error up the stack, I do *not* log it. This prevents the 'Log Scraper' pattern where one error appears 10 times in the logs."

#### Indepth
The "Error return" pattern (`if err != nil { return err }`) ensures that eventually, *someone* handles it. If you log at every level, you get: "Error DB", "Error Handler", "Error Main". Just return the error wrapped with context, and let the top-level handler log it *once*, fully formed.

---

### 440. How would you enforce observability in a Go microservice?
"I use a **Service Chassis** (Template).

A shared library `mycompany/service`.
It initializes Slog, Prometheus, and Otel automatically in `service.Run()`.
This implies every new microservice gets standard metrics, tracing, and logging for free, without the developer needing to configure it manually."

#### Indepth
Observability as Code. Standardize the `logger` constructor. If every team uses their own logger format, you can't build global dashboards. Enforce a shared library that sets up `slog.SetDefault()`, `otel.SetTracerProvider()`, and `promhttp.Handler()` with consistent naming/namespace conventions.
