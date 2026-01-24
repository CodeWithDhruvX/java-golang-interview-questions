package main

import "fmt"

// Pattern: Two Pointers
// Difficulty: Easy/Medium
// Key Concept: Using two indices to traverse an array from different ends or speeds to save space/time.

/*
INTUITION:
Imagine reading a book. Usually, you read left-to-right (One Pointer).
But what if you want to find two numbers that sum up to a target in a sorted list?
- You put one finger at the Start (Small numbers).
- You put one finger at the End (Big numbers).
- If the sum (Start + End) is too big, you need a smaller number. So you move the End finger to the left.
- If the sum is too small, you need a bigger number. So you move the Start finger to the right.
This allows you to find the answer in ONE pass (O(N)), instead of testing every pair (O(N^2)).

PROBLEM:
"Two Sum II - Input Array Is Sorted"
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number.

ALGORITHM:
1. Initialize `left = 0`, `right = length - 1`.
2. Loop while `left < right`:
   - Calculate `sum = arr[left] + arr[right]`.
   - If `sum == target`: Return indices.
   - If `sum > target`: We need a smaller sum. Since array is sorted, moving `right` pointer left gives smaller numbers. Decrement `right`.
   - If `sum < target`: We need a larger sum. Move `left` pointer right. Increment `left`.
*/

func twoSumSorted(arr []int, target int) []int {
	left := 0
	right := len(arr) - 1

	// DRY RUN:
	// Array: [2, 7, 11, 15], Target: 9
	//
	// Step 1: left=0 (val 2), right=3 (val 15).
	// Sum = 2 + 15 = 17.
	// 17 > 9. Too big! We need smaller.
	// Action: right-- (Move right pointer to index 2).
	//
	// Step 2: left=0 (val 2), right=2 (val 11).
	// Sum = 2 + 11 = 13.
	// 13 > 9. Too big!
	// Action: right-- (Move right pointer to index 1).
	//
	// Step 3: left=0 (val 2), right=1 (val 7).
	// Sum = 2 + 7 = 9.
	// 9 == 9. Match!
	// Return [0, 1] (or [1, 2] if 1-indexed).

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			return []int{left + 1, right + 1} // Converting to 1-based index as per typical problem spec
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{-1, -1} // Not found
}

func main() {
	input := []int{2, 7, 11, 15}
	target := 9

	fmt.Printf("Input: %v, Target: %d\n", input, target)
	result := twoSumSorted(input, target)
	fmt.Printf("Indices (1-based): %v\n", result)

	/*
		Why is this O(N)?
		Because the `left` pointer only moves Right, and the `right` pointer only moves Left.
		They meet in the middle. They essentially touch each element once.
		If we used nested loops ("Brute Force"), we would touch pairs multiple times -> O(N^2).
	*/
}
