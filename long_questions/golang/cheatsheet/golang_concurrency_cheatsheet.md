

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

### Explanation
Goroutines are lightweight threads managed by Go's runtime. They start with only 2KB stacks and are multiplexed onto OS threads. The go keyword starts a new goroutine that runs concurrently with the calling function.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are goroutines in Go?
**Your Response:** "Goroutines are Go's lightweight threads managed by the runtime. They start with only 2KB stacks compared to OS threads that start with megabytes. The Go runtime multiplexes many goroutines onto fewer OS threads using the GMP scheduler. I start a goroutine with the go keyword, and it runs concurrently with other goroutines. Goroutines are cheap to create - I can spawn thousands without issues. They communicate through channels rather than shared memory. The runtime handles scheduling, preemption, and stack growth automatically. This makes concurrency in Go much more approachable than traditional threading models."

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

### Explanation
WaitGroup is a synchronization primitive that waits for a collection of goroutines to finish. Add() increments the counter, Done() decrements it, and Wait() blocks until the counter reaches zero.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you wait for multiple goroutines to finish in Go?
**Your Response:** "I use sync.WaitGroup to coordinate multiple goroutines. Before starting each goroutine, I call wg.Add() to increment the counter. Inside each goroutine, I use defer wg.Done() to ensure the counter is decremented when the goroutine exits. Then I call wg.Wait() in the main function, which blocks until all goroutines have called Done(). This pattern is much cleaner than using time.Sleep() or other ad-hoc methods. The WaitGroup ensures I wait exactly as long as needed - no more, no less. It's the standard way to synchronize goroutine completion in Go."

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

### Explanation
Mutex provides mutual exclusion for protecting shared resources from concurrent access. Lock() acquires exclusive access, and Unlock() releases it. Using defer ensures Unlock is always called even if panic occurs.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you prevent race conditions in Go?
**Your Response:** "I use sync.Mutex to prevent race conditions when multiple goroutines access shared data. I wrap critical sections with Lock() and Unlock() calls. The key pattern is to always use defer Unlock() immediately after Lock() to ensure the mutex is released even if a panic occurs. Only one goroutine can hold the lock at a time, preventing simultaneous access. For read-heavy scenarios, I might use sync.RWMutex which allows multiple readers or one exclusive writer. I also run my tests with the race detector using 'go test -race' to catch any race conditions I might have missed. The mutex is Go's fundamental tool for safe shared memory access."

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

### Explanation
Generator pattern creates a function that returns a channel. The function spawns a goroutine that generates values and sends them to the channel, closing it when done. This provides a clean way to produce streams of data.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the generator pattern in Go?
**Your Response:** "The generator pattern creates a function that returns a channel instead of a value. Inside the function, I spawn a goroutine that generates values and sends them to the channel. When generation is complete, I close the channel. This gives the caller a clean interface - they can simply range over the returned channel to receive values. The pattern is great for producing streams of data, reading files line by line, or any scenario where I want to lazy-generate values. It separates the generation logic from the consumption logic and leverages Go's channel semantics for clean coordination. The caller doesn't need to know about the internal goroutine - they just get a channel to read from."

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

### Explanation
Fan-in pattern merges multiple input channels into one output channel. Each input channel is handled by a dedicated goroutine that forwards values. A WaitGroup ensures proper closure when all inputs are exhausted.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement fan-in in Go?
**Your Response:** "I implement fan-in to merge multiple channels into one. I create an output channel and a WaitGroup. For each input channel, I spawn a goroutine that reads from that channel and forwards values to the output. Each goroutine calls Done() when its input closes. I also spawn a goroutine that waits for the WaitGroup and then closes the output. This ensures all values flow through correctly and the output closes only after all inputs are done. The pattern is essential when I need to collect results from multiple concurrent operations into a single stream. It's a fundamental Go concurrency pattern that combines multiple data sources efficiently."

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

### Explanation
Worker pool pattern limits concurrency to a fixed number of workers. Jobs are distributed via a channel, workers process them concurrently, and results are collected. This pattern controls resource usage while maximizing throughput.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a worker pool in Go?
**Your Response:** "I implement a worker pool by creating a jobs channel and results channel, then launching a fixed number of worker goroutines. Each worker continuously reads from the jobs channel until it's closed, processes the job, and sends results to the results channel. I use a WaitGroup to track when all workers finish, and a separate goroutine to close the results channel after the WaitGroup completes. This pattern ensures I never exceed my desired concurrency limit, which is crucial for controlling resources like database connections. The key is using buffered channels to allow the producer to continue working even when workers are busy, and proper channel closure to signal completion."

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

