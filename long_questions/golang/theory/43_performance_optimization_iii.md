# 🏎️ Go Theory Questions: 841–860 Performance Optimization III

## 841. How do you write cache-friendly code in Go?

**Answer:**
**Data Locality**.
CPU caches lines (64 bytes).
1.  **Struct Layout**: Pack fields. `bool` (1 byte) + `float64` (8 bytes) = 16 bytes (padding). Order matters.
2.  **Slices vs Linked Lists**: Slices are contiguous RAM. Lists are scattered. Always prefer Slices.
3.  **Row vs Column**: Iterating `matrix[i][j]` (Row-Major) is fast. `matrix[j][i]` (Column-Major) is cache-thrashing slow.

---

## 842. How do you improve startup time of a Go app?

**Answer:**
1.  **Remove `init()`**: They run sequentially on main thread. Lazy load singletons (`sync.Once`).
2.  **Binary Size**: `go build -ldflags="-s -w"` (Strip DWARF). Smaller binary loads faster.
3.  **Plugins**: Avoid `plugin` package; it adds overhead.
4.  **Reflection**: Reduce heavy reflection at startup (like complex DI containers scanning the whole classpath).

---

## 843. How do you reduce lock contention in Go?

**Answer:**
1.  **Sharding**: Split one Map with one Mutex into 32 Maps with 32 Mutexes (Concurrent Map).
2.  **RWMutex**: If Writes are rare, use `RLock` for reads.
3.  **Channels**: Use channels to serialize access (monitor pattern) - sometimes slower but cleaner.
4.  **Atomics**: `atomic.AddInt64` is CPU instruction level, much faster than Mutex.

---

## 844. How do you identify goroutine leaks?

**Answer:**
**pprof** `goroutine` profile.
Look at the count. "Why do I have 50,000 goroutines?"
Click on the graph. "49,000 are stuck at `channel receive` in `worker.go:45`".
Cause: The channel was never closed, or the context was never cancelled, so the workers are waiting forever.

---

## 845. How do you minimize context switches?

**Answer:**
1.  **Reduce Goroutines**: Don't spawn a goroutine for a 1ms task. Use a Worker Pool.
2.  **Reduce Syscalls**: Batch I/O. Use buffered writers (`bufio`).
3.  **Affinity**: `runtime.LockOSThread()` binds a G to an M (Thread). Useful for CGO/Graphics, usually hurts general web server performance.

---

## 846. How do you use sync.Pool effectively?

**Answer:**
Store **Allocated Buffers**.
`var bufPool = sync.Pool{ New: func() any { return new(bytes.Buffer) } }`.
Get: `b := bufPool.Get().(*bytes.Buffer); b.Reset()`.
Put: `bufPool.Put(b)`.
**Caveat**: `sync.Pool` is cleared on every GC. It is not a permanent cache. It smoothens the allocation curve between GC cycles.

---

## 847. How do you optimize string concatenation?

**Answer:**
Naive: `s += "a"` (O(N^2) if in loop).
Better: `fmt.Sprintf` (Ok, reflective).
Best: **`strings.Builder`**.
`sb.Grow(n)` (Pre-allocate).
`sb.WriteString("a")`.
`sb.String()`.
Builder uses a `[]byte` buffer internally and converts to string using `unsafe` (no copy) when you call `String()`.

---

## 848. How do you use benchmarking to choose better algorithms?

**Answer:**
Don't guess.
Write `BenchmarkBubbleSort` vs `BenchmarkQuickSort`.
Run `go test -bench=. -benchmem`.
Compare `ns/op`.
Sometimes O(N^2) is faster than O(N log N) for small N (e.g., N < 20) due to constant factors and cache locality.

---

## 849. How do you eliminate redundant computations?

**Answer:**
**Memoization**.
Store result of function `f(x)` in a Map `cache[x]`.
Implementation:
`var cache sync.Map`.
Inside function:
`if val, ok := cache.Load(key); ok { return val }`.
`val := compute(); cache.Store(key, val)`.
Useful for pure functions (Factorial, Fibonacci, Regex Compilation).

---

## 850. How do you spot unnecessary interface conversions?

