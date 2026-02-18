# Observability & Monitoring Interview Questions (126-133)

## Production Readiness & Monitoring

### 126. What is observability?
"Monitoring tells you *when* something is wrong ('The server is down').

Observability tells you *why* using the data you've collected. It’s a property of the system—how well can you understand its internal state just by looking at its external outputs (logs, metrics, traces)?

If I have high observability, I don't need to SSH into a box to debug. I can just query my tools and say, 'Ah, the latency spiked because this specific database query locked the table.'"

### 127. Difference between logs, metrics, and traces?
"These are the three pillars of observability.

**Logs**: Detailed, unstructured text events. 'User 123 logged in'. Good for debugging specific errors. High volume, expensive to store.

**Metrics**: Aggregated numbers. 'CPU usage is 80%'. 'Request count is 500/sec'. Good for health dashboards and alerting. Cheap to store.

**Traces**: The journey of a single request across multiple services. 'Request hit API Gateway -> Auth Service -> Product Service -> DB'. Good for finding latency bottlenecks in microservices."

### 128. What are SLAs, SLOs, and SLIs?
"**SLI (Indicator)**: The actual measurement. 'Our current availability is 99.95%.'

**SLO (Objective)**: The goal we set internally. 'We aim for 99.9% availability.' If we dip below this, we stop shipping features and fix bugs.

**SLA (Agreement)**: The legal contract with customers. 'If availability drops below 99.0%, we will refund you 10%.' This is a business promise, usually looser than the SLO."

### 129. How do you monitor a Spring Boot application?
"I use **Spring Boot Actuator**. It exposes endpoints like `/actuator/health`, `/actuator/metrics`, and `/actuator/info` out of the box.

For metrics, I usually add `micrometer-registry-prometheus`. This formats the metrics so Prometheus can scrape them.

Then I visualize everything in **Grafana**. I have dashboards for JVM internals (Heap, GC, Threads), HTTP request rates, and custom business metrics."

### 130. What is distributed tracing?
"In microservices, a single user request might touch 10 different services. If the request is slow, logs alone won't tell you *where* it was slow.

Distributed tracing assigns a unique **Trace ID** to the request at the entry point. This ID is passed in headers (e.g., `X-B3-TraceId`) to every downstream service.

Tools like Zipkin or Jaeger then collect these spans and visualize them as a waterfall chart, showing exactly how long each hop took."

### 131. What are health checks and readiness probes?
"These are critical for Kubernetes.

**Liveness Probe**: 'Am I alive?' If this fails (e.g., deadlock), K8s restarts the pod.

**Readiness Probe**: 'Am I ready to take traffic?' It might be alive but still warming up caches or connecting to the DB. If this fails, K8s stops sending traffic to this pod until it passes.

If you don't separate these, you might restart a pod that is simply busy, or send traffic to a pod that isn't fully initialized yet."

### 132. How do you troubleshoot high latency in production?
"I start with the **Metrics** dashboard. Is it a global issue or just one instance?

Then I check **Traces** (Jaeger) to find the bottleneck. Is it the DB? An external API call? Or CPU processing?

If it's the DB, I check slow query logs.
If it's CPU/Application, I might take a **Thread Dump** to see if threads are blocked.
If it's memory, I check **GC logs**—frequent 'Stop-The-World' pauses often look like latency spikes."

### 133. What is log correlation?
"It's the practice of linking all logs related to a single request using a unique ID.

If I have 1000 users hitting the system, logs are a chaotic stream. By adding a `traceId` or `correlationId` to the MDC (Mapped Diagnostic Context) in my logging framework, every log line automatically includes that ID.

Then I can just search `traceId=12345` in Splunk/Kibana and see the entire story of that request across all services in chronological order."
