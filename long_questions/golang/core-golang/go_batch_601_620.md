## 🔸 Advanced Concurrency Patterns (Questions 601-620)

### Question 601: How do you implement a fan-in pattern in Go?

**Answer:**
Fan-in merges multiple input channels into a single output channel.

```go
func FanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() {
        defer close(c)
        var wg sync.WaitGroup
        wg.Add(2)
        output := func(ch <-chan string) {
            defer wg.Done()
            for v := range ch { c <- v }
        }
        go output(input1)
        go output(input2)
        wg.Wait()
    }()
    return c
}
```

---

### Question 602: How do you implement a fan-out pattern in Go?

**Answer:**
Fan-out distributes work from one channel to multiple worker goroutines.

```go
func FanOut(ch <-chan int, workers int) {
    var wg sync.WaitGroup
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for item := range ch {
                process(id, item)
            }
        }(i)
    }
    wg.Wait()
}
```

---

### Question 603: How do you prevent goroutine leaks in producer-consumer patterns?

**Answer:**
Ensure that if the consumer stops (e.g., error or context cancel), the producer doesn't block forever trying to send to the channel.
**Solution:** Use `select` with `ctx.Done()`.

```go
func produce(ctx context.Context, out chan<- int) {
    for {
        select {
        case <-ctx.Done():
            return // Stop producing
        case out <- rand.Int():
            // Sent
        }
    }
}
```

---

### Question 604: How would you create a semaphore in Go?

**Answer:**
Use a buffered channel. It limits concurrent access to a resource.

```go
type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
    return make(chan struct{}, n)
}

func (s Semaphore) Acquire() { s <- struct{}{} }
func (s Semaphore) Release() { <-s }

// Usage
sem.Acquire()
go doWork()
sem.Release()
```

---

### Question 605: What’s the difference between `sync.WaitGroup` and `sync.Cond`?

**Answer:**
- **`WaitGroup`:** Waits for a *count* of goroutines to finish. (Join pattern).
- **`Cond`:** A condition variable. Allows goroutines to suspend execution and wait for a *signal* (event) to occur (e.g., "Queue is not empty"). Broadscast wakes up *all* waiting goroutines.

---

### Question 606: How do you implement a pub-sub model in Go?

**Answer:**
Maintain a map of topics to list of channels.
When `Publish(topic, msg)` is called, iterate over the list and non-blocking send to subscribers (or use a dedicated goroutine per sub).

---

### Question 607: How do you use a context to timeout multiple goroutines?

**Answer:**
Pass the same `Context` to all of them. `context.WithTimeout` propagates the cancellation signal.

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

go worker(ctx)
go worker(ctx)
// Inside worker:
select {
case <-ctx.Done():
    return // Clean up
}
```

---

### Question 608: How do you build a rate-limiting queue with channels?

**Answer:**
Combine a buffered channel (queue) with a Ticker.

```go
requests := make(chan int, 100)
ticker := time.NewTicker(200 * time.Millisecond) // 5 req/sec

go func() {
    for req := range requests {
        <-ticker.C // Wait for tick
        process(req)
    }
}()
```

---

### Question 609: What is a worker pool, and how do you implement it?

**Answer:**
A fixed number of goroutines that pull tasks from a shared channel. Helps control resource usage.
See Question 602 (Fan-out). Ideally, the number of workers matches `runtime.NumCPU()`.

---

### Question 610: How do you handle backpressure in channel-based designs?

**Answer:**
Backpressure occurs when the consumer is slower than the producer.
1.  **Bounded Channels:** Use unbuffered or small-buffered channels. This blocks the producer, naturally slowing it down.
2.  **Dropping:** `select { case ch <- msg: default: log.Println("Dropping") }` (Load Shedding).

---

### Question 611: How do you gracefully shut down workers?

**Answer:**
1.  **Close Channel:** `close(jobs)`. Workers loop `for j := range jobs` will terminate once channel is empty and closed.
2.  **WaitGroup:** Main waits for workers to exit.

```go
close(jobs)
wg.Wait()
```

---

### Question 612: How do you use `sync.Cond` for event signaling?

**Answer:**
It requires a `Locker` (Mutex).

```go
var mu sync.Mutex
cond := sync.NewCond(&mu)

// Waiter
go func() {
    mu.Lock()
    for !condition {
        cond.Wait() // Unlocks and suspends
    }
    // Condition met
    mu.Unlock()
}()

// Signaler
mu.Lock()
condition = true
cond.Signal() // Wakes one waiter
mu.Unlock()
```

---

### Question 613: How do you prioritize tasks in concurrent processing?

**Answer:**
Go channels are FIFO, no priority support.
**Solution:**
1.  Use two channels: `highPriority` and `lowPriority`.
2.  Worker uses `select` with priority check:

```go
select {
case job := <-highPriority:
    process(job)
default:
    select {
    case job := <-highPriority: process(job)
    case job := <-lowPriority: process(job)
    }
}
```

---

### Question 614: How do you avoid starvation in goroutines?

**Answer:**
Starvation happens if a goroutine never gets CPU time (e.g., tight loops without function calls/preemption points in old Go versions).
**Fix:**
- In newer Go (1.14+), the scheduler is preemptive (asynchronously signals loose loops).
- Avoid holding locks for long durations.
- Ensure all channels are serviced fairly.

---

### Question 615: How do you detect race conditions without `-race` flag?

**Answer:**
It is extremely difficult.
- **Code Review:** Look for shared mutable state accessed without locks.
- **Tools:** `staticcheck` can find some issues.
- **Panic:** Sometimes race conditions cause "concurrent map read and map write" panics at runtime.
**Best Practice:** Always use `-race` in CI/Test.

---

### Question 616: How do you trace execution flow in concurrent systems?

**Answer:**
1.  **Trace ID:** Pass `ctx` with trace ID everywhere.
2.  **Structured Logging:** Log "Worker 5 started job X", "Worker 5 finished job X".
3.  **Go Trace Tool:** `go tool trace` visualizes how goroutines hand off control and block.

---

### Question 617: How do you implement exponential backoff with retries in goroutines?

**Answer:**
Simulate wait before retrying.

```go
func retry(op func() error) {
    delay := 100 * time.Millisecond
    for i := 0; i < 5; i++ {
        err := op()
        if err == nil { return }
        time.Sleep(delay)
        delay *= 2
    }
}
```

---

### Question 618: How do you structure long-running daemons with concurrency?

**Answer:**
Use a "Supervisor" pattern.
- **Main:** Starts services. Listens for OS Signals (SIGINT).
- **Service:** Accepts `ctx`. Runs loop `for { select { case <-ctx.Done(): return } }`.
- **ErrGroup:** Use `errgroup.Group` to manage multiple daemon services. If one crashes, it cancels the context for all.

---

### Question 619: How would you implement circuit breakers in Go?

**Answer:**
State machine: Closed (Normal) -> Open (Error threshold reached, fail fast) -> Half-Open (Test recovery).
Use `Sony/gobreaker` library.
It wraps a function execution. If failures > X%, prevents calls for Y seconds.

---

### Question 620: How do you handle concurrent map access with minimal locking?

**Answer:**
1.  **`sync.RWMutex`:** Allows multiple readers, one writer. Good for Read-Heavy workloads.
2.  **`sync.Map`:** Specialized for append-only caches or when keys are disjoint.
3.  **Sharding:** Create `[32]*sync.Mutex` and `[32]map`. Hash key to pick shard. Reduces lock contention by 32x.

---
