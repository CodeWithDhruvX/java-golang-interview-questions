# 🟤 Monitoring, Logging & Observability

---

### 1. How do you monitor a Kubernetes cluster?

"A production K8s monitoring stack has three layers:

1. **Infrastructure metrics**: Node CPU, memory, disk — via Node Exporter + Prometheus.
2. **K8s metrics**: Pod restarts, deployment status, HPA behavior — via kube-state-metrics + Prometheus.
3. **Application metrics**: Your custom business metrics — exposed via `/metrics` endpoint and scraped by Prometheus.

I visualize everything in **Grafana** with pre-built dashboards (kube-prometheus-stack includes them) and alert via Alertmanager → PagerDuty/Slack."

#### In Depth
The `kube-prometheus-stack` Helm chart is the de facto standard — it bundles Prometheus Operator, Alertmanager, Grafana, Node Exporter, kube-state-metrics, and a set of default rules and dashboards. The Prometheus Operator introduces `ServiceMonitor` and `PodMonitor` CRDs that make scrape target configuration much cleaner than raw Prometheus config files.

---

### 2. What is Prometheus?

"Prometheus is a **time-series database and monitoring system** that scrapes metrics from HTTP endpoints.

It uses a pull model: Prometheus discovers targets (via service discovery — Kubernetes API integration is built-in), scrapes their `/metrics` endpoints at a configured interval, and stores the data in its TSDB.

I use it for everything from cluster health (node memory, pod restarts) to application-level metrics (API request latency, order counts, queue depth)."

#### In Depth
Prometheus's **PromQL** is incredibly powerful for correlating metrics: `sum(rate(http_requests_total[5m])) by (service)` gives per-service request rate. The federation and remote-write features allow aggregating metrics across multiple clusters. For long-term storage, I write to Thanos or Cortex (now migrated to Mimir), which are horizontally scalable, long-term Prometheus-compatible backends.

---

### 3. What is Grafana?

"Grafana is a **visualization and dashboarding platform** that queries Prometheus (and many other datasources) and displays metrics as graphs, heatmaps, and tables.

I build dashboards for every service: one showing the 4 golden signals (latency, traffic, errors, saturation), one for infrastructure health, and one for SLO tracking.

Grafana Alerting (native alerts) can replace Alertmanager for simpler setups, though Alertmanager is more powerful for routing/silencing."

#### In Depth
Grafana supports **provisioning** — dashboards and datasources defined as JSON/YAML files that are loaded at startup. This enables GitOps for dashboards: store dashboard JSON in Git, deploy via ConfigMap, and Grafana picks them up automatically. The `grafana-operator` CRD approach is even cleaner. Grafana also supports multi-tenant setups with organizations and teams for access control.

---

### 4. How do you collect logs in Kubernetes?

"The standard patterns:

1. **Sidecar**: A log-shipper sidecar (Fluent Bit) reads log files from a shared volume and ships to Elasticsearch/Loki.
2. **DaemonSet**: Fluent Bit runs on every node, reads container logs from `/var/log/containers/` and ships them centrally.
3. **Application-level**: App writes structured JSON to stdout, Fluent Bit collects and parses it.

I prefer option 2 (DaemonSet) for simplicity and option 3 (structured logging) for log quality. Never log to files inside a container — stdout/stderr is the K8s-native way."

#### In Depth
Container logs are written by the container runtime to `/var/log/containers/` on the node (symlinked from `/var/log/pods/`). The kubelet manages log rotation via `--container-log-max-size` (default 10Mi) and `--container-log-max-files` (default 5). Logs older than the limit are deleted — this is why persistent log shipping to a central system is critical. Without it, logs from crashed pods are permanently lost.

---

### 5. What is Fluentd?

"Fluentd is a **log collection and routing daemon** that reads logs from various sources and routes them to various destinations.

In K8s, it runs as a DaemonSet, tails container logs, parses them (JSON, regex, etc.), enriches them with K8s metadata (pod name, namespace, labels), and ships to Elasticsearch, Loki, Splunk, S3, etc.

**Fluent Bit** is Fluentd's lightweight sibling — I prefer it for K8s log collection because it uses ~10x less memory while doing 95% of what Fluentd does. Use Fluentd when you need complex log processing."

#### In Depth
Fluent Bit's input plugins support reading from `/var/log/tail`, systemd journal, and directly from the Docker/containerd log driver. Its Kubernetes filter adds metadata. The multiline filter handles stack traces that span multiple log lines. For high throughput, tune Fluent Bit's `Flush` and `Grace` settings and use asynchronous outputs with retry logic to handle backpressure from the destination.

---

### 6. What is the EFK stack?

"EFK = **Elasticsearch + Fluentd/Fluent Bit + Kibana**.

- **Elasticsearch**: Distributed search and analytics engine that stores and indexes logs.
- **Fluentd/Fluent Bit**: Log collector and shipper (DaemonSet on each node).
- **Kibana**: Visualization UI for exploring and searching logs.

This was the standard K8s logging stack for years. Today, many teams are replacing it with **Grafana Loki** because Loki is much cheaper — it indexes only labels (not full-text), stores compressed log chunks in object storage (S3), and integrates natively with Grafana."

#### In Depth
Elasticsearch's main cost in K8s: it needs persistent storage (ILM policies to control disk usage), significant memory (Elasticsearch JVM heap), and careful cluster sizing. For compliance requirements (log retention for 7 years), EFK is still common. For operational logs with short retention (30 days), Loki is dramatically cheaper and simpler to operate.

