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

func main() {
	jobs := make(chan Job, 10)
	wg := sync.WaitGroup{}
	wg.Add(3) // 1 producer + 2 consumers

	// Producer
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			job := Job{
				ID:   i,
				Data: fmt.Sprintf("Task %d", i),
			}
			fmt.Printf("Producer: creating job %d\n", job.ID)
			jobs <- job
			time.Sleep(100 * time.Millisecond)
		}
		close(jobs)
		fmt.Println("Producer: finished, closed jobs channel")
	}()

	// Consumer 1
	go func() {
		defer wg.Done()
		for job := range jobs {
			fmt.Printf("Consumer 1: processing job %d\n", job.ID)
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("Consumer 1: completed job %d\n", job.ID)
		}
		fmt.Println("Consumer 1: finished")
	}()

	// Consumer 2
	go func() {
		defer wg.Done()
		for job := range jobs {
			fmt.Printf("Consumer 2: processing job %d\n", job.ID)
			time.Sleep(250 * time.Millisecond)
			fmt.Printf("Consumer 2: completed job %d\n", job.ID)
		}
		fmt.Println("Consumer 2: finished")
	}()

	wg.Wait()
	fmt.Println("All jobs processed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you ensure all jobs are processed when you have one producer and two consumers?

**Your Response:** I use a single jobs channel that both consumers read from. Go's channel semantics ensure that each job is received by exactly one consumer - there's no duplication or loss of jobs.

When a job is sent to the channel, either Consumer 1 or Consumer 2 will receive it, depending on which is ready first. The channel acts as a work queue with multiple workers pulling from it.

The key is that both consumers use the range pattern on the same channel. When the producer closes the channel, both consumers will finish processing any remaining jobs and then exit naturally.

The WaitGroup tracks all three goroutines (1 producer + 2 consumers). This pattern efficiently distributes work across multiple workers and is commonly used in real systems for parallel processing, like web servers handling requests or background job processors.
