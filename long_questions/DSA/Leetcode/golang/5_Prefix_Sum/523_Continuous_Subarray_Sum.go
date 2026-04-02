package main

import "fmt"

// 523. Continuous Subarray Sum
// Time: O(N), Space: O(N)
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	
	// Map to store prefix sum modulo k and its earliest index
	prefixMod := make(map[int]int)
	prefixMod[0] = -1 // Initialize with sum 0 at index -1
	
	currentSum := 0
	
	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		
		var mod int
		if k != 0 {
			mod = currentSum % k
			if mod < 0 { // Handle negative numbers
				mod += k
			}
		} else {
			mod = currentSum // When k is 0, we need exact sum of 0
		}
		
		// Check if this modulo has been seen before
		if prevIndex, exists := prefixMod[mod]; exists {
			// Check if the subarray length is at least 2
			if i-prevIndex >= 2 {
				return true
			}
		} else {
			// Store the first occurrence of this modulo
			prefixMod[mod] = i
		}
	}
	
	return false
}

func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{23, 2, 4, 6, 7}, 6},
		{[]int{23, 2, 6, 4, 7}, 6},
		{[]int{23, 2, 6, 4, 7}, 13},
		{[]int{0, 0}, 0},
		{[]int{5, 0, 0, 0}, 3},
		{[]int{1, 2, 3}, 5},
		{[]int{1, 2, 12, 1, -1, 2, 1, 2, 10, 1}, 3},
		{[]int{1, 2}, 0},
		{[]int{0, 1, 2, 3}, 0},
		{[]int{5, 1, 2, 3, 1, 2, 3, 1, 2, 3, 5, 5}, 3},
	}
	
	for i, tc := range testCases {
		result := checkSubarraySum(tc.nums, tc.k)
		fmt.Printf("Test Case %d: nums=%v, k=%d -> Continuous subarray sum: %t\n", 
			i+1, tc.nums, tc.k, result)
	}
}
