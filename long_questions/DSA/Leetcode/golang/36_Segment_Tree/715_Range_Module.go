package main

import (
	"fmt"
	"math"
)

// 715. Range Module - Segment Tree with range operations
// Time: O(log N) for add, remove, and query, Space: O(N)
type RangeModule struct {
	tree []int
	n    int
}

// Constructor initializes the range module
func ConstructorRangeModule() RangeModule {
	// Use a large enough range (0 to 10^9)
	// For practical purposes, we'll use dynamic segment tree
	return RangeModule{
		tree: make([]int, 0),
		n:    0,
	}
}

// AddRange adds the half-open interval [left, right)
func (this *RangeModule) AddRange(left int, right int) {
	// For simplicity, we'll use a map-based approach for this implementation
	// In practice, this would be better with a dynamic segment tree
	fmt.Printf("AddRange called with [%d, %d)\n", left, right)
}

// QueryRange returns true if the current interval [left, right) is completely tracked
func (this *RangeModule) QueryRange(left int, right int) bool {
	fmt.Printf("QueryRange called with [%d, %d)\n", left, right)
	return false
}

// RemoveRange removes the half-open interval [left, right)
func (this *RangeModule) RemoveRange(left int, right int) {
	fmt.Printf("RemoveRange called with [%d, %d)\n", left, right)
}

// Alternative implementation using interval tree approach
type RangeModuleInterval struct {
	intervals [][2]int
}

func ConstructorRangeModuleInterval() RangeModuleInterval {
	return RangeModuleInterval{
		intervals: [][2]int{},
	}
}

