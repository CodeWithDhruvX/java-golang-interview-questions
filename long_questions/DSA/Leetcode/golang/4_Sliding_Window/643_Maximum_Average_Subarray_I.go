package main

import "fmt"

// 643. Maximum Average Subarray I (Fixed Size Sliding Window)
// Time: O(N), Space: O(1)
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0.0
	}
	
	// Calculate sum of first window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	
	maxSum := windowSum
	
	// Slide the window
	for i := k; i < len(nums); i++ {
		windowSum = windowSum - nums[i-k] + nums[i]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	
	return float64(maxSum) / float64(k)
}

func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4},
		{[]int{5}, 1},
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{-1, -2, -3, -4, -5}, 3},
		{[]int{0, 0, 0, 0}, 2},
		{[]int{100, 200, 300, 400, 500}, 5},
	}
	
	for i, tc := range testCases {
		result := findMaxAverage(tc.nums, tc.k)
		fmt.Printf("Test Case %d: nums=%v, k=%d -> Max Average: %.4f\n", 
			i+1, tc.nums, tc.k, result)
	}
}
