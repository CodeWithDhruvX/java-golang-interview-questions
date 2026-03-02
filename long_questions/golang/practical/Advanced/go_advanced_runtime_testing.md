# Go Advanced — Runtime Internals, GC, Profiling & Advanced Testing

> **Top Product Company Coverage — Part 1 of 2**
> Topics: GMP Scheduler · Goroutine Internals · GC & Memory · `runtime` package · `pprof` · `escape analysis` · Advanced Testing (`t.Parallel`, mocks, fuzz, benchmarks)

---

## Section 1: GMP Scheduler & Goroutine Internals (Q1–Q18)

### 1. GMP Model — Core Concept
**Q: What are G, M, and P in Go's runtime scheduler?**
```
G (Goroutine) — lightweight execution unit (~2KB stack, created by 'go')
M (Machine)   — OS thread; executes Go code; bounded by GOMAXPROCS
P (Processor) — logical CPU; holds run queues; bridges G and M

Relationship:
  Each M must hold a P to run Go code.
  Each P has a local run queue (LRQ) of Gs.
  A global run queue (GRQ) handles overflow.
  
  M ──holds──▶ P ──picks──▶ G ──executes on──▶ M
```
**A:** Go does NOT use 1:1 thread-per-goroutine. The GMP model uses M:N scheduling — M goroutines multiplexed over N OS threads via P processors. This is why millions of goroutines can run on 8 CPU cores.

---

### 2. GOMAXPROCS and Parallelism
**Q: What is the output and what does it demonstrate?**
```go
package main
import (
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)

func main() {
    // With GOMAXPROCS=1: sequential execution, no true parallelism
    runtime.GOMAXPROCS(1)
    var counter int64
    var wg sync.WaitGroup
    for i := 0; i < 10000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            atomic.AddInt64(&counter, 1)
        }()
    }
    wg.Wait()
    fmt.Println("GOMAXPROCS=1, counter:", counter)

    // With GOMAXPROCS=runtime.NumCPU(): true parallelism
    runtime.GOMAXPROCS(runtime.NumCPU())
    fmt.Println("NumCPU:", runtime.NumCPU())
}
```
**A:** `counter: 10000` in both cases (atomic ensures correctness). With `GOMAXPROCS=1`, goroutines are **concurrent** (interleaved) but not **parallel** (never simultaneously). With `NumCPU`, goroutines run **in parallel** on multiple cores.

---

### 3. Work Stealing
**Q: What is work stealing and why does it matter?**
```
When P's local run queue (LRQ) is empty:
  1. P checks the global run queue (GRQ)
  2. P steals half of another P's LRQ goroutines
  3. P checks the network poller for runnable Gs

Work Steal Example:
  P0 LRQ: [G1, G2, G3, G4]    P1 LRQ: []
  → P1 steals from P0
  P0 LRQ: [G1, G2]             P1 LRQ: [G3, G4]
```
**A:** Work stealing prevents CPU starvation. An idle P doesn't wait — it actively takes work from busy Ps. This distributes load automatically without programmer intervention.

---

### 4. Goroutine Preemption — Cooperative vs Async
**Q: What changed in Go 1.14 regarding goroutine preemption?**
```go
// Pre-Go 1.14: goroutines only preempted at function calls
// Tight loops could starve the scheduler:
func tightLoop() {
    for {
        // no function calls → scheduler never gets control (old behavior)
        i := 0
        i++
    }
}

// Go 1.14+: ASYNC PREEMPTION via signals (SIGURG on Unix)
// The runtime can preempt ANY goroutine at any safe point
// even in tight loops — no cooperation needed
```
**A:** Go 1.14 introduced **asynchronous preemption**. The runtime sends `SIGURG` to OS threads, which pauses goroutines at safe points. Before this, a tight CPU loop could starve other goroutines on the same P.

---

### 5. Goroutine States
**Q: What are the key goroutine states in the scheduler?**
```
_Grunnable  — ready to run, waiting in a P's run queue
_Grunning   — currently executing on an M
_Gwaiting   — blocked (channel, mutex, syscall, sleep)
_Gdead      — finished, resources being recycled
_Gcopystack — stack being copied/grown (STW briefly)

Transitions:
  go func() → _Grunnable (added to P's LRQ)
  scheduler picks it → _Grunning
  ch <- val (blocking) → _Gwaiting
  channel receives → _Grunnable
  func returns → _Gdead → goroutine struct recycled
```
**A:** Understanding goroutine states helps diagnose: goroutine leaks (`_Gwaiting` forever), scheduler starvation (`_Grunnable` not getting CPU time), and deadlocks (all Gs in `_Gwaiting`).

