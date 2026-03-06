# 🏗️ Kafka — Observability, Monitoring & Production Debugging

> **Level:** 🔴 Senior to 🟣 Architect
> **Asked at:** Netflix, Uber, LinkedIn, Amazon, Swiggy

---

## Q1. What are the critical JMX metrics you monitor in a Kafka production cluster and what do they indicate?

"Kafka exposes rich operational metrics via JMX (Java Management Extensions). The most critical ones fall into three groups: **Broker Health**, **Producer Health**, and **Consumer Health**.

**Key Broker Metrics:**

| Metric | What It Means | Alert Threshold |
|---|---|---|
| `kafka.server:type=BrokerTopicMetrics,name=MessagesInPerSec` | Message ingestion rate | Sudden drop → producer issue |
| `kafka.server:type=ReplicaManager,name=UnderReplicatedPartitions` | Partitions not fully replicated | Any value > 0 → ISR degraded |
| `kafka.controller:type=KafkaController,name=ActiveControllerCount` | Number of active controllers | Must be exactly 1; 0 = cluster crisis |
| `kafka.network:type=RequestMetrics,name=RequestsPerSec` | Request throughput per request type | Baseline alerting |
| `kafka.log:type=LogFlushStats,name=LogFlushRateAndTimeMs` | Time to flush log to disk | High → disk I/O bottleneck |

**Key Consumer Metrics:**

| Metric | What It Means |
|---|---|
| `kafka.consumer:type=consumer-fetch-manager-metrics,name=records-lag-max` | Maximum lag across all partitions — the most important consumer metric |
| `kafka.consumer:type=consumer-coordinator-metrics,name=rebalance-rate-per-hour` | How often rebalances occur — high value = instability |

**Prometheus + Grafana Setup:**
Use the `jmx_exporter` Java agent to scrape JMX metrics and expose them on an HTTP endpoint. Prometheus scrapes this endpoint every 15s. Grafana dashboards built on community templates (e.g., the Kafka Overview dashboard) visualize these metrics with pre-configured alerts."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Micrometer elegantly exports these underlying JMX metrics into a `/actuator/prometheus` endpoint in Spring apps running the consumers/producers natively alongside the JVM logic.
* **Golang:** The Go clients like `confluent-kafka-go` export statistics natively (via `Stats()` channel) which must be manually unpacked and exposed to `/metrics` endpoints using `prometheus/client_golang` within your Go microservice.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix, Uber — candidates are expected to know what to look at on a Grafana dashboard when an on-call alert fires at 2 AM, and what the metrics mean operationally.

#### Indepth
**Confluent Control Center vs. Open Source Monitoring:** Confluent's commercial Control Center provides a GUI with built-in consumer lag monitoring, rebalance history, and schema registry metrics. Open-source alternatives like **Kafka UI** (by Provectus) or **AKHQ** provide similar visibility without a commercial license — commonly used at cost-conscious startups.

---

## Q2. How do you debug a sudden spike in consumer lag in production?

"Consumer lag spikes are the most common Kafka production incident. Here is a systematic debugging runbook:

**Step 1 — Identify WHICH consumer group and partition has the lag:**
```bash
kafka-consumer-groups.sh \
  --bootstrap-server broker:9092 \
  --describe --group my-service-group
```
Output shows each partition's `LOG-END-OFFSET`, `CURRENT-OFFSET`, and `LAG`. Look for partitions with disproportionately high lag.

**Step 2 — Rule out producer surge (is the problem speed or volume?):**
Check `MessagesInPerSec` on the Grafana broker dashboard. If the ingestion rate suddenly doubled but consumer throughput stayed flat, the consumers are simply under-provisioned.
- **Fix:** Add more partitions + scale out consumers to match.

**Step 3 — Check for consumer processing slowdown:**
Examine application logs for slow external calls: database timeouts, downstream HTTP 503s, GC pause logs. A single slow downstream service can cascade.
- **Fix:** Add circuit breakers, increase thread pool for consumer processing, or move heavy processing to async.

**Step 4 — Check for rebalances:**
`rebalance-rate-per-hour` metric or logs showing `Revoke partitions` indicate the consumer group is continuously rebalancing — consuming zero messages during each rebalance.
- **Fix:** Tune `max.poll.interval.ms` and `max.poll.records`. Check for a poison pill causing crashes.

