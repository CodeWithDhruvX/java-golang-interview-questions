package main

import (
	"fmt"
	"math"
)

// Pattern: Tree DFS (Depth First Search)
// Difficulty: Easy
// Key Concept: Exploring a tree by going as "deep" as possible before backtracking.

/*
INTUITION:
How do you find the depth of a deep hole (e.g., a well)?
You drop a stone.
- If it hits bottom immediately (No children), depth is 1.
- If it goes down further, you ask "How deep is the left tunnel?" and "How deep is the right tunnel?".
- Your total depth is 1 (yourself) + Max(LeftDepth, RightDepth).

This recursive definition is the heart of DFS.
Depth(node) = 1 + max(Depth(left), Depth(right))

PROBLEM:
"Maximum Depth of Binary Tree"
A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

ALGORITHM:
1. Base Case: If `root == nil`, return 0. (Depth of nothing is 0).
2. Recursive Step:
   - `leftDepth = maxDepth(root.Left)`
   - `rightDepth = maxDepth(root.Right)`
3. Return `1 + max(leftDepth, rightDepth)`.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// Base Case: If we fell off the leaf
	if root == nil {
		return 0
	}

	// Recursive Step
	// DRY RUN:
	// Tree:   3
	//        / \
	//       9  20
	//          / \
	//         15  7
	//
	// Call(3):
	//   -> Call(9):
	//      -> Call(nil) -> 0
	//      -> Call(nil) -> 0
	//      -> Returns 1 + 0 = 1.
	//   -> Call(20):
	//      -> Call(15): Returns 1.
	//      -> Call(7): Returns 1.
	//      -> Returns 1 + Max(1,1) = 2.
	// Call(3) Returns 1 + Max(1, 2) = 3.

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + int(math.Max(float64(leftDepth), float64(rightDepth)))
}

func main() {
	// Build tree: [3,9,20,null,null,15,7]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Printf("Calculating Max Depth...\n")
	depth := maxDepth(root)
	fmt.Printf("Max Depth: %d\n", depth) // Expected: 3
}

/*
Deep Understanding:
Why DFS?
Because we need to reach the LEAF nodes to know the depth. DFS goes straight to leaves.
BFS (Level Order) could also work (counting levels), but DFS code is often cleaner/simpler (3 lines!).
*/
