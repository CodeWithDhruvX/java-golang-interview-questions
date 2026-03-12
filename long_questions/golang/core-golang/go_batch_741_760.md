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

### Explanation
Custom error types in Go are created by defining a struct and implementing the error interface's Error() string method. This allows structured error information with additional fields like error codes, making errors more informative and easier to handle programmatically.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a custom error type in Go?
**Your Response:** "I implement custom error types by defining a struct with relevant fields and implementing the `Error() string` method. For example, I might create an `APIError` struct with fields for error code and message. The Error() method formats these into a readable string. This approach gives me structured errors that I can inspect programmatically - I can check error codes to determine specific handling logic while still having human-readable messages. Custom error types are especially useful in APIs where I need to return specific error codes, or in domain logic where different error types require different handling. The key is that any type with an Error() string method implements the error interface, making it compatible with standard error handling patterns."

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

### Explanation
Error wrapping in Go uses the %w verb with fmt.Errorf to preserve the original error in a wrapper chain. This enables errors.Is and errors.As to inspect the underlying cause while adding context at each layer, maintaining error provenance through the call stack.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you wrap errors in Go?
**Your Response:** "I wrap errors using the `%w` verb with `fmt.Errorf`. When I have an error from a lower layer, I wrap it with additional context like 'failed to open config: %w' followed by the original error. The `%w` verb is special because it preserves the original error in a wrapper chain, unlike `%v` which just converts to string. This allows me to use `errors.Is` to check if the wrapped error matches a specific type, and `errors.As` to extract the original error with its type intact. Error wrapping is crucial for maintaining context as errors bubble up through layers - I add relevant context at each layer while preserving the original cause for proper error handling at the top level."

---

### Question 743: What is `errors.Is()` and `errors.As()` used for?

**Answer:**
- **`errors.Is(err, target)`**: Checks if `err` (or any error in its wrapper chain) matches a specific sentinel value (e.g., `io.EOF`).
- **`errors.As(err, &target)`**: Checks if `err` chain contains a specific **type** (e.g., `*net.OpError`) and assigns it to `target` so you can access fields.

### Explanation
errors.Is checks if an error or any error in its wrapper chain matches a specific sentinel value, while errors.As checks if the chain contains a specific error type and assigns it to a target variable for field access. Both functions traverse the error wrapper chain created with %w.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is `errors.Is()` and `errors.As()` used for?
**Your Response:** "I use `errors.Is` and `errors.As` to inspect wrapped errors. `errors.Is(err, target)` checks if an error or any error in its wrapper chain matches a specific sentinel value like `io.EOF`. This is perfect for checking for known error conditions regardless of how many layers of wrapping have been added. `errors.As(err, &target)` is more powerful - it checks if the error chain contains a specific type and, if found, assigns it to my target variable so I can access its fields. For example, I can use `errors.As` to extract a `*net.OpError` to inspect network-specific details. Both functions traverse the entire error wrapper chain created with the `%w` verb, making them essential for proper error handling in modern Go applications."

---

### Question 744: How do you categorize errors in large Go applications?

**Answer:**
Use Error Types or Codes.
Define a base `AppError` with an `Op` (Operation) and `Kind` (Permission, NotFound, Internal).
This decouples the messiness of low-level errors (SQL/Network) from the HTTP Status Code logic.

### Explanation
Error categorization in large Go applications uses error types or codes with a base AppError containing operation and kind fields. This decouples low-level error details from HTTP status code logic, providing clean separation between technical errors and business logic error handling.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you categorize errors in large Go applications?
**Your Response:** "I categorize errors using error types or codes with a base `AppError` structure. The AppError contains fields like `Op` for the operation that failed and `Kind` for the error category like Permission, NotFound, or Internal. This approach decouples the messy low-level errors - SQL errors, network timeouts, file system issues - from the HTTP status code logic. Instead of sprinkling error handling throughout my application, I map low-level errors to standardized application errors with consistent categories. This makes it much easier to handle errors consistently across different layers and provides clean separation between technical implementation details and business logic error handling. It also makes logging and monitoring much more systematic."

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

