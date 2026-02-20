# 🟡 **11–20: Intermediate (Trees, Graphs, and Basic Optimization)**

### 11. Explain Tree BFS (Level Order) and its typical use case.
"**Tree BFS** (Breadth-First Search) involves exploring a tree level by level, from top to bottom, left to right. Instead of diving deep like DFS, I visit all nodes at the current depth before moving deeper.

I implement this using a **Queue**. I add the root, then in a `while` loop, I pop a node, visit it, and push its children. I use this pattern when I need to find the shortest path in an unweighted tree/graph, or when the problem explicitly asks for level-by-level data, like printing the *Right Side View* of a binary tree."

#### Indepth
BFS operates in **O(N)** time, as it visits every node. However, its space complexity is **O(W)**, where W is the maximum width of the tree (at the bottom level, this can be N/2, so effectively **O(N)**). It is fundamentally different from DFS in its memory footprint shape (Queue vs Stack).

---

### 12. What is the difference between Graph DFS and BFS?
"Both are used to traverse graphs, but they serve different goals. 
**Graph BFS** expands outward like ripples in a pond. It's the go-to for finding the **shortest path** in unweighted networks (like finding the fewest hops between two users).
**Graph DFS** goes deep down one path. It's excellent for checking **connectivity** (can I reach B from A?) or finding cycles in a graph.

The critical difference from Tree traversals is that graphs have cycles, so I must maintain a `visited` set to avoid infinite loops."

#### Indepth
Both algorithms run in **O(V + E)** time, where V is vertices and E is edges. Space is **O(V)** for the queue/stack and the visited set. Knowing when to choose BFS (shortest path) vs DFS (connectivity/backtracking) is a core FAANG interview competency (e.g., *Number of Islands* vs *Word Ladder*).

---

### 13. How does the Union-Find (Disjoint Set) data structure work?
"**Union-Find** is a specialized data structure designed to keep track of elements partitioned into non-overlapping sets. It provides two near-instant operations:
1.  **Find**: Determines which subset an element is in (who is its 'leader').
2.  **Union**: Merges two subsets into a single one.

I use this heavily for network connectivity problems, like determining if adding an edge creates a cycle in a graph, or grouping friends together (e.g., *Number of Provinces*)."

#### Indepth
When implemented with **Path Compression** (flattening the tree during `Find`) and **Union by Rank** (attaching the smaller tree to the larger), the time complexity drops to **O(α(N))**, where α is the inverse Ackermann function, effectively **O(1)** constants time. Space is **O(N)** for the parent array.

---

### 14. What is the Backtracking pattern?
"**Backtracking** is an algorithmic technique for considering all possible combinations to solve a problem incrementally. It's like exploring a maze: I 'Choose' a path, 'Explore' it recursively, and if it leads to a dead-end, I 'Un-choose' (backtrack) and try the next path.

I use this for constraint satisfaction problems like the *Sudoku Solver*, or when I need to generate all *Permutations* or *Subsets* of an array."

#### Indepth
Because Backtracking explores all possibilities, its time complexity is usually exponential, like **O(2^N)** for subsets or **O(N!)** for permutations. Space is **O(N)** due to the recursion stack. The key to mastering this is writing clean `choose -> explore -> unchoose` logic.

---

### 15. Describe 1-Dimensional Dynamic Programming (1D DP).
"**1D Dynamic Programming** is a technique to solve complex, overlapping subproblems by remembering past results rather than recalculating them. 

Imagine climbing stairs: to reach step 10, you must come from step 9 or 8. If I store the number of ways to reach step 9 and 8 in an array (`dp`), I can find step 10 in O(1) time by adding them. I use this for optimization problems with one changing variable, like *Climbing Stairs* or *House Robber*."

#### Indepth
1D DP usually reduces exponential time complexity to **O(N)**. Space complexity is initially **O(N)** for the `dp` array, but if the current state only depends on the previous 1 or 2 states (like Fibonacci calculation), I can reduce the space to **O(1)** by just using a few variables.

---

### 16. How does 2-Dimensional Dynamic Programming (2D DP) differ from 1D DP?
"**2D DP** builds on 1D DP but handles problems where the 'state' depends on two variables instead of one. 

I usually visualize this as filling out a grid. For example, in string matching problems like *Longest Common Subsequence*, the grid represents the indices of string A on the X-axis and string B on the Y-axis. The value of cell `[i][j]` depends on the cells above or to the left of it."

#### Indepth
The complexity reflects the grid size: Time and Space are both **O(N * M)**. This pattern is standard for grid-travel problems (*Unique Paths*), String Edit Distance, and classic Knapsack problems. Sometimes space can be optimized to **O(M)** by only keeping the current and previous row in memory.

---

### 17. What is the Greedy pattern and when is it safe to use?
"A **Greedy** algorithm makes the most optimal, locally 'best' choice at each step without ever looking back or reconsidering past choices. 

I use it when a problem asks for optimization (like min/max) AND I can prove that local optimums lead to a global optimum. A classic example is *Merge Intervals* or interval scheduling: sorting by start/end times and aggressively taking the best next step usually yields the correct answer much faster than DP."

#### Indepth
Greedy algorithms are incredibly fast, usually **O(N log N)** because they almost always require sorting the input first, followed by an **O(N)** traversal. The space is usually **O(1)** or **O(N)** depending on the sorting algorithm. The hardest part is proving mathematically that the greedy choice actually works.

---

### 18. How do you utilize a Heap / Priority Queue?
"A **Heap** (or Priority Queue) is a data structure that allows me to continuously extract the 'best' (minimum or maximum) element from a dynamic dataset in logarithmic time.

I use this pattern anytime a problem asks for 'Top K items', 'K-th largest element', or when I need to merge multiple sorted streams. Instead of sorting an entire array (which is O(N log N)), I can maintain a Heap of size K to find the answer efficiently."

#### Indepth
Heaps offer **O(log N)** time for insertions and extractions. For 'Top K' problems, maintaining a heap of size K takes **O(N log K)** time, which is much faster than full sorting when K is small. Space complexity is **O(K)** or **O(N)** depending on the heap size.

---

### 19. When would you use Bit Manipulation?
"**Bit Manipulation** involves using bitwise operators (AND, OR, XOR, Shift) to directly modify numbers at the binary level. 

I use this for extreme low-level optimizations. A classic interview pattern is using XOR (`^`) because an element XOR'd with itself is 0, which perfectly solves 'find the single non-repeating element' in an array. It's also utilized for representing small sets (masks) efficiently."

#### Indepth
Bit manipulation operations are effectively **O(1)** time constants and **O(1)** space. They are less common in general web development but frequently appear in systems programming roles or FAANG interviews to test fundamental computer science knowledge (e.g., *Counting Bits*, *Power of Two*).

---

### 20. What is a Trie (Prefix Tree) and what problems does it solve?
"A **Trie** is a specialized tree data structure designed for ultra-fast string prefix lookups. Each node represents a character, and moving down the tree spells a word.

Instead of storing 'Cat', 'Car', and 'Cap' separately in a list, a Trie stores the 'C' -> 'a' branch once, saving enormous space and making search instantaneous. I use this pattern exclusively for dictionary implementations, Autocomplete systems, and Spell Checkers."

#### Indepth
Building the Trie takes **O(N * L)** time and space, where N is the number of words and L is the average word length. However, searching for a word or prefix takes only **O(L)** time, completely independent of how many millions of words are in the dictionary. This makes it infinitely better than a Hash Map for prefix-matching tasks.
