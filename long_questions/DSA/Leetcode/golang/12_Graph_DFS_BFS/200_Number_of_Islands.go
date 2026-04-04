package main

import "fmt"

// 200. Number of Islands
// Time: O(M*N), Space: O(M*N) for recursion stack
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	
	m, n := len(grid), len(grid[0])
	islands := 0
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				islands++
				dfs(grid, i, j, m, n)
			}
		}
	}
	
	return islands
}

func dfs(grid [][]byte, i, j, m, n int) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
		return
	}
	
	// Mark as visited
	grid[i][j] = '0'
	
	// Explore all 4 directions
	dfs(grid, i+1, j, m, n)
	dfs(grid, i-1, j, m, n)
	dfs(grid, i, j+1, m, n)
	dfs(grid, i, j-1, m, n)
}

// Helper function to create grid from string
func createGrid(str []string) [][]byte {
	grid := make([][]byte, len(str))
	for i, row := range str {
		grid[i] = []byte(row)
	}
	return grid
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Grid-based DFS/Flood Fill
- **Grid Traversal**: 2D grid treated as graph where cells are nodes
- **Connected Components**: Islands are connected components of '1's
- **DFS Exploration**: Mark visited cells by changing '1' to '0'
- **Four Directions**: Explore up, down, left, right neighbors

## 2. PROBLEM CHARACTERISTICS
- **2D Grid**: Matrix of characters representing land/water
- **Connected Land**: Adjacent (4-directional) '1's form islands
- **Island Counting**: Count number of distinct connected components
- **In-place Modification**: Can modify grid to mark visited cells

## 3. SIMILAR PROBLEMS
- Flood Fill (LeetCode 733) - Similar with color filling
- Surrounded Regions (LeetCode 130) - Similar with border detection
- Max Area of Island (LeetCode 695) - Count area instead of number
- Number of Enclaves (LeetCode 1020) - Similar with border connectivity

## 4. KEY OBSERVATIONS
- **Grid as graph**: Each cell is a node, edges connect adjacent cells
- **Connected components**: Each island is a connected component
- **In-place marking**: Change '1' to '0' to mark visited
- **4-directional adjacency**: Only consider horizontal and vertical neighbors

## 5. VARIATIONS & EXTENSIONS
- **8-directional adjacency**: Include diagonal connections
- **BFS approach**: Use queue instead of recursion
- **Island perimeter**: Calculate perimeter of each island
- **Largest island**: Find island with maximum area

## 6. INTERVIEW INSIGHTS
- Always clarify: "Should I count diagonally connected land?"
- Edge cases: empty grid, all water, all land, single cell
- Time complexity: O(M*N) - visit each cell once
- Space complexity: O(M*N) for recursion stack in worst case

## 7. COMMON MISTAKES
- Not checking grid boundaries properly
- Forgetting to mark cells as visited
- Using wrong adjacency (8-directional vs 4-directional)
- Not handling empty grid case
- Stack overflow on very large grids

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(M*N) time, O(M*N) space
- **Iterative BFS**: Can avoid recursion stack limitations
- **Union Find**: Can use DSU for connected component counting
- **Early termination**: Not applicable (need to scan entire grid)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like counting islands on a map:**
- You have a satellite image showing land and water (grid)
- You want to count how many distinct islands there are
- Each island consists of connected land masses
- As you discover each island, you "paint" it so you don't count it again
- Continue scanning the entire map until all land is accounted for

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: 2D grid of '1's (land) and '0's (water)
2. **Goal**: Count number of distinct islands
3. **Output**: Integer count of islands
4. **Constraint**: Islands are 4-directionally connected

#### Phase 2: Key Insight Recognition
- **"Connected components"** → Each island is a connected component
- **"Grid traversal"** → Need to scan every cell in the grid
- **"Marking visited"** → Change '1' to '0' to avoid recounting
- **"DFS natural fit"** → Recursive exploration of connected land

#### Phase 3: Strategy Development
```
Human thought process:
"I need to count connected groups of land.
I'll scan the entire grid cell by cell.
When I find unvisited land ('1'), I've found a new island.
I'll use DFS to explore and mark all connected land as visited.
Then continue scanning for the next unvisited land cell."
```

#### Phase 4: Edge Case Handling
- **Empty grid**: Return 0
- **All water**: Return 0
- **All land**: Return 1
- **Single cell**: Handle appropriately based on value

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Grid:
11110
11010
11000
00000

Human thinking:
"I'll scan row by row:

Row 0, Col 0: '1' - Found island #1!
Start DFS from (0,0):
- Mark (0,0) as '0'
- Explore neighbors: (0,1), (1,0) are '1's
- Continue exploring all connected land...
Eventually mark entire connected component as '0'

Continue scanning:
Row 0, Col 1: now '0' (visited)
Row 0, Col 2: now '0' (visited)
Row 0, Col 3: '0' (water)
Row 1, Col 0: now '0' (visited)
Row 1, Col 1: now '0' (visited)
Row 1, Col 2: '1' - Found island #2!
Start DFS from (1,2)...

Final count: 3 islands"
```

#### Phase 6: Intuition Validation
- **Why DFS works**: Natural for exploring connected components
- **Why in-place marking**: Efficient way to track visited cells
- **Why O(M*N)**: Each cell visited exactly once
- **Why 4-directional**: Problem specifies horizontal/vertical adjacency

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use BFS?"** → Can, but DFS is more natural for recursive exploration
2. **"Should I use extra space for visited?"** → In-place marking is more efficient
3. **"What about very large grids?"** → Consider iterative BFS to avoid stack overflow
4. **"Can I optimize further?"** → Already optimal for this problem

### Real-World Analogy
**Like counting lakes in a satellite image:**
- You have a satellite image showing water bodies and land
- You want to count how many distinct lakes there are
- Each lake consists of connected water areas
- As you identify each lake, you mark it on your map
- Continue until all water bodies are identified and counted

### Human-Readable Pseudocode
```
function numIslands(grid):
    if grid is empty:
        return 0
    
    islands = 0
    rows = length(grid)
    cols = length(grid[0])
    
    for i from 0 to rows-1:
        for j from 0 to cols-1:
            if grid[i][j] == '1':
                islands++
                dfs(grid, i, j, rows, cols)
    
    return islands

function dfs(grid, i, j, rows, cols):
    if i < 0 or i >= rows or j < 0 or j >= cols or grid[i][j] != '1':
        return
    
    grid[i][j] = '0'  // Mark as visited
    
    dfs(grid, i+1, j, rows, cols)  // Down
    dfs(grid, i-1, j, rows, cols)  // Up
    dfs(grid, i, j+1, rows, cols)  // Right
    dfs(grid, i, j-1, rows, cols)  // Left
```

### Execution Visualization

### Example Grid:
```
11110
11010
11000
00000
```

### DFS Process:
```
Initial grid, islands = 0

Scan (0,0): '1' → islands = 1
DFS from (0,0):
- Mark (0,0) = '0'
- Explore (0,1): '1' → mark '0', explore neighbors
- Explore (1,0): '1' → mark '0', explore neighbors
- Continue until all connected land marked

Grid after first island:
00000
00010
00000
00000

Continue scanning:
(1,2): '1' → islands = 2
DFS from (1,2):
- Mark (1,2) = '0'
- No connected neighbors

Final grid:
00000
00000
00000
00000

Final count: 3 islands
```

### Key Visualization Points:
- **Scanning pattern**: Row-by-row, left-to-right scanning
- **Island discovery**: First '1' in each component triggers new island
- **DFS exploration**: Recursively marks all connected land
- **In-place marking**: Changes '1' to '0' to track visited

### Memory Layout Visualization:
```
Grid Evolution:
11110    →    00000    →    00000
11010         00010         00000
11000         00000         00000
00000         00000         00000

Islands: 0      →    1        →    2        →    3
```

### Time Complexity Breakdown:
- **Each cell visited**: O(1) work per cell
- **Total cells**: M*N
- **Total time**: O(M*N)
- **Space**: O(M*N) in worst case for recursion stack

### Alternative Approaches:

#### 1. Iterative BFS Approach (O(M*N) time, O(M*N) space)
```go
func numIslandsBFS(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    m, n := len(grid), len(grid[0])
    islands := 0
    directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                islands++
                queue := [][2]int{{i, j}}
                grid[i][j] = '0'
                
                for len(queue) > 0 {
                    current := queue[0]
                    queue = queue[1:]
                    
                    for _, dir := range directions {
                        ni, nj := current[0]+dir[0], current[1]+dir[1]
                        if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == '1' {
                            grid[ni][nj] = '0'
                            queue = append(queue, [2]int{ni, nj})
                        }
                    }
                }
            }
        }
    }
    
    return islands
}
```
- **Pros**: No recursion stack, handles very large grids
- **Cons**: Slightly more complex implementation

#### 2. Union Find Approach (O(M*N*α(M*N)) time, O(M*N) space)
```go
func numIslandsUnionFind(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    m, n := len(grid), len(grid[0])
    parent := make([]int, m*n)
    rank := make([]int, m*n)
    
    // Initialize Union Find
    for i := 0; i < m*n; i++ {
        parent[i] = i
        rank[i] = 0
    }
    
    find := func(x int) int {
        for parent[x] != x {
            parent[x] = parent[parent[x]]
            x = parent[x]
        }
        return x
    }
    
    union := func(x, y int) {
        rootX, rootY := find(x), find(y)
        if rootX != rootY {
            if rank[rootX] < rank[rootY] {
                parent[rootX] = rootY
            } else if rank[rootX] > rank[rootY] {
                parent[rootY] = rootX
            } else {
                parent[rootY] = rootX
                rank[rootX]++
            }
        }
    }
    
    // Union adjacent land cells
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                if i > 0 && grid[i-1][j] == '1' {
                    union(i*n+j, (i-1)*n+j)
                }
                if j > 0 && grid[i][j-1] == '1' {
                    union(i*n+j, i*n+(j-1))
                }
            }
        }
    }
    
    // Count unique roots for land cells
    roots := make(map[int]bool)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                roots[find(i*n+j)] = true
            }
        }
    }
    
    return len(roots)
}
```
- **Pros**: Efficient for multiple queries on same grid
- **Cons**: More complex, higher constant factors

#### 3. Iterative DFS with Stack (O(M*N) time, O(M*N) space)
```go
func numIslandsIterativeDFS(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    
    m, n := len(grid), len(grid[0])
    islands := 0
    directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                islands++
                stack := [][2]int{{i, j}}
                grid[i][j] = '0'
                
                for len(stack) > 0 {
                    current := stack[len(stack)-1]
                    stack = stack[:len(stack)-1]
                    
                    for _, dir := range directions {
                        ni, nj := current[0]+dir[0], current[1]+dir[1]
                        if ni >= 0 && ni < m && nj >= 0 && nj < n && grid[ni][nj] == '1' {
                            grid[ni][nj] = '0'
                            stack = append(stack, [2]int{ni, nj})
                        }
                    }
                }
            }
        }
    }
    
    return islands
}
```
- **Pros**: No recursion, similar to DFS but iterative
- **Cons**: Stack management complexity

### Extensions for Interviews:
- **8-directional adjacency**: Include diagonal connections
- **Island perimeter**: Calculate perimeter of each island
- **Largest island**: Find island with maximum area
- **Multiple queries**: Answer multiple island count queries efficiently
- **Dynamic grid**: Handle grid updates and answer queries
*/
func main() {
	// Test cases
	testCases := []struct {
		grid []string
	}{
		{[]string{"11110", "11010", "11000", "00000"}},
		{[]string{"11000", "11000", "00100", "00011"}},
		{[]string{"1"}},
		{[]string{"0"}},
		{[]string{"111", "111", "111"}},
		{[]string{"000", "000", "000"}},
		{[]string{"10101", "01010", "10101"}},
		{[]string{"10000", "01000", "00100", "00010", "00001"}},
		{[]string{"110", "011", "001"}},
		{[]string{}},
	}
	
	for i, tc := range testCases {
		grid := createGrid(tc.grid)
		result := numIslands(grid)
		fmt.Printf("Test Case %d: %v -> Islands: %d\n", i+1, tc.grid, result)
	}
}
