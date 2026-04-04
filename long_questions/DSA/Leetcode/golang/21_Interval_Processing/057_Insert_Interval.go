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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Insert and Merge for Interval Management
- **Insertion Strategy**: Find correct position and insert new interval
- **Merge Logic**: Combine overlapping intervals after insertion
- **Linear Search**: Find insertion point by scanning sorted intervals
- **Three-Phase Process**: Insert before, merge overlapping, insert after

## 2. PROBLEM CHARACTERISTICS
- **Dynamic Insertion**: Add new interval to existing sorted list
- **Overlap Handling**: Merge new interval with any overlapping ones
- **Sorted Maintenance**: Keep intervals in sorted order after insertion
- **In-place Updates**: Modify existing array or create new one

## 3. SIMILAR PROBLEMS
- Merge Intervals (LeetCode 56) - Merge all overlapping intervals
- Non-overlapping Intervals (LeetCode 435) - Remove intervals to make non-overlapping
- Meeting Rooms (LeetCode 252) - Find minimum meeting rooms needed
- Employee Free Time (LeetCode 759) - Find available time slots

## 4. KEY OBSERVATIONS
- **Insertion Point**: Find first interval with end >= new.start
- **Merge Before**: Merge new interval with intervals that end before it starts
- **Merge After**: Merge new interval with intervals that start before it ends
- **Result Construction**: Combine non-overlapping before + merged + non-overlapping after

## 5. VARIATIONS & EXTENSIONS
- **Binary Search Insertion**: O(log N) insertion point finding
- **In-place Modification**: O(1) extra space, but modifies input
- **Range Queries**: Query intervals that overlap with given point
- **Dynamic Operations**: Support multiple insertions and deletions

## 6. INTERVIEW INSIGHTS
- Always clarify: "Intervals already sorted? In-place allowed? Overlap definition?"
- Edge cases: empty list, new interval before/after all, complete overlap
- Time complexity: O(N) linear search, O(log N) with binary search
- Space complexity: O(N) for new array, O(1) for in-place
- Key insight: separate insertion and merging phases

## 7. COMMON MISTAKES
- Wrong insertion point calculation
- Not handling all overlapping intervals correctly
- Forgetting to merge intervals that overlap new interval
- Not maintaining sorted order after insertion
- Off-by-one errors in index calculations

## 8. OPTIMIZATION STRATEGIES
- **Linear Insertion**: O(N) time, O(N) space - simple approach
- **Binary Search Insertion**: O(log N) time, O(N) space - faster insertion
- **In-place Modification**: O(N) time, O(1) space - memory efficient
- **Segment Tree**: O(log N) operations, O(N) space - for many queries

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like inserting a new meeting into an existing schedule:**
- You have a sorted list of existing meetings
- Need to insert a new meeting in the right place
- If the new meeting overlaps with existing ones, merge them
- Result is a clean, non-overlapping schedule that includes the new meeting
- Like inserting a puzzle piece and merging with adjacent pieces

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Sorted intervals array, new interval to insert
2. **Goal**: Insert new interval and merge any overlaps
3. **Constraint**: Maintain sorted order in final result
4. **Output**: New array with merged, non-overlapping intervals

#### Phase 2: Key Insight Recognition
- **"Three-phase natural fit"** → Insert before, merge overlapping, insert after
- **"Linear insertion sufficient"** → Find insertion point by scanning
- **"Separate merge phases"** → Handle intervals before and after insertion point
- **"Sorted order maintenance"** → Critical for correctness

#### Phase 3: Strategy Development
```
Human thought process:
"I need to insert a new interval into sorted intervals and merge overlaps.
I can do this in three phases:
1. Find where to insert: first interval with end >= new.start
2. Add all intervals that end before new.start (no overlap)
3. Merge new interval with all overlapping intervals
4. Add remaining intervals that start after new.end

This ensures I handle all overlaps correctly!"
```

#### Phase 4: Edge Case Handling
- **Empty intervals**: Return just the new interval
- **New interval before all**: Insert at beginning
- **New interval after all**: Insert at end
- **Complete overlap**: Merge with all intervals into one

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Intervals: [[1,3], [6,9]], New: [2,5]

Human thinking:
"Phase 1: Find insertion point
- Scan: [1,3] (end=3) < new.start=2? No, continue
- [6,9] (end=9) >= new.start=2? Yes! Insertion point = before [6,9]

Phase 2: Add intervals ending before new.start
- [1,3] (end=3) < new.start=2? Yes, add to result
- Result: [[1,3]]

Phase 3: Merge new interval with overlapping intervals
- New [2,5] overlaps with [6,9]? new.end=5 >= [6,9].start=6? No
- Actually, [6,9] starts after new interval ends, so no overlap
- Wait, let me recheck overlap logic...

Actually, I think I need to be more careful:
Intervals ending before new.start: none (since [1,3].end=3 >= new.start=2)
Intervals starting before new.end: [1,3] (start=1 <= new.end=5) ✓ overlaps
[6,9] (start=6 <= new.end=5) ✗ no overlap

So merge [2,5] with [1,3]: merged = [1,5]

