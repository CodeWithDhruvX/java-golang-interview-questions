# Go Intermediate — Concurrency Patterns (Mid-Tier Product Companies)

> **30% Gap Coverage — Part 1 of 2**
> Topics: `sync.Mutex`, `sync.RWMutex`, `sync.WaitGroup`, `sync.Once`, `sync.Cond`, `sync/atomic`, Worker Pool, Fan-In/Fan-Out, Pipeline, Goroutine Lifecycle

---

## Section 1: sync.Mutex & sync.RWMutex (Q1–Q12)

### 1. Mutex Protects Shared State
**Q: What is the bug and the fix?**
```go
package main
import (
    "fmt"
    "sync"
)

var counter int

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    counter++ // DATA RACE
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go increment(&wg)
    }
    wg.Wait()
    fmt.Println(counter) // not reliably 1000
}
```
**A:** **Data race** — concurrent goroutines read-modify-write `counter` without synchronization. Final value is unpredictable.
**Fix:** Add `var mu sync.Mutex`, then `mu.Lock(); counter++; mu.Unlock()` inside `increment`.

---

### 2. Mutex Must Not Be Copied
**Q: What is the bug?**
```go
package main
import "sync"

type Safe struct {
    mu sync.Mutex
    v  int
}

func process(s Safe) { // copied by value!
    s.mu.Lock()
    s.v++
    s.mu.Unlock()
}

func main() {
    s := Safe{}
    process(s)
}
```
**A:** `sync.Mutex` contains internal state. Copying it copies that state, creating two mutexes with inconsistent lock states — undefined behavior. Detected by `go vet`.
**Fix:** `func process(s *Safe)` — always pass mutex-containing structs by pointer.

---

### 3. Defer Unlock — Correct Pattern
**Q: Why is this pattern preferred?**
```go
package main
import "sync"

type Cache struct {
    mu   sync.Mutex
    data map[string]string
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock() // always unlocks, even on panic
    c.data[key] = value
}
```
**A:** `defer c.mu.Unlock()` guarantees the mutex is released even if a panic occurs inside the function body. This prevents a deadlock from a leaked lock.

---

### 4. sync.RWMutex — Read vs Write Lock
**Q: What is the output and why is RWMutex better here?**
```go
package main
import (
    "fmt"
    "sync"
)

type Store struct {
    mu   sync.RWMutex
    data map[string]int
}

func (s *Store) Get(key string) int {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return s.data[key]
}

func (s *Store) Set(key string, val int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data[key] = val
}

func main() {
    s := Store{data: map[string]int{}}
    s.Set("x", 42)
    fmt.Println(s.Get("x"))
}
```
**A:** `42`. `RWMutex` allows **multiple concurrent readers** (`RLock`) but only one writer (`Lock`). For read-heavy workloads this is significantly faster than a plain `Mutex`.

---

### 5. Double-Lock Deadlock
**Q: Does this deadlock?**
```go
package main
import "sync"

func main() {
    var mu sync.Mutex
    mu.Lock()
    mu.Lock() // goroutine tries to acquire lock it already holds
}
```
**A:** **Yes, deadlock.** Go's `sync.Mutex` is **not reentrant**. A goroutine cannot lock a mutex it already holds — it will wait forever.

---

### 6. RLock Upgrade Deadlock
**Q: Does this deadlock?**
```go
package main
import "sync"

func main() {
    var rw sync.RWMutex
    rw.RLock()
    rw.Lock() // tries to upgrade RLock to Lock
}
```
**A:** **Yes, deadlock.** You cannot upgrade a read lock to a write lock. `Lock()` waits for all readers to finish, but our goroutine is still holding an `RLock` — circular wait.

---

### 7. Mutex with Condition — sync.Cond
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    cond := sync.NewCond(&mu)
    ready := false

    go func() {
        mu.Lock()
        ready = true
        cond.Signal()
        mu.Unlock()
    }()

    mu.Lock()
    for !ready {
        cond.Wait() // atomically releases mu and suspends
    }
    fmt.Println("ready:", ready)
    mu.Unlock()
}
```
**A:** `ready: true`. `cond.Wait()` atomically releases the mutex and suspends. `cond.Signal()` wakes one waiter. Used for producer-consumer coordination.

---

### 8. sync.Map — Concurrent Safe Map
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var m sync.Map
    m.Store("key", 42)

    val, ok := m.Load("key")
    fmt.Println(val, ok)

    _, ok2 := m.Load("missing")
    fmt.Println(ok2)
}
```
**A:**
```
42 true
false
```
`sync.Map` is safe for concurrent use without additional locking. Best for cases where keys are mostly written once and read many times.

