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
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Worker %d: completed job %d\n", id, job.ID)
	}
}

func main() {
	maxWorkers := 3
	numJobs := 10
	jobs := make(chan Job, maxWorkers) // Small buffer to limit concurrency
	wg := sync.WaitGroup{}

	// Start limited workers
	for i := 1; i <= maxWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs with controlled rate
	go func() {
		for i := 1; i <= numJobs; i++ {
			job := Job{
				ID:   i,
				Data: fmt.Sprintf("Task %d", i),
			}
			jobs <- job
			fmt.Printf("Main: dispatched job %d\n", i)
			time.Sleep(200 * time.Millisecond) // Control dispatch rate
		}
		close(jobs)
	}()

	wg.Wait()
	fmt.Println("All jobs completed with max 3 concurrent goroutines")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you limit the maximum number of goroutines running concurrently?

**Your Response:** I limit concurrency by creating exactly 3 worker goroutines and using a small buffered jobs channel. The key insight is that concurrency is limited by the number of active goroutines, not the channel capacity.

Only 3 workers exist, so at most 3 jobs can be processed simultaneously regardless of how many jobs are in the queue. The small buffer (size 3) allows some queuing but prevents unlimited buildup.

I control the job dispatch rate with a sleep in the sender, which prevents overwhelming the system. The workers continuously pull from the channel, so as soon as one finishes, it can take the next job.

This pattern is crucial for resource management - it prevents system overload by limiting concurrent resource usage. It's commonly used in real systems for database connection pools, API rate limiting, and preventing memory exhaustion from too many concurrent operations.
