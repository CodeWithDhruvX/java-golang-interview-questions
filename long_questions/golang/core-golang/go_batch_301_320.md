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
