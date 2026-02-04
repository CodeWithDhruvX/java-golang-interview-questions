## 🏎️ Performance Optimization (Questions 841-860)

### Question 841: How do you optimize memory usage in Go?

**Answer:**
1.  **Reduce Allocations:** Reuse buffers (`sync.Pool`).
2.  **Struct Layout:** Order fields from largest to smallest (padding minimization).
3.  **Avoid Pointers:** Store values in slices directly (`[]T` not `[]*T`) to help GC scanning.

---

### Question 842: How do you avoid unnecessary allocations?

**Answer:**
(See Q525). Use `make` with capacity. Use `strings.Builder`.

---

### Question 843: How do you reduce GC pressure in Go apps?

**Answer:**
Allocation Rate = GC Frequence.
Allocating less -> Less GC.
Reuse objects.
Use **Ballast** (allocate a large byte array on startup) to artificially increase Heap size, triggering GC less often (Legacy technique, less needed with GOMEMLIMIT).

---

### Question 844: How do you profile heap allocations?

**Answer:**
`go tool pprof -alloc_space` (Total bytes allocated ever).
`go tool pprof -inuse_space` (Bytes currently held).

---

### Question 845: How do you use escape analysis to optimize code?

**Answer:**
`go build -gcflags "-m"`.
Look for `moved to heap`.
Try to keep variables on stack (don't return pointers to local variables if copying is cheap).

---

### Question 846: How do you optimize JSON marshaling in Go?

**Answer:**
Standard `encoding/json` uses reflection (Slow).
Use **easyjson** or **fastjson** (Code generation).
It generates explicit `Encode/Decode` methods for your structs, skipping reflection entirely.

---

### Question 847: How do you write cache-friendly code in Go?

**Answer:**
**Data Locality.**
Access arrays sequentially (`arr[i]`, `arr[i+1]`).
Linked Lists (`*next`) scatter data in RAM, causing CPU Cache Misses.
Slices keeps data contiguous.

---

### Question 848: How do you improve startup time of a Go app?

**Answer:**
1.  Remove `init()` functions that do I/O.
2.  Delay DB connection until first request.
3.  Strip debug symbols (`-s -w`).

---

### Question 849: How do you reduce lock contention in Go?

**Answer:**
(See Q531). Sharding, Atomics, Channels.

---

### Question 850: How do you identify goroutine leaks?

**Answer:**
(See Q534). Pprof goroutine dump.

---

### Question 851: How do you minimize context switches?

**Answer:**
Reduce the number of active threads/goroutines.
Use IO-multiplexing (Netpoller) instead of blocking threads.
Don't use `time.Sleep` in tight loops.

---

### Question 852: How do you use `sync.Pool` effectively?

**Answer:**
Store objects that are **expensive to allocate** but **cheap to reset**.
Example: `bytes.Buffer`, `gzip.Writer`.
Dont put Database connections here.

---

### Question 853: How do you optimize string concatenation?

**Answer:**
Use `strings.Builder`.
It prevents creating a new string for every `+` operation (O(N^2)). Builder is O(N).

---

### Question 854: How do you use benchmarking to choose better algorithms?

**Answer:**
Implement both (e.g., BubbleSort vs QuickSort).
Run `Benchmark`.
Compare `ns/op`.

---

### Question 855: How do you eliminate redundant computations?

**Answer:**
**Memoization.**
Cache result of pure function based on input.

---

### Question 856: How do you spot unnecessary interface conversions?

**Answer:**
Code Review.
If you have `func Do(v interface{})`, calling it with `int` forces an allocation (Boxing).
Use Generics `[T any]` to avoid this overhead.

---

### Question 857: How do you improve performance of I/O-heavy apps?

**Answer:**
- **Buffering:** `bufio.NewWriter`.
- **Parallelism:** Read multiple files concurrently.
- **Async I/O:** (Go does this automatically).

---

### Question 858: How do you handle large slices without GC spikes?

**Answer:**
If a map/slice contains pointers, GC scans it.
If `map[int]int` (no pointers), GC skips it.
Use giant slices of non-pointer structs.

---

### Question 859: How do you reduce reflection usage in Go?

**Answer:**
Avoid `encoding/json` or `fmt.Sprintf` in hot paths.
Write specific code instead of generic.

---

### Question 860: How do you apply zero-copy techniques?

**Answer:**
- **`os.File.ReadFrom`:** Uses `sendfile` syscall (Kernel transfers data Disk -> Net without User space copy).
- **Slicing:** `sub := data[1:5]` shares backing array.

---
