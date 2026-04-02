package main

import "fmt"

// 1. Two Sum
// Time: O(N), Space: O(N)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	
	for i, num := range nums {
		complement := target - num
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}
		numMap[num] = i
	}
	
	return []int{}
}

func main() {
	// Test cases
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
		{[]int{1, 2, 3, 4, 5}, 9},
		{[]int{-1, -2, -3, -4, -5}, -8},
	}
	
	for i, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		fmt.Printf("Test Case %d: nums=%v, target=%d -> Indices: %v\n", i+1, tc.nums, tc.target, result)
	}
}
