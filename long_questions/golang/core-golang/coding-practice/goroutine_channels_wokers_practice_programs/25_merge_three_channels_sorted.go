package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Item struct {
	Source int
	Value  int
	Data   string
}

func source(id int, items chan<- Item, wg *sync.WaitGroup) {
	defer wg.Done()
	values := []int{id, id + 10, id + 20, id + 30} // Different ranges
	
	for _, val := range values {
		time.Sleep(time.Duration(id*50) * time.Millisecond)
		item := Item{
			Source: id,
			Value:  val,
			Data:   fmt.Sprintf("Source %d: %d", id, val),
		}
		items <- item
	}
}

func main() {
	numSources := 3
	items := make(chan Item, 12)
	wg := sync.WaitGroup{}

	// Start 3 sources
	fmt.Println("Starting 3 sources with different value ranges")
	for i := 1; i <= numSources; i++ {
		wg.Add(1)
		go source(i, items, &wg)
	}

	// Wait for sources and close channel
	go func() {
		wg.Wait()
		close(items)
	}()

	// Collect all items
	allItems := []Item{}
	for item := range items {
		allItems = append(allItems, item)
		fmt.Printf("Collected: %s\n", item.Data)
	}

	// Sort by value
	sort.Slice(allItems, func(i, j int) bool {
		return allItems[i].Value < allItems[j].Value
	})

	// Print sorted results
	fmt.Println("\nSorted results:")
	for _, item := range allItems {
		fmt.Printf("Value %d: %s\n", item.Value, item.Data)
	}
}

### How to Explain in Interview (Spoken style format)

**Interviewer:** How do you merge three channels and ensure the output is sorted?

**Your Response:** I merge three channels into one using a shared items channel, then sort the collected results by value. The key insight is separating collection from ordering.

Each source generates items with different value ranges - Source 1 generates 1, 11, 21, 31; Source 2 generates 2, 12, 22, 32; and Source 3 generates 3, 13, 23, 33. All sources write to the same channel.

I collect all items into a slice, then use Go's sort package to sort by the Value field. This ensures the final output is ordered regardless of which source produced which item first.

This pattern is useful when you need to merge parallel data streams but present results in a specific order. It's commonly used in scenarios like merging sorted data from multiple databases, combining parallel API responses, or aggregating time-series data from different sources.
