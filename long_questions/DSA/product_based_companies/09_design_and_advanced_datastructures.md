# Design & Advanced Data Structures (Product-Based Companies)

These questions test your ability to combine multiple core data structures (Maps, Linked Lists, Heaps) to create a custom data structure that meets specific, strict time-complexity requirements. 

## Question 1: LFU Cache (Least Frequently Used)
**Problem Statement:** Design and implement a data structure for a Least Frequently Used (LFU) cache. Implement the `LFUCache` class:
- `LFUCache(int capacity)`: Initializes the object with the capacity of the data structure.
- `int get(int key)`: Gets the value of the `key` if the key exists in the cache. Otherwise, returns `-1`.
- `void put(int key, int value)`: Update the value of the `key` if present, or inserts the `key` if not already present. When the cache reaches its capacity, it should invalidate and remove the **least frequently used** key. If there is a tie, remove the **least recently used** key.

Requires O(1) time complexity for both `get` and `put`.

### Answer:
To achieve O(1), we need multiple HashMaps and a custom Doubly Linked List.
1. `keyNode` Map: Maps `key` -> Node (for O(1) lookups).
2. `freqList` Map: Maps `frequency` -> Doubly Linked List of Nodes with that frequency.
3. Keep track of `minFreq` so we know which list to delete from when at capacity.

**Code Implementation (Java):**
```java
import java.util.HashMap;
import java.util.Map;

class LFUNode {
    int key, val, freq;
    LFUNode prev, next;
    LFUNode(int k, int v) {
        key = k;
        val = v;
        freq = 1;
    }
}

class DLList {
    LFUNode head, tail;
    int size;
    
    DLList() {
        head = new LFUNode(0, 0);
        tail = new LFUNode(0, 0);
        head.next = tail;
        tail.prev = head;
        size = 0;
    }
    
    void addNode(LFUNode node) {
        node.next = head.next;
        node.next.prev = node;
        head.next = node;
        node.prev = head;
        size++;
    }
    
    void removeNode(LFUNode node) {
        node.prev.next = node.next;
        node.next.prev = node.prev;
        size--;
    }
    
    LFUNode removeTail() {
        if (size > 0) {
            LFUNode node = tail.prev;
            removeNode(node);
            return node;
        }
        return null;
    }
}

public class LFUCache {
    int capacity, size, minFreq;
    Map<Integer, LFUNode> keyNode;
    Map<Integer, DLList> freqList;

    public LFUCache(int capacity) {
        this.capacity = capacity;
        this.size = 0;
        this.minFreq = 0;
        this.keyNode = new HashMap<>();
        this.freqList = new HashMap<>();
    }
    
    public int get(int key) {
        if (!keyNode.containsKey(key)) return -1;
        LFUNode node = keyNode.get(key);
        updateNode(node);
        return node.val;
    }
    
    public void put(int key, int value) {
        if (capacity == 0) return;
        
        if (keyNode.containsKey(key)) {
            LFUNode node = keyNode.get(key);
            node.val = value;
            updateNode(node);
        } else {
            if (size == capacity) {
                // Remove LFU/LRU
                DLList list = freqList.get(minFreq);
                LFUNode toRemove = list.removeTail();
                keyNode.remove(toRemove.key);
                size--;
            }
            
            size++;
            minFreq = 1;
            LFUNode newNode = new LFUNode(key, value);
            
            keyNode.put(key, newNode);
            freqList.putIfAbsent(1, new DLList());
            freqList.get(1).addNode(newNode);
        }
    }
    
    private void updateNode(LFUNode node) {
        int oldFreq = node.freq;
        DLList list = freqList.get(oldFreq);
        list.removeNode(node);
        
        if (oldFreq == minFreq && list.size == 0) {
            minFreq++;
        }
        
        node.freq++;
        freqList.putIfAbsent(node.freq, new DLList());
        freqList.get(node.freq).addNode(node);
    }
}
```
**Time Complexity:** O(1) for both `get` and `put`.
**Space Complexity:** O(Capacity)

