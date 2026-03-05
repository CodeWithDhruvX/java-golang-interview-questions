# Hard Patterns & String Algorithms (Product-Based Companies)

This final section covers the absolute hardest, most niche algorithms tested at top-tier FAANG companies (like Google, Citadel, and Meta). These are often asked in L5 (Senior) or highly competitive backend roles.

## Question 1: Course Schedule II (Topological Sort / Kahn's Algorithm)
**Problem Statement:** There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [a, b]` indicates that you must take course `b` first if you want to take course `a`. Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

### Answer:
This is the classic application of **Topological Sort** on a Directed Acyclic Graph (DAG). We use Kahn's Algorithm (BFS with In-Degree counting). We find all nodes with an in-degree of 0 (no prerequisites) and put them in a Queue. As we process them, we remove their outgoing edges (decrement the in-degree of their neighbors). If a neighbor's in-degree hits 0, we add it to the Queue.

**Code Implementation (Java):**
```java
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class CourseScheduleII {
    public int[] findOrder(int numCourses, int[][] prerequisites) {
        int[] inDegree = new int[numCourses];
        List<List<Integer>> adj = new ArrayList<>();
        
        for (int i = 0; i < numCourses; i++) {
            adj.add(new ArrayList<>());
        }
        
        // Build graph and calculate in-degrees
        for (int[] pre : prerequisites) {
            int course = pre[0];
            int required = pre[1];
            adj.get(required).add(course);
            inDegree[course]++;
        }
        
        Queue<Integer> queue = new LinkedList<>();
        // Add all courses with no prerequisites
        for (int i = 0; i < numCourses; i++) {
            if (inDegree[i] == 0) {
                queue.offer(i);
            }
        }
        
        int[] order = new int[numCourses];
        int index = 0;
        
        while (!queue.isEmpty()) {
            int current = queue.poll();
            order[index++] = current;
            
            for (int neighbor : adj.get(current)) {
                inDegree[neighbor]--;
                if (inDegree[neighbor] == 0) {
                    queue.offer(neighbor);
                }
            }
        }
        
        // If index == numCourses, we successfully scheduled all courses (no cycles)
        return index == numCourses ? order : new int[0];
    }
}
```
**Time Complexity:** O(V + E) where V is `numCourses` and E is the length of `prerequisites`.
**Space Complexity:** O(V + E) to store the adjacency list.

---

## Question 2: First Missing Positive (Cyclic Sort)
**Problem Statement:** Given an unsorted integer array `nums`. Return the smallest positive integer that is not present in `nums`. You must implement an algorithm that runs in `O(N)` time and uses `O(1)` auxiliary space.

### Answer:
Since we need `O(1)` space, we cannot use a HashSet. The key insight is that the answer *must* be between `1` and `N+1` (where `N` is the length of the array). We can use the array itself as a hash table via **Cyclic Sort**. We iterate through the array and try to place every number `x` (where `1 <= x <= N`) at its correct index `x - 1`. Finally, we scan the array again to find the first index `i` that does not contain `i + 1`.

**Code Implementation (Java):**
```java
public class FirstMissingPositive {
    public int firstMissingPositive(int[] nums) {
        int n = nums.length;
        int i = 0;
        
        // Cyclic Sort: Put each number in its right place
        // E.g., put 1 at index 0, 2 at index 1...
        while (i < n) {
            // If the number is within the valid range (1 to n) 
            // AND it is not already sitting at its correct index
            if (nums[i] > 0 && nums[i] <= n && nums[nums[i] - 1] != nums[i]) {
                swap(nums, i, nums[i] - 1);
            } else {
                i++;
            }
        }
        
        // Find the first index that doesn't have the correct number
        for (i = 0; i < n; i++) {
            if (nums[i] != i + 1) {
                return i + 1;
            }
        }
        
        // If all numbers 1..N are present, the missing positive is N + 1
        return n + 1;
    }
    
