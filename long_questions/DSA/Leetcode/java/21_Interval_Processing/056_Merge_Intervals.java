import java.util.*;

public class MergeIntervals {
    
    // 56. Merge Intervals
    // Time: O(N log N), Space: O(N)
    public static int[][] merge(int[][] intervals) {
        if (intervals.length <= 1) {
            return intervals;
        }
        
        // Sort intervals by start time
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        
        List<int[]> result = new ArrayList<>();
        int[] current = intervals[0];
        
        for (int i = 1; i < intervals.length; i++) {
            int[] next = intervals[i];
            
            // Check if intervals overlap
            if (current[1] >= next[0]) {
                // Merge intervals
                current[1] = Math.max(current[1], next[1]);
            } else {
                // Add current interval to result and start new current
                result.add(new int[]{current[0], current[1]});
                current = next;
            }
        }
        
        // Add the last interval
        result.add(new int[]{current[0], current[1]});
        
        return result.toArray(new int[result.size()][]);
    }

    // Alternative approach using in-place modification
    public static int[][] mergeInPlace(int[][] intervals) {
        if (intervals.length <= 1) {
            return intervals;
        }
        
        // Sort intervals by start time
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        
        int mergedIndex = 0;
        
        for (int i = 1; i < intervals.length; i++) {
            // If current interval overlaps with the last merged interval
            if (intervals[mergedIndex][1] >= intervals[i][0]) {
                // Merge them
                intervals[mergedIndex][1] = Math.max(intervals[mergedIndex][1], intervals[i][1]);
            } else {
                // Move to next position
                mergedIndex++;
                intervals[mergedIndex] = intervals[i];
            }
        }
        
        return Arrays.copyOf(intervals, mergedIndex + 1);
    }

    // Using a custom interval type
    public static class Interval {
        int start, end;
        
        public Interval(int start, int end) {
            this.start = start;
            this.end = end;
        }
        
        @Override
        public String toString() {
            return "[" + start + "," + end + "]";
        }
    }