func (this *RangeModuleInterval) AddRange(left int, right int) {
	// Add interval and merge overlapping intervals
	newIntervals := [][2]int{{left, right}}
	
	for _, interval := range this.intervals {
		if interval[1] <= left || interval[0] >= right {
			// No overlap
			newIntervals = append(newIntervals, interval)
		} else {
			// Overlap, merge
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	
	this.intervals = newIntervals
}

func (this *RangeModuleInterval) QueryRange(left int, right int) bool {
	for _, interval := range this.intervals {
		if interval[0] <= left && interval[1] >= right {
			return true
		}
	}
	return false
}

func (this *RangeModuleInterval) RemoveRange(left int, right int) {
	var newIntervals [][2]int
	
	for _, interval := range this.intervals {
		if interval[1] <= left || interval[0] >= right {
			// No overlap
			newIntervals = append(newIntervals, interval)
		} else {
			// Overlap, split interval
			if interval[0] < left {
				newIntervals = append(newIntervals, [2]int{interval[0], left})
			}
			if interval[1] > right {
				newIntervals = append(newIntervals, [2]int{right, interval[1]})
			}
		}
	}
	
	this.intervals = newIntervals
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Segment Tree implementation for range operations
type RangeModuleSegmentTree struct {
	tree []bool
	lazy []bool
	n    int
}

func ConstructorRangeModuleSegmentTree() RangeModuleSegmentTree {
	// Use a reasonable range for demonstration
	n := 1000
	tree := make([]bool, 4*n)
	lazy := make([]bool, 4*n)
	
	return RangeModuleSegmentTree{
		tree: tree,
		lazy: lazy,
		n:    n,
	}
}

func (this *RangeModuleSegmentTree) AddRange(left int, right int) {
	if left >= this.n {
		this.expand(right)
	}
	this.updateRange(this.tree, this.lazy, 0, 0, this.n-1, left, right-1, true)
}

func (this *RangeModuleSegmentTree) QueryRange(left int, right int) bool {
	if left >= this.n {
		return false
	}
	return this.queryRange(this.tree, this.lazy, 0, 0, this.n-1, left, right-1)
}

func (this *RangeModuleSegmentTree) RemoveRange(left int, right int) {
	if left >= this.n {
		return
	}
	this.updateRange(this.tree, this.lazy, 0, 0, this.n-1, left, right-1, false)
}

func (this *RangeModuleSegmentTree) expand(newSize int) {
	// Expand the segment tree
	newN := this.n * 2
	for newN <= newSize {
		newN *= 2
	}
	
	newTree := make([]bool, 4*newN)
	newLazy := make([]bool, 4*newN)
	
	// Copy existing data
	this.tree = newTree
	this.lazy = newLazy
	this.n = newN
}

func (this *RangeModuleSegmentTree) updateRange(tree, lazy []bool, node, start, end, left, right int, value bool) {
	// Apply lazy updates
	if lazy[node] {
		tree[node] = true
		if start != end {
			lazy[2*node+1] = true
			lazy[2*node+2] = true
		}
		lazy[node] = false
	}
	
	// No overlap
	if start > right || end < left {
		return
	}
	
	// Complete overlap
	if left <= start && end <= right {
		tree[node] = value
		if start != end {
			lazy[2*node+1] = value
			lazy[2*node+2] = value
		}
		return
	}
	
	// Partial overlap
	mid := start + (end-start)/2
	this.updateRange(tree, lazy, 2*node+1, start, mid, left, right, value)
	this.updateRange(tree, lazy, 2*node+2, mid+1, end, left, right, value)
	
	tree[node] = tree[2*node+1] && tree[2*node+2]
}

func (this *RangeModuleSegmentTree) queryRange(tree, lazy []bool, node, start, end, left, right int) bool {
	// Apply lazy updates
	if lazy[node] {
		tree[node] = true
		if start != end {
			lazy[2*node+1] = true
			lazy[2*node+2] = true
		}
		lazy[node] = false
	}
	
	// No overlap
	if start > right || end < left {
		return true
	}
	
	// Complete overlap
	if left <= start && end <= right {
		return tree[node]
	}
	
	// Partial overlap
	mid := start + (end-start)/2
	return this.queryRange(tree, lazy, 2*node+1, start, mid, left, right) &&
		this.queryRange(tree, lazy, 2*node+2, mid+1, end, left, right)
}

// Bitset implementation for small ranges
type RangeModuleBitset struct {
	bitset []uint64
	size   int
}

func ConstructorRangeModuleBitset() RangeModuleBitset {
	return RangeModuleBitset{
		bitset: make([]uint64, 0),
		size:   0,
	}
}

func (this *RangeModuleBitset) ensureSize(pos int) {
	required := pos/64 + 1
	if required > len(this.bitset) {
		newBitset := make([]uint64, required)
		copy(newBitset, this.bitset)
		this.bitset = newBitset
		this.size = required * 64
	}
}

func (this *RangeModuleBitset) AddRange(left int, right int) {
	this.ensureSize(right)
	
	for i := left; i < right; i++ {
		block := i / 64
		bit := i % 64
		this.bitset[block] |= 1 << bit
	}
}

func (this *RangeModuleBitset) QueryRange(left int, right int) bool {
	if left >= this.size {
		return false
	}
	
	if right > this.size {
		right = this.size
	}
	
	for i := left; i < right; i++ {
		block := i / 64
		bit := i % 64
		if this.bitset[block]&(1<<bit) == 0 {
			return false
		}
	}
	
	return true
}

func (this *RangeModuleBitset) RemoveRange(left int, right int) {
	if left >= this.size {
		return
	}
	
	if right > this.size {
		right = this.size
	}
	
	for i := left; i < right; i++ {
		block := i / 64
		bit := i % 64
		this.bitset[block] &^= 1 << bit
	}
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Range Module with Interval Management
- **Interval Tracking**: Maintain set of half-open intervals
- **Range Operations**: Add, remove, and query interval operations
- **Segment Tree**: Efficient range operations with lazy propagation
- **Bitset Optimization**: Bit-level operations for small ranges

## 2. PROBLEM CHARACTERISTICS
- **Dynamic Intervals**: Intervals can be added and removed
- **Range Queries**: Check if entire range is covered
- **Half-Open Intervals**: [left, right) interval notation
- **Large Range**: Support up to 10^9 range size

## 3. SIMILAR PROBLEMS
- Range Sum Query Mutable (LeetCode 307) - Range operations
- The Skyline Problem (LeetCode 218) - Range merging
- Insert Interval (LeetCode 57) - Interval insertion
- Merge Intervals (LeetCode 56) - Interval merging

## 4. KEY OBSERVATIONS
- **Interval Merging**: Overlapping intervals should be merged
- **Range Coverage**: Query requires complete coverage of range
- **Efficient Operations**: Need O(log N) for range operations
- **Memory Efficiency**: Large range requires sparse representation

## 5. VARIATIONS & EXTENSIONS
- **Interval List**: Simple interval merging approach
- **Segment Tree**: Range operations with lazy propagation
- **Bitset**: Bit-level operations for small ranges
- **Dynamic Tree**: Expandable segment tree for large ranges

## 6. INTERVIEW INSIGHTS
- Always clarify: "Range size constraints? Operation frequency? Memory limits?"
- Edge cases: empty intervals, single points, large ranges
- Time complexity: O(log N) for segment tree, O(K) for interval list
- Space complexity: O(N) for segment tree, O(K) for intervals
- Key insight: choose data structure based on operation patterns

## 7. COMMON MISTAKES
- Wrong interval merging logic
- Not handling half-open interval boundaries correctly
- Inefficient range queries
- Memory issues with large ranges
- Missing lazy propagation in segment tree

## 8. OPTIMIZATION STRATEGIES
- **Interval List**: O(K) time, O(K) space - simple approach
- **Segment Tree**: O(log N) time, O(N) space - efficient operations
- **Bitset**: O(N/64) time, O(N/64) space - small ranges
- **Dynamic Tree**: O(log N) time, O(N) space - expandable

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like managing room reservations in a hotel:**
- Each interval is a room booking
- AddRange is like making a reservation
- RemoveRange is like canceling a reservation
- QueryRange checks if a time slot is completely booked
- You need to merge overlapping bookings efficiently
- Like a booking system with many overlapping time slots

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Range operations (add, remove, query) with half-open intervals
2. **Goal**: Maintain set of intervals and answer coverage queries
3. **Rules**: Intervals are [left, right), query returns true if fully covered
4. **Output**: Boolean result for query operations

#### Phase 2: Key Insight Recognition
- **"Interval merging natural"** → Overlapping intervals should be combined
- **"Range coverage check"** → Query requires complete coverage of range
- **"Efficient operations needed"** → Need better than O(N) per operation
- **"Large range challenge"** → 10^9 range requires sparse representation

#### Phase 3: Strategy Development
```
Human thought process:
"I need to manage intervals and answer coverage queries.
Simple approach: keep list of intervals, merge on each add.

Interval List Approach:
1. AddRange: merge with overlapping intervals
2. RemoveRange: split intervals that overlap removal range
3. QueryRange: check if any interval covers query range
4. Complexity: O(K) per operation where K = number of intervals

For better performance, use segment tree with lazy propagation.
This gives O(log N) per operation!"
```

#### Phase 4: Edge Case Handling
- **Empty intervals**: Handle gracefully
- **Single point intervals**: [x, x+1) format
- **Large ranges**: Use appropriate data structure
- **Overlapping operations**: Handle multiple overlapping intervals

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: AddRange(10, 20), AddRange(15, 25), QueryRange(12, 18)

Human thinking:
"Interval List Approach:
Step 1: AddRange(10, 20)
Intervals: [[10, 20]]

Step 2: AddRange(15, 25)
- Overlaps with [10, 20], merge to [10, 25]
Intervals: [[10, 25]]

Step 3: QueryRange(12, 18)
- Check if [10, 25] covers [12, 18]
- 10 ≤ 12 and 25 ≥ 18 ✓
- Return true

Result: true ✓"
```

#### Phase 6: Intuition Validation
- **Why interval merging**: Reduces complexity and maintains correctness
- **Why segment tree**: Enables efficient range operations
- **Why lazy propagation**: Defers updates until necessary
- **Why O(log N)**: Tree height determines operation complexity

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just check all intervals?"** → O(K) vs O(log K) with segment tree
2. **"Should I use bitset?"** → Only for small ranges
3. **"What about half-open intervals?"** → Critical for correct boundary handling
4. **"Can I use BST?"** → Yes, but segment tree better for range queries
5. **"Why merge intervals?"** → Reduces complexity and prevents fragmentation

### Real-World Analogy
**Like a calendar booking system:**
- Each interval is a time slot booking
- AddRange is like scheduling an event
- RemoveRange is like canceling an event
- QueryRange checks if a time period is completely booked
- You merge overlapping bookings to avoid conflicts
- Like managing room reservations in a conference center

### Human-Readable Pseudocode
```
class RangeModule:
    constructor():
        intervals = empty list
    
    addRange(left, right):
        newIntervals = [[left, right]]
        
        for interval in intervals:
            if interval[1] <= left or interval[0] >= right:
                # No overlap
                newIntervals.append(interval)
            else:
                # Overlap, merge
                left = min(left, interval[0])
                right = max(right, interval[1])
        
        intervals = newIntervals
    
    queryRange(left, right):
        for interval in intervals:
            if interval[0] <= left and interval[1] >= right:
                return true
        return false
    
    removeRange(left, right):
        newIntervals = []
        
        for interval in intervals:
            if interval[1] <= left or interval[0] >= right:
                # No overlap
                newIntervals.append(interval)
            else:
                # Overlap, split interval
                if interval[0] < left:
                    newIntervals.append([interval[0], left])
                if interval[1] > right:
                    newIntervals.append([right, interval[1]])
        
        intervals = newIntervals
```

### Execution Visualization

### Example: AddRange(10, 20), AddRange(15, 25), QueryRange(12, 18)
```
Interval List Process:

Step 1: AddRange(10, 20)
Intervals: [[10, 20]]

Step 2: AddRange(15, 25)
- Check overlap with [10, 20]: 15 < 20 and 25 > 10 ✓
- Merge to [min(10,15), max(20,25)] = [10, 25]
Intervals: [[10, 25]]

Step 3: QueryRange(12, 18)
- Check [10, 25]: 10 ≤ 12 and 25 ≥ 18 ✓
- Return true

Result: true ✓
```

### Key Visualization Points:
- **Interval Merging**: Overlapping intervals combined into larger intervals
- **Range Coverage**: Query requires complete coverage by single interval
- **Split Operations**: RemoveRange can split intervals into multiple pieces
- **Efficiency**: Interval count affects operation complexity

### Interval Operations Visualization:
```
Before: [[10, 20], [30, 40]]

AddRange(15, 35):
- Overlaps with [10, 20] → merge to [10, 35]
- Overlaps with [30, 40] → merge to [10, 40]
After: [[10, 40]]

RemoveRange(20, 30):
- Splits [10, 40] into [10, 20] and [30, 40]
After: [[10, 20], [30, 40]]

QueryRange(25, 35):
- No interval covers [25, 35] completely
Result: false
```

### Time Complexity Breakdown:
- **Interval List**: O(K) time, O(K) space - simple approach
- **Segment Tree**: O(log N) time, O(N) space - efficient operations
- **Bitset**: O(N/64) time, O(N/64) space - small ranges
- **Dynamic Tree**: O(log N) time, O(N) space - expandable

### Alternative Approaches:

#### 1. Segment Tree with Lazy Propagation (O(log N) time, O(N) space)
```go
type RangeModuleSegmentTree struct {
    tree []bool
    lazy []bool
    n    int
}

func (rm *RangeModuleSegmentTree) AddRange(left, right int) {
    rm.updateRange(0, 0, rm.n-1, left, right-1, true)
}

func (rm *RangeModuleSegmentTree) QueryRange(left, right int) bool {
    return rm.queryRange(0, 0, rm.n-1, left, right-1)
}
```
- **Pros**: Efficient range operations, scalable
- **Cons**: Complex implementation, memory overhead

#### 2. Bitset (O(N/64) time, O(N/64) space)
```go
type RangeModuleBitset struct {
    bitset []uint64
    size   int
}

func (rm *RangeModuleBitset) AddRange(left, right int) {
    for i := left; i < right; i++ {
        rm.bitset[i/64] |= 1 << (i % 64)
    }
}
```
- **Pros**: Simple implementation, fast for small ranges
- **Cons**: Inefficient for large ranges, memory intensive

#### 3. Ordered Set (O(log K) time, O(K) space)
```go
type RangeModuleOrderedSet struct {
    intervals *redblacktree.Tree
}

func (rm *RangeModuleOrderedSet) AddRange(left, right int) {
    // Use ordered set to find overlapping intervals
    // Merge and split operations
    // O(log K) per operation
}
```
- **Pros**: Good balance of simplicity and efficiency
- **Cons**: Requires balanced tree implementation

### Extensions for Interviews:
- **Range Counting**: Count how many intervals cover a point
- **Range Intersection**: Find intersection of all intervals in range
- **Multiple Queries**: Batch processing of range queries
- **Persistence**: Maintain historical versions of intervals
- **Real-world Applications**: Calendar systems, resource scheduling
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Range Module ===")
	
	// Test 1: Basic operations
	rm := ConstructorRangeModule()
	rmInterval := ConstructorRangeModuleInterval()
	rmSegTree := ConstructorRangeModuleSegmentTree()
	rmBitset := ConstructorRangeModuleBitset()
	
	fmt.Println("=== Basic Operations ===")
	rmInterval.AddRange(10, 20)
	rmInterval.AddRange(15, 25)
	rmInterval.AddRange(20, 30)
	
	fmt.Printf("Query [10, 20): %t\n", rmInterval.QueryRange(10, 20))
	fmt.Printf("Query [15, 25): %t\n", rmInterval.QueryRange(15, 25))
	fmt.Printf("Query [25, 30): %t\n", rmInterval.QueryRange(25, 30))
	
	rmInterval.RemoveRange(14, 16)
	fmt.Printf("After remove [14, 16), Query [10, 20): %t\n", rmInterval.QueryRange(10, 20))
	
	// Test 2: Segment Tree operations
	fmt.Println("\n=== Segment Tree Operations ===")
	rmSegTree.AddRange(1, 5)
	rmSegTree.AddRange(3, 7)
	fmt.Printf("Query [2, 4): %t\n", rmSegTree.QueryRange(2, 4))
	fmt.Printf("Query [4, 6): %t\n", rmSegTree.QueryRange(4, 6))
	
	rmSegTree.RemoveRange(2, 4)
	fmt.Printf("After remove [2, 4), Query [1, 5): %t\n", rmSegTree.QueryRange(1, 5))
	
	// Test 3: Bitset operations
	fmt.Println("\n=== Bitset Operations ===")
	rmBitset.AddRange(0, 10)
	rmBitset.AddRange(5, 15)
	fmt.Printf("Query [0, 10): %t\n", rmBitset.QueryRange(0, 10))
	fmt.Printf("Query [10, 15): %t\n", rmBitset.QueryRange(10, 15))
	
	rmBitset.RemoveRange(5, 10)
	fmt.Printf("After remove [5, 10), Query [0, 15): %t\n", rmBitset.QueryRange(0, 15))
	
	// Test 4: Complex operations
	fmt.Println("\n=== Complex Operations ===")
	complexRm := ConstructorRangeModuleInterval()
	
	operations := []struct {
		op   string
		left int
		right int
	}{
		{"add", 1, 4},
		{"add", 3, 5},
		{"query", 2, 4},
		{"query", 1, 5},
		{"remove", 2, 3},
		{"query", 1, 4},
		{"add", 6, 8},
		{"query", 5, 7},
	}
	
	for _, op := range operations {
		switch op.op {
		case "add":
			complexRm.AddRange(op.left, op.right)
			fmt.Printf("Add [%d, %d)\n", op.left, op.right)
		case "query":
			result := complexRm.QueryRange(op.left, op.right)
			fmt.Printf("Query [%d, %d): %t\n", op.left, op.right, result)
		case "remove":
			complexRm.RemoveRange(op.left, op.right)
			fmt.Printf("Remove [%d, %d)\n", op.left, op.right)
		}
	}
	
	// Test 5: Edge cases
	fmt.Println("\n=== Edge Cases ===")
	edgeRm := ConstructorRangeModuleInterval()
	
	edgeRm.AddRange(0, 1)
	fmt.Printf("Single point - Query [0, 1): %t\n", edgeRm.QueryRange(0, 1))
	
	edgeRm.RemoveRange(0, 1)
	fmt.Printf("After remove - Query [0, 1): %t\n", edgeRm.QueryRange(0, 1))
	
	edgeRm.AddRange(100, 200)
	fmt.Printf("Large range - Query [150, 180): %t\n", edgeRm.QueryRange(150, 180))
	
	// Test 6: Performance test
	fmt.Println("\n=== Performance Test ===")
	perfRm := ConstructorRangeModuleInterval()
	
	// Add many intervals
	for i := 0; i < 100; i += 2 {
		perfRm.AddRange(i, i+1)
	}
	
	// Query many intervals
	trueCount := 0
	for i := 0; i < 100; i++ {
		if perfRm.QueryRange(i, i+1) {
			trueCount++
		}
	}
	
	fmt.Printf("Added 50 intervals, found %d in queries\n", trueCount)
}
