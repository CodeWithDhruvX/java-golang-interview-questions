package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan int, 20)
	wg := sync.WaitGroup{}
	numProducers := 3

	// Start multiple producers
	for i := 0; i < numProducers; i++ {
		wg.Add(1)
		go func(producerID int) {
			defer wg.Done()
			for j := 1; j <= 4; j++ {
				job := producerID*10 + j
				fmt.Printf("Producer %d: sending job %d\n", producerID, job)
				jobs <- job
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}

	// Single consumer
	go func() {
		wg.Add(1)
		defer wg.Done()
		for job := range jobs {
			fmt.Printf("Consumer: processing job %d\n", job)
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("Consumer: completed job %d\n", job)
		}
	}()

	// Wait for producers and close channel
	go func() {
		wg.Wait()
		close(jobs)
	}()

	// Wait for consumer to finish
	time.Sleep(5 * time.Second)
	fmt.Println("All jobs processed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you handle multiple producers sending to a single consumer?

**Your Response:** I use a shared jobs channel that all producers write to and a single consumer reads from. This is a classic fan-in pattern where multiple data streams converge into one.

Each producer generates jobs with unique IDs based on their producer ID and sends them to the shared channel. The single consumer processes jobs sequentially from this channel.

The key is using a buffered channel to handle the case where producers outpace the consumer. The WaitGroup ensures all producers complete before closing the channel, which signals the consumer to finish.

This pattern is common in real systems like log aggregation, where multiple services send logs to a central processor. It demonstrates understanding of fan-in patterns and channel sharing across multiple goroutines.
