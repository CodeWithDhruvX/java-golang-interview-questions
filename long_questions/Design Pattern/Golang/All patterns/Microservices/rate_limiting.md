# Rate Limiting Pattern

## 🟢 What is it?
The **Rate Limiting Pattern** restricts the number of requests a user/service can make within a specific time window. It protects your service from being overwhelmed (DoS attacks, scraping, or just spike traffic) and ensures fair usage.

Common Algorithms:
1.  **Token Bucket**: Tokens are added to a bucket at a fixed rate. Requests need a token to proceed.
2.  **Leaky Bucket**: Requests queue up and leak out at a constant rate.
3.  **Fixed Window**: "Max 100 reqs per minute" (Can burst at edge of windows).

---

## 🏛️ Real World Analogy
**Club Entry**:
*   The bouncer only lets 1 person in every 10 seconds (Leaky Bucket).
*   OR, the bouncer gives out 50 wristbands an hour. If they are gone, you wait for the next hour (Fixed Window).

---

## 🎯 Strategy to Implement (Token Bucket in Go)

1.  **Bucket**: A struct holding `tokens` (float) and `lastUpdated` (time).
2.  **Refill**: Calculate how much time passed since `lastUpdated`. Add tokens accordingly (e.g., 10 tokens/sec).
3.  **Take**: If `tokens >= 1.0`, decrement and allow. Else, reject.
4.  **Library**: Standard Go library `golang.org/x/time/rate` implements this perfectly.

---

## 💻 Code Example (Using standard library)

```go
package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// 1. Create Limiter
	// Limit: 5 requests per second
	// Burst: 10 requests max at once
	limiter := rate.NewLimiter(5, 10)

	// Simulate 20 incoming requests
	for i := 0; i < 20; i++ {
		// 2. Wait / Allow
		// limiter.Wait(ctx) blocks until a token is available.
		// limiter.Allow() returns false immediately if no token.
		
		ctx := context.Background()
		err := limiter.Wait(ctx) // This blocks if bucket empty!
		
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Request %d processed at %v\n", i, time.Now().Format("15:04:05.000"))
	}
}
```

---

## ✅ When to use?

*   **Public APIs**: prevent one user from using 100% of your server capacity.
*   **Cost Control**: If you pay for an SMS service per message, rate limit your own backend to ensure you don't accidentally send 1M texts.
*   **Brute Force Protection**: Limit login attempts to 5 per minute.
