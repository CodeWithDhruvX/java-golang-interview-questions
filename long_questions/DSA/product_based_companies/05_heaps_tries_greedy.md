# Heaps, Tries, and Greedy Algorithms (Product-Based Companies)

These advanced data structures and algorithms represent the peak of problem-solving tested in rigorous engineering interviews.

## Question 1: Find Median from Data Stream
**Problem Statement:** The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value, and the median is the mean of the two middle values. Implement the `MedianFinder` class:
- `MedianFinder()` initializes the MedianFinder object.
- `void addNum(int num)` adds the integer `num` from the data stream to the data structure.
- `double findMedian()` returns the median of all elements so far.

### Answer:
The optimal approach maintains two heaps: a Max-Heap for the smaller half of numbers and a Min-Heap for the larger half. The Max-Heap is allowed to have at most one more element than the Min-Heap. When adding a number, we first push it to Max-Heap, then pop the max and push it to Min-Heap. If the Min-Heap is larger than the Max-Heap, we balance by moving the min of Min-Heap back to Max-Heap.

**Code Implementation (Java):**
```java
import java.util.Collections;
import java.util.PriorityQueue;

public class MedianFinder {
    private PriorityQueue<Integer> small; // Max-Heap
    private PriorityQueue<Integer> large; // Min-Heap

    public MedianFinder() {
        small = new PriorityQueue<>(Collections.reverseOrder());
        large = new PriorityQueue<>();
    }
    
    public void addNum(int num) {
        small.add(num);
        large.add(small.poll()); // Ensures 'large' contains max of 'small'
        
        // Maintain size property where small is allowed to be larger by 1
        if (small.size() < large.size()) {
            small.add(large.poll());
        }
    }
    
    public double findMedian() {
        if (small.size() > large.size()) {
            return small.peek();
        }
        return (small.peek() + large.peek()) / 2.0;
    }
}
```
**Time Complexity:** O(log N) for `addNum`, O(1) for `findMedian`
**Space Complexity:** O(N)

---

## Question 2: Implement Trie (Prefix Tree)
**Problem Statement:** A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. Implement the `Trie` class with `insert`, `search`, and `startsWith`.

### Answer:
We define a `TrieNode` containing an array of 26 `TrieNode` pointers for children (for lowercase English letters) and a boolean `isEnd` flag. As we process characters, we step down the Trie.

**Code Implementation (Java):**
```java
class TrieNode {
    TrieNode[] children = new TrieNode[26];
    boolean isEnd;
}

public class Trie {
    private TrieNode root;

    public Trie() {
        root = new TrieNode();
    }
    
    public void insert(String word) {
        TrieNode curr = root;
        for (char c : word.toCharArray()) {
            int index = c - 'a';
            if (curr.children[index] == null) {
                curr.children[index] = new TrieNode();
            }
            curr = curr.children[index];
        }
        curr.isEnd = true;
    }
    
    public boolean search(String word) {
        TrieNode node = searchPrefix(word);
        return node != null && node.isEnd;
    }
    
    public boolean startsWith(String prefix) {
        return searchPrefix(prefix) != null;
    }
    
    private TrieNode searchPrefix(String prefix) {
        TrieNode curr = root;
        for (char c : prefix.toCharArray()) {
            int index = c - 'a';
            if (curr.children[index] == null) return null;
            curr = curr.children[index];
        }
        return curr;
    }
}
```
**Time Complexity:** O(M) where M is the keys length for insert/search/startWith
**Space Complexity:** O(M * N) total for building the tree with N words of length M

---

## Question 3: Word Search II (Trie + DFS)
**Problem Statement:** Given an `m x n` board of characters and a list of strings `words`, return all words on the board. Each word must be constructed from letters of sequentially adjacent cells (horizontal/vertical). The same cell may not be used more than once in a word.

### Answer:
Instead of doing pure DFS for every word (which is too slow), we insert all words into a Trie. We then iterate through each cell on the board. If a cell matches a character in the Trie's root children, we perform DFS backtracking, walking down the Trie simultaneously.

**Code Implementation (Java):**
```java
import java.util.ArrayList;
import java.util.List;

class TrieNodeWS2 {
    TrieNodeWS2[] children = new TrieNodeWS2[26];
    String word; // Store the full word at the end node
}

public class WordSearchII {
    public List<String> findWords(char[][] board, String[] words) {
        List<String> res = new ArrayList<>();
        TrieNodeWS2 root = buildTrie(words);
        
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                dfs(board, i, j, root, res);
            }
        }
        return res;
    }
    
    private void dfs(char[][] board, int i, int j, TrieNodeWS2 p, List<String> res) {
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length) return;
        
        char c = board[i][j];
        if (c == '#' || p.children[c - 'a'] == null) return;
        
        p = p.children[c - 'a'];
        if (p.word != null) { // Found a word
            res.add(p.word);
            p.word = null; // Prevent duplicates
        }
        
        board[i][j] = '#'; // Mark visited
        dfs(board, i - 1, j, p, res);
        dfs(board, i + 1, j, p, res);
        dfs(board, i, j - 1, p, res);
        dfs(board, i, j + 1, p, res);
        board[i][j] = c;   // Backtrack
    }
    
    private TrieNodeWS2 buildTrie(String[] words) {
        TrieNodeWS2 root = new TrieNodeWS2();
        for (String w : words) {
            TrieNodeWS2 p = root;
            for (char c : w.toCharArray()) {
                int i = c - 'a';
                if (p.children[i] == null) p.children[i] = new TrieNodeWS2();
                p = p.children[i];
            }
            p.word = w; // Set end of word
        }
        return root;
    }
}
```
**Time Complexity:** O(M * 4^L) where M is cells of board and L maximum length of words. In practice, heavily pruned by the Trie.
**Space Complexity:** O(W * L) for Trie where W is number of words lengths L

---

## Question 4: Jump Game (Greedy Algorithm)
**Problem Statement:** You are given an integer array `nums`. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position. Return `true` if you can reach the last index, or `false` otherwise.

### Answer:
This can be solved very efficiently with a Greedy approach. We keep track of the `maxReach` index we can jump to. As we iterate through the array, if the current index is greater than `maxReach`, it means we can never reach the current position, so we return false. Otherwise, we update `maxReach` to be the maximum of `maxReach` and `i + nums[i]`. If `maxReach` >= last index, we can reach the end.

**Code Implementation (Java):**
```java
public class JumpGame {
    public boolean canJump(int[] nums) {
        int maxReach = 0;
        int target = nums.length - 1;
        
        for (int i = 0; i <= target; i++) {
            if (i > maxReach) {
                return false; // Can't even reach the current index
            }
            maxReach = Math.max(maxReach, i + nums[i]);
            if (maxReach >= target) {
                return true;
            }
        }
        
        return false;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1)
