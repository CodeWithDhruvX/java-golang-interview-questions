package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	unbuffered := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Sender goroutine
	go func() {
		defer wg.Done()
		messages := []string{"Hello", "World", "Go", "Concurrency", "Rocks"}
		
		for _, msg := range messages {
			fmt.Printf("Sending: %s\n", msg)
			unbuffered <- msg // Blocks until receiver is ready
			fmt.Printf("Sent: %s\n", msg)
		}
		close(unbuffered)
	}()

	// Receiver goroutine
	go func() {
		defer wg.Done()
		for msg := range unbuffered {
			fmt.Printf("Received: %s\n", msg)
			time.Sleep(500 * time.Millisecond) // Slow processing
		}
	}()

	wg.Wait()
	fmt.Println("Communication completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** Can you explain the blocking behavior of unbuffered channels?

**Your Response:** Unbuffered channels provide synchronous communication - the sender blocks until a receiver is ready, and the receiver blocks until data is available.

In this example, when the sender tries to send "Hello", it blocks until the receiver is ready to receive. Once received, the sender unblocks and can send the next message. This creates a handshaking mechanism.

The receiver processes each message with a delay, demonstrating how the sender waits during this processing time. This blocking behavior ensures that no data is lost and both goroutines synchronize their execution.

Unbuffered channels are perfect for synchronization scenarios where you want guaranteed delivery and immediate coordination between goroutines. They're fundamental to Go's communication philosophy of "don't communicate by sharing memory, share memory by communicating."
