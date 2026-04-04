package main

import "fmt"

// 994. Rotting Oranges
// Time: O(M*N), Space: O(M*N) for queue
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	
	m, n := len(grid), len(grid[0])
	
	// Queue for BFS
	queue := [][3]int{} // {row, col, time}
	freshOranges := 0
	
	// Initialize queue with rotten oranges and count fresh oranges
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [3]int{i, j, 0})
			} else if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}
	
	// If no fresh oranges, return 0
	if freshOranges == 0 {
		return 0
	}
	
	// If no rotten oranges but fresh oranges exist
	if len(queue) == 0 {
		return -1
	}
	
	// Directions: up, down, left, right
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	minutes := 0
	
	// BFS
	for len(queue) > 0 {
		row, col, time := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]
		
		// Update maximum time
		if time > minutes {
			minutes = time
		}
		
		// Spread rot to adjacent fresh oranges
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n &&
				grid[newRow][newCol] == 1 {
				
				grid[newRow][newCol] = 2 // Rot the orange
				freshOranges--
				queue = append(queue, [3]int{newRow, newCol, time + 1})
			}
		}
	}
	
	// Check if all fresh oranges are rotten
	if freshOranges == 0 {
		return minutes
	}
	
	return -1 // Some fresh oranges remain
}

// Alternative approach using multi-level BFS
func orangesRottingMultiLevel(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	
	m, n := len(grid), len(grid[0])
	
	// Count fresh oranges and initialize queue
	freshOranges := 0
	queue := [][2]int{}
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}
	
	if freshOranges == 0 {
		return 0
	}
	
	if len(queue) == 0 {
		return -1
	}
	
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	minutes := 0
	
	for len(queue) > 0 {
		levelSize := len(queue)
		
		// Process all oranges at current level
		for i := 0; i < levelSize; i++ {
			row, col := queue[0][0], queue[0][1]
			queue = queue[1:]
			
			// Spread rot
			for _, dir := range directions {
				newRow, newCol := row+dir[0], col+dir[1]
				
				if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n &&
					grid[newRow][newCol] == 1 {
					
					grid[newRow][newCol] = 2
					freshOranges--
					queue = append(queue, [2]int{newRow, newCol})
				}
			}
		}
		
		// Only increment minutes if there are more oranges to process
		if len(queue) > 0 {
			minutes++
		}
	}
	
	if freshOranges == 0 {
		return minutes
	}
	
	return -1
}

// DFS approach (less efficient but interesting)
func orangesRottingDFS(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	
	m, n := len(grid), len(grid[0])
	freshOranges := 0
	
	// Count fresh oranges
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				freshOranges++
			}
		}
	}
	
	if freshOranges == 0 {
		return 0
	}
	
	// Simulate minute by minute
	minutes := 0
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	for {
		// Find all oranges that will rot this minute
		toRot := [][2]int{}
		
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 2 {
					// Check adjacent fresh oranges
					for _, dir := range directions {
						newRow, newCol := i+dir[0], j+dir[1]
						
						if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n &&
							grid[newRow][newCol] == 1 {
							
							toRot = append(toRot, [2]int{newRow, newCol})
						}
					}
				}
			}
		}
		
		// If no new oranges to rot, break
		if len(toRot) == 0 {
			break
		}
		
		// Rot the oranges
		for _, pos := range toRot {
			grid[pos[0]][pos[1]] = 2
			freshOranges--
		}
		
		minutes++
	}
	
	if freshOranges == 0 {
		return minutes
	}
	
	return -1
}

