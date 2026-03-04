# ⚡ 09 — Data Structures & Algorithms in JavaScript
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Patterns
- Two Pointers / Sliding Window
- Hash Maps for O(1) lookup
- Stack and Queue patterns
- Binary Search variations
- Tree traversal (DFS/BFS)
- Dynamic Programming
- Graph traversal

---

## 📦 Essential JS Data Structures

```javascript
// Stack (LIFO) — use Array
const stack = [];
stack.push(1);        // [1]
stack.push(2);        // [1, 2]
stack.pop();          // returns 2 → [1]
stack[stack.length-1]; // peek — 1

// Queue (FIFO) — use Array (shift is O(n), for heavy use: linked list or deque)
const queue = [];
queue.push(1);        // enqueue
queue.shift();        // dequeue — O(n)!
// O(1) Queue using two stacks or pointer trick
class Queue {
    #data = [];
    #head = 0;
    enqueue(val) { this.#data.push(val); }
    dequeue()    { return this.#head < this.#data.length ? this.#data[this.#head++] : undefined; }
    get size()   { return this.#data.length - this.#head; }
    peek()       { return this.#data[this.#head]; }
}

// Min-Heap (not built-in — need to implement or use sorted structure)
class MinHeap {
    #heap = [];
    push(val) {
        this.#heap.push(val);
        this.#bubbleUp(this.#heap.length - 1);
    }
    pop() {
        if (this.size === 0) return null;
        const min = this.#heap[0];
        const last = this.#heap.pop();
        if (this.#heap.length) { this.#heap[0] = last; this.#siftDown(0); }
        return min;
    }
    peek()       { return this.#heap[0]; }
    get size()   { return this.#heap.length; }

    #bubbleUp(i) {
        while (i > 0) {
            const parent = (i - 1) >> 1;
            if (this.#heap[parent] <= this.#heap[i]) break;
            [this.#heap[parent], this.#heap[i]] = [this.#heap[i], this.#heap[parent]];
            i = parent;
        }
    }
    #siftDown(i) {
        const n = this.#heap.length;
        while (true) {
            let smallest = i;
            const l = 2*i+1, r = 2*i+2;
            if (l < n && this.#heap[l] < this.#heap[smallest]) smallest = l;
            if (r < n && this.#heap[r] < this.#heap[smallest]) smallest = r;
            if (smallest === i) break;
            [this.#heap[smallest], this.#heap[i]] = [this.#heap[i], this.#heap[smallest]];
            i = smallest;
        }
    }
}
```

---

## ❓ Most Asked Coding Questions

### Q1. Two Sum (HashMap — O(n))

```javascript
// LC 1: Given array and target, find indices of two numbers that add to target

function twoSum(nums, target) {
    const map = new Map(); // value → index

    for (let i = 0; i < nums.length; i++) {
        const complement = target - nums[i];
        if (map.has(complement)) {
            return [map.get(complement), i];
        }
        map.set(nums[i], i);
    }
    return [];
}

// twoSum([2, 7, 11, 15], 9) → [0, 1]
// twoSum([3, 2, 4], 6) → [1, 2]
// Time: O(n), Space: O(n)
```

---

### Q2. Valid Parentheses (Stack)

```javascript
// LC 20: Check if bracket string is valid

function isValid(s) {
    const stack = [];
    const pairs = { ')': '(', ']': '[', '}': '{' };

    for (const ch of s) {
        if ('([{'.includes(ch)) {
            stack.push(ch);
        } else {
            if (stack.pop() !== pairs[ch]) return false;
        }
    }
    return stack.length === 0;
}

// isValid("()[]{}") → true
// isValid("(]")     → false
// isValid("{[]}")   → true
// Time: O(n), Space: O(n)
```

---

### Q3. Longest Substring Without Repeating Characters (Sliding Window)

```javascript
// LC 3: Find length of longest substring without repeating characters

function lengthOfLongestSubstring(s) {
    const seen = new Map(); // char → last seen index
    let maxLen = 0;
    let left   = 0;

    for (let right = 0; right < s.length; right++) {
        const ch = s[right];

        // If char was seen AND is within current window: shrink from left
        if (seen.has(ch) && seen.get(ch) >= left) {
            left = seen.get(ch) + 1;
        }

        seen.set(ch, right);
        maxLen = Math.max(maxLen, right - left + 1);
    }

    return maxLen;
}

// "abcabcbb" → 3 ("abc")
// "bbbbb"    → 1 ("b")
// "pwwkew"   → 3 ("wke")
// Time: O(n), Space: O(min(n, alphabet))
```

