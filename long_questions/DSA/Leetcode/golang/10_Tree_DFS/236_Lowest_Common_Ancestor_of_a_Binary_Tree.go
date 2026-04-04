package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 236. Lowest Common Ancestor of a Binary Tree
// Time: O(N), Space: O(H) where H is tree height (recursion stack)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	
	if left != nil && right != nil {
		return root // p and q are in different subtrees
	}
	
	if left != nil {
		return left
	}
	
	return right
}

// Iterative approach using parent pointers
func lowestCommonAncestorIterative(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	
	// Build parent map and depth
	parent := make(map[*TreeNode]*TreeNode)
	depth := make(map[*TreeNode]int)
	
	stack := []*TreeNode{root}
	parent[root] = nil
	depth[root] = 0
	
	// Build parent pointers and depths
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		
		if node.Left != nil {
			parent[node.Left] = node
			depth[node.Left] = depth[node] + 1
			stack = append(stack, node.Left)
		}
		
		if node.Right != nil {
			parent[node.Right] = node
			depth[node.Right] = depth[node] + 1
			stack = append(stack, node.Right)
		}
	}
	
	// Bring both nodes to the same depth
	for depth[p] > depth[q] {
		p = parent[p]
	}
	for depth[q] > depth[p] {
		q = parent[q]
	}
	
	// Move up both nodes until they meet
	for p != q {
		p = parent[p]
		q = parent[q]
	}
	
	return p
}

