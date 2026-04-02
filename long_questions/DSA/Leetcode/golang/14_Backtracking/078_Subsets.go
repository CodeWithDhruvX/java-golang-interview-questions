package main

import "fmt"

// 78. Subsets
// Time: O(N * 2^N), Space: O(N) for recursion + O(2^N) for result
func subsets(nums []int) [][]int {
	var result [][]int
	current := make([]int, 0, len(nums))
	
	backtrackSubsets(nums, 0, current, &result)
	return result
}

func backtrackSubsets(nums []int, start int, current []int, result *[][]int) {
	// Add current subset to result
	temp := make([]int, len(current))
	copy(temp, current)
	*result = append(*result, temp)
	
	// Generate subsets by including each remaining element
	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])
		backtrackSubsets(nums, i+1, current, result)
		current = current[:len(current)-1]
	}
}

// Iterative approach
func subsetsIterative(nums []int) [][]int {
	result := [][]int{{}} // Start with empty subset
	
	for _, num := range nums {
		newSubsets := make([][]int, 0, len(result))
		
		// For each existing subset, create a new subset with current number
		for _, subset := range result {
			newSubset := make([]int, len(subset))
			copy(newSubset, subset)
			newSubset = append(newSubset, num)
			newSubsets = append(newSubsets, newSubset)
		}
		
		// Add all new subsets to result
		result = append(result, newSubsets...)
	}
	
	return result
}

// Bit manipulation approach
func subsetsBit(nums []int) [][]int {
	n := len(nums)
	result := make([][]int, 0, 1<<n)
	
	for mask := 0; mask < 1<<n; mask++ {
		subset := make([]int, 0)
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 1 {
				subset = append(subset, nums[i])
			}
		}
		result = append(result, subset)
	}
	
	return result
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3},
		{0},
		{},
		{1, 2},
		{1, 2, 3, 4},
		{5},
		{1, 1, 2},
		{-1, -2, -3},
		{100, 200, 300},
		{1},
	}
	
	for i, nums := range testCases {
		result1 := subsets(nums)
		result2 := subsetsIterative(nums)
		result3 := subsetsBit(nums)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Recursive: %d subsets\n", len(result1))
		fmt.Printf("  Iterative: %d subsets\n", len(result2))
		fmt.Printf("  Bit method: %d subsets\n\n", len(result3))
	}
}
