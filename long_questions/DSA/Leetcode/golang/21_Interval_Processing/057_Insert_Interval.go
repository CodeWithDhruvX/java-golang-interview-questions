package main

import "fmt"

// 57. Insert Interval
// Time: O(N), Space: O(N)
func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	i := 0
	n := len(intervals)
	
	// Add all intervals that end before the new interval starts
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}
	
	// Merge all overlapping intervals
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	
	// Add the merged interval
	result = append(result, newInterval)
	
	// Add all remaining intervals
	for i < n {
		result = append(result, intervals[i])
		i++
	}
	
	return result
}

// Alternative approach using binary search to find insertion point
func insertBinarySearch(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	
	// Find the position to insert the new interval
	left, right := 0, len(intervals)-1
	insertPos := len(intervals)
	
	for left <= right {
		mid := left + (right-left)/2
		
		if intervals[mid][0] < newInterval[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	insertPos = left
	
	// Insert the new interval
	intervals = append(intervals, []int{})
	copy(intervals[insertPos+1:], intervals[insertPos:])
	intervals[insertPos] = newInterval
	
	// Merge overlapping intervals
	return mergeIntervals(intervals)
}

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	
	var result [][]int
	current := intervals[0]
	
	for i := 1; i < len(intervals); i++ {
		next := intervals[i]
		
		if current[1] >= next[0] {
			// Merge intervals
			current[1] = max(current[1], next[1])
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

// In-place modification approach
func insertInPlace(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	
	// Find the insertion point
	insertPos := 0
	for insertPos < len(intervals) && intervals[insertPos][0] < newInterval[0] {
		insertPos++
	}
	
	// Insert the new interval
	intervals = append(intervals, []int{})
	copy(intervals[insertPos+1:], intervals[insertPos:])
	intervals[insertPos] = newInterval
	
	// Merge overlapping intervals
	mergedIndex := 0
	
	for i := 1; i < len(intervals); i++ {
		if intervals[mergedIndex][1] >= intervals[i][0] {
			// Merge intervals
			intervals[mergedIndex][1] = max(intervals[mergedIndex][1], intervals[i][1])
		} else {
			// Move to next position
			mergedIndex++
			intervals[mergedIndex] = intervals[i]
		}
	}
	
	return intervals[:mergedIndex+1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test cases
	testCases := []struct {
		intervals   [][]int
		newInterval []int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}},
		{[][]int{}, []int{5, 7}},
		{[][]int{{1, 5}}, []int{2, 3}},
		{[][]int{{1, 5}}, []int{2, 7}},
		{[][]int{{1, 5}}, []int{6, 8}},
		{[][]int{{1, 3}, {6, 9}}, []int{10, 12}},
		{[][]int{{1, 3}, {6, 9}}, []int{0, 0}},
		{[][]int{{2, 4}, {5, 7}, {8, 10}}, []int{1, 11}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{11, 13}},
	}
	
	for i, tc := range testCases {
		// Make copies for different approaches
		intervals1 := make([][]int, len(tc.intervals))
		copy(intervals1, tc.intervals)
		
		intervals2 := make([][]int, len(tc.intervals))
		copy(intervals2, tc.intervals)
		
		intervals3 := make([][]int, len(tc.intervals))
		copy(intervals3, tc.intervals)
		
		result1 := insert(tc.intervals, tc.newInterval)
		result2 := insertBinarySearch(intervals1, tc.newInterval)
		result3 := insertInPlace(intervals3, tc.newInterval)
		
		fmt.Printf("Test Case %d: intervals=%v, new=%v\n", i+1, tc.intervals, tc.newInterval)
		fmt.Printf("  Linear: %v\n", result1)
		fmt.Printf("  Binary: %v\n", result2)
		fmt.Printf("  In-place: %v\n\n", result3)
	}
}
