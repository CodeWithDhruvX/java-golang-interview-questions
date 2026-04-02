package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 98. Validate Binary Search Tree
// Time: O(N), Space: O(H) where H is tree height (recursion stack)
func isValidBST(root *TreeNode) bool {
	return validateBST(root, nil, nil)
}

func validateBST(node *TreeNode, min, max *TreeNode) bool {
	if node == nil {
		return true
	}
	
	// Check current node value against bounds
	if min != nil && node.Val <= min.Val {
		return false
	}
	if max != nil && node.Val >= max.Val {
		return false
	}
	
	// Recursively check left and right subtrees
	return validateBST(node.Left, min, node) && validateBST(node.Right, node, max)
}

// Alternative approach using in-order traversal
func isValidBSTInorder(root *TreeNode) bool {
	var prev *TreeNode
	
	var inorder func(*TreeNode) bool
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		
		// Check left subtree
		if !inorder(node.Left) {
			return false
		}
		
		// Check current node
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		
		// Check right subtree
		return inorder(node.Right)
	}
	
	return inorder(root)
}

// Iterative approach using stack
func isValidBSTIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}
	
	stack := []*TreeNode{}
	var prev *TreeNode
	current := root
	
	for len(stack) > 0 || current != nil {
		// Go as far left as possible
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		
		// Process node
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if prev != nil && current.Val <= prev.Val {
			return false
		}
		prev = current
		
		// Move to right subtree
		current = current.Right
	}
	
	return true
}

// Range-based validation
func isValidBSTRange(root *TreeNode) bool {
	return isValidBSTRangeHelper(root, -2147483648, 2147483647) // 32-bit int range
}

func isValidBSTRangeHelper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	
	if node.Val <= min || node.Val >= max {
		return false
	}
	
	return isValidBSTRangeHelper(node.Left, min, node.Val) &&
		isValidBSTRangeHelper(node.Right, node.Val, max)
}

// Helper function to create BST from array
func createBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	
	root := &TreeNode{Val: nums[0]}
	for i := 1; i < len(nums); i++ {
		insertBST(root, nums[i])
	}
	return root
}

func insertBST(root *TreeNode, val int) {
	if root == nil {
		return
	}
	
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			insertBST(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			insertBST(root.Right, val)
		}
	}
}

// Helper function to create tree from array (not necessarily BST)
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
	fmt.Println("=== Testing Validate Binary Search Tree ===")
	
	testCases := []struct {
		tree       []interface{}
		isBST      bool
		description string
	}{
		{[]interface{}{2, 1, 3}, true, "Valid BST"},
		{[]interface{}{5, 1, 4, nil, nil, 3, 6}, false, "Invalid BST"},
		{[]interface{}{1, 1}, false, "Duplicate values"},
		{[]interface{}{}, true, "Empty tree"},
		{[]interface{}{1}, true, "Single node"},
		{[]interface{}{10, 5, 15, 3, 7, 12, 18}, true, "Perfect BST"},
		{[]interface{}{10, 5, 15, 3, 12, 12, 18}, false, "Right subtree violation"},
		{[]interface{}{10, 5, 15, 1, 6, 12, 20}, true, "Valid BST with varied depths"},
		{[]interface{}{10, 5, 15, 1, 11, 12, 20}, false, "Left subtree violation"},
		{[]interface{}{2, 2, 2}, false, "All same values"},
	}
	
	for i, tc := range testCases {
		root := createTree(tc.tree)
		
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Tree: %v\n", tc.tree)
		
		result1 := isValidBST(root)
		result2 := isValidBSTInorder(root)
		result3 := isValidBSTIterative(root)
		result4 := isValidBSTRange(root)
		
		fmt.Printf("  Recursive: %t\n", result1)
		fmt.Printf("  In-order: %t\n", result2)
		fmt.Printf("  Iterative: %t\n", result3)
		fmt.Printf("  Range-based: %t\n\n", result4)
	}
	
	// Test with manually created BST
	fmt.Println("=== Testing with Created BST ===")
	bst := createBST([]int{50, 30, 70, 20, 40, 60, 80})
	fmt.Printf("Created BST - Valid: %t\n", isValidBST(bst))
	
	// Test with invalid tree
	invalidTree := &TreeNode{
		Val: 10,
		Left: &TreeNode{Val: 5},
		Right: &TreeNode{
			Val: 15,
			Left:  &TreeNode{Val: 6}, // This violates BST property
			Right: &TreeNode{Val: 20},
		},
	}
	fmt.Printf("Invalid tree - Valid: %t\n", isValidBST(invalidTree))
	
	// Test edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	
	// Large BST
	largeBST := createBST([]int{50, 25, 75, 12, 37, 62, 87, 6, 18, 31, 43, 56, 68, 81, 93})
	fmt.Printf("Large BST - Valid: %t\n", isValidBST(largeBST))
	
	// Tree with negative values
	negBST := createBST([]int{0, -10, 10, -20, -5, 5, 20})
	fmt.Printf("BST with negatives - Valid: %t\n", isValidBST(negBST))
	
	// Tree with maximum int values
	maxBST := createBST([]int{1000000, 500000, 1500000})
	fmt.Printf("BST with large values - Valid: %t\n", isValidBST(maxBST))
}
