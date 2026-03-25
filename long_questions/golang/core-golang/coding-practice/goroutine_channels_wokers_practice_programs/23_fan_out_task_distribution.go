package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Data string
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d: processing task %d\n", id, task.ID)
		time.Sleep(time.Duration(task.ID%3+1) * 200 * time.Millisecond)
		fmt.Printf("Worker %d: completed task %d\n", id, task.ID)
	}
}

func main() {
	numWorkers := 4
	numTasks := 12
	tasks := make(chan Task, numTasks)
	wg := sync.WaitGroup{}

	// Fan-out: start multiple workers
	fmt.Printf("Fan-out: distributing %d tasks to %d workers\n", numTasks, numWorkers)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Send tasks
	for i := 1; i <= numTasks; i++ {
		task := Task{
			ID:   i,
			Data: fmt.Sprintf("Work item %d", i),
		}
		tasks <- task
		fmt.Printf("Main: dispatched task %d\n", i)
	}
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("Fan-out completed: all tasks distributed and processed")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** Can you explain the fan-out pattern and how it distributes tasks to multiple workers?

**Your Response:** The fan-out pattern distributes work from a single source to multiple concurrent workers. I create 4 worker goroutines that all read from the same tasks channel.

When tasks are sent to the channel, Go's runtime automatically distributes them to available workers. If multiple workers are waiting, one is chosen randomly to receive the task. This creates natural load balancing without any complex scheduling logic.

The key insight is that the channel acts as a work queue - workers pull tasks as they become available, rather than having tasks pushed to specific workers. This allows for efficient distribution where idle workers automatically get more work.

Fan-out is extremely useful for parallelizing independent tasks. It's commonly used in real systems for parallel API calls, image processing, data transformation pipelines, and any scenario where you have more work than a single processor can handle efficiently.
