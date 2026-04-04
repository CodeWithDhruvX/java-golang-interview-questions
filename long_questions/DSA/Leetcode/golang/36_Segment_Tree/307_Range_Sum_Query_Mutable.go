package main

import "fmt"

// 307. Range Sum Query - Mutable - Segment Tree
// Time: O(log N) for update and query, Space: O(N)
type NumArray struct {
	tree []int
	n    int
}

// Constructor builds segment tree from array
func ConstructorNumArray(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, 2*n)
	
	// Build leaf nodes
	for i := 0; i < n; i++ {
		tree[i+n] = nums[i]
	}
	
	// Build internal nodes
	for i := n - 1; i > 0; i-- {
		tree[i] = tree[2*i] + tree[2*i+1]
	}
	
	return NumArray{tree: tree, n: n}
}

// Update value at index i
func (this *NumArray) Update(i int, val int) {
	// Update leaf node
	pos := i + this.n
	this.tree[pos] = val
	
	// Update internal nodes
	pos /= 2
	for pos > 0 {
		this.tree[pos] = this.tree[2*pos] + this.tree[2*pos+1]
		pos /= 2
	}
}

// Sum range query
func (this *NumArray) SumRange(left int, right int) int {
	// Convert to leaf positions
	left += this.n
	right += this.n
	
	sum := 0
	
	// Query from both ends
	for left <= right {
		if left%2 == 1 {
			sum += this.tree[left]
			left++
		}
		if right%2 == 0 {
			sum += this.tree[right]
			right--
		}
		left /= 2
		right /= 2
	}
	
	return sum
}

// Alternative implementation with recursive segment tree
type NumArrayRecursive struct {
	tree []int
	n    int
}

func ConstructorNumArrayRecursive(nums []int) NumArrayRecursive {
	n := len(nums)
	tree := make([]int, 4*n) // 4*n is safe upper bound
	
	buildTree(nums, tree, 0, 0, n-1)
	
	return NumArrayRecursive{tree: tree, n: n}
}

func buildTree(nums, tree []int, node, start, end int) {
	if start == end {
		tree[node] = nums[start]
		return
	}
	
	mid := start + (end-start)/2
	buildTree(nums, tree, 2*node+1, start, mid)
	buildTree(nums, tree, 2*node+2, mid+1, end)
	
	tree[node] = tree[2*node+1] + tree[2*node+2]
}

func (this *NumArrayRecursive) Update(i int, val int) {
	updateTree(this.tree, 0, 0, this.n-1, i, val)
}

func updateTree(tree []int, node, start, end, idx, val int) {
	if start == end {
		tree[node] = val
		return
	}
	
	mid := start + (end-start)/2
	if idx <= mid {
		updateTree(tree, 2*node+1, start, mid, idx, val)
	} else {
		updateTree(tree, 2*node+2, mid+1, end, idx, val)
	}
	
	tree[node] = tree[2*node+1] + tree[2*node+2]
}

func (this *NumArrayRecursive) SumRange(left int, right int) int {
	return queryTree(this.tree, 0, 0, this.n-1, left, right)
}

func queryTree(tree []int, node, start, end, left, right int) int {
	// No overlap
	if start > right || end < left {
		return 0
	}
	
	// Complete overlap
	if left <= start && end <= right {
		return tree[node]
	}
	
	// Partial overlap
	mid := start + (end-start)/2
	return queryTree(tree, 2*node+1, start, mid, left, right) +
		queryTree(tree, 2*node+2, mid+1, end, left, right)
}

// Lazy propagation segment tree for range updates
type NumArrayLazy struct {
	tree []int
	lazy []int
	n    int
}

func ConstructorNumArrayLazy(nums []int) NumArrayLazy {
	n := len(nums)
	tree := make([]int, 4*n)
	lazy := make([]int, 4*n)
	
	buildTreeLazy(nums, tree, lazy, 0, 0, n-1)
	
	return NumArrayLazy{tree: tree, lazy: lazy, n: n}
}

func buildTreeLazy(nums, tree, lazy []int, node, start, end int) {
	if start == end {
		tree[node] = nums[start]
		return
	}
	
	mid := start + (end-start)/2
	buildTreeLazy(nums, tree, lazy, 2*node+1, start, mid)
	buildTreeLazy(nums, tree, lazy, 2*node+2, mid+1, end)
	
	tree[node] = tree[2*node+1] + tree[2*node+2]
}

