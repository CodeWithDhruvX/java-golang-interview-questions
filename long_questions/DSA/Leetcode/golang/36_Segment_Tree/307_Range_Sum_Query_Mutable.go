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
