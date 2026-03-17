# Priority Category 5: Monitoring & Observability

### 1. How do you integrate Prometheus with Go applications?
**Your Response:**
"I use the official `prometheus/client_golang` library. 
I expose a `/metrics` endpoint on a separate HTTP port. 
Inside the code, I use middleware to automatically capture the 'Golden Signals'—Requests, Errors, and Duration—for every HTTP route. This gives us zero-effort monitoring for every new API we build."

### 2. Explain custom metrics collection in Go services.
**Your Response:**
"Standard metrics tell you if the server is healthy; custom metrics tell you if the *business* is healthy.
For example, I'll define a `prometheus.Counter` for 'Orders Processed' or a `prometheus.Gauge` for 'Active Users.' 
I increment these whenever the business logic executes successfully. This allows our SRE team to build dashboards that show real-time business performance, not just CPU charts."

### 3. How do you implement structured logging in Go?
**Your Response:**
"Plain text logs are an anti-pattern in microservices. I use `uber-go/zap` for **Structured JSON Logging**. 
Every log entry includes fields like `service_name`, `error_code`, and most importantly, `trace_id`. 
This allows us to aggregate millions of logs in **Loki** and query them instantly. I can find every log related to a specific user across ten different services in seconds."

### 4. Design Grafana dashboards for Go microservices.
**Your Response:**
"My ideal dashboard has three levels:
1. **Service Health**: Uptime, CPU/Memory consumption, and Pod restarts.
2. **API Performance**: The RED metrics (Rate, Errors, Duration) for every endpoint.
3. **Business Impact**: Transaction success rates and custom business KPIs.
I design these to be 'At-a-glance'—if a panel turns red, the engineer should know exactly which service is struggling and why."

### 5. How do you integrate Fluentd/Loki for log aggregation?
**Your Response:**
"I use a 'Sidecar' or 'Node-level' agent. 
The Go app simply writes JSON to `stdout`. 
A collector (like Fluent-bit) scrapes those logs, enriches them with Kubernetes metadata (like pod name and namespace), and ships them to **Loki**. 
This separation is key: the Go app doesn't need to know where its logs are going; it just focuses on producing high-quality telemetry."

### 6. Explain distributed tracing in Go microservices.
**Your Response:**
"In microservices, a single bug might involve five different services. I use **OpenTelemetry (OTel)** to solve this. 
Every request starting at the gateway is tagged with a `trace_id`. 
My Go services propagate this ID in their headers. When a request is slow, I look it up in **Jaeger** and see the entire 'waterfall' diagram. I can immediately see that 'Service A' was fast, but its call to 'Service B' was delayed by a slow DB query."