---

### 6. runtime.Gosched() — Yield
**Q: What does Gosched() do?**
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    go func() {
        fmt.Println("goroutine")
    }()
    runtime.Gosched() // yield: give scheduler a chance to run other goroutines
    fmt.Println("main")
}
```
**A:** `runtime.Gosched()` yields the current goroutine's timeslice, allowing other goroutines to run. Returns once the goroutine is rescheduled. Rarely needed in production — prefer `sync.WaitGroup` or channels.

---

### 7. runtime.LockOSThread() — Pin to OS Thread
**Q: When is LockOSThread necessary?**
```go
import "runtime"

func callCGoLibrary() {
    runtime.LockOSThread()   // pin this goroutine to its OS thread
    defer runtime.UnlockOSThread()

    // cgo call that requires thread-local state (OpenGL, GUI, etc.)
    // C library's thread-local storage is now guaranteed stable
    cLibrary.Draw()
}
```
**A:** Some C libraries (OpenGL, GUI toolkits) require calls from the same OS thread. `LockOSThread` pins the goroutine to its M permanently until `UnlockOSThread`. Used in CGO integrations, GUI loops, and thread-local storage scenarios.

---

### 8. Goroutine vs Thread — Quantitative Comparison
**Q: Fill in the comparison table:**
```
                  Goroutine         OS Thread
Stack size:       2KB (grows)       1-8MB (fixed)
Creation cost:    ~0.3μs            ~10μs
Context switch:   ~0.1μs (user)     ~1-10μs (kernel)
Max count:        millions          thousands
Scheduler:        Go runtime (GMP)  OS kernel
Communication:    channels, sync    mutex, semaphore
```
**A:** Goroutines are ~100x cheaper to create and context-switch than OS threads. This is why Go can handle millions of concurrent connections where other languages would exhaust OS thread limits.

---

### 9. Detecting Goroutine Count at Runtime
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func main() {
    fmt.Println("goroutines at start:", runtime.NumGoroutine())

    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            time.Sleep(50 * time.Millisecond)
        }()
    }
    fmt.Println("goroutines during:", runtime.NumGoroutine())
    wg.Wait()
    fmt.Println("goroutines after:", runtime.NumGoroutine())
}
```
**A:** `1 → 101 → 1`. `NumGoroutine()` counts all alive goroutines. High unexpected counts signal goroutine leaks. Monitor this in production via `/debug/vars` or Prometheus.

---

### 10. Stack Growth — Segmented vs Contiguous
**Q: What changed in Go 1.4 and why does it matter?**
```
Go 1.3 and earlier — SEGMENTED STACKS:
  Stack grows by adding new stack segment.
  Problem: "hot split" — function at segment boundary 
  causes expensive growth/shrink repeatedly.

Go 1.4+ — CONTIGUOUS (COPYING) STACKS:
  When stack overflows → allocate 2x larger stack,
  copy ALL stack frames to new memory,
  update all pointers to new locations.
  
  Default initial: 8KB (was 4KB)
  Max default:     1GB (configurable via GOTRACEBACK)
```
**A:** Copying stacks eliminated the hot-split performance problem. The tradeoff: pointers into the stack become invalid during copy — this is why you cannot store goroutine stack addresses in Go.

---

### 11. Network Poller — Non-blocking I/O
**Q: How does Go handle blocking network syscalls without blocking the OS thread?**
```
When goroutine makes a network call (net.Conn.Read):
  1. Go runtime converts to non-blocking syscall
  2. If not ready → goroutine parked in _Gwaiting state
  3. OS thread (M) released → picks up another goroutine (G)
  4. epoll/kqueue/IOCP polls FDs in background
  5. When data ready → goroutine moved back to _Grunnable
  6. Goroutine resumes from where it blocked

Benefits:
  - No OS thread wasted waiting for I/O
  - Thousands of network connections per thread
  - Code looks synchronous but is async underneath
```
**A:** This is Go's key advantage for I/O-heavy services. You write `conn.Read(buf)` synchronously but the runtime transparently parks the goroutine and reuses the OS thread.

---

### 12. Syscall Handling — Thread Detachment
**Q: What happens when a goroutine makes a blocking syscall (not network)?**
```
Blocking syscall (os.File.Read on disk, CGO):
  1. Go runtime detaches M from P before syscall
  2. P remains available — picks up another M (or creates new M)
  3. M is blocked in kernel — cannot run other goroutines
  4. When syscall returns → M tries to reacquire a P
     - If P available → resume normally
     - If no P available → goroutine goes to _Grunnable queue
     - M may park and wait or be reused

This is why: GOMAXPROCS controls P count (logical CPUs),
but there can be MORE OS threads (M) than GOMAXPROCS
if goroutines are blocked in syscalls.
```
**A:** Go creates additional OS threads for blocking syscalls to avoid starving the P. Thread count is bounded by `runtime/debug.SetMaxThreads` (default: 10,000).

