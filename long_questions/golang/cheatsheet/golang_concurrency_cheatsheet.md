# Golang Concurrency Cheatsheet

Quick reference for Goroutines, Channels, WaitGroups, and common concurrency patterns.

---

## 🟢 Basics

### 1. Goroutines
Lightweight threads managed by Go runtime.
```go
func main() {
    // Start a goroutine
    go func() {
        fmt.Println("Hello from goroutine")
    }()
    
    // Wait for it (naive way, see WaitGroup below)
    time.Sleep(time.Second) 
}
```

### 2. WaitGroup (`sync`)
Standard way to wait for a collection of goroutines to finish.
```go
import "sync"

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1) // Increment counter
        go func(id int) {
            defer wg.Done() // Decrement counter
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }
    
    wg.Wait() // Block until counter is 0
}
```

### 3. Mutex (`sync`)
Prevent data races (concurrent read/write) on shared resources.
```go
import "sync"

type SafeCounter struct {
    mu sync.Mutex
    v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
    c.mu.Lock()   // Lock
    defer c.mu.Unlock() // Ensure Unlock happens
    c.v[key]++
}
```

---

## 🟡 Channels

Typed conduits for sending and receiving values.

### 1. Declaration & Init
```go
// Unbuffered: Blocks until sender & receiver are both ready
ch := make(chan int) 

// Buffered: Non-blocking until buffer is full
bufCh := make(chan int, 5) 

// Directional (Function Parameters)
// p <-chan int  (Receive-only)
// p chan<- int  (Send-only)
```

### 2. Operations
```go
ch <- 42      // Send
val := <-ch   // Receive

// Close
close(ch)     // Signals no more values will be sent
// Note: Sending to closed char -> PANIC
//       Receiving from closed -> Returns zero-value immediately
```

### 3. Iterating (Range)
Loops until channel is closed.
```go
for msg := range ch {
    fmt.Println(msg)
}
```

### 4. Select
Wait on multiple channel operations simultaneously.
```go
select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case ch2 <- "message":
    fmt.Println("Sent to ch2")
case <-time.After(time.Second):
    fmt.Println("Timeout") // Useful for SLAs
default:
    fmt.Println("No activity (Non-blocking)")
}
```

---

## 🟣 Common Patterns

### 1. Generator
Function that returns a channel.
```go
func count(n int) <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < n; i++ {
            ch <- i
        }
        close(ch)
    }()
    return ch
}
// Use: for num := range count(5) { ... }
```

### 2. Fan-In (Multiplexing)
Merge multiple channels into one.
```go
func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    output := func(c <-chan int) {
        defer wg.Done()
        for n := range c { out <- n }
    }
    
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }
    
    // Closer goroutine
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}
```

### 3. Worker Pool
Distribute work among fixed number of workers.
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, j)
        time.Sleep(time.Second) // Simulate work
        results <- j * 2
    }
}

func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    
    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs) // Signal no more jobs
    
    // Collect results
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
```

### 4. Done Channel (Cancellation)
Signal goroutines to stop.
```go
func doWork(done <-chan bool) {
    for {
        select {
        case <-done:
            return
        default:
            // Do work...
        }
    }
}

// In main:
// close(done) to stop all workers
```

---

---

## 🟠 Advanced & Production Patterns

### 1. Context Package (`context`)
Essential for managing cancellation signals and deadlines across API boundaries.
```go
import "context"

func main() {
    // Return early if work takes > 50ms
    ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
    defer cancel() // Release resources

    select {
    case <-time.After(1 * time.Second):
        fmt.Println("Overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err()) // prints "context deadline exceeded"
    }
}
```

### 2. Initializing Once (`sync.Once`)
Guarantees a function runs exactly once (thread-safe Singleton).
```go
import "sync"

var once sync.Once
var instance *Config

func GetConfig() *Config {
    once.Do(func() {
        // Runs only once even if called concurrently
        instance = &Config{}
    })
    return instance
}
```

### 3. Atomic Operations (`sync/atomic`)
Lock-free operations for simple counters/flags. Faster than Mutex for simple cases.
```go
import "sync/atomic"

var ops uint64

func main() {
    // Increment
    atomic.AddUint64(&ops, 1)

    // Load (Read)
    val := atomic.LoadUint64(&ops)
    fmt.Println("Ops:", val)
}
```

### 4. Error Groups (`errgroup`)
Managing multiple goroutines where any single error should cancel the rest.
*(Requires `golang.org/x/sync/errgroup`)*
```go
import "golang.org/x/sync/errgroup"

func main() {
    g := new(errgroup.Group)
    urls := []string{"http://google.com", "http://bad-url"}

    for _, url := range urls {
        url := url // Capture loop var
        g.Go(func() error {
            // If this returns error, all other goroutines in 'g' are ignored/cancelled
            return fetch(url) 
        })
    }

    if err := g.Wait(); err != nil {
        fmt.Println("Successfully failed:", err)
    }
}
```

---

## 🔴 Common Pitfalls

1.  **Deadlock**: All goroutines are asleep.
    *   *Cause*: Waiting on a channel that no one is sending to (or vice versa).
2.  **Leaking Goroutines**: Goroutines that never exit.
    *   *Fix*: Ensure they have a way to return (e.g., proper channel closing or context cancellation).
3.  **Race Connection**: Multiple threads accessing variable without locking.
    *   *Check*: Run tests with `go test -race`.
