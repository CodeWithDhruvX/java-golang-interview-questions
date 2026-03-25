package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID      int
	Data    string
	Retries int
}

type Result struct {
	JobID    int
	Success  bool
	Error    string
	Attempts int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		var err error
		attempts := 0
		
		for attempts <= 3 {
			attempts++
			fmt.Printf("Worker %d: attempting job %d (attempt %d)\n", 
				id, job.ID, attempts)
			
			// Simulate work with random failure
			time.Sleep(200 * time.Millisecond)
			if rand.Float32() > 0.3 { // 70% success rate
				// Success
				result := Result{
					JobID:    job.ID,
					Success:  true,
					Attempts: attempts,
				}
				results <- result
				fmt.Printf("Worker %d: job %d succeeded on attempt %d\n", 
					id, job.ID, attempts)
				break
			} else {
				// Failure
				if attempts >= 3 {
					result := Result{
						JobID:    job.ID,
						Success:  false,
						Error:    fmt.Sprintf("Failed after %d attempts", attempts),
						Attempts: attempts,
					}
					results <- result
					fmt.Printf("Worker %d: job %d failed after %d attempts\n", 
						id, job.ID, attempts)
					break
				}
				fmt.Printf("Worker %d: job %d failed, retrying...\n", id, job.ID)
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numWorkers := 3
	numJobs := 10
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	wg := sync.WaitGroup{}

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	for i := 1; i <= numJobs; i++ {
		job := Job{
			ID:   i,
			Data: fmt.Sprintf("Task %d", i),
		}
		jobs <- job
	}
	close(jobs)

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	fmt.Println("Results:")
	successCount := 0
	failureCount := 0
	for result := range results {
		if result.Success {
			fmt.Printf("✓ Job %d: SUCCESS (attempt %d)\n", 
				result.JobID, result.Attempts)
			successCount++
		} else {
			fmt.Printf("✗ Job %d: FAILED - %s\n", 
				result.JobID, result.Error)
			failureCount++
		}
	}
	
	fmt.Printf("\nSummary: %d succeeded, %d failed\n", successCount, failureCount)
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement retry logic in a worker pool with a maximum of 3 retries per job?

**Your Response:** I implement retry logic within each worker using a for loop that attempts the job up to 3 times. Each attempt simulates work with a 70% success rate and 30% failure rate.

If a job succeeds, the worker sends a success result to the results channel. If it fails and hasn't reached the retry limit, the worker retries immediately. After 3 failed attempts, the worker sends a failure result.

The key insight is that retry logic is handled within the worker, not by resending jobs to the channel. This ensures the same worker handles all retries for consistency and prevents other workers from picking up partially-completed jobs.

The Result struct tracks success status, error messages, and total attempts. This pattern is crucial for real-world systems where external operations might fail temporarily but succeed on retry, like API calls or database operations.
