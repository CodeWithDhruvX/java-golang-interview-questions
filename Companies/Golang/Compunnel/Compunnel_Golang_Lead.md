# Compunnel Senior Lead Golang Developer Interview Questions & Answers

**Role:** Senior Lead Golang Developer  
**Focus:** Golang, Microservices, Cloud-Native, APIs, Distributed Systems  
**Excluded Topics:** AWS & GCP specifics (as requested)

---

## 1. Core Golang & Concurrency

### Q1: How does Golang handle concurrency differently than Java or Python? Explain Goroutines and the Scheduler.
**Answer:**
-   **Goroutines vs Threads:** In Java/Python, threads are OS-level and heavy (1-2MB stack). In Go, **goroutines** are user-space threads managed by the Go runtime, starting with a tiny stack (2KB) that grows dynamically. You can run thousands of goroutines on a few OS threads.
-   **M:N Scheduler:** Go uses an M:N scheduler where **M** goroutines are multiplexed onto **N** OS threads.
-   **G-M-P Model:**
    -   **G (Goroutine):** The code to execute.
    -   **M (Machine):** The OS thread.
    -   **P (Processor):** A resource required to execute Go code (manages the run queue of Gs).
-   **Work Stealing:** If a P runs out of Gs, it steals half of the Gs from another P's local run queue, ensuring efficient CPU utilization.

### Q2: Explain the "Worker Pool" pattern in Go and when you would use it.
**Answer:**
A Worker Pool is a pattern to limit the number of active goroutines processing tasks, preventing resource exhaustion (e.g., opening too many DB connections or overwhelming an API).
-   **Implementation:**
    1.  Create a buffered channel to hold tasks (jobs).
    2.  Spawn a fixed number of worker goroutines (e.g., 5 workers).
    3.  Each worker loops over the jobs channel, processing tasks as they arrive.
    4.  Close the jobs channel when all tasks are submitted.
    5.  Use a `sync.WaitGroup` to wait for all workers to finish.
-   **Use Case:** Processing a large queue of images, handling incoming HTTP requests where downstream services have rate limits, or batch processing database records.

### Q3: How do channels work? What is the difference between Buffered and Unbuffered channels?
**Answer:**
-   **Channels:** Pipes that connect concurrent goroutines. You send values into channels from one goroutine and receive into another.
-   **Unbuffered Channel:** Capacity is 0. Sending blocks until another goroutine receives (synchronous). Used for strict synchronization (handshake).
-   **Buffered Channel:** Has a capacity (e.g., `make(chan int, 100)`). Sending only blocks if the buffer is full. Receiving only blocks if the buffer is empty. Used to decouple producers and consumers and handle bursts of traffic.
-   **Closing:** Only the sender should close a channel. Receiving from a closed channel returns the zero value immediately. Sending to a closed channel causes a panic.

### Q4: Explain `context` in Go. How is it used for cancellation and timeouts?
**Answer:**
The `context` package is standard for passing deadlines, cancellation signals, and request-scoped values across API boundaries and between processes.
-   **`context.Background()`:** Base context, usually used in `main` or handling the initial request.
-   **`context.WithCancel(parent)`:** Returns a copy of parent with a `Done` channel that closes when `cancel` is called. Used to stop worker goroutines when the main process stops.
-   **`context.WithTimeout(parent, duration)` / `context.WithDeadline(parent, time)`:** Automatically cancels the context after a duration. Crucial for microservices to prevent hanging requests from consuming resources indefinitely (e.g., DB queries, external API calls).
-   **`context.WithValue(parent, key, val)`:** Passing request-scoped data like UserID or AuthTokens (use sparingly).

---

## 2. Microservices, APIs & Architecture

### Q5: You are migrating a legacy Java monolith to Golang Microservices. What strategies would you use?
**Answer:**
-   **Strangler Fig Pattern:** Gradually replace specific functionalities of the monolith with new microservices.
    1.  Identify a bounded context (e.g., "User Profile" or "Order History").
    2.  Build the new service in Go.
    3.  Use an API Gateway/Load Balancer to route traffic for that specific functionality to the new service while keeping everything else pointing to the monolith.
    4.  Repeat until the monolith is gone.
-   **Anti-Corruption Layer (ACL):** If the new Go service needs to talk to the legacy system, build an ACL adapter to translate between the old and new models, ensuring the new service's domain model remains clean.
-   **Dual Run:** For critical paths, run both old and new code, compare results, and only switch over when confidence is high.

### Q6: REST vs. gRPC vs. GraphQL. When to choose which?
**Answer:**
-   **REST (JSON/HTTP):**
    -   *Pros:* Standard, human-readable, wide tool support, easy to cache, stateless.
    -   *Cons:* Over-fetching/under-fetching, text-based (slower parsing), no strict schema enforcement.
    -   *Use Case:* Public APIs, simple CRUD services, web clients.