### Explanation
Structured error logging in Go uses slog (Go 1.21+) or zap for JSON-formatted logs with key-value pairs. This enables machine-readable logs that are easily searchable and analyzable in logging systems, replacing traditional unstructured text logs.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you log structured errors in Go?
**Your Response:** "I log structured errors using Go's built-in `slog` package or third-party `zap` library. Instead of unstructured text messages, I use key-value pairs: `logger.Error('login failed', slog.String('user', 'bob'), slog.Any('error', err))`. This outputs structured JSON that's machine-readable and easy to search in logging systems. Structured logging makes it much easier to filter logs by specific fields, analyze error patterns, and integrate with monitoring tools. I can search for all failed logins for a specific user, or track error rates by error type. This approach is much more powerful than traditional text-based logging, especially in production environments where I need to quickly diagnose issues across thousands of log entries."

---

### Question 746: How do you use Sentry/Bugsnag with Go?

**Answer:**
Initialize client -> Defer Flush -> Capture.
Use `sentry-go` SDK.
**Middleware:** `sentryhttp` middleware automatically captures panics and sends them to Sentry.

### Explanation
Sentry/Bugsnag integration in Go uses the sentry-go SDK with initialization, deferred flushing, and error capture. The sentryhttp middleware automatically captures panics and HTTP errors, sending them to Sentry for centralized error monitoring and alerting.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use Sentry/Bugsnag with Go?
**Your Response:** "I integrate Sentry using the `sentry-go` SDK following a three-step pattern: initialize the client with configuration like DSN and environment, defer flushing to ensure all events are sent before shutdown, and capture errors or panics. I also use the `sentryhttp` middleware which automatically captures panics and HTTP errors, sending them to Sentry with request context. This gives me centralized error monitoring with rich context including stack traces, user information, and request details. The middleware handles the boilerplate of capturing panics and unhandled exceptions, so I get comprehensive error reporting without manual instrumentation. Sentry helps me track production issues in real-time and get notified about critical errors affecting users."

---

### Question 747: How do you implement centralized error logging?

**Answer:**
Log to `stdout` locally.
In Production, the container runtime (Docker/K8s) captures stdout.
**Fluentd / Filebeat** runs as a daemon, tails the logs, and pushes them to **Elasticsearch / Loki**.

### Explanation
Centralized error logging in Go logs to stdout locally, with container runtimes capturing stdout in production. Fluentd/Filebeat daemons tail logs and push them to centralized systems like Elasticsearch or Loki for aggregation and analysis.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement centralized error logging?
**Your Response:** "I implement centralized error logging by following cloud-native best practices. Locally, I log to stdout for simplicity. In production, container runtimes like Docker or Kubernetes automatically capture stdout. I run Fluentd or Filebeat as daemon processes that tail these logs and push them to centralized systems like Elasticsearch or Loki. This approach follows the twelve-factor app methodology and works seamlessly with container orchestration platforms. The logging agents handle log rotation, parsing, and shipping, while my application just writes structured JSON to stdout. This gives me centralized log aggregation, search capabilities, and retention policies without complex logging infrastructure in my application code. It's the standard pattern for modern microservices architectures."

---

### Question 748: What is the role of stack traces in debugging Go apps?

**Answer:**
Essential for panics.
Standard `error` does not have stack straces.
Use `github.com/pkg/errors` (Wrap) so that when you log, you can print `%+v` to see the exact line number where the error originated, not just "file not found".

### Explanation
Stack traces in Go debugging are essential for panics but not included in standard errors. The github.com/pkg/errors package with Wrap() adds stack traces, and printing with %+v shows exact line numbers where errors originated, providing precise debugging information.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the role of stack traces in debugging Go apps?
**Your Response:** "Stack traces are crucial for debugging panics in Go, but standard errors don't include stack trace information. When I need detailed debugging information, I use the `github.com/pkg/errors` package and its `Wrap()` function to add stack traces to errors. When I log these wrapped errors with `%+v`, I get the exact line number where the error originated, not just generic messages like 'file not found'. This makes debugging much easier because I can see the complete call path leading to the error. While Go's built-in error handling is great for flow control, adding stack traces gives me the forensic information needed to diagnose complex issues in production. This approach bridges the gap between Go's simple error model and the detailed debugging information needed for troubleshooting."

---

### Question 749: How do you implement panic recovery with context?

**Answer:**
Recovered panics lose context (request ID, user).
In your recovery middleware:
1.  Read values from `r.Context()`.
2.  Include them in the log message when `recover()` catches a crash.
3.  Responds with 500 Internal Server Error.

