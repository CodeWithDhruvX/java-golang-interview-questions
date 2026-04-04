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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: BST Node Deletion with Rebalancing
- **Node Search**: Find node to delete using BST property
- **Three Cases**: Handle leaf, one child, two children scenarios
- **Successor/Predecessor**: Replace with inorder successor/predecessor
- **Tree Reconnection**: Maintain BST properties after deletion

## 2. PROBLEM CHARACTERISTICS
- **BST Maintenance**: Preserve BST properties after deletion
- **Structural Changes**: May need to restructure tree
- **Successor Finding**: Find inorder successor (smallest in right subtree)
- **Edge Cases**: Handle root deletion, single node, etc.

## 3. SIMILAR PROBLEMS
- Insert into BST (LeetCode 701) - Similar tree modification
- Validate BST (LeetCode 98) - Check BST properties
- Kth Smallest in BST (LeetCode 230) - Use inorder traversal
- Convert Sorted Array to BST (LeetCode 108) - Tree construction

## 4. KEY OBSERVATIONS
- **Three Deletion Cases**: Leaf, one child, two children
- **Successor Strategy**: Replace with inorder successor for two children
- **Tree Integrity**: Must maintain BST properties throughout
- **Parent Tracking**: Need parent references for iterative approach

## 5. VARIATIONS & EXTENSIONS
- **Predecessor Strategy**: Use inorder predecessor instead of successor
- **Return Deleted Node**: Return both new root and deleted node
- **Multiple Deletions**: Batch delete operations
- **Self-Balancing BST**: AVL or Red-Black tree variants

## 6. INTERVIEW INSIGHTS
- Always clarify: "Return deleted node? Handle duplicates? Tree size?"
- Edge cases: empty tree, single node, root deletion, key not found
- Time complexity: O(H) where H=tree height
- Space complexity: O(H) for recursion stack
- Key insight: handle three cases separately, use successor for two children

## 7. COMMON MISTAKES
- Not handling all three deletion cases
- Wrong successor/predecessor finding
- Not updating parent pointers correctly
- Memory leaks in recursive implementations
- Not handling root deletion properly

## 8. OPTIMIZATION STRATEGIES
- **Recursive Approach**: O(H) time, O(H) space - simple
- **Iterative Approach**: O(H) time, O(1) space - no recursion
- **Successor Caching**: Cache successor finding for repeated operations
- **Parent Tracking**: Maintain parent references for efficiency

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like removing an employee from an organization chart:**
- You have a hierarchical organization (BST)
- Need to remove someone while maintaining reporting structure
- If they have subordinates, need to reassign their reporting
- Like removing a manager and promoting their replacement
- Must keep the organization functional after changes

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: BST root, key to delete
2. **Goal**: Remove node with given key, maintain BST properties
3. **Constraints**: Must preserve BST ordering after deletion
4. **Output**: Modified BST root

#### Phase 2: Key Insight Recognition
- **"Three distinct cases"** → Leaf, one child, two children
- **"Successor strategy"** → Replace with inorder successor for two children
- **"Tree reconnection"** → Need to properly reconnect subtrees
- **"Parent tracking"** → Critical for iterative approach

#### Phase 3: Strategy Development
```
Human thought process:
"I need to delete a node from BST while maintaining properties.
This requires handling three cases:

Case 1: Node is leaf (no children)
- Simply remove the node (set parent's child pointer to nil)

Case 2: Node has one child
- Replace node with its child
- Update parent's pointer to point to grandchild

Case 3: Node has two children
- Find inorder successor (smallest in right subtree)
- Replace node's value with successor's value
- Delete the successor (which will be Case 1 or 2)
- This maintains BST properties

Each case preserves the BST ordering!"
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return nil
- **Key not found**: Return unchanged tree
- **Single node deletion**: Return nil
- **Root deletion**: Handle as special case

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
BST: [5, 3, 7, 2, 4, 6, 8], delete key = 5

Human thinking:
"Find node to delete:
Start at root (5), key = 5 ✓ Found node

Node (5) has two children (3 and 7) → Case 3
Find inorder successor:
- Go right to (7)
- Go left as far as possible: (7) → (6) → nil
- Successor is (6)

Replace and delete:
- Replace root value (5) with successor value (6)
- Delete successor (6) from its position
- Successor (6) has no children → Case 1
- Remove (6) by setting its parent's left pointer to nil

Final tree: [6, 3, 7, 2, 4, nil, 8] ✓ BST maintained"
```