-   **gRPC (Protobuf/HTTP2):**
    -   *Pros:* Binary (compact, fast), strongly typed contracts (.proto), supports streaming (unidirectional/bidirectional), built-in code generation.
    -   *Cons:* Not browser-native (requires gRPC-Web), harder to debug (binary calls).
    -   *Use Case:* Internal microservice-to-microservice communication (low latency), polyglot environments.
-   **GraphQL:**
    -   *Pros:* Client requests exactly what it needs (no over-fetching), single endpoint, strongly typed.
    -   *Cons:* Complex caching, N+1 query problem complexity, can be overkill for simple APIs.
    -   *Use Case:* Mobile apps, heavy frontend data requirements, aggregating data from multiple sources (BFF - Backend For Frontend).

### Q7: How do you handle Distributed Transactions in Microservices?
**Answer:**
Distributed transactions (ACID across services) are hard because of the CAP theorem.
-   **Saga Pattern:** A sequence of local transactions. If one fails, execute compensating transactions (undo steps) for the previous successful steps.
    -   *Choreography:* Events trigger steps (Service A publishes "OrderCreated", Service B listens and does "ReserveStock"). Harder to track system state.
    -   *Orchestration:* A central "Orchestrator" service tells services what to do (Service A -> Orchestrator -> Service B). Easier to manage and visualize.
-   **Two-Phase Commit (2PC):** (Avoid if possible) Coordinator asks all services if they are ready (Prepare), then tells them to Commit. prone to blocking and scaling issues.

### Q8: What is the role of an API Gateway?
**Answer:**
An API Gateway acts as the single entry point for all client requests.
-   **Responsibilities:**
    -   **Routing:** Forwarding requests to the correct implementation service.
    -   **Authentication/Authorization:** Verifying JWTs centrally so each service doesn't have to.
    -   **Rate Limiting:** Protecting services from DDoS or overuse.
    -   **Protocol Translation:** Converting REST frontend requests to gRPC backend calls.
    -   **Response Aggregation:** Combining data from multiple services into one response.

---

## 3. Databases & Messaging

### Q9: SQL (PostgreSQL) vs. NoSQL (Cassandra/MongoDB). How do you choose?
**Answer:**
-   **PostgreSQL (SQL):**
    -   Use when data is structured (relational), requires strict ACID compliance (financial transactions), and complex JOINs are needed.
    -   Golang drivers: `pgx`, `database/sql`.
-   **MongoDB (NoSQL Document):**
    -   Use for flexible schemas, rapid prototyping, content management, or when data is hierarchical (JSON-like).
    -   Golang driver: `mongo-go-driver`.
-   **Cassandra (NoSQL Wide-Column):**
    -   Use for massive write scaling, high availability (peer-to-peer), and time-series data or activity logs where you write huge amounts of data and query by specific keys. Not good for ad-hoc queries or JOINs.

### Q10: How does Kafka fit into an Event-Driven Architecture?
**Answer:**
Kafka is a distributed streaming platform used for high-throughput, fault-tolerant messaging.
-   **Decoupling:** Services don't call each other directly; they publish events ("OrderPlaced") and others consume them ("EmailService", "InventoryService").
-   **Key Concepts:**
    -   **Topic:** Category of messages.
    -   **Partition:** Order is guaranteed only within a partition.
    -   **Consumer Group:** Allows parallel processing. If you have 10 partitions and 10 consumers in a group, each reads from one partition solely.
    -   **Offset:** The position of a consumer in a partition.
-   **Golang Libs:** `sarama`, `confluent-kafka-go`.

---

## 4. Containerization & DevOps (Generalized)

### Q11: Explain the components of Kubernetes.
**Answer:**
-   **Pod:** Smallest unit, runs one or more containers (e.g., your Go app + sidecar).
-   **Deployment:** Manages Pods (replicas, rolling updates, rollbacks).
-   **Service:** Stable network endpoint (VIP) to access a set of Pods (ClusterIP for internal, LoadBalancer for external).
-   **Ingress:** HTTP/HTTPS router (API Gateway) to services.
-   **ConfigMap/Secret:** Injecting configuration (env vars, config files) and secrets without rebuilding the image.
-   **StatefulSet:** For stateful apps like Databases (ensures stable network IDs and persistent storage).

### Q12: How do you build a small, secure Docker image for a Golang app?
**Answer:**
-   **Multi-Stage Build:**
    1.  **Build Stage:** Use a heavy image (e.g., `golang:1.21-alpine`) to compile the code.
    2.  **Runtime Stage:** Copy only the binary to a minimal image like `gcr.io/distroless/static` or `scratch` (empty image).
    -   *Result:* A tiny image (10-20MB) containing only the binary and root CA certificates, with no OS shell (better security).

