# 🆕 **1001–1020: Modern Go Features (v1.22–v1.24)**

### 1001. What is the loop variable scope change in Go 1.22?
"Before 1.22, `for i := range 10` shared the `i` variable across iterations.
This caused the famous `go func() { print(i) }` bug (printing 9, 9, 9...).
In 1.22, `i` is **newly allocated** for each iteration.
We no longer need `i := i` shadowing trick. It just works as expected."

#### Indepth
**The Old Bug**. The classic trap: `for _, v := range items { go func() { process(v) }() }`. All goroutines captured the *same* `v` variable. By the time they ran, the loop had finished and `v` held the last item. The fix was `v := v` inside the loop. Go 1.22 eliminates this entire class of bug by making each iteration's variable a distinct allocation.

---

### 1002. How do you iterate over integers in Go 1.22 (`for i := range n`)?
"Syntactic sugar!
`for i := range 10 { fmt.Println(i) }`.
It prints 0 to 9.
Replaces the verbose `for i := 0; i < 10; i++`.
It’s cleaner and less prone to off-by-one errors."

#### Indepth
**Range over Func**. This is part of a broader "range-over-func" initiative in Go 1.23. The same `range` keyword now works over integers, slices, maps, channels, *and* custom iterator functions. This unification means you learn one syntax and it works everywhere, reducing the cognitive overhead of remembering different loop idioms.

---

### 1003. How does `net/http.ServeMux` support wildcards and methods in Go 1.22?
"Standard library finally got good routing!
`mux.HandleFunc("POST /items/{id}", handler)`.
It supports:
*   **Method Matching**: `POST`, `GET`.
*   **Path Values**: `id := req.PathValue("id")`.
*   **Wildcards**: `/files/{path...}`.
I can finally drop `gorilla/mux` or `chi` for simple projects."

#### Indepth
**Precedence Rules**. The new `ServeMux` has defined precedence: More specific patterns win over less specific ones. `GET /items/{id}` beats `GET /items/`. Method-specific patterns beat method-less ones. This is deterministic and documented, unlike the old `ServeMux` which had subtle ordering bugs when patterns overlapped.

---

### 1004. What is the new `math/rand/v2` package?
"It’s faster and safer.
Global source is now ChaCha8 (cryptographically secure seeding, though still PRNG).
No more `rand.Seed(time.Now())`. It auto-seeds.
Methods like `rand.N(max)` are generic-friendly.
`rand.Intn` is gone; use `rand.IntN`.
Crucially, it removes the global lock contention of v1."

#### Indepth
**ChaCha8 vs PCG**. `math/rand/v2` offers two sources: `rand.New(rand.NewPCG(...))` (fast, for simulations) and `rand.New(rand.NewChaCha8(...))` (cryptographically seeded, default global). The global source uses ChaCha8 to prevent "seed guessing" attacks, even though it's not a CSPRNG. For actual secrets, always use `crypto/rand`.

---

### 1005. What are Go Iterators (`range-over-func`) in Go 1.23?
"Standardized Iterators.
`func(yield func(T) bool)`.
I can use `for v := range mySeq { ... }`.
`mySeq` is just a function that calls `yield(value)` repeatedly.
This unifies iteration over DB rows, API pages, and custom collections without exposing internal state (like `Next()` methods)."

#### Indepth
**Pull vs Push**. Go iterators are "Push" style (the iterator calls `yield`). The alternative is "Pull" style (the consumer calls `Next()`). Go 1.23 also provides `iter.Pull()` to convert a push iterator to a pull iterator when you need to interleave two iterators or need more control. Both styles are now first-class.

---

### 1006. How do you use the `unique` package in Go 1.23?
"It implements **Interning**.
`h := unique.Make("hello")`.
If I have 1 million strings of value 'hello', `unique` stores the string data only once and gives me a lightweight handle (canonical pointer).
Comparing handles (`h1 == h2`) is O(1).
Great for parsers, compilers, or repetitive JSON keys."

#### Indepth
**String Interning**. Without `unique`, two `"hello"` string literals in different packages are separate allocations. With `unique.Make`, they share one. The key benefit is O(1) equality comparison: instead of `strings.Compare` (O(n)), you compare two pointers. This is a massive win in hot paths like symbol tables or HTTP header maps.

---

