# 🟢 **121–135: Observability**

### 121. What is logging?
"Logging is the process of recording discrete events that happen within an application during runtime. 

If my microservice receives a request to process a payment, I write a log: `INFO: User 123 initiated payment for $50`. If the payment fails because the database is unreachable, I write: `ERROR: Failed to connect to DB at 10.0.0.5`.

I rely on logs entirely for debugging. Without them, deciphering why a specific user experienced a bug at 3:00 AM is impossible because the application's internal state is gone."

#### Indepth
Logs should be structured, ideally outputted natively in JSON format rather than plain text. This allows downstream log agents (like Filebeat or Fluentd) to instantly parse fields like `error_code` or `user_id` without relying on brittle, complex Regex parsing rules later.

---

### 122. What is centralized logging?
"Centralized logging aggregates the logs from hundreds of different microservices and physical servers into one single, searchable database.

In a microservices cluster, if I have 50 Order Service pods, I cannot manually SSH into 50 different Linux machines and run `tail -f` to find an error. 

Instead, every pod streams its logs to strictly `stdout`. A cluster-level agent captures this stream and ships it to a central system like the ELK stack (Elasticsearch, Logstash, Kibana) or Datadog. I can then open one web interface and search for a User ID across the entire company's log history."

#### Indepth
Log aggregation pipelines can consume immense network bandwidth and disk space (often generating terabytes of data daily). Strategies include aggressive log rotation, filtering out useless `DEBUG` noise at the origin, and dropping logs into cheap cold storage (like AWS S3) after 7 days for compliance auditing.

---

### 123. What is distributed tracing?
"Distributed tracing is a method to track a single user's request as it travels sequentially across multiple distinct microservices.

If a user clicks 'Checkout', the request hits the API Gateway, then Order Service, then Payment Service, then Inventory Service. If the request takes 5 seconds, how do I know which service was the bottleneck?

Tracing injects a unique `Trace ID` at the Gateway. Every service passes it along in the HTTP headers. Tools like Jaeger or Zipkin collect these traces and generate a beautiful Gantt chart showing exactly how many milliseconds each service spent processing that exact request."

#### Indepth
A Trace represents the entire journey. A Span represents a specific operation within that journey (e.g., "SQL Query to get User" or "HTTP call to Payment"). Proper distributed tracing relies heavily on W3C Trace Context standards to ensure `traceparent` headers propagate cleanly even across different programming languages.

---

### 124. What is metrics?
"Metrics are numerical measurements of an application's behavior recorded over time. 

Unlike logs (which record distinct text events), metrics track continuous absolute numbers: CPU utilization, available memory, active database connections, HTTP 500 error rates, or the number of messages currently in a Kafka queue.

I use metrics for real-time alerting and dashboards. If my 'HTTP 500 error metric' spikes from 1% to 15%, an alert fires to my phone. I then switch to the Logs to figure out *why* it spiked."

#### Indepth
Metrics are staggeringly cheap to store compared to logs because they are just time-series data (a timestamp, a label, and a float value). Systems like Prometheus utilize highly specialized Time-Series Databases (TSDB) capable of ingesting millions of data points per second with minimal CPU footprint.

---

### 125. What is APM?
"APM stands for Application Performance Monitoring. It is a comprehensive suite of tools that combines Logs, Metrics, and Distributed Tracing into a single 'single pane of glass'.

Vendors like Datadog, New Relic, or AppDynamics provide APM agents. I attach their Java agent to my Spring Boot JVM at startup. 

Without writing a single line of custom code, the APM mathematically analyzes my application's bytecode, automatically tracing database queries, detecting slow HTTP calls, and mapping out a gorgeous visual topology of how all my microservices connect to each other."

#### Indepth
While commercial APMs are incredibly powerful, their proprietary agents can sometimes introduce noticeable performance overhead (CPU/RAM spikes) or even cause application crashes due to aggressive bytecode instrumentation. Modern open-source alternatives like OpenTelemetry are rapidly becoming the vendor-neutral standard for APM data collection.

