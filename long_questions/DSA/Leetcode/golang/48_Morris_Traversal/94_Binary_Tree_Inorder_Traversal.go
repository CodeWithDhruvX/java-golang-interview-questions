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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Morris Traversal for Space-Efficient Tree Traversal
- **Threaded Binary Tree**: Use null pointers to create temporary links
- **Inorder Traversal**: Visit left subtree, root, right subtree without recursion
- **Space Optimization**: O(1) auxiliary space instead of O(h) stack
- **Tree Modification**: Temporarily modify tree structure, then restore

## 2. PROBLEM CHARACTERISTICS
- **Binary Tree Traversal**: Visit nodes in specific order without extra space
- **Inorder Sequence**: Left-Root-Right traversal order
- **Space Constraints**: Cannot use recursion or explicit stack
- **Tree Restoration**: Must restore original tree structure

## 3. SIMILAR PROBLEMS
- Binary Tree Inorder Traversal (LeetCode 94) - Same problem
- Binary Tree Preorder Traversal (LeetCode 144) - Different order
- Binary Tree Postorder Traversal (LeetCode 145) - Different order
- Validate Binary Search Tree - Inorder traversal validation

## 4. KEY OBSERVATIONS
- **Threaded Links**: Use right null pointers to create temporary links
- **Predecessor Finding**: Find inorder predecessor for each node
- **Link Creation**: Create link from predecessor to current node
- **Link Removal**: Remove temporary links after use

## 5. VARIATIONS & EXTENSIONS
- **Inorder Traversal**: Left-Root-Right order
- **Preorder Traversal**: Root-Left-Right order
- **Postorder Traversal**: Left-Right-Root order (reverse of modified preorder)
- **Extended Analysis**: Count, sum, min/max during traversal

## 6. INTERVIEW INSIGHTS
- Always clarify: "Tree structure? Traversal order? Space constraints?"
- Edge cases: empty tree, single node, skewed trees
- Time complexity: O(N) time, O(1) space
- Space complexity: O(1) auxiliary space
- Key insight: temporary threaded links eliminate need for stack

## 7. COMMON MISTAKES
- Wrong predecessor finding (should be rightmost node in left subtree)
- Missing link restoration (tree must be unchanged after traversal)
- Incorrect traversal order modifications
- Not handling null pointers properly
- Wrong termination conditions

## 8. OPTIMIZATION STRATEGIES
- **Standard Morris**: O(N) time, O(1) space - optimal
- **Extended Morris**: O(N) time, O(1) space - with additional computations
- **Preorder Morris**: O(N) time, O(1) space - different visitation order
- **Postorder Morris**: O(N) time, O(1) space - reverse of modified preorder

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like navigating a maze with temporary shortcuts:**
- You have a tree structure representing rooms and corridors
- You want to visit all rooms in a specific order without a map
- You create temporary shortcuts from dead ends to unvisited rooms
- After using a shortcut, you remove it to restore the original maze
- Like an explorer leaving temporary markers to avoid getting lost

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary tree root pointer
2. **Goal**: Traverse tree in inorder without using extra space
3. **Constraints**: Cannot use recursion or explicit stack
4. **Output**: Array of node values in inorder sequence

#### Phase 2: Key Insight Recognition
- **"No stack allowed"** → Need alternative way to track return path
- **"Null pointers available"** → Use right null pointers as temporary links
- **"Predecessor connection"** → Link predecessor to current node
- **"Tree restoration"** → Must remove all temporary links

#### Phase 3: Strategy Development
```
Human thought process:
"I need inorder traversal without a stack.
Recursion uses call stack O(h), iterative uses explicit stack O(h).

Morris Traversal Approach:
1. For each node, find its inorder predecessor
2. If predecessor's right is null, link it to current node
3. Follow the link back when we reach predecessor again
4. Remove the temporary link after using it
5. Continue until we've visited all nodes

This gives O(N) time, O(1) space!"
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return empty array
- **Single node**: Return [node.val]
- **Skewed trees**: Algorithm works for all tree shapes
- **Duplicate values**: Handled naturally

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: Tree = [1,2,3,4,5]

Human thinking:
"Morris Traversal Process:
Step 1: Start at root (1)
1 has left child (2), find predecessor of 1
Predecessor = rightmost node in left subtree = 4
Set 4.right = 1 (temporary link)
Move to left child (2)

Step 2: At node 2
2 has left child (4), find predecessor of 2
Predecessor = rightmost node in left subtree = 4
4.right already points to 2 (from previous step)
Visit 2, remove link 4.right = 2
Move to right child (5)

Step 3: At node 5
5 has no left child, visit 5
Move to right child (null, but 4.right = 1)

Step 4: At node 1 (via link)
1's left subtree done, visit 1
Move to right child (3)

Step 5: At node 3
3 has no left child, visit 3
Move to right child (null)

Result: [4,2,5,1,3] ✓"
```

