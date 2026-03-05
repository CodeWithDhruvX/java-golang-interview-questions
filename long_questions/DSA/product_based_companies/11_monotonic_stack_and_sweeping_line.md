# Monotonic Stacks & Sweeping Line (Product-Based Companies)

Two of the most specific, highly recognizable patterns at companies like Meta (Facebook) and Amazon.
- **Monotonic Stacks** are used to solve "Next Greater/Smaller Element" questions.
- **Sweeping Line** (often combined with sorting or Priority Queues) solves complex interval overlap problems (scheduling, geometry).

## Question 1: Daily Temperatures (Monotonic Stack Entry Level)
**Problem Statement:** Given an array of integers `temperatures` represents the daily temperatures, return an array `answer` such that `answer[i]` is the number of days you have to wait after the `i-th` day to get a warmer temperature. If there is no future day for which this is possible, keep `answer[i] == 0` instead.

### Answer:
We use a **decreasing monotonic stack**. The stack will store *indices* of the array. As we iterate through the temperatures, if we find a day that is warmer than the day at the `peek()` of our stack, we know we've found the "next greater temperature" for the `peek()` day. We pop it, calculate the difference in indices, and repeat.

**Code Implementation (Java):**
```java
import java.util.Stack;

public class DailyTemperatures {
    public int[] dailyTemperatures(int[] temperatures) {
        int[] ans = new int[temperatures.length];
        Stack<Integer> stack = new Stack<>(); // Store indices, not values
        
        for (int i = 0; i < temperatures.length; i++) {
            int currentTemp = temperatures[i];
            
            // While current temperature is WARMER than previous ones in stack
            while (!stack.isEmpty() && currentTemp > temperatures[stack.peek()]) {
                int prevDayIndex = stack.pop();
                // Number of days waited is difference in index
                ans[prevDayIndex] = i - prevDayIndex;
            }
            
            // Push current day onto the stack
            stack.push(i);
        }
        
        return ans;
    }
}
```
**Time Complexity:** O(N). Each element is pushed onto and popped from the stack at most once.
**Space Complexity:** O(N) worst case if temperatures are in strictly decreasing order.

---

## Question 2: Largest Rectangle in Histogram (Hard Monotonic Stack)
**Problem Statement:** Given an array of integers `heights` representing the histogram's bar height where the width of each bar is `1`, return the area of the largest rectangle in the histogram.

### Answer:
This is the pinnacle of Monotonic Stack problems. We maintain an **increasing monotonic stack** of indices. For a bar at index `i`, we can expand left until we see a strictly smaller bar, and right until we see a strictly smaller bar.
When we encounter a bar smaller than the `peek()`, it means the `peek()` bar can't expand right anymore. So we pop it, calculate its width using the current index `i` (right boundary) and the *new* `peek()` (left boundary), and compute its area.

**Code Implementation (Java):**
```java
import java.util.Stack;

public class LargestRectangle {
    public int largestRectangleArea(int[] heights) {
        int maxArea = 0;
        Stack<Integer> stack = new Stack<>();
        int n = heights.length;
        
        for (int i = 0; i <= n; i++) {
            // Trick: use a height of 0 at the end to force the stack to empty
            int h = (i == n) ? 0 : heights[i];
            
            // If the current height breaks the increasing property
            while (!stack.isEmpty() && h < heights[stack.peek()]) {
                int height = heights[stack.pop()];
                
                // If stack is empty, it means this popped height is the smallest so far,
                // so the width spans from index 0 all the way to `i-1`.
                // Otherwise, it spans from the next element in the stack up to `i-1`.
                int width = stack.isEmpty() ? i : i - stack.peek() - 1;
                maxArea = Math.max(maxArea, height * width);
            }
            stack.push(i);
        }
        return maxArea;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(N)

---

## Question 3: Meeting Rooms II (Sweeping Line / Min-Heap)
**Problem Statement:** Given an array of meeting time intervals `intervals` where `intervals[i] = [start_i, end_i]`, return the minimum number of conference rooms required.

### Answer:
We can solve this using either a Min-Heap or a Chronological Sweeping array.
**Sweeping Approach (Events):** Split intervals into `(start, +1 room)` and `(end, -1 room)` events. Sort these events by time. If an end and start happen at the same time, process the end first (free the room before taking it). A running sum tracks the active rooms required right now; the max seen is the answer.

**Code Implementation (Java):**
```java
import java.util.Arrays;