func (this *NumArrayLazy) UpdateRange(left, right, val int) {
	updateRangeLazy(this.tree, this.lazy, 0, 0, this.n-1, left, right, val)
}

func updateRangeLazy(tree, lazy []int, node, start, end, left, right, val int) {
	// Apply pending lazy updates
	if lazy[node] != 0 {
		tree[node] += (end - start + 1) * lazy[node]
		if start != end {
			lazy[2*node+1] += lazy[node]
			lazy[2*node+2] += lazy[node]
		}
		lazy[node] = 0
	}
	
	// No overlap
	if start > right || end < left {
		return
	}
	
	// Complete overlap
	if left <= start && end <= right {
		tree[node] += (end - start + 1) * val
		if start != end {
			lazy[2*node+1] += val
			lazy[2*node+2] += val
		}
		return
	}
	
	// Partial overlap
	mid := start + (end-start)/2
	updateRangeLazy(tree, lazy, 2*node+1, start, mid, left, right, val)
	updateRangeLazy(tree, lazy, 2*node+2, mid+1, end, left, right, val)
	
	tree[node] = tree[2*node+1] + tree[2*node+2]
}

func (this *NumArrayLazy) SumRange(left int, right int) int {
	return queryTreeLazy(this.tree, this.lazy, 0, 0, this.n-1, left, right)
}

