# Bulkhead Pattern

## 🟢 What is it?
The **Bulkhead Pattern** isolates elements of an application into pools so that if one fails, the others continue to function. It prevents a crash in one part of the system from bringing down the entire system.

---

## 🏛️ Real World Analogy
**Ship Bulkheads**:
*   Ships are divided into watertight compartments (bulkheads).
*   If the hull is breached in one section, only that section floods.
*   The ship doesn't sink because the other sections stay dry and buoyant.

---

## 🎯 Strategy to Implement

1.  **Separate Connection Pools**: Don't share one DB connection pool for "User Read", "Order Write", and "Analytics". If Analytics locks up all connections, Users can't log in. Create 3 separate pools.
2.  **Separate Goroutine Pools**: Don't run everything on the default scheduler if some tasks are blocking. Create dedicated Worker Pools for different services.
3.  **Resource Limits**: Limit CPU/Memory per container/service.

---

## 💻 Code Example (Separate Worker Pools)

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// A generic worker pool
func createPool(name string, capacity int, jobs <-chan int, wg *sync.WaitGroup) {
	for i := 0; i < capacity; i++ {
		go func(workerID int) {
			for j := range jobs {
				fmt.Printf("[%s] Worker %d processing job %d\n", name, workerID, j)
				time.Sleep(100 * time.Millisecond) // Simulating work
				wg.Done()
			}
		}(i)
	}
}

func main() {
	var wg sync.WaitGroup

	// 1. Critical Pool (User Logins) - Reserved capacity
	criticalJobs := make(chan int, 10)
	createPool("CRITICAL", 5, criticalJobs, &wg)

	// 2. Non-Critical Pool (Image Resizing) - If this gets clogged, logins still work!
	backgroundJobs := make(chan int, 10)
	createPool("BACKGROUND", 2, backgroundJobs, &wg)

	// Submit jobs
	for i := 0; i < 5; i++ {
		wg.Add(2)
		criticalJobs <- i
		backgroundJobs <- i
	}

	close(criticalJobs)
	close(backgroundJobs)
	wg.Wait()
}
```

---

## ✅ When to use?

*   **Mixed Workloads**: You have fast APIs (User Profile) and slow APIS (Report Generation). Separate them so reports don't starve user profiles.
*   **Failover**: If the "Recommendation Service" is down, the "Show Cart" page should still work.