// Helper function to create a binary tree from slice
func createTree(nums []interface{}) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	
	nodes := make([]*TreeNode, len(nums))
	nodeMap := make(map[int]*TreeNode)
	
	for i, val := range nums {
		if val != nil {
			nodes[i] = &TreeNode{Val: val.(int)}
			nodeMap[val.(int)] = nodes[i]
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

## 1. ALGORITHM PATTERN: Tree DFS with Bottom-Up Propagation
- **Recursive Search**: Find nodes p and q in subtrees
- **Base Cases**: Return node if it's p, q, or nil
- **Bottom-Up Propagation**: Return found nodes up the tree
- **LCA Identification**: Current node is LCA if both children return non-nil

## 2. PROBLEM CHARACTERISTICS
- **Binary Tree**: Each node has at most 2 children
- **Ancestor Definition**: Node that is ancestor of both p and q
- **Lowest**: Deepest common ancestor (closest to both nodes)
- **Node References**: Given actual node references, not values

## 3. SIMILAR PROBLEMS
- LCA of BST (LeetCode 235) - Simplified for BST structure
- LCA of Binary Tree III (LeetCode 1644) - With parent pointers
- LCA of k-ary Tree (LeetCode 1676) - Multiple children
- Distance Between Nodes in BST (LeetCode 1740)

## 4. KEY OBSERVATIONS
- **Bottom-up approach**: Results propagate from leaves to root
- **Single pass**: Find both nodes in one traversal
- **LCA condition**: Node where p and q are in different subtrees
- **Early termination**: Can stop when LCA found

## 5. VARIATIONS & EXTENSIONS
- **BST LCA**: Use BST properties for optimization
- **With Parent Pointers**: Can use parent map approach
- **K-ary Tree**: Extend to multiple children
- **Distance Calculation**: Find distance between any two nodes

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are p and q guaranteed to exist? Can they be the same?"
- Edge cases: p is ancestor of q, q is ancestor of p, same node
- Space complexity: O(H) where H is tree height (recursion stack)
- Time complexity: O(N) - visit each node once

## 7. COMMON MISTAKES
- Not handling case where p or q is ancestor of the other
- Using node comparison instead of reference comparison
- Not handling nil nodes properly in base cases
- Forgetting to check both children for LCA condition
- Using value comparison when references are given

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(H) space
- **BST optimization**: Use BST properties for O(log N) in BST
- **Parent pointer approach**: Can use iterative method
- **Early pruning**: Not applicable without additional constraints

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the meeting point of two people in a family tree:**
- You have a family tree showing parent-child relationships
- Two family members need to find their closest common ancestor
- Start from each person and trace up their family lines
- The first person who appears in both family lines is the answer
- This person is the "lowest" (most recent) common ancestor
- If one person is ancestor of the other, that ancestor is the answer

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary tree root and references to nodes p and q
2. **Goal**: Find lowest common ancestor of p and q
3. **Output**: Reference to LCA node
4. **Constraint**: Both p and q are guaranteed to exist in tree

#### Phase 2: Key Insight Recognition
- **"Bottom-up propagation"** → Results flow from leaves to root
- **"Three-way logic"** → Each node returns itself, p, q, or LCA
- **"LCA condition"** → Node where p and q split to different subtrees
- **"Single traversal"** → Find both nodes simultaneously

#### Phase 3: Strategy Development
```
Human thought process:
"I'll search for both nodes p and q at the same time.
For each node, I'll check:
1. If this node is p or q, return it
2. Search left subtree for p or q
3. Search right subtree for p or q
4. If both children return non-nil, this node is the LCA
5. Otherwise, return whichever child found something"
```

#### Phase 4: Edge Case Handling
- **p is ancestor of q**: Return p when found
- **q is ancestor of p**: Return q when found
- **p and q are same**: Return that node
- **Root is LCA**: Handle naturally by algorithm

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:        3
            / \
           5   1
          / \ / \
         6  2 0  8
           / \
          7   4

p = 5, q = 1

Human thinking:
"I'll search from bottom up:

Node 6: not p/q, no children → return nil
Node 7: not p/q, no children → return nil  
Node 4: not p/q, no children → return nil
Node 0: not p/q, no children → return nil
Node 8: not p/q, no children → return nil

Node 2: not p/q, left=7(nil), right=4(nil) → return nil
Node 5: this is p! → return 5

Node 1: this is q! → return 1

Node 3: not p/q, left=5, right=1 → both non-nil!
        This is where p and q split → return 3 (LCA)

Final answer: 3"
```

#### Phase 6: Intuition Validation
- **Why bottom-up works**: LCA is where paths from p and q converge
- **Why single traversal**: Can find both nodes and their split point
- **Why O(N) time**: Each node visited exactly once
- **Why O(H) space**: Recursion depth equals tree height

### Common Human Pitfalls & How to Avoid Them
1. **"Why not find paths separately?"** → That would be O(N²), this is O(N)
2. **"Should I use values instead of references?"** → Problem gives references, use them
3. **"What if one node doesn't exist?"** → Problem guarantees both exist
4. **"Can I optimize for BST?"** → Yes, but this is general binary tree

### Real-World Analogy
**Like finding the first common manager of two employees:**
- You have an organizational chart (tree structure)
- Two employees need to find their closest common manager
- Start from each employee and trace up their management chain
- The first manager who manages both employees is the answer
- This manager is the "lowest" (closest) common manager
- If one employee manages the other, that manager is the answer

### Human-Readable Pseudocode
```
function lowestCommonAncestor(node, p, q):
    if node is nil or node is p or node is q:
        return node
    
    left = lowestCommonAncestor(node.left, p, q)
    right = lowestCommonAncestor(node.right, p, q)
    
    if left != nil and right != nil:
        return node  // p and q are in different subtrees
    
    if left != nil:
        return left
    else:
        return right
```

### Execution Visualization

### Example Tree:
```
        3
       / \
      5   1
     / \ / \
    6  2 0  8
      / \
     7   4

p = 5, q = 1
```

### Recursive Call Flow:
```
LCA(3,5,1)
├── LCA(5,5,1) → returns 5 (found p)
├── LCA(1,5,1) → returns 1 (found q)
└── both children non-nil → return 3 (LCA)
```

### Example where p is ancestor of q:
```
p = 5, q = 4

LCA(3,5,4)
├── LCA(5,5,4)
│   ├── LCA(6,5,4) → nil
│   └── LCA(2,5,4)
│       ├── LCA(7,5,4) → nil
│       └── LCA(4,5,4) → returns 4 (found q)
│   └── left=nil, right=4 → return 4
├── left=5, right=4 → return 5 (p is ancestor)
└── Final answer: 5
```

### Key Visualization Points:
- **Base cases**: Return node if it's p, q, or nil
- **Propagation**: Results bubble up from leaves
- **LCA detection**: Both children return non-nil
- **Ancestor handling**: One child non-nil returns that child

### Memory Layout Visualization:
```
Call Stack (top to bottom):
LCA(3,5,1)
LCA(5,5,1) ← returns 5
LCA(1,5,1) ← returns 1
LCA(3,5,1) ← both children found, returns 3
```

### Time Complexity Breakdown:
- **Each node visited**: O(1) work per node
- **Total nodes**: N
- **Total time**: O(N) - stops when LCA found
- **Space**: O(H) - maximum recursion depth

### Alternative Approaches:

#### 1. Parent Pointer Approach (O(N) time, O(N) space)
```go
func lowestCommonAncestorIterative(root, p, q *TreeNode) *TreeNode {
    // Build parent map and depths
    parent := make(map[*TreeNode]*TreeNode)
    depth := make(map[*TreeNode]int)
    
    stack := []*TreeNode{root}
    parent[root] = nil
    depth[root] = 0
    
    // Build parent pointers and depths
    for len(stack) > 0 {
        node := stack[0]
        stack = stack[1:]
        
        if node.Left != nil {
            parent[node.Left] = node
            depth[node.Left] = depth[node] + 1
            stack = append(stack, node.Left)
        }
        
        if node.Right != nil {
            parent[node.Right] = node
            depth[node.Right] = depth[node] + 1
            stack = append(stack, node.Right)
        }
    }
    
    // Bring both nodes to same depth
    for depth[p] > depth[q] {
        p = parent[p]
    }
    for depth[q] > depth[p] {
        q = parent[q]
    }
    
    // Move up both nodes until they meet
    for p != q {
        p = parent[p]
        q = parent[q]
    }
    
    return p
}
```
- **Pros**: Iterative, no recursion stack
- **Cons**: Uses O(N) extra space for parent map

#### 2. Path Storage Approach (O(N) time, O(H) space)
```go
func lowestCommonAncestorPath(root, p, q *TreeNode) *TreeNode {
    var pathToP, pathToQ []*TreeNode
    
    var findPath func(*TreeNode, *TreeNode, *[]*TreeNode) bool
    findPath = func(node, target *TreeNode, path *[]*TreeNode) bool {
        if node == nil {
            return false
        }
        
        *path = append(*path, node)
        
        if node == target {
            return true
        }
        
        if findPath(node.Left, target, path) || findPath(node.Right, target, path) {
            return true
        }
        
        *path = (*path)[:len(*path)-1] // Backtrack
        return false
    }
    
    findPath(root, p, &pathToP)
    findPath(root, q, &pathToQ)
    
    // Find common ancestor
    i := 0
    for i < len(pathToP) && i < len(pathToQ) && pathToP[i] == pathToQ[i] {
        i++
    }
    
    return pathToP[i-1]
}
```
- **Pros**: Conceptually simple
- **Cons**: Uses extra space for path storage

#### 3. BST Optimization (O(log N) time, O(H) space for BST)
```go
func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    // If both nodes are in left subtree
    if p.Val < root.Val && q.Val < root.Val {
        return lowestCommonAncestorBST(root.Left, p, q)
    }
    
    // If both nodes are in right subtree
    if p.Val > root.Val && q.Val > root.Val {
        return lowestCommonAncestorBST(root.Right, p, q)
    }
    
    // This is the split point
    return root
}
```
- **Pros**: O(log N) time for BST
- **Cons**: Only works for BST, not general binary tree

### Extensions for Interviews:
- **Distance Between Nodes**: Find distance between any two nodes
- **K-ary Tree LCA**: Extend to nodes with multiple children
- **Multiple Queries**: Process multiple LCA queries efficiently
- **LCA with Parent Pointers**: Use parent references for iterative solution
*/
func main() {
	// Test cases
	testCases := []struct {
		tree []interface{}
		p    int
		q    int
	}{
		{[]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, 5, 1},
		{[]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, 5, 4},
		{[]interface{}{1, 2}, 1, 2},
		{[]interface{}{1, 2, nil, nil, 3, 4}, 1, 4},
		{[]interface{}{1, 2, 3, 4, 5}, 2, 4},
		{[]interface{}{1, 2, 3, 4, 5}, 1, 3},
		{[]interface{}{1}, 1, 1},
		{[]interface{}{1, nil, 2}, 1, 2},
		{[]interface{}{37, -34, -48, nil, -100, -100, 48, nil, nil, nil, -54, nil, -71, -22, nil, nil}, -34, -48},
	}
	
	for i, tc := range testCases {
		root := createTree(tc.tree)
		
		var p, q *TreeNode
		if tc.tree != nil {
			// Find nodes p and q
			for _, val := range tc.tree {
				if val != nil {
					if val.(int) == tc.p {
						p = &TreeNode{Val: tc.p}
					}
					if val.(int) == tc.q {
						q = &TreeNode{Val: tc.q}
					}
				}
			}
		}
		
		result1 := lowestCommonAncestor(root, p, q)
		result2 := lowestCommonAncestorIterative(root, p, q)
		
		fmt.Printf("Test Case %d: tree=%v, p=%d, q=%d\n", i+1, tc.tree, tc.p, tc.q)
		if result1 != nil {
			fmt.Printf("  Recursive: %d, Iterative: %d\n", result1.Val, result2.Val)
		} else {
			fmt.Printf("  Recursive: nil, Iterative: nil\n")
		}
		fmt.Println()
	}
}
