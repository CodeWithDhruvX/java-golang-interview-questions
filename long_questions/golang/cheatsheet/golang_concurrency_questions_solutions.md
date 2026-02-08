# Golang Concurrency Questions & Solutions

This guide provides idiomatic solutions to common Golang concurrency interview questions.

---

## 1. Producer-Consumer Worker Pool
**Question:** One goroutine produces numbers, multiple worker goroutines process them, and the main goroutine collects results.

```go
package main

import (
	"fmt"
	"sync"
)

func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func worker(id int, in <-chan int, out chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		out <- fmt.Sprintf("Worker %d processed %d", id, n)
	}
}

func main() {
	jobs := producer(1, 2, 3, 4, 5)
	results := make(chan string, 5)
	var wg sync.WaitGroup

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Close results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}
```

## 2. Fixed-Size Worker Pool
**Question:** Implement a fixed-size worker pool to process N jobs concurrently.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Millisecond * 500) // Simulate work
		results <- j * 2
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for r := range results {
		fmt.Println("Result:", r)
	}
}
```

## 3. Merge Channels (Fan-In)
**Question:** Merge multiple input channels into a single output channel.

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
		for n := range c {
			out <- n
		}
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

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() { c1 <- 1; close(c1) }()
	go func() { c2 <- 2; close(c2) }()

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}
```

## 4. Broadcast Message
**Question:** Broadcast the same message to multiple goroutines.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Closing a channel signals all receivers
	done := make(chan struct{})

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-done // Wait for signal
			fmt.Printf("Worker %d received signal\n", id)
		}(i)
	}

	fmt.Println("Broadcasting signal...")
	close(done) // Broadcast
	wg.Wait()
}
```

## 5. Limiting Concurrency (Semaphore)
**Question:** Limit the number of concurrent goroutines accessing a shared resource.
## The Core Logic

The magic happens in these three specific lines:

1. **The Capacity:** `sem := make(chan struct{}, 3)`
    
    - This creates a "room" that can only hold **3 items** at a time.
        
2. **The Acquire:** `sem <- struct{}{}`
    
    - Before starting work, the goroutine tries to put a piece of data into the channel. If the channel already has 3 items, this line **blocks** and the goroutine waits.
        
3. **The Release:** `<-sem`
    
    - Once the work is done, the goroutine pulls an item out of the channel, making a "slot" available for the next waiting worker.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const maxConcurrent = 3
	sem := make(chan struct{}, maxConcurrent) // Semaphore
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{} // Acquire
			fmt.Printf("Worker %d accessed resource\n", id)
			time.Sleep(500 * time.Millisecond)
			<-sem // Release
			
		}(i)
	}
	wg.Wait()
}
```

## 6. Sequencing Goroutines
**Question:** Print numbers from 1 to N in order using goroutines.

```go
package main

import (
	"fmt"
)

func main() {
	n := 10
	ch := make(chan int)

	// Receiver
	go func() {
		for i := 1; i <= n; i++ {
			fmt.Println(<-ch)
		}
	}()

	// Sender (could be multiple, but channel ensures order if sent sequentially)
	for i := 1; i <= n; i++ {
		ch <- i
	}
}
```
*Note: If multiple goroutines need to print in order (e.g. Worker 1 prints 1, Worker 2 prints 2), you'd need a token passing approach.*

## 7. Graceful Shutdown (Context Timeout)
**Question:** Gracefully shut down all goroutines when a timeout occurs.

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.After(2 * time.Second): // Simulate long work
		fmt.Printf("Worker %d finished\n", id)
	case <-ctx.Done():
		fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers stopped")
}
```

## 8. Stop on First Error
**Question:** Stop all goroutines when any worker returns an error.

```go
package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"errors"
	"time"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	
	// Create 3 tasks
	for i := 0; i < 3; i++ {
		i := i
		g.Go(func() error {
			if i == 1 {
				return errors.New("something went wrong")
			}
			
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(100 * time.Millisecond):
				fmt.Printf("Task %d done\n", i)
				return nil
			}
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Error encountered:", err)
	} else {
		fmt.Println("Success")
	}
}
```

## 9. Fix Deadlock
**Question:** Fix a program that deadlocks due to incorrect channel usage.
*Scenario: Sending to an unbuffered channel without a receiver ready.*

```go
// BROKEN:
// func main() {
//     ch := make(chan int)
//     ch <- 1 // Blocks forever -> Deadlock
//     fmt.Println(<-ch)
// }

// FIXED (Use goroutine):
func main() {
    ch := make(chan int)
    go func() {
        ch <- 1
    }()
    fmt.Println(<-ch)
}

