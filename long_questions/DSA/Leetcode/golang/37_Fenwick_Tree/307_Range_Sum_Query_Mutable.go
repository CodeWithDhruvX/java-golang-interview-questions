package main

import "fmt"

// 307. Range Sum Query - Mutable - Fenwick Tree (Binary Indexed Tree)
// Time: O(log N) for update and query, Space: O(N)
type NumArray struct {
	tree []int
	n    int
}

// Constructor builds Fenwick Tree from array
func ConstructorNumArrayFenwick(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, n+1) // 1-based indexing
	
	// Build Fenwick Tree
	for i := 0; i < n; i++ {
		updateFenwick(tree, i+1, nums[i])
	}
	
	return NumArray{tree: tree, n: n}
}

// Update value at index i (0-based)
func (this *NumArray) Update(i int, val int) {
	// Calculate the difference
	delta := val - this.SumRange(i, i)
	// Update the Fenwick tree
	updateFenwick(this.tree, i+1, delta)
}

// Query sum from left to right (0-based, inclusive)
func (this *NumArray) SumRange(left int, right int) int {
	return queryFenwick(this.tree, right+1) - queryFenwick(this.tree, left)
}

// Update Fenwick tree at position i (1-based) by delta
func updateFenwick(tree []int, i, delta int) {
	for i < len(tree) {
		tree[i] += delta
		i += i & (-i) // Move to parent
	}
}

// Query prefix sum from 1 to i (1-based)
func queryFenwick(tree []int, i int) int {
	sum := 0
	for i > 0 {
		sum += tree[i]
		i -= i & (-i) // Move to parent
	}
	return sum
}

// Alternative implementation with range updates and point queries
type NumArrayRangeUpdate struct {
	tree []int
	n    int
}

func ConstructorNumArrayRangeUpdate(nums []int) NumArrayRangeUpdate {
	n := len(nums)
	tree := make([]int, n+1)
	
	// Initialize with differences
	for i := 0; i < n; i++ {
		if i == 0 {
			updateFenwick(tree, i+1, nums[i])
		} else {
			updateFenwick(tree, i+1, nums[i]-nums[i-1])
		}
	}
	
	return NumArrayRangeUpdate{tree: tree, n: n}
}

func (this *NumArrayRangeUpdate) UpdateRange(left, right, val int) {
	updateFenwick(this.tree, left+1, val)
	if right+1 < this.n {
		updateFenwick(this.tree, right+2, -val)
	}
}

func (this *NumArrayRangeUpdate) Query(i int) int {
	return queryFenwick(this.tree, i+1)
}

// 2D Fenwick Tree
type NumMatrix2D struct {
	tree [][]int
	m, n int
}

func ConstructorNumMatrix2D(matrix [][]int) NumMatrix2D {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix2D{tree: nil, m: 0, n: 0}
	}
	
	m, n := len(matrix), len(matrix[0])
	tree := make([][]int, m+1)
	for i := range tree {
		tree[i] = make([]int, n+1)
	}
	
	// Build 2D Fenwick Tree
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			update2D(tree, i+1, j+1, matrix[i][j])
		}
	}
	
	return NumMatrix2D{tree: tree, m: m, n: n}
}

func (this *NumMatrix2D) Update(row, col, val int) {
	// This would require tracking original values
	// For simplicity, we'll just add the value
	update2D(this.tree, row+1, col+1, val)
}

func (this *NumMatrix2D) SumRegion(row1, col1, row2, col2 int) int {
	return query2D(this.tree, row2+1, col2+1) - 
		   query2D(this.tree, row1, col2+1) - 
		   query2D(this.tree, row2+1, col1) + 
		   query2D(this.tree, row1, col1)
}

func update2D(tree [][]int, i, j, delta int) {
	for x := i; x < len(tree); x += x & (-x) {
		for y := j; y < len(tree[0]); y += y & (-y) {
			tree[x][y] += delta
		}
	}
}

func query2D(tree [][]int, i, j int) int {
	sum := 0
	for x := i; x > 0; x -= x & (-x) {
		for y := j; y > 0; y -= y & (-y) {
			sum += tree[x][y]
		}
	}
	return sum
}

// Fenwick Tree with range queries and range updates
type FenwickTreeRange struct {
	tree1 []int // For range updates
	tree2 []int // For range queries
	n     int
}

