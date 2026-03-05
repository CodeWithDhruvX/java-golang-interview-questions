# 🔭 Observability in Distributed Systems — Product-Based Companies

> **Level:** 🟡 Mid – 🔴 Senior
> **Asked at:** Swiggy, Zepto, Razorpay, Uber, Amazon — any company operating microservices at scale

---

## Q1. What are the three pillars of observability? How do they differ from monitoring?

"Observability is **the ability to understand the internal state of a system by examining its external outputs** — without having to deploy new instrumentation every time a new failure mode appears.

The three pillars: **Logs** (what happened), **Metrics** (how much / how fast), **Traces** (how did this specific request flow through the system).

Monitoring is looking at pre-defined dashboards and alerts. Observability is the property that lets you ask arbitrary new questions about your system's behavior — questions you didn't think to ask when you built it. Well-observable systems let you debug novel failures."

#### Company Context & Level
**Level:** 🟡 Mid | **Asked at:** Swiggy, Zepto, Groww, Razorpay — any company with microservices infra

#### Deep Dive
**Logs — what happened:**
```
Purpose: Record specific events with context.

Good log (structured JSON — machine parseable):
  {
    "timestamp": "2024-03-05T04:30:00Z",
    "level": "ERROR",
    "service": "payment-service",
    "trace_id": "7f83b1657...",
    "user_id": 12345,
    "order_id": "OD789",
    "event": "payment_failed",
    "reason": "bank_timeout",
    "duration_ms": 3024,
    "retry_count": 2
  }

Bad log (unstructured — hard to query):
  "ERROR: Payment failed for user 12345 after 3 seconds"

Stack: Prometheus/Loki, ELK (Elasticsearch + Logstash + Kibana), Datadog Logs.
```

**Metrics — how much / how fast:**
```
Purpose: Numerical measurements over time. Aggregated. Queryable.

Types:
  Counter: monotonically increasing (total_requests, total_errors)
  Gauge: point-in-time value (current_active_connections, memory_usage)
  Histogram: distribution of values (request_duration_ms — tracks p50, p95, p99)
  Summary: similar to histogram, pre-computed quantiles

Key metrics for a payment service:
  payment_requests_total{status="success"} = 1,203,456
  payment_requests_total{status="failure"} = 1,234
  payment_duration_seconds{quantile="0.99"} = 0.842   -- p99 latency: 842ms
  payment_amount_rupees_sum = 98,765,432
  
Error rate = failure_total / (success_total + failure_total) × 100 = 0.10%
Stack: Prometheus + Grafana, Datadog Metrics, AWS CloudWatch.
```

**Traces — how did this request flow:**
```
Purpose: Track a single request across multiple services end-to-end.

Distributed trace for an order placement:
  Trace ID: "7f83b1657aac4b8e"
  
  Spans (child operations):
    [API Gateway]           0ms → 5ms   (5ms)    ✓
    [Order Service]         5ms → 47ms  (42ms)   ✓
      [Inventory Check]     8ms → 20ms  (12ms)   ✓
      [DB Write]           21ms → 40ms  (19ms)   ✓
    [Payment Service]      47ms → 891ms (844ms)  ✓ ← SLOW
      [Bank API Call]      50ms → 889ms (839ms)  ✓ ← ROOT CAUSE
    [Notification Service] 891ms → 900ms (9ms)   ✓

Trace reveals: bank_api_call taking 839ms → payment service p99 inflated
Without tracing: you only see "order placement is slow" — no idea which service
Stack: Jaeger, Zipkin, AWS X-Ray, Datadog APM.
```

---

## Q2. What are SLI, SLO, and SLA? How do you use them to drive architecture decisions?

"SLI (Service Level Indicator) is the actual measured metric. SLO (Service Level Objective) is the target you set internally. SLA (Service Level Agreement) is the contractual commitment to customers with financial penalties.

