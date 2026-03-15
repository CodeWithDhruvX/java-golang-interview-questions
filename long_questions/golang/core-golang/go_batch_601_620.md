## 🔸 Advanced Concurrency Patterns (Questions 601-620)

### Question 601: How do you implement a fan-in pattern in Go?

**Answer:**
Fan-in merges multiple input channels into a single output channel.

```go
func FanIn(input1, input2 <-chan string) <-chan string {
    c := make(chan string)
    go func() {
        defer close(c)
        var wg sync.WaitGroup
        wg.Add(2)
        output := func(ch <-chan string) {
            defer wg.Done()
            for v := range ch { c <- v }
        }
        go output(input1)
        go output(input2)
        wg.Wait()
    }()
    return c
}
```

### Explanation
The fan-in pattern in Go merges multiple input channels into a single output channel. This is implemented by creating a goroutine that reads from all input channels concurrently and writes their values to a single output channel. A WaitGroup ensures all input channels are fully consumed before closing the output channel.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a fan-in pattern in Go?
**Your Response:** "I implement the fan-in pattern by creating a function that takes multiple input channels and returns a single output channel. Inside the function, I launch a goroutine that reads from all input channels concurrently using separate goroutines for each input. Each input goroutine forwards values to the output channel. I use a WaitGroup to track when all input channels are done, then close the output channel. This pattern is useful when I need to merge results from multiple concurrent operations into a single stream. It's the opposite of fan-out, and Go's channels make this pattern very natural to implement with proper coordination using WaitGroups."

---

### Question 602: How do you implement a fan-out pattern in Go?

**Answer:**
Fan-out distributes work from one channel to multiple worker goroutines.

```go
func FanOut(ch <-chan int, workers int) {
    var wg sync.WaitGroup
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for item := range ch {
                process(id, item)
            }
        }(i)
    }
    wg.Wait()
}
```

### Explanation
The fan-out pattern distributes work from one channel to multiple worker goroutines. Multiple workers read from the same channel, allowing concurrent processing of items. Each worker processes items independently, and a WaitGroup ensures all workers complete before the function returns.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a fan-out pattern in Go?
**Your Response:** "I implement the fan-out pattern by creating multiple worker goroutines that all read from the same input channel. I launch a fixed number of workers, each running in its own goroutine, and they all compete to read from the shared channel. This naturally distributes the work among the workers since Go's channels provide fair scheduling. Each worker processes items independently until the channel is closed. I use a WaitGroup to track when all workers have finished processing. This pattern is perfect for parallelizing work when I have more tasks than I can handle with a single goroutine. It's especially effective for CPU-bound tasks where I want to utilize multiple cores."

---

### Question 603: How do you prevent goroutine leaks in producer-consumer patterns?

**Answer:**
Ensure that if the consumer stops (e.g., error or context cancel), the producer doesn't block forever trying to send to the channel.
**Solution:** Use `select` with `ctx.Done()`.

```go
func produce(ctx context.Context, out chan<- int) {
    for {
        select {
        case <-ctx.Done():
            return // Stop producing
        case out <- rand.Int():
            // Sent
        }
    }
}
```

### Explanation
Goroutine leaks in producer-consumer patterns occur when producers block forever on channel sends after consumers have stopped. The solution is to use context cancellation with select statements, allowing producers to detect when consumers are done and exit gracefully instead of blocking indefinitely.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you prevent goroutine leaks in producer-consumer patterns?
**Your Response:** "I prevent goroutine leaks by ensuring producers don't block forever when consumers stop. The key is to use context cancellation with select statements. I pass a context to the producer and use a select that checks both `ctx.Done()` and the channel send operation. When the consumer stops or an error occurs, I cancel the context, which causes the producer to receive the done signal and exit gracefully instead of blocking on the channel send. This approach prevents the common leak scenario where producers keep running even after consumers are gone, consuming resources unnecessarily. Context cancellation provides a clean, idiomatic way to coordinate goroutine lifecycles in Go."

---

### Question 604: How would you create a semaphore in Go?

**Answer:**
Use a buffered channel. It limits concurrent access to a resource.

