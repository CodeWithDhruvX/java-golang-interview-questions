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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Post-Order DP on Trees
- **Bottom-Up Processing**: Process children before parent
- **State Propagation**: Pass computed values up the tree
- **Global Maximum**: Track maximum across all nodes
- **Path Selection**: Choose optimal path at each node

## 2. PROBLEM CHARACTERISTICS
- **Tree Structure**: Binary tree with weighted nodes
- **Path Definition**: Any path from any node to any node
- **Optimization**: Find maximum sum path
- **Substructure**: Optimal substructure property holds

## 3. SIMILAR PROBLEMS
- Diameter of Binary Tree (LeetCode 543) - Longest path in tree
- House Robber III (LeetCode 337) - DP on trees with constraints
- Binary Tree Maximum Path Sum (LeetCode 124) - Same problem
- Balanced Binary Tree (LeetCode 110) - Tree height DP

## 4. KEY OBSERVATIONS
- **Post-Order Natural**: Children must be processed before parent
- **Path Continuity**: Path can continue through at most one child
- **Negative Values**: Can choose to not include negative branches
- **Global vs Local**: Local optimum vs global maximum

## 5. VARIATIONS & EXTENSIONS
- **Path Reconstruction**: Track actual maximum path
- **Path Counting**: Count number of maximum paths
- **Constraints**: Add constraints on node values or path length
- **Multiple Queries**: Handle multiple trees efficiently

## 6. INTERVIEW INSIGHTS
- Always clarify: "Tree size constraints? Negative values? Path definition?"
- Edge cases: empty tree, single node, all negative values
- Time complexity: O(N) for all approaches
- Space complexity: O(N) recursion stack, O(N) for memoization
- Key insight: post-order traversal enables bottom-up DP

## 7. COMMON MISTAKES
- Wrong traversal order (need post-order)
- Not handling negative values correctly
- Missing global maximum tracking
- Incorrect path continuation logic
- Stack overflow for deep trees

## 8. OPTIMIZATION STRATEGIES
- **Post-Order DP**: O(N) time, O(N) space - standard
- **Memoization**: O(N) time, O(N) space - cache results
- **Iterative**: O(N) time, O(N) space - avoid recursion
- **Path Tracking**: O(N) time, O(N) space - reconstruct path

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the most valuable route in a road network:**
- Each intersection is a node with a value (treasure/penalty)
- Roads connect parent and child nodes
- You want the route with maximum total treasure
- At each intersection, you can only continue on one road
- You must visit child intersections before parent (bottom-up)
- Like finding the most profitable delivery route

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary tree with integer values
2. **Goal**: Find maximum sum of any path
3. **Rules**: Path can start/end anywhere, must be connected
4. **Output**: Maximum path sum value

#### Phase 2: Key Insight Recognition
- **"Post-order natural"** → Need child results before parent
- **"Path continuity"** → Path can continue through at most one child
- **"Negative handling"** → Can skip negative branches
- **"Global tracking"** → Maximum might not include root

#### Phase 3: Strategy Development
```
Human thought process:
"I need max path sum in binary tree.
Brute force would try all O(N²) paths.

Post-Order DP Approach:
1. Process children before parent (post-order)
2. For each node, calculate max path through each child
3. Choose max child path (or none if negative)
4. Update global maximum with current node + both children
5. Return max path through one child to parent

This gives O(N) time!"
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return appropriate minimum value
- **Single node**: Return node value
- **All negative**: Return maximum (least negative) value
- **Deep tree**: Handle recursion stack limits

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: Tree: [1,2,3]

Human thinking:
"Post-Order DP Approach:
Process leaf nodes first:
- Node 2 (left leaf): maxPath = 2, globalMax = 2
- Node 3 (right leaf): maxPath = 3, globalMax = 3

Process root node 1:
- Left max path: 2 (from node 2)
- Right max path: 3 (from node 3)
- Max path through one child: 1 + max(2,3) = 4
- Path through both children: 1 + 2 + 3 = 6
- Update globalMax: max(3, 4, 6) = 6
- Return to parent: max path through one child = 4

Final Result: 6 ✓"
```

