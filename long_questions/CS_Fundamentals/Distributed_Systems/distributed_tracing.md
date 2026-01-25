# Distributed Tracing

## 1. The Problem
In a Monolith, debugging is easy: look at one log file.
In Microservices, a single user request might traverse 10 different services. If the request fails or is slow, how do you find the root cause?

## 2. Key Concepts (OpenTracing / OpenTelemetry)

### Trace
A **Trace** represents the journey of a single request through the entire system. It is a Directed Acyclic Graph (DAG) of **Spans**.

### Span
A **Span** represents a single unit of work (e.g., "Service A calling Service B" or "Service B querying Database").
*   Contains: `SpanID`, `ParentSpanID`, `TraceID`, StartTime, EndTime, Tags (Metadata).
*   **Root Span**: The first span in the trace (no parent).

### Context Propagation
The mechanism of passing the `TraceID` and `SpanID` from one service to another.
*   **HTTP Headers**: Usually passed via HTTP headers like `X-B3-TraceId` (Zipkin) or `traceparent` (W3C standard).
*   **Mechanism**:
    1.  Service A receives request. Generates `TraceID=100`, `SpanID=A`.
    2.  Service A calls Service B. Injects header `TraceID=100, ParentSpanID=A`.
    3.  Service B receives header. Generates `SpanID=B`. Logs `TraceID=100`, `ParentSpanID=A`.

## 3. Visualization Tools
*   **Jaeger / Zipkin**: UIs that visualize the trace as a Gantt Chart.
*   **Benefit**: You can visually see which span took the longest (the "Longest Bar").

## 4. Implementation Logic (Pseudo-code)

```python
# Service A
def handle_request(req):
    span = tracer.start_span("operation_A")
    defer span.finish()
    
    # Propagate context
    headers = {}
    tracer.inject(span.context, headers)
    
    # Call Service B
    http.get("http://service-b", headers=headers)
```

```python
# Service B
def handle_request(req):
    # Extract context
    ctx = tracer.extract(req.headers)
    
    # Start child span
    span = tracer.start_span("operation_B", child_of=ctx)
    defer span.finish()
    
    do_work()
```

## 5. Interview Questions
1.  **What is the difference between Logging, Monitoring, and Tracing?**
    *   *Ans*:
        *   **Logging**: "What happened?" (Events, errors).
        *   **Monitoring (Metrics)**: "Is it healthy?" (CPU usage, P99 latency charts).
        *   **Tracing**: "Where did it happen?" (Request flow analysis).
2.  **How do you handle sampling?**
    *   *Ans*: Tracing every single request is expensive (storage/CPU). We use **Sampling** (e.g., trace only 1% of requests).
        *   **Head-based sampling**: Decide at the start of the trace.
        *   **Tail-based sampling**: Decide after the trace is complete (keep only error traces).
