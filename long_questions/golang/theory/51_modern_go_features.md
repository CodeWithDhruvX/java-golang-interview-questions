# 🆕 Go Theory Questions: 1001–1020 Modern Go Features (v1.22–v1.24)

## 1001. What is the loop variable scope change in Go 1.22?

**Answer:**
**Fixed** the most common Go mistake.
Old: `for i := range 3 { go func() { print(i) } }` -> Printed `2 2 2`.
New (1.22): The variable `i` is created **per iteration**.
Prints `0 1 2`.
We no longer need `i := i` inside the loop.

---

## 1002. How do you iterate over integers in Go 1.22 (`for i := range n`)?

**Answer:**
Syntactic sugar.
`for i := range 10 { ... }`.
Iterates 0 to 9.
Replaces `for i := 0; i < 10; i++`.
Clean, readable, and less error-prone (no off-by-one errors).

---

## 1003. How does `net/http.ServeMux` support wildcards and methods in Go 1.22?

**Answer:**
New patterns:
`mux.HandleFunc("POST /items/{id}", handler)`.
- **Method**: `POST` is enforced.
- **Wildcard**: `{id}` matches a segment.
- **Retrieval**: `id := r.PathValue("id")`.
This makes external routers like `chi` or `gorilla/mux` unnecessary for simple services.

---

## 1004. What is the new `math/rand/v2` package?

**Answer:**
Go 1.22.
Faster, cleaner API.
Global state `rand.IntN(100)` is now thread-safe and faster (PCG algorithm).
Removed `Read` (use crypto/rand for bytes).
Top level functions `Float64`, `IntN` are standard. No more `rand.Seed(time.Now())`; it auto-seeds.

---

## 1005. What are Go Iterators (`range-over-func`) in Go 1.23?

**Answer:**
Standard Protocol for iteration.
Type: `func(yield func(V) bool)`.
We can range over a function!
```go
func Fib(yield func(int) bool) {
    a, b := 0, 1
    for {
        if !yield(a) { return }
        a, b = b, a+b
    }
}
for n := range Fib { ... }
```
This standardizes "Lazy Sequences" in Go.

---

## 1006. How do you use the `unique` package in Go 1.23?

**Answer:**
String Interning / Canonicalization.
`h := unique.Make("hello")`.
If you make "hello" again, it returns the same `Handle`.
Comparing Handles is cheap (pointer check).
Useful for parsers or deduplicating millions of identical strings in memory.

---

## 1007. What improvements were made to `time.Timer` garbage collection in Go 1.23?

**Answer:**
Timers are now promptly garbage collected even if not stopped.
Old: `time.After(1h)` inside a tight loop would leak memory until the hour passed.
New: The channel is unreferenced, so the runtime reclaims the timer immediately.
`defer t.Stop()` is less critical (though still good practice).

---

## 1008. What are generic type aliases in Go 1.24?

**Answer:**
We can now alias generic types.
`type Vector[T any] = []T`.
Old: Only defined types (`type Vector[T] []T`) worked, which required casting.
Alias means `Vector[int]` IS `[]int`.
Useful for refactoring generic libraries.

---

## 1009. How do you use the `go tool` directive in `go.mod` (Go 1.24)?

**Answer:**
Tracks executable dependencies.
`tool "golang.org/x/tools/cmd/stringer"`.
`go.mod` manages the version of the tool.
We can run `go tool stringer` without `go install`ing it globally. This replaces the hacky `tools.go` file pattern.

---

## 1010. What is `os.Root` and how does it improve file system isolation (Go 1.24)?

**Answer:**
`root, _ := os.OpenRoot("/tmp/sandbox")`.
`root.Open("file.txt")`.
It guarantees safely that you cannot access `../secret`.
It uses `openat` syscalls.
Prevents **Path Traversal** attacks natively at the OS handle level.

---

## 1011. How do you implement weak pointers in Go 1.24?

**Answer:**
`weak.Pointer[T]`.
Points to an object but doesn't prevent GC.
`p := weak.Make(&obj)`.
`val := p.Value()`.
If `val` is nil, the object was collected.
Useful for Caching keys where you want the entry to disappear if no one else is using the key.

---

## 1012. What is the `omitzero` struct tag option?

**Answer:**
JSON marshaling.
`Field int `json:",omitzero"``.
If Field is 0 (Zero Value), it is omitted.
Different from `omitempty`:
`omitempty` omits if Go definition of "Empty" (False, 0, nil).
`omitzero` is strictly for the zero value of the type. (Subtle difference, mainly for structs).

---

## 1013. How do you use `testing.B.Loop` for benchmarks?

**Answer:**
Replaces `for i := 0; i < b.N; i++`.
```go
for b.Loop() {
    doWork()
}
```
Simplifies benchmark code. The runtime handles the iteration count logic internally.

---

## 1014. How does Go 1.24 support FIPS 140-3 compliance?

**Answer:**
`GOFIPS=1`.
Uses the **FIPS 140-3 verified crypto module** inside the Go runtime (BoringCrypto is being phased/merged).
Ensures your Go app only uses govt-approved algorithms (AES, RSA, etc.) and panics if you try to use MD5. Critical for FedRAMP environments.

---

## 1015. What are usage comparisons for `slices.Concat`?

**Answer:**
Go 1.22+.
`s := slices.Concat(s1, s2, s3)`.
It calculates total size, allocates once, and copies.
Much faster and cleaner than `append(append(s1, s2...), s3...)` which might reallocate 3 times.

---

## 1016. How do you use `runtime.AddCleanup` vs `SetFinalizer`?

**Answer:**
Go 1.24.
`runtime.AddCleanup(ptr, cleanupFunc, arg)`.
Better than Finalizer:
1.  Does not resurrect the object.
2.  Can attach multiple cleanups.
3.  More deterministic (runs closer to collection).
Used for notifying C pools or closing file descriptors attached to Go objects.

---

## 1017. What are the new WASM export capabilities in Go?

**Answer:**
`//go:wasmexport MyFunc`.
Compiles a Go function to a standard Wasm export.
Callable from C/Rust/JS (without `wasm_exec.js` glue if using `GOOS=wasip1`).
Enables Go to write **Envoy Filters** or **database extensions** (Postgres plugins in Wasm).

---

## 1018. How do you debug using `go build -asan`?

**Answer:**
Address Sanitizer (Google).
`go build -asan`.
Detects **Use After Free** in CGO code or unsafe Go code.
Slower execution, but finds memory corruption bugs that the Race Detector misses.

---

## 1019. How do you manage tool dependencies without a `tools.go` file now?

**Answer:**
We verify `tool` directive in `go.mod`.
We run `go get -tool github.com/foo/bar`.
It adds it to `go.mod`.
We execute: `go tool bar`.
No more dummy imports needed to keep the module in the dependency graph.

---

## 1020. What is the anticipated "Flight Recorder" feature?

**Answer:**
`debug/flightrecorder` (Experimental).
Continuous profiling buffer.
Keeps last 1 minute of CPU/Heap profile in RAM.
On crash (or HTTP request), dump the buffer.
Solves the "Heisenbug" problem: "What was happening right before the crash?" without needing `pprof` running 24/7.
