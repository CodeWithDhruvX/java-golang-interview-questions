package main

import (
	"fmt"
	"math"
)

// 124. Binary Tree Maximum Path Sum - DP on Trees
// Time: O(N), Space: O(N) for recursion stack
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return math.MinInt32
	}
	
	maxSum := math.MinInt32
	
	// Post-order traversal
	postOrder(root, &maxSum)
	
	return maxSum
}

func postOrder(node *TreeNode, maxSum *int) {
	if node == nil {
		return
	}
	
	// Traverse left and right
	postOrder(node.Left, maxSum)
	postOrder(node.Right, maxSum)
	
	// Calculate max path sum including current node
	leftMax := 0
	if node.Left != nil {
		leftMax = node.Left.Val
	}
	
	rightMax := 0
	if node.Right != nil {
		rightMax = node.Right.Val
	}
	
	// Current node value plus max of left/right branches
	currentMax := node.Val + max(leftMax, rightMax)
	
	// Update global maximum
	if currentMax > *maxSum {
		*maxSum = currentMax
	}
	
	// Update current node value to be used by parent
	node.Val = currentMax
}

// DP with memoization
func maxPathSumMemo(root *TreeNode) int {
	if root == nil {
		return math.MinInt32
	}
	
	memo := make(map[*TreeNode]int)
	return maxPathSumHelper(root, memo)
}

func maxPathSumHelper(node *TreeNode, memo map[*TreeNode]int) int {
	if node == nil {
		return 0
	}
	
	if val, exists := memo[node]; exists {
		return val
	}
	
	leftMax := 0
	if node.Left != nil {
		leftMax = maxPathSumHelper(node.Left, memo)
	}
	
	rightMax := 0
	if node.Right != nil {
		rightMax = maxPathSumHelper(node.Right, memo)
	}
	
	result := node.Val + max(leftMax, rightMax)
	memo[node] = result
	
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DP with path reconstruction
func maxPathSumWithPath(root *TreeNode) (int, []int) {
	if root == nil {
		return 0, []int{}
	}
	
	maxSum := math.MinInt32
	path := []int{}
	
	postOrderWithPath(root, &maxSum, &path)
	
	return maxSum, path
}

func postOrderWithPath(node *TreeNode, maxSum *int, path *[]int) {
	if node == nil {
		return
	}
	
	// Traverse children
	postOrderWithPath(node.Left, maxSum, path)
	postOrderWithPath(node.Right, maxSum, path)
	
	// Calculate max path sums
	leftMax := 0
	if node.Left != nil {
		leftMax = node.Left.Val
	}
	
	rightMax := 0
	if node.Right != nil {
		rightMax = node.Right.Val
	}
	
	currentMax := node.Val + max(leftMax, rightMax)
	
	// Update global maximum and path
	if currentMax > *maxSum {
		*maxSum = currentMax
		*path = []int{node.Val}
		if leftMax > rightMax && node.Left != nil {
			*path = append(*path, node.Left.Val)
		}
		if rightMax > leftMax && node.Right != nil {
			*path = append(*path, node.Right.Val)
		}
	}
	
	// Update current node value
	node.Val = currentMax
}

// DP with multiple test cases
func maxPathSumMultiple(roots []*TreeNode) []int {
	results := make([]int, len(roots))
	
	for i, root := range roots {
		// Create deep copy to avoid modification
		rootCopy := copyTree(root)
		results[i] = maxPathSum(rootCopy)
	}
	
	return results
}

func copyTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	
	return &TreeNode{
		Val:   node.Val,
		Left:  copyTree(node.Left),
		Right: copyTree(node.Right),
	}
}

// DP with iterative approach
func maxPathSumIterative(root *TreeNode) int {
	if root == nil {
		return math.MinInt32
	}
	
	// Use stack for iterative post-order traversal
	stack := []*TreeNode{root}
	nodeValues := make(map[*TreeNode]int)
	
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		// Push children
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		
		// Process node when both children are processed
		if (node.Right == nil || nodeValues[node.Right] != 0) && 
		   (node.Left == nil || nodeValues[node.Left] != 0) {
			
			leftMax := 0
			if node.Left != nil {
				leftMax = nodeValues[node.Left]
			}
			
			rightMax := 0
			if node.Right != nil {
				rightMax = nodeValues[node.Right]
			}
			
			nodeValues[node] = node.Val + max(leftMax, rightMax)
		}
	}
	
	// Find maximum
	maxSum := math.MinInt32
	for _, val := range nodeValues {
		if val > maxSum {
			maxSum = val
		}
	}
	
	return maxSum
}

