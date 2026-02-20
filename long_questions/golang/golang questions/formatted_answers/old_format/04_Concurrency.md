# 🟣 **61–80: Concurrency and Goroutines**

### 62. What are goroutines?
"Goroutines are **lightweight threads** managed by the Go runtime, not the Operating System. They are the fundamental unit of concurrency in Go.

Technically, they start with a tiny 2KB stack that grows and shrinks dynamically, whereas an OS thread might take 1MB. This efficient memory usage means I can spin up tens of thousands of goroutines on a single machine without crashing it.

I treat them as 'fire-and-forget' functions. When I use the `go` keyword, the function runs independently. However, I always ensure I have a plan for how they will finish or communicate back, otherwise, I risk leaks."

#### Indepth
Goroutines are **multiplexed** onto OS threads (M:N scheduling). The Go Scheduler uses a "work-stealing" algorithm. If a processor runs out of work, it steals goroutines from another processor's queue. This ensures high CPU utilization without manual thread management.

---

### 63. What do you start a goroutine?
"I simply use the keyword `go` followed by a function invocation. It’s syntactically the easiest concurrent model I’ve ever used.

For example, `go processBytes(data)`. This immediately returns control to the main function, scheduling `processBytes` to run on the Go runtime's thread pool.

One catch is that the `main` function doesn't wait for goroutines. If `main` returns, the program exits and kills all running goroutines. So, I always use a **WaitGroup** or a **Done Channel** to ensure important background work completes before the program shuts down."

#### Indepth
Under the hood, `go func()` allocates a structure `g` on the heap and adds it to the local run queue of the current Logic Processor (`P`). It does minimal setup, usually taking nanoseconds, compared to milliseconds for an OS thread spawn.

---

### 64. What is a channel in Go?
"A **channel** is a typed conduit that allows goroutines to communicate and synchronize execution. It’s the embodiment of Go's philosophy: 'Share memory by communicating.'

Under the hood, it’s a thread-safe queue with locking built-in. When I write `ch <- value`, the value is copied into the channel. When I read `val := <-ch`, it’s copied out.

I use them not just for data, but for signaling. A receive operation blocks until data is available, which makes them perfect for coordinating workflow steps without messy implementations of mutexes or condition variables."

#### Indepth
Channels use an internal `hchan` struct protected by a `mutex`. Copying data into/out of the channel involves memory copy (`memmove`). For very large structs, passing pointers over channels is faster, but value passing is safer to prevent race conditions (shared ownership).

---

### 65. What is the difference between buffered and unbuffered channels?
"It comes down to **blocking behavior**. An **unbuffered channel** has no capacity; certain synchronicity is enforced because the sender *must* wait for a receiver to be ready, and vice versa.

A **buffered channel** has a queue size (e.g., `make(chan int, 10)`). The sender only blocks if the buffer is **full**. The receiver only blocks if the buffer is **empty**.

I generally prefer unbuffered channels for strict synchronization (like passing a 'result' or 'done' signal). I use buffered channels when I need to decouple the producer from the consumer—handling 'bursty' traffic where the producer might be temporarily faster than the worker."

#### Indepth
A buffered channel of size 1 can act as a **Mutex**. `ch <- 1` locks, `<-ch` unlocks. However, `sync.Mutex` is optimized for this specific case (using atomic CPU instructions) and is generally faster and clearer for simple critical sections.

---

### 66. How do you close a channel?
"I use the built-in `close(ch)` function. It’s a signal to all receivers that no more values will ever be sent on this channel.

Technically, this sets a flag on the channel structure. Receivers observing a closed channel will immediately return a zero-value and `false` (if checking the comma-ok idiom), rather than blocking.

Crucially, I **only** close channels from the sender side. Closing from the receiver side or closing a channel twice causes a runtime panic. If I have multiple senders, I typically use a separate synchronization mechanism (like a `WaitGroup`) to decide when to close it."

#### Indepth
Closing a channel isn't about freeing memory (GC does that). It's strictly about control flow signals. A common pattern for "N senders, 1 receiver" is to have the receiver close a separate struct `done` channel that the senders listen to, telling them to stop sending.

---

### 67. What happens when you send to a closed channel?
"It **panics** immediately. This is a fatal runtime error.

This strictness ensures integrity: you shouldn't be sending data to a stream that has declared 'I am finished.'

To avoid this, I design my systems so that the 'owner' of the channel (the one writing to it) is the only one responsible for closing it. If `send` might occur after close, strictly wrapping the send in a select with a 'done' context check is a common pattern I use."

#### Indepth
This panic is guaranteed to happen to prevent hard-to-debug data corruption. Unlike a nil pointer (which is a crash due to invalid memory), this is a logic enforcement. There is no `safeSend` function in Go; you must know the state of your system.

---

### 68. How to detect a closed channel while receiving?
"I use the **comma-ok idiom**: `value, ok := <-ch`.

If the channel is open and has data, `ok` is `true`. If the channel is closed and empty, `value` will be the zero-value (like 0 or "") and `ok` will be `false`.

