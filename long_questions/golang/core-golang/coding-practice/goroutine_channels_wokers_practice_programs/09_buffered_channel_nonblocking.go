package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	buffered := make(chan string, 3)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Sender goroutine
	go func() {
		defer wg.Done()
		messages := []string{"Msg1", "Msg2", "Msg3", "Msg4", "Msg5"}
		
		for i, msg := range messages {
			fmt.Printf("Sending %d: %s (buffer len: %d)\n", i+1, msg, len(buffered))
			buffered <- msg // Non-blocking until buffer is full
			fmt.Printf("Sent %d: %s (buffer len: %d)\n", i+1, msg, len(buffered))
		}
		close(buffered)
	}()

	// Receiver goroutine
	go func() {
		defer wg.Done()
		for msg := range buffered {
			fmt.Printf("Received: %s (buffer len: %d)\n", msg, len(buffered))
			time.Sleep(800 * time.Millisecond) // Slow processing
		}
	}()

	wg.Wait()
	fmt.Println("Buffered communication completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do buffered channels differ from unbuffered channels in terms of blocking behavior?

**Your Response:** Buffered channels provide asynchronous communication with a fixed capacity. The sender only blocks when the buffer is full, not when sending each message.

In this example with buffer size 3, the sender can send the first 3 messages immediately without blocking. But when trying to send the 4th message, it blocks until the receiver processes one message and frees buffer space.

The key difference is that buffered channels decouple the sender and receiver - they can work at different speeds. The buffer acts as a queue allowing the sender to continue working while the receiver processes messages asynchronously.

This is useful for producer-consumer scenarios where production and consumption rates differ. The buffer size affects performance - too small causes frequent blocking, too large uses more memory and can hide backpressure issues.
