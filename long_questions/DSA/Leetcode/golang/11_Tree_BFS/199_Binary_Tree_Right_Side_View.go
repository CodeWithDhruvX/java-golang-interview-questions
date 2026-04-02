package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 199. Binary Tree Right Side View
// Time: O(N), Space: O(N)
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	
	var result []int
	queue := []*TreeNode{root}
	
	for len(queue) > 0 {
		levelSize := len(queue)
		
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			
			// Add children to queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			
			// The last node at each level is visible from the right side
			if i == levelSize-1 {
				result = append(result, node.Val)
			}
		}
	}
	
	return result
}

// Recursive approach with depth tracking
func rightSideViewRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	
	var result []int
	
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		
		// If this is the first node we've seen at this depth
		if len(result) == depth {
			result = append(result, node.Val)
		}
		
		// Traverse right subtree first, then left
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
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
		{1, 2, 3, nil, 5, nil, 4},
		{1, 2, 3, 4, 5, nil, nil, nil, nil, nil, nil, nil, 6, 7},
		{1, nil, 3},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, nil, 3, nil, 4, nil, 5},
		{1, 2, 3, nil, nil, nil, 4, nil, nil, 5},
		{1},
		{},
		{1, 2, 3, 4, nil, nil, nil, 5},
	}
	
	for i, nums := range testCases {
		root := createTree(nums)
		result1 := rightSideView(root)
		result2 := rightSideViewRecursive(root)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Iterative: %v\n", result1)
		fmt.Printf("  Recursive: %v\n\n", result2)
	}
}
