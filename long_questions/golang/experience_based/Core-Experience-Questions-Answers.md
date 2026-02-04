# Core Golang Experience-Based Interview Questions

This document transforms "textbook" core Golang questions into **experience-based** inquiries suitable for Senior/Staff roles at product-based companies.

## 1️⃣ Concurrency in Production

**1. Describe a time you encountered a goroutine leak. How did you detect and fix it?**
> **Answer:**
> "I successfully deployed a background worker service, but after 3 days, it crashed with OOM (Out of Memory).
> *   **Detection:** I checked the metrics (Prometheus) and saw `go_goroutines` climbing linearly from 100 to 50,000 over days.
> *   **Investigation:** I took a pprof goroutine profile: `go tool pprof http://.../debug/pprof/goroutine`. It pointed to a function `processEvents`.
> *   **Root Cause:** We were launching a goroutine to send metrics to a slow downstream service without a timeout. The channel send `metricsCh <- data` was blocking forever when the consumer fell behind.
> *   **Fix:** I wrapped the send in a `select` with a `default` case (to drop metrics if full) or added a `context.WithTimeout`."

**2. Have you ever faced a deadlock in a production system?**
> **Answer:**
> "Yes, in a caching layer. We had a `RWMutex` protecting a map.
> *   **Scenario:** A `Get()` function acquired a `RLock()` (read lock), inspecting the cache. Inside `Get()`, if the item was missing, it called `FetchAndSet()`.
> *   **The Bug:** `FetchAndSet()` tried to acquire a `Lock()` (write lock). But the write lock couldn't be granted until the read lock was released. The read lock wouldn't release until `FetchAndSet` successfully finished.
> *   **Result:** The entire application hung.
> *   **Fix:** I refactored the code to release the `RLock` before calling the function that needed the `Lock`, or utilized `double-checked locking` carefully."

**3. When have you chosen Buffered over Unbuffered channels (or vice-versa)?**
> **Answer:**
> "I strictly use **Unbuffered Channels** for synchronization (orchestration) effectively acting as a 'handshake'—I want to know *for sure* the receiver got the message.
> *   **Example:** A graceful shutdown signal.
> I use **Buffered Channels** for decoupling/throughput.
> *   **Example:** A logging worker. I gave it a buffer of 1000. If the logs spike, the application doesn't block immediately.
> *   **Trade-off:** If the app crashes, I lose the 1000 logs in the buffer. I accepted that risk for the performance gain (latency reduction) on the main thread."

**4. How do you manage the lifecycle of thousands of goroutines (Worker Pool)?**
> **Answer:**
> "We needed to process 10,000 S3 files concurrently. Spawning 10,000 goroutines caused the API rate limiter to ban us immediately.
> *   **Solution:** I implemented a 'Bounded Worker Pool'. I created a channel of size `N` (e.g., 50).
> *   **Implementation:** I started 50 workers consuming from a job channel.
> *   **Key Detail:** I used `sync.WaitGroup` to ensure the main function didn't exit until all 50 workers finished draining the channel.
> *   **Outcome:** We processed files at max predictable throughput without overwhelming memory or external quotas."

---

## 2️⃣ Memory Management & Optimization

**1. Have you ever used `sync.Pool`? Why specifically?**
> **Answer:**
> "Yes, in a high-throughput JSON API service.
> *   **Problem:** The GC was using 30% of our CPU. `go tool pprof` showed massive allocations in `json.Unmarshal`. We were creating a new `bytes.Buffer` and struct for every HTTP request (50k RPS).
> *   **Solution:** I introduced `sync.Pool` to reuse the buffers.
> *   **Gotcha:** I had to be extremely careful to `Reset()` the buffer before putting it back in the pool to avoid data leakage between requests.
> *   **Result:** GC CPU time dropped from 30% to 5%, and P99 latency improved significantly."

**2. How do you choose between passing by Value vs. Pointer?**
> **Answer:**
> "The textbook says 'pointers for large structs, values for small'. But in my experience, **Escape Analysis** matters more.
> *   **Scenario:** I was passing a pointer to a small struct `*Config` thinking it was efficient. However, because I returned that pointer from the function, the compiler forced it to escape to the Heap.
> *   **Optimization:** I switched to passing by Value. This kept the struct on the Stack. Stack allocation is essentially free (just moving the stack pointer).
> *   **Rule of Thumb:** I default to Value. I only use Pointers if I *need* to share state (mutability) or if the struct is truly huge (>2KB)."