func queryTreeLazy(tree, lazy []int, node, start, end, left, right int) int {
	// Apply pending lazy updates
	if lazy[node] != 0 {
		tree[node] += (end - start + 1) * lazy[node]
		if start != end {
			lazy[2*node+1] += lazy[node]
			lazy[2*node+2] += lazy[node]
		}
		lazy[node] = 0
	}
	
	// No overlap
	if start > right || end < left {
		return 0
	}
	
	// Complete overlap
	if left <= start && end <= right {
		return tree[node]
	}
	
	// Partial overlap
	mid := start + (end-start)/2
	return queryTreeLazy(tree, lazy, 2*node+1, start, mid, left, right) +
		queryTreeLazy(tree, lazy, 2*node+2, mid+1, end, left, right)
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Segment Tree for Range Sum Queries
- **Iterative Implementation**: Efficient array-based segment tree
- **Recursive Implementation**: Traditional tree-based approach
- **Lazy Propagation**: Optimize range updates with deferred updates
- **Point Updates**: Update individual elements efficiently

## 2. PROBLEM CHARACTERISTICS
- **Mutable Array**: Array elements can be updated
- **Range Sum Queries**: Query sum of subarray ranges
- **Dynamic Operations**: Support both updates and queries
- **Efficient Operations**: Both operations in O(log N) time

## 3. SIMILAR PROBLEMS
- Range Module (LeetCode 715) - Range tracking with segment tree
- The Skyline Problem (LeetCode 218) - Range maximum queries
- Falling Squares (LeetCode 699) - Range height updates
- Count of Smaller Numbers After Self (LeetCode 315) - Segment tree variant

## 4. KEY OBSERVATIONS
- **Binary Tree Structure**: Segment tree is complete binary tree
- **Range Decomposition**: Any range can be decomposed into O(log N) nodes
- **Update Propagation**: Point updates affect O(log N) nodes
- **Lazy Updates**: Range updates can be deferred for efficiency

## 5. VARIATIONS & EXTENSIONS
- **Iterative vs Recursive**: Array-based vs pointer-based implementation
- **Lazy Propagation**: Efficient range updates
- **Different Operations**: Min, max, gcd, custom operations
- **Persistent Segment Tree**: Historical versions of array

## 6. INTERVIEW INSIGHTS
- Always clarify: "Array size constraints? Update/query frequency? Range sizes?"
- Edge cases: single element array, full range queries, invalid indices
- Time complexity: O(log N) for both update and query
- Space complexity: O(N) for iterative, O(4N) for recursive
- Key insight: decompose any range into O(log N) tree nodes

## 7. COMMON MISTAKES
- Wrong indexing in iterative implementation
- Not handling edge cases in range queries
- Missing lazy propagation for range updates
- Incorrect tree building order
- Off-by-one errors in range boundaries

## 8. OPTIMIZATION STRATEGIES
- **Iterative Tree**: O(N) space, O(log N) operations - most efficient
- **Recursive Tree**: O(4N) space, O(log N) operations - intuitive
- **Lazy Propagation**: O(N) space, O(log N) operations - range updates
- **Fenwick Tree**: O(N) space, O(log N) operations - for sum only

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like a hierarchical accounting system:**
- The array is like individual account balances
- Each manager oversees a group of accounts
- Regional managers oversee groups of managers
- The CEO oversees all regional managers
- Updates flow up the hierarchy, queries flow down
- Like organizational structure for efficient reporting

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Initial array, update operations (index, value), sum queries (range)
2. **Goal**: Support efficient updates and range sum queries
3. **Rules**: Array is mutable, operations are dynamic
4. **Output**: Range sum results after updates

#### Phase 2: Key Insight Recognition
- **"Hierarchical grouping natural"** → Group array elements into segments
- **"Binary decomposition"** → Each level doubles the segment size
- **"Update propagation"** → Changes affect parent segments
- **"Range decomposition"** → Any query range covers O(log N) segments

#### Phase 3: Strategy Development
```
Human thought process:
"I need fast updates and range sums.
Naive approach: O(N) for updates, O(N) for queries.

Segment Tree Approach:
1. Build binary tree where each node stores segment sum
2. Leaf nodes store individual array elements
3. Internal nodes store sum of children
4. Update: modify leaf, propagate to parents (O(log N))
5. Query: decompose range into tree nodes (O(log N))

This gives O(log N) for both operations!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Handle gracefully
- **Single element**: Tree with just root
- **Invalid ranges**: Validate input indices
- **Large arrays**: Ensure space efficiency

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [1, 3, 5], query SumRange(0, 2)

Human thinking:
"Iterative Segment Tree Approach:
Step 1: Build tree (size = 2*n = 6)
Tree: [_, _, _, 1, 3, 5] (leaves at positions 3,4,5)
Build internal nodes:
- tree[2] = tree[4] + tree[5] = 3 + 5 = 8
- tree[1] = tree[2] + tree[3] = 8 + 1 = 9
Final tree: [9, 8, _, 1, 3, 5]

Step 2: Query SumRange(0, 2)
Convert to leaf positions: left=3, right=5
- left=3 (odd): include tree[3]=1, left=4
- right=5 (odd): include tree[5]=5, right=4
- left=4, right=4: left <= right, left/2=2, right/2=2
- left=2 (even), right=2 (even): no inclusion
- left=1, right=1: stop

Sum = tree[3] + tree[5] = 1 + 5 = 6 ✓"
```

#### Phase 6: Intuition Validation
- **Why tree structure**: Hierarchical grouping enables efficient operations
- **Why O(log N)**: Each level halves the problem size
- **Why iterative efficient**: Array-based avoids recursion overhead
- **Why lazy propagation**: Defer updates until needed

### Common Human Pitfalls & How to Avoid Them
1. **"Why not simple array?"** → O(N) updates vs O(log N) with tree
2. **"Should I use recursive?"** → Iterative is more efficient, recursive is intuitive
3. **"What about lazy propagation?"** → Essential for range updates
4. **"Can I use Fenwick tree?"** → Yes, but only for sum operations
5. **"Why 4N space?"** → Safe upper bound for recursive tree

### Real-World Analogy
**Like a company's financial reporting system:**
- Individual stores report daily sales (leaf nodes)
- District managers sum store sales (internal nodes)
- Regional managers sum district sales (higher nodes)
- CEO gets total company revenue (root node)
- Sales updates flow up the hierarchy
- Revenue queries can be answered at any level
- Like organizational structure for efficient data aggregation

### Human-Readable Pseudocode
```
class SegmentTree:
    constructor(nums):
        n = length(nums)
        tree = array of size 2*n
        
        # Build leaf nodes
        for i from 0 to n-1:
            tree[n+i] = nums[i]
        
        # Build internal nodes
        for i from n-1 down to 1:
            tree[i] = tree[2*i] + tree[2*i+1]
    
    update(index, value):
        pos = index + n
        tree[pos] = value
        
        # Update parents
        while pos > 1:
            pos = pos // 2
            tree[pos] = tree[2*pos] + tree[2*pos+1]
    
    query(left, right):
        left = left + n
        right = right + n
        sum = 0
        
        while left <= right:
            if left is odd:
                sum += tree[left]
                left += 1
            if right is even:
                sum += tree[right]
                right -= 1
            left = left // 2
            right = right // 2
        
        return sum
```

### Execution Visualization

### Example: nums = [1, 3, 5], query SumRange(0, 2)
```
Iterative Segment Tree Construction:

Step 1: Initialize tree of size 6 (2*n)
tree = [0, 0, 0, 0, 0, 0]

Step 2: Fill leaf nodes (positions 3,4,5)
tree = [0, 0, 0, 1, 3, 5]

Step 3: Build internal nodes
- tree[2] = tree[4] + tree[5] = 3 + 5 = 8
- tree[1] = tree[2] + tree[3] = 8 + 1 = 9
tree = [9, 8, 0, 1, 3, 5]

Query SumRange(0, 2):
Convert: left = 0+3 = 3, right = 2+3 = 5
Iteration 1: left=3(odd)→sum+=1, left=4; right=5(odd)→sum+=5, right=4
Iteration 2: left=4, right=4 → left=2, right=2
Iteration 3: left=2(even), right=2(even) → no additions
Result: sum = 1 + 5 = 6 ✓
```

### Key Visualization Points:
- **Tree Structure**: Complete binary tree in array
- **Leaf Nodes**: Original array elements
- **Internal Nodes**: Sum of children
- **Query Decomposition**: Range covers O(log N) nodes

### Tree Visualization:
```
Array: [1, 3, 5]

Segment Tree:
        9 (sum of all)
       / \
      8   1 (sum of [0])
     / \
    3   5 (sum of [1,2])
   / \
  1   3   5 (individual elements)

Query SumRange(0,2):
Covers nodes: [1] + [3,5] = 6
```

### Time Complexity Breakdown:
- **Iterative Tree**: O(N) build, O(log N) update/query, O(N) space
- **Recursive Tree**: O(N) build, O(log N) update/query, O(4N) space
- **Lazy Propagation**: O(N) build, O(log N) operations, O(N) space
- **Fenwick Tree**: O(N) build, O(log N) operations, O(N) space

### Alternative Approaches:

#### 1. Fenwick Tree (Binary Indexed Tree) (O(N) space, O(log N) operations)
```go
type FenwickTree struct {
    tree []int
    n    int
}

func (ft *FenwickTree) Update(i, delta int) {
    // Point update with delta
    // O(log N) time
}

func (ft *FenwickTree) Query(i int) int {
    // Prefix sum query
    // O(log N) time
}
```
- **Pros**: More memory efficient, simpler implementation
- **Cons**: Only supports sum operations, no range updates

#### 2. Square Root Decomposition (O(N) space, O(√N) operations)
```go
type SqrtDecomposition struct {
    blocks []int
    blockSize int
    nums []int
}

func (sd *SqrtDecomposition) Update(i, val int) {
    // Update in O(√N) time
}

func (sd *SqrtDecomposition) Query(left, right int) int {
    // Query in O(√N) time
}
```
- **Pros**: Simple to understand and implement
- **Cons**: Slower than segment tree for large N

#### 3. Prefix Sum Array (O(N) space, O(1) query, O(N) update)
```go
type PrefixSum struct {
    prefix []int
    nums   []int
}

func (ps *PrefixSum) Query(left, right int) int {
    return ps.prefix[right+1] - ps.prefix[left]
}

func (ps *PrefixSum) Update(i, val int) {
    // O(N) time to update all prefixes
}
```
- **Pros**: Fast queries, simple implementation
- **Cons**: Slow updates, not suitable for dynamic arrays

### Extensions for Interviews:
- **Range Updates**: Add lazy propagation for efficient range updates
- **Different Operations**: Support min, max, gcd operations
- **Persistent Tree**: Maintain historical versions
- **2D Segment Tree**: Extend to 2D range queries
- **Custom Operations**: Support any associative operation
*/
func main() {
	// Test cases
	fmt.Println("=== Testing NumArray (Segment Tree) ===")
	
	// Test 1: Basic operations
	nums := []int{1, 3, 5}
	na := ConstructorNumArray(nums)
	
	fmt.Printf("Initial array: %v\n", nums)
	fmt.Printf("SumRange(0, 2): %d\n", na.SumRange(0, 2))
	fmt.Printf("SumRange(1, 2): %d\n", na.SumRange(1, 2))
	
	na.Update(1, 2)
	fmt.Printf("After Update(1, 2): SumRange(0, 2): %d\n", na.SumRange(0, 2))
	fmt.Printf("After Update(1, 2): SumRange(1, 2): %d\n", na.SumRange(1, 2))
	
	// Test 2: Recursive version
	fmt.Println("\n=== Testing Recursive Version ===")
	naRecursive := ConstructorNumArrayRecursive(nums)
	
	fmt.Printf("SumRange(0, 2): %d\n", naRecursive.SumRange(0, 2))
	naRecursive.Update(0, 10)
	fmt.Printf("After Update(0, 10): SumRange(0, 2): %d\n", naRecursive.SumRange(0, 2))
	
	// Test 3: Lazy propagation version
	fmt.Println("\n=== Testing Lazy Propagation Version ===")
	naLazy := ConstructorNumArrayLazy(nums)
	
	fmt.Printf("SumRange(0, 2): %d\n", naLazy.SumRange(0, 2))
	naLazy.UpdateRange(0, 1, 2)
	fmt.Printf("After UpdateRange(0, 1, 2): SumRange(0, 2): %d\n", naLazy.SumRange(0, 2))
	naLazy.UpdateRange(1, 2, 3)
	fmt.Printf("After UpdateRange(1, 2, 3): SumRange(0, 2): %d\n", naLazy.SumRange(0, 2))
	
	// Test 4: Large array
	fmt.Println("\n=== Testing Large Array ===")
	largeNums := make([]int, 1000)
	for i := range largeNums {
		largeNums[i] = i + 1
	}
	
	largeNa := ConstructorNumArray(largeNums)
	
	fmt.Printf("SumRange(0, 999): %d\n", largeNa.SumRange(0, 999))
	fmt.Printf("SumRange(500, 599): %d\n", largeNa.SumRange(500, 599))
	
	largeNa.Update(500, 1000)
	fmt.Printf("After Update(500, 1000): SumRange(500, 599): %d\n", largeNa.SumRange(500, 599))
	
	// Test 5: Edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	singleNums := []int{5}
	singleNa := ConstructorNumArray(singleNums)
	
	fmt.Printf("Single element - SumRange(0, 0): %d\n", singleNa.SumRange(0, 0))
	singleNa.Update(0, 10)
	fmt.Printf("After update - SumRange(0, 0): %d\n", singleNa.SumRange(0, 0))
	
	// Test 6: Performance comparison
	fmt.Println("\n=== Performance Comparison ===")
	perfNums := make([]int, 100)
	for i := range perfNums {
		perfNums[i] = i % 10
	}
	
	// Test iterative version
	perfNa := ConstructorNumArray(perfNums)
	start := 0
	for i := 0; i < 100; i++ {
		start += perfNa.SumRange(i%10, (i%10)+5)
	}
	
	// Test recursive version
	perfNaRecursive := ConstructorNumArrayRecursive(perfNums)
	for i := 0; i < 100; i++ {
		start += perfNaRecursive.SumRange(i%10, (i%10)+5)
	}
	
	fmt.Printf("Performance test completed for both versions\n")
	
	// Test 7: Range operations with lazy propagation
	fmt.Println("\n=== Testing Range Operations ===")
	rangeNums := []int{1, 2, 3, 4, 5}
	rangeNa := ConstructorNumArrayLazy(rangeNums)
	
	fmt.Printf("Initial SumRange(0, 4): %d\n", rangeNa.SumRange(0, 4))
	rangeNa.UpdateRange(0, 2, 10)
	fmt.Printf("After UpdateRange(0, 2, 10): SumRange(0, 4): %d\n", rangeNa.SumRange(0, 4))
	rangeNa.UpdateRange(2, 4, 5)
	fmt.Printf("After UpdateRange(2, 4, 5): SumRange(0, 4): %d\n", rangeNa.SumRange(0, 4))
}
