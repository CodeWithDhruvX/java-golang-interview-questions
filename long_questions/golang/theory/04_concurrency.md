# 🟢 Go Theory Questions: 61–80 Concurrency and Goroutines

## 61. What are goroutines?

**Answer:**
Goroutines are functions that execute concurrently with other functions. They are extremely lightweight threads managed by the **Go Runtime**, not the Operating System.

Mechanically, they use an **M:N Scheduling** model. The runtime multiplexes thousands of Goroutines (M) onto a small number of OS threads (N). They start with a tiny 2KB stack that grows dynamically, allowing you to spawn tens of thousands of them on a single machine without crashing RAM.

In the real world, we use them for everything: handling individual HTTP requests (1 request = 1 goroutine), background cron jobs, or processing messages from a queue. The main trade-off is ensuring they exit; if you spawn a goroutine that blocks forever, you create a "goroutine leak" that eventually kills your server.

---

## 62. How do you start a goroutine?

**Answer:**
You simply place the keyword `go` in front of any function call: `go process(data)`.

This returns control immediately to the next line of code, while the function runs in the background. The arguments to the function are evaluated **synchronously** at the moment of the call, but the execution of the body happens asynchronously.

I use this often for **Fire-and-Forget** tasks, like sending an email after a user signs up. The HTTP response returns instantly to the user, and the email sends in the background. The risk is that you can't easily handle errors from a goroutine; you must log them internally or send them back via a channel.

---

## 63. What is a channel in Go?

**Answer:**
Channels are the pipes that connect concurrent goroutines. They allow you to send values from one goroutine to another, providing both **Communication** and **Synchronization**.

Under the hood, a channel is a thread-safe queue (ring buffer) protected by a mutex. When you send to a channel and nobody is reading, the runtime **parks** your goroutine (puts it to sleep) efficiently until a receiver arrives.

We use them to implement "Share memory by communicating." Instead of using complex locks to protect a shared variable, we pass the data itself over a channel. This avoids race conditions by design, though it can lead to deadlocks if you aren't careful about who is sending and who is receiving.

---

## 64. What is the difference between buffered and unbuffered channels?

**Answer:**
An **Unbuffered** channel (`make(chan int)`) has zero capacity. It forces a synchronous handoff: the sender blocks until the receiver is ready. It effectively guarantees that "Make" and "Process" happen at the same time.

A **Buffered** channel (`make(chan int, 10)`) has a queue. The sender can drop off a value and keep going, as long as the buffer isn't full.

In practice, I use unbuffered channels for strict synchronization (orchestrating timing) and buffered channels for **rate limiting** or handling **bursty traffic**. Buffered channels decouple the producer from the consumer, but they can hide bugs—if the consumer dies, the producer might fill the buffer and then block, masking the failure for seconds.

---

## 65. How do you close a channel?

**Answer:**
You use the built-in function `close(ch)`. This sends a special signal to all receivers that "no more data is coming."

Mechanically, closing a channel is a broadcast event. It wakes up *all* parked receivers immediately. Any subsequent receive operation returns the **zero value** instantly. It is essential for terminating `range` loops that are reading from the channel.

The main trade-off is the **Panic Risk**. You cannot close a channel twice, and you cannot close a nil channel. Both cause panics. You must strictly ensure that only the **Sender** closes the channel, never the receiver.

---

## 66. What happens when you send to a closed channel?

**Answer:**
It causes a runtime **Panic**. The program crashes immediately.

This is a strict design choice by the Go team to prevent ambiguous data flow. A closed channel means "Data Stream Ended." Writing to it implies you don't know the stream is dead, which is a logic bug.

This usually happens in a **Fan-In** pattern where multiple goroutines write to one channel. If one closes it, the others panic. To fix this, we typically use a `sync.WaitGroup` to wait for all senders to finish, and then have a separate "coordinator" goroutine close the channel only once the WaitGroup is done.

---

## 67. How to detect a closed channel while receiving?

**Answer:**
We use the "comma-ok" idiom: `val, ok := <-ch`.

If `ok` is `true`, the channel is open and `val` is real data.
If `ok` is `false`, the channel is closed and `val` is just the zero value (like `0` or `""`).

Without this, you can't distinguish between receiving the integer `0` (valid data) and the channel being closed. We use this heavily inside `for/select` loops to break out and stop workers when the job queue finishes. Note that `range` loops handle this check automatically for you.

---

## 68. What is the `select` statement in Go?

**Answer:**
`select` is like a `switch` statement for channels. It lets a single goroutine wait on **multiple** channel operations simultaneously.

Whichever channel is ready first (to send or receive) is the case that executes. If multiple are ready, it picks one at pseudo-random.

This is the fundamental building block of Go concurrency patterns. We use it for **Cancellation** (racing a `Done` channel against a `Work` channel) or **Timeouts**. Without `select`, reading from a channel would block indefinitely, making it impossible to write responsive systems that can handle cancellation.

---

## 69. How do you implement timeouts with `select`?

**Answer:**
You add a case with `time.After(duration)`.

`time.After` returns a channel that sends the current time after a delay. The `select` races your operation against this timer. If your API call takes too long, the timer channel "wins," and you execute the timeout logic.

```go
select {
case res := <-apiCall():
    return res
case <-time.After(2 * time.Second):
    return nil, errors.New("timeout")
}
```
This is critical for microservices. It prevents your service from piling up thousands of hanging requests if a downstream database stalls.

---

## 70. What is a `sync.WaitGroup`?

**Answer:**
It’s a thread-safe counter used to wait for a collection of goroutines to finish.

