# Persistent Systems - Golang Developer Interview Preparation

This document is tailored specifically for the **Golang Developer** role at **Persistent Systems**, based on the 2026 hiring trends and job description. It covers key areas: Core Go, Microservices, Cloud Native, and specialized tools.

---

## **1. Core Golang & Concurrency**
*Role Requirement: "Deep understanding of concurrency model (Goroutines, channels) and idiomatic patterns."*

### **Q1: Explain the GMP Model (Goroutine Scheduler).**
*   **Concept:** Go uses an M:N scheduler where **M** OS threads run **N** Goroutines on **P** logical processors.
*   **Key Components:**
    *   **G (Goroutine):** Lightweight thread, contains stack and instruction pointer.
    *   **M (Machine):** OS thread.
    *   **P (Processor):** Resource required to execute Go code (holds the run queue).
*   **Why it matters:** It explains why Go is efficient at handling thousands of concurrent tasks with minimal overhead compared to OS threads.

### **Q2: Buffered vs. Unbuffered Channels?**
*   **Unbuffered:** `make(chan int)`. Blocks the sender until a receiver is ready (Synchronous).
*   **Buffered:** `make(chan int, 5)`. Sender blocks only when the buffer is full. Receiver blocks only when empty.
*   **Use Case:** Use unbuffered for strict synchronization. Use buffered to decouple specific producer-consumer rates or to prevent blocking on bursts.

### **Q3: How do you handle Data Races?**
*   **Detection:** Run tests with `go test -race`.
*   **Prevention:**
    *   Use **Channels** to share memory by communicating.
    *   Use `sync.Mutex` or `sync.RWMutex` to lock shared resources.
    *   Use `sync/atomic` for simple counters/flags.

### **Q4: What is `context` used for?**
*   **Purpose:** Propagate deadlines, cancellation signals, and request-scoped values across API boundaries and between goroutines.
*   **Common Use:**
    *   `ctx, cancel := context.WithTimeout(parent, 5*time.Second)`
    *   Passing `ctx` to DB queries or HTTP requests to ensure they abort if the client disconnects or times out.

---

## **2. Microservices & Architecture**
*Role Requirement: "Build and optimize microservice-based systems."*

### **Q5: gRPC vs. REST - When to use which?**
*(Critical Topic - likely a gap in your repo)*
*   **REST (JSON/HTTP):**
    *   *Pros:* Human-readable, widespread browser support, easy to debug.
    *   *Cons:* Text-based (slower parsing), no strict schema.
*   **gRPC (Protobuf/HTTP2):**
    *   *Pros:* Binary format (smaller/faster), multiplexing (HTTP/2), code generation for strict contracts, supports streaming.
    *   *Cons:* Not browser-native (needs proxy).
*   **Persistent Context:** Expect gRPC for internal service-to-service communication in high-performance backends.

### **Q6: Explain Microservices Design Patterns.**
*   **API Gateway:** Single entry point for clients, handles routing, auth, and rate limiting.
*   **Circuit Breaker:** Prevents cascading failures. If Service A -> Service B fails repeatedly, "trip" the breaker to return errors immediately instead of waiting for timeouts.
*   **Saga Pattern:** Managing distributed transactions (e.g., Order -> Payment -> Inventory) using a sequence of local transactions with compensating actions for rollback.

---

## **3. Messaging (Kafka)**
*Role Requirement: "Integrate event-driven messaging systems like Kafka."*

### **Q7: Kafka Components & Ordering.**
*   **Topic:** Categorized stream of records.
*   **Partition:** Ordered, immutable sequence of records. **Parallelism unit**.
*   **Group:** Consumer Group.
*   **Ordering Guarantee:** Kafka guarantees order **only within a partition**, not across the entire topic.

### **Q8: How to ensure At-Least-Once delivery?**
*   **Producer:** Set `acks=all` (wait for all replicas).
*   **Consumer:** Commit offsets **after** successfully processing the message.
*   **Risk:** Duplicates possible (if consumer crashes after processing but before commit). Application must be **idempotent**.

---

## **4. Cloud Native (Docker & Kubernetes)**
*Role Requirement: "Deploy using Docker, Kubernetes, Helm."*

### **Q9: Docker vs. Virtual Machines.**
*   **VM:** Hypervisor virtualizes hardware. Runs full OS (heavy).
*   **Docker:** Container engine virtualizes the OS Kernel. Shares host kernel, isolated userspace (lightweight, fast boot).

### **Q10: Kubernetes Components.**
*   **Control Plane:** API Server, Etcd (store), Scheduler, Controller Manager.
*   **Node:** Kubelet (agent), Kube-proxy (network rules), Container Runtime.
*   **Pod:** Smallest deployable unit (one or more containers).

### **Q11: What is Helm?**
*   **Definition:** The package manager for Kubernetes.
*   **Concept:** Uses "Charts" (templates) to define, install, and upgrade complex K8s applications (e.g., a chart deploying a Go app + Postgres + Redis).

---

## **5. Testing & Observability (Gaps Filled)**
*Role Requirement: "Ginkgo/Gomega, Prometheus, Grafana."*

### **Q12: Ginkgo & Gomega (BDD Testing)**
*   **Ginkgo:** A BDD (Behavior Driven Development) testing framework.
    *   Uses `Describe`, `Context`, `It` blocks to structure tests like sentences.
*   **Gomega:** Matcher library used with Ginkgo.
    *   `Expect(result).To(Equal(expected))`
*   **Example:**
    ```go
    var _ = Describe("ShoppingCart", func() {
        Context("When adding an item", func() {
            It("Should increase the total count", func() {
                cart.Add("Apple")
                Expect(cart.Count()).To(Equal(1))
            })
        })
    })
    ```

### **Q13: Observability with Prometheus.**
*   **Metric Types:**
    *   **Counter:** Value matches strictly up (e.g., `http_requests_total`).
    *   **Gauge:** Goes up and down (e.g., `memory_usage_bytes`, `goroutines_active`).
    *   **Histogram:** Samples observations (e.g., request duration) and counts them in buckets.
*   **Pull Model:** Prometheus "scrapes" metrics from your Go app's `/metrics` endpoint (using `promhttp` handler) at intervals.

---

## **6. Recommended Practice Tasks**
1.  **Create a gRPC Service:** Define a `.proto` file for a "User Service", generate Go code, and implement a `CreateUser` RPC.
2.  **Add Prometheus Metrics:** Add a counter to the middleware of your current Gorilla Mux project to count HTTP requests.
3.  **Write a Ginkgo Test:** Write a simple BDD test for a calculator function.
