package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Source int
	Data   string
	Value  int
}

func source(id int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Duration(id*100) * time.Millisecond) // Different speeds
		result := Result{
			Source: id,
			Data:   fmt.Sprintf("Source %d - Item %d", id, i),
			Value:  id*10 + i,
		}
		results <- result
		fmt.Printf("Source %d: sent item %d\n", id, i)
	}
}

func main() {
	numSources := 3
	results := make(chan Result, 15)
	wg := sync.WaitGroup{}

	// Start multiple sources (fan-out)
	fmt.Printf("Starting %d sources for fan-in pattern\n", numSources)
	for i := 1; i <= numSources; i++ {
		wg.Add(1)
		go source(i, results, &wg)
	}

	// Wait for sources to finish and close channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Fan-in: collect results from all sources
	fmt.Println("Fan-in: collecting results from all sources")
	allResults := []Result{}
	for result := range results {
		allResults = append(allResults, result)
		fmt.Printf("Received: %s (Value: %d)\n", result.Data, result.Value)
	}

	fmt.Printf("\nCollected %d results from all sources\n", len(allResults))
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How does the fan-in pattern merge results from multiple channels into one?

**Your Response:** The fan-in pattern combines multiple concurrent data streams into a single channel. Instead of each source having its own channel, all sources write to one shared results channel.

I create 3 source goroutines, each generating data at different speeds. They all send their results to the same results channel. The main function reads from this single channel, effectively merging all the data streams.

The key insight is that Go channels handle the multiplexing automatically - multiple goroutines can safely write to the same channel, and the reader receives data as it becomes available from any source.

Fan-in is the counterpart to fan-out. While fan-out distributes work, fan-in collects results. This pattern is extremely common in real systems like aggregating results from parallel API calls, merging logs from multiple services, or collecting data from multiple sensors.
