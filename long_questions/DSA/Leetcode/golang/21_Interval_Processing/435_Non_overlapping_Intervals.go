package main

import (
	"fmt"
	"sort"
)

// 435. Non-overlapping Intervals
// Time: O(N log N), Space: O(1)
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	
	// Sort intervals by end time (greedy approach)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	
	removed := 0
	prevEnd := intervals[0][1]
	
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		
		// Check if current interval overlaps with previous
		if current[0] < prevEnd {
			// Need to remove this interval
			removed++
		} else {
			// No overlap, update prevEnd
			prevEnd = current[1]
		}
	}
	
	return removed
}

// Alternative approach: sort by start time
func eraseOverlapIntervalsByStart(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	
	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	
	removed := 0
	prevEnd := intervals[0][1]
	
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		
		// Check if current interval overlaps with previous
		if current[0] < prevEnd {
			// Need to remove the interval with larger end
			removed++
			if current[1] < prevEnd {
				// Remove previous interval
				prevEnd = current[1]
			}
			// If prevEnd is smaller, we keep it and remove current
		} else {
			// No overlap, update prevEnd
			prevEnd = current[1]
		}
	}
	
	return removed
}

// Dynamic Programming approach (O(N²))
func eraseOverlapIntervalsDP(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	
	// Sort intervals by end time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	
	n := len(intervals)
	dp := make([]int, n)
	
	for i := range dp {
		dp[i] = 1 // Each interval is a valid subsequence by itself
	}
	
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if intervals[i][0] >= intervals[j][1] {
				// No overlap, can extend the subsequence
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	
	// Maximum number of non-overlapping intervals
	maxNonOverlapping := 0
	for _, val := range dp {
		maxNonOverlapping = max(maxNonOverlapping, val)
	}
	
	return n - maxNonOverlapping
}

// Recursive approach with memoization
func eraseOverlapIntervalsRecursive(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	
	// Sort intervals by end time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	
	memo := make(map[int]int)
	
	var solve func(int) int
	solve = func(index int) int {
		if index >= len(intervals) {
			return 0
		}
		
		if val, exists := memo[index]; exists {
			return val
		}
		
		// Option 1: Include current interval
		nextIndex := index + 1
		for nextIndex < len(intervals) && intervals[nextIndex][0] < intervals[index][1] {
			nextIndex++
		}
		
		include := 1 + solve(nextIndex)
		
		// Option 2: Skip current interval
		exclude := solve(index + 1)
		
		result := max(include, exclude)
		memo[index] = result
		return result
	}
	
	maxNonOverlapping := solve(0)
	return len(intervals) - maxNonOverlapping
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
		{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
		{{1, 2}, {1, 2}, {1, 2}},
		{{1, 2}, {2, 3}},
		{{1, 3}, {2, 4}, {3, 5}},
		{{1, 2}, {3, 4}, {5, 6}},
		{{1, 100}, {11, 22}, {1, 11}, {2, 12}},
		{{1, 10}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
		{{}},
		{{1, 2}},
		{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
	}
	
	for i, intervals := range testCases {
		// Make copies for different approaches
		intervals1 := make([][]int, len(intervals))
		copy(intervals1, intervals)
		
		intervals2 := make([][]int, len(intervals))
		copy(intervals2, intervals)
		
		intervals3 := make([][]int, len(intervals))
		copy(intervals3, intervals)
		
		intervals4 := make([][]int, len(intervals))
		copy(intervals4, intervals)
		
		result1 := eraseOverlapIntervals(intervals)
		result2 := eraseOverlapIntervalsByStart(intervals1)
		result3 := eraseOverlapIntervalsDP(intervals2)
		result4 := eraseOverlapIntervalsRecursive(intervals3)
		
		fmt.Printf("Test Case %d: %v\n", i+1, intervals)
		fmt.Printf("  End-time sort: %d\n", result1)
		fmt.Printf("  Start-time sort: %d\n", result2)
		fmt.Printf("  DP: %d\n", result3)
		fmt.Printf("  Recursive: %d\n\n", result4)
	}
}
