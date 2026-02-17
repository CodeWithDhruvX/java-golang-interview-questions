# 🟢 Go Theory Questions: 521–540 Performance Optimization

## 521. How do you benchmark Go code using `testing.B`?

**Answer:**
We write a function starting with `BenchmarkXxx(b *testing.B)`.
Core loop: `for i := 0; i < b.N; i++ { CodeToTest() }`.

`b.N` is dynamically adjusted by the tool until the test runs for enough time (default 1s) to get statistically significant results.
Run with: `go test -bench=. -benchmem`.
The `-benchmem` flag is vital—it shows **Allocations per Operation** (B/op and allocs/op), which are often the culprit for poor performance.

---

## 522. What tools can you use to profile a Go application?

**Answer:**
The standard tool is **pprof**.
It visualizes CPU, Heap, Goroutine, and Block profiles.

We add `import _ "net/http/pprof"` and start an HTTP server.
Then run: `go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30`.
This generates an interactive SVG graph (flame graph) showing which functions consume the most CPU cycles.
For continuous profiling in production, we use tools like **Datadog Continuous Profiler** or **Pyroscope**.

---

## 523. How does memory allocation affect Go performance?

**Answer:**
Allocating memory on the **Heap** is expensive.
1.  **Allocation**: The runtime must find a free span, potentially asking OS for RAM.
2.  **GC Pressure**: Every byte on the heap must be scanned by the Garbage Collector eventually.

High allocation rates = Frequent GC cycles = High CPU usage and latency spikes.
Optimizing Go performance is 80% about reducing heap allocations (Escape Analysis, Sync.Pool, Value Semantics).

---

## 524. How do you detect and fix memory leaks in Go?

**Answer:**
Go is garbage collected, so "leaks" are usually **References that never die**.
Common causes:
1.  **Goroutine Leaks**: Launching a goroutine that blocks forever on a nil channel or without a quit signal.
2.  **Global Maps**: Adding items to a global cache without eviction logic.

Detection:
Look at `go tool pprof -heap` over time. If the "In-Use" memory graph mimics a staircase (up only), you have a leak. Use `pprof -goroutine` to find 100,000 stuck goroutines.

---

## 525. How do you avoid unnecessary allocations in hot paths?

