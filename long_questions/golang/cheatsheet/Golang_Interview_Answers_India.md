# Golang Interview Answers (India-Focused)

This document provides detailed answers to frequently asked Golang interview questions in India, focusing on "WHY" and providing clear examples.

## 1️⃣ Core Golang (MOST ASKED – All Roles)

### Basics

**1. Why Go? Advantages over Java / Node.js**
*   **Why:** Go was designed by Google to solve problems of scale, fast build times, and maintainability.
*   **Advantages:**
    *   **Simplicity:** Minimal keywords compared to Java's verbose syntax.
    *   **Concurrency:** Built-in Goroutines and Channels make concurrency easy and lightweight compared to Java Threads or Node.js Event Loop (which is single-threaded).
    *   **Performance:** Compiled to machine code (fast like C++) but has Garbage Collection (safe like Java).
    *   **Single Binary:** Compiles into a single static binary, making deployment (especially in containers/Kubernetes) extremely simple compared to JVM or Node_modules.

**2. Is Go compiled or interpreted?**
*   **Answer:** **Compiled.**
*   **Why:** Go compiles directly to machine code for the target architecture (e.g., Linux amd64). There is no VM (like JVM) or Interpreter (like Node.js) needed to run the binary. This is why it's so fast.

**3. What is GOROOT vs GOPATH?**
*   **GOROOT:** The directory where Go itself is installed (contains the compiler, standard library). You rarely need to change this.
*   **GOPATH:** The workspace directory (before Go Modules). It contained `src`, `pkg`, and `bin`.
*   **Status:** Since Go 1.11+ (Go Modules), `GOPATH` is less critical for project structure, but still used to store downloaded modules (in `$GOPATH/pkg/mod`) and installed binaries (in `$GOPATH/bin`).

**4. What is a Go module?**
*   **Answer:** A collection of related Go packages that are versioned together as a single unit. It is defined by a `go.mod` file.
*   **Why:** It solves the "dependency hell" allowing you to specify exact versions of libraries your project needs, independent of where your code lives on disk (no longer forced to be inside `GOPATH/src`).

**5. `go mod tidy` vs `go mod vendor`**
*   **`go mod tidy`:** Adds missing modules to `go.mod` and removes unused ones. It ensures your `go.mod` matches your actual code imports.
*   **`go mod vendor`:** Creates a `vendor` directory containing a copy of all dependencies.
    *   **Why use vendor?** To ensure your build is reproducible even if the internet is down or the original repository is deleted.

**6. Difference between `var`, `:=`, `const`**
*   **`var`:** Declares a variable throughout a function or at package level. Can be uninitialized (zero value).
    *   `var x int` (value is 0)
*   **`:=` (Short Declaration):** Declares and initializes a variable. **Only working inside functions.**
    *   `x := 10`
*   **`const`:** Declares a constant value. Must be calculable at **compile time**. cannot change runtime.

### Data Types & Internals

**7. Value types vs Reference types**
*   **Value Types:** `int`, `float`, `bool`, `string`, `struct`, `array`.
    *   **Behavior:** Assigning or passing them copies the data. Changing the copy does NOT affect the original.
*   **Reference Types:** `slice`, `map`, `channel`, `pointer`, `function`, `interface`.
    *   **Behavior:** They point to an underlying data structure. Assigning them copies the *pointer*, so both variables share the same underlying data.

**8. Array vs Slice**
*   **Array:** Fixed size. `[3]int`. The size is part of the type. `[3]int` != `[4]int`. Values are copied when passed.
*   **Slice:** Dynamic size. `[]int`. It is a lightweight wrapper (header) over an array.
*   **Why Slice?** More flexible and memory efficient for most use cases.

**9. How slice works internally (pointer, len, cap)**
*   A slice is a struct with 3 fields:
    *   **Pointer:** Points to the start of the element in the underlying array.
    *   **Length (len):** Number of elements the slice currently holds.
    *   **Capacity (cap):** Number of elements the underlying array *can* hold before reallocating.
