# Prometheus Integration Interview Guide: Spoken Style

This guide provides conversational-style responses for the topics, interview questions, and practice problems related to Prometheus integration in Go microservices.

---

## 🔹 Topics: Core Concepts & Integration Patterns

### 1. Prometheus Client Library & Metric Types
**Interviewer:** "What are the primary metric types in the Prometheus Go client, and when do you use each?"

**Your Response:**
"In the Prometheus Go library, we work with four main types:
- **Counters**: These only go up (or reset to zero). I use them for things like `http_requests_total`.
- **Gauges**: These can go up and down. I use them for current snapshots, like `active_db_connections` or memory usage.
- **Histograms**: These sample observations (like request duration) into configurable buckets. They are great for calculating percentiles like the p99 latency.
- **Summaries**: Similar to histograms, but they calculate quantiles on the client-side. I generally prefer histograms because they allow for easier aggregation across multiple service instances."

### 2. HTTP Middleware for Metrics
**Interviewer:** "How do you implement metrics collection for your HTTP APIs without cluttering your business logic?"

**Your Response:**
"I use Middleware. I create a wrapper that starts a timer, calls the next handler in the chain, and then calculates the elapsed time. I capturing the status code by wrapping the `ResponseWriter`. This way, every single request—regardless of the business logic—automatically records its method, path, status code, and duration into Prometheus. It gives us consistent 'Red metrics' (Rate, Errors, Duration) across the entire service."

### 3. Database & Business Metrics
**Interviewer:** "Beyond standard HTTP metrics, what else should we be monitoring in a Go microservice?"

**Your Response:**
"I focus on two other areas: **Database Health** and **Business KPIs**.
For the DB, I instrument the connection pool—monitoring idle versus active connections—and I record the duration and error rates of specific SQL queries. 
For business metrics, I track 'domain events.' For example, in a checkout service, I'd have a counter for `orders_processed_total` with labels for `payment_method` and `status`. This helps the business team understand the actual value the service is delivering, not just its technical health."

---

## 🔹 Interview Questions: Technical & Design

### 1. How do you implement custom metrics in a Go microservice?
**Your Response:**
"I define my metrics as package-level variables using the `prometheus.New*` functions. During the `init()` phase or a setup function, I register them with the `prometheus.DefaultRegisterer`. Then, inside my logic, I just call methods like `.Inc()` for counters or `.Observe()` for histograms. I always make sure to use labels effectively to keep the metrics granular but manageable."

### 2. How do you handle metric aggregation across multiple microservice instances?
**Your Response:**
"Prometheus handles this on the server-side. Since it's a pull-based system, it scrapes each instance individually. In Grafana, I can then use PromQL to sum or average metrics across all instances of a service. For example, `sum(http_requests_total) by (service)` gives me the total load across the entire cluster, even if we have 50 pods running."

### 3. Design a metrics collection strategy for a microservices architecture.
**Your Response:**
"I'd follow the **'Golden Signals'** (Latency, Traffic, Errors, and Saturation). 
1. I'd use middleware for automated API metrics.
2. I'd use a custom collector for system-level metrics like goroutine counts and memory.
3. I'd implement business-specific counters.
4. Finally, I'd ensure every service exposes these on a `/metrics` endpoint and is discovered automatically via Kubernetes service discovery."

---

## 🔹 Practice Problems: Practical Coding Walkthroughs

### 1. Web Middleware Implementation
**Interviewer:** "Walk me through the code for a Go middleware that tracks HTTP metrics."

**Your Response:**
"I'd create a function that takes an `http.Handler` and returns one. Inside, I wrap the `http.ResponseWriter` in a custom struct to 'spy' on the status code. I capture the start time using `time.Now()`, call `next.ServeHTTP`, and then use `time.Since` to get the duration. Finally, I update a `CounterVec` with the status code and a `HistogramVec` with the duration. It's efficient and completely transparent to the rest of the application."

### 2. System Metrics Collector
**Interviewer:** "How would you create a collector for metrics that Prometheus's default collector doesn't cover?"

**Your Response:**
"I'd implement the `prometheus.Collector` interface, which requires two methods: `Describe` and `Collect`.
In `Describe`, I pass the metric descriptors to the channel. 
In `Collect`, I fetch the actual data—like reading `runtime.MemStats` or calling an OS command—and then 'send' the metric values via the channel. This allows me to expose any dynamic system state as a first-class Prometheus metric."

### 3. Database Query Monitoring
**Interviewer:** "How do you monitor database performance in a way that helps you find slow queries?"

**Your Response:**
"I wrap my database calls in a helper function or use a 'Repository' pattern. Before the query, I start a timer. After the query, I record the duration in a `HistogramVec` with the table name and the operation type (like 'INSERT' or 'SELECT') as labels. If there's an error, I increment a `CounterVec` labeled with the specific error type. This makes it trivial to spot which specific tables or query types are causing bottlenecks or failures."
