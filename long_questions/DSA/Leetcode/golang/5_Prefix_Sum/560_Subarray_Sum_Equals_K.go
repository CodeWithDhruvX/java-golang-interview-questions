package main

import "fmt"

// 560. Subarray Sum Equals K
// Time: O(N), Space: O(N)
func subarraySum(nums []int, k int) int {
	count := 0
	prefixSum := make(map[int]int)
	prefixSum[0] = 1 // Initialize with sum 0 occurring once
	
	currentSum := 0
	
	for _, num := range nums {
		currentSum += num
		
		// Check if (currentSum - k) exists in prefixSum
		if freq, exists := prefixSum[currentSum-k]; exists {
			count += freq
		}
		
		// Update prefixSum with currentSum
		prefixSum[currentSum]++
	}
	
	return count
}

func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 1, 1}, 2},
		{[]int{1, 2, 3}, 3},
		{[]int{1, -1, 0}, 0},
		{[]int{0, 0, 0, 0}, 0},
		{[]int{-1, -1, 1}, 0},
		{[]int{3, 4, 7, -2, 2, 1, 4, 2}, 7},
		{[]int{1}, 1},
		{[]int{1}, 0},
		{[]int{}, 0},
		{[]int{1, 2, 1, 2, 1}, 3},
	}
	
	for i, tc := range testCases {
		result := subarraySum(tc.nums, tc.k)
		fmt.Printf("Test Case %d: nums=%v, k=%d -> Subarrays with sum k: %d\n", 
			i+1, tc.nums, tc.k, result)
	}
}