    public static List<Interval> mergeTyped(List<Interval> intervals) {
        if (intervals.size() <= 1) {
            return intervals;
        }
        
        // Sort intervals by start time
        intervals.sort(Comparator.comparingInt(a -> a.start));
        
        List<Interval> result = new ArrayList<>();
        Interval current = intervals.get(0);
        
        for (int i = 1; i < intervals.size(); i++) {
            Interval next = intervals.get(i);
            
            if (current.end >= next.start) {
                // Merge intervals
                current.end = Math.max(current.end, next.end);
            } else {
                // Add current interval to result and start new current
                result.add(current);
                current = next;
            }
        }
        
        // Add the last interval
        result.add(current);
        
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[][][] testCases = {
            {{1, 3}, {2, 6}, {8, 10}, {15, 18}},
            {{1, 4}, {4, 5}},
            {{1, 3}, {2, 4}, {5, 7}, {6, 8}},
            {{1, 2}, {3, 4}, {5, 6}},
            {{1, 10}, {2, 3}, {4, 5}, {6, 7}},
            {{1, 5}, {2, 3}, {4, 6}},
            {{1, 4}, {0, 2}, {3, 5}},
            {{1, 3}, {5, 7}, {9, 11}},
            {{1, 4}, {2, 5}, {7, 9}},
            {{}}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            // Make copies for different approaches
            int[][] intervals1 = new int[testCases[i].length][];
            for (int j = 0; j < testCases[i].length; j++) {
                intervals1[j] = testCases[i][j].clone();
            }
            
            int[][] intervals2 = new int[testCases[i].length][];
            for (int j = 0; j < testCases[i].length; j++) {
                intervals2[j] = testCases[i][j].clone();
            }
            
            int[][] result1 = merge(testCases[i]);
            int[][] result2 = mergeInPlace(intervals2);
            
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.deepToString(testCases[i]));
            System.out.printf("  Standard: %s\n", Arrays.deepToString(result1));
            System.out.printf("  In-place: %s\n\n", Arrays.deepToString(result2));
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Interval Processing
- **Interval Merging**: Combine overlapping intervals
- **Sorting by Start**: Process intervals chronologically
- **Greedy Merging**: Always merge with current if overlapping
- **Boundary Tracking**: Maintain current merged interval boundaries

## 2. PROBLEM CHARACTERISTICS
- **Interval Representation**: [start, end] pairs
- **Overlap Detection**: current.end >= next.start
- **Merging Strategy**: Extend current interval to include next
- **Sorted Processing**: Process in chronological order

## 3. SIMILAR PROBLEMS
- Insert Interval
- Non-overlapping Intervals
- Meeting Rooms
- Calendar Conflicts

## 4. KEY OBSERVATIONS
- Sorting by start time ensures chronological processing
- Overlap condition: current.end >= next.start
- Merged interval: start = current.start, end = max(current.end, next.end)
- In-place modification saves space
- Time complexity dominated by sorting: O(N log N)

## 5. VARIATIONS & EXTENSIONS
- Different overlap conditions (strict vs inclusive)
- Interval intersection instead of merging
- Finding maximum overlap
- Weighted intervals

## 6. INTERVIEW INSIGHTS
- Clarify: "Are intervals inclusive or exclusive?"
- Edge cases: empty array, single interval, all overlapping
- Time complexity: O(N log N) vs O(N²) naive
- Space complexity: O(1) vs O(N) extra space

## 7. COMMON MISTAKES
- Not sorting intervals first
- Incorrect overlap condition
- Not updating current interval correctly
- Forgetting to add last interval
- Off-by-one errors in boundary checks

## 8. OPTIMIZATION STRATEGIES
- In-place modification saves O(N) space
- Early termination for non-overlapping intervals
- Custom comparator for sorting
- Merge during single pass

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like scheduling meetings:**
- You have meetings (intervals) with start and end times
- You want to merge overlapping meetings into longer meetings
- Process meetings in chronological order
- When current meeting overlaps with next, merge them
- Continue until all meetings processed
- Result is schedule with no conflicts

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of [start, end] intervals
2. **Goal**: Merge all overlapping intervals
3. **Output**: Array of non-overlapping merged intervals

#### Phase 2: Key Insight Recognition
- **"How to detect overlap?"** → current.end >= next.start
- **"How to merge?"** → start = current.start, end = max(current.end, next.end)
- **"What order to process?"** → Sort by start time
- **"Why sorting works?"** → Chronological order ensures correct merging

#### Phase 3: Strategy Development
```
Human thought process:
"I'll merge intervals:
1. Sort intervals by start time
2. Initialize current = first interval
3. For each next interval:
   - If current overlaps next: merge them
   - Else: add current to result, start new current
4. Add final current to result"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty array
- **Single interval**: Return that interval
- **All overlapping**: Return one merged interval
- **No overlapping**: Return all original intervals

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Intervals: [[1,3], [2,6], [8,10], [15,18]]

Human thinking:
"Sort by start time:
[1,3], [2,6], [8,10], [15,18] (already sorted)

Process:
current = [1,3], result = []

next = [2,6]:
- current.end(3) >= next.start(2)? Yes (3 >= 2)
- Merge: [1, max(3,6)] = [1,6]
- current = [1,6]

next = [8,10]:
- current.end(6) >= next.start(8)? No (6 < 8)
- Add [1,6] to result: result = [[1,6]]
- current = [8,10]

next = [15,18]:
- current.end(10) >= next.start(15)? No (10 < 15)
- Add [8,10] to result: result = [[1,6], [8,10]]
- current = [15,18]

Add final current: result = [[1,6], [8,10], [15,18]] ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Sorting ensures chronological processing
- **Why it's efficient**: Single pass after sorting
- **Why it's correct**: All overlapping intervals are merged

### Common Human Pitfalls & How to Avoid Them
1. **"Why not try all pairs?"** → O(N²) vs O(N log N)
2. **"What about unsorted?"** → Overlap detection fails
3. **"How to handle boundaries?"** → Be careful with >= vs >
4. **"What about in-place?"** → Saves space but modifies input

### Real-World Analogy
**Like consolidating meeting rooms:**
- You have room bookings (intervals)
- Some bookings overlap and can be consolidated
- Sort bookings by start time
- When current booking overlaps with next, merge into longer booking
- This frees up rooms for other bookings
- Final schedule has no overlapping bookings

### Human-Readable Pseudocode
```
function merge(intervals):
    if intervals.length <= 1:
        return intervals
    
    // Sort by start time
    sort(intervals, (a, b) -> a[0] - b[0])
    
    result = []
    current = intervals[0]
    
    for i from 1 to intervals.length-1:
        next = intervals[i]
        
        if current[1] >= next[0]:  // Overlap
            current[1] = max(current[1], next[1])
        else:
            result.add(current)
            current = next
    
    result.add(current)
    return result
```

### Execution Visualization

### Example: intervals = [[1,3], [2,6], [8,10], [15,18]]
```
Sorting by start:
[1,3], [2,6], [8,10], [15,18] (already sorted)

Merging Process:
current = [1,3], result = []

next = [2,6]: Overlap (3 >= 2)
→ Merge to [1,6], current = [1,6]

next = [8,10]: No overlap (6 < 8)
→ Add [1,6] to result, current = [8,10]

next = [15,18]: No overlap (10 < 15)
→ Add [8,10] to result, current = [15,18]

Final: [[1,6], [8,10], [15,18]] ✓

Visualization:
[1,3] + [2,6] = [1,6]  (merged)
[8,10] (unchanged)
[15,18] (unchanged)
```

### Key Visualization Points:
- **Sorting** ensures chronological processing
- **Overlap detection**: current.end >= next.start
- **Merging**: Extend current to include next
- **Result building**: Add when no overlap

### Memory Layout Visualization:
```
Processing Evolution:
Step 1: current=[1,3], next=[2,6]
        overlap: 3>=2 ✓
        merge: [1,6]
        current=[1,6]

Step 2: current=[1,6], next=[8,10]
        overlap: 6>=8 ✗
        add: [1,6] to result
        current=[8,10]

Step 3: current=[8,10], next=[15,18]
        overlap: 10>=15 ✗
        add: [8,10] to result
        current=[15,18]

Step 4: current=[15,18], next=none
        add: [15,18] to result

Final: [[1,6], [8,10], [15,18]] ✓
```

### Time Complexity Breakdown:
- **Sorting**: O(N log N) time
- **Merging**: O(N) single pass
- **Total**: O(N log N) time, O(1) extra space
- **Optimal**: Cannot do better than sorting for this problem
- **vs Naive**: O(N²) checking all pairs
*/
