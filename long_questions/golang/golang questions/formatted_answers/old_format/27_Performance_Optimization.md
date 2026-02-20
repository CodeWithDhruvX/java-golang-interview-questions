# 🚀 **521–540: Performance Optimization**

### 521. How do you benchmark Go code using `testing.B`?
"I write a function `BenchmarkMyFunc(b *testing.B)`.
I wrap the logic in a loop: `for i := 0; i < b.N; i++ { MyFunc() }`.
`b.N` matches the runtime's need to get a statistically significant result.
I run it with `go test -bench=. -benchmem`. The `benchmem` flag is crucial—it shows allocations per operation, which is often the first thing I optimize."

#### Indepth
Compiler Optimizations: Go 1.20+ introduced **PGO (Profile-Guided Optimization)**. You build your app, run it in production to generate a `default.pgo` profile, and then `go build -pgo=default.pgo`. The compiler uses this real-world data to inline hot functions more aggressively, boosting performance by 5-10% for free.

---

### 522. What tools can you use to profile a Go application?
"**pprof** is the standard.
I import `_ "net/http/pprof"` and hit `/debug/pprof/`.
It gives me:
*   **CPU Profile**: Where time is spent.
*   **Heap Profile**: Who is allocating memory.
*   **Goroutine Profile**: Detailed stack traces of all goroutines.
*   **Trace**: Visual timeline of the scheduler/GC pauses."

#### Indepth
**Flame Graphs** are the modern way to visualize pprof data. Use `go tool pprof -http=:8080 cpu.prof`. The "Icicle Graph" view shows the call stack width proportional to CPU time. It makes it instantly obvious if 50% of your time is spent in `json.Marshal`.

---

### 523. How does memory allocation affect Go performance?
"Allocating on the **Heap** is expensive because it involves the Garbage Collector.
Allocating on the **Stack** is free (just moving a pointer).
If a variable 'escapes' (e.g., I return a pointer to it), it goes to the Heap.
My goal is 0 allocs in hot paths. I check this with `go build -gcflags='-m'`. High heap churn means frequent GC cycles and latency spikes."

#### Indepth
The **Stack** is growable (starts at 2KB). If a goroutine recurses too deep, the runtime allocates a larger stack and copies data over. This "stack copying" is cheap but not free. The **Heap**, however, needs the Sweeper and Scavenger to reclaim memory, which steals CPU cycles from your app.

---

### 524. How do you detect and fix memory leaks in Go?
"A leak is memory that grows indefinitely.
In Go, it's usually:
1.  **Goroutine Leak**: A goroutine blocked forever on a nil channel.
2.  **Global Map**: Growing a map without deleting keys.
I use `pprof` to confirm `heap_inuse` is rising over days, and `pprof -diff_base` to spot the difference between two snapshots."

#### Indepth
Use `curl http://host/debug/pprof/heap > heap.out`. Then `go tool pprof -alloc_space heap.out`. The `-alloc_space` flag shows *total bytes allocated* since startup (even if freed), which finds methods creating "garbage" pressure. `-inuse_space` only shows what's currently held (leaks).

---

### 525. How do you avoid unnecessary allocations in hot paths?
"I reuse memory.
Instead of `append([]byte, data...)`, I use a pre-allocated buffer from a `sync.Pool`.
Instead of `fmt.Sprintf` (which allocates strings), I use `strconv.AppendInt` to write directly to a byte slice.
I minimize interface conversions, as they often force heap allocation."

#### Indepth
**Interface conversion** (`func(v any)`) often causes an allocation because the runtime must put the concrete value into a box (interface header). Go 1.18+ Generics often eliminate this by monomorphizing the function for the specific type at compile time.

---

### 526. What is escape analysis and how does it impact performance?
"It’s the compiler's decision: **Stack or Heap?**
If I pass a pointer to `fmt.Println`, it escapes to Heap because `fmt` takes `interface{}`.
If I pass a pointer to a local helper function that gets inlined, it stays on Stack.
Understanding this lets me write code that the compiler allows on the stack, reducing GC pressure."

#### Indepth
You can see exactly what escapes with `go build -gcflags="-m -m"`. The output is verbose but tells you *why* something escaped (e.g., "parameter x leaks to ~r0"). Sometimes a simple reordering of code (like not passing a pointer to a closure) can save millions of allocations.

---