---

### 9. sync.Map Range Iteration
**Q: What does Range do?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var m sync.Map
    m.Store("a", 1)
    m.Store("b", 2)
    m.Store("c", 3)

    m.Range(func(key, value interface{}) bool {
        fmt.Println(key, value)
        return true // returning false stops iteration
    })
}
```
**A:** Prints all key-value pairs (order not guaranteed). Returning `false` from the callback stops iteration early — useful for search.

---

### 10. LoadOrStore — Atomic Check-Then-Set
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var m sync.Map

    actual, loaded := m.LoadOrStore("key", "first")
    fmt.Println(actual, loaded)

    actual2, loaded2 := m.LoadOrStore("key", "second")
    fmt.Println(actual2, loaded2)
}
```
**A:**
```
first false
first true
```
First call stores and returns `"first"`, loaded=false. Second call finds key already exists, returns `"first"` unchanged, loaded=true.

---

### 11. Embedding Mutex in Struct — Anti-Pattern
**Q: What is wrong?**
```go
package main
import "sync"

type Counter struct {
    sync.Mutex
    count int
}

func (c *Counter) Inc() {
    c.Lock()
    defer c.Unlock()
    c.count++
}
```
**A:** **This compiles and works**, but embedding `sync.Mutex` promotes `Lock()`/`Unlock()` to the public API of `Counter` — callers can accidentally lock/unlock the internal mutex. Prefer a named field `mu sync.Mutex`.

---

### 12. TryLock (Go 1.18+)
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    mu.Lock()

    ok := mu.TryLock() // non-blocking attempt
    fmt.Println("acquired:", ok)

    mu.Unlock()
    ok2 := mu.TryLock()
    fmt.Println("acquired:", ok2)
    mu.Unlock()
}
```
**A:**
```
acquired: false
acquired: true
```
`TryLock` returns immediately — `false` if the mutex is held, `true` if it successfully acquired the lock.

---

## Section 2: sync/atomic (Q13–Q20)

### 13. Atomic Counter
**Q: Why does this always print 1000?**
```go
package main
import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var count int64
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            atomic.AddInt64(&count, 1)
        }()
    }
    wg.Wait()
    fmt.Println(count)
}
```
**A:** `1000`. `atomic.AddInt64` is a CPU-level atomic operation — no mutex needed for simple integer increments. Fastest way to build a concurrent counter.

---

### 14. atomic.LoadInt64 / StoreInt64
**Q: Why use Load/Store instead of direct read?**
```go
package main
import (
    "fmt"
    "sync/atomic"
)

func main() {
    var flag int64 = 0

    // Another goroutine might do:
    atomic.StoreInt64(&flag, 1)

    // Safe read:
    val := atomic.LoadInt64(&flag)
    fmt.Println(val)
}
```
**A:** `1`. Without atomic ops, a direct read of `flag` from another goroutine is a **data race** on 64-bit values (not guaranteed to be read atomically on all platforms). `LoadInt64`/`StoreInt64` ensure memory visibility.

---

### 15. atomic.CompareAndSwap
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync/atomic"
)

func main() {
    var val int64 = 10

    swapped := atomic.CompareAndSwapInt64(&val, 10, 20) // if val==10, set to 20
    fmt.Println(swapped, val)

    swapped2 := atomic.CompareAndSwapInt64(&val, 10, 30) // val is now 20, not 10
    fmt.Println(swapped2, val)
}
```
**A:**
```
true 20
false 20
```
CAS is the building block for lock-free data structures. It atomically checks the expected value and only updates if it matches.

---

### 16. atomic.Value — Store/Load Any Type
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync/atomic"
)

func main() {
    var av atomic.Value

    av.Store(map[string]int{"hits": 100})

    m := av.Load().(map[string]int)
    fmt.Println(m["hits"])
}
```
**A:** `100`. `atomic.Value` stores and loads any value atomically without a mutex. Ideal for config hot-reloading patterns.

---

### 17. atomic.Value Panic on Type Change
**Q: What happens?**
```go
package main
import (
    "sync/atomic"
)

