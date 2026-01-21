## 🧵 Concurrency Architecture & Design Patterns (Questions 961-980)

### Question 961: How do you architect a pub/sub system in Go?

**Answer:**
(See Q606). Channels + Map.

---

### Question 962: How do you build a pipeline using goroutines?

**Answer:**
Stages: `Gen() -> Sq() -> Print()`.
Each function takes `<-chan In` and returns `<-chan Out`.
Connect them in main: `out = Sq(Gen())`.

---

### Question 963: What is the fan-in/fan-out pattern in Go?

**Answer:**
(See Q601/Q602).

---

### Question 964: How do you limit concurrency using semaphores?

**Answer:**
(See Q604). Buffered channel.

---

### Question 965: How do you implement a worker pool?

**Answer:**
(See Q609).
Crucial for capping resource usage (DB connections, memory).

---

### Question 966: How do you handle retries with backoff in goroutines?

**Answer:**
(See Q617).
Combine `time.Sleep` with `select { case <-ctx.Done(): return }` to ensure retry loop stops if request is cancelled.

---

### Question 967: What is the circuit breaker pattern in Go?

**Answer:**
(See Q619). `sony/gobreaker`.

---

### Question 968: How do you implement message deduplication?

**Answer:**
(See Q625). SHA256 of message content -> Check in Redis.

---

### Question 969: How do you synchronize shared state across goroutines?

**Answer:**
1.  **Start with Channels** (Share memory by communicating).
2.  If complex/performance critical: **Mutex**.
3.  If simple counter: **Atomic**.

---

### Question 970: How do you detect livelocks in Go?

**Answer:**
Harder than Deadlocks.
Goroutines are running (burning CPU) but making no progress (e.g., constantly failing a CAS operation or retrying immediately).
Monitoring CPU usage + No Application Throughput is a sign.

---

### Question 971: How do you timeout long-running operations?

**Answer:**
`select`.
```go
select {
case res := <-work():
    return res
case <-time.After(5 * time.Second):
    return Timeout
}
```

---

### Question 972: How do you use the actor model in Go?

**Answer:**
Go doesn't have native Actors (like Erlang/Akka).
Simulate it:
Struct with a `Run()` loop and an `inbox chan Msg`.
Methods on struct just send to `inbox`.

---

### Question 973: How do you architect loosely coupled goroutines?

**Answer:**
Do not share state variables.
Pass data via Channels.
Use `Context` for cancellation signal propagation.

---

### Question 974: How do you design state machines in Go?

**Answer:**
`type State func() State`
Loop: `currentState = currentState()`.
Each state function returns the next state function. (Rob Pike's Lexer pattern).

---

### Question 975: How do you throttle a job queue in Go?

**Answer:**
(See Q608). Token Bucket.

---

### Question 976: How do you monitor goroutine health?

**Answer:**
Heartbeat channel.
Worker sends `heartbeat <- struct{}{}` every second.
Supervisor kills/restarts worker if no beat for 5s.

---

### Question 977: How do you track context propagation in goroutines?

**Answer:**
Always pass `ctx` as first argument.
Use `pprof` labels `pprof.Do(ctx, labels, func(ctx) { ... })` to attach metadata to goroutines in profiles.

---

### Question 978: How do you implement saga pattern in Go services?

**Answer:**
Distributed Transactions.
Sequence of local transactions.
If Step 3 fails, execute Compensating Transactions (Undo) for Step 2 and Step 1.
Orchestrator based (Central Go code) or Choreography (Events).

---

### Question 979: How do you chain async jobs with error handling?

**Answer:**
Use `errgroup`.
It runs N jobs. If any returns error, it cancels the context for all others and returns the first error.

---

### Question 980: How do you log and trace concurrent tasks?

**Answer:**
Include `GoroutineID` (hacky) or better, a `RequestID` in the closure/logger context.

---
