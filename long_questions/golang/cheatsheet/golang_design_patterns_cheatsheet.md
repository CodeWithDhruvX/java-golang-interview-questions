# Golang Design Patterns Cheatsheet

Idiomatic implementation of common design patterns in Go.

## 🟢 Creational Patterns

### 1. Functional Options Pattern
The standard way to create complex structs with optional configuration.
**Avoids:** Constructor explosion `NewServer(a, b, nil, nil, 10, ...)`

```go
type Server struct {
    host string
    port int
    tls  bool
}

type Option func(*Server)

func NewServer(opts ...Option) *Server {
    // Default config
    s := &Server{
        host: "localhost",
        port: 8080,
    }
    // Apply options
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// Options
func WithPort(port int) Option {
    return func(s *Server) {
        s.port = port
    }
}

func WithTLS() Option {
    return func(s *Server) {
        s.tls = true
    }
}

// Usage
srv := NewServer(WithPort(9000), WithTLS())
```

### 2. Singleton (Thread-Safe)
Using `sync.Once` guarantees execution exactly once.

```go
import "sync"

type Database struct{ conn string }

var (
    instance *Database
    once     sync.Once
)

func GetDatabase() *Database {
    once.Do(func() {
        // Expensive initialization here
        instance = &Database{conn: "postgresql://..."}
    })
    return instance
}
```

---

## 🟡 Structural Patterns

### 1. Middleware (Decorator)
Chaining behavior, commonly used in HTTP servers.

```go
type Handler func(string)

func Logger(next Handler) Handler {
    return func(msg string) {
        fmt.Println("[LOG] Before")
        next(msg)
        fmt.Println("[LOG] After")
    }
}

func Hello(msg string) {
    fmt.Println(msg)
}

// Usage
// protectedHander := Logger(Auth(Hello))
```

### 2. Adapter
Make incompatible interfaces work together.

```go
// Target Interface
type LightningPrinter interface {
    PrintFile(file string)
}

// Adaptee (Legacy)
type OldPrinter struct{}
func (p *OldPrinter) PrintComplex(s string) { fmt.Println("Old:", s) }

// Adapter
type PrinterAdapter struct {
    oldPrinter *OldPrinter
}
func (p *PrinterAdapter) PrintFile(file string) {
    p.oldPrinter.PrintComplex(file)
}
```

---

## 🔴 Concurrency Patterns

### 1. Worker Pool
Process jobs concurrently using a fixed number of goroutines.

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send 5 jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for a := 1; a <= 5; a++ {
        <-results
    }
}
```

### 2. Pipeline
Chain stages of processing channels.

```go
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums { out <- n }
        close(out)
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in { out <- n * n }
        close(out)
    }()
    return out
}

// Usage:
// for n := range sq(gen(2, 3)) { fmt.Println(n) } // 4, 9
```

### 3. Fan-Out / Fan-In
**Fan-Out:** Distribute work to multiple workers (see Worker Pool).
**Fan-In:** Multiplex results from multiple channels into one.

```go
func merge(cs ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    output := func(c <-chan int) {
        for n := range c { out <- n }
        wg.Done()
    }

    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
```

## 🟣 Concurrency Interview Questions
> **[View Solutions](golang_concurrency_questions_solutions.md)**

1. One goroutine produces numbers, multiple worker goroutines process them, and the main goroutine collects results.
2. Implement a fixed-size worker pool to process N jobs concurrently.
3. Merge multiple input channels into a single output channel.
4. Broadcast the same message to multiple goroutines.
5. Limit the number of concurrent goroutines accessing a shared resource.
6. Print numbers from 1 to N in order using goroutines.
7. Gracefully shut down all goroutines when a timeout occurs.
8. Stop all goroutines when any worker returns an error.
9. Fix a program that deadlocks due to incorrect channel usage.
10. Implement a rate limiter allowing N operations per second.
11. Build a pipeline where each stage runs in its own goroutine.
12. Read from a channel with a timeout using select.
13. Identify and fix a goroutine leak.
14. Rewrite a function using send-only and receive-only channels.
15. Implement a thread-safe counter using goroutines and channels.
16. Handle a many-readers, few-writers scenario safely.
17. Process tasks concurrently but return results in the original order.
18. Dynamically scale worker goroutines based on load.
19. Implement a semaphore using a buffered channel.
20. Coordinate two goroutines to alternately print “ping” and “pong”.
21. Pass context through a worker pool and cancel on timeout.
22. Cancel all goroutines when the first error occurs using context.
23. Collect all errors from concurrent goroutines safely.
24. Prevent slow consumers from blocking fast producers.
25. Implement non-blocking send and receive using select and default.
26. Ensure all goroutines exit cleanly when input channel is closed.
27. Safely close a results channel used by multiple goroutines.
28. Process jobs concurrently with retry logic on failure.
29. Limit concurrent HTTP requests using goroutines and channels.
30. Fan out work to multiple goroutines and fan in results safely.