The hierarchy: SLA ← SLO ← SLI. Your SLO must be tighter than your SLA to give you a buffer. We commit to customers a 99.5% uptime SLA (SLA), we target 99.9% internally (SLO), and we measure 99.95% actual uptime (SLI). The gap between SLA and SLO is your operating margin — it absorbs incidents without breaching the customer contract."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Amazon, Google (SRE framework heavily used), Razorpay, Zepto

#### Deep Dive
**Defining SLIs — what to measure:**
```
Availability SLI:
  successful_requests / total_requests (last 30 days)
  = 10,000,000 - 50,000 / 10,000,000 = 99.5%

Latency SLI:
  % of requests completing in < 500ms = 97.3%
  (or raw: p99 = 487ms, p95 = 320ms, p50 = 120ms)

Error rate SLI:
  5xx responses / total responses = 0.05%

Freshness SLI (for data pipelines):
  % of data processed within 15 minutes of ingestion = 99.1%
```

**Error budget — the operational tool:**
```
SLO: 99.9% availability over 30 days
30-day total minutes: 43,200

Error budget = (100% - 99.9%) × 43,200 = 43.2 minutes of allowed downtime

If you've burned 20 minutes of downtime → 23.2 minutes remaining
  → Comfortable territory → can deploy risky changes

If you've burned 40 minutes → error budget almost exhausted
  → Freeze risky deployments until month resets
  → Focus engineering on reliability instead of features

Error budget makes reliability/velocity trade-off quantitative and automatic.
If teams consistently burn budget → architecture is too fragile.
If teams never burn budget → SLO is too conservative → can invest more in features.
```

**Calculating availability from SLO:**
```
99%    → 7.2 hours downtime/month
99.9%  → 43.2 minutes downtime/month  (three nines)
99.99% → 4.32 minutes downtime/month  (four nines)
99.999% → 25.9 seconds downtime/month (five nines — extremely hard)

Four nines (99.99%):
  Every deployment must take < 1 minute (no-downtime rolling deploys)
  Every incident must resolve in < 2 minutes
  Requires: active-active multi-region, automated failover, runbooks for every failure mode
```

**Alerting from SLIs — burn rate alerts:**
```
Naive alert: "page me if error rate > 1% for 5 minutes"
  Problem: 5-minute window catches short spikes (false positives)
           Misses slow burns (2% error rate for 20 minutes → quietly burns budget)

Better: Multi-window burn rate alert
  Alert if error budget burn rate > 14.4x for last 1 hour
  AND error budget burn rate > 3x for last 6 hours
  
  Explanation: 14.4x burn rate for 1 hour = 14.4 × 43.2min × (1h/720h) = burns 0.86% budget
               But sustained will exhaust budget
  
  This catches both fast burns (incidents) and slow burns (gradual degradation).
  
  Tools: Prometheus AlertManager + alerting rules, Datadog SLO alerts.
```

---

## Q3. Explain distributed tracing with OpenTelemetry. How do you propagate trace context across services?

"OpenTelemetry (OTel) is the vendor-neutral standard for instrumenting distributed systems. It provides SDKs for collecting traces, metrics, and logs, and exporting them to any backend (Jaeger, Zipkin, Datadog, New Relic, AWS X-Ray).

Trace context propagation is the key mechanism: when Service A calls Service B, it passes the trace ID and span ID in an HTTP header (`traceparent`). Service B creates a child span under the same trace. This links all spans from all services into one unified trace timeline."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Uber, Swiggy, any company running microservices with distributed tracing

#### Deep Dive
**W3C Trace Context propagation:**
```
HTTP Header: traceparent: 00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01

Format: {version}-{trace-id}-{parent-span-id}-{flags}
  version: 00 (always)
  trace-id: 32-char hex, globally unique across all services for this request
  parent-span-id: 16-char hex, the current service's span ID (becomes parent for next service)
  flags: 01 = sampling on, 00 = sampling off

Service A creates: trace_id = "4bf92f3577b34da6a3ce929d0e0e4736", span_id = "aaaaaaaaaaaaaaa"
Service A calls B: passes traceparent header above
Service B: creates new span with parent = "aaaaaaaaaaaaaaa"
  → Jaeger/Zipkin can now link A and B's spans into one trace
```

