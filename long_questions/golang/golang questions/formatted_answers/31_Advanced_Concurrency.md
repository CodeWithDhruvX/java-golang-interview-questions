# 🔸 **601–620: Advanced Concurrency Patterns**

### 601. How do you implement a fan-in pattern in Go?
"Fan-in merges multiple input channels into one output channel.
multiplexing.
I launch a goroutine for each input:
`for _, ch := range inputs { go func(c) { for v := range c { out <- v }; wg.Done() }(ch) }`.
When all inputs close (`wg.Wait()`), I close the output channel."

#### Indepth
`reflect.Select` can fan-in dynamic channels at runtime, but it's slow (reflection overhead). For high performance, stick to the fixed-concurrency loop shown above. Also, ensure the output channel has a buffer (`make(chan T, numInputs)`) to prevent producers from blocking each other on the final merge.

---

### 602. How do you implement a fan-out pattern in Go?
"Fan-out distributes work from one channel to multiple workers.
`for i := 0; i < numWorkers; i++ { go worker(inputChan) }`.
The workers compete for items from the shared channel.
It automatically load-balances: if Worker A is heavy, Worker B picks up the next item. It's the basis of all worker pools."

#### Indepth
**Bounded Concurrency**. Never just `go worker()` based on input size. If 1,000,000 items arrive, spawning 1M goroutines will kill the scheduler (and memory). Always use a fixed pool size (e.g., `runtime.NumCPU()`) to process an infinite stream of work.

---

### 603. How do you prevent goroutine leaks in producer-consumer patterns?
"A leak happens if a sender blocks forever on a channel no one reads.
Rules:
1.  **Ownership**: The logical owner (producer) closes the channel.
2.  **Context**: Receivers check `ctx.Done()` to exit early.
3.  **Capacity**: Ensure the receiver can drain the channel, or use a non-blocking send with `select`."

#### Indepth
Monitor `runtime.NumGoroutine()`. If this metric climbs steadily over time, you have a leak. Use `pprof` with the `goroutine` profile to find where they are stuck (usually `runtime.gopark` waiting on a channel send/receive that will never happen).

---

### 604. How would you create a semaphore in Go?
"I use a buffered channel.
`sem := make(chan struct{}, capacity)`.
Acquire: `sem <- struct{}{}`.
Release: `<-sem`.
Since the buffer size is fixed, only N goroutines can 'acquire' at once. The N+1th will block. This is how I limit database connections or API concurrency."

#### Indepth
Weighted Semaphores (`golang.org/x/sync/semaphore`) are useful when different tasks consume different amounts of resources. Task A might need 1 unit, Task B needs 5. The standard channel approach only supports weight=1. `Weighted` allows `Acquire(ctx, 5)`.

---

### 605. What’s the difference between sync.WaitGroup and sync.Cond?
"**WaitGroup**: Wait for N events to generic *finish* (Count down).
**Cond**: Wait for a *signal* that a condition has *changed* (Broadcast).
I use WaitGroup 99% of the time.
I use Cond only for complex coordination, like a queue where multiple consumers are sleeping and I need to wake them *all* up when an item arrives."

#### Indepth
**Spurious Wakeups**. `Cond.Wait()` can technically return even if not signaled (though rare). Always wrap `Wait()` in a loop checking the condition: `for !condition { cond.Wait() }`. This ensures correctness even if the OS wakes the thread unexpectedly.

---

### 606. How do you implement a pub-sub model in Go?
"I use a central Broker map.
`subscribers map[string][]chan Msg`.
`func Publish(topic, msg)`: Iterates the slice and sends to each channel.
`func Subscribe(topic)`: Creates a chan, adds to map, returns chan.
I lock the map with `RWMutex` during updates. This is a simple, effective in-process event bus."

#### Indepth
Buffer Bloat. If one subscriber is slow, it blocks the `Publish` loop (and all other subscribers) unless channels are buffered. If they are buffered and fill up, you must decide: Drop the message? (Lossy) or Block? (Slow). Sophisticated buses use a dedicated goroutine per subscriber to isolate them.

---

### 607. How do you use a context to timeout multiple goroutines?
"I create a single context with timeout.
`ctx, cancel := context.WithTimeout(parent, 5*time.Second)`.
I pass this `ctx` to all 10 goroutines.
Inside each: `select { case <-ctx.Done(): return error }`.
When the 5s timer hits, the channel closes, and *all* 10 goroutines receive the signal instantly and abort."

