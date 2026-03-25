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
		
		// Process job with variable time
		time.Sleep(time.Duration(job.ID%4+1) * 200 * time.Millisecond)
		output := fmt.Sprintf("Processed %s", job.Data)
		
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

	// Collect and order results
	orderedResults := make([]Result, numJobs)
	for result := range results {
		orderedResults[result.JobID-1] = result
	}

	// Print in order
	fmt.Println("Ordered Results:")
	for _, result := range orderedResults {
		fmt.Printf("Job %d: %s by Worker %d (took %v)\n", 
			result.JobID, result.Output, result.WorkerID, result.Duration)
	}
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you maintain ordered output in a worker pool where jobs complete at different times?

**Your Response:** I maintain ordered output by separating concurrent processing from result ordering. Workers process jobs concurrently at different speeds, but each result includes the original job ID.

I collect all results in a slice indexed by job ID, which automatically places them in the correct order regardless of completion time. This allows us to benefit from parallel processing while presenting results sequentially.

The key insight is that we don't need jobs to complete in order - we just need to output them in order. Workers can process jobs as fast as they can, and the ordering happens at the collection stage.

This pattern is extremely valuable in real-world scenarios like parallel API calls or batch processing, where you want the performance benefits of concurrent processing but need ordered presentation for user experience or downstream processing requirements.
