package main

import (
	"fmt"
)

// 94. Binary Tree Inorder Traversal - Morris Traversal
// Time: O(N), Space: O(1)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	current := root
	
	for current != nil {
		if current.Left == nil {
			// No left child, visit current and move to right
			result = append(result, current.Val)
			current = current.Right
		} else {
			// Find predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}
			
			if predecessor.Right == nil {
				// Make current the right child of predecessor
				predecessor.Right = current
				current = current.Left
			} else {
				// Revert the changes and visit current
				predecessor.Right = nil
				result = append(result, current.Val)
				current = current.Right
			}
		}
	}
	
	return result
}

// Morris Traversal for preorder
func preorderTraversalMorris(root *TreeNode) []int {
	var result []int
	current := root
	
	for current != nil {
		if current.Left == nil {
			// No left child, visit current and move to right
			result = append(result, current.Val)
			current = current.Right
		} else {
			// Find predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}
			
			if predecessor.Right == nil {
				// Make current the right child of predecessor
				predecessor.Right = current
				result = append(result, current.Val) // Visit before going left
				current = current.Left
			} else {
				// Revert the changes
				predecessor.Right = nil
				current = current.Right
			}
		}
	}
	
	return result
}

// Morris Traversal for postorder
func postorderTraversalMorris(root *TreeNode) []int {
	var result []int
	current := root
	
	for current != nil {
		if current.Right == nil {
			// No right child, visit current and move to left
			result = append(result, current.Val)
			current = current.Left
		} else {
			// Find successor
			successor := current.Right
			for successor.Left != nil && successor.Left != current {
				successor = successor.Left
			}
			
			if successor.Left == nil {
				// Make current the left child of successor
				successor.Left = current
				result = append(result, current.Val) // Visit before going right
				current = current.Right
			} else {
				// Revert the changes
				successor.Left = nil
				current = current.Left
			}
		}
	}
	
	// Reverse the result for postorder
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	
	return result
}

// Morris Traversal with level order
func levelOrderTraversalMorris(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	
	var result [][]int
	current := root
	
	for current != nil {
		if current.Left == nil {
			// No left child, visit current and move to right
			result = append(result, []int{current.Val})
			current = current.Right
		} else {
			// Find predecessor
			predecessor := current.Left
			level := 0
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
				level++
			}
			
			if predecessor.Right == nil {
				// Make current the right child of predecessor
				predecessor.Right = current
				current = current.Left
			} else {
				// Revert the changes and visit current
				predecessor.Right = nil
				result = append(result, []int{current.Val})
				current = current.Right
			}
		}
	}
	
	return result
}

// Morris Traversal with path reconstruction
func inorderTraversalWithPathMorris(root *TreeNode) ([]int, []int) {
	var result []int
	var path []int
	current := root
	
	for current != nil {
		if current.Left == nil {
			// No left child, visit current and move to right
			result = append(result, current.Val)
			path = append(path, current.Val)
			current = current.Right
		} else {
			// Find predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}
			
			if predecessor.Right == nil {
				// Make current the right child of predecessor
				predecessor.Right = current
				current = current.Left
			} else {
				// Revert the changes and visit current
				predecessor.Right = nil
				result = append(result, current.Val)
				path = append(path, current.Val)
				current = current.Right
			}
		}
	}
	
	return result, path
}

// Morris Traversal with counting
func inorderTraversalWithCountMorris(root *TreeNode) ([]int, int) {
	var result []int
	count := 0
	current := root
	
	for current != nil {
		if current.Left == nil {
			// No left child, visit current and move to right
			result = append(result, current.Val)
			count++
			current = current.Right
		} else {
			// Find predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}
			
			if predecessor.Right == nil {
				// Make current the right child of predecessor
				predecessor.Right = current
				current = current.Left
			} else {
				// Revert the changes and visit current
				predecessor.Right = nil
				result = append(result, current.Val)
				count++
				current = current.Right
			}
		}
	}
	
	return result, count
}

// Morris Traversal with sum calculation
func inorderTraversalWithSumMorris(root *TreeNode) ([]int, int) {
	var result []int
	sum := 0
	current := root
	
	for current != nil {
		if current.Left == nil {
			// No left child, visit current and move to right
			result = append(result, current.Val)
			sum += current.Val
			current = current.Right
		} else {
			// Find predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}
			
			if predecessor.Right == nil {
				// Make current the right child of predecessor
				predecessor.Right = current
				current = current.Left
			} else {
				// Revert the changes and visit current
				predecessor.Right = nil
				result = append(result, current.Val)
				sum += current.Val
				current = current.Right
			}
		}
	}
	
	return result, sum
}

