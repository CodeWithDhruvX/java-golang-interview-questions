package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	ID      int
	Value   int
	ProcessTime time.Duration
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan Result, 10)
	wg := sync.WaitGroup{}

	// Start workers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				start := time.Now()
				// Simulate processing with variable time
				time.Sleep(time.Duration(job%5+1) * 100 * time.Millisecond)
				result := Result{
					ID:    job,
					Value: job * job,
					ProcessTime: time.Since(start),
				}
				results <- result
			}
		}(i)
	}

	// Send jobs
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and order results
	orderedResults := make([]Result, 10)
	for result := range results {
		orderedResults[result.ID-1] = result
	}

	// Print in order
	for _, result := range orderedResults {
		fmt.Printf("Job %d: Result=%d, Time=%v\n", 
			result.ID, result.Value, result.ProcessTime)
	}
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you process jobs concurrently but maintain ordered output?

**Your Response:** I use a producer-consumer pattern with multiple workers processing jobs concurrently, but store results with their original IDs to reassemble them in order.

The key insight is separating processing time from output order. Workers process jobs as fast as they can, but each result includes its original job ID. I collect all results in a slice indexed by job ID, which automatically places them in the correct order.

I use a buffered jobs channel for work distribution and a results channel to collect processed data. The WaitGroup ensures all workers complete before closing the results channel.

This pattern is useful in real-world scenarios like parallel API calls where you need concurrent processing but ordered presentation. It demonstrates understanding of concurrent processing with result ordering, which is common in backend systems.
