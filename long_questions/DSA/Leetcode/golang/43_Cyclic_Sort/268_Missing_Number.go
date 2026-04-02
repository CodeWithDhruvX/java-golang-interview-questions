package main

import "fmt"

// 268. Missing Number
// Time: O(N), Space: O(1) - Cyclic Sort approach
func missingNumber(nums []int) int {
	n := len(nums)
	
	// Place each number in its correct position
	i := 0
	for i < n {
		correctPos := nums[i]
		if correctPos < n && nums[i] != nums[correctPos] {
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		} else {
			i++
		}
	}
	
	// Find the first position where number doesn't match index
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	
	return n // All numbers are in correct positions, missing number is n
}

func main() {
	// Test cases
	testCases := [][]int{
		{3, 0, 1},
		{0, 1},
		{9, 6, 4, 2, 3, 5, 7, 0, 1},
		{0},
		{1},
		{},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
		{2, 0, 3, 1},
		{5, 4, 6, 0, 1, 2, 3},
	}
	
	for i, nums := range testCases {
		// Make a copy to preserve original for display
		original := make([]int, len(nums))
		copy(original, nums)
		
		result := missingNumber(nums)
		fmt.Printf("Test Case %d: %v -> Missing number: %d\n", i+1, original, result)
	}
}
