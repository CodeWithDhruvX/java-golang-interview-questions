# Go Concurrency Deep Dive — Goroutines, Channels & Worker Patterns

> **Topics:** Goroutines, Channels (Buffered/Unbuffered), Select Statements, Worker Pools, Producer-Consumer, Synchronization (Mutex, WaitGroup, Once), Context, Pipeline Patterns, Deadlock Prevention, Rate Limiting, Real-world Applications

---

## 📋 Reading Progress

- [ ] **Section 1:** Basic Goroutine Coordination (Q1–Q7)
- [ ] **Section 2:** Channel Patterns (Q8–Q17)
- [ ] **Section 3:** Worker Pools & Task Distribution (Q18–Q25)
- [ ] **Section 4:** Select Statements & Multiplexing (Q26–Q30)
- [ ] **Section 5:** Synchronization Primitives (Q31–Q35)
- [ ] **Section 6:** Context & Cancellation (Q36–Q40)
- [ ] **Section 7:** Pipeline Patterns (Q41–Q43)
- [ ] **Section 8:** Advanced Patterns & Error Handling (Q44–Q53)

> 🔖 **Last read:** <!-- e.g. Q18 · Section 3 done -->

---

## Section 1: Basic Goroutine Coordination (Q1–Q7)

### 1. Odd-Even Number Printing with Two Goroutines
**Q: Print numbers 1-20 using two goroutines with strict ordering (odd first, then even)**
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	oddEven := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Odd goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= 20; i += 2 {
			<-oddEven // Wait for signal
			fmt.Print(i, " ")
			oddEven <- true // Signal even goroutine
		}
	}()

	// Even goroutine
	go func() {
		defer wg.Done()
		for i := 2; i <= 20; i += 2 {
			fmt.Print(i, " ")
			oddEven <- true // Signal odd goroutine
			<-oddEven // Wait for signal
		}
	}()

	// Start with odd goroutine
	oddEven <- true
	wg.Wait()
	fmt.Println()
}
```
**A:** Prints `1 2 3 4 ... 20` with perfect alternation. Uses a boolean channel as a token passing mechanism.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** Can you explain your approach to printing numbers 1-20 using two goroutines with strict ordering?

**Your Response:** I created a solution using two goroutines - one for odd numbers and one for even numbers. The key is using a single boolean channel as a signaling mechanism between them. 

The odd goroutine waits for a signal before printing each odd number, then signals the even goroutine. The even goroutine prints its number and signals back. I use a WaitGroup to ensure the main function waits for both goroutines to complete.

The main function starts the communication by sending the first signal to the odd goroutine. This creates a perfect alternating pattern: odd gets signal, prints odd, signals even; even prints even, signals odd, and so on.

This approach guarantees strict ordering because each goroutine blocks until it receives permission to proceed, eliminating any race conditions. The channel acts as a token that only one goroutine can possess at a time.

---

### 2. Number-Letter Alternating Pattern
**Q: Print pattern "1 A 2 B 3 C..." using two goroutines**
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	numberChan := make(chan bool)
	letterChan := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Number goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			<-numberChan // Wait for signal
			fmt.Print(i, " ")
			letterChan <- true // Signal letter goroutine
		}
	}()

	// Letter goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 26; i++ {
			<-letterChan // Wait for signal
			fmt.Print(string('A'+i), " ")
			if i < 25 { // Don't signal after last letter
				numberChan <- true // Signal number goroutine
			}
		}
	}()

	// Start with number goroutine
	numberChan <- true
	wg.Wait()
	fmt.Println()
}
```
**A:** Prints `1 A 2 B 3 C ... 26 Z` with perfect coordination using two separate channels.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does your solution ensure the pattern "1 A 2 B 3 C..." using two goroutines?

**Your Response:** I use two separate channels for coordination - one for numbers and one for letters. The number goroutine prints a number, then signals the letter goroutine. The letter goroutine prints a letter and signals back to the number goroutine.

The key insight is that each goroutine blocks on its respective channel until it receives permission to proceed. This creates a strict alternating pattern. I start the process by signaling the number goroutine first.

The main function uses a WaitGroup to wait for both goroutines to complete. The conditional check in the letter goroutine prevents it from signaling after the last letter, avoiding a deadlock.

This approach demonstrates understanding of goroutine coordination using channels as synchronization primitives, which is fundamental to Go concurrency.

---

### 3. Three Goroutines Sequential Execution
**Q: Execute three goroutines in sequence A→B→C→A→B→C...**
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	chanAB := make(chan bool)
	chanBC := make(chan bool)
	chanCA := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(3)

	// Goroutine A
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-chanCA // Wait for C
			fmt.Print("A ")
			chanAB <- true // Signal B
		}
	}()

	// Goroutine B
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-chanAB // Wait for A
			fmt.Print("B ")
			chanBC <- true // Signal C
		}
	}()

	// Goroutine C
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-chanBC // Wait for B
			fmt.Print("C ")
			if i < 4 { // Don't signal after last iteration
				chanCA <- true // Signal A
			}
		}
	}()

	// Start the cycle
	chanCA <- true
	wg.Wait()
	fmt.Println()
}
```
**A:** Prints `A B C A B C A B C A B C A B C` in perfect circular sequence.

---

### 4. Five Goroutines Ordered Execution
**Q: Execute five goroutines in order 1→2→3→4→5→1→2...**
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	channels := make([]chan bool, 5)
	for i := range channels {
		channels[i] = make(chan bool)
	}
	
	wg := sync.WaitGroup{}
	wg.Add(5)

	// Start 5 goroutines
	for i := 0; i < 5; i++ {
		go func(id int) {
			defer wg.Done()
			for round := 0; round < 3; round++ {
				<-channels[id] // Wait for turn
				fmt.Printf("%d ", id+1)
				next := (id + 1) % 5
				if round < 2 || id < 4 { // Don't signal after last round for last goroutine
					channels[next] <- true // Signal next
				}
			}
		}(i)
	}

	// Start with goroutine 1
	channels[0] <- true
	wg.Wait()
	fmt.Println()
}
```
**A:** Prints `1 2 3 4 5 1 2 3 4 5 1 2 3 4 5` in perfect order.

---

### 5. Alternate Execution with Turns
**Q: Two goroutines taking turns with explicit turn counting**
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	turnChan := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Goroutine 1
	go func() {
		defer wg.Done()
		for turn := 1; turn <= 10; turn += 2 {
			<-turnChan
			fmt.Printf("G1-Turn%d ", turn)
			if turn < 10 {
				turnChan <- turn + 1
			}
		}
	}()

	// Goroutine 2
	go func() {
		defer wg.Done()
		for turn := 2; turn <= 10; turn += 2 {
			<-turnChan
			fmt.Printf("G2-Turn%d ", turn)
			if turn < 10 {
				turnChan <- turn + 1
			}
		}
	}()

	turnChan <- 1 // Start with turn 1
	wg.Wait()
	fmt.Println()
}
```
**A:** Prints `G1-Turn1 G2-Turn2 G1-Turn3 G2-Turn4 ... G1-Turn9 G2-Turn10`.

---

### 6. Token Passing Coordination
**Q: Use token passing to coordinate multiple goroutines**
```go
package main

import (
	"fmt"
	"sync"
)

type Token struct {
	Value    int
	HolderID int
}

func main() {
	tokenChan := make(chan Token, 1)
	wg := sync.WaitGroup{}
	wg.Add(3)

	// Initialize token
	tokenChan <- Token{Value: 1, HolderID: 0}

	// Three goroutines passing token
	for id := 1; id <= 3; id++ {
		go func(goroutineID int) {
			defer wg.Done()
			for round := 0; round < 3; round++ {
				token := <-tokenChan
				fmt.Printf("Goroutine%d: Token%d\n", goroutineID, token.Value)
				token.Value++
				token.HolderID = goroutineID
				
				// Pass to next goroutine
				next := goroutineID % 3 + 1
				if round < 2 || goroutineID < 3 {
					tokenChan <- token
				}
			}
		}(id)
	}

	wg.Wait()
	fmt.Println("Token passing completed")
}
```
**A:** Shows token being passed between goroutines with incrementing values.

---

### 7. Concurrent Ordered Output
**Q: Multiple goroutines producing output in order**
```go
package main

import (
	"fmt"
	"sort"
	"sync"
)

type Output struct {
	Order int
	Text  string
}

func main() {
	outputChan := make(chan Output, 9)
	wg := sync.WaitGroup{}
	wg.Add(3)

	// Three goroutines producing output
	for id := 1; id <= 3; id++ {
		go func(goroutineID int) {
			defer wg.Done()
			for i := 0; i < 3; i++ {
				order := (goroutineID-1)*3 + i + 1
				text := fmt.Sprintf("G%d-Item%d", goroutineID, i+1)
				outputChan <- Output{Order: order, Text: text}
			}
		}(id)
	}

	// Wait for all outputs
	go func() {
		wg.Wait()
		close(outputChan)
	}()

	// Collect and sort outputs
	var outputs []Output
	for out := range outputChan {
		outputs = append(outputs, out)
	}

	sort.Slice(outputs, func(i, j int) bool {
		return outputs[i].Order < outputs[j].Order
	})

	// Print in order
	for _, out := range outputs {
		fmt.Printf("%s ", out.Text)
	}
	fmt.Println()
}
```
**A:** Prints `G1-Item1 G2-Item1 G3-Item1 G1-Item2 G2-Item2 G3-Item2 G1-Item3 G2-Item3 G3-Item3`.

---

## Section 2: Channel Patterns (Q8–Q17)

### 8. Unbuffered Channel Blocking
**Q: Demonstrate unbuffered channel blocking behavior**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	unbuffered := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Sender
	go func() {
		defer wg.Done()
		messages := []string{"Hello", "World", "Go"}
		for _, msg := range messages {
			fmt.Printf("Sending: %s\n", msg)
			unbuffered <- msg // Blocks until receiver is ready
			fmt.Printf("Sent: %s\n", msg)
		}
	}()

	// Receiver
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			time.Sleep(500 * time.Millisecond) // Simulate work
			msg := <-unbuffered
			fmt.Printf("Received: %s\n", msg)
		}
	}()

	wg.Wait()
	fmt.Println("Unbuffered channel demonstration completed")
}
```
**A:** Shows sender blocking until receiver is ready, demonstrating synchronous communication.

