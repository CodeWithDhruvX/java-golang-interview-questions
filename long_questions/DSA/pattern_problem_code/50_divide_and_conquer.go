package main

import (
	"fmt"
)

// Pattern: Divide and Conquer
// Difficulty: Medium
// Key Concept: Breaking a problem into smaller sub-problems, solving them recursively, and combining the results.

/*
INTUITION:
To sort a deck of cards:
1. Divide deck in half.
2. Hand halves to two friends to sort.
3. Once they return sorted halves, you "merge" them by taking the top card of whichever half is smaller.
This is O(N log N) consistently.

Key Applications: Merge Sort, Quick Sort, Closest Pair of Points, Count Inversions (how unsorted is an array?).

PROBLEM:
LeetCode 912. Sort an Array.
Given an array of integers nums, sort the array in ascending order and return it.

ALGORITHM:
1. Base case: If array length <= 1, return it.
2. Mid = n / 2.
3. Left = MergeSort(0..mid).
4. Right = MergeSort(mid..end).
5. Merge(Left, Right):
   - Two pointers `i` (left), `j` (right).
   - Compare `left[i]` and `right[j]`. Append smaller to Result.
   - Append remaining elements.
*/

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Compare and merge
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	// [5, 2, 3, 1]
	// Split: [5, 2], [3, 1]
	// Split: [5],[2] -> Merge -> [2, 5]
	// Split: [3],[1] -> Merge -> [1, 3]
	// Merge [2, 5] and [1, 3]:
	// 1 < 2 -> [1] (j=1)
	// 2 < 3 -> [1, 2] (i=1)
	// 3 < 5 -> [1, 2, 3] (j=2)
	// Append 5 -> [1, 2, 3, 5]

	nums := []int{5, 2, 3, 1}
	fmt.Printf("Original: %v\n", nums)
	fmt.Printf("Sorted: %v\n", sortArray(nums))
}
