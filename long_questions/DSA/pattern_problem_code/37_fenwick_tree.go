package main

import (
	"fmt"
)

// Pattern: Fenwick Tree (Binary Indexed Tree)
// Difficulty: Hard
// Key Concept: A "compressed" version of Segment Tree that supports Prefix Sums and Updates in O(log N) using bitwise manipulation.

/*
INTUITION:
Segment Tree is great but takes 4N space and a lot of code.
Fenwick Tree uses an array of size N+1 and very short code (3-4 lines for update/query).

Key Idea:
Every number can be represented as sum of powers of 2 (Binary).
Index 12 (1100) stores sum of range covered by 100 (4).
Index 12 is responsible for range (8, 12].
Parent of 12 is 12 - 4 = 8.
To get prefix sum(12), we add tree[12] + prefix_sum(8).

Navigation:
- Go TO Parent (Query): `i -= i & (-i)`
- Go TO Next Node (Update): `i += i & (-i)`

PROBLEM:
LeetCode 307. Range Sum Query - Mutable (Optimized with BIT).

ALGORITHM:
1. `tree` array size N+1.
2. `update(index, delta)`:
   - `index++` (1-based).
   - `while index <= n`:
     - `tree[index] += delta`
     - `index += index & (-index)`
3. `query(index)`: (Prefix sum from 0 to index)
   - `index++` (1-based).
   - `res = 0`.
   - `while index > 0`:
     - `res += tree[index]`
     - `index -= index & (-index)`
4. `SumRange(L, R)` = `query(R) - query(L-1)`.
*/

type FenwickTree struct {
	tree []int
	nums []int
	n    int
}

func ConstructorBIT(nums []int) FenwickTree {
	n := len(nums)
	ft := FenwickTree{
		tree: make([]int, n+1),
		nums: make([]int, n), // Keep copy to calculate delta
		n:    n,
	}
	// Initialize tree
	for i, v := range nums {
		ft.updateInternal(i, v)
		ft.nums[i] = v
	}
	return ft
}

func (this *FenwickTree) Update(index int, val int) {
	delta := val - this.nums[index]
	this.nums[index] = val
	this.updateInternal(index, delta)
}

func (this *FenwickTree) updateInternal(index int, delta int) {
	index++ // 1-based indexing
	for index <= this.n {
		this.tree[index] += delta
		index += index & (-index)
	}
}

func (this *FenwickTree) SumRange(left int, right int) int {
	return this.query(right) - this.query(left-1)
}

func (this *FenwickTree) query(index int) int {
	index++ // 1-based
	sum := 0
	for index > 0 {
		sum += this.tree[index]
		index -= index & (-index)
	}
	return sum
}

func main() {
	// [1, 3, 5]
	// BIT stores sums based on powers of 2.
	nums := []int{1, 3, 5}
	obj := ConstructorBIT(nums)

	fmt.Printf("Sum(0, 2): %d\n", obj.SumRange(0, 2)) // 9

	obj.Update(1, 2)                                  // [1, 2, 5]
	fmt.Printf("Sum(0, 2): %d\n", obj.SumRange(0, 2)) // 8
}
