package main

import "fmt"

// 217. Contains Duplicate
// Time: O(N), Space: O(N)
func containsDuplicate(nums []int) bool {
	numSet := make(map[int]bool)
	
	for _, num := range nums {
		if _, exists := numSet[num]; exists {
			return true
		}
		numSet[num] = true
	}
	
	return false
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 2, 2, 2},
		{},
		{0},
		{-1, -2, -3, -1},
	}
	
	for i, nums := range testCases {
		result := containsDuplicate(nums)
		fmt.Printf("Test Case %d: %v -> Contains duplicate: %t\n", i+1, nums, result)
	}
}