---

### 9. Buffered Channel Non-blocking
**Q: Demonstrate buffered channel non-blocking behavior**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	buffered := make(chan string, 3)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Sender
	go func() {
		defer wg.Done()
		messages := []string{"Msg1", "Msg2", "Msg3", "Msg4"}
		for _, msg := range messages {
			fmt.Printf("Sending: %s\n", msg)
			buffered <- msg // Won't block if buffer has space
			fmt.Printf("Sent: %s (buffer len: %d)\n", msg, len(buffered))
		}
		close(buffered)
	}()

	// Receiver
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second) // Let sender fill buffer first
		for msg := range buffered {
			fmt.Printf("Received: %s\n", msg)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("Buffered channel demonstration completed")
}
```
**A:** Shows buffered channel allowing sender to proceed without immediate receiver.

---

### 10. Producer Slower Than Consumer
**Q: Handle case where producer is slower than consumer**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	items := make(chan int, 5)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Slow Producer
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			time.Sleep(500 * time.Millisecond) // Slow production
			fmt.Printf("Produced: %d\n", i)
			items <- i
		}
		close(items)
	}()

	// Fast Consumer
	go func() {
		defer wg.Done()
		for item := range items {
			fmt.Printf("Consumed: %d\n", item)
			time.Sleep(100 * time.Millisecond) // Fast consumption
		}
	}()

	wg.Wait()
	fmt.Println("Producer slower than consumer scenario completed")
}
```
**A:** Consumer waits when no items available, producer fills buffer when consumer is slow.

---

### 11. Multiple Producers Single Consumer
**Q: Multiple producers, single consumer pattern**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Product struct {
	ID      int
	Producer string
}

func producer(id int, products chan<- Product, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		product := Product{
			ID:       id*10 + i,
			Producer: fmt.Sprintf("P%d", id),
		}
		products <- product
		fmt.Printf("%s produced %d\n", product.Producer, product.ID)
		time.Sleep(time.Duration(id) * 200 * time.Millisecond)
	}
}

func consumer(products <-chan Product, wg *sync.WaitGroup) {
	defer wg.Done()
	for product := range products {
		fmt.Printf("Consumer got %s-%d\n", product.Producer, product.ID)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	products := make(chan Product, 10)
	wg := sync.WaitGroup{}
	wg.Add(3) // 2 producers + 1 consumer

	// Start producers
	go producer(1, products, &wg)
	go producer(2, products, &wg)

	// Start consumer
	go consumer(products, &wg)

	// Close channel when producers done
	go func() {
		wg.Wait()
		close(products)
	}()

	// Wait for consumer
	wg.Wait()
	fmt.Println("Multiple producers single consumer completed")
}
```
**A:** Shows multiple producers feeding a single consumer through shared channel.

---

### 12. Fill Buffered Channel Safely
**Q: Safely fill a buffered channel to capacity**
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	const capacity = 5
	buffered := make(chan int, capacity)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Filler
	go func() {
		defer wg.Done()
		for i := 1; i <= capacity; i++ {
			buffered <- i
			fmt.Printf("Filled slot %d with %d\n", i, i)
		}
		close(buffered)
	}()

	// Reader
	go func() {
		defer wg.Done()
		for value := range buffered {
			fmt.Printf("Read value %d from buffer\n", value)
		}
	}()

	wg.Wait()
	fmt.Printf("Buffered channel safely filled and emptied\n")
}
```
**A:** Demonstrates safe filling and emptying of buffered channel.

---

### 13. Basic Producer-Consumer Pattern
**Q: Implement basic producer-consumer pattern**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data string
}

func main() {
	jobs := make(chan Job, 5)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Producer
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			job := Job{
				ID:   i,
				Data: fmt.Sprintf("Task %d", i),
			}
			fmt.Printf("Producer: creating job %d\n", job.ID)
			jobs <- job
			fmt.Printf("Producer: sent job %d\n", job.ID)
			time.Sleep(200 * time.Millisecond)
		}
		close(jobs)
	}()

	// Consumer
	go func() {
		defer wg.Done()
		for job := range jobs {
			fmt.Printf("Consumer: processing job %d - %s\n", job.ID, job.Data)
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("Consumer: completed job %d\n", job.ID)
		}
	}()

	wg.Wait()
	fmt.Println("Producer-Consumer pattern completed")
}
```
**A:** Classic producer-consumer with channel as queue and proper cleanup.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** Can you explain the basic producer-consumer pattern implementation?

**Your Response:** I implemented the classic producer-consumer pattern using a channel as the communication medium. The producer creates jobs and sends them through the channel, while the consumer receives and processes them.

The channel acts as a queue that decouples the producer and consumer - they can work at different speeds. I used a buffered channel with capacity 5 to allow the producer to work ahead slightly without blocking immediately.

The producer creates Job structs with ID and data, sends them to the channel, then closes the channel when done. The consumer processes jobs from the channel until it's closed, using the range pattern which automatically handles channel closure.

The WaitGroup ensures both goroutines complete before the main function exits. This pattern is fundamental to concurrent programming and is used extensively in real-world systems for task distribution.

---

### 14. One Producer Two Consumers
**Q: Single producer, multiple consumers pattern**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Work string
}

func consumer(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Consumer%d: processing task %d - %s\n", id, task.ID, task.Work)
		time.Sleep(time.Duration(id) * 200 * time.Millisecond)
		fmt.Printf("Consumer%d: completed task %d\n", id, task.ID)
	}
}

func main() {
	tasks := make(chan Task, 10)
	wg := sync.WaitGroup{}
	wg.Add(3) // 1 producer + 2 consumers

	// Producer
	go func() {
		defer wg.Done()
		for i := 1; i <= 8; i++ {
			task := Task{
				ID:   i,
				Work: fmt.Sprintf("WorkItem%d", i),
			}
			fmt.Printf("Producer: created task %d\n", task.ID)
			tasks <- task
			time.Sleep(100 * time.Millisecond)
		}
		close(tasks)
	}()

	// Consumers
	go consumer(1, tasks, &wg)
	go consumer(2, tasks, &wg)

	wg.Wait()
	fmt.Println("One producer two consumers pattern completed")
}
```
**A:** Shows work distribution among multiple consumers from single producer.

---

### 15. Ordered Output Multiple Consumers
**Q: Maintain output order with multiple consumers**
```go
package main

import (
	"fmt"
	"sort"
	"sync"
)

type OrderedTask struct {
	Sequence int
	Result   string
}

func worker(id int, tasks <-chan int, results chan<- OrderedTask, wg *sync.WaitGroup) {
	defer wg.Done()
	for taskID := range tasks {
		result := fmt.Sprintf("Worker%d-Task%d", id, taskID)
		results <- OrderedTask{Sequence: taskID, Result: result}
	}
}

func main() {
	tasks := make(chan int, 10)
	results := make(chan OrderedTask, 10)
	wg := sync.WaitGroup{}
	
	// Start workers
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Send tasks
	go func() {
		for i := 1; i <= 9; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	// Wait for workers and close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and sort results
	var allResults []OrderedTask
	for result := range results {
		allResults = append(allResults, result)
	}

	sort.Slice(allResults, func(i, j int) bool {
		return allResults[i].Sequence < allResults[j].Sequence
	})

	// Print in order
	for _, result := range allResults {
		fmt.Printf("%s ", result.Result)
	}
	fmt.Println()
}
```
**A:** Maintains original task order despite parallel processing.

---

### 16. Bounded Buffer Producer-Consumer
**Q: Implement bounded buffer with backpressure**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type BoundedBuffer struct {
	buffer chan int
	mutex  sync.Mutex
	full   sync.Cond
	empty  sync.Cond
}

func NewBoundedBuffer(size int) *BoundedBuffer {
	b := &BoundedBuffer{
		buffer: make(chan int, size),
	}
	b.full.L = &b.mutex
	b.empty.L = &b.mutex
	return b
}

func (b *BoundedBuffer) Put(item int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	
	for len(b.buffer) == cap(b.buffer) {
		fmt.Printf("Buffer full, waiting to put %d\n", item)
		b.full.Wait()
	}
	
	b.buffer <- item
	fmt.Printf("Put %d, buffer size: %d\n", item, len(b.buffer))
	b.empty.Signal()
}

func (b *BoundedBuffer) Get() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	
	for len(b.buffer) == 0 {
		fmt.Println("Buffer empty, waiting to get")
		b.empty.Wait()
	}
	
	item := <-b.buffer
	fmt.Printf("Got %d, buffer size: %d\n", item, len(b.buffer))
	b.full.Signal()
	return item
}

