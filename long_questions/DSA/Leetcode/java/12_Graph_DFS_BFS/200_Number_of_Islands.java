import java.util.Arrays;

public class NumberOfIslands {
    
    // 200. Number of Islands
    // Time: O(M*N), Space: O(M*N) for recursion stack
    public static int numIslands(char[][] grid) {
        if (grid.length == 0 || grid[0].length == 0) {
            return 0;
        }
        
        int m = grid.length, n = grid[0].length;
        int islands = 0;
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '1') {
                    islands++;
                    dfs(grid, i, j, m, n);
                }
            }
        }
        
        return islands;
    }

    private static void dfs(char[][] grid, int i, int j, int m, int n) {
        if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1') {
            return;
        }
        
        // Mark as visited
        grid[i][j] = '0';
        
        // Explore all 4 directions
        dfs(grid, i + 1, j, m, n);
        dfs(grid, i - 1, j, m, n);
        dfs(grid, i, j + 1, m, n);
        dfs(grid, i, j - 1, m, n);
    }

    // Helper function to create grid from string array
    public static char[][] createGrid(String[] str) {
        char[][] grid = new char[str.length][];
        for (int i = 0; i < str.length; i++) {
            grid[i] = str[i].toCharArray();
        }
        return grid;
    }

    public static void main(String[] args) {
        // Test cases
        String[][] testCases = {
            {"11110", "11010", "11000", "00000"},
            {"11000", "11000", "00100", "00011"},
            {"1"},
            {"0"},
            {"111", "111", "111"},
            {"000", "000", "000"},
            {"10101", "01010", "10101"},
            {"10000", "01000", "00100", "00010", "00001"},
            {"110", "011", "001"},
            {}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            char[][] grid = createGrid(testCases[i]);
            int result = numIslands(grid);
            System.out.printf("Test Case %d: %s -> Islands: %d\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Graph Depth-First Search (DFS)
- **Graph Traversal**: Systematic exploration of all connected components
- **Island Counting**: Count distinct connected components
- **DFS Exploration**: Mark visited and explore all 4 directions
- **Grid Navigation**: 2D array with boundary checks

## 2. PROBLEM CHARACTERISTICS
- **2D Grid**: Character array representing land/water
- **Connected Components**: Count distinct islands
- **4-Direction Movement**: Up, down, left, right exploration
- **In-Place Modification**: Mark visited by modifying grid

## 3. SIMILAR PROBLEMS
- Number of Provinces
- Surrounded Regions
- Max Area of Island
- Pacific Atlantic Water Flow

## 4. KEY OBSERVATIONS
- Each '1' represents land, '0' represents water
- DFS explores all connected land cells
- Need to mark visited to avoid infinite loops
- Boundary checking prevents out-of-bounds access
- Each DFS call finds one complete island

## 5. VARIATIONS & EXTENSIONS
- BFS approach (queue-based)
- Union-Find for dynamic connectivity
- Count island sizes
- Find largest island
- Diagonal movement allowed

## 6. INTERVIEW INSIGHTS
- Clarify: "Are diagonals considered adjacent?"
- Edge cases: empty grid, all land, all water
- Time complexity: O(M×N) vs O(M×N×4) naive
- Space complexity: O(M×N) for recursion stack

## 7. COMMON MISTAKES
- Not marking visited cells
- Missing boundary checks
- Infinite recursion loops
- Modifying original grid incorrectly

## 8. OPTIMIZATION STRATEGIES
- BFS avoids deep recursion stacks
- Iterative DFS with explicit stack
- Union-Find for multiple queries
- Early termination for obvious cases

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like exploring a map:**
- You have a map with land (1) and water (0)
- You want to count how many distinct land masses exist
- For each unexplored land cell, start an expedition
- Explore all connected land cells (4 directions)
- Mark visited cells to avoid re-exploration
- Each complete exploration finds one island

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: 2D grid of characters
2. **Goal**: Count number of distinct islands
3. **Output**: Integer count of islands

#### Phase 2: Key Insight Recognition
- **"What defines an island?"** → Connected group of '1's
- **"How to find connected cells?"** → DFS/BFS exploration
- **"What are the boundaries?"** → Grid edges and visited cells
- **"How to avoid double-counting?"** → Mark visited cells

#### Phase 3: Strategy Development
```
Human thought process:
"I'll scan the entire grid:
1. For each unvisited land cell:
   - Start DFS to explore this island
   - Mark all connected land cells as visited
   - Count this as one island
2. Use 4-direction exploration (up, down, left, right)
3. In-place modification saves space"
```

#### Phase 4: Edge Case Handling
- **Empty grid**: Return 0 (no islands)
- **All water**: Return 0 (no islands)
- **All land**: Return 1 (one big island)
- **Single cell**: Return 1 if it's land, 0 if water

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Grid:
11110
11000
00000

Human thinking:
"Scan entire grid:

Cell (0,0): '1' - unvisited, start DFS
→ Explore (1,0):
   - Can go up? No (out of bounds)
   - Can go down? No (out of bounds)
   - Can go left? No (out of bounds)
   - Can go right? Yes, to (0,1) which is '1'
→ Explore (0,1):
     - Can go up? No
     - Can go down? No
     - Can go left? No
     - Can go right? No
→ Backtrack to (0,0)
→ All directions explored, island count = 1

Cell (0,1): '1' - already visited, skip

Cell (0,2): '0' - water, skip

Cell (1,0): '1' - unvisited, start DFS
→ Explore (1,0):
   - All directions out of bounds or water
→ Island count = 2

Continue scanning...
Final count: 3 islands ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: DFS explores all connected components
- **Why it's correct**: Each island counted exactly once
- **Why it's efficient**: Each cell visited at most once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just count '1's?"** → Need to ensure connectivity
2. **"What about diagonals?"** → Only 4 directions unless specified
3. **"How to track visited?"** → In-place modification or separate visited array
4. **"What about recursion depth?"** → May cause stack overflow for large grids

### Real-World Analogy
**Like counting continents on a world map:**
- You have a satellite image (grid) showing land and water
- You want to count how many distinct continents exist
- For each unexplored land area, send an explorer
- Explorer maps all connected land in that area
- Mark explored areas to avoid redundant exploration
- Each complete mapping reveals one continent

### Human-Readable Pseudocode
```
function numIslands(grid):
    if grid is empty:
        return 0
    
    islands = 0
    
    for i from 0 to grid.length-1:
        for j from 0 to grid[0].length-1:
            if grid[i][j] == '1' and not visited:
                islands += 1
                dfs(grid, i, j)
    
    return islands

function dfs(grid, i, j):
    if out of bounds or grid[i][j] != '1':
        return
        
    mark as visited
    grid[i][j] = '0'
    
    explore all 4 directions:
    dfs(grid, i+1, j)    // up
    dfs(grid, i-1, j)    // down
    dfs(grid, i, j+1)    // left
    dfs(grid, i, j-1)    // right
```

### Execution Visualization

### Example: Grid = ["11110","11000","00000"]
```
Initial Grid:
11110
11000
00000

DFS Process:
Start at (0,0): '1' - unvisited
→ Mark visited, explore 4 directions
→ Only (0,1) reachable, island = 1

Start at (0,1): '1' - visited, skip
Start at (0,2): '0' - water, skip
Start at (1,0): '1' - unvisited
→ Mark visited, explore 4 directions
→ All out of bounds/water, island = 2

Continue scanning...
Final: 3 islands found ✓
```

### Key Visualization Points:
- **DFS exploration** systematically finds connected components
- **In-place marking** saves additional space
- **4-direction movement** covers all adjacent cells
- **Island counting** happens when new unvisited land found

### Memory Layout Visualization:
```
Grid Evolution:
11110    →    01110    →    00110    →    00010    →    00000
11000    →    01000    →    00000    →    00000    →    00000
00000    →    00000    →    00000    →    00000    →    00000

DFS Traversal:
Island 1: (0,0) → (0,1)
Island 2: (1,0)
Island 3: (1,1) → (1,2) → (1,3)

Visited Pattern:
01110    00110    00010    00000
01000    00000    00000    00000
00000    00000    00000    00000    00000
```

### Time Complexity Breakdown:
- **Each Cell**: Visited at most once
- **DFS Calls**: Maximum 4 per cell (4 directions)
- **Total**: O(M×N) time where M×N is grid size
- **Space**: O(M×N) for recursion stack in worst case
- **Optimal**: Cannot do better than O(M×N) for this problem
*/
