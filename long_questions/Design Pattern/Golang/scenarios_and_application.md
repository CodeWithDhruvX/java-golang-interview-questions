# 🧩 Applied Design Patterns: Scenarios & Problem Statements

In FAANG interviews, you are rarely asked "What is the Adapter Pattern?". Instead, you are given a vague problem and expected to **apply** the correct pattern to solve it.

Here is a mapping of common interview scenarios to their solutions.

---

## ⚡ Concurrency Scenarios

### 1. The "Web Crawler" / "File Downloader"
**Problem:** You need to download 10,000 URLs. If you spawn 10,000 goroutines, the system crashes (OOM) or you get banned by sites.
*   ✅ **Solution:** **Worker Pool Pattern**.
*   **Why?** Limits concurrency to a fixed number (e.g., 50 workers) while processing a stream of jobs.

### 2. The "Dashboard Aggregator"
**Problem:** A user loads their dashboard. You need to fetch: 1. User Profile, 2. Recent Orders, 3. Notifications. These are 3 different slow database calls. You want the total time to be `max(t1, t2, t3)`, not `sum(t1, t2, t3)`.
*   ✅ **Solution:** **Fan-Out / Fan-In**.
*   **Why?** Fan-Out to start 3 goroutines. Fan-In to wait for all 3 to finish and combine the results into one JSON response.

### 3. The "Rate Limited API"
**Problem:** You are building a service that scrapes Twitter, but Twitter blocks you if you send more than 10 requests/second.
*   ✅ **Solution:** **Token Bucket (Rate Limiting)** or **Semaphore**.
*   **Why?** Semaphore strictly limits active connections. Token Bucket ensures you respect the "per second" rate.

### 4. The "Data Pipeline"
**Problem:** You have a 100GB CSV file. You need to: 1. Read line, 2. Parse JSON, 3. Sanitize data, 4. Insert into DB. Doing this sequentially for 10M rows is too slow.
*   ✅ **Solution:** **Pipeline Pattern**.
*   **Why?** Create 4 stages connected by channels. Stage 1 reads, Stage 4 writes. All stages run in parallel on different chunks of data.

---

## ☁️ Microservices & Distributed Scenarios

### 5. The "Resilient Payment System"
**Problem:** Your "Checkout" service calls "Payment Service". If "Payment Service" is down or slow, your Checkout service hangs and eventually crashes, taking down the whole site.
*   ✅ **Solution:** **Circuit Breaker Pattern**.
*   **Why?** Detects failures fast. If Payment is down, fail immediately (or queue for later) instead of hanging threads waiting for timeouts.

### 6. The "Distributed Transaction"
**Problem:** User buys a ticket. 1. Deduct money (Bank Svc), 2. Reserve Seat (Ticket Svc). If Reservation fails, you must Refund the money. You can't use SQL transactions across 2 microservices.
*   ✅ **Solution:** **SAGA Pattern**.
*   **Why?** Use a sequence of local transactions with compensating actions (undos) to ensure eventual consistency.

### 7. The "Legacy Integration"
**Problem:** You are rewriting a monolithic Node.js app into Golang Microservices. You need to route `/api/v1/users` to the old Node app and `/api/v2/users` to the new Go app, seamlessly.
*   ✅ **Solution:** **API Gateway (Strangler Fig Pattern)**.
*   **Why?** The Gateway sits in front and routes traffic. You can slowly "strangle" the legacy app by ensuring new routes go to the new service.

### 8. The "Noisy Neighbor"
**Problem:** Your system processes both "High Priority User Logins" and "Low Priority Data Reports". Sometimes the Reports generation eats 100% CPU/DB Connections, and users can't login.
*   ✅ **Solution:** **Bulkhead Pattern**.
*   **Why?** Create separate resource pools (connection pools, worker pools) for Logins vs Reports. If Reports sink, Logins stay afloat.

---

## 🏗️ Classic OOP Scenarios (adapted for Go)

### 9. The "Database Driver"
**Problem:** Your app supports PostgreSQL today, but might support MySQL or SQLite tomorrow. You don't want to rewrite the whole app when you switch DBs.
*   ✅ **Solution:** **Abstract Factory** or **Strategy Pattern**.
*   **Why?** Define an interface `DB { Query() }`. Implement `PostgresDB`, `MySQLDB`. Inject the interface into your service.

### 10. The "Complex Config Object"
**Problem:** Creating a `Server` struct requires 20 parameters (host, port, timeout, tls keys, logger, etc.). Most are optional. `NewServer(a, b, c, d, ...)` is ugly.
*   ✅ **Solution:** **Builder Pattern** (or **Functional Options Pattern** in Go).
*   **Why?** Allows `NewServer(WithPort(8080), WithTimeout(5s))`. Clean and extensible.

### 11. The "Middleware Chain"
**Problem:** For every HTTP request, you need to: 1. Log it, 2. Authenticate it, 3. Gzip compress it.
*   ✅ **Solution:** **Decorator Pattern** (or **Chain of Responsibility**).
*   **Why?** Wrap the `http.Handler` with `LoggingMiddleware(AuthMiddleware(GzipMiddleware(handler)))`.

### 12. The "Undo Button"
**Problem:** You are building a text editor. User types text, then hits Ctrl+Z to undo.
*   ✅ **Solution:** **Memento Pattern** (State snapshot) or **Command Pattern** (Reversible actions).
*   **Why?** Command stores the "diff" of what happened so it can be reversed. Memento stores the snapshot of the previous state.
