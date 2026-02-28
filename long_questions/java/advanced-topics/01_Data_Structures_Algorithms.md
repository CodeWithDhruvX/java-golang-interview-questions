# Data Structures & Algorithms (DSA) for Product-Based Companies

Product-based companies (Amazon, Uber, Atlassian, Google, Microsoft) rely heavily on DSA to filter candidates. You won't just be asked to explain a data structure; you will be asked to solve a problem under time pressure and write clean, optimal code.

Here is a breakdown of the most frequently asked DSA topics and associated mental models.

---

## 1. Arrays & Hashing (The Gatekeepers)

These questions test your basic problem-solving and ability to trade space for time (using a HashMap/HashSet).

### Core Concepts:
*   **Two Pointers:** Moving from opposite ends towards the center, or moving in the same direction at varying speeds.
*   **Sliding Window:** Maintaining a dynamic window (subarray/substring) to track properties like maximum sum, longest substring without repeating characters, etc.
*   **Prefix Sum / Hashing:** Keeping track of accumulated data to answer range queries in O(1) time.

### Must-Know Questions:
*   **Two Sum / Three Sum:** Finding pairs/triplets that add up to a target. (Focus on O(N) or O(N^2) optimization).
*   **Longest Substring Without Repeating Characters:** Classic sliding window problem. O(N) using a HashSet or HashMap.
*   **Container With Most Water:** Two-pointer approach narrowing inwards.
*   **Minimum Window Substring:** Advanced sliding window using character frequency maps.
*   **Product of Array Except Self:** Prefix and suffix array products.

---

## 2. Trees & Graphs (The Heavyweights)

These structures form the backbone of mostly all interview loops.

### Core Concepts:
*   **Trees:** Binary Trees, Binary Search Trees (BST), Heaps/Priority Queues.
*   **Graphs:** Adjacency List representations, Directed/Undirected, Weighted/Unweighted.
*   **DFS (Depth-First Search):** Recursion or Stack. Good for going deep (e.g., detecting cycles, finding paths).
*   **BFS (Breadth-First Search):** Queue. Good for finding the shortest path in an unweighted graph or processing level by level.

### Must-Know Questions (Trees):
*   **Maximum Depth of Binary Tree:** Simple DFS or BFS.
*   **Invert / Flip Binary Tree:** Recursive swapping of left and right children.
*   **Validate Binary Search Tree:** Ensuring left < node < right constraint holds recursively.
*   **Lowest Common Ancestor (LCA) of a Binary Tree:** Finding where two nodes' paths diverge.
*   **Binary Tree Maximum Path Sum:** Post-order traversal keeping track of max sums.

### Must-Know Questions (Graphs):
*   **Number of Islands:** Classic matrix traversal (DFS or BFS).
*   **Course Schedule (Topological Sort):** Detecting cycles in a directed graph using DFS (visited sets) or Kahn's Algorithm (in-degree).
*   **Clone Graph:** Deep copying a graph using a hash map to track visited nodes.
*   **Word Ladder:** Shortest path in an unweighted graph using BFS.
*   **Dijkstra's Algorithm:** Shortest path in a weighted graph (used in Uber/Google router questions).

---

## 3. Dynamic Programming (The Fear Inducers)

DP involves breaking a problem down into smaller overlapping subproblems and caching their results to avoid redundant work.

### Core Concepts:
*   **Top-Down (Memoization):** Recursion + caching results in an array or map.
*   **Bottom-Up (Tabulation):** Iteratively building up solutions from the smallest subproblem using an array/matrix.

### Must-Know Questions:
*   **Climbing Stairs / Fibonacci:** The fundamentals of 1D DP.
*   **Coin Change:** Unbounded knapsack problem. Finding the minimum coins to make an amount.
*   **Longest Increasing Subsequence (LIS):** Classic O(N^2) DP pattern.
*   **Longest Common Subsequence (LCS):** Classic 2D grid DP pattern (comparing two strings).
*   **Word Break:** String segmentation using a boolean DP array.

---

## 4. Linked Lists (The Pointer Puzzles)

Linked lists test your careful manipulation of pointers and handling of edge cases (like null references or single-node lists).

### Core Concepts:
*   **Fast and Slow Pointers:** One pointer moves by 1 step, the other by 2. Used to find the middle or detect cycles.
*   **Dummy Nodes:** Creating a fake node at the head to simplify edge cases when the head itself might be deleted or modified.

### Must-Know Questions:
*   **Reverse a Linked List:** The absolute classic. O(N) time, O(1) space.
*   **Merge Two Sorted Lists:** Pointers racing through two lists.
*   **Linked List Cycle:** Floyd's Tortoise and Hare algorithm.
*   **LRU Cache (Super Important):** Combines a Doubly Linked List with a HashMap for O(1) get and put operations. Often asked as a standalone design question.
*   **Copy List with Random Pointer:** HashMap mapping original nodes to clone nodes.

---

## 5. Sorting and Searching

You rarely need to write a sorting algorithm from scratch, but you must know their time complexities and properties.

### Core Concepts:
*   **Binary Search:** O(log N) search on a *sorted* dataset.
*   **Merge Sort / Quick Sort:** O(N log N) sorting. Know the difference between stable and unstable sorts.

### Must-Know Questions:
*   **Binary Search:** The absolute basics.
*   **Search in Rotated Sorted Array:** Applying binary search when one half is guaranteed to be sorted.
*   **Find Minimum in Rotated Sorted Array:** Using binary search to find the pivot.
*   **Merge Intervals:** Sorting by start time and overlapping ranges.
*   **Kth Largest Element in an Array:** Using a Min-Heap of size K (O(N log K)) or Quickselect (O(N) average).

---

## Tips for Product-Based DSA Interviews:
1.  **Clarify First:** Don't write code immediately. Ask about edge cases (empty inputs, negative numbers, duplicates).
2.  **Brute Force First (If Stuck):** It's better to give a naive O(N^3) solution than nothing at all.
3.  **Think Out Loud:** The interviewer wants to know *how* you think, not just if you memorized the answer.
4.  **Analyze Complexity:** Always state the Time (Big O) and Space complexity before and after you write the code.
5.  **Test Your Code:** Do a dry run with a small sample input after writing to catch off-by-one errors.
