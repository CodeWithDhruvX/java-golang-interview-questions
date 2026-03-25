package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)
	
	// Non-blocking send using select
	fmt.Println("Non-blocking operations with select:")
	
	// Try to send (channel is empty)
	select {
	case ch <- "Hello":
		fmt.Println("Sent: Hello")
	default:
		fmt.Println("Could not send: channel not ready")
	}
	
	// Try to receive (channel has data)
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("Could not receive: no data available")
	}
	
	// Try to receive again (channel is empty)
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("Could not receive: no data available")
	}
	
	// Demonstrate with slow sender
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Delayed message"
	}()
	
	fmt.Println("\nChecking for message without blocking:")
	for i := 0; i < 5; i++ {
		select {
		case msg := <-ch:
			fmt.Printf("Received after %d checks: %s\n", i+1, msg)
			return
		default:
			fmt.Printf("Check %d: no message yet\n", i+1)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you perform non-blocking channel operations using select?

**Your Response:** I use the select statement with a default case to perform non-blocking channel operations. The default case executes immediately if no other channel operations are ready.

When trying to send, if the channel is not ready to receive, the default case executes instead of blocking. Similarly, when trying to receive, if no data is available, the default case executes.

This pattern is useful for checking channel status without blocking, implementing polling loops, or handling multiple operations where you don't want to wait indefinitely.

In the example, I demonstrate both non-blocking send and receive operations. The polling loop shows how you can periodically check for data without blocking the entire goroutine, which is useful in scenarios where you need to do other work while waiting for communication.