*   **Growth:** When you `append` and `len` exceeds `cap`, Go allocates a new, larger array (usually double size), copies elements, and updates the slice pointer.

**10. Map: Reference or Value type? Thread-safe?**
*   **Type:** Reference type.
*   **Thread-Safety:** **NO.** Maps are NOT safe for concurrent use.
*   **Why:** For performance. Adding locks would slow down all map operations.
*   **Fix:** Use `sync.Mutex` or `sync.Map` for concurrent access.
*   **Crash:** Concurrent read/write to a map causes a fatal runtime panic ("concurrent map writes").

### Functions

**11. Defer – execution order**
*   **Answer:** LIFO (Last-In, First-Out).
*   **Why:** Useful for cleanup (closing files, unlocking mutexes) right after acquiring resources, ensuring they run even if the function returns early or panics.

**12. Panic vs Recover**
*   **Panic:** abruptly stops normal execution (like throwing an exception).
*   **Recover:** Used inside a `defer` function to regain control of a panicking goroutine.
*   **Why:** To prevent the entire program from crashing due to a localized error.

### Structs & Interfaces

**13. Struct embedding (Pseudo-Inheritance)**
*   Go allows embedding a struct inside another. The inner struct's fields/methods are promoted to the outer struct.
*   **Note:** This is composition, not inheritance.

**14. Interface satisfaction (Implicit)**
*   **Answer:** A type implements an interface by simply implementing its methods. No `implements` keyword.
*   **Why:** Decouples definition from implementation (Duck Typing: "If it walks like a duck...").

**15. Difference between Interface{} (Empty Interface)**
*   It has 0 methods, so *every* type satisfies it. Used for code that handles unknown types (like `fmt.Println`).

**🔥 Frequently asked: Why Go does not have inheritance?**
*   **Answer:** Go prefers **Composition over Inheritance**.
*   **Why:** Inheritance leads to brittle type hierarchies (fragile base class problem). Composition allows you to build complex types by combining simple, smaller pieces, which is more flexible and maintainable.

---

## 2️⃣ Concurrency (EXTREMELY IMPORTANT)

### Goroutines

**16. Goroutine vs Thread**
*   **Memory:** Goroutines start small (~2KB stack) vs Threads (1MB+).
*   **Cost:** Creating/destroying goroutines is cheap; threads are expensive (OS calls).
*   **Scheduling:** Goroutines are managed by the Go Runtime (M:N scheduling), Threads by the OS Kernel.

**17. How goroutines are scheduled (GMP model)**
*   **G (Goroutine):** The code to run.
*   **M (Machine):** The OS thread.
*   **P (Processor):** Context for scheduling (local queue of Gs).
*   **How it works:** P assigns Gs to M. If a G blocks (syscall), M releases P, and a new M grabs P to keep running other Gs. Stealing: If a P runs out of Gs, it steals half from another P.

### Channels

**18. Buffered vs Unbuffered channels**
*   **Unbuffered:** `make(chan int)`. Sender blocks until Receiver deals with it (Synchronous).
*   **Buffered:** `make(chan int, 5)`. Sender only blocks if buffer is full. Receiver blocks if buffer is empty.

**19. What happens if you:**
*   **Read from closed channel?** Returns zero value (and `false` ok-flag). Does NOT panic.
*   **Write to closed channel?** **PANIC.**
*   **Close channel twice?** **PANIC.**

### Select

**20. `select` statement**
*   Like a `switch` meant for channels. It blocks until one of the cases (send/receive) is ready.
*   **Fan-in Pattern:** Reading from multiple channels into one.

### Sync Package

**21. `sync.Mutex` vs `sync.RWMutex`**
*   **Mutex:** Locks for both read & write. Only one goroutine access.
*   **RWMutex:** Allows multiple **Readers** (if no writer) OR one **Writer**. Better performance for read-heavy data.

**22. Race condition & detection**
*   **What:** Two goroutines access shared memory, and at least one is a write.
*   **Detect:** Run `go run -race main.go`.

