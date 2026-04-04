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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Circular Buffer with Head/Tail Pointers
- **Circular Buffer**: Fixed-size array with wrap-around indexing
- **Head/Tail Pointers**: Track front and rear positions
- **Modulo Arithmetic**: Use % capacity for circular indexing
- **Size Tracking**: Maintain current element count

## 2. PROBLEM CHARACTERISTICS
- **Fixed Capacity**: Queue has maximum size limit
- **FIFO Behavior**: First-in-first-out queue semantics
- **Circular Reuse**: Reuse array positions when elements are removed
- **O(1) Operations**: All operations in constant time

## 3. SIMILAR PROBLEMS
- Implement Queue using Stacks (LeetCode 232) - Queue with two stacks
- Design Double Ended Queue (LeetCode 641) - Deque with both ends
- Design Snake Game (LeetCode 353) - Circular game board
- Design Hit Counter (LeetCode 362) - Circular counter tracking

## 4. KEY OBSERVATIONS
- **Circular Indexing**: Head and tail wrap around using modulo
- **Size Management**: Track current number of elements
- **Empty/Full States**: Special cases when size = 0 or capacity
- **Thread Safety**: Need synchronization for concurrent access

## 5. VARIATIONS & EXTENSIONS
- **Thread Safety**: Add mutex locks for concurrent access
- **Channel Implementation**: Use Go channels for thread-safe queue
- **Statistics Tracking**: Track operation counts and patterns
- **Dynamic Resizing**: Support for growing/shrinking capacity

## 6. INTERVIEW INSIGHTS
- Always clarify: "Thread safety? Capacity limits? Overflow handling?"
- Edge cases: empty queue, full queue, single element
- Time complexity: O(1) for all operations
- Space complexity: O(k) where k=capacity
- Key insight: circular buffer enables O(1) FIFO with fixed array

## 7. COMMON MISTAKES
- Wrong modulo arithmetic causing index out of bounds
- Incorrect size tracking leading to overflow/underflow
- Not handling empty/full edge cases properly
- Race conditions in concurrent implementations
- Off-by-one errors in head/tail calculations

## 8. OPTIMIZATION STRATEGIES
- **Basic Circular Buffer**: O(1) time, O(k) space - standard
- **Thread-Safe Version**: O(1) time, O(k) space - with mutex
- **Channel-Based**: O(1) time, O(k) space - Go idiomatic
- **Statistics Tracking**: O(1) time, O(k) space - with counters

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like a circular parking lot with numbered spots:**
- You have a fixed number of parking spots in a circle (array)
- Cars enter at the entrance and park in the next available spot
- When a car leaves, the next car takes the first available spot
- Head pointer shows where next car should enter
- Tail pointer shows where the last car parked
- Like a circular conveyor belt that wraps around

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Queue with fixed capacity, enqueue/dequeue operations
2. **Goal**: Implement FIFO queue with O(1) operations
3. **Constraints**: Fixed capacity, circular behavior when full
4. **Output**: Efficient circular queue implementation

#### Phase 2: Key Insight Recognition
- **"Circular buffer natural fit"** → Fixed array with wrap-around indexing
- **"Head/tail pointers"** → Track insertion and removal positions
- **"Modulo arithmetic"** → Handle wrap-around with % capacity
- **"Size tracking"** → Need to know current element count

#### Phase 3: Strategy Development
```
Human thought process:
"I need O(1) enqueue/dequeue with fixed capacity.
Regular array would require shifting elements.

Circular Buffer Approach:
1. Fixed-size array as circular buffer
2. Head pointer: where to dequeue from
3. Tail pointer: where to enqueue to
4. Size counter: track current element count
5. Enqueue: add at tail, advance tail, increment size
6. Dequeue: remove from head, advance head, decrement size
7. Use modulo for circular indexing

This gives O(1) operations!"
```

#### Phase 4: Edge Case Handling
- **Empty queue**: Size = 0, head/tail don't matter
- **Full queue**: Size = capacity, reject enqueue
- **Single element**: Head = tail, size = 1
- **Wrap-around**: Handle when head/tail reach end of array

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Operations: Enqueue(1), Enqueue(2), Dequeue(), Enqueue(3), Front()

Human thinking:
"Circular Buffer Approach (capacity=3):
Initial: head=0, tail=-1, size=0, buffer=[_,_,_]

Enqueue(1): tail=( -1+1)%3=0, buffer[0]=1, size=1
State: head=0, tail=0, size=1, buffer=[1,_,_]

Enqueue(2): tail=(0+1)%3=1, buffer[1]=2, size=2
State: head=0, tail=1, size=2, buffer=[1,2,_]

Dequeue(): result=buffer[0]=1, head=(0+1)%3=1, size=1
State: head=1, tail=1, size=1, buffer=[1,2,_]

Enqueue(3): tail=(1+1)%3=2, buffer[2]=3, size=2
State: head=1, tail=2, size=2, buffer=[1,2,3]

Front(): return buffer[1]=2

