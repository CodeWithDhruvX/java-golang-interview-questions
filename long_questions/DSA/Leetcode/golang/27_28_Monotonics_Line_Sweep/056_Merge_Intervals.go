package main

import (
	"fmt"
	"sort"
)

// 56. Merge Intervals - Line Sweep Algorithm
// Time: O(N log N), Space: O(N)
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	
	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	
	// Line sweep algorithm
	var result [][]int
	currentInterval := intervals[0]
	
	for i := 1; i < len(intervals); i++ {
		nextInterval := intervals[i]
		
		// Check if intervals overlap
		if nextInterval[0] <= currentInterval[1] {
			// Merge intervals
			currentInterval[1] = max(currentInterval[1], nextInterval[1])
		} else {
			// Add current interval to result and start new one
			result = append(result, currentInterval)
			currentInterval = nextInterval
		}
	}
	
	result = append(result, currentInterval)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Line sweep with event points
func mergeLineSweep(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	
	// Create events
	type Event struct {
		position int
		isStart  bool
	}
	
	var events []Event
	for _, interval := range intervals {
		events = append(events, Event{interval[0], true})   // Start event
		events = append(events, Event{interval[1], false})  // End event
	}
	
	// Sort events by position, start events before end events at same position
	sort.Slice(events, func(i, j int) bool {
		if events[i].position != events[j].position {
			return events[i].position < events[j].position
		}
		return events[i].isStart // Start events come first
	})
	
	var result [][]int
	activeCount := 0
	currentStart := -1
	
	for _, event := range events {
		if event.isStart {
			if activeCount == 0 {
				currentStart = event.position
			}
			activeCount++
		} else {
			activeCount--
			if activeCount == 0 {
				result = append(result, []int{currentStart, event.position})
			}
		}
	}
	
	return result
}

// Alternative line sweep with counting
func mergeLineSweepCounting(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	
	// Find min and max positions
	minPos, maxPos := intervals[0][0], intervals[0][1]
	for _, interval := range intervals {
		minPos = min(minPos, interval[0])
		maxPos = max(maxPos, interval[1])
	}
	
	// Create difference array
	diff := make([]int, maxPos-minPos+2)
	
	// Add intervals to difference array
	for _, interval := range intervals {
		diff[interval[0]-minPos]++
		diff[interval[1]-minPos]--
	}
	
	// Reconstruct intervals
	var result [][]int
	inInterval := false
	start := -1
	
	for i := range diff {
		if diff[i] > 0 && !inInterval {
			start = i + minPos
			inInterval = true
		} else if diff[i] <= 0 && inInterval {
			result = append(result, []int{start, i + minPos})
			inInterval = false
		}
	}
	
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Line sweep with priority queue
func mergeLineSweepPriorityQueue(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	
	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	
	var result [][]int
	
	// Simple priority queue using slice
	var pq [][]int
	current := 0
	
	for current < len(intervals) || len(pq) > 0 {
		// Add all intervals that start at or before current position
		for current < len(intervals) && (len(pq) == 0 || intervals[current][0] <= pq[0][1]) {
			pq = append(pq, intervals[current])
			current++
		}
		
		// Find minimum end time
		minEnd := pq[0][1]
		minIdx := 0
		for i := 1; i < len(pq); i++ {
			if pq[i][1] < minEnd {
				minEnd = pq[i][1]
				minIdx = i
			}
		}
		
		// Remove intervals that end at or before minEnd
		newPQ := [][]int{}
		for _, interval := range pq {
			if interval[1] > minEnd {
				newPQ = append(newPQ, interval)
			}
		}
		pq = newPQ
		
		// Add merged interval
		if len(pq) == 0 && current < len(intervals) {
			// Start new interval
			if current > 0 {
				result = append(result, []int{intervals[current-1][0], minEnd})
			}
		}
	}
	
	return result
}

// Optimized line sweep with two pointers
func mergeTwoPointers(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	
	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	
	var result [][]int
	merged := intervals[0]
	
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		
		if current[0] <= merged[1] {
			// Overlap, merge
			merged[1] = max(merged[1], current[1])
		} else {
			// No overlap, add merged interval and start new
			result = append(result, merged)
			merged = current
		}
	}
	
	result = append(result, merged)
	return result
}

// Line sweep with interval tree simulation
func mergeIntervalTree(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	
	// Sort by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	
	var result [][]int
	current := intervals[0]
	
	for i := 1; i < len(intervals); i++ {
		next := intervals[i]
		
		// Check overlap
		if next[0] <= current[1] {
			// Merge
			current[1] = max(current[1], next[1])
		} else {
			// Add current and start new
			result = append(result, current)
			current = next
		}
	}
	
	result = append(result, current)
	return result
}