---

### Q4. Maximum Subarray (Kadane's Algorithm)

```javascript
// LC 53: Find contiguous subarray with largest sum

function maxSubArray(nums) {
    let maxSum    = nums[0];
    let currentSum = nums[0];

    for (let i = 1; i < nums.length; i++) {
        // Either extend existing subarray or start fresh from current element
        currentSum = Math.max(nums[i], currentSum + nums[i]);
        maxSum = Math.max(maxSum, currentSum);
    }

    return maxSum;
}

// [-2,1,-3,4,-1,2,1,-5,4] → 6 (subarray [4,-1,2,1])
// Time: O(n), Space: O(1)
```

---

### Q5. Binary Search

```javascript
// LC 704: Classic binary search + variants

function binarySearch(nums, target) {
    let lo = 0, hi = nums.length - 1;

    while (lo <= hi) {
        const mid = lo + ((hi - lo) >> 1); // avoid overflow
        if (nums[mid] === target) return mid;
        if (nums[mid] < target) lo = mid + 1;
        else                    hi = mid - 1;
    }
    return -1;
}

// Search Insert Position (leftmost index where target goes)
function searchInsert(nums, target) {
    let lo = 0, hi = nums.length;
    while (lo < hi) {
        const mid = (lo + hi) >> 1;
        if (nums[mid] < target) lo = mid + 1;
        else                    hi = mid;
    }
    return lo; // lo = hi = insertion point
}

// Find First and Last Position (LC 34)
function searchRange(nums, target) {
    return [findFirst(nums, target), findLast(nums, target)];
}
function findFirst(nums, target) {
    let lo = 0, hi = nums.length - 1, result = -1;
    while (lo <= hi) {
        const mid = (lo + hi) >> 1;
        if (nums[mid] === target) { result = mid; hi = mid - 1; } // keep searching left
        else if (nums[mid] < target) lo = mid + 1;
        else hi = mid - 1;
    }
    return result;
}
function findLast(nums, target) {
    let lo = 0, hi = nums.length - 1, result = -1;
    while (lo <= hi) {
        const mid = (lo + hi) >> 1;
        if (nums[mid] === target) { result = mid; lo = mid + 1; } // keep searching right
        else if (nums[mid] < target) lo = mid + 1;
        else hi = mid - 1;
    }
    return result;
}
```

---

### Q6. Tree Traversals (DFS & BFS)

```javascript
// Binary Tree node
class TreeNode {
    constructor(val, left = null, right = null) {
        this.val = val; this.left = left; this.right = right;
    }
}

// DFS Traversals — recursive
function inorder(root) {   // left → root → right (sorted for BST)
    if (!root) return [];
    return [...inorder(root.left), root.val, ...inorder(root.right)];
}

function preorder(root) {  // root → left → right
    if (!root) return [];
    return [root.val, ...preorder(root.left), ...preorder(root.right)];
}

function postorder(root) { // left → right → root
    if (!root) return [];
    return [...postorder(root.left), ...postorder(root.right), root.val];
}

// BFS — Level Order Traversal (LC 102)
function levelOrder(root) {
    if (!root) return [];
    const result = [];
    const queue  = [root];

    while (queue.length) {
        const levelSize = queue.length;
        const level = [];

        for (let i = 0; i < levelSize; i++) {
            const node = queue.shift();
            level.push(node.val);
            if (node.left)  queue.push(node.left);
            if (node.right) queue.push(node.right);
        }
        result.push(level);
    }
    return result;
}

// Maximum Depth of Binary Tree (LC 104)
function maxDepth(root) {
    if (!root) return 0;
    return 1 + Math.max(maxDepth(root.left), maxDepth(root.right));
}

// Valid BST (LC 98)
function isValidBST(root, min = -Infinity, max = Infinity) {
    if (!root) return true;
    if (root.val <= min || root.val >= max) return false;
    return isValidBST(root.left, min, root.val) &&
           isValidBST(root.right, root.val, max);
}
```

