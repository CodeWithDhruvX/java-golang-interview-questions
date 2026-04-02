package main

import (
	"fmt"
	"sort"
)

// 15. 3Sum
// Time: O(N^2), Space: O(1) (ignoring output space)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)
	
	for i := 0; i < n-2; i++ {
		// Skip duplicates for the first element
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		
		// Two pointers approach for the remaining two elements
		left, right := i+1, n-1
		target := -nums[i]
		
		for left < right {
			sum := nums[left] + nums[right]
			
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				
				// Skip duplicates for the second element
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for the third element
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	
	return result
}

func main() {
	// Test cases
	testCases := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{0, 1, 1},
		{0, 0, 0},
		{-2, 0, 1, 1, 2},
		{-1, -2, -3, -4, -5},
		{1, 2, -2, -1},
		{3, -2, 1, 0, -1, 2, -3},
		{},
		{0},
	}
	
	for i, nums := range testCases {
		result := threeSum(nums)
		fmt.Printf("Test Case %d: %v -> Triplets: %v\n", i+1, nums, result)
	}
}
