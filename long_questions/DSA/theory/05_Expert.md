# 🟣 **41–50: Expert (Niche Algorithms & FAANG Hard Optimizations)**

### 41. Explain the Z-Algorithm.
"The **Z-Algorithm** is a linear-time string matching algorithm similar to KMP, but calculates a 'Z-array' instead of an LPS array. The `Z[i]` value represents the length of the longest substring starting from index `i` that perfectly matches the prefix of the entire string.

I use it by concatenating the `Pattern + '$' + Text`. Then I run the Z-algorithm on this combined string. Any index where `Z[i]` equals the length of my original Pattern is a match! It's slightly easier to conceptualize than KMP because it directly answers 'does a match start exactly here?'"

#### Indepth
The complexity is **O(N + M)** time and **O(N + M)** space (for the concatenated string and Z-array). It requires maintaining a 'Z-box' (a window representing the right-most match found so far) to skip redundant comparisons. Usually interchangeable with KMP for pattern matching length/occurrences.

---

### 42. How are Randomized Algorithms used in interviews?
"A **Randomized Algorithm** uses a random number generator to dictate its control flow. Instead of deterministic worst-case scenarios defeating me, I rely on mathematics guaranteeing average-case brilliance.

The quintessential example is **QuickSelect** (finding the K-th largest element). If I pick a pivot deterministically (like the first element), a sorted array forces an **O(N^2)** worst case. But if I pick a random pivot, probability dictates I will eliminate a large chunk of the array every time, achieving an expected **O(N)** time."

#### Indepth
For QuickSelect, the expected Time is **O(N)**, but technically the worst-case remains **O(N^2)** (though practically impossible). Space is **O(1)**. Interviewers love it because it tests both array partitioning logic and a deep understanding of algorithmic probability.

---

### 43. What is Cyclic Sort and when is it applicable?
"**Cyclic Sort** is an incredibly elegant pattern used exclusively when a problem statement mentions 'an array containing integers in the range 1 to N' (or 0 to N).

Since the elements are constrained to be exact indices of the array, I can sort it without any extra memory. I iterate through the array. If the number at `nums[i]` is, say, 5, I swap it with whatever is at `nums[4]`. I repeat this swap until the number 5 is finally sitting in index 4. Then I move to `i+1`."

#### Indepth
Because every number is swapped into its correct place at most once, the time complexity is strictly **O(N)**. Because it uses no extra arrays or hash maps, the space is strictly **O(1)**. It is essentially a 'magic trick' for problems like *Find the Missing Number* or *First Missing Positive* that demand O(1) space.

---

### 44. Discuss the Floyd-Warshall Algorithm.
"**Floyd-Warshall** is a dynamic programming algorithm used to find the shortest paths between **every mathematically possible pair** of nodes in a weighted graph.

Instead of running Dijkstra's from every single node, FW asks a simple question iteratively: 'Is the path from City A to City B shorter if I route it through City K?'. It wraps this in three nested loops (`for K, for I, for J`) to systematically update an adjacency matrix."

#### Indepth
The three nested loops give it a time complexity of **O(V^3)**. Space is **O(V^2)** for the matrix. Because $V^3$ grows massively, this algorithm is only viable when the graph is very small (usually V <= 400). It's great for 'All-Pairs Shortest Path' or Transitive Closure problems (*Evaluate Division*).

---

### 45. What is a Minimum Spanning Tree (MST)?
"In a connected, weighted graph, a **Minimum Spanning Tree** is a subset of the edges that connects all vertices together without any cycles, using the absolute minimum total edge weight. Imagine connecting 10 islands with bridges using the cheapest possible total budget.

I have two choices:
- **Kruskal's**: Sort all edges by weight. Pick the cheapest edge. If adding it doesn't create a cycle (using a Union-Find), keep it. Repeat until all nodes are connected.
- **Prim's**: Start at a random node. Always pick the cheapest edge connecting the 'visited' group to any 'unvisited' node (using a Priority Queue)."

