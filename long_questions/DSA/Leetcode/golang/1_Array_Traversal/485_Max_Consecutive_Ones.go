package main

import "fmt"

// 485. Max Consecutive Ones
// Time: O(N), Space: O(1)
func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	currentCount := 0
	
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 0
		}
	}
	
	return maxCount
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 1, 0, 1, 1, 1},
		{1, 0, 1, 1, 0, 1},
		{0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 1, 0, 1, 0, 1},
	}
	
	for i, nums := range testCases {
		result := findMaxConsecutiveOnes(nums)
		fmt.Printf("Test Case %d: %v -> Max Consecutive Ones: %d\n", i+1, nums, result)
	}
}
