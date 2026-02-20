# 🏎️ **841–860: Performance Optimization (Part 2)**

### 841. How do you write cache-friendly code in Go?
"I access memory sequentially.
Iterating a slice (`[]int`) is fast because the CPU prefetcher loads the next cache line.
Iterating a linked list is slow (random pointer hopping).
I use **Data-Oriented Design**: struct of arrays instead of array of structs for cache locality."

#### Indepth
**Struct Alignment**. CPU reads memory in 64-byte chunks (cache lines). If your struct has `bool, int64, bool`, it adds padding bytes (7 wasted bytes after first bool, 7 after second). Reorder fields from largest to smallest (`int64, bool, bool`) to minimize padding. This packs more structs into a single cache line, reducing RAM bandwidth.

---

### 842. How do you improve startup time of a Go app?
"1.  **Remove `init()`**: Lazy load global resources (e.g., connect DB on first request).
2.  **Trim Dependencies**: Huge heavy libs add initialization cost.
3.  **Pack**: Use UPX to compress binary if disk read is the bottleneck."

#### Indepth
**Plugin Architecture**. Go binaries are static and huge. If you have 50 features but the user only needs 1, compiling all 50 slows startup. Use Go Plugins (`-buildmode=plugin`) or separate binaries (Micro-architecture) to load features on demand. NOTE: Go Plugins are finicky on Linux/Mac and unsupported on Windows.

---

### 843. How do you reduce lock contention in Go?
"1.  **Shorten Critical Sections**: Do the heavy work *outside* the lock.
2.  **Sharding**: Split one map into 32 maps with 32 locks.
3.  **Atomics**: Use `atomic.AddInt64` instead of `Mutex` for simple counters."

#### Indepth
**RWMutex**. Use `sync.RWMutex` if you have many readers and few writers. But beware: **Writer Starvation**. If new readers keep arriving, the writer might never get the lock. Go 1.19+ fixed this by giving writers priority, but in older versions, a busy read-loop could hang the writer indefinitely.

---

### 844. How do you identify goroutine leaks?
"I use `runtime.NumGoroutine()` in tests.
Before: `n := runtime.NumGoroutine()`.
After: `if runtime.NumGoroutine() > n { fail }`.
Or use `goleak` library which does this automatically and excludes standard library background routines."

#### Indepth
**Debug Handlers**. Go provides a hidden gem: `/debug/pprof/goroutine?debug=2`. It dumps the full stack trace of *every* running goroutine. If you see 10,000 goroutines all "waiting on channel" at line 54, you found your leak. It's much faster than reading code.

---

### 845. How do you minimize context switches?
"Context switches happen when a thread blocks (IO/Lock).
I avoid blocking system calls in tight loops.
I use **Netpoll** (epoll) via standard library for network IO.
I stick to non-blocking channel operations where possible."

#### Indepth
**Processor Affinity**. The OS scheduler moves threads between Cores (Context Switch). This invalidates L1/L2 cache. Used `runtime.LockOSThread()` to pin a sensitive goroutine (like an Audio processing loop) to a specific OS thread, and use `taskset` (Linux) to pin that thread to a specific CPU core for maximum cache locality.

---

### 846. How do you use sync.Pool effectively?
"I put **reset** objects in.
`p.Put(buf)`.
I **must** `buf.Reset()` before putting it back, otherwise the next user gets dirty data.
I don't use it for small objects (`int`) where the overhead of the pool API exceeds the allocation cost."

#### Indepth
**Double Put**. `sync.Pool` has no protection against putting the same pointer twice. If you do `p.Put(x); p.Put(x)`, two goroutines will `Get` the same pointer `x` later and race condition on it. This is a devastating and hard-to-debug bug. Use `go vet` or custom linters to catch this.

---

### 847. How do you optimize string concatenation?
"1.  `+` operator: Good for small, known number of strings.
2.  `strings.Builder`: Best for loops. It minimizes copying.
`b.Grow(n); b.WriteString(s)`.
`bytes.Buffer` converts to `string` at the end (copy). `strings.Builder` returns the underlying byte array as a string (zero copy)."