All operations O(1) ✓"
```

#### Phase 6: Intuition Validation
- **Why circular works**: Reuses array positions efficiently
- **Why head/tail needed**: Track where to insert/remove
- **Why modulo works**: Handles wrap-around automatically
- **Why size tracking**: Distinguish empty/full states

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use regular array?"** → Dequeue would be O(N) shifting
2. **"Should I use linked list?"** → No O(1) random access
3. **"What about dynamic resizing?"** → Different problem specification
4. **"Can I optimize further?"** → O(1) is already optimal
5. **"What about thread safety?"** → Add locks or use channels

### Real-World Analogy
**Like a circular conveyor belt at a factory:**
- You have fixed number of positions on a circular belt (array)
- Items are placed on the belt at the loading point (tail)
- Items are removed from the belt at the unloading point (head)
- Belt continuously moves in a circle
- When belt is full, new items must wait
- When belt is empty, no items available to remove
- Like a circular buffer in embedded systems

### Human-Readable Pseudocode
```
class CircularQueue:
    buffer = array of fixed size k
    head = 0
    tail = -1
    size = 0
    
    function enqueue(val):
        if size == k:
            return false // queue is full
        tail = (tail + 1) % k
        buffer[tail] = val
        size++
        return true
    
    function dequeue():
        if size == 0:
            return error // queue is empty
        val = buffer[head]
        head = (head + 1) % k
        size--
        return val
    
    function front():
        if size == 0:
            return error
        return buffer[head]
    
    function isEmpty():
        return size == 0
    
    function isFull():
        return size == k
```

### Execution Visualization

### Example: Operations: Enqueue(1), Enqueue(2), Dequeue(), Enqueue(3), Front()
```
Circular Buffer Approach (capacity=3):
Initial: head=0, tail=-1, size=0, buffer=[_,_,_]

Enqueue(1): tail=0, buffer[0]=1, size=1
State: head=0, tail=0, size=1, buffer=[1,_,_]

Enqueue(2): tail=1, buffer[1]=2, size=2
State: head=0, tail=1, size=2, buffer=[1,2,_]

Dequeue(): result=buffer[0]=1, head=1, size=1
State: head=1, tail=1, size=1, buffer=[1,2,_]

Enqueue(3): tail=2, buffer[2]=3, size=2
State: head=1, tail=2, size=2, buffer=[1,2,3]

Front(): return buffer[1]=2

Final state: head=1, tail=2, size=2, buffer=[1,2,3]
All operations O(1) ✓
```

### Key Visualization Points:
- **Circular Buffer**: Fixed array with wrap-around indexing
- **Head Pointer**: Where to dequeue from
- **Tail Pointer**: Where to enqueue to
- **Size Tracking**: Current element count
- **Modulo Arithmetic**: Handle circular indexing

### Memory Layout Visualization:
```
State Evolution:
Initial: head=0, tail=-1, size=0, buffer=[_,_,_]

After Enqueue(1): head=0, tail=0, size=1, buffer=[1,_,_]
After Enqueue(2): head=0, tail=1, size=2, buffer=[1,2,_]
After Dequeue(): head=1, tail=1, size=1, buffer=[1,2,_]
After Enqueue(3): head=1, tail=2, size=2, buffer=[1,2,3]

Circular Behavior:
- When tail reaches capacity-1, next enqueue wraps to 0
- When head reaches capacity-1, next dequeue wraps to 0
- Size prevents enqueue when full
- Size prevents dequeue when empty

Operation Complexity:
Enqueue: O(1) time, O(1) space
Dequeue: O(1) time, O(1) space
Front: O(1) time, O(1) space
```

### Time Complexity Breakdown:
- **Enqueue**: O(1) time (modulo + array access), O(1) space
- **Dequeue**: O(1) time (modulo + array access), O(1) space
- **Front**: O(1) time (array access), O(1) space
- **Space**: O(k) where k=capacity

### Alternative Approaches:

#### 1. Two-Stack Implementation (O(1) amortized, O(N) space)
```go
type QueueTwoStacks struct {
    inStack  []int
    outStack []int
}
```
- **Pros**: Dynamic size, no capacity limit
- **Cons**: More complex, potential O(N) operations

#### 2. Linked List Implementation (O(1) enqueue, O(1) dequeue, O(N) space)
```go
type QueueLinkedList struct {
    head *Node
    tail *Node
    size  int
}
```
- **Pros**: Dynamic size, memory efficient
- **Cons**: No O(1) random access, cache unfriendly

#### 3. Channel-Based Implementation (O(1) time, O(k) space)
```go
type QueueChannel struct {
    buffer chan int
    capacity int
}
```
- **Pros**: Thread-safe by default, Go idiomatic
- **Cons**: Fixed capacity, blocking operations

### Extensions for Interviews:
- **Dynamic Resizing**: Support for growing/shrinking capacity
- **Thread Safety**: Mutex locks vs channels
- **Multiple Queues**: Support for priority queues
- **Blocking Operations**: Support for blocking enqueue/dequeue
- **Performance Analysis**: Discuss cache locality and memory usage
*/
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
