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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Greedy Interval Scheduling
- **Greedy Selection**: Choose intervals with earliest end times
- **Non-overlapping Constraint**: Each chosen interval must not overlap previous
- **Sorting Strategy**: Sort by end time to enable greedy selection
- **Maximum Count**: Find maximum number of non-overlapping intervals

## 2. PROBLEM CHARACTERISTICS
- **Interval Selection**: Choose maximum subset of non-overlapping intervals
- **Optimization Goal**: Maximize number of intervals selected
- **Greedy Validity**: Earliest end time strategy is optimal
- **Scheduling Problem**: Classic activity selection problem

## 3. SIMILAR PROBLEMS
- Meeting Rooms (LeetCode 252) - Minimum rooms needed
- Course Schedule III (LeetCode 630) - Maximum courses with deadlines
- Activity Selection Problem - Classic greedy algorithm
- Job Scheduling with Deadlines - Weighted interval selection

## 4. KEY OBSERVATIONS
- **Greedy Optimal**: Choosing earliest finishing intervals maximizes remaining time
- **End Time Sorting**: Essential for greedy strategy to work
- **Non-overlap Check**: Current start >= last selected end
- **Linear Selection**: Single pass through sorted intervals

## 5. VARIATIONS & EXTENSIONS
- **Weighted Intervals**: Maximize total weight instead of count
- **Different Resources**: Multiple machines/rooms available
- **Interval Insertion**: Add new intervals dynamically
- **Range Queries**: Query maximum intervals in given range

## 6. INTERVIEW INSIGHTS
- Always clarify: "Goal is count or sum? Intervals sorted? Overlap definition?"
- Edge cases: empty list, single interval, all overlapping, none overlapping
- Time complexity: O(N log N) for sorting + O(N) for selection
- Space complexity: O(1) extra space
- Key insight: greedy by earliest end time is optimal

## 7. COMMON MISTAKES
- Sorting by start time instead of end time
- Wrong overlap condition (using > instead of >=)
- Not handling edge cases (empty list, single interval)
- Greedy selection without proper sorting
- Confusing with minimum intervals removal problem

## 8. OPTIMIZATION STRATEGIES
- **Greedy by End Time**: O(N log N) time, O(1) space - optimal
- **Dynamic Programming**: O(N²) time, O(N) space - for weighted version
- **Binary Indexed Tree**: O(N log N) for dynamic operations
- **Segment Tree**: O(N log N) for range queries

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like scheduling the maximum number of meetings:**
- You have many meeting options with different times
- Want to attend as many meetings as possible
- Strategy: Always attend the meeting that finishes earliest
- This leaves maximum time for remaining meetings
- Like a greedy scheduler optimizing for quantity

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of intervals [start, end]
2. **Goal**: Find maximum number of non-overlapping intervals
3. **Constraint**: Cannot attend overlapping meetings
4. **Output**: Maximum count of intervals that can be selected

#### Phase 2: Key Insight Recognition
- **"Greedy natural fit"** → Choose earliest finishing meetings
- **"End time sorting"** → Critical for greedy strategy
- **"Non-overlap constraint"** → Current start >= last selected end
- **"Optimal proof"** → Earliest finish maximizes remaining opportunities

#### Phase 3: Strategy Development
```
Human thought process:
"I need to select maximum non-overlapping intervals.
If I sort by end time and always pick the earliest finishing:
1. This leaves maximum time for remaining meetings
2. Any interval that starts after my last selected meeting can be taken
3. This greedy choice is provably optimal!

Algorithm:
1. Sort intervals by end time
2. Initialize count = 0, lastEnd = -infinity
3. For each interval:
   - If interval.start >= lastEnd:
     - Select it: count++, lastEnd = interval.end
4. Return count"
```

#### Phase 4: Edge Case Handling
- **Empty list**: Return 0
- **Single interval**: Return 1
- **All overlapping**: Return 1 (can only pick one)
- **None overlapping**: Return total count

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Intervals: [[1,2], [2,3], [3,4], [1,6]]

Human thinking:
"Sort by end time:
[[1,2], [2,3], [3,4], [1,6]] ✓

Greedy selection:
count = 0, lastEnd = -infinity

[1,2]: start=1 >= -infinity? ✓ Select it
count = 1, lastEnd = 2

[2,3]: start=2 >= lastEnd=2? No overlap ✓ Select it
count = 2, lastEnd = 3

[3,4]: start=3 >= lastEnd=3? No overlap ✓ Select it
count = 3, lastEnd = 4

[1,6]: start=1 >= lastEnd=4? No overlap ✓ Select it
count = 4, lastEnd = 6