```go
type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
    return make(chan struct{}, n)
}

func (s Semaphore) Acquire() { s <- struct{}{} }
func (s Semaphore) Release() { <-s }

// Usage
sem.Acquire()
go doWork()
sem.Release()
```

### Explanation
Semaphores in Go are implemented using buffered channels. The channel capacity represents the maximum number of concurrent accesses. Acquiring sends an empty struct to the channel (blocking if full), and releasing receives from the channel (allowing another goroutine to acquire). This pattern effectively limits concurrent access to resources.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you create a semaphore in Go?
**Your Response:** "I create semaphores in Go using buffered channels, which is a very elegant and idiomatic approach. I define a Semaphore type as a channel of empty structs, where the channel capacity represents the maximum number of concurrent accesses. To acquire the semaphore, I send an empty struct to the channel - this blocks if the channel is full. To release, I receive from the channel, allowing another goroutine to acquire. This pattern effectively limits concurrent access to resources like database connections or API rate limits. The beauty of this approach is that it leverages Go's built-in channel semantics for blocking and coordination, making it simple yet powerful for controlling concurrency."

---

### Question 605: What’s the difference between `sync.WaitGroup` and `sync.Cond`?

**Answer:**
- **`WaitGroup`:** Waits for a *count* of goroutines to finish. (Join pattern).
- **`Cond`:** A condition variable. Allows goroutines to suspend execution and wait for a *signal* (event) to occur (e.g., "Queue is not empty"). Broadscast wakes up *all* waiting goroutines.

### Explanation
WaitGroup is used for waiting for a specific number of goroutines to complete (join pattern). Cond (condition variable) allows goroutines to wait for specific events or conditions to occur. Cond supports Signal() to wake one waiter or Broadcast() to wake all waiters, making it suitable for coordinating around state changes.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between `sync.WaitGroup` and `sync.Cond`?
**Your Response:** "The key difference is their purpose and coordination model. WaitGroup is for waiting for a specific number of goroutines to finish - it's essentially a join pattern where I add to the counter before starting goroutines and wait for them to finish. Cond is a condition variable that allows goroutines to suspend execution and wait for specific events or conditions to occur. With Cond, goroutines wait until they receive a signal that some condition has changed. Cond supports both Signal() to wake one waiting goroutine and Broadcast() to wake all of them. I use WaitGroup when I know exactly how many goroutines I need to wait for, and Cond when I need to coordinate around state changes like 'queue not empty' or 'data available'."

---

### Question 606: How do you implement a pub-sub model in Go?

**Answer:**
Maintain a map of topics to list of channels.
When `Publish(topic, msg)` is called, iterate over the list and non-blocking send to subscribers (or use a dedicated goroutine per sub).

### Explanation
Pub-sub in Go is implemented by maintaining a map where each topic maps to a list of subscriber channels. When publishing, the system iterates through all subscribers for that topic and sends the message. Non-blocking sends or dedicated goroutines per subscriber prevent one slow subscriber from blocking the entire system.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement a pub-sub model in Go?
**Your Response:** "I implement pub-sub by maintaining a map where each topic maps to a list of subscriber channels. When a client wants to subscribe to a topic, I add their channel to the list. When publishing a message, I iterate through all subscribers for that topic and send them the message. I use non-blocking sends or dedicated goroutines per subscriber to prevent one slow subscriber from blocking the entire system. This approach allows multiple subscribers to receive messages independently without interfering with each other. Channels provide the perfect foundation for this pattern since they naturally handle the message passing and buffering needed for pub-sub communication."

---

### Question 607: How do you use a context to timeout multiple goroutines?

**Answer:**
Pass the same `Context` to all of them. `context.WithTimeout` propagates the cancellation signal.

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

