# 🟢 **1–10: Freshers (Core Data Structures & Linear Algorithms)**

### 1. What is the Array Traversal pattern and when do you use it?
"The **Array Traversal** pattern is the most fundamental way to process data. It involves walking through an array from start to finish, examining every single element without skipping any. 

I use it whenever I need to find something in an unsorted dataset, calculate aggregates (like a sum or average), or modify every element. It's essentially using a `for` or `while` loop to visit every index. For example, if I need to find the maximum value in an array or move all zeros to the end, traversal is the starting point."

#### Indepth
This pattern operates in **O(N)** time complexity because it visits each element exactly once, and **O(1)** space complexity as it only requires a few variables for tracking. In FAANG interviews, almost every array problem requires at least one full traversal to understand the data.

---

### 2. How does the Hashing / Frequency Map pattern optimize lookups?
"The **Hashing** or **Frequency Map** pattern replaces slow linear searches with instant lookups. Instead of checking every item to see if it exists, I store the items (or their counts) in a Hash Map (Dictionary) as I see them.

I use this constantly when a problem asks 'Does this exist?', 'Find duplicates', or 'Count occurrences'. For example, in the classic *Two Sum* problem, instead of a nested loop checking every pair, I store each number's required complement in a map, allowing me to find it instantly later."

#### Indepth
Hashing reduces lookup time to **O(1)** on average, making the overall algorithm **O(N)** time complexity. However, it increases space complexity to **O(N)** because of the hash map storage. This time-space tradeoff is one of the most common optimization techniques expected in interviews.

---

### 3. What is the Two Pointers technique and what is its primary use case?
"The **Two Pointers** technique involves using two integer variables (pointers) to keep track of indices in an array or string. Typically, one starts at the beginning (`left = 0`) and one at the end (`right = n-1`), and they move towards each other based on certain conditions.

I primarily use this on **sorted arrays** or strings to find pairs that meet a condition (like a target sum), or to reverse elements in place. It’s perfect for problems like *Valid Palindrome* or *Container With Most Water*."

#### Indepth
This pattern is incredibly efficient. It usually achieves **O(N)** time complexity while maintaining **O(1)** space complexity since it operates in-place. The logic dictating when to move the `left` or `right` pointer is usually the crux of the interview problem.

---

### 4. Explain the Sliding Window pattern and the difference between fixed and variable windows.
"The **Sliding Window** pattern is a way to track a subset of elements (a 'window') over an array or string. As the window moves forward, instead of recalculating everything inside it, I just subtract the element that leaves the window and add the element that enters.

A **fixed window** size never changes—for example, 'find the max sum of exactly 3 consecutive elements'. A **variable window** expands and shrinks dynamically to meet a condition—for example, 'find the longest substring without repeating characters'."

#### Indepth
This pattern turns nested **O(N^2)** loops into a single **O(N)** traversal. The time complexity is **O(N)**, and space is usually **O(1)** or **O(K)** (where K is the size of the character set or window state). It's the standard solution for contiguous subarray or substring problems.

---

### 5. How does the Prefix Sum pattern work?
"The **Prefix Sum** pattern involves creating a new array where each element at index `i` is the sum of all elements in the original array up to that point (`P[i] = P[i-1] + nums[i]`). 

I use this when an interviewer asks me to repeatedly query the sum of a contiguous subarray. Instead of recalculating the sum every time, I calculate the prefix sum array once. Then, the sum of any subarray from index `i` to `j` is just `P[j] - P[i-1]`."

#### Indepth
Building the prefix sum array takes **O(N)** time and **O(N)** space, but it allows for rapid **O(1)** time queries thereafter. This is crucial for problems like *Range Sum Query* or *Subarray Sum Equals K*, where multiple queries would normally cause a Time Limit Exceeded error.

---

### 6. What is a Monotonic Stack and what problem does it solve?
"A **Monotonic Stack** is a stack data structure that enforces its elements to be strictly increasing or strictly decreasing. If a new element arrives that violates the order, existing elements are popped off the stack until the order is restored.

It beautifully solves the 'Next Greater Element' or 'Next Smaller Element' class of problems. For instance, in a stock price array, if I want to know the next day the price will be higher, a decreasing monotonic stack helps resolve that efficiently."

#### Indepth
By only pushing and popping each element at most once, the time complexity is strictly **O(N)**, with **O(N)** space for the stack. It is the optimal underlying pattern for complex problems like *Daily Temperatures* and *Largest Rectangle in Histogram*.

---

### 7. How would you explain Binary Search?
"**Binary Search** is a famously fast algorithm used to find a target value within a **sorted** array. Instead of checking from start to finish, I check the middle element. If my target is smaller, I discard the right half. If it's larger, I discard the left half. I repeat this until I find the target.

I use it whenever the data is sorted or ordered. It turns a linear search into a logarithmic search."

#### Indepth
The continuous halving gives Binary Search a time complexity of **O(log N)** and a space complexity of **O(1)**. In FAANG interviews, you must be comfortable adapting it to find bounds (first/last occurrence) or dealing with rotated sorted arrays.

---

### 8. What is 'Binary Search on Answer Space'?
"**Binary Search on Answer Space** is an advanced application where I don't search through an array of items. Instead, I search for the *answer* itself. 

If an interviewer asks to 'Minimize the Maximum' or 'Maximize the Minimum' (like allocating minimum pages to students), I define the range of possible answers (e.g., min pages to max pages). I then use Binary Search over this range, paired with a helper function `check(x)` that validates if a specific answer `x` is possible."

#### Indepth
The complexity is **O(N log(Range))** because `check(x)` usually takes **O(N)** time, and the binary search takes **O(log(Range))**. This pattern turns very complex optimization problems (like *Koko Eating Bananas*) into manageable search and validation steps.

---

### 9. Describe Linked List Pointer Manipulation.
"**Linked List Pointer Manipulation** involves using carefully managed reference pointers to traverse, modify, or rearrange nodes in a Linked List. Because you can't access nodes by index like an array, you rely entirely on nodes pointing to their `next` (and sometimes `prev`) neighbors.

I use techniques like the **Fast & Slow Pointers** (Floyd's Cycle detection) to find the middle or detect loops, and **Dummy Nodes** to cleanly handle edge cases at the head of a list."

#### Indepth
Most linked list problems operate in **O(N)** time and strictly **O(1)** space. The difficulty usually lies in the careful ordering of pointer updates. If you update a `node.next` too early, you lose the rest of the list. Reversing a linked list is the quintessential exercise for this pattern.

---

### 10. Explain Tree DFS (Depth-First Search) and its traversal types.
"**Tree DFS** is a way of exploring a tree data structure by going as deep as possible down one branch until reaching a leaf node (a dead end), and then backtracking to explore other branches. 

It is usually implemented with recursion and comes in three main flavors:
- **Preorder** (Visit Node, then Left, then Right) - useful for copying the tree.
- **Inorder** (Left, Visit Node, Right) - on a Binary Search Tree, this visits nodes in sorted order.
- **Postorder** (Left, Right, Visit Node) - useful when calculating properties that depend on children, like max depth."

#### Indepth
DFS visits every node once, resulting in **O(N)** time complexity. The space complexity is **O(H)**, where H is the height of the tree, representing the call stack depth during recursion. It is the backbone for problems like *Lowest Common Ancestor* or *Validating a BST*.
