package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 236. Lowest Common Ancestor of a Binary Tree
// Time: O(N), Space: O(H) where H is tree height (recursion stack)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	
	if left != nil && right != nil {
		return root // p and q are in different subtrees
	}
	
	if left != nil {
		return left
	}
	
	return right
}

// Iterative approach using parent pointers
func lowestCommonAncestorIterative(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	
	// Build parent map and depth
	parent := make(map[*TreeNode]*TreeNode)
	depth := make(map[*TreeNode]int)
	
	stack := []*TreeNode{root}
	parent[root] = nil
	depth[root] = 0
	
	// Build parent pointers and depths
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		
		if node.Left != nil {
			parent[node.Left] = node
			depth[node.Left] = depth[node] + 1
			stack = append(stack, node.Left)
		}
		
		if node.Right != nil {
			parent[node.Right] = node
			depth[node.Right] = depth[node] + 1
			stack = append(stack, node.Right)
		}
	}
	
	// Bring both nodes to the same depth
	for depth[p] > depth[q] {
		p = parent[p]
	}
	for depth[q] > depth[p] {
		q = parent[q]
	}
	
	// Move up both nodes until they meet
	for p != q {
		p = parent[p]
		q = parent[q]
	}
	
	return p
}

// Helper function to create a binary tree from slice
func createTree(nums []interface{}) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	
	nodes := make([]*TreeNode, len(nums))
	nodeMap := make(map[int]*TreeNode)
	
	for i, val := range nums {
		if val != nil {
			nodes[i] = &TreeNode{Val: val.(int)}
			nodeMap[val.(int)] = nodes[i]
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
	testCases := []struct {
		tree []interface{}
		p    int
		q    int
	}{
		{[]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, 5, 1},
		{[]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, 5, 4},
		{[]interface{}{1, 2}, 1, 2},
		{[]interface{}{1, 2, nil, nil, 3, 4}, 1, 4},
		{[]interface{}{1, 2, 3, 4, 5}, 2, 4},
		{[]interface{}{1, 2, 3, 4, 5}, 1, 3},
		{[]interface{}{1}, 1, 1},
		{[]interface{}{1, nil, 2}, 1, 2},
		{[]interface{}{37, -34, -48, nil, -100, -100, 48, nil, nil, nil, -54, nil, -71, -22, nil, nil}, -34, -48},
	}
	
	for i, tc := range testCases {
		root := createTree(tc.tree)
		
		var p, q *TreeNode
		if tc.tree != nil {
			// Find nodes p and q
			for _, val := range tc.tree {
				if val != nil {
					if val.(int) == tc.p {
						p = &TreeNode{Val: tc.p}
					}
					if val.(int) == tc.q {
						q = &TreeNode{Val: tc.q}
					}
				}
			}
		}
		
		result1 := lowestCommonAncestor(root, p, q)
		result2 := lowestCommonAncestorIterative(root, p, q)
		
		fmt.Printf("Test Case %d: tree=%v, p=%d, q=%d\n", i+1, tc.tree, tc.p, tc.q)
		if result1 != nil {
			fmt.Printf("  Recursive: %d, Iterative: %d\n", result1.Val, result2.Val)
		} else {
			fmt.Printf("  Recursive: nil, Iterative: nil\n")
		}
		fmt.Println()
	}
}