// DP with path count
func maxPathSumWithCount(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	
	maxSum := math.MinInt32
	pathCount := 0
	
	postOrderWithCount(root, &maxSum, &pathCount)
	
	return maxSum, pathCount
}

func postOrderWithCount(node *TreeNode, maxSum *int, pathCount *int) {
	if node == nil {
		return
	}
	
	postOrderWithCount(node.Left, maxSum, pathCount)
	postOrderWithCount(node.Right, maxSum, pathCount)
	
	leftMax := 0
	if node.Left != nil {
		leftMax = node.Left.Val
	}
	
	rightMax := 0
	if node.Right != nil {
		rightMax = node.Right.Val
	}
	
	currentMax := node.Val + max(leftMax, rightMax)
	
	if currentMax > *maxSum {
		*maxSum = currentMax
		*pathCount = 1
	} else if currentMax == *maxSum {
		*pathCount++
	}
	
	node.Val = currentMax
}

// DP with constraint
func maxPathSumWithConstraint(root *TreeNode, maxNodeValue int) int {
	if root == nil {
		return 0
	}
	
	maxSum := math.MinInt32
	
	postOrderWithConstraint(root, &maxSum, maxNodeValue)
	
	return maxSum
}

func postOrderWithConstraint(node *TreeNode, maxSum *int, maxNodeValue int) {
	if node == nil {
		return
	}
	
	postOrderWithConstraint(node.Left, maxSum, maxNodeValue)
	postOrderWithConstraint(node.Right, maxSum, maxNodeValue)
	
	leftMax := 0
	if node.Left != nil {
		leftMax = node.Left.Val
	}
	
	rightMax := 0
	if node.Right != nil {
		rightMax = node.Right.Val
	}
	
	currentMax := node.Val + max(leftMax, rightMax)
	
	// Apply constraint
	if node.Val > maxNodeValue {
		currentMax = maxNodeValue + max(leftMax, rightMax)
	}
	
	if currentMax > *maxSum {
		*maxSum = currentMax
	}
	
	node.Val = currentMax
}

// DP with different traversal orders
func maxPathSumDifferentOrders(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 0
	}
	
	// Post-order (bottom-up)
	postOrderRoot := copyTree(root)
	postOrderSum := maxPathSum(postOrderRoot)
	
	// Pre-order (top-down)
	preOrderRoot := copyTree(root)
	preOrderSum := maxPathSumPreOrder(preOrderRoot)
	
	// In-order
	inOrderRoot := copyTree(root)
	inOrderSum := maxPathSumInOrder(inOrderRoot)
	
	return postOrderSum, preOrderSum, inOrderSum
}

func maxPathSumPreOrder(node *TreeNode) int {
	if node == nil {
		return 0
	}
	
	leftMax := maxPathSumPreOrder(node.Left)
	rightMax := maxPathSumPreOrder(node.Right)
	
	return node.Val + max(leftMax, rightMax)
}

func maxPathSumInOrder(node *TreeNode) int {
	if node == nil {
		return 0
	}
	
	leftMax := maxPathSumInOrder(node.Left)
	rightMax := maxPathSumInOrder(node.Right)
	
	return node.Val + max(leftMax, rightMax)
}