### Explanation
Panic recovery with context in Go requires reading values from request context before recovery, including them in log messages when recover() catches crashes, and responding with 500 errors. This preserves request context like request ID and user information that would otherwise be lost.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement panic recovery with context?
**Your Response:** "I implement panic recovery with context by creating middleware that preserves request information before recovery. When a panic occurs, the recovered panic loses context like request ID and user information. In my recovery middleware, I first read values from `r.Context()` like request ID, user ID, and other relevant context. Then when `recover()` catches a panic, I include this context in the log message and respond with a 500 Internal Server Error. This approach ensures that even when my application crashes, I have the necessary context to trace the issue back to specific requests and users. Without preserving this context, debugging production panics would be much harder because I wouldn't know which request triggered the crash. This pattern is essential for building robust web services that can recover gracefully from panics."

---

### Question 750: How do you differentiate retryable vs fatal errors?

**Answer:**
Define a `Temporary() bool` interface.
If `net.Error.Temporary() == true` (timeout), retry.
If DB returns "Unique Constraint Violation", do not retry (Fatal).

### Explanation
Differentiating retryable vs fatal errors in Go uses a Temporary() bool interface. Network errors with Temporary() true (like timeouts) are retryable, while database constraint violations are fatal and shouldn't be retried to avoid repeated failures.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you differentiate retryable vs fatal errors?
**Your Response:** "I differentiate retryable from fatal errors by implementing a `Temporary() bool` interface or using existing ones like `net.Error`. For network errors, if `net.Error.Temporary() == true`, it's likely a timeout or transient issue that I should retry. For database errors like 'Unique Constraint Violation', these are fatal errors that won't be fixed by retrying - retrying would just cause the same error repeatedly. I implement this logic in my retry mechanisms: check if the error implements Temporary() and only retry if it returns true. For fatal errors, I fail fast and return the error to the caller. This approach prevents endless retry loops on permanent failures while still being resilient to temporary network issues. The key is understanding which errors are transient versus permanent and handling each appropriately."

---

### Question 751: How do you expose Prometheus metrics in Go?

**Answer:**
Use `github.com/prometheus/client_golang`.
1.  Define `var opsProcessed = promauto.NewCounter(...)`.
2.  Increment `opsProcessed.Inc()`.
3.  Serve `http.Handle("/metrics", promhttp.Handler())`.

### Explanation
Prometheus metrics in Go use the prometheus/client_golang library. Metrics are defined with promauto.NewCounter(), incremented with Inc(), and exposed via HTTP at /metrics endpoint using promhttp.Handler() for Prometheus scraping.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you expose Prometheus metrics in Go?
**Your Response:** "I expose Prometheus metrics using the `prometheus/client_golang` library. First, I define metrics like counters, gauges, or histograms using `promauto.NewCounter()` which automatically registers them. Throughout my code, I increment these metrics with `Inc()` or other methods. Finally, I expose an HTTP endpoint at `/metrics` using `http.Handle('/metrics', promhttp.Handler())`. Prometheus then scrapes this endpoint periodically to collect the metrics. This approach gives me comprehensive monitoring of my application's performance, error rates, and business metrics. The library handles all the complexity of metric formatting and exposition, so I just focus on defining the right metrics and updating them at the appropriate points in my code."

---

### Question 752: How do you set up OpenTelemetry in Go?

**Answer:**
1.  Create a TraceProvider (Exporter -> Jaeger/OTLP).
2.  Use `otel.SetTracerProvider`.
3.  In code: `tracer := otel.Tracer("name")`.
4.  `ctx, span := tracer.Start(ctx, "funcName")`.
5.  `defer span.End()`.

### Explanation
OpenTelemetry in Go requires creating a TraceProvider with exporters to Jaeger/OTLP, setting it globally with otel.SetTracerProvider, getting a tracer, starting spans with context, and ending spans with defer to capture trace data.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you set up OpenTelemetry in Go?
**Your Response:** "I set up OpenTelemetry in Go by following a five-step process. First, I create a TraceProvider configured with an exporter that sends traces to Jaeger or an OTLP-compatible collector. Second, I set this as the global tracer provider using `otel.SetTracerProvider`. In my application code, I get a tracer using `otel.Tracer('name')`. When I want to trace an operation, I create a span with `ctx, span := tracer.Start(ctx, 'operationName')`. Finally, I defer `span.End()` to ensure the span is properly closed when the operation completes. This setup gives me distributed tracing across my services, allowing me to follow requests through multiple microservices and identify performance bottlenecks. The context propagation happens automatically, so traces flow across service boundaries without additional code."