---

### 7. What is Loki in Kubernetes logging?

"Loki is a **horizontally scalable, highly available log aggregation system** from Grafana Labs.

Philosophy: index only labels (like Prometheus), stream the log content as compressed chunks to object storage. This makes it much cheaper than full-text indexing (Elasticsearch) while being fast for label-filtered queries.

I query Loki with **LogQL**: `{app="payments-api", level="error"} |= "timeout"` to find error logs with 'timeout' in them. Grafana has native Loki support."

#### In Depth
Loki's architecture: **Promtail** (or Fluent Bit Loki output) ships logs. **Loki** ingests, compresses, and stores chunks in S3/GCS. **Querier** retrieves and queries chunks. The chunk format makes Loki very storage-efficient (10-30x cheaper than ELK for the same data). The limitation: no full-text inverted index means searches that aren't label-filtered require scanning all chunks — slow for needle-in-haystack searches.

---

### 8. What is kube-state-metrics?

"kube-state-metrics generates **Prometheus metrics about the state of Kubernetes objects** (not about container resource usage).

Examples of metrics it exposes: `kube_pod_status_ready`, `kube_deployment_status_replicas_unavailable`, `kube_node_status_condition`. These let you alert on 'any deployment has unavailable replicas' or 'node is NotReady'.

It runs as a single Deployment and connects to the API server to watch resources. It's essential for cluster-level alerting in Prometheus."

#### In Depth
kube-state-metrics differs from the Metrics Server: kube-state-metrics reports on the K8s object model (is this deployment healthy?), while Metrics Server reports on resource consumption (how much CPU is this pod using?). Both are needed in a complete monitoring setup. kube-state-metrics also exports labels from K8s resources as Prometheus labels — this allows hierarchical queries correlating application labels to infrastructure.

---

### 9. How do you implement alerting in Kubernetes?

"Complete alerting pipeline: **Prometheus collects metrics → Alertmanager receives firing alerts → routes to correct channels**.

I define `PrometheusRule` CRDs with expressions like:
```yaml
- alert: HighPodRestartRate
  expr: rate(kube_pod_container_status_restarts_total[15m]) > 0.1
  for: 5m
  labels: {severity: warning}
```

Alertmanager routes by labels: `severity: critical` → PagerDuty (wake someone up), `severity: warning` → Slack channel."

#### In Depth
Alert quality matters enormously. Common mistakes: too many alerts (alert fatigue), alerting on symptoms instead of causes, missing `for` clause (fires immediately on one data point). I follow the **USE method** (Utilization, Saturation, Errors) for infrastructure and **RED method** (Rate, Errors, Duration) for services. Every alert should have a runbook link in its annotations describing how to respond.

---

### 10. What is OpenTelemetry?

"OpenTelemetry (OTel) is the **CNCF standard for application observability** — it provides SDKs for collecting traces, metrics, and logs in a vendor-neutral way.

Instead of instrumenting for Jaeger OR Datadog OR Dynatrace, you instrument once with OTel, and then configure exporters to send to any backend.

I use OTel in Go services to trace service-to-service calls with actual latency breakdowns. When a request times out, I can see exactly which downstream call caused the delay."

#### In Depth
OTel's architecture: SDK (in your app) → OTLP protocol → OTel Collector (deployed as DaemonSet or sidecar) → backends (Jaeger, Prometheus, Loki, commercial APMs). The Collector handles batching, retry, and routing — isolating your app code from observability backend changes. Auto-instrumentation via eBPF (Odigos, Pixie) is emerging as a zero-code option that instruments pods at the kernel level.

---

### 11. What is the difference between whitebox and blackbox monitoring?

"**Whitebox monitoring**: The application exposes its internal state (metrics, traces, logs). You know if a request failed because the app told you. Examples: Prometheus metrics from `/metrics`, structured logs, distributed traces.

**Blackbox monitoring**: External probes test the application from the outside — like a real user. You know the system is down because a probe can't reach it. Examples: Prometheus blackbox-exporter (HTTP checks, DNS checks, TCP checks).

In production I use both: whitebox for deep diagnosis, blackbox as the 'is it alive from outside?' safety net. The blackbox-exporter is my final guard — if it fires, users are definitely impacted."

#### In Depth
Blackbox monitoring catches issues that whitebox misses: DNS failures, load balancer misconfigurations, TLS certificate expiry, and network path failures. The Prometheus blackbox-exporter runs HTTP/HTTPS/DNS/TCP/ICMP probes and exposes the results as Prometheus metrics. It can check SSL certificate validity days ahead of expiry — a critical production alert.

---

### 12. How do you correlate logs and traces in Kubernetes?

"The standard approach is **trace ID injection**: the tracing SDK generates a trace ID for each request. This trace ID is added as a structured field in every log line emitted during that request.

With Grafana, I can go from a Loki log showing an error → click the trace ID → jump to Jaeger/Tempo showing the full distributed trace.

This makes debugging production issues dramatically faster — you go from 'there was an error' to 'the gRPC call to payments-service took 5 seconds because it was waiting on a DB lock'."

#### In Depth
The OpenTelemetry SDK provides a `trace.SpanFromContext(ctx).SpanContext()` method to extract trace and span IDs from the current context. Structured logging libraries like `zap` or `slog` can inject these as fields. In Grafana, the **Derived Fields** feature in Loki datasource config automatically creates clickable links from log fields containing trace IDs to your Tempo/Jaeger datasource.

---