#### Indepth
Both run in **O(E log E)** or **O(E log V)** time. Space is **O(V + E)**. Kruskal's is generally easier to code if I already have a Union-Find template memorized. Typical problems include *Min Cost to Connect All Points*.

---

### 46. What is Reservoir Sampling?
"**Reservoir Sampling** is a randomized algorithm used for randomly selecting `K` items from a continuous stream of data where the total size `N` is unknown (and you can't load the whole stream into memory).

If `K=1` (pick one random item), I unconditionally save the first item. When the second item arrives, I keep it with a $1/2$ probability. The third item with $1/3$, the N-th item with $1/N$. Mathematically, this guarantees that at any point, every item seen so far had an equal $1/N$ probability of being in my 'reservoir'."

#### Indepth
Time complexity is **O(N)** as I only iterate the stream once. Space is **O(K)** (the reservoir size). It is a highly specific but critical algorithm for Big Data interviews or questions like *Linked List Random Node*.

---

### 47. How do you find Strongly Connected Components (SCCs)?
"In a directed graph, a **Strongly Connected Component** is a maximal set of vertices where every vertex can reach every other vertex. If you leave an SCC, you cannot trace a path back into it.

I usually use **Tarjan's Algorithm** for this. It uses a single DFS pass to assign 'discovery times' to nodes. I maintain a `low_link` value representing the lowest discovery time reachable from that node. If a node's `low_link` equals its own `discovery_time` after visiting all neighbors, it is the 'root' of an SCC."

#### Indepth
Tarjan's runs efficiently in **O(V + E)** time and uses **O(V)** space for the tracking arrays and stack. Understanding Tarjan's (or alternatively Kosaraju's two-pass DFS) is the pinnacle of graph connectivity questions on LeetCode hard problems (*Critical Connections in a Network*).

---

### 48. What is a Morris Traversal?
"**Morris Traversal** is a way to traverse a Binary Tree in perfectly linear time **without using a stack or recursion**.

Standard traversal requires **O(H)** space to remember where to return to. Morris bypasses this by temporarily modifying the tree. When at a node, it finds the right-most node in the left subtree and creates a temporary pointer back to the current node. It uses these built-in 'threads' to find its way back up, destroying the temporary links as it goes to restore the tree."

#### Indepth
Time is **O(N)** (each edge is traversed at most 3 times). Space is strictly **O(1)**. It is purely an optimization pattern used when an interviewer forces an explicit O(1) space constraint on tree traversals, like in *Recover Binary Search Tree*.

---

### 49. How do Non-Comparison Sorting algorithms work?
"Standard algorithms like Merge or Quick sort compare two elements directly (`if a > b`), creating a mathematical speed limit of **O(N log N)**.

**Non-Comparison Sorts** (Bucket Sort, Counting Sort, Radix Sort) bypass this by using the inherent *value* of the data as an index. If I know my array only contains numbers 0-100, I can just create an array of size 101, tally the occurrences, and print them out. No comparisons needed!"

#### Indepth
Time complexity plummets to **O(N)**. The trade-off is Space Complexity, which becomes **O(K)**, where K is the range of possible values. If the numbers range from 0 to 1 Billion, non-comparison sorts require too much memory and become useless (e.g., *Maximum Gap*, *Sort Colors*).

---

### 50. Explain the Divide and Conquer strategy.
"**Divide and Conquer** is the recursive strategy of breaking a large problem into multiple non-overlapping subproblems, solving them, and merging the results.

The most famous example is **Merge Sort**: dividing an array in half until sizes are 1, then merging the sorted halves. In advanced interviews, it's used for complex geometry or array counting. For example, in *Count Inversions*, instead of an $O(N^2)$ brute force, I modify Merge Sort to count how many items from the right half are smaller than items in the left half during the merge step."

#### Indepth
Time complexity is extremely consistent at **O(N log N)** (or similar, based on the Master Theorem). Space is usually **O(N)** or **O(log N)**. While D&C is often overshadowed by DP, recognizing when subproblems are **independent** (D&C) vs overlapping (DP) is a hallmark of a Senior/Expert engineer.
