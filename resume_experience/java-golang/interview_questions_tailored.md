# Tailored Interview Questions for Senior Backend Engineer
Based on your resume (Golang, Java, Angular, Azure, Kubernetes), here are the most likely and impactful interview questions you might face.

## 1. Golang (Core & Advanced)
**Context:** You have strong recent experience with Gin, Gorilla/Mux, and microservices.

*   **Concurrency & Goroutines:**
    *   **Q:** How does the Go scheduler handle Goroutines compared to OS threads? Explain the M:N scheduler model.
    *   **Q:** What is a race condition? How do you detect it in Go ( `-race` flag)?
    *   **Q:** Explain the difference between unbuffered and buffered channels. When would you use one over the other?
    *   **Q:** How do you handle graceful shutdowns in a Go microservice (handling `SIGTERM`/`SIGINT` context cancellation)?

*   **Memory Management:**
    *   **Q:** How does Garbage Collection work in Go (Tricolor Mark-and-Sweep)? Have you ever had to tune `GOGC`?
    *   **Q:** Explain the difference between `new()` and `make()`.
    *   **Q:** When does a variable escape to the heap? (Escape Analysis).

*   **Frameworks (Gin/Gorilla):**
    *   **Q:** You used Gin and Gorilla/Mux. What are the key performance differences? Why might you choose a standard `net/http` router over Gin for a simple service?
    *   **Q:** How do you implement middleware in Gin (e.g., for logging, auth, panic recovery)?

## 2. Java & Spring Boot
**Context:** listed as a core skill, though less recent than Go in the last year.

*   **Spring Boot Internals:**
    *   **Q:** How does Spring Boot auto-configuration work? (@EnableAutoConfiguration, `spring.factories`).
    *   **Q:** Explain the Bean lifecycle in Spring. What is the difference between `@Component`, `@Service`, and `@Repository`?
    *   **Q:** How do you handle transaction management in Spring (`@Transactional` propagation levels)?

*   **Java Core:**
    *   **Q:** Explain the changes in memory management from Java 8 to Java 17+ (G1GC, ZGC).
    *   **Q:** How does a `ConcurrentHashMap` achieve thread safety compared to `Hashtable` or `Collections.synchronizedMap`?

## 3. Microservices & System Design
**Context:** Experience with "Distributed systems", "Insight Governance", "KPMG Audit Platform".

*   **Architecture:**
    *   **Q:** You migrated/built services for an Audit Platform. How did you handle data consistency across microservices? (Saga Pattern vs 2PC).
    *   **Q:** How do you design for failure? (Circuit Breaker, Retries with Exponential Backoff).
    *   **Q:** Compare gRPC vs REST. Why did you choose gRPC for your internal food delivery project but REST for others?

*   **Scalability:**
    *   **Q:** How would you scale the "Insight Governance" platform if the data volume increased by 100x? (Database partitioning, caching strategies with Redis).
    *   **Q:** How do you handle distributed tracing? You mentioned Jaeger—how do you propagate context across service boundaries in Go?

## 4. Kubernetes, Docker & Cloud (Azure)
**Context:** Strong Azure AKS experience.

*   **Kubernetes:**
    *   **Q:** Explain the difference between a `Deployment`, `StatefulSet`, and `DaemonSet`.
    *   **Q:** How do you handle configuration management in K8s? (ConfigMaps vs Secrets).
    *   **Q:** What is a Sidecar pattern? Have you used it (e.g., for logging or service mesh)?
    *   **Q:** How do you maximize resource utilization in AKS? (Requests vs Limits, HPA - Horizontal Pod Autoscaler).

*   **Docker:**
    *   **Q:** How do you optimize Docker build images for Golang? (Multi-stage builds to create standard slim images).

## 5. Database (SQL & NoSQL)
**Context:** PostgreSQL, MySQL, MongoDB, Redis.

*   **Q:** You used both SQL (Postgres) and NoSQL (Mongo). What factors decide which one to choose for a new service?
*   **Q:** Explain ACID properties. How does Redis persistence work (RDB vs AOF)?
*   **Q:** How do you optimize a slow query in PostgreSQL? (Explain analyze, Index types like B-Tree vs GIN).

## 6. Frontend (Angular)
**Context:** "Built and maintained Angular-based admin console".

*   **Q:** Explain the change detection mechanism in Angular (Zone.js).
*   **Q:** What are RxJS Observables vs Promises? When to use `switchMap` vs `mergeMap`?
*   **Q:** How do you optimize Angular performance? (Lazy loading modules, OnPush change detection).

## 7. Project-Specific Questions (The "Hook")
**Context:** These questions target the specific bullets in your resume.

*   **On "Insight Governance":**
    *   **Q:** You "Integrated PowerBI dashboards into the UI". How did you handle the authentication/embedding securely?
    *   **Q:** You "improved operational insights by 15%". How did you measure this? What metrics did you track?

*   **On "KPMG Audit Platform":**
    *   **Q:** You "built reusable Golang microservice frameworks". What exactly was reusable? (e.g., logging wrappers, error handling, DB connectors).
    *   **Q:** "Rule-based calculations"—did you implement a rules engine? If so, was it hardcoded or dynamic? (e.g., did you use an AST or a library?)