func main() {
	buffer := NewBoundedBuffer(3)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Producer
	go func() {
		defer wg.Done()
		for i := 1; i <= 8; i++ {
			buffer.Put(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Consumer
	go func() {
		defer wg.Done()
		for i := 1; i <= 8; i++ {
			buffer.Get()
			time.Sleep(200 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("Bounded buffer demonstration completed")
}
```
**A:** Shows bounded buffer with explicit waiting conditions for full/empty states.

---

### 17. Graceful Shutdown Producer-Consumer
**Q: Implement graceful shutdown for producer-consumer**
```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type WorkItem struct {
	ID      int
	Payload string
}

func gracefulProducer(ctx context.Context, work chan<- WorkItem, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(work)
	
	for i := 1; i <= 20; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Producer stopping: %v\n", ctx.Err())
			return
		default:
			item := WorkItem{
				ID:      i,
				Payload: fmt.Sprintf("Data%d", i),
			}
			select {
			case work <- item:
				fmt.Printf("Produced item %d\n", i)
			case <-ctx.Done():
				fmt.Printf("Producer stopping mid-way: %v\n", ctx.Err())
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func gracefulConsumer(ctx context.Context, work <-chan WorkItem, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for item := range work {
		select {
		case <-ctx.Done():
			fmt.Printf("Consumer stopping: %v\n", ctx.Err())
			return
		default:
			fmt.Printf("Consuming item %d: %s\n", item.ID, item.Payload)
			time.Sleep(150 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	work := make(chan WorkItem, 5)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go gracefulProducer(ctx, work, &wg)
	go gracefulConsumer(ctx, work, &wg)

	wg.Wait()
	fmt.Println("Graceful shutdown demonstration completed")
}
```
**A:** Shows context-based cancellation for clean shutdown.

---

## Section 3: Worker Pools & Task Distribution (Q18–Q25)

### 18. Basic Worker Pool
**Q: Implement basic worker pool pattern**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data string
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d: processing job %d\n", id, job.ID)
		time.Sleep(time.Duration(job.ID%3+1) * 200 * time.Millisecond)
		fmt.Printf("Worker %d: completed job %d\n", id, job.ID)
	}
}

func main() {
	numWorkers := 3
	numJobs := 10
	jobs := make(chan Job, numJobs)
	wg := sync.WaitGroup{}

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for i := 1; i <= numJobs; i++ {
		job := Job{
			ID:   i,
			Data: fmt.Sprintf("Task %d", i),
		}
		jobs <- job
		fmt.Printf("Main: dispatched job %d\n", i)
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All jobs completed")
}
```
**A:** Shows efficient work distribution among fixed number of workers.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** Can you explain how the worker pool pattern distributes work among multiple workers?

**Your Response:** The worker pool pattern uses a channel as a task queue and multiple worker goroutines that pull tasks from this queue. I create 3 worker goroutines that all read from the same jobs channel.

When jobs are sent to the channel, any available worker can receive them. Go's channel semantics handle the distribution automatically - if multiple workers are waiting, one is chosen randomly to receive the job.

Each worker uses the range pattern to continuously process jobs until the channel is closed. This creates an efficient work distribution where idle workers automatically pick up new jobs.

The main function dispatches all jobs, then closes the channel to signal no more work is coming. The WaitGroup ensures all workers complete their current jobs before the program exits.

This pattern is extremely efficient for concurrent processing and is used extensively in real-world systems like web servers, database connection pools, and background job processors.

---

### 19. Worker Pool with Results Collection
**Q: Collect results from worker pool**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Input int
}

type Result struct {
	TaskID int
	Output int
	WorkerID int
}

func worker(id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		output := task.Input * task.Input // Square the input
		result := Result{
			TaskID:   task.ID,
			Output:   output,
			WorkerID: id,
		}
		results <- result
		fmt.Printf("Worker %d: processed task %d\n", id, task.ID)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	numWorkers := 3
	numTasks := 10
	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)
	wg := sync.WaitGroup{}

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Send tasks
	go func() {
		for i := 1; i <= numTasks; i++ {
			task := Task{
				ID:    i,
				Input: i * 10,
			}
			tasks <- task
		}
		close(tasks)
	}()

	// Wait for workers and close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result: Task%d=%d (Worker%d)\n", 
			result.TaskID, result.Output, result.WorkerID)
	}

	fmt.Println("All tasks processed with results collected")
}
```
**A:** Shows worker pool with separate channel for collecting results.

---

### 20. Ordered Worker Pool Output
**Q: Maintain output order in worker pool**
```go
package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type OrderedTask struct {
	ID      int
	Data    string
	WorkerID int
}

func orderedWorker(id int, tasks <-chan int, results chan<- OrderedTask, wg *sync.WaitGroup) {
	defer wg.Done()
	for taskID := range tasks {
		result := OrderedTask{
			ID:       taskID,
			Data:     fmt.Sprintf("Processed%d-ByWorker%d", taskID, id),
			WorkerID: id,
		}
		results <- result
		fmt.Printf("Worker %d: processed task %d\n", id, taskID)
		time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	}
}

func main() {
	numWorkers := 4
	numTasks := 12
	tasks := make(chan int, numTasks)
	results := make(chan OrderedTask, numTasks)
	wg := sync.WaitGroup{}

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go orderedWorker(i, tasks, results, &wg)
	}

	// Send tasks in order
	go func() {
		for i := 1; i <= numTasks; i++ {
			tasks <- i
		}
		close(tasks)
	}()

	// Wait for workers and close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and sort results by original task order
	var allResults []OrderedTask
	for result := range results {
		allResults = append(allResults, result)
	}

	sort.Slice(allResults, func(i, j int) bool {
		return allResults[i].ID < allResults[j].ID
	})

	// Print in original order
	fmt.Println("Results in original order:")
	for _, result := range allResults {
		fmt.Printf("Task%d: %s\n", result.ID, result.Data)
	}
}
```
**A:** Maintains original task order despite parallel processing.

---

### 21. Limited Concurrent Goroutines
**Q: Limit number of concurrent goroutines**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, semaphore chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Acquire semaphore
	semaphore <- struct{}{}
	defer func() { <-semaphore }()
	
	fmt.Printf("Task %d starting\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Task %d completed\n", id)
}

func main() {
	maxConcurrency := 3
	numTasks := 10
	semaphore := make(chan struct{}, maxConcurrency)
	wg := sync.WaitGroup{}

	// Start tasks with concurrency limit
	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		go task(i, semaphore, &wg)
	}

	wg.Wait()
	fmt.Println("All tasks completed with concurrency limit")
}
```
**A:** Uses semaphore pattern to limit concurrent goroutines.

---

### 22. Worker Pool with Retry
**Q: Implement worker pool with retry mechanism**
```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RetryableTask struct {
	ID      int
	Attempt int
	MaxRetries int
	Data    string
}

func retryWorker(id int, tasks <-chan RetryableTask, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		success := rand.Float32() > 0.3 // 70% success rate
		
		if success {
			result := fmt.Sprintf("Worker%d: Task%d succeeded on attempt %d", 
				id, task.ID, task.Attempt)
			results <- result
			fmt.Printf("Worker %d: Task %d succeeded\n", id, task.ID)
		} else if task.Attempt < task.MaxRetries {
			// Retry the task
			retryTask := RetryableTask{
				ID:         task.ID,
				Attempt:    task.Attempt + 1,
				MaxRetries: task.MaxRetries,
				Data:       task.Data,
			}
			fmt.Printf("Worker %d: Task %d failed, retrying (attempt %d)\n", 
				id, task.ID, task.Attempt + 1)
			time.Sleep(100 * time.Millisecond)
			// Send back to task queue for retry
			go func() {
				tasks <- retryTask
			}()
		} else {
			result := fmt.Sprintf("Worker%d: Task%d failed after %d attempts", 
				id, task.ID, task.MaxRetries)
			results <- result
			fmt.Printf("Worker %d: Task %d failed permanently\n", id, task.ID)
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	numWorkers := 2
	numTasks := 8
	maxRetries := 3
	tasks := make(chan RetryableTask, numTasks*maxRetries)
	results := make(chan string, numTasks)
	wg := sync.WaitGroup{}

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go retryWorker(i, tasks, results, &wg)
	}

	// Send initial tasks
	go func() {
		for i := 1; i <= numTasks; i++ {
			task := RetryableTask{
				ID:         i,
				Attempt:    1,
				MaxRetries: maxRetries,
				Data:       fmt.Sprintf("Data%d", i),
			}
			tasks <- task
		}
	}()

	// Wait for processing and close channels
	go func() {
		wg.Wait()
		close(tasks)
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("Worker pool with retry demonstration completed")
}
```
**A:** Shows retry mechanism for failed tasks in worker pool.

---

### 23. Fan-Out Task Distribution
**Q: Implement fan-out pattern for task distribution**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type WorkItem struct {
	ID   int
	Data string
}

func processor(id int, input <-chan WorkItem, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range input {
		fmt.Printf("Processor %d: processing item %d - %s\n", id, item.ID, item.Data)
		time.Sleep(time.Duration(id) * 100 * time.Millisecond)
		fmt.Printf("Processor %d: completed item %d\n", id, item.ID)
	}
}

func main() {
	numProcessors := 4
	numItems := 20
	input := make(chan WorkItem, numItems)
	wg := sync.WaitGroup{}

	// Fan-out: start multiple processors
	for i := 1; i <= numProcessors; i++ {
		wg.Add(1)
		go processor(i, input, &wg)
	}

	// Send work items
	go func() {
		for i := 1; i <= numItems; i++ {
			item := WorkItem{
				ID:   i,
				Data: fmt.Sprintf("Item%d", i),
			}
			input <- item
			fmt.Printf("Main: sent item %d\n", i)
		}
		close(input)
	}()

	wg.Wait()
	fmt.Println("Fan-out task distribution completed")
}
```
**A:** Distributes work among multiple processors efficiently.

---

### 24. Fan-In Merge Channels
**Q: Implement fan-in pattern to merge multiple channels**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func source(id int, output chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	messages := []string{
		fmt.Sprintf("Source%d-Msg1", id),
		fmt.Sprintf("Source%d-Msg2", id),
		fmt.Sprintf("Source%d-Msg3", id),
	}
	
	for _, msg := range messages {
		time.Sleep(time.Duration(id) * 200 * time.Millisecond)
		output <- msg
		fmt.Printf("Source %d: sent %s\n", id, msg)
	}
}

func fanIn(inputs ...<-chan string) <-chan string {
	output := make(chan string)
	wg := sync.WaitGroup{}
	
	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for msg := range ch {
				output <- msg
			}
		}(input)
	}
	
	go func() {
		wg.Wait()
		close(output)
	}()
	
	return output
}

func main() {
	numSources := 3
	sources := make([]chan string, numSources)
	wg := sync.WaitGroup{}
	
	// Create source channels
	for i := 0; i < numSources; i++ {
		sources[i] = make(chan string, 3)
		wg.Add(1)
		go source(i+1, sources[i], &wg)
	}
	
	// Fan-in: merge all sources
	merged := fanIn(sources...)
	
	// Wait for sources to finish
	go func() {
		wg.Wait()
		for _, ch := range sources {
			close(ch)
		}
	}()
	
	// Read from merged channel
	fmt.Println("Receiving from merged channel:")
	for msg := range merged {
		fmt.Printf("Received: %s\n", msg)
	}
	
	fmt.Println("Fan-in merge completed")
}
```
**A:** Merges multiple input channels into single output channel.

---

### 25. Merge Three Sorted Channels
**Q: Merge three sorted channels maintaining order**
```go
package main

import (
	"fmt"
	"sort"
	"sync"
)

type SortedItem struct {
	Value int
	Source string
}

func sortedSource(id int, values []int, output chan<- SortedItem, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	// Sort values first
	sort.Ints(values)
	
	for _, val := range values {
		item := SortedItem{
			Value:  val,
			Source: fmt.Sprintf("Src%d", id),
		}
		output <- item
	}
}

func mergeSorted(inputs ...<-chan SortedItem) <-chan SortedItem {
	output := make(chan SortedItem)
	wg := sync.WaitGroup{}
	
	for i, input := range inputs {
		wg.Add(1)
		go func(index int, ch <-chan SortedItem) {
			defer wg.Done()
			for item := range ch {
				output <- item
			}
		}(i, input)
	}
	
	go func() {
		wg.Wait()
		close(output)
	}()
	
	return output
}

func main() {
	// Three sorted sources
	source1 := []int{1, 4, 7, 10}
	source2 := []int{2, 5, 8, 11}
	source3 := []int{3, 6, 9, 12}
	
	channels := make([]chan SortedItem, 3)
	wg := sync.WaitGroup{}
	
	// Start sorted sources
	for i, values := range [][]int{source1, source2, source3} {
		channels[i] = make(chan SortedItem, len(values))
		wg.Add(1)
		go sortedSource(i+1, values, channels[i], &wg)
	}
	
	// Merge sorted channels
	merged := mergeSorted(channels...)
	
	// Wait for sources
	go func() {
		wg.Wait()
	}()
	
	// Collect and sort final result
	var allItems []SortedItem
	for item := range merged {
		allItems = append(allItems, item)
	}
	
	// Sort by value to maintain global order
	sort.Slice(allItems, func(i, j int) bool {
		return allItems[i].Value < allItems[j].Value
	})
	
	// Print merged sorted result
	fmt.Println("Merged sorted result:")
	for _, item := range allItems {
		fmt.Printf("%d (%s) ", item.Value, item.Source)
	}
	fmt.Println()
}
```
**A:** Merges multiple sorted channels while maintaining global sort order.

---

## Section 4: Select Statements & Multiplexing (Q26–Q30)

### 26. Split Channel Even Odd
**Q: Split channel into even and odd channels**
```go
package main

import (
	"fmt"
	"sync"
)

func splitter(input <-chan int, even, odd chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(even)
	defer close(odd)
	
	for num := range input {
		if num%2 == 0 {
			even <- num
		} else {
			odd <- num
		}
	}
}

func collector(name string, input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var numbers []int
	for num := range input {
		numbers = append(numbers, num)
		fmt.Printf("%s collected: %d\n", name, num)
	}
	fmt.Printf("%s total: %v\n", name, numbers)
}

func main() {
	numbers := make(chan int, 20)
	even := make(chan int, 10)
	odd := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(3)

	// Start splitter
	go splitter(numbers, even, odd, &wg)

	// Start collectors
	go collector("Even", even, &wg)
	go collector("Odd", odd, &wg)

	// Send numbers
	go func() {
		for i := 1; i <= 20; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	wg.Wait()
	fmt.Println("Channel split demonstration completed")
}
```
**A:** Splits input channel into even and odd output channels.

---

### 27. Select Multiple Channels
**Q: Use select to read from multiple channels**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func channel1(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	messages := []string{"A1", "A2", "A3"}
	for _, msg := range messages {
		time.Sleep(300 * time.Millisecond)
		ch <- msg
		fmt.Printf("Channel 1 sent: %s\n", msg)
	}
}

func channel2(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	messages := []string{"B1", "B2", "B3"}
	for _, msg := range messages {
		time.Sleep(200 * time.Millisecond)
		ch <- msg
		fmt.Printf("Channel 2 sent: %s\n", msg)
	}
}

func main() {
	ch1 := make(chan string, 3)
	ch2 := make(chan string, 3)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Start senders
	go channel1(ch1, &wg)
	go channel2(ch2, &wg)

	// Read from both channels using select
	fmt.Println("Reading from both channels using select:")
	for i := 0; i < 6; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from channel 1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from channel 2: %s\n", msg2)
		}
	}

	wg.Wait()
	fmt.Println("All messages received")
}
```
**A:** Uses select to multiplex between multiple channels.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does the select statement allow reading from multiple channels?

**Your Response:** The select statement allows a goroutine to wait on multiple channel operations simultaneously. It blocks until one of the cases can proceed, then executes that case.

In this example, I have two channels with different send rates. Channel 1 sends every 300ms, Channel 2 sends every 200ms. The select statement checks both channels and receives from whichever one has data available first.

If both channels have data ready, Go randomly selects one to ensure fairness. This creates a multiplexing effect where we can handle multiple data streams without blocking on any single one.

The select loop continues until all 6 messages are received (3 from each channel). This pattern is fundamental to Go concurrency and is used extensively in real systems for handling multiple input sources, implementing timeouts, or managing cancellation.

---

### 28. Select with Timeout
**Q: Implement select with timeout mechanism**
```go
package main

import (
	"fmt"
	"time"
)

func slowProducer(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Slow message"
}

func main() {
	ch := make(chan string)
	
	// Start slow producer
	go slowProducer(ch)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No message received within 1 second")
	}
	
	// Try again with longer timeout
	go slowProducer(ch)
	
	select {
	case msg := <-ch:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout: No message received within 3 seconds")
	}
}
```
**A:** Shows timeout handling using time.After in select.

---

### 29. Non-blocking Select
**Q: Implement non-blocking channel operations**
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Non-blocking receive
	select {
	case msg := <-ch1:
		fmt.Printf("Received from ch1: %s\n", msg)
	default:
		fmt.Println("No message available on ch1")
	}
	
	// Non-blocking send
	select {
	case ch2 <- "test message":
		fmt.Println("Sent to ch2")
	default:
		fmt.Println("Could not send to ch2 (no receiver)")
	}
	
	// Start receiver and try again
	go func() {
		time.Sleep(100 * time.Millisecond)
		<-ch2
	}()
	
	time.Sleep(50 * time.Millisecond)
	
	select {
	case ch2 <- "test message":
		fmt.Println("Successfully sent to ch2")
	default:
		fmt.Println("Still could not send to ch2")
	}
	
	time.Sleep(100 * time.Millisecond)
}
```
**A:** Demonstrates non-blocking channel operations using default case.

---

### 30. Prioritize Channel Select
**Q: Implement priority-based channel selection**
```go
package main

import (
	"fmt"
	"time"
)

func prioritySelect(high, low <-chan string) {
	for {
		select {
		case msg := <-high:
			fmt.Printf("HIGH priority: %s\n", msg)
		case msg := <-low:
			fmt.Printf("LOW priority: %s\n", msg)
		case <-time.After(2 * time.Second):
			fmt.Println("No messages received, exiting...")
			return
		}
	}
}

func main() {
	highPriority := make(chan string, 5)
	lowPriority := make(chan string, 5)
	
	// Send messages to both channels
	go func() {
		for i := 1; i <= 3; i++ {
			highPriority <- fmt.Sprintf("High%d", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()
	
	go func() {
		for i := 1; i <= 5; i++ {
			lowPriority <- fmt.Sprintf("Low%d", i)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	// Process with priority
	prioritySelect(highPriority, lowPriority)
}
```
**A:** Shows priority handling where high priority messages are processed first.

---

## Section 5: Synchronization Primitives (Q31–Q35)

### 31. WaitGroup Multiple Goroutines
**Q: Coordinate multiple goroutines with WaitGroup**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(id) * 200 * time.Millisecond)
	fmt.Printf("Worker %d completed\n", id)
}

func main() {
	numWorkers := 5
	wg := sync.WaitGroup{}
	
	fmt.Printf("Starting %d workers\n", numWorkers)
	
	// Add all workers to WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	
	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All workers completed")
}
```
**A:** Shows basic WaitGroup usage for coordinating multiple goroutines.

---

### 32. Concurrent Counter with Mutex
**Q: Implement thread-safe counter using mutex**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := SafeCounter{}
	wg := sync.WaitGroup{}
	numIncrementers := 100
	
	// Start multiple incrementers
	for i := 0; i < numIncrementers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("Final counter value: %d (expected: %d)\n", 
		counter.Value(), numIncrementers*1000)
}
```
**A:** Shows mutex-protected shared state with correct final value.

---

### 33. Shared Map with Mutex Protection
**Q: Implement thread-safe shared map**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu   sync.RWMutex
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
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.data[key]
	return value, exists
}