#### Phase 6: Intuition Validation
- **Why post-order works**: Children must be processed before parent
- **Why choose one child**: Path can't split in both directions
- **Why global tracking**: Maximum path might not go through root
- **Why O(N)**: Each node processed exactly once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not pre-order?"** → Need child results before parent
2. **"Should I include both children?"** → Path can't split
3. **"What about negative values?"** → Can skip negative branches
4. **"Can I optimize further?"** → O(N) is optimal for tree traversal
5. **"What about path reconstruction?"** → Need additional tracking

### Real-World Analogy
**Like finding the most profitable delivery route:**
- Each location is a node with profit/loss value
- Roads connect parent and child locations
- You want the route with maximum total profit
- At each location, you can only continue to one next location
- You must evaluate child locations before parent (bottom-up)
- Like logistics optimization for delivery networks

### Human-Readable Pseudocode
```
function maxPathSum(root):
    if root == null:
        return 0
    
    globalMax = -infinity
    
    postOrder(node):
        if node == null:
            return 0
        
        # Process children first
        leftMax = postOrder(node.left)
        rightMax = postOrder(node.right)
        
        # Max path through current node (can only continue one direction)
        currentMax = node.val + max(leftMax, rightMax, 0)
        
        # Update global maximum (can use both children)
        globalMax = max(globalMax, 
                      node.val + leftMax + rightMax,
                      currentMax)
        
        return currentMax
    
    postOrder(root)
    return globalMax
```

### Execution Visualization

### Example: Tree = [1,2,3]
```
Post-Order DP Process:

Step 1: Process leaf nodes
Node 2 (left leaf):
- leftMax = 0, rightMax = 0
- currentMax = 2 + max(0,0,0) = 2
- globalMax = max(-∞, 2+0+0, 2) = 2
- return 2

Node 3 (right leaf):
- leftMax = 0, rightMax = 0
- currentMax = 3 + max(0,0,0) = 3
- globalMax = max(2, 3+0+0, 3) = 3
- return 3

Step 2: Process root node 1
- leftMax = 2 (from node 2)
- rightMax = 3 (from node 3)
- currentMax = 1 + max(2,3,0) = 4
- globalMax = max(3, 1+2+3, 4) = 6
- return 4

Final Result: 6 ✓
```

### Key Visualization Points:
- **Bottom-Up Processing**: Children before parent
- **Path Selection**: Choose optimal child path
- **Global Tracking**: Maximum might be anywhere
- **Negative Handling**: Skip negative branches

### Tree Traversal Visualization:
```
Processing Order:
    1
   / \
  2   3

Step 1: Process 2 (left leaf)
Step 2: Process 3 (right leaf)  
Step 3: Process 1 (root)

Value Propagation:
Node 2 → return 2
Node 3 → return 3
Node 1 → leftMax=2, rightMax=3, globalMax=6
```

### Time Complexity Breakdown:
- **Post-Order DP**: O(N) time, O(N) space - standard
- **Memoization**: O(N) time, O(N) space - cache results
- **Iterative**: O(N) time, O(N) space - avoid recursion
- **Path Reconstruction**: O(N) time, O(N) space - extra tracking

### Alternative Approaches:

#### 1. Brute Force (O(N²) time, O(N) space)
```go
func maxPathSumBruteForce(root *TreeNode) int {
    // Try all possible paths
    // For each node as start, find max path ending there
    // O(N²) paths in worst case
    // ... implementation details omitted
}
```
- **Pros**: Simple to understand
- **Cons**: Too slow for large trees

#### 2. BFS with DP (O(N) time, O(N) space)
```go
func maxPathSumBFS(root *TreeNode) int {
    // Level-order traversal with DP
    // More complex but avoids recursion
    // ... implementation details omitted
}
```
- **Pros**: No recursion stack issues
- **Cons**: More complex implementation

#### 3. Divide and Conquer (O(N) time, O(N) space)
```go
func maxPathSumDivideConquer(root *TreeNode) int {
    // Divide tree into subtrees, combine results
    // Similar to post-order but different perspective
    // ... implementation details omitted
}
```
- **Pros**: Natural recursive formulation
- **Cons**: Same as post-order in complexity

### Extensions for Interviews:
- **Path Reconstruction**: Track actual maximum path nodes
- **Path Counting**: Count number of maximum sum paths
- **Constraints**: Add constraints on path length or node values
- **Multiple Trees**: Process multiple trees efficiently
- **Variations**: Handle different path definitions (must include root, etc.)
*/
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