func main() {
    var av atomic.Value
    av.Store(42)
    av.Store("hello") // panic!
}
```
**A:** **Panic.** `atomic.Value` requires all stored values to be the **same concrete type**. Storing a different type after the first store panics.

---

### 18. Mutex vs Atomic — When to Use Which
**Q: Which is more appropriate here?**
```go
// Option A: Mutex
type Stats struct {
    mu      sync.Mutex
    total   int64
    errors  int64
    latency float64 // float can't be atomically updated
}

// Option B: Atomic for simple integers
var total  int64
var errors int64
// atomic.AddInt64(&total, 1)
// atomic.AddInt64(&errors, 1)
```
**A:** Use **atomic** for simple integer counters (fastest, no contention). Use **Mutex** when: updating multiple related fields together, updating non-integer types (float64, structs, slices), or when consistency between fields matters.

---

### 19. Spinlock Anti-Pattern
**Q: What is the problem with this?**
```go
package main
import (
    "runtime"
    "sync/atomic"
)

var locked int32

func spinLock() {
    for !atomic.CompareAndSwapInt32(&locked, 0, 1) {
        runtime.Gosched() // yield
    }
}

func spinUnlock() {
    atomic.StoreInt32(&locked, 0)
}
```
**A:** This implements a spinlock — it works but is **CPU-wasteful** under contention (busy-waiting). In Go, prefer `sync.Mutex` which uses the runtime scheduler and blocks efficiently instead of spinning.

---

### 20. atomic.Pointer (Go 1.19+)
**Q: What does this safely do?**
```go
package main
import (
    "fmt"
    "sync/atomic"
)

func main() {
    type Config struct{ Debug bool }

    var p atomic.Pointer[Config]
    p.Store(&Config{Debug: false})

    cfg := p.Load()
    fmt.Println(cfg.Debug)

    p.Store(&Config{Debug: true})
    fmt.Println(p.Load().Debug)
}
```
**A:**
```
false
true
```
`atomic.Pointer[T]` is a type-safe atomic pointer (Go 1.19+). Useful for atomically swapping configuration structs without a mutex.

---

## Section 3: Advanced Goroutine & Channel Patterns (Q21–Q40)

### 21. Done Channel — Goroutine Cancellation
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "time"
)

func worker(done <-chan struct{}) {
    for {
        select {
        case <-done:
            fmt.Println("worker stopped")
            return
        default:
            // do work
        }
        time.Sleep(10 * time.Millisecond)
    }
}

func main() {
    done := make(chan struct{})
    go worker(done)
    time.Sleep(50 * time.Millisecond)
    close(done) // signal all goroutines to stop
    time.Sleep(20 * time.Millisecond)
}
```
**A:** `worker stopped`. Closing a `done` channel is idiomatic for broadcasting cancellation to any number of goroutines simultaneously.

---

### 22. Pipeline Pattern
**Q: What is the output?**
```go
package main
import "fmt"

func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func main() {
    for v := range square(generate(2, 3, 4)) {
        fmt.Print(v, " ")
    }
}
```
**A:** `4 9 16 `. Classic Go pipeline: each stage reads from one channel and writes to the next. Stages run concurrently.

---

### 23. Fan-Out Pattern
**Q: What does this pattern do?**
```go
package main
import (
    "fmt"
    "sync"
)

func fanOut(in <-chan int, workers int) []<-chan int {
    outs := make([]<-chan int, workers)
    for i := range outs {
        out := make(chan int)
        outs[i] = out
        go func(o chan int) {
            for v := range in {
                o <- v * 2
            }
            close(o)
        }(out)
    }
    return outs
}

func main() {
    in := make(chan int, 3)
    in <- 1; in <- 2; in <- 3
    close(in)

    outs := fanOut(in, 2)
    var wg sync.WaitGroup
    for _, ch := range outs {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c {
                fmt.Print(v, " ")
            }
        }(ch)
    }
    wg.Wait()
}
```
**A:** Prints doubled values (order may vary). Fan-Out distributes work from one channel across multiple worker goroutines.

---

### 24. Fan-In (Merge) Pattern
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    output := func(c <-chan int) {
        defer wg.Done()
        for v := range c {
            out <- v
        }
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }
    go func() { wg.Wait(); close(out) }()
    return out
}

