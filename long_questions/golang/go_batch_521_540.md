## 🚀 Performance Optimization (Questions 521-540)

### Question 521: How do you benchmark Go code using `testing.B`?

**Answer:**
Benchmarks function start with `Benchmark` in `_test.go` files and take `*testing.B`.
You must execute the target code `b.N` times.

```go
func BenchmarkConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = fmt.Sprintf("str%d", i)
    }
}
```
Run: `go test -bench=. -benchmem`

---

### Question 522: What tools can you use to profile a Go application?

**Answer:**
1.  **pprof:** The standard built-in profiler (CPU, Heap, Goroutine, Block, Mutex).
2.  **trace:** `go tool trace`. Visualizes the runtime behavior (scheduler, garbage collector, goroutine execution) over time.
3.  **fgprof:** Captures both On-CPU and Off-CPU (I/O waiting) time.

---

### Question 523: How does memory allocation affect Go performance?

**Answer:**
- **Stack Allocation:** Extremely fast. Freed automatically when function returns.
- **Heap Allocation:** Slower. Requires Garbage Collector (GC) to track and free.
- **Impact:** High heap allocation rate -> Frequent GC cycles -> Higher Latency/CPU usage.
**Goal:** Keep short-lived variables on the Stack.

---

### Question 524: How do you detect and fix memory leaks in Go?

**Answer:**
Go has a GC, but leaks occur if references are kept unintentionally.
**Detection:**
1.  Capture Heap Profile (`go tool pprof http://localhost:8080/debug/pprof/heap`).
2.  Use `--base` flag to compare two profiles over time:
    `go tool pprof --base profile1.pb.gz profile2.pb.gz`
3.  Look for `inuse_space` growing continuously.

**Common Causes:** Unstopped `time.Ticker`, Goroutines stuck waiting on nil channels, Global maps growing indefinitely.

---

### Question 525: How do you avoid unnecessary allocations in hot paths?

**Answer:**
1.  **Preallocate:** Use `make([]T, 0, capacity)` to avoid slice resizing.
2.  **Reuse Buffers:** Use `sync.Pool` or simply pass a pre-allocated buffer to functions.
3.  **Strings:** Use `strings.Builder` instead of `+`.
4.  **Zero-Allocation Libraries:** Use `fasthttp` or `zerolog` principles (avoiding boxing/interface{} conversion).

---

### Question 526: What is escape analysis and how does it impact performance?

**Answer:**
Escape Analysis is the compiler phase that decides if a variable can live on the **Stack** or must "escape" to the **Heap**.
- If a pointer to a variable is returned from a function, it *escapes* to Heap.
- If a variable is passed to `fmt.Println` (takes interface{}), it often escapes.
**Check:** `go build -gcflags="-m"` to see optimization decisions.

---

### Question 527: How do you use pprof to trace CPU usage?

**Answer:**
1.  **Expose:** Import `net/http/pprof` and run a web server.
2.  **Capture:** `go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30`
3.  **Analyze:**
    - `top`: Text list of function consuming most CPU.
    - `web`: Generates an SVG graph.
    - `list CheckSum`: Shows assembly/source usage for a specific function.

---

### Question 528: How do you optimize slice operations for speed?

**Answer:**
1.  **Capacity:** allocate capacity upfront (`make([]int, 0, 1000)`).
2.  **Filter without Alloc:** Use in-place filtering.
    ```go
    // Filter odd numbers
    n := 0
    for _, x := range data {
        if x%2 == 0 {
            data[n] = x
            n++
        }
    }
    data = data[:n]
    ```
3.  **Copy:** Use `copy()` builtin, it uses optimized `memmove`.

---

### Question 529: What is object pooling and how is it implemented in Go?

**Answer:**
Object pooling reuses heavy objects instead of allocating/deallocating them repeatedly to reduce GC pressure.
**Implementation:** `sync.Pool`.

```go
var bufPool = sync.Pool{
    New: func() interface{} { return new(bytes.Buffer) },
}

func Log() {
    b := bufPool.Get().(*bytes.Buffer)
    b.Reset()
    defer bufPool.Put(b)
    
    b.WriteString("Log entry...")
}
```

---

### Question 530: How does GC tuning affect latency in Go services?

**Answer:**
Go 1.19+ introduced `GOMEMLIMIT`.
- **GOGC (Default 100):** GC triggers when heap grows by 100%. Lowering it triggers GC more often (lower RAM, higher CPU).
- **GOMEMLIMIT:** Sets a soft memory cap. The GC becomes aggressive as it approaches this limit, preventing Out-Of-Memory (OOM) crashes while staying lazy when memory is abundant.