**Step 5 — Verify `records-lag-max` is trending down:**
After applying fixes, watch the lag metric in Grafana. The lag should decrease as the consumer catches up. If it plateaus, consumption rate still equals production rate and you need to scale further."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** Spring's auto-configured ThreadPools mean you can easily tweak `concurrency` property inside `@KafkaListener` to spin up more internal worker threads, assuming enough partitions exist locally without scaling VMs.
* **Golang:** Increasing throughput natively involves just spawning more goroutines behind an internal channel off taking events from a single reader. Scaling horizontally implies running more pod replicas with identical Reader `GroupID` configurations to pull from unassigned partitions.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Swiggy, Amazon, LinkedIn — this is a classic senior SRE/backend interview question testing systematic debugging rather than memorized facts. Companies want engineers who can follow a structured runbook.

#### Indepth
**Lag by Time vs. Lag by Messages:** Consumer lag in messages (by offset delta) can be misleading if message sizes vary. One lagged offset may represent 100MB or 10 bytes. Use `MaxLag` combined with `MessagesInPerSec` rate to calculate the real estimated time-to-catch-up: `lag_messages / consumption_rate_per_second = seconds_to_recover`.

---

## Q3. How do you implement distributed tracing across Kafka-based microservices?

"In a synchronous REST chain, distributed tracing is straightforward — inject a `trace-id` header in the HTTP request and propagate it downstream. In async Kafka pipelines, the producer and consumer run in different processes at different times, making propagation manual but essential for debugging latency issues.

**Implementing OpenTelemetry Trace Propagation with Kafka:**

**Producer Side — Inject trace context into Kafka headers:**
```java
// Using OpenTelemetry API
Span span = tracer.spanBuilder("kafka-produce").startSpan();
try (Scope scope = span.makeCurrent()) {
    ProducerRecord<String, String> record = new ProducerRecord<>("orders", key, payload);
    
    // Inject W3C TraceContext headers into Kafka record headers
    openTelemetry.getPropagators().getTextMapPropagator().inject(
        Context.current(), record.headers(), 
        (headers, key, value) -> headers.add(key, value.getBytes())
    );
    kafkaTemplate.send(record);
} finally {
    span.end();
}
```

**Consumer Side — Extract trace context and continue the trace:**
```java
@KafkaListener(topics = "orders")
public void consume(ConsumerRecord<String, String> record) {
    // Extract trace context from Kafka headers
    Context extractedContext = openTelemetry.getPropagators().getTextMapPropagator()
        .extract(Context.current(), record.headers(), 
            (headers, key) -> new String(headers.lastHeader(key).value()));
    
    Span span = tracer.spanBuilder("kafka-consume")
        .setParent(extractedContext)
        .startSpan();
    // ... process and end span
}
```

This ensures the Jaeger/Zipkin/Tempo trace waterfall graph shows one end-to-end trace from HTTP request → Kafka produce → Kafka consume → downstream DB call, with accurate inter-service latency breakdown."

#### 💻 Language Specifics (Java Spring Boot & Golang)
* **Java Spring Boot:** `micrometer-tracing` combined with `spring-kafka` auto-propagates `traceparent` contexts into Kafka Record Headers invisibly, requiring practically zero boilerplate Java code.
* **Golang:** Both `segmentio` and `confluent` clients natively support header propagation. Using `go.opentelemetry.io`, you utilize `otel.GetTextMapPropagator().Inject(ctx, &MessageHeaderCarrier{})` when producing, and `Extract` the slice of `kafka.Header` structures on the reader side.

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix, Uber — at scale, debugging 'why did this order take 30 seconds' across 5 Kafka hops is impossible without distributed tracing. This shows the candidate understands modern observability beyond simple logging.

#### Indepth
**Baggage Propagation:** Beyond trace IDs, OpenTelemetry supports **baggage** — arbitrary key-value pairs propagated across service boundaries. Examples: `user.tier=premium`, `request.region=in-south`. Downstream Kafka consumers can read baggage to make routing or SLA decisions without re-fetching context from a database.
---