func toChan(vals ...int) <-chan int {
    c := make(chan int, len(vals))
    for _, v := range vals {
        c <- v
    }
    close(c)
    return c
}

func main() {
    merged := merge(toChan(1, 2), toChan(3, 4))
    for v := range merged {
        fmt.Print(v, " ")
    }
}
```
**A:** `1 2 3 4` (order may vary). Fan-In merges multiple channels into one. The merged channel closes only after all inputs close.

---

### 25. Worker Pool Pattern
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

func workerPool(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        results <- j * j
    }
}

func main() {
    jobs := make(chan int, 5)
    results := make(chan int, 5)

    var wg sync.WaitGroup
    for w := 0; w < 3; w++ { // 3 workers
        wg.Add(1)
        go workerPool(jobs, results, &wg)
    }

    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    go func() { wg.Wait(); close(results) }()

    var sum int
    for r := range results {
        sum += r
    }
    fmt.Println(sum) // 1+4+9+16+25
}
```
**A:** `55`. Worker pool: fixed number of goroutines process jobs from a shared channel. Essential pattern for rate-limiting and resource control.

---

### 26. Timeout on Channel Receive
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)

    select {
    case msg := <-ch:
        fmt.Println(msg)
    case <-time.After(100 * time.Millisecond):
        fmt.Println("timeout!")
    }
}
```
**A:** `timeout!`. `time.After` returns a channel that receives after the duration — used with `select` to implement operation timeouts.

---

### 27. Channel as Semaphore
**Q: What does this pattern control?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    sem := make(chan struct{}, 3) // max 3 concurrent
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            sem <- struct{}{} // acquire
            fmt.Printf("working: %d\n", id)
            <-sem // release
        }(i)
    }
    wg.Wait()
}
```
**A:** At most 3 goroutines execute the guarded section concurrently. A buffered channel of size N is a classic **counting semaphore** — limits parallel execution.

---

### 28. Or-Done Pattern
**Q: What does this safely read from a channel with cancellation?**
```go
func orDone(done, c <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for {
            select {
            case <-done:
                return
            case v, ok := <-c:
                if !ok {
                    return
                }
                select {
                case out <- v:
                case <-done:
                }
            }
        }
    }()
    return out
}
```
**A:** Safely reads from channel `c` but stops when `done` is closed — combining cancellation with normal channel processing. The double-select ensures we don't block writing to `out` when done fires.

---

### 29. Goroutine Leak — Stuck Send
**Q: What is the goroutine leak?**
```go
package main
import "fmt"

func doWork() <-chan int {
    ch := make(chan int)
    go func() {
        // if nobody receives, this goroutine leaks forever
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
    return ch
}

func main() {
    ch := doWork()
    fmt.Println(<-ch) // reads only one value
    // goroutine is stuck waiting to send the remaining 9 values
}
```
**A:** The goroutine leaks — it is stuck trying to send values that nobody reads. Fix: use a `done` channel or buffered channel, or ensure the caller drains fully.

---

### 30. select with Multiple Ready Channels
**Q: Is the choice deterministic?**
```go
package main
import "fmt"

func main() {
    ch1 := make(chan string, 1)
    ch2 := make(chan string, 1)
    ch1 <- "one"
    ch2 <- "two"

    select {
    case msg := <-ch1:
        fmt.Println(msg)
    case msg := <-ch2:
        fmt.Println(msg)
    }
}
```
**A:** **Not deterministic.** When multiple cases in a `select` are ready simultaneously, Go picks one **uniformly at random**. Each run may print `one` or `two`.

---

### 31. Closing a Channel Broadcasts to All Receivers
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    ch := make(chan struct{})
    var wg sync.WaitGroup

    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            <-ch // all block here
            fmt.Printf("goroutine %d unblocked\n", id)
        }(i)
    }

    close(ch) // unblocks ALL goroutines simultaneously
    wg.Wait()
}
```
**A:** All 3 goroutines print their message (order varies). Closing a channel is the only way to **broadcast** to multiple receivers at once.

---

### 32. WaitGroup Add Must Precede go
**Q: What is the race condition?**
```go
package main
import "sync"

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        go func() {
            wg.Add(1) // WRONG: Add called inside goroutine
            defer wg.Done()
        }()
    }
    wg.Wait() // may reach 0 before Add is called
}
```
**A:** `wg.Wait()` may complete before goroutines call `wg.Add(1)`, causing the main function to return early. **Always call `wg.Add(n)` before starting goroutines.**

---

### 33. Goroutine with Panic Recovery
**Q: Does main survive?**
```go
package main
import "fmt"

