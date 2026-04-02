package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 102. Binary Tree Level Order Traversal
// Time: O(N), Space: O(N)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	
	var result [][]int
	queue := []*TreeNode{root}
	
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]int, 0, levelSize)
		
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)
			
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		
		result = append(result, currentLevel)
	}
	
	return result
}

// Recursive approach
func levelOrderRecursive(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	
	var result [][]int
	
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		
		// Ensure result has enough levels
		for len(result) <= level {
			result = append(result, []int{})
		}
		
		result[level] = append(result[level], node.Val)
		
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	
	dfs(root, 0)
	return result
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

func main() {
	// Test cases
	testCases := [][]interface{}{
		{3, 9, 20, nil, nil, 15, 7},
		{1},
		{},
		{1, 2, 3, 4, 5, 6, 7},
		{1, nil, 2, nil, 3, nil, 4, nil, 5},
		{1, 2, 3, 4, nil, nil, nil, 5},
		{1, 2, 3, nil, nil, nil, 4, nil, nil, 5},
		{1, 2, nil, 3, nil, 4, nil, 5},
	}
	
	for i, nums := range testCases {
		root := createTree(nums)
		result1 := levelOrder(root)
		result2 := levelOrderRecursive(root)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Iterative: %v\n", result1)
		fmt.Printf("  Recursive: %v\n\n", result2)
	}
}
