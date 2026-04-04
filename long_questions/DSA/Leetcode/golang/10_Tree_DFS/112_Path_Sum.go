package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 112. Path Sum
// Time: O(N), Space: O(H) where H is tree height (recursion stack)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	
	// If we reach a leaf node
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	
	// Recursively check left and right subtrees
	remainingSum := targetSum - root.Val
	return hasPathSum(root.Left, remainingSum) || hasPathSum(root.Right, remainingSum)
}

// Iterative approach using stack
func hasPathSumIterative(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	
	stack := []struct {
		node *TreeNode
		sum  int
	}{{root, root.Val}}
	
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		node := current.node
		sum := current.sum
		
		// If we reach a leaf node
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				return true
			}
		}
		
		if node.Right != nil {
			stack = append(stack, struct {
				node *TreeNode
				sum  int
			}{node.Right, sum + node.Right.Val})
		}
		
		if node.Left != nil {
			stack = append(stack, struct {
				node *TreeNode
				sum  int
			}{node.Left, sum + node.Left.Val})
		}
	}
	
	return false
}

// Helper function to create a binary tree from slice
func createTree(nums []interface{}) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	
	nodes := make([]*TreeNode, len(nums))
	for i, val := range nums {
		if val == nil {
			nodes[i] = nil
		} else {
			nodes[i] = &TreeNode{Val: val.(int)}
		}
	}
	
	for i := 0; i < len(nums); i++ {
		left := 2*i + 1
		right := 2*i + 2
		
		if left < len(nums) {
			nodes[i].Left = nodes[left]
		}
		if right < len(nums) {
			nodes[i].Right = nodes[right]
		}
	}
	
	return nodes[0]
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Tree DFS with Path Accumulation
- **Recursive Traversal**: Visit each node exactly once
- **Path Sum Tracking**: Subtract node value from target as we descend
- **Leaf Check**: Verify if remaining sum equals leaf value
- **Backtracking**: Natural through recursion stack unwinding

## 2. PROBLEM CHARACTERISTICS
- **Root-to-Leaf Path**: Must start at root and end at leaf
- **Sum Constraint**: Path values must equal target exactly
- **Binary Tree**: Each node has at most 2 children
- **Existence Check**: Return true/false, not the actual path

## 3. SIMILAR PROBLEMS
- Path Sum II (LeetCode 113) - Return all valid paths
- Path Sum III (LeetCode 437) - Any path starting anywhere
- Path Sum IV (LeetCode 666) - Tree encoded as array
- Sum Root to Leaf Numbers (LeetCode 129)

## 4. KEY OBSERVATIONS
- **Target reduction**: Subtract current node value from target
- **Leaf definition**: Node with no children
- **Early termination**: Can stop when leaf reached
- **Path uniqueness**: Only need one valid path

## 5. VARIATIONS & EXTENSIONS
- **Return All Paths**: Collect all valid root-to-leaf paths
- **Any Path**: Allow paths starting from any node
- **Negative Values**: Handle negative numbers in tree
- **Multiple Queries**: Process multiple target sums

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can tree have negative values? What constitutes a leaf?"
- Edge cases: empty tree, single node, negative values
- Space complexity: O(H) where H is tree height (recursion stack)
- Time complexity: O(N) - visit each node once

## 7. COMMON MISTAKES
- Not checking leaf condition properly
- Using addition instead of subtraction for target
- Forgetting to handle empty tree case
- Not considering negative values affect on pruning
- Confusing path sum with subtree sum

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(H) space
- **Iterative with stack**: Can avoid recursion but more complex
- **Early pruning**: Not applicable without additional constraints
- **Memoization**: Not beneficial for tree structure

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like navigating a maze with a budget:**
- You have a maze with paths (tree structure)
- You start at the entrance (root) with a budget (target sum)
- Each room (node) costs some amount (node value)
- You want to reach an exit (leaf) spending exactly your budget
- As you enter each room, subtract its cost from your remaining budget
- If you reach an exit with exactly zero budget, you found a valid path

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary tree root and target sum
2. **Goal**: Determine if any root-to-leaf path sums to target
3. **Output**: Boolean indicating path existence
4. **Constraint**: Path must go from root to leaf (no stopping mid-way)

