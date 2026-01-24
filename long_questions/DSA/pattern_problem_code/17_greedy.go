package main

import (
	"fmt"
	"sort"
)

// Pattern: Greedy
// Difficulty: Medium
// Key Concept: Making the locally optimal choice at each step. For Intervals: Sorting allows us to just look at "Prev" vs "Curr".

/*
INTUITION:
"Merge Intervals"
You have meetings: [1,3], [2,6], [8,10], [15,18].
Which ones overlap?
To know this efficiently, we MUST sort them by Start Time.
Sorted: [1,3], [2,6], [8,10], [15,18]. (Already sorted here).

1. Take [1,3].
2. Look at next: [2,6].
   - Does 2 start before 3 ends? Yes (2 < 3).
   - They Overlap! Merge them.
   - New merged interval: [1, max(3, 6)] -> [1, 6].
3. Look at next: [8,10].
   - Does 8 start before 6 ends? No.
   - No overlap. Add [1,6] to results. Start new current interval [8,10].

Crucial "Greedy" logic: Since they are sorted, if I don't overlap with the *immediate* previous one, I definitely won't overlap with any earlier ones.

PROBLEM:
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals.

ALGORITHM:
1. Sort intervals by `start`.
2. Initialize `results` with the first interval.
3. Iterate from 1 to N:
   - Compare `currStart` with `lastEnd` of results.
   - If Overlap (`currStart <= lastEnd`):
     - Update `lastEnd = max(lastEnd, currEnd)`.
   - Else:
     - Append current interval to results.
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// Step 1: Sort by Start Time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	// Step 2: Iterate and Merge
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := result[len(result)-1]

		if current[0] <= lastMerged[1] {
			// Overlap! Extend the end time
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
				// Note: Slice logic in Go means modifying 'lastMerged' indirectly implies
				// we need to re-assign if we weren't using pointers.
				// But here result[index] is a slice header. Modifying the content of the header is tricky.
				// SAFER WAY in Go slices:
				result[len(result)-1][1] = current[1]
			}
		} else {
			// No overlap. Add new interval.
			result = append(result, current)
		}
	}

	return result
}

func main() {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Printf("Input: %v\n", input)

	res := merge(input)
	fmt.Printf("Merged: %v\n", res) // Expected: [[1 6] [8 10] [15 18]]
}
