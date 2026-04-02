package main

import "fmt"

// 287. Find the Duplicate Number
// Time: O(N), Space: O(1) - Floyd's Tortoise and Hare Algorithm
func findDuplicate(nums []int) int {
	// Phase 1: Find the intersection point
	slow := nums[0]
	fast := nums[0]
	
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	
	// Phase 2: Find the entrance to the cycle
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	
	return slow
}

// Alternative solution using cyclic sort (modifies the array)
func findDuplicateCyclicSort(nums []int) int {
	i := 0
	n := len(nums)
	
	for i < n {
		correctPos := nums[i] - 1
		if nums[i] != nums[correctPos] {
			nums[i], nums[correctPos] = nums[correctPos], nums[i]
		} else {
			i++
		}
	}
	
	// The duplicate will be at the position where the number doesn't match index+1
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return nums[i]
		}
	}
	
	return -1 // Should never reach here for valid input
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 3, 4, 2, 2},
		{3, 1, 3, 4, 2},
		{1, 1},
		{2, 2, 2, 2, 2},
		{1, 4, 4, 3, 2},
		{5, 4, 3, 2, 1, 5},
		{3, 1, 2, 3, 4, 5},
		{2, 5, 9, 6, 9, 3, 8, 9, 7, 1},
	}
	
	for i, nums := range testCases {
		// Make a copy for Floyd's algorithm (doesn't modify array)
		numsCopy1 := make([]int, len(nums))
		copy(numsCopy1, nums)
		
		// Make a copy for cyclic sort (modifies array)
		numsCopy2 := make([]int, len(nums))
		copy(numsCopy2, nums)
		
		result1 := findDuplicate(numsCopy1)
		result2 := findDuplicateCyclicSort(numsCopy2)
		
		fmt.Printf("Test Case %d: %v -> Duplicate (Floyd): %d, Duplicate (Cyclic): %d\n", 
			i+1, nums, result1, result2)
	}
}