### 527. How do you use `pprof` to trace CPU usage?
"I capture a profile:
`curl -o cpu.prof http://localhost:6060/debug/pprof/profile?seconds=30`.
Then `go tool pprof cpu.prof`.
I use `top` to see the function consuming the most cycles.
Usually, it’s **serialization** (JSON/Protobuf) or excessive map lookups. I focus 80% of my optimization effort there."

#### Indepth
**Syscalls** are expensive. If `pprof` shows a lot of time in `syscall.Read` or `syscall.Write`, you are doing too many small I/O operations. Wrap your `net.Conn` or `os.File` in a `bufio.Reader/Writer`. This aggregates small IOs into 4KB chunks, drastically reducing context switches.

---

### 528. How do you optimize slice operations for speed?
"**Pre-allocate**.
`make([]T, 0, 1000)` prevents 10 reallocations/copies as I append.
**Copy vs Re-slice**:
`b = a[:2]` is fast (same backing array).
`copy(b, a)` is slower but safer if I want the original huge array to be GC'd.
I avoid `append` inside strict loops where simple index assignment `s[i] = val` covers it."

#### Indepth
Slice growth strategy: Go doubles the capacity until 1024 elements, then grows by ~25%. If you know the size, `make([]T, len)` is always faster. Also, be careful of **Memory Leaks in Slices**: `small := huge[:2]` keeps the entire `huge` array in memory. Use `copy` to detach the small part.

---

### 529. What is object pooling and how is it implemented in Go?
"It reuses objects to reduce GC pressure.
`var p = sync.Pool{ New: func() any { return new(Buffer) } }`.
`buf := p.Get().(*Buffer)`.
`buf.Reset()`.
`p.Put(buf)`.
It’s not a cache (GC can clear it anytime). It’s purely for reducing the allocation rate of short-lived, heavy objects like buffers."

