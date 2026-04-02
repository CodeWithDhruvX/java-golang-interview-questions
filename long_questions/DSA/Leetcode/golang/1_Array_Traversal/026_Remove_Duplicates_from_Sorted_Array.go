package main

import "fmt"

// 26. Remove Duplicates from Sorted Array
// Time: O(N), Space: O(1)
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	uniqueIndex := 0
	
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}
	
	return uniqueIndex + 1
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{1, 2, 3, 4, 5},
		{1, 1, 1, 1},
		{},
		{2},
	}
	
	for i, nums := range testCases {
		// Make a copy to preserve original for display
		original := make([]int, len(nums))
		copy(original, nums)
		
		length := removeDuplicates(nums)
		result := nums[:length]
		fmt.Printf("Test Case %d: %v -> Length: %d, Unique elements: %v\n", i+1, original, length, result)
	}
}