**🔥 Indian Interview Question: How to limit number of goroutines?**
*   **Answer:** Use a **Buffered Channel** as a semaphore (worker pool pattern).
    ```go
    sem := make(chan struct{}, 5) // Limit to 5
    for i := 0; i < 100; i++ {
        sem <- struct{}{} // Block if full
        go func() {
            defer func() { <-sem }() // Release
            doWork()
        }()
    }
    ```

---

## 3️⃣ Golang + MySQL

### Database Basics

**23. Is `sql.DB` a connection or pool?**
*   **Answer:** It is a **Connection Pool**, not a single connection.
*   **Why:** You should create it once (singleton) and reuse it. Do not `Open()` for every query.

**24. `Prepare` vs `Query` vs `Exec`**
*   **Exec:** For INSERT/UPDATE/DELETE (no rows returned).
*   **Query:** For SELECT (returns rows).
*   **Prepare:** Prepares a statement for repeated execution (security + performance).

### Performance

**25. Connection Pooling Logic**
*   `SetMaxOpenConns(n)`: Max active connections.
*   `SetMaxIdleConns(n)`: Max connections kept open in pool waiting for use.
*   `SetConnMaxLifetime(d)`: Close connection after duration `d` (prevents stale connections).

**26. SQL Injection Prevention**
*   **Answer:** Use **Parameterized Queries** (Placeholders `?` or `$1`).
    *   **Bad:** `fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", name)`
    *   **Good:** `db.Query("SELECT * FROM users WHERE name = ?", name)`

**27. Transactions**
*   Use `tx, err := db.Begin()`.
*   Defer `tx.Rollback()` (in case of error/panic).
*   Call `tx.Commit()` only if everything succeeds.

**🔥 Frequently asked: How do you handle database timeout?**
*   **Answer:** Use **Context**.
    ```go
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    db.QueryContext(ctx, "SELECT ...")
    ```
*   **Why:** If DB is slow, the context cancels the query, freeing up connection and goroutine.

---

## 4️⃣ Golang + REST API (Most Common)

### net/http

**28. `http.Handler` vs `http.HandlerFunc`**
*   **`http.Handler`:** An interface with a `ServeHTTP(w, r)` method.
*   **`http.HandlerFunc`:** A helper type that adapts a regular function `func(w, r)` to satisfy the `http.Handler` interface.

**29. Middleware implementation**
*   **Pattern:** Chain of handlers. A function that takes a `http.Handler` and returns a `http.Handler`.
    ```go
    func LoggingMiddleware(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            log.Println(r.URL.Path)
            next.ServeHTTP(w, r)
        })
    }
    ```

**30. Context propagation**
*   **How:** `ctx := r.Context()`. Pass this `ctx` to DB calls or other services.
*   **Why:** To handle cancellation/timeouts across the entire request chain.

### JSON

**31. Struct tags (`json:"name,omitempty"`)**
*   **`json:"name"`**: Rename the field in JSON output.
*   **`omitempty`**: Do not output the field if it has a zero value (0, "", nil, false).
*   **`-`**: Ignore this field (do not marshal/unmarshal).

### Authentication

**32. JWT Implementation**
*   **Create:** `jwt.NewWithClaims` -> `token.SignedString(secretKey)`.
*   **Parse:** `jwt.ParseWithClaims` -> verify signature -> check expiration.
*   **Middleware:** Extract token from `Authorization: Bearer <token>` header -> validate -> store user info in context.

**33. Token Expiration Handling**
*   Client receives 401 Unauthorized -> Client uses Refresh Token to get new Access Token -> Retry request.

### API Design

**34. REST vs RPC**
*   **REST:** Resource-based (NOUNS), Stateless, Standard HTTP verbs (GET, POST).
*   **RPC:** Action-based (VERBS), Function calls over network.

**35. Idempotent APIs**
*   **Definition:** Making the same request multiple times has the same effect as making it once.
    *   **GET, PUT, DELETE:** Idempotent.
    *   **POST:** Not idempotent (creates new resource every time).

---

## 5️⃣ Golang + gRPC (Product Companies Love This)

### gRPC Basics

