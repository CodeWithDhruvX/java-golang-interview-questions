# 🟠 **21–30: Advanced (Specific Properties & Advanced Techniques)**

### 21. Explain the Interval Processing pattern.
"**Interval Processing** deals with datasets representing ranges of time or space, like `[start, end]`. It's practically mandatory for any calendar, scheduling, or 'overlapping' type problems. 

The core technique is almost always to **sort** the intervals by their start time. Once sorted, I can iterate through them and easily compare the `current_end` with the `next_start` to see if they overlap (i.e., if `curr_end >= next_start`). If they do, I merge them; if not, I push the current one to my results and move on."

#### Indepth
Because this pattern requires sorting, the time complexity is governed by **O(N log N)**. The space complexity is **O(N)** to hold the sorted intervals or the merged results. Dealing with edge cases (like intervals completely swallowing other intervals) is what makes these problems tricky (e.g., *Merge Intervals* or *Insert Interval*).

---

### 22. What is Topological Sort and what graph problems does it solve?
"**Topological Sort** is a linear ordering of vertices in a Directed Acyclic Graph (DAG) such that for every directed edge `U -> V`, vertex `U` comes before `V` in the ordering. 

I use this whenever a problem implies **dependencies**. For instance, 'You must take Course A before Course B', or 'Build Package X before Package Y'. If there is a cycle (a circular dependency), a topological sort is impossible, making this a great way to detect cycles."

#### Indepth
I usually implement this using **Kahn's Algorithm** (BFS with an In-Degree array). I count how many incoming edges each node has. Nodes with 0 in-degrees are 'ready' and go into a queue. As I process them, I decrement the in-degrees of their neighbors. Time and Space complexity are both **O(V + E)**.

---

### 23. How do you approach Matrix Traversal (2D BFS/DFS)?
"**Matrix Traversal** is applying Graph BFS or DFS to a 2D grid (usually an array of arrays). Instead of explicit edges, a node's 'neighbors' are simply the cells Up, Down, Left, and Right (and occasionally diagonals).

I use this for flood-fill operations, finding the shortest path in a maze, or grouping connected elements. A key implementation detail is always checking bounds (`0 <= r < rows` and `0 <= c < cols`) before visiting a neighbor, and maintaining a `visited` set to prevent infinite loops."

#### Indepth
Traversing the matrix visits each cell at most a constant number of times, making the time complexity **O(N * M)** (Rows * Cols). Space complexity is also **O(N * M)** in the worst case for the BFS Queue or DFS recursion stack (e.g., *Number of Islands*, *Rotting Oranges*).

---

### 24. What are the key properties of a Binary Search Tree (BST) used in interviews?
"A **Binary Search Tree** has a strict property: for any given node, all values in its Left subtree are smaller, and all values in its Right subtree are larger. 

I leverage this heavily for two things:
1.  **Fast Lookups/Insertions**: Moving left or right discards half the remaining tree, mimicking Binary Search.
2.  **Inorder Traversal**: An Inorder traversal (Left, Root, Right) of a valid BST will *always* yield the elements in strictly sorted, ascending order."

#### Indepth
Operations on a balanced BST take **O(log N)** time. However, in the worst case (a skewed tree that looks like a linked list), it degrades to **O(N)**. Most interview problems rely on the Inorder property (e.g., *Validate BST*, *Kth Smallest Element*).

---

### 25. How do you handle String Parsing / Decoding problems?
"**String Parsing** problems often involve deeply nested structures or special formatting, like parsing `3[a2[c]]` into `accaccacc`.

I almost exclusively use a **Stack** for these. I iterate character by character. When I encounter an opening bracket `[`, I push the current string state and numbers to the stack. When I hit a closing bracket `]`, I pop the state, multiply the string, and append it to the previous state. It perfectly handles arbitrary levels of nesting."

#### Indepth
Because each character is processed and pushed/popped a constant number of times, the time complexity is **O(N)** (relative to the length of the output string). Space is **O(N)** for the stack. This is the foundation for writing mini-compilers or calculators (e.g., *Basic Calculator*, *Decode String*).

