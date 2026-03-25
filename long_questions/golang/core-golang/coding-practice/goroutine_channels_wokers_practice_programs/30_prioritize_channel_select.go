package main

import (
	"fmt"
	"sync"
	"time"
)

func prioritySender(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	messages := []string{"HIGH: Critical", "HIGH: Urgent", "HIGH: Important"}
	for _, msg := range messages {
		time.Sleep(400 * time.Millisecond)
		ch <- msg
		fmt.Printf("Priority sent: %s\n", msg)
	}
}

func normalSender(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	messages := []string{"NORMAL: Regular", "NORMAL: Standard", "NORMAL: Routine"}
	for _, msg := range messages {
		time.Sleep(300 * time.Millisecond)
		ch <- msg
		fmt.Printf("Normal sent: %s\n", msg)
	}
}

func main() {
	priorityCh := make(chan string, 3)
	normalCh := make(chan string, 3)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Start senders
	go prioritySender(priorityCh, &wg)
	go normalSender(normalCh, &wg)

	// Read with priority using select
	fmt.Println("Reading with priority (priority channel checked first):")
	
	for i := 0; i < 6; i++ {
		select {
		case priorityMsg := <-priorityCh:
			fmt.Printf("🔴 PRIORITY: %s\n", priorityMsg)
		case normalMsg := <-normalCh:
			fmt.Printf("🔵 NORMAL: %s\n", normalMsg)
		}
	}

	wg.Wait()
	fmt.Println("All messages processed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you prioritize one channel over another using select?

**Your Response:** I prioritize channels by ordering the cases in the select statement. Go evaluates select cases in source order, so the priority channel is checked first.

In this example, I have a priority channel and a normal channel. When both channels have data available, the priority channel case (listed first) will be chosen. This creates a priority system where critical messages are processed before regular ones.

The select statement still provides fairness - if only the normal channel has data, it will be processed. But when both have data, priority is given to the first case.

This pattern is extremely useful in real systems for handling different priority levels of work, like processing critical system alerts before regular logs, or handling VIP customer requests before regular ones. It demonstrates understanding of select case ordering for implementing priority queues.