**36. What is gRPC?**
*   **Answer:** A high-performance RPC framework by Google.
*   **Key features:** Uses **HTTP/2** for transport and **Protocol Buffers** (Protobuf) for serialization.

**37. HTTP/2 Benefits**
*   **Multiplexing:** Multiple requests over a single TCP connection (no Head-of-Line blocking).
*   **Header Compression:** HPACK reduces overhead.
*   **Server Push:** Server can send data before client asks.
*   **Binary Protocol:** efficient parsing.

**38. Protobuf vs JSON**
*   **Protobuf:** Binary format. Smaller size, faster serialization/deserialization. Schema-enforced (`.proto`).
*   **JSON:** Text format. Human readable but larger and slower. Schema-les.

### Protobuf

**39. Field Numbering Importance**
*   Fields are identified by **Tag Number**, not name. Important for backward compatibility. NEVER change the tag number of an existing field.

**40. Backward Compatibility Rules**
*   New fields: Old code ignores them.
*   Deleted fields: Old code sees default values.
*   **Don't** reuse tag numbers.

### gRPC Types

**41. Unary RPC:** Simple Request -> Response (1:1).
**42. Server Streaming:** Request -> Stream of Responses (1:N) (e.g., Live stock updates).
**43. Client Streaming:** Stream of Requests -> 1 Response (N:1) (e.g., Uploading big file).
**44. Bidirectional Streaming:** Stream -> Stream (N:N) (e.g., Chat app).

**🔥 Very common: Why gRPC is faster than REST?**
*   **Answer:**
    1.  **Protocol:** HTTP/2 (Multiplexing) vs HTTP/1.1.
    2.  **Payload:** Protobuf (Binary, Compressed) is much smaller than JSON.
    3.  **Parsing:** Binary parsing is CPU efficient compared to text parsing (JSON).

---

## 6️⃣ Golang + Azure (Cloud Roles)

### Azure Basics

**45. Azure App Service vs AKS**
*   **App Service:** PaaS. Easy deploy. Good for simple web apps. Scalable but less control.
*   **AKS (Kubernetes Service):** CaaS/Orchestrator. Full control over containers, networking, scaling. Good for microservices.
*   **Decision:** Simple App -> App Service. Complex Microservices -> AKS.

### Go + Azure

**46. Deploy Go App**
*   **App Service:** Build binary -> Zip deploy OR push Docker image to Azure Container Registry (ACR) -> App Service pulls it.
*   **AKS:** Build Docker image -> Push to ACR -> Apply Kubernetes manifests (Deployment, Service).

### Storage

**47. Azure Blob Storage**
*   Use `azblob` SDK.
*   Connect using Connection String or Managed Identity (Better).

### Messaging

**48. Service Bus vs Event Hub**
*   **Service Bus:** Enterprise Messaging. Queues & Topics. Exactly-once delivery guarantee. (Order Processing).
*   **Event Hub:** Big Data Streaming. Ingests millions of events/sec. (Telemetry, Logging).

### Security

**🔥 Common question: How do you securely store DB credentials?**
*   **Answer:** Use **Azure Key Vault**.
*   **How:**
    1.  Enable **Managed Identity** for the Go App.
    2.  Grant the App's identity access to Key Vault.
    3.  Go App uses Azure SDK to fetch secret at runtime (no secrets in code/env vars).

---

## 7️⃣ Golang + Microservices (Senior Roles)

### Architecture

**49. Monolith vs Microservices**
*   **Monolith:** Single codebase, single deployments, shared database. Hard to scale specific parts.
*   **Microservices:** Use **independent deployment** of small services, each with its own DB. Communicates via HTTP/gRPC.

**50. Service Discovery**
*   **What:** How Service A finds Service B (IP + Port change dynamically).
*   **Tools:** Consul, Etcd, or **Kubernetes DNS** (CoreDNS). Services refer to each other by name (e.g., `http://payment-service`), and DNS resolves it.

**51. API Gateway**
*   **Role:** Single entry point for clients (Web/Mobile).
*   **Tasks:** Auth, Rate Limiting, Routing, SSL Termination. (e.g., Kong, Nginx, Azure API Gateway).

