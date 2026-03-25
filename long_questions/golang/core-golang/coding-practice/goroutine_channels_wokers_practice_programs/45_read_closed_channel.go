package main

import (
	"fmt"
	"time"
)

func readFromClosedChannel() {
	fmt.Println("=== Reading from Closed Channel ===")
	
	ch := make(chan string, 3)
	
	// Send some data
	ch <- "Hello"
	ch <- "World"
	ch <- "Go"
	
	// Close the channel
	close(ch)
	
	// Reading from closed channel
	fmt.Println("Reading from closed channel:")
	
	// Method 1: Using range (recommended)
	fmt.Println("Method 1: Using range")
	for value := range ch {
		fmt.Printf("Received: %s\n", value)
	}
	
	// Method 2: Using comma-ok pattern
	fmt.Println("\nMethod 2: Using comma-ok pattern")
	ch2 := make(chan string, 2)
	ch2 <- "Test1"
	ch2 <- "Test2"
	close(ch2)
	
	for {
		value, ok := <-ch2
		if !ok {
			fmt.Println("Channel closed, stopping")
			break
		}
		fmt.Printf("Received: %s\n", value)
	}
	
	// Method 3: Reading after close returns zero value
	fmt.Println("\nMethod 3: Zero value behavior")
	var value string
	value, ok := <-ch2
	fmt.Printf("Value: '%s', OK: %v\n", value, ok) // "", false
	
	// Demonstrate with different types
	fmt.Println("\nDifferent channel types:")
	intCh := make(chan int, 1)
	intCh <- 42
	close(intCh)
	
	intVal, intOk := <-intCh
	fmt.Printf("Int channel - Value: %d, OK: %v\n", intVal, intOk)
	
	intVal2, intOk2 := <-intCh
	fmt.Printf("Int channel after empty - Value: %d, OK: %v\n", intVal2, intOk2) // 0, false
}

func main() {
	readFromClosedChannel()
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does Go handle reading from a closed channel?

**Your Response:** Go provides safe mechanisms for reading from closed channels. When a channel is closed, you can still read any remaining values, but subsequent reads return the zero value for the channel type.

I demonstrate three approaches: First, using range which automatically stops when the channel is closed after reading all remaining values. This is the recommended approach.

Second, using the comma-ok pattern where the second boolean indicates whether the value came from the channel (true) or is just the zero value because the channel is closed (false).

Third, showing that reading from an empty closed channel returns the zero value - empty string for strings, 0 for integers, nil for pointers.

The key insight is that closing a channel doesn't lose data - it signals no more values will come. This allows graceful shutdown patterns where producers close channels and consumers finish processing remaining items.

This behavior is fundamental to Go's communication patterns and enables clean pipeline shutdowns and proper resource cleanup.
