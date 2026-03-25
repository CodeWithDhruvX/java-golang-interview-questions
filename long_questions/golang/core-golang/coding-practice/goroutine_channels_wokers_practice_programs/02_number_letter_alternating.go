package main

import (
	"fmt"
	"sync"
)

func main() {
	numberChan := make(chan bool)
	letterChan := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Number goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			<-numberChan // Wait for signal
			fmt.Print(i, " ")
			letterChan <- true // Signal letter goroutine
		}
	}()

	// Letter goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 26; i++ {
			<-letterChan // Wait for signal
			fmt.Print(string('A'+i), " ")
			if i < 25 { // Don't signal after last letter
				numberChan <- true // Signal number goroutine
			}
		}
	}()

	// Start with number goroutine
	numberChan <- true
	wg.Wait()
	fmt.Println()
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does your solution ensure the pattern "1 A 2 B 3 C..." using two goroutines?

**Your Response:** I use two separate channels for coordination - one for numbers and one for letters. The number goroutine prints a number, then signals the letter goroutine. The letter goroutine prints a letter and signals back to the number goroutine.

The key insight is that each goroutine blocks on its respective channel until it receives permission to proceed. This creates a strict alternating pattern. I start the process by signaling the number goroutine first.

The main function uses a WaitGroup to wait for both goroutines to complete. The conditional check in the letter goroutine prevents it from signaling after the last letter, avoiding a deadlock.

This approach demonstrates understanding of goroutine coordination using channels as synchronization primitives, which is fundamental to Go concurrency.