#### Indepth
Go 1.21 introduced `context.AfterFunc(ctx, func())`. It allows you to schedule a cleanup function to run *immediately* when the context is cancelled, without needing to spin up a dedicated goroutine to wait on `<-ctx.Done()`. This saves resources in high-concurrency timeout logic.

---

### 608. How do you build a rate-limiting queue with channels?
"I use a `time.Ticker` alongside the job channel.
Worker loop:
`for job := range jobs { <-ticker.C; process(job) }`.
The worker *must* wait for a tick before taking the next job.
If the ticker is 100ms, the worker can only process 10 jobs/second. This smooths out bursty traffic."

#### Indepth
**Token Bucket**. For burstier/flexible limits (e.g., "avg 10/sec, but allow burst of 50"), use `golang.org/x/time/rate`. `limiter.Wait(ctx)` blocks until a token is available. Tickers are strict interval based; Token Buckets allow borrowing time.

---

### 609. What is a worker pool, and how do you implement it?
"1.  `jobs := make(chan Job, 100)`
2.  Spawn N workers: `go worker(jobs, results)`.
3.  Feed jobs into `jobs`.
4.  Close `jobs` when done.
It keeps my active goroutine count constant (N) regardless of the number of items (M), preventing memory explosion."

#### Indepth
Dynamic Resizing. You might want to scale workers based on queue depth. This is hard with the standard loop pattern. You need a manager goroutine that monitors `len(jobs)` and launches new workers (up to a max) or sends a "poison pill" to kill idle workers when traffic drops.

---

### 610. How do you handle backpressure in channel-based designs?
"If the producer is faster than consumer, the buffered channel fills up.
Once full, the producer **blocks** on send (`ch <- item`).
This naturally slows down the producer (Backpressure).
If I can't block the producer (like an HTTP handler), I must start dropping items or returning 503 errors (`select { case ch <- item: default: return 503 }`)."

#### Indepth
**Ring Buffer**. For drop-oldest behavior (logging), channels are bad (blocking). Use a Ring Buffer. If full, overwrite the read pointer. This guarantees the producer never blocks, but the consumer might lose old data. This is how `log/syslog` often works.

---

### 611. How do you gracefully shut down workers?
"I close the `jobs` channel.
`close(jobs)`.
Workers: `for job := range jobs { ... }`.
The loop terminates when the channel is empty and closed. The workers exit naturally.
I wrap this with a `WaitGroup` to ensure the main thread waits for them to cleanly finish current tasks."

#### Indepth
`context.WithCancel` is the modern way to signal shutdown, especially if you have multiple layers of workers. Closing the channel works for the *immediate* consumer, but a propagated Context cancellation reaches the database driver, HTTP client, and file reader simultaneously, stopping the entire pipeline.

---

### 612. How do you use sync.Cond for event signaling?
"1.  Lock the associated Mutex.
2.  Check condition.
3.  If not met, `cond.Wait()` (this releases lock and sleeps).
4.  When another goroutine changes state, it calls `cond.Signal()` (wake one) or `cond.Broadcast()` (wake all).
It’s reusable and broadcasts a 'state change', unlike channels which pass values."

#### Indepth
Critically, `cond.Signal()` doesn't transfer ownership or data. It just wakes a thread. That thread must then re-acquire the lock and check the data. It is meant for "Something changed, go look" scenarios, not "Here is the data" (use Channels for that).

---

### 613. How do you prioritize tasks in concurrent processing?
"Go channels are FIFO. No priority.
To implement priority, I use **two** channels: `high` and `low`.
Worker:
`select { case job := <-high: do(job); default: }`
`select { case job := <-high: do(job); case job := <-low: do(job); }`.
I check the high channel *first* (non-blocking). If empty, I wait on both."

#### Indepth
**Double Select Trick**. The example logic biases slightly but `select` is random when both are ready. To strictly enforce priority, you need two select blocks:
`select { case v := <-high: return v; default: }`
`select { case v := <-high: return v; case v := <-low: return v; }`
This guarantees checking `high` before entering the random wait.

---