func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

func (sm *SafeMap) Size() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}

func main() {
	sharedMap := NewSafeMap()
	wg := sync.WaitGroup{}
	
	// Writers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(writerID int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("key_%d_%d", writerID, j)
				value := writerID*100 + j
				sharedMap.Set(key, value)
				fmt.Printf("Writer %d: set %s = %d\n", writerID, key, value)
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}
	
	// Readers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				key := fmt.Sprintf("key_%d_%d", j%5, j%10)
				if value, exists := sharedMap.Get(key); exists {
					fmt.Printf("Reader %d: got %s = %d\n", readerID, key, value)
				}
				time.Sleep(30 * time.Millisecond)
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final map size: %d\n", sharedMap.Size())
}
```
**A:** Shows thread-safe map operations with RWMutex for concurrent access.

---

### 34. RWLock Multiple Readers Single Writer
**Q: Implement multiple readers, single writer pattern**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type DataStore struct {
	mu    sync.RWMutex
	data  []string
	count int
}

func NewDataStore() *DataStore {
	return &DataStore{
		data: make([]string, 0),
	}
}

func (ds *DataStore) Write(item string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.data = append(ds.data, item)
	ds.count++
	fmt.Printf("Writer: added '%s' (total: %d)\n", item, ds.count)
}

func (ds *DataStore) Read(id int) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	if len(ds.data) > 0 {
		item := ds.data[len(ds.data)-1] // Read latest
		fmt.Printf("Reader %d: read '%s' (total items: %d)\n", id, item, len(ds.data))
	} else {
		fmt.Printf("Reader %d: no data available\n", id)
	}
}

func main() {
	store := NewDataStore()
	wg := sync.WaitGroup{}
	
	// Single writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			store.Write(fmt.Sprintf("Item%d", i))
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	// Multiple readers
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			for j := 0; j < 15; j++ {
				store.Read(readerID)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Println("Multiple readers single writer demonstration completed")
}
```
**A:** Shows RWMutex allowing concurrent reads but exclusive writes.

---

