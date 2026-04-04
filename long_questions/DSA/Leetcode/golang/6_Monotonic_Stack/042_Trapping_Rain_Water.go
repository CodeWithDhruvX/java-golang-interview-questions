package main

import "fmt"

// 42. Trapping Rain Water
// Time: O(N), Space: O(N)
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	
	// Calculate left max for each position
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	
	// Calculate right max for each position
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	
	// Calculate trapped water
	water := 0
	for i := 0; i < n; i++ {
		water += min(leftMax[i], rightMax[i]) - height[i]
	}
	
	return water
}

// Alternative solution using two pointers (O(1) space)
func trapTwoPointers(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0
	
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}
	
	return water
}

// Helper functions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Two-Pass with Left/Right Maximums
- **Left Maximum Array**: Track highest bar to the left of each position
- **Right Maximum Array**: Track highest bar to the right of each position
- **Water Calculation**: Water at position = min(leftMax, rightMax) - height
- **Two-Pointer Optimization**: O(1) space version using two pointers

## 2. PROBLEM CHARACTERISTICS
- **Water Trapping**: Calculate water trapped between bars
- **Boundary Constraint**: Water level limited by shorter boundary
- **Local Optimum**: Each position's water depends on local maxima
- **Global Solution**: Sum of local water amounts

## 3. SIMILAR PROBLEMS
- Largest Rectangle in Histogram (LeetCode 84) - Similar bar-based problem
- Container With Most Water (LeetCode 11) - Two-pointer water problem
- Trapping Rain Water II (LeetCode 407) - 2D version
- Rain Water Trapping Variations - Different constraints

## 4. KEY OBSERVATIONS
- **Water level formula**: Water[i] = min(maxLeft[i], maxRight[i]) - height[i]
- **Boundary dependence**: Water at any position limited by boundaries
- **Independent calculation**: Each position can be calculated independently
- **Two-pass natural fit**: Need left and right information

## 5. VARIATIONS & EXTENSIONS
- **2D version**: Trapping rain water in 2D matrix
- **Multiple queries**: Answer queries for specific ranges
- **Streaming data**: Handle data that arrives over time
- **Variable boundaries**: Different boundary conditions

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can height array be modified? What about empty input?"
- Edge cases: empty array, single element, monotonic arrays
- Time complexity: O(N) time, O(N) space or O(1) space with two pointers
- Space optimization: Two-pointer approach reduces space to O(1)

## 7. COMMON MISTAKES
- Not handling empty array case
- Using O(N²) nested loop approach
- Not calculating left/right maxima correctly
- Forgetting to take minimum of left and right maxima
- Not handling negative heights (though typically non-negative)

## 8. OPTIMIZATION STRATEGIES
- **Two-pass approach**: O(N) time, O(N) space
- **Two-pointer optimization**: O(N) time, O(1) space
- **Early termination**: Not applicable (need full processing)
- **In-place calculation**: Use output array for intermediate storage

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like water filling between walls:**
- You have a series of walls of different heights
- When it rains, water gets trapped between higher walls
- At any position, water level is limited by the shorter wall on either side
- The amount of water trapped is the difference between water level and ground
- You need to calculate this for every position

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of non-negative integers representing wall heights
2. **Goal**: Calculate total water trapped after rain
3. **Constraint**: Water can only be trapped between higher walls
4. **Output**: Total units of water trapped

#### Phase 2: Key Insight Recognition
- **"Boundary dependence"** → Water at each position depends on boundaries
- **"Left/right maxima"** → Need highest wall on each side
- **"Minimum principle"** → Water level limited by shorter boundary
- **"Two-pass natural fit"** → Calculate left and right maxima separately

#### Phase 3: Strategy Development
```
Human thought process:
"I need to calculate water trapped at each position.
For any position, water level is determined by the shorter of the tallest walls
to its left and right.
I'll first calculate the tallest wall to the left of each position.
Then I'll calculate the tallest wall to the right of each position.
Finally, I'll sum min(leftMax, rightMax) - height for all positions."
```

#### Phase 4: Edge Case Handling
- **Empty array**: No water can be trapped, return 0
- **Single element**: No space for water, return 0
- **Monotonic arrays**: Increasing or decreasing arrays trap no water
- **Equal heights**: Flat surfaces trap no water

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]

Human thinking:
"First, I'll calculate left max for each position:
Position 0: leftMax = 0
Position 1: leftMax = max(0, 1) = 1
Position 2: leftMax = max(1, 0) = 1
Position 3: leftMax = max(1, 2) = 2
...and so on

Then right max for each position:
Position 11: rightMax = 1
Position 10: rightMax = max(1, 2) = 2
Position 9: rightMax = max(2, 1) = 2
Position 8: rightMax = max(2, 2) = 2
...and so on

Finally, calculate water at each position:
Position 0: min(0, 3) - 0 = 0
Position 1: min(1, 3) - 1 = 0
Position 2: min(1, 3) - 0 = 1
Position 3: min(2, 3) - 2 = 0
Position 4: min(2, 3) - 1 = 1
...and so on