**Answer:**
1.  **Preallocate Slices**: `make([]int, 0, 1000)` prevents resizing/copying.
2.  **Reuse Buffers**: Use `sync.Pool` for `bytes.Buffer` or `[]byte`.
3.  **Avoid Interfaces**: Passing concrete types avoids interface wrapping (if compiler can't inline).
4.  **Zerolog approach**: Use libraries designed for zero-allocs (structs instead of maps, appending bytes instead of strings).

---

## 526. What is escape analysis and how does it impact performance?

**Answer:**
Escape Analysis is the compiler's decision: "Can this variable live on the Stack?"

**Stack**: Fast allocation (move pointer), auto-free on return.
**Heap**: Slow allocation, GC overhead.

We inspect it: `go build -gcflags="-m"`.
If a variable "Escapes to Heap", it costs performance. We optimize code to prevent escape (e.g., passing pointers down the stack is fine, passing pointers *up* the stack forces escape).

---

## 527. How do you use `pprof` to trace CPU usage?

**Answer:**
We capture a profile: `curl -o cpu.prof .../profile`.
We open it: `go tool pprof -http=:8080 cpu.prof`.

We look at the **Flame Graph**.
X-axis: Population (how often it's on CPU).
Y-axis: Stack Depth.
The "Wide" bars are the functions consuming the most time. If `runtime.mallocgc` is wide, we have an allocation problem. If `syscall.Read` is wide, we are I/O bound.

---

## 528. How do you optimize slice operations for speed?

**Answer:**
1.  **Pre-allocate**: `make([]T, 0, capacity)`.
2.  **filtering**: Instead of `new := []int{}`, iterate and swap elements in-place or maintain an index `n` to rewrite the slice over itself (`a[:n]`).
3.  **Copy**: Use `copy()` builtin, which compiles to `memmove` (very fast assembly).
4.  **Pointer Slices**: If structs are large (e.g., 200 bytes), use `[]*Struct` to avoid moving massive amounts of data during sorting/growing.

---

## 529. What is object pooling and how is it implemented in Go?

**Answer:**
Object Pooling recycles allocated objects so we don't trash the GC.
Implementation: `sync.Pool`.

`Get()` retrieves an object (or creates new if empty).
`Put()` returns it.
**Warning**: `sync.Pool` can be drained by the GC at any time. Do not store persistent state (like database connections) in it. Use it for ephemeral buffers, gzip writers, or temporary request context objects.

---

## 530. How does GC tuning affect latency in Go services?

**Answer:**
The main knob is `GOGC` (default 100).
It controls the ratio of new garbage to live data before triggering GC.
`GOGC=100`: GC runs when heap grows 100% (doubles).
`GOGC=200`: GC runs less often (more RAM usage, better throughput).

Go 1.19+ added `GOMEMLIMIT`. We set this to 90% of container memory. The GC will relax `GOGC` aggressively until it approaches the limit, preventing OOM kills while minimizing GC cycles.

---

## 531. How do you measure and reduce goroutine contention?

**Answer:**
Contention happens when many goroutines fight for a single Lock.
Measure: `go build -race` (detects races) or `pprof -mutex`.

Wait times in the mutex profile show contention.
Fix:
1.  **Sharding**: Split one map with one mutex into 32 maps with 32 mutexes (`concurrent-map`).
2.  **Channels**: Use channels to serialize access (Actor pattern).
3.  **RWMutex**: If read-heavy, allow concurrent readers.

---

## 532. What is lock contention and how to identify it in Go?

**Answer:**
Lock Contention is when CPU time is wasted **Waiting** for a lock, not working.
We identify it via the **Block Profile**.
`runtime.SetBlockProfileRate(1)`.
In pprof: `go tool pprof block.prof`.
It shows exactly where goroutines are parking.
If `GenericLock` is high, replace the Mutex with `atomic` operations (CAS) if possible, or redesign the data flow to be lock-free.

---

## 533. How do you batch DB operations for better throughput?

**Answer:**
Round-trips (Network Latency) kill throughput.
Instead of 1,000 inserts (1,000ms latency total), do 1 batch insert (10ms).

In Go:
Use a **Buffered Channel**.
Workers read from channel, fill a slice `batch := make([]Item, 0, 100)`.
When `len(batch) == 100` OR `ticker.C` fires (e.g., every 500ms), flush the batch to DB.
This creates a hybrid Size/Time trigger, ensuring low latency for low traffic and high throughput for high traffic.

---

## 534. How would you profile goroutine leaks?

**Answer:**
We check `runtime.NumGoroutine()`. If it grows linearly without shrinking, it's a leak.
To pinpoint: use `pprof`.
Click "Goroutines". See the stack traces.
If you see 5,000 goroutines stuck at `net.http.ReadRequest`, you are leaking connections.
If stuck at `channel receive`, you have orphaned workers.

---

## 535. What are the downsides of excessive goroutines?

**Answer:**
Goroutines are cheap (2KB), but not free.
1.  **Memory**: 1 million goroutines = 2GB RAM minimum (stacks).
2.  **Scheduler Load**: The scheduler has to manage 1M entities. Scanning them (or their stacks) during GC takes CPU.
3.  **Panic Risk**: If you launch a goroutine per HTTP request without a limit (Worker Pool), a spike can OOM the server. always bound concurrency.

---

## 536. How would you measure and fix cold starts in Go Lambdas?

**Answer:**
Go is great for cold starts (~200ms vs Java's 2s).
Measure: AWS X-Ray or logs.
Fix:
1.  **Binary Size**: Use `-ldflags="-s -w"`, strip symbols. Smaller binary = faster load from S3.
2.  **Init Code**: Avoid heavy lifting (DB connection, loading big config) in `init()`. Do it lazily on the first request if possible, or use **Provisioned Concurrency**.

---

## 537. How do you decide between a map vs slice for performance?

**Answer:**
**Small N (< 20-50)**: Slice is faster. Linear scan (O(N)) is better than Hashing (O(1)) because of CPU Cache Lokality and no overhead.
**Large N**: Map is faster.
**Sparse Data**: Map.
**Dense Data**: Slice.
Always benchmark. Iterating a map is also random and slower than iterating a slice. If you mostly Read/Iterate, use Slice. If you mostly Lookup by ID, use Map.

---

## 538. How would you write a memory-efficient parser in Go?

**Answer:**
We **Stream**, never load the whole file.
JSON: `json.Decoder` (Stream) vs `json.Unmarshal` (RAM).

For extreme efficiency (e.g., Parsing 10GB XML):
Create a state machine reading byte-by-byte (`bufio.Scanner`).
Reuse the same `[]byte` buffer for the current token.
Once processed, overwrite the buffer. This keeps RAM usage constant (e.g., 4KB) regardless of the file size (1TB).

---

## 539. How do you use channels efficiently under heavy load?

**Answer:**
1.  **Buffered Channels**: Provide "give" so sender doesn't block immediately on receiver latency.
2.  **Batching**: Send `[]T` instead of `T` over the channel to reduce lock contention on the channel itself.
3.  **Closer**: Ensure producers close the channel so consumers terminate.
**Anti-Pattern**: Using channels for *everything*. They are slower than Mutexes. For simple state protection, use Mutex. Use Channels for orchestration/signaling.

---

## 540. When should you use sync.Pool?

**Answer:**
Use `sync.Pool` when:
1.  You have **Short-Lived** objects (Request Scope).
2.  You allocate **Many** of them (High Frequency).
3.  The objects are **expensive** to allocate (Large Buffers, complex Structs).

Do not use it for:
1.  Long-lived objects.
2.  Tiny objects (int, small struct) — the locking overhead of the Pool is higher than the allocation cost.
3.  Objects that must be "Clean" (you must manually Reset the object before reuse, or you leak data between users).
