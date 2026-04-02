package main

import "fmt"

// 209. Minimum Size Subarray Sum (Variable Size Sliding Window)
// Time: O(N), Space: O(1)
func minSubArrayLen(target int, nums []int) int {
	left := 0
	currentSum := 0
	minLength := len(nums) + 1 // Initialize with a value larger than any possible result
	
	for right := 0; right < len(nums); right++ {
		currentSum += nums[right]
		
		// Shrink the window from the left as much as possible
		for currentSum >= target {
			currentLength := right - left + 1
			if currentLength < minLength {
				minLength = currentLength
			}
			currentSum -= nums[left]
			left++
		}
	}
	
	if minLength == len(nums)+1 {
		return 0 // No subarray found
	}
	
	return minLength
}

func main() {
	// Test cases
	testCases := []struct {
		target int
		nums   []int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}},
		{4, []int{1, 4, 4}},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}},
		{15, []int{1, 2, 3, 4, 5}},
		{100, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{5, []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}},
		{3, []int{1, 1, 1, 1, 1, 1, 1}},
		{1, []int{2, 3, 1, 2, 4, 3}},
	}
	
	for i, tc := range testCases {
		result := minSubArrayLen(tc.target, tc.nums)
		fmt.Printf("Test Case %d: target=%d, nums=%v -> Min length: %d\n", 
			i+1, tc.target, tc.nums, result)
	}
}