#### Phase 2: Key Insight Recognition
- **"Target reduction strategy"** → Subtract node values as we descend
- **"Leaf verification"** → Only check sum at leaf nodes
- **"Recursive propagation"** → Pass remaining sum to children
- **"Early success"** → Return true immediately when valid path found

#### Phase 3: Strategy Development
```
Human thought process:
"I need to check if there's a path from root to leaf that sums to target.
As I traverse down the tree, I'll subtract each node's value from the target.
When I reach a leaf, I'll check if the remaining target equals the leaf's value.
If any path works, I'll return true immediately."
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return false (no path exists)
- **Single node**: Check if node value equals target
- **Negative values**: Still works, just affects intermediate sums
- **Large target**: May cause integer overflow in some languages

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:     5
         / \
        4   8
       /   / \
      11  13  4
     /  \      \
    7    2      1

Target: 22

Human thinking:
"I'll traverse from root, tracking remaining sum:

Root (5): remaining = 22 - 5 = 17
Go left to (4): remaining = 17 - 4 = 13
Go left to (11): remaining = 13 - 11 = 2
Go left to (7): remaining = 2 - 7 = -5
Leaf (7): remaining ≠ 0, backtrack

Back to (11): go right to (2): remaining = 2 - 2 = 0
Leaf (2): remaining = 0 ✓ Found valid path!
Return true"
```

#### Phase 6: Intuition Validation
- **Why subtraction works**: Maintains remaining budget for path
- **Why leaf check only**: Path must end at leaf per problem
- **Why O(N) time**: Each node visited at most once
- **Why recursion natural**: Tree structure lends itself to DFS

### Common Human Pitfalls & How to Avoid Them
1. **"Why not sum up instead of subtract?"** → Both work, subtraction is cleaner
2. **"Should I check internal nodes too?"** → No, path must end at leaf
3. **"What about negative values?"** → Algorithm still works correctly
4. **"Can I prune branches early?"** → Not without additional constraints

### Real-World Analogy
**Like planning a road trip with a fuel budget:**
- You have a map of connected cities (tree structure)
- You start at home (root) with a fuel budget (target sum)
- Each road segment (node) consumes some fuel (node value)
- You want to reach a destination city (leaf) using exactly all fuel
- As you travel to each city, subtract fuel consumption from budget
- If you reach a destination with empty tank, you found a valid route

### Human-Readable Pseudocode
```
function hasPathSum(node, target):
    if node is nil:
        return false
    
    // If we're at a leaf node
    if node.left is nil and node.right is nil:
        return node.value == target
    
    // Check subtrees with reduced target
    remainingTarget = target - node.value
    return hasPathSum(node.left, remainingTarget) or 
           hasPathSum(node.right, remainingTarget)
```

### Execution Visualization

### Example Tree and Target = 22:
```
     5
    / \
   4   8
  /   / \
 11  13  4
/  \      \
7   2      1
```

### Recursive Call Flow:
```
hasPathSum(5, 22)
├── hasPathSum(4, 17)
│   └── hasPathSum(11, 13)
│       ├── hasPathSum(7, 2) → false (7 ≠ 2)
│       └── hasPathSum(2, 2) → true (2 = 2) ✓
└── returns true (found valid path)
```

### Key Visualization Points:
- **Target reduction**: Subtract node value at each level
- **Leaf check**: Only verify when reaching leaf node
- **Early success**: Return true immediately when found
- **Backtracking**: Natural through recursion return

### Memory Layout Visualization:
```
Call Stack (top to bottom):
hasPathSum(5, 22)
hasPathSum(4, 17)
hasPathSum(11, 13)
hasPathSum(2, 2) ← returns true
```

