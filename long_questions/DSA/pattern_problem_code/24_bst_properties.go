package main

import (
	"fmt"
	"math"
)

// Pattern: Binary Search Tree (BST) Properties
// Difficulty: Medium
// Key Concept: Utilizing the ordering property (Left < Root < Right) to solve problems efficiently.

/*
INTUITION:
"Validate Binary Search Tree"
Is this tree a valid BST?
      5
     / \
    1   4
       / \
      3   6
Wait! 4 is right of 5, but 4 < 5. This is INVALID. The Right subtree must contain ONLY nodes > Root.

Key Logic:
As we go down the tree, we narrow the "Valid Range" (Min, Max).
- Root (5): Valid Range (-inf, +inf).
- Go Left (1): Must be < 5. Range (-inf, 5).
- Go Right (4): Must be > 5. Range (5, +inf).
  - Check Node 4. Is 4 in (5, +inf)? NO! 4 is not > 5. Return False.

ALGORITHM:
Recursive DFS with constraints (low, high).
1. `validate(node, low, high)`
2. If `node == nil`, Valid.
3. If `node.val` not in `(low, high)`, Invalid.
4. Recurse Left: `validate(node.Left, low, node.val)`
5. Recurse Right: `validate(node.Right, node.val, high)`
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, low, high int64) bool {
	// Base Case: Empty trees are valid
	if node == nil {
		return true
	}

	val := int64(node.Val)

	// Constraint Check
	if val <= low || val >= high {
		return false
	}

	// Recursive Step
	// Left child must be smaller than current node (New High = val)
	// Right child must be bigger than current node (New Low = val)
	return validate(node.Left, low, val) && validate(node.Right, val, high)
}

func main() {
	// Invalid Tree: [5,1,4,null,null,3,6]
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}

	fmt.Printf("Validating BST...\n")
	res := isValidBST(root)
	fmt.Printf("Is Valid: %v\n", res) // Expected: false
}