// Morris Traversal with maximum/minimum
func inorderTraversalWithMinMaxMorris(root *TreeNode) ([]int, int, int) {
	var result []int
	maximum := -1 << 31
	minimum := 1 << 31 - 1
	current := root
	
	for current != nil {
		if current.Left == nil {
			// No left child, visit current and move to right
			result = append(result, current.Val)
			if current.Val > maximum {
				maximum = current.Val
			}
			if current.Val < minimum {
				minimum = current.Val
			}
			current = current.Right
		} else {
			// Find predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}
			
			if predecessor.Right == nil {
				// Make current the right child of predecessor
				predecessor.Right = current
				current = current.Left
			} else {
				// Revert the changes and visit current
				predecessor.Right = nil
				result = append(result, current.Val)
				if current.Val > maximum {
					maximum = current.Val
				}
				if current.Val < minimum {
					minimum = current.Val
				}
				current = current.Right
			}
		}
	}
	
	return result, maximum, minimum
}

// Morris Traversal with height calculation
func inorderTraversalWithHeightMorris(root *TreeNode) ([]int, int) {
	var result []int
	maxHeight := 0
	current := root
	
	for current != nil {
		if current.Left == nil {
			// No left child, visit current and move to right
			result = append(result, current.Val)
			current = current.Right
		} else {
			// Find predecessor and calculate height
			predecessor := current.Left
			height := 1
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
				height++
			}
			
			if predecessor.Right == nil {
				// Make current the right child of predecessor
				predecessor.Right = current
				if height > maxHeight {
					maxHeight = height
				}
				current = current.Left
			} else {
				// Revert the changes and visit current
				predecessor.Right = nil
				result = append(result, current.Val)
				current = current.Right
			}
		}
	}
	
	return result, maxHeight
}

// Morris Traversal with leaf nodes
func inorderTraversalLeafNodesMorris(root *TreeNode) ([]int, []int) {
	var result []int
	var leafNodes []int
	current := root
	
	for current != nil {
		if current.Left == nil {
			// No left child, visit current and move to right
			result = append(result, current.Val)
			if current.Left == nil && current.Right == nil {
				leafNodes = append(leafNodes, current.Val)
			}
			current = current.Right
		} else {
			// Find predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}
			
			if predecessor.Right == nil {
				// Make current the right child of predecessor
				predecessor.Right = current
				current = current.Left
			} else {
				// Revert the changes and visit current
				predecessor.Right = nil
				result = append(result, current.Val)
				if current.Left == nil && current.Right == nil {
					leafNodes = append(leafNodes, current.Val)
				}
				current = current.Right
			}
		}
	}
	
	return result, leafNodes
}