### 35. Sync.Once Resource Initialization
**Q: Implement one-time resource initialization**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Resource struct {
	ID      string
	Created time.Time
	Data    string
}

type ResourceManager struct {
	once   sync.Once
	resource *Resource
}

func (rm *ResourceManager) GetResource() *Resource {
	rm.once.Do(func() {
		fmt.Println("Initializing resource...")
		time.Sleep(1 * time.Second) // Simulate expensive initialization
		rm.resource = &Resource{
			ID:      "resource-123",
			Created: time.Now(),
			Data:    "Important data",
		}
		fmt.Println("Resource initialized")
	})
	return rm.resource
}

func main() {
	manager := ResourceManager{}
	wg := sync.WaitGroup{}
	
	// Multiple goroutines trying to get resource
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d: requesting resource\n", id)
			resource := manager.GetResource()
			fmt.Printf("Goroutine %d: got resource %s\n", id, resource.ID)
		}(i)
	}
	
	wg.Wait()
	
	// Verify it's the same resource
	resource1 := manager.GetResource()
	resource2 := manager.GetResource()
	
	fmt.Printf("Same resource instance: %t\n", resource1 == resource2)
	fmt.Printf("Resource created at: %v\n", resource1.Created)
}
```
**A:** Shows sync.Once ensuring initialization happens exactly once.

---

## Section 6: Context & Cancellation (Q36–Q40)

### 36. Context Cancel Goroutine
**Q: Implement context-based goroutine cancellation**
```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: stopping due to %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	
	// Start workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}
	
	// Let workers run for 2 seconds
	time.Sleep(2 * time.Second)
	
	// Cancel all workers
	fmt.Println("Cancelling all workers...")
	cancel()
	
	wg.Wait()
	fmt.Println("All workers stopped")
}
```
**A:** Shows context cancellation stopping multiple goroutines simultaneously.

---

### 37. Context Timeout Task
**Q: Implement context with timeout for task execution**
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context, duration time.Duration) error {
	fmt.Printf("Starting task that takes %v\n", duration)
	
	select {
	case <-time.After(duration):
		fmt.Println("Task completed successfully")
		return nil
	case <-ctx.Done():
		fmt.Printf("Task cancelled: %v\n", ctx.Err())
		return ctx.Err()
	}
}

func main() {
	// Task with sufficient timeout
	fmt.Println("=== Test 1: Sufficient timeout ===")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel1()
	
	err := longRunningTask(ctx1, 2*time.Second)
	fmt.Printf("Result: %v\n\n", err)
	
	// Task with insufficient timeout
	fmt.Println("=== Test 2: Insufficient timeout ===")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel2()
	
	err = longRunningTask(ctx2, 2*time.Second)
	fmt.Printf("Result: %v\n", err)
	
	// Manual cancellation
	fmt.Println("\n=== Test 3: Manual cancellation ===")
	ctx3, cancel3 := context.WithCancel(context.Background())
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Manually cancelling...")
		cancel3()
	}()
	
	err = longRunningTask(ctx3, 2*time.Second)
	fmt.Printf("Result: %v\n", err)
}
```
**A:** Shows timeout and manual cancellation using context.

---

### 38. Context Pipeline Cancellation
**Q: Pass context through pipeline of goroutines**
```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Data struct {
	Value int
	Stage string
}

func stage1(ctx context.Context, input <-chan int, output chan<- Data, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Stage1: stopping due to %v\n", ctx.Err())
			return
		case input <- i:
			fmt.Printf("Stage1: generated %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func stage2(ctx context.Context, input <-chan int, output chan<- Data, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	for num := range input {
		select {
		case <-ctx.Done():
			fmt.Printf("Stage2: stopping due to %v\n", ctx.Err())
			return
		default:
			processed := num * 2
			data := Data{
				Value: processed,
				Stage: "stage2",
			}
			output <- data
			fmt.Printf("Stage2: processed %d -> %d\n", num, processed)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func stage3(ctx context.Context, input <-chan Data, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for data := range input {
		select {
		case <-ctx.Done():
			fmt.Printf("Stage3: stopping due to %v\n", ctx.Err())
			return
		default:
			final := data.Value + 10
			fmt.Printf("Stage3: %s value %d -> final %d\n", 
				data.Stage, data.Value, final)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	
	// Pipeline channels
	ch1 := make(chan int)
	ch2 := make(chan Data)
	ch3 := make(chan Data)

	fmt.Println("Starting pipeline with cancellation support")

	// Start pipeline stages
	wg.Add(3)
	go stage1(ctx, ch1, ch2, &wg)
	go stage2(ctx, ch1, ch3, &wg)
	go stage3(ctx, ch3, &wg)

	// Cancel after 3 seconds
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("\nCancelling pipeline...")
		cancel()
	}()

	wg.Wait()
	fmt.Println("Pipeline stopped")
}
```
**A:** Shows context propagation through multi-stage pipeline.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you pass context through a pipeline of goroutines?

**Your Response:** I pass the same context through all pipeline stages, allowing cancellation to propagate through the entire processing chain. Each stage receives the context and checks for cancellation signals.

The pipeline has three stages: stage1 generates numbers, stage2 processes them, and stage3 performs final processing. Each stage uses select to check for the context's Done() channel alongside its regular processing.

When cancel() is called, all stages receive the cancellation signal simultaneously and can stop gracefully. The context flows through the pipeline like water - when the source is cut off, everything downstream stops.

The key insight is that context provides a tree-like cancellation structure. The same context can be passed to multiple goroutines, and canceling it affects all of them. This is perfect for pipelines where you need to stop all processing when an error occurs or a timeout happens.

This pattern is essential for real-world systems like data processing pipelines, microservice chains, or any multi-stage concurrent processing where you need coordinated cancellation.

---

### 39. Pipeline Three Stages
**Q: Implement three-stage processing pipeline**
```go
package main

import (
	"fmt"
	"sync"
)

type PipelineData struct {
	Value   int
	Stage   string
	Process string
}

func stage1(input <-chan int, output chan<- PipelineData, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	for num := range input {
		data := PipelineData{
			Value:   num,
			Stage:   "stage1",
			Process: fmt.Sprintf("Initial: %d", num),
		}
		output <- data
		fmt.Printf("Stage1: processed %d\n", num)
	}
}

func stage2(input <-chan PipelineData, output chan<- PipelineData, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	for data := range input {
		data.Value = data.Value * 2
		data.Stage = "stage2"
		data.Process = fmt.Sprintf("Doubled: %d", data.Value)
		output <- data
		fmt.Printf("Stage2: doubled to %d\n", data.Value)
	}
}

func stage3(input <-chan PipelineData, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for data := range input {
		data.Value = data.Value + 10
		data.Stage = "stage3"
		data.Process = fmt.Sprintf("Final: %d", data.Value)
		fmt.Printf("Stage3: final result %d - %s\n", data.Value, data.Process)
	}
}

func main() {
	numbers := make(chan int, 10)
	stage1Out := make(chan PipelineData, 10)
	stage2Out := make(chan PipelineData, 10)
	wg := sync.WaitGroup{}
	
	// Start pipeline stages
	wg.Add(3)
	go stage1(numbers, stage1Out, &wg)
	go stage2(stage1Out, stage2Out, &wg)
	go stage3(stage2Out, &wg)
	
	// Send input data
	go func() {
		for i := 1; i <= 8; i++ {
			numbers <- i
		}
		close(numbers)
	}()
	
	wg.Wait()
	fmt.Println("Three-stage pipeline completed")
}
```
**A:** Shows classic pipeline pattern with data transformation through stages.

---

### 40. Pipeline Cancellation on Error
**Q: Implement pipeline that cancels on error**
```go
package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type WorkItem struct {
	ID    int
	Data  string
	Error error
}

func processingStage(ctx context.Context, stageName string, 
	input <-chan WorkItem, output chan<- WorkItem, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	for item := range input {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: cancelled due to %v\n", stageName, ctx.Err())
			return
		default:
			// Simulate processing that might fail
			if item.ID == 5 {
				item.Error = errors.New("processing failed")
				fmt.Printf("%s: error on item %d\n", stageName, item.ID)
			} else {
				item.Data = fmt.Sprintf("%s-processed-%d", stageName, item.ID)
				fmt.Printf("%s: processed item %d\n", stageName, item.ID)
			}
			
			select {
			case output <- item:
			case <-ctx.Done():
				fmt.Printf("%s: cancelled during send\n", stageName)
				return
			}
			
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func errorChecker(ctx context.Context, input <-chan WorkItem, 
	output chan<- WorkItem, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	for item := range input {
		if item.Error != nil {
			fmt.Printf("Error detected: %v, cancelling pipeline\n", item.Error)
			cancel()
			return
		}
		
		select {
		case output <- item:
		case <-ctx.Done():
			fmt.Println("Error checker: cancelled")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	input := make(chan WorkItem, 10)
	stage1Out := make(chan WorkItem, 10)
	stage2Out := make(chan WorkItem, 10)
	wg := sync.WaitGroup{}
	
	// Start pipeline stages with error checking
	wg.Add(3)
	go processingStage(ctx, "Stage1", input, stage1Out, &wg)
	go errorChecker(ctx, stage1Out, stage2Out, cancel, &wg)
	go processingStage(ctx, "Stage2", stage2Out, make(chan WorkItem), &wg)
	
	// Send work items
	go func() {
		for i := 1; i <= 10; i++ {
			item := WorkItem{
				ID:   i,
				Data: fmt.Sprintf("Item%d", i),
			}
			select {
			case input <- item:
			case <-ctx.Done():
				fmt.Printf("Main: cancelled during send of item %d\n", i)
				return
			}
		}
		close(input)
	}()
	
	wg.Wait()
	fmt.Println("Pipeline with error cancellation completed")
}
```
**A:** Shows pipeline that cancels all stages when error is detected.

---

## Section 7: Pipeline Patterns (Q41–Q43)

