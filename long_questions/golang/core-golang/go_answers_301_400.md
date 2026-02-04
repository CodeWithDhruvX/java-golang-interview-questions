## 🟢 Go Internals and Runtime (Questions 301-320)

### Question 301: How does the Go scheduler work?

**Answer:**
Go uses an M:N scheduler that multiplexes M goroutines onto N OS threads:

**Key components:**
```go
// G - Goroutine
// M - Machine (OS thread)
// P - Processor (context for execution)
```

**Scheduling model:**
```
┌─────────────────────────────────────┐
│  Global Run Queue                   │
└─────────────────────────────────────┘
         │
         ├──> P0 ──> M0 ──> OS Thread
         │    │
         │    └──> Local Run Queue: [G1, G2, G3]
         │
         ├──> P1 ──> M1 ──> OS Thread
         │    │
         │    └──> Local Run Queue: [G4, G5]
         │
         └──> P2 ──> M2 ──> OS Thread
              │
              └──> Local Run Queue: [G6, G7, G8]
```

**How it works:**
```go
// 1. Each P has a local queue of goroutines
// 2. M runs goroutines from P's local queue
// 3. If local queue is empty, M steals from other Ps
// 4. GOMAXPROCS controls number of Ps

import "runtime"

func main() {
    // Set number of processors:
    runtime.GOMAXPROCS(4)
    
    // Create goroutines:
    for i := 0; i < 1000; i++ {
        go func(id int) {
            // Work distributed across Ps
            doWork(id)
        }(i)
    }
}
```

**Preemption:**
```go
// Cooperative preemption at function calls
// Preemptive preemption (Go 1.14+) for long-running goroutines

func infiniteLoop() {
    for {
        // Go 1.14+: Can be preempted even without function calls
    }
}
```

---

### Question 302: What is M:N scheduling in Golang?

**Answer:**
M:N scheduling maps M goroutines to N OS threads:

**Advantages:**
- Lightweight goroutines (2KB stack)
- Efficient context switching
- Better CPU utilization

**Example:**
```go
// Create 10,000 goroutines on just a few OS threads:
for i := 0; i < 10000; i++ {
    go func(id int) {
        time.Sleep(time.Second)
        fmt.Println(id)
    }(i)
}

// Check thread count:
fmt.Printf("OS Threads: %d\n", runtime.NumCPU())
fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
```

**Work stealing:**
```go
// If P1's queue is empty, it steals from P2:
P1: []           P2: [G1, G2, G3, G4]
              Steal →
P1: [G4, G3]     P2: [G1, G2]
```

**Blocking operations:**
```go
// Syscall blocks M, not P:
func readFile() {
    data, _ := os.ReadFile("file.txt")  // M blocks
    // P is handed off to another M
    process(data)
}
```

---

### Question 303: How does the Go garbage collector work?

**Answer:**
Go uses a concurrent, tri-color mark-and-sweep GC:

**Phases:**
```go
// 1. Mark Setup (STW - Stop The World)
//    - Enables write barrier
//    - Scans stack roots

// 2. Concurrent Mark
//    - Application runs while marking reachable objects
//    - Uses write barrier to track new allocations

// 3. Mark Termination (STW)
//    - Rescan, finalize marking

// 4. Concurrent Sweep
//    - Reclaim unmarked objects
```

**Tri-color marking:**
```go
White: Not yet scanned
Gray:  Scanned, but references not scanned
Black: Scanned along with all references

// Process:
1. All objects start white
2. Roots (globals, stack) → gray
3. Gray objects → scan → black
4. Gray references → gray
5. When no gray objects remain, white objects are garbage
```

**Monitor GC:**
```go
import "runtime"

func monitorGC() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("GC runs: %d\n", m.NumGC)
    fmt.Printf("Pause total: %v\n", time.Duration(m.PauseTotalNs))
    fmt.Printf("Heap alloc: %v MB\n", m.Alloc/1024/1024)
}

// Enable GC trace:
// GODEBUG=gctrace=1 go run main.go
```

**Tuning:**
```go
import "runtime/debug"

// Set GC percentage (default 100):
debug.SetGCPercent(200)  // Less frequent GC

// Set memory limit (Go 1.19+):
debug.SetMemoryLimit(8 * 1024 * 1024 * 1024)  // 8GB
```

---

### Question 304: What are STW (stop-the-world) events in GC?

**Answer:**
STW pauses all application goroutines during GC:

**When STW occurs:**
```go
1. Mark setup phase
2. Mark termination phase
```

**Measure STW pauses:**
```bash
GODEBUG=gctrace=1 ./app

# Output:
# gc 45 @10.012s 0%: 0.025+0.78+0.019 ms clock
#                    ^STW   ^Mark ^STW
```

**Monitor pauses:**
```go
import "runtime"

func trackGCPauses() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    // Last 256 GC pauses:
    for i := 0; i < int(m.NumGC); i++ {
        pause := m.PauseNs[(i+255)%256]
        if pause > 1000000 {  // > 1ms
            fmt.Printf("Long pause: %v ms\n", pause/1000000)
        }
    }
}
```

**Reduce STW time:**
```go
// 1. Reduce heap size
// 2. Increase GOGC
// 3. Use sync.Pool
// 4. Avoid global pointers

// Example - avoid pointer-heavy structs:
// BAD:
type Data struct {
    Values []*int  // Many pointers to scan
}

// GOOD:
type Data struct {
    Values []int   // No pointers, faster GC
}
```

---

### Question 305: How are goroutines implemented under the hood?

**Answer:**
Goroutines are lightweight user-space threads:

**Structure:**
```go
type g struct {
    stack       stack   // Stack memory
    stackguard0 uintptr // Stack overflow detection
    m           *m      // Current M (thread)
    sched       gobuf   // Scheduling info (PC, SP)
    atomicstatus uint32 // State: running, runnable, waiting
    goid        int64   // Goroutine ID
}
```

**Stack:**
```go
// Starts with 2KB stack (vs 1MB for OS thread)
// Grows dynamically via stack copying

func recursiveFunc(n int) {
    if n == 0 {
        return
    }
    var buf [1024]byte  // Stack grows if needed
    recursiveFunc(n - 1)
}
```

**States:**
```go
_Gidle      // Just allocated
_Grunnable  // Ready to run
_Grunning   // Executing
_Gsyscall   // In syscall
_Gwaiting   // Blocked (channel, lock, etc.)
_Gdead      // Finished execution
```

**Creation overhead:**
```go
func BenchmarkGoroutine(b *testing.B) {
    for i := 0; i < b.N; i++ {
        done := make(chan bool)
        go func() {
            done <- true
        }()
        <-done
    }
}
// ~1-2 microseconds per goroutine
```

---

### Question 306: How does stack growth work in Go?

**Answer:**
Go stacks grow dynamically using stack copying:

**Initial size:**
```go
// Each goroutine starts with 2KB stack (Go 1.4+)
// Previously: segmented stacks (hot split problem)
```

**Growth process:**
```go
1. Function prologue checks stack space
2. If insufficient → allocate larger stack (2x current)
3. Copy old stack to new stack
4. Update pointers
5. Continue execution
```

**Example triggering growth:**
```go
func deep(n int) {
    if n == 0 {
        return
    }
    
    // Large local variable forces stack growth:
    var buf [10000]byte
    buf[0] = 1
    
    deep(n - 1)
}

func main() {
    deep(100)  // Stack grows multiple times
}
```

**View stack info:**
```go
import "runtime/debug"

func printStack() {
    // Get stack trace:
    debug.PrintStack()
    
    // Or programmatically:
    buf := make([]byte, 1024)
    n := runtime.Stack(buf, false)
    fmt.Printf("Stack: %s\n", buf[:n])
}
```

**Stack shrinking:**
```go
// Stack shrinks when usage drops below 1/4
// Prevents memory waste after deep recursion
```

---

### Question 307: What is the difference between blocking and non-blocking channels internally?

**Answer:**

**Unbuffered channel (blocking):**
```go
ch := make(chan int)

// Sender blocks until receiver ready:
go func() {
    ch <- 42  // Blocks here
    fmt.Println("Sent!")
}()

val := <-ch  // Sender unblocks when this executes
```

**Internal structure:**
```go
type hchan struct {
    qcount   uint           // Current items in queue
    dataqsiz uint           // Buffer size
    buf      unsafe.Pointer // Buffer array
    sendx    uint           // Send index
    recvx    uint           // Receive index
    recvq    waitq          // Blocked receivers
    sendq    waitq          // Blocked senders
    lock     mutex          // Protects all fields
}
```

**Buffered channel (non-blocking until full):**
```go
ch := make(chan int, 3)

// Won't block (buffer not full):
ch <- 1
ch <- 2
ch <- 3

// Now blocks (buffer full):
ch <- 4  // Blocks here
```

**Non-blocking with select:**
```go
select {
case ch <- value:
    fmt.Println("Sent")
default:
    fmt.Println("Would block, skipped")
}
```

**Performance difference:**
```go
// Unbuffered: Direct handoff (faster when synchronized)
// Buffered: Queue operations (faster when producer/consumer mismatch)

func BenchmarkUnbuffered(b *testing.B) {
    ch := make(chan int)
    go func() {
        for range ch {}
    }()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        ch <- i
    }
}
```

---

### Question 308: What is GOMAXPROCS and how does it affect execution?

**Answer:**
GOMAXPROCS sets the number of OS threads for parallel execution:

**Default value:**
```go
import "runtime"

// Default = number of CPU cores
fmt.Println(runtime.GOMAXPROCS(0))  // Query without changing
```

**Set GOMAXPROCS:**
```go
// Method 1: Programmatically
runtime.GOMAXPROCS(4)  // Use 4 threads

// Method 2: Environment variable
// GOMAXPROCS=4 go run main.go
```

**Effect on performance:**
```go
func compute() {
    for i := 0; i < 1000000000; i++ {
        _ = i * i
    }
}

func testConcurrency() {
    start := time.Now()
    
    var wg sync.WaitGroup
    for i := 0; i < 8; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            compute()
        }()
    }
    wg.Wait()
    
    fmt.Printf("Time: %v\n", time.Since(start))
}

// With GOMAXPROCS=1: ~8s (serial)
// With GOMAXPROCS=8: ~1s (parallel)
```

**When to adjust:**
```go
// ✅ CPU-bound tasks: Set to number of cores
runtime.GOMAXPROCS(runtime.NumCPU())

// ✅ I/O-bound tasks: Can use more than cores
runtime.GOMAXPROCS(runtime.NumCPU() * 2)

// ❌ Usually don't need to change (default is good)
```

**Container awareness:**
```go
// Go 1.5+: Automatically detects CPU quota in containers
// No need to manually set GOMAXPROCS in Docker/K8s
```

---

### Question 309: How does Go manage memory fragmentation?

