package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(id int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: stopping due to %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	fmt.Println("Starting workers with cancellation support")

	// Start workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, ctx, &wg)
	}

	// Let workers run for a while
	time.Sleep(2 * time.Second)
	
	fmt.Println("\nCancelling all workers...")
	cancel()

	// Wait for workers to stop
	wg.Wait()
	fmt.Println("All workers stopped")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you cancel a goroutine using context?

**Your Response:** I use context.WithCancel to create a cancellable context that can signal goroutines to stop gracefully. The context carries a cancellation signal that goroutines can monitor.

Each worker uses a select statement to check for either work to do or the cancellation signal. When cancel() is called, the context's Done() channel is closed, which unblocks all waiting goroutines.

The key insight is that context provides a standardized way to propagate cancellation signals across goroutine boundaries. Instead of creating custom cancellation channels, context provides a consistent pattern that works throughout the Go ecosystem.

When a goroutine receives the cancellation signal, it can clean up resources and exit gracefully. This is much better than forcing goroutines to terminate, as it ensures proper cleanup and prevents resource leaks.

This pattern is fundamental to real-world Go applications for handling shutdowns, timeouts, and request cancellation in distributed systems.