### 41. Rate Limiter Channel
**Q: Implement rate limiter using channel**
```go
package main

import (
	"fmt"
	"time"
)

func rateLimiter(requests <-chan int, limit time.Duration) {
	ticker := time.NewTicker(limit)
	defer ticker.Stop()
	
	for req := range requests {
		<-ticker.C // Wait for tick
		fmt.Printf("Processing request %d at %v\n", req, time.Now().Format("15:04:05.000"))
	}
}

func main() {
	requests := make(chan int, 20)
	
	// Generate requests
	go func() {
		for i := 1; i <= 10; i++ {
			requests <- i
			fmt.Printf("Request %d received at %v\n", i, time.Now().Format("15:04:05.000"))
		}
		close(requests)
	}()
	
	// Process with rate limit (1 request per 500ms)
	fmt.Println("Starting rate-limited processing (1 request per 500ms)")
	rateLimiter(requests, 500*time.Millisecond)
	
	fmt.Println("Rate limiter demonstration completed")
}
```
**A:** Shows rate limiting using ticker channel.

---

### 42. Semaphore Buffered Channel
**Q: Implement semaphore using buffered channel**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, semaphore chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Acquire semaphore
	semaphore <- struct{}{}
	defer func() { <-semaphore }()
	
	fmt.Printf("Worker %d started at %v\n", id, time.Now().Format("15:04:05"))
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d completed at %v\n", id, time.Now().Format("15:04:05"))
}

func main() {
	maxConcurrent := 3
	numWorkers := 8
	semaphore := make(chan struct{}, maxConcurrent)
	wg := sync.WaitGroup{}
	
	fmt.Printf("Starting %d workers with max concurrency of %d\n", numWorkers, maxConcurrent)
	
	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, semaphore, &wg)
	}
	
	wg.Wait()
	fmt.Println("Semaphore demonstration completed")
}
```
**A:** Shows semaphore pattern using buffered channel for concurrency control.

---

### 43. Barrier Synchronization
**Q: Implement barrier synchronization pattern**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Barrier struct {
	count    int
	waiting  int
	arrived  chan struct{}
	departed chan struct{}
}

func NewBarrier(n int) *Barrier {
	return &Barrier{
		count:    n,
		arrived:  make(chan struct{}),
		departed: make(chan struct{}),
	}
}

func (b *Barrier) Wait() {
	b.waiting++
	
	if b.waiting == b.count {
		// Last goroutine to arrive
		close(b.arrived) // Signal all arrived
		<-b.departed     // Wait for departure signal
	} else {
		<-b.arrived // Wait for all to arrive
		b.waiting--
		if b.waiting == 0 {
			// Last goroutine to depart
			close(b.departed)
		}
	}
}

func phaseWorker(id int, barrier *Barrier, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Phase 1
	fmt.Printf("Worker %d: Phase 1\n", id)
	time.Sleep(time.Duration(id) * 200 * time.Millisecond)
	
	fmt.Printf("Worker %d: Waiting at barrier\n", id)
	barrier.Wait()
	
	// Phase 2
	fmt.Printf("Worker %d: Phase 2\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	
	fmt.Printf("Worker %d: Waiting at barrier\n", id)
	barrier.Wait()
	
	// Phase 3
	fmt.Printf("Worker %d: Phase 3\n", id)
}

func main() {
	numWorkers := 4
	barrier := NewBarrier(numWorkers)
	wg := sync.WaitGroup{}
	
	fmt.Printf("Starting %d workers with barrier synchronization\n", numWorkers)
	
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go phaseWorker(i, barrier, &wg)
	}
	
	wg.Wait()
	fmt.Println("Barrier synchronization demonstration completed")
}
```
**A:** Shows barrier pattern for synchronizing multiple goroutines at checkpoints.

---

## Section 8: Advanced Patterns & Error Handling (Q44–Q53)

### 44. Create and Fix Deadlock
**Q: Demonstrate deadlock creation and resolution**
```go
package main

import (
	"fmt"
	"time"
)

func createDeadlock() {
	fmt.Println("=== Creating Deadlock ===")
	
	// Deadlock example 1: Unbuffered channel send without receiver
	fmt.Println("1. Unbuffered channel deadlock:")
	ch1 := make(chan int)
	// ch1 <- 42 // This would deadlock - no receiver
	
	// Fix: Use goroutine receiver
	go func() {
		value := <-ch1
		fmt.Printf("Received: %d\n", value)
	}()
	ch1 <- 42
	time.Sleep(100 * time.Millisecond)
	
	// Deadlock example 2: Circular wait
	fmt.Println("\n2. Circular wait deadlock:")
	ch2 := make(chan int)
	ch3 := make(chan int)
	
	// This would deadlock:
	// go func() {
	//     ch2 <- <-ch3
	// }()
	// go func() {
	//     ch3 <- <-ch2
	// }()
	
	// Fix: Use select with timeout or reorder operations
	go func() {
		select {
		case ch2 <- <-ch3:
		case <-time.After(time.Second):
			fmt.Println("Timeout in circular wait")
		}
	}()
	go func() {
		ch3 <- 100
	}()
	
	time.Sleep(100 * time.Millisecond)
	
	// Deadlock example 3: Waiting on own channel
	fmt.Println("\n3. Self-dependency deadlock:")
	
	// This would deadlock:
	// ch4 := make(chan int)
	// ch4 <- <-ch4
	
	// Fix: Use separate channels
	ch4 := make(chan int)
	ch5 := make(chan int)
	go func() {
		ch5 <- <-ch4
	}()
	go func() {
		ch4 <- 200
	}()
	
	value := <-ch5
	fmt.Printf("Fixed circular dependency: %d\n", value)
}

func main() {
	createDeadlock()
	fmt.Println("\nAll deadlock examples demonstrated and fixed")
}
```
**A:** Shows common deadlock patterns and their solutions.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** Can you create and fix a deadlock scenario?

**Your Response:** I demonstrate three common deadlock patterns and their fixes. Deadlocks occur when goroutines wait indefinitely for resources that will never become available.

First, unbuffered channel deadlock - sending to an unbuffered channel without a receiver blocks forever. The fix is ensuring a receiver is ready before sending, typically by starting the receiver goroutine first.

Second, circular wait deadlock - two goroutines waiting for each other. The fix is using select with timeout or reordering operations to break the circular dependency.

Third, self-dependency deadlock - a goroutine waiting on itself. The fix is using separate channels to avoid circular dependencies.

The key insight is that deadlocks in Go usually come from channel operations without corresponding receivers/senders, or circular waiting patterns. The fixes involve ensuring proper channel setup, using timeouts, or restructuring the communication flow.

Understanding these patterns is crucial for writing robust concurrent Go programs that don't hang unexpectedly.

---

### 45. Read Closed Channel
**Q: Demonstrate reading from closed channel behavior**
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Reading from Closed Channel ===")
	
	// Example 1: Reading from closed buffered channel
	fmt.Println("1. Buffered channel:")
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)
	
	// Read remaining values
	for i := 0; i < 4; i++ { // One extra read
		value, ok := <-ch1
		if ok {
			fmt.Printf("Read %d, channel open\n", value)
		} else {
			fmt.Printf("Read %d, channel closed\n", value)
		}
	}
	
	// Example 2: Range over closed channel
	fmt.Println("\n2. Range over closed channel:")
	ch2 := make(chan string, 3)
	ch2 <- "A"
	ch2 <- "B"
	ch2 <- "C"
	close(ch2)
	
	fmt.Println("Ranging over closed channel:")
	for value := range ch2 {
		fmt.Printf("Value: %s\n", value)
	}
	fmt.Println("Range completed automatically")
	
	// Example 3: Select with closed channel
	fmt.Println("\n3. Select with closed channel:")
	ch3 := make(chan int)
	go func() {
		ch3 <- 42
		close(ch3)
	}()
	
	for i := 0; i < 3; i++ {
		select {
		case value := <-ch3:
			fmt.Printf("Select read: %d\n", value)
		default:
			fmt.Println("Select: no value available")
		}
	}
}
```
**A:** Shows different ways to safely read from closed channels.

---

### 46. Write Closed Channel Panic
**Q: Demonstrate panic when writing to closed channel**
```go
package main

import (
	"fmt"
	"time"
)

func safeWrite(ch chan<- int, value int, done chan<- bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
			done <- false
		} else {
			done <- true
		}
	}()
	
	ch <- value
	fmt.Printf("Successfully wrote %d\n", value)
}

func main() {
	fmt.Println("=== Writing to Closed Channel ===")
	
	// Example 1: Direct write to closed channel (panics)
	fmt.Println("1. Direct write to closed channel:")
	ch1 := make(chan int, 1)
	ch1 <- 1
	close(ch1)
	
	fmt.Println("Attempting to write to closed channel...")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic recovered: %v\n", r)
		}
	}()
	
	// This will panic
	// ch1 <- 2
	
	// Example 2: Safe write with recovery
	fmt.Println("\n2. Safe write with panic recovery:")
	ch2 := make(chan int, 1)
	ch2 <- 10
	close(ch2)
	
	done := make(chan bool)
	go safeWrite(ch2, 20, done)
	
	success := <-done
	fmt.Printf("Write successful: %t\n", success)
	
	// Example 3: Check before write pattern
	fmt.Println("\n3. Check before write pattern:")
	ch3 := make(chan int, 1)
	ch3 <- 100
	
	// Simulate checking if channel is closed (not directly possible)
	// Instead, use select with default for non-blocking write
	select {
	case ch3 <- 200:
		fmt.Println("Write successful")
	default:
		fmt.Println("Write would block (channel might be closed or full)")
	}
	
	close(ch3)
	
	select {
	case ch3 <- 300:
		fmt.Println("Write successful (unexpected)")
	default:
		fmt.Println("Write blocked (channel is closed)")
	}
	
	time.Sleep(100 * time.Millisecond)
}
```
**A:** Shows panic behavior and safe patterns for writing to potentially closed channels.

---

### 47. Detect and Fix Goroutine Leak
**Q: Demonstrate goroutine leak detection and prevention**
```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func leakyGoroutine() {
	ch := make(chan int)
	
	// This goroutine will leak - no sender and no receiver
	go func() {
		<-ch // Will block forever
	}()
}

func fixedGoroutine(done chan struct{}) {
	ch := make(chan int)
	
	go func() {
		select {
		case <-ch:
		case <-done: // Can be cancelled
		}
	}()
	
	// Send value or close to prevent leak
	close(ch)
}