public class MeetingRoomsII {
    // Array sorting Sweep Line Approach
    public int minMeetingRooms(int[][] intervals) {
        if (intervals == null || intervals.length == 0) return 0;
        
        int n = intervals.length;
        int[] starts = new int[n];
        int[] ends = new int[n];
        
        for (int i = 0; i < n; i++) {
            starts[i] = intervals[i][0];
            ends[i] = intervals[i][1];
        }
        
        Arrays.sort(starts);
        Arrays.sort(ends);
        
        int rooms = 0;
        int maxRooms = 0;
        int s = 0, e = 0;
        
        // Similar to processing events on a timeline
        while (s < n) {
            if (starts[s] < ends[e]) {
                // A meeting starts before the earliest ending meeting finishes
                rooms++;
                maxRooms = Math.max(maxRooms, rooms);
                s++;
            } else {
                // Earliest ending meeting finishes. Free a room.
                rooms--;
                e++;
            }
        }
        
        return maxRooms;
    }
}
```
**Time Complexity:** O(N log N) for sorting.
**Space Complexity:** O(N) for `starts` and `ends` arrays.

---

## Question 4: The Skyline Problem (Hard Sweeping Line)
**Problem Statement:** A city's skyline is the outer contour of the silhouette formed by all the buildings. Given the locations and heights of all the buildings, return the skyline formed by these buildings collectively. (`buildings[i] = [left, right, height]`).

### Answer:
We define events: "Start of building" (push height to active pool) and "End of building" (remove height from active pool). 
We sweep across the x-axis. As our active set of heights grows/shrinks, if the **maximum height** across all currently active buildings changes, that x-coordinate marks a new point on the skyline contour.
Using a `PriorityQueue` allows finding the current `max` height efficiently, though a `TreeMap` is optimal so `remove` operations take `O(log N)` instead of `O(N)`.

**Code Implementation (Java):**
```java
import java.util.*;

public class TheSkylineProblem {
    public List<List<Integer>> getSkyline(int[][] buildings) {
        List<List<Integer>> result = new ArrayList<>();
        List<int[]> heights = new ArrayList<>();
        
        // Transform buildings to start and end height events
        for (int[] b : buildings) {
            // Negative height distinguishes start events from end events 
            // and ensures starts process before ends at the same x coordinate
            heights.add(new int[]{b[0], -b[2]}); // Start
            heights.add(new int[]{b[1], b[2]});  // End
        }
        
        // Sort by x coordinate. If same x, sort by height event
        Collections.sort(heights, (a, b) -> {
            if (a[0] != b[0]) return a[0] - b[0];
            return a[1] - b[1];
        });
        
        // Max Heap (TreeMap allows O(log N) deletion of non-top elements)
        // Key: Height, Value: Count of buildings at this height
        TreeMap<Integer, Integer> activeHeights = new TreeMap<>(Collections.reverseOrder());
        activeHeights.put(0, 1); // Ground is always active
        int prevMaxHeight = 0;
        
        for (int[] h : heights) {
            int x = h[0];
            int height = h[1];
            
            if (height < 0) { // It's a start event
                activeHeights.put(-height, activeHeights.getOrDefault(-height, 0) + 1);
            } else { // It's an end event
                int count = activeHeights.get(height);
                if (count == 1) {
                    activeHeights.remove(height);
                } else {
                    activeHeights.put(height, count - 1);
                }
            }
            
            int currentMaxHeight = activeHeights.firstKey();
            if (prevMaxHeight != currentMaxHeight) {
                // The skyline changed, record this point
                result.add(Arrays.asList(x, currentMaxHeight));
                prevMaxHeight = currentMaxHeight;
            }
        }
        
        return result;
    }
}
```
**Time Complexity:** O(N log N) processing overall
**Space Complexity:** O(N) max heights stored at once
