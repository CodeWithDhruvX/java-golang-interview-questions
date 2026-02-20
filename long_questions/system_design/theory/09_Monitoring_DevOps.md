# 🔴 Monitoring, Logging & DevOps — Questions 81–90

> **Level:** 🔴 Senior / SRE
> **Asked at:** Google (SRE role), Netflix, LinkedIn, Amazon (SRE/Platform teams), Hotstar

---

### 81. How to monitor a distributed system?
"Monitoring a distributed system requires visibility into three signal types — the **three pillars of observability**: Metrics, Logs, and Traces.

**Metrics** answer 'is the system healthy right now?' — P99 latency, error rate, throughput, CPU/memory. **Logs** answer 'what happened?' — structured event logs for auditing and debugging. **Traces** answer 'where did this specific request go?' — a distributed trace shows the path of a single request across 10 microservices: Service A called B at 12ms, B called DB at 45ms, B called C at 8ms.

No single tool gives you all three. My standard stack: **Prometheus + Grafana** for metrics, **ELK Stack (Elasticsearch, Logstash, Kibana)** or Loki for logs, **Jaeger or Zipkin** for distributed tracing."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Google SRE, Netflix (famous for their observability tooling), LinkedIn, Amazon

#### Indepth
Observability stack architecture:
```
Services → [Prometheus scrape] → Prometheus → Grafana dashboards
        → [structured logs] → Fluentd/Filebeat → Elasticsearch → Kibana
        → [trace spans] → OpenTelemetry SDK → Jaeger/Tempo → Grafana
```

