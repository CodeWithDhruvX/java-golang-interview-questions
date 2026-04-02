package main

import "fmt"

// 238. Product of Array Except Self
// Time: O(N), Space: O(1) (excluding output array)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	
	// First pass: calculate left products
	leftProduct := 1
	for i := 0; i < n; i++ {
		result[i] = leftProduct
		leftProduct *= nums[i]
	}
	
	// Second pass: calculate right products and multiply with left products
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}
	
	return result
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
		{1, 2, 3, 4, 5},
		{0, 0},
		{1, 0},
		{2, 3, 0, 4},
		{-1, -2, -3, -4},
		{1, 1, 1, 1},
		{5},
		{},
	}
	
	for i, nums := range testCases {
		result := productExceptSelf(nums)
		fmt.Printf("Test Case %d: %v -> Product except self: %v\n", 
			i+1, nums, result)
	}
}
