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
}
