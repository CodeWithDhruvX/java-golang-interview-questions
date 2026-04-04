package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 543. Diameter of Binary Tree
// Time: O(N), Space: O(H) where H is tree height (recursion stack)
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		
		leftDepth := dfs(node.Left)
		rightDepth := dfs(node.Right)
		
		// Update max diameter found so far
		currentDiameter := leftDepth + rightDepth
		if currentDiameter > maxDiameter {
			maxDiameter = currentDiameter
		}
		
		// Return the maximum depth from this node
		return max(leftDepth, rightDepth) + 1
	}
	
	dfs(root)
	return maxDiameter
}

// Alternative approach returning both depth and diameter
func diameterOfBinaryTreeAlternative(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		
		leftDepth, leftDiameter := dfs(node.Left)
		rightDepth, rightDiameter := dfs(node.Right)
		
		// Current node diameter
		currentDiameter := leftDepth + rightDepth
		
		// Maximum diameter in subtrees
		maxDiameter := max(max(leftDiameter, rightDiameter), currentDiameter)
		
		// Return depth and max diameter
		return max(leftDepth, rightDepth) + 1, maxDiameter
	}
	
	_, maxDiameter := dfs(root)
	return maxDiameter
}

// Helper function to create a binary tree from slice
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

## 1. ALGORITHM PATTERN: Tree DFS with Global Maximum
- **Recursive Traversal**: Visit each node exactly once
- **Depth Calculation**: Compute depth of subtrees recursively
- **Diameter Update**: Update global max using left + right depths
- **Depth Return**: Return max depth + 1 for parent calculations

## 2. PROBLEM CHARACTERISTICS
- **Binary Tree**: Each node has at most 2 children
- **Diameter Definition**: Longest path between any two nodes
- **Path Through Root**: May or may not pass through root
- **Edge Count**: Diameter measured in edges (not nodes)

## 3. SIMILAR PROBLEMS
- Maximum Depth of Binary Tree (LeetCode 104)
- Longest Path in Binary Tree (LeetCode 124)
- Binary Tree Maximum Path Sum (LeetCode 124)
- Tree Height and Diameter calculations

## 4. KEY OBSERVATIONS
- **Global maximum**: Need to track max across all nodes
- **Local calculation**: Diameter through current node = leftDepth + rightDepth
- **Depth propagation**: Return depth for parent calculations
- **Single pass**: Can compute both depth and diameter in one traversal

## 5. VARIATIONS & EXTENSIONS
- **Node Count**: Count nodes instead of edges in diameter
- **Path Retrieval**: Return actual longest path
- **K-ary Tree**: Extend to nodes with multiple children
- **Multiple Queries**: Process multiple diameter queries

## 6. INTERVIEW INSIGHTS
- Always clarify: "Is diameter measured in nodes or edges?"
- Edge cases: empty tree, single node, unbalanced tree
- Space complexity: O(H) where H is tree height (recursion stack)
- Time complexity: O(N) - visit each node once

## 7. COMMON MISTAKES
- Not using global variable for maximum tracking
- Confusing diameter with height/depth
- Counting nodes instead of edges (or vice versa)
- Forgetting to add 1 when returning depth
- Not handling empty tree case properly

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(H) space
- **Two-pass approach**: Can compute depths first, then diameter
- **Iterative with stack**: Can avoid recursion but more complex
- **Morris traversal**: O(1) space but very complex

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the longest distance between any two people in a company:**
- You have an organizational chart (tree structure)
- You want to find the longest chain of command between any two employees
- For each manager, the longest path through them is: 
  longest path down left side + longest path down right side
- Keep track of the maximum such path found across all managers
- The company diameter is the maximum path length found

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary tree root node
2. **Goal**: Find longest path between any two nodes (in edges)
3. **Output**: Integer representing diameter length
4. **Constraint**: Path may not pass through root

#### Phase 2: Key Insight Recognition
- **"Local to global"** → Local diameter = leftDepth + rightDepth
- **"Depth propagation"** → Return max depth for parent calculations
- **"Global tracking"** → Need variable to track maximum across all nodes
- **"Single traversal"** → Can compute both depth and diameter together

