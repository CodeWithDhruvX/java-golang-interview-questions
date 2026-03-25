package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func slowTask(ctx context.Context, taskID int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	start := time.Now()
	fmt.Printf("Task %d: starting\n", taskID)
	
	// Simulate work that might take too long
	select {
	case <-time.After(3 * time.Second):
		// Task completed
		fmt.Printf("Task %d: completed in %v\n", taskID, time.Since(start))
	case <-ctx.Done():
		// Task timed out
		fmt.Printf("Task %d: timed out after %v due to %v\n", 
			taskID, time.Since(start), ctx.Err())
	}
}

func main() {
	// Create context with 2 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Important to prevent resource leak
	
	wg := sync.WaitGroup{}

	fmt.Println("Starting tasks with 2 second timeout")

	// Start multiple tasks
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go slowTask(ctx, i, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks handled")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement a timeout for a task using context?

**Your Response:** I use context.WithTimeout to create a context that automatically cancels after a specified duration. The context carries both a timeout signal and cancellation reason.

Each task uses a select statement to race between the actual work (simulated with time.After) and the context's Done() channel. If the work completes before the timeout, the task succeeds. If the timeout occurs first, the context is cancelled and the task can handle the timeout gracefully.

The key insight is that context.WithTimeout handles the timeout logic automatically - I don't need to create timers or manage cancellation manually. The context automatically cancels after 2 seconds, and all tasks listening to it receive the signal simultaneously.

I use defer cancel() to ensure the context resources are cleaned up even if the main function exits early. This pattern is essential for preventing resource leaks and is widely used in real systems for API calls, database operations, and any task that might hang.
