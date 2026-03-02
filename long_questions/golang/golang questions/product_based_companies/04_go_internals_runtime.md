# 🔬 04 — Go Internals & Runtime
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Go's GC (tri-color mark-and-sweep)
- Escape analysis
- The GMP scheduler in depth
- Stack vs heap allocation
- Memory model
- Interface representation in memory

---

## ❓ Most Asked Questions

### Q1. How does Go's Garbage Collector work?

```
Go uses a concurrent, tri-color mark-and-sweep GC:

Phases:
1. Mark Setup (STW ~microseconds)
   – Stop-the-world briefly to enable write barriers
   – Enable write barrier to track new allocations

2. Concurrent Mark
   – Goroutines continue running
   – GC marks all live objects using tri-color algorithm:
     • White: not yet visited (candidates for collection)
     • Gray: discovered but children not scanned
     • Black: fully scanned (reachable)
   – Write barrier ensures new allocations are tracked

3. Mark Termination (STW ~microseconds)
   – Final cleanup, disable write barriers

4. Sweep (concurrent)
   – Reclaim memory of white (unreachable) objects
   – Happens concurrently with user code

GC is triggered when: heap doubles in size (configurable with GOGC)
```

```go
import "runtime"
// Force GC
runtime.GC()
// Get memory stats
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Printf("HeapAlloc: %v MB\n", m.HeapAlloc/1024/1024)
fmt.Printf("NumGC: %v\n", m.NumGC)
```

---

### Q2. What is escape analysis?

```go
// Escape analysis determines where variables are allocated:
// - Stack: fast, automatically freed when function returns
// - Heap: slower, must be GC'd

// Variables ESCAPE to heap when:
// 1. Returned by pointer
func newUser() *User {
    u := User{Name: "Alice"}  // u escapes to heap
    return &u
}

// 2. Captured by a goroutine closure
func launch() {
    x := 42
    go func() { fmt.Println(x) }()  // x escapes
}

// 3. Interface boxing
var i interface{} = User{Name: "Alice"}  // may escape

// 4. Slice/map element when slice/map itself escapes

// Check escape analysis:
// go build -gcflags="-m" ./...
// go build -gcflags="-m -m" ./...  (more verbose)
```

---

### Q3. How does Go's scheduler work? Explain GMP in depth.

```
GMP: Goroutine, Machine (OS thread), Processor (logical CPU)

Structure:
- P has a local run queue (LRQ) of ~256 goroutines
- Global run queue (GRQ) for overflow
- When P's LRQ is empty, it steals from another P's LRQ (work stealing)
- M executes G's as directed by P

Scheduling triggers:
- Function calls (preemption checkpoints)
- System calls (M parks, P detaches and finds new M)
- channel operations (G may block/unblock)
- time.Sleep, sync.Mutex.Lock, etc.

Syscall handling:
- When G makes blocking syscall, M takes it into OS
- P detaches from M and finds another M (or creates one)
- When syscall returns, G re-enters scheduler
```

```go
// View goroutine state
import "runtime/debug"
debug.PrintStack()

// goroutine dump (in panic output)
// Shows: goroutine N [state]: ...
```

---

### Q4. What is the difference between Stack and Heap in Go?

| | Stack | Heap |
|--|-------|------|
| Allocation | Compile-time known size | Dynamic |
| Speed | Very fast (just move SP) | Slower (GC involved) |
| Lifetime | Function scope | Until GC |
| Size | ~1MB (configurable) | Limited by RAM |
| Thread-safety | Per-goroutine | Shared, needs sync |

```go
// Stack allocation (no GC pressure)
func sum(a, b int) int {
    result := a + b  // result on stack
    return result    // value copied on return
}

// Heap allocation (GC must track)
func newSlice(n int) []int {
    return make([]int, n)  // always on heap for large n
}

// Check with: go build -gcflags="-m" ./...
// Output: "./main.go:4:9: sum result does not escape"
//         "./main.go:8:14: make([]int, n) escapes to heap"
```

---

### Q5. How are interfaces represented in memory?

```go
// An interface value has two words:
// | type pointer | data pointer |
// (itab)         (value or pointer to value)

type Animal interface { Sound() string }
type Dog struct{ Name string }
func (d Dog) Sound() string { return "woof" }

var a Animal = Dog{Name: "Rex"}
// Memory: [*itab for (Dog, Animal)] [pointer to Dog value]

// nil interface: both words are nil
var a2 Animal  // nil — isNil = true

// non-nil interface with nil value
var d *Dog = nil
a3 := Animal(d)  // NOT nil — has type info, value is nil
// This is the "nil interface gotcha"!
```

---

### Q6. What are write barriers and why does Go use them?

```
Write barriers are hooks inserted by the compiler when pointer writes occur.
They notify the GC about pointer changes during concurrent mark phase.

Without write barriers:
- GC might miss objects created/moved after marking started
- Could free live objects → use-after-free bugs

Go's hybrid write barrier (since Go 1.14):
- Marks both the old and new value of a pointer write as grey
- Prevents GC from missing live objects

Performance impact: ~10-20% overhead during GC
```

---

### Q7. What are finalizers in Go and when should you use them?

```go
import "runtime"

type FileHandle struct { fd int }

func OpenFile(path string) *FileHandle {
    fh := &FileHandle{fd: openSyscall(path)}

    // Register cleanup when GC collects fh
    runtime.SetFinalizer(fh, func(f *FileHandle) {
        closeSyscall(f.fd)
        fmt.Println("file handle finalized")
    })
    return fh
}

// Caveats:
// - Not guaranteed to run promptly or at all (if program exits)
// - Cannot be used with types that contain cycles
// - Prefer explicit Close() pattern instead
// - Use for wrapping non-Go resources (file descriptors, C memory)
```

---

### Q8. How do you tune the GC with GOGC and GOMEMLIMIT?

```bash
# GOGC — target heap growth percentage (default 100 = double)
# Lower value = more frequent GC = less memory, more CPU
# Higher value = less frequent GC = more memory, less CPU

GOGC=50 ./app    # GC when heap grows 50% (aggressive)
GOGC=200 ./app   # GC when heap doubles twice (lazy)
GOGC=off ./app   # disable GC (latency-critical, manual runtime.GC())

# GOMEMLIMIT (Go 1.19+) — soft memory limit for the process
GOMEMLIMIT=500MiB ./app

# In code
import "runtime/debug"
debug.SetGCPercent(50)       // same as GOGC=50
debug.SetMemoryLimit(500 << 20)  // 500MB
```
