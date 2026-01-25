# Semaphore Pattern

## 🟢 What is it?
The **Semaphore Pattern** is used to control access to a shared resource by maintaining a set of permits. If a permit is available, access is granted. If not, the request blocks until a permit is released. 

In Go, we don't need a complex mutex or condition variable to build this; a **Buffered Channel** acts as a perfect counting semaphore.

---

## 🏛️ Real World Analogy
**Nightclub Bouncer**:
*   The club has a capacity of **50 people**.
*   The Bouncer (Semaphore) has a clicker count.
*   If the count is < 50, you can enter (Acquire).
*   If the count is 50, you wait in line (Block) until someone leaves (Release).

---

## 🎯 Strategy to Implement

1.  **Define Limit**: Decide the max concurrency (e.g., `max = 5`).
2.  **Create Channel**: `sem := make(chan struct{}, max)` (Buffered channel of size `max`).
3.  **Acquire**: Before starting a heavy task, send into the channel: `sem <- struct{}{}`. If the buffer is full, this line blocks.
4.  **Do Work**: Run your logic.
5.  **Release**: After work is done (usually in `defer`), read from the channel: `<-sem`. This frees up a slot in the buffer.

---

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	requests := 15
	concurrencyLimit := 3

	// 1. Create a Semaphore (Buffered Channel)
	sem := make(chan struct{}, concurrencyLimit)

	for i := 1; i <= requests; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			// 2. Acquire Token
			fmt.Printf("Request %d waiting...\n", id)
			sem <- struct{}{} // Blocks here if channel is full (3 items)
			
			// 3. Critical Section / Heavy Work
			fmt.Printf("Request %d acquired semaphore. Processing...\n", id)
			time.Sleep(2 * time.Second) // Simulate expensive I/O
			
			// 4. Release Token
			<-sem
			fmt.Printf("Request %d released semaphore.\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All requests processed.")
}
```

---

## ✅ When to use?

*   **API Rate Limiting**: You want to scrape a website but they ban you if you make more than 5 rps (requests per second) concurrent.
*   **Database Connections**: Your DB crashes if you open 1000 connections at once. You use a semaphore to limit it to 50 active queries.
*   **File Open Limit**: Preventing "Too many open files" error by limiting concurrent file reads.