#### Indepth
**Pre-allocation**. Always use `var b strings.Builder; b.Grow(n)`. If you don't call `Grow`, the builder starts small (e.g., 64 bytes) and reallocates/copies every time it fills up (doubling strategy). Guessing the size (even roughly) eliminates these expensive reallocations.

---

### 848. How do you use benchmarking to choose better algorithms?
"I implement both.
`BenchmarkBubbleSort` vs `BenchmarkQuickSort`.
I test with different N (10, 1000, 1M).
Sometimes O(N^2) is faster than O(N log N) for N < 20 because of CPU caching and branch prediction constants."

#### Indepth
**Sub-benchmarks**. Use `b.Run("size=10", func(b...) { ... })`. This creates a hierarchical view. `BenchmarkSort/size=10`, `BenchmarkSort/size=1000`. This isolates the variable (Input Size) and shows the growth curve clearly in the output, allowing you to spot the "tipping point" where one algo beats another.

---

### 849. How do you eliminate redundant computations?
"**Memoization**.
I store results of expensive functions in a map.
Check map -> Return.
Else Compute -> Store.
I use `singleflight` to prevent 100 concurrent requests from computing the same value simultaneously."

#### Indepth
**Pre-computation**. If a value depends only on constants (e.g., `CRC32 Table`), compute it in `init()` or `var x = func() {...}()`. Don't compute it on every request. For expensive math (e.g., Sin/Cos tables), generate a `.go` file with the precomputed lookup table using `go generate`. Lookup is O(1) vs Calculation O(N).

---

### 850. How do you spot unnecessary interface conversions?
"Escape Analysis (`go build -gcflags="-m"`).
If I assign a concrete type to an interface, it *might* escape to heap if called dynamically.
I profile with `pprof`. If `runtime.convT2E` shows up, I'm converting types too much (boxing)."

#### Indepth
**Boxing Cost**. Assigning `int` to `interface{}` creates a new allocation on the heap (to hold the int value and method table), unless the compiler can optimize it. If you have a map `map[string]interface{}` and store millions of ints, you get millions of tiny allocations. Use `map[string]int` or a custom union struct if performance matters.

---

### 851. How do you improve performance of I/O-heavy apps?
"I use **Buffering**.
`bufio.NewWriter(file)`.
Instead of writing 1 byte 1000 times (1000 syscalls), it writes 4KB once.
For network, I use **Pipelining** (sending multiple requests without waiting for individual responses)."