### 614. How do you avoid starvation in goroutines?
"Starvation happens if a high-priority task hogs the CPU.
In the priority example above, if `high` is always full, `low` never runs.
Fix: Every 10th loop, check `low` even if `high` has data.
Or allow the Go runtime scheduler to preempt goroutines (which it does every 10ms automatically)."

#### Indepth
Preemption (Go 1.14+) fixed most starvation issues by allowing the runtime to pause tight loops (`for {}`). Before this, a tight loop could hang the scheduler on a CPU core. However, explicit `runtime.Gosched()` is still useful in cooperatively multitasking systems to "yield" the CPU voluntarily.

---

### 615. How do you detect race conditions without `-race` flag?
"It's extremely hard.
Static analysis (`go vet`) finds some lock copying issues.
Code Review: Look for shared maps/slices accessed by multiple goroutines without a mutex.
But realistically? I can't. The `-race` flag is unique and essential. I run it in CI heavily."

#### Indepth
The Race Detector uses the C/C++ ThreadSanitizer (TSan). It keeps a "shadow state" of memory to track last-write timestamps. This is why it has high overhead. It can detect races even if they didn't cause a crash *in that specific run*, as long as the code path was executed.

---

### 616. How do you trace execution flow in concurrent systems?
"I use **Distributed Tracing** (TraceID) even inside a monolith.
I pass `ctx` everywhere.
Log lines include `trace_id`.
This lets me grep logs for a single request across multiple goroutines.
Without this, 1000 interleaved logs from 100 requests are unreadable."

#### Indepth
**Context Propagation**. Libraries like OpenTelemetry automatically extract the TraceID from incoming HTTP headers (`traceparent`) and stash it in the `context.Context`. Your `slog` or `zap` logger should extract this from the context automatically (`logger.WithContext(ctx).Info(...)`).

---

### 617. How do you implement exponential backoff with retries in goroutines?
"Loop with sleep.
`delay := 1 * time.Second`
`for i:=0; i<max; i++ { err := do(); if err == nil return; time.Sleep(delay); delay *= 2 }`.
I always check `ctx.Done()` during sleep so the retry loop handles cancellation immediately."

#### Indepth
Use `math.Pow(2, i)` plus **Jitter**. Pure exponential backoff (`1s, 2s, 4s, 8s`) can cause synchronized retry storms. Always add `random(0, 1000ms)` to the delay. This spreads out the retries so your database doesn't get hit by 1000 requests all exactly 4 seconds after a restart.

---

### 618. How do you structure long-running daemons with concurrency?
"I use an **ErrGroup**.
`g, ctx := errgroup.WithContext(ctx)`.
`g.Go(func() { return server.ListenAndServe() })`.
`g.Go(func() { return consumer.Run() })`.
If *any* goroutine returns an error, the context is canceled, signaling *all* others to shutdown. It’s the perfect supervisor pattern."

#### Indepth
`oklog/run` is another popular alternative. It handles signal trapping (SIGTERM) and actor groups strictly. However, `errgroup` is standard (experimental stdlib) and arguably easier. Just remember: `errgroup` waits for *all* goroutines to return, so if one hangs, the `Wait()` hangs.

---

### 619. How would you implement circuit breakers in Go?
"I wrap the critical call.
Identify failure: `if err != nil { failures++ }`.
Trip: `if failures > threshold { state = Open }`.
In `Open` state, return logical error immediately.
After timeout, allow 1 test request (Half-Open).
I use a mutex (or atomic) to protect the state variable."

#### Indepth
`sony/gobreaker` is the industry standard Go implementation. It tracks consecutive failures or failure ratios. Crucially, don't just count *errors*; count *5xx errors*. A 404 Not Found shouldn't trip the circuit breaker, but a 503 Timeout definitely should.

---

### 620. How do you handle concurrent map access with minimal locking?
"Standard `sync.Map` or `RWMutex`.
With `RWMutex`: `RLock()` for reads (parallel), `Lock()` for writes (exclusive).
If writes are rare, this is very fast.
If writes are frequent, I shard the map (32 maps with 32 locks) to reduce contention on any single bucket."

#### Indepth
`sync.Map` is optimized for two cases: (1) Keys are written once and read many times (cache), or (2) Disjoint sets of keys are used by different goroutines. For general N-way Read/Write, a standard `map` + `RWMutex` is often faster and type-safe (generics). `sync.Map` uses `interface{}`/`any` which loses type safety.
