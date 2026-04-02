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
