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

	// Producer with graceful shutdown
	go func() {
		defer wg.Done()
		defer close(jobs) // Ensure channel is closed
		
		for i := 1; i <= 10; i++ {
			job := Job{
				ID:   i,
				Data: fmt.Sprintf("Task %d", i),
			}
			
			select {
			case jobs <- job:
				fmt.Printf("Producer: sent job %d\n", job.ID)
				time.Sleep(200 * time.Millisecond)
			case <-time.After(1 * time.Second):
				fmt.Printf("Producer: timeout sending job %d, shutting down\n", job.ID)
				return
			}
		}
		fmt.Println("Producer: finished all jobs")
	}()

	// Consumer with graceful shutdown
	go func() {
		defer wg.Done()
		jobCount := 0
		
		for job := range jobs {
			jobCount++
			fmt.Printf("Consumer: processing job %d\n", job.ID)
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("Consumer: completed job %d\n", job.ID)
		}
		
		fmt.Printf("Consumer: processed %d jobs, shutting down gracefully\n", jobCount)
	}()

	wg.Wait()
	fmt.Println("System shutdown complete")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement graceful shutdown in a producer-consumer system?

**Your Response:** I implement graceful shutdown by ensuring proper channel closure and using defer statements to guarantee cleanup.

The producer uses defer close(jobs) to ensure the channel is always closed, even if it exits early. I also add a timeout mechanism using select to prevent the producer from blocking indefinitely during shutdown.

The consumer handles shutdown naturally by using the range pattern over the channel - when the channel is closed and all jobs are processed, the range loop exits cleanly.

The WaitGroup ensures the main function waits for both producer and consumer to complete their work before exiting. This prevents abrupt termination and ensures all in-flight jobs are processed.

This pattern is crucial for real-world services where you need to handle shutdown signals gracefully, ensuring no data loss and clean resource cleanup.
