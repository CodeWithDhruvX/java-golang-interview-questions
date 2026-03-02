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
