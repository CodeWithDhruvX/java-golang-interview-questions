package main

import (
	"fmt"
	"math"
)

// Pattern: DP on Trees
// Difficulty: Hard
// Key Concept: Computing values for subtrees (Post-order traversal) and combining them to form the answer for the root.

/*
INTUITION:
In a tree, many problems can be solved by deciding: "If I knew the answer for my left child and right child, could I compute the answer for myself?"

Example: Diameter of Tree, Max Path Sum.

For "Max Path Sum":
A "path" can go from any node to any node.
For a specific node 'Root', the max path passing THROUGH it could be:
1. MaxLen(LeftChild) + Root + MaxLen(RightChild) (The path turns here)
2. MaxLen(LeftChild) + Root (Path continues up to parent)
3. MaxLen(RightChild) + Root (Path continues up to parent)
4. Just Root (If children are negative)

Key distinction:
- When *returning* to the parent, we can only pick ONE branch (left or right). We can't split.
- When *calculating* the global maximum, we can connect Left+Root+Right (forming a ^ shape).

PROBLEM:
LeetCode 124. Binary Tree Maximum Path Sum.
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

ALGORITHM:
1. Use a global variable `globalMax` initialized to min integer.
2. Define a recursive function `dfs(node)`:
   - Base case: If node is nil, return 0.
   - Recursive step:
     - leftMax = max(dfs(node.left), 0). (Ignore negative paths).
     - rightMax = max(dfs(node.right), 0).
   - Update global maximum: `globalMax = max(globalMax, leftMax + rightMax + node.val)` (This represents the ^ path).
   - Return value to parent: `node.val + max(leftMax, rightMax)` (Can only extend one way).

COMPLEXITY:
Time: O(N) - Visit every node once.
Space: O(H) - Recursion stack height.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var globalMax int

func maxPathSum(root *TreeNode) int {
	globalMax = math.MinInt32
	dfs(root)
	return globalMax
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// Calculate max path sum starting from left child
	// If it's negative, we ignore it (take 0).
	left := max(dfs(node.Left), 0)

	// Calculate max path sum starting from right child
	right := max(dfs(node.Right), 0)

	// Update global max (considering the path that TURNS at this node)
	currentPathSum := node.Val + left + right
	if currentPathSum > globalMax {
		globalMax = currentPathSum
	}

	// Return the max path sum that can be EXTENDED to the parent
	return node.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	/*
	    -10
	    /  \
	   9   20
	      /  \
	     15   7
	*/
	// Max path: 15 -> 20 -> 7 (Sum: 42)

	root := &TreeNode{Val: -10}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Printf("Max Path Sum: %d\n", maxPathSum(root))

	/*
	   1
	  / \
	 2   3
	*/
	// Max path: 2 -> 1 -> 3 (Sum: 6)
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 3}
	fmt.Printf("Max Path Sum: %d\n", maxPathSum(root2))
}
