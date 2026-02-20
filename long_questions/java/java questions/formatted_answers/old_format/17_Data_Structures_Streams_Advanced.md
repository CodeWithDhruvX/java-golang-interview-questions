# 17. Data Structures (Streams & Advanced)

**Q: BST vs Heap**
> "Both are trees, but they have different rules.
>
> **Binary Search Tree (BST)** is ordered. Everything to the left of a node is smaller; everything to the right is larger. It's built for **Searching** (O(log n)).
>
> **Heap (Min/Max)** only guarantees that the parent is smaller (or larger) than its children. It doesn't care about left vs. right. It's built for **Fast Access to the Extremes** (finding the min or max is O(1)). You rarely search in a heap; you just grab the top element."

**Indepth:**
> **Self-Balancing**: A standard BST can degenerate into a linked list (O(n)) if you insert sorted data (1, 2, 3, 4). Real-world implementations use **Red-Black Memory** or **AVL Trees** to keep the tree balanced (O(log n)).


---

**Q: Map vs FlatMap**
> "Think of **Map** as a 1-to-1 transformation. You have a list of `Person` objects, and you want a list of their names. One person in, one name out.
> `Stream<Person> -> map() -> Stream<String>`
>
> **FlatMap** is a 1-to-Many transformation that also 'flattens' the result. If you have a list of `Writer` objects, and each writer has a list of `Books`, using `map()` would give you a `Stream<List<Book>>` (a stream of lists). That's messy.
> `flatMap()` takes those inner lists and pours them all out into a single, continuous `Stream<Book>`."

**Indepth:**
> **Nulls**: `flatMap` effectively filters out empty results. If a function returns an empty stream, it adds nothing to the outcome. This is safer than mapping to null.


---

**Q: Reduce() method**
> "**Reduce** takes a stream of elements and combines them into a single result.
>
> It needs two things:
> 1.  **Identity**: The starting value (e.g., `0` for specific sum, or `""` for string concat).
> 2.  **Accumulator**: A function that takes the 'running total' and the 'next element' and combines them.
>
> Example: `numbers.stream().reduce(0, (a, b) -> a + b)` adds everything up.
> Use it when you want to boil a whole list down to one number or object."

**Indepth:**
> **Parallelism**: `reduce` is designed for parallelism. If the operation is associative `(a+b)+c == a+(b+c)`, the stream can split, reduce chunks in parallel, and combine the results.


---

**Q: Parallel Stream vs Sequential Stream**
> "**Sequential Stream** runs on a single thread. It processes items one by one. It's safe, predictable, and usually fast enough.
>
> **Parallel Stream** splits the data into chunks and processes them on multiple threads (using the Fork/Join pool).
> *   **Pro**: Can be much faster for massive datasets or CPU-intensive tasks.
> *   **Con**: It has overhead (managing threads). If your task is small, parallel is actually *slower*. Also, if your operations aren't thread-safe, you'll get random bugs."

**Indepth:**
> **Common Pool**: Note that *all* parallel streams in the JVM share the same `ForkJoinPool.commonPool()`. If one task blocks (e.g., checks a slow website), it can starve every other parallel stream in your application.


---

**Q: Sliding Window Technique**
> "This isn't a Java class; it's an algorithm pattern.
>
> Imagine you need to find the maximum sum of any 3 consecutive limits in an array.
>
> *   **Naive way**: Loop through every element, and for each one, look at the next 2. That's O(n*k).
> *   **Sliding Window**: You create a 'window' of size 3. Calculated the sum. Then, slide the window one step right.
> instead of re-calculating the whole sum, you just **subtract** the element leaving on the left and **add** the element entering on the right. That makes it O(n)."

**Indepth:**
> **Range**: Sliding Window is essentially optimizing a nested loop by reusing the previous computation. It turns O(N*K) into O(N).


---

**Q: Two-Pointer Technique**
> "This is used for searching in sorted arrays or strings.
>
> Example: Find two numbers in a sorted array that add up to a target.
> Instead of nested loops (O(n^2)), you put one pointer at the **Start** and one at the **End**.
> *   If `sum > target`, move the End pointer left (to get a smaller sum).
> *   If `sum < target`, move the Start pointer right (to get a larger sum).
>
> It reduces the complexity to O(n)."

**Indepth:**
> **Requirement**: This technique almost exclusively relies on the input being **Sorted**. If the array is unsorted, you can't decide which pointer to move, and the logic falls apart.