func main() {
	// Test cases
	fmt.Println("=== Testing Binary Tree Maximum Path Sum - DP on Trees ===")
	
	// Helper function to create test trees
	createTree := func(vals []interface{}) *TreeNode {
		if len(vals) == 0 {
			return nil
		}
		
		nodes := make([]*TreeNode, len(vals))
		for i, val := range vals {
			if val != nil {
				nodes[i] = &TreeNode{Val: val.(int)}
			}
		}
		
		for i := 0; i < len(vals); i++ {
			left := 2*i + 1
			right := 2*i + 2
			
			if left < len(vals) && vals[left] != nil {
				nodes[i].Left = nodes[left]
			}
			if right < len(vals) && vals[right] != nil {
				nodes[i].Right = nodes[right]
			}
		}
		
		return nodes[0]
	}
	
	testCases := []struct {
		tree       []interface{}
		description string
	}{
		{[]interface{}{1, 2, 3}, "Simple tree"},
		{[]interface{}{-10, 9, 20, nil, nil, 15, 7}, "Complex tree"},
		{[]interface{}{5, 4, 8, 11, nil, 13, 4, 1, 6, nil, 9}, "Large tree"},
		{[]interface{}{1}, "Single node"},
		{[]interface{}{}, "Empty tree"},
		{[]interface{}{2, 1, 3, 4, 5}, "Increasing values"},
		{[]interface{}{5, 4, 3, 2, 1}, "Decreasing values"},
		{[]interface{}{0, -1, -2, -3, -4}, "Negative values"},
		{[]interface{}{100, 200, 300, 400}, "Large values"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Tree: %v\n", tc.tree)
		
		root := createTree(tc.tree)
		
		result1 := maxPathSum(root)
		result2 := maxPathSumMemo(root)
		result3 := maxPathSumIterative(root)
		
		fmt.Printf("  Post-order DP: %d\n", result1)
		fmt.Printf("  Memoized DP: %d\n", result2)
		fmt.Printf("  Iterative DP: %d\n", result3)
		
		// Test path reconstruction
		maxSum, path := maxPathSumWithPath(root)
		fmt.Printf("  Max sum with path: %d, path: %v\n", maxSum, path)
		
		// Test different orders
		postSum, preSum, inSum := maxPathSumDifferentOrders(root)
		fmt.Printf("  Post-order: %d, Pre-order: %d, In-order: %d\n", postSum, preSum, inSum)
		
		fmt.Println()
	}
	
	// Test multiple trees
	fmt.Println("=== Multiple Trees Test ===")
	roots := []*TreeNode{
		createTree([]interface{}{1, 2, 3}),
		createTree([]interface{}{4, 5, 6}),
		createTree([]interface{}{7, 8, 9}),
	}
	
	results := maxPathSumMultiple(roots)
	for i, result := range results {
		fmt.Printf("Tree %d max path sum: %d\n", i+1, result)
	}
	
	// Test path count
	fmt.Println("\n=== Path Count Test ===")
	testRoot := createTree([]interface{}{1, 2, 3, 4, 5})
	maxSum, pathCount := maxPathSumWithCount(testRoot)
	fmt.Printf("Max sum: %d, Path count: %d\n", maxSum, pathCount)
	
	// Test constraint
	fmt.Println("\n=== Constraint Test ===")
	constrainedRoot := createTree([]interface{}{100, 200, 300, 400})
	constrainedResult := maxPathSumWithConstraint(constrainedRoot, 500)
	fmt.Printf("Constrained result: %d\n", constrainedResult)
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	// Create large tree
	largeVals := make([]interface{}, 1000)
	for i := 0; i < 1000; i++ {
		largeVals[i] = i % 1000
	}
	
	largeRoot := createTree(largeVals)
	fmt.Printf("Large tree with %d nodes\n", len(largeVals))
	
	result := maxPathSum(largeRoot)
	fmt.Printf("Large tree result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// All negative values
	allNegative := createTree([]interface{}{-1, -2, -3, -4, -5})
	fmt.Printf("All negative: %d\n", maxPathSum(allNegative))
	
	// Mixed positive and negative
	mixed := createTree([]interface{}{-10, 5, -3, 8, -2, 7})
	fmt.Printf("Mixed values: %d\n", maxPathSum(mixed))
	
	// Single large value
	singleLarge := createTree([]interface{}{1000000})
	fmt.Printf("Single large value: %d\n", maxPathSum(singleLarge))
	
	// Deep tree
	deepRoot := &TreeNode{Val: 1}
	deepRoot.Left = &TreeNode{Val: 2}
	deepRoot.Right = &TreeNode{Val: 3}
	deepRoot.Left.Left = &TreeNode{Val: 4}
	deepRoot.Left.Right = &TreeNode{Val: 5}
	deepRoot.Right.Left = &TreeNode{Val: 6}
	deepRoot.Right.Right = &TreeNode{Val: 7}
	deepRoot.Left.Left.Left = &TreeNode{Val: 8}
	deepRoot.Left.Left.Right = &TreeNode{Val: 9}
	
	fmt.Printf("Deep tree: %d\n", maxPathSum(deepRoot))
}
