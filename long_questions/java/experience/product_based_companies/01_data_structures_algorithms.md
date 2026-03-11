# 🧮 01 — Data Structures & Algorithms in Java
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

> Companies like **Amazon, Flipkart, Paytm, Swiggy, Zomato, Razorpay, PhonePe, Meesho** focus heavily on DSA implemented in Java idioms.

---

## 🔑 Must-Know Topics
- Arrays, ArrayDeque, PriorityQueue
- Linked lists, stacks, queues using Java Collections
- Binary trees and BST operations
- Graph algorithms (BFS, DFS)
- HashMap/HashSet for algorithmic patterns
- Sorting and searching
- Common LeetCode-pattern problems in Java

---

## ❓ Most Asked Questions

### Q1. Implement a Stack using Java

```java
// Option 1: Use Deque (recommended — ArrayDeque is faster than Stack class)
Deque<Integer> stack = new ArrayDeque<>();
stack.push(1);   // adds to front
stack.push(2);
stack.push(3);
int top = stack.peek();   // 3 — look without removing
int val = stack.pop();    // 3 — remove from front
boolean empty = stack.isEmpty();

// Option 2: Generic Stack from scratch
public class Stack<T> {
    private final Deque<T> deque = new ArrayDeque<>();

    public void push(T item) { deque.push(item); }
    public T pop() {
        if (isEmpty()) throw new EmptyStackException();
        return deque.pop();
    }
    public T peek() { return deque.peek(); }
    public boolean isEmpty() { return deque.isEmpty(); }
    public int size() { return deque.size(); }
}

// AVOID: java.util.Stack — synchronized, legacy, slow
// USE: ArrayDeque for stack/queue in all interview solutions
```

---

### 🎯 How to Explain in Interview

"For implementing a stack in Java, I prefer using ArrayDeque over the legacy Stack class. ArrayDeque is faster because it's not synchronized, and it provides all the stack operations through the Deque interface. The key operations are push() to add to the top, pop() to remove from the top, and peek() to look at the top element without removing it. If I need to implement a custom stack, I can wrap ArrayDeque and add error handling. The important thing to remember is that Stack is synchronized and slower, while ArrayDeque is the modern, preferred approach for both stacks and queues in interview solutions."

---

### Q2. Implement a Queue using Java

```java
// Standard queue — LinkedList or ArrayDeque
Queue<String> queue = new LinkedList<>();
queue.offer("task1");   // enqueue (prefer offer over add — no exception)
queue.offer("task2");
String head = queue.peek();    // "task1" — look without removing
String next = queue.poll();    // "task1" — remove from front (null if empty)

// Priority Queue (Min-Heap by default)
PriorityQueue<Integer> minHeap = new PriorityQueue<>();
minHeap.offer(5); minHeap.offer(1); minHeap.offer(3);
System.out.println(minHeap.poll()); // 1 — always extracts minimum

// Max-Heap
PriorityQueue<Integer> maxHeap = new PriorityQueue<>(Comparator.reverseOrder());
maxHeap.offer(5); maxHeap.offer(1); maxHeap.offer(3);
System.out.println(maxHeap.poll()); // 5

// Custom object heap — sort Tasks by priority
PriorityQueue<Task> taskQueue = new PriorityQueue<>(
    Comparator.comparingInt(Task::getPriority).reversed()); // high priority first
```

---

### 🎯 How to Explain in Interview

"For queues in Java, I use LinkedList or ArrayDeque for standard FIFO queues with offer() to enqueue and poll() to dequeue. I prefer offer() over add() because it doesn't throw exceptions when the queue is full. For priority queues, PriorityQueue implements a heap - by default it's a min-heap that always extracts the smallest element. I can create a max-heap by providing a reverse comparator. The beauty is that I can also use custom objects with PriorityQueue by providing a Comparator that defines the priority ordering. This makes PriorityQueue perfect for problems like finding the k-th largest element or implementing Dijkstra's algorithm."

---

### Q3. Reverse a Linked List

```java
public class ListNode {
    int val;
    ListNode next;
    ListNode(int val) { this.val = val; }
}

// Iterative — O(n) time, O(1) space
public ListNode reverseList(ListNode head) {
    ListNode prev = null, curr = head;
    while (curr != null) {
        ListNode next = curr.next;
        curr.next = prev;
        prev = curr;
        curr = next;
    }
    return prev;
}

// Recursive — O(n) time, O(n) stack space
public ListNode reverseListRecursive(ListNode head) {
    if (head == null || head.next == null) return head;
    ListNode newHead = reverseListRecursive(head.next);
    head.next.next = head;
    head.next = null;
    return newHead;
}
```

---

### 🎯 How to Explain in Interview

