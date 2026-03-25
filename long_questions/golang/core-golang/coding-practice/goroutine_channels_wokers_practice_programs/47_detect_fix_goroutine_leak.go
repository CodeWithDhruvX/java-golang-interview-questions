package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func detectGoroutineLeak() {
	fmt.Println("=== Detecting and Fixing Goroutine Leaks ===")
	
	// Show initial goroutine count
	initialCount := runtime.NumGoroutine()
	fmt.Printf("Initial goroutine count: %d\n", initialCount)
	
	// Example 1: Goroutine leak - forgotten channel
	fmt.Println("\n1. Goroutine leak example:")
	leakyFunction()
	time.Sleep(100 * time.Millisecond)
	leakCount := runtime.NumGoroutine()
	fmt.Printf("Goroutines after leak: %d (leaked: %d)\n", 
		leakCount, leakCount-initialCount)
	
	// Example 2: Fixed version
	fmt.Println("\n2. Fixed version:")
	fixedFunction()
	time.Sleep(100 * time.Millisecond)
	fixedCount := runtime.NumGoroutine()
	fmt.Printf("Goroutines after fix: %d\n", fixedCount)
	
	// Example 3: Detecting leaks with WaitGroup
	fmt.Println("\n3. Using WaitGroup to prevent leaks:")
	preventLeakWithWaitGroup()
}

func leakyFunction() {
	// This goroutine leaks because it waits on a channel that will never receive
	ch := make(chan bool)
	go func() {
		<-ch // This goroutine will wait forever
		fmt.Println("This will never print")
	}()
	
	// We forgot to send to ch or close it
}

func fixedFunction() {
	// Fixed version - ensure goroutine can exit
	done := make(chan bool)
	go func() {
		select {
		case <-done:
			fmt.Println("Goroutine received signal, exiting cleanly")
		case <-time.After(50 * time.Millisecond):
			fmt.Println("Goroutine timed out, exiting")
		}
	}()
	
	// Signal the goroutine to exit
	close(done)
}

func preventLeakWithWaitGroup() {
	wg := sync.WaitGroup{}
	
	// Start multiple workers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Do some work
			fmt.Printf("Worker %d: working\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Worker %d: done\n", id)
		}(i)
	}
	
	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All workers completed, no leaks")
}

func monitorGoroutines() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	
	for range ticker.C {
		count := runtime.NumGoroutine()
		fmt.Printf("Current goroutine count: %d\n", count)
		
		if count > 100 { // Arbitrary threshold
			fmt.Printf("WARNING: High goroutine count detected: %d\n", count)
			// In production, you might trigger alerts or logging here
		}
	}
}

func main() {
	detectGoroutineLeak()
	
	fmt.Println("\n=== Goroutine Leak Prevention Tips ===")
	fmt.Println("1. Always ensure goroutines have a way to exit")
	fmt.Println("2. Use context.WithCancel for cancellation")
	fmt.Println("3. Use WaitGroup to track goroutine completion")
	fmt.Println("4. Avoid indefinite waits without timeouts")
	fmt.Println("5. Monitor goroutine count in production")
	fmt.Println("6. Use tools like 'go tool pprof' to detect leaks")
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you detect and fix goroutine leaks?

**Your Response:** Goroutine leaks occur when goroutines never exit and continue consuming memory. I detect them using runtime.NumGoroutine() to monitor the goroutine count over time.

The most common cause is goroutines waiting indefinitely on channels that will never receive or close. I demonstrate this with a leaky function that starts a goroutine waiting on a channel, but never sends to it.

To fix leaks, I ensure every goroutine has an exit path. I use several patterns: context cancellation for graceful shutdown, timeouts with select, and WaitGroups to track completion.

The key insight is that goroutine management is like resource management - every goroutine you start must have a clear termination condition. In production, I monitor goroutine counts and set up alerts for unusual growth.

For complex systems, I use profiling tools like pprof to identify leaking goroutines and their stack traces. This helps pinpoint exactly where goroutines are getting stuck.