---

### 13. internal/race — Race Detector Internals
**Q: What is the output and how does -race work?**
```go
// go run -race main.go
package main

import "fmt"

func main() {
    c := make(chan bool)
    m := make(map[string]string)
    go func() {
        m["key"] = "goroutine write" // concurrent unsynchronized write
        c <- true
    }()
    m["key"] = "main write"          // concurrent unsynchronized write
    <-c
    fmt.Println(m["key"])
}
```
**A:** With `-race`: prints `WARNING: DATA RACE` with goroutine stacks. The race detector uses **ThreadSanitizer (TSan)** shadow memory to track every memory access. 5-10x slowdown; ~5-10x memory overhead. Always run tests with `-race` before release.

---

### 14. runtime.Stack — Goroutine Stack Dump
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "runtime"
)

func inner() string {
    buf := make([]byte, 4096)
    n := runtime.Stack(buf, false) // false = current goroutine only
    return string(buf[:n])
}

func outer() { fmt.Println(inner()) }

func main() { outer() }
```
**A:** Prints the current goroutine's stack trace like:
```
goroutine 1 [running]:
main.inner(...)
main.outer(...)
main.main(...)
```
`runtime.Stack(buf, true)` dumps ALL goroutines — essential for diagnosing leaks and deadlocks.

---

### 15. Finalizers — runtime.SetFinalizer
**Q: What does this do and what are the caveats?**
```go
package main
import (
    "fmt"
    "runtime"
)

type Resource struct{ name string }

func main() {
    r := &Resource{name: "file"}
    runtime.SetFinalizer(r, func(res *Resource) {
        fmt.Println("finalizing:", res.name)
    })
    r = nil              // r unreachable
    runtime.GC()         // force GC — finalizer may run
    runtime.Gosched()    // give GC goroutines a chance
}
```
**A:** May print `finalizing: file`. Caveats: (1) finalizer runs in a separate goroutine, (2) ordering not guaranteed, (3) delays GC of object by one cycle, (4) NOT guaranteed to run before program exits. **Prefer `defer Close()`** over finalizers.

---

### 16. runtime.KeepAlive — Prevent Premature Finalization
**Q: When is KeepAlive needed?**
```go
import "runtime"

func process(fd *File) {
    p := fd.pointer()  // raw pointer from fd, used in C code
    cLibrary.UsePointer(p)
    // Without KeepAlive: GC may collect fd here (it's "unused" in Go)
    // and run finalizer (closing the file) while C code uses p
    runtime.KeepAlive(fd) // fd remains live until this point
}
```
**A:** When a Go object's finalizer could run prematurely (before CGO finishes using its raw pointer), `runtime.KeepAlive` tells the GC: "keep this object alive at least until this point." Critical in CGO code.

---

### 17. runtime.GOARCH and GOOS
**Q: What is the output and what is the use case?**
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.GOOS)    // "linux", "windows", "darwin"
    fmt.Println(runtime.GOARCH)  // "amd64", "arm64", "386"
    fmt.Println(runtime.Version()) // "go1.21.0"
}
```
**A:** Output varies by platform. Used for: conditional compilation (`//go:build linux`), platform-specific code paths, logging in production diagnostics, and build scripts.

---

### 18. Goroutine Leak Detection Pattern
**Q: How do you detect goroutine leaks in tests?**
```go
package mypackage_test

import (
    "runtime"
    "testing"
    "time"
)

func TestNoGoroutineLeak(t *testing.T) {
    before := runtime.NumGoroutine()

    // Run function under test
    stopCh := make(chan struct{})
    go startWorker(stopCh) // function being tested
    time.Sleep(10 * time.Millisecond)
    close(stopCh) // trigger shutdown
    time.Sleep(50 * time.Millisecond) // allow goroutine to exit

    after := runtime.NumGoroutine()
    if after > before {
        t.Errorf("goroutine leak: before=%d, after=%d", before, after)
    }
}
```
**A:** Compare `NumGoroutine()` before and after. The `goleak` package (`go.uber.org/goleak`) automates this — widely used at Uber, Google, Stripe.

---

## Section 2: Garbage Collector Internals (Q19–Q30)

