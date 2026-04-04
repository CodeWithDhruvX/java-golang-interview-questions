package main

import "fmt"

// 11. Container With Most Water
// Time: O(N), Space: O(1)
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0
	
	for left < right {
		// Calculate current area
		width := right - left
		height1 := height[left]
		height2 := height[right]
		currentHeight := height1
		if height2 < height1 {
			currentHeight = height2
		}
		
		currentArea := width * currentHeight
		if currentArea > maxWater {
			maxWater = currentArea
		}
		
		// Move the pointer with smaller height
		if height1 < height2 {
			left++
		} else {
			right--
		}
	}
	
	return maxWater
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Two Pointers with Greedy Approach
- **Left Pointer**: Starts at the beginning of the array
- **Right Pointer**: Starts at the end of the array
- **Area Calculation**: width × min(height[left], height[right])
- **Greedy Movement**: Move pointer with smaller height inward

## 2. PROBLEM CHARACTERISTICS
- **Container Formation**: Two vertical lines form container sides
- **Area Maximization**: Want to maximize width × height
- **Height Constraint**: Container height limited by shorter side
- **Width Reduction**: Moving pointers inward reduces width

## 3. SIMILAR PROBLEMS
- Trapping Rain Water (LeetCode 42)
- 3Sum Closest (LeetCode 16)
- Two Sum II (LeetCode 167)
- Sort Colors (LeetCode 75)

## 4. KEY OBSERVATIONS
- **Width Decreases**: As pointers move inward, width always decreases
- **Height Matters**: Only taller lines can compensate for width loss
- **Greedy Choice**: Moving shorter line gives chance for taller line
- **Optimal Substructure**: Local optimal choices lead to global optimum

## 5. VARIATIONS & EXTENSIONS
- **3D Container**: Find maximum volume container with 3+ lines
- **Obstacles**: Some positions cannot be used as container sides
- **Multiple Containers**: Find k largest non-overlapping containers
- **Circular Array**: Container can wrap around array boundaries

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can lines be at same position? What about zero height?"
- Edge cases: empty array, single element, all same heights
- Space complexity: O(1) - only two pointers and variables
- Time complexity: O(N) - single pass with two pointers

## 7. COMMON MISTAKES
- Moving both pointers instead of just the smaller one
- Not calculating area correctly (using max height instead of min)
- Forgetting to update maximum area after each calculation
- Using nested loops instead of two-pointer approach
- Not handling edge cases (empty array, etc.)

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(1) space
- **Early termination**: If remaining possible area ≤ current max
- **Parallel processing**: Not applicable due to pointer dependency
- **Cache optimization**: Sequential memory access is already optimal

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like choosing the best container for water:**
- You have many vertical sticks of different heights
- You want to pick two sticks to hold the most water
- The water level is limited by the shorter stick
- The distance between sticks determines width
- You start with the widest possible container
- Then you try to find better ones by moving inward

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of heights representing vertical lines
2. **Goal**: Find two lines that form container with max area
3. **Area Formula**: (right - left) × min(height[left], height[right])
4. **Constraint**: Must use exactly two lines

#### Phase 2: Key Insight Recognition
- **"Width vs Height trade-off"** → Wider containers have less height potential
- **"Greedy elimination"** → Shorter line can never be part of better container
- **"Two-pointer efficiency"** → Can explore all possibilities in O(N)
- **"Monotonic movement"** → Pointers only move inward, never outward

#### Phase 3: Strategy Development
```
Human thought process:
"I'll start with the widest possible container (first and last lines).
The area is limited by the shorter line.
If I move the taller line inward, width decreases but height doesn't increase,
so I'll definitely get less area.
But if I move the shorter line inward, I might find a taller line
that compensates for the width loss.
So I'll always move the shorter line and keep track of the maximum area."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 (no container possible)
- **Single element**: Return 0 (need two lines)
- **All same heights**: Maximum area is with outermost lines
- **Zero heights**: Return 0 (no container can hold water)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [1, 8, 6, 2, 5, 4, 8, 3, 7]
left = 0 (height=1), right = 8 (height=7)

Human thinking:
"Let's start with the widest container:
Area = (8-0) × min(1,7) = 8 × 1 = 8
Max so far: 8

Left line (1) is shorter than right line (7).
If I move the right line, width decreases but height stays 1.
But if I move the left line, I might find a taller line.
Let's move left to position 1:

left = 1 (height=8), right = 8 (height=7)
Area = (8-1) × min(8,7) = 7 × 7 = 49
Max so far: 49

Now right line (7) is shorter than left line (8).
Move right to position 7:

left = 1 (height=8), right = 7 (height=3)
Area = (7-1) × min(8,3) = 6 × 3 = 18
Max stays: 49

Right line (3) is shorter, move right to position 6:

left = 1 (height=8), right = 6 (height=8)
Area = (6-1) × min(8,8) = 5 × 8 = 40
Max stays: 49

Both lines have same height, move either. Let's move left:

left = 2 (height=6), right = 6 (height=8)
Area = (6-2) × min(6,8) = 4 × 6 = 24
Max stays: 49

Continue this process...
Final maximum area found: 49"
```