---

## Question 2: Insert Delete GetRandom O(1)
**Problem Statement:** Implement the `RandomizedSet` class:
- `bool insert(int val)` Inserts an item `val` into the set if not present.
- `bool remove(int val)` Removes an item `val` from the set if present.
- `int getRandom()` Returns a random element from the current set of elements (all equally likely).

You must implement the functions with an average `O(1)` time complexity.

### Answer:
To achieve O(1) random access, we need a dynamic array (ArrayList). But removing from an array in O(1) is impossible unless we remove the *last* element. So, we use a HashMap to map values to their indices in the list. During removal, we swap the element-to-remove with the last element in the list, update the map, and remove the last element.

**Code Implementation (Java):**
```java
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Random;

public class RandomizedSet {
    private Map<Integer, Integer> map;
    private List<Integer> list;
    private Random rand;

    public RandomizedSet() {
        map = new HashMap<>();
        list = new ArrayList<>();
        rand = new Random();
    }
    
    public boolean insert(int val) {
        if (map.containsKey(val)) return false;
        map.put(val, list.size()); // Store value to index
        list.add(val);
        return true;
    }
    
    public boolean remove(int val) {
        if (!map.containsKey(val)) return false;
        
        int indexToRemove = map.get(val);
        int lastElement = list.get(list.size() - 1);
        
        // Swap last element into the index of element dropping out
        list.set(indexToRemove, lastElement);
        map.put(lastElement, indexToRemove);
        
        // Remove the true last element
        list.remove(list.size() - 1);
        map.remove(val);
        
        return true;
    }
    
    public int getRandom() {
        return list.get(rand.nextInt(list.size()));
    }
}
```
**Time Complexity:** O(1) average for all operations. (Amortized array resizing).
**Space Complexity:** O(N) where N is elements.

---

## Question 3: Design Add and Search Words Data Structure
**Problem Statement:** Design a data structure that supports adding new words and finding if a string matches any previously added string.
Implement the `WordDictionary` class:
- `void addWord(word)` Adds `word` to the data structure.
- `bool search(word)` Returns `true` if there is any string matched in the structure, `false` otherwise. `word` may contain dots `.` where dots can be matched with any letter.

### Answer:
This is heavily based on a **Trie (Prefix Tree)**. `addWord` is standard Trie insertion. For `search`, when we hit a `.`, we must recursively check all 26 possible children of the current TrieNode to see if any branch leads to a valid word match.

**Code Implementation (Java):**
```java
class TrieNodeWD {
    TrieNodeWD[] children = new TrieNodeWD[26];
    boolean isEnd = false;
}

public class WordDictionary {
    private TrieNodeWD root;

    public WordDictionary() {
        root = new TrieNodeWD();
    }
    
    public void addWord(String word) {
        TrieNodeWD curr = root;
        for (char c : word.toCharArray()) {
            int index = c - 'a';
            if (curr.children[index] == null) {
                curr.children[index] = new TrieNodeWD();
            }
            curr = curr.children[index];
        }
        curr.isEnd = true;
    }
    
    public boolean search(String word) {
        return searchInNode(word, 0, root);
    }
    
    private boolean searchInNode(String word, int index, TrieNodeWD node) {
        if (index == word.length()) {
            return node.isEnd;
        }
        
        char c = word.charAt(index);
        if (c == '.') {
            // Wildcard: Check all 26 possible children
            for (TrieNodeWD child : node.children) {
                if (child != null && searchInNode(word, index + 1, child)) {
                    return true;
                }
            }
            return false;
        } else {
            // Normal character check
            TrieNodeWD child = node.children[c - 'a'];
            if (child == null) return false;
            return searchInNode(word, index + 1, child);
        }
    }
}
```
**Time Complexity:** `addWord`: O(N). `search`: O(N) if no dots, O(26^M) in worst case purely dots.
**Space Complexity:** O(N) total characters added.
