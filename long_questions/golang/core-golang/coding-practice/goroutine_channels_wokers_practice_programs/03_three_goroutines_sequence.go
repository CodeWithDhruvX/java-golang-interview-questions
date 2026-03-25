package main

import (
	"fmt"
	"sync"
)

func main() {
	channels := []chan bool{
		make(chan bool), // Channel for G1
		make(chan bool), // Channel for G2
		make(chan bool), // Channel for G3
	}
	wg := sync.WaitGroup{}
	wg.Add(3)

	// Goroutine 1: prints 1, 4, 7, 10...
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 3 {
			<-channels[0] // Wait for signal
			fmt.Print("G1:", i, " ")
			channels[1] <- true // Signal G2
		}
	}()

	// Goroutine 2: prints 2, 5, 8, 11...
	go func() {
		defer wg.Done()
		for i := 2; i <= 11; i += 3 {
			<-channels[1] // Wait for signal
			fmt.Print("G2:", i, " ")
			channels[2] <- true // Signal G3
		}
	}()

	// Goroutine 3: prints 3, 6, 9, 12...
	go func() {
		defer wg.Done()
		for i := 3; i <= 12; i += 3 {
			<-channels[2] // Wait for signal
			fmt.Print("G3:", i, " ")
			if i < 12 { // Don't signal after last iteration
				channels[0] <- true // Signal G1
			}
		}
	}()

	// Start the chain with G1
	channels[0] <- true
	wg.Wait()
	fmt.Println()
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you coordinate three goroutines to print numbers in sequence?

**Your Response:** I created a circular signaling pattern using three separate channels. Each goroutine has its own dedicated channel where it waits for permission to proceed.

G1 prints numbers 1, 4, 7... then signals G2. G2 prints 2, 5, 8... then signals G3. G3 prints 3, 6, 9... then signals back to G1, creating a circular chain.

The key is that each goroutine blocks on its own channel until the previous goroutine signals it. This ensures perfect ordering: G1→G2→G3→G1→G2→G3 and so on.

I start the chain by signaling G1 first. The conditional check in G3 prevents signaling after the last iteration to avoid deadlock. The WaitGroup ensures the main function waits for completion.

This pattern demonstrates understanding of multi-goroutine coordination and can be extended to any number of goroutines.