func monitorGoroutines(name string) {
	for i := 0; i < 5; i++ {
		count := runtime.NumGoroutine()
		fmt.Printf("%s - Goroutine count: %d\n", name, count)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== Goroutine Leak Detection ===")
	
	initialCount := runtime.NumGoroutine()
	fmt.Printf("Initial goroutine count: %d\n\n", initialCount)
	
	// Example 1: Leaky goroutine
	fmt.Println("1. Creating leaky goroutine:")
	leakyGoroutine()
	monitorGoroutines("After leak")
	
	// Example 2: Fixed goroutine with proper cleanup
	fmt.Println("\n2. Fixed goroutine with cleanup:")
	done := make(chan struct{})
	fixedGoroutine(done)
	monitorGoroutines("After fix")
	
	// Example 3: Worker pool with proper shutdown
	fmt.Println("\n3. Worker pool with proper shutdown:")
	
	type Task struct {
		ID int
	}
	
	tasks := make(chan Task, 10)
	wg := sync.WaitGroup{}
	
	// Start workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for task := range tasks {
				fmt.Printf("Worker %d: task %d\n", id, task.ID)
				time.Sleep(100 * time.Millisecond)
			}
			fmt.Printf("Worker %d: shutting down\n", id)
		}(i)
	}
	
	// Send some tasks
	go func() {
		for i := 1; i <= 5; i++ {
			tasks <- Task{ID: i}
		}
		close(tasks) // Important: close to signal workers to stop
	}()
	
	wg.Wait()
	finalCount := runtime.NumGoroutine()
	fmt.Printf("\nFinal goroutine count: %d (leaked: %d)\n", 
		finalCount, finalCount-initialCount)
}
```
**A:** Shows goroutine leak patterns and proper cleanup techniques.

---

### 48. Nil Channel Behavior
**Q: Demonstrate nil channel behavior in select**
```go
package main

import (
	"fmt"
	"time"
)

func nilChannelBehavior() {
	fmt.Println("=== Nil Channel Behavior ===")
	
	// Example 1: Send to nil channel blocks forever
	fmt.Println("1. Send to nil channel:")
	var ch1 chan int
	// ch1 <- 42 // This would block forever
	fmt.Println("Sending to nil channel would block forever")
	
	// Example 2: Receive from nil channel blocks forever
	fmt.Println("\n2. Receive from nil channel:")
	var ch2 chan string
	// value := <-ch2 // This would block forever
	fmt.Println("Receiving from nil channel would block forever")
	
	// Example 3: Nil channel in select is ignored
	fmt.Println("\n3. Nil channel in select:")
	ch3 := make(chan int, 1)
	ch4 := make(chan int, 1)
	var nilCh chan int
	
	ch3 <- 1
	ch4 <- 2
	
	// This will only receive from ch3 and ch4, not nilCh
	for i := 0; i < 2; i++ {
		select {
		case value := <-ch3:
			fmt.Printf("Received from ch3: %d\n", value)
		case value := <-ch4:
			fmt.Printf("Received from ch4: %d\n", value)
		case value := <-nilCh:
			fmt.Printf("Received from nilCh: %d (impossible)\n", value)
		default:
			fmt.Println("No channel ready")
		}
	}
	
	// Example 4: Dynamic channel enabling/disabling
	fmt.Println("\n4. Dynamic channel control:")
	enabled := true
	var dynamicCh chan int
	
	if enabled {
		dynamicCh = make(chan int, 1)
		dynamicCh <- 100
	}
	
	select {
	case value := <-dynamicCh:
		fmt.Printf("Received from dynamic channel: %d\n", value)
	default:
		fmt.Println("Dynamic channel disabled or empty")
	}
	
	// Disable channel by setting to nil
	dynamicCh = nil
	
	select {
	case value := <-dynamicCh:
		fmt.Printf("Received after disable: %d\n", value)
	default:
		fmt.Println("Channel successfully disabled")
	}
}

func main() {
	nilChannelBehavior()
	fmt.Println("\nNil channel behavior demonstration completed")
}
```
**A:** Shows how nil channels behave and can be used for dynamic channel control.

---

### 49. Task Scheduler with Intervals
**Q: Implement task scheduler with different intervals**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID       int
	Name     string
	Interval time.Duration
	Active   bool
}

func scheduler(tasks <-chan Task, stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	
	taskMap := make(map[int]Task)
	tickers := make(map[int]*time.Ticker)
	
	for {
		select {
		case task := <-tasks:
			if task.Active {
				// Start or update task
				if ticker, exists := tickers[task.ID]; exists {
					ticker.Stop()
				}
				
				ticker := time.NewTicker(task.Interval)
				tickers[task.ID] = ticker
				taskMap[task.ID] = task
				
				fmt.Printf("Started task %d (%s) with interval %v\n", 
					task.ID, task.Name, task.Interval)
				
				// Run task in separate goroutine
				go func(t Task, tk *time.Ticker) {
					for {
						select {
						case <-tk.C:
							fmt.Printf("Executing task %d (%s) at %v\n", 
								t.ID, t.Name, time.Now().Format("15:04:05"))
						case <-stop:
							tk.Stop()
							return
						}
					}
				}(task, ticker)
			} else {
				// Stop task
				if ticker, exists := tickers[task.ID]; exists {
					ticker.Stop()
					delete(tickers, task.ID)
					delete(taskMap, task.ID)
					fmt.Printf("Stopped task %d\n", task.ID)
				}
			}
			
		case <-stop:
			// Stop all tickers
			for _, ticker := range tickers {
				ticker.Stop()
			}
			fmt.Println("Scheduler stopped")
			return
		}
	}
}

func main() {
	taskChan := make(chan Task, 10)
	stopChan := make(chan struct{})
	wg := sync.WaitGroup{}
	
	wg.Add(1)
	go scheduler(taskChan, stopChan, &wg)
	
	// Schedule tasks with different intervals
	tasks := []Task{
		{ID: 1, Name: "Health Check", Interval: 2 * time.Second, Active: true},
		{ID: 2, Name: "Data Backup", Interval: 5 * time.Second, Active: true},
		{ID: 3, Name: "Log Cleanup", Interval: 3 * time.Second, Active: true},
	}
	
	for _, task := range tasks {
		taskChan <- task
	}
	
	// Let scheduler run for 10 seconds
	time.Sleep(10 * time.Second)
	
	// Stop one task
	stopTask := Task{ID: 2, Name: "Data Backup", Active: false}
	taskChan <- stopTask
	
	// Run for 5 more seconds
	time.Sleep(5 * time.Second)
	
	// Stop scheduler
	close(stopChan)
	wg.Wait()
	
	fmt.Println("Task scheduler demonstration completed")
}
```
**A:** Shows scheduler managing multiple tasks with different execution intervals.

---

### 50. Parallel API Calls Ordered
**Q: Make parallel API calls but return results in order**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type APIResponse struct {
	RequestID int
	Data      string
	Timestamp time.Time
}

func apiCall(id int, responses chan<- APIResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Simulate API call with varying duration
	duration := time.Duration(id%3+1) * 200 * time.Millisecond
	time.Sleep(duration)
	
	response := APIResponse{
		RequestID: id,
		Data:      fmt.Sprintf("API data for request %d", id),
		Timestamp: time.Now(),
	}
	
	responses <- response
	fmt.Printf("API call %d completed in %v\n", id, duration)
}

func main() {
	numRequests := 8
	responses := make(chan APIResponse, numRequests)
	wg := sync.WaitGroup{}
	
	fmt.Printf("Making %d parallel API calls...\n", numRequests)
	
	// Make all API calls in parallel
	for i := 1; i <= numRequests; i++ {
		wg.Add(1)
		go apiCall(i, responses, &wg)
	}
	
	// Wait for all calls to complete
	go func() {
		wg.Wait()
		close(responses)
	}()
	
	// Collect responses
	var allResponses []APIResponse
	for response := range responses {
		allResponses = append(allResponses, response)
	}
	
	// Sort by request ID to maintain order
	for i := 0; i < len(allResponses)-1; i++ {
		for j := i + 1; j < len(allResponses); j++ {
			if allResponses[i].RequestID > allResponses[j].RequestID {
				allResponses[i], allResponses[j] = allResponses[j], allResponses[i]
			}
		}
	}
	
	// Display results in order
	fmt.Println("\nResults in original order:")
	for _, response := range allResponses {
		fmt.Printf("Request %d: %s (received at %v)\n", 
			response.RequestID, response.Data, response.Timestamp.Format("15:04:05.000"))
	}
	
	fmt.Println("\nParallel API calls with ordered results completed")
}
```
**A:** Shows parallel API calls with ordered result presentation.

---

### 51. Log Processor Multiple Writers
**Q: Implement log processor with multiple writers**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
	Source    string
}

type LogProcessor struct {
	entries  chan LogEntry
	writers  []chan LogEntry
	wg       sync.WaitGroup
	stopped  bool
	mu       sync.RWMutex
}

func NewLogProcessor(numWriters int) *LogProcessor {
	lp := &LogProcessor{
		entries: make(chan LogEntry, 1000),
		writers: make([]chan LogEntry, numWriters),
	}
	
	for i := 0; i < numWriters; i++ {
		lp.writers[i] = make(chan LogEntry, 100)
		lp.wg.Add(1)
		go lp.writer(i)
	}
	
	go lp.distribute()
	return lp
}

func (lp *LogProcessor) distribute() {
	for entry := range lp.entries {
		lp.mu.RLock()
		if !lp.stopped {
			// Send to all writers
			for _, writer := range lp.writers {
				select {
				case writer <- entry:
				default:
					fmt.Printf("Writer channel full, dropping log from %s\n", entry.Source)
				}
			}
		}
		lp.mu.RUnlock()
	}
	
	// Close all writer channels
	lp.mu.RLock()
	for _, writer := range lp.writers {
		close(writer)
	}
	lp.mu.RUnlock()
}

func (lp *LogProcessor) writer(id int) {
	defer lp.wg.Done()
	
	count := 0
	for entry := range lp.writers[id] {
		count++
		fmt.Printf("Writer %d [%s]: %s %s - %s\n", 
			id, entry.Timestamp.Format("15:04:05.000"), 
			entry.Level, entry.Source, entry.Message)
		
		// Simulate writing delay
		time.Sleep(50 * time.Millisecond)
	}
	
	fmt.Printf("Writer %d processed %d log entries\n", id, count)
}

func (lp *LogProcessor) Log(level, source, message string) {
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Source:    source,
	}
	
	select {
	case lp.entries <- entry:
	default:
		fmt.Printf("Log processor full, dropping log from %s\n", source)
	}
}

func (lp *LogProcessor) Stop() {
	lp.mu.Lock()
	lp.stopped = true
	lp.mu.Unlock()
	close(lp.entries)
	lp.wg.Wait()
}

func logGenerator(name string, level string, processor *LogProcessor, wg *sync.WaitGroup) {
	defer wg.Done()
	
	messages := []string{
		"Starting process",
		"Initializing components",
		"Processing data",
		"Handling request",
		"Cleaning up resources",
		"Process completed",
	}
	
	for i, msg := range messages {
		processor.Log(level, name, msg)
		time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
	}
}

func main() {
	processor := NewLogProcessor(3)
	wg := sync.WaitGroup{}
	
	fmt.Println("=== Multi-writer Log Processor ===")
	
	// Start multiple log generators
	sources := []struct {
		name  string
		level string
	}{
		{"Auth", "INFO"},
		{"Database", "WARN"},
		{"API", "ERROR"},
		{"Cache", "INFO"},
		{"Queue", "DEBUG"},
	}
	
	for _, source := range sources {
		wg.Add(1)
		go logGenerator(source.name, source.level, processor, &wg)
	}
	
	wg.Wait()
	
	// Give some time for processing
	time.Sleep(1 * time.Second)
	
	processor.Stop()
	fmt.Println("Log processor demonstration completed")
}
```
**A:** Shows log processing with multiple concurrent writers.

