package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 102. Binary Tree Level Order Traversal
// Time: O(N), Space: O(N)
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	
	var result [][]int
	queue := []*TreeNode{root}
	
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]int, 0, levelSize)
		
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)
			
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		
		result = append(result, currentLevel)
	}
	
	return result
}

// Recursive approach
func levelOrderRecursive(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	
	var result [][]int
	
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		
		// Ensure result has enough levels
		for len(result) <= level {
			result = append(result, []int{})
		}
		
		result[level] = append(result[level], node.Val)
		
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
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

## 1. ALGORITHM PATTERN: Breadth-First Search (BFS) Level Traversal
- **Queue-based Traversal**: Process nodes level by level
- **Level Size Tracking**: Use queue length to identify level boundaries
- **Batch Processing**: Process all nodes at current level before moving to next
- **Child Enqueuing**: Add children of current level nodes to queue

## 2. PROBLEM CHARACTERISTICS
- **Binary Tree**: Each node has at most 2 children
- **Level Order**: Visit nodes from top to bottom, left to right
- **2D Output**: Return array of arrays, one subarray per level
- **Complete Traversal**: Visit every node exactly once

## 3. SIMILAR PROBLEMS
- Binary Tree Zigzag Level Order Traversal (LeetCode 103)
- Binary Tree Right Side View (LeetCode 199)
- Average of Levels in Binary Tree (LeetCode 637)
- N-ary Tree Level Order Traversal (LeetCode 429)

## 4. KEY OBSERVATIONS
- **Queue management**: Queue holds nodes of current and next levels
- **Level boundary**: queue length at start = nodes in current level
- **Sequential processing**: Process all nodes at current level together
- **Memory efficiency**: Only one level in queue at any time

## 5. VARIATIONS & EXTENSIONS
- **Zigzag Traversal**: Alternate direction at each level
- **Right Side View**: Only last node of each level
- **Level Averages**: Calculate average of each level
- **N-ary Tree**: Extend to nodes with multiple children

## 6. INTERVIEW INSIGHTS
- Always clarify: "Should I return empty array for null tree?"
- Edge cases: empty tree, single node, unbalanced tree
- Space complexity: O(W) where W is maximum tree width
- Time complexity: O(N) - visit each node once

## 7. COMMON MISTAKES
- Not tracking level size properly, mixing levels
- Using wrong queue operations (FIFO vs LIFO)
- Not handling empty tree case
- Forgetting to add children to queue
- Using slice operations that are O(n) instead of O(1)

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(W) space
- **Pre-allocate slices**: Use make with capacity for efficiency
- **Two queues**: Can use separate queues for current/next levels
- **Recursive alternative**: DFS with level tracking

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like processing a company organization level by level:**
- You have an organizational chart (tree structure)
- You want to process employees by management level
- Start with the CEO (root), then all VPs, then all directors, etc.
- At each level, gather everyone before moving to the next level down
- For each person, note their direct reports for the next level
- Continue until you reach all employees at the bottom

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary tree root node
2. **Goal**: Return nodes grouped by level from top to bottom
3. **Output**: 2D array where each subarray represents one level
4. **Constraint**: Process left-to-right within each level

#### Phase 2: Key Insight Recognition
- **"Level by level"** → Need to process all nodes at same depth together
- **"Queue natural fit"** → BFS naturally processes by distance from root
- **"Level boundary"** → Queue size tells us nodes in current level
- **"Batch processing"** → Process entire level before moving to next

#### Phase 3: Strategy Development
```
Human thought process:
"I need to visit the tree level by level.
I'll use a queue to keep track of nodes to visit.
For each level, I'll process all nodes currently in the queue,
adding their children to the queue for the next level.
The queue size tells me how many nodes are in the current level."
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return empty array
- **Single node**: Return array with one subarray containing the root
- **Unbalanced tree**: Handle naturally with queue processing
- **Very wide tree**: Consider memory usage for large queues

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:     3
         / \
        9  20
           /  \
         15   7

Human thinking:
"I'll process level by level:

Initial queue: [3]
Level 0: process 3, add 9 and 20 to queue
Queue becomes: [9, 20]
Result: [[3]]

Level 1: queue size = 2, process 9 and 20
Process 9: no children
Process 20: add 15 and 7 to queue
Queue becomes: [15, 7]
Result: [[3], [9, 20]]

Level 2: queue size = 2, process 15 and 7
Process 15: no children
Process 7: no children
Queue becomes: []
Result: [[3], [9, 20], [15, 7]]

Done! Final result: [[3], [9, 20], [15, 7]]"
```

#### Phase 6: Intuition Validation
- **Why queue works**: FIFO naturally processes by distance from root
- **Why level size matters**: Distinguishes current level from next level
- **Why O(N) time**: Each node visited exactly once
- **Why O(W) space**: Queue holds at most one level's worth of nodes

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use recursion?"** → Can, but BFS is more natural for level order
2. **"Should I use two queues?"** → Possible, but single queue with size tracking is simpler
3. **"What about very wide trees?"** → Space complexity is O(W) where W is max width
4. **"Can I optimize further?"** → Already optimal for this problem

### Real-World Analogy
**Like processing customer service tickets by priority level:**
- You have a ticket system organized by priority levels (tree levels)
- Level 0: Critical tickets, Level 1: High priority, Level 2: Medium, etc.
- Process all critical tickets first, then all high priority, etc.
- For each ticket, note any follow-up tickets it creates (children)
- Move to next priority level only after current level is complete
- This ensures fair processing within each priority level

### Human-Readable Pseudocode
```
function levelOrder(root):
    if root is nil:
        return []
    
    result = []
    queue = [root]
    
    while queue is not empty:
        levelSize = length(queue)
        currentLevel = []
        
        for i from 0 to levelSize-1:
            node = queue.dequeue()
            currentLevel.append(node.value)
            
            if node.left is not nil:
                queue.enqueue(node.left)
            if node.right is not nil:
                queue.enqueue(node.right)
        
        result.append(currentLevel)
    
    return result
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

### BFS Process:
```
Initial: queue=[3], result=[]

Level 0: queue=[3], size=1
- Process 3: add 9, 20 to queue
- queue=[9, 20], result=[[3]]

Level 1: queue=[9, 20], size=2  
- Process 9: no children
- Process 20: add 15, 7 to queue
- queue=[15, 7], result=[[3], [9, 20]]

Level 2: queue=[15, 7], size=2
- Process 15: no children  
- Process 7: no children
- queue=[], result=[[3], [9, 20], [15, 7]]

Done!
```

### Key Visualization Points:
- **Level boundary**: queue size at start of each iteration
- **Batch processing**: All nodes at same level processed together
- **Child enqueuing**: Children added for next level processing
- **Sequential order**: Left-to-right within each level maintained

### Memory Layout Visualization:
```
Queue state evolution:
[3] → [9, 20] → [15, 7] → []
Level 0    Level 1     Level 2   Empty
```

### Time Complexity Breakdown:
- **Each node visited**: O(1) work per node
- **Total nodes**: N
- **Total time**: O(N)
- **Space**: O(W) - maximum queue size (tree width)

### Alternative Approaches:

#### 1. Recursive DFS with Level Tracking (O(N) time, O(H) space)
```go
func levelOrderRecursive(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var result [][]int
    
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, level int) {
        if node == nil {
            return
        }
        
        // Ensure result has enough levels
        for len(result) <= level {
            result = append(result, []int{})
        }
        
        result[level] = append(result[level], node.Val)
        
        dfs(node.Left, level+1)
        dfs(node.Right, level+1)
    }
    
    dfs(root, 0)
    return result
}
```
- **Pros**: No queue, uses recursion stack
- **Cons**: O(H) space, may cause stack overflow

#### 2. Two Queue Approach (O(N) time, O(W) space)
```go
func levelOrderTwoQueues(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var result [][]int
    currentLevel := []*TreeNode{root}
    nextLevel := []*TreeNode{}
    
    for len(currentLevel) > 0 {
        levelValues := make([]int, 0, len(currentLevel))
        
        for _, node := range currentLevel {
            levelValues = append(levelValues, node.Val)
            
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }
            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
        }
        
        result = append(result, levelValues)
        currentLevel, nextLevel = nextLevel, []*TreeNode{}
    }
    
    return result
}
```
- **Pros**: Clear separation of levels
- **Cons**: More memory for two queues

#### 3. Sentinel Approach (O(N) time, O(W) space)
```go
func levelOrderSentinel(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var result [][]int
    queue := []*TreeNode{root, nil} // nil marks level end
    
    currentLevel := []int{}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if node == nil {
            // Level finished
            result = append(result, currentLevel)
            currentLevel = []int{}
            
            if len(queue) > 0 {
                queue = append(queue, nil) // Add sentinel for next level
            }
        } else {
            currentLevel = append(currentLevel, node.Val)
            
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
- **Pros**: No need to track queue size
- **Cons**: Extra sentinel nodes, slightly more complex

### Extensions for Interviews:
- **Zigzag Traversal**: Alternate direction at each level
- **Right Side View**: Only last node of each level
- **Level Averages**: Calculate average of each level
- **N-ary Tree**: Extend to nodes with multiple children
- **Bottom-up Level Order**: Return levels from bottom to top
*/
func main() {
	// Test cases
	testCases := [][]interface{}{
		{3, 9, 20, nil, nil, 15, 7},
		{1},
		{},
		{1, 2, 3, 4, 5, 6, 7},
		{1, nil, 2, nil, 3, nil, 4, nil, 5},
		{1, 2, 3, 4, nil, nil, nil, 5},
		{1, 2, 3, nil, nil, nil, 4, nil, nil, 5},
		{1, 2, nil, 3, nil, 4, nil, 5},
	}
	
	for i, nums := range testCases {
		root := createTree(nums)
		result1 := levelOrder(root)
		result2 := levelOrderRecursive(root)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Iterative: %v\n", result1)
		fmt.Printf("  Recursive: %v\n\n", result2)
	}
}
