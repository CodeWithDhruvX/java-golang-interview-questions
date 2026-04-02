package main

import (
	"fmt"
	"sort"
)

// 56. Merge Intervals
// Time: O(N log N), Space: O(N)
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	
	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	
	var result [][]int
	current := intervals[0]
	
	for i := 1; i < len(intervals); i {
		next := intervals[i]
		
		// Check if intervals overlap
		if current[1] >= next[0] {
			// Merge intervals
			current[1] = max(current[1], next[1])
		} else {
			// Add current interval to result and start new current
			result = append(result, []int{current[0], current[1]})
			current = next
		}
		i++
	}
	
	// Add the last interval
	result = append(result, []int{current[0], current[1]})
	
	return result
}

// Alternative approach using in-place modification
func mergeInPlace(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	
	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	
	mergedIndex := 0
	
	for i := 1; i < len(intervals); i++ {
		// If current interval overlaps with the last merged interval
		if intervals[mergedIndex][1] >= intervals[i][0] {
			// Merge them
			intervals[mergedIndex][1] = max(intervals[mergedIndex][1], intervals[i][1])
		} else {
			// Move to next position
			mergedIndex++
			intervals[mergedIndex] = intervals[i]
		}
	}
	
	return intervals[:mergedIndex+1]
}

// Using a custom interval type
type Interval struct {
	Start, End int
}

func mergeTyped(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	
	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	
	var result []Interval
	current := intervals[0]
	
	for i := 1; i < len(intervals); i++ {
		next := intervals[i]
		
		if current.End >= next.Start {
			// Merge intervals
			current.End = max(current.End, next.End)
		} else {
			// Add current interval to result and start new current
			result = append(result, current)
			current = next
		}
	}
	
	// Add the last interval
	result = append(result, current)
	
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test cases
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 3}, {2, 4}, {5, 7}, {6, 8}},
		{{1, 2}, {3, 4}, {5, 6}},
		{{1, 10}, {2, 3}, {4, 5}, {6, 7}},
		{{1, 5}, {2, 3}, {4, 6}},
		{{1, 4}, {0, 2}, {3, 5}},
		{{1, 3}, {5, 7}, {9, 11}},
		{{1, 4}, {2, 5}, {7, 9}},
		{{}},
	}
	
	for i, intervals := range testCases {
		// Make copies for different approaches
		intervals1 := make([][]int, len(intervals))
		copy(intervals1, intervals)
		
		intervals2 := make([][]int, len(intervals))
		copy(intervals2, intervals)
		
		result1 := merge(intervals)
		result2 := mergeInPlace(intervals1)
		
		fmt.Printf("Test Case %d: %v\n", i+1, intervals)
		fmt.Printf("  Standard: %v\n", result1)
		fmt.Printf("  In-place: %v\n\n", result2)
	}
}
