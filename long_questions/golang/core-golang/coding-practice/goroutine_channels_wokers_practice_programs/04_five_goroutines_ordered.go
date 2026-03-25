package main

import (
	"fmt"
	"sync"
)

func main() {
	channels := make([]chan bool, 5)
	for i := range channels {
		channels[i] = make(chan bool)
	}
	wg := sync.WaitGroup{}
	wg.Add(5)

	// Create 5 goroutines
	for i := 0; i < 5; i++ {
		go func(goroutineID int) {
			defer wg.Done()
			for num := goroutineID + 1; num <= 50; num += 5 {
				<-channels[goroutineID] // Wait for signal
				fmt.Print(num, " ")
				nextID := (goroutineID + 1) % 5
				if num < 50 || goroutineID < 4 { // Don't signal after last number
					channels[nextID] <- true // Signal next goroutine
				}
			}
		}(i)
	}

	// Start with first goroutine
	channels[0] <- true
	wg.Wait()
	fmt.Println()
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How would you scale the coordination pattern to five goroutines while maintaining order?

**Your Response:** I scaled the circular signaling pattern to five goroutines using an array of channels. Each goroutine has its own channel and waits for permission before printing.

The logic is similar to the three-goroutine solution but generalized. Goroutine 0 prints 1, 6, 11... then signals goroutine 1. Goroutine 1 prints 2, 7, 12... then signals goroutine 2, and so on. Goroutine 4 signals back to goroutine 0, completing the circle.

I use modulo arithmetic `(goroutineID + 1) % 5` to calculate the next goroutine to signal. The conditional check ensures we don't signal after the last number to prevent deadlock.

This approach demonstrates scalability - the same pattern works for any number of goroutines. It shows understanding of concurrent programming patterns that can be extended to real-world scenarios with multiple workers.