func safeGo(f func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("recovered:", r)
            }
        }()
        f()
    }()
}

func main() {
    safeGo(func() { panic("oops") })
    // give goroutine time to run
    fmt.Scanln()
}
```
**A:** `recovered: oops`. Wrapping goroutine launches in a helper that recovers panics prevents the whole program from crashing — production pattern for background workers.

---

### 34. Buffered Channel as Task Queue
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    queue := make(chan func(), 3)
    queue <- func() { fmt.Println("task 1") }
    queue <- func() { fmt.Println("task 2") }
    queue <- func() { fmt.Println("task 3") }
    close(queue)

    for task := range queue {
        task()
    }
}
```
**A:**
```
task 1
task 2
task 3
```
A buffered channel of functions is a lightweight task queue. Values are buffered without needing a receiver immediately.

---

### 35. sync.Once for Singleton
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

type DB struct{ name string }

var (
    instance *DB
    once     sync.Once
)

func getDB() *DB {
    once.Do(func() {
        instance = &DB{name: "main"}
        fmt.Println("DB initialized")
    })
    return instance
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            db := getDB()
            _ = db
        }()
    }
    wg.Wait()
}
```
**A:** `DB initialized` printed **exactly once**, regardless of how many goroutines call `getDB()`. `sync.Once` is the idiomatic Go singleton pattern.

---

### 36. Goroutine ID (Not in stdlib)
**Q: Does Go expose goroutine IDs?**
```go
// There is NO standard way to get goroutine ID in Go.
// This is intentional — the Go team discourages goroutine-local storage.
// For debugging, use: runtime/pprof or go tool trace
import "runtime"

func goroutineCount() int {
    return runtime.NumGoroutine()
}
```
**A:** Go deliberately does **not** expose goroutine IDs in the standard library (unlike thread IDs in Java). `runtime.NumGoroutine()` gives the count but not IDs.

---

### 37. select Starving a Case
**Q: What is the problem?**
```go
package main
import "fmt"

func main() {
    highPriority := make(chan int, 100)
    lowPriority  := make(chan int, 100)

    for i := 0; i < 100; i++ {
        highPriority <- i
        lowPriority  <- i
    }

    for i := 0; i < 100; i++ {
        select {
        case v := <-highPriority:
            fmt.Print("H:", v, " ")
        case v := <-lowPriority:
            fmt.Print("L:", v, " ")
        }
    }
}
```
**A:** Both channels print randomly — Go `select` gives no priority. To enforce priority, check `highPriority` first with a non-blocking select before falling to `lowPriority`.

---

### 38. for-select Loop with Done
**Q: What is the canonical pattern for a long-running goroutine?**
```go
func processLoop(done <-chan struct{}, data <-chan int) {
    for {
        select {
        case <-done:
            return // graceful shutdown
        case v, ok := <-data:
            if !ok {
                return // channel closed
            }
            _ = v // process v
        }
    }
}
```
**A:** This is the canonical **for-select loop** pattern in Go: handles both a `done` signal and incoming data. The `ok` check handles closed channels. Used in every production Go service.

---

### 39. Goroutine Stack Size
**Q: What is the initial goroutine stack size and how does it grow?**
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    // Each goroutine starts with a small stack (typically 2KB-8KB)
    // It grows dynamically up to 1GB (default max) via stack copying
    
    var stats runtime.MemStats
    runtime.ReadMemStats(&stats)
    
    before := stats.StackInuse
    
    // Spawn many goroutines
    done := make(chan struct{})
    for i := 0; i < 10000; i++ {
        go func() { <-done }()
    }
    
    runtime.ReadMemStats(&stats)
    fmt.Printf("Stack in use: ~%dKB for 10k goroutines\n", 
        (stats.StackInuse-before)/1024)
    close(done)
}
```
**A:** Goroutines start with ~2KB stack (vs 1MB for OS threads). They grow and shrink dynamically via **stack copying** (Go changed from segmented stacks to contiguous copying stacks in Go 1.4). This is why you can run millions of goroutines.

---

