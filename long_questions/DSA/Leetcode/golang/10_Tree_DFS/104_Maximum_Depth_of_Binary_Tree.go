package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 104. Maximum Depth of Binary Tree
// Time: O(N), Space: O(H) where H is tree height (recursion stack)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	
	return max(leftDepth, rightDepth) + 1
}

// Iterative BFS approach
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	
	queue := []*TreeNode{root}
	depth := 0
	
	for len(queue) > 0 {
		levelSize := len(queue)
		depth++
		
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	
	return depth
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Tree Depth-First Search (DFS)
- **Recursive Traversal**: Visit each node exactly once
- **Base Case**: Return 0 for nil nodes
- **Recursive Case**: 1 + max(leftDepth, rightDepth)
- **Depth Calculation**: Accumulate depth as we return from recursion

## 2. PROBLEM CHARACTERISTICS
- **Binary Tree**: Each node has at most 2 children
- **Depth Definition**: Number of nodes from root to deepest leaf
- **Tree Structure**: May be unbalanced, empty, or single node
- **Recursive Nature**: Natural fit for DFS approach

## 3. SIMILAR PROBLEMS
- Minimum Depth of Binary Tree (LeetCode 111)
- Balanced Binary Tree (LeetCode 110)
- Invert Binary Tree (LeetCode 226)
- Same Tree (LeetCode 100)

## 4. KEY OBSERVATIONS
- **Depth propagation**: Each node contributes 1 to depth
- **Maximum selection**: Take max of left and right subtrees
- **Leaf termination**: Depth calculation stops at leaves
- **Linear traversal**: Visit each node exactly once

## 5. VARIATIONS & EXTENSIONS
- **Minimum Depth**: Find shortest path to leaf
- **Balanced Check**: Verify height difference ≤ 1
- **Level Order**: BFS alternative for depth
- **N-ary Tree**: Extend to nodes with multiple children

## 6. INTERVIEW INSIGHTS
- Always clarify: "Is this a binary tree? Can it be empty?"
- Edge cases: empty tree, single node, unbalanced tree
- Space complexity: O(H) where H is tree height (recursion stack)
- Time complexity: O(N) - visit each node once

## 7. COMMON MISTAKES
- Not handling nil root case properly
- Forgetting to add 1 for current node
- Using sum instead of max for depth
- Stack overflow on very deep trees
- Not considering iterative BFS alternative

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(H) space
- **Iterative BFS**: Can use queue to avoid recursion stack
- **Tail recursion**: Not applicable in Go for tree traversal
- **Early pruning**: Not applicable for depth calculation

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like counting the levels in a company hierarchy:**
- You have an organizational chart (tree structure)
- You want to know the longest chain of command
- Start from the CEO (root) and count down to the deepest employee
- For each manager, find their longest reporting chain
- The company depth is 1 + the longest chain among all departments
- Keep track of the maximum depth as you explore the organization

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary tree root node
2. **Goal**: Find maximum depth (number of nodes from root to deepest leaf)
3. **Output**: Integer representing maximum depth
4. **Constraint**: Tree may be empty or unbalanced

#### Phase 2: Key Insight Recognition
- **"Recursive structure"** → Tree depth = 1 + max(subtree depths)
- **"Base case"** → Empty tree has depth 0
- **"Divide and conquer"** → Solve left and right subtrees independently
- **"Combine results"** → Take maximum and add current node

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the deepest level of the tree.
For any node, the depth is 1 (for itself) plus the maximum
of the depths of its left and right subtrees.
If a node is nil, its depth is 0.
I'll recursively calculate depths and combine results."
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return 0 (no nodes)
- **Single node**: Return 1 (just the root)
- **Unbalanced tree**: Handle different depths naturally
- **Very deep tree**: Consider stack overflow limits

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:     3
         / \
        9   20
           /  \
          15   7

Human thinking:
"I'll calculate depth from bottom up:

Node 9: no children → depth = 1
Node 15: no children → depth = 1
Node 7: no children → depth = 1

Node 20: left depth=1, right depth=1
         max(1,1) + 1 = 2

Node 3: left depth=1, right depth=2
        max(1,2) + 1 = 3

Final answer: 3"
```

#### Phase 6: Intuition Validation
- **Why recursion works**: Tree has natural recursive structure
- **Why max is used**: Want deepest path, not sum of all paths
- **Why O(N) time**: Each node visited exactly once
- **Why O(H) space**: Recursion depth equals tree height

### Common Human Pitfalls & How to Avoid Them
1. **"Why not count edges instead of nodes?"** → Problem specifies nodes, clarify in interview
2. **"Should I use BFS instead?"** → BFS works but uses more memory for wide trees
3. **"What about very deep trees?"** → Consider iterative approach to avoid stack overflow
4. **"Can I optimize further?"** → Already optimal O(N) time

### Real-World Analogy
**Like finding the longest family generation chain:**
- You have a family tree showing parent-child relationships
- You want to know the longest generation line
- Start from the oldest ancestor and trace down
- For each person, find their longest descendant line
- The family depth is 1 + the longest line among all children
- Keep track of the maximum generation count found

### Human-Readable Pseudocode
```
function maxDepth(node):
    if node is nil:
        return 0
    
    leftDepth = maxDepth(node.left)
    rightDepth = maxDepth(node.right)
    
    return max(leftDepth, rightDepth) + 1
```

### Execution Visualization

### Example Tree:
```
     3
    / \
   9  20
     /  \
    15   7
```

### Recursive Call Stack:
```
maxDepth(3)
├── maxDepth(9) → returns 1
├── maxDepth(20)
│   ├── maxDepth(15) → returns 1
│   ├── maxDepth(7) → returns 1
│   └── max(1,1) + 1 = 2
└── max(1,2) + 1 = 3
```

### Key Visualization Points:
- **Base case**: nil nodes return 0
- **Leaf nodes**: Return 1 (no children)
- **Internal nodes**: Return 1 + max of children
- **Root calculation**: Final depth from entire tree

### Memory Layout Visualization:
```
Call Stack (top to bottom):
maxDepth(3) ← waiting for children
maxDepth(20) ← waiting for children  
maxDepth(15) ← returns 1
```

### Time Complexity Breakdown:
- **Each node visited**: O(1) work per node
- **Total nodes**: N
- **Total time**: O(N)
- **Space**: O(H) - maximum recursion depth

### Alternative Approaches:

#### 1. Iterative BFS (O(N) time, O(W) space where W is max width)
```go
func maxDepthBFS(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    queue := []*TreeNode{root}
    depth := 0
    
    for len(queue) > 0 {
        levelSize := len(queue)
        depth++
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    
    return depth
}
```
- **Pros**: No recursion stack, works for very deep trees
- **Cons**: Uses more memory for wide trees

#### 2. Iterative DFS with Stack (O(N) time, O(H) space)
```go
func maxDepthDFS(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    stack := []struct {
        node  *TreeNode
        depth int
    }{{root, 1}}
    
    maxDepth := 0
    
    for len(stack) > 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        node := current.node
        depth := current.depth
        
        if node.Left == nil && node.Right == nil {
            maxDepth = max(maxDepth, depth)
        }
        
        if node.Right != nil {
            stack = append(stack, struct {
                node  *TreeNode
                depth int
            }{node.Right, depth + 1})
        }
        
        if node.Left != nil {
            stack = append(stack, struct {
                node  *TreeNode
                depth int
            }{node.Left, depth + 1})
        }
    }
    
    return maxDepth
}
```
- **Pros**: No recursion, same space complexity
- **Cons**: More complex implementation

#### 3. Morris Traversal (O(N) time, O(1) space)
```go
func maxDepthMorris(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    maxDepth := 0
    current := root
    depth := 0
    
    for current != nil {
        if current.Left == nil {
            current = current.Right
            depth++
        } else {
            predecessor := current.Left
            depthTemp := 1
            
            for predecessor.Right != nil && predecessor.Right != current {
                predecessor = predecessor.Right
                depthTemp++
            }
            
            if predecessor.Right == nil {
                predecessor.Right = current
                current = current.Left
                depth++
            } else {
                predecessor.Right = nil
                current = current.Right
                maxDepth = max(maxDepth, depth)
                depth = depthTemp
            }
        }
    }
    
    return maxDepth
}
```
- **Pros**: O(1) space, no recursion or stack
- **Cons**: Very complex, modifies tree structure

### Extensions for Interviews:
- **Minimum Depth**: Find shortest path to leaf
- **Balanced Tree**: Check if height-balanced
- **N-ary Tree**: Extend to multiple children
- **Maximum Width**: Find widest level of tree
*/
func main() {
	// Test cases
	testCases := [][]interface{}{
		{3, 9, 20, nil, nil, 15, 7},
		{1, nil, 2},
		{},
		{1},
		{1, 2, 3, 4, 5, nil, nil, 6, 7, nil, nil, nil, nil, 8},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, nil, nil, nil, 5},
		{1, nil, 2, nil, 3, nil, 4, nil, 5},
		{0},
		{1, 2},
	}
	
	for i, nums := range testCases {
		root := createTree(nums)
		result1 := maxDepth(root)
		result2 := maxDepthBFS(root)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Recursive: %d, BFS: %d\n\n", result1, result2)
	}
}
