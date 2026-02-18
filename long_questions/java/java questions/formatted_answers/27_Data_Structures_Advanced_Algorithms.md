# 27. Data Structures (Advanced Algorithms & Features)

**Q: Sliding Window Technique**
> "Imagine you have an array and a window of size K.
> Instead of re-calculating the sum of the window every time you move it, you just:
> 1.  **Subtract** the element leaving the window.
> 2.  **Add** the new element entering the window.
>
> This turns an O(N*K) operation into a linear O(N) operation. Crucial for performance."

**Indepth:**
> **Generalization**: This technique isn't just for sums. It works for string problems too (e.g., "Longest Substring Without Repeating Characters"). You extend the window right and contract from the left if a rule is violated.


---

**Q: Two-Pointer Technique**
> "Used for searching pairs in a sorted array (e.g., 'Find two numbers that sum to Target').
> *   Pointer A at the **Start**.
> *   Pointer B at the **End**.
>
> If sum > target, move B left (decrease sum).
> If sum < target, move A right (increase sum).
>
> It solves the problem in one pass (O(N)) instead of nested loops (O(N^2))."

**Indepth:**
> **Three Sum**: The standard "3Sum" problem (find three numbers summing to 0) involves sorting the array first, locking one number, and then running the Two-Pointer technique on the rest.


---

**Q: Prefix Sum Arrays**
> "If you need to calculate the sum of a sub-array (from index L to R) thousands of times, don't loop every time.
>
> Pre-calculate a 'Prefix Sum' array where `P[i]` is the sum of all elements from 0 to `i`.
> Then, `Sum(L, R) = P[R] - P[L-1]`.
> It makes range queries instant (O(1))."

**Indepth:**
> **2D Prefix Sum**: This concept extends to 2D matrices. `Sum(r1, c1, r2, c2)` can also be calculated in O(1) time by pre-calculating a 2D cumulative sum grid.


---

**Q: Java 9/11 String & Collection Features**
> "**String.repeat(n)**: 'abc'.repeat(3) -> 'abcabcabc'.
> **String.isBlank()**: Checks if a string is empty OR just whitespace. Better than `isEmpty()`.
>
> **List.of() / Set.of() / Map.of()**:
> Creates immutable collections in one line.
> `var map = Map.of(\"Key\", \"Val\");`. Clean and safe."

**Indepth:**
> **Var**: Remember `var` is only for local variables. You can't use it for fields or method parameters. It infers the type from the right-hand side `var list = new ArrayList<String>()`.


---

**Q: Collectors: groupingBy & partitioningBy**
> "**groupingBy**: Like SQL GROUP BY. Classification Function -> List of Items.
> `Map<Dept, List<Emp>> byDept = stream.collect(groupingBy(Emp::getDept));`
>
> **partitioningBy**: Special case where the key is just Boolean (True/False).
> `Map<Boolean, List<Emp>> passing = stream.collect(partitioningBy(e -> e.score > 50));`
> Key 'true' has passing students, 'false' has failing ones."

**Indepth:**
> **Cascading**: You can collect recursively. `groupingBy(Dept, groupingBy(City))` creates a nested Map `Map<Dept, Map<City, List<Emp>>>`.


---

**Q: LRU Cache Implementation**
> "You need to store Key-Value pairs, but also know the 'Order of Use'.
> Use a **LinkedHashMap**.
>
> In the constructor, set the 'accessOrder' flag to `true`.
> Override `removeEldestEntry()`. If `size() > capacity`, return true.
> Java will automatically delete the least recently used item for you."

**Indepth:**
> **O(1)**: Why is it O(1)? The LinkedHashMap keeps a doubly-linked list of entries. Moving an entry to the tail involves changing 4 pointers. It doesn't require shifting elements like an ArrayList.


---

