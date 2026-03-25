package main

import (
	"fmt"
	"time"
)

func createDeadlock() {
	fmt.Println("=== Creating Deadlock ===")
	
	// Deadlock example 1: Unbuffered channel send without receiver
	fmt.Println("1. Unbuffered channel deadlock:")
	ch1 := make(chan int)
	// ch1 <- 42 // This would deadlock - no receiver
	
	// Fix: Use goroutine receiver
	go func() {
		value := <-ch1
		fmt.Printf("Received: %d\n", value)
	}()
	ch1 <- 42
	time.Sleep(100 * time.Millisecond)
	
	// Deadlock example 2: Circular wait
	fmt.Println("\n2. Circular wait deadlock:")
	ch2 := make(chan int)
	ch3 := make(chan int)
	
	// This would deadlock:
	// go func() {
	//     ch2 <- <-ch3
	// }()
	// go func() {
	//     ch3 <- <-ch2
	// }()
	
	// Fix: Use select with timeout or reorder operations
	go func() {
		select {
		case ch2 <- <-ch3:
		case <-time.After(time.Second):
			fmt.Println("Timeout in circular wait")
		}
	}()
	go func() {
		ch3 <- 100
	}()
	
	time.Sleep(100 * time.Millisecond)
	
	// Deadlock example 3: Waiting on own channel
	fmt.Println("\n3. Self-dependency deadlock:")
	
	// This would deadlock:
	// ch4 := make(chan int)
	// ch4 <- <-ch4
	
	// Fix: Use separate channels
	ch4 := make(chan int)
	ch5 := make(chan int)
	go func() {
		ch5 <- <-ch4
	}()
	go func() {
		ch4 <- 200
	}()
	
	value := <-ch5
	fmt.Printf("Fixed circular dependency: %d\n", value)
}

func main() {
	createDeadlock()
	fmt.Println("\nAll deadlock examples demonstrated and fixed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** Can you create and fix a deadlock scenario?

**Your Response:** I demonstrate three common deadlock patterns and their fixes. Deadlocks occur when goroutines wait indefinitely for resources that will never become available.

First, unbuffered channel deadlock - sending to an unbuffered channel without a receiver blocks forever. The fix is ensuring a receiver is ready before sending, typically by starting the receiver goroutine first.

Second, circular wait deadlock - two goroutines waiting for each other. The fix is using select with timeout or reordering operations to break the circular dependency.

Third, self-dependency deadlock - a goroutine waiting on itself. The fix is using separate channels to avoid circular dependencies.

The key insight is that deadlocks in Go usually come from channel operations without corresponding receivers/senders, or circular waiting patterns. The fixes involve ensuring proper channel setup, using timeouts, or restructuring the communication flow.

Understanding these patterns is crucial for writing robust concurrent Go programs that don't hang unexpectedly.
