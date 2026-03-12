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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug and the fix?
**Your Response:** This has a data race bug. Multiple goroutines increment the same counter variable without synchronization. The increment operation isn't atomic - it involves reading the value, incrementing it, and writing it back. When multiple goroutines do this simultaneously, some increments get lost, so the final value is unpredictable. The fix is to add a mutex to protect the shared state - lock before incrementing and unlock after.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug?
**Your Response:** The bug is that the Safe struct is being copied by value when passed to process(). Since sync.Mutex contains internal state, copying it creates two mutexes with inconsistent lock states, leading to undefined behavior. Go's vet tool detects this. The fix is to always pass mutex-containing structs by pointer instead of by value.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why is this pattern preferred?
**Your Response:** Using defer for Unlock() is preferred because it guarantees the mutex is always released, even if a panic occurs inside the function. Without defer, if the code panics after Lock() but before Unlock(), the mutex remains locked forever, causing a deadlock for any other goroutine trying to acquire it.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output and why is RWMutex better here?
**Your Response:** This prints 42. RWMutex is better here because it allows multiple concurrent readers. Multiple goroutines can call Get() simultaneously with RLock(), but only one goroutine can call Set() with Lock(). For read-heavy workloads where writes are infrequent, this provides much better performance than a regular Mutex which would only allow one goroutine at a time.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this deadlock?
**Your Response:** Yes, this deadlocks. Go's sync.Mutex is not reentrant, which means the same goroutine cannot lock it twice. When the goroutine calls Lock() the second time, it blocks forever waiting for itself to release the lock. This is a common gotcha for developers coming from languages where mutexes are reentrant.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this deadlock?
**Your Response:** Yes, this deadlocks. You cannot upgrade a read lock to a write lock in Go's RWMutex. When we call Lock() while holding RLock(), it waits for all readers to finish, but our goroutine is itself a reader, creating a circular wait. The solution is to release the read lock first, then acquire the write lock.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints ready: true. sync.Cond is a condition variable that allows goroutines to wait for a specific condition. cond.Wait() atomically releases the mutex and blocks the goroutine. When cond.Signal() is called, it wakes up one waiting goroutine. This is commonly used for producer-consumer patterns where one goroutine needs to wait for another to complete some work.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints 42 true then false. sync.Map is a concurrent-safe map that doesn't require external locking. Store() adds a key-value pair, and Load() retrieves it with a boolean indicating if the key exists. It's optimized for workloads where keys are written once and read many times, making it perfect for caches or configuration maps.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does Range do?
**Your Response:** Range iterates over all key-value pairs in the sync.Map. The order is not guaranteed. The callback function receives each key and value, and if it returns false, the iteration stops early. This is useful for searching or when you don't need to iterate through all entries.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints first false then first true. LoadOrStore is an atomic operation that either loads an existing value or stores a new one if it doesn't exist. The first call stores "first" and reports loaded=false. The second call finds the existing value and returns it with loaded=true. This is great for cache-like patterns where you want to initialize a value only once.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is wrong?
**Your Response:** While this compiles and works, embedding sync.Mutex promotes the Lock() and Unlock() methods to the public API of Counter. This means external code can accidentally manipulate the internal mutex, which can lead to deadlocks or inconsistent state. It's better to use a named field like mu sync.Mutex to keep the mutex private and prevent accidental misuse.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints acquired: false then acquired: true. TryLock is a non-blocking attempt to acquire a mutex. It returns immediately with false if the mutex is already held, or true if it successfully acquires the lock. This is useful when you don't want to block and need to check if you can acquire the lock without waiting.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why does this always print 1000?
**Your Response:** This prints 1000 because atomic.AddInt64 performs an atomic increment operation at the CPU level. Unlike regular integer addition which involves read-modify-write steps that can be interleaved across goroutines, atomic operations are indivisible. This makes it the fastest way to build a concurrent counter without using mutexes.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why use Load/Store instead of direct read?
**Your Response:** This prints 1. On some platforms, reading a 64-bit value isn't atomic - it might read the high 32 bits and then the low 32 bits, potentially getting a mixed value if another goroutine writes simultaneously. LoadInt64 and StoreInt64 ensure the entire 64-bit value is read or written atomically, preventing data races and ensuring proper memory visibility across goroutines.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints true 20 then false 20. CompareAndSwap atomically checks if the current value equals the expected value (10), and only if it does, updates it to the new value (20). The first call succeeds because val is 10, so it updates to 20 and returns true. The second call fails because val is now 20, not 10, so it returns false without updating. CAS is fundamental for building lock-free data structures.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints 100. atomic.Value is a special type that can store and load any value atomically without needing a mutex. It's particularly useful for configuration hot-reloading patterns where you want to atomically swap an entire configuration struct. Here we store a map and then load it back, type-asserting it to the expected type.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens?
**Your Response:** This panics. atomic.Value has a strict requirement: all values stored must be of the same concrete type. The first Store establishes the type, and any subsequent Store with a different type will panic. This is a design choice to ensure type safety and avoid runtime surprises when loading values.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which is more appropriate here?
**Your Response:** For Option A with a Mutex, use it when you need to update multiple fields together to maintain consistency, or when updating non-atomic types like float64. For Option B with atomic operations, use it for simple independent integer counters - it's faster and has no contention. Atomic operations are ideal for simple metrics and flags, while Mutex is better for complex state updates.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the problem with this?
**Your Response:** This is a spinlock implementation that wastes CPU. While it works, under contention it continuously loops checking the lock value (busy-waiting), burning CPU cycles. Go's sync.Mutex is better because it uses the runtime scheduler to block the goroutine efficiently when the lock is held, allowing the CPU to do other work. Spinlocks are only appropriate for very short critical sections on systems where context switching is expensive.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this safely do?
**Your Response:** This prints false then true. atomic.Pointer is a type-safe atomic pointer introduced in Go 1.19. It allows you to atomically store and load pointers of a specific type. This is perfect for atomically swapping configuration structs without using a mutex. The type parameter ensures compile-time type safety, unlike atomic.Value which requires type assertions.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints worker stopped. The done channel is a cancellation pattern. The worker goroutine runs in a loop, checking the done channel in a select statement. When main closes the done channel, all goroutines waiting on it unblock immediately. This is the idiomatic way to broadcast cancellation to multiple goroutines in Go.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints 4 9 16. This is a pipeline pattern where each stage processes data concurrently. generate() creates a channel and sends numbers to it. square() reads from that channel, squares each number, and sends the result to its output channel. The stages run concurrently, with each goroutine processing independently. This is a fundamental pattern in Go for data processing pipelines.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this pattern do?
**Your Response:** This implements the fan-out pattern, distributing work across multiple worker goroutines. The fanOut function creates multiple output channels, each handled by a separate goroutine that doubles values from the input channel. This parallelizes work processing - instead of one worker handling all values sequentially, multiple workers handle subsets concurrently, improving throughput for CPU-bound operations.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints 1 2 3 4, though the order may vary. The merge function implements fan-in, combining multiple input channels into one output channel. It launches a goroutine for each input channel that forwards values to the output. The output channel only closes after all input goroutines finish. This is the complement to fan-out, useful when you need to collect results from parallel workers.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints 55. This is a worker pool pattern with 3 fixed workers processing 5 jobs. Each worker squares the job number and sends the result to the results channel. The main function collects all results and sums them. Worker pools are essential for controlling resource usage - instead of spawning unlimited goroutines, you have a fixed pool that processes jobs sequentially, preventing resource exhaustion.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints timeout!. The time.After function returns a channel that will receive a value after 100ms. In the select statement, since no value is ever sent on ch, the timeout case fires first. This is the standard pattern for implementing timeouts in Go operations, preventing them from blocking forever.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this pattern control?
**Your Response:** This controls concurrency to at most 3 goroutines simultaneously. The buffered channel acts as a semaphore - goroutines send to it to acquire a slot and receive from it to release. Since the channel capacity is 3, only 3 goroutines can be in the critical section at once. This is perfect for rate-limiting resource access like database connections or API calls.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this safely read from a channel with cancellation?
**Your Response:** The orDone function safely reads from channel c while respecting cancellation. It uses a double-select pattern - the outer select checks for the done signal, and the inner select ensures we don't block when trying to forward values to the output channel. If done closes while we're trying to send, we immediately exit. This pattern prevents goroutine leaks when combining cancellation with channel operations.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the goroutine leak?
**Your Response:** The goroutine leaks because it's trying to send 10 values but the caller only reads one. After sending the first value, the goroutine blocks on the second send, waiting for a receiver that never comes. This creates a goroutine leak - the goroutine is stuck forever consuming memory. The fix is to either use a buffered channel, ensure the caller drains all values, or provide a done channel for cancellation.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Is the choice deterministic?
**Your Response:** No, this is not deterministic. When multiple cases in a select are ready simultaneously, Go picks one uniformly at random. Since both ch1 and ch2 have buffered values ready, each run may print either "one" or "two". This randomness is intentional to prevent bias toward any particular case. If you need deterministic behavior, you should use a different pattern like prioritized selects.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** All 3 goroutines print their message, though the order varies. This demonstrates that closing a channel broadcasts to all receivers. When a channel is closed, all goroutines waiting to receive from it unblock immediately and receive the zero value. This is the only way to broadcast a signal to multiple goroutines simultaneously in Go.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the race condition?
**Your Response:** The race condition is that wg.Wait() might complete before any goroutine calls wg.Add(1). Since wg.Add() is called inside the goroutine, there's a window where main could check the counter, find it's zero, and return immediately. Always call wg.Add() before starting the goroutines to ensure the counter is incremented before any goroutine can call wg.Done().

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does main survive?
**Your Response:** Yes, main survives and prints recovered: oops. The safeGo helper wraps the goroutine in a recover function, catching any panics and printing them instead of letting the program crash. This is a production pattern for background workers where you want to handle panics gracefully without taking down the entire application.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints task 1, task 2, task 3. The buffered channel holds functions as values. We enqueue three tasks, close the channel, then iterate over it executing each function. This creates a simple task queue where tasks are buffered without requiring an immediate receiver. It's lightweight and doesn't need a separate worker goroutine.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints DB initialized exactly once. sync.Once ensures that the function passed to Do() is executed exactly one time, no matter how many goroutines call it simultaneously. It uses atomic operations internally to track execution state. This is the idiomatic way to implement singletons in Go, perfect for expensive one-time initialization like database connections.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does Go expose goroutine IDs?
**Your Response:** No, Go deliberately doesn't expose goroutine IDs in the standard library. This is intentional - the Go team discourages goroutine-local storage patterns that are common with thread IDs. Instead, you should use explicit parameters or channels to pass data. You can get the count of goroutines with runtime.NumGoroutine(), but not individual IDs.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the problem?
**Your Response:** The problem is that select doesn't prioritize channels - it randomly picks between highPriority and lowPriority when both have values. If you need priority handling, you should first check highPriority with a non-blocking select, and only if that's empty, fall back to a select that includes lowPriority. This gives you explicit control over priority rather than relying on random selection.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the canonical pattern for a long-running goroutine?
**Your Response:** This is the canonical for-select loop pattern used in every production Go service. It continuously checks two things: the done channel for graceful shutdown, and the data channel for incoming work. The ok check handles when the data channel is closed. This pattern ensures goroutines can exit cleanly and handle all edge cases properly.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the initial goroutine stack size and how does it grow?
**Your Response:** Goroutines start with a very small stack, typically 2KB compared to 1MB for OS threads. They grow and shrink dynamically using stack copying. When a goroutine needs more stack space, Go allocates a new, larger stack and copies the old one. This efficient memory usage is why Go can run millions of goroutines while other languages struggle with thousands of threads.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output and what does GOMAXPROCS control?
**Your Response:** The output shows the number of CPU cores and current GOMAXPROCS setting. GOMAXPROCS controls how many OS threads can execute Go code simultaneously. Before Go 1.5, the default was 1, meaning goroutines were concurrent but not parallel. Now the default matches the number of CPU cores, enabling true parallelism. Setting it to 1 would make goroutines concurrent but running on a single thread.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints error: task failed. The errgroup.Group runs multiple goroutines concurrently and waits for all to complete. It returns the first non-nil error encountered, or nil if all succeed. This is the standard pattern for parallel operations like API calls or database queries where you want to fail fast if any operation fails.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this enforce?
**Your Response:** This enforces a rate limit of one request every 100ms. The time.Ticker channel delivers a value at regular intervals. By receiving from ticker.C before processing each request, we ensure exactly one request is processed per interval. This is the idiomatic way to implement rate limiting in Go, preventing overwhelming downstream services.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this pattern implement?
**Your Response:** This implements exponential backoff retry. It calls the function up to the specified number of attempts, doubling the sleep duration between attempts. This pattern is standard for network calls and external services - it gives the service time to recover while not overwhelming it with immediate retries. The exponential increase prevents thundering herd problems.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What problem does singleflight solve?
**Your Response:** Singleflight solves the cache stampede or thundering herd problem. When multiple goroutines request the same key simultaneously, only one actual call is made. All waiting goroutines receive the same result. This is essential for high-traffic services where a cache miss could trigger thousands of simultaneous database calls. It ensures the expensive operation runs only once.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this limit?
**Your Response:** This limits parallelism to 3 tasks. The semaphore channel acts as a counter - goroutines send to it to acquire a slot and receive to release. The len(sem) shows how many are currently active. This is simpler than a full worker pool when you just need to limit concurrency without the complexity of job distribution.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and what does it enforce?
**Your Response:** Yes, this compiles. Directional channels enforce usage at compile time. The producer function takes a send-only channel (chan<-), so it can only send to it, not receive. The consumer takes a receive-only channel (<-chan), so it can only receive. This prevents accidental misuse and makes the API self-documenting about the intended data flow.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does the heartbeat channel convey?
**Your Response:** The heartbeat channel pulses at regular intervals to show the goroutine is still alive and processing. Consumers can monitor this channel - if heartbeats stop arriving, it indicates the goroutine may be stalled or deadlocked. This pattern is useful for health checks and monitoring long-running goroutines where you need to know they're still functioning.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is guaranteed to stop the goroutine?
**Your Response:** The done channel guarantees the goroutine can always exit. The generator runs in a loop, but each iteration checks the done channel. When main closes done, the generator receives the signal and returns immediately. This is a critical pattern - every goroutine must have a termination path to prevent leaks.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What flag detects data races?
**Your Response:** The -race flag enables Go's race detector. It instruments the program to detect when multiple goroutines access the same memory concurrently without proper synchronization. Running tests or programs with -race will report data races at runtime. This is a mandatory step before shipping concurrent code to production - it catches subtle bugs that might only appear under specific timing conditions.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints context deadline exceeded. The context has a 100ms timeout, but the task takes 200ms. The doTask function checks the context's Done channel - when the timeout expires, it returns the context's error. This is the foundation of Go's context package for handling cancellation and timeouts across API boundaries.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can `a` and `done` values print out of order or hang in this code without channels/mutexes?
**Your Response:** Yes, this can print out of order or hang forever. Go's memory model doesn't guarantee that observing the write to done means you'll also observe the write to a. The compiler or CPU might reorder instructions. Without proper synchronization like channels or mutexes to establish a happens-before relationship, there's no guarantee of memory visibility across goroutines.

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

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does this pattern prevent memory leaks and garbage collection pressure when handling many requests?
**Your Response:** This is the Leaky Buffer pattern for managing memory efficiently. The buffered channel acts as a free list - buffers are reused when available, reducing GC pressure. During traffic spikes, it allocates new buffers as needed. When traffic subsides, excess buffers are dropped and garbage collected naturally. This prevents unbounded memory growth while still reusing buffers efficiently. It's a simpler alternative to sync.Pool for managing reusable resources.

---

*End of Part 1 — Concurrency Patterns (52 questions)*
*See `go_intermediate_context_interfaces.md` for Part 2: Context, Interfaces, Generics, Error Handling*
