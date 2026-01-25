# Barrier Pattern

## 🟢 What is it?
A **Barrier** is a synchronization primitive that prevents a group of goroutines from proceeding until they have all reached a certain point of execution. "We all move forward together, or nobody moves."

In Go, the `sync.WaitGroup` is the simplest form of a barrier (waiting for everyone to finish). For more complex "cyclic" barriers (where workers meet up multiple times), you might need a custom implementation or a channel structure.

---

## 🏛️ Real World Analogy
**Tour Bus**:
*   The bus is at a rest stop.
*   It cannot leave until **all 30 passengers** are back on board.
*   The driver counts them one by one. If 29 are there, everyone waits for the last person.
*   Once the count hits 30 (Barrier Reached), the bus moves to the next destination.

---

## 🎯 Strategy to Implement (Simple WaitGroup Barrier)

1.  **Initialize**: partial `sync.WaitGroup`.
2.  **Add**: `wg.Add(N)` where N is the number of workers.
3.  **Pass**: Pass the pointer `&wg` to workers (or use closure).
4.  **Done**: Workers call `wg.Done()` when they reach the barrier point.
5.  **Wait**: The main goroutine (or a controller) calls `wg.Wait()` to block until the counter is zero.

---

## 💻 Code Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this worker is done
	
	fmt.Printf("Worker %d starting... (takes %v)\n", id, duration)
	time.Sleep(duration) // Simulate work
	fmt.Printf("Worker %d arrived at barrier!\n", id)
}

func main() {
	var wg sync.WaitGroup
	
	fmt.Println("Manager: Waiting for all workers to finish task A...")

	// We have 3 workers
	wg.Add(3)
	
	// Start them with different speeds
	go worker(1, 2*time.Second, &wg)
	go worker(2, 4*time.Second, &wg)
	go worker(3, 1*time.Second, &wg)

	// Barrier: This blocks until counter is 0
	wg.Wait()

	fmt.Println("Manager: All workers arrived! Moving to next phase.")
}
```

---

## ✅ When to use?

*   **Parallel Initialization**: You have 3 microservices to fetch config from. You need *all* configs before the app can start serving traffic.
*   **Batch Processing**: You split a 1GB file into 10 chunks. You must wait for *all* 10 chunks to be processed before zipping the result.