### 19. Go GC — Tri-Color Mark-and-Sweep
**Q: Explain the algorithm:**
```
Go uses Concurrent Tri-Color Mark-and-Sweep GC:

Colors:
  White — not yet visited (candidates for collection)
  Grey  — discovered, children not yet scanned
  Black — fully scanned (reachable, not a candidate)

Phase 1: MARK — concurrent with mutators (your code)
  1. Start: all objects White, roots (globals, stacks) → Grey
  2. For each Grey object: scan children → Grey,
     then mark object Black
  3. Repeat until no Grey objects remain
  4. All remaining White objects = unreachable = garbage

Phase 2: SWEEP — concurrent
  Free memory occupied by White objects.

Write Barrier (Dijkstra + Yuasa hybrid):
  When your code modifies a pointer during marking,
  the write barrier ensures new pointers are greyed
  → prevents GC from missing live objects.
```
**A:** Go's GC is mostly concurrent (runs alongside your code) with two short STW (stop-the-world) pauses: one at mark start and one at mark termination. Target: < 1ms STW latency.

---

### 20. GC Trigger and GOGC
**Q: What controls when the GC runs?**
```go
package main
import (
    "fmt"
    "runtime"
    "runtime/debug"
)

func main() {
    // GOGC=100 (default): GC triggers when heap doubles
    // GOGC=200: GC triggers when heap triples (less frequent, more memory)
    // GOGC=off: disable GC entirely
    // GOGC=50:  GC triggers at 50% growth (more frequent, less memory)

    fmt.Println("GOGC default:", debug.SetGCPercent(-1)) // returns old value
    debug.SetGCPercent(100) // restore

    // Force a GC cycle:
    runtime.GC()

    var stats runtime.MemStats
    runtime.ReadMemStats(&stats)
    fmt.Printf("HeapAlloc: %dKB\n", stats.HeapAlloc/1024)
    fmt.Printf("NumGC: %d\n", stats.NumGC)
}
```
**A:** `GOGC=100` (default) means GC runs when live heap grows by 100% from last collection. Use `GOGC=off` for latency-sensitive tasks (then call `runtime.GC()` manually). Go 1.19+ adds `GOMEMLIMIT` for soft memory cap.

---

### 21. GOMEMLIMIT (Go 1.19+)
**Q: What does GOMEMLIMIT solve?**
```go
import "runtime/debug"

func main() {
    // GOMEMLIMIT=512MiB (env var) OR:
    debug.SetMemoryLimit(512 * 1024 * 1024) // 512MB soft limit

    // With GOMEMLIMIT: GC becomes more aggressive when approaching
    // the limit, preventing OOM kills even if GOGC=100 would allow
    // heap to grow beyond limit.
    //
    // Production pattern: set GOMEMLIMIT to ~90% of container memory limit
    // GOMEMLIMIT=900MiB with a 1GB container limit
}
```
**A:** Before `GOMEMLIMIT`, a container with 1GB RAM limit would be OOM-killed if heap grew to 1.5GB between GC cycles. `GOMEMLIMIT` tells the GC to run more aggressively BEFORE hitting the OS limit.

---

### 22. Escape Analysis — heap vs stack
**Q: Which variables escape to heap?**
```go
package main
import "fmt"

// Run: go build -gcflags="-m" to see escape analysis output

func stackAlloc() int {
    x := 42      // stays on stack — doesn't escape
    return x     // value copied out, x dies with function
}

func heapAlloc() *int {
    x := 42      // ESCAPES to heap: address returned
    return &x    // go build -m: "./main.go: x escapes to heap"
}

func interfaceEscape(v interface{}) {
    // v's underlying value escapes: interface boxing forces heap alloc
    fmt.Println(v) // fmt.Println takes interface{} — reflection needed → heap
}

func main() {
    n := stackAlloc()
    p := heapAlloc()
    fmt.Println(n, *p)
}
```
**A:** A value escapes to heap when: (1) its address is returned/stored, (2) it's stored in an interface, (3) it's captured by a closure and outlives the frame, (4) it's too large for the stack. Use `-gcflags="-m"` to inspect.

---

### 23. Reducing GC Pressure — Object Pooling
**Q: What is the output and why is sync.Pool useful?**
```go
package main
import (
    "fmt"
    "sync"
)

var bufPool = sync.Pool{
    New: func() interface{} {
        fmt.Println("allocating new buffer")
        return make([]byte, 4096)
    },
}

func process() {
    buf := bufPool.Get().([]byte) // reuse or allocate
    defer bufPool.Put(buf)        // return to pool after use
    // use buf...
    _ = buf
}

func main() {
    process()
    process() // second call: reuses buffer from pool
    process()
}
```
**A:** `allocating new buffer` printed once (or more depending on GC). `sync.Pool` caches objects between GC cycles. GC clears the pool — objects are not permanent. Reduces GC pressure for short-lived, frequently allocated objects (buffers, parsers, decoders).