---

### Q7. Flatten Nested Array / Deep Clone (Recursive)

```javascript
// Flatten nested array to any depth
function flattenDeep(arr) {
    return arr.reduce((acc, val) =>
        Array.isArray(val) ? acc.concat(flattenDeep(val)) : acc.concat(val), []
    );
}
// flattenDeep([1, [2, [3, [4]]]]) → [1, 2, 3, 4]

// Using generator (memory efficient)
function* flatGen(arr) {
    for (const item of arr) {
        if (Array.isArray(item)) yield* flatGen(item);
        else yield item;
    }
}
[...flatGen([1, [2, [3]]])]; // [1, 2, 3]

// Deep clone (handles circular refs with a WeakMap)
function deepClone(val, seen = new WeakMap()) {
    if (val === null || typeof val !== 'object') return val;
    if (val instanceof Date) return new Date(val);
    if (val instanceof RegExp) return new RegExp(val);
    if (seen.has(val)) return seen.get(val); // handle circular refs

    const clone = Array.isArray(val) ? [] : Object.create(Object.getPrototypeOf(val));
    seen.set(val, clone);

    for (const key of Reflect.ownKeys(val)) {
        clone[key] = deepClone(val[key], seen);
    }
    return clone;
}
```

---

### Q8. Top K Frequent Elements (Bucket Sort)

```javascript
// LC 347: Given array, return k most frequent elements

function topKFrequent(nums, k) {
    // Step 1: count frequencies
    const freq = new Map();
    for (const n of nums) freq.set(n, (freq.get(n) ?? 0) + 1);

    // Step 2: bucket sort — index is frequency
    const buckets = Array.from({ length: nums.length + 1 }, () => []);
    for (const [num, count] of freq) buckets[count].push(num);

    // Step 3: collect from highest frequency bucket
    const result = [];
    for (let i = buckets.length - 1; i >= 0 && result.length < k; i--) {
        result.push(...buckets[i]);
    }
    return result.slice(0, k);
}

// [1,1,1,2,2,3], k=2 → [1, 2]
// Time: O(n), Space: O(n) — beats O(n log n) sorting approach
```

---

### Q9. Merge Intervals (LC 56)

```javascript
function merge(intervals) {
    intervals.sort((a, b) => a[0] - b[0]); // sort by start

    const merged = [intervals[0]];

    for (let i = 1; i < intervals.length; i++) {
        const current = intervals[i];
        const last    = merged[merged.length - 1];

        if (current[0] <= last[1]) {
            // Overlapping: extend last interval if needed
            last[1] = Math.max(last[1], current[1]);
        } else {
            merged.push(current);
        }
    }
    return merged;
}

// [[1,3],[2,6],[8,10],[15,18]] → [[1,6],[8,10],[15,18]]
// Time: O(n log n), Space: O(n)
```

---

### Q10. Climbing Stairs / Fibonacci (Dynamic Programming)

```javascript
// LC 70: n stairs, climb 1 or 2 at a time — how many ways?

// Naive (recursive) — O(2^n)
function climbNaive(n) {
    if (n <= 1) return 1;
    return climbNaive(n - 1) + climbNaive(n - 2);
}

// Memoization — O(n) time, O(n) space
function climbMemo(n, memo = {}) {
    if (n <= 1) return 1;
    if (memo[n]) return memo[n];
    return memo[n] = climbMemo(n - 1, memo) + climbMemo(n - 2, memo);
}

// Bottom-up DP — O(n) time, O(1) space
function climbStairs(n) {
    if (n <= 1) return 1;
    let prev = 1, curr = 1;
    for (let i = 2; i <= n; i++) {
        [prev, curr] = [curr, prev + curr];
    }
    return curr;
}

// climbStairs(5) → 8  (1+1+1+1+1, 1+1+1+2, 1+1+2+1, 1+2+1+1, 2+1+1+1, 1+2+2, 2+1+2, 2+2+1)

// General DP template
function dp(n) {
    const cache = new Array(n + 1).fill(0);
    cache[0] = base_case_0;
    cache[1] = base_case_1;
    for (let i = 2; i <= n; i++) {
        cache[i] = recurrence(cache[i-1], cache[i-2]);
    }
    return cache[n];
}
```