Golden Signals (Google SRE's framework — 4 metrics to monitor):
1. **Latency:** Time to serve a request. Distinguish successful vs error latency (an error returning in 1ms skews the healthy latency metric).
2. **Traffic:** Requests per second (QPS/RPS / TPS). A sudden drop is as alarming as a spike.
3. **Errors:** Rate of failed requests (500s, timeouts, explicit errors).
4. **Saturation:** How "full" is the service? CPU %, memory %, queue depth, DB connection pool utilization. Predicts degradation before it happens.

**USE Method (for resource monitoring):** Utilization, Saturation, Errors. Apply to every resource: CPU, Memory, Disk, Network, DB connections.

---

### 82. What are metrics, logs, and traces?
"The three pillars of observability — each answers a different question about system behavior.

**Metrics** are numerical measurements aggregated over time. CPU at 87%, 1200 requests/second, P99 latency 340ms. They're compact, fast to query, and perfect for dashboards and alerting. Prometheus stores metrics as time series.

**Logs** are discrete structured events: `{ timestamp, severity, service, message, context }`. They answer 'what exactly happened' when something went wrong. They're verbose and expensive to store/query at scale — use them purposefully, not with `console.log` sprinkled everywhere.

**Traces** link an entire request journey across services using a **trace_id**. Every service call, DB query, and cache hit within one user request is a span. The trace shows the timeline, duration, and parent-child relationships of all spans. Without traces, debugging 'why is this API slow?' in a microservices system is guesswork."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Any company running microservices — Google, Amazon, Flipkart DevEx/SRE, Swiggy, Hotstar

#### Indepth
OpenTelemetry is the emerging standard that unifies all three:
- **OTLP (OpenTelemetry Protocol):** Single SDK to emit metrics, logs, and traces
- **Collector:** Vendor-neutral agent that receives, processes, and exports to backends
- **Backends:** Any combination of Prometheus/Thanos (metrics), Loki/ELK (logs), Jaeger/Tempo (traces)

Structured logging best practices:
```json
{
  "timestamp": "2025-02-20T09:15:30Z",
  "level": "ERROR",
  "service": "payment-service",
  "trace_id": "a3f4b2c1d0e5...",
  "span_id": "b1c2d3e4f5...",
  "user_id": "user_123",
  "order_id": "ord_456",
  "message": "Payment gateway timeout",
  "duration_ms": 5002,
  "error_code": "GATEWAY_TIMEOUT"
}
```

Correlating logs with traces: include `trace_id` in every log line. Jump from a Grafana trace to the exact log lines for that specific request. This is the killer feature of unified observability.

---

### 83. What is Prometheus and Grafana?
"Prometheus is a **pull-based metrics collection and storage system** built for monitoring cloud-native and microservices applications. Grafana is the **visualization layer** that turns Prometheus data into dashboards.

Prometheus works by **scraping** — every N seconds (default 15s), Prometheus HTTP-GETs a `/metrics` endpoint on each service and stores the returned time series data in its local TSDB. Services expose metrics in a standard text format. This pull model means Prometheus knows about dead services (scrape fails) — no need for special failure detection.

PromQL (Prometheus Query Language) is what powers Grafana dashboards. A query like `rate(http_requests_total{status="500"}[5m])` gives me the 5-minute rolling rate of 500 errors, grouped by any label."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Companies adopting Kubernetes and microservices — Swiggy, Meesho, Razorpay, Google (contributed to Prometheus), CNCF ecosystem companies

#### Indepth
Prometheus ecosystem:
- **Exporters:** Special agents that translate existing systems' metrics into Prometheus format. `node_exporter` for Linux host metrics, `postgres_exporter` for PostgreSQL, `jmx_exporter` for JVM.
- **Service Discovery:** Prometheus integrates with Kubernetes — automatically discovers pods with annotation `prometheus.io/scrape: "true"`. No manual config per service.
- **Alertmanager:** Routes Prometheus alerts to PagerDuty, Slack, email. Supports grouping (consolidate 50 alerts into one), silencing (suppress alerts during maintenance), and inhibition (don't alert about downstream effects if root cause is already alerting).
- **Thanos / Cortex / VictoriaMetrics:** Prometheus is single-server by default. For multi-region or multi-cluster, Thanos adds global query capability, long-term storage (S3), and high availability. VictoriaMetrics is a high-performance drop-in replacement.
- **Grafana:** Organizations create shared dashboards for each service. USE methodology dashboards per infrastructure component. RED methodology (Rate, Errors, Duration) dashboards per service.

---

### 84. What is centralized logging?
"Centralized logging aggregates log output from all services and servers into a single queryable store. Without it, debugging requires SSHing into individual servers and grepping — impossible at the scale of 100+ microservices.

My standard setup: services write structured JSON logs to stdout → a log shipper (Fluentd, Filebeat, or Fluent Bit as a DaemonSet in Kubernetes) reads those logs and ships to **Elasticsearch** → **Kibana** provides the query interface. This is the ELK or EFK stack.

A query like 'all errors for user_id=123 in the last hour, across all services, sorted by timestamp' becomes trivial. Without centralized logging, it's impossible."

#### 🏢 Company Context
**Level:** 🟡 Mid – 🔴 Senior | **Asked at:** Any company with multiple services — Swiggy, Meesho, Razorpay, Hotstar

#### Indepth
Logging architecture alternatives:
- **ELK Stack (Enterprise):** Elasticsearch (storage/search) + Logstash (transform/filter) + Kibana (UI). Expensive at scale (Elasticsearch costs RAM).
- **EFK Stack (Kubernetes-native):** Replace Logstash with Fluentd (lighter, Kubernetes DaemonSet). Fluentd collects pod logs from `/var/log/containers`.
- **Grafana Loki (Cloud-native):** "Prometheus for logs." Indexes only metadata labels (not full text) — much cheaper storage than Elasticsearch. Log content searched via regex (slower for full-text). Best for Kubernetes where labels are the natural index.
- **AWS CloudWatch Logs / Google Cloud Logging:** Managed solutions. No operations overhead.

Log pipeline in Kubernetes:
```
Pod stdout → /var/log/containers/{pod}.log → Fluent Bit DaemonSet
             → Kafka (buffer, prevent log loss if ES is slow)
             → Logstash (transform, parse, enrich)
             → Elasticsearch (5-day hot retention → 30-day warm → S3 archive)
             → Kibana (search, dashboards, alerts)
```

**Log sampling:** At very high throughput (100K logs/second), storing everything is expensive. Sample non-error logs (store 1 in 100). Always store 100% of ERROR and WARN logs. This reduces storage 100x while keeping full visibility for anomalies.

---

### 85. How to detect system bottlenecks?
"A bottleneck is the component that limits the overall system's throughput — the slowest step in the chain. Detecting it requires systematic profiling at every layer.

My approach: start with **symptoms** (high latency, error spikes, CPU/memory saturation) from dashboards. Then drill down with **distributed traces** — they show which service and which operation within it is taking the most time. If the trace shows the DB span is 400ms of a 500ms request, the DB is the bottleneck.

Then profile the bottleneck: is it slow queries (use `EXPLAIN ANALYZE`)? Lock contention? Missing indexes? I/O bound? CPU bound? Memory pressure causing swapping? Each root cause has a specific fix."

#### 🏢 Company Context
**Level:** 🔴 Senior / SRE | **Asked at:** Any company doing performance engineering — Google perf teams, Netflix, Amazon (Prime Day prep), Hotstar (pre-IPL launch load testing)

#### Indepth
Systematic bottleneck detection methodology:

1. **Resource utilization check (USE):**
   - CPU: `top`, `htop`, `mpstat`. High CPU → profile with `pprof` (Go), `async-profiler` (Java)
   - Memory: `free -h`, heap dumps. High GC pressure (Java) → memory leak or large allocation
   - Disk I/O: `iostat -x 1`. High `%util` → DB or log writes saturating disk
   - Network: `iftop`, `ss -ti`. High retransmissions → network saturation or packet loss

2. **Application-level profiling:**
   - Go: `go tool pprof http://localhost:6060/debug/pprof/profile` — CPU flame graphs
   - Java: async-profiler + JFR. N+1 query detection in ORM framework.

3. **Database profiling:**
   - PostgreSQL: `pg_stat_statements` — shows cumulative query time. `EXPLAIN (ANALYZE, BUFFERS)` for individual query plans.
   - Slow query log (MySQL): queries exceeding `long_query_time` threshold.

4. **Load testing + incremental capacity:**
   - Use Gatling, k6, or Locust. Ramp up from 100 to 10K concurrent users. Watch where metrics degrade. The degradation point is the bottleneck.
   - Hotstar load-tested for 10M concurrent streams before IPL (Indian Premier League) 2018 after the previous year's crash.

---

### 86. How do you debug a distributed system?
"Debugging a distributed system is fundamentally different from single-process debugging because the bug might involve timing between services, a race condition across nodes, or an intermittent network issue.

My first tool is **distributed tracing** — find the trace_id for the failing request and see the full span tree. This immediately shows which service is failing or slow. Then query the centralized logs filtered by that trace_id to see the exact log lines from every service involved.

If it's a recurring issue, I look for patterns in metrics: does it correlate with a specific time? A specific upstream service? A specific data pattern? Correlation usually reveals the root cause. If it's intermittent and hard to reproduce, I add more structured logging around the suspected area with the specific context (user_id, request payload summary)."

#### 🏢 Company Context
**Level:** 🔴 Senior / Principal | **Asked at:** Google SRE, Netflix, LinkedIn — distributed systems debugging is a key SRE skill

#### Indepth
Distributed debugging toolkit:
1. **Correlation IDs:** Every request gets a unique ID (UUID or trace_id) at the API gateway. Propagate it through all service calls as a header (`X-Trace-ID`). Log it in every service. Enables joining logs from 10 services for one request.
2. **Distributed tracing (Jaeger/Zipkin/Tempo):** Visualize the call tree. See timing gaps between services (is the network slow? or service slow?). Find the "problem span."
3. **Structured logs + centralized search:** Kibana / Grafana Loki query by trace_id to read every log event from every service for that specific request.
4. **Chaos engineering to reproduce:** If the bug only happens under high load, reproduce it in a load test. k6 or Gatling can simulate the exact traffic pattern.
5. **Canary analysis:** Deploy new version to 1% of traffic. Compare error rates, latency between canary and production. If canary is bad, rollback with minimal blast radius.

**Death by a Thousand Microservices (distributed debugging nightmare):** When 15 services are involved, a single request failure requires checking 15 trace spans and 15 log sources. This is the real cost of microservices. Good observability tooling (single pane of glass in Grafana) is not optional — it's what makes microservices operable.

---

### 87. What is canary deployment?
"Canary deployment is a technique for **gradually rolling out a new version of software** to a small subset of users before deploying to everyone — 'canary in a coal mine' metaphor.

The process: deploy the new version to 1% of servers (or route 1% of traffic to it). Monitor closely — compare error rates, latency, and business metrics against the existing 99%. If the canary looks healthy, gradually expand: 5% → 10% → 25% → 50% → 100%. If the canary shows problems, immediately route 100% back to the old version.

Canary deployments reduced Netflix's incident rate significantly. Instead of discovering a bug when it hits 100% of users, you find it at 1% — when impact is minimal and rollback is trivial."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Netflix, Amazon (Flywheel), Google, Flipkart, Swiggy — any company deploying frequently

#### Indepth
Canary implementation methods:
1. **Load balancer weight shifting:** LB routes 99% to stable, 1% to canary. Gradually shift weights. Nginx `upstream` weight directive or AWS ALB weighted target groups.
2. **Kubernetes canary (Ingress-based):** Deploy canary as a separate Deployment with fewer replicas. Ingress rule routes small % to canary service. Istio/Flagger automates this with metrics-based promotion.
3. **Feature flags:** New code is deployed everywhere, but gated behind a flag that's enabled for 1% of users. No infra-level routing needed — just flag evaluation per request.
4. **Flagger (automatic canary):** K8s operator that automates canary releases. Analyzes Prometheus metrics. If error rate > threshold: automatically rolls back. If metrics look good for N minutes: promotes to 100%.

**Canary vs A/B Testing:**
- Canary: releasing same feature to % of users — detect bugs. Risk management.
- A/B Testing: releasing two different versions — measure which performs better on business metrics. Product experimentation.
These are different purposes but often use similar infrastructure.

---

### 88. What is blue-green deployment?
"Blue-green deployment maintains **two identical production environments** — 'blue' (current live) and 'green' (the new version being staged). When ready to deploy, you switch the router to send 100% of traffic to green. Blue becomes idle.

The advantage: rollback is instantaneous — switch traffic back to blue. There's no gradual ramp-up (unlike canary). You test green with production traffic in one shot. This is perfect for database migrations or releases that are hard to run in two versions simultaneously.

The disadvantage: you need double the infrastructure (both environments must be production-sized and ready). Also, database changes between versions must be backward-compatible — blue and green might briefly share the same DB during switchover."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Amazon (AWS pioneered this), Netflix, any CI/CD-mature company

#### Indepth
Blue-Green implementation:
- **DNS-based switch:** Update DNS CNAME from blue's IP to green's IP. Instant but DNS propagation latency (based on TTL). Risk: cached DNS entries still hitting blue briefly.
- **Load balancer switch:** Update ALB target group from blue cluster to green cluster. Truly instant (no DNS caching).
- **AWS CodeDeploy Blue/Green:** Automated — creates green ASG (Auto Scaling Group), deploys, runs health checks, shifts LB, terminates blue.
- **Kubernetes Blue/Green:** Two Deployments (blue and green). Switch K8s Service selector: `kubectl patch service myapp -p '{"spec":{"selector":{"version":"green"}}}'`.

**Database challenge:** If the green version requires a schema change (e.g., adds a non-null column), the blue version won't know about it. During the brief period both might run → incompatibility. Solution: always do additive, backward-compatible DB changes first (in a prior deployment), then switch app code.

**Blue-Green + Canary combination:** Some teams use both — do a canary (1% → 100%) within the new (green) environment to validate before the blue→green cut.

---

### 89. How do you handle configuration in distributed systems?
"Configuration management in distributed systems is far more complex than a single application's `.env` file — you need centralized, versioned, auditable config that can be changed at runtime without restarts.

My preferred approach: a **centralized config service** (Consul, etcd, AWS Parameter Store, or HashiCorp Vault) stores all configuration. Services fetch config at startup and subscribe to change notifications. When config changes, services hot-reload without restart.

The critical rule: **never hardcode secrets in code or environment variables** in plain text. Passwords, API keys, and certificates go in a secrets manager (HashiCorp Vault, AWS Secrets Manager) with fine-grained access control and rotation."

#### 🏢 Company Context
**Level:** 🔴 Senior | **Asked at:** Large-scale microservices companies — Amazon, Google, Flipkart infra teams, Swiggy, Razorpay

#### Indepth
Configuration tiers:
1. **Static config (build-time):** Baked into Docker image. `VERSION=2.1.3`. Never changes at runtime.
2. **Environment config (deploy-time):** Environment variables set at container start. Database URLs, service endpoints. Changed by re-deploying container.
3. **Dynamic config (runtime):** Feature flags, rate limits, timeouts. Stored in config service. Services watch for changes and apply without restart. Redis, Consul, etcd, LaunchDarkly.
4. **Secrets (runtime, encrypted):** Database passwords, API keys, TLS certificates. Vault injects secrets as env vars or files at pod startup. Secrets rotated automatically; pods re-fetched new secrets.

**Feature flags as config:** Feature flags (LaunchDarkly, Unleash) are a form of dynamic config. Toggle features per user, per region, per % of traffic — without deployment. This is the safest way to release high-risk features: deploy the code (dark launch), enable the flag gradually.

**Config as Code:** Store config in Git (GitOps). Config changes go through PR review. Automated synchronization (ArgoCD, Flux) applies config to the cluster. Full audit trail, rollback via git revert.

---

### 90. What is chaos engineering?
"Chaos engineering is the **practice of deliberately injecting failures into a production system** to verify that it's genuinely resilient — not just theoretically resilient.

Netflix invented the concept with **Chaos Monkey** — a service that randomly kills production EC2 instances during business hours. The reasoning: if production instances are going to die in random ways at random times, better to know about resilience gaps *before* a real outage than during one.

The principle: 'Chaos engineering is the discipline of experimenting on a system in order to build confidence in the system's capability to withstand turbulent conditions in production' (Principles of Chaos Engineering, 2016)."

#### 🏢 Company Context
**Level:** 🔴 Senior / Principal | **Asked at:** Netflix (invented it!), Amazon, Google (DiRT — Disaster Recovery Testing), LinkedIn, Flipkart

#### Indepth
Chaos Engineering progression:
1. **Chaos Monkey:** Kills random EC2 instances. Tests: does auto-scaling replace them? Does the LB detect and reroute? Do microservices handle the lost dependency?
2. **Chaos Gorilla:** Takes out an entire AWS Availability Zone. Tests multi-AZ redundancy.
3. **Chaos Kong:** Takes out an entire AWS Region. Tests cross-region failover (Netflix's most extreme test).
4. **Latency Injection:** Adds artificial latency (e.g., +500ms) to calls between specific services. Tests circuit breakers and timeout configurations.
5. **Fault Injection:** Returns errors from specific services to test error handling and fallback mechanisms.

Tools:
- **Netflix Simian Army** (Chaos Monkey, Janitor Monkey, Conformity Monkey, etc.)
- **Gremlin:** SaaS chaos engineering platform. Fine-grained fault injection (CPU stress, memory, network, time-skew).
- **Chaos Toolkit:** Open source. YAML-defined experiments.
- **Litmus (Kubernetes):** K8s-native chaos experiments. Pod kill, node drain, network partition within a cluster.
- **AWS Fault Injection Simulator (FIS):** Managed chaos engineering service from AWS.

**Game Days:** Scheduled exercises where the team deliberately breaks production systems and practices recovery procedures. Not automated — run by a designated "chaos team." Tests team response, runbooks, and communication as much as the technical resilience.
