package main

import "fmt"

// Pattern: Binary Search
// Difficulty: Easy
// Key Concept: Divide and Conquer on a sorted array. Eliminate half the search space every step.

/*
INTUITION:
I'm thinking of a number between 1 and 100.
If you guess 50 and I say "Too High", you instantly know the answer isn't 51-100. You just eliminated 50 numbers in one guess.
You then guess 25 (half of the remaining 1-49).
This logarithmic reduction O(log N) is incredibly fast.
1 billion items takes only ~30 guesses.

PROBLEM:
Given a sorted array `nums` and a `target`, return the index of `target`. Return -1 if not found.

ALGORITHM:
1. `low = 0`, `high = len - 1`.
2. Loop while `low <= high`:
   - `mid = low + (high - low) / 2`. (Safe way to calculate average to avoid integer overflow).
   - If `arr[mid] == target`: Found it! Return `mid`.
   - If `arr[mid] < target`: Middle is too small. We need bigger numbers. Go Right (`low = mid + 1`).
   - If `arr[mid] > target`: Middle is too big. We need smaller numbers. Go Left (`high = mid - 1`).
3. If loop exits, number isn't there. Return -1.
*/

func binarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	// DRY RUN:
	// Input: [-1, 0, 3, 5, 9, 12], Target: 9
	//
	// Search 1:
	// low=0, high=5.
	// mid = 0 + (5-0)/2 = 2.
	// val = nums[2] = 3.
	// 3 < 9. Too small!
	// Action: low = mid + 1 = 3. (Discard indices 0,1,2).
	//
	// Search 2:
	// low=3, high=5.
	// mid = 3 + (5-3)/2 = 4.
	// val = nums[4] = 9.
	// 9 == 9. Found!
	// Return 4.

	for low <= high {
		mid := low + (high-low)/2 // Prevents overflow for very large N

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	fmt.Printf("Nums: %v\nTarget: %d\n", nums, target)
	index := binarySearch(nums, target)
	fmt.Printf("Index: %d\n", index) // Expected: 4
}
