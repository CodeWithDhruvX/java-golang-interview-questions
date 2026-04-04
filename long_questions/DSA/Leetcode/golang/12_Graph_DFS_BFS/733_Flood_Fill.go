package main

import "fmt"

// 733. Flood Fill
// Time: O(M*N), Space: O(M*N) for recursion stack
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return image
	}
	
	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}
	
	dfs(image, sr, sc, originalColor, color)
	return image
}

func dfs(image [][]int, row, col, originalColor, newColor int) {
	m, n := len(image), len(image[0])
	
	// Boundary check and color check
	if row < 0 || row >= m || col < 0 || col >= n || image[row][col] != originalColor {
		return
	}
	
	// Fill the current cell
	image[row][col] = newColor
	
	// Recursively fill all 4 directions
	dfs(image, row+1, col, originalColor, newColor)
	dfs(image, row-1, col, originalColor, newColor)
	dfs(image, row, col+1, originalColor, newColor)
	dfs(image, row, col-1, originalColor, newColor)
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Region-based DFS/Flood Fill
- **Grid Traversal**: 2D grid treated as graph where cells are nodes
- **Connected Region**: Fill all connected cells with same original color
- **DFS Exploration**: Recursively explore and fill connected cells
- **Four Directions**: Explore up, down, left, right neighbors

## 2. PROBLEM CHARACTERISTICS
- **2D Grid**: Matrix of integers representing colors
- **Starting Point**: Fill starts from specific coordinates (sr, sc)
- **Color Replacement**: Replace original color with new color
- **Connected Region**: Only fill cells connected to starting point

## 3. SIMILAR PROBLEMS
- Number of Islands (LeetCode 200) - Similar with island counting
- Surrounded Regions (LeetCode 130) - Similar with border detection
- Max Area of Island (LeetCode 695) - Similar with area calculation
- Island Perimeter (LeetCode 463) - Similar with boundary detection

## 4. KEY OBSERVATIONS
- **Grid as graph**: Each cell is a node, edges connect adjacent cells
- **Connected region**: Fill only cells connected to starting point
- **Color matching**: Only fill cells with original color
- **In-place modification**: Can modify grid directly

## 5. VARIATIONS & EXTENSIONS
- **8-directional flood fill**: Include diagonal connections
- **BFS approach**: Use queue instead of recursion
- **Multiple starting points**: Fill from multiple seeds
- **Boundary detection**: Stop at specific boundaries

## 6. INTERVIEW INSIGHTS
- Always clarify: "Should I count diagonally connected cells?"
- Edge cases: empty grid, single cell, original equals new color
- Time complexity: O(M*N) - visit each cell at most once
- Space complexity: O(M*N) for recursion stack in worst case

## 7. COMMON MISTAKES
- Not checking if original color equals new color
- Not checking grid boundaries properly
- Forgetting to check original color before filling
- Not handling empty grid case
- Stack overflow on very large grids

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(M*N) time, O(M*N) space
- **Iterative BFS**: Can avoid recursion stack limitations
- **Early termination**: Stop if original color equals new color
- **Boundary optimization**: Check boundaries efficiently

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like using the paint bucket tool in an image editor:**
- You have an image with different colored regions (grid)
- You click on a specific pixel (starting point)
- The paint bucket fills all connected pixels of the same color
- The fill spreads to all adjacent pixels of the original color
- The fill stops when it hits different colors or image boundaries

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: 2D grid, starting coordinates (sr, sc), new color
2. **Goal**: Fill connected region starting from (sr, sc) with new color
3. **Output**: Modified grid with region filled
4. **Constraint**: Only fill cells connected to starting point

#### Phase 2: Key Insight Recognition
- **"Connected region"** → Need to explore all connected cells
- **"Color matching"** → Only fill cells with original color
- **"DFS natural fit"** → Recursive exploration of connected region
- **"In-place modification"** → Can modify grid directly

#### Phase 3: Strategy Development
```
Human thought process:
"I need to fill a connected region starting from a point.
I'll use DFS to explore all connected cells with the same original color.
For each cell, I'll change its color and explore its neighbors.
I'll stop when I hit cells with different colors or boundaries.
If the original color equals the new color, no work is needed."
```

#### Phase 4: Edge Case Handling
- **Empty grid**: Return original grid
- **Original equals new**: Return original grid unchanged
- **Single cell**: Fill if it matches original color
- **Out of bounds**: Handled by boundary checks

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Grid:
1 1 1
1 1 0
1 0 1

Start: (1, 1), New Color: 2

Human thinking:
"I'll start filling from (1,1):

Original color at (1,1) = 1
New color = 2 (different, so proceed)

DFS from (1,1):
- Fill (1,1) = 2
- Explore neighbors: (0,1), (2,1), (1,0), (1,2)
  * (0,1) = 1 → fill, explore its neighbors
  * (2,1) = 0 → different color, stop
  * (1,0) = 1 → fill, explore its neighbors  
  * (1,2) = 0 → different color, stop

Continue exploring:
- From (0,1): fill (0,0), (0,2)
- From (1,0): fill (2,0)

Final grid:
2 2 2
2 2 0
2 0 1"
```

#### Phase 6: Intuition Validation
- **Why DFS works**: Natural for exploring connected regions
- **Why color check**: Only fill cells with original color
- **Why O(M*N)**: Each cell visited at most once
- **Why 4-directional**: Problem specifies horizontal/vertical adjacency

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use BFS?"** → Can, but DFS is more natural for recursive exploration
2. **"Should I check original color first?"** → Yes, avoid unnecessary work
3. **"What about very large grids?"** → Consider iterative BFS to avoid stack overflow
4. **"Can I optimize further?"** → Already optimal for this problem

### Real-World Analogy
**Like using a paint bucket tool in graphics software:**
- You have an image with different colored regions
- You click on a specific pixel to start filling
- The paint bucket fills all connected pixels of the same color
- The fill spreads outward until it hits different colors
- This is exactly how flood fill algorithms work in image editors

### Human-Readable Pseudocode
```
function floodFill(image, sr, sc, newColor):
    if image is empty:
        return image
    
    originalColor = image[sr][sc]
    if originalColor == newColor:
        return image
    
    dfs(image, sr, sc, originalColor, newColor)
    return image

function dfs(image, row, col, originalColor, newColor):
    if row < 0 or row >= rows or col < 0 or col >= cols:
        return
    
    if image[row][col] != originalColor:
        return
    
    image[row][col] = newColor
    
    dfs(image, row+1, col, originalColor, newColor)  // Down
    dfs(image, row-1, col, originalColor, newColor)  // Up
    dfs(image, row, col+1, originalColor, newColor)  // Right
    dfs(image, row, col-1, originalColor, newColor)  // Left
```

### Execution Visualization

### Example Grid:
```
1 1 1
1 1 0  
1 0 1
```

### Flood Fill Process:
```
Start: (1,1), New Color: 2, Original Color: 1

Initial state:
1 1 1
1 1 0
1 0 1

DFS from (1,1):
- Fill (1,1) = 2
- Explore (0,1): fill = 2, explore neighbors
- Explore (1,0): fill = 2, explore neighbors
- Explore (2,1): color 0, stop
- Explore (1,2): color 0, stop

Continue exploring:
- From (0,1): fill (0,0) = 2, (0,2) = 2
- From (1,0): fill (2,0) = 2

Final grid:
2 2 2
2 2 0
2 0 1
```

### Key Visualization Points:
- **Starting point**: Fill begins from specified coordinates
- **Color matching**: Only cells with original color are filled
- **Connected region**: Fill spreads to all connected cells
- **Boundary stopping**: Different colors act as boundaries

### Memory Layout Visualization:
```
Grid Evolution:
1 1 1    →    2 2 2    →    2 2 2
1 1 0         2 2 0         2 2 0
1 0 1         1 0 1         2 0 1

Fill spreads from center outward
```

### Time Complexity Breakdown:
- **Each cell visited**: O(1) work per cell
- **Connected region size**: K cells in worst case M*N
- **Total time**: O(M*N) in worst case
- **Space**: O(M*N) in worst case for recursion stack

### Alternative Approaches:

#### 1. Iterative BFS Approach (O(M*N) time, O(M*N) space)
```go
func floodFillBFS(image [][]int, sr int, sc int, color int) [][]int {
    if len(image) == 0 || len(image[0]) == 0 {
        return image
    }
    
    originalColor := image[sr][sc]
    if originalColor == color {
        return image
    }
    
    m, n := len(image), len(image[0])
    queue := [][2]int{{sr, sc}}
    image[sr][sc] = color
    directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        for _, dir := range directions {
            nr, nc := current[0]+dir[0], current[1]+dir[1]
            if nr >= 0 && nr < m && nc >= 0 && nc < n && image[nr][nc] == originalColor {
                image[nr][nc] = color
                queue = append(queue, [2]int{nr, nc})
            }
        }
    }
    
    return image
}
```
- **Pros**: No recursion stack, handles very large grids
- **Cons**: Slightly more complex implementation

#### 2. Iterative DFS with Stack (O(M*N) time, O(M*N) space)
```go
func floodFillIterativeDFS(image [][]int, sr int, sc int, color int) [][]int {
    if len(image) == 0 || len(image[0]) == 0 {
        return image
    }
    
    originalColor := image[sr][sc]
    if originalColor == color {
        return image
    }
    
    m, n := len(image), len(image[0])
    stack := [][2]int{{sr, sc}}
    directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    
    for len(stack) > 0 {
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if image[current[0]][current[1]] == originalColor {
            image[current[0]][current[1]] = color
            
            for _, dir := range directions {
                nr, nc := current[0]+dir[0], current[1]+dir[1]
                if nr >= 0 && nr < m && nc >= 0 && nc < n {
                    stack = append(stack, [2]int{nr, nc})
                }
            }
        }
    }
    
    return image
}
```
- **Pros**: No recursion, similar to DFS but iterative
- **Cons**: Stack management complexity

#### 3. Union Find Approach (O(M*N*α(M*N)) time, O(M*N) space)
```go
func floodFillUnionFind(image [][]int, sr int, sc int, color int) [][]int {
    if len(image) == 0 || len(image[0]) == 0 {
        return image
    }
    
    originalColor := image[sr][sc]
    if originalColor == color {
        return image
    }
    
    m, n := len(image), len(image[0])
    parent := make([]int, m*n)
    
    // Initialize Union Find
    for i := 0; i < m*n; i++ {
        parent[i] = i
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
            parent[rootY] = rootX
        }
    }
    
    // Union all connected cells with original color
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if image[i][j] == originalColor {
                if i > 0 && image[i-1][j] == originalColor {
                    union(i*n+j, (i-1)*n+j)
                }
                if j > 0 && image[i][j-1] == originalColor {
                    union(i*n+j, i*n+(j-1))
                }
            }
        }
    }
    
    // Find all cells in the same component as starting point
    startRoot := find(sr*n+sc)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if image[i][j] == originalColor && find(i*n+j) == startRoot {
                image[i][j] = color
            }
        }
    }
    
    return image
}
```
- **Pros**: Theoretical interest, good for multiple queries
- **Cons**: Overly complex for single flood fill operation

### Extensions for Interviews:
- **8-directional flood fill**: Include diagonal connections
- **Multiple starting points**: Fill from multiple seeds simultaneously
- **Boundary detection**: Stop at specific boundary conditions
- **Counting regions**: Count number of distinct regions
- **Largest region**: Find region with maximum area
*/
func main() {
	// Test cases
	testCases := []struct {
		image    [][]int
		sr, sc   int
		color    int
	}{
		{
			[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			1, 1, 2,
		},
		{
			[][]int{{0, 0, 0}, {0, 0, 0}},
			0, 0, 0,
		},
		{
			[][]int{{0}},
			0, 0, 1,
		},
		{
			[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			0, 0, 2,
		},
		{
			[][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}},
			1, 1, 2,
		},
		{
			[][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}},
			2, 2, 3,
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			0, 0, 10,
		},
	}
	
	for i, tc := range testCases {
		// Make a copy to preserve original for display
		original := make([][]int, len(tc.image))
		for j := range tc.image {
			original[j] = make([]int, len(tc.image[j]))
			copy(original[j], tc.image[j])
		}
		
		result := floodFill(tc.image, tc.sr, tc.sc, tc.color)
		fmt.Printf("Test Case %d: image=%v, start=(%d,%d), color=%d\n", 
			i+1, original, tc.sr, tc.sc, tc.color)
		fmt.Printf("  Result: %v\n\n", result)
	}
}