**Answer:**
Escape Analysis/CPU Profiling.
`runtime.convT2E` or `runtime.convT2I`.
If you see these high in CPU profile, you are boxing values into interfaces.
Example: `sort.Slice` (Interface) vs `slices.Sort` (Generics).
Generics eliminate the interface conversion overhead, inlining the comparison function.

---

## 851. How do you improve performance of I/O-heavy apps?

**Answer:**
1.  **Asynchrony**: Netpoller (Go default).
2.  **Buffering**: `bufio.NewWriter(f)` (4KB buffer) reduces syscalls 4000x.
3.  **IO_URING**: (Linux). Libraries like `rio` use the new kernel async IO ring.
4.  **Zero Copy**: `Splice` / `Sendfile`.

---

## 852. How do you handle large slices without GC spikes?

**Answer:**
GC scans slices. If slice has 10M pointers, GC checks all 10M.
Optimization: **Slice of Values** (structs without pointers).
`[]int` -> GC ignored (it's just bytes).
`[]*int` -> GC scans.
If you need a map of 100M items, use an Off-Heap library or a sharded map of int-keyed structs to hide pointers from the collector.

---

## 853. How do you reduce reflection usage in Go?

**Answer:**
Reflection is runtime inspection (Slow).
1.  **Code Gen**: Generate the specific code (Protobuf, EasyJSON).
2.  **Interfaces**: Use defined interfaces instead of `interface{}`/Reflection where possible.
3.  **Type Assertion**: `if s, ok := v.(Stringer); ok` is faster than `reflect.ValueOf(v).Method("String")`.

---

## 854. How do you apply zero-copy techniques?

**Answer:**
1.  **Slicing**: `b := data[10:20]` creates a view, no copy.
2.  **Unsafe**: Cast `[]byte` to `string`.
`*(*string)(unsafe.Pointer(&bytes))`
(Dangerous: string is immutable, if you modify bytes later, you break the string).
3.  **Network**: `ReadFrom` / `WriteTo` interfaces often bypass userspace buffers.

---

## 855. How do you build a custom Go compiler plugin?

**Answer:**
Go compiler is written in Go.
We don't usually write "Plugins". We write **Analysers** or **Linters**.
If modifying compilation:
We fork `go/src/cmd/compile`.
Or use `goyacc` for custom parsers.
Real compiler plugins (loading shared objects) are supported via `plugin` package but are extremely limited (must match exact Go version/dependencies).

---

## 856. What is SSA (Static Single Assignment) form in Go?

**Answer:**
The Go Compiler's backend uses SSA.
Intermediate Representation where every variable is assigned exactly once.
`x = 1; x = 2` becomes `x1 = 1; x2 = 2`.
This makes optimization passes simpler:
- **Dead Code Elimination**.
- **Bounds Check Elimination**.
- **Register Allocation**.
You can see it: `GOSSAFUNC=main go build`. It generates an HTML file showing the SSA passes.

---

## 857. How does Go handle type inference?

**Answer:**
**Unification Algorithm**.
`x := 10` -> 10 is int, so x is int.
For Generics: `func F[T any](p T)`.
Call `F(10)`. Compiler sees 10 is int. T unifies to int.
It primarily happens within function bodies or generic instantiation. It does *not* do global type inference (like functional languages) to keep compilation fast.

---

## 858. What is escape analysis in Go?

**Answer:**
Compiler phase: "Does this variable outlive the function?".
Yes -> Heap (GC).
No -> Stack (Cheap).
Logic:
- Ref taken and passed to another goroutine? -> Escape.
- Ref passed to interface method? -> Escape (usually).
- Large size? -> Escape.
We optimize by ensuring short-lived objects stay on stack.

---

## 859. How does inlining affect performance in Go?

**Answer:**
**Inlining**: Copy function body to call site.
Pros: Removes function call overhead (PUSH/POP stack, registers). Enables further optimizations (Dead code).
Cons: Increases binary size (Instruction Cache pressure).
Go inlines small "leaf" functions (cost < 80).
`go build -gcflags="-m -m"` shows inlining decisions.

---

## 860. What are build constraints and how do they work?

**Answer:**
(See Q 786).
Metadata telling compiler *when* to include a file.
`//go:build linux && amd64`.
Logic:
- `&&` AND
- `||` OR
- `!` NOT
Old syntax: `// +build linux` is deprecated.
This allows a single source tree to support Windows, Linux, Wasm, and JS seamlessly.
