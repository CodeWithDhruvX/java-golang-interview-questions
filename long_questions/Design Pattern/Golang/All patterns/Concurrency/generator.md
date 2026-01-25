# Generator Pattern

## 🟢 What is it?
The **Generator Pattern** is a function that returns a **receive-only channel** (`<-chan T`). Inside the function, it spawns a goroutine that generates values and sends them to the channel.

This creates a "stream" of data that the consumer can read from without worrying about how the data is produced or how the concurrency is managed.

---

## 🏛️ Real World Analogy
**Candy Dispenser**:
*   You don't go into the kitchen and cook the candy yourself.
*   You just turn a handle (read from channel), and a piece of candy pops out.
*   The machine (Generator) hides the complexity of making and storing the candy.

---

## 🎯 Strategy to Implement

1.  **Function Signature**: Create a function that returns `<-chan T`.
2.  **Make Channel**: Inside, create `c := make(chan T)`.
3.  **Goroutine**: Launch `go func() { ... }`.
4.  **Loop & Send**: Inside the goroutine, generate data and send to `c`.
5.  **Cleanup**: Ensure to `close(c)` when done to prevent deadlocks for the range loop consumer.
6.  **Return**: Return the channel immediately (non-blocking).

---

## 💻 Code Example

```go
package main

import (
	"fmt"
	"math/rand"
)

// Generator: Returns a channel that emits 'n' random numbers
func randomGenerator(n int) <-chan int {
	c := make(chan int)
	
	// Start the work in the background
	go func() {
		defer close(c) // Always close when done!
		for i := 0; i < n; i++ {
			c <- rand.Intn(100)
		}
	}()
	
	return c
}

func main() {
	// 1. Create the generator
	// The function returns instantly, the work happens in background
	stream := randomGenerator(5)

	// 2. Consume the stream
	// We don't know "how" the numbers are made, we just consume them.
	for num := range stream {
		fmt.Println("Received:", num)
	}
}
```

---

## ✅ When to use?

*   **streaming Data**: Reading lines from a huge log file and passing them line-by-line to a processor.
*   **Infinite Sequences**: Generating an infinite stream of ID numbers (UIDs).
*   **Encapsulation**: You want to hide the complexity of channel creation and goroutine management from the API user.
