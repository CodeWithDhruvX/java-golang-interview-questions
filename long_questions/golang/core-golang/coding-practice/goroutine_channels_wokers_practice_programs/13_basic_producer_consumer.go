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
	jobs := make(chan Job, 5)
	wg := sync.WaitGroup{}
	wg.Add(2)

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
			fmt.Printf("Producer: sent job %d\n", job.ID)
			time.Sleep(200 * time.Millisecond)
		}
		close(jobs)
	}()

	// Consumer
	go func() {
		defer wg.Done()
		for job := range jobs {
			fmt.Printf("Consumer: processing job %d - %s\n", job.ID, job.Data)
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("Consumer: completed job %d\n", job.ID)
		}
	}()

	wg.Wait()
	fmt.Println("Producer-Consumer pattern completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** Can you explain the basic producer-consumer pattern implementation?

**Your Response:** I implemented the classic producer-consumer pattern using a channel as the communication medium. The producer creates jobs and sends them through the channel, while the consumer receives and processes them.

The channel acts as a queue that decouples the producer and consumer - they can work at different speeds. I used a buffered channel with capacity 5 to allow the producer to work ahead slightly without blocking immediately.

The producer creates Job structs with ID and data, sends them to the channel, then closes the channel when done. The consumer processes jobs from the channel until it's closed, using the range pattern which automatically handles channel closure.

The WaitGroup ensures both goroutines complete before the main function exits. This pattern is fundamental to concurrent programming and is used extensively in real-world systems for task distribution.
