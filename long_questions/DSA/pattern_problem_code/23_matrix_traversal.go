package main

import "fmt"

// Pattern: Matrix Traversal (BFS / DFS)
// Difficulty: Medium
// Key Concept: Navigating a Grid using BFS for shortest path or spreading logic.

/*
INTUITION:
"Rotting Oranges"
You have a grid. 0=Empty, 1=Fresh, 2=Rotten.
Every minute, a rotten orange makes its 4 neighbors rotten.
How long until ALL fresh oranges are rotten?

Think of it like a Virus spreading.
- It starts at ALL rotten oranges simultaneously (Multi-source BFS).
- Minute 1: All immediate neighbors get infected.
- Minute 2: Neighbors of those neighbors get infected.
We traverse layer by layer (BFS).

PROBLEM:
Return minimum minutes to rot all oranges. If impossible, return -1.

ALGORITHM:
1. Scan grid.
   - Count Fresh oranges.
   - Add all Rotten ones to a Queue (Time 0).
2. BFS Loop (Level by level):
   - For each rotten orange in info for THIS minute:
     - Check neighbors (Up, Down, Left, Right).
     - If neighbor is Fresh:
       - Make it Rotten.
       - Decrease Fresh count.
       - Add to Queue.
   - Increment Time.
3. If Fresh count == 0, return Time. Else -1.
*/

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	queue := [][]int{}
	freshCount := 0

	// Step 1: Init (Multi-source)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			} else if grid[r][c] == 1 {
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0
	}

	minutes := 0
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// Step 2: BFS
	for len(queue) > 0 && freshCount > 0 {
		// Process this WHOLE minute layer
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			orange := queue[0]
			queue = queue[1:]
			r, c := orange[0], orange[1]

			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				// Check bounds and if fresh
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
					// Infect it
					grid[nr][nc] = 2
					freshCount--
					queue = append(queue, []int{nr, nc})
				}
			}
		}
		minutes++
	}

	if freshCount > 0 {
		return -1
	}
	return minutes
}

func main() {
	// Grid:
	// 2 1 1
	// 1 1 0
	// 0 1 1
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	fmt.Printf("Initial Grid: %v\n", grid)
	res := orangesRotting(grid)
	fmt.Printf("Minutes: %d\n", res) // Expected: 4
}
