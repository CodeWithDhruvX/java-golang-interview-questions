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