I rely on this heavily inside `for range` loops. A `for msg := range ch` loop automatically breaks when the channel is closed, which is the cleanest way to drain a worker queue until shutdown."

#### Indepth
Note that a closed channel returns the *zero value* forever. It doesn't block. This is a common bug source: if you have a `select` loop and one channel closes, that case becomes non-blocking and generates an infinite loop of zero values, spiking CPU to 100%. Always set the channel variable to `nil` after seeing it close to disable that case in the select.

---

### 69. What is the `select` statement in Go?
"The `select` statement is like a `switch`, but exclusively for channel operations. It allows a goroutine to wait on multiple communication operations simultaneously.

It blocks until one of its cases can proceed. If multiple cases are ready (e.g., data arrived on two different channels), `select` picks one **at random**. This randomness prevents one active channel from starvation.

I use it in almost every long-running backend service to handle **cancellation** alongside business logic—waiting for `<-ctx.Done()` in one case and `<-jobQueue` in another."

#### Indepth
The Go runtime randomizes the polling order of cases in a `select` statement to ensure fairness. If it was deterministic (always checking top-to-bottom), the top case could starve the bottom ones if it always had high traffic.

---

### 70. How do you implement timeouts with `select`?
"I add a `case` that receives from `time.After(duration)`.

`time.After` returns a channel that sends the current time after the delay. So my structure looks like:
`select { case res := <-resultCh: handle(res); case <-time.After(2 * time.Second): return ErrTimeout }`.

Usage-wise, this is critical for production systems. I never want to block indefinitely on a network call or an external service. Hard timeouts ensure my system stays responsive even if dependencies hang."

#### Indepth
Beware that `time.After` leaks resources if the select picks another case, until the timer fires. In high-throughput loops, this accumulates millions of timers. Use `time.NewTimer` and explicitly `Stop()` it to prevent memory leaks in tight loops.

---

### 71. What is a `sync.WaitGroup`?
"It’s a synchronization primitive that waits for a collection of goroutines to finish. Think of it as a thread-safe counter.

I call `wg.Add(1)` before starting a task. The worker calls `wg.Done()` (decrement) when finished. The main thread calls `wg.Wait()`, blocking until the counter hits zero.

It’s my go-to tool for **fan-out/fan-in** patterns—like spawning 10 scrapers to fetch URLs in parallel and waiting for all of them to finish before aggregating the results."

#### Indepth
`WaitGroup` must be passed by pointer (`*sync.WaitGroup`), never by value! Copying a WaitGroup copies its internal counter state, leading to a deadlock where the worker signals the *copy*, but the main thread waits on the *original*. `go vet` catches this.

---

### 72. How does `sync.Mutex` work?
"A `Mutex` (Mutual Exclusion) creates a critical section where only *one* goroutine can execute at a time. It prevents race conditions on shared memory.

I use `mu.Lock()` to claim exclusive access and `mu.Unlock()` to release it—almost always in a `defer` statement right after locking to ensure I don't forget it if a panic occurs.

While channels are great for data flow, I prefer `Mutex` for **state**. If I just need to safely increment a counter or update a map in a struct, a Mutex is often faster and simpler than setting up a dedicated channel-coordinating goroutine."

#### Indepth
`sync.Mutex` has two modes: **Normal** and **Starvation**. In Starvation mode (triggered if a waiter waits >1ms), the mutex ownership is handed directly to the first waiter, bypassing new arriving goroutines. This prevents "tail latency" spikes where a request waits forever.

---

### 73. What is `sync.Once`?
"It’s a utility that ensures a piece of code is executed **exactly once**, regardless of how many goroutines call it simultaneously.

Internally, it uses an atomic counter and a mutex. The first caller wins the race and executes the function; subsequent callers wait until it finishes, then skip execution.

I use this for **lazy singleton initialization**—like connecting to a database or loading a heavy configuration file only when the first request arrives, rather than at application startup."

#### Indepth
`sync.Once` is implemented using an atomic variable `done`. The "fast path" just checks the atomic integer (cheap). Only if it's 0 does it acquire the slow mutex to execute the function. This makes `Do()` extremely low-overhead for frequent calls.

---

### 74. How do you avoid race conditions?
"The primary mantra is: **'Do not communicate by sharing memory; instead, share memory by communicating.'**

This means I prefer passing data copies over **channels** rather than multiple threads modifying the same pointer. If I *must* share memory (like a global cache), I protect it strictly with `sync.Mutex` or `sync.RWMutex`.

To verify my safety, I always run my tests with the `-race` flag (`go test -race`). The Go race detector is incredibly good at finding unlocked concurrent access that I missed during code review."

#### Indepth
The `-race` flag instruments code with "happens-before" annotations. It tracks memory access at runtime. It increases memory usage by 5-10x and execution time by 2-20x, so don't run it in production, but *always* run it in CI/CD pipeline.

---

### 75. What is the Go memory model?
"It’s the specification that defines **visibility**—guaranteeing when a variable written by one goroutine is visible to another.

