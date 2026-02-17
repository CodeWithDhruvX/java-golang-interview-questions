# 🟢 Go Theory Questions: 601–620 Advanced Concurrency Patterns

## 601. How do you implement a fan-in pattern in Go?

**Answer:**
Fan-In merges multiple input channels into one output channel.

```go
func FanIn(inputs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    for _, ch := range inputs {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for n := range c { out <- n }
        }(ch)
    }
    go func() { wg.Wait(); close(out) }()
    return out
}
```
This is useful when you have sharded workers (e.g., 3 downloaders) and want to process their results in a single aggregator.

---

## 602. How do you implement a fan-out pattern in Go?

**Answer:**
Fan-Out distributes work from one channel to multiple workers.

```go
jobs := make(chan Job, 100)
for i := 0; i < NumWorkers; i++ {
    go worker(jobs)
}
```
The key is that all workers range over the **same** channel. The runtime automatically load-balances the items. If Worker 1 is busy, Worker 2 picks up the next item.

---

## 603. How do you prevent goroutine leaks in producer-consumer patterns?

**Answer:**
Leaks happen if the receiver stops, but the producer keeps sending (blocks forever), or vice-versa.

1.  **Cancellation**: Pass `context.Context` to the producer. If `ctx.Done()` is closed, return immediately.
2.  **Close**: The Producer must close the channel when done, signaling consumers to exit.
3.  **Select**: Use `select { case ch <- item: ... case <-done: return }` to ensure the send operation is interruptible.

---

## 604. How would you create a semaphore in Go?

**Answer:**
We use a **Buffered Channel**.

`sem := make(chan struct{}, MaxConcurrency)`
Acquire: `sem <- struct{}{}` (Blocks if full).
Release: `<-sem`.

We wrap the critical section:
```go
sem <- struct{}{}
go func() {
    defer func() { <-sem }()
    process()
}()
```
For advanced features (weighted semaphore, try-acquire), use `golang.org/x/sync/semaphore`.

---

## 605. What’s the difference between sync.WaitGroup and sync.Cond?

**Answer:**
**WaitGroup**: Waits for a group of goroutines to *finish*. (Count down to zero). Use for Batch Processing.

**Cond**: Waits for an *event* or *state change*. (Signal/Broadcast).
Use `Cond` when many goroutines are waiting for a condition (e.g., "Buffer is not empty") and you want to wake them up efficiently without busy spinning.

---

## 606. How do you implement a pub-sub model in Go?

**Answer:**
(See Q 586).
A central Broker struct holds a map of topics to subscribers.

`map[string][]chan Payload`
`Subscribe(topic)` adds a chan.
`Publish(topic, msg)` iterates the slice and sends.
Crucially, usually we use a **Non-Blocking Send** in the publish loop to prevent one slow subscriber from freezing the entire broker.

---

## 607. How do you use a context to timeout multiple goroutines?

**Answer:**
You create a context with timeout:
`ctx, cancel := context.WithTimeout(parent, 5*time.Second)`
`defer cancel()`

You pass this `ctx` to all goroutines.
Inside them:
```go
select {
case <-workCh:
    // ...
case <-ctx.Done():
    return // Clean up and exit
}
```
When the 5s timer hits, the channel closes. All goroutines receive the signal almost simultaneously and terminate.

---

## 608. How do you build a rate-limiting queue with channels?

**Answer:**
We use `time.Ticker` creates the "heartbeat".

```go
ticker := time.NewTicker(200 * time.Millisecond) // 5 req/sec
for req := range requests {
    <-ticker.C // Wait for tick
    process(req)
}
```
This forces the loop to run at max speed of the ticker.
Token Bucket is more flexible (allows bursts), but Ticker is the simplest "Leaky Bucket" implementation (strict interval enforcement).

---

## 609. What is a worker pool, and how do you implement it?

**Answer:**
A Worker Pool restricts concurrency.

Structure:
1.  **Job Channel**: `chan Job`.
2.  **Result Channel**: `chan Result`.
3.  **Dispatcher**: Spawns N workers.
4.  **Worker**: Loop `range jobs`.

This decoupling allows the Producer to push 1M jobs instantly (buffered), while the consumers process them at a steady, safe rate (e.g., 50 concurrent DB connections).

---

## 610. How do you handle backpressure in channel-based designs?

**Answer:**
Backpressure means "Stop sending, I'm full."
Unbuffered channels provide **Natural Backpressure**. The sender blocks until the receiver is ready.

