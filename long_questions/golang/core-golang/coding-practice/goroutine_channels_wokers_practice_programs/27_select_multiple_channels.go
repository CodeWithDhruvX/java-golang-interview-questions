package main

import (
	"fmt"
	"sync"
	"time"
)

func channel1(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	messages := []string{"A1", "A2", "A3"}
	for _, msg := range messages {
		time.Sleep(300 * time.Millisecond)
		ch <- msg
		fmt.Printf("Channel 1 sent: %s\n", msg)
	}
}

func channel2(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	messages := []string{"B1", "B2", "B3"}
	for _, msg := range messages {
		time.Sleep(200 * time.Millisecond)
		ch <- msg
		fmt.Printf("Channel 2 sent: %s\n", msg)
	}
}

func main() {
	ch1 := make(chan string, 3)
	ch2 := make(chan string, 3)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Start senders
	go channel1(ch1, &wg)
	go channel2(ch2, &wg)

	// Read from both channels using select
	fmt.Println("Reading from both channels using select:")
	for i := 0; i < 6; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from channel 1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from channel 2: %s\n", msg2)
		}
	}

	wg.Wait()
	fmt.Println("All messages received")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does the select statement allow reading from multiple channels?

**Your Response:** The select statement allows a goroutine to wait on multiple channel operations simultaneously. It blocks until one of the cases can proceed, then executes that case.

In this example, I have two channels with different send rates. Channel 1 sends every 300ms, Channel 2 sends every 200ms. The select statement checks both channels and receives from whichever one has data available first.

If both channels have data ready, Go randomly selects one to ensure fairness. This creates a multiplexing effect where we can handle multiple data streams without blocking on any single one.

The select loop continues until all 6 messages are received (3 from each channel). This pattern is fundamental to Go concurrency and is used extensively in real systems for handling multiple input sources, implementing timeouts, or managing cancellation.
