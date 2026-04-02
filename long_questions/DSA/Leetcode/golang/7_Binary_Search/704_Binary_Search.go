package main

import "fmt"

// 704. Binary Search
// Time: O(log N), Space: O(1)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
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
		{[]int{-1, 0, 3, 5, 9, 12}, 9},
		{[]int{-1, 0, 3, 5, 9, 12}, 2},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{}, 1},
		{[]int{1}, 1},
		{[]int{1}, 0},
		{[]int{-10, -5, 0, 5, 10}, -5},
		{[]int{-10, -5, 0, 5, 10}, 0},
		{[]int{2, 4, 6, 8, 10}, 7},
	}
	
	for i, tc := range testCases {
		result := search(tc.nums, tc.target)
		fmt.Printf("Test Case %d: nums=%v, target=%d -> Index: %d\n", 
			i+1, tc.nums, tc.target, result)
	}
}
