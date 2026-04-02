package main

import "fmt"

// 55. Jump Game
// Time: O(N), Space: O(1)
func canJump(nums []int) bool {
	maxReach := 0
	
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test cases
	testCases := [][]int{
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
		{0},
		{1},
		{2, 0, 0},
		{1, 1, 1, 1, 1},
		{3, 2, 1, 0, 4, 5},
		{2, 0, 0, 0, 1},
		{1, 2, 3},
		{100, 0, 0, 0, 0},
	}
	
	for i, nums := range testCases {
		result := canJump(nums)
		fmt.Printf("Test Case %d: %v -> Can jump: %t\n", i+1, nums, result)
	}
}
