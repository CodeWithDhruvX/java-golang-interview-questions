package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan int, 5) // Buffered channel
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Producer - sends jobs quickly
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			fmt.Printf("Producing job %d\n", i)
			jobs <- i // Non-blocking until buffer full
			fmt.Printf("Produced job %d (buffer: %d/5)\n", i, len(jobs))
			time.Sleep(100 * time.Millisecond) // Fast producer
		}
		close(jobs)
	}()

	// Consumer - processes slowly
	go func() {
		defer wg.Done()
		for job := range jobs {
			fmt.Printf("Consuming job %d (buffer: %d/5)\n", job, len(jobs))
			time.Sleep(500 * time.Millisecond) // Slow processing
			fmt.Printf("Completed job %d\n", job)
		}
	}()

	wg.Wait()
	fmt.Println("Producer-consumer completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does the buffered channel handle the fast producer and slow consumer scenario?

**Your Response:** The buffered channel acts as a cushion between the fast producer and slow consumer. With buffer size 5, the producer can send up to 5 jobs without blocking, buying time for the consumer to catch up.

Initially, the producer fills the buffer quickly. Once the buffer is full, the producer blocks and waits for the consumer to process jobs and free up buffer space. This creates natural backpressure - the producer automatically slows down when the consumer can't keep up.

The buffer size is crucial here. Too small and the producer blocks frequently; too large and we waste memory and might hide performance issues. Size 5 provides a good balance for this scenario.

This pattern is common in real systems like message queues, where producers and consumers operate at different rates. The buffer smooths out the rate differences while maintaining flow control through backpressure.