**Answer:**
Go uses size classes and spans to minimize fragmentation:

**Size classes:**
```go
// Go allocates from predefined size classes:
// 8, 16, 24, 32, 48, 64, 80, ..., 32768 bytes

// Example:
var s []byte
s = make([]byte, 50)  // Allocated from 64-byte class
```

**Memory structure:**
```go
┌─────────────────────────────────────┐
│  Heap                               │
│  ├─ Spans (8KB, 16KB, ...)         │
│  │  ├─ Size class 1 (8 bytes)     │
│  │  ├─ Size class 2 (16 bytes)    │
│  │  └─ Size class N               │
│  └─ Large objects (>32KB)         │
└─────────────────────────────────────┘
```

**mcache (per-P cache):**
```go
// Each P has mcache with spans for each size class
// Fast allocation without locks

type mcache struct {
    alloc [numSpanClasses]*mspan  // Per-size-class spans
}
```

**Allocation path:**
```go
// Small object (<32KB):
1. Try P's mcache
2. If full → get span from mcentral
3. If empty → get span from mheap
4. If needed → request from OS

// Large object (>32KB):
1. Directly allocate from mheap
```

**Reduce fragmentation:**
```go
// 1. Use sync.Pool for frequently allocated objects
var pool = sync.Pool{
    New: func() interface{} {
        return &Object{}
    },
}

// 2. Pre-allocate slices:
data := make([]Item, 0, expectedSize)

// 3. Reuse buffers:
var buf bytes.Buffer
buf.Reset()  // Clear instead of allocating new
```

---

### Question 310: How are maps implemented internally in Go?

**Answer:**
Go maps use hash tables with buckets:

**Structure:**
```go
type hmap struct {
    count     int          // Number of elements
    flags     uint8        // Iterator flags
    B         uint8        // log2(# of buckets)
    noverflow uint16       // Approx overflow buckets
    hash0     uint32       // Hash seed
    buckets   unsafe.Pointer  // Array of buckets
    oldbuckets unsafe.Pointer // Old buckets during growth
}

type bmap struct {
    tophash [8]uint8     // Top 8 bits of hash
    // Followed by 8 keys, 8 values, overflow pointer
}
```

**Hash and bucket selection:**
```go
m := make(map[string]int)
m["key"] = 42

// Internally:
1. hash := hash("key")
2. bucket := hash & (1<<B - 1)  // Select bucket
3. tophash := hash >> (64 - 8)  // Top 8 bits
4. Search bucket for tophash match
5. If found, compare full key
```

**Growth (doubling):**
```go
// When load factor > 6.5 or too many overflow buckets:
m := make(map[int]int)

for i := 0; i < 1000; i++ {
    m[i] = i  // Grows at ~52, ~104, ~208, ~416, ~832
}
```

**Iteration order:**
```go
// Intentionally randomized to prevent dependency:
m := map[string]int{"a": 1, "b": 2, "c": 3}

for k, v := range m {
    fmt.Println(k, v)  // Random order each run
}
```

**Pre-sizing:**
```go
// BAD:
m := make(map[string]int)
for i := 0; i < 10000; i++ {
    m[fmt.Sprint(i)] = i  // Multiple reallocations
}

// GOOD:
m := make(map[string]int, 10000)  // Pre-allocate
for i := 0; i < 10000; i++ {
    m[fmt.Sprint(i)] = i  // No reallocation
}
```

---

### Question 311: How does slice backing array reallocation work?

**Answer:**
Slices grow by copying to a larger backing array:

**Growth strategy:**
```go
// < 1024 elements: double capacity
// ≥ 1024 elements: grow by 25%

s := make([]int, 0)
capacities := []int{}

for i := 0; i < 2000; i++ {
    oldCap := cap(s)
    s = append(s, i)
    if cap(s) != oldCap {
        capacities = append(capacities, cap(s))
    }
}

fmt.Println(capacities)
// [1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 1280, 1696, 2304]
```

**Reallocation example:**
```go
s1 := make([]int, 0, 4)
s1 = append(s1, 1, 2, 3, 4)

// s1 is full, next append triggers reallocation:
s2 := append(s1, 5)

fmt.Printf("s1: %p, cap=%d\n", s1, cap(s1))  // Old array
fmt.Printf("s2: %p, cap=%d\n", s2, cap(s2))  // New array (cap=8)
```

**Shared backing array gotcha:**
```go
original := []int{1, 2, 3, 4, 5}
slice1 := original[:3]    // [1, 2, 3]
slice2 := original[2:]    // [3, 4, 5]

slice1[2] = 999  // Modifies original!

fmt.Println(slice2)  // [999, 4, 5]
```

**Avoid unnecessary allocations:**
```go
// BAD:
var result []int
for i := 0; i < 10000; i++ {
    result = append(result, i)  // Many reallocations
}

// GOOD:
result := make([]int, 0, 10000)
for i := 0; i < 10000; i++ {
    result = append(result, i)  // No reallocation
}
```

**Copy to avoid sharing:**
```go
original := []int{1, 2, 3, 4, 5}
copied := make([]int, len(original))
copy(copied, original)  // Independent backing array
```

---

### Question 312: What is the zero value concept in Go?

**Answer:**
Every type has a default zero value when declared:

**Basic types:**
```go
var i int        // 0
var f float64    // 0.0
var b bool       // false
var s string     // ""
var p *int       // nil
```

**Composite types:**
```go
var arr [3]int          // [0, 0, 0]
var slice []int         // nil (not empty slice)
var m map[string]int    // nil (cannot write to nil map)
var ch chan int         // nil
var fn func()           // nil
```

**Structs:**
```go
type Person struct {
    Name string
    Age  int
    Tags []string
}

var p Person
// p = Person{Name: "", Age: 0, Tags: nil}
```

**Useful zero values:**
```go
// sync.Mutex: Ready to use
var mu sync.Mutex
mu.Lock()  // Works!

// bytes.Buffer: Ready to use
var buf bytes.Buffer
buf.WriteString("hello")  // Works!

// sync.WaitGroup: Ready to use
var wg sync.WaitGroup
wg.Add(1)  // Works!
```

**Dangerous zero values:**
```go
// nil map: Cannot write
var m map[string]int
m["key"] = 1  // PANIC!

// Solution: Initialize
m = make(map[string]int)
m["key"] = 1  // OK

// nil slice: Can append (it allocates)
var s []int
s = append(s, 1)  // OK, creates new slice
```

**Checking for zero values:**
```go
func IsZero(v interface{}) bool {
    return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

// Or for specific types:
if len(slice) == 0 { }  // nil or empty
if m == nil { }         // nil map
if s == "" { }          // empty string
```

---

### Question 313: How does Go avoid data races with its memory model?

**Answer:**
Go's memory model defines synchronization guarantees:

**Happens-before relationship:**
```go
// If event A happens-before event B:
// A's effects are visible to B

var a, b int

func goroutine1() {
    a = 1        // Event A
    b = 2        // Event B
}

func goroutine2() {
    print(b)     // Event C
    print(a)     // Event D
}

// NO guarantee that C sees B, or D sees A
// Could print: 0, 0 or 2, 0 or 0, 1 or 2, 1
```

**Synchronization primitives:**

**1. Channels:**
```go
var a string
var done = make(chan bool)

func setup() {
    a = "hello world"
    done <- true  // Send happens-before receive
}

func main() {
    go setup()
    <-done       // Receive happens after send
    print(a)     // Guaranteed to see "hello world"
}
```

**2. Mutex:**
```go
var l sync.Mutex
var a string

func f() {
    a = "hello"
    l.Unlock()  // Unlock happens-before next Lock
}

func g() {
    l.Lock()    // Lock happens-after previous Unlock
    print(a)    // Guaranteed to see "hello"
}
```

**3. Once:**
```go
var once sync.Once
var a string

func setup() {
    a = "initialized"
}

func doIt() {
    once.Do(setup)  // setup completes happens-before once.Do returns
    print(a)        // Guaranteed to see "initialized"
}
```

**Race detector:**
```bash
# Build with race detector:
go run -race main.go
go test -race ./...
go build -race

# Detects:
# - Conflicting memory accesses
# - At least one is a write
# - Not synchronized
```

**Safe patterns:**
```go
// ✅ Channel communication
ch := make(chan int)
go func() { ch <- data }()
val := <-ch

// ✅ Mutex protection
mu.Lock()
shared = value
mu.Unlock()

// ✅ Atomic operations
atomic.StoreInt32(&counter, value)

// ❌ Unprotected access
var shared int
go func() { shared = 1 }()  // RACE!
print(shared)               // RACE!
```

---

### Question 314: What is escape analysis and how can you visualize it?

**Answer:**
Escape analysis determines if a variable can be stack-allocated or must heap-allocate:

**Run escape analysis:**
```bash
go build -gcflags="-m" main.go

# Or more verbose:
go build -gcflags="-m -m" main.go
```

**Example outputs:**
```go
package main

func stackAlloc() int {
    x := 42  // Does not escape
    return x
}

func heapAlloc() *int {
    x := 42       // Escapes to heap
    return &x
}

func makesSlice() []int {
    s := make([]int, 100)  // Escapes to heap
    return s
}

func main() {
    stackAlloc()
    heapAlloc()
    makesSlice()
}
```

**Build output:**
```bash
$ go build -gcflags="-m" main.go
./main.go:7:2: x escapes to heap
./main.go:12:11: make([]int, 100) escapes to heap
```

**Common escape scenarios:**
```go
// 1. Return pointer to local:
func escape1() *int {
    x := 1
    return &x  // x escapes
}

// 2. Assign to interface:
func escape2() {
    x := 1
    var i interface{} = x  // x escapes
}

// 3. Send to channel:
func escape3() {
    x := 1
    ch := make(chan *int)
    ch <- &x  // x escapes
}

// 4. Large structs:
func escape4() {
    var buf [1024 * 1024]byte  // Too large, escapes
}

// 5. Closures capturing variables:
func escape5() func() int {
    x := 1
    return func() int {
        return x  // x escapes
    }
}
```

**Optimization tips:**
```go
// ✅ Return value, not pointer:
func noEscape1() int {
    x := 1
    return x  // x on stack
}

// ✅ Use small structs:
type Small struct {
    a, b int
}  // Fits on stack

// ✅ Avoid interface{} when possible
func process(x int) { }  // x on stack

// ❌ interface{} forces heap:
func process(x interface{}) { }  // x escapes
```

---

### Question 315: How are method sets determined in Go?

**Answer:**
Method sets define which methods are available on a type:

**Rules:**
```go
type T struct {
    value int
}

// Value receiver method:
func (t T) ValueMethod() {
    fmt.Println(t.value)
}

// Pointer receiver method:
func (t *T) PointerMethod() {
    t.value++
}
```