// ALTERNATIVE FIX (Buffer):
func main2() {
    ch := make(chan int, 1)
    ch <- 1
    fmt.Println(<-ch)
}
```

## 10. Rate Limiter
**Question:** Implement a rate limiter allowing N operations per second.

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

	// Rate limiter ticker: 1 request every 200ms
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("Processed request", req, time.Now())
	}
}

## 11. Pipeline Pattern
**Question:** Build a pipeline where each stage runs in its own goroutine.

```go
package main

import "fmt"

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func squarer(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func main() {
	// Set up the pipeline
	// gen -> sq -> main
	for n := range squarer(generator(1, 2, 3, 4)) {
		fmt.Println(n)
	}
}
```

## 12. Select Timeout
**Question:** Read from a channel with a timeout using select.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout encountered")
	}
}
```

## 13. Goroutine Leak Fix
**Question:** Identify and fix a goroutine leak.
*Scenario: A goroutine blocked sending to a channel that no one reads.*

```go
package main

import "fmt"

// LEAKY:
// func process() {
//     ch := make(chan int) // Unbuffered
//     go func() { ch <- 1 }() // Blocks forever if no receiver
//     return // Returns immediately, leaker hangs
// }

// FIXED:
func process() <-chan int {
	ch := make(chan int)
	go func() {
		// Ensure someone reads, or use context/done channel to cancel
		defer close(ch)
		ch <- 1
	}()
	return ch
}

func main() {
	// Consume to prevent leak
	for n := range process() {
		fmt.Println(n)
	}
}
```

## 14. Channel Directions
**Question:** Rewrite a function using send-only and receive-only channels.

```go
package main

import "fmt"

// Takes receive-only channel
func consumer(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}

// Takes send-only channel
func producer(out chan<- int) {
	out <- 1
	out <- 2
	close(out)
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
```

## 15. Thread-Safe Counter
**Question:** Implement a thread-safe counter using goroutines and channels (or Mutex).

```go
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	fmt.Println("Count:", c.Value())
}
```

## 16. RWMutex (Many Readers)
**Question:** Handle a many-readers, few-writers scenario safely.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func (m *SafeMap) Read(key string) int {
	m.mu.RLock() // Multiple readers allowed
	defer m.mu.RUnlock()
	return m.data[key]
}

func (m *SafeMap) Write(key string, val int) {
	m.mu.Lock() // Exclusive lock
	defer m.mu.Unlock()
	m.data[key] = val
}

func main() {
	m := SafeMap{data: make(map[string]int)}
	m.Write("foo", 1)

	// Multiple readers
	for i := 0; i < 10; i++ {
		go fmt.Println(m.Read("foo"))
	}
	time.Sleep(time.Second)
}
```

## 17. Ordered Results (Fan-In/Fan-Out)
**Question:** Process tasks concurrently but return results in the original order.

```go
package main

import (
	"fmt"
	"sync"
)

type Result struct {
	id  int
	val int
}

func main() {
	numJobs := 5
	results := make([]Result, numJobs)
	var wg sync.WaitGroup

	for i := 0; i < numJobs; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Process concurrent
			res := id * id 
			// Store in proper index (sync via index assignment is safe if indices unique)
			results[id] = Result{id, res}
		}(i)
	}

	wg.Wait()
	// Iterate in order
	for _, r := range results {
		fmt.Printf("Job %d: %d\n", r.id, r.val)
	}
}
```

## 18. Dynamic Worker Pool
**Question:** Dynamically scale worker goroutines based on load.
*Note: This is complex; a simple approach uses a semaphore to limit max scale.*

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int)
	
	// Spawner
	go func() {
		for j := range jobs {
			j := j
			// Spin up a new worker per job, relying on runtime scheduler
			// Or check a counter/semaphore to decide if we should spawn
			go func() {
				fmt.Println("Processing", j)
				time.Sleep(100 * time.Millisecond)
			}()
		}
	}()

	for i := 0; i < 10; i++ {
		jobs <- i
	}
	time.Sleep(2 * time.Second)
}
```

## 19. Buffered Channel Semaphore
**Question:** Implement a semaphore using a buffered channel.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	sem := make(chan struct{}, 2) // Max 2 concurrent

	for i := 0; i < 5; i++ {
		sem <- struct{}{} // Acquire token
		go func(id int) {
			fmt.Println("Worker", id, "running")
			time.Sleep(time.Second)
			<-sem // Release token
		}(i)
	}
	time.Sleep(6 * time.Second)
}
```

## 20. Ping Pong
**Question:** Coordinate two goroutines to alternately print “ping” and “pong”.

```go
package main

import (
	"fmt"
	"time"
)

func player(name string, table chan int) {
	for {
		ball := <-table
		ball++
		fmt.Println(name, ball)
		time.Sleep(500 * time.Millisecond)
		table <- ball
	}
}

