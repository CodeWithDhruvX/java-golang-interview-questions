package main

import (
	"fmt"
	"math"
)

// Pattern: Interval Processing
// Difficulty: Medium
// Key Concept: Handling ranges systematically (Non-overlapping vs Overlapping).

/*
INTUITION:
"Insert Interval"
You have a sorted calendar of meetings: [1,3], [6,9].
You want to schedule a new meeting: [2,5].
Does it fit? Can we merge it?

Logic:
1. **No Overlap (Before)**: Meeting [1,3] ends at 3. New [2,5] starts at 2. 3 > 2. Overlap!
   Wait, if we compare strictly: Intervals ending *before* my new start are safe. Add them.

2. **Overlap (Merging)**:
   While intervals overlap with my new meeting (End >= NewStart), merge them into one GIANT meeting.
   NewStart = min(Old, New), NewEnd = max(Old, New).
   [1,3] overlaps [2,5] -> Merge -> [1,5].

3. **No Overlap (After)**: Everything else is strictly after my merged meeting. Add them.

PROBLEM:
Insert `newInterval` into `intervals` (sorted non-overlapping) and merge if necessary.

ALGORITHM:
1. Skip/Add all intervals ending before `newInterval` starts.
2. While intervals overlap `newInterval`, merge them (Update `newInterval` min/max).
3. Add the merged `newInterval`.
4. Add the rest of the intervals.
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i := 0
	n := len(intervals)

	// Step 1: Add all intervals that come strictly BEFORE the new interval
	// e.g., interval=[1,2], new=[4,8]. 2 < 4. Safe to add.
	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// Step 2: Merge overlapping intervals
	// While interval starts before newInterval ends
	// e.g., interval=[3,5], new=[4,8]. 3 < 8. Overlap!
	for i < n && intervals[i][0] <= newInterval[1] {
		// New Start = Min(current start, new start)
		newInterval[0] = int(math.Min(float64(newInterval[0]), float64(intervals[i][0])))
		// New End = Max(current end, new end)
		newInterval[1] = int(math.Max(float64(newInterval[1]), float64(intervals[i][1])))
		i++
	}
	// Add the final merged interval
	result = append(result, newInterval)

	// Step 3: Add the rest
	// e.g. interval=[9,10]. Safe.
	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}

	fmt.Printf("Existing: %v\nNew: %v\n", intervals, newInterval)

	// Dry Run:
	// 1. [1,3] vs [2,5]. 3 >= 2? Yes. Don't skip.
	// 2. Merge [1,3] and [2,5].
	//    start = min(1,2) = 1.
	//    end = max(3,5) = 5.
	//    New = [1,5].
	//    Next: [6,9]. 6 <= 5? No. Stop merging.
	// 3. Add [1,5]. Result: [[1,5]].
	// 4. Add rest: [6,9]. Result: [[1,5], [6,9]].

	res := insert(intervals, newInterval)
	fmt.Printf("Result: %v\n", res) // Expected: [[1 5] [6 9]]
}
