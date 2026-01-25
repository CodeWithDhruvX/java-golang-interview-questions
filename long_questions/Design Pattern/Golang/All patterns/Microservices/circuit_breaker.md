# Circuit Breaker Pattern

## 🟢 What is it?
The **Circuit Breaker Pattern** prevents an application from repeatedly trying to execute an operation that's likely to fail. It wraps a protected function call in a monitor object that tracks failures.

*   **Closed (Normal)**: Requests pass through. If failures exceed a threshold, it trips to Open.
*   **Open (Broken)**: Requests fail immediately (fast failure) without calling the downstream service. After a timeout, it moves to Half-Open.
*   **Half-Open (Test)**: Allows a limited number of requests through. If they succeed, it resets to Closed. If they fail, it goes back to Open.

---

## 🏛️ Real World Analogy
**Electrical Circuit Breaker**:
*   If you plug too many appliances into one outlet, the wiring gets hot.
*   The breaker "trips" (Open) to cut the power and prevent a fire.
*   You cannot turn the power back on until you fix the issue.
*   Once fixed, you flip the switch back (Closed).

---

## 🎯 Strategy to Implement

1.  **State Management**: maintain state (Closed, Open, Half-Open).
2.  **Counts**: Track total requests, consecutive failures, and last failure time.
3.  **Mutex**: Protect state transitions with `sync.Mutex`.
4.  **Execute Wrapper**: Create a method `Execute(func() error) error`.
    *   If Open: check timeout. If passed, allow 1 "probe" request (Half-Open). Else return "Circuit Open Error".
    *   If Closed: Run function. If error, increment failure count. If successes, reset failure count.
5.  **Libraries**: In Go, we often use `github.com/sony/gobreaker` or `github.com/afex/hystrix-go`.

---

## 💻 Code Example (Manual Implementation)

```go
package main

import (
    "errors"
    "fmt"
    "sync"
    "time"
)

type State int
const (
    StateClosed State = iota
    StateOpen
    StateHalfOpen
)

type CircuitBreaker struct {
    mu           sync.Mutex
    state        State
    failureCount int
    threshold    int
    lastFailure  time.Time
    timeout      time.Duration
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        threshold: threshold,
        timeout:   timeout,
        state:     StateClosed,
    }
}

func (cb *CircuitBreaker) Execute(action func() error) error {
    cb.mu.Lock()
    
    // 1. Check if Open
    if cb.state == StateOpen {
        if time.Since(cb.lastFailure) > cb.timeout {
             // 2. Retry Logic (Half-Open)
            fmt.Println("State: Half-Open (Probing...)")
            cb.state = StateHalfOpen
        } else {
            cb.mu.Unlock()
            return errors.New("circuit breaker is OPEN")
        }
    }
    cb.mu.Unlock()

    // 3. Executing Action
    err := action()

    cb.mu.Lock()
    defer cb.mu.Unlock()

    if err != nil {
        // 4. Handle Failure
        cb.failureCount++
        cb.lastFailure = time.Now()
        
        if cb.failureCount >= cb.threshold {
            cb.state = StateOpen
            fmt.Println("State: Open (Tripped!)")
        }
        return err
    }

    // 5. Handle Success
    cb.failureCount = 0
    cb.state = StateClosed
    fmt.Println("State: Closed (Reset)")
    return nil
}

// Simulating a flaky service
func flakyService() error {
    if time.Now().Unix()%2 == 0 {
        return errors.New("service failed")
    }
    return nil
}

func main() {
    cb := NewCircuitBreaker(2, 2*time.Second)

    for i := 0; i < 10; i++ {
        err := cb.Execute(func() error {
            return flakyService()
        })
        if err != nil {
            fmt.Printf("Request %d failed: %v\n", i, err)
        } else {
            fmt.Printf("Request %d success\n", i)
        }
        time.Sleep(500 * time.Millisecond)
    }
}
```

---

## ✅ When to use?

*   **External APIs**: When calling Third-party APIs (Stripe, Twilio) that might go down.
*   **Database**: To stop hammering a database that is already struggling or timing out.
*   **Microservices**: Prevents cascading failures where one slow service brings down the entire mesh.
