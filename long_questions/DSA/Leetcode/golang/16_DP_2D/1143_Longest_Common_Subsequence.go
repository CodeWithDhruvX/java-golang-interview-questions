package main

import "fmt"

// 1143. Longest Common Subsequence
// Time: O(M*N), Space: O(M*N)
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	
	// dp[i][j] = length of LCS of text1[0:i] and text2[0:j]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	
	return dp[m][n]
}

// Space optimized version: O(min(M,N)) space
func longestCommonSubsequenceOptimized(text1 string, text2 string) int {
	// Ensure text1 is the shorter string for space optimization
	if len(text1) > len(text2) {
		text1, text2 = text2, text1
	}
	
	m, n := len(text1), len(text2)
	prev := make([]int, m+1)
	current := make([]int, m+1)
	
	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if text1[i-1] == text2[j-1] {
				current[i] = prev[i-1] + 1
			} else {
				current[i] = max(prev[i], current[i-1])
			}
		}
		// Swap prev and current
		prev, current = current, prev
	}
	
	return prev[m]
}

// Function to actually reconstruct the LCS
func longestCommonSubsequenceReconstruct(text1 string, text2 string) string {
	m, n := len(text1), len(text2)
	
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	
	// Reconstruct the LCS
	var result []byte
	i, j := m, n
	
	for i > 0 && j > 0 {
		if text1[i-1] == text2[j-1] {
			result = append(result, text1[i-1])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}
	
	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	
	return string(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test cases
	testCases := []struct {
		text1 string
		text2 string
	}{
		{"abcde", "ace"},
		{"abc", "abc"},
		{"abc", "def"},
		{"", ""},
		{"", "abc"},
		{"abc", ""},
		{"AGGTAB", "GXTXAYB"},
		{"abcdef", "acbcf"},
		{"XMJYAUZ", "MZJAWXU"},
		{"ABCD", "ABDC"},
	}
	
	for i, tc := range testCases {
		result1 := longestCommonSubsequence(tc.text1, tc.text2)
		result2 := longestCommonSubsequenceOptimized(tc.text1, tc.text2)
		result3 := longestCommonSubsequenceReconstruct(tc.text1, tc.text2)
		
		fmt.Printf("Test Case %d: \"%s\", \"%s\"\n", i+1, tc.text1, tc.text2)
		fmt.Printf("  Length: %d, Optimized: %d, LCS: \"%s\"\n\n", 
			result1, result2, result3)
	}
}