---

### 26. What does the Data Structure Design pattern entail?
"**Data Structure Design** problems are less about discovering an algorithm and more about combining existing basic structures to achieve specific performance goals. 

The interviewer might say 'Design a cache with O(1) get and O(1) put'. To achieve O(1) lookup, I need a Hash Map. To evict the oldest item in O(1), I need an ordered structure where I can remove from the middle instantly—which means a Doubly Linked List. Combining them solves the *LRU Cache* problem."

#### Indepth
These problems test object-oriented principles, state management, and deep knowledge of data structure internals. The goal is almost always to achieve **O(1)** or **O(log N)** for the requested operations. Examples include *Min Stack* or *Insert Delete GetRandom O(1)*.

---

### 27. Could you explain the Line Sweep pattern?
"**Line Sweep** is a geometric or interval-based pattern. Imagine a vertical line scanning across the X-axis from left to right. Instead of checking every single coordinate point continuously, the line only stops at 'events'—the start or end of an interval or rectangle.

I use this when dealing with multiple overlapping intervals where I need to know the 'skyline' or maximum overlap at any point. I sort the events by their X-coordinate and process them sequentially, maintaining a data structure (like a Heap or active set) representing what the line currently intersects."

#### Indepth
Because it requires sorting the events, the time complexity is bounded by **O(N log N)**. Space is **O(N)** to store the events and the active state. It is the elegant solution to the notoriously difficult *The Skyline Problem* and *Meeting Rooms II*.

---

### 28. How does Event-Based Sweeping relate to Line Sweep?
"**Event-Based Sweeping** is the discrete version of Line Sweep. If I want to know the maximum number of people in a store at any time, I don't look at every minute of the day. 

Instead, I break down 'Person A stayed from 1:00 to 3:00' into two events: `(1:00, +1 person)` and `(3:00, -1 person)`. I dump all arrivals and departures into one array, sort them by time, and keep a running sum. The peak of that running sum is my answer."

#### Indepth
Similar to Line Sweep, it operates in **O(N log N)** time due to sorting, and **O(N)** space. It perfectly handles capacity/overlap problems like *Car Pooling* or *My Calendar II*, effectively turning multi-interval overlap checking into a simple sorted prefix sum.

---

### 29. What is the Meet-in-the-Middle pattern?
"**Meet-in-the-Middle** is a clever divide-and-conquer strategy used when a problem size is slightly too large for a pure brute-force approach, but small enough that dividing it helps. 

If I need to find subsets in an array of size 40, $2^{40}$ is too slow. I divide the array into two halves of size 20. I generate all $2^{20}$ sums for both halves independently. Then, I sort one half and use Binary Search or Two Pointers to find the matching pair between the two halves."

#### Indepth
This pattern dramatically reduces time complexity from **O(2^N)** to **O(2^(N/2) * log(2^(N/2)))** (which simplifies to **O(N * 2^(N/2))**). However, it requires **O(2^(N/2))** space to store the generated subsets. It is crucial for specific subset-sum optimizations (e.g., *Closest Subsequence Sum*).

---

### 30. Explain State Compression DP (Bitmask DP).
"**Bitmask DP** is dynamic programming where the 'state' represents a set of items (like 'visited nodes' or 'used variables'). Instead of using an array `[True, False, True]` as the state key—which is clunky and slow to hash—I compress it into a single integer. 

For 3 items, `[True, False, True]` is binary `101`, which is the decimal number `5`. I can use bitwise operations to check or set states (`mask | (1 << i)`). I use this for optimization problems like the Traveling Salesman Problem where N is small (usually N <= 20)."

#### Indepth
The complexity usually looks like Time **O(2^N * N^2)** and Space **O(2^N)**. The `2^N` factor perfectly represents all possible subsets (From 000...0 to 111...1). It is essential for problems like *Smallest Sufficient Team* and *Partition to K Equal Sum Subsets*.
