package main

import "fmt"

// 198. House Robber
// Time: O(N), Space: O(1)
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	
	prev2, prev1 := nums[0], max(nums[0], nums[1])
	
	for i := 2; i < len(nums); i++ {
		current := max(prev1, prev2+nums[i])
		prev2, prev1 = prev1, current
	}
	
	return prev1
}

// DP array approach
func robDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	
	return dp[len(nums)-1]
}

// Recursive with memoization
func robMemo(nums []int) int {
	memo := make(map[int]int)
	return robHelper(nums, len(nums)-1, memo)
}

func robHelper(nums []int, index int, memo map[int]int) int {
	if index < 0 {
		return 0
	}
	if index == 0 {
		return nums[0]
	}
	
	if val, exists := memo[index]; exists {
		return val
	}
	
	result := max(robHelper(nums, index-1, memo), robHelper(nums, index-2, memo)+nums[index])
	memo[index] = result
	return result
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
		{1, 2, 3, 1},
		{2, 7, 9, 3, 1},
		{2, 1, 1, 2},
		{1},
		{1, 2},
		{5, 5, 10, 100, 10, 5},
		{2, 7, 9, 3, 1, 5, 6, 8},
		{100, 1, 1, 100},
		{4, 1, 2, 7, 5, 3, 1},
	}
	
	for i, nums := range testCases {
		result1 := rob(nums)
		result2 := robDP(nums)
		result3 := robMemo(nums)
		
		fmt.Printf("Test Case %d: %v -> Iterative: %d, DP: %d, Memo: %d\n", 
			i+1, nums, result1, result2, result3)
	}
}
