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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Sort and Merge for Interval Overlapping
- **Sorting Strategy**: Sort intervals by start time
- **Linear Scan**: Single pass through sorted intervals
- **Overlap Detection**: Check if current.end >= next.start
- **Merging Logic**: Extend current interval when overlapping

## 2. PROBLEM CHARACTERISTICS
- **Interval Merging**: Combine overlapping intervals into non-overlapping
- **Sorting Required**: Need intervals in chronological order
- **Overlap Logic**: Intervals overlap if current.end >= next.start
- **Linear Processing**: After sorting, single linear pass suffices

## 3. SIMILAR PROBLEMS
- Insert Interval (LeetCode 57) - Insert new interval into merged list
- Non-overlapping Intervals (LeetCode 435) - Remove intervals to make non-overlapping
- Meeting Rooms (LeetCode 252) - Find minimum meeting rooms needed
- Employee Free Time (LeetCode 759) - Find available time slots

## 4. KEY OBSERVATIONS
- **Sorting Essential**: Must sort by start time first
- **Overlap Condition**: current.end >= next.start indicates overlap
- **Merge Logic**: New merged interval has current.start and max(current.end, next.end)
- **Linear After Sort**: O(N) processing after O(N log N) sorting

## 5. VARIATIONS & EXTENSIONS
- **Different Overlap Definitions**: Strict vs non-strict overlap
- **Interval Insertion**: Add new interval to already merged list
- **Range Queries**: Query how many intervals overlap with point
- **Dynamic Updates**: Support insertions and deletions

## 6. INTERVIEW INSIGHTS
- Always clarify: "Overlap definition? Interval format? Empty input?"
- Edge cases: empty list, single interval, all overlapping, none overlapping
- Time complexity: O(N log N) for sorting + O(N) for merging
- Space complexity: O(N) for result + O(1) or O(N) for sorting
- Key insight: sorting enables linear merge pass

## 7. COMMON MISTAKES
- Not sorting intervals first
- Wrong overlap condition (using > instead of >=)
- Not handling edge cases (empty list, single interval)
- Forgetting to add last interval to result
- Incorrect merge logic for end times

## 8. OPTIMIZATION STRATEGIES
- **Sort and Merge**: O(N log N) time, O(N) space - optimal
- **In-place Merge**: O(N log N) time, O(1) extra space
- **Counting Sort**: O(N) time if range is small and bounded
- **Early Termination**: Not applicable (need process all intervals)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like merging overlapping time slots on a calendar:**
- You have meetings scheduled at different times
- Some meetings overlap and need to be combined
- Sort meetings by start time first
- Then scan through, merging any that overlap
- Like consolidating back-to-back or overlapping meetings

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of intervals [start, end]
2. **Goal**: Merge overlapping intervals into minimal non-overlapping set
3. **Constraint**: Intervals are closed [start, end] or open/closed
4. **Output**: Array of merged, non-overlapping intervals

#### Phase 2: Key Insight Recognition
- **"Sorting natural fit"** → Need chronological order
- **"Linear scan sufficient"** → After sorting, single pass works
- **"Overlap detection"** → current.end >= next.start means overlap
- **"Merge logic"** → Combine start and max(end) for merged interval

#### Phase 3: Strategy Development
```
Human thought process:
"I need to merge overlapping intervals.
First, I must sort them by start time so they're in order.
Then I can scan through once:
- Keep track of current merged interval
- For each next interval:
  - If it overlaps with current (current.end >= next.start):
    - Merge by extending current.end = max(current.end, next.end)
  - Else:
    - Add current to result
    - Start new current = next
- Add final current to result

This gives me all merged intervals!"
```

#### Phase 4: Edge Case Handling
- **Empty list**: Return empty array
- **Single interval**: Return as-is
- **All overlapping**: Should merge into one big interval
- **No overlapping**: Return sorted intervals unchanged

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Intervals: [[1,3], [2,6], [8,10], [15,18]]

Human thinking:
"First, sort by start time:
[[1,3], [2,6], [8,10], [15,18]] ✓

Now scan and merge:
current = [1,3]

Next = [2,6]:
- Does [1,3] overlap [2,6]? 3 >= 2 ✓ Yes
- Merge: current = [1, max(3,6)] = [1,6]

Next = [8,10]:
- Does [1,6] overlap [8,10]? 6 >= 8 ✗ No
- Add [1,6] to result: [[1,6]]
- current = [8,10]

Next = [15,18]:
- Does [8,10] overlap [15,18]? 10 >= 15 ✗ No
- Add [8,10] to result: [[1,6], [8,10]]
- current = [15,18]