**Method sets:**
```go
// Type T:
// - ValueMethod only

// Type *T:
// - ValueMethod (promoted)
// - PointerMethod

var v T
v.ValueMethod()     // ✅ OK
v.PointerMethod()   // ✅ OK (auto-addressed: (&v).PointerMethod())

var p *T = &T{}
p.ValueMethod()     // ✅ OK (auto-dereferenced: (*p).ValueMethod())
p.PointerMethod()   // ✅ OK
```

**Interface satisfaction:**
```go
type Incrementer interface {
    Increment()
}

type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++
}

// Usage:
var i Incrementer

c := Counter{}
i = c   // ❌ Compile error: Counter doesn't implement Incrementer
i = &c  // ✅ OK: *Counter implements Incrementer

// Why? Because Increment has pointer receiver,
// only *Counter has it in method set
```

**Rule of thumb:**
```go
// Value receiver:
// - Method doesn't modify receiver
// - Small receiver (copy is cheap)
// - Primitive types, small structs

func (t Time) After(u Time) bool { }

// Pointer receiver:
// - Method modifies receiver
// - Large receiver (avoid copy)
// - Consistency (if one method uses pointer, all should)

func (b *Buffer) Write(p []byte) (n int, err error) { }
```

**Method set table:**
```
Type      | Has methods with receiver
          | Value (T) | Pointer (*T)
----------|-----------|--------------
T         |    ✅     |      ❌
*T        |    ✅     |      ✅
```

---

### Question 316: What is the difference between pointer receiver and value receiver at runtime?

**Answer:**

**Value receiver (copy):**
```go
type Counter struct {
    count int
}

func (c Counter) IncrementValue() {
    c.count++  // Modifies copy, not original
}

func main() {
    c := Counter{count: 0}
    c.IncrementValue()
    fmt.Println(c.count)  // Still 0!
}
```

**Pointer receiver (reference):**
```go
func (c *Counter) IncrementPointer() {
    c.count++  // Modifies original
}

func main() {
    c := Counter{count: 0}
    c.IncrementPointer()
    fmt.Println(c.count)  // Now 1
}
```

**Performance impact:**
```go
type LargeStruct struct {
    data [1024]int
}

// BAD - Copies 8KB on every call:
func (ls LargeStruct) Process() { }

// GOOD - Copies only 8 bytes (pointer):
func (ls *LargeStruct) Process() { }

// Benchmark:
func BenchmarkValue(b *testing.B) {
    ls := LargeStruct{}
    for i := 0; i < b.N; i++ {
        ls.Process()  // Slow (copies 8KB)
    }
}

func BenchmarkPointer(b *testing.B) {
    ls := &LargeStruct{}
    for i := 0; i < b.N; i++ {
        ls.Process()  // Fast (copies 8 bytes)
    }
}
```

**Memory allocation:**
```go
type Small struct {
    value int
}

// Value receiver: May stay on stack
func (s Small) Get() int {
    return s.value
}

// Pointer receiver: May force heap allocation
func (s *Small) Set(v int) {
    s.value = v
}

// Check with: go build -gcflags="-m"
```

**When to use each:**
```go
// ✅ Value receiver:
// - Small types (< 16 bytes)
// - Read-only operations
// - Primitive types, time.Time, etc.

type Point struct {
    X, Y int
}

func (p Point) Distance() float64 {
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

// ✅ Pointer receiver:
// - Need to modify receiver
// - Large structs
// - Implement interfaces (consistency)

type Database struct {
    // Large struct
}

func (db *Database) Query(sql string) ([]Row, error) {
    // ...
}
```

---

### Question 317: How does Go handle panics internally?

**Answer:**
Panics unwind the stack until recovered or program terminates:

**Panic mechanism:**
```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    
    panic("something went wrong")
    
    // This won't execute
    fmt.Println("After panic")
}
```

**Stack unwinding:**
```go
func level3() {
    panic("error at level 3")
}

func level2() {
    defer fmt.Println("Defer in level2")
    level3()
    fmt.Println("After level3")  // Not executed
}

func level1() {
    defer fmt.Println("Defer in level1")
    level2()
    fmt.Println("After level2")  // Not executed
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
            debug.PrintStack()
        }
    }()
    
    level1()
}

// Output:
// Defer in level2
// Defer in level1
// Recovered: error at level 3
// [stack trace]
```

**Internal structure:**
```go
type _panic struct {
    argp      unsafe.Pointer  // Pointer to arguments
    arg       interface{}     // Panic argument
    link      *_panic         // Link to older panic
    recovered bool            // Was recover called?
    aborted   bool            // Panic aborted?
}
```

**Multiple panics:**
```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered first:", r)
            panic("second panic")  // Can panic again
        }
    }()
    
    panic("first panic")
}
```

**Panic in goroutine:**
```go
func main() {
    go func() {
        panic("goroutine panic")  // Crashes entire program!
    }()
    
    time.Sleep(time.Second)
    fmt.Println("Main")  // Won't print
}

// Solution: Always recover in goroutines:
func main() {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                log.Println("Recovered in goroutine:", r)
            }
        }()
        panic("goroutine panic")
    }()
    
    time.Sleep(time.Second)
    fmt.Println("Main")  // Will print
}
```

**When to panic:**
```go
// ✅ Initialization errors:
func init() {
    if config == nil {
        panic("config not loaded")
    }
}

// ✅ Programming errors:
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}

// ❌ Normal errors:
// Use errors, not panics!
func readFile(name string) ([]byte, error) {
    data, err := os.ReadFile(name)
    if err != nil {
        return nil, err  // Don't panic!
    }
    return data, nil
}
```

---

### Question 318: What is reflection implemented in Go?

**Answer:**
Reflection uses interface{} type information to inspect and manipulate values:

**Basic reflection:**
```go
import "reflect"

func inspectType(x interface{}) {
    t := reflect.TypeOf(x)
    v := reflect.ValueOf(x)
    
    fmt.Printf("Type: %v\n", t)
    fmt.Printf("Kind: %v\n", t.Kind())
    fmt.Printf("Value: %v\n", v)
}

// Usage:
inspectType(42)              // Type: int, Kind: int
inspectType("hello")         // Type: string, Kind: string
inspectType([]int{1, 2, 3})  // Type: []int, Kind: slice
```

**Modify values:**
```go
func modify(x interface{}) {
    v := reflect.ValueOf(x)
    
    if v.Kind() != reflect.Ptr {
        panic("not a pointer")
    }
    
    elem := v.Elem()
    if !elem.CanSet() {
        panic("cannot set")
    }
    
    switch elem.Kind() {
    case reflect.Int:
        elem.SetInt(42)
    case reflect.String:
        elem.SetString("modified")
    }
}

// Usage:
var i int = 10
modify(&i)
fmt.Println(i)  // 42
```

**Struct reflection:**
```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func inspectStruct(x interface{}) {
    t := reflect.TypeOf(x)
    
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Field: %s, Type: %s, Tag: %s\n",
            field.Name,
            field.Type,
            field.Tag.Get("json"))
    }
}

// Usage:
p := Person{Name: "Alice", Age: 30}
inspectStruct(p)
// Field: Name, Type: string, Tag: name
// Field: Age, Type: int, Tag: age
```

**Call methods dynamically:**
```go
type Calculator struct{}

func (c Calculator) Add(a, b int) int {
    return a + b
}

func callMethod(obj interface{}, methodName string, args ...interface{}) {
    v := reflect.ValueOf(obj)
    method := v.MethodByName(methodName)
    
    // Convert args to reflect.Value:
    in := make([]reflect.Value, len(args))
    for i, arg := range args {
        in[i] = reflect.ValueOf(arg)
    }
    
    // Call method:
    results := method.Call(in)
    
    // Print results:
    for _, result := range results {
        fmt.Println(result.Interface())
    }
}

// Usage:
calc := Calculator{}
callMethod(calc, "Add", 10, 20)  // 30
```

**Performance cost:**
```go
func BenchmarkDirect(b *testing.B) {
    x := 42
    for i := 0; i < b.N; i++ {
        _ = x
    }
}

func BenchmarkReflection(b *testing.B) {
    x := 42
    v := reflect.ValueOf(x)
    for i := 0; i < b.N; i++ {
        _ = v.Int()
    }
}
// Reflection is ~10-100x slower
```

---

### Question 319: What is type identity in Go?

**Answer:**
Type identity determines if two types are the same:

**Identical types:**
```go
// Same type name:
var a int
var b int
// a and b have identical types

// Same underlying type:
type MyInt int
type YourInt int

var x MyInt
var y YourInt
// x and y have DIFFERENT types (not identical)
```

**Named vs unnamed types:**
```go
// Named types:
type Person struct {
    Name string
}

type Employee struct {
    Name string
}

var p Person
var e Employee
// Different types even if same structure

// Unnamed types (identical if same structure):
var a struct{ Name string }
var b struct{ Name string }
// a and b have identical types
```

**Type conversion:**
```go
type Celsius float64
type Fahrenheit float64

var c Celsius = 100
var f Fahrenheit

// Need explicit conversion (not identical types):
f = Fahrenheit(c)  // OK
// f = c           // Error!

// But can assign to float64:
var temp float64
temp = float64(c)  // OK
```

**Function types:**
```go
type Handler func(string) error
type Processor func(string) error

var h Handler
var p Processor

// Not identical (different type names):
// p = h  // Error!

// But same signature:
p = Processor(h)  // OK with conversion

// Identical unnamed function types:
var fn1 func(string) error
var fn2 func(string) error
fn1 = fn2  // OK (identical unnamed types)
```

**Interface types:**
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type MyReader interface {
    Read(p []byte) (n int, err error)
}

var r Reader
var mr MyReader

// Not identical (different type names):
// mr = r  // Error!

// Identical unnamed interfaces:
var i1 interface{ Read([]byte) (int, error) }
var i2 interface{ Read([]byte) (int, error) }
i1 = i2  // OK
```

**Assignability:**
```go
// T is assignable to T (identical)
// T is assignable to interface if T implements it
// x is assignable to T if x is untyped constant

// Examples:
var x int = 42       // untyped constant → int
var y float64 = 42   // untyped constant → float64

type MyInt int
var a MyInt = 42     // untyped constant → MyInt
// var b MyInt = x   // Error! (x is typed as int)
```

---

### Question 320: How are interface values represented in memory?

**Answer:**
Interfaces are represented as (type, value) pairs:

**Interface structure:**
```go
type iface struct {
    tab  *itab           // Type information + method table
    data unsafe.Pointer  // Actual value
}

type itab struct {
    inter *interfacetype  // Interface type
    _type *_type          // Concrete type
    fun   [1]uintptr      // Method pointers
}
```

**Example:**
```go
var x interface{} = 42

