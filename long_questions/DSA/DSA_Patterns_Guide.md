
# 🗺️ Structured Learning Roadmap (6-Week Plan)
To master these 50 patterns efficiently, follow this phase-by-phase curriculum. Do not jump to Advanced patterns until you master the Basics.

### 🗓️ Phase 1: The Foundation (Weeks 1-2)
*Goal: Master linear manipulation and basic lookups.*
1.  **Array Traversal** (Pattern 1)
2.  **Hashing / Frequency Map** (Pattern 2)
3.  **Two Pointers** (Pattern 3)
4.  **Sliding Window** (Pattern 4)
5.  **Prefix Sum** (Pattern 5)
6.  **Linked List Pointer Manipulation** (Pattern 9)

### 🗓️ Phase 2: Fundamental Data Structures (Week 3)
*Goal: Understand LIFO, FIFO, and Hierarchies.*
1.  **Monotonic Stack** (Pattern 6)
2.  **String Parsing / Stacks** (Pattern 25)
3.  **Tree DFS (Pre/In/Post)** (Pattern 10)
4.  **Tree BFS (Level Order)** (Pattern 11)
5.  **Binary Search Tree Properties** (Pattern 24)

### 🗓️ Phase 3: Search & Graphs (Week 4)
*Goal: Navigate complex spaces.*
1.  **Binary Search** (Pattern 7)
2.  **Binary Search on Answer Space** (Pattern 8)
3.  **Graph DFS / BFS** (Pattern 12)
4.  **Matrix Traversal** (Pattern 23)
5.  **Topological Sort** (Pattern 22)
6.  **Union-Find** (Pattern 13)

### 🗓️ Phase 4: Optimization & Heaps (Week 5)
*Goal: Find the "Best" or "K-th" element.*
1.  **Heap / Priority Queue** (Pattern 18)
2.  **Greedy** (Pattern 17)
3.  **Backtracking** (Pattern 14)
4.  **Interval Processing** (Pattern 21)
5.  **Trie** (Pattern 20)

### 🗓️ Phase 5: Dynamic Programming (Week 6)
*Goal: Solved complex overlapping subproblems.*
1.  **DP (1D)** (Pattern 15)
2.  **DP (2D)** (Pattern 16)
3.  **DP on Trees** (Pattern 32) (Advanced)
4.  **DP on Graphs** (Pattern 33) (Advanced)

### 🗓️ Phase 6: Specialist & Advanced (Bonus)
*Goal: Crack L5+ / Hard Interviews.*
*   **Bit Manipulation** (Pattern 19)
*   **Segment Tree / Fenwick Tree** (Patterns 36, 37)
*   **Shortest Path** (Pattern 35)
*   **Monotonics & Line Sweep** (Patterns 27, 28)

---
# 📘 The Ultimate DSA Patterns Guide
> *From "Totally New" to "FAANG Aspirant" — A Deep Dive into 42 Key Algorithms.*

---

# ✅ Core & Common DSA Patterns (MUST-KNOW)

## 1. Array Traversal
### 👶 The Concept (ELI5)
Imagine you have a row of mailboxes. To check your mail, you walk from the first mailbox to the last one, checking them one by one. You don't skip any.

### 🧠 Deep Dive
*   **How it works**: Use a loop (for/while) to visit every element in the array at least once.
*   **Complexity**: Time **O(N)**, Space **O(1)**.
*   **Snippet**: `for i in range(len(arr)): print(arr[i])`

### 🏢 FAANG Context
*   **When to use**: When you need to process or search for an element in an unsorted dataset.
*   **Common Problems**: *Find Max Element*, *Move Zeros*, *Remove Duplicates*.

---

## 2. Hashing / Frequency Map
### 👶 The Concept (ELI5)
Imagine a library. Instead of checking every book to find "Harry Potter", you look it up in the catalog (Index Card) which tells you exactly where it is. It's instant.

### 🧠 Deep Dive
*   **How it works**: Use a Hash Map (Dictionary) to store keys and values. This allows **O(1)** lookups.
*   **Complexity**: Time **O(N)** to build, **O(1)** to query. Space **O(N)**.
*   **Snippet**: `count = {}; for x in nums: count[x] += 1`

### 🏢 FAANG Context
*   **When to use**: When you need to check "Does this exist?" or count frequencies instantly.
*   **Common Problems**: *Two Sum*, *Contains Duplicate*, *Group Anagrams*.

---

## 3. Two Pointers
### 👶 The Concept (ELI5)
Imagine reading a scroll. You put one finger at the start and one finger at the end. You move them towards each other to read the sentence from both sides at once.

### 🧠 Deep Dive
*   **How it works**: Initialize two counters, usually `left = 0` and `right = n-1`. Move them based on a condition (e.g., if sum is too small, move left; if too big, move right).
*   **Complexity**: Time **O(N)**, Space **O(1)**.
*   **Snippet**: `while left < right: if check(): left++ else: right--`

### 🏢 FAANG Context
*   **When to use**: Sorted arrays, finding pairs, or reversing segments.
*   **Common Problems**: *Two Sum II*, *3Sum*, *Valid Palindrome*, *Container With Most Water*.

---

## 4. Sliding Window
### 👶 The Concept (ELI5)
Imagine looking through a rectangular window at a long train. You can only see 3 cars at a time. As the train moves, one car leaves your view and a new one enters. You don't look at the whole train at once, just the "window".

