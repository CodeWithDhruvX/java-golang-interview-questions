package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	bufferSize := 3
	buffered := make(chan int, bufferSize)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Filler goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= 8; i++ {
			select {
			case buffered <- i:
				fmt.Printf("Successfully sent %d (buffer: %d/%d)\n", i, len(buffered), bufferSize)
			case <-time.After(100 * time.Millisecond):
				fmt.Printf("Timeout sending %d (buffer full: %d/%d)\n", i, len(buffered), bufferSize)
			}
		}
		close(buffered)
	}()

	// Consumer goroutine
	go func() {
		defer wg.Done()
		for value := range buffered {
			fmt.Printf("Received %d (buffer: %d/%d)\n", value, len(buffered), bufferSize)
			time.Sleep(500 * time.Millisecond) // Slow consumption
		}
	}()

	wg.Wait()
	fmt.Println("Channel operations completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you safely handle sending to a potentially full buffered channel?

**Your Response:** I use a select statement with a timeout case to handle the scenario where the buffer might be full. This prevents the sender from blocking indefinitely.

When trying to send, the select statement attempts the channel send first. If the buffer is full, it falls back to the timeout case after 100ms. This allows the sender to handle the full buffer condition gracefully instead of blocking forever.

The pattern demonstrates understanding of non-blocking communication using select. The timeout provides a way to detect and handle backpressure situations where the consumer can't keep up.

This approach is safer than just blocking because it allows the sender to implement alternative strategies when the buffer is full, like retrying later, dropping messages, or alerting about system overload.
