package main

import "fmt"

// Pattern: Tree BFS (Breadth First Search)
// Difficulty: Medium
// Key Concept: Exploring a tree level-by-level (horizontal) instead of depth-first (vertical).

/*
INTUITION:
Imagine pouring water on top of a pyramid.
First, the top stone gets wet (Level 0).
Then, the water flows to the two stones below it (Level 1).
Then, it flows to the 4 stones below that (Level 2).
The water never reaches Level 2 before finishing Level 1.

To implement this "Level" order, we use a QUEUE (FIFO - First In, First Out).
- We enter a level. We count how many nodes are in it (`size`).
- We process exactly that many nodes.
- As we process them, we add their children to the back of the queue (for the NEXT level).

PROBLEM:
"Binary Tree Level Order Traversal"
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

ALGORITHM:
1. If root is nil, return empty.
2. Initialize Queue with `root`.
3. Loop while Queue is not empty:
   - Calculate `levelSize = len(queue)`. (Crucial! Snapshot the size).
   - Loop `i` from 0 to `levelSize`:
     - Pop node from front.
     - Add value to current level result.
     - Push Left Child (if exists).
     - Push Right Child (if exists).
   - Add current level result to final answer.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	// DRY RUN:
	// Tree:   3
	//        / \
	//       9  20
	//          / \
	//         15  7
	//
	// Init: Q=[3]
	//
	// Level 1:
	//   Size=1. Loop 1 time.
	//   Pop 3. Res=[[3]].
	//   Add 3.Left(9), 3.Right(20). Q=[9, 20].
	//
	// Level 2:
	//   Size=2. Loop 2 times.
	//   Iter 1: Pop 9. Res=[[3], [9]]. Add children (nil). Q=[20].
	//   Iter 2: Pop 20. Res=[[3], [9, 20]]. Add 20.Left(15), 20.Right(7). Q=[15, 7].
	//
	// Level 3:
	//   Size=2. Loop 2 times.
	//   Iter 1: Pop 15. Q=[7].
	//   Iter 2: Pop 7. Q=[].
	//   Res=[[3], [9, 20], [15, 7]].
	//
	// Queue Empty. Done.

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		// Process ONLY the nodes that were present at the start of this level
		for i := 0; i < levelSize; i++ {
			node := queue[0]  // Peek
			queue = queue[1:] // Dequeue

			currentLevel = append(currentLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

func main() {
	// Build tree: [3,9,20,null,null,15,7]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Printf("Level Order Traversal:\n")
	res := levelOrder(root)
	fmt.Printf("%v\n", res) // Expected: [[3], [9, 20], [15, 7]]
}
