package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d: starting\n", id)
	time.Sleep(time.Duration(id) * 200 * time.Millisecond)
	fmt.Printf("Worker %d: completed\n", id)
}

func main() {
	numWorkers := 5
	wg := sync.WaitGroup{}

	fmt.Printf("Starting %d goroutines with WaitGroup\n", numWorkers)
	
	// Add all workers to WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all workers to complete
	fmt.Println("Main: waiting for all workers to complete")
	wg.Wait()
	fmt.Println("Main: all workers completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does WaitGroup help you wait for multiple goroutines to complete?

**Your Response:** WaitGroup is a synchronization primitive that allows you to wait for a collection of goroutines to finish. It works like a counter that tracks how many goroutines are still running.

I use wg.Add(1) before starting each worker to increment the counter. Each worker calls wg.Done() when it finishes to decrement the counter. The main function calls wg.Wait() which blocks until the counter reaches zero.

The key insight is that WaitGroup provides a clean way to coordinate completion without using channels or complex signaling. It's specifically designed for the common pattern of "start multiple goroutines and wait for them all to finish."

This pattern is fundamental to Go concurrent programming and is used extensively in real systems for coordinating parallel work, like processing multiple files, making parallel API calls, or running background tasks.
