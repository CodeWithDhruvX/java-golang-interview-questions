package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 98. Validate Binary Search Tree
// Time: O(N), Space: O(H) where H is tree height (recursion stack)
func isValidBST(root *TreeNode) bool {
	return validateBST(root, nil, nil)
}

func validateBST(node *TreeNode, min, max *TreeNode) bool {
	if node == nil {
		return true
	}
	
	// Check current node value against bounds
	if min != nil && node.Val <= min.Val {
		return false
	}
	if max != nil && node.Val >= max.Val {
		return false
	}
	
	// Recursively check left and right subtrees
	return validateBST(node.Left, min, node) && validateBST(node.Right, node, max)
}

// Alternative approach using in-order traversal
func isValidBSTInorder(root *TreeNode) bool {
	var prev *TreeNode
	
	var inorder func(*TreeNode) bool
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		
		// Check left subtree
		if !inorder(node.Left) {
			return false
		}
		
		// Check current node
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		
		// Check right subtree
		return inorder(node.Right)
	}
	
	return inorder(root)
}

// Iterative approach using stack
func isValidBSTIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}
	
	stack := []*TreeNode{}
	var prev *TreeNode
	current := root
	
	for len(stack) > 0 || current != nil {
		// Go as far left as possible
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		
		// Process node
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if prev != nil && current.Val <= prev.Val {
			return false
		}
		prev = current
		
		// Move to right subtree
		current = current.Right
	}
	
	return true
}

// Range-based validation
func isValidBSTRange(root *TreeNode) bool {
	return isValidBSTRangeHelper(root, -2147483648, 2147483647) // 32-bit int range
}