go worker(ctx)
go worker(ctx)
// Inside worker:
select {
case <-ctx.Done():
    return // Clean up
}
```

### Explanation
Context timeout in Go works by creating a context with timeout and passing it to all goroutines. The context automatically cancels when the timeout expires, and all goroutines monitoring ctx.Done() receive the cancellation signal simultaneously, allowing coordinated shutdown across multiple goroutines.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use a context to timeout multiple goroutines?
**Your Response:** "I use context to timeout multiple goroutines by creating a single context with timeout using `context.WithTimeout()` and passing it to all goroutines. The context automatically cancels when the timeout expires, and all goroutines monitoring `ctx.Done()` receive the cancellation signal simultaneously. Inside each worker, I use a select statement that checks for `ctx.Done()` and performs cleanup when the timeout occurs. This approach provides coordinated shutdown across multiple goroutines without complex manual coordination. The context propagates the cancellation signal through the call chain, ensuring all related operations stop gracefully when the timeout expires. It's the idiomatic way to handle timeouts in concurrent Go programs."

---

### Question 608: How do you build a rate-limiting queue with channels?

**Answer:**
Combine a buffered channel (queue) with a Ticker.

```go
requests := make(chan int, 100)
ticker := time.NewTicker(200 * time.Millisecond) // 5 req/sec

go func() {
    for req := range requests {
        <-ticker.C // Wait for tick
        process(req)
    }
}()
```

### Explanation
Rate-limiting queues in Go combine a buffered channel for the queue with a time.Ticker for rate control. The ticker emits signals at a fixed interval, and each request waits for a tick before being processed. This ensures a maximum processing rate regardless of how many requests are queued.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you build a rate-limiting queue with channels?
**Your Response:** "I build rate-limiting queues by combining a buffered channel for the queue with a time.Ticker for rate control. I create a buffered channel to hold incoming requests and a ticker that emits signals at fixed intervals - for example, every 200ms for 5 requests per second. In the processing goroutine, I read from the request channel but wait for a ticker signal before processing each request. This ensures that no matter how many requests are queued, they're processed at a maximum rate determined by the ticker interval. The buffered channel allows for burst handling while the ticker enforces the long-term rate limit. This approach is simple, efficient, and leverages Go's built-in timing primitives."

---

### Question 609: What is a worker pool, and how do you implement it?

**Answer:**
A fixed number of goroutines that pull tasks from a shared channel. Helps control resource usage.
See Question 602 (Fan-out). Ideally, the number of workers matches `runtime.NumCPU()`.

### Explanation
Worker pools consist of a fixed number of goroutines that continuously pull tasks from a shared channel. This pattern controls resource usage by limiting concurrent operations. The optimal number of workers typically matches the number of CPU cores for CPU-bound tasks, but can be higher for I/O-bound operations.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a worker pool, and how do you implement it?
**Your Response:** "A worker pool is a pattern where I create a fixed number of goroutines that continuously pull tasks from a shared channel. This helps control resource usage by limiting how many operations run concurrently. I implement it by launching worker goroutines that all read from the same task channel. For CPU-bound tasks, I typically set the number of workers to match `runtime.NumCPU()`, but for I/O-bound operations, I might use more workers. This pattern prevents the system from being overwhelmed by too many concurrent operations while still utilizing available resources efficiently. It's essentially the fan-out pattern with a bounded number of workers, providing predictable resource usage and performance."

---

### Question 610: How do you handle backpressure in channel-based designs?

**Answer:**
Backpressure occurs when the consumer is slower than the producer.
1.  **Bounded Channels:** Use unbuffered or small-buffered channels. This blocks the producer, naturally slowing it down.
2.  **Dropping:** `select { case ch <- msg: default: log.Println("Dropping") }` (Load Shedding).

### Explanation
Backpressure in channel-based designs occurs when consumers can't keep up with producers. Bounded channels handle this by blocking producers when the channel is full, naturally slowing down production. Load shedding uses select with a default case to drop messages when the channel is full, preventing system overload.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle backpressure in channel-based designs?
**Your Response:** "I handle backpressure using two main approaches. First, I use bounded channels with small or zero buffers, which naturally apply backpressure by blocking producers when the channel is full. This forces producers to slow down when consumers can't keep up. Second, I implement load shedding using a select statement with a default case - if the channel is full, I drop the message and log it rather than blocking. The approach I choose depends on the requirements - bounded channels for guaranteed delivery where it's better to wait than lose data, or load shedding for systems where it's acceptable to lose some data to maintain responsiveness. Bounded channels provide natural flow control, while load shedding protects the system from being overwhelmed."

---

### Question 611: How do you gracefully shut down workers?

**Answer:**
1.  **Close Channel:** `close(jobs)`. Workers loop `for j := range jobs` will terminate once channel is empty and closed.
2.  **WaitGroup:** Main waits for workers to exit.

```go
close(jobs)
wg.Wait()
```

### Explanation
Graceful shutdown of workers involves closing the jobs channel, which causes workers reading with range to terminate when the channel is empty and closed. A WaitGroup ensures the main goroutine waits for all workers to finish before exiting completely.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you gracefully shut down workers?
**Your Response:** "I gracefully shut down workers by closing the jobs channel and using a WaitGroup to track completion. When I'm ready to shut down, I call `close(jobs)` which signals that no more work will be coming. Workers that are using `for j := range jobs` will automatically terminate when the channel is empty and closed. I use a WaitGroup to track when all workers have finished their current work and exited. The main goroutine waits on `wg.Wait()` before proceeding with shutdown. This approach ensures all in-progress work is completed while preventing new work from being accepted, providing a clean and coordinated shutdown process without losing any in-flight tasks."

---

### Question 612: How do you use `sync.Cond` for event signaling?

**Answer:**
It requires a `Locker` (Mutex).

```go
var mu sync.Mutex
cond := sync.NewCond(&mu)

