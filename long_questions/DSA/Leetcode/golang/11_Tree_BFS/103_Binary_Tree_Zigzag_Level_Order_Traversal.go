package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 103. Binary Tree Zigzag Level Order Traversal
// Time: O(N), Space: O(N)
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	
	var result [][]int
	queue := []*TreeNode{root}
	leftToRight := true
	
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]int, 0, levelSize)
		
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			
			if leftToRight {
				currentLevel = append(currentLevel, node.Val)
			} else {
				// Insert at beginning for right-to-left traversal
				currentLevel = append([]int{node.Val}, currentLevel...)
			}
			
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		
		result = append(result, currentLevel)
		leftToRight = !leftToRight
	}
	
	return result
}

// Alternative approach using two stacks
func zigzagLevelOrderStacks(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	
	var result [][]int
	currentStack := []*TreeNode{root}
	nextStack := []*TreeNode{}
	leftToRight := true
	
	for len(currentStack) > 0 {
		levelSize := len(currentStack)
		currentLevel := make([]int, 0, levelSize)
		
		for i := 0; i < levelSize; i++ {
			node := currentStack[len(currentStack)-1]
			currentStack = currentStack[:len(currentStack)-1]
			currentLevel = append(currentLevel, node.Val)
			
			if leftToRight {
				if node.Left != nil {
					nextStack = append(nextStack, node.Left)
				}
				if node.Right != nil {
					nextStack = append(nextStack, node.Right)
				}
			} else {
				if node.Right != nil {
					nextStack = append(nextStack, node.Right)
				}
				if node.Left != nil {
					nextStack = append(nextStack, node.Left)
				}
			}
		}
		
		result = append(result, currentLevel)
		currentStack, nextStack = nextStack, []*TreeNode{}
		leftToRight = !leftToRight
	}
	
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

## 1. ALGORITHM PATTERN: BFS with Direction Alternation
- **Queue-based Traversal**: Process nodes level by level
- **Direction Toggle**: Alternate between left-to-right and right-to-left
- **Flexible Insertion**: Use prepend/append based on current direction
- **Level Processing**: Complete each level before direction change

## 2. PROBLEM CHARACTERISTICS
- **Binary Tree**: Each node has at most 2 children
- **Zigzag Pattern**: Alternate direction at each level
- **2D Output**: Return array of arrays, one subarray per level
- **Direction Memory**: Track current traversal direction

## 3. SIMILAR PROBLEMS
- Binary Tree Level Order Traversal (LeetCode 102)
- Binary Tree Right Side View (LeetCode 199)
- Average of Levels in Binary Tree (LeetCode 637)
- N-ary Tree Level Order Traversal (LeetCode 429)

## 4. KEY OBSERVATIONS
- **Direction flag**: Boolean to track current traversal direction
- **Insertion strategy**: Append for LTR, prepend for RTL
- **Child order**: Always enque children left-to-right regardless of direction
- **Level boundary**: Queue size still identifies level boundaries

## 5. VARIATIONS & EXTENSIONS
- **Custom Pattern**: Any custom traversal pattern
- **Spiral Traversal**: Different spiral patterns
- **Multiple Queues**: Use separate queues for different approaches
- **Stack Implementation**: Use stacks instead of queue for natural RTL

## 6. INTERVIEW INSIGHTS
- Always clarify: "Should the first level be left-to-right?"
- Edge cases: empty tree, single node, unbalanced tree
- Space complexity: O(W) where W is maximum tree width
- Time complexity: O(N) - visit each node once

## 7. COMMON MISTAKES
- Forgetting to toggle direction flag
- Using wrong insertion order for RTL levels
- Mixing up child enqueuing order with output order
- Not handling empty tree case properly
- Using inefficient slice operations for prepend

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(W) space
- **Two stacks approach**: More natural for alternating directions
- **Pre-allocate slices**: Use make with capacity for efficiency
- **Direction optimization**: Use index-based approach instead of prepend

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like reading a book with alternating reading directions:**
- You have a book with chapters organized in levels (tree structure)
- Read the first chapter (root) left-to-right normally
- Read the next level right-to-left (like reading Hebrew/Arabic)
- Alternate reading direction for each subsequent level
- For each chapter, note the chapters it references for the next level
- Continue until you reach all chapters at the bottom

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary tree root node
2. **Goal**: Return nodes grouped by level with alternating directions
3. **Output**: 2D array where each subarray represents one level
4. **Constraint**: First level left-to-right, then alternate

#### Phase 2: Key Insight Recognition
- **"Direction alternation"** → Need to toggle between LTR and RTL
- **"Insertion strategy"** → Different append/prepend based on direction
- **"Queue still works"** → BFS still processes by distance from root
- **"Output order vs enqueue order"** → Children always enqueued LTR

#### Phase 3: Strategy Development
```
Human thought process:
"I need to visit the tree level by level, but alternate directions.
I'll use a queue for BFS and a direction flag.
For left-to-right levels, I'll append values normally.
For right-to-left levels, I'll prepend values to reverse the order.
Children are always enqueued left-to-right for the next level."
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return empty array
- **Single node**: Return array with one subarray containing the root
- **Unbalanced tree**: Handle naturally with queue processing
- **Direction initialization**: Start with left-to-right

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:     3
         / \
        9  20
           /  \
         15   7

Human thinking:
"I'll process level by level with alternating directions:

Initial queue: [3], direction = LTR
Level 0 (LTR): process 3, add 9 and 20 to queue
Append 3 to result: [3]
Queue becomes: [9, 20]
Toggle direction: RTL

Level 1 (RTL): queue size = 2, process 9 and 20
Process 9: prepend 9 → [9]
Process 20: prepend 20 → [20, 9]
Add 15 and 7 to queue
Queue becomes: [15, 7]
Result: [[3], [20, 9]]
Toggle direction: LTR

Level 2 (LTR): queue size = 2, process 15 and 7
Process 15: append 15 → [15]
Process 7: append 7 → [15, 7]
Queue becomes: []
Result: [[3], [20, 9], [15, 7]]

Done! Final result: [[3], [20, 9], [15, 7]]"
```

#### Phase 6: Intuition Validation
- **Why queue works**: BFS still processes by distance from root
- **Why direction matters**: Only affects output order, not traversal order
- **Why O(N) time**: Each node visited exactly once
- **Why O(W) space**: Queue holds at most one level's worth of nodes

### Common Human Pitfalls & How to Avoid Them
1. **"Why not reverse the queue?"** → That would affect traversal, not just output
2. **"Should I enqueue children differently?"** → No, always enqueue LTR for next level
3. **"What about prepend efficiency?"** → In Go, prepend is O(n), consider alternatives
4. **"Can I use stacks instead?"** → Yes, two stacks approach is more natural

### Real-World Analogy
**Like processing documents in alternating reading directions:**
- You have a document system organized in hierarchical levels
- Level 0: Executive summary (read normally LTR)
- Level 1: Department reports (read RTL for variety)
- Level 2: Team details (read LTR again)
- For each document, reference subordinate documents for next level
- Alternate reading style keeps the review process engaging

### Human-Readable Pseudocode
```
function zigzagLevelOrder(root):
    if root is nil:
        return []
    
    result = []
    queue = [root]
    leftToRight = true
    
    while queue is not empty:
        levelSize = length(queue)
        currentLevel = []
        
        for i from 0 to levelSize-1:
            node = queue.dequeue()
            
            if leftToRight:
                currentLevel.append(node.value)
            else:
                currentLevel.prepend(node.value)
            
            if node.left is not nil:
                queue.enqueue(node.left)
            if node.right is not nil:
                queue.enqueue(node.right)
        
        result.append(currentLevel)
        leftToRight = not leftToRight
    
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

### Zigzag BFS Process:
```
Initial: queue=[3], result=[], direction=LTR

Level 0 (LTR): queue=[3], size=1
- Process 3: append 3 → [3], add 9, 20 to queue
- queue=[9, 20], result=[[3]]
- direction=RTL

Level 1 (RTL): queue=[9, 20], size=2
- Process 9: prepend 9 → [9]
- Process 20: prepend 20 → [20, 9], add 15, 7 to queue
- queue=[15, 7], result=[[3], [20, 9]]
- direction=LTR

Level 2 (LTR): queue=[15, 7], size=2
- Process 15: append 15 → [15]
- Process 7: append 7 → [15, 7]
- queue=[], result=[[3], [20, 9], [15, 7]]

Done!
```

### Key Visualization Points:
- **Direction toggle**: Boolean flag switches after each level
- **Insertion strategy**: Append for LTR, prepend for RTL
- **Child enqueuing**: Always left-to-right regardless of direction
- **Level boundary**: Queue size still identifies current level

### Memory Layout Visualization:
```
Queue: [3] → [9, 20] → [15, 7] → []
Direction: LTR → RTL → LTR
Output: [3] → [20, 9] → [15, 7]
```

### Time Complexity Breakdown:
- **Each node visited**: O(1) work per node
- **Prepend operations**: O(k) where k is level size (worst case O(N²) total)
- **Total time**: O(N²) in worst case due to prepend, O(N) average
- **Space**: O(W) - maximum queue size (tree width)

### Alternative Approaches:

#### 1. Two Stacks Approach (O(N) time, O(W) space)
```go
func zigzagLevelOrderStacks(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var result [][]int
    currentStack := []*TreeNode{root}
    nextStack := []*TreeNode{}
    leftToRight := true
    
    for len(currentStack) > 0 {
        levelSize := len(currentStack)
        currentLevel := make([]int, 0, levelSize)
        
        for i := 0; i < levelSize; i++ {
            node := currentStack[len(currentStack)-1]
            currentStack = currentStack[:len(currentStack)-1]
            currentLevel = append(currentLevel, node.Val)
            
            if leftToRight {
                if node.Left != nil {
                    nextStack = append(nextStack, node.Left)
                }
                if node.Right != nil {
                    nextStack = append(nextStack, node.Right)
                }
            } else {
                if node.Right != nil {
                    nextStack = append(nextStack, node.Right)
                }
                if node.Left != nil {
                    nextStack = append(nextStack, node.Left)
                }
            }
        }
        
        result = append(result, currentLevel)
        currentStack, nextStack = nextStack, []*TreeNode{}
        leftToRight = !leftToRight
    }
    
    return result
}
```
- **Pros**: Natural alternating directions, O(N) time
- **Cons**: More complex with two stacks

#### 2. Index-based Approach (O(N) time, O(W) space)
```go
func zigzagLevelOrderIndex(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var result [][]int
    queue := []*TreeNode{root}
    leftToRight := true
    
    for len(queue) > 0 {
        levelSize := len(queue)
        currentLevel := make([]int, levelSize)
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            if leftToRight {
                currentLevel[i] = node.Val
            } else {
                currentLevel[levelSize-1-i] = node.Val
            }
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        result = append(result, currentLevel)
        leftToRight = !leftToRight
    }
    
    return result
}
```
- **Pros**: O(N) time, no expensive prepend operations
- **Cons**: Need to pre-allocate slice with known size

#### 3. Deque Approach (O(N) time, O(W) space)
```go
func zigzagLevelOrderDeque(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var result [][]int
    deque := []*TreeNode{root}
    leftToRight := true
    
    for len(deque) > 0 {
        levelSize := len(deque)
        currentLevel := make([]int, 0, levelSize)
        
        for i := 0; i < levelSize; i++ {
            var node *TreeNode
            if leftToRight {
                node = deque[0]
                deque = deque[1:]
            } else {
                node = deque[len(deque)-1]
                deque = deque[:len(deque)-1]
            }
            
            currentLevel = append(currentLevel, node.Val)
            
            if leftToRight {
                if node.Left != nil {
                    deque = append(deque, node.Left)
                }
                if node.Right != nil {
                    deque = append(deque, node.Right)
                }
            } else {
                if node.Right != nil {
                    deque = append([]*TreeNode{node.Right}, deque...)
                }
                if node.Left != nil {
                    deque = append([]*TreeNode{node.Left}, deque...)
                }
            }
        }
        
        result = append(result, currentLevel)
        leftToRight = !leftToRight
    }
    
    return result
}
```
- **Pros**: Efficient operations from both ends
- **Cons**: More complex deque management

### Extensions for Interviews:
- **Custom Patterns**: Any custom traversal pattern
- **Spiral Traversal**: Different spiral patterns
- **Level Reversal**: Reverse specific levels
- **N-ary Tree**: Extend to nodes with multiple children
- **Memory Optimization**: Minimize space usage for large trees
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
		result1 := zigzagLevelOrder(root)
		result2 := zigzagLevelOrderStacks(root)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Queue method: %v\n", result1)
		fmt.Printf("  Stack method: %v\n\n", result2)
	}
}
