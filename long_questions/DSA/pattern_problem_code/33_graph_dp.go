package main

import (
	"fmt"
)

// Pattern: DP on Graphs
// Difficulty: Hard
// Key Concept: Performing DFS/BFS on a graph (or grid) with Memoization to compute path properties (often Longest Path).

/*
INTUITION:
In a grid, finding the "Longest Increasing Path" is deceptively similar to "Number of Islands" (BFS/DFS), but there's a catch:
If you just do BFS, you might revisit nodes inefficiently or struggle to find the "longest" since BFS finds "shortest".
If you do naive DFS, you recompute the same subpaths over and over.

Solution: Remember the result.
"If I know the longest path starting from cell (2,3) is 5, I don't need to re-walk it when I arrive at (2,3) from (2,2)."
This is DFS + Memoization.
Since the path must be strictly increasing, there are NO CYCLES possible. It acts like a DAG (Directed Acyclic Graph).

PROBLEM:
LeetCode 329. Longest Increasing Path in a Matrix.
Given an m x n integers matrix, return the length of the longest increasing path.
From each cell, you can either move in four directions: left, right, up, or down. You may not move diagonally or move outside the boundary.

ALGORITHM:
1. Initialize `memo[m][n]` with 0.
2. Iterate through every cell `(i, j)` in the matrix.
3. Call `dfs(i, j)`.
   - If `memo[i][j]` != 0, return it.
   - Initialize `maxLen = 1`.
   - For particular neighbor `(ni, nj)`:
     - If `matrix[ni][nj] > matrix[i][j]`:
       - `len = 1 + dfs(ni, nj)`
       - `maxLen = max(maxLen, len)`
   - `memo[i][j] = maxLen`
   - Return `maxLen`.
4. The answer is the maximum value in the entire `memo` table (or tracked globally).
*/

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	maxLen := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			len := dfs(matrix, i, j, m, n, memo)
			if len > maxLen {
				maxLen = len
			}
		}
	}
	return maxLen
}

func dfs(matrix [][]int, i, j, m, n int, memo [][]int) int {
	if memo[i][j] != 0 {
		return memo[i][j]
	}

	maxPath := 1 // Minimum path is the cell itself of length 1

	for _, d := range dirs {
		ni, nj := i+d[0], j+d[1]
		// Check bounds and Increasing condition
		if ni >= 0 && ni < m && nj >= 0 && nj < n && matrix[ni][nj] > matrix[i][j] {
			length := 1 + dfs(matrix, ni, nj, m, n, memo)
			if length > maxPath {
				maxPath = length
			}
		}
	}

	memo[i][j] = maxPath
	return maxPath
}

func main() {
	// 9 9 4
	// 6 6 8
	// 2 1 1
	// Longest: 1 -> 2 -> 6 -> 9 is Val 9? No.
	// 1 -> 2 -> 6 -> 9. Length 4.
	// Grid:
	// [9,9,4]
	// [6,6,8]
	// [2,1,1]
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	fmt.Printf("Max Length: %d\n", longestIncreasingPath(matrix))

	// 3 4 5
	// 3 2 6
	// 2 2 1
	// Path: 1->2->6->9 not possible.
	// Path: 3->4->5->6? No.
	// 1->2->2(X)
	// 1->2(middle)->6->? No.
	// Longest: 1 -> 2 -> 6 (Len 3)?
	// Or 3 -> 4 -> 5 (Len 3)?
	// Or 1 -> 2 -> 3 (left col) -> 4 -> 5? increasing...
	// Let's rely on code.
	matrix2 := [][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}
	fmt.Printf("Max Length: %d\n", longestIncreasingPath(matrix2)) // Expect 4 (1->2->6 or 1->2->3->4->5?)
	// Actually for Matrix 2:
	// 1(at 2,2) -> 2(at 1,1) -> 3(at 1,0) -> 4(at 0,1) -> 5(at 0,2). Len 5 ?
	// Wait 2(1,1) < 3(1,0) ? No 2 < 3. Yes.
	// 1 < 2. Yes.
	// So 1->2->3->4->5. Length 5?
	// But 2->6 (Len 2).
	// Code will solve.
}
