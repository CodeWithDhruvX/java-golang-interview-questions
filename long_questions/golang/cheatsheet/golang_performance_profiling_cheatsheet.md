# Golang Performance & Profiling Cheatsheet

Optimization techniques and tooling for high-performance Go applications.

## 🟢 Profiling (pprof)

### 1. Enabling pprof
Add this one-liner to your `main()` to expose profiling endpoints.
```go
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    // ... app logic ...
}
```

### 2. Capturing Profiles
Run `go tool pprof` against the running app.

**CPU Profile (Busy loops):**
```bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

**Heap Profile (Memory Leaks):**
```bash
go tool pprof http://localhost:6060/debug/pprof/heap
```

**Goroutine Blocking (Wait times):**
```bash
go tool pprof http://localhost:6060/debug/pprof/block
```

### 3. Analyzing Profiles
Inside the pprof interactive shell:
- `top`: Show functions consuming most resources.
- `list <Function>`: Show line-by-line cost.
- `web`: Open visualization in browser (requires Graphviz).

---

## 🟡 Execution Tracing

Visualizes the timeline of every goroutine, syscall, and GC event.

```go
func main() {
    f, _ := os.Create("trace.out")
    defer f.Close()
    trace.Start(f)
    defer trace.Stop()
    
    // ... workload ...
}
```

**View Trace:**
```bash
go tool trace trace.out
```
*Great for debugging latency spikes and scheduler delays.*

---

## 🔴 Compiler Optimizations

### 1. Escape Analysis
Check what variables move to the heap.
```bash
go build -gcflags="-m" main.go
```
*Goal: Keep objects on stack (zero allocation cost) by not returning pointers to local variables where possible.*

### 2. Inlining
The compiler replaces function calls with the function body to save overhead.
- **Rule:** Functions must be simple (no loops, select, switch, defer, recover).
- **Check:** `go build -gcflags="-m -m"` (2x m for inlining details).

### 3. Bounds Check Elimination (BCE)
The compiler skips array index checks if it can prove safety.
```go
// Has bounds check
func f(s []int) {
    _ = s[0]
    _ = s[1]
    _ = s[2]
}

// No bounds check (Hinting)
func g(s []int) {
    _ = s[2] // Compiler checks once here
    _ = s[0] // Safe
    _ = s[1] // Safe
}
```
**Check:** `go build -gcflags="-d=ssa/check_bce/debug=1"`

---

## 🔵 Memory Optimization Tricks

### 1. Struct Field Alignment (Padding)
Order matters! Place larger fields first to avoid padding.
```go
// Bad: 24 bytes (on 64-bit)
type Bad struct {
    A bool    // 1 byte + 7 padding
    B float64 // 8 bytes
    C int32   // 4 bytes + 4 padding
}

// Good: 16 bytes
type Good struct {
    B float64 // 8 bytes
    C int32   // 4 bytes
    A bool    // 1 byte + 3 padding
}
```
*Tool: `fieldalignment` from `golang.org/x/tools/go/analysis/passes/fieldalignment`*

### 2. Strings vs Bytes
- Conversion `string(bytes)` allocates memory (copy).
- **Constraint:** Functions accepting `[]byte` are often faster than `string` if you are modifying data.

### 3. String Concatenation
**Bad:** `s += "str"` (O(N^2) reallocation)
**Good:** `strings.Builder`

```go
var b strings.Builder
b.Grow(100) // Pre-allocate
for i := 0; i < 100; i++ {
    b.WriteString("str")
}
return b.String()
```