**Q: Trie (Prefix Tree)**
> "A tree where edges distinct characters.
> Looking up 'Google' means following the path G -> O -> O -> G -> L -> E.
>
> **Superpower**: Prefix Search.
> 'Find all words starting with PRE'. In a Hash Map, you have to verify every single key. In a Trie, you just walk down P-R-E and return the whole subtree. Extremely fast for autocomplete."

**Indepth:**
> **End Marker**: A Trie node typically needs a boolean flag `isEndOfWord`. Otherwise, you can't distinguish between the word "bat" and "batch" (since "bat" is a prefix of "batch").


---

**Q: BFS vs DFS**
> "**BFS (Breadth First)**: Uses a **Queue**.
> Ripples out layer by layer.
> Good for: Shortest Path in unweighted graphs.
>
> **DFS (Depth First)**: Uses a **Stack** (or Recursion).
> Dives deep into one path, then backtracks.
> Good for: Mazes, topological sort, detecting cycles."

**Indepth:**
> **Memory**: BFS stores the "frontier" of nodes. In a wide graph, the Queue can grow massive, consuming lots of RAM. DFS only stores the current path, so it uses less memory (proportional to height).


---

**Q: ConcurrentHashMap vs Hashtable**
> "**Hashtable** locks the **whole map** for every write. It's a bottleneck.
>
> **ConcurrentHashMap** uses **Lock Stripping** (in Java 7) or **CAS (Compare-And-Swap)** (in Java 8+).
> It allows multiple threads to read and write to different parts ('buckets') of the map simultaneously without blocking each other. It is the gold standard for thread-safe maps."

**Indepth:**
> **Iteration**: `ConcurrentHashMap` iterators are "weakly consistent". They won't throw usage errors, but they may verify elements added *after* the iterator started.


---

**Q: BlockingQueue (Producer-Consumer)**
> "If you are implementing Producer-Consumer, **do not write wait/notify code yourself**. You will get it wrong.
>
> Use `ArrayBlockingQueue` or `LinkedBlockingQueue`.
> *   `put()`: Blocks if queue is Full.
> *   `take()`: Blocks if queue is Empty.
> It handles all the concurrency logic internally."

**Indepth:**
> **Poison Pill**: To shut down a consumer thread gracefully, a common pattern is to put a special "Poison Pill" object into the queue. When the consumer takes it, it knows it's time to exit the loop.


---

**Q: CopyOnWriteArrayList**
> "A thread-safe list where **every write (add/remove) makes a copy of the entire underlying array**.
>
> **Use Case**: Read-Heavy scenarios (like a list of Event Listeners).
> **Readers** never block and see a consistent snapshot.
> **Writers** pay a high cost (copying the array).
> Don't use it if you add/remove elements frequently."

**Indepth:**
> **Snapshots**: Because the iterator works on an array *snapshot*, you can iterate through the list while another thread deletes everything from it, and your iterator will happily finish printing the old data.


---

**Q: WeakHashMap**
> "In a normal HashMap, if you put a Key in, that Key stays in memory forever (or until removed).
>
> In **WeakHashMap**, the Key is held by a 'Weak Reference'.
> If the Key object is **only** referenced by this map (and nowhere else in your app), the Garbage Collector will delete it, and the entry will vanish from the map.
>
> **Use Case**: Caches where you want entries to auto-expire if the application stops using the key."

**Indepth:**
> **Tomcat**: Web stats specifically utilize WeakHashMaps to store session data or classloader references to ensure they don't prevent applications from undeploying.


---

**Q: Cycle Detection in LinkedList**
> "**Floyd's Cycle-Finding Algorithm** (Tortoise and Hare).
> 1.  Slow pointer moves 1 step.
> 2.  Fast pointer moves 2 steps.
>
> If there is a loop, the Fast pointer will eventually 'lap' the Slow pointer and they will meet (`slow == fast`).
> If Fast reaches `null`, there is no cycle."

**Indepth:**
> **Start of Cycle**: To find *where* the cycle begins: Once fast/slow meet, reset Slow to Head. Move both one step at a time. The point where they meet again is the start of the loop.

