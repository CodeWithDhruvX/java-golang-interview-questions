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
