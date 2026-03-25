
not working code:
// You can edit this code!
// Click here and start typing.
// two goroutine and one channel make sure that ordeer of the element 1 to 10
package main

import "sync"

//"sync"

func consumer(ele <-chan int) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			for n:=range ele{
				ele<-n
			}
		}()

	}
}

func main() {
	odds := make(chan int, 10)
	evens := make(chan int, 10)
	// odd numbers

	go func() {
		for i := 0; i < 10; i++ {
			///wg.Add(1)
			if i%2 == 1 {
				//defer wg.Done()
				result <- i
			}

		}
		//wg.Wait()
		close(result)
	}()

	

}
how to fix these program in golang



package main

import (
	"fmt"
	"sync"
)

func main() {
	// Unbuffered channel: make(chan int)
	// This acts as a synchronous lock between the two routines.
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1: The Producer (Odds)
	go func() {
		defer wg.Done()
		for i := 1; i <= 9; i += 2 {
			// This line blocks until the other goroutine reads from 'ch'
			ch <- i
		}
		// Close the channel to tell the receiver we are done sending odds
		close(ch)
	}()

	// Goroutine 2: The Consumer (Evens)
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			// 1. Wait for the odd number to arrive
			odd, ok := <-ch
			if ok {
				fmt.Printf("Odd:  %d\n", odd)
			}

			// 2. Print the even number immediately after
			fmt.Printf("Even: %d\n", i)
		}
	}()

	wg.Wait()
}



