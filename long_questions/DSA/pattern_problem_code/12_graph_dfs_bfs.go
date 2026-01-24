package main

import "fmt"

// Pattern: Graph DFS / BFS (Matrix Traversal)
// Difficulty: Medium
// Key Concept: Exploring connected components in a grid using recursion (DFS).

/*
INTUITION:
You are looking at a map of logic (1s) and water (0s).
An "Island" is a cluster of 1s.
How do you count them?
Imagine "sinking" the islands.
When you fly over and see land (1):
1. Count it as an island (+1).
2. Land on it and run around (DFS) to every connected piece of land.
3. As you run, you "sink" the land (turn 1s into 0s) so you don't count it again later.
4. When you have explored the whole island, take off again and look for the next patch of land.

PROBLEM:
"Number of Islands"
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

ALGORITHM:
1. Loop through every cell `(r, c)`.
2. If `grid[r][c] == '1'`:
   - Found a new island! `count++`.
   - Call `dfs(r, c)` to sink the entire connected island.
3. `dfs(r, c)`:
   - Base Case: If out of bounds or water ('0'), return.
   - Sink: Mark `grid[r][c] = '0'`.
   - Recurse: Call DFS on Up, Down, Left, Right neighbors.
*/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Loop through every cell
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// If we find land, it's a new island
			if grid[r][c] == '1' {
				count++
				// Visit (sink) the whole island so we don't count its parts again
				dfsSink(grid, r, c)
			}
		}
	}
	return count
}

func dfsSink(grid [][]byte, r, c int) {
	rows := len(grid)
	cols := len(grid[0])

	// Base Cases:
	// 1. Out of bounds row
	// 2. Out of bounds col
	// 3. Already water (0)
	if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == '0' {
		return
	}

	// Action: Sink the current cell
	grid[r][c] = '0'

	// Recurse: Visit all 4 neighbors
	dfsSink(grid, r+1, c) // Down
	dfsSink(grid, r-1, c) // Up
	dfsSink(grid, r, c+1) // Right
	dfsSink(grid, r, c-1) // Left
}

func main() {
	// Grid:
	// 1 1 0 0 0
	// 1 1 0 0 0
	// 0 0 1 0 0
	// 0 0 0 1 1
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Printf("Counting Islands...\n")
	result := numIslands(grid)

	// DRY RUN Intuition:
	// (0,0) is '1'. Count=1. Sink (0,0), (0,1), (1,0), (1,1).
	// Loop continues... (0,1) is now '0' (skipped).
	// (2,2) is '1'. Count=2. Sink (2,2).
	// (3,3) is '1'. Count=3. Sink (3,3), (3,4).
	// Total 3.

	fmt.Printf("Islands: %d\n", result) // Expected: 3
}
