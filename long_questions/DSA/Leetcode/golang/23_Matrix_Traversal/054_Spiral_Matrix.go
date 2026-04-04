package main

import "fmt"

// 54. Spiral Matrix
// Time: O(M*N), Space: O(1) (excluding output)
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	
	top, bottom := 0, m-1
	left, right := 0, n-1
	
	for top <= bottom && left <= right {
		// Traverse from left to right (top row)
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++
		
		// Traverse from top to bottom (right column)
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--
		
		// Traverse from right to left (bottom row) if still valid
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom--
		}
		
		// Traverse from bottom to top (left column) if still valid
		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}
	
	return result
}

// Alternative approach using direction vectors
func spiralOrderDirections(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	
	// Directions: right, down, left, up
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIndex := 0
	
	// Boundaries
	top, bottom := 0, m-1
	left, right := 0, n-1
	
	row, col := 0, 0
	
	for len(result) < m*n {
		result = append(result, matrix[row][col])
		
		// Calculate next position
		nextRow := row + directions[dirIndex][0]
		nextCol := col + directions[dirIndex][1]
		
		// Check if we need to change direction
		if nextRow < top || nextRow > bottom || nextCol < left || nextCol > right {
			// Update boundaries and change direction
			switch dirIndex {
			case 0: // Moving right
				top++
			case 1: // Moving down
				right--
			case 2: // Moving left
				bottom--
			case 3: // Moving up
				left++
			}
			
			dirIndex = (dirIndex + 1) % 4
			nextRow = row + directions[dirIndex][0]
			nextCol = col + directions[dirIndex][1]
		}
		
		row, col = nextRow, nextCol
	}
	
	return result
}

// Recursive approach
func spiralOrderRecursive(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	
	var result []int
	spiralHelper(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, &result)
	return result
}

