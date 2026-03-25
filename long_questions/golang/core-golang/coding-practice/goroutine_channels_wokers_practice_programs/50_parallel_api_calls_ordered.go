package main

import (
	"fmt"
	"sync"
	"time"
)

type APIResponse struct {
	RequestID int
	Data      string
	Latency   time.Duration
	Error     error
}

func apiCall(requestID int, responseChan chan<- APIResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Simulate API call with variable latency
	latency := time.Duration(requestID%5+1) * 200 * time.Millisecond
	time.Sleep(latency)
	
	response := APIResponse{
		RequestID: requestID,
		Data:      fmt.Sprintf("API response for request %d", requestID),
		Latency:   latency,
	}
	
	responseChan <- response
	fmt.Printf("API call %d completed in %v\n", requestID, latency)
}

func main() {
	numRequests := 10
	responseChan := make(chan APIResponse, numRequests)
	wg := sync.WaitGroup{}
	
	fmt.Printf("Making %d parallel API calls with ordered output\n", numRequests)
	
	// Start all API calls in parallel
	startTime := time.Now()
	for i := 1; i <= numRequests; i++ {
		wg.Add(1)
		go apiCall(i, responseChan, &wg)
	}
	
	// Wait for all calls to complete
	go func() {
		wg.Wait()
		close(responseChan)
	}()
	
	// Collect responses and maintain order
	responses := make([]APIResponse, numRequests)
	for response := range responseChan {
		responses[response.RequestID-1] = response
	}
	
	// Print in order
	fmt.Println("\nOrdered results:")
	for i, response := range responses {
		fmt.Printf("Request %d: %s (took %v)\n", 
			i+1, response.Data, response.Latency)
	}
	
	totalTime := time.Since(startTime)
	fmt.Printf("\nTotal time: %v (vs %v sequential)\n", 
		totalTime, time.Duration(numRequests)*300*time.Millisecond)
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you make parallel API calls but maintain ordered response output?

**Your Response:** I separate parallel execution from ordered output by making all API calls concurrently but collecting responses with their original request IDs.

All API calls start simultaneously in separate goroutines. Each call simulates network latency with different durations to show real-world variability. When a call completes, it sends its response to a shared channel with the request ID included.

I collect all responses in a slice indexed by request ID, which automatically places them in the correct order regardless of completion time. The key insight is that we don't need responses to arrive in order - we just need to output them in order.

This pattern provides the performance benefit of parallel execution while maintaining the user experience of sequential, predictable output. It's extremely useful in real systems like:
- Dashboard widgets loading data from multiple APIs
- Microservice aggregations
- Batch data processing
- Any scenario where you need parallel performance but ordered presentation

The total time shows the efficiency gain - instead of waiting for each API call sequentially, we only wait for the slowest call.
