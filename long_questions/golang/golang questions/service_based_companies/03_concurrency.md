# 📙 03 — Concurrency in Go
> **Most Asked in Service-Based Companies** | 🟡 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Goroutines vs OS threads
- Buffered vs unbuffered channels
- `sync.WaitGroup`, `sync.Mutex`, `sync.Once`
- `select` statement
- `context.Context` for cancellation
- Race conditions and the `-race` flag

---

## ❓ Most Asked Questions

### Q1. What are goroutines? How are they different from threads?

```go
// Start a goroutine with 'go' keyword
go func() {
    fmt.Println("running in goroutine")
}()

// Multiple goroutines
for i := 0; i < 5; i++ {
    go func(n int) {
        fmt.Println("goroutine", n)
    }(i)  // pass i as argument to avoid closure capture issues
}
```

| | Goroutine | OS Thread |
|--|-----------|-----------|
| Stack size | ~2KB (grows dynamically) | ~1–8MB fixed |
| Creation cost | Very cheap | Expensive |
| Scheduling | Go runtime (cooperative + preemptive) | OS kernel |
| Count | Millions possible | Thousands max |
| Communication | Channels | Shared memory + locks |

---

### Q2. What are channels? What is the difference between buffered and unbuffered?

```go
// Unbuffered channel — sender blocks until receiver is ready
ch := make(chan int)
go func() { ch <- 42 }()  // blocks until main reads
val := <-ch                // synchronous

// Buffered channel — sender blocks only when buffer is full
bch := make(chan int, 3)   // buffer size 3
bch <- 1  // doesn't block
bch <- 2  // doesn't block
bch <- 3  // doesn't block
// bch <- 4  // would block — buffer full

fmt.Println(<-bch)  // 1
```

| | Unbuffered | Buffered |
|--|-----------|---------|
| Capacity | 0 | > 0 |
| Sender blocks | Until receiver ready | Until buffer full |
| Use case | Synchronization | Rate limiting, work queues |

---

### Q3. How does `sync.WaitGroup` work?

```go
var wg sync.WaitGroup

for i := 1; i <= 5; i++ {
    wg.Add(1)  // increment counter before launch
    go func(id int) {
        defer wg.Done()  // decrement when done
        fmt.Printf("Worker %d done\n", id)
    }(i)
}

wg.Wait()  // blocks until counter reaches 0
fmt.Println("all workers done")
```

> **Important:** Call `wg.Add(1)` BEFORE the goroutine starts, not inside it.

---

### Q4. How does `sync.Mutex` prevent race conditions?

```go
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

counter := &SafeCounter{}
var wg sync.WaitGroup
for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        counter.Increment()
    }()
}
wg.Wait()
fmt.Println(counter.Value())  // always 1000
```

---

### Q5. What is `sync.RWMutex` and when to use it?

```go
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

// Multiple readers can read simultaneously
func (c *Cache) Get(key string) (string, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    v, ok := c.data[key]
    return v, ok
}

// Only one writer at a time; blocks all readers
func (c *Cache) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}
```
> **Use RWMutex when:** reads are much more frequent than writes (e.g., caching, config).

---

### Q6. What is the `select` statement?

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() { time.Sleep(1 * time.Second); ch1 <- "one" }()
go func() { time.Sleep(2 * time.Second); ch2 <- "two" }()

// select picks whichever channel is ready first
select {
case msg := <-ch1:
    fmt.Println("received:", msg)
case msg := <-ch2:
    fmt.Println("received:", msg)
}

// With default — non-blocking
select {
case v := <-ch1:
    fmt.Println(v)
default:
    fmt.Println("no message ready")
}

// With timeout
select {
case v := <-ch1:
    fmt.Println(v)
case <-time.After(500 * time.Millisecond):
    fmt.Println("timed out")
}
```

---

### Q7. How do you close a channel and how to detect it?

```go
ch := make(chan int)

// Producer goroutine
go func() {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)  // ONLY sender should close
}()

// Consumer: range loop — stops when closed
for v := range ch {
    fmt.Println(v)
}

// Check if closed
v, ok := <-ch
if !ok {
    fmt.Println("channel is closed")
}
```
> **Rules:** Never close a channel from receiver; never close a nil channel; don't close twice.

---

### Q8. What is `sync.Once` used for?

```go
var (
    instance *Database
    once     sync.Once
)

// Thread-safe singleton
func GetDB() *Database {
    once.Do(func() {
        instance = &Database{conn: "postgresql://..."}
        fmt.Println("DB initialized once")
    })
    return instance
}
```
> `sync.Once` ensures the function runs exactly once, even with concurrent calls.

---

### Q9. How do you use `context.Context` for cancellation?

```go
// Create cancellable context
ctx, cancel := context.WithCancel(context.Background())
defer cancel()  // always call cancel

go func(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("goroutine stopped:", ctx.Err())
            return
        default:
            fmt.Println("working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}(ctx)

time.Sleep(2 * time.Second)
cancel()  // cancels the goroutine

// Context with timeout
ctx2, cancel2 := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel2()
```

---

### Q10. How do you implement a simple worker pool?

```go
func workerPool(jobs <-chan int, results chan<- int, workerID int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        results <- j * j  // process job
        fmt.Printf("Worker %d processed job %d\n", workerID, j)
    }
}

func main() {
    jobs    := make(chan int, 100)
    results := make(chan int, 100)
    var wg sync.WaitGroup

    // Start 3 workers
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go workerPool(jobs, results, w, &wg)
    }

    // Send 9 jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Wait and collect
    go func() {
        wg.Wait()
        close(results)
    }()

    for r := range results {
        fmt.Println("Result:", r)
    }
}
```

---

### Q11. How do you detect race conditions?

```go
// Run with race detector
// go run -race main.go
// go test -race ./...

// Example of a race condition
var counter int
var wg sync.WaitGroup
for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        counter++  // DATA RACE — concurrent read/write
    }()
}
wg.Wait()
// Use sync.Mutex or sync/atomic to fix
```
