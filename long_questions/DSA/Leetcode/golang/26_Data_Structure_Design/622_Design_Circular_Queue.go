package main

import (
	"fmt"
	"sync"
	"time"
)

// 622. Design Circular Queue
// Time: O(1) for all operations, Space: O(k) where k is capacity
type MyCircularQueue struct {
	buffer     []int
	head       int
	tail       int
	size       int
	capacity   int
	mutex      sync.Mutex
}

// Constructor initializes the queue with a specified capacity
func ConstructorMyCircularQueue(k int) MyCircularQueue {
	return MyCircularQueue{
		buffer:   make([]int, k),
		head:     0,
		tail:     -1,
		size:     0,
		capacity: k,
	}
}

// EnQueue inserts an element into the circular queue. Returns true if the operation is successful.
func (this *MyCircularQueue) EnQueue(value int) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	
	if this.IsFull() {
		return false
	}
	
	this.tail = (this.tail + 1) % this.capacity
	this.buffer[this.tail] = value
	this.size++
	
	return true
}

// DeQueue deletes an element from the circular queue. Returns true if the operation is successful.
func (this *MyCircularQueue) DeQueue() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	
	if this.IsEmpty() {
		return false
	}
	
	this.head = (this.head + 1) % this.capacity
	this.size--
	
	return true
}

// Front gets the front item from the queue.
func (this *MyCircularQueue) Front() int {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	
	if this.IsEmpty() {
		return -1
	}
	
	return this.buffer[this.head]
}

// Rear gets the last item from the queue.
func (this *MyCircularQueue) Rear() int {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	
	if this.IsEmpty() {
		return -1
	}
	
	return this.buffer[this.tail]
}

// IsEmpty checks whether the circular queue is empty or not.
func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

// IsFull checks whether the circular queue is full or not.
func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.capacity
}

// Alternative implementation without mutex (single-threaded)
type MyCircularQueueSimple struct {
	buffer   []int
	head     int
	tail     int
	size     int
	capacity int
}

func ConstructorMyCircularQueueSimple(k int) MyCircularQueueSimple {
	return MyCircularQueueSimple{
		buffer:   make([]int, k),
		head:     0,
		tail:     -1,
		size:     0,
		capacity: k,
	}
}

func (this *MyCircularQueueSimple) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	
	this.tail = (this.tail + 1) % this.capacity
	this.buffer[this.tail] = value
	this.size++
	
	return true
}

func (this *MyCircularQueueSimple) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	
	this.head = (this.head + 1) % this.capacity
	this.size--
	
	return true
}

func (this *MyCircularQueueSimple) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[this.head]
}

func (this *MyCircularQueueSimple) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[this.tail]
}

func (this *MyCircularQueueSimple) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueueSimple) IsFull() bool {
	return this.size == this.capacity
}

// Version with detailed tracking and statistics
type MyCircularQueueWithStats struct {
	buffer      []int
	head        int
	tail        int
	size        int
	capacity    int
	enqueueCount int
	dequeueCount int
	frontCount   int
	rearCount    int
}

func ConstructorMyCircularQueueWithStats(k int) MyCircularQueueWithStats {
	return MyCircularQueueWithStats{
		buffer:   make([]int, k),
		head:     0,
		tail:     -1,
		size:     0,
		capacity: k,
	}
}

func (this *MyCircularQueueWithStats) EnQueue(value int) bool {
	this.enqueueCount++
	
	if this.IsFull() {
		return false
	}
	
	this.tail = (this.tail + 1) % this.capacity
	this.buffer[this.tail] = value
	this.size++
	
	return true
}

func (this *MyCircularQueueWithStats) DeQueue() bool {
	this.dequeueCount++
	
	if this.IsEmpty() {
		return false
	}
	
	this.head = (this.head + 1) % this.capacity
	this.size--
	
	return true
}

func (this *MyCircularQueueWithStats) Front() int {
	this.frontCount++
	
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[this.head]
}

func (this *MyCircularQueueWithStats) Rear() int {
	this.rearCount++
	
	if this.IsEmpty() {
		return -1
	}
	return this.buffer[this.tail]
}

func (this *MyCircularQueueWithStats) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueueWithStats) IsFull() bool {
	return this.size == this.capacity
}

func (this *MyCircularQueueWithStats) GetStats() (int, int, int, int) {
	return this.enqueueCount, this.dequeueCount, this.frontCount, this.rearCount
}

func (this *MyCircularQueueWithStats) GetBuffer() []int {
	result := make([]int, 0, this.size)
	for i := 0; i < this.size; i++ {
		idx := (this.head + i) % this.capacity
		result = append(result, this.buffer[idx])
	}
	return result
}

// Thread-safe version with channel-based implementation
type MyCircularQueueChannel struct {
	buffer   chan int
	capacity int
}

func ConstructorMyCircularQueueChannel(k int) MyCircularQueueChannel {
	return MyCircularQueueChannel{
		buffer:   make(chan int, k),
		capacity: k,
	}
}

func (this *MyCircularQueueChannel) EnQueue(value int) bool {
	select {
	case this.buffer <- value:
		return true
	default:
		return false // Channel is full
	}
}

func (this *MyCircularQueueChannel) DeQueue() bool {
	select {
	case <-this.buffer:
		return true
	default:
		return false // Channel is empty
	}
}

func (this *MyCircularQueueChannel) Front() int {
	select {
	case value := <-this.buffer:
		// Put it back
		select {
		case this.buffer <- value:
			return value
		default:
			return -1 // Shouldn't happen
		}
	default:
		return -1
	}
}

