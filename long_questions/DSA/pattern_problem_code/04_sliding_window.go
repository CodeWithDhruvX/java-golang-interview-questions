package main

import (
	"fmt"
	"math"
)

// Pattern: Sliding Window
// Difficulty: Medium
// Key Concept: Maintain a "window" of array elements. Slide it forward like a frame, saving re-calculation time.

/*
INTUITION:
Imagine looking at a long train through a small window in your house.
You can only see 3 cars at a time.
As the train moves:
- One car leaves your view (Subtract it).
- One car enters your view (Add it).
You don't need to recount all the cars in the window every time. You just update the sum.
This turns a nested loop O(N*K) operation into a linear O(N) operation.

PROBLEM:
"Maximum Sum Subarray of Size K"
Given an array of integers and a number k, find the maximum sum of a subarray of size k.

ALGORITHM:
1. Initialize `windowSum = 0`, `maxSum = 0`.
2. Grow the window: Loop `i` from 0 to N.
3. Add `arr[i]` to `windowSum`.
4. If `i >= k - 1`: (This means the window has reached size K)
   - Update `maxSum` if `windowSum` is the best we've seen.
   - Shrink the window: Subtract the element that is falling out (`arr[i - (k-1)]`).
   - (The loop continues, effectively sliding the window right by 1).
*/

func maxSumSubarray(arr []int, k int) int {
	maxSum := math.MinInt
	windowSum := 0

	// DRY RUN:
	// Array: [2, 1, 5, 1, 3, 2], k=3
	//
	// i=0: Val=2. WinSum=2. (Not size 3 yet)
	// i=1: Val=1. WinSum=3. (Not size 3 yet)
	// i=2: Val=5. WinSum=3+5=8.
	//        -> Win size is 3 (indices 0,1,2).
	//        -> Update MaxSum = 8.
	//        -> Slide: Remove index 0 (Val 2). WinSum = 8 - 2 = 6.
	//
	// i=3: Val=1. WinSum=6+1=7.
	//        -> Win size is 3 (indices 1,2,3).
	//        -> MaxSum is 8 vs 7. Keep 8.
	//        -> Slide: Remove index 1 (Val 1). WinSum = 7 - 1 = 6.
	//
	// i=4: Val=3. WinSum=6+3=9.
	//        -> New MaxSum = 9!
	//        -> Slide: Remove index 2 (Val 5). WinSum = 9 - 5 = 4.

	startWindow := 0
	for endWindow := 0; endWindow < len(arr); endWindow++ {
		windowSum += arr[endWindow] // Add right element

		// Check if we hit the window size
		if endWindow >= k-1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			// Slide the window forward
			// Subtract the element going out of the window from the left
			windowSum -= arr[startWindow]
			// Move the left pointer
			startWindow++
		}
	}

	return maxSum
}

func main() {
	input := []int{2, 1, 5, 1, 3, 2}
	k := 3

	fmt.Printf("Input: %v, k: %d\n", input, k)
	result := maxSumSubarray(input, k)
	fmt.Printf("Max Subarray Sum: %d\n", result) // Expected: 9 (subarray [5, 1, 3])
}