---

### 24. sync.Pool Caveats
**Q: What are the constraints on sync.Pool?**
```go
// WRONG: don't cache state between requests in Pool
// Pool objects can be evicted by GC at any time

// CORRECT usage pattern:
var jsonDecoderPool = sync.Pool{
    New: func() interface{} { return new(json.Decoder) },
}

func decodeRequest(r io.Reader) (*Request, error) {
    dec := jsonDecoderPool.Get().(*json.Decoder)
    *dec = *json.NewDecoder(r)  // reset state before use!
    defer jsonDecoderPool.Put(dec)
    var req Request
    return &req, dec.Decode(&req)
}
```
**A:** Pool constraints: (1) GC clears pool between cycles — not a persistent cache, (2) objects can be accessed by any goroutine — don't store goroutine-local state, (3) **always reset state** before reuse — pooled objects carry previous use's data.

---

### 25. runtime.MemStats — Memory Profiling
**Q: What do the key MemStats fields mean?**
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)

    fmt.Printf("HeapAlloc:    %dKB  (live heap objects)\n",   m.HeapAlloc/1024)
    fmt.Printf("HeapSys:      %dKB  (heap reserved from OS)\n", m.HeapSys/1024)
    fmt.Printf("HeapIdle:     %dKB  (unused, returned to OS)\n", m.HeapIdle/1024)
    fmt.Printf("HeapInuse:    %dKB  (actively used)\n",        m.HeapInuse/1024)
    fmt.Printf("NumGC:        %d    (total GC cycles)\n",      m.NumGC)
    fmt.Printf("PauseTotalNs: %dμs  (total STW time)\n",       m.PauseTotalNs/1000)
    fmt.Printf("Mallocs:      %d    (total allocations)\n",    m.Mallocs)
    fmt.Printf("Frees:        %d    (total freed)\n",          m.Frees)
}
```
**A:** `HeapAlloc` is the most watched metric — current live heap. High `Mallocs - Frees` indicates accumulation. High `PauseTotalNs` signals GC pressure. Export these to Prometheus in production.

---

### 26. Heap Profiling with pprof
**Q: What is the standard pattern to expose pprof in a service?**
```go
package main
import (
    "log"
    "net/http"
    _ "net/http/pprof" // side-effect import registers /debug/pprof/* handlers
)

func main() {
    // Existing server:
    http.HandleFunc("/api", apiHandler)

    // pprof served at:
    // /debug/pprof/           — index
    // /debug/pprof/heap       — heap profile
    // /debug/pprof/goroutine  — goroutine dump
    // /debug/pprof/cpu        — CPU profile (30-second sample)
    // /debug/pprof/trace      — execution trace
    log.Fatal(http.ListenAndServe(":6060", nil))
}

// Collect: go tool pprof http://localhost:6060/debug/pprof/heap
// Visualize: go tool pprof -http=:8080 profile.pb.gz
```
**A:** Blank import `_ "net/http/pprof"` auto-registers all profiling endpoints. In production: expose on a separate internal port (not public). This pattern is used at every major Go company.

---

### 27. CPU Profiling in Tests
**Q: What does this collect?**
```go
package main_test

import (
    "os"
    "runtime/pprof"
    "testing"
)

func BenchmarkHeavyFunc(b *testing.B) {
    // Enable CPU profiling for this benchmark:
    // go test -bench=. -cpuprofile=cpu.prof
    // go tool pprof cpu.prof
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        heavyFunc()
    }
}

func TestWithProfile(t *testing.T) {
    f, _ := os.Create("cpu.prof")
    defer f.Close()
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    
    // code under measurement
    heavyFunc()
}
```
**A:** CPU profile samples goroutine stacks every 100μs, showing which functions consume the most CPU. `go tool pprof` → `top10` shows hottest functions. `web` opens a flame graph.

---

### 28. Allocation Profiling — inuse_objects vs alloc_objects
**Q: What is the difference?**
```bash
# inuse_objects (default): objects currently in heap
go tool pprof http://localhost:6060/debug/pprof/heap

# alloc_objects: all allocations since program start (even freed)
go tool pprof -alloc_objects http://localhost:6060/debug/pprof/heap
```
**A:** `inuse_objects` → find **memory leaks** (objects that should have been freed). `alloc_objects` → find **GC pressure** (objects that ARE freed but allocated so frequently they stress GC). Both are critical for optimization.

---

### 29. go tool trace — Execution Tracing
**Q: What does the execution tracer reveal?**
```go
package main
import (
    "os"
    "runtime/trace"
    "fmt"
)

