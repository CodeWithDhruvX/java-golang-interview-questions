package main

import "fmt"

// 62. Unique Paths
// Time: O(M*N), Space: O(M*N)
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	
	// dp[i][j] = number of ways to reach cell (i,j)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	
	// Initialize first row and first column
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	
	// Fill the dp table
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	
	return dp[m-1][n-1]
}

// Space optimized version: O(N) space
func uniquePathsOptimized(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	
	// Use 1D array representing current row
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	
	return dp[n-1]
}

// Mathematical approach using combinatorics
func uniquePathsMath(m int, n int) int {
	// Total steps = (m-1) + (n-1)
	// Choose (m-1) down steps or (n-1) right steps
	// C(total, m-1) or C(total, n-1)
	
	total := m + n - 2
	down := m - 1
	right := n - 1
	
	// Use min(down, right) to reduce computation
	if down > right {
		down, right = right, down
	}
	
	result := 1
	for i := 1; i <= down; i++ {
		result = result * (total - down + i) / i
	}
	
	return result
}

func main() {
	// Test cases
	testCases := []struct {
		m int
		n int
	}{
		{3, 7},
		{3, 2},
		{1, 1},
		{1, 5},
		{5, 1},
		{2, 2},
		{10, 10},
		{7, 3},
		{3, 3},
		{100, 1},
		{1, 100},
	}
	
	for i, tc := range testCases {
		result1 := uniquePaths(tc.m, tc.n)
		result2 := uniquePathsOptimized(tc.m, tc.n)
		result3 := uniquePathsMath(tc.m, tc.n)
		
		fmt.Printf("Test Case %d: m=%d, n=%d -> DP: %d, Optimized: %d, Math: %d\n", 
			i+1, tc.m, tc.n, result1, result2, result3)
	}
}