#### Phase 6: Intuition Validation
- **Why three cases work**: Covers all possible node configurations
- **Why successor works**: Smallest in right subtree maintains ordering
- **Why BST preserved**: Inorder successor is next element in sequence
- **Why O(H)**: Only traverse one path down the tree

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just set to nil?"** → Need to handle children properly
2. **"Should I use predecessor?"** → Either works, be consistent
3. **"What about duplicates?"** → Clarify duplicate handling policy
4. **"Can I optimize further?"** → O(H) is already optimal
5. **"What about iterative vs recursive?"** → Trade-offs between simplicity and memory

### Real-World Analogy
**Like removing a product from a sorted inventory system:**
- Products are organized in a hierarchical sorted structure
- Need to remove a product while maintaining sorted order
- If product has subcategories, need to reorganize them
- Like removing a category and promoting its subcategory
- Must keep the inventory system functional after changes

### Human-Readable Pseudocode
```
function deleteNode(root, key):
    if root is null:
        return null
    
    if key < root.val:
        root.left = deleteNode(root.left, key)
    else if key > root.val:
        root.right = deleteNode(root.right, key)
    else:
        // Node found - handle three cases
        if root.left is null:
            return root.right
        else if root.right is null:
            return root.left
        else:
            // Two children - find successor
            successor = findMin(root.right)
            root.val = successor.val
            root.right = deleteNode(root.right, successor.val)
    
    return root

function findMin(node):
    while node.left is not null:
        node = node.left
    return node
```

### Execution Visualization

### Example: BST = [5, 3, 7, 2, 4, 6, 8], delete key = 5
```
Initial Tree:
        5
       /   \
      3     7
     / \   / \
    2   4 6   8

Step 1: Find node to delete
- Start at root (5), key = 5 ✓ Found node

Step 2: Determine case
- Node (5) has two children → Case 3

Step 3: Find inorder successor
- Go right to (7)
- Go left as far as possible: (7) → (6) → nil
- Successor is (6)

Step 4: Replace and delete
- Replace root value (5) with successor value (6)
- Delete successor (6) from its position
- Successor (6) has no children → Case 1
- Remove (6) by setting its parent's left pointer to nil

Final Tree:
        6
       /   \
      3     7
     / \     \
    2   4     8

BST properties maintained ✓
```

### Key Visualization Points:
- **Three Cases**: Leaf, one child, two children
- **Successor Finding**: Smallest in right subtree
- **Tree Reconnection**: Proper pointer updates
- **BST Preservation**: Ordering maintained throughout

### Memory Layout Visualization:
```
Recursive Stack During Deletion:
Call Stack:              Operation:
deleteNode(5, 5)         Found node, Case 3
├── findMin(7)            Find successor
│   ├── findMin(6)        Successor found = 6
│   └── return 6
├── deleteNode(7, 6)      Delete successor
│   ├── 6 < 7, go left
│   ├── deleteNode(6, 6) Found node, Case 1
│   │   └── return null
│   └── return 7 (with 6.left = nil)
└── return 5 (with val = 6)

Tree State Evolution:
Initial:    5          After replacement:   6
           / \                      / \
          3   7                  3   7
         / \ / \                / \   \
        2   4 6   8            2   4     8
```

### Time Complexity Breakdown:
- **Node Search**: O(H) time to find node to delete
- **Successor Finding**: O(H) time in worst case
- **Tree Restructuring**: O(1) time for pointer updates
- **Total**: O(H) time, O(H) space for recursion stack

### Alternative Approaches:

#### 1. Iterative Deletion (O(H) time, O(1) space)
```go
func deleteNodeIterative(root *TreeNode, key int) *TreeNode {
    // Find node and its parent iteratively
    // Handle three cases with parent pointers
    // No recursion, uses O(1) extra space
    // ... implementation details omitted
}
```
- **Pros**: No recursion, memory efficient
- **Cons**: More complex implementation

#### 2. Predecessor Strategy (O(H) time, O(H) space)
```go
func deleteNodeWithPredecessor(root *TreeNode, key int) *TreeNode {
    // Similar to successor but use inorder predecessor
    // Find largest in left subtree instead
    // ... implementation details omitted
}
```
- **Pros**: Alternative approach, same complexity
- **Cons**: Different successor/predecessor choice

#### 3. Return Deleted Node (O(H) time, O(H) space)
```go
func deleteNodeWithReturn(root *TreeNode, key int) (*TreeNode, *TreeNode) {
    // Return both new root and deleted node
    // Useful for applications needing deleted node
    // ... implementation details omitted
}
```
- **Pros**: Provides deleted node for further processing
- **Cons**: Slightly more complex return handling

### Extensions for Interviews:
- **Duplicate Handling**: Handle nodes with equal values
- **Multiple Deletions**: Batch delete operations
- **Self-Balancing BST**: AVL or Red-Black tree variants
- **Memory Management**: Proper cleanup in languages without GC
- **Performance Analysis**: Discuss worst-case scenarios
*/
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
