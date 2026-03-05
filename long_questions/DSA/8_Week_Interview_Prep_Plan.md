# 8-Week Intensive Interview Preparation Plan (Java/Golang focus)

This plan integrates the highest-yield DSA patterns (from the provided `service_based_companies` and `product_based_companies` folders), LeetCode/NeetCode practice, System Design basics, and Language Fundamentals.

**Prerequisites:**
- Dedicate 2-3 hours on weekdays, 4-5 hours on weekends.
- Use **LeetCode** for DSA practice. Focus on the "NeetCode 150" or "Blind 75" lists.
- Stick to one language (Java or Golang) for the entire 8 weeks for algorithmic coding to build muscle memory.

---

## Phase 1: Core DSA Patterns & Language Fundamentals (Weeks 1-3)
*Goal: Build an intuitive understanding of basic data structures and how to manipulate them quickly in your chosen language.*

### Week 1: Arrays, Strings, and Hashing
*   **Day 1-2:** Read through `service_based_companies/01_arrays_strings.md` and `service_based_companies/03_sorting_searching.md`. Ensure you completely understand the time/space complexity logic.
*   **Day 3-5 (LeetCode Practice):** Two Sum, Valid Anagram, Contains Duplicate, Product of Array Except Self, Top K Frequent Elements.
*   **Day 6-7 (Language Deep Dive):** 
    *   **Java:** HashMap internal working, Garbage Collection basics, `String` vs `StringBuilder`.
    *   **Golang:** Slices vs Arrays internals, Maps, string immutability.

### Week 2: Two Pointers, Sliding Window, and Linked Lists
*   **Day 1-2:** Read through `product_based_companies/01_arrays_strings_two_pointers.md` and `service_based_companies/02_linked_lists_math.md`. 
*   **Day 3-5 (LeetCode Practice):** Valid Palindrome, 3Sum, Container With Most Water, Best Time to Buy and Sell Stock, Longest Substring Without Repeating Characters, Reverse Linked List, Linked List Cycle.
*   **Day 6-7 (CS Fundamentals):** Operating Systems basics (Threads vs Processes, Mutexes, Semaphores, Deadlocks).

### Week 3: Stacks, Queues, and Binary Search
*   **Day 1-2:** Read through `product_based_companies/02_linked_lists_stacks_queues.md`. Study monotonic stacks and deque structures.
*   **Day 3-5 (LeetCode Practice):** Valid Parentheses, Evaluate Reverse Polish Notation, Daily Temperatures, Binary Search, Search a 2D Matrix, Koko Eating Bananas.
*   **Day 6-7 (Language Deep Dive - Concurrency):**
    *   **Java:** `Thread`, `Runnable`, `Callable`, `ExecutorService`, `synchronized`, `volatile`, `CountDownLatch`.
    *   **Golang:** Goroutines, Channels (buffered vs unbuffered), `sync.WaitGroup`, `sync.Mutex`, Select statements.

---

## Phase 2: Advanced DSA & High-Level System Design (HLD) (Weeks 4-6)
*Goal: Master tree/graph traversals, understand how distributed systems work, and conquer dynamic programming fears.*

### Week 4: Trees and Tries
*   **Day 1-2:** Read `product_based_companies/03_trees_graphs.md` (Tree sections) and `product_based_companies/05_heaps_tries_greedy.md` (Trie section).
*   **Day 3-5 (LeetCode Practice):** Invert Binary Tree, Maximum Depth of Binary Tree, Diameter of Binary Tree, Lowest Common Ancestor, Validate Binary Search Tree, Implement Trie (Prefix Tree).
*   **Day 6-7 (System Design - Intro):** CAP Theorem, Horizontal vs Vertical Scaling, Load Balancing (Layer 4 vs Layer 7), Consistent Hashing basics.

### Week 5: Graphs and Heaps (Priority Queues)
*   **Day 1-2:** Review Graph sections in `product_based_companies/03_trees_graphs.md`. Understand Kahn's Algorithm (Topological Sort) and difference between BFS/DFS. Read Heap section in `05_heaps_tries_greedy.md`.
*   **Day 3-5 (LeetCode Practice):** Number of Islands, Max Area of Island, Clone Graph, Course Schedule, Kth Largest Element in a Stream, Find Median from Data Stream.
*   **Day 6-7 (System Design - Data & Infrastructure):** Relational (SQL) vs Non-Relational (NoSQL) databases, ACID vs BASE, Caching strategies (Redis/Memcached, Write-Through vs Write-Back).

### Week 6: Dynamic Programming (1D & 2D) and Backtracking
*   **Day 1-2:** Read `product_based_companies/04_dynamic_programming_backtracking.md`. Focus heavily on the recursive tree mapping to memoization, and then to a bottom-up 1D/2D array approach.
*   **Day 3-5 (LeetCode Practice):** Climbing Stairs, House Robber, Coin Change, Longest Increasing Subsequence, Combinations, Subsets, Word Search.
*   **Day 6-7 (System Design - Communication):** REST vs gRPC vs GraphQL, API Gateways, Message Queues (Kafka, RabbitMQ), WebSockets.

---

## Phase 3: Mock Interviews & Specialization (Weeks 7-8)
*Goal: Simulate real interview environments, refine your verbal communication, and plug specific knowledge gaps.*

### Week 7: System Design Deep Dives & Low-Level Design (LLD)
*   **Day 1-2 (LLD / OOP):** Practice designing a Parking Lot, URL Shortener, or a Movie Ticket Booking System. Focus on Class Diagrams, SOLID principles, and Design Patterns (Singleton, Factory, Observer, Strategy).
*   **Day 3-4 (HLD Practice):** Outline the architecture for Design Twitter, Design WhatsApp, or Design Netflix. Focus on the flow of data from Client -> Load Balancer -> Web Server -> App Server -> Cache/DB.
*   **Day 5-7 (Mock Coding Interviews):** Do 3 mock interviews (use Pramp, interviewing.io, or a friend). **Crucial constraint:** Do not use an IDE. Code in a Google Doc or plain text editor and talk out loud while you type.

### Week 8: Revision and Hard Problems
*   **Day 1-2 (Greedy & Advanced Constraints):** Read the Greedy section of `product_based_companies/05_heaps_tries_greedy.md`. Practice Jump Game, Gas Station, Merge Intervals.
*   **Day 3 (Behavioral Prep):** Prepare the STAR methodology (Situation, Task, Action, Result) for questions like: "Tell me about a time you had a conflict with a teammate" or "Describe your most challenging architectural decision."
*   **Day 4-6 (Review):** Re-read all markdown files provided in `service_based_companies/` and `product_based_companies/`. Review the LeetCode problems you struggled with the most. Do 1 LeetCode Hard problem (e.g., Trapping Rain Water, Median of Two Sorted Arrays).
*   **Day 7:** Rest. Don't code the day before your primary interview.

---

## 💡 Daily Tips to guarantee success:
1.  **The 20-Minute Rule:** If you are stuck on a LeetCode problem for more than 20 minutes without writing any pseudo-code, **look at the solution**. Do not waste hours on one problem. Study the optimal solution, understand the pattern, code it yourself, and revisit it in 3 days.
2.  **Talk Out Loud:** Whiteboard interviews test communication as much as coding. Practice explaining your time and space complexity to a rubber duck or your webcam.
3.  **Pattern Over Memorization:** Never memorize code. Memorize the fact that "Top K elements" usually implies a Heap, and "Longest substring" implies a Sliding Window.
