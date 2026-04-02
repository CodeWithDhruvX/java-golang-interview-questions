package main

import "fmt"

// 33. Search in Rotated Sorted Array
// Time: O(log N), Space: O(1)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if nums[mid] == target {
			return mid
		}
		
		// Determine which side is sorted
		if nums[left] <= nums[mid] {
			// Left side is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Right side is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	
	return -1
}

func main() {
	// Test cases
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
		{[]int{1}, 0},
		{[]int{1}, 1},
		{[]int{5, 1, 3}, 5},
		{[]int{4, 5, 6, 7, 8, 9, 1, 2, 3}, 9},
		{[]int{2, 3, 4, 5, 6, 7, 8, 9, 1}, 1},
		{[]int{6, 7, 0, 1, 2, 3, 4, 5}, 0},
		{[]int{6, 7, 0, 1, 2, 3, 4, 5}, 5},
		{[]int{1, 3, 5}, 2},
	}
	
	for i, tc := range testCases {
		result := search(tc.nums, tc.target)
		fmt.Printf("Test Case %d: nums=%v, target=%d -> Index: %d\n", 
			i+1, tc.nums, tc.target, result)
	}
}
