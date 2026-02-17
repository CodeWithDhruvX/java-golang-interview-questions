# 🟢 Go Theory Questions: 201–220 Performance Optimization

## 201. How do you avoid unnecessary allocations in Go?

**Answer:**
Allocations are expensive because they tax the Garbage Collector. To avoid them, we follow three rules.

First, **Stack Allocation**: We allow the compiler to keep variables on the stack by avoiding returning pointers to local variables where possible.
Second, **Sizing**: We pre-allocate slices `make([]int, 0, 100)` to avoid resize-copy operations.
Third, **Reuse**: We use `sync.Pool` for short-lived, heavy objects like bytes buffers, resetting and recycling them instead of creating new ones.

---

## 202. How do you reduce GC (Garbage Collector) pressure?

**Answer:**
GC pressure is purely a function of how much "trash" you generate. The GC runs more often if you allocate more.

To reduce it, we switch from Pointers to Values. A slice of pointers `[]*Point` forces the GC to scan every single pointer. A slice of structs `[]Point` is a single block of memory—the GC scans it as one unit.

We also avoid "churn"—creating temporary objects in loops. By moving variable declarations outside hot loops or reusing buffers, we can often cut GC CPU usage by half.

---

## 203. What is String Interning and when to use it?

**Answer:**
String Interning is a technique where you deduplicate strings in memory.

If you read a CSV file with 1 million rows, and the "Country" column is "USA" 500,000 times, standard Go will allocate 500,000 separate string headers. Interning checks a map: "Have I seen 'USA' before?" If yes, it returns the pointer to the existing string.

This saves massive amounts of RAM. Go doesn't do this automatically for runtime strings (only compile-time constants), so we implement it manually map for High-Cardinality repetitious data.

---

## 204. How do you profile Heap Allocations?

**Answer:**
We use `pprof` with the `-alloc_space` flag.

`go tool pprof -http=:8080 mem.profile`.

We look at the "Allocated Space" view, not just "In Use Space." This distinction is vital. A function might create 1GB of data and immediately delete it. It won't show up in "In Use" (RAM usage is low), but it triggers the GC constantly (CPU usage is high). The allocation profile reveals these "trash generators."

---

## 205. How does inlining affect performance?

**Answer:**
Inlining is when the compiler takes the code from a small function and pastes it directly into the caller, replacing the function call.

This removes the **Function Call Overhead** (stack creation, register saving). More importantly, it enables further optimizations—the compiler can see exactly what's happening and maybe remove dead code.

Go inlines small functions automatically. You can check what's being inlined with `go build -gcflags="-m"`. We help the compiler by keeping utility functions short and simple.

---

## 206. What is Profile Guided Optimization (PGO)?

**Answer:**
PGO (available in Go 1.20+) is a technique where you build your app, run it under load, capture a production profile, and then loose that profile back into the build process.

`go build -pgo=cpu.pprof`.

The compiler uses this "cheat sheet" to know which functions are actually called most often. It can then aggressively inline the hot paths and ignore the cold paths. We typically see 5-10% "free" performance gains just by adding this step to our CI pipeline.

---

## 207. How do you optimize JSON Marshaling?

**Answer:**
Standard `encoding/json` uses reflection, which is slow.

For hot paths, we optimized by avoiding reflection. We use **Code Generation** libraries like `easyjson` or `fastjson`.

These tools generate static `Marshal()` methods for your structs at compile time. Instead of inspecting fields at runtime, they write `buffer.WriteString("id:"); buffer.WriteInt(s.ID)`. This is typically 2x-4x faster and creates significantly less garbage.

---

## 208. How do you use sync.Pool effectively?

**Answer:**
`sync.Pool` is a thread-safe bucket of reusable objects.

When you need an object, you call `Get()`. If the pool is empty, it makes a new one. When done, you `Put()` it back.

The catch is that the Pool is emptied completely during every Garbage Collection. It is **not** a cache. It is a mechanism to amortize allocation costs between GC cycles. It’s perfect for `bytes.Buffer` or `gzip.Writer` instances that are expensive to initialize but cheap to reset.

---

## 209. What is False Sharing in concurrency?

**Answer:**
False Sharing acts as a "silent performance killer" on multi-core CPUs.

It happens when two different atomic variables sit next to each other in the same **Cache Line** (usually 64 bytes). If Core A writes to Variable 1, it invalidates the whole cache line. Core B, trying to read Variable 2, is forced to reload memory even though Variable 2 didn't change.

To fix this in high-performance counters, we use **Memory Padding**. We add `_ [56]byte` fields between the variables to force them onto different cache lines.

---

## 210. How do you handle large slices without GC spikes?

**Answer:**
If you have a massive slice (say, 10 million items), the GC has to scan it. If the slice contains pointers (like `[]*User`), the GC must chase every single pointer.

To optimize, we flatten the data. We switch to `[]User` (value type). Now the slice is just one big block of bytes with no external pointers.