#### Phase 6: Intuition Validation
- **Why it works**: We never miss the optimal pair by eliminating shorter lines
- **Why it's efficient**: Each step eliminates one line from consideration
- **Why it's correct**: Greedy choice is provably optimal for this problem

### Common Human Pitfalls & How to Avoid Them
1. **"Should I move both pointers?"** → No, only move the shorter one
2. **"What about equal heights?"** → Can move either pointer
3. **"Should I check all pairs?"** → No, two-pointer eliminates need for O(N²)
4. **"Is local optimum global optimum?"** → Yes, due to problem structure

### Real-World Analogy
**Like finding the best fishing spot between two banks:**
- You have a river with varying bank heights
- You want to cast your fishing line between two points
- The water level is limited by the lower bank
- Distance between banks determines your fishing area
- Start with the widest span, then move inward looking for better spots
- Always move from the lower bank hoping to find higher ground

### Human-Readable Pseudocode
```
function maxContainer(heights):
    left = 0
    right = length(heights) - 1
    maxArea = 0
    
    while left < right:
        width = right - left
        currentHeight = min(heights[left], heights[right])
        currentArea = width * currentHeight
        maxArea = max(maxArea, currentArea)
        
        if heights[left] < heights[right]:
            left = left + 1
        else:
            right = right - 1
    
    return maxArea
```

### Execution Visualization

### Example: [1, 8, 6, 2, 5, 4, 8, 3, 7]
```
Initial: left=0(1), right=8(7), maxArea=0

Step 1: left=0(1), right=8(7)
→ width=8, height=min(1,7)=1, area=8
→ maxArea=8
→ Move left (1 < 7): left=1

Step 2: left=1(8), right=8(7)
→ width=7, height=min(8,7)=7, area=49
→ maxArea=49
→ Move right (8 > 7): right=7

Step 3: left=1(8), right=7(3)
→ width=6, height=min(8,3)=3, area=18
→ maxArea=49
→ Move right (8 > 3): right=6

Step 4: left=1(8), right=6(8)
→ width=5, height=min(8,8)=8, area=40
→ maxArea=49
→ Move either (equal): left=2

Step 5: left=2(6), right=6(8)
→ width=4, height=min(6,8)=6, area=24
→ maxArea=49
→ Move left (6 < 8): left=3

Continue until left >= right...
Final maxArea = 49
```

### Key Visualization Points:
- **Width decreases**: right-left gets smaller each step
- **Height optimization**: Always move from shorter line
- **Area tracking**: Keep maximum found so far
- **Termination**: Stop when pointers meet or cross

### Memory Layout Visualization:
```
Array: [1][8][6][2][5][4][8][3][7]
Index:  0  1  2  3  4  5  6  7  8
       ^                    ^
     left=0              right=8
       area=8×1=8

          ^              ^
        left=1          right=8  
          area=7×7=49 (max)
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each element visited at most once
- **Constant Space**: O(1) - only two pointers and variables
- **Greedy Optimization**: Each step eliminates one line permanently
- **No Additional Data Structures**: Pure pointer manipulation

### Alternative Approaches:

#### 1. Brute Force (O(N²))
```go
func maxArea(height []int) int {
    maxWater := 0
    for i := 0; i < len(height); i++ {
        for j := i + 1; j < len(height); j++ {
            width := j - i
            currentHeight := height[i]
            if height[j] < currentHeight {
                currentHeight = height[j]
            }
            currentArea := width * currentHeight
            if currentArea > maxWater {
                maxWater = currentArea
            }
        }
    }
    return maxWater
}
```
- **Pros**: Simple, checks all possibilities
- **Cons**: O(N²) time complexity

#### 2. Stack-Based Approach (Not applicable)
- Stack approach doesn't work well for this problem
- Two-pointer greedy approach is optimal

### Extensions for Interviews:
- **3D Container**: Find maximum volume with 3+ lines
- **Multiple Queries**: Process many queries efficiently
- **Dynamic Updates**: Handle height changes efficiently
- **Circular Array**: Container can wrap around
*/
func main() {
	// Test cases
	testCases := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
		{4, 3, 2, 1, 4},
		{1, 2, 1},
		{2, 3, 4, 5, 18, 17, 6},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 3, 2, 5, 25, 24, 5},
	}
	
	for i, height := range testCases {
		result := maxArea(height)
		fmt.Printf("Test Case %d: %v -> Max Area: %d\n", i+1, height, result)
	}
}
