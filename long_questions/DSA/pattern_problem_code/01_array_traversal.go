package main

import "fmt"

// Pattern: Array Traversal
// Difficulty: Beginner
// Key Concept: Visiting every element sequentially to process or find a value.

/*
INTUITION:
The most fundamental operation in DSA is visiting every item in a collection.
Think of it like a mailman delivering mail. He must go to House 1, then House 2, then House 3...
He cannot teleport continuously. He must traverse the route.

We use a simple loop (for loop) to iterate from index 0 to N-1.
In Go, `range` makes this cleaner, but understanding the `i` index is crucial.

PROBLEM:
Find the "Maximum Element" in an array.

ALGORITHM:
1. Assume the first element is the maximum (`currentMax = arr[0]`).
2. Walk through the rest of the array (from index 1 to End).
3. For each element `num`:
   - Ask: "Is `num` bigger than my `currentMax`?"
   - If YES: Update `currentMax = num`.
   - If NO: Keep walking.
4. At the end, `currentMax` holds the largest value encountered.
*/

func findMaxElement(arr []int) int {
	// EDGE CASE: If array is empty, handle gracefully (or panic/error depending on requirements)
	if len(arr) == 0 {
		return 0 // Or -1, or error
	}

	// Step 1: Initialize currentMax with the first element.
	// We don't use 0 because the array might contain negative numbers (e.g., [-5, -2]).
	currentMax := arr[0]

	// LINE-BY-LINE DRY RUN:
	// Array: [3, 1, 9, 2]
	// Init: currentMax = 3

	// Step 2: Iterate through the array.
	// We can start from index 1 since we already considered index 0.
	for i := 1; i < len(arr); i++ {
		// Capture the current value strictly for readability (optional)
		val := arr[i]

		// Step 3: Compare
		// i=1: val=1. Is 1 > 3? No. currentMax remains 3.
		// i=2: val=9. Is 9 > 3? Yes. currentMax becomes 9.
		// i=3: val=2. Is 2 > 9? No. currentMax remains 9.
		if val > currentMax {
			currentMax = val
		}
	}

	// Step 4: Return result
	return currentMax
}

func main() {
	// Test Cases
	input := []int{10, 5, 20, 8, 15}

	fmt.Printf("Input: %v\n", input)
	result := findMaxElement(input)
	fmt.Printf("Max Element: %d\n", result) // Expected: 20

	// Dry Run Visualization in Output
	// Start: Max=10
	// Visit 5: 5 > 10? No.
	// Visit 20: 20 > 10? Yes. Max=20.
	// Visit 8: 8 > 20? No.
	// Visit 15: 15 > 20? No.
	// End: 20
}