End of list, add current: [[1,6], [8,10], [15,18]] ✓"
```

#### Phase 6: Intuition Validation
- **Why sorting works**: Chronological order enables linear processing
- **Why O(N log N)**: Sorting dominates, merge pass is linear
- **Why overlap logic works**: current.end >= next.start captures all overlaps
- **Why optimal**: Can't do better than sorting for this problem

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just check all pairs?"** → O(N²) vs O(N log N) with sorting
2. **"Should I use different overlap condition?"** → Depends on interval definition
3. **"What about in-place modification?"** → Works but more error-prone
4. **"Can I optimize further?"** → Sorting is already optimal
5. **"What about different interval formats?"** → Clarify start/end inclusivity

### Real-World Analogy
**Like consolidating overlapping meetings on a calendar:**
- You have a list of meetings with start/end times
- Some meetings overlap and should be combined
- Sort meetings by when they start
- Go through chronologically, merging any that overlap
- Result is a clean schedule with no overlapping time blocks
- Each merged block represents the total time span of related meetings

### Human-Readable Pseudocode
```
function mergeIntervals(intervals):
    if intervals is empty or has one interval:
        return intervals
    
    sort intervals by start time
    
    result = []
    current = intervals[0]
    
    for each interval in intervals[1:]:
        if current.end >= interval.start:
            // Overlap - merge them
            current.end = max(current.end, interval.end)
        else:
            // No overlap - add current and start new
            result.append(current)
            current = interval
    
    result.append(current)
    return result
```

### Execution Visualization

### Example: Intervals = [[1,3], [2,6], [8,10], [15,18]]
```
Sorting Phase:
Original: [[1,3], [2,6], [8,10], [15,18]]
Sorted:   [[1,3], [2,6], [8,10], [15,18]] ✓

Merging Phase:
Step 1: current = [1,3]
Step 2: next = [2,6], overlap? 3 >= 2 ✓
         Merge: current = [1,6]
Step 3: next = [8,10], overlap? 6 >= 8 ✗
         Add [1,6] to result, current = [8,10]
Step 4: next = [15,18], overlap? 10 >= 15 ✗
         Add [8,10] to result, current = [15,18]
End: Add [15,18] to result

Final result: [[1,6], [8,10], [15,18]] ✓
```

### Key Visualization Points:
- **Sorting First**: Essential for chronological processing
- **Overlap Detection**: current.end >= next.start
- **Merge Logic**: Keep current.start, extend current.end
- **Linear Processing**: Single pass after sorting

### Memory Layout Visualization:
```
Interval Processing State:
intervals = [[1,3], [2,6], [8,10], [15,18]]

Step-by-step evolution:
result = [], current = [1,3]
result = [], current = [1,6]    (merged with [2,6])
result = [[1,6]], current = [8,10]  (no overlap)
result = [[1,6], [8,10]], current = [15,18] (no overlap)
result = [[1,6], [8,10], [15,18]] ✓
```

### Time Complexity Breakdown:
- **Sorting**: O(N log N) time
- **Merging Pass**: O(N) time
- **Total Time**: O(N log N) dominated by sorting
- **Space**: O(N) for result + O(1) or O(N) for sorting
- **Optimal**: Sorting is required for this problem

### Alternative Approaches:

#### 1. Brute Force Pair Comparison (O(N²) time, O(N) space)
```go
func mergeBruteForce(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }
    
    // Sort intervals by start time first
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    var result [][]int
    i := 0
    for i < len(intervals) {
        current := intervals[i]
        
        // Find all intervals that overlap with current
        j := i + 1
        for j < len(intervals) && intervals[j][0] <= current[1] {
            current[1] = max(current[1], intervals[j][1])
            j++
        }
        
        result = append(result, current)
        i = j
    }
    
    return result
}
```
- **Pros**: Simple to understand
- **Cons**: O(N²) time in worst case

#### 2. In-place Modification (O(N log N) time, O(1) space)
```go
func mergeInPlace(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }
    
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    mergedIndex := 0
    for i := 1; i < len(intervals); i++ {
        if intervals[mergedIndex][1] >= intervals[i][0] {
            intervals[mergedIndex][1] = max(intervals[mergedIndex][1], intervals[i][1])
        } else {
            mergedIndex++
            intervals[mergedIndex] = intervals[i]
        }
    }
    
    return intervals[:mergedIndex+1]
}
```
- **Pros**: O(1) extra space
- **Cons**: Modifies input, more error-prone

#### 3. Counting Sort (O(N) time, O(R) space)
```go
func mergeCountingSort(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }
    
    // Use counting sort if start times are in small range
    maxStart := 0
    for _, interval := range intervals {
        maxStart = max(maxStart, interval[0])
    }
    
    // Counting sort implementation
    // ... (omitted for brevity)
    
    return mergeSorted(intervals)
}
```
- **Pros**: O(N) time for small ranges
- **Cons**: Only works for bounded integer ranges

### Extensions for Interviews:
- **Different Overlap Definitions**: Strict vs non-strict overlap
- **Interval Insertion**: Add new interval to already merged list
- **Range Queries**: Query how many intervals overlap with point
- **Dynamic Updates**: Support insertions and deletions
- **Performance Analysis**: Discuss worst-case scenarios
*/
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
