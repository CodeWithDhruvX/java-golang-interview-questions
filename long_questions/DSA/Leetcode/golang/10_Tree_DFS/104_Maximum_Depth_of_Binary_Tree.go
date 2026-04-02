package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 104. Maximum Depth of Binary Tree
// Time: O(N), Space: O(H) where H is tree height (recursion stack)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	
	return max(leftDepth, rightDepth) + 1
}

// Iterative BFS approach
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	queue := []*TreeNode{root}
	depth := 0
	
	for len(queue) > 0 {
		levelSize := len(queue)
		depth++
		
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	
	return depth
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test cases
	testCases := [][]interface{}{
		{3, 9, 20, nil, nil, 15, 7},
		{1, nil, 2},
		{},
		{1},
		{1, 2, 3, 4, 5, nil, nil, 6, 7, nil, nil, nil, nil, 8},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, nil, nil, nil, 5},
		{1, nil, 2, nil, 3, nil, 4, nil, 5},
		{0},
		{1, 2},
	}
	
	for i, nums := range testCases {
		root := createTree(nums)
		result1 := maxDepth(root)
		result2 := maxDepthBFS(root)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Recursive: %d, BFS: %d\n\n", result1, result2)
	}
}
