# 🟢 **296–305: Observability & SRE**

### 296. What are the three pillars of Observability?
"The three pillars are **Metrics, Logs, and Traces**. 

- **Metrics** are numerical data aggregated over time (e.g., CPU Usage: 80%, Request Rate: 500/sec). They are cheap to store and alert on when things go wrong. Tools: Prometheus, Datadog.
- **Logs** are immutable strings describing discrete events (e.g., 'User 123 failed login'). They are expensive to search but provide deep context on *why* something went wrong. Tools: ELK stack (Elasticsearch, Logstash, Kibana).
- **Traces** track the progression of a single request as it hops across multiple microservices. They tell you exactly *where* the bottleneck is. Tools: Jaeger, Zipkin, OpenTelemetry.

A mature system unifies these. An alert on a *Metric* leads to a slow *Trace*, which contains the exact *Log* line causing the error."

#### Indepth
OpenTelemetry (OTel) is the modern industry standard. Previously, developers had to include Datadog agents, Prometheus exporters, and Zipkin tracers individually. OTel provides a single, vendor-agnostic SDK. You instrument your code once with OTel, and the OTel Collector routes the data to Datadog for traces, Prometheus for metrics, and Elastic for logs.

---

### 297. How does Distributed Tracing work?
"When an API Gateway receives a brand new HTTP request, it generates a unique ID (the **Trace ID**) and initiates a Root Span.

When the Gateway calls the Order Service, it physically injects this `TraceID` into the HTTP headers (e.g., standard `traceparent` header in W3C format). 

The Order Service reads the header, creates a child 'Span', does its work, and passes the exact same `TraceID` downstream to the Payment Service. 

Behind the scenes, all services unilaterally send their Spans to a central tracing server (like Jaeger). Since all spans share the same `TraceID`, Jaeger stitches them together visually into a Gantt chart, showing me exactly that the Order service took 50ms, the DB took 20ms, and the Payment service took 800ms (revealing the bottleneck)."

#### Indepth
Integrating traces with logs is critical. If your logging framework (like Logback or unstructured JSON logs in Go) is configured properly (using MDC - Mapped Diagnostic Context in Java), every single `log.info()` or `log.error()` is automatically tagged with the current `TraceID`. This allows you to instantly pull up all logs across 5 different services that correspond exactly to a specific user's failed request.

---

### 298. What is the difference between Service Level Agreements (SLA), SLOs, and SLIs?
"These are SRE (Site Reliability Engineering) concepts for measuring reliability.

- **SLI (Service Level Indicator):** A direct, factual measurement of performance. Example: '99.5% of HTTP GET requests to `/orders` completed with a 200 OK in the last 5 minutes.'
- **SLO (Service Level Objective):** The internal target your engineering team commits to hitting. Example: 'We aim for our SLI to be successfully above 99.9% measured over a rolling 30-day window.'
- **SLA (Service Level Agreement):** The legally binding contract with the customer. Example: 'If we drop below 99.5% uptime, we will refund the client 10% of their monthly bill.' SLAs are always looser than SLOs to provide a safety margin."

#### Indepth
If your SLO is 99.9% uptime per month, that mathematically equals an **Error Budget** of 43.8 minutes of allowable downtime per month. SRE teams track this closely. If an engineering team burns through their Error Budget by deploying buggy features in Week 1, CI/CD is frozen, and the team is forced to halt feature development and exclusively ship reliability tasks until the budget recovers.

---

### 299. How do you implement structured logging?
"Traditional logging like `log.info("User {} bought item {}", userId, itemId)` generates unstructured strings. Searching for all purchases of 'item_55' in an ELK stack requires slow, complex RegEx parsing of millions of lines of text.

**Structured Logging** means the application strictly logs JSON objects, not strings.

```json
{"timestamp": "2024-03-01T10:00:00Z", "level": "INFO", "message": "Purchase successful", "user_id": 123, "item_id": "item_55"}
```

Because Elasticsearch natively digests JSON, I can query `item_id=item_55 AND level=ERROR` and get results in milliseconds without any parsing pipelines."

#### Indepth
Implementing structured logging often reveals massive log volume issues. Because you are now logging large JSON payloads instead of short strings, disk I/O and network bandwidth to Logstash drastically increase. You must use log sampling architectures or drop DEBUG logs actively at the edge.

---

### 300. What is a synthetic transaction in monitoring?
"Relying purely on server metrics (like checking if the internal Payment API is returning HTTP 200) is a false sense of security. The Payment API might be 'healthy', but a Javascript error on the frontend checkout button might mean zero users can actually buy anything.

**Synthetic Monitoring** means writing a script (often a headless Chrome browser script) that simulates a real user navigating the UI, logging in, adding an item to the cart, and clicking checkout via a test credit card.

We configure Datadog or New Relic to run this script automatically every 3 minutes from 5 different geographic regions globally. If the script fails, it pages the on-call engineer immediately, proving that the actual user journey is broken."

#### Indepth
Synthetic tests create garbage data in production (e.g., thousands of fake 'Test Orders'). You must uniquely flag synthetic test accounts or inject special HTTP headers (`X-Synthetic-Test: true`) so that analytics dashboards and revenue reporting pipelines explicitly ignore them.