// Memory layout:
// x = iface{
//     tab: &itab{
//         inter: interface{} type info
//         _type: int type info
//         fun: []
//     },
//     data: pointer to 42
// }
```

**With methods:**
```go
type Reader interface {
    Read(p []byte) (int, error)
}

type File struct{}

func (f *File) Read(p []byte) (int, error) {
    return 0, nil
}

var r Reader = &File{}

// Memory layout:
// r = iface{
//     tab: &itab{
//         inter: Reader interface info
//         _type: *File type info
//         fun: [&File.Read]  // Method pointer
//     },
//     data: pointer to File instance
// }
```

**Nil interface vs nil value:**
```go
var i interface{}
fmt.Println(i == nil)  // true

var p *int
i = p
fmt.Println(i == nil)  // false! (type=*int, value=nil)

// Explanation:
// First: i = (nil, nil)          → nil
// Second: i = (*int, nil)        → not nil (has type)
```

**Empty vs non-empty interface:**
```go
// Empty interface (interface{}):
type eface struct {
    _type *_type          // Just type info
    data  unsafe.Pointer  // Value
}

// Non-empty interface (has methods):
type iface struct {
    tab  *itab           // Type + methods
    data unsafe.Pointer  // Value
}
```

**Type assertion cost:**
```go
var i interface{} = 42

// Direct access:
x := i.(int)  // Check tab._type == int, then extract data

// With check:
x, ok := i.(int)  // Same, but no panic on failure

// Type switch:
switch v := i.(type) {
case int:    // Compare tab._type
case string: // Compare tab._type
}
```

**Interface satisfaction:**
```go
type Walker interface {
    Walk()
}

type Runner interface {
    Run()
}

type Person struct{}

func (p Person) Walk() {}
func (p Person) Run()  {}

// Person satisfies both:
var w Walker = Person{}  // itab with Walk pointer
var r Runner = Person{}  // itab with Run pointer

// Different itabs for same value!
```

---

*[Questions 321-360 will cover DevOps, Docker, Streaming, and Async Processing]*
## 🟡 DevOps, Docker, and Deployment (Questions 321-340)

### Question 321: How do you containerize a Go application?

**Answer:**
Containerizing a Go app involves creating a `Dockerfile` to build and run the binary.

**Basic Dockerfile:**
```dockerfile
# Start from a small base image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Command to run
CMD ["./main"]
```

**Build and Run:**
```bash
docker build -t my-go-app .
docker run -p 8080:8080 my-go-app
```

---

### Question 322: What is a multi-stage Docker build and how does it help with Go?

**Answer:**
Multi-stage builds allow you to use a large image for building (with compilers/tools) and a tiny image for running (just the binary). This drastically reduces image size.

**Example:**
```dockerfile
# Stage 1: Build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/myapp .
CMD ["./myapp"]
```

**Benefits:**
- Removes source code, compilers, and cache from final image.
- Reduces image size from ~800MB (golang image) to ~10MB (alpine + binary).
- Improves security (fewer tools installed in runtime container).

---

### Question 323: How do you reduce the size of a Go Docker image?

**Answer:**
1. **Use Multi-stage builds:** Separate build and runtime environments.
2. **Use `scratch` or `alpine` base images:**
   - `scratch`: Empty image (smallest possible).
   - `alpine`: Minimal Linux (~5MB).
3. **Strip debug information:**
   ```bash
   go build -ldflags="-s -w" -o myapp
   ```
   - `-s`: Disable symbol table.
   - `-w`: Disable DWARF generation.
4. **Compress binary (optional):** Use `upx`.

**Optimized Dockerfile:**
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o myapp .

FROM scratch
COPY --from=builder /app/myapp /myapp
# Copy CA certificates for HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/myapp"]
```

---

### Question 324: How do you handle secrets in Go apps deployed via Docker?

**Answer:**
Never hardcode secrets. Use environment variables or secret managers.

**1. Environment Variables:**
```go
func main() {
    dbPass := os.Getenv("DB_PASSWORD")
    if dbPass == "" {
        log.Fatal("DB_PASSWORD not set")
    }
}
```
Run with: `docker run -e DB_PASSWORD=secret myapp`

**2. Docker Secrets (Swarm/K8s):**
Mount secrets as files (e.g., `/run/secrets/db_password`).
```go
func readSecret(name string) string {
    content, _ := os.ReadFile("/run/secrets/" + name)
    return strings.TrimSpace(string(content))
}
```

**3. External Secret Managers:**
Use SDKs for AWS Secrets Manager, HashiCorp Vault, etc.

---

### Question 325: How do you use environment variables in Go?

**Answer:**
Use the `os` package or libraries like `godotenv` or `viper`.

**Standard library:**
```go
import "os"

// Get value
port := os.Getenv("PORT")
if port == "" {
    port = "8080" // Default
}

// Set value (for current process)
os.Setenv("APP_ENV", "production")

// Expand string
conn := os.ExpandEnv("postgres://user:$DB_PASSWORD@localhost:5432/db")
```

**Using `godotenv` (Load from .env file):**
```go
import "github.com/joho/godotenv"

func main() {
    godotenv.Load() // Loads .env file
    s3Bucket := os.Getenv("S3_BUCKET")
}
```

---

### Question 326: How do you compile a static Go binary for Alpine Linux?

**Answer:**
Alpine uses `musl` libc instead of `glibc`. To ensure compatibility or avoid dependency issues, build a statically linked binary.

**Command:**
```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .
```

- `CGO_ENABLED=0`: Disables CGO, ensuring a pure Go binary (static by default).
- `GOOS=linux`: Targets Linux.
- `-a`: Force rebuild of packages.

**Why?**
- Allows the binary to run on `scratch` images.
- Avoids "no such file or directory" errors when moving binaries between distros.

---

### Question 327: What is scratch image in Docker and why is it used with Go?

**Answer:**
`scratch` is an explicitly empty Docker image. It contains no filesystem, no shell, no libraries.

**Why use it with Go?**
- **Size:** Results in the smallest possible Docker image (size = binary size).
- **Security:** Small attack surface (no shell, no tools to exploit).

**Gotchas:**
- **No shell:** Can't use `sh`, `bash`, or `exec` inside.
- **Missing SSL certs:** HTTPS requests will fail. You must copy `/etc/ssl/certs/ca-certificates.crt` from a builder stage.
- **Missing Timezone data:** Time functions might work, but loading locations (e.g., `time.LoadLocation("America/New_York")`) requires copying `zoneinfo`.

---

### Question 328: How do you manage config files in Go across environments?

**Answer:**
Best practice is to prioritize: Flag > Env Var > Config File > Default.

**Using `viper`:**
```go
import "github.com/spf13/viper"

func initConfig() {
    viper.SetConfigName("config") // config.yaml
    viper.AddConfigPath(".")      // Look in current dir
    viper.AutomaticEnv()          // Read env vars

    // Set defaults
    viper.SetDefault("port", 8080)

    if err := viper.ReadInConfig(); err != nil {
        log.Println("No config file found, using defaults")
    }
}

func main() {
    port := viper.GetInt("port")
}
```

**Pattern:**
- Local dev: use `config.yaml`
- Production: use Environment Variables (override config file).

---

### Question 329: How do you build Go binaries for different OS/arch?

**Answer:**
Go supports cross-compilation out of the box using `GOOS` and `GOARCH` environment variables.

**Examples:**
```bash
# Windows (64-bit)
GOOS=windows GOARCH=amd64 go build -o myapp.exe

# Linux (64-bit)
GOOS=linux GOARCH=amd64 go build -o myapp-linux

# macOS (M1/M2 - ARM64)
GOOS=darwin GOARCH=arm64 go build -o myapp-mac

# Linux (ARM - Raspberry Pi)
GOOS=linux GOARCH=arm go build -o myapp-pi
```

**Check supported platforms:**
```bash
go tool dist list
```

---

### Question 330: How do you use GoReleaser?

**Answer:**
GoReleaser automates building, packaging, and releasing Go binaries to GitHub/GitLab.

**Steps:**
1. **Install:** `brew install goreleaser`
2. **Init:** `goreleaser init` (creates `.goreleaser.yaml`)
3. **Configure `.goreleaser.yaml`:**
   ```yaml
   builds:
     - env:
         - CGO_ENABLED=0
       goos:
         - linux
         - windows
         - darwin
       goarch:
         - amd64
         - arm64
   ```
4. **Release:**
   ```bash
   git tag -a v1.0.0 -m "First release"
   git push origin v1.0.0
   goreleaser release --clean
   ```

**What it does:**
- Cross-compiles for defined targets.
- Creates archives (.tar.gz, .zip).
- Generates checksums.
- Creates GitHub Release and uploads artifacts.
- Can build Docker images and update Homebrew taps.

---

### Question 331: What is a Docker healthcheck for a Go app?

**Answer:**
A healthcheck ensures the container is actually ready to serve traffic, not just running.

**In Dockerfile:**
```dockerfile
# Check every 30s, timeout 3s
HEALTHCHECK --interval=30s --timeout=3s \
  CMD curl -f http://localhost:8080/health || exit 1
```
*Note: This requires `curl` to be installed in the image.*

**Go Implementation:**
```go
func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        // Check DB connection, cache, etc.
        if err := db.Ping(); err != nil {
            w.WriteHeader(http.StatusServiceUnavailable)
            return
        }
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })
    http.ListenAndServe(":8080", nil)
}
```

**For `scratch` images:**
Use a separate Go binary for healthchecking or the built-in `grpc-health-probe` if using gRPC.

---

### Question 332: How do you log container stdout/stderr from Go?

**Answer:**
By default, Docker captures stdout and stderr. In Go, simply print to these streams.

**Standard Log:**
```go
log.Println("This goes to stderr") // Default logger uses stderr
```

**Explicit Streams:**
```go
fmt.Fprintln(os.Stdout, "Info log")
fmt.Fprintln(os.Stderr, "Error log")
```

**Structured Logging (JSON):**
For production/parsing (ELK, Datadog), use JSON logs.
```go
import "log/slog" // Go 1.21+

logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("User logged in", "user_id", 42)
// Output: {"time":"...","level":"INFO","msg":"User logged in","user_id":42}
```

**Docker logs command:**
```bash
docker logs -f my-container
```

---

### Question 333: How do you set up autoscaling for Go services?

**Answer:**
Autoscaling is typically handled by the orchestrator (Kubernetes), not the Go app itself.

**Horizontal Pod Autoscaler (HPA) in K8s:**
Scales pods based on CPU/Memory or custom metrics.

1. **Expose Metrics:**
   Use specific endpoints (e.g., `/metrics` with **Prometheus**).
   ```go
   // Expose "active_requests" metric
   ```