### Observability

**52. Logging, Metrics, Tracing**
*   **Logging:** Events (Zap/Logrus). "What happened?"
*   **Metrics:** Aggregates (Prometheus + Grafana). "How many requests/sec?"
*   **Tracing:** Request path across services (OpenTelemetry, Jaeger). "Where is the latency?"

### Resilience

**53. Circuit Breaker**
*   **Concept:** Prevents cascading failures. If Service B is failing, Service A stops calling it immediately (Open State) and returns error/fallback, instead of waiting for timeouts.

**54. Refund with Backoff**
*   **Concept:** If a call fails, retry. But wait longer each time (1s, 2s, 4s...) to avoid hammering a recovering service.

---

## 8️⃣ Golang Testing (Often Ignored, Still Asked)

**55. Table-Driven Tests**
*   **Why:** Go idiom to avoid repetitive test code.
    ```go
    tests := []struct {
        input    int
        expected int
    }{
        {1, 2},
        {2, 4},
    }
    for _, tc := range tests {
        // Run test
    }
    ```

**56. Mocking**
*   **How:** Use **Interfaces**.
    *   Inject an interface (e.g., `Database`) into your service.
    *   In tests, pass a `MockDatabase` struct that satisfies the interface but returns fake data.
*   **Tools:** `github.com/stretchr/testify/mock` or `gomock`.

**57. Benchmarking**
*   **Command:** `go test -bench .`
*   **Function:** `func BenchmarkMyFunc(b *testing.B)`
*   **Usage:** To compare performance of two implementations.

---

## 9️⃣ Golang Performance & Internals (Senior Level)

**58. Garbage Collector (GC)**
*   **Algorithm:** **Non-Generational Concurrent Tri-color Mark and Sweep.**
*   **Goal:** Low latency (stops the world for very short time).
*   **Tuning:** `GOGC` variable (default 100).

**59. Escape Analysis**
*   **Concept:** Compiler decision: "Should this variable live on **Stack** or **Heap**?"
*   **Stack:** Fast allocation/deallocation (when function returns).
*   **Heap:** Slower, needs GC. Happens if a reference is returned from a function or passed to a huge struct.
*   **Check:** `go build -gcflags="-m"`

**60. Memory Leaks in Go?**
*   **Possible?** Yes.
*   **Causes:**
    1.  **Goroutine Leaks:** Starting a goroutine that never exits (blocked on channel/nil map).
    2.  **Unclosed Resources:** `defer body.Close()` (HTTP response bodies).
    3.  **Substring:** Keep a small slice of a huge array (keeps the huge array in memory).

**61. CPU Profiling**
*   Use `net/http/pprof`.
*   Capture profile and view with `go tool pprof`.

---

## 🔟 HR + Scenario-Based (Indian Interviews LOVE These)

**62. How did you debug a production issue?**
*   **Scenario:** High latency.
*   ** Steps:**
    1.  Checked **Metrics** (Grafana) -> Saw CPU spike.
    2.  Checked **Logs** (ELK/Splunk) -> Found specific endpoint errors.
    3.  Used **Tracing** (Jaeger) -> Found DB query taking 5s.
    4.  **Fix:** Added missing index to DB.

**63. How to handle high traffic?**
*   **Horizontal Scaling:** Add more pods (HPA in Kubernetes).
*   **Cache:** Use Redis for read-heavy data.
*   **Async:** Move heavy tasks (emails, report generation) to a **Queue** (Kafka/RabbitMQ) and process in background.
*   **DB:** Read Replicas.

**64. Explain a Race Condition you faced?**
*   **Scenario:** Two goroutines updating a global counter (or map) without a lock.
*   **Fix:** Added `sync.Mutex` around the critical section.

**65. Handling partial failures (Microservices)?**
*   **Scenario:** Payment service is down, but User service is up.
*   **Solution:** **Graceful Degradation**.
    *   Allow user to login/browse, but disable "Checkout" button.
    *   Show "Service unavailable" for that specific part, don't crash the whole page.

---
**Good Luck! 🚀**