Phase 4: Add remaining intervals
- [6,9] starts after merged.end=5, so add as-is
- Final result: [[1,5], [6,9]] ✓"
```

#### Phase 6: Intuition Validation
- **Why three-phase works**: Separates concerns clearly
- **Why linear search works**: Intervals are already sorted
- **Why merge logic works**: Combines all overlapping intervals correctly
- **Why O(N)**: Single pass through intervals

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just append and sort?"** → O(N log N) vs O(N) for sorted input
2. **"Should I use binary search?"** → Faster insertion but more complex
3. **"What about in-place modification?"** → Saves space but error-prone
4. **"Can I optimize further?"** → Linear scan is already optimal for single insertion
5. **"What about multiple insertions?"** → Consider data structure for repeated operations

### Real-World Analogy
**Like scheduling a new meeting in an existing calendar:**
- You have a calendar with existing meetings in chronological order
- Need to add a new meeting at the right time slot
- If the new meeting overlaps with existing ones, combine them into a longer meeting
- Keep the calendar sorted and non-overlapping
- The result shows the updated schedule with the new meeting properly placed

### Human-Readable Pseudocode
```
function insertInterval(intervals, newInterval):
    if intervals is empty:
        return [newInterval]
    
    result = []
    i = 0
    n = len(intervals)
    
    // Phase 1: Add intervals ending before new interval starts
    while i < n and intervals[i][1] < newInterval[0]:
        result.append(intervals[i])
        i++
    
    // Phase 2: Merge new interval with overlapping intervals
    while i < n and intervals[i][0] <= newInterval[1]:
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    
    result.append(newInterval)
    
    // Phase 3: Add remaining intervals
    while i < n:
        result.append(intervals[i])
        i++
    
    return result
```

### Execution Visualization

### Example: Intervals = [[1,3], [6,9]], New = [2,5]
```
Insertion Process:
Phase 1: Find intervals ending before new.start=2
- [1,3]: end=3 < 2? ✗ No
- [6,9]: end=9 < 2? ✗ No
- No intervals end before new.start

Phase 2: Merge new with overlapping intervals
- Check [1,3]: start=1 <= new.end=5? ✓ Overlaps
  Merge: new = [min(2,1)=1, max(5,3)=5] = [1,5]
- Check [6,9]: start=6 <= new.end=5? ✗ No overlap
- Final merged: [1,5]

Phase 3: Add remaining intervals
- [6,9]: starts after merged.end=5, add as-is

Final result: [[1,5], [6,9]] ✓
```

### Key Visualization Points:
- **Three Phases**: Before, merge, after
- **Overlap Detection**: interval.start <= new.end AND interval.end >= new.start
- **Merge Logic**: min of starts, max of ends
- **Sorted Order**: Maintained throughout process

### Memory Layout Visualization:
```
Interval State During Insertion:
intervals = [[1,3], [6,9]], new = [2,5]

Phase 1: No intervals end before 2
result = []

Phase 2: Merge overlapping intervals
Check [1,3]: overlaps with [2,5]? ✓
merged = [1,5]
Check [6,9]: overlaps with [2,5]? ✗
result = [[1,5]]

Phase 3: Add remaining intervals
Add [6,9]: starts after merged.end=5
result = [[1,5], [6,9]] ✓
```

### Time Complexity Breakdown:
- **Linear Search**: O(N) time, O(N) space
- **Binary Search**: O(log N) for insertion + O(N) for merging
- **In-place**: O(N) time, O(1) extra space
- **Total**: O(N) time, O(N) space for new array

### Alternative Approaches:

#### 1. Append and Sort (O(N log N) time, O(N) space)
```go
func insertAndSort(intervals [][]int, newInterval []int) [][]int {
    intervals = append(intervals, newInterval)
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    return mergeIntervals(intervals)
}
```
- **Pros**: Simple to implement
- **Cons**: O(N log N) even if input was already sorted

#### 2. Binary Search Insertion (O(log N) time, O(N) space)
```go
func insertBinarySearch(intervals [][]int, newInterval []int) [][]int {
    if len(intervals) == 0 {
        return [][]int{newInterval}
    }
    
    // Find insertion position using binary search
    left, right := 0, len(intervals)
    for left < right {
        mid := left + (right-left)/2
        if intervals[mid][1] < newInterval[0] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    // Insert at position 'left'
    result := make([][]int, len(intervals)+1)
    copy(result, intervals[:left])
    result[left] = newInterval
    copy(result[left+1:], intervals[left:])
    
    return mergeIntervals(result)
}
```
- **Pros**: O(log N) insertion point finding
- **Cons**: More complex implementation

#### 3. Segment Tree (O(log N) operations, O(N) space)
```go
type SegmentTree struct {
    tree []Interval
    n    int
}

func (st *SegmentTree) Insert(interval Interval) {
    // Insert interval into segment tree
    // Merge with overlapping intervals in tree
}

func (st *SegmentTree) Query(point int) []Interval {
    // Query intervals that overlap with given point
}
```
- **Pros**: O(log N) for multiple operations
- **Cons**: Complex implementation, overhead for single operation

### Extensions for Interviews:
- **Binary Search Insertion**: O(log N) insertion point finding
- **In-place Modification**: O(1) extra space, but modifies input
- **Range Queries**: Query intervals that overlap with given point
- **Dynamic Operations**: Support multiple insertions and deletions
- **Performance Analysis**: Discuss worst-case scenarios
*/
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
