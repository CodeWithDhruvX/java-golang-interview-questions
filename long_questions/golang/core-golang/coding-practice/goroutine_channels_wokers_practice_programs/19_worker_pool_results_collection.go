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
	JobID    int
	WorkerID int
	Output   string
	Duration time.Duration
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		start := time.Now()
		
		// Process the job
		time.Sleep(time.Duration(job.ID%3+1) * 200 * time.Millisecond)
		output := fmt.Sprintf("Processed %s by worker %d", job.Data, id)
		
		result := Result{
			JobID:    job.ID,
			WorkerID: id,
			Output:   output,
			Duration: time.Since(start),
		}
		
		results <- result
	}
}

func main() {
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
	for result := range results {
		fmt.Printf("Job %d: %s (took %v)\n", 
			result.JobID, result.Output, result.Duration)
	}
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does your worker pool collect and return results from processed jobs?

**Your Response:** I use a separate results channel that workers send their processed results to. This creates a clean separation between job distribution and result collection.

Each worker processes a job and creates a Result struct containing the job ID, worker ID, processed output, and processing duration. The worker sends this result to the results channel.

The main function collects all results from the results channel. I use a separate goroutine to wait for all workers to complete and then close the results channel, which signals that no more results will come.

This pattern is powerful because it decouples job processing from result handling. The results channel can be buffered to handle cases where result processing is slower than job processing, and multiple goroutines could consume results if needed.

This approach is commonly used in real systems where you need to track which worker processed which job and collect performance metrics.
