package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 450. Delete Node in a BST
// Time: O(H), Space: O(1)
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	
	// Find the node to delete
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// Node found
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// Node has two children
			// Find inorder successor (smallest in right subtree)
			minNode := findMin(root.Right)
			root.Val = minNode.Val
			// Delete the inorder successor
			root.Right = deleteNode(root.Right, minNode.Val)
		}
	}
	
	return root
}

// Find minimum node in BST
func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// Alternative version using inorder predecessor
func deleteNodeWithPredecessor(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	
	if key < root.Val {
		root.Left = deleteNodeWithPredecessor(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNodeWithPredecessor(root.Right, key)
	} else {
		// Node found
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// Find inorder predecessor (largest in left subtree)
			maxNode := findMax(root.Left)
			root.Val = maxNode.Val
			// Delete the inorder predecessor
			root.Left = deleteNodeWithPredecessor(root.Left, maxNode.Val)
		}
	}
	
	return root
}

// Find maximum node in BST
func findMax(node *TreeNode) *TreeNode {
	for node.Right != nil {
		node = node.Right
	}
	return node
}

// Iterative version
func deleteNodeIterative(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	
	var parent *TreeNode
	current := root
	
	// Find the node to delete and its parent
	for current != nil && current.Val != key {
		parent = current
		if key < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	
	if current == nil {
		return root // Node not found
	}
	
	// Case 1: Node has no children
	if current.Left == nil && current.Right == nil {
		if parent == nil {
			return nil // Deleting root
		}
		if parent.Left == current {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	}
	
	// Case 2: Node has one child
	if current.Left == nil || current.Right == nil {
		var child *TreeNode
		if current.Left != nil {
			child = current.Left
		} else {
			child = current.Right
		}
		
		if parent == nil {
			return child // Deleting root
		}
		
		if parent.Left == current {
			parent.Left = child
		} else {
			parent.Right = child
		}
	}
	
	// Case 3: Node has two children
	if current.Left != nil && current.Right != nil {
		// Find inorder successor
		successor := current.Right
		successorParent := current
		
		for successor.Left != nil {
			successorParent = successor
			successor = successor.Left
		}
		
		// Copy successor's value
		current.Val = successor.Val
		
		// Delete successor
		if successorParent.Left == successor {
			successorParent.Left = successor.Right
		} else {
			successorParent.Right = successor.Right
		}
	}
	
	return root
}

// Version that returns the deleted node
func deleteNodeWithReturn(root *TreeNode, key int) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	
	if key < root.Val {
		root.Left, deleted := deleteNodeWithReturn(root.Left, key)
		return root, deleted
	} else if key > root.Val {
		root.Right, deleted := deleteNodeWithReturn(root.Right, key)
		return root, deleted
	} else {
		// Node found
		if root.Left == nil {
			return root.Right, root
		} else if root.Right == nil {
			return root.Left, root
		} else {
			// Node has two children
			minNode := findMin(root.Right)
			root.Val = minNode.Val
			root.Right, deleted := deleteNodeWithReturn(root.Right, minNode.Val)
			return root, deleted
		}
	}
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

// Helper function to copy a tree
func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	
	newRoot := &TreeNode{Val: root.Val}
	newRoot.Left = copyTree(root.Left)
	newRoot.Right = copyTree(root.Right)
	
	return newRoot
}

func main() {
	// Test cases
	fmt.Println("=== Testing Delete Node in BST ===")
	
	testCases := []struct {
		initialTree []interface{}
		deleteKey   int
		description string
	}{
		{[]interface{}{5, 3, 6, 2, 4, nil, 7}, 3, "Delete node with two children"},
		{[]interface{}{5, 3, 6, 2, 4, nil, 7}, 2, "Delete leaf node"},
		{[]interface{}{5, 3, 6, 2, 4, nil, 7}, 6, "Delete node with one child"},
		{[]interface{}{5, 3, 6, 2, 4, nil, 7}, 5, "Delete root node"},
		{[]interface{}{}, 10, "Delete from empty tree"},
		{[]interface{}{10}, 10, "Delete single node tree"},
		{[]interface{}{10, 5, 15}, 10, "Delete root with two children"},
		{[]interface{}{10, 5, 15, 2, 7, 12, 18}, 15, "Delete right subtree node"},
		{[]interface{}{10, 5, 15, 2, 7, 12, 18}, 2, "Delete left subtree node"},
		{[]interface{}{10, 5, 15, 2, 7, 12, 18}, 20, "Delete non-existent node"},
	}
	
	for i, tc := range testCases {
		root := createTree(tc.initialTree)
		
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Initial tree: %v\n", tc.initialTree)
		fmt.Printf("  Delete key: %d\n", tc.deleteKey)
		
		// Test recursive version
		root1 := copyTree(root)
		result1 := deleteNode(root1, tc.deleteKey)
		
		// Test with predecessor
		root2 := copyTree(root)
		result2 := deleteNodeWithPredecessor(root2, tc.deleteKey)
		
		// Test iterative version
		root3 := copyTree(root)
		result3 := deleteNodeIterative(root3, tc.deleteKey)
		
		fmt.Printf("  Recursive result: %v\n", treeToArray(result1))
		fmt.Printf("  Predecessor result: %v\n", treeToArray(result2))
		fmt.Printf("  Iterative result: %v\n\n", treeToArray(result3))
	}
	
	// Test complex scenarios
	fmt.Println("=== Testing Complex Scenarios ===")
	
	// Build a complex BST
	complexBST := createBST([]int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45, 55, 65, 75, 85})
	fmt.Printf("Complex BST initial: %v\n", treeToArray(complexBST))
	
	// Delete various nodes
	deleteKeys := []int{50, 30, 70, 20, 80}
	for _, key := range deleteKeys {
		complexBST = deleteNode(complexBST, key)
		fmt.Printf("After deleting %d: %v\n", key, treeToArray(complexBST))
	}
	
	// Test edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	
	// Delete from single node tree
	singleNode := &TreeNode{Val: 10}
	singleNode = deleteNode(singleNode, 10)
	fmt.Printf("Single node tree after deletion: %v\n", treeToArray(singleNode))
	
	// Delete root repeatedly
	root := createBST([]int{10, 5, 15, 2, 7, 12, 18})
	for i := 0; i < 5; i++ {
		if root != nil {
			rootVal := root.Val
			root = deleteNode(root, rootVal)
			fmt.Printf("After deleting root %d: %v\n", rootVal, treeToArray(root))
		}
	}
	
	// Test with deleted node return
	fmt.Println("\n=== Testing with Deleted Node Return ===")
	testRoot := createBST([]int{5, 3, 6, 2, 4, nil, 7})
	testRoot, deleted := deleteNodeWithReturn(testRoot, 3)
	
	if deleted != nil {
		fmt.Printf("Deleted node value: %d\n", deleted.Val)
	}
	fmt.Printf("Tree after deletion: %v\n", treeToArray(testRoot))
}
