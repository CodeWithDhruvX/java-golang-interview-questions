package main

import "fmt"

// 329. Longest Increasing Path in a Matrix
// Time: O(M*N), Space: O(M*N) for memoization
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	
	m, n := len(matrix), len(matrix[0])
	
	// Memoization cache
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 indicates not computed
		}
	}
	
	maxLength := 0
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	var dfs func(int, int) int
	dfs = func(row, col int) int {
		if memo[row][col] != -1 {
			return memo[row][col]
		}
		
		maxPath := 1 // Current cell itself
		
		// Explore all 4 directions
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			
			// Check boundaries and increasing condition
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n &&
				matrix[newRow][newCol] > matrix[row][col] {
				
				pathLength := 1 + dfs(newRow, newCol)
				if pathLength > maxPath {
					maxPath = pathLength
				}
			}
		}
		
		memo[row][col] = maxPath
		return maxPath
	}
	
	// Find longest path starting from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			currentLength := dfs(i, j)
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}
	
	return maxLength
}

// Topological sort approach
func longestIncreasingPathTopological(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	
	m, n := len(matrix), len(matrix[0])
	
	// Build graph and calculate in-degree
	adj := make(map[string][]string)
	inDegree := make(map[string]int)
	
	// Create nodes for each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			node := fmt.Sprintf("%d,%d", i, j)
			inDegree[node] = 0
		}
	}
	
	// Build edges from smaller to larger values
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			current := fmt.Sprintf("%d,%d", i, j)
			
			for _, dir := range directions {
				newRow, newCol := i+dir[0], j+dir[1]
				
				if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n &&
					matrix[newRow][newCol] > matrix[i][j] {
					
					neighbor := fmt.Sprintf("%d,%d", newRow, newCol)
					adj[current] = append(adj[current], neighbor)
					inDegree[neighbor]++
				}
			}
		}
	}
	
	// Topological sort using BFS (Kahn's algorithm)
	queue := []string{}
	
	// Find nodes with no incoming edges
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}
	
	maxLength := 0
	
	for len(queue) > 0 {
		levelSize := len(queue)
		
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			
			// Process neighbors
			for _, neighbor := range adj[current] {
				inDegree[neighbor]--
				if inDegree[neighbor] == 0 {
					queue = append(queue, neighbor)
				}
			}
		}
		
		maxLength++
	}
	
	return maxLength
}

// BFS approach
func longestIncreasingPathBFS(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	
	m, n := len(matrix), len(matrix[0])
	
	// DP table to store longest path from each cell
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	// Process cells in order of their values
	cells := make([][2]int, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cells = append(cells, [2]int{i, j})
		}
	}
	
	// Sort cells by their values (simple bubble sort for demonstration)
	for i := 0; i < len(cells)-1; i++ {
		for j := 0; j < len(cells)-i-1; j++ {
			if matrix[cells[j][0]][cells[j][1]] > matrix[cells[j+1][0]][cells[j+1][1]] {
				cells[j], cells[j+1] = cells[j+1], cells[j]
			}
		}
	}
	
	maxLength := 1
	
	for _, cell := range cells {
		i, j := cell[0], cell[1]
		dp[i][j] = 1 // At least the cell itself
		
		// Check all neighbors
		for _, dir := range directions {
			prevRow, prevCol := i+dir[0], j+dir[1]
			
			if prevRow >= 0 && prevRow < m && prevCol >= 0 && prevCol < n &&
				matrix[prevRow][prevCol] < matrix[i][j] {
				
				if dp[prevRow][prevCol]+1 > dp[i][j] {
					dp[i][j] = dp[prevRow][prevCol] + 1
				}
			}
		}
		
		if dp[i][j] > maxLength {
			maxLength = dp[i][j]
		}
	}
	
	return maxLength
}

