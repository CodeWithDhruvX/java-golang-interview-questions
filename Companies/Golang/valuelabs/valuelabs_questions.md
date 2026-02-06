numbers : {1,2,3,4,5}



**/



//multiple go routine with their square behaviour

// main goroutine printing the square

// don't use the time package

// don't shared variable accros goroutine (used the channels)


To solve this, we will use a **producer-consumer pattern**. One goroutine will feed the numbers into a channel, multiple worker goroutines will calculate the squares, and the main goroutine will collect and print the results.

Since we aren't using the `time` package, we rely on **`sync.WaitGroup`** to ensure the main function doesn't exit before the work is finished.

### The Implementation

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Channels for communication
	jobs := make(chan int, len(numbers))
	results := make(chan int, len(numbers))

	var wg sync.WaitGroup

	// 1. Start worker goroutines
	// We'll spawn 3 workers to handle the squaring logic
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range jobs {
				results <- n * n
			}
		}()
	}

	// 2. Send jobs to the workers
	for _, n := range numbers {
		jobs <- n
	}
	close(jobs) // Closing tells workers there is no more data

	// 3. Wait for workers and close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// 4. Main goroutine prints the squares
	fmt.Println("Calculated Squares:")
	for res := range results {
		fmt.Printf("Result: %d\n", res)
	}
}

```

---

### How it Works

* **Communication over Sharing:** Instead of sharing a slice and a mutex, we use the `jobs` channel to distribute tasks and the `results` channel to collect them.
* **The `sync.WaitGroup`:** This is crucial since we aren't using `time.Sleep`. It keeps track of how many goroutines are still processing and allows the program to wait gracefully.
* **Closing Channels:** * Closing `jobs` signals to the workers that the "input stream" has ended.
* Closing `results` (inside a separate goroutine) signals the `main` loop that all calculations are finished, allowing the program to exit cleanly.



Would you like me to modify this to maintain the **original order** of the numbers, or perhaps show you how to use a buffered channel to handle much larger datasets?