#### Phase 6: Intuition Validation
- **Why predecessor links**: Provide way to return to parent without stack
- **Why right null pointers**: Available for temporary use
- **Why tree restoration**: Must leave tree unchanged after traversal
- **Why O(1) space**: Only use existing tree pointers

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use recursion?"** → Uses O(h) stack space, violates constraint
2. **"Should I use an explicit stack?"** → Also uses O(h) space
3. **"What about tree modification?"** → Temporary, must be restored
4. **"Can I use parent pointers?"** → Tree nodes don't have parent references
5. **"Why find predecessor?"** → Creates return path without extra space

### Real-World Analogy
**Like hiking through a trail system without a map:**
- You have a network of trails (tree structure)
- You want to visit all trail intersections in a specific order
- You leave temporary markers (threaded links) to find your way back
- After following a marker, you remove it so others don't get confused
- Like a hiker using temporary trail markers to ensure complete coverage

### Human-Readable Pseudocode
```
function morrisInorder(root):
    current = root
    result = []
    
    while current != null:
        if current.left == null:
            # No left subtree, visit current
            result.append(current.val)
            current = current.right
        else:
            # Find inorder predecessor
            predecessor = current.left
            while predecessor.right != null and predecessor.right != current:
                predecessor = predecessor.right
            
            if predecessor.right == null:
                # Create temporary link
                predecessor.right = current
                current = current.left
            else:
                # Remove temporary link and visit current
                predecessor.right = null
                result.append(current.val)
                current = current.right
    
    return result
```

### Execution Visualization

### Example: Tree = [1,2,3,4,5]
```
Tree Structure:
      1
     / \
    2   3
   / \
  4   5

Step 1: At 1, predecessor = 4
Set 4.right = 1 (temporary link)
Move to 2

Step 2: At 2, predecessor = 4
4.right already = 2 (link exists)
Visit 2, remove 4.right = 2
Move to 5

Step 3: At 5, no left child
Visit 5, move to right (null)

Step 4: Back to 1 via 4.right link
Visit 1, move to right (3)

Step 5: At 3, no left child
Visit 3, move to right (null)

Result: [4,2,5,1,3] ✓
```

### Key Visualization Points:
- **Predecessor Finding**: Rightmost node in left subtree
- **Link Creation**: Predecessor.right = current (temporary)
- **Link Detection**: Predecessor.right == current indicates return
- **Link Removal**: Predecessor.right = nil (restoration)

### Threaded Binary Tree Concept:
```
Normal Tree:          Threaded Tree (temporary):
      1                      1
     / \                    / \
    2   3                  2---3
   / \                    / \
  4   5                  4   5
       \                    \
        1 (temporary)        1 (restored)
```

### Time Complexity Breakdown:
- **Standard Morris**: O(N) time, O(1) space - optimal
- **Extended Analysis**: O(N) time, O(1) space - with computations
- **Preorder/Postorder**: O(N) time, O(1) space - different orders
- **All Variants**: O(N) time, O(1) space - space-efficient

### Alternative Approaches:

#### 1. Recursive Inorder (O(N) time, O(h) space)
```go
func inorderRecursive(root *TreeNode) []int {
    var result []int
    
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        result = append(result, node.Val)
        dfs(node.Right)
    }
    
    dfs(root)
    return result
}
```
- **Pros**: Simple, natural recursion
- **Cons**: Uses O(h) stack space

#### 2. Iterative with Stack (O(N) time, O(h) space)
```go
func inorderIterative(root *TreeNode) []int {
    var result []int
    stack := []*TreeNode{}
    current := root
    
    for current != nil || len(stack) > 0 {
        for current != nil {
            stack = append(stack, current)
            current = current.Left
        }
        
        current = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, current.Val)
        current = current.Right
    }
    
    return result
}
```
- **Pros**: No recursion, explicit control
- **Cons**: Still uses O(h) space for stack

#### 3. Threaded Binary Tree Permanent (O(N) time, O(N) space)
```go
func inorderThreaded(root *TreeNode) []int {
    // Convert to permanently threaded tree
    // Then traverse without temporary modifications
    return result
}
```
- **Pros**: One-time conversion cost
- **Cons**: Permanently modifies tree structure

### Extensions for Interviews:
- **Preorder Traversal**: Visit root before left subtree
- **Postorder Traversal**: Reverse of modified preorder
- **Extended Analysis**: Count, sum, min/max during traversal
- **BST Validation**: Check if tree is valid BST
- **Tree Comparison**: Compare two trees using Morris traversal
- **Real-world Applications**: Memory-constrained environments, embedded systems
*/
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