func spiralHelper(matrix [][]int, top, bottom, left, right int, result *[]int) {
	if top > bottom || left > right {
		return
	}
	
	// Traverse from left to right (top row)
	for col := left; col <= right; col++ {
		*result = append(*result, matrix[top][col])
	}
	
	// Traverse from top to bottom (right column)
	for row := top + 1; row <= bottom; row++ {
		*result = append(*result, matrix[row][right])
	}
	
	// Traverse from right to left (bottom row) if still valid
	if top < bottom {
		for col := right - 1; col >= left; col-- {
			*result = append(*result, matrix[bottom][col])
		}
	}
	
	// Traverse from bottom to top (left column) if still valid
	if left < right {
		for row := bottom - 1; row > top; row-- {
			*result = append(*result, matrix[row][left])
		}
	}
	
	// Recursively process inner matrix
	spiralHelper(matrix, top+1, bottom-1, left+1, right-1, result)
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Spiral Matrix Traversal
- **Boundary Shrinking**: Process matrix in concentric layers
- **Direction Cycling**: Right → Down → Left → Up → repeat
- **Layer Processing**: Process outer layer, then inner layer
- **Boundary Updates**: Shrink boundaries after each side

## 2. PROBLEM CHARACTERISTICS
- **2D Matrix**: Rectangular grid of elements
- **Spiral Order**: Clockwise traversal from outer to inner
- **Boundary Management**: Track current layer boundaries
- **Complete Coverage**: Visit every element exactly once

## 3. SIMILAR PROBLEMS
- Rotate Image (LeetCode 48) - Matrix rotation
- Set Matrix Zeroes (LeetCode 73) - Mark rows/columns
- Word Search (LeetCode 79) - DFS in matrix
- Matrix Diagonal Traversal - Diagonal pattern problems

## 4. KEY OBSERVATIONS
- **Layer Structure**: Matrix consists of concentric rectangular layers
- **Direction Pattern**: Four directions repeat consistently
- **Boundary Convergence**: Top/bottom and left/right converge to center
- **Termination Condition**: Boundaries cross or meet

## 5. VARIATIONS & EXTENSIONS
- **Counter-clockwise**: Different direction order
- **Spiral Inward**: Start from center, spiral outward
- **Different Shapes**: Non-rectangular matrices
- **Custom Patterns**: Zigzag, diagonal patterns

## 6. INTERVIEW INSIGHTS
- Always clarify: "Matrix dimensions? Empty matrix? Single element?"
- Edge cases: empty matrix, single row/column, odd/even dimensions
- Time complexity: O(M*N) where M=rows, N=columns
- Space complexity: O(1) extra space (excluding output)
- Key insight: process in layers, shrink boundaries

## 7. COMMON MISTAKES
- Wrong boundary updates (off-by-one errors)
- Not handling single row/column cases
- Incorrect direction cycling
- Missing termination condition
- Not handling empty matrix properly

## 8. OPTIMIZATION STRATEGIES
- **Layer Processing**: O(M*N) time, O(1) space - optimal
- **Direction Vectors**: O(M*N) time, O(1) space - cleaner code
- **Recursive**: O(M*N) time, O(L) space where L=number of layers
- **In-place**: O(M*N) time, O(1) space - modifies input

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like peeling an onion layer by layer:**
- You have a rectangular onion (matrix)
- Peel the outer layer in clockwise direction
- Then peel the next inner layer
- Continue until you reach the center
- Like unwrapping a present in spiral pattern

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: 2D matrix of elements
2. **Goal**: Return elements in clockwise spiral order
3. **Pattern**: Right → Down → Left → Up, then repeat
4. **Output**: Array of elements in spiral order

#### Phase 2: Key Insight Recognition
- **"Layer structure"** → Matrix has concentric rectangular layers
- **"Boundary shrinking"** → Each layer reduces boundaries
- **"Direction cycling"** → Four consistent directions
- **"Termination"** → Boundaries converge to center

#### Phase 3: Strategy Development
```
Human thought process:
"I need to traverse matrix in spiral order.
I can process it layer by layer:

1. Start with outer boundaries: top=0, bottom=m-1, left=0, right=n-1
2. Traverse top row left→right, then shrink top boundary
3. Traverse right column top→bottom, then shrink right boundary
4. Traverse bottom row right→left, then shrink bottom boundary
5. Traverse left column bottom→top, then shrink left boundary
6. Repeat until boundaries cross

This naturally creates the spiral pattern!"
```

#### Phase 4: Edge Case Handling
- **Empty matrix**: Return empty array
- **Single element**: Return that element
- **Single row**: Traverse left→right only
- **Single column**: Traverse top→bottom only

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Matrix: [[1,2,3], [4,5,6], [7,8,9]]

Human thinking:
"Initial boundaries: top=0, bottom=2, left=0, right=2

Layer 1 (outer):
1. Top row: [1,2,3] (left→right), top=1
2. Right column: [6,9] (top→bottom), right=1
3. Bottom row: [8,7] (right→left), bottom=1
4. Left column: [4] (bottom→top), left=1

Layer 2 (inner):
Boundaries now: top=1, bottom=1, left=1, right=1
Only element [5] remains

Final result: [1,2,3,6,9,8,7,4,5] ✓"
```

#### Phase 6: Intuition Validation
- **Why layering works**: Each layer is independent rectangle
- **Why boundary shrinking works**: Reduces problem size systematically
- **Why direction order works**: Creates clockwise spiral
- **Why O(M*N)**: Each element visited exactly once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use indices?"** → Need boundary management for spiral
2. **"Should I use recursion?"** → Iterative is simpler, recursion adds overhead
3. **"What about different spiral directions?"** → Just change direction order
4. **"Can I optimize further?"** → O(M*N) is already optimal
5. **"What about non-square matrices?"** → Boundaries handle any rectangle

### Real-World Analogy
**Like reading a spiral-bound notebook:**
- You have pages with text in spiral pattern
- Read the outer ring clockwise, then move to inner ring
- Continue until you reach the center
- Each ring is independent, boundaries shrink naturally
- Like following a path that spirals inward

### Human-Readable Pseudocode
```
function spiralOrder(matrix):
    if matrix is empty:
        return []
    
    m, n = matrix dimensions
    result = []
    top, bottom = 0, m-1
    left, right = 0, n-1
    
    while top <= bottom && left <= right:
        // Traverse top row left→right
        for col from left to right:
            result.append(matrix[top][col])
        top++
        
        // Traverse right column top→bottom
        for row from top to bottom:
            result.append(matrix[row][right])
        right--
        
        // Traverse bottom row right→left (if still valid)
        if top <= bottom:
            for col from right down to left:
                result.append(matrix[bottom][col])
            bottom--
        
        // Traverse left column bottom→top (if still valid)
        if left <= right:
            for row from bottom down to top:
                result.append(matrix[row][left])
            left++
    
    return result
```

### Execution Visualization

### Example: Matrix = [[1,2,3], [4,5,6], [7,8,9]]
```
Initial State:
Matrix:
1 2 3
4 5 6
7 8 9

Boundaries: top=0, bottom=2, left=0, right=2
Result: []

Layer 1 Processing:
Step 1 - Top row (left→right): [1,2,3]
Result: [1,2,3], top=1

Step 2 - Right column (top→bottom): [6,9]
Result: [1,2,3,6,9], right=1

Step 3 - Bottom row (right→left): [8,7]
Result: [1,2,3,6,9,8,7], bottom=1

Step 4 - Left column (bottom→top): [4]
Result: [1,2,3,6,9,8,7,4], left=1

Layer 2 Processing:
Boundaries: top=1, bottom=1, left=1, right=1
Single element: [5]

Step 1 - Top row: [5]
Result: [1,2,3,6,9,8,7,4,5], top=2

Termination: top=2 > bottom=1, left=2 > right=1 ✓
Final result: [1,2,3,6,9,8,7,4,5] ✓
```

### Key Visualization Points:
- **Boundary Management**: Four boundaries shrink systematically
- **Direction Order**: Right → Down → Left → Up
- **Layer Independence**: Each layer processed independently
- **Termination**: Boundaries converge to center

### Memory Layout Visualization:
```
Matrix State During Spiral:
1 2 3
4 5 6
7 8 9

Boundary Evolution:
Initial: top=0, bottom=2, left=0, right=2
After top row: top=1, bottom=2, left=0, right=2
After right col: top=1, bottom=2, left=0, right=1
After bottom row: top=1, bottom=1, left=0, right=1
After left col: top=1, bottom=1, left=1, right=1

Final: top=2, bottom=1, left=2, right=1 (crossed)
Result built: [1,2,3,6,9,8,7,4,5] ✓
```

### Time Complexity Breakdown:
- **Element Access**: O(M*N) time (each element visited once)
- **Boundary Updates**: O(1) per element
- **Space**: O(1) extra space (excluding output array)
- **Total**: O(M*N) time, O(1) space

### Alternative Approaches:

#### 1. Direction Vector Approach (O(M*N) time, O(1) space)
```go
func spiralOrderDirections(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []
    }
    
    m, n := len(matrix), len(matrix[0])
    result := []int{}
    
    // Direction vectors: right, down, left, up
    directions := [][2]int{{0,1}, {1,0}, {0,-1}, {-1,0}}
    dirIndex := 0
    
    // Boundaries
    top, bottom := 0, m-1
    left, right := 0, n-1
    
    row, col := 0, 0
    
    for len(result) < m*n {
        result = append(result, matrix[row][col])
        
        // Calculate next position
        nextRow := row + directions[dirIndex][0]
        nextCol := col + directions[dirIndex][1]
        
        // Check if we need to change direction
        if nextRow < top || nextRow > bottom || nextCol < left || nextCol > right {
            // Update boundaries and change direction
            switch dirIndex {
            case 0: top++      // Moving right
            case 1: right--     // Moving down
            case 2: bottom--    // Moving left
            case 3: left++      // Moving up
            }
            dirIndex = (dirIndex + 1) % 4
            nextRow = row + directions[dirIndex][0]
            nextCol = col + directions[dirIndex][1]
        }
        
        row, col = nextRow, nextCol
    }
    
    return result
}
```
- **Pros**: Cleaner code, no explicit loops per direction
- **Cons**: More complex boundary checking logic

#### 2. Recursive Layer Approach (O(M*N) time, O(L) space)
```go
func spiralOrderRecursive(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []
    }
    
    var result []int
    spiralHelper(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, &result)
    return result
}

