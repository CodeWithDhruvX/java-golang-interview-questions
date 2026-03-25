package main

import (
	"fmt"
	"sync"
	"time"
)

type Barrier struct {
	count    int
	waiting  int
	release  chan struct{}
	mu       sync.Mutex
	wg       sync.WaitGroup
}

func NewBarrier(count int) *Barrier {
	return &Barrier{
		count:   count,
		release: make(chan struct{}),
	}
}

func (b *Barrier) Wait() {
	b.mu.Lock()
	b.waiting++
	
	if b.waiting == b.count {
		// Last goroutine to arrive - release all
		close(b.release)
		b.mu.Unlock()
		return
	}
	
	b.mu.Unlock()
	
	// Wait for release
	<-b.release
}

func workerPhase1(id int, barrier *Barrier, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Printf("Worker %d: starting phase 1\n", id)
	time.Sleep(time.Duration(id) * 200 * time.Millisecond)
	fmt.Printf("Worker %d: completed phase 1, waiting at barrier\n", id)
	
	barrier.Wait()
	
	fmt.Printf("Worker %d: passed barrier, starting phase 2\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("Worker %d: completed phase 2\n", id)
}

func main() {
	numWorkers := 4
	barrier := NewBarrier(numWorkers)
	wg := sync.WaitGroup{}
	
	fmt.Printf("Starting %d workers with barrier synchronization\n", numWorkers)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workerPhase1(i, barrier, &wg)
	}

	wg.Wait()
	fmt.Println("All workers completed both phases")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement a barrier that ensures all goroutines complete phase 1 before any start phase 2?

**Your Response:** I implement a barrier using a counter and a release channel. The barrier tracks how many goroutines have arrived and releases them all when the count reaches the required number.

Each worker calls Wait() after completing phase 1. The Wait method increments the waiting count. The last goroutine to arrive (when waiting equals count) closes the release channel, which unblocks all waiting goroutines simultaneously.

The key insight is that a closed channel broadcasts to all waiting goroutines - they all receive the signal and can proceed to phase 2 together. This ensures no goroutine starts phase 2 until all have completed phase 1.

I use a mutex to protect the shared waiting count from race conditions. This pattern is essential for parallel algorithms that have distinct phases, like parallel matrix operations, map-reduce workflows, or any multi-stage concurrent processing where phase boundaries must be strictly enforced.