#### Phase 3: Strategy Development
```
Human thought process:
"For each node, I need to know:
1. The longest path that goes through this node (left depth + right depth)
2. The maximum depth from this node to its deepest leaf
I'll compute both recursively and keep track of the global maximum.
The diameter is the maximum path found at any node."
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return 0 (no edges)
- **Single node**: Return 0 (no edges)
- **Unbalanced tree**: Handle naturally with depth calculations
- **Very deep tree**: Consider stack overflow limits

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:        1
            / \
           2   3
          / \
         4   5

Human thinking:
"I'll compute depth and diameter from bottom up:

Node 4: depth=1, diameter=0
Node 5: depth=1, diameter=0

Node 2: left depth=1, right depth=1
       local diameter = 1 + 1 = 2
       max depth = max(1,1) + 1 = 2
       global max = 2

Node 3: depth=1, diameter=0
       global max stays 2

Node 1: left depth=2, right depth=1
       local diameter = 2 + 1 = 3
       max depth = max(2,1) + 1 = 3
       global max = 3

Final answer: 3 (path 4-2-1-3 has 3 edges)"
```

#### Phase 6: Intuition Validation
- **Why global variable needed**: Diameter may be anywhere, not just at root
- **Why depth calculation**: Needed for parent diameter calculations
- **Why O(N) time**: Each node visited exactly once
- **Why O(H) space**: Recursion depth equals tree height

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just find longest root-to-leaf path?"** → Diameter may not pass through root
2. **"Should I count nodes or edges?"** → Clarify in interview, this solution counts edges
3. **"Can I compute depths first?"** → Yes, but this single-pass is more efficient
4. **"What about very deep trees?"** → Consider iterative approach to avoid stack overflow

### Real-World Analogy
**Like finding the longest communication chain in a network:**
- You have a computer network (tree structure)
- You want to find the longest path between any two computers
- For each router, the longest path through it is:
  longest path to one side + longest path to other side
- Keep track of the maximum such path across all routers
- The network diameter is the maximum communication path found

### Human-Readable Pseudocode
```
globalMaxDiameter = 0

function diameterOfBinaryTree(node):
    if node is nil:
        return 0
    
    leftDepth = diameterOfBinaryTree(node.left)
    rightDepth = diameterOfBinaryTree(node.right)
    
    // Update global maximum diameter
    currentDiameter = leftDepth + rightDepth
    globalMaxDiameter = max(globalMaxDiameter, currentDiameter)
    
    // Return maximum depth from this node
    return max(leftDepth, rightDepth) + 1
```

### Execution Visualization

### Example Tree:
```
        1
       / \
      2   3
     / \
    4   5
```

### Recursive Call Flow:
```
diameter(1)
├── diameter(2)
│   ├── diameter(4) → depth=1, diameter=0
│   ├── diameter(5) → depth=1, diameter=0
│   ├── local diameter = 1+1=2, global max=2
│   └── return depth=2
├── diameter(3) → depth=1, diameter=0
├── local diameter = 2+1=3, global max=3
└── return depth=3

Final answer: 3
```

### Key Visualization Points:
- **Local diameter**: leftDepth + rightDepth at each node
- **Global tracking**: Maximum across all nodes
- **Depth return**: max(left, right) + 1 for parent
- **Bottom-up**: Calculations flow from leaves to root

### Memory Layout Visualization:
```
Call Stack (top to bottom):
diameter(1) ← waiting for children
diameter(2) ← waiting for children
diameter(4) ← returns depth=1
```

### Time Complexity Breakdown:
- **Each node visited**: O(1) work per node
- **Total nodes**: N
- **Total time**: O(N)
- **Space**: O(H) - maximum recursion depth

### Alternative Approaches:

#### 1. Two-Pass Approach (O(N) time, O(N) space)
```go
func diameterOfBinaryTreeTwoPass(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    // First pass: compute depths for all nodes
    depths := make(map[*TreeNode]int)
    var computeDepth func(*TreeNode) int
    computeDepth = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        
        leftDepth := computeDepth(node.Left)
        rightDepth := computeDepth(node.Right)
        
        depth := max(leftDepth, rightDepth) + 1
        depths[node] = depth
        return depth
    }
    
    computeDepth(root)
    
    // Second pass: compute diameters using precomputed depths
    maxDiameter := 0
    var computeDiameter func(*TreeNode)
    computeDiameter = func(node *TreeNode) {
        if node == nil {
            return
        }
        
        leftDepth := 0
        if node.Left != nil {
            leftDepth = depths[node.Left]
        }
        
        rightDepth := 0
        if node.Right != nil {
            rightDepth = depths[node.Right]
        }
        
        currentDiameter := leftDepth + rightDepth
        maxDiameter = max(maxDiameter, currentDiameter)
        
        computeDiameter(node.Left)
        computeDiameter(node.Right)
    }
    
    computeDiameter(root)
    return maxDiameter
}
```
- **Pros**: Clear separation of concerns
- **Cons**: O(N) extra space for depth map

