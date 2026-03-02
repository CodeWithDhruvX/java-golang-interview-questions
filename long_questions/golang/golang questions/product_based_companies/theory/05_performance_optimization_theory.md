# 🗣️ Theory — Performance Optimization in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How do you profile a Go application?"

> *"Go has excellent built-in profiling via the `pprof` package. The easiest way is to add `import _ 'net/http/pprof'` to your server — this registers profiling endpoints under `/debug/pprof`. Then you can use `go tool pprof` to connect to a running service and capture a CPU profile, heap snapshot, or goroutine dump. For CPU profiling: `go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30` — this samples your program for 30 seconds and gives you a call graph. The `top` command shows where your CPU time is going. `web` opens a flame graph — the wider a box, the more CPU it consumed. For tests, `go test -cpuprofile cpu.out` saves a profile you can analyze later."*

---

## Q: "What is `sync.Pool` and when should you use it?"

> *"sync.Pool is a cache of allocated objects that can be reused across goroutines. It's designed for objects that are frequently created and discarded — like buffers, encoders, or request contexts. Instead of allocating and garbage collecting them repeatedly, you Get one from the pool, use it, and Put it back. This reduces the number of heap allocations in hot paths, which directly reduces GC pressure. Important caveats: Pool objects can be collected at any GC cycle — so never store state in a pooled object that needs to outlive the operation. Always reset the object before use. Common uses in production: `bytes.Buffer` pools in web servers, `json.Encoder` pools, database row scanner pools."*

---

## Q: "How do you reduce heap allocations in Go?"

> *"Allocations are the enemy of performance — every allocation is more work for the GC. The key strategies: first, pre-allocate slices with `make([]T, 0, expectedSize)` when you know roughly how many elements you'll add, to avoid repeated resizing. Second, use `strings.Builder` for string concatenation instead of `+=` — which creates a new string on every operation. Third, prefer value types over pointer types for small objects — returning `Point{x,y}` instead of `*Point` keeps it on the stack. Fourth, avoid interface boxing in hot loops — storing concrete values in interfaces causes heap allocation. And fifth, use sync.Pool for frequently recycled objects."*

---

## Q: "How do you detect goroutine leaks and why are they a problem?"

> *"A goroutine leak is when goroutines are started but never terminate — they block forever on channels, mutexes, or I/O. They don't consume CPU but they do consume memory — about 2-8KB each to start, possibly growing. With thousands of leaking goroutines, you'll run out of memory. Detection: `runtime.NumGoroutine()` can alert you if the count unexpectedly grows. In tests, the `goleak` library from Uber checks that goroutine counts returned to baseline after a test. In production, expose the goroutine count as a metric and alert on it. Prevention: always pass a context to goroutines so you can cancel them; always close channels when you're done so range-over-channel loops can exit; always have a case for `ctx.Done()` in select blocks."*

---

## Q: "What is the difference between `ns/op`, `B/op`, and `allocs/op` in benchmarks?"

> *"When you run `go test -bench=. -benchmem`, you get three metrics per benchmark. `ns/op` is nanoseconds per operation — the raw speed. Lower is better. `B/op` is bytes allocated per operation on the heap — how memory-hungry each call is. Lower means less GC pressure. `allocs/op` is the number of heap allocations per operation — how many times the allocator was called. Even if total bytes is small, many small allocations can cause GC overhead. The goal is usually to drive `allocs/op` toward zero for hot paths. If you see 0 allocs/op, the operation is entirely stack-allocated — ideal for high-frequency paths."*

---

## Q: "How does GOGC tuning work and when would you change it?"

> *"GOGC controls when garbage collection triggers. With the default of 100, the GC runs when the heap has grown 100% since the last collection — i.e., when it doubles. Lowering GOGC means more frequent collections — you use less memory but spend more CPU on GC. Raising it means less frequent collections — lower GC CPU overhead but higher memory usage. For a latency-sensitive service where you want to minimize GC pause frequency, you might raise GOGC and set a `GOMEMLIMIT` instead — let the heap grow larger between collections but cap total memory usage. For a memory-constrained environment, lower GOGC. In practice, most services run well with the defaults; only tune when you have clear profiling data showing GC is a bottleneck."*