---

### 126. What is SLI?
"Service Level Indicator (SLI) is a carefully defined, quantitative measure of some aspect of the level of service that is provided. 

It is the actual factual metric I am observing in reality. 

For example, an SLI could be 'the percentage of HTTP GET requests that return a 200 OK status code' or 'the 99th percentile latency of the checkout API'. It is pure data—the raw numbers my Prometheus server is gathering."

#### Indepth
Good SLIs usually follow the "Four Golden Signals" defined by Google SREs: Latency, Traffic, Errors, and Saturation. If you cannot objectively measure an SLI, you cannot possibly hold yourself accountable to any customer guarantees.

---

### 127. What is SLO?
"Service Level Objective (SLO) is a target value or range of values for a specific service level, measured by an SLI.

It is the internal goal my engineering team strives for. 

If my SLI is 'percentage of successful HTTP requests', my SLO might be '99.9% of requests must be successful over a 30-day window'. If our actual SLI drops to 99.8%, the engineering team halts new feature development and focuses entirely on fixing bugs to restore the system's reliability."

#### Indepth
An SLO determines the 'Error Budget'. If your target is 99.9% uptime per month (roughly 43 minutes of allowed downtime), that 43 minutes is your budget. If a bad K8s deployment consumes 40 minutes of that budget on day 2, the team knows they must freeze non-critical deployments for the rest of the month.

---

### 128. What is SLA?
"Service Level Agreement (SLA) is an explicit or implicit business contract with your users that includes consequences if you miss the agreed-upon SLOs.

While SLOs are internal engineering goals, SLAs involve lawyers and money.

If my SaaS company signs an SLA guaranteeing a client '99.99% API Uptime', and we suffer a 2-hour outage, the SLA dictates that we owe that client a financial penalty or free service credits for breaching the contract."

#### Indepth
Because SLAs trigger financial penalties, business leaders always set the external SLA slightly more conservatively than the internal engineering SLO. If the SLA is 99.9%, the engineering SLO is internally enforced strictly at 99.95% to ensure a comfortable safety margin.

---

### 129. What is health check?
"A health check is a dedicated HTTP endpoint (usually `/health`) exposed by a microservice to report its current operational status.

Instead of writing complex monitoring scripts, an orchestrator (like Kubernetes) or a load balancer simply pings this endpoint every 5 seconds. 

If the endpoint returns `200 OK`, K8s keeps routing traffic to it. If the endpoint starts returning `503 Service Unavailable` (maybe because the app lost its database connection), K8s instantly removes it from the load balancer rotation, shielding users from seeing errors."

#### Indepth
In Spring Boot, the Actuator `spring-boot-starter-actuator` dependency handles this automatically. It exposes a `/actuator/health` endpoint that cumulatively checks the status of the database connection, Redis, and disk space natively out-of-the-box.

---

### 130. What is liveness vs readiness probe?
"These are two distinct health check concepts utilized by Kubernetes.

**Liveness Probe** answers: 'Is this application dead?' If the code encounters a deadlock and freezes indefinitely, the liveness probe fails. Kubernetes responds by brutally terminating the Pod and restarting it from scratch.

**Readiness Probe** answers: 'Is this application ready to receive user traffic?' If a Spring Boot application takes 15 seconds to boot up and warm its caches, the readiness probe fails for those first 15 seconds. K8s leaves the Pod alone but actively refuses to route any HTTP traffic to it until the probe returns `200 OK`."

#### Indepth
Confusing the two is a catastrophic mistake. If an application temporarily loses its database connection, it should fail its *Readiness* probe (so K8s stops sending it traffic), but it should NOT fail its *Liveness* probe. If liveness fails, K8s restarts the Pod violently, which accomplishes nothing but adding CPU strain since the database is still offline.

---

### 131. How to monitor using Prometheus?
"Prometheus is a time-series monitoring system that uses a 'pull-based' architecture.

Instead of my 50 microservices constantly pushing metrics *to* Prometheus, I expose an HTTP endpoint (`/metrics`) on every microservice containing raw text data about its CPU and memory. 