---

### Q11. Connected Components / Graph BFS (LC 200 — Number of Islands)

```javascript
// LC 200: Count islands in a 2D grid ('1' = land, '0' = water)

function numIslands(grid) {
    if (!grid.length) return 0;
    const rows = grid.length, cols = grid[0].length;
    let count = 0;

    function bfs(r, c) {
        const queue = [[r, c]];
        grid[r][c] = '0'; // mark visited

        while (queue.length) {
            const [row, col] = queue.shift();
            const directions = [[1,0],[-1,0],[0,1],[0,-1]];

            for (const [dr, dc] of directions) {
                const nr = row + dr, nc = col + dc;
                if (nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] === '1') {
                    grid[nr][nc] = '0';
                    queue.push([nr, nc]);
                }
            }
        }
    }

    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (grid[r][c] === '1') {
                count++;
                bfs(r, c);
            }
        }
    }
    return count;
}
// Time: O(rows × cols), Space: O(rows × cols) for queue

// General graph DFS (adjacency list)
function dfs(graph, node, visited = new Set()) {
    if (visited.has(node)) return;
    visited.add(node);
    console.log(node);
    for (const neighbor of graph[node] ?? []) {
        dfs(graph, neighbor, visited);
    }
}
```

---

### Q12. Implement LRU Cache (LC 146)

```javascript
// Constraints: O(1) get and put — use Map (insertion-ordered) + doubly linked list trick

class LRUCache {
    #capacity;
    #cache = new Map();

    constructor(capacity) {
        this.#capacity = capacity;
    }

    get(key) {
        if (!this.#cache.has(key)) return -1;
        // Move to end (most recently used)
        const val = this.#cache.get(key);
        this.#cache.delete(key);
        this.#cache.set(key, val);
        return val;
    }

    put(key, value) {
        if (this.#cache.has(key)) this.#cache.delete(key);
        else if (this.#cache.size >= this.#capacity) {
            // Delete LRU: first key inserted (Map preserves insertion order)
            this.#cache.delete(this.#cache.keys().next().value);
        }
        this.#cache.set(key, value);
    }
}

// const lru = new LRUCache(2);
// lru.put(1, 1); lru.put(2, 2);
// lru.get(1);    // → 1 (moves 1 to end: [2,1])
// lru.put(3, 3); // capacity full, evict LRU(2): [1,3]
// lru.get(2);    // → -1 (was evicted)
// Key insight: JavaScript Map preserves insertion order → use delete+re-insert to move to end
```

---

### Q13. Quick Tips — Common JS interview gotchas for DSA

```javascript
// ✅ Sorting numbers (sort default is lexicographic!)
[10, 1, 5].sort();              // [1, 10, 5] ❌ wrong!
[10, 1, 5].sort((a, b) => a - b); // [1, 5, 10] ✅

// ✅ Integer division (JS has no integer type)
Math.floor(7 / 2);   // 3
(7 / 2) | 0;         // 3 (bitwise OR — faster)
7 >> 1;              // 3 (right shift — fastest)

// ✅ Infinity for comparison initialization
let maxVal = -Infinity;
let minVal = Infinity;

// ✅ 2D array creation
const grid = Array.from({ length: m }, () => new Array(n).fill(0));
// NOT: Array(m).fill(Array(n).fill(0)) — shares references!

// ✅ Character codes
'a'.charCodeAt(0);     // 97
String.fromCharCode(97); // 'a'
'a'.charCodeAt(0) - 'a'.charCodeAt(0); // 0 (index in alphabet)

// ✅ Reverse a string
"hello".split("").reverse().join(""); // "olleh"

// ✅ Check if number is power of 2
n > 0 && (n & (n - 1)) === 0; // bit trick

// ✅ Modulo for circular index
(index + 1) % length; // next index, wraps around

// ✅ Object.is vs === for edge cases
Object.is(NaN, NaN); // true (safe NaN check)
Object.is(-0, 0);    // false (distinguishes -0 and +0)
NaN === NaN;         // false
-0 === 0;            // true
```