// Waiter
go func() {
    mu.Lock()
    for !condition {
        cond.Wait() // Unlocks and suspends
    }
    // Condition met
    mu.Unlock()
}()

// Signaler
mu.Lock()
condition = true
cond.Signal() // Wakes one waiter
mu.Unlock()
```

### Explanation
sync.Cond is a condition variable that requires a Mutex for synchronization. Waiters lock the mutex, check the condition, and call cond.Wait() which atomically unlocks and suspends the goroutine. Signalers lock the mutex, change the condition, and call Signal() to wake one waiter or Broadcast() to wake all.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you use `sync.Cond` for event signaling?
**Your Response:** "I use sync.Cond for event signaling by creating it with a Mutex. Waiters lock the mutex, check if their condition is met, and if not, call cond.Wait() which atomically unlocks the mutex and suspends the goroutine. When the condition changes, the signaler locks the mutex, updates the condition, and calls cond.Signal() to wake one waiting goroutine or cond.Broadcast() to wake all of them. The key is that cond.Wait() automatically handles the unlock/suspend and lock/resume coordination, preventing race conditions. This pattern is useful when goroutines need to wait for specific state changes, like 'queue not empty' or 'data available', providing efficient coordination without busy waiting."

---

### Question 613: How do you prioritize tasks in concurrent processing?

**Answer:**
Go channels are FIFO, no priority support.
**Solution:**
1.  Use two channels: `highPriority` and `lowPriority`.
2.  Worker uses `select` with priority check:

```go
select {
case job := <-highPriority:
    process(job)
default:
    select {
    case job := <-highPriority: process(job)
    case job := <-lowPriority: process(job)
    }
}
```

### Explanation
Go channels don't support priority natively since they're FIFO. Priority processing is implemented using separate channels for different priority levels. Workers use nested select statements - first checking high priority, then using a default case to check both channels if no high priority items are available.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you prioritize tasks in concurrent processing?
**Your Response:** "Since Go channels are FIFO and don't support priority natively, I implement priority processing using separate channels for different priority levels. I create multiple channels like `highPriority` and `lowPriority`, then workers use nested select statements. The outer select checks for high priority items first. If no high priority items are available, the default case runs an inner select that checks both channels, still preferring high priority items but falling back to low priority ones. This approach ensures high priority tasks are processed as soon as they're available, while low priority tasks only get processed when no high priority tasks exist. It's a clean way to implement priority queuing using Go's channel primitives."

---

### Question 614: How do you avoid starvation in goroutines?

**Answer:**
Starvation happens if a goroutine never gets CPU time (e.g., tight loops without function calls/preemption points in old Go versions).
**Fix:**
- In newer Go (1.14+), the scheduler is preemptive (asynchronously signals loose loops).
- Avoid holding locks for long durations.
- Ensure all channels are serviced fairly.

### Explanation
Goroutine starvation occurs when a goroutine never gets CPU time, historically due to tight loops without preemption points. Modern Go (1.14+) has a preemptive scheduler that helps prevent this. Additional prevention includes avoiding long-held locks and ensuring fair channel servicing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you avoid starvation in goroutines?
**Your Response:** "I avoid goroutine starvation by leveraging modern Go's preemptive scheduler and following good practices. Since Go 1.14, the scheduler is preemptive and can interrupt tight loops, which has largely solved the classic starvation problem. I still avoid holding locks for long durations to prevent blocking other goroutines, and I ensure all channels are serviced fairly by not giving preference to any particular channel in select statements. I also avoid patterns where one goroutine could monopolize CPU time. The combination of the modern preemptive scheduler and careful design around lock duration and channel fairness effectively prevents starvation in most Go applications."

---

### Question 615: How do you detect race conditions without `-race` flag?

**Answer:**
It is extremely difficult.
- **Code Review:** Look for shared mutable state accessed without locks.
- **Tools:** `staticcheck` can find some issues.
- **Panic:** Sometimes race conditions cause "concurrent map read and map write" panics at runtime.
**Best Practice:** Always use `-race` in CI/Test.

### Explanation
Detecting race conditions without the race flag is extremely difficult. Code review can identify shared mutable state without proper synchronization. Tools like staticcheck find some patterns. Runtime panics sometimes indicate races. The best practice is always using the race detector in CI and testing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you detect race conditions without `-race` flag?
**Your Response:** "Detecting race conditions without the race flag is extremely difficult and not reliable. I can try to identify potential issues through code review by looking for shared mutable state accessed without proper locks or synchronization. Tools like staticcheck can find some problematic patterns, and sometimes race conditions cause runtime panics like 'concurrent map read and map write'. However, these methods are not comprehensive. The best practice is to always use the `-race` flag in CI pipelines and during testing. The race detector is the most reliable tool for finding race conditions, and trying to detect them without it is like trying to find a needle in a haystack. I emphasize using the proper tools rather than relying on manual detection methods."

---

### Question 616: How do you trace execution flow in concurrent systems?

**Answer:**
1.  **Trace ID:** Pass `ctx` with trace ID everywhere.
2.  **Structured Logging:** Log "Worker 5 started job X", "Worker 5 finished job X".
3.  **Go Trace Tool:** `go tool trace` visualizes how goroutines hand off control and block.

### Explanation
Tracing execution flow in concurrent systems involves passing trace IDs through context for correlation, structured logging for tracking operations, and Go's trace tool for visualization. Context propagation maintains trace context across goroutine boundaries, while the trace tool provides detailed runtime behavior visualization.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you trace execution flow in concurrent systems?
**Your Response:** "I trace execution flow in concurrent systems using three main approaches. First, I pass trace IDs through context everywhere, which maintains correlation across goroutine boundaries and function calls. Second, I use structured logging to log key events like 'Worker 5 started job X' and 'Worker 5 finished job X' with the trace ID. Third, I use Go's trace tool with `go tool trace` which visualizes how goroutines hand off control and block. The combination of context propagation, structured logging, and trace visualization gives me a complete picture of how work flows through the concurrent system. This approach helps me understand the execution order, identify bottlenecks, and debug complex concurrent interactions."

---

### Question 617: How do you implement exponential backoff with retries in goroutines?

**Answer:**
Simulate wait before retrying.

```go
func retry(op func() error) {
    delay := 100 * time.Millisecond
    for i := 0; i < 5; i++ {
        err := op()
        if err == nil { return }
        time.Sleep(delay)
        delay *= 2
    }
}
```

### Explanation
Exponential backoff with retries in goroutines implements a delay between retry attempts that doubles each time. Starting with a small delay, each failed attempt increases the wait time exponentially, providing increasing backoff pressure while allowing for eventual recovery from transient failures.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you implement exponential backoff with retries in goroutines?
**Your Response:** "I implement exponential backoff with retries by creating a retry function that doubles the delay between attempts. I start with a small delay like 100 milliseconds, and after each failed attempt, I sleep for the current delay duration, then double the delay for the next attempt. This creates an exponential backoff pattern - 100ms, 200ms, 400ms, 800ms, and so on. This approach is perfect for handling transient failures like network issues or temporary service unavailability. The exponential delay prevents overwhelming the struggling service while still allowing for quick recovery from brief glitches. I typically set a maximum number of retries to prevent infinite loops."

---

### Question 618: How do you structure long-running daemons with concurrency?

**Answer:**
Use a "Supervisor" pattern.
- **Main:** Starts services. Listens for OS Signals (SIGINT).
- **Service:** Accepts `ctx`. Runs loop `for { select { case <-ctx.Done(): return } }`.
- **ErrGroup:** Use `errgroup.Group` to manage multiple daemon services. If one crashes, it cancels the context for all.

### Explanation
Long-running daemons use a supervisor pattern where the main process starts services and listens for OS signals. Services accept context and run in loops checking for cancellation. ErrGroup manages multiple services, where one failure cancels all services through context propagation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you structure long-running daemons with concurrency?
**Your Response:** "I structure long-running daemons using a supervisor pattern. The main process starts all the services and listens for OS signals like SIGINT for graceful shutdown. Each service accepts a context and runs in a loop that checks for `ctx.Done()` to know when to shut down. I use an errgroup.Group to manage multiple daemon services - if one service crashes, it cancels the context for all services, ensuring coordinated shutdown. This pattern provides clean startup, graceful shutdown, and failure handling. The context propagation ensures all services respond to shutdown signals quickly, while the errgroup provides proper error handling and coordination between services."

---

### Question 619: How would you implement circuit breakers in Go?

**Answer:**
State machine: Closed (Normal) -> Open (Error threshold reached, fail fast) -> Half-Open (Test recovery).
Use `Sony/gobreaker` library.
It wraps a function execution. If failures > X%, prevents calls for Y seconds.

### Explanation
Circuit breakers implement a state machine with three states: Closed (normal operation), Open (failure threshold reached, fail fast), and Half-Open (testing recovery). Libraries like gobreaker wrap function execution and automatically transition between states based on failure rates and time thresholds.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How would you implement circuit breakers in Go?
**Your Response:** "I implement circuit breakers using a state machine approach with three states: Closed for normal operation, Open when the failure threshold is reached, and Half-Open to test recovery. I typically use the `Sony/gobreaker` library which provides a robust implementation. The circuit breaker wraps function execution and tracks failure rates. When failures exceed a threshold, it opens the circuit and fails fast without calling the actual function. After a timeout, it enters half-open state to test if the service has recovered. This pattern prevents cascading failures and provides resilience when calling external services. It's essential for microservices architectures where downstream services might become temporarily unavailable."

---

### Question 620: How do you handle concurrent map access with minimal locking?

**Answer:**
1.  **`sync.RWMutex`:** Allows multiple readers, one writer. Good for Read-Heavy workloads.
2.  **`sync.Map`:** Specialized for append-only caches or when keys are disjoint.
3.  **Sharding:** Create `[32]*sync.Mutex` and `[32]map`. Hash key to pick shard. Reduces lock contention by 32x.

### Explanation
Concurrent map access with minimal locking uses three approaches. RWMutex allows multiple concurrent readers with exclusive writers. sync.Map is optimized for specific access patterns like append-only caches. Sharding divides the map into multiple segments with separate locks, reducing contention by spreading access across multiple locks.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle concurrent map access with minimal locking?
**Your Response:** "I handle concurrent map access with minimal locking using three approaches depending on the use case. For read-heavy workloads, I use `sync.RWMutex` which allows multiple concurrent readers while ensuring exclusive access for writers. For specific patterns like append-only caches or when keys are disjoint, I use `sync.Map` which is optimized for these scenarios. For high-contention scenarios, I implement sharding by creating multiple maps with separate locks - I hash the key to pick which shard to use, reducing lock contention by a factor of 32 or more. The choice depends on the access pattern - RWMutex for general read-heavy use, sync.Map for specialized patterns, and sharding for high-contention scenarios where performance is critical."

---
