# Worker Pool Pattern

## 🟢 What is it?
The **Worker Pool Pattern** (also known as a Thread Pool) involves creating a fixed number of worker goroutines to process a stream of jobs. Instead of spawning a new goroutine for every single task (which can exhaust system resources like memory or file descriptors), you "pool" a set of workers that pick jobs off a queue (channel).

This is **essential** in Go for controlling concurrency and preventing resource exhaustion.

---

## 🏛️ Real World Analogy
Think of a **Bank Teller Line**:
*   **Without a Pool**: Imagine if every time a customer walked in, the bank built a brand new teller counter and hired a new teller just for that one person. The process of hiring/building (spawning goroutines) takes time, and eventually, the bank runs out of space.
*   **With a Pool**: The bank has **5 fixed tellers**. Customers wait in a single line (the channel). As soon as a teller is free, they call "Next!" and process the transaction. This keeps the bank orderly and efficient regardless of how many customers arrive.

---

## 🎯 Strategy to Implement

1.  **Job Struct**: Define a struct that holds the data needed to process a single unit of work.
2.  **Result Struct**: Define a struct to hold the output/error of the processing.
3.  **Job Channel**: Create a buffered channel to hold incoming jobs.
4.  **Result Channel**: Create a buffered channel to collect results.
5.  **Worker Function**: A function that loops over the `jobs` channel, processes data, and sends output to the `results` channel.
6.  **Dispatcher**: Spawn a fixed number (e.g., `numWorkers`) of worker goroutines.
7.  **Submission**: Send all jobs to the `jobs` channel and close it.
8.  **Collection**: Read from the `results` channel.

---

## 💻 Code Example

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1. Job represents the work to be done
type Job struct {
	ID    int
	Value int
}

// 2. Result represents the outcome
type Result struct {
	JobID  int
	Output int
	Err    error
}

// 5. Worker Function
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j.ID)
		
		// Simulate expensive processing
		time.Sleep(time.Millisecond * 500)
		output := j.Value * 2 // Example work: doubling the value

		// Send result
		results <- Result{JobID: j.ID, Output: output, Err: nil}
		
		fmt.Printf("Worker %d finished job %d\n", id, j.ID)
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	// 3 & 4. Create Channels
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// 6. Dispatcher: Start the worker pool
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// 7. Submission: Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: rand.Intn(100)}
	}
	close(jobs) // Important: Close jobs channel so workers know when to stop

	// 8. Collection: Collect results
	// We know exactly how many results to expect, so we can loop numJobs times.
	// Alternatively, use a sync.WaitGroup to close the results channel.
	for a := 1; a <= numJobs; a++ {
		res := <-results
		fmt.Printf("Result received: JobID %d, Output %d\n", res.JobID, res.Output)
	}
}
```

---

## ✅ When to use?

*   **Rate Limiting**: When you need to limit the number of concurrent connections to a database or external API (e.g., only 50 requests at a time).
*   **CPU Bound Tasks**: When you have millions of small tasks (like image processing) and you only want to use as many workers as you have CPU cores.
*   **Queue Processing**: Processing items from a message queue (like RabbitMQ or Kafka) where you want a fixed number of consumers.
