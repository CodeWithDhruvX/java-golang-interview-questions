package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func worker(id int, counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter.Increment()
		if i%100 == 0 {
			fmt.Printf("Worker %d: count = %d\n", id, counter.Value())
		}
	}
}

func main() {
	counter := &Counter{}
	numWorkers := 5
	wg := sync.WaitGroup{}

	fmt.Printf("Starting %d workers incrementing shared counter\n", numWorkers)
	
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, counter, &wg)
	}

	wg.Wait()
	fmt.Printf("Final count: %d\n", counter.Value())
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you protect a shared counter from race conditions using Mutex?

**Your Response:** I use a Mutex (mutual exclusion lock) to ensure only one goroutine can access the counter at a time. The Counter struct has a Mutex field that protects the value field.

The Increment method locks the mutex before modifying the value and unlocks it afterward using defer. This prevents race conditions where multiple goroutines might try to increment simultaneously, which could lead to lost increments.

The Value method also locks the mutex to ensure a consistent read while other goroutines might be writing. This prevents reading partially updated values.

The key insight is that any shared mutable state must be protected by synchronization primitives. Without the mutex, concurrent increments would corrupt the final count. This pattern is fundamental to safe concurrent programming in Go.
