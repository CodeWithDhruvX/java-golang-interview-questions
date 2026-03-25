package main

import (
	"fmt"
	"sync"
	"time"
)

type DataStore struct {
	mu   sync.RWMutex
	data []string
}

func NewDataStore() *DataStore {
	return &DataStore{
		data: make([]string, 0),
	}
}

func (ds *DataStore) Read(index int) (string, bool) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	
	if index >= len(ds.data) {
		return "", false
	}
	return ds.data[index], true
}

func (ds *DataStore) Write(item string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.data = append(ds.data, item)
}

func (ds *DataStore) Length() int {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return len(ds.data)
}

func reader(id int, store *DataStore, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		length := store.Length()
		if length > 0 {
			index := i % length
			if value, exists := store.Read(index); exists {
				fmt.Printf("Reader %d: read '%s' at index %d\n", id, value, index)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Reader %d: completed\n", id)
}

func writer(id int, store *DataStore, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		item := fmt.Sprintf("Item_%d_%d", id, i)
		store.Write(item)
		fmt.Printf("Writer %d: wrote '%s'\n", id, item)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Printf("Writer %d: completed\n", id)
}

func main() {
	store := NewDataStore()
	wg := sync.WaitGroup{}

	fmt.Println("Demonstrating RWMutex: multiple readers, single writer")

	// Start multiple readers (can read concurrently)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go reader(i, store, &wg)
	}

	// Start single writer (exclusive access for writes)
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go writer(i, store, &wg)
	}

	wg.Wait()
	fmt.Printf("Final data store length: %d\n", store.Length())
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does RWMutex differ from regular Mutex and when would you use it?

**Your Response:** RWMutex (Read-Write Mutex) provides different locking for readers and writers. Multiple goroutines can hold read locks simultaneously, but only one can hold a write lock, and write locks are exclusive.

I use RLock() for read operations, which allows multiple readers to access the data concurrently. For write operations, I use Lock() which provides exclusive access, blocking both readers and other writers.

This is more efficient than a regular Mutex when you have many more reads than writes, because readers don't block each other. In the example, 5 readers can read simultaneously, while writers get exclusive access when they need to modify the data.

The key insight is that RWMutex optimizes for read-heavy workloads. It's perfect for scenarios like configuration data, caches, or reference data that's read frequently but updated rarely. This pattern demonstrates understanding of lock granularity and performance optimization in concurrent systems.
