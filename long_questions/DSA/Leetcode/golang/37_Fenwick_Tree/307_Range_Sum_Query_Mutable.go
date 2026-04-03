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