### Explanation
Done channel pattern uses channel closure for cancellation. When the done channel is closed, all receivers unblock immediately, allowing coordinated shutdown of multiple goroutines.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you cancel goroutines in Go?
**Your Response:** "I use the done channel pattern for cancellation. I create a channel and pass it to all worker goroutines. Each worker uses a select statement to check for both work and the done channel. When I want to cancel all workers, I simply close the done channel. This causes all workers waiting on <-done to immediately unblock and exit. This pattern is simple but effective for basic cancellation. For more complex scenarios, I might use the context package which provides timeouts and cancellation propagation. The done channel approach is perfect for shutdown signals or any scenario where I need to coordinate stopping multiple goroutines simultaneously."

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

### Explanation
Context package manages cancellation signals, deadlines, and timeouts across API boundaries. context.WithTimeout creates a context that cancels automatically after a duration, and ctx.Done() provides a channel for cancellation notification.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use the context package in Go?
**Your Response:** "I use the context package to manage cancellation and timeouts across API boundaries. I create contexts with timeouts or cancellation capabilities using context.WithTimeout() or context.WithCancel(). I pass the context to functions that might need to be cancelled. Each function checks ctx.Done() in a select statement to detect cancellation. When the timeout expires or I call cancel(), all functions receive on ctx.Done() and can clean up and exit. The context propagates through the call chain, allowing cancellation to cascade through multiple function calls. This pattern is essential for building robust, cancelable operations in Go, especially for HTTP servers and long-running tasks."

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

### Explanation
sync.Once guarantees that a function executes exactly once, even if called concurrently from multiple goroutines. It's commonly used for thread-safe singleton initialization and one-time setup operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement thread-safe singletons in Go?
**Your Response:** "I use sync.Once to implement thread-safe singletons. I create a global once variable and call once.Do() with the initialization function. The first goroutine to call Do() executes the function, while all subsequent calls block until the first completes, then simply return without executing. This guarantees the initialization runs exactly once, even with concurrent access. It's much cleaner than using double-checked locking with mutexes. I use this pattern for database connections, configuration loading, or any expensive initialization that should only happen once. The sync.Once handles all the complexity of making singleton initialization thread-safe and efficient."

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

### Explanation
Atomic operations provide lock-free synchronization for simple operations like counters and flags. They're faster than mutexes for basic use cases and include operations like Add, Load, Store, and CompareAndSwap.

### How to Explain in Interview (Spoken style format)
**Interviewer:** When would you use atomic operations in Go?
**Your Response:** "I use atomic operations for simple synchronization needs like counters and flags. They're lock-free and faster than mutexes for basic operations. I use atomic.AddUint64() for incrementing counters, atomic.LoadUint64() for reading values, and atomic.StoreUint64() for writing. For more complex operations like compare-and-swap, I use atomic.CompareAndSwapUint64(). Atomic operations are ideal when I only need to synchronize a single variable. For anything more complex involving multiple variables, I'd use a mutex. The key benefit is performance - atomic operations avoid the overhead of locking and are implemented using CPU instructions that guarantee atomicity. They're perfect for metrics, reference counting, or simple state management."

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

### Explanation
errgroup coordinates multiple goroutines with error handling. If any goroutine returns an error, the group cancels the context and g.Wait() returns the first error. This provides fail-fast behavior for concurrent operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle errors in concurrent operations?
**Your Response:** "I use the errgroup package to coordinate multiple goroutines with error handling. I create an errgroup and spawn goroutines using g.Go(). If any goroutine returns an error, the errgroup cancels the context, causing all other goroutines to exit gracefully. When I call g.Wait(), it returns the first error encountered. This pattern is perfect for fan-out operations where I want to fail fast if any operation fails. The errgroup handles all the complexity of goroutine coordination and error propagation. For more advanced scenarios, I use errgroup.WithContext() which provides a cancellable context that's automatically cancelled when any goroutine fails. This makes concurrent error handling much cleaner and more reliable."

---

## 🔴 Common Pitfalls

1.  **Deadlock**: All goroutines are asleep.
    *   *Cause*: Waiting on a channel that no one is sending to (or vice versa).
2.  **Leaking Goroutines**: Goroutines that never exit.
    *   *Fix*: Ensure they have a way to return (e.g., proper channel closing or context cancellation).
3.  **Race Connection**: Multiple threads accessing variable without locking.
    *   *Check*: Run tests with `go test -race`.
