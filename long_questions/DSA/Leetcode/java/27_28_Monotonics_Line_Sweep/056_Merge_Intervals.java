import java.util.*;

public class MergeIntervals {
    
    // 56. Merge Intervals - Line Sweep Algorithm
    // Time: O(N log N), Space: O(N)
    public int[][] merge(int[][] intervals) {
        if (intervals.length <= 1) {
            return intervals;
        }
        
        // Sort intervals by start time
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        
        // Line sweep algorithm
        List<int[]> result = new ArrayList<>();
        int[] currentInterval = intervals[0];
        
        for (int i = 1; i < intervals.length; i++) {
            int[] nextInterval = intervals[i];
            
            // Check if intervals overlap
            if (nextInterval[0] <= currentInterval[1]) {
                // Merge intervals
                currentInterval[1] = Math.max(currentInterval[1], nextInterval[1]);
            } else {
                // Add current interval to result and start new one
                result.add(currentInterval);
                currentInterval = nextInterval;
            }
        }
        
        result.add(currentInterval);
        return result.toArray(new int[result.size()][]);
    }
    
    // Line sweep with event points
    public int[][] mergeLineSweep(int[][] intervals) {
        if (intervals.length <= 1) {
            return intervals;
        }
        
        // Create events
        class Event {
            int position;
            boolean isStart;
            
            Event(int position, boolean isStart) {
                this.position = position;
                this.isStart = isStart;
            }
        }
        
        List<Event> events = new ArrayList<>();
        for (int[] interval : intervals) {
            events.add(new Event(interval[0], true));   // Start event
            events.add(new Event(interval[1], false));  // End event
        }
        
        // Sort events by position, start events before end events at same position
        Collections.sort(events, (a, b) -> {
            if (a.position != b.position) {
                return Integer.compare(a.position, b.position);
            }
            return a.isStart ? -1 : 1; // Start events come first
        });
        
        List<int[]> result = new ArrayList<>();
        int activeCount = 0;
        int currentStart = -1;
        
        for (Event event : events) {
            if (event.isStart) {
                if (activeCount == 0) {
                    currentStart = event.position;
                }
                activeCount++;
            } else {
                activeCount--;
                if (activeCount == 0) {
                    result.add(new int[]{currentStart, event.position});
                }
            }
        }
        
        return result.toArray(new int[result.size()][]);
    }
    
    // Alternative line sweep with counting
    public int[][] mergeLineSweepCounting(int[][] intervals) {
        if (intervals.length <= 1) {
            return intervals;
        }
        
        // Find min and max positions
        int minPos = intervals[0][0], maxPos = intervals[0][1];
        for (int[] interval : intervals) {
            minPos = Math.min(minPos, interval[0]);
            maxPos = Math.max(maxPos, interval[1]);
        }
        
        // Create difference array
        int[] diff = new int[maxPos - minPos + 2];
        
        // Add intervals to difference array
        for (int[] interval : intervals) {
            diff[interval[0] - minPos]++;
            diff[interval[1] - minPos]--;
        }
        
        // Reconstruct intervals
        List<int[]> result = new ArrayList<>();
        boolean inInterval = false;
        int start = -1;
        
        for (int i = 0; i < diff.length; i++) {
            if (diff[i] > 0 && !inInterval) {
                start = i + minPos;
                inInterval = true;
            } else if (diff[i] <= 0 && inInterval) {
                result.add(new int[]{start, i + minPos});
                inInterval = false;
            }
        }
        
        return result.toArray(new int[result.size()][]);
    }
    
    // Line sweep with priority queue
    public int[][] mergeLineSweepPriorityQueue(int[][] intervals) {
        if (intervals.length <= 1) {
            return intervals;
        }
        
        // Sort intervals by start time
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        
        List<int[]> result = new ArrayList<>();
        
        // Simple priority queue using PriorityQueue
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> Integer.compare(a[1], b[1]));
        int current = 0;
        