#### Indepth
**io_uring**. On modern Linux (5.10+), `epoll` (Go's Netpoll) is good, but `io_uring` is better for File IO. Go doesn't use `io_uring` by default yet. Libraries like `rio` allow you to use asynchronous file IO (submit request, get callback later) which is significantly faster for database-like workloads than blocking Syscalls.

---

### 852. How do you handle large slices without GC spikes?
"If I have a huge cache `[]Item` (1GB), the GC scans it every cycle.
Optimization: Use a map of `int` keys (pointers cause scanning) or store data off-heap using **cgo** or **syscall.Mmap**.
Or explicitly set `GOGC=off` and manage GC manually (risky)."

#### Indepth
**Ballast**. A trick used by Twitch. Allocate a giant byte slice `make([]byte, 10<<30)` (10GB) at startup and keep it alive. This forced the GC target (Heap * 2) to be huge (20GB). The GC runs less frequently because the "small" garbage (100MB) doesn't trigger the doubling threshold relative to the 10GB ballast. Obsolete with `GOGC` tuning in Go 1.19 (`SetMemoryLimit`).

---

### 853. How do you reduce reflection usage in Go?
"Reflection is slow because it can't be optimized by the compiler.
I generate code instead (`go generate`).
DAO libraries that use reflection are convenient but slow. I prefer generating the SQL scanning code for each struct at build time."

#### Indepth
**Modern Reflection**. `reflect` isn't *always* slow. `reflect.Type` operations are fast (cached). `reflect.Value` operations differ. `Current Go` optimizes common reflection patterns. But `reflect.Call` (invoking a function by name) is still ~5-10x slower than direct call. Use it for setup/config, never for the hot loop.

---

### 854. How do you apply zero-copy techniques?
"I use `splice` (on Linux) to move data between file descriptors (Network <-> Pipe <-> File) without copying to User Space.
Go exposes this via `io.Copy` specializations.
Also, slicing `b[:n]` instead of `append` prevents data movement in memory."

#### Indepth
**sendfile**. `io.Copy` automatically uses the `sendfile` syscall if Source is a `File` and Dest is a `TCPConn`. This delegates the transfer to the OS Kernel (DMA). The data goes Disk -> Kernel Buffer -> NIC without ever touching your Go program's RAM or CPU. This is how Nginx/Go static file servers achieve 100Gbps.

---

### 855. How do you avoid false sharing in Go?
"False Sharing: Two atomics sit on the same Cache Line (64 bytes).
Core 1 writes A. Core 2 writes B.
They invalidate each other's L1 cache constantly.
Fix: **Padding**.
`type Padded struct { A uint64; _ [56]byte; B uint64 }`."

#### Indepth
**Cache Line Size**. `64 bytes` is standard on x86/ARM. But some architectures (Apple M1/M2) have `128 byte` cache lines. If you optimize for 64, you might still false-share on Mac. Go's `cpu.CacheLinePad` helps abstract this, but manual padding ensures isolation across platforms.

---

### 856. How do you optimize regular expressions in Go?
"Go's `regexp` is safe (O(n)) but slower than PCRE.
I avoid `regexp` inside loops. `MustCompile` outside.
If simple, I use `strings` package (`Contains`, `Split`). It's 10x-100x faster."

#### Indepth
**Pre-compilation**. Always use `var re = regexp.MustCompile(...)` at the global package level (or `init`). compiling a regex is expensive (it builds a state machine). If you do `regexp.MatchString` inside a loop, it re-compiles the regex *every iteration*. This is a top-3 performance killer in Go apps.

---

### 857. How do you use standard library `sort` efficiently?
"`sort.Slice` uses reflection (slower).
`sort.Ints` is faster.
Defining `Len/Less/Swap` methods on my type is fastest (no reflection).
If identifying top K items, I don't sort the whole slice. I use a Heap (O(N log K))."

#### Indepth
**pdqsort**. As of Go 1.19, the underlying sort algorithm changed from Quicksort to **pdqsort** (Pattern-Defeating Quicksort). It detects patterns (already sorted, reverse sorted) and runs in O(N). It is significantly faster for real-world data which is often partially sorted.

---

### 858. How do you benchmark memory allocations per function?
"`go test -bench=. -benchmem`.
Look at `allocs/op`.
My goal is 0 for hot path functions.
If 1 alloc/op, it might be the return value escaping or an interface conversion."

#### Indepth
**Allocs/op**. `0 allocs/op` isn't always best. A stack allocation that is copied 10 times might be slower than 1 heap allocation shared 10 times. Focus on `ns/op`. But generally, allocations kill throughput because they trigger **GC**. Lower allocs = Less GC work = Higher Throughput.

---

### 859. How do you optimize JSON unmarshaling?
"Standard `encoding/json` uses reflection.
I use **code generation** libraries like `easyjson` or `fastjson`.
They generate `UnmarshalJSON` methods that parse bytes directly without overhead.
Performance gain is typically 2x-5x."

#### Indepth
**Safety Trade-off**. `easyjson` / `fastjson` are fast because they skip validity checks (e.g., duplicate keys, detailed error messages) and avoid reflection. They are perfect for trusted internal streams. For public APIs, use standard lib `encoding/json` to ensure the input is valid standard JSON, preventing weird parsing bugs.

---

### 860. How do you use PGO (Profile Guided Optimization)?
"New in Go 1.20+.
1. Run app in prod, capture `cpu.prof`.
2. Check in `default.pgo`.
3. `go build -pgo=auto`.
The compiler sees which functions are called most and makes inlining decisions based on real usage, boosting perf by ~5-10%."

#### Indepth
**Iterative Builds**. PGO creates a chicken-and-egg problem. You need a profile to build the optimized binary, but you need the binary to get a profile. Pipeline: Build v1 (Standard) -> Deploy -> Collect Profile -> Commit to Repo -> Build v2 (Optimized with v1's profile). Repeat. The profile matches "close enough" even if code changes slightly.
