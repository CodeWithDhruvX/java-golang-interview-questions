package main

import (
	"fmt"
)

// Pattern: Morris Traversal
// Difficulty: Hard
// Key Concept: Traversing a Binary Tree in O(N) time and O(1) space (without stack or recursion) by creating temporary links.

/*
INTUITION:
Normal Inorder Traversal (Recursive or Stack) takes O(H) space.
Morris Traversal achieves O(1) space.
Idea: When we go Left, we need a way to come back to Root.
Usually Stack stores the Root.
Morris says: "Before going Left, let's find the 'predecessor' (Rightmost node of Left Subtree). Let's make its Right Child point to Root."
Now, when we finish the Left Subtree, we will end up at the Predecessor. Its Right pointer leads us back to Root!
Once we are back, we cut the link to restore the tree.

ALGORITHM:
1. Initialize `curr` as root.
2. While `curr` is not nil:
   - If `curr.left` is nil:
     - Visit `curr`.
     - `curr = curr.right`.
   - Else:
     - Find `predecessor` (Rightmost node of `curr.left`).
     - If `predecessor.right` is nil: (First time visiting)
       - `predecessor.right = curr` (Make thread).
       - `curr = curr.left`.
     - If `predecessor.right` is `curr`: (Second time visiting - returning from left)
       - `predecessor.right = nil` (Cut thread).
       - Visit `curr`.
       - `curr = curr.right`.

PROBLEM:
LeetCode 94. Binary Tree Inorder Traversal.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	curr := root

	for curr != nil {
		if curr.Left == nil {
			// No left child, visit this node and go right
			res = append(res, curr.Val)
			curr = curr.Right
		} else {
			// Find predecessor
			prev := curr.Left
			for prev.Right != nil && prev.Right != curr {
				prev = prev.Right
			}

			if prev.Right == nil {
				// Establish link
				prev.Right = curr
				curr = curr.Left
			} else {
				// Link already exists (we returned from left)
				// Remove link
				prev.Right = nil
				// Visit node
				res = append(res, curr.Val)
				// Go right
				curr = curr.Right
			}
		}
	}
	return res
}

func main() {
	/*
	   1
	    \
	     2
	    /
	   3
	*/
	// Inorder: 1 3 2
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	fmt.Printf("Inorder: %v\n", inorderTraversal(root))
}
