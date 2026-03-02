# 🚀 05 — Performance Optimization in Go
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- `pprof` profiling (CPU + memory)
- Benchmark tests (`testing.B`)
- Reducing allocations (`sync.Pool`)
- String builder vs concatenation
- Avoiding interface boxing in hot paths
- GC tuning for latency-sensitive apps

---

## ❓ Most Asked Questions

### Q1. How do you profile a Go application using `pprof`?

```go
import (
    "net/http"
    _ "net/http/pprof"  // registers /debug/pprof endpoints
)

func main() {
    // Add pprof HTTP server (in separate goroutine for production)
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    // Your application code...
}
```

```bash
# CPU profiling — 30 second sample
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# Memory/heap profiling
go tool pprof http://localhost:6060/debug/pprof/heap

# Goroutine dump
curl http://localhost:6060/debug/pprof/goroutine?debug=2

# In pprof REPL:
# top        — top CPU consumers
# top -cum   — top cumulative consumers
# web        — open flame graph in browser
# list <fn>  — show annotated source for function

# CPU profile from tests
go test -cpuprofile cpu.prof -memprofile mem.prof -bench . ./...
go tool pprof cpu.prof
```

---

### Q2. How do you write meaningful benchmark tests?

```go
// Basic benchmark
func BenchmarkStringConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        s := ""
        for j := 0; j < 100; j++ {
            s += "x"  // bad: O(n²) — creates new string each time
        }
        _ = s
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var sb strings.Builder
        for j := 0; j < 100; j++ {
            sb.WriteByte('x')  // good: O(n)
        }
        _ = sb.String()
    }
}

// With memory reporting
// go test -bench=. -benchmem -benchtime=5s ./...
// Output:
// BenchmarkStringConcat-8    100000    12345 ns/op    4096 B/op    99 allocs/op
// BenchmarkStringBuilder-8  1000000     1234 ns/op     128 B/op     1 allocs/op

// Reset timer (exclude setup from benchmark)
func BenchmarkWithSetup(b *testing.B) {
    data := setupLargeData()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        processData(data)
    }
}
```

---

### Q3. How does `sync.Pool` help with performance?

```go
// sync.Pool — reuses objects to reduce GC pressure
var bufferPool = sync.Pool{
    New: func() interface{} {
        return &bytes.Buffer{}
    },
}

func processRequest(data []byte) string {
    // Get from pool instead of allocating
    buf := bufferPool.Get().(*bytes.Buffer)
    buf.Reset()               // clear before use
    defer bufferPool.Put(buf) // return to pool when done

    buf.Write(data)
    buf.WriteString(" processed")
    return buf.String()
}

// Real-world examples:
// - HTTP response writers
// - JSON encoders/decoders
// - gzip writers
// - Database row scanners
```

> **Warning:** Pool objects can be collected by GC at any GC cycle. Don't store state that must survive between calls.

---

### Q4. How do you reduce allocations in hot paths?

```go
// ❌ Avoid: interface boxing causes heap allocation
func sumInterface(nums []interface{}) int {
    total := 0
    for _, n := range nums {
        total += n.(int)  // unbox — allocation per call context
    }
    return total
}

// ✅ Better: use concrete types
func sum(nums []int) int {
    total := 0
    for _, n := range nums { total += n }
    return total
}

// ❌ Avoid: append in loop without pre-allocation
func collectBad(n int) []int {
    var result []int
    for i := 0; i < n; i++ { result = append(result, i) }  // many reallocations
    return result
}

// ✅ Better: pre-allocate
func collectGood(n int) []int {
    result := make([]int, 0, n)  // pre-allocate capacity
    for i := 0; i < n; i++ { result = append(result, i) }
    return result
}

// ❌ Avoid: string concatenation in loop
for _, s := range words { result += s + " " }

// ✅ Better: strings.Builder
var sb strings.Builder
sb.Grow(estimatedSize)
for _, s := range words { sb.WriteString(s); sb.WriteByte(' ') }
result := sb.String()
```

---

### Q5. How do you detect goroutine leaks?

```go
// Goroutine leak: goroutine blocked forever, never returns

// ❌ Common leak patterns:
// 1. Sending to unbuffered channel with no receiver
go func() { ch <- value }()  // if no one reads ch, goroutine leaks

// 2. Range over channel that never closes
go func() {
    for v := range ch { process(v) }  // leaks if ch never closed
}()

// Detection:
import "runtime"
before := runtime.NumGoroutine()
doWork()
after := runtime.NumGoroutine()
if after > before { fmt.Println("goroutine leak:", after-before) }

// Use goleak package in tests
import "go.uber.org/goleak"
func TestNoGoroutineLeak(t *testing.T) {
    defer goleak.VerifyNone(t)
    // ... test code
}

// Fix: always use context for cancellation
func worker(ctx context.Context, ch <-chan int) {
    for {
        select {
        case <-ctx.Done():
            return  // clean exit
        case v, ok := <-ch:
            if !ok { return }
            process(v)
        }
    }
}
```

---

### Q6. How do you benchmark memory allocations?

```bash
# Run benchmarks with memory stats
go test -bench=BenchmarkMyFunc -benchmem ./...

# Output columns:
# ns/op     — nanoseconds per operation
# B/op      — bytes allocated per operation  ← watch this
# allocs/op — heap allocations per operation ← minimize this

# Escape analysis to understand where allocations come from
go build -gcflags="-m" ./...

# Trace allocations in detail
go test -trace trace.out ./...
go tool trace trace.out
```
