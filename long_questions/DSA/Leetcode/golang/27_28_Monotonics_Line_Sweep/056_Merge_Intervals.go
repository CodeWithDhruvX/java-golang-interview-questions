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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Line Sweep for Interval Merging
- **Event-Based Processing**: Convert intervals to start/end events
- **Sorting**: Sort events by position, start before end
- **Active Count Tracking**: Track number of overlapping intervals
- **Interval Reconstruction**: Build merged intervals from active periods

## 2. PROBLEM CHARACTERISTICS
- **Interval Overlap**: Need to merge overlapping intervals
- **Sorting Required**: Natural ordering for line sweep
- **Event Processing**: Convert intervals to discrete events
- **Output Construction**: Build merged intervals from active periods

## 3. SIMILAR PROBLEMS
- Insert Interval (LeetCode 57) - Insert interval into sorted list
- Meeting Rooms (LeetCode 252) - Find minimum meeting rooms
- Calendar Integration (LeetCode 759) - Merge worker schedules
- Skyline Problem (LeetCode 218) - Complex interval merging

## 4. KEY OBSERVATIONS
- **Line Sweep Natural**: Events naturally model interval overlaps
- **Sorting Critical**: Proper event ordering essential
- **Active Count**: When count goes 0→1, start new interval
- **End Detection**: When count goes 1→0, end current interval

## 5. VARIATIONS & EXTENSIONS
- **Priority Queue**: Handle intervals with different priorities
- **Counting Approach**: Use difference array for range queries
- **Tree-Based**: Use interval trees for dynamic operations
- **Segment Trees**: Handle range updates and queries

## 6. INTERVIEW INSIGHTS
- Always clarify: "Interval format? Overlap definition? Output order?"
- Edge cases: empty array, single interval, all overlapping
- Time complexity: O(N log N) for sorting, O(N) for sweep
- Space complexity: O(N) for events and result
- Key insight: line sweep converts overlap detection to counting

## 7. COMMON MISTAKES
- Wrong event sorting (end before start at same position)
- Incorrect active count management
- Not handling zero-length intervals properly
- Missing edge cases (empty, single interval)
- Off-by-one errors in interval boundaries

## 8. OPTIMIZATION STRATEGIES
- **Basic Line Sweep**: O(N log N) time, O(N) space - standard
- **Counting Approach**: O(N log N) time, O(N) space - for range queries
- **Two Pointers**: O(N log N) time, O(1) space - for sorted input
- **Priority Queue**: O(N log N) time, O(N) space - for dynamic operations

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like tracking meeting room occupancy:**
- You have meetings with start/end times (intervals)
- Convert each meeting to "meeting starts" and "meeting ends" events
- Sort all events chronologically
- Sweep through timeline, counting active meetings
- When count goes from 0→1, meeting starts
- When count goes from 1→0, meeting ends
- Merge overlapping meetings into continuous blocks

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of intervals [start, end]
2. **Goal**: Merge overlapping intervals
3. **Constraints**: Intervals may overlap, be nested, or be separate
4. **Output**: Array of non-overlapping merged intervals

#### Phase 2: Key Insight Recognition
- **"Line sweep natural fit"** → Events model interval overlaps perfectly
- **"Active counting"** → Track how many intervals are active
- **"Transition points"** → Start when count 0→1, end when 1→0
- **"Event ordering"** → Start events before end events at same position

#### Phase 3: Strategy Development
```
Human thought process:
"I need to merge overlapping intervals.
Direct pairwise comparison is O(N²).

Line Sweep Approach:
1. Convert intervals to start/end events
2. Sort events by position (start before end at same position)
3. Sweep through events:
   - Track active interval count
   - When count 0→1: start new merged interval
   - When count 1→0: end current merged interval
4. Build result from start/end transitions

This handles all overlap patterns!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty array
- **Single interval**: Return array with single interval
- **All overlapping**: Return single merged interval
- **No overlap**: Return original intervals (sorted)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Intervals: [[1,3], [2,6], [8,10], [15,18]]

Human thinking:
"Line Sweep Approach:
1. Create events:
   (1,start), (3,end), (2,start), (6,end), (8,start), (10,end), (15,start), (18,end)

2. Sort events:
   (1,start), (2,start), (3,end), (6,end), (8,start), (10,end), (15,start), (18,end)

3. Sweep:
   pos=1: start, count=1, start=[1]
   pos=2: start, count=2
   pos=3: end, count=1
   pos=6: end, count=0, end=[6] → add [1,6]
   pos=8: start, count=1, start=[8]
   pos=10: end, count=0, end=[10] → add [8,10]
   pos=15: start, count=1, start=[15]
   pos=18: end, count=0, end=[18] → add [15,18]

Result: [[1,6], [8,10], [15,18]] ✓"
```

#### Phase 6: Intuition Validation
- **Why events work**: Convert continuous intervals to discrete points
- **Why sorting works**: Process events in chronological order
- **Why counting works**: Active count tells if we're inside or outside intervals
- **Why transitions matter**: Count changes indicate interval boundaries