The core rule is: without explicit synchronization (channels, mutexes, waiter groups), there is **no guarantee** that Goroutine B sees the write from Goroutine A. Compilers reorder code, and CPU caches delay writes.

This implies that 'clever' lock-free code using plain boolean flags is usually broken. I stick to the standard synchronization primitives which act as memory barriers, flushing caches and enforcing order."

#### Indepth
Go's memory model is weaker than Java's (volatile) but stronger than C++'s relaxed atomics. It guarantees that a send on a channel *happens before* the receive completes. This causality chain is what allows generic synchronization without understanding CPU cache line invalidation.

---

### 76. How do you use `context.Context` for cancellation?
"The `Context` is the standard way to carry cancellation signals across API boundaries and goroutines.

I create a context with `ctx, cancel := context.WithCancel(parent)`. When I call `cancel()`, the `ctx.Done()` channel closes. All downstream functions listening to this channel stop their work and return immediately.

I use this for every request path. If a user closes their browser tab, I cancel the context, which stops the DB query and saves resources. It’s a required pattern for any robust Go server."

#### Indepth
`context` values are immutable. `WithCancel` returns a *new* context child. This forms a tree. Cancelling the parent cancels all children. Use `context.Background()` for the root and `context.TODO()` when you're unsure (refactoring placeholders).

---

### 77. How to pass data between goroutines?
"I primarily use **channels**. They are the safest and most idiomatic way to transfer ownership of data.

If I have a Producer generating 'Jobs' and a Consumer processing them, a buffered channel acts as the perfect glue. It handles the locking and queuing logic automatically.

For simple configuration data or read-only 'context', I might pass immutable structs or use `context.Context` values, but for the actual flow of business data, channels are my default choice."

#### Indepth
Don't use channels for everything. If you have "referentially transparent" data (just values), passing them as function arguments is faster (register passing vs channel locking). Channels are for specific *coordination* or *handoff* points.

---

### 78. What is the `runtime.GOMAXPROCS()` function?
"It controls the number of **OS threads** that can execute Go code simultaneously. By default, it equals the number of CPU cores available.

Changing it limits how much actual parallelism the runtime utilizes. If I set it to 1, my program becomes strictly concurrent but single-threaded (no true parallel execution).

I rarely touch this in code. However, in **Kubernetes**, I verify it matches the CPU quota (using `automaxprocs`). If my container has 2 CPUs but the Node has 64, Go might spin up 64 threads, causing excessive context switching and throttling. Matching it to the quota fixes this."

#### Indepth
Prior to Go 1.5, GOMAXPROCS defaulted to 1. This historic artifact is why some old tutorials say "Go isn't parallel by default". Today, it defaults to `runtime.NumCPU()`. Manipulating this is rarely needed unless you are running in a containerized environment with "fractional CPUS" (like 0.5 CPU).

---

### 79. How do you detect deadlocks in Go?
"A deadlock happens when all goroutines are waiting for each other, and no one can proceed. Go has a built-in detector that panics with `fatal error: all goroutines are asleep - deadlock!` if the runtime sees zero executable goroutines.

However, this only catches complete global deadlocks. **Partial deadlocks** (where 2 threads are stuck but others run) are harder.

I detect those by using `pprof` to inspect the stack traces of stuck goroutines or by using timeouts on all channel/lock operations. If a lock takes >5 seconds, I panic or log a stack trace to debug why it was never released."

#### Indepth
You can retrieve the stack trace of all goroutines programmatically using `runtime.Stack(buf, true)`. Sending `SIGQUIT` (Ctrl+\) to a running Go program performs a "core dump" to stderr, which is invaluable for debugging stuck processes in production.

---

### 80. What are worker pools and how do you implement them?
"A worker pool is a pattern to limit concurrency. Instead of spawning `go func()` for *every* incoming request (which could exhaust RAM), I start a fixed number of workers (e.g., 5).

I create a `jobs` channel. I start 5 goroutines that `range` over the `jobs` channel. The main handler simply sends requests into the channel.

This ensures my database is never hit by more than 5 concurrent queries, acting as a natural backpressure mechanism. If the buffer fills, the callers wait, rather than the system crashing."

#### Indepth
A more advanced pattern is the **Semaphore** pattern using a buffered channel. `sem := make(chan struct{}, 5)`. Before starting a job, `sem <- struct{}{}`. When done, `<-sem`. This limits active goroutines without pre-allocating a fixed pool of workers.

---

### 81. How to write concurrent-safe data structures?
"I wrap the data structure (like a map) in a struct that has a `sync.RWMutex`.

I ensure that every read acquires a **Read Lock** (`RLock`) and every write acquires a **Write Lock** (`Lock`). Alternatively, I can use `sync.Map` for specific use cases like append-only caches, but usually, a standard map with a mutex is clearer.

The key is strict discipline: **never** expose the underlying map directly. Only access it through methods that handle the locking."

#### Indepth
`sync.Map` is specialized for cases where keys are stable (write once, read many) or disjoint (different goroutines access different keys). For general "R/W cache" use cases, a `map` + `RWMutex` is often 2x faster due to lower interface overhead (no type assertions).
