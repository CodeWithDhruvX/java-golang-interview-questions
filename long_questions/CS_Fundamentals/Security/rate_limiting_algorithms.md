# Rate Limiting Algorithms

## 1. Why Rate Limit?
1.  **Prevent Abuse**: Denial of Service (DoS) attacks.
2.  **Cost Control**: Limit API usage for tiered pricing tiers.
3.  **Fairness**: Prevent one user from hogging all resources.

## 2. Algorithms

### A. Token Bucket
*   **Concept**: A bucket holds tokens. Tokens are added at a fixed rate (e.g., 10 tokens/sec). Each request consumes 1 token. If bucket is empty, request is rejected.
*   **Pros**: Allows **Bursts**. If bucket is full (capacity 100), you can send 100 requests instantly.
*   **Use Case**: General API limiting (AWS, Stripe).

### B. Leaky Bucket
*   **Concept**: Requests enter a bucket (queue). The bucket "leaks" (processes) requests at a constant fixed rate. If bucket overflows, requests are dropped.
*   **Pros**: **Smooths out traffic**. Output rate is constant.
*   **Cons**: No bursts allowed.
*   **Use Case**: NGINX traffic shaping, packet switching.

### C. Fixed Window Counter
*   **Concept**: Count requests in a fixed time window (e.g., 100 req per minute). Reset counter at star of next minute (12:00, 12:01).
*   **Cons**: **Edge Case Spike**. You can send 100 req at 12:00:59 and 100 req at 12:01:00. In that 1 second, you handled 200 requests (double the limit).

### D. Sliding Window Log
*   **Concept**: Store timestamp of every request. Remove timestamps older than window.
*   **Pros**: Perfectly accurate.
*   **Cons**: **High Memory usage**. Storing timestamps for millions of requests is expensive.

### E. Sliding Window Counter (Hybrid)
*   **Concept**: Approximates the count using the previous window and current window weighted average.
*   **Pros**: Accurate and Memory efficient.
*   **Formula**: `Count = CurrentWindowCount + (PreviousWindowCount * OverlapPercentage)`

## 3. Implementation (Go - Token Bucket)
A simple implementation using channels.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Bucket capacity 5, refill rate 1 per 200ms
	requests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		requests <- i
	}
	close(requests)

	ticker := time.NewTicker(200 * time.Millisecond)
	
	for req := range requests {
		// This blocks until ticker ticks
		<-ticker.C
		fmt.Println("Processed request", req, time.Now())
	}
}
```

## 4. Distributed Rate Limiting
If you have 10 servers, how do you enforce a global limit?
*   **Sticky Sessions**: Rate limit locally on each server. (Bad for load balancing).
*   **Redis (Centralized)**: Use Redis `INCR` and `EXPIRE`.
    *   *Issue*: Race conditions.
    *   *Solution*: Lua Scripts in Redis to make `Read-Check-Decr` atomic.

## 5. Interview Questions
1.  **How to handle HTTP 429 (Too Many Requests)?**
    *   *Ans*: Client should implement **Exponential Backoff**. Wait 1s, then 2s, then 4s before retrying. 
    *   Server should return `Retry-After` header.
2.  **Difference between Throttling and Rate Limiting?**
    *   *Ans*: Rate Limiting = "Reject after X". Throttling = "Slow down to X" (Queueing).