        while (current < intervals.length || !pq.isEmpty()) {
            // Add all intervals that start at or before current position
            while (current < intervals.length && (pq.isEmpty() || intervals[current][0] <= pq.peek()[1])) {
                pq.offer(intervals[current]);
                current++;
            }
            
            // Find minimum end time
            int minEnd = pq.peek()[1];
            
            // Remove intervals that end at or before minEnd
            while (!pq.isEmpty() && pq.peek()[1] <= minEnd) {
                pq.poll();
            }
            
            // Add merged interval
            if (pq.isEmpty() && current < intervals.length) {
                // Start new interval
                if (current > 0) {
                    result.add(new int[]{intervals[current - 1][0], minEnd});
                }
            }
        }
        
        return result.toArray(new int[result.size()][]);
    }
    
    // Optimized line sweep with two pointers
    public int[][] mergeTwoPointers(int[][] intervals) {
        if (intervals.length <= 1) {
            return intervals;
        }
        
        // Sort intervals by start time
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        
        List<int[]> result = new ArrayList<>();
        int[] merged = intervals[0];
        
        for (int i = 1; i < intervals.length; i++) {
            int[] current = intervals[i];
            
            if (current[0] <= merged[1]) {
                // Overlap, merge
                merged[1] = Math.max(merged[1], current[1]);
            } else {
                // No overlap, add merged interval and start new
                result.add(merged);
                merged = current;
            }
        }
        
        result.add(merged);
        return result.toArray(new int[result.size()][]);
    }
    
    // Line sweep with interval tree simulation
    public int[][] mergeIntervalTree(int[][] intervals) {
        if (intervals.length <= 1) {
            return intervals;
        }
        
        // Sort by start time
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        
        List<int[]> result = new ArrayList<>();
        int[] current = intervals[0];
        
        for (int i = 1; i < intervals.length; i++) {
            int[] next = intervals[i];
            
            // Check overlap
            if (next[0] <= current[1]) {
                // Merge
                current[1] = Math.max(current[1], next[1]);
            } else {
                // Add current and start new
                result.add(current);
                current = next;
            }
        }
        
        result.add(current);
        return result.toArray(new int[result.size()][]);
    }
    
    // Line sweep with detailed events
    public class MergeResult {
        int[][] intervals;
        List<String> explanation;
        
        MergeResult(int[][] intervals, List<String> explanation) {
            this.intervals = intervals;
            this.explanation = explanation;
        }
    }
    
    public MergeResult mergeLineSweepDetailed(int[][] intervals) {
        List<String> explanation = new ArrayList<>();
        
        if (intervals.length <= 1) {
            explanation.add("0 or 1 intervals, returning as is");
            return new MergeResult(intervals, explanation);
        }
        
        explanation.add(String.format("Processing %d intervals", intervals.length));
        
        // Create events
        class Event {
            int position;
            boolean isStart;
            int interval;
            
            Event(int position, boolean isStart, int interval) {
                this.position = position;
                this.isStart = isStart;
                this.interval = interval;
            }
        }
        
        List<Event> events = new ArrayList<>();
        for (int i = 0; i < intervals.length; i++) {
            events.add(new Event(intervals[i][0], true, i));
            events.add(new Event(intervals[i][1], false, i));
            explanation.add(String.format("Created events for interval %d: start=%d, end=%d", 
                i, intervals[i][0], intervals[i][1]));
        }
        
        // Sort events
        Collections.sort(events, (a, b) -> {
            if (a.position != b.position) {
                return Integer.compare(a.position, b.position);
            }
            return a.isStart ? -1 : 1;
        });
        
        explanation.add("Sorted events by position (start before end at same position)");
        
        List<int[]> result = new ArrayList<>();
        int activeCount = 0;
        int currentStart = -1;
        
        for (int i = 0; i < events.size(); i++) {
            Event event = events.get(i);
            if (event.isStart) {
                if (activeCount == 0) {
                    currentStart = event.position;
                    explanation.add(String.format("Event %d: Starting new interval at %d", 
                        i, currentStart));
                }
                activeCount++;
            } else {
                activeCount--;
                if (activeCount == 0) {
                    result.add(new int[]{currentStart, event.position});
                    explanation.add(String.format("Event %d: Ending interval at %d, result: %s", 
                        i, event.position, Arrays.toString(new int[]{currentStart, event.position})));
                }
            }
        }
        
        return new MergeResult(result.toArray(new int[result.size()][]), explanation);
    }
    
    public static void main(String[] args) {
        MergeIntervals mi = new MergeIntervals();
        
        // Test cases
        System.out.println("=== Testing Merge Intervals - Line Sweep ===");
        
        int[][][][] testCases = {
            {{1, 3}, {2, 6}, {8, 10}, {15, 18}},
            {{1, 4}, {4, 5}},
            {{1, 10}, {2, 3}, {4, 5}, {6, 7}},
            {{1, 3}, {5, 7}, {9, 11}},
            {{1, 5}},
            {},
            {{1, 2}, {3, 4}, {2, 3}},
            {{1, 4}, {0, 2}, {3, 5}},
            {{1, 100}, {2, 3}, {4, 5}, {6, 7}, {8, 9}, {10, 11}},
            {{1, 3}, {2, 4}, {3, 5}, {4, 6}}
        };
        
        String[] descriptions = {
            "Standard case",
            "Touching intervals",
            "Nested intervals",
            "Non-overlapping",
            "Single interval",
            "Empty array",
            "Sequential with overlap",
            "Mixed order",
            "Large interval with many small",
            "Chain overlapping"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("  Input: %s\n", Arrays.deepToString(testCases[i]));
            
            int[][] result1 = mi.merge(testCases[i]);
            int[][] result2 = mi.mergeLineSweep(testCases[i]);
            int[][] result3 = mi.mergeLineSweepCounting(testCases[i]);
            int[][] result4 = mi.mergeTwoPointers(testCases[i]);
            
            System.out.printf("  Standard: %s\n", Arrays.deepToString(result1));
            System.out.printf("  Line Sweep: %s\n", Arrays.deepToString(result2));
            System.out.printf("  Counting: %s\n", Arrays.deepToString(result3));
            System.out.printf("  Two Pointers: %s\n\n", Arrays.deepToString(result4));
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Line Sweep Explanation ===");
        int[][] testIntervals = {{1, 3}, {2, 6}, {8, 10}, {15, 18}};
        MergeResult result = mi.mergeLineSweepDetailed(testIntervals);
        
        System.out.printf("Result: %s\n", Arrays.deepToString(result.intervals));
        for (String step : result.explanation) {
            System.out.printf("  %s\n", step);
        }
        
        // Performance test
        System.out.println("\n=== Performance Test ===");
        
        int[][] largeIntervals = new int[10000][];
        for (int i = 0; i < 10000; i++) {
            largeIntervals[i] = new int[]{i, i + 10};
        }
        
        System.out.printf("Large test with %d intervals\n", largeIntervals.length);
        
        int[][] largeResult = mi.mergeLineSweep(largeIntervals);
        System.out.printf("Line sweep result length: %d\n", largeResult.length);
        
        largeResult = mi.mergeTwoPointers(largeIntervals);
        System.out.printf("Two pointers result length: %d\n", largeResult.length);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        
        // Overlapping at same point
        int[][] samePoint = {{1, 3}, {3, 5}, {3, 7}};
        System.out.printf("Same point overlap: %s\n", Arrays.deepToString(mi.mergeLineSweep(samePoint)));
        
        // Single point intervals
        int[][] singlePoints = {{1, 1}, {2, 2}, {3, 3}};
        System.out.printf("Single points: %s\n", Arrays.deepToString(mi.mergeLineSweep(singlePoints)));
        
        // Very large numbers
        int[][] largeNumbers = {{1000000, 2000000}, {1500000, 2500000}, {3000000, 4000000}};
        System.out.printf("Large numbers: %s\n", Arrays.deepToString(mi.mergeLineSweep(largeNumbers)));
        
        // Negative numbers
        int[][] negativeNumbers = {{-5, -3}, {-4, -2}, {-1, 1}};
        System.out.printf("Negative numbers: %s\n", Arrays.deepToString(mi.mergeLineSweep(negativeNumbers)));
        
        // Test priority queue approach
        System.out.println("\n=== Priority Queue Test ===");
        int[][] pqTest = {{1, 4}, {2, 5}, {3, 6}, {7, 8}};
        System.out.printf("Priority queue: %s\n", Arrays.deepToString(mi.mergeLineSweepPriorityQueue(pqTest)));
    }
}
