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

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you pass context through a pipeline of goroutines?

**Your Response:** I pass the same context through all pipeline stages, allowing cancellation to propagate through the entire processing chain. Each stage receives the context and checks for cancellation signals.

The pipeline has three stages: stage1 generates numbers, stage2 processes them, and stage3 performs final processing. Each stage uses select to check for the context's Done() channel alongside its regular processing.

When cancel() is called, all stages receive the cancellation signal simultaneously and can stop gracefully. The context flows through the pipeline like water - when the source is cut off, everything downstream stops.

The key insight is that context provides a tree-like cancellation structure. The same context can be passed to multiple goroutines, and canceling it affects all of them. This is perfect for pipelines where you need to stop all processing when an error occurs or a timeout happens.

This pattern is essential for real-world systems like data processing pipelines, microservice chains, or any multi-stage concurrent processing where you need coordinated cancellation.