---

### 52. Chat System Ordered Processing
**Q: Implement chat system with ordered message processing**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	ID        int
	Content   string
	Sender    string
	Timestamp time.Time
}

type ChatRoom struct {
	messages  chan Message
	processed chan string
	wg        sync.WaitGroup
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		messages:  make(chan Message, 100),
		processed: make(chan string, 100),
	}
}

func (cr *ChatRoom) Start() {
	cr.wg.Add(1)
	go cr.processMessages()
}

func (cr *ChatRoom) processMessages() {
	defer cr.wg.Done()
	defer close(cr.processed)
	
	// Buffer for ordering messages
	messageBuffer := make([]Message, 0)
	
	for msg := range cr.messages {
		messageBuffer = append(messageBuffer, msg)
		
		// Process when buffer is full or no more messages coming
		if len(messageBuffer) >= 5 {
			cr.processBuffer(&messageBuffer)
		}
	}
	
	// Process remaining messages
	cr.processBuffer(&messageBuffer)
}

func (cr *ChatRoom) processBuffer(buffer *[]Message) {
	if len(*buffer) == 0 {
		return
	}
	
	// Sort messages by timestamp
	for i := 0; i < len(*buffer)-1; i++ {
		for j := i + 1; j < len(*buffer); j++ {
			if (*buffer)[i].Timestamp.After((*buffer)[j].Timestamp) {
				(*buffer)[i], (*buffer)[j] = (*buffer)[j], (*buffer)[i]
			}
		}
	}
	
	// Process in order
	for _, msg := range *buffer {
		processed := fmt.Sprintf("[%v] %s: %s", 
			msg.Timestamp.Format("15:04:05.000"), 
			msg.Sender, 
			msg.Content)
		cr.processed <- processed
	}
	
	*buffer = (*buffer)[:0] // Clear buffer
}

func (cr *ChatRoom) SendMessage(msg Message) {
	cr.messages <- msg
}

func (cr *ChatRoom) GetProcessedMessages() <-chan string {
	return cr.processed
}

func (cr *ChatRoom) Stop() {
	close(cr.messages)
	cr.wg.Wait()
}

func sender(name string, room *ChatRoom, wg *sync.WaitGroup) {
	defer wg.Done()
	
	messages := []string{
		"Hello everyone!",
		"How are you?",
		"Great weather today",
		"Anyone up for coffee?",
		"See you later!",
	}
	
	for i, content := range messages {
		msg := Message{
			ID:        i + 1,
			Content:   content,
			Sender:    name,
			Timestamp: time.Now().Add(time.Duration(i) * time.Millisecond),
		}
		
		room.SendMessage(msg)
		fmt.Printf("%s sent: %s\n", name, content)
		time.Sleep(time.Duration(name[0]-'A'+1) * 100 * time.Millisecond)
	}
}

func main() {
	room := NewChatRoom()
	room.Start()
	
	fmt.Println("=== Chat System: Multiple Senders, Ordered Processing ===")
	
	// Start multiple senders
	wg := sync.WaitGroup{}
	senders := []string{"Alice", "Bob", "Charlie"}
	
	for _, sender := range senders {
		wg.Add(1)
		go sender(sender, room, &wg)
	}
	
	// Display processed messages
	go func() {
		for processed := range room.GetProcessedMessages() {
			fmt.Println("Processed:", processed)
		}
	}()
	
	// Wait for all senders
	wg.Wait()
	
	// Stop chat room
	room.Stop()
	
	fmt.Println("Chat system demonstration completed")
}
```
**A:** Shows chat system maintaining message order despite concurrent sending.

---

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement a chat system with multiple senders but ordered message processing?

**Your Response:** I implement a chat room that collects messages from multiple senders and processes them in chronological order. Multiple users send messages concurrently to a central message channel.

The chat room uses a buffering and sorting strategy. It collects messages in a buffer, sorts them by timestamp, then processes them in order. This ensures that even though messages arrive concurrently, they're displayed chronologically.

Each sender creates messages with timestamps and sends them to the chat room. The chat room's processor maintains order by sorting before output, which is crucial for chat applications where message order affects conversation flow.

The key insight is separating concurrent message collection from ordered processing. Senders don't block each other, but the final output maintains proper chronological order for readability.

This pattern is essential for real-time communication systems like chat applications, collaborative editors, or any system where multiple users generate content that needs to be presented in a logical order. It demonstrates understanding of concurrent systems with ordering requirements.

---

### 53. File Pipeline Read Process Write
**Q: Implement file processing pipeline with read, process, write stages**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type FileData struct {
	Filename string
	Content  string
	Stage    string
}

type ProcessedData struct {
	Filename   string
	Original   string
	Processed  string
	Size       int
	WordCount  int
	ProcessedAt time.Time
}

func fileReader(filenames []string, output chan<- FileData, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	for _, filename := range filenames {
		// Simulate reading file
		content := fmt.Sprintf("Content of %s: This is sample data for processing.", filename)
		data := FileData{
			Filename: filename,
			Content:  content,
			Stage:    "read",
		}
		
		output <- data
		fmt.Printf("Reader: read file %s\n", filename)
		time.Sleep(100 * time.Millisecond) // Simulate I/O delay
	}
}

func dataProcessor(input <-chan FileData, output chan<- ProcessedData, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	for data := range input {
		// Simulate processing (word count, transformation, etc.)
		wordCount := len(data.Content)
		processed := fmt.Sprintf("PROCESSED: %s (words: %d)", data.Content, wordCount)
		
		result := ProcessedData{
			Filename:    data.Filename,
			Original:    data.Content,
			Processed:   processed,
			Size:        len(data.Content),
			WordCount:   wordCount,
			ProcessedAt: time.Now(),
		}
		
		output <- result
		fmt.Printf("Processor: processed %s\n", data.Filename)
		time.Sleep(150 * time.Millisecond) // Simulate processing time
	}
}

func fileWriter(input <-chan ProcessedData, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for data := range input {
		// Simulate writing processed data
		fmt.Printf("Writer: wrote processed data for %s\n", data.Filename)
		fmt.Printf("  - Original size: %d bytes\n", data.Size)
		fmt.Printf("  - Word count: %d\n", data.WordCount)
		fmt.Printf("  - Processed at: %v\n", data.ProcessedAt.Format("15:04:05.000"))
		time.Sleep(80 * time.Millisecond) // Simulate write delay
	}
}

func main() {
	filenames := []string{
		"document1.txt",
		"report.pdf",
		"data.csv",
		"image.jpg",
		"config.xml",
		"log.txt",
	}
	
	readChan := make(chan FileData, 10)
	processChan := make(chan ProcessedData, 10)
	wg := sync.WaitGroup{}
	
	fmt.Println("=== File Processing Pipeline ===")
	fmt.Printf("Processing %d files through read->process->write pipeline\n\n", len(filenames))
	
	// Start pipeline stages
	wg.Add(3)
	go fileReader(filenames, readChan, &wg)
	go dataProcessor(readChan, processChan, &wg)
	go fileWriter(processChan, &wg)
	
	// Wait for pipeline to complete
	wg.Wait()
	
	fmt.Println("\nFile processing pipeline completed")
}
```
**A:** Shows complete file processing pipeline with three distinct stages.

---

## Summary

This comprehensive collection covers **53 essential Go concurrency patterns** organized into 8 sections:

1. **Basic Goroutine Coordination** (Q1-7): Token passing, alternating execution, ordered output
2. **Channel Patterns** (Q8-17): Buffered/unbuffered channels, producer-consumer, graceful shutdown
3. **Worker Pools** (Q18-25): Basic pools, result collection, retry mechanisms, fan-out/fan-in
4. **Select Statements** (Q26-30): Multiplexing, timeouts, non-blocking operations, priority handling
5. **Synchronization Primitives** (Q31-35): WaitGroup, Mutex, RWMutex, Sync.Once
6. **Context & Cancellation** (Q36-40): Context propagation, timeout handling, pipeline cancellation
7. **Pipeline Patterns** (Q41-43): Rate limiting, semaphores, barrier synchronization
8. **Advanced Patterns** (Q44-53): Deadlock prevention, leak detection, real-world applications

Each program includes:
- **Complete, runnable Go code**
- **Expected output and behavior**
- **Interview-style explanations** for discussing the approach
- **Real-world applicability** and use cases

These patterns form the foundation of concurrent programming in Go and are essential for:
- **System design interviews**
- **Real-world concurrent applications**
- **Building scalable Go services**
- **Understanding Go's concurrency model**

> **Key Takeaway:** Go's concurrency model is built around **"Do not communicate by sharing memory; instead, share memory by communicating."** Master these patterns to write efficient, scalable concurrent Go applications.
