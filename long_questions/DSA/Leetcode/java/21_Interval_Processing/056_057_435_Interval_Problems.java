import java.util.*;

public class IntervalProcessing {
    
    // 56. Merge Intervals
    // Time: O(N log N), Space: O(N)
    public static int[][] merge(int[][] intervals) {
        if (intervals == null || intervals.length <= 1) {
            return intervals;
        }
        
        // Sort intervals by start time
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        
        List<int[]> merged = new ArrayList<>();
        int[] current = intervals[0];
        
        for (int i = 1; i < intervals.length; i++) {
            int[] interval = intervals[i];
            
            if (interval[0] <= current[1]) {
                // Overlapping intervals, merge them
                current[1] = Math.max(current[1], interval[1]);
            } else {
                // Non-overlapping, add current and start new
                merged.add(current);
                current = interval;
            }
        }
        
        merged.add(current);
        
        return merged.toArray(new int[merged.size()][]);
    }

    // 57. Insert Interval
    // Time: O(N), Space: O(N)
    public static int[][] insert(int[][] intervals, int[] newInterval) {
        if (intervals == null || intervals.length == 0) {
            return new int[][]{newInterval};
        }
        
        List<int[]> result = new ArrayList<>();
        int i = 0;
        int n = intervals.length;
        
        // Add all intervals that end before new interval starts
        while (i < n && intervals[i][1] < newInterval[0]) {
            result.add(intervals[i]);
            i++;
        }
        
        // Merge overlapping intervals
        while (i < n && intervals[i][0] <= newInterval[1]) {
            newInterval[0] = Math.min(newInterval[0], intervals[i][0]);
            newInterval[1] = Math.max(newInterval[1], intervals[i][1]);
            i++;
        }
        
        result.add(newInterval);
        
        // Add remaining intervals
        while (i < n) {
            result.add(intervals[i]);
            i++;
        }
        
        return result.toArray(new int[result.size()][]);
    }

    // 435. Non-overlapping Intervals
    // Time: O(N log N), Space: O(1)
    public static int eraseOverlapIntervals(int[][] intervals) {
        if (intervals == null || intervals.length <= 1) {
            return 0;
        }
        
        // Sort intervals by end time
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[1], b[1]));
        
        int count = 0;
        int end = intervals[0][1];
        
        for (int i = 1; i < intervals.length; i++) {
            if (intervals[i][0] < end) {
                // Overlapping, remove current interval
                count++;
            } else {
                // Non-overlapping, update end
                end = intervals[i][1];
            }
        }
        
        return count;
    }

    public static void main(String[] args) {
        // Test cases for merge
        int[][][] testCases1 = {
            {{1, 3}, {2, 6}, {8, 10}, {15, 18}},
            {{1, 4}, {4, 5}},
            {{1, 3}, {2, 4}, {3, 5}},
            {{1, 2}, {3, 4}, {5, 6}},
            {{1, 10}, {2, 3}, {4, 5}, {6, 7}},
            {{1, 3}},
            {},
            {{1, 4}, {0, 2}, {3, 5}},
            {{1, 3}, {5, 7}, {9, 11}},
            {{1, 4}, {2, 3}, {3, 6}}
        };
        
        // Test cases for insert
        Object[][] testCases2 = {
            {new int[][]{{1, 3}, {6, 9}}, new int[]{2, 5}},
            {new int[][]{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, new int[]{4, 8}},
            {new int[][]{}, new int[]{5, 7}},
            {new int[][]{{1, 5}}, new int[]{2, 3}},
            {new int[][]{{1, 3}, {6, 9}}, new int[]{10, 12}},
            {new int[][]{{1, 5}}, new int[]{6, 8}},
            {new int[][]{{1, 3}, {4, 6}}, new int[]{2, 5}},
            {new int[][]{{1, 2}, {3, 5}, {6, 7}}, new int[]{4, 6}},
            {new int[][]{{1, 5}}, new int[]{0, 0}},
            {new int[][]{{1, 3}, {6, 9}}, new int[]{0, 10}}
        };
        
        // Test cases for eraseOverlapIntervals
        int[][][] testCases3 = {
            {{1, 2}, {2, 3}, {3, 4}, {1, 3}},
            {{1, 2}, {2, 3}, {3, 4}, {4, 5}},
            {{1, 2}, {1, 2}, {1, 2}},
            {{1, 100}, {11, 22}, {1, 11}, {2, 12}},
            {{1, 2}, {3, 4}, {5, 6}},
            {{1, 3}, {2, 4}, {3, 5}},
            {{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}},
            {{1, 10}, {2, 3}, {4, 5}, {6, 7}},
            {{1, 2}, {1, 3}, {2, 3}, {3, 4}},
            {{1, 5}, {6, 10}, {11, 15}}
        };
        
        System.out.println("Merge Intervals:");
        for (int i = 0; i < testCases1.length; i++) {
            int[][] intervals = testCases1[i];
            int[][] result = merge(intervals);
            System.out.printf("Test Case %d: %s -> %s\n", 
                i + 1, Arrays.deepToString(intervals), Arrays.deepToString(result));
        }
        
        System.out.println("\nInsert Interval:");
        for (int i = 0; i < testCases2.length; i++) {
            int[][] intervals = (int[][]) testCases2[i][0];
            int[] newInterval = (int[]) testCases2[i][1];
            int[][] result = insert(intervals, newInterval);
            System.out.printf("Test Case %d: %s, new=%s -> %s\n", 
                i + 1, Arrays.deepToString(intervals), Arrays.toString(newInterval), 
                Arrays.deepToString(result));
        }
        
        System.out.println("\nErase Overlap Intervals:");
        for (int i = 0; i < testCases3.length; i++) {
            int[][] intervals = testCases3[i];
            int result = eraseOverlapIntervals(intervals);
            System.out.printf("Test Case %d: %s -> %d intervals to remove\n", 
                i + 1, Arrays.deepToString(intervals), result);
        }
    }
}