#### Indepth
`sync.Pool` is local to each **P** (Processor). This means accessing it is mostly lock-free (stealing from other P's involves a lock). Don't put "expensive to create" database connections in a `sync.Pool` (use a real pool for that). `sync.Pool` is specifically for memory buffers that are cheap to zero-out but expensive to GC.

---

### 530. How does GC tuning affect latency in Go services?
"By default, `GOGC=100`. The GC runs when the heap doubles.
If I have 64GB of RAM, I set `GOGC=200` to run GC half as often.
Even better, I use `GOMEMLIMIT=10GiB` (Go 1.19+).
This tells the GC: 'Be lazy until we hit 10GB, then go hard.' It prevents OOM kills while maximizing throughput."

#### Indepth
`GOGC` is a tradeoff. `GOGC=off` disables GC entirely (dangerous, but useful for short CLI scripts). The "Soft Memory Limit" (`GOMEMLIMIT`) is a game changer for Kubernetes. It allows you to use all available RAM (efficiency) without crashing, as the GC will panic only as a last resort if it can't reclaim enough space.

---

### 531. How do you measure and reduce goroutine contention?
"Contention happens when many goroutines fight for the same Mutex.
I use `go test -bench=. -mutexprofile=mutex.out`.
Visualizing this shows 'Time spent waiting for lock'.
Fixes:
1.  **Sharded Locks**: Split one map into 32 maps/locks.
2.  **Channels**: Serialize access via a single owner goroutine instead of locking."

#### Indepth
**Mutex Spinning**: Before putting a goroutine to sleep (OS context switch), the runtime spins the CPU for a few nanoseconds hoping the lock becomes free. This burns CPU but reduces latency. If you see high CPU but low system load, it might be mutex spinning (contention).

---

### 532. What is lock contention and how to identify it in Go?
"It means CPU is idle because threads are blocked on `Mutex.Lock`.
I identify it via **Block Profile**.
`http://localhost:6060/debug/pprof/block`.
If I see `sync.(*Mutex).Lock` taking 50% of time, I have a hot lock. Adding more CPU cores won't help; it will make it worse."

#### Indepth
The **Block Profiler** is disabled by default because it has overhead. Enable it with `runtime.SetBlockProfileRate(1)` (captures every blocking event) for debugging, but turn it off or sample heavily (`Rate(10000)`) in production.

---

### 533. How do you batch DB operations for better throughput?
"Round-trips kill performance.
Instead of `INSERT` 1000 times (1000 network calls), I buffer them.
`INSERT INTO users VALUES (...), (...), ...`.
I use a Channel + Ticker. When the channel has 100 items or 1s passes, I flush the batch. This increases throughput by 10-50x."

#### Indepth
Latency vs Throughput. Batching improves throughput but harms latency (the first item waits for the batch to fill). Use a **hybrid trigger**: "Flush if 100 items OR 50ms passed". This ensures that even low-traffic periods don't suffer from high latency.

---

### 534. How would you profile goroutine leaks?
"If `runtime.NumGoroutine()` keeps rising, I'm leaking.
I check `http://localhost:6060/debug/pprof/goroutine?debug=2`.
It dumps stack traces of *all* goroutines.
I search for blocks. If 10,000 goroutines are stuck at `line 50` waiting on a channel that nobody writes to, that's my bug."

#### Indepth
Use `goleak` (by Uber) in your unit tests. `defer goleak.VerifyNone(t)`. It checks if any new goroutines were spawned during the test and not cleaned up. This catches leaks at the PR stage before they hit production.

---

### 535. What are the downsides of excessive goroutines?
"Goroutines are cheap (2KB) but not free.
1M goroutines = 2GB RAM just for stacks.
But **Scheduler** cost is higher:
Scanning 1M stacks during GC is slow.
Scheduling 1M ready-to-run goroutines causes cache thrashing.
I always use a worker pool to cap concurrency (e.g., 10k active jobs)."

#### Indepth
Context Switching. If you have 2 CPU cores and 1000 active (running, not waiting) goroutines, the OS/Scheduler wastes time switching between them. The ideal number of **active** threads is `runtime.GOMAXPROCS` (usually equal to CPU cores). Everything else should be waiting (IO) or queued.

---

### 536. How would you measure and fix cold starts in Go Lambdas?
"Go is fast, but `init()` functions run sequentially.
I measure locally using `time` or AWS X-Ray.
Fixes:
1.  Remove huge dependencies (AWS SDK v2 is modular; imports only what I need).
2.  Avoid network calls in `init()`.
3.  Use standard `net/http` over heavy frameworks like Fiber if <10ms startup is needed."

#### Indepth
**init()** functions are the silent killer of startup time. They run sequentially on a single thread. Avoid complex logic (DB connections, S3 fetches) in `init()`. Do them lazily or in `main()`. Use `GODEBUG=gctrace=1` to see if GC runs during startup.

---

### 537. How do you decide between a map vs slice for performance?
"**Slice**: O(N) lookup. Better cache locality. Fast for small N (< 20).
**Map**: O(1) lookup. Hashing overhead + pointer chasing.
For a list of 5 items, a linear scan of a slice is effectively faster than a map lookup due to CPU caching and lack of hashing cost."

#### Indepth
Map collision/hashing logic is complex (`SwissTable`). For small integers or enums, consider using a `[]T` as a lookup table (index = enum value). It's simpler and much faster. Only use maps when the key space is sparse or non-integer.

---

### 538. How would you write a memory-efficient parser in Go?
"I avoid `string` allocations.
I create a `Scanner` that uses indices on the underlying `[]byte` source.
I emit tokens as `[]byte` slices of the original buffer (zero-copy).
`simdjson-go` processes JSON at GB/s speeds by using this technique plus SIMD instructions."

#### Indepth
"Zero-Copy" parsing is dangerous if not handled: the `[]byte` result points to the huge original buffer, keeping it alive. If you only need a small chunk for long-term storage, copy it `string(b)`. If it's ephemeral (request duration), zero-copy is the way to go.

---

### 539. How do you use channels efficiently under heavy load?
"Unbuffered channels cause synchronization (blocking) for every message.
I use **buffered channels** to decouple producer/consumer speeds.
But I don't set the buffer too high (10k), as it hides backpressure and consumes RAM.
I avoid `select` where possible—it involves complex locking logic. A simple `range channel` on a single consumer is faster."

#### Indepth
**Channel Latency**: Sending on a channel involves a lock and potentially a scheduler call. For extremely high-performance (millions of ops/sec), `channels` might be the bottleneck. Lock-free ring buffers (using atomic CAS) are faster but much harder to implement correctly.

---

### 540. When should you use sync.Pool?
"Only for **long-lived** objects that are **expensive to allocate** and **frequently created**.
Good: `*bytes.Buffer`, `*gzip.Writer`.
Bad: `*int`, small structs.
If the object is tiny, the overhead of the Pool (locking/interface conv) outweighs the allocation savings. Plus, I must reset the object perfectly to avoid data bleeding."

#### Indepth
**Resets are critical**. If you put a buffer back in the pool containing user A's private data, and then pull it out for user B, you have a security incident. Always `Reset()` *before* putting it back, or *immediately* after taking it out. `defer p.Put(buf)` is a good pattern, but make sure `buf.Reset()` happens.