2. **Configure HPA:**
   ```yaml
   apiVersion: autoscaling/v2
   kind: HorizontalPodAutoscaler
   metadata:
     name: my-go-app
   spec:
     scaleTargetRef:
       apiVersion: apps/v1
       kind: Deployment
       name: my-go-app
     minReplicas: 2
     maxReplicas: 10
     metrics:
     - type: Resource
       resource:
         name: cpu
         target:
           type: Utilization
           averageUtilization: 50
   ```
   *Scales up if CPU usage > 50%.*

---

### Question 334: How would you containerize a gRPC Go service?

**Answer:**
Similar to HTTP, but needs to expose the gRPC port.

**Dockerfile:**
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o grpc-server server.go

FROM alpine:latest
COPY --from=builder /src/grpc-server /app/
EXPOSE 50051
CMD ["/app/grpc-server"]
```

**Testing:**
Since gRPC is not HTTP (usually), you can't just `curl` it.
- Use **`grpcurl`** for command line testing.
- Include **gRPC Health Probe** in the container for K8s readiness probes.

```dockerfile
# Install health probe
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.19 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
```

---

### Question 335: How to deploy Go microservices in Kubernetes?

**Answer:**
Deployment involves creating K8s manifests:

**1. Deployment.yaml:**
Defines the pods, replicas, and container image.
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: go-api
  template:
    metadata:
      labels:
        app: go-api
    spec:
      containers:
      - name: go-api
        image: myregistry/go-api:v1
        ports:
        - containerPort: 8080
        env:
        - name: DB_URL
          valueFrom:
            secretKeyRef:
              name: db-secrets
              key: url
```

**2. Service.yaml:**
Exposes the pods to the cluster.
```yaml
apiVersion: v1
kind: Service
metadata:
  name: go-api-svc
spec:
  selector:
    app: go-api
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
```

**Steps:**
1. Build image: `docker build -t ...`
2. Push image: `docker push ...`
3. Apply manifests: `kubectl apply -f k8s/`

---

### Question 336: How do you write Helm charts for a Go app?

**Answer:**
Helm creates reusable, templated K8s manifests.

**Structure:**
```
my-chart/
  Chart.yaml    # Meta info
  values.yaml   # Default config
  templates/    # Manifest templates
    deployment.yaml
    service.yaml
```

**values.yaml:**
```yaml
replicaCount: 2
image:
  repository: my-go-app
  tag: "1.0.0"
service:
  port: 80
```

**templates/deployment.yaml:**
```yaml
apiVersion: apps/v1
kind: Deployment
spec:
  replicas: {{ .Values.replicaCount }}
  ...
      containers:
        - image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
```

**Install:**
```bash
helm install my-release ./my-chart
```
This allows deploying the same Go app to Dev, Staging, and Prod with different `values.yaml` files.

---

### Question 337: How do you monitor a Go service in production?

**Answer:**
Comprehensive monitoring includes:

1. **Metrics (Quantitative):**
   - Use **Prometheus** for time-series data.
   - Track: Request rate, Error rate, Latency (RED method), Memory, Goroutines.

2. **Logs (Qualitative):**
   - Use Structured Logging (JSON).
   - Aggregate via ELK stack (Elasticsearch, Logstash, Kibana) or Loki.

3. **Tracing (Context):**
   - Use **OpenTelemetry** (Jaeger/Zipkin) to trace requests across microservices.

4. **Health Checks:**
   - Liveness/Readiness probes for K8s.

**Go Code:**
Expose `/metrics` endpoint for Prometheus scraping and inject trace IDs into logs.

---

### Question 338: How do you use Prometheus with a Go app?

**Answer:**
Use the `prometheus/client_golang` library.

**Code:**
```go
package main

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define a metric
var reqCounter = prometheus.NewCounter(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total number of HTTP requests",
    },
)

func init() {
    // Register metric
    prometheus.MustRegister(reqCounter)
}

func handler(w http.ResponseWriter, r *http.Request) {
    reqCounter.Inc() // Increment
    w.Write([]byte("Hello"))
}

func main() {
    http.HandleFunc("/", handler)
    // Expose metrics endpoint
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":8080", nil)
}
```
Prometheus server then scrapes `http://app:8080/metrics`.

---

### Question 339: How do you enable structured logging in production?

**Answer:**
Structured logging (JSON) is essential for machine parsing.

**Using `log/slog` (Go 1.21+):**
```go
import (
    "log/slog"
    "os"
)

func main() {
    // Determine format based on env
    var handler slog.Handler
    if os.Getenv("ENV") == "production" {
        handler = slog.NewJSONHandler(os.Stdout, nil)
    } else {
        handler = slog.NewTextHandler(os.Stdout, nil)
    }
    
    logger := slog.New(handler)
    slog.SetDefault(logger)

    // Log with fields
    slog.Info("Request processed", 
        "path", "/api/v1/user",
        "status", 200,
        "duration_ms", 45,
    )
}
```
**Output (Prod):**
`{"time":"...","level":"INFO","msg":"Request processed","path":"/api/v1/user","status":200,"duration_ms":45}`

---

### Question 340: How do you handle log rotation in containerized Go apps?

**Answer:**
**You usually don't.**

In containerized environments (Docker/K8s):
1. **App Responsibility:** The app should simply write to `stdout`/`stderr`. It should **not** manage files or rotation.
2. **Platform Responsibility:** The container runtime (Docker daemon, Kubelet) captures these logs.
   - **Docker:** Configure logging driver (e.g., `json-file` with `max-size` and `max-file`).
   - **K8s:** Often uses a sidecar (e.g., Fluentd, Filebeat) to ship logs to a central backend (Elasticsearch, Splunk).

**If running on raw VM (Systemd):**
Use **Lumberjack** library in Go to manage file rotation application-side.
```go
import "gopkg.in/natefinch/lumberjack.v2"

log.SetOutput(&lumberjack.Logger{
    Filename:   "/var/log/myapp/foo.log",
    MaxSize:    100, // megabytes
    MaxBackups: 3,
    MaxAge:     28,  // days
    Compress:   true,
})
```
But for Docker, stick to standard output.
## 🔵 Streaming, Messaging, and Asynchronous Processing (Questions 341-360)

### Question 341: How do you consume messages from Kafka in Go?

**Answer:**
Use a library like `github.com/IBM/sarama` or `github.com/segmentio/kafka-go`.

**Using `kafka-go` (Simpler API):**
```go
import (
    "context"
    "fmt"
    "github.com/segmentio/kafka-go"
)

func consume() {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:  []string{"localhost:9092"},
        Topic:    "my-topic",
        GroupID:  "my-group",
        MinBytes: 10e3, // 10KB
        MaxBytes: 10e6, // 10MB
    })
    defer r.Close()

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            break
        }
        fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
    }
}
```

---

### Question 342: How do you publish messages to a RabbitMQ topic?

**Answer:**
Use `github.com/rabbitmq/amqp091-go`.

**Steps:**
1. Connect to RabbitMQ.
2. Open a channel.
3. Declare an exchange (Topic type).
4. Publish message with routing key.

**Code:**
```go
import "github.com/rabbitmq/amqp091-go"

func publish() {
    conn, _ := amqp091.Dial("amqp://guest:guest@localhost:5672/")
    defer conn.Close()

    ch, _ := conn.Channel()
    defer ch.Close()

    err := ch.ExchangeDeclare(
        "logs_topic", // name
        "topic",      // type
        true,         // durable
        false,        // auto-deleted
        false,        // internal
        false,        // no-wait
        nil,          // arguments
    )

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err = ch.PublishWithContext(ctx,
        "logs_topic", // exchange
        "kern.critical", // routing key
        false,        // mandatory
        false,        // immediate
        amqp091.Publishing{
            ContentType: "text/plain",
            Body:        []byte("Kernel panic!"),
        })
}
```

---

### Question 343: What is the idiomatic way to implement a message handler in Go?

**Answer:**
Use an interface-based design to decouple the handler logic from the transport (Kafka/RabbitMQ/HTTP).

**Pattern:**
```go
// 1. Define Handler Interface
type MessageHandler interface {
    Handle(ctx context.Context, msg []byte) error
}

// 2. Concrete Implementation
type OrderProcessor struct{}

func (op *OrderProcessor) Handle(ctx context.Context, msg []byte) error {
    var order Order
    json.Unmarshal(msg, &order)
    return processOrder(order)
}

// 3. Worker (Transport Layer)
func StartConsumer(handler MessageHandler, messages <-chan []byte) {
    for msg := range messages {
        go func(m []byte) {
            if err := handler.Handle(context.Background(), m); err != nil {
                log.Println("Error handling message:", err)
                // Nack/Retry logic here
            }
        }(msg)
    }
}
```
This makes unit testing the `OrderProcessor` trivial without mocking Kafka.

---

### Question 344: How would you implement a worker pool pattern?

**Answer:**
A worker pool limits concurrency to a fixed number of goroutines.

**Implementation:**
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        time.Sleep(time.Second) // Simulate work
        results <- j * 2
    }
}