Prometheus acts like a web scraper. Every 15 seconds, it reaches out to all 50 IP addresses, reads that text data, and stores it efficiently. I find this incredibly robust because if Prometheus crashes, my microservices aren't affected—they don't care, they are just exposing a passive endpoint."

#### Indepth
PromQL (Prometheus Query Language) allows engineers to mathematically manipulate metrics. You don't just ask for 'Error Rate'; you write a query like `rate(http_requests_total{status="500"}[5m])` to dynamically calculate the per-second error rate over a 5-minute sliding window.

---

### 132. How to visualize using Grafana?
"Grafana is an open-source analytics and interactive visualization web application. It is almost always paired directly with Prometheus.

While Prometheus acts as the raw database holding billions of metric data points, Grafana acts as the beautiful frontend dashboard. 

I connect Grafana directly to my Prometheus data source. I can then drag and drop charts, creating a 'Microservice Overview Dashboard' that shows live line graphs of API latency, pie charts of HTTP response codes, and gauges showing JVM heap memory usage. It’s what goes on the TV monitors in the engineering office."

#### Indepth
Grafana extends aggressively beyond Prometheus. A single Grafana dashboard can simultaneously execute a PromQL query for live CPU graphs, query an Elasticsearch database to show a table of recent ERROR logs, and hit an AWS CloudWatch API to show SQS queue depth, acting as the ultimate aggregator.

---

### 133. What is tracing using Jaeger?
"Jaeger is an open-source, end-to-end distributed tracing system. 

When a trace is initiated at my API Gateway, OpenTelemetry code running in my microservices generates 'Spans' (metadata representing how long a specific method took) and asynchronously flushes these spans via UDP to a Jaeger Collector.

I open the Jaeger UI, paste in a specific Correlation ID, and Jaeger visually reconstructs the entire request timeline. If a request took 5 seconds, Jaeger clearly highlights that the Inventory Service took 10ms, whereas the Payment Service's specific SQL query took 4990ms, instantly pinpointing the exact line of code causing the latency."

#### Indepth
Jaeger architecture scales massively through Kafka. Instead of Jaeger Collectors writing directly to Cassandra/Elasticsearch databases (which could bottleneck), agents often push trace spans into Kafka topics. Jaeger Ingesters then pull from Kafka, buffering the massive write loads gracefully.

---

### 134. What is log aggregation?
"Log aggregation is the automated process of collecting logs from heterogeneous, distributed sources, parsing them, and normalizing them into a single, cohesive repository.

Because a simple 'Login' flow might traverse NGINX, a Node.js frontend layer, a Spring Boot backend, and a PostgreSQL database, their logs are all formatted entirely differently. 

An aggregation pipeline (like Logstash or Fluentd) grabs these diverse files. It uses regex (like Grok patterns) to extract the timestamp, severity level, and message string, converts them all into a standard JSON format, and pushes them into a search engine like Elasticsearch."

#### Indepth
Without log aggregation, debugging requires executing parallel SSH bash scripts using `grep` across dozens of machines simultaneously—an antiquated strategy totally incompatible with ephemeral Kubernetes pods that delete their logs the moment they crash.

---

### 135. What is correlation tracing?
"Correlation tracing is the practice of linking Log entries directly to Trace IDs. 

When observing a slow request in Jaeger, I see the Payment service took 4 seconds, but Jaeger only shows me the latency, not *why* it was slow. 

If my logging framework (like Logback in Java) is configured correctly, it automatically injects the Jaeger `TraceID` into every single log statement. In Datadog or Kibana, I can click exactly on that 4-second span, and it instantly filters my database to show me the exact 5 lines of application log output generated during that specific 4-second window."

#### Indepth
This is the holy grail of modern Observability (the unification of Metrics, Logs, and Traces). When a Metric alerts you to a spike, you find the Trace that is slow, and pivot seamlessly to the precise Logs associated with that Trace ID, vastly accelerating Mean Time To Resolution (MTTR) during production outages.
