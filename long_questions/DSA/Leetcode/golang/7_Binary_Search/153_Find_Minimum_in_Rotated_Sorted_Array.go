package main

import "fmt"

// 153. Find Minimum in Rotated Sorted Array
// Time: O(log N), Space: O(1)
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	
	// If array is not rotated
	if nums[left] <= nums[right] {
		return nums[left]
	}
	
	for left < right {
		mid := left + (right-left)/2
		
		// Check if mid is the minimum element
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		
		// Check if mid+1 is the minimum element
		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		
		// Decide which half to search
		if nums[mid] >= nums[left] {
			// Left half is sorted, search right half
			left = mid + 1
		} else {
			// Right half is sorted, search left half
			right = mid - 1
		}
	}
	
	return nums[left]
}

func main() {
	// Test cases
	testCases := [][]int{
		{3, 4, 5, 1, 2},
		{4, 5, 6, 7, 0, 1, 2},
		{11, 13, 15, 17, 2, 5, 6, 8, 10},
		{1, 2, 3, 4, 5},
		{2, 1},
		{1},
		{5, 1, 2, 3, 4},
		{2, 3, 4, 5, 1},
		{3, 4, 5, 6, 7, 0, 1, 2},
		{10, 20, 30, 40, 5, 6, 7},
	}
	
	for i, nums := range testCases {
		result := findMin(nums)
		fmt.Printf("Test Case %d: %v -> Minimum: %d\n", i+1, nums, result)
	}
}