// Helper function to create board from strings
func createMatrix(matrixStr []string) [][]int {
	matrix := make([][]int, len(matrixStr))
	for i, row := range matrixStr {
		matrix[i] = make([]int, len(row))
		for j, char := range row {
			matrix[i][j] = int(char - '0') // Assuming single digits
		}
	}
	return matrix
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Longest Increasing Path in Matrix
- **DFS with Memoization**: Depth-first search with caching to avoid recomputation
- **Topological Sort**: Build DAG of increasing edges and find longest path
- **Dynamic Programming**: Process cells in order to build longest path
- **Graph Construction**: Convert matrix to directed graph of increasing edges

## 2. PROBLEM CHARACTERISTICS
- **Matrix Graph**: 2D grid where each cell is a node
- **Increasing Constraint**: Can only move to cells with larger values
- **Path Finding**: Find longest path following increasing rule
- **4-Direction Movement**: Can move horizontally/vertically to adjacent cells

## 3. SIMILAR PROBLEMS
- Longest Increasing Subsequence (LeetCode 300) - 1D version
- Number of Longest Increasing Subsequence (LeetCode 673) - Count paths
- Longest Path in DAG - General graph version
- Matrix Path Problems - Various matrix traversal patterns

## 4. KEY OBSERVATIONS
- **DAG Structure**: Matrix with increasing constraint forms a DAG
- **Multiple Solutions**: DFS with memoization, topological sort, DP all work
- **Memoization Critical**: Without caching, exponential complexity
- **Processing Order**: Topological sort enables DP approach

## 5. VARIATIONS & EXTENSIONS
- **Diagonal Movement**: Allow 8-directional movement
- **Equal Values**: Handle cells with same values
- **Path Counting**: Count number of longest paths
- **Multiple Queries**: Answer multiple path length queries

## 6. INTERVIEW INSIGHTS
- Always clarify: "Movement directions? Equal values? Matrix size?"
- Edge cases: empty matrix, single cell, all equal values
- Time complexity: O(M*N) with memoization, O(M*N*log(M*N)) with topological sort
- Space complexity: O(M*N) for memoization/DP
- Key insight: matrix with increasing constraint is a DAG

## 7. COMMON MISTAKES
- Not using memoization (exponential time)
- Wrong boundary checking in DFS
- Not handling equal values correctly
- Incorrect graph construction for topological sort
- Missing base cases in recursion

## 8. OPTIMIZATION STRATEGIES
- **DFS with Memoization**: O(M*N) time, O(M*N) space
- **Topological Sort**: O(M*N*log(M*N)) time, O(M*N) space
- **Dynamic Programming**: O(M*N) time, O(M*N) space
- **Early Pruning**: Stop when no increasing neighbors

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the longest hiking trail with elevation gain:**
- You have a grid of elevations (matrix values)
- Can only move to higher elevations
- Want to find the longest possible trail
- Like a mountain climber seeking the longest ascent route
- Each step must go to a higher point

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: 2D matrix of integers
2. **Goal**: Find length of longest strictly increasing path
3. **Constraints**: Move only to adjacent cells with larger values
4. **Output**: Maximum path length

#### Phase 2: Key Insight Recognition
- **"DAG natural fit"** → Increasing constraint creates directed acyclic graph
- **"Memoization essential"** → Without caching, exponential recomputation
- **"Multiple approaches"** → DFS, topological sort, DP all work
- **"Processing order"** → Topological sort enables DP solution

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find longest increasing path in matrix.
This is like finding longest path in a DAG:

DFS with Memoization Approach:
1. For each cell, DFS to find longest path starting there
2. Cache results to avoid recomputation
3. Return max of all cached results

Topological Sort Approach:
1. Build graph where edges go from smaller to larger values
2. Topologically sort nodes (process in increasing order)
3. DP: longest path to node = 1 + max(longest paths to predecessors)
4. Return max of all DP values

Both approaches give O(M*N) time!"
```

#### Phase 4: Edge Case Handling
- **Empty matrix**: Return 0
- **Single cell**: Return 1
- **All equal values**: Return 1 (no increasing moves)
- **Matrix with one row/column**: Handle correctly

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Matrix: [
  [9, 9, 4],
  [6, 6, 8],
  [2, 1, 1]
]

Human thinking:
"DFS with Memoization from (0,2) = 4:
- Can move to (1,2) = 8 (larger) or (2,1) = 6 (larger)
- Try (1,2) = 8:
  - From (1,2), can move to (2,2) = 1 (smaller) - no
  - Can move to (1,1) = 6 (smaller) - no
  - Can move to (0,2) = 4 (smaller) - no
  - Can move to (2,2) = 1 (smaller) - no
  - Path length: 2
- Try (2,1) = 6:
  - From (2,1), can move to (2,2) = 1 (smaller) - no
  - Can move to (1,1) = 6 (equal) - no
  - Can move to (1,0) = 6 (equal) - no
  - Can move to (2,0) = 2 (smaller) - no
  - Path length: 2
- Cache result: longest from (0,2) is 2

Continue for all cells, find maximum = 4 ✓"
```

#### Phase 6: Intuition Validation
- **Why DAG works**: Increasing constraint prevents cycles
- **Why memoization works**: Each cell result computed once
- **Why topological sort works**: Processes nodes in dependency order
- **Why O(M*N)**: Each cell processed once with optimal approach

### Common Human Pitfalls & How to Avoid Them
1. **"Why not try all paths?"** → Exponential without memoization
2. **"Should I use BFS?"** → BFS finds shortest, not longest
3. **"What about equal values?"** → Clarify if equal allowed
4. **"Can I optimize further?"** → O(M*N) is already optimal
5. **"What about diagonal moves?"** → Clarify movement constraints

### Real-World Analogy
**Like planning the longest uphill hiking route:**
- You have a topographic map (matrix of elevations)
- Can only move to higher elevations
- Want to find the longest possible ascent trail
- Each step must go to a higher point
- Like a mountain climber seeking the longest continuous ascent

### Human-Readable Pseudocode
```
function longestIncreasingPath(matrix):
    if matrix is empty:
        return 0
    
    m, n = matrix dimensions
    memo = array[m][n] initialized to -1
    
    max_length = 0
    
    for i from 0 to m-1:
        for j from 0 to n-1:
            path_length = dfs(i, j)
            max_length = max(max_length, path_length)
    
    return max_length

function dfs(row, col):
    if memo[row][col] != -1:
        return memo[row][col]
    
    max_path = 1 // Current cell
    
    directions = [[-1,0], [1,0], [0,-1], [0,1]] // up, down, left, right
    
    for dir in directions:
        new_row, new_col = row + dir[0], col + dir[1]
        
        if new_row >= 0 and new_row < m and new_col >= 0 and new_col < n:
            if matrix[new_row][new_col] > matrix[row][col]:
                path_length = 1 + dfs(new_row, new_col)
                max_path = max(max_path, path_length)
    
    memo[row][col] = max_path
    return max_path
```

### Execution Visualization

### Example: Matrix = [[9,9,4],[6,6,8],[2,1,1]]
```
Matrix Visualization:
9 9 4
6 6 8
2 1 1

DFS from (2,0) = 2:
Step 1: At (2,0), value = 2
- Can move to (1,0) = 6 (larger) or (2,1) = 1 (smaller)
- Try (1,0) = 6:
  - From (1,0), can move to (0,0) = 9 (larger), (1,1) = 6 (equal)
  - Try (0,0) = 9:
    - From (0,0), no larger neighbors
    - Path length: 2 (2→6→9)
  - Try (1,1) = 6: equal, no move
  - Max from (1,0): 2
- Try (2,1) = 1: smaller, no move
- Cache result: longest from (2,0) is 2

Continue exploring all cells...
Longest path found: 4 (2→6→9) ✓
```

### Key Visualization Points:
- **Increasing Constraint**: Only move to larger values
- **Memoization**: Cache results to avoid recomputation
- **DFS Exploration**: Systematic path building
- **Maximum Tracking**: Track longest path found

### Memory Layout Visualization:
```
Matrix State During DFS:
9 9 4
6 6 8
2 1 1

Current Path: 2→6→9
Visited Cells: (2,0), (1,0), (0,0)
Current Position: (0,0) = 9
Path Length: 3
Next: No larger neighbors from (0,0)

Memo Cache State:
(2,0): 2  (computed)
(1,0): 2  (computed)
(0,0): 3  (computed)
Other cells: -1 (not computed yet)

Result: max of all cached values = 4 ✓
```

### Time Complexity Breakdown:
- **Without Memoization**: O(4^(M*N)) exponential time
- **With Memoization**: O(M*N) time, O(M*N) space
- **Topological Sort**: O(M*N*log(M*N)) time, O(M*N) space
- **DP Approach**: O(M*N) time, O(M*N) space

### Alternative Approaches:

#### 1. Topological Sort with DP (O(M*N*log(M*N)) time, O(M*N) space)
```go
func longestIncreasingPathTopological(matrix [][]int) int {
    // Build graph from smaller to larger values
    // Topologically sort nodes
    // DP: longest path to each node
    // Return max DP value
    // ... implementation details omitted
}
```
- **Pros**: No recursion, deterministic order
- **Cons**: More complex graph construction

#### 2. BFS for Longest Path (O(M*N*4^L) time, O(M*N) space)
```go
func longestIncreasingPathBFS(matrix [][]int) int {
    // BFS with DP to track longest paths
    // More complex than DFS for this problem
    // ... implementation details omitted
}
```
- **Pros**: Can find all longest paths
- **Cons**: More complex, overkill for single longest

#### 3. Dynamic Programming by Value Order (O(M*N) time, O(M*N) space)
```go
func longestIncreasingPathDP(matrix [][]int) int {
    // Sort cells by value, process in order
    // DP: longest path ending at each cell
    // ... implementation details omitted
}
```
- **Pros**: Clean DP formulation
- **Cons**: Requires sorting step, more complex

### Extensions for Interviews:
- **Diagonal Movement**: Allow 8-directional movement
- **Equal Values**: Handle cells with same values
- **Path Counting**: Count number of longest paths
- **Multiple Queries**: Answer multiple path length queries
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []struct {
		matrix    []string
		description string
	}{
		{[]string{"9", "9", "4", "6", "6", "8", "2"}, "Single column"},
		{[]string{"3", "4", "5"}, "3", "2", "6", "2", "2", "1"}, "Standard case"},
		{[]string{"1"}, "Single cell"},
		{[]string{"1", "1"}, "Two cells same value"},
		{[]string{"1", "2"}, "Two cells increasing"},
		{[]string{"2", "1"}, "Two cells decreasing"},
		{[]string{"7", "7", "7"}, "All same values"},
		{[]string{"1", "2", "3"}, {"6", "5", "4"}, {"7", "8", "9"}, "3x3 matrix"},
		{[]string{"3", "4", "5", "6"}, "2", "2", "2", "2"}, "Complex case"},
	}
	
	for i, tc := range testCases {
		matrix := createMatrix(tc.matrix)
		
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Matrix: %v\n", tc.matrix)
		
		result1 := longestIncreasingPath(matrix)
		result2 := longestIncreasingPathTopological(matrix)
		result3 := longestIncreasingPathBFS(matrix)
		
		fmt.Printf("  DFS + Memo: %d\n", result1)
		fmt.Printf("  Topological: %d\n", result2)
		fmt.Printf("  BFS: %d\n\n", result3)
	}
}