### 🧠 Deep Dive
*   **How it works**: Maintain a subset of items within a "window" (range). Instead of recalculating the whole range, you subtract the element leaving and add the element entering.
*   **Complexity**: Time **O(N)**, Space **O(1)** or **O(K)**.
*   **Snippet**: `curr += arr[i] - arr[i-k]`
*   **Types**: **Fixed** (window size is constant) vs **Variable** (window expands/shrinks).

### 🏢 FAANG Context
*   **When to use**: Subarray or Substring problems (e.g., "Find the longest substring...").
*   **Common Problems**: *Longest Substring Without Repeating Characters*, *Minimum Window Substring*, *Max Consecutive Ones*.

---

## 5. Prefix Sum
### 👶 The Concept (ELI5)
Imagine you are saving money.
Day 1: $10 (Total: $10)
Day 2: $20 (Total: $30)
Day 3: $5 (Total: $35)
This running total is the "Prefix Sum". If you want to know how much you saved between Day 2 and 3, you just do `Total(Day 3) - Total(Day 1)`.

### 🧠 Deep Dive
*   **How it works**: Create a new array `P` where `P[i] = P[i-1] + nums[i]`. Allows range sum queries in **O(1)**.
*   **Complexity**: Time **O(N)** to build, **O(1)** per query. Space **O(N)**.
*   **Snippet**: `P[i] = P[i-1] + nums[i]`

### 🏢 FAANG Context
*   **When to use**: When you need to calculate sum of subarrays frequently.
*   **Common Problems**: *Range Sum Query*, *Subarray Sum Equals K*, *Product of Array Except Self*.

---

## 6. Monotonic Stack
### 👶 The Concept (ELI5)
Imagine a line of people ordered by height. When a new person comes, if they are taller than the person in front, the shorter person leaves the line. This ensures the line is always increasing (or decreasing) in height.

### 🧠 Deep Dive
*   **How it works**: Maintain a stack that keeps elements strictly increasing or decreasing. Popping elements helps find the "Next Greater Element" or "Next Smaller Element".
*   **Complexity**: Time **O(N)**, Space **O(N)**.
*   **Snippet**: `while stack and stack[-1] < curr: stack.pop()`

### 🏢 FAANG Context
*   **When to use**: "Next Greater Element", "Previous Smaller Element", or histogram problems.
*   **Common Problems**: *Next Greater Element*, *Daily Temperatures*, *Largest Rectangle in Histogram*.

---

## 7. Binary Search
### 👶 The Concept (ELI5)
Imagine guessing a number between 1 and 100. You guess 50. I say "Too high". You immediately throw away 51-100. Now you guess 25. You cut the problem in half every time.

### 🧠 Deep Dive
*   **How it works**: In a **sorted** range, compare the middle element. If target is lower, search left half. If higher, search right half.
*   **Complexity**: Time **O(log N)**, Space **O(1)**.
*   **Snippet**: `mid = (low + high) // 2`

### 🏢 FAANG Context
*   **When to use**: Search in a **sorted** array or search space.
*   **Common Problems**: *Binary Search*, *Search in Rotated Sorted Array*, *First Bad Version*.

---

## 8. Binary Search on Answer Space
### 👶 The Concept (ELI5)
Instead of searching for a value in a list, you search for the *answer* itself. "Can I manage with 5 trucks?" (No). "Can I manage with 10?" (Yes). "Can I manage with 7?" (Yes). The answer is somewhere between 5 and 7.

### 🧠 Deep Dive
*   **How it works**: Define a range `[min, max]` for the possible answer. Write a `check(x)` function that returns `True/False` if `x` is valid. Binary search over this range using `check(x)`.
*   **Complexity**: Time **O(N log(Range))**.
*   **Snippet**: `if check(mid): ans = mid; high = mid - 1`

### 🏢 FAANG Context
*   **When to use**: "Minimize the Maximum" or "Maximize the Minimum" problems.
*   **Common Problems**: *Koko Eating Bananas*, *Split Array Largest Sum*, *Capacity to Ship Packages*.

---

## 9. Linked List Pointer Manipulation
### 👶 The Concept (ELI5)
Imagine a treasure hunt where each clue leads to the next location. You are at Clue A. It points to Clue B. Clue B points to Clue C. You can't jump to C without visiting B first.

### 🧠 Deep Dive
*   **How it works**: Use pointers (`curr`, `next`, `prev`, `fast`, `slow`) to navigate nodes.
*   **Complexity**: Time **O(N)**, Space **O(1)**.
*   **Techniques**: Fast/Slow pointers (detect cycle), Dummy node (simplify edge cases).

### 🏢 FAANG Context
*   **When to use**: Any problem involving Linked Lists.
*   **Common Problems**: *Reverse Linked List*, *Merge Two Sorted Lists*, *Linked List Cycle*, *Middle of Linked List*.

---

## 10. Tree DFS (Preorder / Inorder / Postorder)
### 👶 The Concept (ELI5)
Imagine exploring a maze. You go down one path as deep as possible until you hit a dead end, then you backtrack and try the next path.
*   **Pre**: Me first, then kids.
*   **In**: Left kid, Me, Right kid.
*   **Post**: Kids first, then Me.

### 🧠 Deep Dive
*   **How it works**: Use recursion or a stack. Visit node, then recursively visit children.
*   **Complexity**: Time **O(N)**, Space **O(H)** (Height of tree).
*   **Snippet**: `def dfs(node): if not node: return; visit(node); dfs(node.left); dfs(node.right)`

### 🏢 FAANG Context
*   **When to use**: Exploring all nodes, validating BST, Path sums.
*   **Common Problems**: *Maximum Depth of Binary Tree*, *Validate BST*, *Path Sum*, *Lowest Common Ancestor*.

