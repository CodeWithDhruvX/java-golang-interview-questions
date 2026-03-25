package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, exists := sm.data[key]
	return value, exists
}

func (sm *SafeMap) Increment(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key]++
}

func writer(id int, safeMap *SafeMap, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key_%d", i%10)
		safeMap.Increment(key)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("Writer %d: completed\n", id)
}

func reader(id int, safeMap *SafeMap, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("key_%d", i%10)
		if value, exists := safeMap.Get(key); exists {
			fmt.Printf("Reader %d: %s = %d\n", id, key, value)
		}
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Printf("Reader %d: completed\n", id)
}

func main() {
	safeMap := NewSafeMap()
	wg := sync.WaitGroup{}

	fmt.Println("Starting concurrent map operations with Mutex protection")

	// Start writers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go writer(i, safeMap, &wg)
	}

	// Start readers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go reader(i, safeMap, &wg)
	}

	wg.Wait()
	fmt.Println("All operations completed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you protect a shared map from concurrent access using Mutex?

**Your Response:** I create a SafeMap struct that wraps a regular map and protects all access with a Mutex. The map itself is not thread-safe, so every read and write operation must be protected.

The Set, Get, and Increment methods all lock the mutex before accessing the underlying map and unlock it afterward using defer. This ensures that only one goroutine can modify or read the map at a time, preventing race conditions.

I demonstrate this with multiple writers incrementing map values and multiple readers accessing those values concurrently. Without the mutex protection, concurrent map writes could cause data corruption or panics.

The key insight is that Go's built-in map type is not safe for concurrent access - you must provide your own synchronization. This pattern is essential for any shared mutable state in concurrent Go programs.
