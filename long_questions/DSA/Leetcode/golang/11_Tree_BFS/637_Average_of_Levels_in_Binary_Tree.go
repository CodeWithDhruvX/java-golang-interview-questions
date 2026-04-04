package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 637. Average of Levels in Binary Tree
// Time: O(N), Space: O(N)
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	
	var result []float64
	queue := []*TreeNode{root}
	
	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0
		
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val
			
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		
		average := float64(levelSum) / float64(levelSize)
		result = append(result, average)
	}
	
	return result
}

// Recursive approach
func averageOfLevelsRecursive(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	
	var result []float64
	
	var dfs func(*TreeNode, int, *[]int)
	dfs = func(node *TreeNode, level int, sums *[]int) {
		if node == nil {
			return
		}
		
		// Ensure sums slice has enough elements
		for len(*sums) <= level {
			*sums = append(*sums, 0)
		}
		
		(*sums)[level] += node.Val
		
		dfs(node.Left, level+1, sums)
		dfs(node.Right, level+1, sums)
	}
	
	// First pass: calculate sums at each level
	sums := []int{}
	dfs(root, 0, &sums)
	
	// Second pass: calculate counts at each level
	counts := make([]int, len(sums))
	
	var countNodes func(*TreeNode, int)
	countNodes = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		
		counts[level]++
		countNodes(node.Left, level+1)
		countNodes(node.Right, level+1)
	}
	
	countNodes(root, 0)
	
	// Calculate averages
	for i := 0; i < len(sums); i++ {
		average := float64(sums[i]) / float64(counts[i])
		result = append(result, average)
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

## 1. ALGORITHM PATTERN: BFS with Level Aggregation
- **Queue-based Traversal**: Process nodes level by level
- **Level Summation**: Accumulate sum of node values at each level
- **Level Counting**: Count nodes at each level for average calculation
- **Floating Point Division**: Calculate average as float64

## 2. PROBLEM CHARACTERISTICS
- **Binary Tree**: Each node has at most 2 children
- **Level Averages**: Calculate average value for each tree level
- **Floating Point Output**: Return array of floating point averages
- **Complete Traversal**: Visit every node exactly once

## 3. SIMILAR PROBLEMS
- Binary Tree Level Order Traversal (LeetCode 102)
- Binary Tree Right Side View (LeetCode 199)
- Average of Levels in N-ary Tree (LeetCode 638)
- Find Largest Value in Each Tree Row (LeetCode 515)

## 4. KEY OBSERVATIONS
- **Level boundary**: Queue size at start = nodes in current level
- **Sum accumulation**: Add node values during level processing
- **Average calculation**: sum / count for each level
- **Precision handling**: Use float64 for accurate division

## 5. VARIATIONS & EXTENSIONS
- **Maximum per Level**: Find maximum value in each level
- **Minimum per Level**: Find minimum value in each level
- **Median per Level**: Find median value in each level
- **N-ary Tree**: Extend to nodes with multiple children

## 6. INTERVIEW INSIGHTS
- Always clarify: "What precision should be used for averages?"
- Edge cases: empty tree, single node, negative values
- Space complexity: O(W) where W is maximum tree width
- Time complexity: O(N) - visit each node once

## 7. COMMON MISTAKES
- Not tracking level size properly, mixing levels
- Using integer division instead of floating point
- Not handling empty tree case
- Forgetting to reset sum for each level
- Precision issues with floating point arithmetic

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(W) space
- **Pre-allocate slices**: Use make with capacity for efficiency
- **Two-pass approach**: Can separate sum and count calculations
- **Recursive alternative**: DFS with level tracking

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like calculating average salary by department level:**
- You have a company organization chart (tree structure)
- You want to know the average salary at each management level
- Process all employees at the same level together
- For each level, sum up all salaries and count employees
- Calculate average by dividing sum by count
- Continue until you reach all employees at the bottom

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary tree root node
2. **Goal**: Calculate average value for each tree level
3. **Output**: Array of floating point averages
4. **Constraint**: Process levels from top to bottom

#### Phase 2: Key Insight Recognition
- **"Level processing"** → Need to process one level at a time
- **"Sum and count"** → Need both sum and count for average calculation
- **"Queue natural fit"** → BFS naturally processes by distance from root
- **"Floating point"** → Must use float64 for accurate division

#### Phase 3: Strategy Development
```
Human thought process:
"I need to calculate the average value at each tree level.
I'll use BFS to process level by level.
For each level, I'll sum up all node values and count the nodes.
Then I'll calculate the average by dividing sum by count.
I'll repeat this for all levels until I reach the leaves."
```

#### Phase 4: Edge Case Handling
- **Empty tree**: Return empty array
- **Single node**: Return array with one average (the node value)
- **Negative values**: Handle naturally with sum and count
- **Large values**: Consider potential overflow, use float64

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tree:     3
         / \
        9  20
           /  \
         15   7

Human thinking:
"I'll calculate averages level by level:

Initial queue: [3]
Level 0: process 3, sum=3, count=1
Average = 3.0/1 = 3.0
Add 9 and 20 to queue
Queue becomes: [9, 20]
Result: [3.0]

Level 1: queue size = 2, process 9 and 20
Sum = 9 + 20 = 29, count = 2
Average = 29.0/2 = 14.5
Add 15 and 7 to queue
Queue becomes: [15, 7]
Result: [3.0, 14.5]

Level 2: queue size = 2, process 15 and 7
Sum = 15 + 7 = 22, count = 2
Average = 22.0/2 = 11.0
Queue becomes: []
Result: [3.0, 14.5, 11.0]

Done! Final result: [3.0, 14.5, 11.0]"
```

#### Phase 6: Intuition Validation
- **Why BFS works**: Naturally processes by distance from root
- **Why sum and count needed**: Average requires both total and count
- **Why O(N) time**: Each node visited exactly once
- **Why O(W) space**: Queue holds at most one level's worth of nodes

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use integer division?"** → Need floating point for accurate averages
2. **"Should I use recursion?"** → Can, but BFS is more intuitive for level-based
3. **"What about very large values?"** → Use float64 to avoid overflow
4. **"Can I optimize further?"** → Already optimal for this problem

### Real-World Analogy
**Like calculating average test scores by grade level:**
- You have a school organized by grade levels (tree structure)
- You want to know the average test score for each grade
- Process all students in the same grade together
- For each grade, sum up all test scores and count students
- Calculate average by dividing total score by number of students
- Continue from kindergarten through all grades

### Human-Readable Pseudocode
```
function averageOfLevels(root):
    if root is nil:
        return []
    
    result = []
    queue = [root]
    
    while queue is not empty:
        levelSize = length(queue)
        levelSum = 0
        
        for i from 0 to levelSize-1:
            node = queue.dequeue()
            levelSum += node.value
            
            if node.left is not nil:
                queue.enqueue(node.left)
            if node.right is not nil:
                queue.enqueue(node.right)
        
        average = float(levelSum) / float(levelSize)
        result.append(average)
    
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

### BFS Average Calculation Process:
```
Initial: queue=[3], result=[]

Level 0: queue=[3], size=1
- Process 3: sum=3, count=1
- Average = 3.0/1 = 3.0
- Add 9, 20 to queue
- queue=[9, 20], result=[3.0]

Level 1: queue=[9, 20], size=2
- Process 9: sum=9
- Process 20: sum=9+20=29, count=2
- Average = 29.0/2 = 14.5
- Add 15, 7 to queue
- queue=[15, 7], result=[3.0, 14.5]

Level 2: queue=[15, 7], size=2
- Process 15: sum=15
- Process 7: sum=15+7=22, count=2
- Average = 22.0/2 = 11.0
- queue=[], result=[3.0, 14.5, 11.0]

Done! Final result: [3.0, 14.5, 11.0]
```

### Key Visualization Points:
- **Level boundary**: queue size at start of each iteration
- **Sum accumulation**: Add node values during level processing
- **Count tracking**: Level size provides count automatically
- **Average calculation**: sum / count as floating point

### Memory Layout Visualization:
```
Queue: [3] → [9, 20] → [15, 7] → []
Sum:   3       29        22
Count: 1       2         2
Avg:   3.0     14.5      11.0
```

### Time Complexity Breakdown:
- **Each node visited**: O(1) work per node
- **Total nodes**: N
- **Total time**: O(N)
- **Space**: O(W) - maximum queue size (tree width)

### Alternative Approaches:

#### 1. Recursive DFS with Level Tracking (O(N) time, O(H) space)
```go
func averageOfLevelsRecursive(root *TreeNode) []float64 {
    if root == nil {
        return []float64{}
    }
    
    var sums []int
    var counts []int
    
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, level int) {
        if node == nil {
            return
        }
        
        // Ensure slices have enough elements
        for len(sums) <= level {
            sums = append(sums, 0)
            counts = append(counts, 0)
        }
        
        sums[level] += node.Val
        counts[level]++
        
        dfs(node.Left, level+1)
        dfs(node.Right, level+1)
    }
    
    dfs(root, 0)
    
    // Calculate averages
    result := make([]float64, len(sums))
    for i := 0; i < len(sums); i++ {
        result[i] = float64(sums[i]) / float64(counts[i])
    }
    
    return result
}
```
- **Pros**: No queue, uses recursion stack
- **Cons**: O(H) space, may cause stack overflow

#### 2. Two-Pass Approach (O(N) time, O(N) space)
```go
func averageOfLevelsTwoPass(root *TreeNode) []float64 {
    if root == nil {
        return []float64{}
    }
    
    // First pass: collect all nodes by level
    var levels [][]int
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
        
        levels = append(levels, currentLevel)
    }
    
    // Second pass: calculate averages
    result := make([]float64, len(levels))
    for i, level := range levels {
        sum := 0
        for _, val := range level {
            sum += val
        }
        result[i] = float64(sum) / float64(len(level))
    }
    
    return result
}
```
- **Pros**: Clear separation of collection and calculation
- **Cons**: O(N) extra space for storing all levels

#### 3. Stream Processing with Rolling Average (O(N) time, O(W) space)
```go
func averageOfLevelsStreaming(root *TreeNode) []float64 {
    if root == nil {
        return []float64{}
    }
    
    var result []float64
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        levelSum := 0
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            levelSum += node.Val
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        // Calculate and store average immediately
        average := float64(levelSum) / float64(levelSize)
        result = append(result, average)
    }
    
    return result
}
```
- **Pros**: Memory efficient, processes on the fly
- **Cons**: Same as basic approach, just different naming

### Extensions for Interviews:
- **Maximum per Level**: Find maximum value in each level
- **Minimum per Level**: Find minimum value in each level
- **Median per Level**: Find median value in each level
- **N-ary Tree**: Extend to nodes with multiple children
- **Standard Deviation**: Calculate standard deviation per level
*/
func main() {
	// Test cases
	testCases := [][]interface{}{
		{3, 9, 20, nil, nil, 15, 7},
		{3, 9, 20, 15, 7},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, 5, nil, nil, nil, 6, 7},
		{1, nil, 2, nil, 3, nil, 4, nil, 5},
		{1, 2, 3, 4, nil, nil, nil, 5},
		{1, 2, 3, nil, nil, nil, 4, nil, nil, 5},
		{1},
		{},
	}
	
	for i, nums := range testCases {
		root := createTree(nums)
		result1 := averageOfLevels(root)
		result2 := averageOfLevelsRecursive(root)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Iterative: [")
		for j, avg := range result1 {
			if j > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%.1f", avg)
		}
		fmt.Printf("]\n")
		
		fmt.Printf("  Recursive: [")
		for j, avg := range result2 {
			if j > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%.1f", avg)
		}
		fmt.Printf("]\n\n")
	}
}