func spiralHelper(matrix [][]int, top, bottom, left, right int, result *[]int) {
    if top > bottom || left > right {
        return
    }
    
    // Traverse top row
    for col := left; col <= right; col++ {
        *result = append(*result, matrix[top][col])
    }
    
    // Traverse right column
    for row := top + 1; row <= bottom; row++ {
        *result = append(*result, matrix[row][right])
    }
    
    // Traverse bottom row (if different from top)
    if top < bottom {
        for col := right - 1; col >= left; col-- {
            *result = append(*result, matrix[bottom][col])
        }
    }
    
    // Traverse left column (if different from right)
    if left < right {
        for row := bottom - 1; row > top; row-- {
            *result = append(*result, matrix[row][left])
        }
    }
    
    // Recursively process inner layer
    spiralHelper(matrix, top+1, bottom-1, left+1, right-1, result)
}
```
- **Pros**: Natural recursive formulation
- **Cons**: O(L) space for recursion stack

#### 3. In-place Modification (O(M*N) time, O(1) space)
```go
func spiralOrderInPlace(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []
    }
    
    result := make([]int, 0, len(matrix)*len(matrix[0]))
    // ... modify matrix in-place to mark visited
    // ... implementation details omitted
    return result
}
```
- **Pros**: O(1) extra space
- **Cons**: Modifies input matrix, more complex

### Extensions for Interviews:
- **Counter-clockwise Spiral**: Different direction order
- **Spiral Outward**: Start from center, spiral outward
- **Different Patterns**: Zigzag, diagonal traversals
- **Custom Shapes**: Non-rectangular matrices
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := [][][]int{
		{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
		{{1, 2, 3}, {4, 5, 6}},
		{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
		{{1}},
		{{1, 2, 3, 4, 5}},
		{{1}, {2}, {3}, {4}, {5}},
		{{1, 2}, {3, 4}},
		{{1, 2, 3}, {4, 5}, {6, 7, 8}},
		{{}},
	}
	
	for i, matrix := range testCases {
		fmt.Printf("Test Case %d: %v\n", i+1, matrix)
		
		result1 := spiralOrder(matrix)
		result2 := spiralOrderDirections(matrix)
		result3 := spiralOrderRecursive(matrix)
		
		fmt.Printf("  Standard: %v\n", result1)
		fmt.Printf("  Directions: %v\n", result2)
		fmt.Printf("  Recursive: %v\n\n", result3)
	}
}