#### 2. Iterative with Stack (O(N) time, O(H) space)
```go
func diameterOfBinaryTreeIterative(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    maxDiameter := 0
    stack := []struct {
        node     *TreeNode
        visited  bool
        depth    int
    }{{root, false, 0}}
    
    depths := make(map[*TreeNode]int)
    
    for len(stack) > 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        node := current.node
        
        if !current.visited {
            // Push back as visited
            stack = append(stack, struct {
                node     *TreeNode
                visited  bool
                depth    int
            }{node, true, 0})
            
            // Push children
            if node.Right != nil {
                stack = append(stack, struct {
                    node     *TreeNode
                    visited  bool
                    depth    int
                }{node.Right, false, 0})
            }
            if node.Left != nil {
                stack = append(stack, struct {
                    node     *TreeNode
                    visited  bool
                    depth    int
                }{node.Left, false, 0})
            }
        } else {
            // Calculate depth and diameter
            leftDepth := 0
            rightDepth := 0
            
            if node.Left != nil {
                leftDepth = depths[node.Left]
            }
            if node.Right != nil {
                rightDepth = depths[node.Right]
            }
            
            currentDiameter := leftDepth + rightDepth
            maxDiameter = max(maxDiameter, currentDiameter)
            
            depths[node] = max(leftDepth, rightDepth) + 1
        }
    }
    
    return maxDiameter
}
```
- **Pros**: No recursion stack
- **Cons**: More complex implementation

#### 3. Return Both Values (O(N) time, O(H) space)
```go
func diameterOfBinaryTreeAlternative(root *TreeNode) int {
    var dfs func(*TreeNode) (int, int)
    dfs = func(node *TreeNode) (int, int) {
        if node == nil {
            return 0, 0
        }
        
        leftDepth, leftDiameter := dfs(node.Left)
        rightDepth, rightDiameter := dfs(node.Right)
        
        // Current node diameter
        currentDiameter := leftDepth + rightDepth
        
        // Maximum diameter in subtrees
        maxDiameter := max(max(leftDiameter, rightDiameter), currentDiameter)
        
        // Return depth and max diameter
        return max(leftDepth, rightDepth) + 1, maxDiameter
    }
    
    _, maxDiameter := dfs(root)
    return maxDiameter
}
```
- **Pros**: No global variable, clean interface
- **Cons**: Slightly more complex return handling

### Extensions for Interviews:
- **Node Count Diameter**: Count nodes instead of edges
- **Path Retrieval**: Return actual longest path nodes
- **K-ary Tree**: Extend to multiple children
- **Multiple Queries**: Process multiple diameter queries
- **Binary Tree Maximum Path Sum**: Similar pattern with value sums
*/
func main() {
	// Test cases
	testCases := [][]interface{}{
		{1, 2, 3, 4, 5},
		{1, 2},
		{1, 2, nil, nil, 3, 4, nil, nil, nil, 5},
		{4, -7, -3, nil, nil, -9, -3, 9, -7, -4, nil, 6, nil, -6, -6, nil, 0, 6, 5, nil, 9, nil, nil, -1, -4, nil, nil, nil, -2},
		{1},
		{},
		{1, nil, 2, nil, 3, nil, 4, nil, 5},
		{1, 2, 3, 4, nil, nil, nil, 5},
		{1, 2, 3, nil, nil, nil, 4, nil, nil, 5},
	}
	
	for i, nums := range testCases {
		root := createTree(nums)
		result1 := diameterOfBinaryTree(root)
		result2 := diameterOfBinaryTreeAlternative(root)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Method 1: %d, Method 2: %d\n\n", result1, result2)
	}
}
