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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: BST Insertion with Traversal
- **BST Property**: Left subtree < root < right subtree
- **Node Navigation**: Traverse down tree to find insertion point
- **Leaf Insertion**: Insert new node as leaf
- **Multiple Approaches**: Iterative, recursive, with path tracking

## 2. PROBLEM CHARACTERISTICS
- **Dynamic Structure**: Tree grows with each insertion
- **Search Operation**: Need to find correct position
- **Pointer Updates**: Update parent's left/right pointers
- **Empty Tree Handling**: Special case for first insertion

## 3. SIMILAR PROBLEMS
- Delete Node in BST (LeetCode 450) - Tree modification
- Validate BST (LeetCode 98) - Check BST properties
- Kth Smallest in BST (LeetCode 230) - BST traversal
- Convert Sorted Array to BST (LeetCode 108) - Tree construction

## 4. KEY OBSERVATIONS
- **Search Path**: Follow BST property to find insertion point
- **Leaf Insertion**: New nodes always become leaves
- **Parent Tracking**: Critical for iterative approach
- **Duplicate Handling**: Policy for equal values (left or right)

## 5. VARIATIONS & EXTENSIONS
- **Duplicate Policy**: Insert duplicates to left or right subtree
- **Balanced BST**: Self-balancing variants (AVL, Red-Black)
- **Path Tracking**: Return insertion path for debugging
- **Batch Insertion**: Insert multiple elements efficiently

## 6. INTERVIEW INSIGHTS
- Always clarify: "Duplicate handling? Tree size? Return path?"
- Edge cases: empty tree, single node, duplicates, large values
- Time complexity: O(H) where H=tree height
- Space complexity: O(1) for iterative, O(H) for recursive
- Key insight: follow BST property to find insertion point

## 7. COMMON MISTAKES
- Wrong comparison direction (val < current.val vs >)
- Not handling empty tree case
- Not updating parent pointers correctly
- Infinite loops in iterative implementation
- Wrong duplicate handling policy

## 8. OPTIMIZATION STRATEGIES
- **Iterative Approach**: O(H) time, O(1) space - optimal
- **Recursive Approach**: O(H) time, O(H) space - simple
- **Path Tracking**: O(H) time, O(H) space - useful for debugging
- **Balanced BST**: O(log N) time, O(N) space for self-balancing

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like adding a new book to a sorted library:**
- You have a library organized by book numbers (BST)
- Need to place new book in correct position
- Walk through the library comparing book numbers
- Find the empty shelf where the new book belongs
- Like following a sorted filing system to find the right spot

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: BST root, value to insert
2. **Goal**: Insert value while maintaining BST properties
3. **Constraints**: Left < root < right for all nodes
4. **Output**: Modified BST root

#### Phase 2: Key Insight Recognition
- **"BST navigation"** → Follow comparison to find insertion point
- **"Leaf insertion"** → New nodes always become leaves
- **"Parent tracking"** → Need parent reference for iterative
- **"Duplicate policy"** → Decide where equal values go

#### Phase 3: Strategy Development
```
Human thought process:
"I need to insert a value into BST while maintaining properties.
This requires finding the correct insertion point:

Iterative Approach:
1. Start at root
2. Compare value with current node:
   - If value < current.val: go left
   - If value > current.val: go right
   - If equal: handle duplicates (policy dependent)
3. Continue until find empty spot (nil child pointer)
4. Insert new node at that position
5. Return root (unchanged unless tree was empty)

Recursive Approach:
1. Base case: if node is nil, create new node
2. Compare value with current node
3. Recurse left or right based on comparison
4. Return updated node

Both approaches maintain BST properties!"
```

#### Phase 4: Edge Case Handling
- **Empty tree**: New node becomes root
- **Single node**: Insert as left or right child
- **Duplicates**: Follow policy (left or right)
- **Large values**: Handle integer overflow if needed

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
BST: [5, 3, 7, 2, 4], insert value = 6

Human thinking:
"Start at root (5):
Compare 6 with 5: 6 > 5, go right to (7)
Compare 6 with 7: 6 < 7, go left to (6)
Compare 6 with 6: equal, handle duplicates
Policy: insert duplicates to right
Current node (6) has no right child
Insert 6 as right child of (6)

