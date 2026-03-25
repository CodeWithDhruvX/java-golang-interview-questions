package main

import (
	"fmt"
	"sync"
	"time"
)

func slowSender(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second) // Slow sender
	ch <- "Slow message"
	fmt.Println("Slow sender: message sent")
}

func main() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Start slow sender
	go slowSender(ch, &wg)

	// Read with timeout using select
	fmt.Println("Waiting for message with 2 second timeout...")
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: No message received within 2 seconds")
	}

	// Try again without timeout
	fmt.Println("\nWaiting for message without timeout...")
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	}

	wg.Wait()
	fmt.Println("Program completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement a timeout using the select statement?

**Your Response:** I implement a timeout using the select statement with a time.After case. The time.After function returns a channel that sends a value after the specified duration.

In the select statement, I have two cases: one for receiving from the channel and one for the timeout. Whichever case completes first executes. If the message doesn't arrive within 2 seconds, the timeout case triggers.

This pattern is extremely useful for preventing goroutines from blocking indefinitely when waiting for responses that might never come. It's commonly used in real systems for API calls, database queries, or any operation that might hang.

The key insight is that select allows you to wait on multiple conditions simultaneously, making it perfect for implementing timeouts, cancellations, and other time-based operations in concurrent systems.
