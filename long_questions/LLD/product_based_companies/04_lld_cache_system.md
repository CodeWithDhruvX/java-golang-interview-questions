# Low-Level Design (LLD) - Cache System (LRU / LFU)

## Problem Statement
Design an in-memory cache component that supports fast `get()` and `put()` operations, along with an eviction policy. The standard cache size is restricted. Implement a generic Cache and specifically an LRU (Least Recently Used) cache.

## Requirements
*   **Operations:** Must implement `put(key, value)` and `get(key)` in `O(1)` average time complexity.
*   **Eviction:** When cache capacity is full, an eviction algorithm (LRU, LFU, FIFO) must evict a key before adding the new key.
*   **Thread Safety:** The cache should handle concurrent reading and writing.

## Core Entities / Classes

1.  **Cache Interface:** `get(K)`, `put(K, V)`, `remove(K)`.
2.  **Capacity Constraint:** Max size limit.
3.  **Storage Engine:** Generally a Hash Map (`Map<K, Node>`) for `O(1)` access.
4.  **Eviction Algorithm Mechanism:** A Doubly Linked List is traditionally used alongside the Hash Map for LRU.
    *   `Node`: Prev, Next, Key, Value.

## Key Design Patterns Applicable
*   **Strategy Pattern:** Injecting the Eviction Policy (LRU, LFU) at runtime.
*   **Factory Pattern:** `CacheBuilder` or `CacheFactory` to instantiate cache instances easily.
*   **Observer/Listener Pattern (Optional):** Emitting cache hit/miss/eviction events to a metrics collector.

## Code Snippet (LRU Concept - Doubly Linked List + HashMap)

```java
class Node {
    int key, val;
    Node prev, next;
    public Node(int k, int v) { this.key = k; this.val = v; }
}

public class LRUCache {
    private final int capacity;
    private final Map<Integer, Node> map;
    private final Node head, tail; // Dummy nodes

    public LRUCache(int cap) {
        this.capacity = cap;
        this.map = new HashMap<>();
        head = new Node(0, 0);
        tail = new Node(0, 0);
        head.next = tail;
        tail.prev = head;
    }

    public int get(int key) {
        if (!map.containsKey(key)) return -1;
        Node node = map.get(key);
        remove(node); // bring to front
        insert(node);
        return node.val;
    }

    // Insert, Remove side operations...
}
```

## Follow-up Questions for Candidate
1.  How do you make this cache thread-safe? (Discuss `ConcurrentHashMap` vs `ReentrantWriteReadLock`).
2.  How would you implement an LFU (Least Frequently Used) cache? (Requires priority queue or two HashMaps and a Doubly linked list array).
3.  How can you implement a Time-To-Live (TTL) eviction strategy?