"Reversing a linked list is a classic interview problem. The iterative approach uses three pointers: prev, curr, and next. I walk through the list, reversing each link as I go by setting curr.next to prev. This runs in O(n) time with O(1) space. The recursive approach is more elegant but uses O(n) stack space due to the recursion depth. In interviews, I usually show the iterative solution first because it's more space-efficient, then mention the recursive approach as an alternative. The key insight is that I need to carefully handle the pointer updates to avoid losing the rest of the list while reversing."

---

### Q4. Two Sum — HashMap pattern

```java
// O(n) with HashMap
public int[] twoSum(int[] nums, int target) {
    Map<Integer, Integer> seen = new HashMap<>();  // value → index
    for (int i = 0; i < nums.length; i++) {
        int complement = target - nums[i];
        if (seen.containsKey(complement)) {
            return new int[]{seen.get(complement), i};
        }
        seen.put(nums[i], i);
    }
    return new int[]{};
}

// Variant: Two Sum in sorted array — two pointers
public int[] twoSumSorted(int[] nums, int target) {
    int left = 0, right = nums.length - 1;
    while (left < right) {
        int sum = nums[left] + nums[right];
        if (sum == target) return new int[]{left + 1, right + 1}; // 1-indexed
        else if (sum < target) left++;
        else right--;
    }
    return new int[]{};
}
```

---

### 🎯 How to Explain in Interview

"The Two Sum problem demonstrates the power of hash maps for O(n) solutions. For the unsorted array version, I use a HashMap to store numbers I've seen and their indices. As I iterate through the array, I check if the complement (target - current number) is already in the map. If it is, I've found the pair. If not, I add the current number to the map. This gives me O(n) time with O(n) space. For the sorted array version, I can use the two-pointer technique - start with left at the beginning and right at the end, then move them inward based on whether the sum is too small or too large. This gives O(n) time with O(1) space. The key is choosing the right approach based on whether the array is sorted."

---

### Q5. Binary Search

```java
// Standard binary search — O(log n)
public int binarySearch(int[] nums, int target) {
    int left = 0, right = nums.length - 1;
    while (left <= right) {
        int mid = left + (right - left) / 2;  // avoids int overflow
        if (nums[mid] == target)      return mid;
        else if (nums[mid] < target)  left = mid + 1;
        else                          right = mid - 1;
    }
    return -1;
}

// Find first occurrence (leftmost)
public int findFirst(int[] nums, int target) {
    int left = 0, right = nums.length - 1, result = -1;
    while (left <= right) {
        int mid = left + (right - left) / 2;
        if (nums[mid] == target) { result = mid; right = mid - 1; } // keep going left
        else if (nums[mid] < target) left = mid + 1;
        else right = mid - 1;
    }
    return result;
}

// Java built-in
int idx = Arrays.binarySearch(nums, target);  // must be sorted; negative = not found
```

---

### 🎯 How to Explain in Interview

"Binary search is fundamental for searching in sorted arrays. The key is maintaining the search boundaries with left and right pointers, and calculating mid as left + (right - left) / 2 to avoid integer overflow. I compare the middle element with the target and adjust the boundaries accordingly. For finding the first occurrence of a duplicate, I continue searching left even after finding a match. Java provides Arrays.binarySearch() which returns the index if found, or a negative insertion point if not found. Binary search runs in O(log n) time, making it much faster than linear search for large sorted datasets. It's the foundation for many advanced algorithms and a must-know for interviews."

---

### Q6. BFS and DFS on a Graph

```java
// Graph represented as adjacency list
Map<Integer, List<Integer>> graph = new HashMap<>();

// BFS — level-order traversal, shortest path in unweighted graph
public List<Integer> bfs(int start) {
    List<Integer> visited = new ArrayList<>();
    Set<Integer> seen = new HashSet<>();
    Queue<Integer> queue = new LinkedList<>();
    queue.offer(start);
    seen.add(start);
    while (!queue.isEmpty()) {
        int node = queue.poll();
        visited.add(node);
        for (int neighbor : graph.getOrDefault(node, List.of())) {
            if (!seen.contains(neighbor)) {
                seen.add(neighbor);
                queue.offer(neighbor);
            }
        }
    }
    return visited;
}

// DFS — recursive
Set<Integer> visitedDFS = new HashSet<>();

public void dfs(int node) {
    visitedDFS.add(node);
    System.out.println(node);
    for (int neighbor : graph.getOrDefault(node, List.of())) {
        if (!visitedDFS.contains(neighbor)) {
            dfs(neighbor);
        }
    }
}
```

---

### 🎯 How to Explain in Interview