func main() {
    f, _ := os.Create("trace.out")
    defer f.Close()
    trace.Start(f)
    defer trace.Stop()

    // Code under analysis
    ch := make(chan int)
    go func() { ch <- 42 }()
    fmt.Println(<-ch)
}
// Visualize: go tool trace trace.out
```
**A:** `go tool trace` shows: goroutine scheduling latency, GC events, network/syscall wait times, goroutine creation/destruction, and P utilization over time. Essential for diagnosing scheduler hiccups and latency spikes.

---

### 30. Benchmark Memory Reporting
**Q: What is the output and what do B/op and allocs/op mean?**
```go
package mypackage

import "testing"

func BenchmarkConcat(b *testing.B) {
    b.ReportAllocs() // enable memory reporting
    for i := 0; i < b.N; i++ {
        s := ""
        for j := 0; j < 100; j++ {
            s += "x" // allocates new string each iteration
        }
        _ = s
    }
}

// Run: go test -bench=BenchmarkConcat -benchmem
// Output:
// BenchmarkConcat-8    10000    95000 ns/op    4944 B/op    99 allocs/op
```
**A:** `B/op` = bytes allocated per operation. `allocs/op` = number of heap allocations per operation. Use `strings.Builder` to reduce allocations. Optimizing allocations is the most impactful Go performance work.

---

## Section 3: Advanced Testing (Q31–Q45)

### 31. t.Parallel() — Parallel Tests
**Q: What does t.Parallel() do?**
```go
package mypackage_test

import "testing"

func TestA(t *testing.T) {
    t.Parallel() // this test can run with other parallel tests
    // ... test A logic
}

func TestB(t *testing.T) {
    t.Parallel()
    // ... test B logic
}

// Run: go test -parallel 4 ./...
// Both TestA and TestB run concurrently
```
**A:** `t.Parallel()` signals that this test can run concurrently with other parallel tests. The test pauses at `t.Parallel()` call until all non-parallel tests complete, then resumes. Speeds up I/O-bound test suites significantly.

---

### 32. Table-Driven with t.Run — Subtests
**Q: What is the output structure?**
```go
package mypackage_test

import "testing"

func divide(a, b float64) float64 {
    if b == 0 { return 0 }
    return a / b
}

func TestDivide(t *testing.T) {
    tests := []struct {
        name string
        a, b, want float64
    }{
        {"positive", 10, 2, 5},
        {"zero divisor", 10, 0, 0},
        {"negative", -6, 2, -3},
    }
    for _, tc := range tests {
        tc := tc // capture for parallel
        t.Run(tc.name, func(t *testing.T) {
            t.Parallel()
            got := divide(tc.a, tc.b)
            if got != tc.want {
                t.Errorf("divide(%v,%v) = %v; want %v", tc.a, tc.b, got, tc.want)
            }
        })
    }
}
```
**A:** `t.Run` creates subtests: `TestDivide/positive`, `TestDivide/zero_divisor`, `TestDivide/negative`. Run individual: `go test -run TestDivide/positive`. `tc := tc` captures the loop variable for parallel safety (pre Go 1.22).

---

### 33. TestMain — Test Suite Setup/Teardown
**Q: When do you use TestMain?**
```go
package mypackage_test

import (
    "fmt"
    "os"
    "testing"
)

func TestMain(m *testing.M) {
    // Setup: runs ONCE before all tests
    fmt.Println("setting up test suite")
    db := setupTestDB()

    // Run all tests
    code := m.Run()

    // Teardown: runs ONCE after all tests
    fmt.Println("tearing down test suite")
    db.Close()

    os.Exit(code) // must call os.Exit with the test result code
}
```
**A:** `TestMain` wraps the entire test suite. Used for: starting test databases/servers, loading test fixtures, configuring global state. **Must call `os.Exit(m.Run())`** — missing `os.Exit` causes deferred teardown to block and tests to report wrong exit code.

---

### 34. Interface Mocking with Testify
**Q: What is the standard mock pattern?**
```go
package service_test

import (
    "testing"
    "github.com/stretchr/testify/mock"
)

type UserRepo interface {
    FindByID(id int) (*User, error)
}

type MockUserRepo struct { mock.Mock }

func (m *MockUserRepo) FindByID(id int) (*User, error) {
    args := m.Called(id)
    return args.Get(0).(*User), args.Error(1)
}