Maximum non-overlapping = 4 ✓"
```

#### Phase 6: Intuition Validation
- **Why greedy works**: Earliest finish maximizes remaining time
- **Why end time sorting**: Essential for greedy choice
- **Why optimal**: Can be proven by exchange argument
- **Why O(N log N)**: Sorting dominates, selection is linear

### Common Human Pitfalls & How to Avoid Them
1. **"Why not sort by start time?"** → Wrong greedy choice
2. **"Should I use DP?"** → O(N²) vs O(N log N), unnecessary for unweighted
3. **"What about weighted intervals?"** → Greedy fails, need DP
4. **"Can I optimize further?"** → Greedy is already optimal
5. **"What about minimum removals?"** → Different problem (LeetCode 435)

### Real-World Analogy
**Like scheduling the maximum number of appointments:**
- You have many possible appointments with different time slots
- Want to book as many appointments as possible
- Strategy: Always book the appointment that finishes earliest
- This leaves maximum time for remaining appointments
- Like a busy professional trying to maximize client meetings

### Human-Readable Pseudocode
```
function maxNonOverlappingIntervals(intervals):
    if intervals is empty:
        return 0
    
    sort intervals by end time
    
    count = 0
    lastEnd = -infinity
    
    for each interval in intervals:
        if interval.start >= lastEnd:
            count++
            lastEnd = interval.end
    
    return count
```

### Execution Visualization

### Example: Intervals = [[1,2], [2,3], [3,4], [1,6]]
```
Sorting Phase:
Original: [[1,2], [2,3], [3,4], [1,6]]
Sorted:   [[1,2], [2,3], [3,4], [1,6]] ✓

Greedy Selection:
count = 0, lastEnd = -∞

[1,2]: start=1 >= -∞? ✓ Select
count = 1, lastEnd = 2

[2,3]: start=2 >= 2? ✓ Select
count = 2, lastEnd = 3

[3,4]: start=3 >= 3? ✓ Select
count = 3, lastEnd = 4

[1,6]: start=1 >= 4? No overlap ✓ Select
count = 4, lastEnd = 6

Final result: 4 non-overlapping intervals ✓
```

### Key Visualization Points:
- **End Time Sorting**: Critical for greedy strategy
- **Non-overlap Check**: interval.start >= lastEnd
- **Greedy Choice**: Always pick earliest finishing
- **Linear Selection**: Single pass after sorting

### Memory Layout Visualization:
```
Interval Selection State:
intervals = [[1,2], [2,3], [3,4], [1,6]]

Step-by-step selection:
count = 0, lastEnd = -∞

Process [1,2]: start=1 >= -∞ ✓
count = 1, lastEnd = 2

Process [2,3]: start=2 >= 2 ✓
count = 2, lastEnd = 3

Process [3,4]: start=3 >= 3 ✓
count = 3, lastEnd = 4

Process [1,6]: start=1 >= 4 ✓
count = 4, lastEnd = 6

Final: count = 4 ✓
```

### Time Complexity Breakdown:
- **Sorting**: O(N log N) time
- **Greedy Selection**: O(N) time
- **Total Time**: O(N log N) dominated by sorting
- **Space**: O(1) extra space (in-place sorting possible)
- **Optimal**: Greedy is optimal for unweighted case

### Alternative Approaches:

#### 1. Dynamic Programming (O(N²) time, O(N) space)
```go
func maxNonOverlappingDP(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }
    
    // Sort by start time for DP
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    n := len(intervals)
    dp := make([]int, n)
    
    for i := 0; i < n; i++ {
        dp[i] = 1 // Each interval by itself
        for j := 0; j < i; j++ {
            if intervals[j][1] <= intervals[i][0] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
    }
    
    return max(dp...)
}
```
- **Pros**: Works for weighted intervals
- **Cons**: O(N²) time, unnecessary for unweighted case

#### 2. Sort by Start Time with Greedy (O(N log N) time, O(1) space)
```go
func maxNonOverlappingStartSort(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    count := 0
    lastEnd := -1
    
    for _, interval := range intervals {
        if interval[0] > lastEnd {
            count++
            lastEnd = interval[1]
        }
    }
    
    return count
}
```
- **Pros**: Same complexity
- **Cons**: Less intuitive, may miss optimal in some cases

#### 3. Recursive with Memoization (O(N²) time, O(N) space)
```go
func maxNonOverlappingRecursive(intervals [][]int) int {
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
        
        // Option 1: Skip current interval
        skip := solve(index + 1)
        
        // Option 2: Take current interval
        take := 1
        nextIndex := index + 1
        for nextIndex < len(intervals) && intervals[nextIndex][0] < intervals[index][1] {
            nextIndex++
        }
        take += solve(nextIndex)
        
        memo[index] = max(skip, take)
        return memo[index]
    }
    
    return solve(0)
}
```
- **Pros**: More flexible for variations
- **Cons**: O(N²) time, more complex

### Extensions for Interviews:
- **Weighted Intervals**: Maximize total weight instead of count
- **Different Resources**: Multiple machines/rooms available
- **Interval Insertion**: Add new intervals dynamically
- **Range Queries**: Query maximum intervals in given range
- **Performance Analysis**: Discuss worst-case scenarios
*/
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
