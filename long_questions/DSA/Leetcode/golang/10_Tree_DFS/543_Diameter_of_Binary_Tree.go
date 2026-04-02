package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 543. Diameter of Binary Tree
// Time: O(N), Space: O(H) where H is tree height (recursion stack)
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		
		leftDepth := dfs(node.Left)
		rightDepth := dfs(node.Right)
		
		// Update max diameter found so far
		currentDiameter := leftDepth + rightDepth
		if currentDiameter > maxDiameter {
			maxDiameter = currentDiameter
		}
		
		// Return the maximum depth from this node
		return max(leftDepth, rightDepth) + 1
	}
	
	dfs(root)
	return maxDiameter
}

// Alternative approach returning both depth and diameter
func diameterOfBinaryTreeAlternative(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		
		leftDepth, leftDiameter := dfs(node.Left)
		rightDepth, rightDiameter := dfs(node.Right)
		
		// Current node diameter
		currentDiameter := leftDepth + rightDepth
		
		// Maximum diameter in subtrees
		maxDiameter := max(max(leftDiameter, rightDiameter), currentDiameter)
		
		// Return depth and max diameter
		return max(leftDepth, rightDepth) + 1, maxDiameter
	}
	
	_, maxDiameter := dfs(root)
	return maxDiameter
}

// Helper function to create a binary tree from slice
func createTree(nums []interface{}) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	
	nodes := make([]*TreeNode, len(nums))
	for i, val := range nums {
		if val != nil {
			nodes[i] = &TreeNode{Val: val.(int)}
		}
	}
	
	for i := 0; i < len(nums); i++ {
		if nums[i] != nil {
			left := 2*i + 1
			right := 2*i + 2
			
			if left < len(nums) && nums[left] != nil {
				nodes[i].Left = nodes[left]
			}
			if right < len(nums) && nums[right] != nil {
				nodes[i].Right = nodes[right]
			}
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
		{1, 2, 3, 4, 5},
		{1, 2},
		{1, 2, nil, nil, 3, 4, nil, nil, nil, 5},
		{4, -7, -3, nil, nil, -9, -3, 9, -7, -4, nil, 6, nil, -6, -6, nil, 0, 6, 5, nil, 9, nil, nil, -1, -4, nil, nil, nil, -2},
		{1},
		{},
		{1, nil, 2, nil, 3, nil, 4, nil, 5},
		{1, 2, 3, 4, nil, nil, nil, 5},
		{1, 2, 3, nil, nil, nil, 4, nil, nil, 5},
	}
	
	for i, nums := range testCases {
		root := createTree(nums)
		result1 := diameterOfBinaryTree(root)
		result2 := diameterOfBinaryTreeAlternative(root)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Method 1: %d, Method 2: %d\n\n", result1, result2)
	}
}