### 1007. What improvements were made to `time.Timer` garbage collection in Go 1.23?
"Timers are now instantly collectible when unreferenced.
Previously, `time.After(1h)` would leak memory until the hour passed, even if the select finished earlier.
Now, the runtime cleans up the timer immediately if the channel is unreachable.
I don't need `defer timer.Stop()` as religiously anymore (though still good practice)."

#### Indepth
**The Old Leak**. Before 1.23, `select { case <-time.After(1*time.Minute): ... }` in a loop was a classic memory leak. Each iteration created a new timer that lived for a full minute. In a high-throughput loop, you'd accumulate thousands of live timers. The fix was `t := time.NewTimer(d); defer t.Stop()`. Go 1.23 makes the naive version safe.

---

### 1008. What are generic type aliases in Go 1.24?
"I can alias a generic type.
`type MyList[T] = []T`.
Previously, aliases (`=`) only worked on concrete types.
This allows refactoring generic code across packages without breaking the API."

#### Indepth
**Migration Use Case**. Imagine you have `pkg/a` with `type Stack[T any] struct{...}`. You want to move it to `pkg/collections`. Before 1.24, you'd have to update every import. Now: `package a; type Stack[T any] = collections.Stack[T]`. Old code still compiles. You can migrate callsites gradually.

---

### 1009. How do you use the `go tool` directive in `go.mod` (Go 1.24)?
"`go.mod`: `tool golang.org/x/tools/cmd/stringer`.
I can run `go tool stringer`.
It manages tool dependencies **versioning** inside `go.mod`.
No more `tools.go` hack! It ensures every dev uses the exact same linter/generator version."

---

### 1010. What is `os.Root` and how does it improve file system isolation (Go 1.24)?
"`root, _ := os.OpenRoot("/tmp/sandbox")`.
`root.Open("file.txt")`.
If I try `root.Open("../etc/passwd")`, it fails.
It guarantees operations are confined to the directory tree, preventing path traversal attacks safely at the OS level (using `openat` syscalls)."

#### Indepth
**Chroot Alternative**. `os.Root` is safer than `chroot`. `chroot` requires root privileges and can be escaped. `os.Root` uses `openat`/`RESOLVE_NO_ESCAPE` kernel flags, which are enforced by the kernel itself. It works in unprivileged containers and is the correct way to implement sandboxed file access in Go.

---

### 1011. How do you implement weak pointers in Go 1.24?
"`weak.Make(&obj)`.
It returns a pointer that doesn't prevent GC.
If GC runs and `obj` is only held by weak pointers, it collects `obj`.
`ptr := w.Pointer()`. It returns `nil` if collected.
Useful for Caches where I want entries to disappear if memory is tight, without manual eviction."

#### Indepth
**Finalizer Comparison**. `SetFinalizer` had a "resurrection" problem: if the finalizer stored `obj` somewhere, it prevented collection. `weak.Make` avoids this. The GC simply nullifies the weak pointer without running user code. This makes it safe to use in concurrent data structures like `sync.Map`-based caches.

---

### 1012. What is the `omitzero` struct tag option?
"`json:"field,omitzero"`.
Similar to `omitempty`, but checks if the value is the **Zero Value** for its type.
It avoids the confusion where `omitempty` hides `0` or `false` (valid values) when using pointers. `omitzero` is smarter and clearer."

#### Indepth
**The `omitempty` Trap**. `omitempty` omits `0`, `false`, `""`, and `nil`. This means you can't distinguish "user explicitly set score to 0" from "user didn't provide score". The workaround was `*int` (pointer). `omitzero` checks the `IsZero() bool` method if present, allowing types to define their own zero-value semantics.

---

### 1013. How do you use `testing.B.Loop` for benchmarks?
"`for b.Loop() { DoWork() }`.
Replaces `for i := 0; i < b.N; i++`.
It handles the looping logic internally.
It makes setup/teardown logic cleaner (everything outside the loop is setup) and avoids the easy mistake of using `i` incorrectly."

#### Indepth
**Timer Reset**. `b.Loop()` also correctly handles timer reset. With the old `b.N` loop, if you had setup code inside the loop, you'd call `b.ResetTimer()` to exclude it. `b.Loop()` implicitly excludes everything before the first `b.Loop()` call and after the last, making benchmark timing more accurate by default.

---