---

### Question 531: How do you measure and reduce goroutine contention?

**Answer:**
**Measure:**
1.  `go test -bench=. -cpuprofile=cpu.out`
2.  `go tool pprof cpu.out` -> Look for `runtime.futex` or `syscall` overhead.
3.  Enable Mutex Profiling: `runtime.SetMutexProfileFraction(1)`.

**Reduce:**
- Shard locks (break one big mutex into many smaller ones).
- Use Channels instead of Mutexes (sometimes, not always faster).
- Use `sync/atomic` for counters.

---

### Question 532: What is lock contention and how to identify it in Go?

**Answer:**
It happens when many goroutines try to acquire the same `sync.Mutex` simultaneously, causing them to sleep/context switch waiting for the lock.
**Identify:**
- Use the **Block Profiler** or **Mutex Profiler**.
- Access `debug/pprof/mutex` or `debug/pprof/block` endpoints.
- High `Lock` time relative to CPU time indicates contention.

---

### Question 533: How do you batch DB operations for better throughput?

**Answer:**
Database Round Tritt (RTT) is expensive.
**Batching:** Combine 1000 INSERTs into one SQL statement.
```sql
INSERT INTO table (val) VALUES (1), (2), (3)...;
```
Or use `pgx.Batch` (Postgres driver) to queue multiple queries and send them in one network packet.
This dramatically increases Write throughput.

---

### Question 534: How would you profile goroutine leaks?

**Answer:**
1.  Check the total count: `runtime.NumGoroutine()`.
2.  If rising, grab a profile: `http://localhost:6060/debug/pprof/goroutine?debug=2`.
3.  This dumps the stack trace of **every** currently running goroutine.
4.  Look for large clumps of goroutines stuck at the same line (e.g., waiting on a channel read `chan receive`).

---

### Question 535: What are the downsides of excessive goroutines?

**Answer:**
Although lightweight (2KB stack), they are not free.
1.  **Memory:** 1,000,000 goroutines ≈ 2GB-4GB RAM minimum.
2.  **Scheduler:** The runtime scheduler has to manage them. Too many runnable goroutines causing execution to trash (context switch storms).
3.  **Panic Risk:** If you spawn a goroutine per HTTP request without a limit, a traffic spike can OOM the server. (Solution: Worker Pool / Semaphore).

---

### Question 536: How would you measure and fix cold starts in Go Lambdas?

**Answer:**
**Measure:** AWS X-Ray or looking at `Init Duration` in logs.
**Fix:**
1.  **Binary Size:** Strip DWARF (`-s -w`), avoid large dependencies.
2.  **Init Logic:** Move heavy initialization from `var _ = init()` or global scope to `lazy` loading (initialize only on first request).
3.  **Chain:** Avoid `init()` chains that perform network calls.

---

### Question 537: How do you decide between a map vs slice for performance?

**Answer:**
- **Small N (<= 20-50):** Slice linear scan is often faster than Map lookup because of CPU cache locality and zero hashing overhead.
- **Large N:** Map is O(1), Slice is O(N). Map wins.
- **Micro-Optimization:** Avoid maps for composite keys (structs) if possible as hashing is slower.

---

### Question 538: How would you write a memory-efficient parser in Go?

**Answer:**
1.  **Streaming:** Use `bufio.Scanner` or `json.Decoder` to process stream chunk-by-chunk instead of `ioutil.ReadAll`.
2.  **Zero-Copy:** Use `[]byte` and slice indexing rather than converting to `string`.
3.  **Reuse:** Recycle structs or buffers as you parse lines.

---

### Question 539: How do you use channels efficiently under heavy load?

**Answer:**
1.  **Buffered Channels:** Use buffers to decouple producer/consumer bursts.
2.  **Batching:** Send `[]Item` (batch) over channel instead of single `Item` to reduce locking/context-switching overhead per item.
3.  **Close:** Dont close channels unless necessary to signal completion. Closing is expensive and requires synchronization.

---

### Question 540: When should you use `sync.Pool`?

**Answer:**
Only when:
1.  You have a **Hot Path**.
2.  You verify via profiling that **Garbage Collection (GC)** is a bottleneck.
3.  The objects are large/complex to allocate.
**Warning:** `sync.Pool` is cleared on every GC cycle. Do not use it for long-term storage or database connections.

---