func main() {
    const numJobs = 100
    const numWorkers = 5

    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // 1. Start workers
    for w := 1; w <= numWorkers; w++ {
        go worker(w, jobs, results)
    }

    // 2. Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs) // Signal no more jobs

    // 3. Collect results
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
```

---

### Question 345: How do you use the context package for cancellation in streaming apps?

**Answer:**
Pass `context.Context` to all long-running operations. Monitor `ctx.Done()` to stop processing immediately.

**Example:**
```go
func StreamProcessor(ctx context.Context, stream <-chan Data) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Tracking stream stopped:", ctx.Err())
            return
        case data := <-stream:
            process(data)
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    go StreamProcessor(ctx, dataChannel)
    
    // ... main continues ...
}
```
If the timeout hits or `cancel()` is called, the `StreamProcessor` exits immediately, preventing goroutine leaks.

---

### Question 346: How do you retry failed messages in Go?

**Answer:**
Use **Exponential Backoff** with **Jitter**.

**Library:** `github.com/cenkalti/backoff`

**Manual Implementation:**
```go
func processWithRetry(operation func() error) error {
    maxRetries := 5
    baseDelay := 100 * time.Millisecond

    for i := 0; i < maxRetries; i++ {
        err := operation()
        if err == nil {
            return nil
        }
        
        // Exponential backoff: 100ms, 200ms, 400ms...
        delay := baseDelay * time.Duration(1<<i)
        
        // Add jitter (randomness) to prevent thundering herd
        jitter := time.Duration(rand.Intn(50)) * time.Millisecond
        
        log.Printf("Retry %d after error: %v. Waiting %v", i+1, err, delay+jitter)
        time.Sleep(delay + jitter)
    }
    return fmt.Errorf("operation failed after %d retries", maxRetries)
}
```

---

### Question 347: What is dead-letter queue (DLQ) and how do you use it?

**Answer:**
A DLQ is a standard queue where "bad" messages (that failed processing after max retries) are sent for manual inspection.

**Strategy:**
1. Consumer attempts to process message.
2. Implementation fails → Retry 3 times.
3. Still fails → Publish message to `topic-dlq`.
4. Acknowledge original message to remove it from main queue.
5. **Alerting:** Monitor DLQ depth to alert developers.

**Go Code Snippet:**
```go
if err := process(msg); err != nil {
    if retries >= maxRetries {
        // Publish to DLQ
        producer.Publish("my-topic-dlq", msg.Body)
        // Ack on main topic
        msg.Ack() 
    } else {
        retries++
        // Nack to retry later
        msg.Nack()
    }
}
```

---

### Question 348: How do you handle idempotency in message consumers?

**Answer:**
Idempotency ensures processing the same message multiple times has the same effect as processing it once.

**Strategies:**
1. **Database Uniqueness:** Use the Message ID as a Primary Key/Unique Constraint in the DB.
   - If `INSERT` fails with "Duplicate Key", ignore the message.
2. **Redis Deduplication:**
   - Check if `MessageID` exists in Redis.
   - If not, process and set `MessageID` with TTL.

**Example (DB Approach):**
```go
func processOrder(db *sql.DB, startOrder StartOrderMsg) error {
    tx, _ := db.Begin()
    
    // Check if processed
    var exists bool
    tx.QueryRow("SELECT exists(SELECT 1 FROM processed_msgs WHERE id=$1)", startOrder.MsgID).Scan(&exists)
    
    if exists {
        return nil // Already processed, safe to ack
    }

    // Process logic...
    
    // Mark as processed
    tx.Exec("INSERT INTO processed_msgs (id) VALUES ($1)", startOrder.MsgID)
    
    return tx.Commit()
}
```

---

### Question 349: How do you implement exponential backoff in Go?

**Answer:**
Wait time increases exponentially with each failure ($Base \times 2^{Attempt}$).

```go
func Retry(attempts int, sleep time.Duration, f func() error) error {
    if err := f(); err != nil {
        if attempts--; attempts > 0 {
            // Jitter for robustness
            jitter := time.Duration(rand.Int63n(int64(sleep))) / 2
            sleep = sleep + jitter/2

            time.Sleep(sleep)
            return Retry(attempts, 2*sleep, f)
        }
        return err
    }
    return nil
}
```

---

### Question 350: How do you stream logs to a file/socket in real-time?

**Answer:**
Use `io.MultiWriter` to write to multiple destinations (Console + File/Socket).

```go
func main() {
    // 1. Open file
    file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    
    // 2. Open Socket (e.g., Logstash)
    conn, _ := net.Dial("tcp", "localhost:5000")

    // 3. Create MultiWriter
    logger := log.New(io.MultiWriter(os.Stdout, file, conn), "INFO: ", log.LstdFlags)

    logger.Println("This log goes to Console, File, and Socket!")
}
```

---

### Question 351: How do you work with WebSockets in Go?

**Answer:**
Use `github.com/gorilla/websocket` (standard) or `nhooyr.io/websocket` (minimal).

**Server Example (Gorilla):**
```go
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWS(w http.ResponseWriter, r *http.Request) {
    conn, _ := upgrader.Upgrade(w, r, nil) // Upgrade HTTP to WS
    defer conn.Close()

    for {
        // Read message
        mt, message, err := conn.ReadMessage()
        if err != nil { break }
        
        log.Printf("Received: %s", message)

        // Echo back
        conn.WriteMessage(mt, message)
    }
}
```

---

### Question 352: How do you handle bi-directional streaming in gRPC?

**Answer:**
Define `stream` in both request and response in Protobuf.

**Proto:**
```protobuf
service ChatService {
  rpc Chat(stream ChatMessage) returns (stream ChatMessage);
}
```

**Go Implementation:**
```go
func (s *server) Chat(stream pb.ChatService_ChatServer) error {
    for {
        // 1. Receive
        in, err := stream.Recv()
        if err == io.EOF { return nil }
        if err != nil { return err }

        log.Printf("Got: %s", in.Message)

        // 2. Send
        err = stream.Send(&pb.ChatMessage{Message: "Reply: " + in.Message})
        if err != nil { return err }
    }
}
```

---

### Question 353: What is Server-Sent Events (SSE) and how is it done in Go?

**Answer:**
SSE sends one-way real-time updates from Server to Client over HTTP. It's simpler than WebSockets.

**Implementation:**
1. Set Headers: `Content-Type: text/event-stream`.
2. Flush writes immediately.

```go
func sseHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")
    w.Header().Set("Connection", "keep-alive")

    flusher, ok := w.(http.Flusher)
    if !ok { return }

    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case t := <-ticker.C:
            // Format: "data: <payload>\n\n"
            fmt.Fprintf(w, "data: The time is %s\n\n", t.Format(time.RFC3339))
            flusher.Flush() // Send to client immediately
        case <-r.Context().Done():
            return // Client disconnected
        }
    }
}
```

---

### Question 354: How do you manage fan-in/fan-out channel patterns?

**Answer:**

**Fan-Out:** Distribute work to multiple workers.
```go
func fanOut(ch <-chan int, workers int) {
    for i := 0; i < workers; i++ {
        go worker(ch) 
    }
}
```

**Fan-In:** Multiplex multiple channels into one.
```go
func fanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() { 
        for { c <- <-input1 } 
    }()
    go func() { 
        for { c <- <-input2 } 
    }()
    return c
}
```

**Better Fan-In (using `select` + `sync.WaitGroup` to close):**
Usually requires a "merge" function that loops over all inputs and sends to output.

---

### Question 355: How would you implement throttling on async tasks?

**Answer:**
Use a **Buffered Channel** as a semaphore (Token Bucket).

```go
// Limit to 10 concurrent requests
var semaphore = make(chan struct{}, 10)

func process(req Request) {
    semaphore <- struct{}{} // Acquire token (blocks if full)
    
    go func() {
        defer func() { <-semaphore }() // Release token
        heavyOperation(req)
    }()
}
```

**Or using `golang.org/x/time/rate`:**
```go
limiter := rate.NewLimiter(rate.Every(100*time.Millisecond), 10) // 10 reqs/sec

func handler() {
    if err := limiter.Wait(ctx); err != nil {
        return
    }
    // Proceed
}
```

---

### Question 356: How do you avoid data races when consuming messages?

**Answer:**
1. **Don't share memory:** Pass copies of data.
2. **Immutable Data:** If sharing read-only data, it's safe.
3. **Loop Variable Capture:** (Common Go Pitfall pre-1.22)

**BAD:**
```go
for msg := range messages {
    go func() {
        process(msg) // Race! 'msg' changes for all goroutines
    }()
}
```

**GOOD:**
```go
for msg := range messages {
    go func(m Message) {
        process(m)
    }(msg) // Pass by value
}
```
*Note: Go 1.22 fix this loop variable issue automatically.*

---

### Question 357: How would you implement a message queue from scratch in Go?

**Answer:**
For interview/simple use, use Channels + Mutex.

```go
type SimpleQueue struct {
    mu    sync.Mutex
    items []string
    cond  *sync.Cond
}

func NewQueue() *SimpleQueue {
    q := &SimpleQueue{}
    q.cond = sync.NewCond(&q.mu)
    return q
}

func (q *SimpleQueue) Enqueue(item string) {
    q.mu.Lock()
    defer q.mu.Unlock()
    q.items = append(q.items, item)
    q.cond.Signal() // Wake up a consumer
}

func (q *SimpleQueue) Dequeue() string {
    q.mu.Lock()
    defer q.mu.Unlock()
    
    for len(q.items) == 0 {
        q.cond.Wait() // Block until data available
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    return item
}
```

---

### Question 358: How do you implement ordered message processing in Go?

**Answer:**
In Kafka, ordering is only guaranteed **per partition**.

**Strategy:**
1. **Partitioning:** Ensure related messages (e.g., updates for UserID: 123) always go to the same partition using Partition Keys.
2. **Single Consumer per Partition:**
   If you use a worker pool inside a consumer, you lose ordering.
   
   **To fix worker pool ordering:**
   - Hash the content (e.g., UserID) to select a specific worker channel.
   
```go
// Dispatcher
workerChans := make([]chan Msg, 10) // 10 workers

func dispatch(msg Msg) {
    // Consistent Hashing
    workerID := hash(msg.Key) % 10
    workerChans[workerID] <- msg
}
```
Now, all messages for User 123 go to Worker 7 sequentially.

---

### Question 359: How do you handle large stream ingestion (100K+ msgs/sec)?

**Answer:**
1. **Batching:** Don't write 1 by 1. Accumulate 1000 messages or wait 500ms, then write.
2. **Workers:** Use a worker pool to parallelize deserialization/validation.
3. **Zero-Allocation:** Use `sync.Pool` to reuse objects.
4. **Asynchronous Ack:** If exact durability isn't critical (fire & forget), ack immediately.

**Batching Example:**
```go
func batchWriter(ch <-chan Item) {
    batch := make([]Item, 0, 1000)
    ticker := time.NewTicker(1 * time.Second)
    
    for {
        select {
        case item := <-ch:
            batch = append(batch, item)
            if len(batch) >= 1000 {
                flush(batch)
                batch = batch[:0]
            }
        case <-ticker.C:
            if len(batch) > 0 {
                flush(batch)
                batch = batch[:0]
            }
        }
    }
}
```

---

### Question 360: How do you persist in-flight streaming data?

**Answer:**
When crashing, any data in memory (`channels`) is lost.

**Solutions:**
1. **Write-Ahead Log (WAL):** Write to disk (append-only file) before processing.
2. **Ack Last:** Don't Acknowledge (commit offset) to Kafka until *after* DB write is confirmed.
3. **Graceful Shutdown:**
   - Stop accepting new messages.
   - Wait `Waitgroup`.Wait() for workers to finish.
   - Timeout if taking too long.

```go
func shutdown() {
    close(messages) // Stop producers
    wg.Wait()       // Wait for in-flight
    // Now safe to exit
}
```

---
## 🟣 System Design and Observability (Questions 361-380)

### Question 361: How do you design a rate limiter in Go?

**Answer:**
A rate limiter controls the number of requests a user can make in a given timeframe.

**Token Bucket Algorithm (Memory-based):**
Using `golang.org/x/time/rate`.

```go
package main

import (
    "golang.org/x/time/rate"
    "net/http"
    "time"
)

var limiter = rate.NewLimiter(1, 5) // 1 req/sec, burst of 5

func limit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Too many requests", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello"))
    })
    http.ListenAndServe(":8080", limit(mux))
}
```

**Distributed Rate Limiter (Redis):**
For multiple instances, use Redis (Fixed Window or Sliding Window).
- **Key:** `rate_limit:{user_id}`
- **Value:** Count
- **TTL:** 1 second

Increment key; if > limit, block.

---

### Question 362: How do you implement distributed tracing in Go?

**Answer:**
Use **OpenTelemetry (OTel)** to propagate context (Trace ID, Span ID) across services.

**Setup with Jaeger:**

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/trace"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() *sdktrace.TracerProvider {
    exporter, _ := jaeger.New(jaeger.WithCollectorEndpoint())
    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exporter),
        sdktrace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String("my-service"),
        )),
    )
    otel.SetTracerProvider(tp)
    return tp
}

func handler(w http.ResponseWriter, r *http.Request) {
    ctx, span := otel.Tracer("my-service").Start(r.Context(), "handler-span")
    defer span.End()
    
    // Pass ctx to downstream functions/services
    doWork(ctx)
}
```

