package main

import (
	"fmt"
	"time"
)

func nilChannelBehavior() {
	fmt.Println("=== Nil Channel Behavior ===")
	
	// Example 1: Send to nil channel blocks forever
	fmt.Println("1. Send to nil channel:")
	var ch1 chan int // nil channel
	
	go func() {
		fmt.Println("Goroutine: attempting to send to nil channel")
		// ch1 <- 42 // This would block forever
		fmt.Println("This will never print")
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main: goroutine is blocked on nil channel send")
	
	// Example 2: Receive from nil channel blocks forever
	fmt.Println("\n2. Receive from nil channel:")
	var ch2 chan string
	
	go func() {
		fmt.Println("Goroutine: attempting to receive from nil channel")
		// value := <-ch2 // This would block forever
		fmt.Printf("This will never print: %s\n", value)
	}()
	
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main: goroutine is blocked on nil channel receive")
	
	// Example 3: Nil channels in select
	fmt.Println("\n3. Nil channels in select statements:")
	var ch3 chan int
	var ch4 chan int
	
	// Initialize ch4
	ch4 = make(chan int, 1)
	ch4 <- 100
	
	go func() {
		for i := 0; i < 3; i++ {
			select {
			case <-ch3: // nil channel, never ready
				fmt.Println("Received from ch3 (impossible)")
			case <-ch4:
				fmt.Printf("Received from ch4: %d\n", <-ch4)
			default:
				fmt.Printf("Iteration %d: no channel ready\n", i+1)
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	time.Sleep(1 * time.Second)
	
	// Example 4: Practical use - disabling channels in select
	fmt.Println("\n4. Practical use: dynamically enable/disable channels")
	dynamicChannelControl()
}

func dynamicChannelControl() {
	var dataCh chan int
	var controlCh chan int
	
	controlCh = make(chan int, 1)
	
	go func() {
		for i := 0; i < 5; i++ {
			select {
			case dataCh <- i: // dataCh is nil, this case never selected
				fmt.Printf("Sent %d to dataCh\n", i)
			case ctrl := <-controlCh:
				fmt.Printf("Received control: %d\n", ctrl)
				if ctrl == 1 {
					// Enable data channel
					dataCh = make(chan int, 2)
					fmt.Println("Data channel enabled")
				} else if ctrl == 0 {
					// Disable data channel
					dataCh = nil
					fmt.Println("Data channel disabled")
				}
			default:
				fmt.Printf("Iteration %d: waiting...\n", i+1)
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	// Send control signals
	time.Sleep(300 * time.Millisecond)
	controlCh <- 1 // Enable data channel
	
	time.Sleep(500 * time.Millisecond)
	controlCh <- 0 // Disable data channel
	
	time.Sleep(500 * time.Millisecond)
}

func main() {
	nilChannelBehavior()
	
	fmt.Println("\n=== Nil Channel Summary ===")
	fmt.Println("1. Send to nil channel: blocks forever")
	fmt.Println("2. Receive from nil channel: blocks forever")
	fmt.Println("3. Close nil channel: panic")
	fmt.Println("4. In select: nil channels are never ready")
	fmt.Println("5. Practical use: dynamically disable channel cases")
	fmt.Println("6. Nil channels are useful for channel management patterns")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do nil channels behave in Go and what are their practical uses?

**Your Response:** Nil channels in Go have special blocking behavior - sending to or receiving from a nil channel blocks forever. This might seem like a bug, but it's actually a useful feature.

When I send to or receive from a nil channel, the operation blocks indefinitely. In select statements, nil channels are never considered ready, so their cases are never selected.

The practical application is dynamic channel control. I can set a channel to nil to temporarily disable it in a select statement. This is useful for patterns like:
- Temporarily disabling certain operations
- Implementing channel enable/disable logic
- Creating conditional channel behavior
- Managing multiple communication paths

For example, I can have a data channel that starts as nil (disabled), then enable it by assigning it a real channel when needed, and disable it again by setting it back to nil.

The key insight is that nil channels aren't errors - they're a tool for sophisticated channel management patterns. This demonstrates deep understanding of Go's channel semantics and advanced concurrency patterns.
