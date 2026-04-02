package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 701. Insert into a Binary Search Tree
// Time: O(H) where H is tree height, Space: O(1)
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	
	current := root
	for {
		if val < current.Val {
			if current.Left == nil {
				current.Left = &TreeNode{Val: val}
				break
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &TreeNode{Val: val}
				break
			}
			current = current.Right
		}
	}
	
	return root
}

// Recursive version
func insertIntoBSTRecursive(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	
	if val < root.Val {
		root.Left = insertIntoBSTRecursive(root.Left, val)
	} else {
		root.Right = insertIntoBSTRecursive(root.Right, val)
	}
	
	return root
}

// Version that returns the new root (useful when root is nil)
func insertIntoBSTWithRoot(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val: val}
	
	if root == nil {
		return newNode
	}
	
	parent := findInsertParent(root, val)
	if val < parent.Val {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}
	
	return root
}

func findInsertParent(root *TreeNode, val int) *TreeNode {
	parent := root
	current := root
	
	for current != nil {
		parent = current
		if val < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	
	return parent
}

// Version with path tracking
func insertIntoBSTWithPath(root *TreeNode, val int) (*TreeNode, []int) {
	if root == nil {
		return &TreeNode{Val: val}, []int{}
	}
	
	var path []int
	current := root
	
	for {
		path = append(path, current.Val)
		
		if val < current.Val {
			if current.Left == nil {
				current.Left = &TreeNode{Val: val}
				path = append(path, val)
				break
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &TreeNode{Val: val}
				path = append(path, val)
				break
			}
			current = current.Right
		}
	}
	
	return root, path
}

// Version with duplicate handling (insert duplicates to right)
func insertIntoBSTWithDuplicates(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	
	current := root
	for {
		if val < current.Val {
			if current.Left == nil {
				current.Left = &TreeNode{Val: val}
				break
			}
			current = current.Left
		} else {
			// Insert duplicates to the right
			if current.Right == nil {
				current.Right = &TreeNode{Val: val}
				break
			}
			current = current.Right
		}
	}
	
	return root
}

// Helper function to create BST from array
func createBSTFromSorted(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	
	return createBSTFromSortedHelper(nums, 0, len(nums)-1)
}

func createBSTFromSortedHelper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	
	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}
	
	root.Left = createBSTFromSortedHelper(nums, left, mid-1)
	root.Right = createBSTFromSortedHelper(nums, mid+1, right)
	
	return root
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

// Helper function to print tree structure
func printTree(root *TreeNode, level int) {
	if root == nil {
		return
	}
	
	// Print right subtree first
	printTree(root.Right, level+1)
	
	// Print current node
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Println(root.Val)
	
	// Print left subtree
	printTree(root.Left, level+1)
}

func main() {
	// Test cases
	fmt.Println("=== Testing Insert into Binary Search Tree ===")
	
	testCases := []struct {
		initialTree []interface{}
		insertVal   int
		description string
	}{
		{[]interface{}{4, 2, 7, 1, 3}, 5, "Insert into existing tree"},
		{[]interface{}{}, 10, "Insert into empty tree"},
		{[]interface{}{10}, 5, "Insert smaller value"},
		{[]interface{}{10}, 15, "Insert larger value"},
		{[]interface{}{10, 5, 15}, 12, "Insert between values"},
		{[]interface{}{10, 5, 15, 2, 7, 12, 18}, 1, "Insert as new minimum"},
		{[]interface{}{10, 5, 15, 2, 7, 12, 18}, 20, "Insert as new maximum"},
		{[]interface{}{10, 5, 15, 2, 7, 12, 18}, 6, "Insert in left subtree"},
		{[]interface{}{10, 5, 15, 2, 7, 12, 18}, 14, "Insert in right subtree"},
		{[]interface{}{10, 5, 15}, 10, "Insert duplicate value"},
	}
	
	for i, tc := range testCases {
		root := createTree(tc.initialTree)
		
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Initial tree: %v\n", tc.initialTree)
		fmt.Printf("  Insert value: %d\n", tc.insertVal)
		
		// Test iterative version
		root1 := copyTree(root)
		result1 := insertIntoBST(root1, tc.insertVal)
		
		// Test recursive version
		root2 := copyTree(root)
		result2 := insertIntoBSTRecursive(root2, tc.insertVal)
		
		// Test with path tracking
		root3 := copyTree(root)
		result3, path := insertIntoBSTWithPath(root3, tc.insertVal)
		
		// Test with duplicates
		root4 := copyTree(root)
		result4 := insertIntoBSTWithDuplicates(root4, tc.insertVal)
		
		fmt.Printf("  Iterative result (in-order): %v\n", treeToArray(result1))
		fmt.Printf("  Recursive result (in-order): %v\n", treeToArray(result2))
		fmt.Printf("  Path taken: %v\n", path)
		fmt.Printf("  With duplicates (in-order): %v\n\n", treeToArray(result4))
	}
	
	// Test building BST from sorted array
	fmt.Println("=== Building BST from Sorted Array ===")
	sortedArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bst := createBSTFromSorted(sortedArray)
	fmt.Printf("Sorted array: %v\n", sortedArray)
	fmt.Printf("BST in-order (should be same): %v\n", treeToArray(bst))
	
	// Insert into balanced BST
	fmt.Println("\n=== Inserting into Balanced BST ===")
	fmt.Printf("Before insertion: %v\n", treeToArray(bst))
	bst = insertIntoBST(bst, 11)
	fmt.Printf("After inserting 11: %v\n", treeToArray(bst))
	bst = insertIntoBST(bst, 0)
	fmt.Printf("After inserting 0: %v\n", treeToArray(bst))
	
	// Test with many insertions
	fmt.Println("\n=== Testing Multiple Insertions ===")
	var root *TreeNode
	insertValues := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	
	for _, val := range insertValues {
		root = insertIntoBST(root, val)
	}
	
	fmt.Printf("Inserted values: %v\n", insertValues)
	fmt.Printf("Final BST in-order: %v\n", treeToArray(root))
	
	// Verify BST property
	isValid := validateBST(root, nil, nil)
	fmt.Printf("BST property valid: %t\n", isValid)
	
	// Test edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	
	// Very large values
	largeRoot := insertIntoBST(nil, 1000000)
	largeRoot = insertIntoBST(largeRoot, 2000000)
	largeRoot = insertIntoBST(largeRoot, 500000)
	fmt.Printf("Large values BST: %v\n", treeToArray(largeRoot))
	
	// Negative values
	negRoot := insertIntoBST(nil, 0)
	negRoot = insertIntoBST(negRoot, -10)
	negRoot = insertIntoBST(negRoot, 10)
	negRoot = insertIntoBST(negRoot, -5)
	negRoot = insertIntoBST(negRoot, 5)
	fmt.Printf("Negative values BST: %v\n", treeToArray(negRoot))
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

// Helper function to create tree from array
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

// BST validation helper
func validateBST(node *TreeNode, min, max *TreeNode) bool {
	if node == nil {
		return true
	}
	
	if min != nil && node.Val <= min.Val {
		return false
	}
	if max != nil && node.Val >= max.Val {
		return false
	}
	
	return validateBST(node.Left, min, node) && validateBST(node.Right, node, max)
}