---

### Question 363: How do you handle distributed transactions?

**Answer:**
Distributed transactions (across microservices) cannot use ACID properties of a single DB.

**Patterns:**
1. **Saga Pattern (Choreography):**
   - Service A publishes `OrderCreated`.
   - Service B listens, reserves inventory, publishes `InventoryReserved`.
   - If Service B fails, it publishes `InventoryFailed`.
   - Service A listens to fail event and executes **Compensating Transaction** (undo).

2. **Saga Pattern (Orchestration):**
   - A central "Orchestrator" service tells A, then B, then C what to do.
   - If any fail, the orchestrator triggers rollbacks.

3. **Two-Phase Commit (2PC):** (Avoid in microservices - blocking & slow).

---

### Question 364: What is the Circuit Breaker pattern and how to implement it?

**Answer:**
Prevents cascading failures when a downstream service is down.

**States:**
- **Closed:** Normal operation.
- **Open:** Fails immediately (after N errors).
- **Half-Open:** Allows 1 test request. If success -> Closed; else -> Open.

**Library:** `github.com/sony/gobreaker`

```go
var cb *gobreaker.CircuitBreaker

func init() {
    st := gobreaker.Settings{
        Name:        "HTTP_GET",
        MaxRequests: 1,
        Timeout:     5 * time.Second,
        ReadyToTrip: func(counts gobreaker.Counts) bool {
             return counts.ConsecutiveFailures > 3
        },
    }
    cb = gobreaker.NewCircuitBreaker(st)
}

func callService() ([]byte, error) {
    body, err := cb.Execute(func() (interface{}, error) {
        resp, err := http.Get("http://example.com")
        if err != nil { return nil, err }
        return io.ReadAll(resp.Body)
    })
    return body.([]byte), err
}
```

---

### Question 365: How do you design a notification system in Go?

**Answer:**
A system to send Email, SMS, Push notifications.

**Architecture:**
1. **API Gateway:** Receives `POST /notify`.
2. **Message Queue (Kafka/RabbitMQ):** Decouples API from senders.
   - Topics: `notify-email`, `notify-sms`.
3. **Workers:** 
   - `EmailWorker`: Consumes `notify-email`, calls SendGrid/SES.
   - `SMSWorker`: Consumes `notify-sms`, calls Twilio.
4. **Retry Logic:** Exponential backoff for failed external calls.
5. **Rate Limiting:** Throttle sends per user.

**Go Interface:**
```go
type Notifier interface {
    Send(user User, msg string) error
}

type EmailNotifier struct {} // impl...
type SMSNotifier struct {}   // impl...
```

---

### Question 366: How do you handle configuration hot-reloading?

**Answer:**
Reload config without restarting the app.

**Using `viper`:**
```go
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
    fmt.Println("Config file changed:", e.Name)
    // Re-read or update global config struct
    viper.Unmarshal(&AppConfig)
    updateConnectionPool() // if needed
})
```

**Kubernetes:**
- Update `ConfigMap`.
- K8s updates mounted file.
- Viper/fsnotify detects change -> Reload.

---

### Question 367: How do you implement health checks for microservices?

**Answer:**
Expose `/health` endpoints.

1. **Liveness Probe (Am I running?):**
   - Returns 200 OK if process is up.
   - If 500/Timeout, K8s restarts pod.

2. **Readiness Probe (Can I serve traffic?):**
   - Checks dependencies (DB, Cache).
   - Returns 200 OK only if connected.
   - If 500, K8s stops sending traffic to this pod.

**Implementation:**
```go
func health(w http.ResponseWriter, r *http.Request) {
    if err := db.Ping(); err != nil {
        w.WriteHeader(503)
        return
    }
    w.WriteHeader(200)
    w.Write([]byte("OK"))
}
```

---

### Question 368: How do you design a URL shortener in Go?

**Answer:**
**components:**
- **API:** `POST /shorten`, `GET /{short_code}`.
- **DB:** SQL (Postgres) or NoSQL (DynamoDB/Redis).
- **ID Generation:**
  - Base62 encoding of an auto-increment integer ID.
  - Or K-Sortable Unique ID (KSUID/Snowflake).

**Code Logic:**
1. `POST`: Insert URL to DB -> Get ID (1001) -> Base62(1001) = "g7" -> Return `site.com/g7`.
2. `GET /g7`: Base62Decode("g7") -> 1001 -> Select from DB -> `http.Redirect(301)`.

**Concurrency:**
- DB handles collision (Unique constraint).
- Read-heavy -> Cache in Redis (`g7` -> `full_url`).

---

### Question 369: How do you debug high CPU usage in a Go app?

**Answer:**
Use `pprof`.

1. **Enable pprof:**
   ```go
   import _ "net/http/pprof"
   go func() { log.Println(http.ListenAndServe("localhost:6060", nil)) }()
   ```

2. **Capture Profile:**
   ```bash
   go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
   ```

3. **Analyze:**
   - `top`: Shows functions using most CPU.
   - `list <func>`: Shows source code with CPU time per line.
   - `web`: Visual graph in browser.

**Common Culprits:**
- Tight loops.
- Excessive Garbage Collection (check allocations).
- Serialization (JSON) in hot paths.

---

### Question 370: How do you debug a memory leak in Go?

**Answer:**
Memory leaks in Go are usually:
1. **Goroutine Leaks:** Goroutines blocked forever (waiting on nil channel, unclosed body).
2. **Retained References:** Global map growing forever, Slice with small window into large array.

**Debugging:**
1. **Capture Heap Profile:**
   ```bash
   go tool pprof http://localhost:6060/debug/pprof/heap
   ```
2. **Compare Profiles (diff):**
   Capture heap at T1 and T2, then `pprof -base heap1 heap2`.
3. **Analyze:** Look for objects with high `inuse_space` or `inuse_objects`.

---

### Question 371: What is the Sidecar pattern and how is it used with Go?

**Answer:**
Deploying a helper container alongside the main Go application container in the same Pod.

**Use Cases:**
- **Logging:** Sidecar reads log files, forwards to Splunk.
- **Proxy:** Envoy/Istio sidecar handles mTLS, circuit breaking, tracing (Service Mesh).
- **Config:** Sidecar watches remote config and updates local file.

**Go Impact:**
- Go app talks to `localhost` for these services (e.g., `localhost:5432` proxy to DB).
- Simplifies Go code (offloads infrastructure concerns).

---

### Question 372: How do you design a job scheduler in Go?

**Answer:**
**Simple (In-Memory):**
- Ticker + Goroutine.
```go
ticker := time.NewTicker(1 * time.Hour)
for range ticker.C {
    go cleanup()
}
```

**Distributed (Robust):**
- Use a library like `gocron` or `robfig/cron`.
- **Leader Election:** Only one instance runs the job (using Redis lock or Etcd).
- **Persistent Queue:** If job fails, retry later (patterns like RabbitMQ Delayed Exchange).

**Leader Election Pattern:**
```go
// Redis SetNX (Set if Not Exists) with TTL
success, _ := redis.SetNX(ctx, "leader_lock", "my-id", 10*time.Second).Result()
if success {
    runJob()
    // Refresh lock periodically (Heartbeat)
}
```

---

### Question 373: How do you implement API versioning?

**Answer:**
1. **URL Path (Most Common):** `/api/v1/users`
   - Easy to see and route.
   - Using `http.NewServeMux`:
     ```go
     mux.Handle("/api/v1/", v1Handler)
     mux.Handle("/api/v2/", v2Handler)
     ```

2. **Header:** `Accept: application/vnd.myapi.v1+json`
   - Cleaner URLs.
   - Harder to test in browser.

3. **Query Param:** `/users?v=1`

**Best Practice:**
Keep logic separate (`package v1`, `package v2`) to avoid `if v1 {...} else {...}` spaghetti code.

---

### Question 374: What is Context Propagation?

**Answer:**
Passing request-scoped values (User ID, Trace ID, Auth Token) through the call chain using `context.Context`.

**Example:**
1. Middleware extracts `Trace-ID` header -> puts in `ctx`.
2. Handler calls Database -> passes `ctx`.
3. DB Driver (if instrumented) uses `ctx` to log Trace ID or handle cancellation.

**Custom Context Value:**
```go
type key int
const userKey key = 0

// Set
ctx = context.WithValue(ctx, userKey, "Alice")

// Get
user, ok := ctx.Value(userKey).(string)
```
*Note: Only use context for request-scoped data, not for optional parameters.*

---

### Question 375: How to handle 10 million concurrent WebSocket connections?

**Answer:**
The "C10M" problem.

1. **Kernel Tuning:**
   - Increase Max Open Files (`ulimit -n`).
   - Tune TCP buffer sizes (`net.ipv4.tcp_rmem`, `wmem`).

2. **Go Optimization:**
   - **Goroutines:** 10M goroutines ≈ 20-40GB RAM (2KB each). Doable on large implementation.
   - **Epoll (Advanced):** Instead of 1 goroutine per conn, use `syscall.Epoll` (on Linux) to manage connections event-driven (Library: `gnet` or `evio`).

3. **Architecture:**
   - **User Level Sharding:** Load balancer distributes users to different server nodes.
   - **State:** Keep connection state minimal.

---

### Question 376: How do you secure internal gRPC services?

**Answer:**
1. **mTLS (Mutual TLS):**
   - Both Client and Server authenticate each other using certificates.
   - Best for zero-trust internal networks.

2. **Token Authentication (Interceptors):**
   - Client sends JWT in Metadata (`authorization` header).
   - Server Interceptor validates JWT.

**Code (Unary Interceptor):**
```go
func authInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
    md, _ := metadata.FromIncomingContext(ctx)
    token := md["authorization"][0]
    
    if !valid(token) {
        return nil, status.Error(codes.Unauthenticated, "invalid token")
    }
    return handler(ctx, req)
}
```

---

### Question 377: Explain the Outbox Pattern.

**Answer:**
Ensures **Atomicity** when writing to DB and publishing an Event.

**Problem:** 
- `DB.Insert(Order)` succeeds.
- `Kafka.Publish(OrderCreated)` fails.
- Inconsistency!