### 40. GOMAXPROCS Effect
**Q: What is the output and what does GOMAXPROCS control?**
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // query without changing

    // Set to 1: goroutines are concurrent but not parallel
    old := runtime.GOMAXPROCS(1)
    fmt.Println("old GOMAXPROCS:", old)
}
```
**A:** Output depends on machine (e.g., `CPUs: 8`, `GOMAXPROCS: 8`). `GOMAXPROCS` controls **how many OS threads can execute Go code simultaneously**. Default is `runtime.NumCPU()` since Go 1.5.

---

## Section 4: Advanced Patterns (Q41–Q52)

### 41. errgroup for Parallel Tasks with Error Collection
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
    "golang.org/x/sync/errgroup"
)

func main() {
    var g errgroup.Group

    g.Go(func() error { return nil })
    g.Go(func() error { return errors.New("task failed") })
    g.Go(func() error { return nil })

    if err := g.Wait(); err != nil {
        fmt.Println("error:", err)
    }
}
```
**A:** `error: task failed`. `errgroup.Group` runs goroutines concurrently and returns the **first non-nil error** after all complete. Standard pattern for parallel API calls / DB queries.

---

### 42. Rate Limiting Pattern with Ticker
**Q: What does this enforce?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.NewTicker(100 * time.Millisecond)
    defer limiter.Stop()

    for req := range requests {
        <-limiter.C
        fmt.Println("request", req, "at", time.Now().Format("15:04:05.000"))
    }
}
```
**A:** One request processed every 100ms. `time.Ticker` is the idiomatic Go rate limiter — processes exactly N requests per second.

---

### 43. Retry Pattern with Goroutines
**Q: What does this pattern implement?**
```go
func withRetry(attempts int, sleep time.Duration, fn func() error) error {
    for i := 0; i < attempts; i++ {
        if err := fn(); err == nil {
            return nil
        }
        time.Sleep(sleep)
        sleep *= 2 // exponential backoff
    }
    return fmt.Errorf("all %d attempts failed", attempts)
}
```
**A:** Exponential backoff retry — calls `fn()` up to `attempts` times, doubling the sleep duration each time. Standard production pattern for network calls and external services.

---

### 44. Singleflight — Deduplicate Concurrent Requests
**Q: What problem does singleflight solve?**
```go
import "golang.org/x/sync/singleflight"

var sfg singleflight.Group

func fetchData(key string) (interface{}, error) {
    result, err, shared := sfg.Do(key, func() (interface{}, error) {
        // Only ONE of these runs even if 100 goroutines call with same key
        return expensiveDBCall(key)
    })
    _ = shared // true if result was shared
    return result, err
}
```
**A:** **Cache stampede / thundering herd prevention.** If 100 goroutines simultaneously request `key`, only ONE actual call is made. All 100 goroutines receive the same result. Essential for high-traffic services.

---

### 45. Bounded Parallelism with Semaphore Channel
**Q: What does this limit?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    const maxParallel = 3
    sem := make(chan struct{}, maxParallel)
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            sem <- struct{}{}           // acquire slot
            defer func() { <-sem }()   // release slot

            fmt.Printf("running task %d (parallel: %d)\n", id, len(sem))
        }(i)
    }
    wg.Wait()
}
```
**A:** At most 3 tasks run in parallel at any time. `len(sem)` shows current active count. This is the **semaphore pattern** — controls resource usage without a full worker pool.

---

### 46. Channel Direction Enforcement
**Q: Does this compile and what does it enforce?**
```go
package main
import "fmt"

func producer(out chan<- int) {
    for i := 0; i < 3; i++ {
        out <- i
    }
    close(out)
    // <-out  // compile error: cannot receive from send-only channel
}

func consumer(in <-chan int) {
    for v := range in {
        fmt.Println(v)
        // in <- 99  // compile error: cannot send to receive-only channel
    }
}

func main() {
    ch := make(chan int, 3)
    producer(ch)
    consumer(ch)
}
```
**A:** **Yes, compiles.** Directional channels prevent misuse at compile time. `chan<-` is send-only, `<-chan` is receive-only. Self-documenting and type-safe.

---