### Common Human Pitfalls & How to Avoid Them
1. **"Why not pairwise comparison?"** → O(N²) vs O(N log N)
2. **"Should I use two pointers?"** → Only works if input already sorted
3. **"What about nested intervals?"** → Line sweep handles naturally
4. **"Can I optimize further?"** → O(N log N) is optimal for unsorted input
5. **"What about floating point?"** → Same approach works with decimal values

### Real-World Analogy
**Like managing meeting room schedules:**
- You have meeting requests with start/end times
- Convert each meeting to "meeting starts" and "meeting ends" notifications
- Sort all notifications chronologically
- Track how many meetings are currently active
- When first meeting starts, mark room as occupied
- When all meetings end, mark room as free
- Merge consecutive occupied periods into continuous blocks
- Like building a master schedule from individual bookings

### Human-Readable Pseudocode
```
function mergeIntervals(intervals):
    if len(intervals) <= 1:
        return intervals
    
    // Create events
    events = []
    for interval in intervals:
        events.append((interval[0], 'start'))
        events.append((interval[1], 'end'))
    
    // Sort events by position, start before end
    sort(events, key=lambda e: (e.position, not e.is_start))
    
    result = []
    active_count = 0
    current_start = -1
    
    for event in events:
        if event.is_start:
            if active_count == 0:
                current_start = event.position
            active_count += 1
        else:
            active_count -= 1
            if active_count == 0:
                result.append([current_start, event.position])
    
    return result
```

### Execution Visualization

### Example: Intervals = [[1,3], [2,6], [8,10], [15,18]]
```
Event Creation:
[1,3] → (1,start), (3,end)
[2,6] → (2,start), (6,end)
[8,10] → (8,start), (10,end)
[15,18] → (15,start), (18,end)

Sorted Events:
(1,start), (2,start), (3,end), (6,end), (8,start), (10,end), (15,start), (18,end)

Line Sweep Process:
pos=1: start, count=1, current_start=1
pos=2: start, count=2
pos=3: end, count=1
pos=6: end, count=0, add [1,6] to result
pos=8: start, count=1, current_start=8
pos=10: end, count=0, add [8,10] to result
pos=15: start, count=1, current_start=15
pos=18: end, count=0, add [15,18] to result

Final Result: [[1,6], [8,10], [15,18]] ✓
```

### Key Visualization Points:
- **Event Creation**: Each interval becomes two events
- **Event Sorting**: Chronological processing order
- **Active Count**: Tracks overlapping intervals
- **Transitions**: Count changes mark interval boundaries

### Memory Layout Visualization:
```
Event Processing Flow:
Initial: count=0, current_start=-1, result=[]

pos=1: (1,start) → count=1, current_start=1
pos=2: (2,start) → count=2
pos=3: (3,end) → count=1
pos=6: (6,end) → count=0, result=[[1,6]]
pos=8: (8,start) → count=1, current_start=8
pos=10: (10,end) → count=0, result=[[1,6], [8,10]]
pos=15: (15,start) → count=1, current_start=15
pos=18: (18,end) → count=0, result=[[1,6], [8,10], [15,18]]

Active Count Evolution:
0→1→2→1→0→1→0→1→0
Each transition marks interval boundary
```

### Time Complexity Breakdown:
- **Event Creation**: O(N) time, O(N) space
- **Event Sorting**: O(N log N) time, O(N) space
- **Line Sweep**: O(N) time, O(1) additional space
- **Total**: O(N log N) time, O(N) space
- **Optimal**: Cannot do better than sorting unsorted intervals

### Alternative Approaches:

#### 1. Two Pointers (O(N log N) time, O(1) space)
```go
func mergeTwoPointers(intervals [][]int) [][]int {
    // Sort intervals first, then merge in one pass
    // Only works if we can modify or sort input
    // ... implementation details omitted
}
```
- **Pros**: O(1) extra space after sorting
- **Cons**: Requires sorted input or sorting step

#### 2. Counting Array (O(N log N) time, O(N) space)
```go
func mergeCounting(intervals [][]int) [][]int {
    // Use difference array for range queries
    // Good for multiple range queries
    // ... implementation details omitted
}
```
- **Pros**: Efficient for multiple queries
- **Cons**: More complex, coordinate compression needed

#### 3. Interval Tree (O(N log N) time, O(N) space)
```go
func mergeIntervalTree(intervals [][]int) [][]int {
    // Build interval tree for dynamic operations
    // Overkill for simple merging
    // ... implementation details omitted
}
```
- **Pros**: Supports dynamic operations
- **Cons**: Complex implementation, overhead for simple case

### Extensions for Interviews:
- **Dynamic Operations**: Support for insert/delete interval queries
- **Range Queries**: Answer "how many intervals overlap point X"
- **Priority Intervals**: Handle intervals with different priorities
- **Multi-dimensional**: Extend to 2D rectangle merging
- **Performance Analysis**: Discuss cache behavior and practical considerations
*/
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