func (this *MyCircularQueueChannel) Rear() int {
	if len(this.buffer) == 0 {
		return -1
	}
	
	// Read all elements and put them back to find the last one
	var lastValue int
	for i := 0; i < len(this.buffer); i++ {
		value := <-this.buffer
		lastValue = value
		this.buffer <- value
	}
	
	return lastValue
}

func (this *MyCircularQueueChannel) IsEmpty() bool {
	return len(this.buffer) == 0
}

func (this *MyCircularQueueChannel) IsFull() bool {
	return len(this.buffer) == this.capacity
}

func main() {
	// Test cases
	fmt.Println("=== Testing MyCircularQueue ===")
	
	// Test 1: Basic operations
	cq := ConstructorMyCircularQueue(3)
	
	fmt.Printf("EnQueue 1: %t\n", cq.EnQueue(1))
	fmt.Printf("EnQueue 2: %t\n", cq.EnQueue(2))
	fmt.Printf("EnQueue 3: %t\n", cq.EnQueue(3))
	fmt.Printf("EnQueue 4 (full): %t\n", cq.EnQueue(4))
	
	fmt.Printf("Rear: %d\n", cq.Rear())
	fmt.Printf("Front: %d\n", cq.Front())
	
	fmt.Printf("DeQueue: %t\n", cq.DeQueue())
	fmt.Printf("Front after dequeue: %d\n", cq.Front())
	fmt.Printf("EnQueue 4: %t\n", cq.EnQueue(4))
	fmt.Printf("Rear: %d\n", cq.Rear())
	
	// Test 2: Edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	emptyCq := ConstructorMyCircularQueue(2)
	
	fmt.Printf("Empty queue - Front: %d, Rear: %d\n", emptyCq.Front(), emptyCq.Rear())
	fmt.Printf("Empty queue - DeQueue: %t\n", emptyCq.DeQueue())
	fmt.Printf("Empty queue - IsEmpty: %t, IsFull: %t\n", emptyCq.IsEmpty(), emptyCq.IsFull())
	
	emptyCq.EnQueue(10)
	fmt.Printf("After enqueue - Front: %d, Rear: %d\n", emptyCq.Front(), emptyCq.Rear())
	
	// Test 3: Simple version
	fmt.Println("\n=== Testing Simple Version ===")
	simpleCq := ConstructorMyCircularQueueSimple(2)
	
	simpleCq.EnQueue(100)
	simpleCq.EnQueue(200)
	fmt.Printf("Simple queue - Front: %d, Rear: %d\n", simpleCq.Front(), simpleCq.Rear())
	fmt.Printf("IsFull: %t\n", simpleCq.IsFull())
	
	// Test 4: Statistics version
	fmt.Println("\n=== Testing Statistics Version ===")
	statsCq := ConstructorMyCircularQueueWithStats(3)
	
	statsCq.EnQueue(10)
	statsCq.EnQueue(20)
	statsCq.EnQueue(30)
	statsCq.DeQueue()
	statsCq.Front()
	statsCq.Rear()
	
	enqueues, dequeues, fronts, rears := statsCq.GetStats()
	fmt.Printf("Stats - Enqueues: %d, Dequeues: %d, Fronts: %d, Rears: %d\n", 
		enqueues, dequeues, fronts, rears)
	fmt.Printf("Buffer: %v\n", statsCq.GetBuffer())
	
	// Test 5: Channel version
	fmt.Println("\n=== Testing Channel Version ===")
	channelCq := ConstructorMyCircularQueueChannel(2)
	
	fmt.Printf("Channel EnQueue 1: %t\n", channelCq.EnQueue(1))
	fmt.Printf("Channel EnQueue 2: %t\n", channelCq.EnQueue(2))
	fmt.Printf("Channel EnQueue 3 (full): %t\n", channelCq.EnQueue(3))
	
	fmt.Printf("Channel Front: %d, Rear: %d\n", channelCq.Front(), channelCq.Rear())
	
	fmt.Printf("Channel DeQueue: %t\n", channelCq.DeQueue())
	fmt.Printf("Channel Front after dequeue: %d\n", channelCq.Front())
	
	// Test 6: Circular behavior
	fmt.Println("\n=== Testing Circular Behavior ===")
	circularCq := ConstructorMyCircularQueue(3)
	
	// Fill the queue
	for i := 1; i <= 3; i++ {
		circularCq.EnQueue(i * 10)
	}
	fmt.Printf("Full queue - Front: %d, Rear: %d\n", circularCq.Front(), circularCq.Rear())
	
	// Remove and add to test circular behavior
	circularCq.DeQueue()
	circularCq.EnQueue(40)
	fmt.Printf("After circular operation - Front: %d, Rear: %d\n", circularCq.Front(), circularCq.Rear())
	
	circularCq.DeQueue()
	circularCq.DeQueue()
	circularCq.EnQueue(50)
	circularCq.EnQueue(60)
	fmt.Printf("After more circular ops - Front: %d, Rear: %d\n", circularCq.Front(), circularCq.Rear())
	
	// Test 7: Concurrent access (basic test)
	fmt.Println("\n=== Testing Concurrent Access ===")
	concurrentCq := ConstructorMyCircularQueue(5)
	
	// Start goroutines to enqueue and dequeue
	done := make(chan bool)
	
	// Enqueue goroutine
	go func() {
		for i := 0; i < 10; i++ {
			concurrentCq.EnQueue(i)
			time.Sleep(time.Millisecond)
		}
		done <- true
	}()
	
	// Dequeue goroutine
	go func() {
		for i := 0; i < 10; i++ {
			concurrentCq.DeQueue()
			time.Sleep(time.Millisecond * 2)
		}
		done <- true
	}()
	
	<-done
	<-done
	
	fmt.Printf("Concurrent test completed - Size: %d, Front: %d\n", 
		5-len(concurrentCq.buffer), concurrentCq.Front())
}