---

## 11. Tree BFS (Level Order)
### 👶 The Concept (ELI5)
Imagine pouring water on the top of a pyramid. It wets the top stone first. Then it flows to all stones on the second level. Then all stones on the third level. Level by level.

### 🧠 Deep Dive
*   **How it works**: Use a **Queue**. Add root. While queue is not empty, pop a node, visit it, and add its children to the back of the queue.
*   **Complexity**: Time **O(N)**, Space **O(W)** (Max width of tree).
*   **Snippet**: `queue = deque([root]); while queue: node = queue.popleft()`

### 🏢 FAANG Context
*   **When to use**: Finding shortest path in unweighted graph/tree, printing level by level.
*   **Common Problems**: *Binary Tree Level Order Traversal*, *Right Side View*, *Zigzag Traversal*.

---

## 12. Graph DFS / BFS
### 👶 The Concept (ELI5)
Same as Tree DFS/BFS, but for cities connected by roads.
*   **BFS**: Expanding ripples in a pond (Shortest path).
*   **DFS**: A maze runner going deep (Checking connectivity).
*   **Key**: You must keep track of "Visited" cities so you don't go in circles!

### 🧠 Deep Dive
*   **How it works**:
    *   **DFS**: Stack/Recursion + `visited = set()`.
    *   **BFS**: Queue + `visited = set()`.
*   **Complexity**: Time **O(V + E)**.

### 🏢 FAANG Context
*   **When to use**: Connectivity (Is there a path?), Shortest Path (BFS), Cycle Detection (DFS).
*   **Common Problems**: *Number of Islands*, *Clone Graph*, *Word Ladder*.

---

## 13. Union-Find (Disjoint Set)
### 👶 The Concept (ELI5)
Imagine a party with many small groups of friends.
1.  **Find**: "Which group do you belong to?" (Who is your leader?)
2.  **Union**: "Hey group A, merge with group B!" (Now you share the same leader).

### 🧠 Deep Dive
*   **How it works**: Maintains a parent array where `parent[i]` is the parent of `i`. Two optimizations: **Path Compression** (points directly to ultimate leader) and **Union by Rank** (attach small tree to big tree).
*   **Complexity**: Time **O(α(N))** (Inverse Ackermann function, practically **O(1)**).
*   **Snippet**: `def find(x): if x != p[x]: p[x] = find(p[x]); return p[x]`

### 🏢 FAANG Context
*   **When to use**: Network connectivity, grouping items, Kruskal's Algorithm.
*   **Common Problems**: *Number of Provinces*, *Redundant Connection*, *Accounts Merge*.

---

## 14. Backtracking
### ðŸ‘¶ The Concept (ELI5)
Imagine you are exploring a cave system. You go down a tunnel. If itâ€™s a dead end, you walk back ("backtrack") to the last intersection and try a different tunnel. You systematically try every path until you find the treasure or prove it's not there.

### ðŸ§  Deep Dive
*   **How it works**: Recursion. "Choose" an option, "Explore" recursivley, then "Un-choose" (backtrack) to restore state for the next option.
*   **Complexity**: Exponential Time **O(2^N)** or **O(N!)**.
*   **Snippet**: `def backtrack(path): if goal: res.append(path); return; for choice in options: path.add(choice); backtrack(); path.remove(choice)`

### ðŸ¢ FAANG Context
*   **When to use**: "Find all combinations", "Permutations", "Sudoku Solver".
*   **Common Problems**: *Permutations*, *Subsets*, *Combination Sum*, *N-Queens*.

---

## 15. Dynamic Programming (1D)
### ðŸ‘¶ The Concept (ELI5)
Imagine climbing a staircase. To get to step 10, you must come from step 9 or step 8. If you know how many ways to reach step 9 and step 8, you just add them up. You don't recount from the bottom every time. You remember the past logic.

### ðŸ§  Deep Dive
*   **How it works**: Break problem into overlapping subproblems. Memorize answers in an array `dp`. `dp[i]` depends on `dp[i-1]`, `dp[i-2]`, etc.
*   **Complexity**: Time **O(N)**, Space **O(N)** (can optimize to **O(1)**).
*   **Snippet**: `dp[i] = dp[i-1] + dp[i-2]`

### ðŸ¢ FAANG Context
*   **When to use**: Optimization problems (Min/Max/Count ways) with one changing variable.
*   **Common Problems**: *Climbing Stairs*, *House Robber*, *Coin Change*, *Longest Increasing Subsequence*.

---

## 16. Dynamic Programming (2D)
### ðŸ‘¶ The Concept (ELI5)
Imagine a robot moving on a grid from top-left to bottom-right. The cost to reach any square depends on the cost to reach the square above it and the square to its left. You fill out the grid one by one.

### ðŸ§  Deep Dive
*   **How it works**: Similar to 1D, but state involves two variables (e.g., indices `i` and `j` of two strings, or row/col of a grid).
*   **Complexity**: Time **O(N*M)**, Space **O(N*M)**.
*   **Snippet**: `dp[i][j] = text1[i] == text2[j] ? 1 + dp[i-1][j-1] : max(dp[i-1][j], dp[i][j-1])`

### ðŸ¢ FAANG Context
*   **When to use**: Grid paths, String matching (Edit Distance, LCS), Knapsack type problems.
*   **Common Problems**: *Unique Paths*, *Longest Common Subsequence*, *Edit Distance*.

---