func isValidBSTRangeHelper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	
	if node.Val <= min || node.Val >= max {
		return false
	}
	
	return isValidBSTRangeHelper(node.Left, min, node.Val) &&
		isValidBSTRangeHelper(node.Right, node.Val, max)
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
	if root == nil {
		return
	}
	
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{Val: val}
		} else {
			insertBST(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else {
			insertBST(root.Right, val)
		}
	}
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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: BST Validation with Range Constraints
- **Range Validation**: Each node must be within valid range
- **Recursive Bounds**: Pass min/max constraints down the tree
- **In-order Traversal**: BST produces sorted sequence
- **Property Checking**: Left < root < right for all nodes

## 2. PROBLEM CHARACTERISTICS
- **BST Properties**: Left subtree < root < right subtree
- **Global Constraints**: Must hold for entire subtree, not just immediate children
- **Tree Traversal**: Need to visit all nodes
- **Validation**: Check if tree satisfies BST rules

## 3. SIMILAR PROBLEMS
- Recover Binary Search Tree (LeetCode 99) - Fix BST violations
- Kth Smallest Element in BST (LeetCode 230) - BST property utilization
- Convert Sorted Array to BST (LeetCode 108) - BST construction
- Binary Tree Inorder Traversal (LeetCode 94) - BST property demonstration

## 4. KEY OBSERVATIONS
- **Range Propagation**: Each node inherits constraints from ancestors
- **In-order Property**: Valid BST produces strictly increasing sequence
- **Local vs Global**: Local check insufficient, need global constraints
- **Multiple Approaches**: Range validation, in-order traversal, iterative

## 5. VARIATIONS & EXTENSIONS
- **Allow Duplicates**: Handle equal values (left ≤ root < right)
- **Custom Ranges**: Different comparison functions
- **Tree Reconstruction**: Build BST from validation failures
- **Multiple Trees**: Validate forest of BSTs

## 6. INTERVIEW INSIGHTS
- Always clarify: "Duplicate values allowed? Tree size? Data type limits?"
- Edge cases: empty tree, single node, duplicate values
- Time complexity: O(N) where N=number of nodes
- Space complexity: O(H) where H=tree height (recursion stack)
- Key insight: need global constraints, not just local checks

## 7. COMMON MISTAKES
- Only checking immediate children (insufficient)
- Wrong boundary conditions (≤ vs <)
- Not handling empty tree correctly
- Integer overflow in range validation
- Missing duplicate value handling

## 8. OPTIMIZATION STRATEGIES
- **Range Validation**: O(N) time, O(H) space - optimal
- **In-order Traversal**: O(N) time, O(H) space - elegant
- **Iterative In-order**: O(N) time, O(H) space - no recursion
- **Early Termination**: Stop on first violation

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like validating a family tree with age rules:**
- Each person must be older than all ancestors in left lineage
- Each person must be younger than all ancestors in right lineage
- Rules must hold for entire family history, not just immediate parents
- Like checking if a hierarchical organization follows seniority rules
- Each level inherits constraints from above

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary tree with integer values
2. **Goal**: Determine if tree is a valid BST
3. **Rules**: Left subtree < root < right subtree for all nodes
4. **Output**: Boolean indicating BST validity

#### Phase 2: Key Insight Recognition
- **"Global constraints needed"** → Local checks insufficient
- **"Range propagation"** → Each node inherits min/max bounds
- **"In-order property"** → Valid BST produces sorted sequence
- **"Multiple approaches"** → Range validation, in-order, iterative

#### Phase 3: Strategy Development
```
Human thought process:
"I need to validate BST properties.
This requires global constraints:

Range Validation Approach:
1. Root can be any value (min=-∞, max=+∞)
2. Left child must be < root value (max=root.val)
3. Right child must be > root value (min=root.val)
4. Propagate constraints down the tree
5. Violation occurs if node violates inherited bounds

In-order Traversal Approach:
1. Valid BST produces strictly increasing sequence
2. Perform in-order traversal
3. Check if sequence is strictly increasing
4. Any non-increasing pair indicates violation

Both approaches ensure global BST properties!"
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return true (trivially valid)
- **Single node**: Return true (no violations possible)
- **Duplicate values**: Clarify if allowed (typically not)
- **Integer limits**: Handle min/max int values

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree: [10, 5, 15, 6, 20] (invalid BST)

Human thinking:
"Range Validation:
Root (10): range (-∞, +∞) ✓
Left child (5): range (-∞, 10) ✓
Right child (15): range (10, +∞) ✓

Left subtree of 5:
- No left child
- Right child (6): range (5, 10) ✓

Right subtree of 15:
- Left child (6): range (10, 15) ✗ VIOLATION!
  6 is not > 10 (min constraint)

In-order Traversal:
Sequence: 5, 6, 10, 15, 20
Wait, this looks increasing...
Actually tree is: [10, 5, 15, null, null, 6, 20]
In-order: 5, 10, 6, 15, 20
6 < 10 ✗ VIOLATION!

Both approaches detect the violation ✓"
```

#### Phase 6: Intuition Validation
- **Why range validation works**: Propagates global constraints
- **Why in-order works**: BST property guarantees sorted sequence
- **Why O(N)**: Each node visited once
- **Why O(H) space**: Recursion depth equals tree height

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just check children?"** → Need global constraints
2. **"Should I use BFS?"** → BFS doesn't help with BST validation
3. **"What about duplicates?"** → Clarify duplicate handling policy
4. **"Can I optimize further?"** → O(N) is already optimal
5. **"What about iterative approach?"** → Possible but more complex

### Real-World Analogy
**Like validating a company hierarchy with salary bands:**
- Each manager's salary must be within allowed range
- Subordinates must earn less than managers above them
- Rules must hold for entire organization, not just immediate reporting
- Like checking if salary bands are properly enforced
- Each level inherits constraints from organizational structure

### Human-Readable Pseudocode
```
function isValidBST(root):
    return validateBST(root, -∞, +∞)

function validateBST(node, min_val, max_val):
    if node is null:
        return true
    
    // Check current node against inherited bounds
    if node.val <= min_val or node.val >= max_val:
        return false
    
    // Validate subtrees with updated bounds
    left_valid = validateBST(node.left, min_val, node.val)
    right_valid = validateBST(node.right, node.val, max_val)
    
    return left_valid and right_valid

// Alternative: In-order traversal
function isValidBSTInorder(root):
    prev_val = -∞
    
    function inorder(node):
        if node is null:
            return true
        
        if not inorder(node.left):
            return false
        
        if node.val <= prev_val:
            return false
        prev_val = node.val
        
        return inorder(node.right)
    
    return inorder(root)
```

### Execution Visualization

### Example: Tree = [10, 5, 15, null, null, 6, 20] (Invalid)
```
Tree Structure:
        10
       /  \
      5    15
           /  \
          6   20

Range Validation:
Root (10): range (-∞, +∞) ✓
Left child (5): range (-∞, 10) ✓
Right child (15): range (10, +∞) ✓

Left subtree of 5: No violations ✓

Right subtree of 15:
Left child (6): range (10, 15) ✗ VIOLATION!
6 ≤ 10 (violates min constraint)

In-order Traversal:
Visit order: 5, 10, 6, 15, 20
Sequence: [5, 10, 6, 15, 20]
6 < 10 ✗ VIOLATION!

Both methods detect invalid BST ✓
```

### Key Visualization Points:
- **Range Propagation**: Constraints flow down the tree
- **Global Validation**: Each node checked against all ancestors
- **In-order Property**: BST produces sorted sequence
- **Violation Detection**: Multiple ways to detect violations

### Memory Layout Visualization:
```
Recursive Stack During Validation:
Call Stack:          Min/Max Ranges:
validateBST(10, -∞, +∞)
├── validateBST(5, -∞, 10)
│   ├── validateBST(null, -∞, 5) ✓
│   └── validateBST(null, 5, 10) ✓
└── validateBST(15, 10, +∞)
    ├── validateBST(6, 10, 15) ✗ VIOLATION!
    └── validateBST(20, 15, +∞) ✓

In-order Traversal State:
Visited: [5, 10]
Current: 6
Previous: 10
Comparison: 6 ≤ 10 ✗ VIOLATION!
```

### Time Complexity Breakdown:
- **Range Validation**: O(N) time (each node visited once)
- **In-order Traversal**: O(N) time (each node visited once)
- **Space**: O(H) where H=tree height (recursion stack)
- **Optimal**: Cannot do better than visiting all nodes

### Alternative Approaches:

#### 1. In-order Traversal (O(N) time, O(H) space)
```go
func isValidBSTInorder(root *TreeNode) bool {
    var prev *TreeNode
    
    var inorder func(*TreeNode) bool
    inorder = func(node *TreeNode) bool {
        if node == nil {
            return true
        }
        
        if !inorder(node.Left) {
            return false
        }
        
        if prev != nil && node.Val <= prev.Val {
            return false
        }
        prev = node
        
        return inorder(node.Right)
    }
    
    return inorder(root)
}
```
- **Pros**: Elegant, uses BST property directly
- **Cons**: Requires global state (previous node)

#### 2. Iterative In-order (O(N) time, O(H) space)
```go
func isValidBSTIterative(root *TreeNode) bool {
    stack := []*TreeNode{}
    var prev *TreeNode
    current := root
    
    for len(stack) > 0 || current != nil {
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }
        
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if prev != nil && current.Val <= prev.Val {
            return false
        }
        prev = current
        
        current = current.Right
    }
    
    return true
}
```
- **Pros**: No recursion, more memory efficient
- **Cons**: More complex implementation

#### 3. Range with Integer Bounds (O(N) time, O(H) space)
```go
func isValidBSTRange(root *TreeNode) bool {
    return isValidBSTRangeHelper(root, -2147483648, 2147483647)
}

func isValidBSTRangeHelper(node *TreeNode, min, max int) bool {
    if node == nil {
        return true
    }
    
    if node.Val <= min || node.Val >= max {
        return false
    }
    
    return isValidBSTRangeHelper(node.Left, min, node.Val) &&
           isValidBSTRangeHelper(node.Right, node.Val, max)
}
```
- **Pros**: No pointer arithmetic, simpler bounds
- **Cons**: Limited by integer range

### Extensions for Interviews:
- **Allow Duplicates**: Handle equal values (left ≤ root < right)
- **Custom Comparators**: Support different comparison functions
- **Tree Reconstruction**: Build valid BST from violations
- **Multiple Trees**: Validate forest of BSTs
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Validate Binary Search Tree ===")
	
	testCases := []struct {
		tree       []interface{}
		isBST      bool
		description string
	}{
		{[]interface{}{2, 1, 3}, true, "Valid BST"},
		{[]interface{}{5, 1, 4, nil, nil, 3, 6}, false, "Invalid BST"},
		{[]interface{}{1, 1}, false, "Duplicate values"},
		{[]interface{}{}, true, "Empty tree"},
		{[]interface{}{1}, true, "Single node"},
		{[]interface{}{10, 5, 15, 3, 7, 12, 18}, true, "Perfect BST"},
		{[]interface{}{10, 5, 15, 3, 12, 12, 18}, false, "Right subtree violation"},
		{[]interface{}{10, 5, 15, 1, 6, 12, 20}, true, "Valid BST with varied depths"},
		{[]interface{}{10, 5, 15, 1, 11, 12, 20}, false, "Left subtree violation"},
		{[]interface{}{2, 2, 2}, false, "All same values"},
	}
	
	for i, tc := range testCases {
		root := createTree(tc.tree)
		
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Tree: %v\n", tc.tree)
		
		result1 := isValidBST(root)
		result2 := isValidBSTInorder(root)
		result3 := isValidBSTIterative(root)
		result4 := isValidBSTRange(root)
		
		fmt.Printf("  Recursive: %t\n", result1)
		fmt.Printf("  In-order: %t\n", result2)
		fmt.Printf("  Iterative: %t\n", result3)
		fmt.Printf("  Range-based: %t\n\n", result4)
	}
	
	// Test with manually created BST
	fmt.Println("=== Testing with Created BST ===")
	bst := createBST([]int{50, 30, 70, 20, 40, 60, 80})
	fmt.Printf("Created BST - Valid: %t\n", isValidBST(bst))
	
	// Test with invalid tree
	invalidTree := &TreeNode{
		Val: 10,
		Left: &TreeNode{Val: 5},
		Right: &TreeNode{
			Val: 15,
			Left:  &TreeNode{Val: 6}, // This violates BST property
			Right: &TreeNode{Val: 20},
		},
	}
	fmt.Printf("Invalid tree - Valid: %t\n", isValidBST(invalidTree))
	
	// Test edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	
	// Large BST
	largeBST := createBST([]int{50, 25, 75, 12, 37, 62, 87, 6, 18, 31, 43, 56, 68, 81, 93})
	fmt.Printf("Large BST - Valid: %t\n", isValidBST(largeBST))
	
	// Tree with negative values
	negBST := createBST([]int{0, -10, 10, -20, -5, 5, 20})
	fmt.Printf("BST with negatives - Valid: %t\n", isValidBST(negBST))
	
	// Tree with maximum int values
	maxBST := createBST([]int{1000000, 500000, 1500000})
	fmt.Printf("BST with large values - Valid: %t\n", isValidBST(maxBST))
}