**3. Describe deep diving into a memory leak using pprof.**
> **Answer:**
> "I used `go tool pprof -accumulate_info -alloc_space http://localhost:8080/debug/pprof/heap`.
> *   **Observation:** I noticed `inuse_space` was stable, but `alloc_space` was huge. This meant we were churning memory (high GC pressure) rather than leaking it.
> *   **Drill Down:** The `top` command showed `strings.Split` was the culprit. We were parsing a massive log file in memory.
> *   **Fix:** I switched to `bufio.Scanner` to stream the file line-by-line instead of loading the whole string and splitting it."

---

## 3️⃣ Interfaces, Generics & Design

**1. Go 1.18+ Generics: When did you decide to USE them vs. AVOID them?**
> **Answer:**
> "I use Generics for **data structures** (like a `Set[T]` or `ConcurrentMap[K, V]`) and **utility functions** (like `Map/Filter/Reduce` on slices).
> *   **Where I Avoided it:** I didn't use it for our business logic 'Service' interfaces.
> *   **Why:** A `Service[T]` made mocking harder and the code harder to read for new juniors. Go interfaces (`Accept Interfaces, Return Structs`) are usually sufficient for polymorphism. Generics are for *type safety* without reflection, not for OOP-style inheritance."

**2. Explain the 'Typed Nil' problem in a production context.**
> **Answer:**
> "I caused a panic with this once.
> *   **Code:** `func GetError() error { var myErr *MyCustomError = nil; return myErr }`
> *   **The Bug:** The caller checked `if err != nil`. Since `myErr` (a pointer) was wrapped in an interface `error`, the interface itself was *not nil* (it had a type `*MyCustomError` but a value `nil`).
> *   **The Crash:** The code proceeded to access methods on the nil error.
> *   **Fix:** I learned to always return an explicit `nil` literal (`return nil`), not a nil pointer of a concrete type, when satisfying an interface."

---

## 4️⃣ Error Handling at Scale

**1. How do you handle errors in a mono-repo/large microservice architecture?**
> **Answer:**
> "We moved away from just `return err`.
> *   **Wrapping:** We use `fmt.Errorf("fetching user failed: %w", err)` to add context (stack-trace-like history) while preserving the original error for inspection.
> *   **Sentinel Errors:** For domain logic, we define `var ErrUserNotFound = errors.New(...)`.
> *   **The Check:** We use `errors.Is(err, ErrUserNotFound)` so our HTTP handler knows to return 404.
> *   **Experience:** We avoided using custom struct errors for everything because it created tight coupling between packages. Sentinel errors + Wrapping gave us the best balance."

**2. How do you handle panicked goroutines in production?**
> **Answer:**
> "If a goroutine panics, it crashes the *entire program*, not just that goroutine.
> *   **Strategy:** For critical background jobs and HTTP middleware, I use a `defer/recover` block.
> *   **Implementation:** `defer func() { if r := recover(); r != nil { log.Error("Panic recovered", r) } }()`
> *   **Caveat:** I don't use this everywhere. If my specific business logic panics (e.g., nil pointer), I generally *want* the app to crash and restart (Crash-only software) rather than run in a corrupted state, unless it's the main request handler."

---

## 5️⃣ Real-World Networking & Context

**1. How do you propagate request IDs across microservices?**
> **Answer:**
> "We use `context.Context`.
> *   **Middleware:** An ingress middleware generates a `X-Request-ID`. It puts it into the Go Context using a private key type `ctx = context.WithValue(ctx, key, reqID)`.
> *   **Outbound:** When calling Service B, we have an HTTP client interceptor that extracts the ID from the Context and injects it into the HTTP headers.
> *   **Value:** This let us search Splunk/Datadog for one ID and see the trace across 5 different services."

**2. What happens if you ignore `ctx.Done()` in a long-running DB query?**
> **Answer:**
> "We had an incident where the API client timed out (disconnected), but our backend kept processing the massive report generation for another 30 seconds.
> *   **Impact:** We wasted CPU/DB resources on a result no one was listening for.
> *   **Fix:** In our tight loops, I added `select { case <-ctx.Done(): return ctx.Err(); default: work() }`. For DB queries, I switched to `db.QueryContext(ctx, ...)` so the driver kills the connection immediately when the client vanishes."

---