### 1014. How does Go 1.24 support FIPS 140-3 compliance?
"`GOFIPS=1`.
The standard `crypto` library transparently switches to a FIPS-verified backend (if built with the FIPS toolchain).
It allows government-compliant apps without rewriting code to use specific FIPS SDKs."

#### Indepth
**FIPS 140-3**. This is a US government standard for cryptographic modules. Required for federal contracts, healthcare (HIPAA), and finance. Previously, Go teams had to use BoringCrypto (a Google fork) or external C libraries. Native FIPS support means Go is now a first-class citizen for compliance-heavy industries.

---

### 1015. What are usage comparisons for `slices.Concat`?
"`slices.Concat(s1, s2, s3)`.
It’s optimized.
Allocates the exact final size once.
Memcopies efficiently.
Much faster and cleaner than `append(append(s1, s2...), s3...)`."  

#### Indepth
**Allocation Count**. `append(append(s1, s2...), s3...)` may allocate twice: once for `s1+s2`, once for `(s1+s2)+s3`. `slices.Concat` pre-calculates the total length `len(s1)+len(s2)+len(s3)`, allocates once, and copies. For large slices, this halves memory pressure and eliminates one GC cycle.

---

### 1016. How do you use `runtime.AddCleanup` vs `SetFinalizer`?
"`AddCleanup(ptr, cleanupFunc, arg)`.
It’s the modern, safer `SetFinalizer`.
It allows attaching a cleanup to an object without the resurrection issues of Finalizers.
It's deterministic enough for resource tracking (but still not for urgent resource release—use `defer` for that)."

#### Indepth
**Finalizer Ordering**. `SetFinalizer` had a critical flaw: if object A's finalizer referenced object B, and B was also being finalized, the order was undefined. `AddCleanup` is designed to avoid this by not allowing the cleanup function to access the object being cleaned up (it receives a separate `arg` instead).

---

### 1017. What are the new WASM export capabilities in Go?
"`//go:wasmexport MyFunc`.
Compiles to a WASM function that the host (JS/Rust) can call directly.
No more `js.Global().Set(...)` hackery.
It follows the WASM Component Model standards, making Go WASM modules compatible with other languages."

#### Indepth
**WASM Component Model**. This is the future of WASM interoperability. Instead of sharing raw memory (the old way, requiring manual pointer arithmetic), the Component Model defines high-level types (strings, records, lists). `//go:wasmexport` generates the correct ABI, allowing a Rust host to call a Go function with a `string` argument naturally.

---

### 1018. How do you debug using `go build -asan`?
"**Address Sanitizer** (from C/C++ world).
Detects use-after-free, buffer overflows (in unsafe code/cgo).
Run tests: `go test -asan`.
It adds instrumentation. Slower execution, but catches memory corruption bugs that `-race` might miss."

#### Indepth
**CGO Safety**. `-asan` is most valuable when using CGO. Pure Go is memory-safe by design. But CGO calls C code, which can have buffer overflows. `-asan` instruments both the Go and C sides of the boundary, catching bugs like "C function wrote 5 bytes into a 4-byte Go buffer" that would otherwise cause silent data corruption.

---

### 1019. How do you manage tool dependencies without a `tools.go` file now?
"See Q1009.
I use the `tool` directive in `go.mod`.
`go get -tool golang.org/x/tools/gopls`.
It separates 'application dependencies' (imports) from 'development tools' (binaries).
`go tool` runs them."

#### Indepth
**The Old Hack**. The `tools.go` pattern was: create a file with `//go:build ignore` and blank imports of tools. This forced `go mod tidy` to track them. It was a community workaround, not an official feature. The `go.mod` `tool` directive is the official, clean solution that finally makes tool versioning a first-class concern.

---

### 1020. What is the anticipated "Flight Recorder" feature?
"A lightweight, always-on tracer.
It records execution events to a circular buffer.
If the program crashes, I can dump the buffer.
It tells me what happened in the last 1s before the crash (like a Black Box).
It's designed to be cheap enough for production use."

#### Indepth
**Continuous Profiling**. The Flight Recorder is complementary to tools like Parca or Pyroscope (Continuous Profiling). Those sample CPU/Memory over time. The Flight Recorder captures the exact sequence of events leading to a crash. Together, they give you "What was the system doing overall?" (profiling) AND "What happened in the last second?" (flight recorder).

---
