# 🗣️ Theory — Advanced Concurrency & Goroutines
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "Can you explain the GMP model in Go?"

> *"GMP stands for Goroutine, Machine, and Processor — the three entities in Go's runtime scheduler. A Goroutine is the unit of work — lightweight, starting at 2KB of stack. An M is an OS thread that actually runs code — Go typically has as many Ms as CPU cores. A P is a logical processor that holds a queue of goroutines and is assigned to an M. The scheduler is the glue: each P has a local run queue of goroutines. An M needs a P to run goroutines. When a goroutine blocks on I/O or a syscall, the P detaches and finds another M — or creates one — so the CPU stays busy. This is how Go runs millions of goroutines on a handful of OS threads."*

---

## Q: "What is a pipeline in Go? How do you implement one?"

> *"A pipeline is a composition of concurrent processing stages connected by channels. Each stage is a goroutine that reads from an input channel, transforms the data, and sends to an output channel. The first stage generates data; the last stage consumes it. The middle stages transform it. Because each stage is independent and channels act as the connectors, stages can run concurrently — while stage 2 is processing item 1, stage 3 is processing item 0. It's the channel-based equivalent of Unix pipes. For production pipelines, you need to handle cancellation — pass a context and check its Done channel in each stage."*

---

## Q: "Explain fan-out and fan-in patterns."

> *"Fan-out means distributing work from one channel to multiple goroutines — like forking. One producer puts work into a channel and N workers all read from it. This parallelizes processing. Fan-in is the inverse — merging results from multiple goroutines into a single channel. You'd have N worker goroutines each writing to their own output channel, and a fan-in goroutine that reads from all of them and writes to one merged channel. Together they form a scatter-gather pattern: fan-out to scatter work to parallel workers, fan-in to gather results. Often paired with sync.WaitGroup to know when all workers are done."*

---

## Q: "What is a semaphore in Go and how do you implement one?"

> *"A semaphore limits the number of goroutines that can access a resource concurrently. In Go, the idiomatic implementation is a buffered channel. You create a channel with capacity equal to the limit — `sem := make(chan struct{}, 3)`. Acquiring the semaphore means sending to it — `sem <- struct{}{}` — which blocks when full. Releasing means receiving — `<-sem`. This is elegant because the buffered channel's blocking behavior is exactly semaphore semantics. Common use: limiting concurrent HTTP requests, limiting concurrent database connections in a batch job, or rate-limiting access to an external API."*

---

## Q: "What is `errgroup` and how is it better than just using WaitGroup?"

> *"errgroup from `golang.org/x/sync` is WaitGroup plus error propagation. The problem with raw WaitGroup is: when a goroutine fails, how do you get the error back to the coordinator? You'd need an error channel or a mutex-protected slice. errgroup handles this cleanly. You call `g.Go(func() error { ... })` for each goroutine. `g.Wait()` blocks until all complete and returns the first non-nil error. If any goroutine fails, the context is automatically cancelled — all goroutines that check `ctx.Done()` will see the cancellation signal and stop, preventing wasted work after a failure."*

---

## Q: "How do you identify and fix deadlocks in Go?"

> *"A deadlock is when all goroutines are blocked, waiting for each other in a cycle — nobody can proceed. Go's runtime actually detects this and panics with 'all goroutines are asleep — deadlock!'. For subtler deadlocks — like a goroutine that's undetected by the runtime — you use the race detector `go run -race main.go` and pprof goroutine dumps. The common causes: acquiring two mutexes in different orders from different goroutines; sending to an unbuffered channel with no receiver; a channel that's never closed so a range-over-channel never exits. Prevention: consistent lock ordering, always having a receiver for every send, always closing channels when done."*

---

## Q: "What is the difference between `sync.Mutex` and `sync.RWMutex`? When do you choose which?"

> *"A regular Mutex is an exclusive lock — only one goroutine can hold it at a time, whether reading or writing. An RWMutex is a readers-writer lock — multiple readers can hold `RLock` simultaneously, but a writer needs an exclusive `Lock` that blocks all readers and other writers. Choose RWMutex when: reads are much more frequent than writes, the read operation takes non-trivial time, and you have significant concurrency. A cache is the classic example. Choose plain Mutex when operations are very fast, or when reads and writes are roughly equal — because RWMutex has higher overhead and the advantage only shows up when reads truly dominate."*
