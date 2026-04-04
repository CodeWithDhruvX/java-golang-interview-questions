package main

import "fmt"

// 62. Unique Paths
// Time: O(M*N), Space: O(M*N)
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	
	// dp[i][j] = number of ways to reach cell (i,j)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	
	// Initialize first row and first column
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	
	// Fill the dp table
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	
	return dp[m-1][n-1]
}

// Space optimized version: O(N) space
func uniquePathsOptimized(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	
	// Use 1D array representing current row
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	
	return dp[n-1]
}

// Mathematical approach using combinatorics
func uniquePathsMath(m int, n int) int {
	// Total steps = (m-1) + (n-1)
	// Choose (m-1) down steps or (n-1) right steps
	// C(total, m-1) or C(total, n-1)
	
	total := m + n - 2
	down := m - 1
	right := n - 1
	
	// Use min(down, right) to reduce computation
	if down > right {
		down, right = right, down
	}
	
	result := 1
	for i := 1; i <= down; i++ {
		result = result * (total - down + i) / i
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: 2D Dynamic Programming for Grid Paths
- **2D DP Table**: dp[i][j] represents number of ways to reach cell (i,j)
- **Grid Movement**: Only right and down moves allowed
- **Combinatorial Foundation**: Each path is a sequence of right/down moves
- **Base Cases**: First row and first column initialized to 1

## 2. PROBLEM CHARACTERISTICS
- **Grid Navigation**: Count paths from top-left to bottom-right
- **Movement Constraints**: Only right and down moves
- **Path Counting**: Count all possible unique paths
- **Grid Size**: M rows × N columns

## 3. SIMILAR PROBLEMS
- Unique Paths II (LeetCode 63) - Grid with obstacles
- Minimum Path Sum (LeetCode 64) - Path optimization with weights
- Dungeon Game (LeetCode 174) - Path with dynamic health
- Robot Path Planning - Various grid navigation problems

## 4. KEY OBSERVATIONS
- **Path Structure**: Each path consists of (M-1) down and (N-1) right moves
- **Combinatorial Nature**: Total moves = (M-1)+(N-1), choose positions for downs
- **DP Recurrence**: dp[i][j] = dp[i-1][j] + dp[i][j-1]
- **Base Cases**: First row and column have only one path each

## 5. VARIATIONS & EXTENSIONS
- **Grid Obstacles**: Some cells are blocked
- **Different Movement**: Allow diagonal or other directions
- **Path Constraints**: Maximum/minimum path length or sum
- **Multiple Start/End**: Different start and end positions

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are there obstacles? What are grid dimensions?"
- Edge cases: 1×1 grid, single row/column, large grids
- Time complexity: O(M×N) for DP, O(1) for combinatorial
- Space complexity: O(M×N) for DP, O(1) for combinatorial
- Consider integer overflow for large grids

## 7. COMMON MISTAKES
- Not handling 1×1 grid correctly
- Wrong base case initialization
- Integer overflow for large grids
- Using recursion without memoization (exponential)
- Off-by-one errors in DP table access

## 8. OPTIMIZATION STRATEGIES
- **2D DP**: O(M×N) time, O(M×N) space
- **Space Optimization**: O(N) space using rolling array
- **Combinatorial**: O(1) time using binomial coefficients
- **Mathematical Formula**: C(M+N-2, M-1) or C(M+N-2, N-1)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like counting routes in a city grid:**
- You're at the top-left corner of a city grid
- You can only move right (east) or down (south)
- You want to reach the bottom-right corner
- Each intersection has a certain number of ways to reach it
- The number grows as you move further from start

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Grid dimensions M×N
2. **Goal**: Count paths from (0,0) to (M-1,N-1)
3. **Constraints**: Only right and down moves
4. **Output**: Number of unique paths

#### Phase 2: Key Insight Recognition
- **"DP natural fit"** → Each cell depends on top and left neighbors
- **"Combinatorial approach"** → Paths are sequences of moves
- **"Base cases"** → First row/column have single path
- **"Mathematical formula"** → Choose positions for down moves

#### Phase 3: Strategy Development
```
Human thought process:
"I need to count all paths from start to end.
Each cell can be reached from top or left.
So paths to current cell = paths to top + paths to left.
First row and column are special - only one way to reach them.
I can fill the grid systematically using this recurrence."
```

#### Phase 4: Edge Case Handling
- **1×1 grid**: Only one path (start equals end)
- **Single row**: Only one path (all moves right)
- **Single column**: Only one path (all moves down)
- **Large grids**: Use combinatorial approach to avoid overflow

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
M = 3, N = 7 (3×7 grid)

Human thinking:
"I'll build a 3×7 DP table:
Initialize first row and column with 1s.

Row 0: [1, 1, 1, 1, 1, 1, 1, 1]
Row 1: [1, 2, 3, 4, 5, 6, 7, 8]
Row 2: [1, 3, 6, 10, 15, 21, 28, 36]

Let me trace some cells:
dp[1][1] = dp[0][1] + dp[1][0] = 1 + 1 = 2
dp[1][2] = dp[0][2] + dp[1][1] = 1 + 2 = 3
dp[2][3] = dp[1][3] + dp[2][2] = 4 + 6 = 10

Final result: dp[2][6] = 28 paths ✓"

Combinatorial check:
Total moves = (3-1) + (7-1) = 2 + 6 = 8
Choose 2 positions for down moves: C(8,2) = 28 ✓"
```

#### Phase 6: Intuition Validation
- **Why DP works**: Each cell's path count depends only on previous cells
- **Why combinatorial works**: Paths are unique sequences of moves
- **Why O(M×N)**: Each cell computed once
- **Why base cases**: First row/column have only one approach path

### Common Human Pitfalls & How to Avoid Them
1. **"Why not brute force?"** → Exponential paths, too slow
2. **"Should I use recursion?"** → Only with memoization
3. **"What about overflow?"** → Use 64-bit integers
4. **"Can I optimize further?"** → Combinatorial approach is optimal

### Real-World Analogy
**Like planning routes in a street grid:**
- You're at the northwest corner of a city
- You can only travel east or south
- You want to count all possible routes to southeast corner
- Each intersection can be reached from north or west
- The number of routes increases as you go further

### Human-Readable Pseudocode
```
function uniquePaths(m, n):
    if m <= 0 or n <= 0:
        return 0
    
    dp = create 2D array m×n
    for i from 0 to m-1:
        dp[i][0] = 1
    for j from 0 to n-1:
        dp[0][j] = 1
    
    for i from 1 to m-1:
        for j from 1 to n-1:
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
    
    return dp[m-1][n-1]

function uniquePathsMath(m, n):
    total = m + n - 2
    down = m - 1
    right = n - 1
    
    // Use min(down, right) to reduce computation
    if down > right:
        down, right = right, down
    
    result = 1
    for i from 1 to down:
        result = result * (total - down + i) / i
    
    return result
```

### Execution Visualization

### Example: M = 3, N = 3
```
DP Table Progression:
Initial: First row and column = 1

[1, 1, 1]
[1, 2, 3]
[1, 3, 6]

Cell-by-cell computation:
dp[1][1] = dp[0][1] + dp[1][0] = 1 + 1 = 2
dp[1][2] = dp[0][2] + dp[1][1] = 1 + 2 = 3
dp[2][1] = dp[1][1] + dp[2][0] = 2 + 1 = 3
dp[2][2] = dp[1][2] + dp[2][1] = 3 + 3 = 6

Final result: dp[2][2] = 6 paths ✓
```

### Key Visualization Points:
- **DP Recurrence**: Each cell = top + left
- **Base Cases**: First row/column = 1
- **Path Growth**: Numbers increase diagonally
- **Final Result**: Bottom-right corner

### Memory Layout Visualization:
```
Path Counting Visualization:
3×3 Grid with paths to each cell:

Start → [1] → [1] → [1]
  ↓       ↓       ↓
 [1] → [2] → [3]
  ↓       ↓       ↓
 [1] → [3] → [6] ← End

Each number shows total paths to that cell.
```

### Time Complexity Breakdown:
- **DP Approach**: O(M×N) time, O(M×N) space
- **Space Optimized**: O(M×N) time, O(N) space
- **Combinatorial**: O(1) time, O(1) space
- **Recursive without Memo**: O(2^(M+N)) time, exponential

### Alternative Approaches:

#### 1. Recursive with Memoization (O(M×N) time, O(M×N) space)
```go
func uniquePathsRecursive(m, n int) int {
    memo := make(map[string]int)
    return uniquePathsHelper(0, 0, m-1, n-1, memo)
}

func uniquePathsHelper(row, col, targetRow, targetCol int, memo map[string]int) int {
    if row == targetRow && col == targetCol {
        return 1
    }
    if row > targetRow || col > targetCol {
        return 0
    }
    
    key := fmt.Sprintf("%d,%d", row, col)
    if val, exists := memo[key]; exists {
        return val
    }
    
    paths := uniquePathsHelper(row+1, col, targetRow, targetCol, memo) +
            uniquePathsHelper(row, col+1, targetRow, targetCol, memo)
    
    memo[key] = paths
    return paths
}
```
- **Pros**: Intuitive, follows problem description
- **Cons**: Recursion overhead, same complexity as DP

#### 2. Iterative with Space Optimization (O(M×N) time, O(N) space)
```go
func uniquePathsOptimized(m, n int) int {
    if m <= 0 || n <= 0 {
        return 0
    }
    
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[j] = dp[j] + dp[j-1]
        }
    }
    
    return dp[n-1]
}
```
- **Pros**: Reduced space usage
- **Cons**: Slightly less intuitive

#### 3. Pure Combinatorial (O(1) time, O(1) space)
```go
func uniquePathsCombinatorial(m, n int) int {
    total := m + n - 2
    down := m - 1
    right := n - 1
    
    // C(total, down) = total! / (down! * right!)
    result := 1
    for i := 1; i <= down; i++ {
        result = result * (total - down + i) / i
    }
    
    return result
}
```
- **Pros**: Optimal time and space
- **Cons**: Requires careful overflow handling

### Extensions for Interviews:
- **Grid with Obstacles**: Some cells are blocked
- **Different Movement**: Allow diagonal or other directions
- **Path Constraints**: Maximum/minimum path length
- **Multiple Queries**: Answer many path count queries
- **Modular Arithmetic**: Return result modulo large prime
*/
func main() {
	// Test cases
	testCases := []struct {
		m int
		n int
	}{
		{3, 7},
		{3, 2},
		{1, 1},
		{1, 5},
		{5, 1},
		{2, 2},
		{10, 10},
		{7, 3},
		{3, 3},
		{100, 1},
		{1, 100},
	}
	
	for i, tc := range testCases {
		result1 := uniquePaths(tc.m, tc.n)
		result2 := uniquePathsOptimized(tc.m, tc.n)
		result3 := uniquePathsMath(tc.m, tc.n)
		
		fmt.Printf("Test Case %d: m=%d, n=%d -> DP: %d, Optimized: %d, Math: %d\n", 
			i+1, tc.m, tc.n, result1, result2, result3)
	}
}
