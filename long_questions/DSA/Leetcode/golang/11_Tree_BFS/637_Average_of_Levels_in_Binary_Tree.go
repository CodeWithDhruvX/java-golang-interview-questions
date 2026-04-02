package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 637. Average of Levels in Binary Tree
// Time: O(N), Space: O(N)
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	
	var result []float64
	queue := []*TreeNode{root}
	
	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0
		
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val
			
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		
		average := float64(levelSum) / float64(levelSize)
		result = append(result, average)
	}
	
	return result
}

// Recursive approach
func averageOfLevelsRecursive(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	
	var result []float64
	
	var dfs func(*TreeNode, int, *[]int)
	dfs = func(node *TreeNode, level int, sums *[]int) {
		if node == nil {
			return
		}
		
		// Ensure sums slice has enough elements
		for len(*sums) <= level {
			*sums = append(*sums, 0)
		}
		
		(*sums)[level] += node.Val
		
		dfs(node.Left, level+1, sums)
		dfs(node.Right, level+1, sums)
	}
	
	// First pass: calculate sums at each level
	sums := []int{}
	dfs(root, 0, &sums)
	
	// Second pass: calculate counts at each level
	counts := make([]int, len(sums))
	
	var countNodes func(*TreeNode, int)
	countNodes = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		
		counts[level]++
		countNodes(node.Left, level+1)
		countNodes(node.Right, level+1)
	}
	
	countNodes(root, 0)
	
	// Calculate averages
	for i := 0; i < len(sums); i++ {
		average := float64(sums[i]) / float64(counts[i])
		result = append(result, average)
	}
	
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
		{3, 9, 20, 15, 7},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, 5, nil, nil, nil, 6, 7},
		{1, nil, 2, nil, 3, nil, 4, nil, 5},
		{1, 2, 3, 4, nil, nil, nil, 5},
		{1, 2, 3, nil, nil, nil, 4, nil, nil, 5},
		{1},
		{},
	}
	
	for i, nums := range testCases {
		root := createTree(nums)
		result1 := averageOfLevels(root)
		result2 := averageOfLevelsRecursive(root)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Iterative: [")
		for j, avg := range result1 {
			if j > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%.1f", avg)
		}
		fmt.Printf("]\n")
		
		fmt.Printf("  Recursive: [")
		for j, avg := range result2 {
			if j > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%.1f", avg)
		}
		fmt.Printf("]\n\n")
	}
}
