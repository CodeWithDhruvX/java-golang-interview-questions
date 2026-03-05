# Binary Search & Intervals (Product-Based Companies)

At top product-based companies (Google, Meta, Amazon), simple Binary Search is rarely asked. Instead, they heavily test **Binary Search on Answer Space** (where you binary search the *result* rather than the array index) and **Interval processing** (Merge, Insert, Overlaps).

## Question 1: Koko Eating Bananas (Binary Search on Answer Space)
**Problem Statement:** Koko loves to eat bananas. There are `n` piles of bananas, the `i-th` pile has `piles[i]` bananas. The guards have gone and will come back in `h` hours. Koko can decide her bananas-per-hour eating speed of `k`. Each hour, she chooses some pile of bananas and eats `k` bananas from that pile. If the pile has less than `k` bananas, she eats all of them instead and will not eat any more bananas during this hour. Return the minimum integer `k` such that she can eat all the bananas within `h` hours.

### Answer:
This is a classic "Binary Search on Answer Range" problem. The minimum eating speed `k` is 1, and the maximum is the maximum bananas in any single pile. We binary search between `[1, max(piles)]`. For each mid (potential speed `k`), we check if Koko can eat all bananas within `h` hours. 

**Code Implementation (Java):**
```java
public class KokoEatingBananas {
    public int minEatingSpeed(int[] piles, int h) {
        int left = 1;
        int right = 1;
        for (int pile : piles) {
            right = Math.max(right, pile);
        }

        while (left < right) {
            int mid = left + (right - left) / 2;
            if (canEatAll(piles, h, mid)) {
                right = mid; // Try to find a smaller speed
            } else {
                left = mid + 1; // Current speed is too slow
            }
        }
        return left;
    }

    private boolean canEatAll(int[] piles, int h, int k) {
        int hoursNeeded = 0;
        for (int pile : piles) {
            hoursNeeded += Math.ceil((double) pile / k);
        }
        return hoursNeeded <= h;
    }
}
```
**Time Complexity:** O(N log M) where N is the number of piles and M is the maximum bananas in a pile.
**Space Complexity:** O(1)

---

## Question 2: Median of Two Sorted Arrays
**Problem Statement:** Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return the median of the two sorted arrays. The overall run time complexity should be `O(log (m+n))`.

### Answer:
To achieve logarithmic time, we use Binary Search on the smaller array. We try to partition both arrays such that the number of elements in the left half of the combined array equals the right half, and every element on the left is less than or equal to every element on the right.

**Code Implementation (Java):**
```java
public class MedianTwoSortedArrays {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        if (nums1.length > nums2.length) {
            return findMedianSortedArrays(nums2, nums1); // Ensure nums1 is smaller
        }
        
        int x = nums1.length;
        int y = nums2.length;
        
        int low = 0;
        int high = x;
        
        while (low <= high) {
            int partitionX = (low + high) / 2;
            int partitionY = (x + y + 1) / 2 - partitionX;
            
            int maxLeftX = (partitionX == 0) ? Integer.MIN_VALUE : nums1[partitionX - 1];
            int minRightX = (partitionX == x) ? Integer.MAX_VALUE : nums1[partitionX];
            
            int maxLeftY = (partitionY == 0) ? Integer.MIN_VALUE : nums2[partitionY - 1];
            int minRightY = (partitionY == y) ? Integer.MAX_VALUE : nums2[partitionY];
            
            if (maxLeftX <= minRightY && maxLeftY <= minRightX) {
                // We have partitioned array at correct place
                if ((x + y) % 2 == 0) {
                    return ((double)Math.max(maxLeftX, maxLeftY) + Math.min(minRightX, minRightY)) / 2;
                } else {
                    return (double)Math.max(maxLeftX, maxLeftY);
                }
            } else if (maxLeftX > minRightY) { // We are too far on right side for partitionX. Go on left side.
                high = partitionX - 1;
            } else { // We are too far on left side for partitionX. Go on right side.
                low = partitionX + 1;
            }
        }
        throw new IllegalArgumentException();
    }
}
```
**Time Complexity:** O(log(min(m, n)))
**Space Complexity:** O(1)

---

## Question 3: Merge Intervals
**Problem Statement:** Given an array of `intervals` where `intervals[i] = [start_i, end_i]`, merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

### Answer:
First, we sort the intervals based on their start times. Then, we iterate through the sorted intervals. If the current interval's start time is less than or equal to the previous interval's end time, they overlap, so we merge them by updating the end time. Otherwise, they do not overlap, and we add the current interval to the result.

**Code Implementation (Java):**
```java
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class MergeIntervals {
    public int[][] merge(int[][] intervals) {
        if (intervals.length <= 1) return intervals;
        
        // Sort by ascending starting point
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        
        List<int[]> result = new ArrayList<>();
        int[] currentInterval = intervals[0];
        result.add(currentInterval);
        
        for (int[] interval : intervals) {
            int currentEnd = currentInterval[1];
            int nextStart = interval[0];
            int nextEnd = interval[1];
            
            if (currentEnd >= nextStart) { // Overlap, merge them
                currentInterval[1] = Math.max(currentEnd, nextEnd);
            } else { // No overlap, add to result
                currentInterval = interval;
                result.add(currentInterval);
            }
        }
        
        return result.toArray(new int[result.size()][]);
    }
}
```
**Time Complexity:** O(N log N) for sorting
**Space Complexity:** O(N) or O(log N) depending on sorting algorithm implementation. Output list logic takes O(N).

---

## Question 4: Insert Interval
**Problem Statement:** You are given an array of non-overlapping intervals `intervals` where `intervals[i] = [start_i, end_i]` represent the start and the end of the `i-th` interval and `intervals` is sorted in ascending order by `start_i`. You are also given an interval `newInterval = [start, end]` that represents the start and end of another interval. Insert `newInterval` into `intervals` such that `intervals` is still sorted in ascending order by `start_i` and `intervals` still does not have any overlapping intervals (merge overlapping intervals if necessary).

### Answer:
Iterate over the intervals. There are three cases for each interval compared to the `newInterval`:
1. It ends before `newInterval` starts -> Add to result.
2. It starts after `newInterval` ends -> Add `newInterval` to result (if not added yet), then add current interval.
3. It overlaps with `newInterval` -> Merge them by updating the start and end of `newInterval`, but don't add to result yet.

**Code Implementation (Java):**
```java
import java.util.ArrayList;
import java.util.List;

public class InsertInterval {
    public int[][] insert(int[][] intervals, int[] newInterval) {
        List<int[]> result = new ArrayList<>();
        int i = 0;
        int n = intervals.length;
        
        // Add all intervals ending before newInterval starts
        while (i < n && intervals[i][1] < newInterval[0]) {
            result.add(intervals[i]);
            i++;
        }
        
        // Merge all overlapping intervals into one newInterval
        while (i < n && intervals[i][0] <= newInterval[1]) {
            newInterval[0] = Math.min(newInterval[0], intervals[i][0]);
            newInterval[1] = Math.max(newInterval[1], intervals[i][1]);
            i++;
        }
        result.add(newInterval);
        
        // Add all remaining intervals
        while (i < n) {
            result.add(intervals[i]);
            i++;
        }
        
        return result.toArray(new int[result.size()][]);
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(N) for output list