You call `Add(1)` when you start a task, `Done()` (which decrements the counter) when it finishes, and `Wait()` to block the main thread until the counter hits zero.

It is the standard tool for **Scatter-Gather** patterns. If I need to fetch data from 10 different URLs in parallel, I spawn 10 goroutines, each with a `defer wg.Done()`, and `wg.Wait()` at the end to aggregate the results. The most common bug is calling `Add(1)` *inside* the goroutine—this is a race condition. You must call `Add` *before* spawning the generic.

---

## 71. How does `sync.Mutex` work?

**Answer:**
A Mutex (Mutual Exclusion) locks a section of code so only **one** goroutine can execute it at a time.

`Lock()` blocks other goroutines. `Unlock()` opens the gate. It protects shared variables (like maps or counters) from data corruption during concurrent access.

In the real world, we use this to protect internal state in structs, usually creating **Thread-Safe** objects. The performance is generally better than channels for simple state updates, but the risk is **Deadlocks**—if you forget to unlock (use `defer Unlock()`!) or if you try to lock the same mutex twice in the same goroutine.

---

## 72. What is `sync.Once`?

**Answer:**
`sync.Once` guarantees that a function executes **exactly once**, no matter how many goroutines call it simultaneously.

It uses an atomic counter and a lock internally to ensure safety. The first goroutine runs the code; all subsequent callers block and then return immediately once the first one finishes.

We use it primarily for **Lazy Initialization** (Singletons). For example, initializing a complex database connection pool only when the first request actually needs it. It’s cleaner and safer than manual `if initialized == false` checks, which are prone to race conditions.

---

## 73. How do you avoid race conditions?

**Answer:**
The best way is to **Avoid Shared State** entirely. "Share memory by communicating, don't communicate by sharing memory." Passing portions of data over channels is safer than locking a global variable.

If you must share state (like a cache), use `sync.Mutex` or `atomic` primitives to Serialize access.

Finally, effective tooling is key. We ALWAYS run tests with `go test -race`. The **Race Detector** instruments the code to catch race conditions at runtime. It has high overhead, so we don't run it in production, but it is mandatory in CI/CD.

---

## 74. What is the Go memory model?

**Answer:**
The Go Memory Model defines the strict rules of **Visibility**. It answers the question: "If I write to variable X in Goroutine A, when is Goroutine B guaranteed to see it?"

The answer is: **Never**, unless you use explicit synchronization. The compiler and CPU are free to reorder instructions for performance.

Synchronization primitives like Channels and Mutexes act as **Memory Barriers**. A channel send "Happens-Before" the receive completes, guaranteeing that any writes done before the send are visible to the receiver. Without these barriers, you are coding in "undefined behavior" land.

---

## 75. How do you use `context.Context` for cancellation?

**Answer:**
We pass `context.Context` as the **first argument** to every function in the call stack.

To cancel a tree of operations, we create a context with `ctx, cancel := context.WithCancel(parent)`. When we call `cancel()`, the `ctx.Done()` channel closes.

Functions listen to this channel in a `select`. If it closes, they abort their work immediately. This is standard in HTTP servers—if a user closes their browser tab, the `request.Context` is cancelled, propagating down to the database driver to kill the running SQL query and free up resources.

---

## 76. How to pass data between goroutines?

**Answer:**
We use **Channels**. This is the idiomatic "Go Way."

By sending a pointer over a channel, you effectively transfer ownership of that data to the receiving goroutine. The sender should stop touching it to avoid race conditions.

For high-throughput pipelines (like log processing), we might use buffered channels to allow the producer to run slightly ahead of the consumer. This decoupling allows the system to absorb small bursts of traffic without blocking the entire application.

---

## 77. What is the `runtime.GOMAXPROCS()` function?

**Answer:**
It limits the number of OS threads that can execute Go code simultaneously. By default, it equals the number of CPU cores.

We generally leave this alone. However, in containerized environments (Docker/Kubernetes) with CPU quotas, older versions of Go sometimes saw the physical host's 64 cores instead of the container's 2-core quota, leading to severe throttling.

In those cases, we used to set `GOMAXPROCS` manualy (or use `automaxprocs` lib). Setting it to `1` is also a useful debugging trick to serialize execution and reproduce certain race conditions deterministically.

---

## 78. How do you detect deadlocks in Go?
**Answer:**
For **Global Deadlocks** (where every single goroutine is asleep), the Go Runtime detects it automatically and panics with `fatal error: all goroutines are asleep`.

For **Partial Deadlocks** (where just a few threads are stuck dependent on each other), we use `pprof`.

We enable the **Block Profiler**. We can then inspect the `/debug/pprof/block` endpoint to see stack traces of goroutines that have been waiting on locks for an unusually long time. This helps us find the "cycle" (A waits for B, B waits for A) and fix the locking order.

---

## 79. What are worker pools and how do you implement them?

**Answer:**
A Worker Pool is a pattern to limit concurrency.

Instead of spawning a new goroutine for every single task (which can lead to unbounded memory growth), you start a fixed number of workers (e.g., 5). They all consume from a single shared channel.

This creates **Backpressure**. If the workers are busy, the channel fills up, and the producer eventually blocks. This prevents a massive spike in traffic from crashing your server with Out-Of-Memory errors. It effectively smoothes out the load.

---

## 80. How to write concurrent-safe data structures?

**Answer:**
The Golden Rule is **Encapsulation**.

You hide the mutable data (like a `map`) as a **private field** inside a struct. You protect it with a `sync.RWMutex`.

You then expose only public methods (`Get`, `Set`) that handle the locking internally. This prevents the caller from ever accessing the map without a lock. If you rely on the caller to "remember to lock," they will eventually forget, and your program will crash.
