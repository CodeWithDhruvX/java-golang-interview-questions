package main

import (
	"fmt"
	"sync"
)

func main() {
	oddChan := make(chan bool)
	evenChan := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Odd goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= 20; i += 2 {
			<-oddChan // Wait for signal to print odd
			fmt.Print(i, " ")
			evenChan <- true // Signal even goroutine
		}
	}()

	// Even goroutine
	go func() {
		defer wg.Done()
		for i := 2; i <= 20; i += 2 {
			<-evenChan // Wait for signal to print even
			fmt.Print(i, " ")
			if i < 20 {
				oddChan <- true // Signal odd goroutine (except for last number)
			}
		}
	}()

	// Start with odd goroutine
	oddChan <- true
	wg.Wait()
	fmt.Println()
}


// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	// Unbuffered channel: make(chan int)
// 	// This acts as a synchronous lock between the two routines.
// 	ch := make(chan int)
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	// Goroutine 1: The Producer (Odds)
// 	go func() {
// 		defer wg.Done()
// 		for i := 1; i <= 9; i += 2 {
// 			// This line blocks until the other goroutine reads from 'ch'
// 			ch <- i
// 		}
// 		// Close the channel to tell the receiver we are done sending odds
// 		close(ch)
// 	}()

// 	// Goroutine 2: The Consumer (Evens)
// 	go func() {
// 		defer wg.Done()
// 		for i := 2; i <= 10; i += 2 {
// 			// 1. Wait for the odd number to arrive
// 			odd, ok := <-ch
// 			if ok {
// 				fmt.Printf("Odd:  %d\n", odd)
// 			}

// 			// 2. Print the even number immediately after
// 			fmt.Printf("Even: %d\n", i)
// 		}
// 	}()

// 	wg.Wait()
// }