### Q13: Basic Observability Setup for a Go Service?
**Answer:**
-   **Metrics (Prometheus):** Expose a `/metrics` endpoint using `prometheus/client_golang` to report request count, latency, memory usage (Go runtime metrics).
-   **Logging:** Use structured logging (JSON) with libraries like `zap` or `logrus` so logs can be parsed by ELK/Splunk.
-   **Tracing:** Use OpenTelemetry (OTel) to pass trace IDs across microservices to visualize the request path (Distributed Tracing) in Jaeger/Zipkin.

---

---

## 5. Advanced Scenarios & Troubleshooting (Lead Level)

### Q14: How do you debug a Memory Leak or High CPU usage in a production Go service?
**Answer:**
-   **pprof:** Go has a built-in profiler tool called `pprof`.
    1.  **Enable pprof:** Import `_ "net/http/pprof"` and start an HTTP server.
    2.  **Generate Profile:** `go tool pprof http://localhost:6060/debug/pprof/heap` (for memory) or `/profile` (for CPU).
    3.  **Analyze:** Use `top` to see functions consuming the most resources or `list <func_name>` to see line-by-line usage.
    4.  **Common Leaks:** Goroutines that never exit (blocked on channel), large global maps that never get cleaned up, or unclosed HTTP response bodies.

### Q15: You need to design a Rate Limiter for your API. How would you approach it?
**Answer:**
-   **Token Bucket Algorithm:**
    -   Visualize a bucket that gets filled with tokens at a constant rate.
    -   Each request removes a token. If the bucket is empty, the request is rejected (429 Too Many Requests).
-   **Implementation:**
    -   **Single Instance:** Go's `golang.org/x/time/rate` package.
    -   **Distributed (Microservices):** Use Redis (with Lua scripts/increments) to maintain a shared counter across all instances of your service.
    -   **Middleware:** Implement this as a middleware in Gin/Gorilla Mux to check the rate limit before hitting the business logic.

### Q16: How do you handle configuration management in a cloud-native Go app?
**Answer:**
-   **12-Factor App Principle:** Store config in the environment.
-   **Libraries:** `Viper` or `kelseyhightower/envconfig`.
-   **Precedence:**
    1.  Default values in code.
    2.  Config file (config.yaml) for non-sensitive data.
    3.  Environment Variables (override config files).
    4.  Command-line flags.
-   **Secrets:** Never commit secrets to git. Inject them via Kubernetes Secrets or HashiCorp Vault into environment variables at runtime.

### Q17: Database Migration Strategies for Zero Downtime?
**Answer:**
When changing a database schema (e.g., renaming a column) without stopping the app:
1.  **Expand:** Add the new column `new_name`.
2.  **Dual Write:** Update code to write to *both* `old_name` and `new_name`.
3.  **Backfill:** Run a script to copy data from `old_name` to `new_name` for existing records.
4.  **Contract:** Switch code to read from `new_name` only.
5.  **Cleanup:** Remove `old_name` column.
-   **Tools:** `golang-migrate/migrate` or `goose` to manage SQL migration files (up/down scripts).

## 6. Leadership & Soft Skills (Lead Role Specifics)

### Q18: How do you mentor junior Go developers?
**Answer:**
-   **Code Reviews:** Focus on *why* something is done, not just syntax. Point out Go idioms (e.g., "accept interfaces, return structs", "handle errors explicitly").
-   **Pair Programming:** Work together on complex tasks to share knowledge on debugging and design.
-   **Design Docs:** Encourage them to write a mini-design doc before coding to think about edge cases.
-   **Safe Environment:** Create a culture where it's okay to make mistakes in non-production environments to learn.

### Q19: How do you decide between building a new Microservice vs. adding to an existing one?
**Answer:**
-   **Single Responsibility Principle:** Does the new functionality fit the domain of the existing service?
-   **Coupling:** Will the new feature require frequent changes to the existing service's data model?
-   **Scalability:** Does this new feature have vastly different scaling requirements (e.g., heavy CPU vs. I/O)?
-   **Team Ownership:** Is there a clear team that owns this new domain?
-   *Rule of Thumb:* Start with a "Modular Monolith" (well-structured code in one service) and split only when the complexity or scaling needs demand it. Premature microservices lead to "Distributed Monoliths" (latency, complexity).

### Q20: Explain "Dependency Injection" in Go and why it matters.
**Answer:**
-   Go doesn't need complex frameworks (like Spring in Java).
-   **Manual DI:** Pass dependencies (Database connections, Logger, Config) as arguments to your service constructors or struct fields.
-   **Interface Decoupling:** Define an interface for the dependency (e.g., `UserRepository`). The service depends on the interface, not the concrete implementation (`PostgresUserRepo`).
-   **Testing:** This allows you to easily inject a `MockUserRepo` during unit tests to test business logic in isolation without a real database.
-   **Wire:** Google's `wire` tool generates DI code at compile-time if the graph gets too complex, but simple manual injection is preferred for clarity.