If using buffered channels, monitoring `len(ch)` vs `cap(ch)` can act as a signal.
If `len(ch) > 0.9 * cap(ch)`, we can reject new API requests (503 Service Unavailable) or tell the producer to sleep. This prevents the system from crashing under load (Load Shedding).

---

## 611. How do you gracefully shut down workers?

**Answer:**
1.  Stop sending: `close(jobsChannel)`.
2.  Workers finish current item, loop terminates (`range` exits).
3.  Workers call `wg.Done()`.
4.  Main thread waits `wg.Wait()`.

This ensures we don't kill a worker in the middle of a transaction (corrupting data).

---

## 612. How do you use sync.Cond for event signaling?

**Answer:**
`c := sync.NewCond(&sync.Mutex{})`.

**Waiter**:
`c.L.Lock(); for !ready { c.Wait() }; doWork(); c.L.Unlock()`.
Note the `for !ready`. `Wait()` releases the lock and suspends. When it wakes, it re-acquires. We check the condition again because it might have changed (spurious wakeup).

**Signaler**:
`c.L.Lock(); ready = true; c.Broadcast(); c.L.Unlock()`.
`Broadcast` wakes *all* waiters. `Signal` wakes one.

---

## 613. How do you prioritize tasks in concurrent processing?

**Answer:**
Go channels don't support priority. A queue is FIFO.
To implement Priority:
1.  **Two Channels**: `highCh`, `lowCh`.
2.  **Worker Select**:
```go
select {
case job := <-highCh: process(job)
case job := <-lowCh: process(job)
}
```
**Problem**: Go's select is random if both are ready.
**Fix**: Non-blocking check for high priority first.
```go
select {
case job := <-highCh: process(job)
default:
    select {
    case job := <-highCh: process(job)
    case job := <-lowCh: process(job)
    }
}
```

---

## 614. How do you avoid starvation in goroutines?

**Answer:**
Starvation happens when a high-priority process hogs the CPU/Lock, and low-priority ones never run.

In Go, the scheduler includes a "time slice" mechanism to prevent CPU starvation.
For resource starvation (Locks), use `sync.Mutex` (which is vaguely FIFO for waiters).
Avoid spin-locks.
In Priority Queue designs (see Q 613), ensure you service `lowCh` occasionally even if `highCh` is full (e.g., process 1 low for every 10 high).

---

## 615. How do you detect race conditions without `-race` flag?

**Answer:**
You generally **can't** reliably.
But signs include:
1.  "Impossible" values (counter = 1001 when max is 1000).
2.  Random crashes (`concurrent map read and map write`).
3.  Heisenbugs (bugs that vanish when you add `fmt.Println`).

Code Review Audit: Look for shared usage of `map` or `slice` across goroutines without `mu.Lock()`. But rely on the automated `-race` detector in CI; manual detection is error-prone.

---

## 616. How do you trace execution flow in concurrent systems?

**Answer:**
We use **Distributed Tracing** (TraceID) even inside a single process.
Pass `ctx` containing the SpanID to every goroutine.
`ctx, span := tracer.Start(ctx, "worker")`.

When viewing the Trace in Jaeger, you see the "Main" span spawn 10 "Worker" child spans in parallel. This visualizes exactly when they started, how long they overlapped, and which one caused the delay.

---

## 617. How do you implement exponential backoff with retries in goroutines?

**Answer:**
Simple recursive or loop approach.

```go
func retry(op func() error) {
    wait := 100 * time.Millisecond
    for {
        if err := op(); err == nil { return }
        time.Sleep(wait)
        wait *= 2
        if wait > 10*time.Second { wait = 10*time.Second } // Cap
    }
}
```
If calling this inside a goroutine, ensure it respects `context.Context` so the retry loop aborts if the parent request is canceled.

---

## 618. How do you structure long-running daemons with concurrency?

**Answer:**
We use the **ErrGroup** pattern (`golang.org/x/sync/errgroup`).

`g, ctx := errgroup.WithContext(context.Background())`

`g.Go(func() error { return runHTTPServer() })`
`g.Go(func() error { return runMetricsServer() })`
`g.Go(func() error { return consumeKafka() })`

`g.Wait()` blocks until *any* of them returns an error. If one dies, the Context is canceled, triggering the others to shutdown. This ensures the daemon acts as a cohesive unit—if one critical limb fails, the whole body restarts.