func TestGetUser(t *testing.T) {
    mockRepo := &MockUserRepo{}
    mockRepo.On("FindByID", 42).Return(&User{Name: "Alice"}, nil)

    svc := NewUserService(mockRepo)
    user, err := svc.GetUser(42)

    mockRepo.AssertExpectations(t)  // verifies all expected calls were made
    assert.NoError(t, err)
    assert.Equal(t, "Alice", user.Name)
}
```
**A:** Testify mocks implement interfaces and record/verify calls. `mockRepo.AssertExpectations(t)` fails the test if expected calls weren't made.

---

### 35. Fuzz Testing (Go 1.18+)
**Q: What does this fuzz test do?**
```go
package mypackage_test

import (
    "testing"
    "unicode/utf8"
)

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func FuzzReverse(f *testing.F) {
    // Seed corpus
    f.Add("hello")
    f.Add("世界")
    f.Add("")

    f.Fuzz(func(t *testing.T, s string) {
        rev := reverseString(s)
        // Property: double-reverse must equal original
        if reverseString(rev) != s {
            t.Errorf("double reverse failed for %q", s)
        }
        // Property: result must be valid UTF-8
        if !utf8.ValidString(rev) {
            t.Errorf("reversed string is invalid UTF-8")
        }
    })
}

// Run: go test -fuzz=FuzzReverse -fuzztime=30s
```
**A:** Fuzz testing generates random inputs from the seed corpus, trying to find inputs that violate the stated properties. Found crashes are saved as corpus files for regression. Google uses fuzzing extensively.

---

### 36. Benchmark with b.SetBytes
**Q: What is the output and what does MB/s mean?**
```go
package mypackage_test

import (
    "testing"
    "strings"
)

func BenchmarkBuilder(b *testing.B) {
    const size = 1000
    b.SetBytes(int64(size)) // bytes processed per operation
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        var sb strings.Builder
        for j := 0; j < size; j++ {
            sb.WriteByte('x')
        }
        _ = sb.String()
    }
}
// Output: BenchmarkBuilder-8  200000  6000 ns/op  166.7 MB/s  1024 B/op  3 allocs/op
```
**A:** `b.SetBytes(n)` enables throughput reporting (MB/s = n*1e9/ns_per_op). Use when benchmarking I/O-like operations (parsing, serialization, hashing) to express performance in terms of data throughput.

---

### 37. Golden File Testing
**Q: What is the golden file pattern?**
```go
package mypackage_test
import (
    "os"
    "testing"
)

func TestHTMLOutput(t *testing.T) {
    got := renderHTML(inputData)
    goldenFile := "testdata/output.golden"

    if *update { // go test -update flag
        os.WriteFile(goldenFile, []byte(got), 0644)
        return
    }

    want, _ := os.ReadFile(goldenFile)
    if got != string(want) {
        t.Errorf("output mismatch:\ngot:  %s\nwant: %s", got, want)
    }
}
```
**A:** Golden files store expected output as files in `testdata/`. Run `go test -update` to regenerate them. Used for: HTML rendering, code generation, CLI output tests where expected output is large or complex.

---

### 38. t.Cleanup — Register Teardown Per Test
**Q: Why use t.Cleanup instead of defer?**
```go
func TestWithTempDir(t *testing.T) {
    dir, err := os.MkdirTemp("", "test-*")
    if err != nil {
        t.Fatal(err)
    }
    t.Cleanup(func() {
        os.RemoveAll(dir) // runs after test AND all subtests complete
    })

    t.Run("subtest", func(t *testing.T) {
        // dir still exists here
        _ = dir
    })
    // t.Cleanup runs AFTER subtest returns
}
```
**A:** `t.Cleanup` is preferred over `defer` in tests because it runs after ALL subtests complete, not just when the current `t` scope exits. Essential for shared resources used by subtests.

---

### 39. httptest Package — Testing HTTP Handlers
**Q: What is the output?**
```go
package api_test
import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello, " + name))
}

func TestGreetHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/greet?name=Go", nil)
    rec := httptest.NewRecorder()

    greetHandler(rec, req)

    if rec.Code != http.StatusOK {
        t.Errorf("expected 200, got %d", rec.Code)
    }
    if rec.Body.String() != "Hello, Go" {
        t.Errorf("unexpected body: %q", rec.Body.String())
    }
}
```
**A:** `httptest.NewRequest` and `httptest.NewRecorder` let you test HTTP handlers without starting a real server. `ResponseRecorder` captures status code, headers, and body for assertion.

---

### 40. Benchmark Comparison — go test -benchstat
**Q: What is the workflow for benchmark comparison?**
```bash
# Before optimization:
go test -bench=BenchmarkProcess -count=5 ./... > before.txt

# Apply optimization, then:
go test -bench=BenchmarkProcess -count=5 ./... > after.txt

# Compare with benchstat:
go install golang.org/x/perf/cmd/benchstat@latest
benchstat before.txt after.txt