Total water = sum of all water amounts"
```

#### Phase 6: Intuition Validation
- **Why left/right maxima work**: Capture boundary constraints
- **Why minimum principle**: Water level limited by shorter boundary
- **Why O(N) time**: Each array processed once
- **Why O(1) space possible**: Two-pointer optimization

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use nested loops?"** → O(N²) time, too slow for large inputs
2. **"Should I use stack?"** → More complex, two-pass is simpler
3. **"What about two pointers?"** → Yes, can optimize space to O(1)
4. **"Can I optimize further?"** → Two-pointer approach is already optimal

### Real-World Analogy
**Like calculating water storage in a series of dams:**
- You have a series of dams of different heights along a river
- When water flows, it gets trapped between higher dams
- At any point, water level is limited by the shorter dam upstream or downstream
- The amount of water stored is the difference between water level and riverbed
- Engineers need to calculate total storage capacity

### Human-Readable Pseudocode
```
function trap(height):
    if height is empty:
        return 0
    
    leftMax = array of size len(height)
    rightMax = array of size len(height)
    
    // Calculate left max
    leftMax[0] = height[0]
    for i from 1 to len(height)-1:
        leftMax[i] = max(leftMax[i-1], height[i])
    
    // Calculate right max
    rightMax[len(height)-1] = height[len(height)-1]
    for i from len(height)-2 down to 0:
        rightMax[i] = max(rightMax[i+1], height[i])
    
    // Calculate trapped water
    water = 0
    for i from 0 to len(height)-1:
        water += min(leftMax[i], rightMax[i]) - height[i]
    
    return water
```

### Execution Visualization

### Example: height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
```
Left Max Calculation:
[0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3]

Right Max Calculation:
[3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 1]

Water Calculation:
[min(0,3)-0, min(1,3)-1, min(1,3)-0, min(2,3)-2, min(2,3)-1, 
 min(2,3)-0, min(2,3)-1, min(3,3)-3, min(3,2)-2, min(3,2)-1,
 min(3,2)-2, min(3,1)-1]

= [0, 0, 1, 0, 1, 2, 1, 0, 0, 1, 0, 0]

Total water = 6
```

### Key Visualization Points:
- **Left max accumulation**: Building maximum from left to right
- **Right max accumulation**: Building maximum from right to left
- **Water level calculation**: Minimum of left and right maxima
- **Total accumulation**: Sum of all water amounts

### Memory Layout Visualization:
```
Height:     [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
Left Max:   [0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3]
Right Max:  [3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 1]
Water:      [0, 0, 1, 0, 1, 2, 1, 0, 0, 1, 0, 0]

At position 2: leftMax=1, rightMax=3, height=0
Water level = min(1,3) = 1
Trapped water = 1 - 0 = 1
```

### Time Complexity Breakdown:
- **Two-pass approach**: O(N) time, O(N) space
- **Two-pointer approach**: O(N) time, O(1) space
- **Nested loops**: O(N²) time, O(1) space (too slow)
- **Stack approach**: O(N) time, O(N) space (more complex)

### Alternative Approaches:

#### 1. Two-Pointer Approach (O(N) time, O(1) space)
```go
func trapTwoPointers(height []int) int {
    left, right := 0, len(height)-1
    leftMax, rightMax := 0, 0
    water := 0
    
    for left < right {
        if height[left] < height[right] {
            if height[left] >= leftMax {
                leftMax = height[left]
            } else {
                water += leftMax - height[left]
            }
            left++
        } else {
            if height[right] >= rightMax {
                rightMax = height[right]
            } else {
                water += rightMax - height[right]
            }
            right--
        }
    }
    
    return water
}
```
- **Pros**: O(1) space, optimal time
- **Cons**: More complex logic to understand

#### 2. Stack Approach (O(N) time, O(N) space)
```go
func trapStack(height []int) int {
    stack := []int{}
    water := 0
    
    for i, h := range height {
        for len(stack) > 0 && h > height[stack[len(stack)-1]] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            
            if len(stack) > 0 {
                distance := i - stack[len(stack)-1] - 1
                boundedHeight := min(h, height[stack[len(stack)-1]]) - height[top]
                water += distance * boundedHeight
            }
        }
        stack = append(stack, i)
    }
    
    return water
}
```
- **Pros**: Single pass, intuitive water trapping
- **Cons**: O(N) space, more complex implementation

#### 3. Brute Force (O(N²) time, O(1) space)
```go
func trapBrute(height []int) int {
    water := 0
    
    for i := 1; i < len(height)-1; i++ {
        leftMax := 0
        for j := 0; j <= i; j++ {
            leftMax = max(leftMax, height[j])
        }
        
        rightMax := 0
        for j := i; j < len(height); j++ {
            rightMax = max(rightMax, height[j])
        }
        
        water += min(leftMax, rightMax) - height[i]
    }
    
    return water
}
```
- **Pros**: Simple to understand
- **Cons**: O(N²) time, too slow for large inputs

### Extensions for Interviews:
- **2D Version**: Trapping rain water in 2D matrix
- **Multiple Queries**: Answer queries for specific ranges efficiently
- **Streaming Data**: Handle data that arrives over time
- **Variable Boundaries**: Different boundary conditions
- **Maximum Water**: Find position with maximum trapped water
*/
func main() {
	// Test cases
	testCases := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
		{2, 0, 2},
		{3, 0, 0, 2, 0, 4},
		{0, 0, 0, 0},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{2, 2, 2, 2},
		{1},
		{},
		{0, 2, 0},
		{4, 2, 0, 3, 2, 5, 2, 1, 5, 2},
	}
	
	for i, height := range testCases {
		result1 := trap(height)
		result2 := trapTwoPointers(height)
		fmt.Printf("Test Case %d: %v -> Water (O(N) space): %d, Water (O(1) space): %d\n", 
			i+1, height, result1, result2)
	}
}
