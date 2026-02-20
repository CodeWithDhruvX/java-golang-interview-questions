# 🔴 **31–40: Senior (Advanced DP, Trees, and Optimization Algorithms)**

### 31. What is Digit DP and when do you use it?
"**Digit DP** is a highly specialized form of Dynamic Programming used exclusively for counting numbers within a specific range `[L, R]` that satisfy a certain property (e.g., 'Numbers where the sum of digits is 10' or 'Numbers without consecutive 1s in binary').

Instead of iterating from L to R, which is too slow (O(R)), I build the number digit by digit from left to right. My DP state usually tracks the current `index` (digit position), a boolean `is_tight` (to ensure I don't exceed the upper bound R), and whatever property I am tracking (like a running sum)."

#### Indepth
The complexity drops phenomenally from **O(R)** to **O(Digits * 10 * State)**. Since an integer up to $10^{18}$ only has 18 digits, this algorithm runs in near-constant time. It is a defining pattern for Senior-level mathematical counting problems (e.g., *Number of Digit One*).

---

### 32. Explain the DP on Trees pattern.
"**DP on Trees** combines Tree Traversals (specifically Post-order DFS) with Dynamic Programming. I use this when the answer for a parent node depends on the fully calculated answers of its children.

For example, to find the maximum sum of non-adjacent nodes in a tree (*House Robber III*), my DFS function returns two values for every node: the max sum *if* I include this node, and the max sum *if* I exclude it. The root node combines these values from its children to get the final answer."

#### Indepth
Because it's a standard DFS gathering results from the bottom up, it operates in **O(N)** time and **O(1)** or **O(N)** space (for the recursion stack). Recognizing when a tree problem goes beyond standard traversal and requires returning complex state back up the tree is the core of this pattern.

---

### 33. How does DP on Graphs differ from standard graph traversals?
"**DP on Graphs** usually involves finding the longest or shortest path in a graph where traditional algorithms (like Dijkstra) might be too slow or don't fit the constraints. 

It is most commonly used on **Directed Acyclic Graphs (DAGs)**. If the graph is a DAG, I can find a Topological Sort order first. Once ordered, I can use 1D DP to calculate things like the *Longest Increasing Path in a Matrix*. The topological order guarantees I never process a node before its prerequisites are met."

#### Indepth
For DAGs, this pattern runs in **O(V + E)** time, which is optimally fast. Space is **O(V)** for the DP array. When combined with Bitmask DP for problems like the Traveling Salesman, it shifts to **O(2^V * V^2)**.

---

### 34. What is BFS with State (Multi-source / 0-1 BFS)?
"Standard BFS uses a simple Queue to find the shortest path in an unweighted graph. **BFS with State** introduces complexities.
- **Multi-source BFS**: Instead of starting from one point, I add multiple starting nodes to the Queue at time `t=0` (e.g., *Rotting Oranges* spreading from multiple sources simultaneously).
- **0-1 BFS**: When edge weights are strictly 0 or 1, I use a Deque (Double-Ended Queue). If I take a 0-weight edge, I push the node to the *front* of the deque. If it's a 1-weight edge, I push it to the *back*. This guarantees I always process the cheapest paths first without needing a slow Priority Queue."

#### Indepth
Multi-source BFS maintains **O(V + E)** time. 0-1 BFS is a massive optimization because it solves shortest-path problems in **O(V + E)** time instead of Dijkstra's **O(E log V)**. Understanding deque manipulation is crucial here (e.g., *Shortest Path in a Grid with Obstacles Elimination*).

---

### 35. Compare Dijkstra's and Bellman-Ford algorithms for Shortest Path.
"Both find the shortest path in a weighted graph.
**Dijkstra's Algorithm** is a Greedy approach using a Priority Queue. It constantly expands the currently known shortest path. However, it completely fails if the graph has **negative edge weights**.
**Bellman-Ford Algorithm** handles negative weights and can detect negative weight cycles. It works by 'relaxing' all edges $V-1$ times. If I can still relax an edge on the $V$-th pass, I know there is a negative cycle."

#### Indepth
Dijkstra runs in **O(E log V)** time and is the default choice for routing (like Google Maps). Bellman-Ford runs in **O(V * E)** time, making it significantly slower but necessary for financial arbitrage or networks where 'costs' can be negative. Both require **O(V)** space for distance arrays.

---

### 36. What is a Segment Tree used for?
"A **Segment Tree** is a highly specialized binary tree used to handle Range Queries and Range Updates dynamically. 

If I have an array and I repeatedly need the sum (or minimum/maximum) of elements between indices `L` and `R`, a Prefix Sum array works... *unless* the array values are constantly changing. A Segment Tree allows me to both **update** a value and **query** a range in logarithmic time."

#### Indepth
Building the tree takes **O(N)** time and **O(N)** space. Both Point Updates and Range Queries take exactly **O(log N)** time. This is strictly superior to an Array (Update **O(1)**, Query **O(N)**) or Prefix Sum (Update **O(N)**, Query **O(1)**) when both operations are frequent (e.g., *Range Sum Query - Mutable*).

---

### 37. How does a Fenwick Tree (Binary Indexed Tree) compare to a Segment Tree?
"A **Fenwick Tree** (or BIT) solves the exact same problem as a Segment Tree—dynamic range queries and updates (specifically prefix sums)—but it is incredibly space-efficient and shorter to code.

It represents the tree implicitly using an array and relies on clever bitwise operations (`i += i & (-i)`) to navigate between parent and child nodes. It essentially stores overlapping prefix sums based on the binary representation of the indices."

#### Indepth
Like a Segment Tree, Updates and Queries are **O(log N)**, and Space is **O(N)**. However, the constant factors in a Fenwick Tree are much smaller, meaning it runs faster in practice. The trade-off is that it is primarily limited to reversible operations (like addition/subtraction), whereas a Segment Tree can handle Min/Max queries easily.

---

### 38. Explain Mo's Algorithm.
"**Mo's Algorithm** is an advanced technique for processing **offline range queries**. 'Offline' means I am given all the queries up front before I have to output the answers.

If I process queries like `[1, 100]`, then `[90, 95]`, then `[2, 99]`, moving my left/right pointers back and forth is incredibly slow. Mo's Algorithm sorts the queries into 'blocks' or 'buckets' of size $\sqrt{N}$. By strategically reordering them, the pointers move smoothly across the array, minimizing redundant calculations."

#### Indepth
By sorting the queries, it bounds the total pointer movement, achieving a time complexity of **$O((N+Q) \sqrt{N})$**, where Q is the number of queries. It's heavily used in competitive programming and high-end interviews for complex counting over ranges (e.g., *Distinct numbers in range*).

---

### 39. What is the Rolling Hash (Rabin-Karp) pattern?
"**Rolling Hash** is a string matching technique designed to find a substring within a larger string efficiently. 

Instead of comparing 'ABCD' to every 4-letter slice of a document character by character, I calculate a mathematical Hash value for 'ABCD'. As I slide my window across the document, I 'roll' the hash—I subtract the value of the character leaving the window and add the value of the new character entering. This updates the hash in **O(1)** instead of recalculating it."

#### Indepth
This drops the string-matching time complexity from **O(N * M)** to **O(N + M)** in the average case. It requires **O(1)** space. It is the premier algorithm for plagiarism detection or finding highly repeated substrings (e.g., *Longest Duplicate Substring*).

---

### 40. How does the KMP (Knuth-Morris-Pratt) Algorithm work?
"**KMP** is another advanced string matching algorithm. It excels because it never compares the same character twice. 

If I am searching for 'AABAAX' inside a text and I mismatch at 'X', conventional logic starts over at the second character. KMP realizes that the 'AABAA' I just checked has a repeating prefix and suffix ('AA'). It uses a pre-calculated LPS ('Longest Prefix Suffix') array to jump the search forward intelligently, skipping redundant checks."

#### Indepth
Building the LPS array takes **O(M)** time, and the search takes **O(N)**, yielding a strictly **O(N + M)** time complexity without the hash collision risks of Rabin-Karp. Space is **O(M)** for the LPS array. It's historically significant but often replaced by easier algorithms in fast-paced interviews unless explicitly requested (e.g., *Implement strStr()*).
