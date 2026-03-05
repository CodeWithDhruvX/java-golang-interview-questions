# Sorting and Searching (Service-Based Companies)

Searching and basic Sorting algorithms are highly requested in initial technical rounds to check if candidates understand algorithm complexity and basic mechanics.

## Question 1: Binary Search
**Problem Statement:** Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise, return `-1`.

### Answer:
Since the array is sorted, we can use binary search by continuously dividing the search interval in half.

**Code Implementation (Java):**
```java
public class BinarySearch {
    public int search(int[] nums, int target) {
        int left = 0;
        int right = nums.length - 1;
        
        while (left <= right) {
            int mid = left + (right - left) / 2; // Prevent integer overflow
            
            if (nums[mid] == target) {
                return mid;
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return -1;
    }
}
```
**Time Complexity:** O(log N)
**Space Complexity:** O(1)

---

## Question 2: Bubble Sort
**Problem Statement:** Implement the Bubble Sort algorithm to sort an array of integers in ascending order.

### Answer:
Bubble sort repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order. An optimized bubble sort adds a boolean flag to stop early if no swaps are made in a pass.

**Code Implementation (Java):**
```java
public class BubbleSort {
    public void bubbleSort(int[] arr) {
        int n = arr.length;
        boolean swapped;
        
        for (int i = 0; i < n - 1; i++) {
            swapped = false;
            // The last i elements are already sorted
            for (int j = 0; j < n - i - 1; j++) {
                if (arr[j] > arr[j + 1]) {
                    int temp = arr[j];
                    arr[j] = arr[j + 1];
                    arr[j + 1] = temp;
                    swapped = true;
                }
            }
            // If no elements were swapped, it means the array is sorted
            if (!swapped) break;
        }
    }
}
```
**Time Complexity:** O(N^2) worst/average case, O(N) best case
**Space Complexity:** O(1)

---

## Question 3: Search Insert Position
**Problem Statement:** Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

### Answer:
This is a pure variation of Binary Search. Instead of just returning `-1` if the element isn't found, you return the `left` pointer, which naturally points to the correct insertion spot when the loop terminates.

**Code Implementation (Java):**
```java
public class SearchInsertPosition {
    public int searchInsert(int[] nums, int target) {
        int left = 0;
        int right = nums.length - 1;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (nums[mid] == target) {
                return mid;
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        return left;
    }
}
```
**Time Complexity:** O(log N)
**Space Complexity:** O(1)

---

## Question 4: Merge Sorted Array
**Problem Statement:** You are given two integer arrays `nums1` and `nums2`, sorted in non-decreasing order. Merge `nums2` into `nums1` as one sorted array. `nums1` has a size large enough to hold additional elements from `nums2`.

### Answer:
Instead of creating a new array, we can start filling `nums1` from the back. We keep three pointers: one for the last active element of `nums1`, one for the last element of `nums2`, and one for the end of the total length in `nums1`.

**Code Implementation (Java):**
```java
public class MergeSortedArray {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int p1 = m - 1;
        int p2 = n - 1;
        int p = m + n - 1;
        
        // While there are elements in nums2
        while (p2 >= 0) {
            if (p1 >= 0 && nums1[p1] > nums2[p2]) {
                nums1[p] = nums1[p1];
                p1--;
            } else {
                nums1[p] = nums2[p2];
                p2--;
            }
            p--;
        }
    }
}
```
**Time Complexity:** O(M + N)
**Space Complexity:** O(1)
