package main

import (
	"fmt"
	"sync"
	"time"
)

type Number struct {
	Value int
	Source string
}

func splitter(numbers <-chan Number, even chan<- Number, odd chan<- Number, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(even)
	defer close(odd)
	
	for num := range numbers {
		if num.Value%2 == 0 {
			even <- num
			fmt.Printf("Splitter: %d -> even channel\n", num.Value)
		} else {
			odd <- num
			fmt.Printf("Splitter: %d -> odd channel\n", num.Value)
		}
	}
}

func processor(id string, numbers <-chan Number, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range numbers {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("Processor %s: processed %d (%s)\n", 
			id, num.Value, num.Source)
	}
}

func main() {
	numbers := make(chan Number, 10)
	even := make(chan Number, 5)
	odd := make(chan Number, 5)
	wg := sync.WaitGroup{}

	// Start splitter
	wg.Add(1)
	go splitter(numbers, even, odd, &wg)

	// Start processors
	wg.Add(2)
	go processor("EVEN", even, &wg)
	go processor("ODD", odd, &wg)

	// Send numbers
	go func() {
		for i := 1; i <= 10; i++ {
			num := Number{
				Value:  i,
				Source: fmt.Sprintf("Input %d", i),
			}
			numbers <- num
			fmt.Printf("Main: sent %d\n", i)
		}
		close(numbers)
	}()

	wg.Wait()
	fmt.Println("Channel splitting completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you split one channel into two separate channels based on content?

**Your Response:** I use a splitter goroutine that reads from the input channel and routes each number to either the even or odd channel based on its value.

The splitter examines each number using the modulo operator - if the number is divisible by 2, it goes to the even channel; otherwise, it goes to the odd channel. The splitter also handles closing both output channels when it's done.

Two separate processor goroutines read from the even and odd channels independently. This allows parallel processing of the two data streams after the split.

This pattern is useful for categorizing and routing data based on content. It's commonly used in real systems for log routing (error vs info), message filtering, or data preprocessing where different types of data need different handling. The pattern demonstrates understanding of channel fan-out for data routing.