---

### Question 753: How do you trace gRPC requests in Go?

**Answer:**
Use middleware: `go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc`.
Client: `grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor())`.
Server: `grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor())`.
Matches trace headers automatically across services.

### Explanation
gRPC request tracing in Go uses OpenTelemetry middleware that automatically instruments unary and streaming calls. Client and server interceptors handle trace context propagation, automatically matching trace headers across services without manual intervention.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you trace gRPC requests in Go?
**Your Response:** "I trace gRPC requests using OpenTelemetry middleware from `go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc`. On the client side, I add `grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor())` to my dial options. On the server side, I add `grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor())` to my server options. The middleware automatically handles trace context propagation - it extracts trace headers from incoming requests, creates child spans, and injects trace headers into outgoing requests. This means I get complete distributed tracing across my gRPC services without manually handling trace headers. The middleware captures request/response details, latency, and errors automatically. This approach is much cleaner than manual instrumentation and works for both unary and streaming gRPC calls."

---

### Question 754: How do you record and export application traces?

**Answer:**
Spans are recorded in memory (buffered).
The Exporter pushes them via HTTP/gRPC to a Collector (Jaeger/Tempo) immediately or periodically (BatchSpanProcessor is more efficient).

### Explanation
Application traces are recorded in memory buffers and exported via HTTP/gRPC to collectors like Jaeger or Tempo. Exporters can send spans immediately or use BatchSpanProcessor for more efficient periodic batch sending to reduce overhead.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you record and export application traces?
**Your Response:** "Application traces are recorded in memory buffers and then exported to trace collectors. The spans I create are stored in memory until they're exported. The exporter pushes these spans via HTTP or gRPC to a collector like Jaeger or Tempo. I can configure the exporter to send spans immediately as they're completed, but for better performance, I use BatchSpanProcessor which batches spans and sends them periodically. This reduces the overhead of trace collection by amortizing the network cost across multiple spans. The collector then stores these traces and provides querying and visualization capabilities. This batched approach is much more efficient for high-throughput applications while still providing near real-time visibility into request flows across my services."

---

### Question 755: How do you handle slow endpoints in production Go apps?

**Answer:**
1.  **Timeout:** `http.TimeoutHandler`.
2.  **Instrumentation:** Log duration > 500ms.
3.  **Profiling:** Capture CPU profile during slowness.
4.  **Tracing:** Identify if it's DB, 3rd Party API, or CPU calculation.

### Explanation
Slow endpoints in production are handled with timeout handlers, duration logging, CPU profiling during slowness, and tracing to identify bottlenecks. This multi-pronged approach helps diagnose whether slowness is due to database queries, external APIs, or CPU-intensive operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle slow endpoints in production Go apps?
**Your Response:** "I handle slow endpoints using a multi-pronged approach. First, I implement timeout protection using `http.TimeoutHandler` to prevent requests from running indefinitely. Second, I add instrumentation to log requests that take longer than 500ms, creating an alert threshold. Third, I set up profiling to automatically capture CPU profiles when slowness is detected, so I can analyze what's causing the bottleneck. Fourth, I use distributed tracing to identify whether the slowness is coming from database queries, third-party API calls, or CPU-intensive calculations. This combination of timeout protection, monitoring, profiling, and tracing gives me both immediate mitigation and long-term diagnostic capabilities. I can quickly identify the root cause and take appropriate action, whether it's optimizing queries, adding caching, or scaling resources."

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

### Explanation
Custom labels/tags in logs are added using slog or zap by creating child loggers with context. The logger.With() method creates a new logger that automatically includes specified tags in all subsequent log entries, ensuring consistent context throughout request processing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you add custom labels/tags to logs?
**Your Response:** "I add custom labels to logs using structured logging libraries like `slog` or `zap`. I create a child logger with context using `logger.With('request_id', id, 'user_id', uid)`. This returns a new logger that automatically includes these tags in all subsequent log entries. When I call `reqLogger.Info('step 1')` and `reqLogger.Info('step 2')`, both logs will have the request_id and user_id tags automatically attached. This approach ensures consistent context throughout a request's lifecycle without manually adding tags to every log statement. It's especially useful for tracing single requests through complex workflows, making it easy to filter and search logs for specific operations or users. The child logger inherits all configuration from the parent while adding the contextual tags."

---

### Question 757: How do you redact sensitive data in logs?

