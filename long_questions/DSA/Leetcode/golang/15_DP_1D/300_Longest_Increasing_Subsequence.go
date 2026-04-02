package main

import (
	"fmt"
	"math"
)

// 300. Longest Increasing Subsequence
// Time: O(N log N), Space: O(N)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	// tails[i] will be the smallest ending value of an increasing subsequence of length i+1
	tails := make([]int, 0)
	
	for _, num := range nums {
		// Binary search to find the insertion position
		left, right := 0, len(tails)
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}
		
		// If num is larger than all elements, append it
		if left == len(tails) {
			tails = append(tails, num)
		} else {
			// Replace the existing element
			tails[left] = num
		}
	}
	
	return len(tails)
}

// O(N^2) DP approach
func lengthOfLISDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	dp := make([]int, len(nums))
	maxLength := 1
	
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLength = max(maxLength, dp[i])
	}
	
	return maxLength
}

// Patience sorting approach (same as binary search method but more intuitive)
func lengthOfLISPatience(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	piles := make([][]int, 0)
	
	for _, num := range nums {
		// Find the leftmost pile where top card >= num
		left, right := 0, len(piles)
		for left < right {
			mid := left + (right-left)/2
			if piles[mid][len(piles[mid])-1] >= num {
				right = mid
			} else {
				left = mid + 1
			}
		}
		
		if left == len(piles) {
			// Create new pile
			piles = append(piles, []int{num})
		} else {
			// Place on existing pile
			piles[left] = append(piles[left], num)
		}
	}
	
	return len(piles)
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
		{10, 9, 2, 5, 3, 7, 101, 18},
		{0, 1, 0, 3, 2, 3},
		{7, 7, 7, 7, 7, 7, 7},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 3, 6, 7, 9, 4, 10, 5, 6},
		{2, 2, 2, 2, 2},
		{1},
		{},
		{3, 10, 2, 1, 20},
	}
	
	for i, nums := range testCases {
		result1 := lengthOfLIS(nums)
		result2 := lengthOfLISDP(nums)
		result3 := lengthOfLISPatience(nums)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Binary Search: %d, DP O(N²): %d, Patience: %d\n\n", 
			result1, result2, result3)
	}
}
