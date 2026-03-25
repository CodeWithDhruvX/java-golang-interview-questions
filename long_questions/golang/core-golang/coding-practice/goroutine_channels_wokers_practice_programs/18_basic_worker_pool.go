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

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d: processing job %d\n", id, job.ID)
		time.Sleep(time.Duration(job.ID%3+1) * 200 * time.Millisecond)
		fmt.Printf("Worker %d: completed job %d\n", id, job.ID)
	}
}

func main() {
	numWorkers := 3
	numJobs := 10
	jobs := make(chan Job, numJobs)
	wg := sync.WaitGroup{}

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for i := 1; i <= numJobs; i++ {
		job := Job{
			ID:   i,
			Data: fmt.Sprintf("Task %d", i),
		}
		jobs <- job
		fmt.Printf("Main: dispatched job %d\n", i)
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All jobs completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** Can you explain how the worker pool pattern distributes work among multiple workers?

**Your Response:** The worker pool pattern uses a channel as a task queue and multiple worker goroutines that pull tasks from this queue. I create 3 worker goroutines that all read from the same jobs channel.

When jobs are sent to the channel, any available worker can receive them. Go's channel semantics handle the distribution automatically - if multiple workers are waiting, one is chosen randomly to receive the job.

Each worker uses the range pattern to continuously process jobs until the channel is closed. This creates an efficient work distribution where idle workers automatically pick up new jobs.

The main function dispatches all jobs, then closes the channel to signal no more work is coming. The WaitGroup ensures all workers complete their current jobs before the program exits.

This pattern is extremely efficient for concurrent processing and is used extensively in real-world systems like web servers, database connection pools, and background job processors.
