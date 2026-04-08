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

### Explanation
Worker pool pattern limits concurrency to a fixed number of workers. Jobs are distributed via a channel, workers process them concurrently, and results are collected. This pattern controls resource usage while maximizing throughput.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a worker pool in Go?
**Your Response:** "I implement a worker pool by creating a jobs channel and results channel, then launching a fixed number of worker goroutines. Each worker continuously reads from the jobs channel until it's closed, processes the job, and sends results to the results channel. I use a WaitGroup to track when all workers finish, and a separate goroutine to close the results channel after the WaitGroup completes. This pattern ensures I never exceed my desired concurrency limit, which is crucial for controlling resources like database connections. The key is using buffered channels to allow the producer to continue working even when workers are busy, and proper channel closure to signal completion."

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

### Explanation
Fan-in pattern merges multiple input channels into one output channel. Each input channel is handled by a dedicated goroutine that forwards values. A WaitGroup ensures proper closure when all inputs are exhausted.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you merge multiple channels in Go?
**Your Response:** "I use the fan-in pattern to merge multiple channels. I create an output channel and a WaitGroup. For each input channel, I spawn a goroutine that reads from that channel and forwards values to the output. Each goroutine calls Done() when its input closes. I also spawn a goroutine that waits for the WaitGroup and then closes the output. This ensures all values flow through correctly and the output closes only after all inputs are done. The pattern is essential when I need to collect results from multiple concurrent operations into a single stream. It's a fundamental Go concurrency pattern that combines multiple data sources efficiently."

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

### Explanation
Semaphore pattern uses a buffered channel to limit concurrent access. The channel capacity represents the maximum concurrent operations. Goroutines acquire by sending and release by receiving, blocking when the limit is reached.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a semaphore in Go?
**Your Response:** "I implement semaphores using buffered channels where the capacity represents the maximum concurrent operations. Before accessing the shared resource, a goroutine sends to the channel to acquire a permit. If the channel is full, the goroutine blocks until another finishes and receives from the channel. I use defer to ensure the permit is always released. This pattern is elegant because it uses Go's built-in blocking behavior to control concurrency. I typically use struct{} as the element since I only need the signaling, not the data. This approach is perfect for rate limiting, connection pooling, or any scenario where I need to control resource usage."

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

### Explanation
Broadcast signaling uses channel closure - when a channel is closed, all receivers immediately unblock and receive the zero value. This provides an efficient way to signal multiple goroutines simultaneously.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you broadcast a signal to multiple goroutines?
**Your Response:** "I use channel closure for broadcasting. When I close a channel, all goroutines waiting to receive from it immediately unblock. I create a channel, spawn multiple goroutines that wait to receive from it, and when I want to broadcast, I simply close the channel. This is much more efficient than sending individual messages. The pattern works because closing a channel is a broadcast operation - all waiting receivers are notified simultaneously. I typically use a struct{} channel since I'm only interested in the signal, not the data. This approach is perfect for shutdown signals, coordination, or any scenario where I need to notify multiple goroutines at once."

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

### Explanation
Graceful shutdown uses context cancellation. A context with timeout/cancellation is passed to workers, which use select to check both work completion and ctx.Done(). This allows clean shutdown with resource cleanup.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement graceful shutdown in Go?
**Your Response:** "I implement graceful shutdown using context cancellation. I create a context with timeout or cancellation capability and pass it to all worker goroutines. Each worker uses a select statement to check for both work completion and the context's Done() channel. When the timeout occurs or I call cancel(), the context is cancelled, causing all workers to receive on Done() and exit gracefully. This approach allows workers to clean up resources, save state, or complete in-progress operations before shutting down. The context pattern is the idiomatic way to handle cancellation and timeouts in Go, providing clean coordination across multiple goroutines."

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