// Helper function to create grid from strings
func createGrid(gridStr []string) [][]int {
	grid := make([][]int, len(gridStr))
	for i, row := range gridStr {
		grid[i] = make([]int, len(row))
		for j, char := range row {
			grid[i][j] = int(char - '0') // Assuming single digits
		}
	}
	return grid
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Multi-Source BFS for Spread Simulation
- **Multi-Source BFS**: Start from all rotten oranges simultaneously
- **Level Processing**: Process all oranges rotting at same minute
- **Spread Simulation**: Rot spreads to adjacent fresh oranges
- **Time Tracking**: Track minutes until all fresh oranges rot

## 2. PROBLEM CHARACTERISTICS
- **Grid Spread**: Rot spreads from rotten to fresh oranges
- **Multi-Source**: Multiple starting points (all rotten oranges)
- **4-Direction Spread**: Rot spreads to adjacent cells
- **Time Calculation**: Minutes until all fresh oranges rot

## 3. SIMILAR PROBLEMS
- 01 Matrix (LeetCode 542) - Update matrix with distance
- Walls and Gates (LeetCode 286) - Multi-source BFS from gates
- Zombie in Matrix (LeetCode 286) - Similar spread pattern
- Contagion Spread - Disease spread simulation

## 4. KEY OBSERVATIONS
- **BFS Natural Fit**: Spread happens in layers/levels
- **Multi-Source**: All rotten oranges start spreading simultaneously
- **Level Processing**: All oranges rotting at same time form a level
- **Termination**: When no more fresh oranges or rotting stops

## 5. VARIATIONS & EXTENSIONS
- **Diagonal Spread**: Allow 8-directional rotting
- **Variable Spread Speed**: Different rotting rates
- **Obstacles**: Cells that block rot spread
- **Multiple Rot Types**: Different types of rot with different spread patterns

## 6. INTERVIEW INSIGHTS
- Always clarify: "Grid size? Empty cells? Spread directions?"
- Edge cases: no fresh oranges, no rotten oranges, all empty
- Time complexity: O(M*N) where M=rows, N=columns
- Space complexity: O(M*N) for queue in worst case
- Key insight: multi-source BFS for simultaneous spread

## 7. COMMON MISTAKES
- Not handling no fresh oranges case (should return 0)
- Not handling no rotten oranges with fresh oranges (should return -1)
- Wrong minute counting (incrementing for each orange vs each level)
- Not updating grid state properly
- Missing boundary checks

## 8. OPTIMIZATION STRATEGIES
- **Multi-Source BFS**: O(M*N) time, O(M*N) space - optimal
- **Multi-Level BFS**: O(M*N) time, O(M*N) space - cleaner
- **Simulation**: O(M*N*L) time, O(1) space - less efficient
- **Early Termination**: Stop when all fresh oranges rotted

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like disease spread or fire spreading:**
- You have multiple infection sources (rotten oranges)
- Disease spreads to adjacent healthy cells each minute
- All sources spread simultaneously
- Want to know when all healthy cells are infected
- Like watching a contagion spread in real-time

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: 2D grid with 0=empty, 1=fresh, 2=rotten oranges
2. **Goal**: Find minutes until all fresh oranges become rotten
3. **Rules**: Rot spreads to adjacent cells each minute
4. **Output**: Minutes to rot all fresh oranges, or -1 if impossible

#### Phase 2: Key Insight Recognition
- **"Multi-source BFS natural fit"** → All rotten oranges start spreading
- **"Level processing"** → All oranges rotting at same time form a level
- **"Simultaneous spread"** → Not sequential from each source
- **"Termination conditions"** → No fresh left or spread stops

#### Phase 3: Strategy Development
```
Human thought process:
"I need to simulate rot spread from multiple sources.
This is like multi-source BFS:

1. Initialize queue with all rotten oranges (time=0)
2. Count fresh oranges
3. BFS by levels:
   - Process all oranges at current time
   - Spread rot to adjacent fresh oranges
   - Add newly rotten oranges with time+1
   - Increment time after each level
4. Stop when no fresh oranges or queue empty
5. Return time if all rotted, else -1

This correctly simulates simultaneous spread!"
```

#### Phase 4: Edge Case Handling
- **No fresh oranges**: Return 0 (nothing to rot)
- **No rotten oranges**: Return -1 (no spread source)
- **All empty**: Return 0 (nothing to rot)
- **Single orange**: Handle based on its state

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Grid: [
  [2,1,1],
  [1,1,0],
  [0,1,1]
]

Human thinking:
"Initial state:
- Rotten oranges: (0,0) at time 0
- Fresh oranges: 6 total
- Queue: [(0,0,0)]

Minute 0 (time=0):
Process (0,0):
- Spread to (0,1) and (1,0) (both fresh)
- Queue: [(0,1,1), (1,0,1)]
- Fresh oranges remaining: 4

Minute 1 (time=1):
Process (0,1) and (1,0):
- From (0,1): spread to (0,2) (fresh)
- From (1,0): spread to (1,1) (fresh)
- Queue: [(0,2,2), (1,1,2)]
- Fresh oranges remaining: 2

Minute 2 (time=2):
Process (0,2) and (1,1):
- From (0,2): spread to (1,2) (empty)
- From (1,1): spread to (2,1) (fresh)
- Queue: [(2,1,3)]
- Fresh oranges remaining: 1

Minute 3 (time=3):
Process (2,1):
- Spread to (2,2) (fresh)
- Queue: [(2,2,4)]
- Fresh oranges remaining: 0

All fresh oranges rotted! Return 4 minutes ✓"
```

#### Phase 6: Intuition Validation
- **Why multi-source BFS works**: All sources spread simultaneously
- **Why level processing works**: All oranges rotting at same time
- **Why O(M*N)**: Each cell processed once
- **Why termination conditions work**: Handles all cases correctly

### Common Human Pitfalls & How to Avoid Them
1. **"Why not sequential BFS?"** → Would be wrong (sources don't wait)
2. **"Should I use DFS?"** → DFS explores depth, not time levels
3. **"What about diagonal spread?"** → Clarify movement constraints
4. **"Can I optimize further?"** → O(M*N) is already optimal
5. **"What about obstacles?"** → Clarify if empty cells block spread

### Real-World Analogy
**Like watching a disease spread in a population:**
- You have multiple infected individuals (rotten oranges)
- Disease spreads to adjacent healthy people each day
- All infected people spread simultaneously
- Want to know when entire population is infected
- Like epidemiology tracking contagion spread

### Human-Readable Pseudocode
```
function orangesRotting(grid):
    if grid is empty:
        return -1
    
    m, n = grid dimensions
    queue = []
    fresh_oranges = 0
    
    // Initialize queue with all rotten oranges
    for i from 0 to m-1:
        for j from 0 to n-1:
            if grid[i][j] == 2:
                queue.append([i, j, 0]) // row, col, time
            else if grid[i][j] == 1:
                fresh_oranges++
    
    if fresh_oranges == 0:
        return 0
    
    if len(queue) == 0:
        return -1
    
    directions = [[-1,0], [1,0], [0,-1], [0,1]] // up, down, left, right
    minutes = 0
    
    while queue is not empty:
        level_size = len(queue)
        
        // Process all oranges rotting at current time
        for k from 0 to level_size-1:
            row, col, time = queue.pop_front()
            
            // Spread rot to adjacent fresh oranges
            for dir in directions:
                new_row, new_col = row + dir[0], col + dir[1]
                
                if new_row >= 0 and new_row < m and new_col >= 0 and new_col < n:
                    if grid[new_row][new_col] == 1:
                        grid[new_row][new_col] = 2
                        fresh_oranges--
                        queue.append([new_row, new_col, time + 1])
        
        minutes = time + 1
    
    if fresh_oranges == 0:
        return minutes - 1
    else:
        return -1
```

### Execution Visualization

### Example: Grid = [[2,1,1],[1,1,0],[0,1,1]]
```
Initial State:
2 1 1
1 1 0
0 1 1

Queue: [(0,0,0)] (rotten orange at time 0)
Fresh oranges: 6

Minute 0 (process time 0):
Process (0,0):
- Spread to (0,1) and (1,0)
- Grid becomes: 2 2 1 / 2 1 0 / 0 1 1
- Queue: [(0,1,1), (1,0,1)]
- Fresh oranges: 4

Minute 1 (process time 1):
Process (0,1) and (1,0):
- From (0,1): spread to (0,2)
- From (1,0): spread to (1,1)
- Grid becomes: 2 2 2 / 2 2 0 / 0 1 1
- Queue: [(0,2,2), (1,1,2)]
- Fresh oranges: 2

Minute 2 (process time 2):
Process (0,2) and (1,1):
- From (0,2): spread to (1,2) (empty)
- From (1,1): spread to (2,1)
- Grid becomes: 2 2 2 / 2 2 0 / 0 2 1
- Queue: [(2,1,3)]
- Fresh oranges: 1

Minute 3 (process time 3):
Process (2,1):
- Spread to (2,2)
- Grid becomes: 2 2 2 / 2 2 0 / 0 2 2
- Queue: [(2,2,4)]
- Fresh oranges: 0

All fresh oranges rotted! Return 4 minutes ✓
```

### Key Visualization Points:
- **Multi-Source**: All rotten oranges start spreading
- **Level Processing**: All oranges rotting at same time
- **Simultaneous Spread**: Correctly models real-world spread
- **Termination**: When no fresh oranges remain

### Memory Layout Visualization:
```
Grid Evolution:
Time 0:     Time 1:     Time 2:     Time 3:     Time 4:
2 1 1       2 2 1       2 2 2       2 2 2       2 2 2
1 1 0  →    2 1 0  →    2 2 0  →    2 2 0  →    2 2 0
0 1 1       0 1 1       0 2 1       0 2 2       0 2 2

Queue Evolution:
Time 0: [(0,0,0)]
Time 1: [(0,1,1), (1,0,1)]
Time 2: [(0,2,2), (1,1,2)]
Time 3: [(2,1,3)]
Time 4: [(2,2,4)]

Fresh Orange Count: 6 → 4 → 2 → 1 → 0 ✓
```

### Time Complexity Breakdown:
- **Initial Scan**: O(M*N) time to find all rotten/fresh oranges
- **BFS Processing**: O(M*N) time (each cell processed once)
- **Queue Operations**: O(1) per operation
- **Total**: O(M*N) time, O(M*N) space for queue
- **Optimal**: Each cell visited once

### Alternative Approaches:

#### 1. Multi-Level BFS (O(M*N) time, O(M*N) space)
```go
func orangesRottingMultiLevel(grid [][]int) int {
    // Process all nodes at same level together
    // Cleaner minute counting
    // ... implementation details omitted
}
```
- **Pros**: Cleaner minute counting
- **Cons**: Same complexity as standard BFS

#### 2. Simulation Approach (O(M*N*L) time, O(1) space)
```go
func orangesRottingSimulation(grid [][]int) int {
    // Simulate minute by minute
    // Less efficient but more intuitive
    // ... implementation details omitted
}
```
- **Pros**: Very intuitive
- **Cons**: O(M*N*L) where L=minutes, less efficient

#### 3. Union-Find with Time Tracking (O(M*N*α(M*N)) time, O(M*N) space)
```go
func orangesRottingUnionFind(grid [][]int) int {
    // Use union-find to track connected components
    // Track minimum time for each component
    // More complex, overkill for this problem
    // ... implementation details omitted
}
```
- **Pros**: Can handle complex spread patterns
- **Cons**: Overkill, more complex than needed

### Extensions for Interviews:
- **Diagonal Spread**: Allow 8-directional rotting
- **Variable Spread Speed**: Different rotting rates
- **Obstacles**: Cells that block rot spread
- **Multiple Rot Types**: Different types of rot with different spread patterns
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []struct {
		grid    []string
		description string
	}{
		{[]string{"2", "1", "1"}, "Simple case"},
		{[]string{"2", "1", "1", "0", "1", "1"}, "More complex"},
		{[]string{"0", "2"}, "No fresh oranges"},
		{[]string{"1"}, "Single fresh orange"},
		{[]string{"2"}, "Single rotten orange"},
		{[]string{"0", "1", "2", "0", "1", "2"}, "Mixed pattern"},
		{[]string{"1", "1", "1"}, "All fresh, no rotten"},
		{[]string{"2", "2", "2"}, "All rotten"},
		{[]string{"0", "0", "0"}, "All empty"},
		{[]string{"1", "2", "0", "1", "2", "1", "2", "1"}, "Large grid"},
	}
	
	for i, tc := range testCases {
		grid := createGrid(tc.grid)
		
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Grid: %v\n", tc.grid)
		
		// Make copies for different approaches
		grid1 := createGrid(tc.grid)
		grid2 := createGrid(tc.grid)
		grid3 := createGrid(tc.grid)
		
		result1 := orangesRotting(grid1)
		result2 := orangesRottingMultiLevel(grid2)
		result3 := orangesRottingDFS(grid3)
		
		fmt.Printf("  BFS with time: %d\n", result1)
		fmt.Printf("  Multi-level BFS: %d\n", result2)
		fmt.Printf("  Simulation: %d\n\n", result3)
	}
}
