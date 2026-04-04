package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 199. Binary Tree Right Side View
// Time: O(N), Space: O(N)
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	
	var result []int
	queue := []*TreeNode{root}
	
	for len(queue) > 0 {
		levelSize := len(queue)
		
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			
			// Add children to queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			
			// The last node at each level is visible from the right side
			if i == levelSize-1 {
				result = append(result, node.Val)
			}
		}
	}
	
	return result
}

// Recursive approach with depth tracking
func rightSideViewRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	
	var result []int
	
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		
		// If this is the first node we've seen at this depth
		if len(result) == depth {
			result = append(result, node.Val)
		}
		
		// Traverse right subtree first, then left
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	
	dfs(root, 0)
	return result
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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: BFS Level-based Selection
- **Queue-based Traversal**: Process nodes level by level
- **Level Boundary**: Use queue length to identify level boundaries
- **Rightmost Selection**: Capture last node of each level
- **Complete Traversal**: Visit every node exactly once

## 2. PROBLEM CHARACTERISTICS
- **Binary Tree**: Each node has at most 2 children
- **Right Side View**: Only visible nodes from right perspective
- **Level Processing**: Process one level at a time
- **1D Output**: Return array of visible nodes

## 3. SIMILAR PROBLEMS
- Binary Tree Level Order Traversal (LeetCode 102)
- Binary Tree Left Side View (LeetCode - variant)
- Binary Tree Zigzag Level Order Traversal (LeetCode 103)
- Average of Levels in Binary Tree (LeetCode 637)

## 4. KEY OBSERVATIONS
- **Rightmost visibility**: Last node in each level is visible from right
- **Level boundary**: Queue size at start = nodes in current level
- **Index tracking**: Last node in level = index levelSize-1
- **Sequential processing**: Process all nodes at current level together

## 5. VARIATIONS & EXTENSIONS
- **Left Side View**: Only first node of each level
- **Top View**: Nodes visible from top perspective
- **Bottom View**: Nodes visible from bottom perspective
- **N-ary Tree**: Extend to nodes with multiple children

## 6. INTERVIEW INSIGHTS
- Always clarify: "What should be returned for empty tree?"
- Edge cases: empty tree, single node, unbalanced tree
- Space complexity: O(W) where W is maximum tree width
- Time complexity: O(N) - visit each node once

## 7. COMMON MISTAKES
- Not tracking level size properly, missing rightmost node
- Using wrong index for rightmost node (should be levelSize-1)
- Not handling empty tree case
- Forgetting to add children to queue
- Confusing right side view with rightmost node in tree

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(W) space
- **Recursive alternative**: DFS with depth-first right traversal
- **Early termination**: Not applicable (need to visit all levels)
- **Memory optimization**: Use single queue efficiently

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like looking at a building from the right side:**
- You have a building with floors organized in levels (tree structure)
- From the right side, you can only see the rightmost room on each floor
- For each floor, process all rooms from left to right
- The last room you encounter on each floor is visible from the right
- Note which rooms are visible from your right-side perspective
- Continue until you reach all floors at the bottom

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary tree root node
2. **Goal**: Return nodes visible from right side
3. **Output**: Array of visible nodes from top to bottom
4. **Constraint**: Only rightmost node of each level is visible

#### Phase 2: Key Insight Recognition
- **"Rightmost visibility"** → Last node in each level is visible
- **"Level processing"** → Need to process one level at a time
- **"Queue natural fit"** → BFS naturally processes by distance from root
- **"Index tracking"** → Can track position within level

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find what's visible from the right side of the tree.
For each level, only the rightmost node is visible.
I'll use BFS to process level by level.
For each level, I'll process all nodes and capture the last one.
The last node in each level becomes part of the right side view."
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return empty array
- **Single node**: Return array with just the root
- **Unbalanced tree**: Handle naturally with queue processing
- **Right-leaning tree**: All nodes might be visible

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:        1
            / \
           2   3
          / \   \
         4   5   6

Human thinking:
"I'll process level by level to find rightmost nodes:

Initial queue: [1]
Level 0: process 1, add 2 and 3 to queue
Rightmost node (index 0): 1
Queue becomes: [2, 3]
Result: [1]

Level 1: queue size = 2, process 2 and 3
Process 2: add 4 and 5 to queue
Process 3: add 6 to queue
Rightmost node (index 1): 3
Queue becomes: [4, 5, 6]
Result: [1, 3]

Level 2: queue size = 3, process 4, 5, and 6
Process 4: no children
Process 5: no children
Process 6: no children
Rightmost node (index 2): 6
Queue becomes: []
Result: [1, 3, 6]