func main() {
	table := make(chan int)
	go player("ping", table)
	go player("pong", table)

	table <- 0 // Serve
	time.Sleep(3 * time.Second)
	<-table // Grab ball
}
```

## 21. Context Cancellation in Worker Pool
**Question:** Pass context through a worker pool and cancel on timeout.

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d cancelled\n", id)
			return
		case j, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("Worker %d processing %d\n", id, j)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	jobs := make(chan int, 100)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, &wg)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}
```

## 22. Cancel on First Error (Context)
**Question:** Cancel all goroutines when the first error occurs using context.

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	// CancelFunc is usually called by the first error-er
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cleanup

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return // Early exit
			case <-time.After(time.Duration(id) * 100 * time.Millisecond):
				if id == 1 {
					fmt.Println("Error in task", id)
					cancel() // Cancel others
				} else {
					fmt.Println("Task", id, "done")
				}
			}
		}(i)
	}
	wg.Wait()
}
```

## 23. Collect All Errors
**Question:** Collect all errors from concurrent goroutines safely.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	errCh := make(chan error, 3) // Buffered for non-blocking send
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if id%2 != 0 {
				errCh <- fmt.Errorf("error from %d", id)
			}
		}(i)
	}

	wg.Wait()
	close(errCh)

	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	fmt.Println("Errors collected:", errs)
}
```

## 24. Slow Consumer (Drop Pattern)
**Question:** Prevent slow consumers from blocking fast producers.
*Solution: Use buffered channel and select with default to drop messages.*

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	// Slow Consumer
	go func() {
		for n := range ch {
			fmt.Println("Consumed:", n)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Fast Producer
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
			fmt.Println("Sent:", i)
		default:
			fmt.Println("Dropped/Blocked:", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
```

## 25. Non-Blocking Select
**Question:** Implement non-blocking send and receive using select and default.

```go
package main

import "fmt"

func main() {
	ch := make(chan string)

	// Non-blocking result
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received")
	}

	// Non-blocking send
	select {
	case ch <- "hello":
		fmt.Println("Sent message")
	default:
		fmt.Println("Could not send (receiver not ready)")
	}
}
```

## 26. Drain Function
**Question:** Ensure all goroutines exit cleanly when input channel is closed.

```go
package main

import (
	"fmt"
	"sync"
)

func worker(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Range loop exits automatically when 'in' is closed
	for n := range in {
		fmt.Println("Processing", n)
	}
	fmt.Println("Worker stopping")
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ch, &wg)

	ch <- 1
	ch <- 2
	close(ch) // Signals worker to stop

	wg.Wait()
}
```

## 27. Safe Result Channel Closure
**Question:** Safely close a results channel used by multiple goroutines.
*Problem: Multiple senders can't close the channel. Who closes it?*
*Solution: Use a WaitGroup to track senders, close after Wait.*

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	results := make(chan int)
	var wg sync.WaitGroup

	// Multiple Producers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			results <- id
		}(i)
	}

	// Closer Goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	// Consumer
	for r := range results {
		fmt.Println("Result:", r)
	}
}
```

## 28. Concurrent Retry Logic
**Question:** Process jobs concurrently with retry logic on failure.

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		retries := 0
		for {
			err := process(j)
			if err == nil {
				fmt.Printf("Job %d processed by %d\n", j, id)
				break
			}
			retries++
			if retries > 3 {
				fmt.Printf("Job %d failed after retries\n", j)
				break
			}
			time.Sleep(10 * time.Millisecond) // Backoff
		}
	}
}

func process(j int) error {
	if rand.Float32() < 0.5 {
		return fmt.Errorf("random fail")
	}
	return nil
}

func main() {
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
}
```

## 29. Limit HTTP Requests (Semaphore)
**Question:** Limit concurrent HTTP requests using goroutines and channels.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	urls := []string{"http://a.com", "http://b.com", "http://c.com", "http://d.com"}
	maxConcurrent := 2
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			
			sem <- struct{}{} // Acquire
			fmt.Println("Fetching", u)
			time.Sleep(time.Second) // Simulate fetch
			<-sem // Release
			
		}(url)
	}
	wg.Wait()
}
```

## 30. Fan-Out / Fan-In Safe
**Question:** Fan out work to multiple goroutines and fan in results safely.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	// Source
	in := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ { in <- i }
		close(in)
	}()

	// Fan-Out (3 workers)
	outs := make([]<-chan int, 3)
	for i := 0; i < 3; i++ {
		outs[i] = worker(in)
	}

	// Fan-In
	for n := range merge(outs...) {
		fmt.Println("Result:", n)
	}
}

func worker(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch <-chan int) {
			for n := range ch { out <- n }
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
```
```
