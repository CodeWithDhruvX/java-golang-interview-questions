package main

import "fmt"

// 46. Permutations
// Time: O(N * N!), Space: O(N!) for result + O(N) for recursion
func permute(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	current := make([]int, len(nums))
	
	backtrackPermute(nums, used, current, 0, &result)
	return result
}

func backtrackPermute(nums []int, used []bool, current []int, pos int, result *[][]int) {
	if pos == len(nums) {
		// Make a copy of current permutation
		temp := make([]int, len(nums))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			current[pos] = nums[i]
			
			backtrackPermute(nums, used, current, pos+1, result)
			
			used[i] = false
		}
	}
}

// Alternative approach using swapping
func permuteSwap(nums []int) [][]int {
	var result [][]int
	backtrackSwap(nums, 0, &result)
	return result
}

func backtrackSwap(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		// Make a copy of current permutation
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}
	
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		backtrackSwap(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3},
		{0, 1},
		{1},
		{},
		{1, 2, 3, 4},
		{1, 1, 2},
		{1, 2, 2},
		{5, 6, 7},
		{1, 2, 3, 4, 5},
	}
	
	for i, nums := range testCases {
		// Make copies for both methods
		nums1 := make([]int, len(nums))
		copy(nums1, nums)
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		
		result1 := permute(nums1)
		result2 := permuteSwap(nums2)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Used array: %d permutations\n", len(result1))
		fmt.Printf("  Swap method: %d permutations\n\n", len(result2))
	}
}