### 47. Heartbeat Pattern
**Q: What does the heartbeat channel convey?**
```go
func withHeartbeat(done <-chan struct{}, interval time.Duration) (<-chan struct{}, <-chan int) {
    heartbeat := make(chan struct{}, 1)
    work := make(chan int)
    go func() {
        ticker := time.NewTicker(interval)
        defer ticker.Stop()
        for {
            select {
            case <-done:
                return
            case <-ticker.C:
                select {
                case heartbeat <- struct{}{}: // non-blocking send
                default:
                }
            }
        }
    }()
    return heartbeat, work
}
```
**A:** The `heartbeat` channel pulses at `interval` — used to confirm a goroutine is still alive. Consumers can detect a stalled goroutine if heartbeats stop arriving. Pattern from "Concurrency in Go" book.

---

### 48. Preventing Goroutine Leaks — Always Have an Exit
**Q: What is guaranteed to stop the goroutine?**
```go
package main
import (
    "fmt"
    "time"
)

func generateForever(done <-chan struct{}) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        i := 0
        for {
            select {
            case <-done:
                return // guaranteed exit
            case out <- i:
                i++
            }
        }
    }()
    return out
}

func main() {
    done := make(chan struct{})
    nums := generateForever(done)
    for i := 0; i < 5; i++ {
        fmt.Println(<-nums)
    }
    close(done) // stops generator goroutine
    time.Sleep(10 * time.Millisecond)
}
```
**A:** `0 1 2 3 4`. The `done` channel guarantees the goroutine can always exit. Rule: **every goroutine must have a way to terminate**.

---

### 49. Data Race Detection
**Q: What flag detects data races?**
```go
// Run with: go run -race main.go
// Or: go test -race ./...

package main
import "fmt"

func main() {
    c := make(chan bool)
    m := make(map[string]string)
    go func() {
        m["1"] = "a" // unsynchronized write
        c <- true
    }()
    m["2"] = "b" // concurrent write
    <-c
    fmt.Println(m)
}
```
**A:** The `-race` flag enables the **race detector** — instruments memory accesses and reports races at runtime. Mandatory step before any concurrent code ships to production.

---

### 50. Context Propagation Basics (Preview)
**Q: What is the output?**
```go
package main
import (
    "context"
    "fmt"
    "time"
)

func doTask(ctx context.Context) error {
    select {
    case <-time.After(200 * time.Millisecond):
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
    defer cancel()

    err := doTask(ctx)
    fmt.Println(err)
}
```
**A:** `context deadline exceeded`. The task needs 200ms but the context times out at 100ms. This is the foundation of the context package — covered in depth in Part 2.

---

### 51. Memory Model: Happens-Before Guarantee
**Q: Can `a` and `done` values print out of order or hang in this code without channels/mutexes?**
```go
package main

import "fmt"

var a string
var done bool

func setup() {
    a = "hello, world"
    done = true
}

func main() {
    go setup()
    for !done {
    }
    fmt.Println(a)
}
```
**A:** **Yes, or it might hang forever.** Go's memory model does not guarantee that a write to `done` observed by `main` implies that the write to `a` is also observed. The compiler or CPU may reorder instructions. To establish a **"happens-before"** relationship and ensure memory visibility across goroutines, you must use explicit synchronization (like closing a channel, or `sync.Mutex`).

---

### 52. Leaky Buffer Pattern
**Q: How does this pattern prevent memory leaks and garbage collection pressure when handling many requests?**
```go
var freeList = make(chan *[]byte, 100)
// For GC efficiency in heavy networking/file IO

func getBuffer() *[]byte {
    select {
    case b := <-freeList:
        return b // reuse existing buffer
    default:
        b := make([]byte, 1024)
        return &b // allocate new if none free
    }
}

func releaseBuffer(b *[]byte) {
    select {
    case freeList <- b:
        // returned to free list
    default:
        // free list full, let GC handle it
    }
}
```
**A:** This is the **Leaky Buffer** pattern. A buffered channel acts as a free list to reuse memory allocations (reducing GC pressure). If the server gets a sudden spike, it allocates new buffers (the `default` case in `getBuffer`). When the spike ends, excess buffers are dropped (the `default` case in `releaseBuffer`) and garbage collected natively, so memory doesn't grow unbounded forever. It's a simpler, channel-based alternative to `sync.Pool`.

---

*End of Part 1 — Concurrency Patterns (52 questions)*
*See `go_intermediate_context_interfaces.md` for Part 2: Context, Interfaces, Generics, Error Handling*