"BFS and DFS are two fundamental graph traversal algorithms. BFS uses a queue and explores level by level, making it perfect for finding the shortest path in unweighted graphs. I keep track of visited nodes to avoid cycles and process each node's neighbors before moving to the next level. DFS uses recursion or a stack and explores as deep as possible along each branch before backtracking. DFS is great for problems like detecting cycles or exploring all possible paths. In Java, I represent graphs using adjacency lists with HashMap, and use HashSet for tracking visited nodes. The choice between BFS and DFS depends on the problem - BFS for shortest paths, DFS for exhaustive exploration. Both run in O(V+E) time where V is vertices and E is edges."

---

### Q7. Sliding Window — Maximum sum subarray of size K

```java
public int maxSumSubarray(int[] nums, int k) {
    if (nums.length < k) return -1;
    int windowSum = 0;
    for (int i = 0; i < k; i++) windowSum += nums[i];
    int maxSum = windowSum;
    for (int i = k; i < nums.length; i++) {
        windowSum += nums[i] - nums[i - k];
        maxSum = Math.max(maxSum, windowSum);
    }
    return maxSum;
}

// Longest substring without repeating characters
public int lengthOfLongestSubstring(String s) {
    Map<Character, Integer> charIndex = new HashMap<>();
    int maxLen = 0, left = 0;
    for (int right = 0; right < s.length(); right++) {
        char c = s.charAt(right);
        if (charIndex.containsKey(c) && charIndex.get(c) >= left) {
            left = charIndex.get(c) + 1;
        }
        charIndex.put(c, right);
        maxLen = Math.max(maxLen, right - left + 1);
    }
    return maxLen;
}
```

---

### 🎯 How to Explain in Interview

"Sliding window is a powerful technique for problems involving subarrays or substrings. For the maximum sum subarray of size K, I first compute the sum of the first K elements, then slide the window by subtracting the element leaving the window and adding the new element entering the window. This gives me O(n) time instead of O(n*K) for the naive approach. For longest substring without repeating characters, I maintain a window with left and right pointers, and a HashMap to track the last seen position of each character. When I encounter a duplicate within the current window, I move the left pointer past the previous occurrence. Sliding window reduces many O(n²) problems to O(n) by avoiding reprocessing of elements that are already in the current window."

---

### Q8. Merge Intervals

```java
public int[][] merge(int[][] intervals) {
    if (intervals.length <= 1) return intervals;

    Arrays.sort(intervals, (a, b) -> a[0] - b[0]);  // sort by start

    List<int[]> merged = new ArrayList<>();
    int[] current = intervals[0];

    for (int i = 1; i < intervals.length; i++) {
        if (intervals[i][0] <= current[1]) {
            current[1] = Math.max(current[1], intervals[i][1]);  // merge
        } else {
            merged.add(current);
            current = intervals[i];
        }
    }
    merged.add(current);
    return merged.toArray(new int[0][]);
}
// Input:  [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
```

---

### 🎯 How to Explain in Interview

"Merge intervals is a classic sorting and greedy algorithm problem. I first sort the intervals by their start time, then iterate through them merging overlapping intervals. If the current interval's start is less than or equal to the previous interval's end, they overlap and I merge them by extending the end to the maximum of both ends. If they don't overlap, I add the previous interval to the result and start a new current interval. The key insight is that after sorting, I only need to check each interval against the immediately preceding one. This runs in O(n log n) time due to the sort, and O(n) space for the result. This pattern applies to many scheduling and resource allocation problems."

---

### Q9. LRU Cache Implementation

```java
// O(1) get and put using LinkedHashMap
public class LRUCache {
    private final int capacity;
    private final LinkedHashMap<Integer, Integer> cache;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        // accessOrder=true — reorders on access (most recently used at tail)
        this.cache = new LinkedHashMap<>(capacity, 0.75f, true) {
            @Override
            protected boolean removeEldestEntry(Map.Entry<Integer, Integer> eldest) {
                return size() > capacity;  // auto-evict when full
            }
        };
    }

    public int get(int key) { return cache.getOrDefault(key, -1); }
    public void put(int key, int value) { cache.put(key, value); }
}

// From scratch — HashMap + Doubly Linked List (interview answer)
// For concurrency: use LinkedHashMap wrapped in synchronized or ConcurrentLinkedHashMap
```

---

### 🎯 How to Explain in Interview

"LRU cache is a perfect example of combining data structures for O(1) operations. I use LinkedHashMap with accessOrder set to true, which automatically reorders entries on access. When the cache exceeds capacity, I override removeEldestEntry() to remove the least recently used item. The get() and put() operations are both O(1) because LinkedHashMap provides constant-time access and reordering. For interviews, I should also mention the from-scratch implementation using a HashMap for O(1) lookup and a doubly linked list to maintain access order. The key insight is that LinkedHashMap already implements the LRU behavior when accessOrder is true, making it an elegant solution. For concurrent access, I can wrap it in synchronized or use concurrent collections."

---
