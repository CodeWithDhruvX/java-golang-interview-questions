# High-Level Design (HLD): Observability and Tracing

"How do you monitor this system?" is a guaranteed question in any modern HLD interview. "Just reading the logs" is no longer an acceptable answer for distributed microservices.

## 1. What are the Three Pillars of Observability?
**Answer:**
Observability is a measure of how well internal states of a system can be inferred from knowledge of its external outputs. It relies on:
1.  **Logs:** Immutable records of discrete events that happened over time (e.g., "User 123 failed login"). High detail, but hard to correlate across a massive system.
2.  **Metrics:** Numeric representations of data measured over intervals of time. (e.g., CPU Usage at 9:00 AM was 85%, HTTP 500 error rate is 2%). Great for alerting and dashboards.
3.  **Traces:** Represent the end-to-end journey of a single request as it traverses a distributed system. Tells you exactly where the bottleneck occurred.

## 2. How do you implement Distributed Tracing? (e.g., Jaeger, Zipkin, OpenTelemetry)
**Answer:**
When a user clicks "Checkout," the request might hit the API Gateway -> Order Service -> Inventory Service -> Payment Service. If it takes 5 seconds, how do you know which service was slow?
*   **The Solution (Trace ID and Span IDs):**
    1.  When a request enters the API Gateway, the gateway generates a unique UUID called a **Trace ID** (e.g., `Trace-123`).
    2.  The gateway injects this Trace ID into HTTP Headers (`X-B3-TraceId`) before forwarding the request to the Order Service.
    3.  Every microservice involved must extract the Trace ID from incoming requests and inject it into any outgoing requests it makes.
    4.  Every individual unit of work (e.g., "DB query in Order Service", "HTTP call to Payment Service") generates a **Span ID**. A Span contains start time, end time, and the parent Trace ID.
    5.  Services asynchronously fire these Spans to a central tracing backend (like Jaeger).
    6.  The backend GUI pieces together all Spans sharing `Trace-123` to generate a visual waterfall chart of exactly how that single transaction flowed through the entire network, pinpointing latency exactly.

## 3. How do you design an effective Metrics and Alerting pipeline?
**Answer:**
*   **Collection Model:**
    *   *Push-based (e.g., StatsD, InfluxDB):* Applications actively send metrics over network UDP to a time-series database.
    *   *Pull-based (e.g., Prometheus):* Applications expose an HTTP endpoint (e.g., `/metrics`). A central Prometheus server scrapes this endpoint every 15 seconds. Pull is often preferred in dynamic environments (like Kubernetes) as it prevents the central server from being overwhelmed by aggressive clients.
*   **Storage:** A Time-Series Database (TSDB). These are highly optimized for appending sequential timestamped data and querying aggregates over time windows. 
*   **Alerting Rules:** Avoid alerting on CPU being at 80% (noise). Alert on **SLI/SLO breaches** (e.g., "The P99 latency for the checkout endpoint has exceeded 2 seconds for 5 consecutive minutes"). Use tools like PagerDuty or Alertmanager to route the alert to the on-call engineer.

## 4. What is structured logging and why is it necessary?
**Answer:**
*   **Unstructured Log:** `logger.info("User 456 failed to purchase item 99 in 400ms")`
*   **Structured Log:** `logger.info(json_serialize({ "event": "purchase_failed", "user_id": 456, "item_id": 99, "duration_ms": 400 }))`
*   **Why?** In an ELK or EFK stack (Elasticsearch, Fluentd, Kibana), searching for "how many users failed purchasing item 99" using regex on unstructured text is painfully slow. With structured JSON logs, Logstash easily parses the fields, and Elasticsearch indexes `user_id` and `item_id` directly, making complex analytical queries instant.