**Sampling strategies — not tracing every request:**
```
At 100K RPS, tracing every request = 100K traces/second = huge storage cost.

Head-based sampling (decision at request start):
  Random: sample 1% of all requests. Simple. Misses rare errors.
  Rate-limited: sample up to 100 traces/second regardless of traffic.
  
Tail-based sampling (decision after request completes):
  Sample 100% of requests that had an error
  Sample 100% of requests taking > 1 second
  Sample 1% of healthy, fast requests
  
  Why tail-based is better: captures all interesting requests.
  Why it's hard: requires buffering completed spans before sampling decision.
  
Tools for tail-based: OpenTelemetry Collector, Tempo (Grafana).
```

**Go service with OpenTelemetry (practical example):**
```go
// Initialize tracer
func initTracer() *trace.TracerProvider {
    exporter, _ := jaeger.New(jaeger.WithCollectorEndpoint(
        jaeger.WithEndpoint("http://jaeger:14268/api/traces"),
    ))
    tp := trace.NewTracerProvider(
        trace.WithBatcher(exporter),
        trace.WithResource(resource.NewWithAttributes(
            semconv.ServiceNameKey.String("payment-service"),
        )),
        trace.WithSampler(trace.TraceIDRatioBased(0.1)), // 10% sampling
    )
    otel.SetTracerProvider(tp)
    otel.SetTextMapPropagator(propagation.TraceContext{}) // W3C propagation
    return tp
}

// Instrument HTTP handler
func paymentHandler(w http.ResponseWriter, r *http.Request) {
    // Extract trace context from incoming request headers
    ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
    
    tracer := otel.Tracer("payment-service")
    ctx, span := tracer.Start(ctx, "process-payment")
    defer span.End()
    
    span.SetAttributes(
        attribute.String("payment.order_id", r.Header.Get("X-Order-ID")),
        attribute.String("payment.method", "upi"),
    )
    
    // Call bank API — passes trace context in outgoing request
    result, err := callBankAPI(ctx, amount)
    if err != nil {
        span.RecordError(err)
        span.SetStatus(codes.Error, err.Error())
    }
}
```

---

## Q4. How do you implement centralized logging for microservices? What is the ELK stack?

"In a microservices system, each service writes logs to its own stdout. To debug a user's order failure, you need logs from the API gateway, Order Service, Inventory Service, and Payment Service — all scattered across different pods/containers. Centralized logging aggregates all logs into one queryable system.

The ELK stack: **Elasticsearch** (storage + search), **Logstash** (log processing pipeline), **Kibana** (visualization/dashboards). Modern alternative: the EFK stack (replace Logstash with Fluentd/Fluent Bit) — lighter weight, better for Kubernetes."

#### Company Context & Level
**Level:** 🟡 Mid | **Asked at:** Swiggy, Zepto, Flipkart, service companies implementing DevOps

#### Deep Dive
**Log pipeline in Kubernetes:**
```
Pod container → writes logs to stdout
Kubernetes: captures stdout → stores in /var/log/containers/ on node

Fluent Bit (runs as DaemonSet on every node):
  → Reads /var/log/containers/
  → Parses JSON logs
  → Adds Kubernetes metadata: {pod_name, namespace, node, container}
  → Sends to Elasticsearch

Elasticsearch:
  → Stores indexed in time-series indexes: logs-2024.03.05
  → Full-text search + structured field queries

Kibana:
  → "Show me all logs where trace_id = '7f83b165' and level = 'ERROR'"
  → Discover view: browse raw logs
  → Dashboard: error rate over time, top error messages, latency histograms
```