**Answer:**
1.  **Custom Type:** `type Password string` with a `String()` method that returns "***".
2.  **Filter:** Implement a `slog.Handler` that checks keys (like "password", "token") and replaces values before writing.

### Explanation
Sensitive data redaction in logs uses custom types with String() methods returning masked values, or slog handlers that filter sensitive keys. This prevents passwords, tokens, and other sensitive information from appearing in log outputs while maintaining structured logging.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you redact sensitive data in logs?
**Your Response:** "I redact sensitive data in logs using two main approaches. First, I create custom types like `type Password string` with a `String()` method that returns '***' instead of the actual value. When the logging framework tries to serialize this type, it calls the String() method and gets the masked value. Second, I implement a custom `slog.Handler` that checks for sensitive keys like 'password', 'token', or 'secret' and replaces their values before writing to the log output. This handler-based approach is more flexible as it can handle any sensitive field without requiring custom types. Both approaches ensure that sensitive information never appears in logs, which is crucial for security and compliance. I prefer the handler approach for comprehensive coverage across all log entries."

---

### Question 758: How do you detect memory leaks using Go tools?

**Answer:**
pprof Heap Profile.
Look for high `inuse_objects` counts that constantly increase and never drop after GC runs.

### Explanation
Memory leak detection in Go uses pprof heap profiles to identify high inuse_objects counts that continuously increase without dropping after garbage collection runs. This pattern indicates objects being allocated but never freed, suggesting memory leaks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you detect memory leaks using Go tools?
**Your Response:** "I detect memory leaks using Go's pprof heap profiling. I run a heap profile and look for patterns where the `inuse_objects` count constantly increases and never drops, even after garbage collection runs. This indicates objects being allocated but never freed, which is the classic sign of a memory leak. I can use `go tool pprof` to analyze the heap profile and identify exactly which types of objects are leaking. I might also enable GC debug logging or use runtime.GC() to force garbage collection and see if memory usage drops. The key is looking for sustained growth patterns rather than temporary spikes. Memory leaks in Go often come from goroutines that never exit, cached references that are never cleared, or circular references in complex data structures. The pprof tools help me pinpoint the exact source of the leak."

---

### Question 759: How do you instrument performance counters in Go?

**Answer:**
`expvar` package (simple JSON Map of counters at `/debug/vars`).
Or Prometheus Counters/Gauges/Histograms.

### Explanation
Performance counters in Go use the expvar package for simple JSON counters exposed at /debug/vars, or Prometheus metrics with counters, gauges, and histograms for more sophisticated monitoring and alerting capabilities.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you instrument performance counters in Go?
**Your Response:** "I instrument performance counters using either Go's built-in `expvar` package or Prometheus metrics. The `expvar` package is simple - I just publish variables that are automatically exposed as a JSON map at `/debug/vars`. This is great for basic counters and gauges. For more sophisticated monitoring, I use Prometheus metrics which provide counters for monotonically increasing values, gauges for current values, and histograms for distributions like request latencies. Prometheus metrics integrate better with monitoring systems and support alerting. I choose based on complexity - for simple internal tools, expvar is sufficient, but for production services that need comprehensive monitoring and alerting, Prometheus is the better choice. Both approaches give me visibility into application performance and health metrics."

---

### Question 760: How do you implement a tracing middleware?

**Answer:**
```go
func TraceMid(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Start Span
        ctx, span := tracer.Start(r.Context(), r.URL.Path)
        defer span.End()
        
        r = r.WithContext(ctx)
        next.ServeHTTP(w, r)
        
        // Record Status Code
    })
}
```

### Explanation
Tracing middleware in Go creates spans using tracer.Start() with request context, updates the request with tracing context, calls the next handler, and records status codes. The defer span.End() ensures span completion regardless of request outcome.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a tracing middleware?
**Your Response:** "I implement tracing middleware as an HTTP handler wrapper. The middleware starts by creating a span using `tracer.Start(r.Context(), r.URL.Path)` with the request context and URL path as the operation name. I defer `span.End()` to ensure the span is properly closed. I update the request with the tracing context using `r.WithContext(ctx)` so downstream handlers can access the span. After calling the next handler, I record additional information like the response status code on the span. This middleware automatically traces all HTTP requests, creating spans that include timing, request metadata, and response information. The context propagation ensures that any spans created within the request handlers become child spans, giving me a complete trace of the request flow through my service."
