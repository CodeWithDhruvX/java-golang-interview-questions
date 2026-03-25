package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tokenChan := make(chan bool, 1)
	wg := sync.WaitGroup{}
	numWorkers := 3
	wg.Add(numWorkers)

	// Initialize with token
	tokenChan <- true

	// Create multiple goroutines
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				// Wait for token
				token := <-tokenChan
				
				// Critical section - only one goroutine at a time
				fmt.Printf("Worker %d has token, executing task %d\n", workerID, j+1)
				time.Sleep(100 * time.Millisecond) // Simulate work
				
				// Pass token to next worker
				tokenChan <- token
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does the token passing pattern ensure only one goroutine executes at a time?

**Your Response:** The token passing pattern uses a single boolean channel as a mutual exclusion mechanism. The token represents permission to enter the critical section.

Each goroutine must acquire the token before executing its critical code. It reads from the channel (blocking until token is available), performs its work, then writes the token back to the channel for the next goroutine.

This ensures that only one goroutine can be in its critical section at any given time because there's only one token available. All other goroutines block waiting for the token.

It's essentially implementing a mutex using channels. The pattern is useful when you need to coordinate access to shared resources or ensure sequential execution of critical sections across multiple goroutines.