# Output:
# name             old time/op  new time/op  delta
# BenchmarkProcess  95μs ± 2%   40μs ± 1%   -57.89% (p=0.008 n=5+5)
```
**A:** `benchstat` computes statistical significance of benchmark improvements. `-count=5` runs each benchmark 5 times to reduce noise. This is the professional workflow for performance optimization at Google/Uber.

---

### 41. testify/assert vs require
**Q: What is the difference?**
```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestUser(t *testing.T) {
    user, err := getUser(42)

    require.NoError(t, err)       // STOPS test on failure (like t.Fatal)
    require.NotNil(t, user)       // STOPS test if nil — safe to dereference below

    assert.Equal(t, "Alice", user.Name)  // CONTINUES on failure (like t.Error)
    assert.Equal(t, 30, user.Age)        // reports all failures
}
```
**A:** `require` = stop immediately (like `t.Fatal`). `assert` = continue collecting failures (like `t.Error`). Use `require` for preconditions (nil checks, error checks) before assertions that would panic on nil.

---

### 42. go test -cover and coverprofile
**Q: What is the standard coverage workflow?**
```bash
# Run tests with coverage:
go test -coverprofile=coverage.out ./...

# View coverage percentage:
go tool cover -func=coverage.out

# Open visual HTML report:
go tool cover -html=coverage.out

# Fail if coverage below threshold (CI):
go test -coverprofile=coverage.out ./... && \
  go tool cover -func=coverage.out | grep total | awk '{print $3}' | \
  awk -F'%' '{if ($1 < 80) exit 1}'
```
**A:** Coverage profiles identify untested code paths. HTML report highlights red (untested) and green (covered) lines. 80%+ coverage is a common CI gate. Integration tests use `-coverpkg=./...` to include indirect package coverage.

---

### 43. Dependency Injection for Testability
**Q: Which design is more testable?**
```go
// BAD: hardcoded dependency — cannot mock in tests
type OrderService struct{}
func (s *OrderService) Create(item string) error {
    db := sql.Open("postgres", os.Getenv("DB_URL")) // cannot substitute
    return db.Exec("INSERT INTO orders ...")
}

// GOOD: dependency injected via interface
type DB interface {
    Exec(query string, args ...interface{}) (sql.Result, error)
}

type OrderService struct { db DB }

func NewOrderService(db DB) *OrderService { return &OrderService{db: db} }

func (s *OrderService) Create(item string) error {
    return s.db.Exec("INSERT INTO orders (item) VALUES (?)", item)
}
// In tests: inject mock DB. In production: inject *sql.DB.
```
**A:** Dependency injection via interfaces is the core Go testability pattern. The interface `DB` is defined where it's used (not where it's implemented), keeping coupling minimal.

---

### 44. Build Tags for Integration Tests
**Q: How do you separate unit and integration tests?**
```go
//go:build integration

package mypackage_test

import "testing"

// Only compiled when: go test -tags=integration ./...
// NOT compiled in regular: go test ./...

func TestRealDatabase(t *testing.T) {
    // Test that requires a real running database
    db := connectToRealDB(t)
    defer db.Close()
    // ... real DB tests
}
```
**A:** Build tags (`//go:build integration`) keep slow/infrastructure-dependent tests separate. CI pipeline runs `go test ./...` for fast unit tests and `go test -tags=integration` for slower integration tests. Note: `//go:build` replaces `// +build` syntax (Go 1.17+).

---

### 45. Testcontainers Pattern — Real Dependencies in Tests
**Q: What does this pattern provide?**
```go
package db_test
import (
    "context"
    "testing"
    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/modules/postgres"
)

func TestWithRealPostgres(t *testing.T) {
    ctx := context.Background()
    
    // Spin up a real PostgreSQL container
    pgContainer, err := postgres.RunContainer(ctx,
        testcontainers.WithImage("postgres:15"),
        postgres.WithDatabase("testdb"),
    )
    if err != nil { t.Fatal(err) }
    defer pgContainer.Terminate(ctx)

    connStr, _ := pgContainer.ConnectionString(ctx)
    db := connectDB(t, connStr)
    
    // Test against real Postgres — no mocks
    runMigrations(db)
    testOrderCreation(t, db)
}
```
**A:** Testcontainers starts real Docker containers for tests. Used at Uber, Stripe, and major Go shops for high-confidence integration tests without a persistent test infrastructure.

---

*End of Part 1 — Runtime Internals, GC, Profiling & Advanced Testing (45 questions)*
*See `go_advanced_system_design_patterns.md` for Part 2: System Design Pattern Practical Coding*