// Line sweep with detailed events
func mergeLineSweepDetailed(intervals [][]int) ([][]int, []string) {
	var explanation []string
	
	if len(intervals) <= 1 {
		explanation = append(explanation, "0 or 1 intervals, returning as is")
		return intervals, explanation
	}
	
	explanation = append(explanation, fmt.Sprintf("Processing %d intervals", len(intervals)))
	
	// Create events
	type Event struct {
		position int
		isStart  bool
		interval int
	}
	
	var events []Event
	for i, interval := range intervals {
		events = append(events, Event{interval[0], true, i})
		events = append(events, Event{interval[1], false, i})
		explanation = append(explanation, fmt.Sprintf("Created events for interval %d: start=%d, end=%d", 
			i, interval[0], interval[1]))
	}
	
	// Sort events
	sort.Slice(events, func(i, j int) bool {
		if events[i].position != events[j].position {
			return events[i].position < events[j].position
		}
		return events[i].isStart
	})
	
	explanation = append(explanation, "Sorted events by position (start before end at same position)")
	
	var result [][]int
	activeCount := 0
	currentStart := -1
	
	for i, event := range events {
		if event.isStart {
			if activeCount == 0 {
				currentStart = event.position
				explanation = append(explanation, fmt.Sprintf("Event %d: Starting new interval at %d", 
					i, currentStart))
			}
			activeCount++
		} else {
			activeCount--
			if activeCount == 0 {
				result = append(result, []int{currentStart, event.position})
				explanation = append(explanation, fmt.Sprintf("Event %d: Ending interval at %d, result: %v", 
					i, event.position, []int{currentStart, event.position}))
			}
		}
	}
	
	return result, explanation
}

func main() {
	// Test cases
	fmt.Println("=== Testing Merge Intervals - Line Sweep ===")
	
	testCases := []struct {
		intervals   [][]int
		description string
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, "Standard case"},
		{[][]int{{1, 4}, {4, 5}}, "Touching intervals"},
		{[][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}}, "Nested intervals"},
		{[][]int{{1, 3}, {5, 7}, {9, 11}}, "Non-overlapping"},
		{[][]int{{1, 5}}, "Single interval"},
		{[][]int{}, "Empty array"},
		{[][]int{{1, 2}, {3, 4}, {2, 3}}, "Sequential with overlap"},
		{[][]int{{1, 4}, {0, 2}, {3, 5}}, "Mixed order"},
		{[][]int{{1, 100}, {2, 3}, {4, 5}, {6, 7}, {8, 9}, {10, 11}}, "Large interval with many small"},
		{[][]int{{1, 3}, {2, 4}, {3, 5}, {4, 6}}, "Chain overlapping"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Input: %v\n", tc.intervals)
		
		result1 := merge(tc.intervals)
		result2 := mergeLineSweep(tc.intervals)
		result3 := mergeLineSweepCounting(tc.intervals)
		result4 := mergeTwoPointers(tc.intervals)
		
		fmt.Printf("  Standard: %v\n", result1)
		fmt.Printf("  Line Sweep: %v\n", result2)
		fmt.Printf("  Counting: %v\n", result3)
		fmt.Printf("  Two Pointers: %v\n\n", result4)
	}
	
	// Detailed explanation
	fmt.Println("=== Detailed Line Sweep Explanation ===")
	testIntervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result, explanation := mergeLineSweepDetailed(testIntervals)
	
	fmt.Printf("Result: %v\n", result)
	for _, step := range explanation {
		fmt.Printf("  %s\n", step)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largeIntervals := make([][]int, 10000)
	for i := 0; i < 10000; i++ {
		largeIntervals[i] = []int{i, i + 10}
	}
	
	fmt.Printf("Large test with %d intervals\n", len(largeIntervals))
	
	result = mergeLineSweep(largeIntervals)
	fmt.Printf("Line sweep result length: %d\n", len(result))
	
	result = mergeTwoPointers(largeIntervals)
	fmt.Printf("Two pointers result length: %d\n", len(result))
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Overlapping at same point
	samePoint := [][]int{{1, 3}, {3, 5}, {3, 7}}
	fmt.Printf("Same point overlap: %v\n", mergeLineSweep(samePoint))
	
	// Single point intervals
	singlePoints := [][]int{{1, 1}, {2, 2}, {3, 3}}
	fmt.Printf("Single points: %v\n", mergeLineSweep(singlePoints))
	
	// Very large numbers
	largeNumbers := [][]int{{1000000, 2000000}, {1500000, 2500000}, {3000000, 4000000}}
	fmt.Printf("Large numbers: %v\n", mergeLineSweep(largeNumbers))
	
	// Negative numbers
	negativeNumbers := [][]int{{-5, -3}, {-4, -2}, {-1, 1}}
	fmt.Printf("Negative numbers: %v\n", mergeLineSweep(negativeNumbers))
	
	// Test priority queue approach
	fmt.Println("\n=== Priority Queue Test ===")
	pqTest := [][]int{{1, 4}, {2, 5}, {3, 6}, {7, 8}}
	fmt.Printf("Priority queue: %v\n", mergeLineSweepPriorityQueue(pqTest))
}
