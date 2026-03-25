package main

import (
	"fmt"
	"sync"
)

func main() {
	turnChan := make(chan bool, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Goroutine A
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-turnChan // Wait for turn
			fmt.Println("Goroutine A - Turn", i+1)
			turnChan <- true // Pass turn to B
		}
	}()

	// Goroutine B
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine B - Turn", i+1)
			turnChan <- true // Pass turn to A
			<-turnChan // Wait for turn
		}
	}()

	// Start with Goroutine A
	turnChan <- true
	wg.Wait()
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement strict turn-based execution between two goroutines?

**Your Response:** I use a single buffered channel as a turn-taking mechanism. The channel acts like a token that only one goroutine can possess at a time.

Goroutine A waits for the turn token, prints its message, then passes the token back to Goroutine B. Goroutine B prints its message and passes the token to A. This creates a strict alternating pattern.

I used a buffered channel with capacity 1 to allow the initial turn to be placed without blocking. The WaitGroup ensures both goroutines complete before the main function exits.

This pattern is essentially a mutex implementation using channels - it ensures mutual exclusion and fair scheduling. It's a fundamental concurrency pattern that demonstrates understanding of synchronization primitives.
