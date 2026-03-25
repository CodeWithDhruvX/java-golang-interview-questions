package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type WorkItem struct {
	ID    int
	Data  string
	Error error
}

func stage1(ctx context.Context, output chan<- WorkItem, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Stage1: cancelled - %v\n", ctx.Err())
			return
		case output <- WorkItem{ID: i, Data: fmt.Sprintf("Item %d", i)}:
			fmt.Printf("Stage1: produced item %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func stage2(ctx context.Context, input <-chan WorkItem, output chan<- WorkItem, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)
	
	for item := range input {
		select {
		case <-ctx.Done():
			fmt.Printf("Stage2: cancelled - %v\n", ctx.Err())
			return
		default:
			// Simulate random errors
			if rand.Intn(5) == 0 { // 20% chance of error
				item.Error = fmt.Errorf("processing error for item %d", item.ID)
				fmt.Printf("Stage2: ERROR on item %d\n", item.ID)
			} else {
				item.Data = fmt.Sprintf("Processed %s", item.Data)
				fmt.Printf("Stage2: processed item %d\n", item.ID)
			}
			output <- item
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func stage3(ctx context.Context, input <-chan WorkItem, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for item := range input {
		select {
		case <-ctx.Done():
			fmt.Printf("Stage3: cancelled - %v\n", ctx.Err())
			return
		default:
			if item.Error != nil {
				fmt.Printf("Stage3: detected error, cancelling pipeline: %v\n", item.Error)
				cancel() // Cancel the entire pipeline
				return
			}
			
			fmt.Printf("Stage3: outputting %s\n", item.Data)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	
	// Pipeline channels
	ch1 := make(chan WorkItem, 3)
	ch2 := make(chan WorkItem, 3)

	fmt.Println("Starting pipeline with error-based cancellation")

	// Start pipeline stages
	wg.Add(3)
	go stage1(ctx, ch1, &wg)
	go stage2(ctx, ch1, ch2, &wg)
	go stage3(ctx, ch2, cancel, &wg)

	wg.Wait()
	fmt.Println("Pipeline completed (cancelled due to error)")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you implement pipeline cancellation when an error occurs?

**Your Response:** I implement error-based cancellation by having stages detect errors and call the cancel function to stop the entire pipeline. The context carries the cancellation signal to all stages.

Stage 2 simulates random processing errors. When an error occurs, it's attached to the WorkItem and passed downstream. Stage 3 detects these errors and calls cancel() to stop the entire pipeline.

Each stage uses select to check for the context's Done() channel. When cancel() is called, all stages receive the cancellation signal and can stop gracefully. This prevents further processing after a critical error.

The key insight is that context provides a coordinated shutdown mechanism - when one stage detects an error, it can signal all other stages to stop immediately. This is much more efficient than letting the pipeline continue processing when you know it will fail.

This pattern is crucial for real-world systems where errors should stop further processing to conserve resources and prevent cascading failures. It's commonly used in data pipelines, request processing, and any multi-stage operation where early failure should stop all work.