    private void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
}
```
**Time Complexity:** O(N). Even though there is a `while` loop inside a `while` loop (via the swap moving things around without incrementing `i`), each element is swapped into its correct place at most once.
**Space Complexity:** O(1)

---

## Question 3: Find the Index of the First Occurrence in a String (KMP Algorithm)
**Problem Statement:** Given two strings `needle` and `haystack`, return the index of the first occurrence of `needle` in `haystack`, or `-1` if `needle` is not part of `haystack`. Expected O(N+M) time complexity.

### Answer:
While Java's `indexOf` or a brute-force approach works for easy constraints, FAANG demands the **Knuth-Morris-Pratt (KMP)** algorithm for guaranteed linear time string matching. We pre-process the `needle` to create an `LPS` (Longest Prefix Suffix) array. This array tells us how far we can "jump back" in the `needle` when a mismatch occurs, without re-scanning characters in the `haystack`.

**Code Implementation (Java):**
```java
public class KMPStringMatching {
    public int strStr(String haystack, String needle) {
        if (needle.isEmpty()) return 0;
        
        int n = haystack.length();
        int m = needle.length();
        int[] lps = computeLPS(needle); // Pre-process needle
        
        int i = 0; // Index for haystack
        int j = 0; // Index for needle
        
        while (i < n) {
            if (haystack.charAt(i) == needle.charAt(j)) {
                i++;
                j++;
                if (j == m) {
                    return i - m; // Match found
                }
            } else {
                if (j > 0) {
                    // Mismatch happened, but we saw matching prefix before
                    // Jump back using LPS array instead of resetting j to 0
                    j = lps[j - 1];
                } else {
                    // Mismatch at the very first character of needle
                    i++;
                }
            }
        }
        return -1;
    }
    
    private int[] computeLPS(String pattern) {
        int m = pattern.length();
        int[] lps = new int[m];
        int len = 0; // Length of previous longest prefix suffix
        int i = 1;
        
        while (i < m) {
            if (pattern.charAt(i) == pattern.charAt(len)) {
                len++;
                lps[i] = len;
                i++;
            } else {
                if (len > 0) {
                    len = lps[len - 1];
                } else {
                    lps[i] = 0;
                    i++;
                }
            }
        }
        return lps;
    }
}
```
**Time Complexity:** O(N + M) where N is length of `haystack` and M is length of `needle`.
**Space Complexity:** O(M) for the LPS array.

---

## Question 4: Range Sum Query - Mutable (Segment Tree)
**Problem Statement:** Given an integer array `nums`, handle multiple queries of the following types:
1. `update(index, val)`: Updates the value of `nums[index]` to be `val`.
2. `sumRange(left, right)`: Returns the sum of the elements of `nums` between indices `left` and `right` inclusive.

### Answer:
If there were no updates, Prefix Sums `O(1)` would work. If there were no range queries, standard arrays `O(1)` update would work. To do both efficiently, we need a **Segment Tree** (or Fenwick Tree). A Segment Tree is a binary tree where the root represents the sum of the whole array, and children represent sums of halves. 

**Code Implementation (Java):**
```java
public class NumArray {
    int[] tree;
    int n;

    public NumArray(int[] nums) {
        if (nums.length > 0) {
            n = nums.length;
            tree = new int[n * 2]; // 2N size is sufficient for iterative segment tree
            buildTree(nums);
        }
    }
    
    private void buildTree(int[] nums) {
        // Insert leaf nodes in the second half of the tree array
        for (int i = n, j = 0;  i < 2 * n; i++,  j++) {
            tree[i] = nums[j];
        }
        // Build the tree by calculating parents from leaves up to root
        for (int i = n - 1; i > 0; --i) {
            tree[i] = tree[i * 2] + tree[i * 2 + 1];
        }
    }
    
    public void update(int pos, int val) {
        pos += n; // Get leaf node index
        tree[pos] = val; // Update leaf
        
        // Bubble up the difference to the root
        while (pos > 0) {
            int left = pos;
            int right = pos;
            if (pos % 2 == 0) {
                right = pos + 1; // It's a left child, its sibling is pos+1
            } else {
                left = pos - 1; // It's a right child, its sibling is pos-1
            }
            // Update parent
            tree[pos / 2] = tree[left] + tree[right];
            pos /= 2;
        }
    }
    
    public int sumRange(int l, int r) {
        // Offset to leaf nodes
        l += n;
        r += n;
        int sum = 0;
        
        while (l <= r) {
            if ((l % 2) == 1) { // If l is a right child
                sum += tree[l];
                l++; // Move to next boundary
            }
            if ((r % 2) == 0) { // If r is a left child
                sum += tree[r];
                r--; // Move to previous boundary
            }
            l /= 2;
            r /= 2;
        }
        return sum;
    }
}
```
**Time Complexity:** `buildTree`: O(N). `update` and `sumRange`: O(log N).
**Space Complexity:** O(N) for the segment tree array.