### Time Complexity Breakdown:
- **Each node visited**: O(1) work per node
- **Total nodes**: N
- **Total time**: O(N) - may stop early if path found
- **Space**: O(H) - maximum recursion depth

### Alternative Approaches:

#### 1. Iterative with Stack (O(N) time, O(H) space)
```go
func hasPathSumIterative(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
    stack := []struct {
        node *TreeNode
        sum  int
    }{{root, root.Val}}
    
    for len(stack) > 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        node := current.node
        sum := current.sum
        
        // If we reach a leaf node
        if node.Left == nil && node.Right == nil {
            if sum == targetSum {
                return true
            }
        }
        
        if node.Right != nil {
            stack = append(stack, struct {
                node *TreeNode
                sum  int
            }{node.Right, sum + node.Right.Val})
        }
        
        if node.Left != nil {
            stack = append(stack, struct {
                node *TreeNode
                sum  int
            }{node.Left, sum + node.Left.Val})
        }
    }
    
    return false
}
```
- **Pros**: No recursion stack, same complexity
- **Cons**: More complex implementation

#### 2. Path Collection (O(N) time, O(H) space)
```go
func pathSum(root *TreeNode, targetSum int) [][]int {
    var result [][]int
    var current []int
    
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, remaining int) {
        if node == nil {
            return
        }
        
        current = append(current, node.Val)
        
        if node.Left == nil && node.Right == nil && remaining == node.Val {
            // Found valid path, make a copy
            path := make([]int, len(current))
            copy(path, current)
            result = append(result, path)
        }
        
        dfs(node.Left, remaining-node.Val)
        dfs(node.Right, remaining-node.Val)
        
        current = current[:len(current)-1] // Backtrack
    }
    
    dfs(root, targetSum)
    return result
}
```
- **Pros**: Returns all valid paths
- **Cons**: More space for path storage

#### 3. BFS with Path Tracking (O(N) time, O(W) space)
```go
func hasPathSumBFS(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
    queue := []struct {
        node *TreeNode
        sum  int
    }{{root, root.Val}}
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        node := current.node
        sum := current.sum
        
        if node.Left == nil && node.Right == nil && sum == targetSum {
            return true
        }
        
        if node.Left != nil {
            queue = append(queue, struct {
                node *TreeNode
                sum  int
            }{node.Left, sum + node.Left.Val})
        }
        
        if node.Right != nil {
            queue = append(queue, struct {
                node *TreeNode
                sum  int
            }{node.Right, sum + node.Right.Val})
        }
    }
    
    return false
}
```
- **Pros**: Level-order traversal
- **Cons**: Uses more memory for wide trees

### Extensions for Interviews:
- **Path Sum II**: Return all valid root-to-leaf paths
- **Path Sum III**: Find any path (not just root-to-leaf)
- **Negative Values**: Handle trees with negative numbers
- **Multiple Targets**: Check multiple target sums efficiently
- **Path Count**: Count number of valid paths instead of boolean
*/
func main() {
	// Test cases
	testCases := []struct {
		tree      []interface{}
		targetSum int
	}{
		{[]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}, 22},
		{[]interface{}{1, 2, 3}, 5},
		{[]interface{}{1, 2, 3}, 4},
		{[]interface{}{}, 0},
		{[]interface{}{1}, 1},
		{[]interface{}{1}, 2},
		{[]interface{}{-2, nil, -3}, -5},
		{[]interface{}{1, -2, 3, -4, nil, nil, nil, 5}, -1},
		{[]interface{}{0}, 0},
		{[]interface{}{1, 2}, 3},
	}
	
	for i, tc := range testCases {
		root := createTree(tc.tree)
		result1 := hasPathSum(root, tc.targetSum)
		result2 := hasPathSumIterative(root, tc.targetSum)
		
		fmt.Printf("Test Case %d: tree=%v, target=%d\n", i+1, tc.tree, tc.targetSum)
		fmt.Printf("  Recursive: %t, Iterative: %t\n\n", result1, result2)
	}
}