**Solution (Outbox):**
1. Start Transaction.
2. `DB.Insert(Order)`.
3. `DB.Insert(Outbox, {Event: OrderCreated})`.
4. Commit Transaction. (Atomic).
5. **Separate Process (CDC/Poller):** Reads `Outbox` table and publishes to Kafka.
6. Delete from `Outbox` after publish.

---

### Question 378: How do you implement checking for race conditions in CI?

**Answer:**
Use the Go Race Detector.

**Command:**
```bash
go test -race ./...
```

**In CI (GitHub Actions):**
```yaml
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
      - run: go test -race -v ./...
```
**Caveat:** Race detector slows down execution (5-10x) and uses more memory. Run it on integration tests or a subset of tests if full suite is too slow.

---

### Question 379: How do you optimize Go garbage collection for low latency?

**Answer:**
1. **Reduce Allocations:**
   - Use `sync.Pool`.
   - Pre-allocate slices/maps (`make([]int, 0, 1000)`).
   - Use value types instead of pointers where appropriate (less work for scanner).

2. **GOGC Tuning:**
   - `GOGC=200`: GC runs less often (uses more RAM).
   - `GOGC=off`: Disable manual GC (dangerous).

3. **Memory Ballast (Legacy trick):**
   - Allocate a huge byte array (e.g., 1GB) on startup.
   - Increases heap size, so GC triggers less frequently (since it triggers based on % growth).
   - *Note: Go 1.19 `SetMemoryLimit` (Soft Limit) is the modern way.*

---

### Question 380: What is Semantic Versioning in Go modules?

**Answer:**
Go modules strictly follow SemVer (`vMajor.Minor.Patch`).

- **v1.x.x:** Public API is stable.
- **v2.0.0:** Breaking changes.
  - **Import Path changes:** `github.com/user/lib/v2`.
  - Directory structure usually involves a `v2` folder or `go.mod` change.
  
**Pseudo-version:**
Using a specific commit hash: `v0.0.0-20230101-abcdef123456`.

**Direct/Indirect:**
- `// indirect` in `go.mod`: Dependency of a dependency.

---
## 🔴 Security and Advanced Testing (Questions 381-400)

### Question 381: How do you prevent SQL injection in Go?

**Answer:**
Always use **Parameterized Queries** (Prepared Statements). Never concatenate strings into queries.

**Vulnerable:**
```go
query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", unsafeInput)
db.Query(query) // DANGER!
```

**Secure:**
```go
// Use ? placeholders (or $1 for Postgres)
query := "SELECT * FROM users WHERE name = ?"
db.Query(query, safeInput)
```
The database driver escapes the input safely.

**ORMs:**
GORM and sqlx handle this automatically if used correctly (avoid `db.Raw` with string concatenation).

---

### Question 382: What are some common security vulnerabilities in Go apps?

**Answer:**
1. **Goroutine Leaks:** Denial of Service (DoS) by exhausting resources.
2. **Directory Traversal:** Using `filepath.Join` with user input without cleaning.
3. **Data Races:** Concurrent access causing corrupted state.
4. **Insecure Randomness:** Using `math/rand` for tokens (use `crypto/rand`).
5. **Cross-Site Scripting (XSS):** Rendering User Input in HTML templates without escaping.

---

### Question 383: How do you implement Secure Password Hashing?

**Answer:**
Use **bcrypt** or **Argon2**. Never use MD5 or SHA1/SHA256 directly.

**Library:** `golang.org/x/crypto/bcrypt`

```go
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

---

### Question 384: How do you validate input in Go APIs?

**Answer:**
1. **Manual Validation:** Checks in handlers.
2. **Validator Library:** `github.com/go-playground/validator` (standard with Gin).

```go
type User struct {
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"gte=18"`
}

func validateUser(u User) error {
    validate := validator.New()
    return validate.Struct(u)
}
```

---

### Question 385: How do you implement JWT authentication?

**Answer:**
**Library:** `github.com/golang-jwt/jwt/v5`

**Sign (Login):**
```go
func createToken(user string) (string, error) {
    claims := jwt.MapClaims{
        "user": user,
        "exp":  time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte("my-secret-key"))
}
```

**Verify (Middleware):**
Parse the token string, provide the key callback, and check validity.

---

### Question 386: How do you handle secrets securely in Go?

**Answer:**
1. **Never commit to Git.**
2. **Environment Variables:** `os.Getenv("SECRET")`.
3. **Hashicorp Vault:** Fetch secrets at runtime.
4. **Memory Safety:** Use `[]byte` instead of `string` for sensitive data (can be wiped), though Go's GC makes this hard to guarantee.
5. **Avoid Logging:** Sanitize logs (don't log struct with "Password" field).

---

### Question 387: What is CSRF and how to mitigate it in Go?

**Answer:**
**CSRF (Cross-Site Request Forgery):** Attacker tricks user into performing action on trusted site (using cookies).

**Mitigation:**
1. **CSRF Tokens:**
   - Server sends random token in Cookie/HTML.
   - Client must send it back in Header `X-CSRF-Token`.
   - Library: `github.com/gorilla/csrf`.

2. **SameSite Cookies:**
   - Set `SameSite=Strict` or `Lax` on Auth cookies.

```go
http.SetCookie(w, &http.Cookie{
    Name:     "session_token",
    Value:    "xyz",
    SameSite: http.SameSiteStrictMode,
    Secure:   true,
})
```

---

### Question 388: How do you test RESTful APIs in Go?

**Answer:**
Use `net/http/httptest`.

```go
func TestHealthCheck(t *testing.T) {
    // 1. Create Request
    req, _ := http.NewRequest("GET", "/health", nil)
    
    // 2. Create ResponseRecorder
    w := httptest.NewRecorder()
    
    // 3. Call Handler
    HealthHandler(w, req)
    
    // 4. Assert
    if w.Code != http.StatusOK {
        t.Errorf("Expected 200, got %d", w.Code)
    }
    if w.Body.String() != "OK" {
        t.Errorf("Unexpected body: %v", w.Body.String())
    }
}
```

---

### Question 389: What are Table-Driven Tests?

**Answer:**
The idiomatic way to write tests in Go. Data-driven approach.

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -2, -3},
        {"mixed", -1, 1, 0},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            result := Add(tc.a, tc.b)
            if result != tc.expected {
                t.Errorf("got %d, want %d", result, tc.expected)
            }
        })
    }
}
```

---

### Question 390: How do you mock HTTP calls in tests?

**Answer:**
1. **Interface Injection:** Define `HTTPClient` interface, mock implementation.
2. **httptest.Server:** Start a real internal server that mocks external responses.

```go
func TestExternalCall(t *testing.T) {
    // Mock Server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(200)
        w.Write([]byte(`{"status": "ok"}`))
    }))
    defer server.Close()

    // Use server.URL as API endpoint
    api := NewAPI(server.URL)
    result := api.Call()
    
    if result != "ok" { t.Fail() }
}
```

---

### Question 391: What are Flaky Tests and how to identify them?

**Answer:**
Tests that sometimes pass and sometimes fail without code changes.

**Causes:**
- Race conditions.
- Network dependencies (external API calls).
- Random number generation.
- Logic depending on time (e.g., `time.Now()`).

**Fixes:**
- Run with `-race`.
- Use Mocks for external calls.
- Run repeatedly: `go test -count=100 ./...`.

---

### Question 392: How do you generate test coverage reports?

**Answer:**
```bash
# Run tests and generate profile
go test -coverprofile=coverage.out ./...

# View detailed HTML report
go tool cover -html=coverage.out
```
This opens a browser showing exact lines covered (green) and missed (red).

---

### Question 393: What is Golden File Testing?

**Answer:**
Used for complex outputs (large JSON, HTML, Images). Instead of hardcoding the expected string, compare against a saved file.

```go
func TestJSON(t *testing.T) {
    got := generateBigJSON()
    
    if *update { // Flag to update golden file
        os.WriteFile("testdata/golden.json", got, 0644)
    }
    
    want, _ := os.ReadFile("testdata/golden.json")
    if string(got) != string(want) {
        t.Errorf("Mismatch found")
    }
}
```
Run `go test -update` to regenerate the expected output.

---

### Question 394: How do you mock database interactions?

**Answer:**
1. **Library: `go-sqlmock`:** Mock `sql/driver` behavior (rows, errors) without a real DB.
2. **Docker (Testcontainers):** Spin up a real Postgres container for integration tests (slower but more accurate).
3. **Repository Pattern:** Mock the `UserRepository` interface.

**Using sqlmock:**
```go
db, mock, _ := sqlmock.New()
mock.ExpectQuery("SELECT name FROM users").
    WithArgs(1).
    WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Alice"))

user, _ := GetUser(db, 1)
```

---

### Question 395: What is Fuzz Testing in Go?

**Answer:**
Native in Go 1.18+. Generates random inputs to find edge cases/crashes.

```go
func FuzzReverse(f *testing.F) {
    f.Add("hello") // Seed corpus
    
    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
    })
}
```
Run: `go test -fuzz=Fuzz`

---

### Question 396: How do you benchmark Go code?

**Answer:**
Write a function starting with `Benchmark` in `_test.go`.

```go
func BenchmarkConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = fmt.Sprintf("hello %s", "world")
    }
}
```
Run: `go test -bench=.`
Look for `ns/op` (nanoseconds per operation) and `allocs/op`.

---

### Question 397: How do you handle Time in unit tests?

**Answer:**
Don't use `time.Now()` directly in logic.
Dependency Inject a `Clock` interface.

```go
type Clock interface {
    Now() time.Time
}

type RealClock struct{}
func (RealClock) Now() time.Time { return time.Now() }

type MockClock struct { t time.Time }
func (m MockClock) Now() time.Time { return m.t }
```
Now you can freeze time in tests!

---

### Question 398: What is `go:embed` and how does it help?

**Answer:**
Go 1.16+ feature to bundle static assets (HTML, SQL, Config) into the binary.

```go
import "embed"

//go:embed templates/*
var templateFS embed.FS

//go:embed config.json
var configFile []byte
```
Useful for:
- Single-binary deployments (Binary + Frontend).
- Database migration files.

---

### Question 399: What is the purpose of `init()` function?

**Answer:**
Runs automatically before `main()`.

**Uses:**
- Initializing global variables.
- Registering drivers (e.g., `image`, `database/sql`).

**Caveats:**
- Hard to test (side effects).
- Order of execution between files is alphabetical (tricky).
- Avoid complex logic or dependencies in `init()`.

---

### Question 400: Explain the Go memory model (briefly).

**Answer:**
Specifies conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes in another goroutine.

**Key Rule:** **"HAPPENS BEFORE"**
- A send on a channel **happens before** the corresponding receive completes.
- A lock release **happens before** the next acquire.
- `go` statement **happens before** the goroutine's execution.

If you don't use synchronization (Channels/Mutex), you have a **Data Race**, and visibility is not guaranteed!

---
