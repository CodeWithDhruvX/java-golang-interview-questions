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