However, if even that is too slow (scan time), we might move the data **off-heap**—allocating memory manually via `mmap` (using syscalls) where the GC doesn't look. This is extreme and rarely needed, but it's the ultimate escape hatch.

---

## 211. How do you optimize string concatenation?

**Answer:**
String concatenation using `+` creates a new string every time. `s += "a"` inside a loop is O(N^2) complexity because it recopies the whole string repeatedly.

We use `strings.Builder`.

It maintains an internal buffer and allows you to append bytes efficiently. When you call `.String()`, it essentially casts the buffer to a string without reallocating (sharing the memory), making it the most efficient way to build strings dynamically.

---

## 212. How do you minimize Context Switches?

**Answer:**
Context switches happen when the OS pauses one thread to run another.

In Go, goroutine switches are cheap, but not free. The main cause of excessive switching is **Lock Contention**. If 1000 goroutines all try to grab the same Mutex, they all wake up, fight, sleep, and wake up again.

To optimize, we switch to **Sharding**. Instead of one Big Map with one Lock, we use 32 smaller maps each with their own lock. This distributes the contention, allowing more threads to run in parallel without blocking.

---

## 213. How does `defer` impact performance?

**Answer:**
Historically, `defer` had a noticeable overhead (allocating a closure).

Since Go 1.14+, `defer` is **open-coded**. The compiler essentially injects the deferred code directly at the return points. In modern Go, the overhead is practically zero (nanoseconds).

You should basically never avoid `defer` for performance reasons anymore unless you are writing a tight loop processing pixels in an image. For I/O and locks, the safety benefit outweighs the negligible cost.

---

## 214. How do you use the `unsafe` package for zero-copy?

**Answer:**
Standard Go requires casting `[]byte` to `string` to copy the memory, because strings are immutable.

With `unsafe`, we can trick the compiler. We create a string header that points to the byte slice's underlying array.

`func BytesToString(b []byte) string`. This is **dangerous**. If you modify the byte slice later, the "immutable" string will change, potentially crashing the runtime or causing security bugs. We uses this only in the absolute hottest paths of web frameworks or parsers.

---

## 215. How do you identify Goroutine Leaks?

**Answer:**
A goroutine leak is when you start a `go func` that never returns, usually because it's blocked on a channel read where no message ever comes.

We identify them using `runtime.NumGoroutine()` in tests. We check the count before and after a test case. If `after > before`, we leaked something.

We also use library `goleak` in our test suite. It automatically fails the test if any unexpected goroutines are left running.

---

## 216. How do you improve startup time of a Go app?

**Answer:**
Go apps start fast, but initialization can slow them down.

The main culprit is usually `init()` functions doing I/O—connecting to databases or fetching secrets synchronously.

We optimize by applying **Lazy Loading**. We don't connect to the DB in `init()`. We connect on the first request, or we start the connection in a background goroutine while the HTTP server boots up. This lets Kubernetes mark the pod as "Started" quickly.

---

## 217. How do you write cache-friendly code?

**Answer:**
CPUs love **Spatial Locality**—accessing memory sequentially.

Linked lists are cache-unfriendly (random jumps). Slice of pointers `[]*Node` are also cache-unfriendly.

We optimize by using Slice of Values `[]Node`. We iterate `for i := range nodes`. This allows the CPU pre-fetcher to pull the next chunk of data into L1 cache before the code asks for it, often resulting in a 10x speedup for numerical processing.

---

## 218. How to optimize map keys?

**Answer:**
Using strings as map keys is convenient but involves hashing the string every lookup.

If performance is critical, we try to use **Integers** or **Structs of Integers**.

If we must use strings, and the strings are long, hashing becomes expensive. We might compute the hash once, store it `type KeyedData struct { hash uint64; data string }`, and use the integer hash as the map key (handling collisions manually). This is how Go's internal type map works.

---

## 219. What is Escape Analysis debugging?

**Answer:**
It's the process of asking the compiler "Why did you put this on the heap?"

We run `go build -gcflags="-m -m"`. The output tells us: "parameter x escapes to heap".

Sometimes it's subtle—like passing a variable to `fmt.Println` (which takes `any`). Since `any` is an interface, the value must be boxed on the heap. Recognizing this, we might switch to a specialized logger that takes concrete types to avoid that allocation in a hot loop.

---

## 220. How do you use alignment to save memory?

**Answer:**
Structs are padded for memory alignment.

`struct { bool; int64; bool }` takes 24 bytes. (1 byte bool + 7 bytes padding + 8 bytes int + 1 byte bool + 7 bytes padding).

If we reorder to `struct { int64; bool; bool }`, it takes 16 bytes. The two bools fit in the tail padding of the int64 (or start a new word efficiently).

For a struct used 100 million times in memory, checking field order with tools like `fieldalignment` can save gigabytes of RAM.
