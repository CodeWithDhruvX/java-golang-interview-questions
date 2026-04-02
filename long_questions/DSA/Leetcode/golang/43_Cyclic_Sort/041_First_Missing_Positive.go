package main

import "fmt"

// 41. First Missing Positive
// Time: O(N), Space: O(1) - Cyclic Sort approach
func firstMissingPositive(nums []int) int {
	n := len(nums)
	
	// Place each positive number in its correct position
	i := 0
	for i < n {
		correctPos := nums[i] - 1
		if nums[i] > 0 && nums[i] <= n && nums[i] != nums[correctPos] {
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		} else {
			i++
		}
	}
	
	// Find the first position where number doesn't match index+1
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	
	return n + 1 // All numbers are in correct positions
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 0},
		{3, 4, -1, 1},
		{7, 8, 9, 11, 12},
		{1, 2, 3},
		{2, 3, 4},
		{-1, -2, -3},
		{0},
		{1},
		{2},
		{3, 4, 2, 1},
		{1, 1},
		{2, 2},
		{1, 2, 6, 3, 5, 4},
		{1, 2, 0, 2, 1},
	}
	
	for i, nums := range testCases {
		// Make a copy to preserve original for display
		original := make([]int, len(nums))
		copy(original, nums)
		
		result := firstMissingPositive(nums)
		fmt.Printf("Test Case %d: %v -> First missing positive: %d\n", i+1, original, result)
	}
}