### Explanation
The errgroup pattern coordinates multiple goroutines and cancels all when any returns an error. errgroup.WithContext creates a group with a cancellable context that's automatically cancelled when any goroutine fails.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you stop all goroutines when one fails?
**Your Response:** "I use the errgroup package to coordinate multiple goroutines with error handling. I create an errgroup with a context using errgroup.WithContext(), then spawn goroutines using g.Go(). If any goroutine returns an error, the errgroup automatically cancels the context, causing all other goroutines to receive on ctx.Done() and exit gracefully. When I call g.Wait(), it returns the first error encountered. This pattern is perfect for fan-out operations where I want to fail fast if any operation fails. The errgroup handles all the complexity of goroutine coordination and error propagation, making concurrent error handling much cleaner and more reliable."

## 🟢 7. Rate Limiter
**Pattern:** Block on a `time.Ticker` channel.
```go
package main
import 
(
      "fmt"
      "time"
)

func main() {
      request:=[]int{1,2,3,4,5}
      limiter:=time.Tick(200*time.Millisecond)

      for i:=range request {
            <-limiter
            process(i)
      }      

}

func process(req int){
      fmt.Printf("Processing request %d at %v \n",req,time.Now().Format("15:04:05.000"))
}
```

### Explanation
Rate limiting uses time.Ticker to control operation frequency. Each operation blocks waiting for a ticker tick before proceeding, ensuring a maximum rate of operations per time period.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement rate limiting in Go?
**Your Response:** "I implement rate limiting using time.Ticker. I create a ticker that fires at the rate I want to allow - for example, 200ms between operations for 5 requests per second. Before processing each request, I wait to receive from the ticker channel. This blocks until the next tick is available, ensuring I never exceed my desired rate. The ticker approach is simple and effective for basic rate limiting. For more complex scenarios, I might use a token bucket algorithm or the golang.org/x/time/rate package, but the ticker pattern works well for most use cases. It's a clean way to prevent overwhelming downstream services or APIs."

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

### Explanation
Timeout pattern uses select to race an operation against time.After(). The first case to complete wins, allowing operations to be cancelled if they take too long without blocking indefinitely.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement timeouts in Go?
**Your Response:** "I implement timeouts using select with time.After(). I create a select statement with two cases - one for the operation I want to perform, and another for time.After() with my timeout duration. Whichever case completes first wins. If the operation completes before the timeout, I get the result. If the timeout fires first, I handle the timeout case. This pattern prevents my program from blocking indefinitely on slow operations. It's the idiomatic way to add timeouts to channel operations, HTTP requests, or any operation that might hang. The select statement makes the timeout logic clean and readable."

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

### Explanation
Pipeline pattern chains processing stages where each stage is a goroutine that transforms data from input to output channels. Each stage runs concurrently, enabling efficient data flow processing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a pipeline pattern in Go?
**Your Response:** "I implement pipelines by chaining stages where each stage is a function that takes an input channel and returns an output channel. Each stage runs its own goroutine that reads from the input channel, processes the data, and sends results to the output channel. I chain these stages together - for example, generator feeds into squarer, which feeds into printer. This creates a concurrent processing pipeline where each stage works independently. The key is that each stage closes its output channel when done, allowing the next stage to know when processing is complete. This pattern is perfect for data processing workflows where I want to apply multiple transformations to a stream of data efficiently."

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

### Explanation
Ordered results pattern uses a pre-allocated slice with index-based storage. Each goroutine writes to its designated index position, avoiding race conditions while maintaining original order despite concurrent processing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you process tasks concurrently but return results in order?
**Your Response:** "I use a pre-allocated slice with index-based storage. I create a results slice with the exact size I need, then spawn goroutines that each receive their index as a parameter. Each goroutine writes its result to its designated index position in the slice. Since each goroutine writes to a different memory location, there are no race conditions. After all goroutines finish using a WaitGroup, I can iterate through the results slice in order. This approach gives me the performance benefits of concurrent processing while maintaining the original order. It's much cleaner than trying to coordinate ordering through channels or complex synchronization."

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