## 17. Greedy
### ðŸ‘¶ The Concept (ELI5)
Imagine you are given a 1-hour break to eat. There are many plates of food. A "Greedy" strategy means you just eat the biggest, most delicious thing *right now* without worrying if it makes you too full for dessert later. You make the locally optimal choice.

### ðŸ§  Deep Dive
*   **How it works**: At each step, make the choice that looks best at the moment. Never look back.
*   **Complexity**: Generally **O(N log N)** (if sorting needed).
*   **Snippet**: `sort(events); count=0; end=-1; for s,e in events: if s >= end: count++; end=e;`

### ðŸ¢ FAANG Context
*   **When to use**: Interval scheduling, resource allocation where local optimal leads to global optimal.
*   **Common Problems**: *Merge Intervals*, *Partition Labels*, *Jump Game*.

---

## 18. Heap / Priority Queue
### ðŸ‘¶ The Concept (ELI5)
Imagine a waiting room where the sickest patient always sees the doctor next, regardless of when they arrived. The "Priority Queue" always keeps the most important item at the front.

### ðŸ§  Deep Dive
*   **How it works**: Binary Heap. `pop()` gives Min (or Max) in **O(log N)**. `push()` adds in **O(log N)**.
*   **Complexity**: Time **O(log N)** per operation.
*   **Snippet**: `heapq.heappush(h, x); val = heapq.heappop(h)`

### ðŸ¢ FAANG Context
*   **When to use**: "Top K items", "Merge K sorted lists", "Median of stream".
*   **Common Problems**: *Kth Largest Element in an Array*, *Merge K Sorted Lists*, *Find Median from Data Stream*.

---

## 19. Bit Manipulation
### ðŸ‘¶ The Concept (ELI5)
Computers speak in 0s and 1s. Bit manipulation is like talking to the computer in its native language. It's extremely fast and used for "on/off" switches.

### ðŸ§  Deep Dive
*   **How it works**: Use AND (`&`), OR (`|`), XOR (`^`), Shift (`<<`, `>>`).
*   **Complexity**: Time **O(1)** (usually number of bits, which is constant 32/64).
*   **Snippet**: `x ^ x = 0`; `x & (x-1)` drops lowest set bit.

### ðŸ¢ FAANG Context
*   **When to use**: Low-level optimizations, unique element finding, subsets.
*   **Common Problems**: *Single Number*, *Number of 1 Bits*, *Counting Bits*.

---

## 20. Trie (Prefix Tree)
### ðŸ‘¶ The Concept (ELI5)
Imagine a dictionary where words are stored as a tree. "Cat", "Car", and "Cap" all share the same starting branch "C" -> "a". You don't store "Ca" three times. You save space and search fast.

### ðŸ§  Deep Dive
*   **How it works**: Tree where keys are characters. Good for prefix lookups.
*   **Complexity**: Time **O(L)** where L is word length.
*   **Snippet**: `node = root; for char in word: node = node.children[char]`

### ðŸ¢ FAANG Context
*   **When to use**: Autocomplete, Spell checker, Prefix matching.
*   **Common Problems**: *Implement Trie*, *Word Search II*, *Design Search Autocomplete System*.

---

## 21. Interval Processing
### ðŸ‘¶ The Concept (ELI5)
Imagine a calendar. You have meetings from 9-10, 10-11, and 8-12. Interval processing helps you figure out: "Am I double booked?" (Overlap) or "Combine adjacent meetings" (Merge).

### ðŸ§  Deep Dive
*   **How it works**: **Sort** by start time. Iterate and compare `current_end` with `next_start`.
*   **Complexity**: Time **O(N log N)** (due to sorting).
*   **Snippet**: `sort(intervals); if curr_end >= next_start: merge()`

### ðŸ¢ FAANG Context
*   **When to use**: Scheduling, calendar merging.
*   **Common Problems**: *Merge Intervals*, *Insert Interval*, *Non-overlapping Intervals*.

---

## 22. Topological Sort
### ðŸ‘¶ The Concept (ELI5)
Imagine getting dressed.
1. Underwear
2. Pants
3. Belt
You can't put on the belt before the pants. Topological sort gives you the correct order to do tasks that depend on each other.