Done! Final result: [1, 3, 6]"
```

#### Phase 6: Intuition Validation
- **Why BFS works**: Naturally processes by distance from root
- **Why rightmost matters**: From right side, only rightmost node visible
- **Why O(N) time**: Each node visited exactly once
- **Why O(W) space**: Queue holds at most one level's worth of nodes

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just traverse right subtree?"** → Left subtree might have rightmost nodes
2. **"Should I use DFS instead?"** → Can, but BFS is more intuitive for level-based
3. **"What about very wide trees?"** → Space complexity is O(W) where W is max width
4. **"Can I optimize further?"** → Already optimal for this problem

### Real-World Analogy
**Like taking a photo of a building from the right side:**
- You have a multi-story building with rooms on each floor (tree structure)
- From the right side, you can only see the rightmost room on each floor
- For each floor, walk through all rooms from left to right
- The last room you see on each floor appears in your photo
- Take photos of each floor from right to left
- The final photo shows the right side view of the building

### Human-Readable Pseudocode
```
function rightSideView(root):
    if root is nil:
        return []
    
    result = []
    queue = [root]
    
    while queue is not empty:
        levelSize = length(queue)
        
        for i from 0 to levelSize-1:
            node = queue.dequeue()
            
            if i == levelSize-1:  // Last node in level
                result.append(node.value)
            
            if node.left is not nil:
                queue.enqueue(node.left)
            if node.right is not nil:
                queue.enqueue(node.right)
    
    return result
```

### Execution Visualization

### Example Tree:
```
        1
       / \
      2   3
     / \   \
    4   5   6
```

### BFS Right Side View Process:
```
Initial: queue=[1], result=[]

Level 0: queue=[1], size=1
- Process 1: i=0, i==0 (last), add 1 to result, add 2,3 to queue
- queue=[2, 3], result=[1]

Level 1: queue=[2, 3], size=2
- Process 2: i=0, not last, add 4,5 to queue
- Process 3: i=1, i==1 (last), add 3 to result, add 6 to queue
- queue=[4, 5, 6], result=[1, 3]

Level 2: queue=[4, 5, 6], size=3
- Process 4: i=0, not last, no children
- Process 5: i=1, not last, no children
- Process 6: i=2, i==2 (last), add 6 to result, no children
- queue=[], result=[1, 3, 6]

Done! Final result: [1, 3, 6]
```

### Key Visualization Points:
- **Level boundary**: queue size at start of each iteration
- **Rightmost detection**: i == levelSize-1 identifies last node
- **Child enqueuing**: Children added for next level processing
- **Result building**: Only rightmost nodes added to result

### Memory Layout Visualization:
```
Queue: [1] → [2, 3] → [4, 5, 6] → []
Index: 0     0,1       0,1,2
Rightmost: 1     3         6
Result: [1] → [1, 3] → [1, 3, 6]
```

### Time Complexity Breakdown:
- **Each node visited**: O(1) work per node
- **Total nodes**: N
- **Total time**: O(N)
- **Space**: O(W) - maximum queue size (tree width)

### Alternative Approaches:

#### 1. Recursive DFS with Right-First Traversal (O(N) time, O(H) space)
```go
func rightSideViewRecursive(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    var result []int
    
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, depth int) {
        if node == nil {
            return
        }
        
        // If this is the first node we've seen at this depth
        if len(result) == depth {
            result = append(result, node.Val)
        }
        
        // Traverse right subtree first, then left
        dfs(node.Right, depth+1)
        dfs(node.Left, depth+1)
    }
    
    dfs(root, 0)
    return result
}
```
- **Pros**: No queue, uses recursion stack, natural right-first approach
- **Cons**: O(H) space, may cause stack overflow

#### 2. Two Queue Approach (O(N) time, O(W) space)
```go
func rightSideViewTwoQueues(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    var result []int
    currentLevel := []*TreeNode{root}
    nextLevel := []*TreeNode{}
    
    for len(currentLevel) > 0 {
        // Add rightmost node of current level to result
        rightmost := currentLevel[len(currentLevel)-1]
        result = append(result, rightmost.Val)
        
        for _, node := range currentLevel {
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }
            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
        }
        
        currentLevel, nextLevel = nextLevel, []*TreeNode{}
    }
    
    return result
}
```
- **Pros**: Clear separation of levels, simpler rightmost detection
- **Cons**: More memory for two queues

#### 3. Sentinel Approach (O(N) time, O(W) space)
```go
func rightSideViewSentinel(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    var result []int
    queue := []*TreeNode{root, nil} // nil marks level end
    lastNode := root
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if node == nil {
            // Level finished, add last node to result
            result = append(result, lastNode.Val)
            
            if len(queue) > 0 {
                queue = append(queue, nil) // Add sentinel for next level
            }
        } else {
            lastNode = node // Track last node seen
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    
    return result
}
```
- **Pros**: No need to track queue size or index
- **Cons**: Extra sentinel nodes, need to track last node

### Extensions for Interviews:
- **Left Side View**: Only first node of each level
- **Top View**: Nodes visible from top perspective
- **Bottom View**: Nodes visible from bottom perspective
- **N-ary Tree**: Extend to nodes with multiple children
- **Multiple Views**: Generate multiple views in single traversal
*/
func main() {
	// Test cases
	testCases := [][]interface{}{
		{1, 2, 3, nil, 5, nil, 4},
		{1, 2, 3, 4, 5, nil, nil, nil, nil, nil, nil, nil, 6, 7},
		{1, nil, 3},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, nil, 3, nil, 4, nil, 5},
		{1, 2, 3, nil, nil, nil, 4, nil, nil, 5},
		{1},
		{},
		{1, 2, 3, 4, nil, nil, nil, 5},
	}
	
	for i, nums := range testCases {
		root := createTree(nums)
		result1 := rightSideView(root)
		result2 := rightSideViewRecursive(root)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Iterative: %v\n", result1)
		fmt.Printf("  Recursive: %v\n\n", result2)
	}
}