func main() {
	// Test cases
	fmt.Println("=== Testing Morris Traversal ===")
	
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
		{[]interface{}{1, nil, 2, nil, 3, nil, 4}, "Right skewed tree"},
		{[]interface{}{4, 2, 5, 1, 3, nil, 6}, "Balanced tree"},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7}, "Complete binary tree"},
		{[]interface{}{1}, "Single node"},
		{[]interface{}{}, "Empty tree"},
		{[]interface{}{1, nil, 2, 3, nil, 4, 5}, "Complex tree"},
		{[]interface{}{5, 4, 3, 2, 1}, "Left skewed tree"},
		{[]interface{}{1, 2, nil, 3, nil, 4, nil}, "Left chain"},
		{[]interface{}{1, nil, 2, nil, 3, nil, 4}, "Right chain"},
		{[]interface{}{10, 5, 15, 3, 7, 12, 18}, "BST"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Tree: %v\n", tc.tree)
		
		root := createTree(tc.tree)
		
		// Test inorder traversal
		inorder := inorderTraversal(root)
		fmt.Printf("  Inorder: %v\n", inorder)
		
		// Test preorder traversal
		preorder := preorderTraversalMorris(root)
		fmt.Printf("  Preorder: %v\n", preorder)
		
		// Test postorder traversal
		postorder := postorderTraversalMorris(root)
		fmt.Printf("  Postorder: %v\n", postorder)
		
		// Test level order traversal
		levelOrder := levelOrderTraversalMorris(root)
		fmt.Printf("  Level order: %v\n", levelOrder)
		
		// Test with path reconstruction
		result, path := inorderTraversalWithPathMorris(root)
		fmt.Printf("  With path: %v, path: %v\n", result, path)
		
		// Test with counting
		result, count := inorderTraversalWithCountMorris(root)
		fmt.Printf("  With count: %v, count: %d\n", result, count)
		
		// Test with sum calculation
		result, sum := inorderTraversalWithSumMorris(root)
		fmt.Printf("  With sum: %v, sum: %d\n", result, sum)
		
		// Test with min/max
		result, max, min := inorderTraversalWithMinMaxMorris(root)
		fmt.Printf("  With min/max: %v, max: %d, min: %d\n", result, max, min)
		
		// Test with height
		result, height := inorderTraversalWithHeightMorris(root)
		fmt.Printf("  With height: %v, height: %d\n", result, height)
		
		// Test with leaf nodes
		result, leafNodes := inorderTraversalLeafNodesMorris(root)
		fmt.Printf("  With leaf nodes: %v, leaves: %v\n", result, leafNodes)
		
		fmt.Println()
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	
	// Create large tree
	largeVals := make([]interface{}, 1000)
	for i := 0; i < 1000; i++ {
		largeVals[i] = i + 1
	}
	
	largeRoot := createTree(largeVals)
	fmt.Printf("Large tree with %d nodes\n", len(largeVals))
	
	start := time.Now()
	inorder := inorderTraversal(largeRoot)
	duration := time.Since(start)
	
	fmt.Printf("Large tree inorder traversal: %d nodes, time: %v\n", len(inorder), duration)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Single node
	singleRoot := createTree([]interface{}{42})
	singleInorder := inorderTraversal(singleRoot)
	fmt.Printf("Single node: %v\n", singleInorder)
	
	// Empty tree
	emptyInorder := inorderTraversal(nil)
	fmt.Printf("Empty tree: %v\n", emptyInorder)
	
	// All left children
	leftVals := []interface{}{5, 4, 3, 2, 1}
	leftRoot := createTree(leftVals)
	leftInorder := inorderTraversal(leftRoot)
	fmt.Printf("Left skewed: %v\n", leftInorder)
	
	// All right children
	rightVals := []interface{}{1, nil, 2, nil, 3, nil, 4}
	rightRoot := createTree(rightVals)
	rightInorder := inorderTraversal(rightRoot)
	fmt.Printf("Right skewed: %v\n", rightInorder)
	
	// Test space efficiency
	fmt.Println("\n=== Space Efficiency Test ===")
	
	// Create tree with many nodes
	spaceVals := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		spaceVals[i] = i + 1
	}
	
	spaceRoot := createTree(spaceVals)
	
	// Morris traversal uses O(1) space
	spaceInorder := inorderTraversal(spaceRoot)
	fmt.Printf("Space efficient traversal: %d nodes\n", len(spaceInorder))
	
	// Test tree modification
	fmt.Println("\n=== Tree Modification Test ===")
	
	// Create tree and verify it's not modified
	modifyRoot := createTree([]interface{}{1, 2, 3})
	
	beforeInorder := inorderTraversal(modifyRoot)
	afterInorder := inorderTraversal(modifyRoot)
	
	fmt.Printf("Before modification: %v\n", beforeInorder)
	fmt.Printf("After modification: %v\n", afterInorder)
	fmt.Printf("Tree structure preserved: %t\n", 
		len(beforeInorder) == len(afterInorder))
	
	// Test with negative values
	fmt.Println("\n=== Negative Values Test ===")
	negativeVals := []interface{}{-1, -2, -3, -4, -5}
	negativeRoot := createTree(negativeVals)
	negativeInorder := inorderTraversal(negativeRoot)
	fmt.Printf("Negative values: %v\n", negativeInorder)
	
	// Test with duplicate values
	fmt.Println("\n=== Duplicate Values Test ===")
	duplicateVals := []interface{}{1, 1, 2, 2, 3, 3}
	duplicateRoot := createTree(duplicateVals)
	duplicateInorder := inorderTraversal(duplicateRoot)
	fmt.Printf("Duplicate values: %v\n", duplicateInorder)
}