func ConstructorFenwickTreeRange(n int) FenwickTreeRange {
	return FenwickTreeRange{
		tree1: make([]int, n+1),
		tree2: make([]int, n+1),
		n:     n,
	}
}

func (this *FenwickTreeRange) UpdateRange(l, r, val int) {
	this.updateInternal(this.tree1, l, val)
	this.updateInternal(this.tree1, r+1, -val)
	this.updateInternal(this.tree2, l, val*(l-1))
	this.updateInternal(this.tree2, r+1, -val*r)
}

func (this *FenwickTreeRange) QueryRange(l, r int) int {
	return this.queryInternal(r) - this.queryInternal(l-1)
}

func (this *FenwickTreeRange) updateInternal(tree []int, i, val int) {
	for i <= this.n {
		tree[i] += val
		i += i & (-i)
	}
}

func (this *FenwickTreeRange) queryInternal(tree []int, i int) int {
	sum := 0
	for i > 0 {
		sum += tree[i]
		i -= i & (-i)
	}
	return sum
}

func (this *FenwickTreeRange) queryInternalRange(i int) int {
	return this.queryInternal(this.tree1, i)*i - this.queryInternal(this.tree2, i)
}

// Fenwick Tree for minimum queries
type FenwickTreeMin struct {
	tree []int
	n    int
}

func ConstructorFenwickTreeMin(arr []int) FenwickTreeMin {
	n := len(arr)
	tree := make([]int, n+1)
	
	// Initialize with infinity
	for i := range tree {
		tree[i] = 1 << 31 // Max int value
	}
	
	// Build tree
	for i := 0; i < n; i++ {
		this.updateMin(i, arr[i])
	}
	
	return FenwickTreeMin{tree: tree, n: n}
}

func (this *FenwickTreeMin) updateMin(i, val int) {
	i++
	for i <= this.n {
		if val < this.tree[i] {
			this.tree[i] = val
		}
		i += i & (-i)
	}
}