### ðŸ§  Deep Dive
*   **How it works**: Graph algorithm on DAG (Directed Acyclic Graph). Uses Degree array (Kahn's Algo) or DFS.
*   **Complexity**: Time **O(V + E)**.
*   **Snippet**: `queue = [nodes where in_degree == 0]; while queue: process(node); reduce_neighbors_degree()`

### ðŸ¢ FAANG Context
*   **When to use**: Build systems, Course scheduling, Dependency resolution.
*   **Common Problems**: *Course Schedule*, *Course Schedule II*, *Alien Dictionary*.

---

## 23. Matrix Traversal (2D BFS/DFS)
### ðŸ‘¶ The Concept (ELI5)
Imagine a chessboard or a grid map. You want to move a knight from A1 to H8, or flood-fill a color in Paint. You move Up, Down, Left, Right to explore neighbors.

### ðŸ§  Deep Dive
*   **How it works**: Graph traversal on a grid. Nodes are cells `(r, c)`. Neighbors are `(r+1, c), (r-1, c)...`.
*   **Complexity**: Time **O(N*M)**.
*   **Snippet**: `directions = [(0,1), (0,-1), (1,0), (-1,0)]; for dr, dc in directions: nr, nc = r+dr, c+dc`

### ðŸ¢ FAANG Context
*   **When to use**: Shortest path in grid (BFS), Connected areas (DFS).
*   **Common Problems**: *Number of Islands*, *Rotting Oranges*, *Word Search*.

---

## 24. Binary Search Tree (BST) Properties
### ðŸ‘¶ The Concept (ELI5)
A special tree where everything to the left is smaller, and everything to the right is bigger. It allows you to find number "50" by going Left or Right, just like Binary Search but dynamically.

### ðŸ§  Deep Dive
*   **How it works**: Utilizes the property: `Left < Root < Right`. Inorder traversal yields sorted values.
*   **Complexity**: Search/Insert/Delete **O(log N)** (balanced) or **O(N)** (skewed).
*   **Snippet**: `if val < root.val: go_left() else: go_right()`

### ðŸ¢ FAANG Context
*   **When to use**: Maintaining ordered dynamic data.
*   **Common Problems**: *Validate BST*, *Kth Smallest Element in BST*, *LCA of BST*.

---

## 25. String Parsing / Decoding
### ðŸ‘¶ The Concept (ELI5)
Imagine reading a sentence with special codes like "3[a]2[bc]". You need to read character by character, remember the numbers, and expand it to "aaabcbc".

### ðŸ§  Deep Dive
*   **How it works**: Often uses a **Stack** to handle nested brackets/structures. Or simple iteration with state variables.
*   **Complexity**: Time **O(N)**.
*   **Snippet**: `stack = []; for char in s: if char == ']': resolve(stack)`

### ðŸ¢ FAANG Context
*   **When to use**: Calculator, Decoders, Parsing tags.
*   **Common Problems**: *Decode String*, *Basic Calculator*, *Valid Parentheses*.

---

## 26. Data Structure Design
### ðŸ‘¶ The Concept (ELI5)
Imagine building a custom vending machine. You don't just use it; you build the internal gears. "I need a machine that gives me a soda in 0.1 seconds (O(1)) and remembers the last soda I bought."

### ðŸ§  Deep Dive
*   **How it works**: Combine existing structures (HashMap + Doubly Linked List for LRU) to achieve specific time complexities.
*   **Complexity**: Varies. Usually aim for **O(1)**.
*   **Snippet**: `class LRUCache: def __init__(self): self.map = {}; self.list = DList()`

### ðŸ¢ FAANG Context
*   **When to use**: Implementing caches, specialized containers.
*   **Common Problems**: *LRU Cache*, *Min Stack*, *Design Twitter*.

---

# âš¡ Advanced / High-Impact Patterns (GOOD TO KNOW)

## 27. Line Sweep
### ðŸ‘¶ The Concept (ELI5)
Imagine a vertical line moving across a plane from left to right. As it moves, it hits "events" (start of a rectangle, end of a rectangle). You process these events in order instead of checking every coordinate.

### ðŸ§  Deep Dive
*   **How it works**: Sort events by coordinate (usually X-axis). Iterate through sorted events. Maintain active state (active intervals).
*   **Complexity**: Time **O(N log N)** (sorting).
*   **Snippet**: `events.sort(); for x, type in events: process(type)`

### ðŸ¢ FAANG Context
*   **When to use**: "Number of airplanes in the sky", "Skyline problem", "Overlapping intervals".
*   **Common Problems**: *The Skyline Problem*, *Meeting Rooms II*.

---

## 28. Event-Based Sweeping
### ðŸ‘¶ The Concept (ELI5)
Similar to Line Sweep. Instead of checking every second of time, you only wake up when something interesting happens (an "Event").

### ðŸ§  Deep Dive
*   **How it works**: Decompose intervals into `(start, +1)` and `(end, -1)`. Sort all limit points.
*   **Complexity**: Time **O(N log N)**.
*   **Snippet**: `events = [(s, 1), (e, -1)]; sort(events)`

### ðŸ¢ FAANG Context
*   **When to use**: Counting overlaps at specific points.
*   **Common Problems**: *My Calendar II*, *Car Pooling*.

---

## 29. Meet-in-the-Middle
### ðŸ‘¶ The Concept (ELI5)
Imagine cutting a tunnel through a mountain. Itâ€™s faster if two teams start from opposite sides and meet in the middle, rather than one team digging the whole way.

### ðŸ§  Deep Dive
*   **How it works**: Split the search space into two halves. Generate all possibilities for both halves (**O(2^(N/2))**). Then efficiently match them (using Hash Map or Two Pointers).
*   **Complexity**: Reduces **O(2^N)** to **O(2^(N/2))**.
*   **Snippet**: `A = generate(set1); B = generate(set2); find_pairs(A, B)`

### ðŸ¢ FAANG Context
*   **When to use**: Constraints are small (N ~ 40), but not small enough for brute force (2^40 is too big, 2^20 is okay).
*   **Common Problems**: *Partition Array Into Two Arrays to Minimize Sum Difference*, *Closest Subsequence Sum*.

---

## 30. State Compression DP (Bitmask DP)
### ðŸ‘¶ The Concept (ELI5)
Imagine keeping a checklist of 10 chores. Instead of writing "Chore 1 done, Chore 2 not done...", you just write "10"... which in binary is 1010 (Chore 1 and 3 done). You fit the whole checklist into a single integer.

### ðŸ§  Deep Dive
*   **How it works**: Use an integer as a bitmask to represent the set of visited nodes or used items.
*   **Complexity**: Time **O(2^N * N^2)**.
*   **Snippet**: `dp[mask | (1<<i)] = ...`

### ðŸ¢ FAANG Context
*   **When to use**: Traveling Salesman Problem, Assigning N jobs to N people (N <= 20).
*   **Common Problems**: *Smallest Sufficient Team*, *Partition to K Equal Sum Subsets*.

---

## 31. Digit DP
### ðŸ‘¶ The Concept (ELI5)
Imagine counting how many house numbers between 1 and 1000 have the digit '7'. Instead of checking 1, 2, 3... you build the number digit by digit, deciding "Can I put a 7 here?".

### ðŸ§  Deep Dive
*   **How it works**: Construct numbers position by position. State: `(index, tight_constraint, leading_zeros, ...).`
*   **Complexity**: Time **O(Digits * 10 * State)**.
*   **Snippet**: `memo = {}; dfs(idx, is_tight, ...)`

### ðŸ¢ FAANG Context
*   **When to use**: Count numbers in range `[L, R]` satisfying a property (e.g., sum of digits is X).
*   **Common Problems**: *Number of Digit One*, *Non-negative Integers without Consecutive Ones*.

---

## 32. DP on Trees
### ðŸ‘¶ The Concept (ELI5)
Imagine a company hierarchy. The boss wants to know "How much sales did my entire team make?". He asks his direct reports. They ask their reports. The answer bubbles up from the interns to the CEO.

### ðŸ§  Deep Dive
*   **How it works**: Post-order traversal (DFS). Compute answer for children, combine them to get answer for root.
*   **Complexity**: Time **O(N)**.
*   **Snippet**: `left = dfs(node.left); right = dfs(node.right); return max(left, right) + node.val`

### ðŸ¢ FAANG Context
*   **When to use**: Maximum Independent Set on Tree, Tree Diameter, Path Sum.
*   **Common Problems**: *Binary Tree Maximum Path Sum*, *House Robber III*, *Diameter of Binary Tree*.

---

## 33. DP on Graphs
### ðŸ‘¶ The Concept (ELI5)
Finding the shortest path or longest path in a maze that might have cycles or complex connections.

### ðŸ§  Deep Dive
*   **How it works**: Often combined with Topological Sort (for DAGs) or Bitmask (for TSP).
*   **Complexity**: **O(V+E)** for DAG.
*   **Snippet**: `dist[u] = min(dist[u], dist[v] + weight)`

### ðŸ¢ FAANG Context
*   **When to use**: Longest path in DAG, counting paths.
*   **Common Problems**: *Longest Increasing Path in a Matrix*, *Cheapest Flights Within K Stops*.

---

## 34. BFS with State (Multi-source / 0-1 BFS)
### ðŸ‘¶ The Concept (ELI5)
*   **Multi-source**: Drop 5 stones in a pond. Watch the ripples meet.
*   **0-1 BFS**: Some roads are paved (0 cost), some are dirt (1 cost). You want the cheapest route. Use a Deque.

### ðŸ§  Deep Dive
*   **How it works**:
    *   **Multi-source**: Add all sources to Queue at `t=0`.
    *   **0-1 BFS**: Use Deque. Push 0-cost to front, 1-cost to back.
*   **Complexity**: Time **O(V+E)**.
*   **Snippet**: `q.appendleft(v) if weight == 0 else q.append(v)`

### ðŸ¢ FAANG Context
*   **When to use**: Shortest path in 0/1 weight graphs, rotten oranges spreading.
*   **Common Problems**: *01 Matrix*, *Rotting Oranges*, *Shortest Path in a Grid with Obstacles Elimination*.

---

## 35. Shortest Path (Dijkstra / Bellman-Ford)
### ðŸ‘¶ The Concept (ELI5)
Google Maps. Finding the fastest route from A to B considering traffic (weights).
*   **Dijkstra**: No negative roads. Greedy expansion.
*   **Bellman-Ford**: Handles negative roads (time travel?).

### ðŸ§  Deep Dive
*   **How it works**:
    *   **Dijkstra**: Priority Queue `(dist, node)`.
    *   **Bellman**: Relax all edges V-1 times.
*   **Complexity**: Dijkstra **O(E log V)**. Bellman **O(VE)**.
*   **Snippet**: `heapq.heappush(pq, (0, start))`

### ðŸ¢ FAANG Context
*   **When to use**: Weighted graphs.
*   **Common Problems**: *Network Delay Time*, *Path with Maximum Probability*.

---

## 36. Segment Tree
### ðŸ‘¶ The Concept (ELI5)
Imagine a long receipt. You want to know the sum of items 5 to 10 quickly. A Segment Tree pre-calculates the sum of halves, quarters, etc., so you combine just a few numbers instead of adding one by one.

### ðŸ§  Deep Dive
*   **How it works**: Binary tree where each node represents an interval. Root is `[0, N-1]`.
*   **Complexity**: Build **O(N)**, Query/Update **O(log N)**.
*   **Snippet**: `update(pos, val); query(left, right)`

### ðŸ¢ FAANG Context
*   **When to use**: Range Queries with Updates (Sum, Min, Max).
*   **Common Problems**: *Range Sum Query - Mutable*, *Count of Smaller Numbers After Self*.

---

## 37. Fenwick Tree (BIT)
### ðŸ‘¶ The Concept (ELI5)
A magical array that lets you calculate prefix sums and update values very fast. Itâ€™s like a Segment Tree but easier to code (using bitwise tricks).

### ðŸ§  Deep Dive
*   **How it works**: Uses `i & (-i)` to navigate parent/children indices.
*   **Complexity**: Update/Query **O(log N)**. Space **O(N)**.
*   **Snippet**: `while i < n: tree[i] += val; i += i & (-i)`

### ðŸ¢ FAANG Context
*   **When to use**: Prefix sums with updates.
*   **Common Problems**: *Range Sum Query - Mutable*, *Reverse Pairs*.

---

## 38. Mo's Algorithm
### ðŸ‘¶ The Concept (ELI5)
Imagine you have 100 questions about a dataset. Instead of answering them in order (1, 2, 3), you reorder them so you don't have to move your focus back and forth too much.

### ðŸ§  Deep Dive
*   **How it works**: Sort queries by blocks. Move `L` and `R` pointers gradually.
*   **Complexity**: Time **O((N+Q) * sqrt(N))**.
*   **Snippet**: `sort(queries, key=lambda q: (q.l // block, q.r))`

### ðŸ¢ FAANG Context
*   **When to use**: Offline range queries (you know all queries beforehand).
*   **Common Problems**: *Count of Range Sum*, *Distinct numbers in range*.

---

## 39. Rolling Hash (Rabin-Karp)
### ðŸ‘¶ The Concept (ELI5)
Imagine trying to match a DNA sequence. Instead of comparing letter by letter, you turn the sequence into a number (Hash). If the numbers match, the sequence probably matches. When you move to the next position, you update the number quickly instead of recalculating (Rolling).

### ðŸ§  Deep Dive
*   **How it works**: `NewHash = (OldHash - OldChar) * Base + NewChar`.
*   **Complexity**: Time **O(N)** average.
*   **Snippet**: `h = (h * 26 + char) % mod`

### ðŸ¢ FAANG Context
*   **When to use**: String matching, finding duplicate substrings.
*   **Common Problems**: *Repeated DNA Sequences*, *Longest Duplicate Substring*.

---

## 40. KMP Algorithm
### ðŸ‘¶ The Concept (ELI5)
Smart string search. If you match "ABCD" but fail at "E", you know you don't need to start over at "B". You know "B" doesn't match "A". KMP tells you exactly how far to jump back.

### ðŸ§  Deep Dive
*   **How it works**: Build a "Longest Prefix Suffix" (LPS) array. Use it to skip comparisons.
*   **Complexity**: Time **O(N + M)**.
*   **Snippet**: `while j > 0 and s[i] != p[j]: j = lps[j-1]`

### ðŸ¢ FAANG Context
*   **When to use**: Pattern matching in text.
*   **Common Problems**: *Implement strStr()*, *Shortest Palindrome*.

---

## 41. Z-Algorithm
### ðŸ‘¶ The Concept (ELI5)
Similar to KMP but calculates a "Z-array" where `Z[i]` is the length of the substring starting at `i` that matches the prefix of the string.

### ðŸ§  Deep Dive
*   **How it works**: Maintain a "Z-box" (window of match).
*   **Complexity**: Time **O(N)**.
*   **Snippet**: `if i <= R: z[i] = min(R - i + 1, z[i - L])`

### ðŸ¢ FAANG Context
*   **When to use**: String matching, Pattern occurrences.
*   **Common Problems**: *Implement strStr()*, *String Search*.

---

## 42. Randomized Algorithms
### ðŸ‘¶ The Concept (ELI5)
Sometimes guessing is faster than thinking. If you need to pick a leader, just pick one randomly. If you pick enough times, you'll be right.

### ðŸ§  Deep Dive
*   **How it works**: Use `random()` to make decisions (e.g., Pivot in Quicksort, QuickSelect).
*   **Complexity**: Expected **O(N)**, Worst case **O(N^2)** (but very rare).
*   **Snippet**: `pivot = random.choice(nums)`

### ðŸ¢ FAANG Context
*   **When to use**: QuickSort, K-th largest element (QuickSelect).
*   **Common Problems**: *Kth Largest Element in an Array*, *Random Pick with Weight*.

---

# ðŸŽ Bonus: Missing "Hidden Gem" Patterns
> *Patterns not in the original list but critical for top-tier (L5+) interviews.*

## 43. Cyclic Sort
### ðŸ‘¶ The Concept (ELI5)
Imagine a classroom with numbered seats 1 to N. Reviewing the students, you see student #5 sitting in seat #2. You tell him "Go to seat #5!". He goes there, displacing whoever was there. You repeat this until everyone is in their correct seat.

### ðŸ§  Deep Dive
*   **How it works**: Iterate array. If `nums[i]` is not at index `nums[i]-1`, swap it to its correct position. Repeat swap until correct number is at `i`.
*   **Complexity**: Time **O(N)** (each number swapped at most once). Space **O(1)**.
*   **Snippet**: `while i < n: j = nums[i] - 1; if nums[i] != nums[j]: swap(i, j) else: i++`

### ðŸ¢ FAANG Context
*   **When to use**: "Array containing numbers 1 to N", "Find missing/duplicate number in O(1) space".
*   **Common Problems**: *Missing Number*, *Find All Duplicates in an Array*, *First Missing Positive*.

---

## 44. Floyd-Warshall Algorithm
### ðŸ‘¶ The Concept (ELI5)
Imagine you have a map of cities. You want to know the shortest path between **every pair** of cities. You systematically ask: "Can I get from A to B faster if I go through C?". You ask this for every possible city C.

### ðŸ§  Deep Dive
*   **How it works**: 3 Nested Loops. `dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])`.
*   **Complexity**: Time **O(V^3)**.
*   **Snippet**: `for k: for i: for j: d[i][j] = min(d[i][j], d[i][k] + d[k][j])`

### ðŸ¢ FAANG Context
*   **When to use**: All-Pairs Shortest Path, Small Graph (V <= 400), Transitive Closure.
*   **Common Problems**: *Find the City With the Smallest Number of Neighbors at a Threshold Distance*, *Evaluate Division*.

---

## 45. Minimum Spanning Tree (Prim's / Kruskal's)
### ðŸ‘¶ The Concept (ELI5)
You need to connect 10 islands with bridges so everyone can visit everyone else, but you have a limited budget. You want to pick the cheapest set of bridges that connects them all without creating any useless loops.

### ðŸ§  Deep Dive
*   **How it works**:
    *   **Kruskal's**: Sort edges by weight. Add if it doesn't form a cycle (Union-Find).
    *   **Prim's**: Start at one node. Greedily add cheapest edge connecting "visited" set to "unvisited" set (Priority Queue).
*   **Complexity**: **O(E log E)** or **O(E log V)**.
*   **Snippet**: `sort(edges); for u,v,w in edges: if union(u,v): cost += w`

### ðŸ¢ FAANG Context
*   **When to use**: Network design, Clustering, connecting points with min cost.
*   **Common Problems**: *Min Cost to Connect All Points*, *Cheapest Flights*.

---

## 46. Reservoir Sampling
### ðŸ‘¶ The Concept (ELI5)
Imagine a never-ending river of fish. You only have a bucket that holds 1 fish. You want to catch a "random" fish, but you don't know how many fish are in the river. Algorithm: Keep the first fish. For the second fish, flip a coin; if heads, swap it. For the Nth fish, keep it with probability `1/N`.

### ðŸ§  Deep Dive
*   **How it works**: Pick `k` items from a stream of unknown size. For `i-th` item, pick it with probability `k/i`.
*   **Complexity**: Time **O(N)**, Space **O(1)**.
*   **Snippet**: `if random(0, i) < k: reservoir[j] = stream[i]`

### ðŸ¢ FAANG Context
*   **When to use**: "Random element from stream", "Random line from a huge file".
*   **Common Problems**: *Linked List Random Node*, *Random Pick Index*.

---

## 47. Strongly Connected Components (Tarjan's / Kosaraju's)
### ðŸ‘¶ The Concept (ELI5)
In a city with one-way streets, a "Strongly Connected Component" is a neighborhood where you can get from any house to any other house and back. If you leave the neighborhood, you might never be able to return.

### ðŸ§  Deep Dive
*   **How it works**:
    *   **Kosaraju**: DFS Order -> Transpose Graph -> DFS again.
    *   **Tarjan**: Single DFS using `discovery_time` and `low_link` values.
*   **Complexity**: Time **O(V+E)**.
*   **Snippet**: `if low[u] == disc[u]: pop_stack_until_u()`

### ðŸ¢ FAANG Context
*   **When to use**: Critical connections (Bridges), Network analysis.
*   **Common Problems**: *Critical Connections in a Network*, *Course Schedule IV*.

---

## 48. Morris Traversal
### ðŸ‘¶ The Concept (ELI5)
Traversing a tree usually requires a map (Stack/Recursion) to remember where you came from. Morris Traversal is like leaving a temporary thread trail (modifying the tree pointers) so you can find your way back without a map.

### ðŸ§  Deep Dive
*   **How it works**: Create a temporary link from the "right-most node of the left child" back to the current node. allows traversing in **O(N)** time with **O(1)** space.
*   **Complexity**: Time **O(N)**, Space **O(1)**.
*   **Snippet**: `pre = cur.left; while pre.right: pre = pre.right; pre.right = cur`

### ðŸ¢ FAANG Context
*   **When to use**: O(1) Space limitation for Tree Traversal.
*   **Common Problems**: *Binary Tree Inorder Traversal*, *Recover Binary Search Tree*.

---

## 49. Non-Comparison Sorting (Bucket / Counting / Radix)
### ðŸ‘¶ The Concept (ELI5)
Imagine sorting a pile of mail by Zip Code. You don't compare every letter to every other letter. You just toss them into bins: "100xx" bin, "200xx" bin. Then you just collect the piles. Itâ€™s strictly faster than comparing.

### ðŸ§  Deep Dive
*   **How it works**: Use the *value* of the element to determine its position.
    *   **Counting**: Count frequency of each value.
    *   **Bucket**: Scatter to buckets â†’ Sort buckets â†’ Gather.
*   **Complexity**: Time **O(N)**. Space **O(K)** (range of values).
*   **Snippet**: `buckets[val / range].append(val); for b in buckets: sort(b)`

### ðŸ¢ FAANG Context
*   **When to use**: "Sort an array in O(N)", "Maximum Gap", "Top K Frequent".
*   **Common Problems**: *Sort Colors*, *Top K Frequent Elements*, *Maximum Gap*.

---

## 50. Divide and Conquer
### ðŸ‘¶ The Concept (ELI5)
"Divide and Rule". How do you count the number of people in a stadium? Split the stadium in half. Ask two friends to count each half. They split their halves too. Eventually, someone just counts one person. Then you add up the answers.

### ðŸ§  Deep Dive
*   **How it works**: Recursively break problem into non-overlapping subproblems.
    1.  **Divide**: Split into subproblems.
    2.  **Conquer**: Solve subproblems (often recursively).
    3.  **Combine**: Merge solutions.
*   **Complexity**: Defined by Master Theorem (e.g., **O(N log N)** for Merge Sort).
*   **Snippet**: `mid = n // 2; left = solve(0, mid); right = solve(mid, n); return merge(left, right)`

### ðŸ¢ FAANG Context
*   **When to use**: "Count Inversions", "Merge K Sorted Lists", "Majority Element" (D&C variant).
*   **Common Problems**: *Sort an Array (Merge Sort)*, *Count of Range Sum*, *Beautiful Array*.
