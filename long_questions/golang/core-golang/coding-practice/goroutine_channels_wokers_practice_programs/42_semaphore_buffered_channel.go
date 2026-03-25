package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	permits chan struct{}
}

func NewSemaphore(capacity int) *Semaphore {
	return &Semaphore{
		permits: make(chan struct{}, capacity),
	}
}

func (s *Semaphore) Acquire() {
	s.permits <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.permits
}

func (s *Semaphore) TryAcquire() bool {
	select {
	case s.permits <- struct{}{}:
		return true
	default:
		return false
	}
}

func worker(id int, semaphore *Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()
	
	fmt.Printf("Worker %d: waiting for permit\n", id)
	semaphore.Acquire()
	defer semaphore.Release()
	
	fmt.Printf("Worker %d: acquired permit, working\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d: finished, released permit\n", id)
}

func main() {
	// Semaphore with capacity 3 (max 3 concurrent workers)
	semaphore := NewSemaphore(3)
	wg := sync.WaitGroup{}
	
	fmt.Println("Starting semaphore with capacity 3 (max 3 concurrent workers)")

	// Start more workers than semaphore capacity
	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go worker(i, semaphore, &wg)
	}

	wg.Wait()
	fmt.Println("All workers completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement a semaphore using a buffered channel?

**Your Response:** I implement a semaphore using a buffered channel where each element represents a permit. The channel capacity equals the maximum number of concurrent operations allowed.

The Acquire method sends to the channel, which blocks if the channel is full (no permits available). The Release method receives from the channel, freeing up a permit for other workers. This creates a natural counting semaphore.

With capacity 3, only 3 workers can acquire permits simultaneously. When a 4th worker tries to acquire, it blocks until one of the first 3 workers releases its permit.

The key insight is that buffered channels naturally implement semaphores - the channel buffer represents available permits, sending acquires permits, and receiving releases them. This pattern is perfect for limiting concurrent access to resources like database connections, file handles, or API rate limits.

I also provide a TryAcquire method using non-blocking select, which is useful when you don't want to block if no permits are available.