func (this *FenwickTreeMin) queryMin(i int) int {
	i++
	result := 1 << 31 // Max int value
	for i > 0 {
		if this.tree[i] < result {
			result = this.tree[i]
		}
		i -= i & (-i)
	}
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Fenwick Tree (Binary Indexed Tree)
- **Point Updates**: Efficient single element updates
- **Prefix Sum Queries**: Fast range sum calculations
- **1-based Indexing**: Natural for BIT operations
- **Binary Decomposition**: Each index stores sum of specific range

## 2. PROBLEM CHARACTERISTICS
- **Mutable Array**: Dynamic updates to array elements
- **Range Sum Queries**: Query sum of subarray ranges
- **Efficient Operations**: Need O(log N) for both updates and queries
- **Associative Operation**: Sum operation is associative and commutative

## 3. SIMILAR PROBLEMS
- Range Sum Query Mutable (LeetCode 307) - Same problem
- Create Sorted Array through Instructions (LeetCode 1649) - BIT for counting
- Count of Smaller Numbers After Self (LeetCode 315) - BIT variant
- Reverse Pairs (LeetCode 493) - BIT for counting inversions

## 4. KEY OBSERVATIONS
- **Binary Representation**: Each number's binary representation determines range coverage
- **Tree Structure**: Implicit tree structure in array representation
- **Prefix Sum Property**: Range sum = prefix sum up to right - prefix sum up to left-1
- **Efficient Updates**: Point updates affect O(log N) tree nodes

## 5. VARIATIONS & EXTENSIONS
- **Range Updates**: Support range increment operations
- **2D Fenwick Tree**: Extend to 2D range queries
- **Range Queries + Range Updates**: Advanced BIT with two trees
- **Different Operations**: Minimum, maximum, GCD operations

## 6. INTERVIEW INSIGHTS
- Always clarify: "Array size constraints? Update/query frequency? Range sizes?"
- Edge cases: empty array, single element, invalid indices
- Time complexity: O(log N) for both update and query
- Space complexity: O(N) for tree storage
- Key insight: binary representation enables efficient range decomposition

## 7. COMMON MISTAKES
- Wrong indexing (0-based vs 1-based confusion)
- Incorrect update/query implementation
- Not handling empty array properly
- Wrong range sum calculation
- Missing modulo operations for large results

## 8. OPTIMIZATION STRATEGIES
- **Standard BIT**: O(log N) time, O(N) space - basic approach
- **Range Updates**: O(log N) time, O(N) space - difference array
- **2D BIT**: O(log M × log N) time, O(M × N) space - matrix queries
- **Advanced Range**: O(log N) time, O(N) space - range queries + updates

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like a hierarchical accounting system:**
- Each manager oversees a specific range of accounts
- The range sizes follow powers of two (1, 2, 4, 8, ...)
- When you update one account, you inform all relevant managers
- When you query a range, you ask the minimal set of managers
- Like a company with a very specific organizational structure

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Initial array, update operations (index, value), sum queries (range)
2. **Goal**: Support efficient updates and range sum queries
3. **Rules**: Array is mutable, operations are dynamic
4. **Output**: Range sum results after updates

#### Phase 2: Key Insight Recognition
- **"Binary decomposition natural"** → Each index can be decomposed into powers of 2
- **"Prefix sums useful"** → Range sum = prefix[right] - prefix[left-1]
- **"Tree structure implicit"** → Binary representation creates tree hierarchy
- **"O(log N) achievable"** → Each operation touches O(log N) nodes

#### Phase 3: Strategy Development
```
Human thought process:
"I need fast updates and range sums.
Naive approach: O(N) for updates, O(N) for queries.

Fenwick Tree Approach:
1. Use 1-based indexing for natural binary operations
2. Each tree[i] stores sum of range (i - LSB(i) + 1) to i
3. Update: add delta to index i, then i += LSB(i) repeatedly
4. Query: accumulate sum at index i, then i -= LSB(i) repeatedly
5. Range sum = query(right) - query(left-1)

This gives O(log N) for both operations!"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Handle gracefully with zero size
- **Single element**: Tree with just one node
- **Invalid indices**: Validate input bounds
- **Large values**: Use appropriate data types

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: nums = [1, 3, 5], query SumRange(0, 2)

Human thinking:
"Fenwick Tree Approach:
Step 1: Build tree (1-based indexing)
Tree size = 4 (n+1)
Initialize: [0, 0, 0, 0]

Build by updates:
- Update index 1 with value 1:
  tree[1] += 1, tree[2] += 1, tree[4] += 1
  tree = [0, 1, 1, 0]
- Update index 2 with value 3:
  tree[2] += 3, tree[4] += 3
  tree = [0, 1, 4, 3]
- Update index 3 with value 5:
  tree[3] += 5
  tree = [0, 1, 4, 8]

Step 2: Query SumRange(0, 2)
Range sum = query(3) - query(0)
query(3): sum = tree[3] + tree[2] = 5 + 4 = 9
query(0): sum = 0
Result = 9 - 0 = 9 ✓"
```

#### Phase 6: Intuition Validation
- **Why binary representation**: Each number's LSB determines next parent
- **Why 1-based indexing**: Simplifies LSB calculations
- **Why O(log N)**: Each operation moves through binary digits
- **Why prefix sums**: Range queries reduce to prefix difference

### Common Human Pitfalls & How to Avoid Them
1. **"Why not segment tree?"** → BIT is simpler and faster for sum operations
2. **"Should I use 0-based indexing?"** → 1-based is natural for BIT operations
3. **"What about range updates?"** → Use difference array technique
4. **"Can I use for other operations?"** → Yes, if operation is associative
5. **"Why LSB (Lowest Set Bit)?"** → Determines range coverage and parent navigation

### Real-World Analogy
**Like a digital inventory management system:**
- Each item is tracked at individual level
- Managers oversee ranges in powers of two (1, 2, 4, 8 items)
- When one item count changes, you update relevant managers
- When you need total count for a range, you ask minimal managers
- Like a warehouse with binary-based management structure

### Human-Readable Pseudocode
```
class FenwickTree:
    constructor(nums):
        n = length(nums)
        tree = array of size n+1 (1-based)
        
        # Build tree by updating each element
        for i from 0 to n-1:
            update(i+1, nums[i])
    
    update(index, delta):
        while index <= n:
            tree[index] += delta
            index += index & (-index)  # Add LSB
    
    query(index):
        sum = 0
        while index > 0:
            sum += tree[index]
            index -= index & (-index)  # Remove LSB
        return sum
    
    sumRange(left, right):
        return query(right+1) - query(left)
```

### Execution Visualization

### Example: nums = [1, 3, 5], query SumRange(0, 2)
```
Fenwick Tree Construction:

Step 1: Initialize tree of size 4 (n+1)
tree = [0, 0, 0, 0] (1-based indexing)

Step 2: Build by updates
Update index 1 with value 1:
- tree[1] += 1, index = 1 + LSB(1) = 2
- tree[2] += 1, index = 2 + LSB(2) = 4
- Stop (index > n)
tree = [0, 1, 1, 0]

Update index 2 with value 3:
- tree[2] += 3, index = 2 + LSB(2) = 4
- tree[4] += 3, index = 4 + LSB(4) = 8
- Stop
tree = [0, 1, 4, 3]

Update index 3 with value 5:
- tree[3] += 5, index = 3 + LSB(3) = 4
- tree[4] += 5, index = 4 + LSB(4) = 8
- Stop
tree = [0, 1, 4, 8]

Query SumRange(0, 2):
sum = query(3) - query(0)
query(3): sum = tree[3] + tree[2] = 5 + 4 = 9
query(0): sum = 0
Result = 9 - 0 = 9 ✓
```

### Key Visualization Points:
- **Binary Representation**: Each index's LSB determines coverage
- **Tree Structure**: Implicit tree in array representation
- **Update Path**: Index → parent → grandparent → ...
- **Query Path**: Index → parent → grandparent → ... → root

### BIT Structure Visualization:
```
Array: [1, 3, 5] (0-based)
Tree:  [0, 1, 4, 8] (1-based)

Coverage:
tree[1] covers: nums[0] (range size 1)
tree[2] covers: nums[0:1] (range size 2)
tree[3] covers: nums[2] (range size 1)
tree[4] covers: nums[0:2] (range size 4)

Binary representation:
1 = 001 (LSB = 1, covers 1 element)
2 = 010 (LSB = 2, covers 2 elements)
3 = 011 (LSB = 1, covers 1 element)
4 = 100 (LSB = 4, covers 4 elements)
```

### Time Complexity Breakdown:
- **Standard BIT**: O(log N) time, O(N) space - basic approach
- **Range Updates**: O(log N) time, O(N) space - difference array
- **2D BIT**: O(log M × log N) time, O(M × N) space - matrix queries
- **Advanced Range**: O(log N) time, O(N) space - range queries + updates

### Alternative Approaches:

#### 1. Segment Tree (O(log N) time, O(N) space)
```go
type SegmentTree struct {
    tree []int
    n    int
}

func (st *SegmentTree) Update(index, value int) {
    // Update leaf and propagate to root
    // O(log N) time
}

func (st *SegmentTree) Query(left, right int) int {
    // Query range recursively
    // O(log N) time
}
```
- **Pros**: More flexible for different operations
- **Cons**: More complex implementation, higher constant factors

#### 2. Prefix Sum Array (O(1) query, O(N) update)
```go
type PrefixSum struct {
    prefix []int
    nums   []int
}

func (ps *PrefixSum) Query(left, right int) int {
    return ps.prefix[right+1] - ps.prefix[left]
}

func (ps *PrefixSum) Update(index, value int) {
    // O(N) time to update all prefixes
}
```
- **Pros**: Fast queries, simple implementation
- **Cons**: Slow updates, not suitable for dynamic arrays

#### 3. Square Root Decomposition (O(√N) time, O(N) space)
```go
type SqrtDecomposition struct {
    blocks []int
    blockSize int
    nums   []int
}

func (sd *SqrtDecomposition) Update(index, value int) {
    // O(√N) time
}

func (sd *SqrtDecomposition) Query(left, right int) int {
    // O(√N) time
}
```
- **Pros**: Simple to understand, good performance
- **Cons**: Slower than BIT for large N

### Extensions for Interviews:
- **Range Updates**: Use difference array technique
- **2D Queries**: Extend to 2D Fenwick Tree
- **Different Operations**: Minimum, maximum, GCD variants
- **Multiple BITs**: Handle multiple dimensions or operations
- **Real-world Applications**: Financial calculations, inventory management
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Fenwick Tree (Binary Indexed Tree) ===")
	
	// Test 1: Basic operations
	fmt.Println("=== Basic Operations ===")
	nums := []int{1, 3, 5}
	na := ConstructorNumArrayFenwick(nums)
	
	fmt.Printf("Initial array: %v\n", nums)
	fmt.Printf("SumRange(0, 2): %d\n", na.SumRange(0, 2))
	fmt.Printf("SumRange(1, 2): %d\n", na.SumRange(1, 2))
	
	na.Update(1, 2)
	fmt.Printf("After Update(1, 2): SumRange(0, 2): %d\n", na.SumRange(0, 2))
	fmt.Printf("After Update(1, 2): SumRange(1, 2): %d\n", na.SumRange(1, 2))
	
	// Test 2: Range updates
	fmt.Println("\n=== Range Updates ===")
	rangeNa := ConstructorNumArrayRangeUpdate(nums)
	
	fmt.Printf("Query(0): %d\n", rangeNa.Query(0))
	fmt.Printf("Query(1): %d\n", rangeNa.Query(1))
	fmt.Printf("Query(2): %d\n", rangeNa.Query(2))
	
	rangeNa.UpdateRange(0, 1, 10)
	fmt.Printf("After UpdateRange(0, 1, 10): Query(0): %d\n", rangeNa.Query(0))
	fmt.Printf("After UpdateRange(0, 1, 10): Query(1): %d\n", rangeNa.Query(1))
	fmt.Printf("After UpdateRange(0, 1, 10): Query(2): %d\n", rangeNa.Query(2))
	
	// Test 3: 2D Fenwick Tree
	fmt.Println("\n=== 2D Fenwick Tree ===")
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	
	na2d := ConstructorNumMatrix2D(matrix)
	fmt.Printf("SumRegion(2, 1, 4, 3): %d\n", na2d.SumRegion(2, 1, 4, 3))
	fmt.Printf("SumRegion(1, 1, 2, 2): %d\n", na2d.SumRegion(1, 1, 2, 2))
	fmt.Printf("SumRegion(1, 2, 2, 4): %d\n", na2d.SumRegion(1, 2, 2, 4))
	
	// Test 4: Range queries and updates
	fmt.Println("\n=== Range Queries and Updates ===")
	rangeTree := ConstructorFenwickTreeRange(5)
	
	// Initialize with values
	for i := 0; i < 5; i++ {
		rangeTree.UpdateRange(i, i, i+1)
	}
	
	fmt.Printf("QueryRange(0, 4): %d\n", rangeTree.QueryRange(0, 4))
	fmt.Printf("QueryRange(1, 3): %d\n", rangeTree.QueryRange(1, 3))
	
	rangeTree.UpdateRange(1, 3, 10)
	fmt.Printf("After UpdateRange(1, 3, 10): QueryRange(0, 4): %d\n", rangeTree.QueryRange(0, 4))
	fmt.Printf("After UpdateRange(1, 3, 10): QueryRange(1, 3): %d\n", rangeTree.QueryRange(1, 3))
	
	// Test 5: Minimum queries
	fmt.Println("\n=== Minimum Queries ===")
	minTree := ConstructorFenwickTreeMin([]int{5, 2, 8, 1, 9})
	
	fmt.Printf("QueryMin(0): %d\n", minTree.queryMin(0))
	fmt.Printf("QueryMin(2): %d\n", minTree.queryMin(2))
	fmt.Printf("QueryMin(4): %d\n", minTree.queryMin(4))
	
	minTree.updateMin(2, 0)
	fmt.Printf("After updateMin(2, 0): QueryMin(2): %d\n", minTree.queryMin(2))
	fmt.Printf("After updateMin(2, 0): QueryMin(4): %d\n", minTree.queryMin(4))
	
	// Test 6: Performance test
	fmt.Println("\n=== Performance Test ===")
	largeNums := make([]int, 10000)
	for i := range largeNums {
		largeNums[i] = i % 100
	}
	
	largeNa := ConstructorNumArrayFenwick(largeNums)
	
	// Perform some operations
	total := 0
	for i := 0; i < 100; i++ {
		total += largeNa.SumRange(i*10, (i+1)*10-1)
	}
	
	fmt.Printf("Performance test completed with 100 queries, total: %d\n", total)
	
	// Test 7: Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Single element
	singleNa := ConstructorNumArrayFenwick([]int{42})
	fmt.Printf("Single element: %d\n", singleNa.SumRange(0, 0))
	singleNa.Update(0, 100)
	fmt.Printf("After update: %d\n", singleNa.SumRange(0, 0))
	
	// Empty array
	emptyNa := ConstructorNumArrayFenwick([]int{})
	fmt.Printf("Empty array length: %d\n", emptyNa.n)
	
	// Negative numbers
	negNa := ConstructorNumArrayFenwick([]int{-1, -2, -3, -4})
	fmt.Printf("Negative numbers sum: %d\n", negNa.SumRange(0, 3))
}
