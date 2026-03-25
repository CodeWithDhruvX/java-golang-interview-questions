package main

import (
	"fmt"
	"runtime"
	"time"
)

func writeToClosedChannel() {
	fmt.Println("=== Writing to Closed Channel (Panic Handling) ===")
	
	// Demonstrate panic when writing to closed channel
	fmt.Println("1. Direct write to closed channel (will panic):")
	ch := make(chan string, 2)
	ch <- "Hello"
	close(ch)
	
	// This would panic:
	// ch <- "World" // panic: send on closed channel
	
	fmt.Println("Channel closed, attempting safe write patterns...")
	
	// Safe pattern 1: Use defer and recover
	fmt.Println("\n2. Safe write with panic recovery:")
	safeWriteWithRecover(ch)
	
	// Safe pattern 2: Use sync.Once to ensure single close
	fmt.Println("\n3. Using sync.Once to prevent double close:")
	safeWriteWithOnce()
	
	// Safe pattern 3: Channel ownership pattern
	fmt.Println("\n4. Channel ownership pattern:")
	channelOwnershipPattern()
}

func safeWriteWithRecover(ch chan string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	
	// Attempt to write (will panic but be recovered)
	ch <- "This will panic"
	fmt.Println("This line won't be reached")
}

func safeWriteWithOnce() {
	ch := make(chan string, 2)
	var once sync.Once
	
	// Writer function
	writer := func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("Message %d", i)
			select {
			case ch <- msg:
				fmt.Printf("Sent: %s\n", msg)
			default:
				fmt.Printf("Channel full or closed, skipping: %s\n", msg)
			}
		}
	}
	
	// Closer function
	closer := func() {
		once.Do(func() {
			fmt.Println("Closing channel (only once)")
			close(ch)
		})
	}
	
	go writer()
	time.Sleep(100 * time.Millisecond)
	closer()
	time.Sleep(100 * time.Millisecond)
	closer() // Won't close again
}

func channelOwnershipPattern() {
	// Producer owns and closes the channel
	producer := func() chan string {
		ch := make(chan string, 3)
		
		go func() {
			defer close(ch) // Producer closes when done
			for i := 0; i < 3; i++ {
				ch <- fmt.Sprintf("Item %d", i)
				time.Sleep(100 * time.Millisecond)
			}
		}()
		
		return ch // Consumer receives ownership
	}
	
	// Consumer only reads
	consumer := func(ch <-chan string) {
		for item := range ch {
			fmt.Printf("Consumed: %s\n", item)
		}
		fmt.Println("Consumer finished")
	}
	
	ch := producer()
	consumer(ch)
}

func main() {
	writeToClosedChannel()
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you handle the panic that occurs when writing to a closed channel?

**Your Response:** Writing to a closed channel causes a panic in Go, but there are several patterns to handle this safely.

First, I can use defer and recover to catch the panic if writing to a potentially closed channel. This allows graceful handling but isn't ideal for normal flow control.

Second, I use sync.Once to ensure a channel is only closed once, preventing multiple close operations that could cause issues.

Third, and most importantly, I use the channel ownership pattern where only the producer goroutine owns and closes the channel. Consumers receive a read-only view (<-chan) and never attempt to write or close.

The key insight is that channel ownership should be clear - one goroutine owns the channel (writes to it and closes it), while others only receive from it. This prevents accidental writes to closed channels.

For cases where you must write to potentially closed channels, use select with a default case for non-blocking attempts, or implement proper shutdown signaling using context or separate done channels.
