package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 230. Kth Smallest Element in a BST
// Time: O(H + K) where H is tree height, Space: O(H)
func kthSmallest(root *TreeNode, k int) int {
	var result int
	count := 0
	
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}
		
		// Traverse left subtree
		inorder(node.Left)
		
		// Process current node
		count++
		if count == k {
			result = node.Val
			return
		}
		
		// Traverse right subtree
		inorder(node.Right)
	}
	
	inorder(root)
	return result
}

// Iterative approach using stack
func kthSmallestIterative(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	
	stack := []*TreeNode{}
	current := root
	count := 0
	
	for len(stack) > 0 || current != nil {
		// Go as far left as possible
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		
		// Process node
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		count++
		if count == k {
			return current.Val
		}
		
		// Move to right subtree
		current = current.Right
	}
	
	return -1 // k is larger than number of nodes
}

// Optimized approach with early termination
func kthSmallestOptimized(root *TreeNode, k int) int {
	var result int
	var count int
	
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}
		
		// Check if we can skip entire left subtree
		leftSize := countNodes(node.Left)
		if count+leftSize+1 >= k {
			// The kth element is in this subtree
			if count+leftSize+1 == k {
				result = node.Val
				count = k // Stop further processing
				return
			}
			
			// Search in left subtree
			inorder(node.Left)
			if count >= k {
				return
			}
			
			// Process current node
			count++
			if count == k {
				result = node.Val
				return
			}
			
			// Search in right subtree
			inorder(node.Right)
		} else {
			// Skip entire left subtree and current node
			count += leftSize + 1
			// Search in right subtree
			inorder(node.Right)
		}
	}
	
	inorder(root)
	return result
}

// Helper function to count nodes in subtree
func countNodes(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + countNodes(node.Left) + countNodes(node.Right)
}

// Approach using reverse inorder for kth largest
func kthLargest(root *TreeNode, k int) int {
	var result int
	count := 0
	
	var reverseInorder func(*TreeNode)
	reverseInorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}
		
		// Traverse right subtree first
		reverseInorder(node.Right)
		
		// Process current node
		count++
		if count == k {
			result = node.Val
			return
		}
		
		// Traverse left subtree
		reverseInorder(node.Left)
	}
	
	reverseInorder(root)
	return result
}

// Approach that returns both kth smallest and largest
func kthSmallestAndLargest(root *TreeNode, k int) (int, int) {
	var smallResult, largeResult int
	var smallCount, largeCount int
	
	// For kth smallest
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || smallCount >= k {
			return
		}
		
		inorder(node.Left)
		smallCount++
		if smallCount == k {
			smallResult = node.Val
			return
		}
		inorder(node.Right)
	}
	
	// For kth largest
	var reverseInorder func(*TreeNode)
	reverseInorder = func(node *TreeNode) {
		if node == nil || largeCount >= k {
			return
		}
		
		reverseInorder(node.Right)
		largeCount++
		if largeCount == k {
			largeResult = node.Val
			return
		}
		reverseInorder(node.Left)
	}
	
	inorder(root)
	reverseInorder(root)
	
	return smallResult, largeResult
}

// Approach using augmented BST (with node counts)
type AugmentedNode struct {
	Val   int
	Left  *AugmentedNode
	Right *AugmentedNode
	Size  int // Size of subtree including this node
}

func kthSmallestAugmented(root *AugmentedNode, k int) int {
	return kthSmallestAugmentedHelper(root, k)
}

func kthSmallestAugmentedHelper(node *AugmentedNode, k int) int {
	if node == nil {
		return -1
	}
	
	leftSize := 0
	if node.Left != nil {
		leftSize = node.Left.Size
	}
	
	if k <= leftSize {
		// kth element is in left subtree
		return kthSmallestAugmentedHelper(node.Left, k)
	} else if k == leftSize+1 {
		// Current node is the kth element
		return node.Val
	} else {
		// kth element is in right subtree
		return kthSmallestAugmentedHelper(node.Right, k-leftSize-1)
	}
}

// Helper function to build augmented BST
func buildAugmentedBST(nums []int) *AugmentedNode {
	if len(nums) == 0 {
		return nil
	}
	
	root := &AugmentedNode{Val: nums[0]}
	for i := 1; i < len(nums); i++ {
		insertAugmentedBST(root, nums[i])
	}
	updateSize(root)
	return root
}

func insertAugmentedBST(root *AugmentedNode, val int) {
	current := root
	for {
		if val < current.Val {
			if current.Left == nil {
				current.Left = &AugmentedNode{Val: val}
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &AugmentedNode{Val: val}
				return
			}
			current = current.Right
		}
	}
}