---

**Q: Java 9+ Collection Factory Methods**
> "Before Java 9, creating a small immutable list was verbose: `Arrays.asList()` or `Collections.unmodifiableList()`.
>
> Now we have:
> *   `List.of("A", "B")`
> *   `Set.of("A", "B")`
> *   `Map.of("Key1", "Val1", "Key2", "Val2")`
>
> These return heavily optimized, **immutable** collections. You can't add nulls, and you can't resize them. They are perfect for configuration and constants."

**Indepth:**
> **Dupes**: `Set.of()` will throw an `IllegalArgumentException` if you pass duplicate elements (`Set.of("A", "A")`). It validates uniqueness at creation time.


---

**Q: Collectors.groupingBy()**
> "This is arguably the most useful method in the Stream API. It works like the `GROUP BY` clause in SQL.
>
> If you have a list of `Employee` objects and you want to group them by Department:
> `employees.stream().collect(Collectors.groupingBy(Employee::getDepartment));`
>
> The result is a `Map<Department, List<Employee>>`. You can even cascade collectors to count them:
> `groupingBy(Employee::getDepartment, Collectors.counting())`."

**Indepth:**
> **Downstream**: The second argument to `groupingBy` is a "downstream collector". You can group, and *then* map, reduce, or count the values in each group. `groupingBy(City, mapping(Person::getName, toList()))`.


---

**Q: LRU Cache Implementation**
> "An **LRU (Least Recently Used) Cache** throws away the oldest used items when it gets full.
>
> To implement this efficiently (O(1) access and removal), you need two data structures working together:
> 1.  **HashMap**: For instant lookups (Key -> Node).
> 2.  **Doubly Linked List**: To maintain the order.
>
> **The Trick**:
> *   When you access an item, you move it to the *Head* of the list (mark as recently used).
> *   When you add an item and the cache is full, you simply remove the *Tail* of the list (least recently used) and remove it from the Map.
>
> In Java, `LinkedHashMap` actually has this logic built-in if you override the `removeEldestEntry` method."

**Indepth:**
> **Access Order**: `LinkedHashMap` has a special constructor `(capacity, loadFactor, accessOrder)`. If `accessOrder` is true, iterating the map visits the most recently accessed elements last.


---

**Q: Trie (Prefix Tree)**
> "A **Trie** is a special tree used for storing strings, like a dictionary for autocomplete.
>
> *   The root is empty.
> *   Each node represents a character.
> *   To store 'CAT', you go Root -> C -> A -> T.
>
> **Why use it?**
> If you have a million words, checking if a word starts with 'pre' is super slow in a List. In a Trie, it takes just 3 tiny steps (P -> R -> E). It's incredibly fast for prefix-based searches."

**Indepth:**
> **Memory**: A Trie can actually *save* memory if many strings share common prefixes ("internet", "interest", "international"). The node "inter" is stored only once.


---

**Q: BFS vs DFS**
> "These are the two ways to walk through a Graph or Tree.
>
> **BFS (Breadth-First Search)**: Explores layer by layer. It visits all neighbors before going deeper.
> *   **Data Structure**: Uses a **Queue**.
> *   **Use Case**: Finding the *shortest path* in an unweighted graph.
>
> **DFS (Depth-First Search)**: Goes as deep as possible down one path before backtracking.
> *   **Data Structure**: Uses a **Stack** (or Recursion).
> *   **Use Case**: Solving mazes, checking for cycles, or pathfinding where you want *any* solution, not necessarily the shortest."

**Indepth:**
> **Recursion Risk**: DFS using recursion can crash with `StackOverflowError` if the graph is too deep. For deep graphs, use an explicit `Stack` object instead.


---

**Q: ConcurrentHashMap vs Hashtable**
> "**Hashtable** is the dinosaur. It is thread-safe, but it locks the **entire map** for every operation. If one thread is reading, no one else can write. It's a major bottleneck.
>
> **ConcurrentHashMap** is the modern replacement. It locks **segments** (buckets), not the whole map.
> Two threads can safely write to different buckets at the exact same time without waiting for each other. Reads are generally lock-free. It is much, much faster."

**Indepth:**
> **Nulls**: Neither Hashtable nor ConcurrentHashMap allow `null` keys or values. HashMap *does* allow one null key. This is a historical quirk.

