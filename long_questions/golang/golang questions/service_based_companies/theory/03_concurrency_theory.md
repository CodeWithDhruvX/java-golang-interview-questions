# 🗣️ Theory — Concurrency in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What are goroutines and how are they different from threads?"

> *"Goroutines are Go's lightweight unit of concurrency. You start one by just putting the `go` keyword in front of a function call — `go doWork()`. What makes goroutines special compared to OS threads is their cost. An OS thread typically uses 1 to 8 megabytes of stack and takes significant time to create and context-switch. A goroutine starts at around 2 kilobytes and grows as needed, dynamically. The Go runtime manages thousands or even millions of goroutines, multiplexing them onto a small number of OS threads. So goroutines are much cheaper, faster to create, and enable a programming model where 'one goroutine per request' is actually feasible."*

---

## Q: "What are channels in Go? What's the difference between buffered and unbuffered?"

> *"Channels are Go's mechanism for goroutines to communicate — they follow the philosophy 'don't communicate by sharing memory, share memory by communicating'. An unbuffered channel — `make(chan int)` — is synchronous. When one goroutine sends, it blocks until another goroutine receives, and vice versa. It acts as a rendezvous point. A buffered channel — `make(chan int, 5)` — has a queue. The sender only blocks when the buffer is full, and the receiver only blocks when the buffer is empty. Buffered channels are useful for decoupling producer and consumer speeds, or for rate limiting work."*

---

## Q: "What is a WaitGroup and when do you use it?"

> *"A `sync.WaitGroup` is like a counter that lets your main goroutine wait for a collection of goroutines to finish. You call `wg.Add(1)` before launching each goroutine, `wg.Done()` inside each goroutine when it finishes — typically with a defer — and `wg.Wait()` in the main goroutine to block until the counter reaches zero. The common mistake with WaitGroups is calling `wg.Add(1)` inside the goroutine itself — that creates a race condition because the goroutine might not even have started by the time you call `wg.Wait()`. Always add before launching."*

---

## Q: "How does a Mutex work? When would you use RWMutex instead?"

> *"A Mutex — mutual exclusion lock — prevents multiple goroutines from accessing the same data simultaneously. You call `Lock()` before accessing shared data and `Unlock()` after — always in a defer so you can't forget it. An `RWMutex` is a read-write mutex that allows multiple concurrent readers or one exclusive writer. Use `RLock` and `RUnlock` for reads, `Lock` and `Unlock` for writes. If your shared state is read far more often than it's written — like a config or a cache — RWMutex gives you better throughput because readers don't block each other."*

---

## Q: "What is the `select` statement? How is it different from a switch?"

> *"The `select` statement is like a `switch` but for channels. It waits on multiple channel operations and proceeds with whichever one is ready first. If multiple are ready at the same time, it picks one randomly — which is by design to avoid starvation. If none are ready and there's a `default` case, it executes that without blocking — making it non-blocking. Without a default, `select` blocks until one channel is ready. The most common patterns: using `select` with `time.After()` for timeouts, and using it with `ctx.Done()` for graceful cancellation."*

---

## Q: "How do you close a channel properly? What are the rules?"

> *"There are strict rules around closing channels. First, only the sender should close a channel — never the receiver, because the sender knows when there's no more data. Second, never close a channel more than once — that panics. Third, never send to a closed channel — also panics. The idiomatic pattern is: the producer goroutine closes the channel after sending all values, and the consumer uses `for v := range ch` which automatically stops when the channel is closed. If you need to check manually, you use the two-value receive: `v, ok := <-ch` where `ok` is false when the channel is closed and empty."*

---

## Q: "What is `sync.Once` used for?"

> *"sync.Once guarantees that a function executes exactly once, no matter how many goroutines call it concurrently. It's the idiomatic Go way to implement a thread-safe singleton. The classic use case is lazy initialization — you want to create a database connection pool or load a config file exactly once, the first time it's needed. Inside `once.Do(func() { ... })`, the function runs exactly once. Subsequent calls to `once.Do()` are no-ops, even from different goroutines. It's simpler and safer than trying to implement double-checked locking yourself."*

---

## Q: "What is `context.Context` and why is it important?"

> *"Context is Go's standard way to propagate cancellation, deadlines, and request-scoped values through a call chain. Think about an HTTP request — it comes in, triggers database queries, calls other services. If the client disconnects, you want everything in that chain to cancel cleanly. That's what context does. You pass `ctx context.Context` as the first argument to every function that might need cancellation. You check `ctx.Done()` for cancellation signals. You create child contexts with `WithTimeout` or `WithCancel`. It's considered best practice to never store a context in a struct — always pass it as a function parameter."*

---

## Q: "What is a worker pool and why do you need one?"

> *"Launching a goroutine for every single task sounds great since goroutines are cheap — but if you have 100,000 tasks, you'd create 100,000 goroutines, which is wasteful and can overwhelm your resources. A worker pool limits concurrency: you create a fixed number of goroutines — say, 10 workers — and they all read from a shared jobs channel. As tasks come in, available workers pick them up. This gives you controlled, bounded parallelism. It's one of the most common concurrency patterns in production Go code, especially for processing queues, making HTTP requests in bulk, or doing database operations concurrently."*
