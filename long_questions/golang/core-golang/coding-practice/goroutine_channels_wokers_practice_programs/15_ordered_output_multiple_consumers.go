package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data string
}

type Result struct {
	ID       int
	Data     string
	Consumer int
}

func main() {
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)
	wg := sync.WaitGroup{}
	numConsumers := 2

	// Producer
	go func() {
		for i := 1; i <= 10; i++ {
			job := Job{
				ID:   i,
				Data: fmt.Sprintf("Task %d", i),
			}
			jobs <- job
		}
		close(jobs)
	}()

	// Multiple consumers
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go func(consumerID int) {
			defer wg.Done()
			for job := range jobs {
				// Process job
				time.Sleep(time.Duration(job.ID%3+1) * 200 * time.Millisecond)
				result := Result{
					ID:       job.ID,
					Data:     fmt.Sprintf("Processed %s", job.Data),
					Consumer: consumerID,
				}
				results <- result
			}
		}(i)
	}

	// Wait for consumers to finish
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
		fmt.Printf("Job %d: %s (by Consumer %d)\n", 
			result.ID, result.Data, result.Consumer)
	}
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you maintain output order with multiple consumers processing at different speeds?

**Your Response:** I separate concurrent processing from ordered output by having consumers send results with their original job IDs to a results channel, then reassemble them in order.

The consumers process jobs concurrently at different speeds, but each result includes the original job ID. I collect all results in a slice indexed by job ID, which automatically places them in the correct order regardless of processing time.

The key insight is that we don't need to process in order - we just need to output in order. This allows us to benefit from parallel processing while maintaining the appearance of sequential execution.

This pattern is extremely useful in real-world scenarios like parallel API calls where you want the performance of concurrent processing but need ordered presentation. It demonstrates understanding of concurrent processing with result ordering, which is crucial for many backend systems.
