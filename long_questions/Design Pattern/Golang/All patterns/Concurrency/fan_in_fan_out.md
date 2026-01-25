# Fan-Out / Fan-In Pattern

## 🟢 What is it?
This pattern has two parts:
1.  **Fan-Out**: Starting multiple goroutines to handle input from a single channel. This distributes the workload.
2.  **Fan-In**: Multiplexing or combining multiple result channels into a single channel. This is often used to consolidate results from distributed tasks.

---

## 🏛️ Real World Analogy
**Call Center Operations**:
*   **Fan-Out**: A single phone number (input) routes calls to 50 support agents (workers). The work is "fanned out" to whoever is available.
*   **Fan-In**: Each agent types notes into their own computer. All these notes are "fanned in" to a single central database log.

---

## 🎯 Strategy to Implement

### Fan-Out
Simply looping over a channel and launching a goroutine for each item, or launching fixed workers (like a Worker Pool) is effectively Fan-Out.

### Fan-In
1.  **Multiple Inputs**: You have `N` channels effectively producing data.
2.  **Single Output**: You want one channel to read from.
3.  **Multiplexer Function**: Create a function that takes `...<-chan T` (variadic input channels).
4.  **WaitGroup**: Use a `sync.WaitGroup` to wait for all input channels to close.
5.  **Output Channel**: Create a single output channel.
6.  **Goroutines per Input**: For each input channel, start a goroutine that reads from it and sends to the output channel.
7.  **Closer**: Start one separate goroutine that waits for the WaitGroup to finish, then closes the output channel.

---

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Producer: Generates data on a channel
func producer(id int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < 3; i++ {
			time.Sleep(time.Millisecond * 100) // Simulate work
			out <- fmt.Sprintf("Producer %d: Item %d", id, i)
		}
	}()
	return out
}

// Fan-In: Combines multiple channels into one
func fanIn(inputs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	// Function to copy values from one channel to the 'out' channel
	output := func(c <-chan string) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(inputs))
	// Start a goroutine for each input channel
	for _, c := range inputs {
		go output(c)
	}

	// Start a separate goroutine to close 'out' once all inputs are done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// 1. Fan-Out: Start multiple producers (workers)
	// In a real scenario, this might be taking 1 job channel and splitting it,
	// or just starting multiple independent services.
	ch1 := producer(1)
	ch2 := producer(2)
	ch3 := producer(3)

	// 2. Fan-In: Combine results
	// We read from one single channel, even though 3 producers are running.
	merged := fanIn(ch1, ch2, ch3)

	// Consume the merged output
	for msg := range merged {
		fmt.Println(msg)
	}
}
```

---

## ✅ When to use?

*   **Microservices Aggregation**: When you call User Service, Order Service, and Analytics Service in parallel (Fan-Out), and want to return a single JSON response to the client once all are done (Fan-In).
*   **Data Processing Pipelines**: When multiple stagess of a pipeline produce results that need to be serialized into a single file.
