# ⚡ 02 — Advanced Concurrency & Goroutines
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- GMP model (Goroutine, Machine, Processor)
- Advanced channel patterns (fan-in, fan-out, pipeline)
- `sync/atomic` operations
- Semaphore pattern
- `errgroup` for concurrent error handling
- Context propagation in concurrent systems
- Deadlocks and how to prevent them

---

## ❓ Most Asked Questions

### Q1. Explain the GMP (Goroutine-Machine-Processor) Model

```
G (Goroutine) — lightweight execution unit (~2KB stack, grows dynamically)
M (Machine)   — OS thread managed by Go runtime
P (Processor) — logical processor, holds run queue of goroutines

- GOMAXPROCS sets number of P's (default = number of CPU cores)
- Each P runs on one M at a time
- Goroutines are assigned to P's run queue
- When G blocks (I/O, syscall), P detaches from M and finds another M
- This allows millions of goroutines with only handful of OS threads
```

```go
import "runtime"
fmt.Println(runtime.GOMAXPROCS(0))   // current value
runtime.GOMAXPROCS(4)                // set to 4
fmt.Println(runtime.NumGoroutine())  // current goroutine count
```

---

### Q2. Implement Pipeline Pattern

```go
// Stage 1: Generate numbers
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums { out <- n }
    }()
    return out
}

// Stage 2: Square numbers
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in { out <- n * n }
    }()
    return out
}

// Stage 3: Filter even numbers
func filterEven(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            if n%2 == 0 { out <- n }
        }
    }()
    return out
}

// Compose pipeline
func main() {
    nums := generate(1, 2, 3, 4, 5)
    squared := square(nums)
    evens := filterEven(squared)
    for v := range evens { fmt.Println(v) }  // 4, 16
}
```

---

### Q3. Implement Fan-Out / Fan-In

```go
// Fan-Out: distribute work from one channel to multiple goroutines
func fanOut(in <-chan int, numWorkers int) []<-chan int {
    channels := make([]<-chan int, numWorkers)
    for i := 0; i < numWorkers; i++ {
        ch := make(chan int)
        channels[i] = ch
        go func() {
            defer close(ch)
            for v := range in { ch <- v * v }
        }()
    }
    return channels
}

// Fan-In: merge multiple channels into one
func fanIn(channels ...<-chan int) <-chan int {
    merged := make(chan int)
    var wg sync.WaitGroup

    merge := func(ch <-chan int) {
        defer wg.Done()
        for v := range ch { merged <- v }
    }

    wg.Add(len(channels))
    for _, ch := range channels { go merge(ch) }

    go func() { wg.Wait(); close(merged) }()
    return merged
}
```

---

### Q4. Implement a Semaphore in Go

```go
// Semaphore using buffered channel to limit concurrency
type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
    return make(Semaphore, n)  // capacity = max concurrent goroutines
}

func (s Semaphore) Acquire() { s <- struct{}{} }  // blocks when full
func (s Semaphore) Release() { <-s }               // frees a slot

// Limit to 3 concurrent HTTP requests
sem := NewSemaphore(3)
var wg sync.WaitGroup

for _, url := range urls {
    wg.Add(1)
    go func(u string) {
        defer wg.Done()
        sem.Acquire()
        defer sem.Release()
        fetchURL(u)
    }(url)
}
wg.Wait()
```

---

### Q5. How do you use `sync/atomic` for lock-free operations?

```go
import "sync/atomic"

var counter int64

// Atomic increment — no mutex needed
atomic.AddInt64(&counter, 1)

// Atomic read
val := atomic.LoadInt64(&counter)

// Atomic compare-and-swap (CAS) — foundation of lock-free structures
old := int64(10)
new := int64(20)
swapped := atomic.CompareAndSwapInt64(&counter, old, new)
// swapped=true means: counter was 10, now set to 20

// Atomic store
atomic.StoreInt64(&counter, 100)

// Real use: lock-free singleton check
var initialized int32
func ensureInit() {
    if atomic.CompareAndSwapInt32(&initialized, 0, 1) {
        // Only one goroutine runs this
        initExpensiveResource()
    }
}
```

---

### Q6. How do you use `errgroup` for concurrent error handling?

```go
import "golang.org/x/sync/errgroup"

func fetchAll(urls []string) error {
    g, ctx := errgroup.WithContext(context.Background())

    results := make([]string, len(urls))

    for i, url := range urls {
        i, url := i, url  // capture loop variables
        g.Go(func() error {
            // Respect context cancellation
            select {
            case <-ctx.Done():
                return ctx.Err()
            default:
            }
            resp, err := http.Get(url)
            if err != nil { return fmt.Errorf("fetch %s: %w", url, err) }
            defer resp.Body.Close()
            body, _ := io.ReadAll(resp.Body)
            results[i] = string(body)
            return nil
        })
    }

    // Wait returns first non-nil error; ctx is cancelled when any goroutine errors
    if err := g.Wait(); err != nil {
        return err
    }
    return nil
}
```

---

### Q7. How do you avoid deadlocks?

```go
// Common deadlock scenario
// ch1 := make(chan int)
// ch2 := make(chan int)
// go func() { ch1 <- 1; v := <-ch2 }()
// go func() { ch2 <- 2; v := <-ch1 }()  // DEADLOCK

// Prevention rules:
// 1. Always acquire locks in the same order
// 2. Use select with default for non-blocking channel ops
// 3. Set timeouts with context
// 4. Don't hold a lock while waiting for a channel

// Detect deadlocks with -race and runtime detection
// go run -race main.go
// The Go runtime detects deadlocks: "all goroutines are asleep - deadlock!"

// Safe pattern: use select with timeout
result := make(chan int)
go func() { result <- compute() }()

select {
case v := <-result:
    fmt.Println("got:", v)
case <-time.After(5 * time.Second):
    fmt.Println("timed out")
}
```

---

### Q8. Implement a Publisher-Subscriber system using channels

```go
type PubSub struct {
    mu          sync.RWMutex
    subscribers map[string][]chan string
}

func NewPubSub() *PubSub {
    return &PubSub{subscribers: make(map[string][]chan string)}
}

func (ps *PubSub) Subscribe(topic string) <-chan string {
    ch := make(chan string, 10)
    ps.mu.Lock()
    ps.subscribers[topic] = append(ps.subscribers[topic], ch)
    ps.mu.Unlock()
    return ch
}

func (ps *PubSub) Publish(topic, message string) {
    ps.mu.RLock()
    defer ps.mu.RUnlock()
    for _, ch := range ps.subscribers[topic] {
        select {
        case ch <- message:
        default:  // skip if subscriber is slow
        }
    }
}

// Usage
ps := NewPubSub()
sub1 := ps.Subscribe("orders")
sub2 := ps.Subscribe("orders")
go ps.Publish("orders", "order#123 placed")
fmt.Println(<-sub1)  // order#123 placed
fmt.Println(<-sub2)  // order#123 placed
```