**Structured logging — essential for searchability:**
```go
// BAD: String interpolation — cannot query individual fields
log.Printf("Payment failed for user %d, amount %d, error: %s", userID, amount, err)

// GOOD: Structured fields — each field is queryable in Elasticsearch
logger.Error("payment_failed",
    zap.Int64("user_id", userID),
    zap.Int64("amount_rupees", amount),
    zap.String("error_code", "bank_timeout"),
    zap.String("order_id", orderID),
    zap.String("trace_id", traceID),
    zap.Int("retry_count", retryCount),
    zap.Duration("duration", time.Since(start)),
)

// Kibana query: user_id:12345 AND error_code:"bank_timeout" AND @timestamp:[now-1h TO now]
// → Returns all failed payments for user 12345 in the last hour
```

**Log levels and what to log at each level:**
```
DEBUG:  Detailed flow for local development. Never in production (too noisy, cost).
INFO:   Important business events: "order_placed", "payment_success", "user_registered"
WARN:   Unexpected but recoverable: "retrying_db_connection", "cache_miss_fallback_to_db"
ERROR:  Something failed but request still served: "payment_failed", "inventory_reserve_failed"
FATAL:  Unrecoverable — service cannot continue (program exits after logging)

Production log level: INFO. Set to DEBUG temporarily during incident investigation.
```

---

## Q5. What do you monitor in a production microservices system? What alerts would you set up?

"I follow the **USE method** (Utilization, Saturation, Errors) for infrastructure and the **RED method** (Rate, Errors, Duration) for services.

RED for every service: Rate (requests per second), Errors (error rate %), Duration (p50/p95/p99 latency).

These three metrics, for every service, tell you immediately if something is wrong and identify the faulty service. It's the minimum viable observability baseline."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** SRE/Platform roles, senior backend at Amazon, Swiggy, Razorpay

#### Deep Dive
**Grafana dashboard — RED metrics per service:**
```
Panel 1: Request Rate (RPS)
  sum(rate(http_requests_total{service="payment"}[5m])) by (status_code)
  → See if traffic is normal, spiking, or dropped to zero (downstream outage)

Panel 2: Error Rate (%)
  sum(rate(http_requests_total{service="payment", status=~"5.."}[5m]))
  / sum(rate(http_requests_total{service="payment"}[5m])) × 100
  → Alert if > 1% for the p2p minutes

Panel 3: Latency Distribution (p50, p95, p99)
  histogram_quantile(0.99, rate(http_request_duration_seconds_bucket{service="payment"}[5m]))
  → Alert if p99 > 2 seconds

Panel 4: Saturation — Infrastructure
  CPU: avg(rate(container_cpu_usage_seconds_total[5m])) by (pod) / request_cpu_limit
  Memory: container_memory_working_set_bytes{pod=~"payment.*"} / container_spec_memory_limit_bytes
  → Alert if CPU > 85% or Memory > 90%

Panel 5: Database
  Slow queries: pg_stat_statements_mean_exec_time > 100ms
  Active connections: pg_stat_activity_count approaching max_connections
  Replication lag: pg_replication_slots_confirmed_flush_lsn
```

**Alerting runbook (what to page, what to silently log):**
```
PAGE ON-CALL IMMEDIATELY (P1):
  - Error rate > 5% for 5 minutes
  - p99 latency > 5 seconds for 5 minutes
  - Service not responding (health check failure for 3 consecutive checks)
  - Database connection pool exhausted

PAGE WITH 30-MINUTE DELAY (P2):
  - Error rate > 1% for 15 minutes
  - p99 latency > 2 seconds for 10 minutes
  - CPU > 90% for 10 minutes
  - Disk usage > 85%

SLACK NOTIFICATION ONLY (no page):
  - p99 > 1 second
  - Cache hit rate dropped below 80%
  - DB slow queries > 10 per minute
  - Error budget burn rate elevated

NEVER ALERT ON:
  - p50 latency (too noisy, hides tail latencies that matter)
  - Absolute request counts (traffic varies by time of day)
  - Individual exception occurrences (alert on rates, not instances)
```
