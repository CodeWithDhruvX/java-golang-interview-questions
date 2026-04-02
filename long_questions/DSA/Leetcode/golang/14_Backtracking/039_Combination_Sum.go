package main

import (
	"fmt"
	"sort"
)

// 39. Combination Sum
// Time: O(2^T/M), Space: O(T/M) where T is target, M is minimum candidate
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // Sort to help with pruning
	var result [][]int
	current := make([]int, 0, len(candidates))
	
	backtrackCombinationSum(candidates, target, 0, current, &result)
	return result
}

func backtrackCombinationSum(candidates []int, target, start int, current []int, result *[][]int) {
	if target == 0 {
		// Found a valid combination
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	
	if target < 0 {
		return // Exceeded target
	}
	
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break // No need to continue as candidates are sorted
		}
		
		// Include candidates[i]
		current = append(current, candidates[i])
		backtrackCombinationSum(candidates, target-candidates[i], i, current, result) // i (not i+1) because we can reuse
		current = current[:len(current)-1] // Backtrack
	}
}

func main() {
	// Test cases
	testCases := []struct {
		candidates []int
		target    int
	}{
		{[]int{2, 3, 6, 7}, 7},
		{[]int{2, 3, 5}, 8},
		{[]int{2}, 1},
		{[]int{1}, 1},
		{[]int{1}, 2},
		{[]int{2, 3, 6, 7, 8, 10}, 10},
		{[]int{4, 5, 6, 7, 8}, 11},
		{[]int{3, 5, 7}, 0},
		{[]int{2, 4, 6, 8}, 16},
		{[]int{1, 2, 3, 4, 5}, 7},
	}
	
	for i, tc := range testCases {
		result := combinationSum(tc.candidates, tc.target)
		fmt.Printf("Test Case %d: candidates=%v, target=%d\n", i+1, tc.candidates, tc.target)
		fmt.Printf("  Combinations: %v\n\n", result)
	}
}
