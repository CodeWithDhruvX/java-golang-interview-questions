import java.util.*;

public class SpiralMatrix {
    
    // 54. Spiral Matrix
    // Time: O(M*N), Space: O(1) (excluding output)
    public static List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> result = new ArrayList<>();
        if (matrix.length == 0 || matrix[0].length == 0) {
            return result;
        }
        
        int m = matrix.length, n = matrix[0].length;
        int top = 0, bottom = m - 1;
        int left = 0, right = n - 1;
        
        while (top <= bottom && left <= right) {
            // Traverse from left to right (top row)
            for (int col = left; col <= right; col++) {
                result.add(matrix[top][col]);
            }
            top++;
            
            // Traverse from top to bottom (right column)
            for (int row = top; row <= bottom; row++) {
                result.add(matrix[row][right]);
            }
            right--;
            
            // Traverse from right to left (bottom row) if still valid
            if (top <= bottom) {
                for (int col = right; col >= left; col--) {
                    result.add(matrix[bottom][col]);
                }
                bottom--;
            }
            
            // Traverse from bottom to top (left column) if still valid
            if (left <= right) {
                for (int row = bottom; row >= top; row--) {
                    result.add(matrix[row][left]);
                }
                left++;
            }
        }
        
        return result;
    }

    // Alternative approach using direction vectors
    public static List<Integer> spiralOrderDirections(int[][] matrix) {
        List<Integer> result = new ArrayList<>();
        if (matrix.length == 0 || matrix[0].length == 0) {
            return result;
        }
        
        int m = matrix.length, n = matrix[0].length;
        
        // Directions: right, down, left, up
        int[][] directions = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};
        int dirIndex = 0;
        
        // Boundaries
        int top = 0, bottom = m - 1;
        int left = 0, right = n - 1;
        
        int row = 0, col = 0;
        
        while (result.size() < m * n) {
            result.add(matrix[row][col]);
            
            // Calculate next position
            int nextRow = row + directions[dirIndex][0];
            int nextCol = col + directions[dirIndex][1];
            
            // Check if we need to change direction
            if (nextRow < top || nextRow > bottom || nextCol < left || nextCol > right) {
                // Update boundaries and change direction
                switch (dirIndex) {
                    case 0: // Moving right
                        top++;
                        break;
                    case 1: // Moving down
                        right--;
                        break;
                    case 2: // Moving left
                        bottom--;
                        break;
                    case 3: // Moving up
                        left++;
                        break;
                }
                
                dirIndex = (dirIndex + 1) % 4;
                nextRow = row + directions[dirIndex][0];
                nextCol = col + directions[dirIndex][1];
            }
            
            row = nextRow;
            col = nextCol;
        }
        
        return result;
    }

    // Recursive approach
    public static List<Integer> spiralOrderRecursive(int[][] matrix) {
        List<Integer> result = new ArrayList<>();
        if (matrix.length == 0 || matrix[0].length == 0) {
            return result;
        }
        
        spiralHelper(matrix, 0, matrix.length - 1, 0, matrix[0].length - 1, result);
        return result;
    }

    private static void spiralHelper(int[][] matrix, int top, int bottom, int left, int right, List<Integer> result) {
        if (top > bottom || left > right) {
            return;
        }
        
        // Traverse from left to right (top row)
        for (int col = left; col <= right; col++) {
            result.add(matrix[top][col]);
        }
        
        // Traverse from top to bottom (right column)
        for (int row = top + 1; row <= bottom; row++) {
            result.add(matrix[row][right]);
        }
        
        // Traverse from right to left (bottom row) if still valid
        if (top < bottom) {
            for (int col = right - 1; col >= left; col--) {
                result.add(matrix[bottom][col]);
            }
        }
        
        // Traverse from bottom to top (left column) if still valid
        if (left < right) {
            for (int row = bottom - 1; row > top; row--) {
                result.add(matrix[row][left]);
            }
        }
        
        // Recursively process inner matrix
        spiralHelper(matrix, top + 1, bottom - 1, left + 1, right - 1, result);
    }

    public static void main(String[] args) {
        // Test cases
        int[][][] testCases = {
            {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
            {{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
            {{1, 2, 3}, {4, 5, 6}},
            {{1, 2}, {3, 4}, {5, 6}, {7, 8}},
            {{1}},
            {{1, 2, 3, 4, 5}},
            {{1}, {2}, {3}, {4}, {5}},
            {{1, 2}, {3, 4}},
            {{1, 2, 3}, {4, 5}, {6, 7, 8}},
            {}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.deepToString(testCases[i]));
            
            List<Integer> result1 = spiralOrder(testCases[i]);
            List<Integer> result2 = spiralOrderDirections(testCases[i]);
            List<Integer> result3 = spiralOrderRecursive(testCases[i]);
            
            System.out.printf("  Standard: %s\n", result1);
            System.out.printf("  Directions: %s\n", result2);
            System.out.printf("  Recursive: %s\n\n", result3);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Matrix Traversal
- **Spiral Traversal**: Layer-by-layer matrix traversal
- **Boundary Management**: Track current matrix boundaries
- **Direction Cycling**: Right → Down → Left → Up pattern
- **Layer Processing**: Process outer layer, then move inward

## 2. PROBLEM CHARACTERISTICS
- **2D Matrix**: Rectangular array of integers
- **Spiral Order**: Clockwise traversal from outer to inner
- **Boundary Shrink**: Boundaries move inward after each layer
- **Complete Coverage**: Visit all elements exactly once

## 3. SIMILAR PROBLEMS
- Rotate Matrix
- Zigzag Traversal
- Diagonal Traversal
- Matrix Search in Sorted Order

## 4. KEY OBSERVATIONS
- Four directions: right, down, left, up (in that order)
- Boundaries shrink after each complete cycle
- Each layer processes perimeter then moves inward
- Time complexity: O(M×N) for M×N matrix
- Space complexity: O(1) for output storage

## 5. VARIATIONS & EXTENSIONS
- Counter-clockwise spiral
- Spiral from center outward
- Different starting positions
- Non-square matrices

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I handle empty matrix?"
- Edge cases: single row, single column, empty matrix
- Time complexity: O(M×N) vs O(M×N log M×N) recursive
- Space complexity: O(1) vs O(M×N) for recursive

## 7. COMMON MISTAKES
- Incorrect boundary updates
- Off-by-one errors in loops
- Not handling all four directions
- Forgetting to shrink boundaries properly
- Infinite loops in boundary conditions

## 8. OPTIMIZATION STRATEGIES
- Iterative approach avoids recursion stack
- Direction vector approach simplifies logic
- Early termination for small matrices
- In-place boundary updates

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like peeling an onion:**
- You have a matrix (onion) with multiple layers
- Start from outer layer and peel it off completely
- Move to the next inner layer and repeat
- Each layer is peeled in four directions (right, down, left, up)
- Continue until you reach the center
- The peeled layers give you the spiral order

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: 2D matrix of integers
2. **Goal**: Return elements in clockwise spiral order
3. **Output**: List of integers in spiral traversal

#### Phase 2: Key Insight Recognition
- **"How to traverse spiral?"** → Process layers from outside to inside
- **"What are the directions?"** → Right → Down → Left → Up → repeat
- **"How to track boundaries?"** → top, bottom, left, right variables
- **"When to move inward?"** → After completing one full cycle

#### Phase 3: Strategy Development
```
Human thought process:
"I'll traverse in layers:
1. Initialize boundaries: top=0, bottom=m-1, left=0, right=n-1
2. While boundaries are valid:
   - Traverse top row left→right
   - Traverse right column top→bottom
   - Traverse bottom row right→left
   - Traverse left column bottom→top
   - Shrink boundaries inward
3. This processes each layer completely"
```

#### Phase 4: Edge Case Handling
- **Empty matrix**: Return empty list
- **Single row**: Just traverse left→right
- **Single column**: Just traverse top→bottom
- **Odd dimensions**: Handle center element properly

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Matrix:
1 2 3
4 5 6
7 8 9

Human thinking:
"Let's traverse in layers:

Initial boundaries: top=0, bottom=2, left=0, right=2

Layer 1 (outer):
- Top row: 1→2→3
- Right column: 6→9
- Bottom row: 8→7
- Left column: 4→5
- Shrink: top=1, bottom=1, left=1, right=1

Layer 2 (inner):
- Top row: 5
- Right column: 6
- Bottom row: 8
- Left column: 4
- Shrink: top=2, bottom=0, left=2, right=0

Boundaries crossed, stop ✓

Result: [1,2,3,6,9,8,7,4,5] ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Layer-by-layer ensures complete coverage
- **Why it's efficient**: Each element visited exactly once
- **Why it's correct**: Boundary management prevents missing elements

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use recursion?"** → Stack overhead and complexity
2. **"What about direction vectors?"** → Simplifies boundary updates
3. **"How to handle odd dimensions?"** → Center element handled separately
4. **"What about boundaries?"** → Must shrink after each layer

### Real-World Analogy
**Like reading a book page by page:**
- You have a book page (matrix) with text
- You read it in spiral pattern: left to right, top to bottom, etc.
- Each complete circuit around the page is one layer
- Then you move to the inner text and continue
- This is exactly how you read a spiral-bound document

### Human-Readable Pseudocode
```
function spiralOrder(matrix):
    if matrix is empty:
        return []
    
    result = []
    top = 0, bottom = m-1, left = 0, right = n-1
    
    while top <= bottom and left <= right:
        // Traverse top row
        for col from left to right:
            result.add(matrix[top][col])
        top++
        
        // Traverse right column
        for row from top to bottom:
            result.add(matrix[row][right])
        right--
        
        // Traverse bottom row
        if top <= bottom:
            for col from right to left:
                result.add(matrix[bottom][col])
            bottom--
        
        // Traverse left column
        if left <= right:
            for row from bottom to top:
                result.add(matrix[row][left])
            left++
    
    return result
```

### Execution Visualization

### Example: 3×3 matrix
```
Matrix:
1 2 3
4 5 6
7 8 9

Spiral Traversal:

Layer 1 (outer):
→ 1→2→3 (top row)
→ 6→9 (right column)
→ 8→7 (bottom row)
→ 4→5 (left column)

Layer 2 (center):
→ 5 (center element)

Result: [1,2,3,6,9,8,7,4,5] ✓

Visualization:
┌───┐
│1 2 3│
│4   5 6│
│7   8 9│
└───┘

Spiral path: 1→2→3→6→9→8→7→4→5
```

### Key Visualization Points:
- **Layer processing** ensures complete coverage
- **Boundary management** tracks current traversal limits
- **Direction cycling** follows right→down→left→up pattern
- **Inward movement** after each complete layer

### Memory Layout Visualization:
```
Boundary Evolution:
Initial: top=0, bottom=2, left=0, right=2
After Layer 1: top=1, bottom=1, left=1, right=1
After Layer 2: top=2, bottom=0, left=2, right=0

Traversal Pattern:
Right→Down→Left→Up→Right→Down→Left→Up→...

Element Access Order:
[0,0] [0,1] [0,2] [1,2] [2,2] [2,1] [1,1] [1,0] [0,0]
```

### Time Complexity Breakdown:
- **Each Element**: Visited exactly once
- **Boundary Operations**: O(1) per element
- **Total**: O(M×N) time where M×N is matrix size
- **Space**: O(1) for output (excluding result storage)
- **Optimal**: Cannot do better than O(M×N) for this problem
- **vs Recursive**: O(M×N) time but O(M×N) space for call stack
*/
