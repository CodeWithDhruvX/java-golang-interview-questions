# Golang Concurrency Interview Cheatsheet (Quick Crack)

High-yield patterns and solutions for concurrency questions.

---

## 🟢 1. Worker Pool (Limit Concurrent Processing)
**Pattern:**  
1. Create `jobs` channel (buffered) & `results` channel.
2. Launch `N` workers (goroutines).
3. Workers range over `jobs`.
4. Send jobs, then `close(jobs)`.
5. Wait/Close `results`.

```go
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        results <- j * 2
    }
}
// Main: Launch workers, send jobs, close jobs, wait, close results.
```

## 🟢 2. Fan-In (Merge Channels)
**Pattern:** Launch a goroutine for each input channel that forwards to one output channel. Wait for all to finish, then close output.

```go
func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    wg.Add(len(cs))
    
    for _, c := range cs {
        go func(c <-chan int) {
            defer wg.Done()
            for n := range c { out <- n }
        }(c)
    }
    
    go func() { wg.Wait(); close(out) }() // Closer goroutine
    return out
}
```

## 🟢 3. Semaphore (Limit Concurrent Access)
**Pattern:** Use a buffered channel of size `N`. Send to acquire, receive to release.

```go
sem := make(chan struct{}, 3) // Max 3 concurrent
for i := 0; i < 10; i++ {
    go func() {
        sem <- struct{}{} // Acquire
        defer func() { <-sem }() // Release
        // ... Critical Section ...
    }()
}
```

## 🟢 4. Broadcast Signal
**Pattern:** Closing a channel signals ALL listeners immediately.

```go
done := make(chan struct{})
// Workers:
go func() {
    <-done // Block until closed
    fmt.Println("Start work")
}()
// Main:
close(done) // Broadcasts start to all
```

## 🟢 5. Graceful Shutdown (Context)
**Pattern:** Pass `context.Context`. Workers check `ctx.Done()`.

```go
func worker(ctx context.Context) {
    select {
    case <-time.After(time.Second):
        fmt.Println("Work done")
    case <-ctx.Done(): // Timeout or Cancel scenarios
        fmt.Println("Cancelled:", ctx.Err())
    }
}
// Main: ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
```

## 🟢 6. Stop on First Error (ErrGroup)
**Pattern:** Use `errgroup` to cancel all other goroutines if one fails.

```go
g, ctx := errgroup.WithContext(context.Background())
g.Go(func() error {
    if err := doTask1(); err != nil { return err } // Cancels ctx
    return nil
})
g.Go(func() error {
    select {
    case <-ctx.Done(): return ctx.Err() // Check cancellation
    case <-doTask2(): return nil
    }
})
if err := g.Wait(); err != nil { fmt.Println("Failed:", err) }
```

## 🟢 7. Rate Limiter
**Pattern:** Block on a `time.Ticker` channel.

```go
limiter := time.Tick(200 * time.Millisecond) // 5 req/sec
for req := range requests {
    <-limiter // Wait for tick
    process(req)
}
```

## 🟢 8. Timeout (Select)
**Pattern:** Race your operation against `time.After()`.

```go
select {
case res := <-ch:
    fmt.Println(res)
case <-time.After(2 * time.Second):
    fmt.Println("Timeout")
}
```

## 🟢 9. Pipeline
**Pattern:** Series of stages where each stage is a function that takes an input channel and returns an output channel.

```go
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in { out <- n * n }
    }()
    return out
}
// Main: for n := range sq(gen(1, 2, 3)) { ... }
```

## 🟢 10. Ordered Results (Fan-Out/Fan-In)
**Pattern:** Use a slice and index to store results in order, rather than appending to a shared slice (race) or reading from channel (unordered).

```go
results := make([]int, n)
var wg sync.WaitGroup
for i := 0; i < n; i++ {
    wg.Add(1)
    go func(index int) { // Pass index
        defer wg.Done()
        results[index] = process(index) // Safe: distinct memory location
    }(i)
}
wg.Wait()
```

## 🟢 11. Thread-Safe Singleton
**Pattern:** `sync.Once`.

```go
var once sync.Once
var instance *Config
func GetConfig() *Config {
    once.Do(func() {
        instance = &Config{}
    })
    return instance
}
```

## 🟢 12. Ping Pong
**Pattern:** Two goroutines sharing one unbuffered channel.

```go
func player(name string, table chan int) {
    for {
        ball := <-table
        ball++
        table <- ball
    }
}
// Main: table <- 0 (Serve)
```

## 🟢 13. Goroutine Leak Prevention
**Rule:** NEVER start a goroutine without knowing how it will stop.
- **Consumer:** Should exit when channel is closed (`range`).
- **Producer:** Should close channel when done.
- **Blocked Send:** Ensure receiver exists or use buffer/select/context.

## 🟢 14. Deadlock Fixes
1. **Send to nil channel:** Blocks forever.
2. **Current Goroutine Wait:** Don't wait for yourself (e.g., `main` waiting on unbuffered channel sent by `main`). Use `go func()`.
