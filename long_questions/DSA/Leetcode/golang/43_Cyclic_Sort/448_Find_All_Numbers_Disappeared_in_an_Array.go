package main

import "fmt"

// 448. Find All Numbers Disappeared in an Array
// Time: O(N), Space: O(1) - Cyclic Sort approach
func findDisappearedNumbers(nums []int) []int {
	i := 0
	n := len(nums)
	
	// Place each number in its correct position
	for i < n {
		correctPos := nums[i] - 1 // Numbers are from 1 to n
		if nums[i] != nums[correctPos] {
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		} else {
			i++
		}
	}
	
	// Find all positions where number doesn't match index+1
	var result []int
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			result = append(result, i+1)
		}
	}
	
	return result
}

func main() {
	// Test cases
	testCases := [][]int{
		{4, 3, 2, 7, 8, 2, 3, 1},
		{1, 1},
		{2, 2},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 1, 2, 2, 3, 3, 4, 4},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{1},
		{},
	}
	
	for i, nums := range testCases {
		// Make a copy to preserve original for display
		original := make([]int, len(nums))
		copy(original, nums)
		
		result := findDisappearedNumbers(nums)
		fmt.Printf("Test Case %d: %v -> Disappeared numbers: %v\n", i+1, original, result)
	}
}
