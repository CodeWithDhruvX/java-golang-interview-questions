package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 112. Path Sum
// Time: O(N), Space: O(H) where H is tree height (recursion stack)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	
	// If we reach a leaf node
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	
	// Recursively check left and right subtrees
	remainingSum := targetSum - root.Val
	return hasPathSum(root.Left, remainingSum) || hasPathSum(root.Right, remainingSum)
}

// Iterative approach using stack
func hasPathSumIterative(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	
	stack := []struct {
		node *TreeNode
		sum  int
	}{{root, root.Val}}
	
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		node := current.node
		sum := current.sum
		
		// If we reach a leaf node
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				return true
			}
		}
		
		if node.Right != nil {
			stack = append(stack, struct {
				node *TreeNode
				sum  int
			}{node.Right, sum + node.Right.Val})
		}
		
		if node.Left != nil {
			stack = append(stack, struct {
				node *TreeNode
				sum  int
			}{node.Left, sum + node.Left.Val})
		}
	}
	
	return false
}

// Helper function to create a binary tree from slice
func createTree(nums []interface{}) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	
	nodes := make([]*TreeNode, len(nums))
	for i, val := range nums {
		if val == nil {
			nodes[i] = nil
		} else {
			nodes[i] = &TreeNode{Val: val.(int)}
		}
	}
	
	for i := 0; i < len(nums); i++ {
		left := 2*i + 1
		right := 2*i + 2
		
		if left < len(nums) {
			nodes[i].Left = nodes[left]
		}
		if right < len(nums) {
			nodes[i].Right = nodes[right]
		}
	}
	
	return nodes[0]
}

func main() {
	// Test cases
	testCases := []struct {
		tree      []interface{}
		targetSum int
	}{
		{[]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}, 22},
		{[]interface{}{1, 2, 3}, 5},
		{[]interface{}{1, 2, 3}, 4},
		{[]interface{}{}, 0},
		{[]interface{}{1}, 1},
		{[]interface{}{1}, 2},
		{[]interface{}{-2, nil, -3}, -5},
		{[]interface{}{1, -2, 3, -4, nil, nil, nil, 5}, -1},
		{[]interface{}{0}, 0},
		{[]interface{}{1, 2}, 3},
	}
	
	for i, tc := range testCases {
		root := createTree(tc.tree)
		result1 := hasPathSum(root, tc.targetSum)
		result2 := hasPathSumIterative(root, tc.targetSum)
		
		fmt.Printf("Test Case %d: tree=%v, target=%d\n", i+1, tc.tree, tc.targetSum)
		fmt.Printf("  Recursive: %t, Iterative: %t\n\n", result1, result2)
	}
}
