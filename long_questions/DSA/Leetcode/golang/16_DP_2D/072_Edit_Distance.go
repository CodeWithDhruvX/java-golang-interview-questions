package main

import (
	"fmt"
	"math"
)

// 72. Edit Distance
// Time: O(M*N), Space: O(M*N)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	
	// dp[i][j] = minimum operations to convert word1[0:i] to word2[0:j]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Initialize base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i // delete all characters
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // insert all characters
	}
	
	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // no operation needed
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j],   // delete
					dp[i][j-1],   // insert
					dp[i-1][j-1], // replace
				)
			}
		}
	}
	
	return dp[m][n]
}

// Space optimized version: O(min(M,N)) space
func minDistanceOptimized(word1 string, word2 string) int {
	// Ensure word1 is the shorter string for space optimization
	if len(word1) > len(word2) {
		word1, word2 = word2, word1
	}
	
	m, n := len(word1), len(word2)
	prev := make([]int, m+1)
	current := make([]int, m+1)
	
	// Initialize first row
	for i := 0; i <= m; i++ {
		prev[i] = i
	}
	
	for j := 1; j <= n; j++ {
		current[0] = j
		for i := 1; i <= m; i++ {
			if word1[i-1] == word2[j-1] {
				current[i] = prev[i-1]
			} else {
				current[i] = 1 + min(
					prev[i],     // delete
					current[i-1], // insert
					prev[i-1],   // replace
				)
			}
		}
		// Swap prev and current
		prev, current = current, prev
	}
	
	return prev[m]
}

// Function to reconstruct the actual operations
func minDistanceWithPath(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Initialize base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	
	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j],
					dp[i][j-1],
					dp[i-1][j-1],
				)
			}
		}
	}
	
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}

func main() {
	// Test cases
	testCases := []struct {
		word1 string
		word2 string
	}{
		{"horse", "ros"},
		{"intention", "execution"},
		{"", ""},
		{"", "abc"},
		{"abc", ""},
		{"abc", "abc"},
		{"kitten", "sitting"},
		{"flaw", "lawn"},
		{"algorithm", "altruistic"},
		{"sunday", "saturday"},
	}
	
	for i, tc := range testCases {
		result1 := minDistance(tc.word1, tc.word2)
		result2 := minDistanceOptimized(tc.word1, tc.word2)
		result3 := minDistanceWithPath(tc.word1, tc.word2)
		
		fmt.Printf("Test Case %d: \"%s\" -> \"%s\"\n", i+1, tc.word1, tc.word2)
		fmt.Printf("  Standard: %d, Optimized: %d, With path: %d\n\n", 
			result1, result2, result3)
	}
}
