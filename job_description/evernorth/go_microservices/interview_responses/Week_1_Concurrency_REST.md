# Week 1: Go Fundamentals & Core Microservices

### 🔹 Topic: Go Concurrency (Goroutines, Channels, Context)

**Interviewer:** "How do you handle concurrency in Go, and why is the `context` package so important in a microservices environment?"

**Your Response:**
"In Go, concurrency is a first-class citizen. I primarily use **Goroutines** for lightweight threading and **Channels** to communicate between them—following the mantra 'Don't communicate by sharing memory; share memory by communicating.' 

For synchronization, I use `sync.WaitGroup` when I need to wait for a collection of goroutines to finish, and `sync.Mutex` if I'm protecting a shared resource.

The `context` package is critical because, in microservices, a single request often triggers multiple downstream calls. If a user cancels a request or a timeout occurs, I use `context` to propagate that signal across all goroutines. This prevents 'goroutine leaks' and ensures we aren't wasting resources on work that's no longer needed."

### 🔹 Interview Focus: High-Yield Questions

**1. How do you implement worker pools in Go?**
**Your Response:** "I create a pool of goroutines that stay alive and listen on a shared 'jobs' channel. The main routine sends tasks into that channel, and whichever worker is free picks it up. Once they finish, they might send the result back through another 'results' channel. This is great for limiting resource usage when you have thousands of tasks but only want to run, say, 10 at a time."

**2. Explain the difference between buffered and unbuffered channels.**
**Your Response:** "An **unbuffered channel** is like a hand-off; the sender blocks until the receiver is ready to take the value. A **buffered channel** has a capacity; the sender only blocks when the buffer is full. I use unbuffered channels for strict synchronization and buffered channels when I want to decouple the producer's speed from the consumer's, though I'm careful not to set the buffer too large to avoid hiding latency issues."

**3. How do you handle cancellation in concurrent operations?**
**Your Response:** "I use `context.WithCancel` or `context.WithTimeout`. Inside the goroutine, I use a `select` statement to listen to the `ctx.Done()` channel. If that channel closes, I stop what I'm doing and clean up. This is the standard way to ensure our services stay responsive."

**4. What are race conditions and how do you prevent them?**
**Your Response:** "A race condition happens when two goroutines try to access and modify the same variable at the same time. I prevent them by using channels for communication or `sync.Mutex` for locking. During development, I always run my tests with the `-race` flag to catch these early."

**5. How does `context.Context` work in microservices?**
**Your Response:** "It carries deadlines, cancellation signals, and request-scoped values across API boundaries. For example, if Service A calls Service B, I pass the context along. If Service A's request times out, Service B will see the 'canceled' signal through the context and stop its work immediately."

### 🔹 Week 1 Practice Problems: Spoken Walkthroughs

**1. Concurrent Web Crawler:**
"To build this, I'd use a worker pool. I'd have a channel for URLs to visit and a `WaitGroup` to track active crawling. Each worker picks a URL, fetches it, and sends discovered links back to the jobs channel (using a map to avoid duplicates, protected by a mutex)."

**2. Rate Limiter Middleware:**
"I'd implement this using a 'token bucket' algorithm. I'd use a `map[string]*rate.Limiter` (from the `x/time/rate` package) indexed by IP. For every request, the middleware checks if the IP has a token; if not, it returns a 429 Too Many Requests."

**3. Simple API Gateway:**
"The gateway acts as the single entry point. I'd use `httputil.NewSingleHostReverseProxy` in Go. Based on the URL path, I'd route the request to the appropriate internal service, perhaps adding JWT validation or logging middleware at the gateway level."

**4. Worker Pool System:**
"I'd define a `Job` interface and a `Worker` struct. The `Dispatcher` would manage a pool of workers and assign jobs from a queue. This keeps the system stable under high load by controlling the concurrency level."

**5. Circuit Breaker Pattern:**
"I'd use a library like `gobreaker`. I'd wrap my downstream HTTP calls in a function that the circuit breaker monitors. If the failure rate hits a threshold (say 50%), the breaker 'opens,' immediately failing future requests for a cooldown period to give the downstream service time to recover."
