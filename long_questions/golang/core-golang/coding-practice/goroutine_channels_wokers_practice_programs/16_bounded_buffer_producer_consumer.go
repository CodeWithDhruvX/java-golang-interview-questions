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
	bufferSize := 3
	jobs := make(chan Job, bufferSize)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Producer with bounded buffer
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			job := Job{
				ID:   i,
				Data: fmt.Sprintf("Task %d", i),
			}
			
			fmt.Printf("Producer: trying to send job %d (buffer: %d/%d)\n", 
				i, len(jobs), bufferSize)
			jobs <- job // Blocks if buffer is full
			fmt.Printf("Producer: sent job %d (buffer: %d/%d)\n", 
				i, len(jobs), bufferSize)
			
			time.Sleep(150 * time.Millisecond) // Production rate
		}
		close(jobs)
		fmt.Println("Producer: finished")
	}()

	// Consumer
	go func() {
		defer wg.Done()
		for job := range jobs {
			fmt.Printf("Consumer: processing job %d (buffer: %d/%d)\n", 
				job.ID, len(jobs), bufferSize)
			time.Sleep(400 * time.Millisecond) // Slower consumption
			fmt.Printf("Consumer: completed job %d\n", job.ID)
		}
		fmt.Println("Consumer: finished")
	}()

	wg.Wait()
	fmt.Println("Bounded buffer producer-consumer completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does the bounded buffer prevent the producer from overwhelming the system?

**Your Response:** The bounded buffer acts as a natural throttle through backpressure. With buffer size 3, the producer can only send 3 jobs ahead of the consumer before blocking.

When the buffer fills up, the producer automatically blocks and waits for the consumer to process jobs and free up space. This creates self-regulating flow control - the producer automatically slows down when the consumer can't keep up.

The key insight is that the buffer size determines how much the producer can work ahead. Too small and the producer blocks frequently; too large and we waste memory and might overwhelm the consumer.

This pattern is fundamental to real-world systems like message queues and network protocols, where bounded buffers prevent producers from overwhelming consumers and provide natural backpressure mechanisms for system stability.