func updateSize(node *AugmentedNode) {
	if node == nil {
		return
	}
	
	updateSize(node.Left)
	updateSize(node.Right)
	
	leftSize := 0
	rightSize := 0
	if node.Left != nil {
		leftSize = node.Left.Size
	}
	if node.Right != nil {
		rightSize = node.Right.Size
	}
	
	node.Size = 1 + leftSize + rightSize
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
	current := root
	for {
		if val < current.Val {
			if current.Left == nil {
				current.Left = &TreeNode{Val: val}
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &TreeNode{Val: val}
				return
			}
			current = current.Right
		}
	}
}

// Helper function to convert tree to array (in-order traversal)
func treeToArray(root *TreeNode) []int {
	var result []int
	inorderTraversal(root, &result)
	return result
}

func inorderTraversal(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	
	inorderTraversal(node.Left, result)
	*result = append(*result, node.Val)
	inorderTraversal(node.Right, result)
}

func main() {
	// Test cases
	fmt.Println("=== Testing Kth Smallest Element in BST ===")
	
	testCases := []struct {
		initialTree []interface{}
		k          int
		description string
	}{
		{[]interface{}{3, 1, 4, nil, 2}, 1, "Find smallest element"},
		{[]interface{}{3, 1, 4, nil, 2}, 2, "Find second smallest"},
		{[]interface{}{3, 1, 4, nil, 2}, 3, "Find third smallest"},
		{[]interface{}{5, 3, 6, 2, 4, nil, nil, 1}, 4, "Find 4th smallest"},
		{[]interface{}{5, 3, 6, 2, 4, nil, nil, 1}, 7, "Find largest element"},
		{[]interface{}{}, 1, "Empty tree"},
		{[]interface{}{10}, 1, "Single node"},
		{[]interface{}{10, 5, 15, 2, 7, 12, 18}, 5, "Medium tree"},
		{[]interface{}{10, 5, 15, 2, 7, 12, 18}, 3, "Find third in medium tree"},
	}
	
	for i, tc := range testCases {
		root := createTree(tc.initialTree)
		
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Tree: %v\n", tc.initialTree)
		fmt.Printf("  k = %d\n", tc.k)
		
		result1 := kthSmallest(root, tc.k)
		result2 := kthSmallestIterative(root, tc.k)
		result3 := kthSmallestOptimized(root, tc.k)
		
		fmt.Printf("  Recursive: %d\n", result1)
		fmt.Printf("  Iterative: %d\n", result2)
		fmt.Printf("  Optimized: %d\n\n", result3)
	}
	
	// Test kth largest
	fmt.Println("=== Testing Kth Largest Element ===")
	largeRoot := createBST([]int{10, 5, 15, 2, 7, 12, 18, 1, 3, 6, 8, 11, 13, 16, 20})
	fmt.Printf("BST: %v\n", treeToArray(largeRoot))
	
	for k := 1; k <= 5; k++ {
		largest := kthLargest(largeRoot, k)
		fmt.Printf("kth largest (k=%d): %d\n", k, largest)
	}
	
	// Test both kth smallest and largest
	fmt.Println("\n=== Testing Both Kth Smallest and Largest ===")
	for k := 1; k <= 3; k++ {
		small, large := kthSmallestAndLargest(largeRoot, k)
		fmt.Printf("k=%d: smallest=%d, largest=%d\n", k, small, large)
	}
	
	// Test augmented BST
	fmt.Println("\n=== Testing Augmented BST ===")
	augmentedRoot := buildAugmentedBST([]int{10, 5, 15, 2, 7, 12, 18, 1, 3, 6, 8, 11, 13, 16, 20})
	
	for k := 1; k <= 5; k++ {
		result := kthSmallestAugmented(augmentedRoot, k)
		fmt.Printf("Augmented BST kth smallest (k=%d): %d\n", k, result)
	}
	
	// Test edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	
	// Very large k
	singleRoot := &TreeNode{Val: 10}
	fmt.Printf("Single node, k=5: %d\n", kthSmallest(singleRoot, 5))
	
	// Large tree
	largeBST := createBST([]int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45, 55, 65, 75, 85})
	fmt.Printf("Large tree, k=8: %d\n", kthSmallest(largeBST, 8))
	
	// Test with negative values
	negRoot := createBST([]int{0, -10, 10, -20, -5, 5, 15})
	fmt.Printf("Tree with negatives: %v\n", treeToArray(negRoot))
	fmt.Printf("kth smallest (k=3): %d\n", kthSmallest(negRoot, 3))
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	perfRoot := createBST([]int{100, 50, 150, 25, 75, 125, 175, 12, 37, 62, 87, 112, 137, 162, 187})
	
	start := 0
	for k := 1; k <= 20; k++ {
		start += kthSmallest(perfRoot, k)
	}
	
	fmt.Printf("Performance test completed with 20 queries\n")
}
