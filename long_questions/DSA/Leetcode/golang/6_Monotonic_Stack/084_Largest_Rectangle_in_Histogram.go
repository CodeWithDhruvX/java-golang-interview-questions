package main

import "fmt"

// 84. Largest Rectangle in Histogram
// Time: O(N), Space: O(N)
func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := []int{-1} // Initialize with -1 to handle edge case
	maxArea := 0
	
	for i := 0; i < n; i++ {
		// While current height is less than height at stack top
		for stack[len(stack)-1] != -1 && heights[i] < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i - stack[len(stack)-1] - 1
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	
	// Process remaining bars in stack
	for stack[len(stack)-1] != -1 {
		height := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		width := n - stack[len(stack)-1] - 1
		area := height * width
		if area > maxArea {
			maxArea = area
		}
	}
	
	return maxArea
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Monotonic Stack for Rectangle Areas
- **Monotonic Increasing Stack**: Maintain stack of increasing heights
- **Area Calculation**: Calculate area when popping from stack
- **Width Determination**: Use current index and new stack top for width
- **Sentinel Handling**: Use -1 sentinel to handle edge cases

## 2. PROBLEM CHARACTERISTICS
- **Rectangle Area**: Find largest rectangle in histogram
- **Bar Constraints**: Rectangle must be formed by contiguous bars
- **Height Limitation**: Rectangle height limited by shortest bar in range
- **Contiguous Range**: Must use consecutive histogram bars

## 3. SIMILAR PROBLEMS
- Trapping Rain Water (LeetCode 42) - Similar bar-based problem
- Largest Rectangle in Binary Matrix (LeetCode 85) - 2D extension
- Maximal Rectangle (LeetCode 221) - Rectangle in binary matrix
- Next Greater Element II (LeetCode 503) - Circular monotonic stack

## 4. KEY OBSERVATIONS
- **Stack property**: Stack maintains increasing heights
- **Area formula**: Area = height × width where width = right - left - 1
- **Pop condition**: Pop when current height < stack top height
- **Width calculation**: Use current index as right boundary

## 5. VARIATIONS & EXTENSIONS
- **2D version**: Largest rectangle in binary matrix
- **Circular histogram**: Circular arrangement of bars
- **Multiple queries**: Query largest rectangle in subranges
- **Streaming data**: Handle data that arrives over time

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can heights be zero? What about empty input?"
- Edge cases: empty array, single bar, all equal heights
- Time complexity: O(N) time, O(N) space
- Stack invariant: Always maintain increasing heights

## 7. COMMON MISTAKES
- Not handling empty array case
- Using nested loops (O(N²) time)
- Incorrect width calculation
- Not processing remaining stack elements
- Forgetting sentinel value initialization

## 8. OPTIMIZATION STRATEGIES
- **Monotonic stack**: O(N) time, O(N) space
- **Single pass**: Process each element once
- **Efficient pop**: Calculate area immediately when popping
- **Sentinel optimization**: Use -1 to simplify boundary handling

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the largest billboard you can build:**
- You have buildings of different heights along a street
- You want to build the largest possible billboard using consecutive buildings
- The billboard height is limited by the shortest building in the range
- You scan from left to right, keeping track of potential building sites
- When you find a shorter building, you calculate areas for taller ones

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of non-negative integers representing bar heights
2. **Goal**: Find area of largest rectangle formed by contiguous bars
3. **Constraint**: Rectangle height limited by shortest bar in range
4. **Output**: Maximum rectangle area

#### Phase 2: Key Insight Recognition
- **"Monotonic stack natural fit"** → Maintain increasing heights
- **"Area on pop"** → Calculate rectangle when height decreases
- **"Width calculation"** → Use indices to determine width
- **"Sentinel handling"** → Use -1 to simplify edge cases

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the largest rectangle in the histogram.
I'll use a stack to keep track of increasing bar heights.
When I encounter a shorter bar, I know the taller bars can't extend further.
I'll pop from the stack and calculate areas using the current position as right boundary.
The width is determined by the current position and the new stack top.
I'll process all remaining bars at the end."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 (no bars, no rectangle)
- **Single bar**: Return height of that bar
- **All equal heights**: Return height × number of bars
- **Zero heights**: Skip zero height bars in calculations

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: [2, 1, 5, 6, 2, 3]

Human thinking:
"I'll process with a stack initialized with -1:

i=0, height=2: Stack [-1], push 0 → Stack [-1, 0]
i=1, height=1: 1 < height[0]=2, pop 0
  height=2, width=1-(-1)-1=1, area=2×1=2, maxArea=2
  Stack [-1], push 1 → Stack [-1, 1]
i=2, height=5: Stack [-1, 1], push 2 → Stack [-1, 1, 2]
i=3, height=6: Stack [-1, 1, 2], push 3 → Stack [-1, 1, 2, 3]
i=4, height=2: 2 < height[3]=6, pop 3
  height=6, width=4-2-1=1, area=6×1=6, maxArea=6
  2 < height[2]=5, pop 2
  height=5, width=4-1-1=2, area=5×2=10, maxArea=10
  Stack [-1, 1], push 4 → Stack [-1, 1, 4]
i=5, height=3: Stack [-1, 1, 4], push 5 → Stack [-1, 1, 4, 5]

Process remaining stack:
pop 5: height=3, width=6-4-1=1, area=3×1=3
pop 4: height=2, width=6-1-1=4, area=2×4=8
pop 1: height=1, width=6-(-1)-1=6, area=1×6=6

Final maxArea = 10 ✓"
```

#### Phase 6: Intuition Validation
- **Why monotonic stack works**: Maintains potential rectangle boundaries
- **Why area on pop**: When height decreases, taller bars can't extend further
- **Why width formula**: Current index - new stack top - 1 gives correct width
- **Why O(N) time**: Each bar pushed and popped at most once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use nested loops?"** → O(N²) time, too slow for large inputs
2. **"Should I use divide and conquer?"** → More complex, stack is simpler
3. **"What about width calculation?"** → Use current index and stack top
4. **"Can I optimize further?"** → Monotonic stack is already optimal

### Real-World Analogy
**Like finding the largest possible building foundation:**
- You have plots of land with different height restrictions along a street
- You want to build the largest possible building foundation using consecutive plots
- The building height is limited by the shortest plot in the range
- As you survey the street, you note potential building sites
- When you find a lower restriction, you finalize calculations for higher ones

### Human-Readable Pseudocode
```
function largestRectangleArea(heights):
    stack = [-1]  // Initialize with sentinel
    maxArea = 0
    
    for i from 0 to len(heights)-1:
        while stack.top != -1 and heights[i] < heights[stack.top]:
            height = heights[stack.top]
            stack.pop()
            width = i - stack.top - 1
            area = height * width
            maxArea = max(maxArea, area)
        stack.push(i)
    
    // Process remaining bars
    while stack.top != -1:
        height = heights[stack.top]
        stack.pop()
        width = len(heights) - stack.top - 1
        area = height * width
        maxArea = max(maxArea, area)
    
    return maxArea
```

### Execution Visualization

### Example: heights = [2, 1, 5, 6, 2, 3]
```
Stack Evolution and Area Calculations:
Initial: Stack = [-1], maxArea = 0

i=0, h=2: Stack = [-1, 0]
i=1, h=1: Pop 0 (height=2, width=1, area=2), maxArea=2, Stack = [-1, 1]
i=2, h=5: Stack = [-1, 1, 2]
i=3, h=6: Stack = [-1, 1, 2, 3]
i=4, h=2: Pop 3 (height=6, width=1, area=6), maxArea=6
          Pop 2 (height=5, width=2, area=10), maxArea=10
          Stack = [-1, 1, 4]
i=5, h=3: Stack = [-1, 1, 4, 5]

Final Processing:
Pop 5: height=3, width=1, area=3
Pop 4: height=2, width=4, area=8
Pop 1: height=1, width=6, area=6

Final maxArea = 10
```

### Key Visualization Points:
- **Stack invariant**: Always maintains increasing heights
- **Area calculation**: Height × width when popping
- **Width determination**: Current index - new stack top - 1
- **Sentinel handling**: -1 simplifies boundary calculations

### Memory Layout Visualization:
```
Heights: [2, 1, 5, 6, 2, 3]
Index:    0  1  2  3  4  5

Stack at i=3: [-1, 1, 2, 3]
              ↑  ↑  ↑  ↑
              -1 1  2  3 (indices)

When popping height=6 at index 3:
New stack top = 2
Width = 3 - 2 - 1 = 1
Area = 6 × 1 = 6

Rectangle spans from index 3 to 3 (just bar 3)
```

### Time Complexity Breakdown:
- **Monotonic stack**: O(N) time, O(N) space
- **Nested loops**: O(N²) time, O(1) space (too slow)
- **Divide and conquer**: O(N log N) time, O(log N) space (more complex)
- **Segment tree**: O(N log N) time, O(N) space (overkill)

### Alternative Approaches:

#### 1. Divide and Conquer (O(N log N) time, O(log N) space)
```go
func largestRectangleAreaDivide(heights []int) int {
    return divideAndConquer(heights, 0, len(heights)-1)
}

func divideAndConquer(heights []int, left, right int) int {
    if left > right {
        return 0
    }
    if left == right {
        return heights[left]
    }
    
    minIndex := left
    for i := left; i <= right; i++ {
        if heights[i] < heights[minIndex] {
            minIndex = i
        }
    }
    
    leftArea := divideAndConquer(heights, left, minIndex-1)
    rightArea := divideAndConquer(heights, minIndex+1, right)
    crossArea := heights[minIndex] * (right - left + 1)
    
    return max(leftArea, rightArea, crossArea)
}
```
- **Pros**: Conceptually simple
- **Cons**: O(N²) worst case, more complex

#### 2. Brute Force (O(N²) time, O(1) space)
```go
func largestRectangleAreaBrute(heights []int) int {
    maxArea := 0
    
    for i := 0; i < len(heights); i++ {
        minHeight := heights[i]
        for j := i; j < len(heights); j++ {
            minHeight = min(minHeight, heights[j])
            area := minHeight * (j - i + 1)
            maxArea = max(maxArea, area)
        }
    }
    
    return maxArea
}
```
- **Pros**: Simple to understand
- **Cons**: O(N²) time, too slow for large inputs

#### 3. Segment Tree (O(N log N) time, O(N) space)
```go
type SegmentNode struct {
    minIndex int
    left     *SegmentNode
    right    *SegmentNode
}

func largestRectangleAreaSegment(heights []int) int {
    root := buildSegmentTree(heights, 0, len(heights)-1)
    return querySegmentTree(heights, root, 0, len(heights)-1)
}
```
- **Pros**: Efficient for multiple queries
- **Cons**: Complex implementation, overkill for single query

### Extensions for Interviews:
- **2D Version**: Largest rectangle in binary matrix
- **Circular Histogram**: Bars arranged in circle
- **Multiple Queries**: Query largest rectangle in subranges
- **Streaming Data**: Handle data that arrives over time
- **Count Rectangles**: Count all possible rectangles above certain area
*/
func main() {
	// Test cases
	testCases := [][]int{
		{2, 1, 5, 6, 2, 3},
		{2, 4},
		{1, 1, 1, 1},
		{4, 2, 0, 3, 2, 5},
		{6, 5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5, 6},
		{2, 1, 2, 3, 1},
		{0},
		{},
		{1},
		{1000, 1000, 1000},
	}
	
	for i, heights := range testCases {
		result := largestRectangleArea(heights)
		fmt.Printf("Test Case %d: %v -> Largest rectangle area: %d\n", 
			i+1, heights, result)
	}
}
