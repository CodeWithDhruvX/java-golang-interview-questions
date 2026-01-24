package main

import (
	"fmt"
)

// Pattern: Segment Tree
// Difficulty: Hard
// Key Concept: Binary tree where each node represents an interval. Used for Range Queries (Sum, Max, Min) with Updates in O(log N).

/*
INTUITION:
We have an array [1, 3, 5, 7, 9, 11].
We want to:
1. Find sum of specific range (e.g., index 2 to 4).
2. Update a value at specific index 3.

Prefix Sum array handles (1) fast, but (2) is O(N).
Simple Array handles (2) fast, but (1) is O(N).

Segment Tree balances them: Both are O(log N).
Root represents sum(0..5).
Left Child: sum(0..2). Right Child: sum(3..5).
... down to leaves.

When we update index 3:
We change the leaf. Then we walk up the tree updating parents. O(log N).
When we query [2, 4]:
We combine values from nodes that structurally cover [2, 4] (e.g., node for [2,2] + node for [3,4]). O(log N).

PROBLEM:
LeetCode 307. Range Sum Query - Mutable.
Given an integer array nums, handle multiple queries of the following types:
1. Update: Update the value of nums[index] to be val.
2. SumRange: Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.

ALGORITHM:
1. Use array `tree` of size 4*N.
2. `Build(node, start, end)`: recursion.
   - `mid = (start + end) / 2`.
   - Build Left, Build Right.
   - `tree[node] = tree[2*node] + tree[2*node+1]`.
3. `Update(node, start, end, idx, val)`:
   - If leaf: update `tree[node]`.
   - Else: recurse left or right. Update current node sum.
4. `Query(node, start, end, L, R)`:
   - If range matches exactly `[start, end] == [L, R]`, return `tree[node]`.
   - Else: Split query. Return `Query(Left) + Query(Right)`.
*/

type NumArray struct {
	tree []int
	nums []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	if n == 0 {
		return NumArray{}
	}
	tree := make([]int, 4*n)
	na := NumArray{tree: tree, nums: nums, n: n}
	na.build(1, 0, n-1)
	return na
}

func (this *NumArray) build(node, start, end int) {
	if start == end {
		this.tree[node] = this.nums[start]
		return
	}
	mid := (start + end) / 2
	this.build(2*node, start, mid)
	this.build(2*node+1, mid+1, end)
	this.tree[node] = this.tree[2*node] + this.tree[2*node+1]
}

func (this *NumArray) Update(index int, val int) {
	if this.n == 0 {
		return
	}
	this.update(1, 0, this.n-1, index, val)
}

func (this *NumArray) update(node, start, end, idx, val int) {
	if start == end {
		this.tree[node] = val
		this.nums[idx] = val // Update original too if needed
		return
	}
	mid := (start + end) / 2
	if idx <= mid {
		this.update(2*node, start, mid, idx, val)
	} else {
		this.update(2*node+1, mid+1, end, idx, val)
	}
	this.tree[node] = this.tree[2*node] + this.tree[2*node+1]
}

func (this *NumArray) SumRange(left int, right int) int {
	if this.n == 0 {
		return 0
	}
	return this.query(1, 0, this.n-1, left, right)
}

func (this *NumArray) query(node, start, end, L, R int) int {
	if R < start || L > end {
		return 0 // Out of range
	}
	if L <= start && end <= R {
		return this.tree[node] // Fully inside range
	}
	mid := (start + end) / 2
	leftSum := this.query(2*node, start, mid, L, R)
	rightSum := this.query(2*node+1, mid+1, end, L, R)
	return leftSum + rightSum
}

func main() {
	// Nums: [1, 3, 5]
	// Build:
	//       Root(9) [0-2]
	//      /      \
	//   L(4)[0-1]  R(5)[2-2]
	//   /   \
	// LL(1) LR(3)

	nums := []int{1, 3, 5}
	obj := Constructor(nums)
	fmt.Printf("Initial Sum(0, 2): %d\n", obj.SumRange(0, 2)) // 1+3+5 = 9

	obj.Update(1, 2)                                          // Arr becomes [1, 2, 5]
	fmt.Printf("Updated Sum(0, 2): %d\n", obj.SumRange(0, 2)) // 1+2+5 = 8
}
