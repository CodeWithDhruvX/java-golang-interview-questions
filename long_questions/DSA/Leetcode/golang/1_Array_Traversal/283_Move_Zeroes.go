package main

import "fmt"

// 283. Move Zeroes
// Time: O(N), Space: O(1)
func moveZeroes(nums []int) {
	lastNonZeroFoundAt := 0
	
	// Move all non-zero elements to the front
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroFoundAt] = nums[i]
			lastNonZeroFoundAt++
		}
	}
	
	// Fill the remaining positions with zeros
	for i := lastNonZeroFoundAt; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	// Test cases
	testCases := [][]int{
		{0, 1, 0, 3, 12},
		{0},
		{1, 2, 3, 4},
		{0, 0, 1, 0, 2, 0, 3},
		{4, 0, 5, 0, 3, 0, 1},
	}
	
	for i, nums := range testCases {
		// Make a copy to preserve original for display
		original := make([]int, len(nums))
		copy(original, nums)
		
		moveZeroes(nums)
		fmt.Printf("Test Case %d: %v -> After moving zeroes: %v\n", i+1, original, nums)
	}
}
