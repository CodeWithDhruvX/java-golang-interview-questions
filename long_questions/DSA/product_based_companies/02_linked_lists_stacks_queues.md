# Linked Lists, Stacks, and Queues (Product-Based Companies)

At top product companies, Linked Lists aren't just reversed; they involve complex multi-node manipulations. Stacks and Queues often drive monotonic logic or caching algorithms (like LRU/LFU). 

## Question 1: LRU Cache
**Problem Statement:** Design a data structure that follows the constraints of a Least Recently Used (LRU) cache. Implement the `LRUCache` class:
- `LRUCache(int capacity)` Initialize the LRU cache with positive size `capacity`.
- `int get(int key)` Return the value of the `key` if the key exists, otherwise return `-1`.
- `void put(int key, int value)` Update the value of the `key` if the `key` exists. Otherwise, add the `key-value` pair to the cache. If the number of keys exceeds the `capacity` from this operation, evict the least recently used key.

Both `get` and `put` must run in O(1) average time complexity.

### Answer:
To achieve O(1) time complexity for both `get` and `put`, we combine a HashMap (for O(1) access) and a Doubly Linked List (for O(1) removal/addition of nodes at the ends). The most recently used item is moved to the head, and the least recently used item sits at the tail.

**Code Implementation (Java):**
```java
import java.util.HashMap;

class Node {
    int key, value;
    Node prev, next;
    Node(int k, int v) { key = k; value = v; }
}

public class LRUCache {
    private int capacity;
    private HashMap<Integer, Node> map;
    private Node head, tail;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        map = new HashMap<>();
        head = new Node(0, 0); // Dummy head
        tail = new Node(0, 0); // Dummy tail
        head.next = tail;
        tail.prev = head;
    }
    
    public int get(int key) {
        if (map.containsKey(key)) {
            Node node = map.get(key);
            remove(node); // Remove from current position
            insert(node); // Move to head (most recently used)
            return node.value;
        }
        return -1;
    }
    
    public void put(int key, int value) {
        if (map.containsKey(key)) {
            remove(map.get(key));
        }
        if (map.size() == capacity) {
            remove(tail.prev); // Evict least recently used (tail's prev)
        }
        insert(new Node(key, value)); // Insert new node at head
    }
    
    private void remove(Node node) {
        map.remove(node.key);
        node.prev.next = node.next;
        node.next.prev = node.prev;
    }
    
    private void insert(Node node) {
        map.put(node.key, node);
        node.next = head.next;
        node.next.prev = node;
        head.next = node;
        node.prev = head;
    }
}
```
**Time Complexity:** O(1) for both `get` and `put`
**Space Complexity:** O(capacity)

---

## Question 2: Merge k Sorted Lists
**Problem Statement:** You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order. Merge all the linked-lists into one sorted linked-list and return it.

### Answer:
An elegant and optimal way to solve this is using a Min-Heap (PriorityQueue). We put the head of each linked list into the heap. Then we continuously extract the minimum node, append it to our result list, and if that extracted node has a `next` node, we push the `next` node into the heap.

**Code Implementation (Java):**
```java
import java.util.PriorityQueue;

class ListNode {
    int val;
    ListNode next;
    ListNode(int x) { val = x; }
}

public class MergeKSortedLists {
    public ListNode mergeKLists(ListNode[] lists) {
        if (lists == null || lists.length == 0) return null;
        
        PriorityQueue<ListNode> pq = new PriorityQueue<>((a, b) -> a.val - b.val);
        
        for (ListNode list : lists) {
            if (list != null) {
                pq.add(list);
            }
        }
        
        ListNode dummy = new ListNode(0);
        ListNode current = dummy;
        
        while (!pq.isEmpty()) {
            ListNode node = pq.poll();
            current.next = node;
            current = current.next;
            
            if (node.next != null) {
                pq.add(node.next);
            }
        }
        
        return dummy.next;
    }
}
```
**Time Complexity:** O(N log k) where N is the total number of nodes and k is the number of linked lists.
**Space Complexity:** O(k) for the priority queue.

---

## Question 3: Sliding Window Maximum
**Problem Statement:** You are given an array of integers `nums`, there is a sliding window of size `k` which is moving from the very left of the array to the very right. You can only see the `k` numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

### Answer:
We can solve this efficiently in O(N) time using a Monotonic Deque (Double-ended Queue). We store indices in the deque. The deque maintains elements in monotonically decreasing order so that the maximum element of the current window is always at the front.

**Code Implementation (Java):**
```java
import java.util.Deque;
import java.util.LinkedList;

public class SlidingWindowMaximum {
    public int[] maxSlidingWindow(int[] nums, int k) {
        if (nums == null || k <= 0) return new int[0];
        int n = nums.length;
        int[] result = new int[n - k + 1];
        int ri = 0;
        
        Deque<Integer> q = new LinkedList<>();
        for (int i = 0; i < nums.length; i++) {
            // Remove indices out of current window
            while (!q.isEmpty() && q.peek() < i - k + 1) {
                q.poll();
            }
            // Remove smaller numbers in k range as they are useless
            while (!q.isEmpty() && nums[q.peekLast()] < nums[i]) {
                q.pollLast();
            }
            q.offer(i);
            
            // The window has started forming
            if (i >= k - 1) {
                result[ri++] = nums[q.peek()];
            }
        }
        return result;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(k)

---

## Question 4: Reverse Nodes in k-Group
**Problem Statement:** Given the `head` of a linked list, reverse the nodes of the list `k` at a time, and return the modified list. If the number of nodes is not a multiple of `k` then left-out nodes, in the end, should remain as it is.

### Answer:
We check if there are `k` nodes available to reverse. If there are, we reverse them using typical linked list reversal logic and connect the previously reversed segment to the current reversing segment's head. Recursion or an iterative approach with `dummy` nodes works well here.

**Code Implementation (Java):**
```java
public class ReverseKGroup {
    public ListNode reverseKGroup(ListNode head, int k) {
        if (head == null || k == 1) return head;
        
        ListNode dummy = new ListNode(0);
        dummy.next = head;
        
        ListNode curr = dummy, nex = dummy, pre = dummy;
        int count = 0;
        
        while (curr.next != null) {
            curr = curr.next;
            count++;
        }
        
        while (count >= k) {
            curr = pre.next;
            nex = curr.next;
            for (int i = 1; i < k; i++) {
                curr.next = nex.next;
                nex.next = pre.next;
                pre.next = nex;
                nex = curr.next;
            }
            pre = curr;
            count -= k;
        }
        return dummy.next;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1)
