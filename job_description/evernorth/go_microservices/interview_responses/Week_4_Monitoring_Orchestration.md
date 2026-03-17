# Week 4: Monitoring, Orchestration & Integration

### 🔹 Topic: Prometheus & Observability

**Interviewer:** "How do you ensure your Go microservices are observable in production?"

**Your Response:**
"I follow the 'Three Pillars of Observability': **Logging, Metrics, and Tracing**.
For metrics, I integrate the **Prometheus** client library. I export standard metrics like request duration and error rates, plus custom 'business metrics'—like 'number of successful checkouts.' 

For logging, I use structured JSON logging (with libraries like `zap` or `zerolog`) so they are easily searchable in **Loki** or ELK. 
For tracing, I use **OpenTelemetry** to track a request as it hops across different services, which is vital for debugging latency in a microservices setup."

### 🔹 Topic: Netflix Conductor (Orchestration)

**Interviewer:** "When would you use a workflow orchestrator like Conductor versus simple inter-service calls?"

**Your Response:**
"I use Conductor when I have long-running processes or complex business workflows that involve multiple steps and potential failures. Simple REST calls work for quick requests, but if a process takes hours or needs complex 'retry-on-failure' logic, Conductor is better. 

In Go, I implement **Workers** that poll Conductor for tasks. This keeps my services decoupled; the service doesn't need to know the 'next step'—the orchestrator manages that. It also handles state management and retries automatically."

### 🔹 Interview Focus: Monitoring & Integration

**1. Types of metrics to track?**
**Your Response:** "I focus on the 'GOLDEN SIGNALS': **Latency** (response time), **Traffic** (RPM), **Errors** (4xx/5xx rates), and **Saturation** (CPU/Memory). I also add domain metrics that matter to the business."

**2. How to implement custom metrics?**
**Your Response:** "I define a `prometheus.Counter` or `prometheus.Histogram` at the top level of my package. Then, inside my business logic, I call `.Inc()` or `.Observe()` whenever the relevant event occurs. I make sure to register these with the Prometheus registry."

**3. Self-healing microservices?**
**Your Response:** "This is a combination of Kubernetes and Go logic. K8s handles restarts through health probes. In Go, I implement 'Graceful Shutdown.' When the service receives a SIGTERM, it stops accepting new requests, finishes the ones in progress, closes DB connections, and then exits. This prevents dropped requests during a rollout."

### 🔹 Week 4 Practice Problems: Spoken Walkthroughs

**1. Prometheus metrics collection system:**
"I'd create a middleware that wraps every HTTP handler. It records the start time, executes the handler, and then records the duration and status code in a Prometheus Histogram. I'd expose these at `/metrics` for Prometheus to scrape."

**2. Conductor workflow implementation:**
"I'd define a JSON workflow in Conductor. Then in Go, I'd use the Conductor SDK to write a worker that polls for a specific 'task type.' When it gets a task, it executes the logic and sends back a 'Completed' or 'Failed' status."

**3. End-to-end monitoring system:**
"I'd set up Prometheus to scrape my Go pods, Loki to collect logs via Fluentd, and a Grafana dashboard to visualize it all. The dashboard would have panels for the Golden Signals and a 'Health Status' light for each service."
