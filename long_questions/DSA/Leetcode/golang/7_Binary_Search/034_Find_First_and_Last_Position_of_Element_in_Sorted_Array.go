package main

import "fmt"

// 34. Find First and Last Position of Element in Sorted Array
// Time: O(log N), Space: O(1)
func searchRange(nums []int, target int) []int {
	return []int{findFirstOccurrence(nums, target), findLastOccurrence(nums, target)}
}

func findFirstOccurrence(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if nums[mid] == target {
			result = mid
			right = mid - 1 // Continue searching left half
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return result
}

func findLastOccurrence(nums []int, target int) int {
	left, right := 0, len(nums)-1
	result := -1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if nums[mid] == target {
			result = mid
			left = mid + 1 // Continue searching right half
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return result
}

func main() {
	// Test cases
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8},
		{[]int{5, 7, 7, 8, 8, 10}, 6},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1}, 0},
		{[]int{2, 2, 2, 2, 2}, 2},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{-3, -2, -1, 0, 1}, -1},
		{[]int{1, 3, 5, 7, 9}, 4},
	}
	
	for i, tc := range testCases {
		result := searchRange(tc.nums, tc.target)
		fmt.Printf("Test Case %d: nums=%v, target=%d -> Range: %v\n", 
			i+1, tc.nums, tc.target, result)
	}
}