Final tree: [5, 3, 7, 2, 4, 6, nil] ✓ BST maintained"
```

#### Phase 6: Intuition Validation
- **Why BST navigation works**: Comparison leads to correct position
- **Why leaf insertion works**: New nodes don't disrupt existing structure
- **Why O(H)**: Only traverse one path down the tree
- **Why parent tracking matters**: Need to update parent's child pointer

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just append?"** → Need to maintain BST ordering
2. **"Should I use recursion?"** → Both work, iterative avoids recursion depth
3. **"What about duplicates?"** → Clarify duplicate handling policy
4. **"Can I optimize further?"** → O(H) is already optimal
5. **"What about balanced trees?"** → Different problem (self-balancing)

### Real-World Analogy
**Like adding a new employee to a company hierarchy:**
- You have an organization chart sorted by employee ID
- Need to place new employee in correct position
- Follow the hierarchy comparing employee IDs
- Find the empty position where new employee belongs
- Like following a sorted filing system to find the right spot

### Human-Readable Pseudocode
```
function insertIntoBST(root, val):
    if root is null:
        return new Node(val)
    
    current = root
    while true:
        if val < current.val:
            if current.left is null:
                current.left = new Node(val)
                break
            current = current.left
        else if val > current.val:
            if current.right is null:
                current.right = new Node(val)
                break
            current = current.right
        else:
            // Handle duplicates based on policy
            current.right = new Node(val)
            break
    
    return root

// Recursive version
function insertIntoBSTRecursive(root, val):
    if root is null:
        return new Node(val)
    
    if val < root.val:
        root.left = insertIntoBSTRecursive(root.left, val)
    else:
        root.right = insertIntoBSTRecursive(root.right, val)
    
    return root
```

### Execution Visualization

### Example: BST = [5, 3, 7, 2, 4], insert value = 6
```
Initial Tree:
        5
       /   \
      3     7
     / \   
    2   4

Step 1: Start at root (5)
Compare 6 with 5: 6 > 5, go right to (7)

Step 2: At node (7)
Compare 6 with 7: 6 < 7, go left to (6)

Step 3: At node (6)
Compare 6 with 6: equal, handle duplicates
Policy: insert duplicates to right
Current node (6) has no right child

Step 4: Insert new node
Insert 6 as right child of (6)

Final Tree:
        5
       /   \
      3     7
     / \   /
    2   4 6

BST properties maintained ✓
```

### Key Visualization Points:
- **BST Navigation**: Follow comparisons to find insertion point
- **Leaf Insertion**: New nodes always become leaves
- **Parent Updates**: Critical for maintaining structure
- **Duplicate Handling**: Policy-dependent placement

### Memory Layout Visualization:
```
Tree State During Insertion:
Initial:    After insertion:
    5           5
   / \          / \
  3   7        3   7
 / \            / \ /
2   4          2   4 6

Pointer Updates:
- Node (7).left was nil, now points to (6)
- Node (6).right was nil, now points to new node
- All other pointers unchanged

BST Properties:
- Left subtree: [2,3,4] < 5 ✓
- Right subtree: [6,7] > 5 ✓
- All subtrees maintain BST property ✓
```

### Time Complexity Breakdown:
- **Search Path**: O(H) time to find insertion point
- **Node Creation**: O(1) time
- **Pointer Update**: O(1) time
- **Total**: O(H) time, O(1) space for iterative, O(H) for recursive

### Alternative Approaches:

#### 1. Recursive Insertion (O(H) time, O(H) space)
```go
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
```
- **Pros**: Simple, elegant code
- **Cons**: Uses recursion stack, potential stack overflow

#### 2. Path Tracking (O(H) time, O(H) space)
```go
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
```
- **Pros**: Returns insertion path for debugging
- **Cons**: Extra space for path storage

#### 3. Self-Balancing BST (O(log N) time, O(N) space)
```go
type AVLNode struct {
    Val   int
    Left  *AVLNode
    Right *AVLNode
    Height int
}

func insertIntoAVL(root *AVLNode, val int) *AVLNode {
    // Insert like regular BST
    // Then rebalance to maintain AVL properties
    // More complex but guarantees O(log N) operations
    // ... implementation details omitted
}
```
- **Pros**: Guaranteed O(log N) operations
- **Cons**: Much more complex implementation

### Extensions for Interviews:
- **Duplicate Policies**: Different strategies for equal values
- **Balanced Trees**: Discuss AVL, Red-Black trees
- **Batch Insertion**: Insert multiple elements efficiently
- **Performance Analysis**: Compare different approaches
- **Memory Management**: Proper cleanup considerations
*/
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
