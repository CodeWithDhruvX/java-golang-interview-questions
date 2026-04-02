package main

import "fmt"

// 10. Regular Expression Matching
// Time: O(M*N), Space: O(M*N)
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	
	// dp[i][j] = true if s[0:i] matches p[0:j]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	
	// Empty string matches empty pattern
	dp[0][0] = true
	
	// Empty string matches patterns like a*, a*b*, a*b*c*
	for j := 2; j <= n; j += 2 {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	
	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				// Direct match or wildcard
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// Two cases:
				// 1. '*' matches zero occurrence of preceding character
				dp[i][j] = dp[i][j-2]
				// 2. '*' matches one or more occurrences of preceding character
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}
	
	return dp[m][n]
}

// Recursive with memoization
func isMatchMemo(s string, p string) bool {
	memo := make(map[string]bool)
	return isMatchHelper(s, p, memo)
}

func isMatchHelper(s string, p string, memo map[string]bool) bool {
	key := s + "|" + p
	if val, exists := memo[key]; exists {
		return val
	}
	
	if len(p) == 0 {
		return len(s) == 0
	}
	
	// Check if next character in pattern is '*'
	if len(p) >= 2 && p[1] == '*' {
		// Case 1: '*' matches zero characters
		if isMatchHelper(s, p[2:], memo) {
			memo[key] = true
			return true
		}
		
		// Case 2: '*' matches one or more characters
		if len(s) > 0 && (p[0] == s[0] || p[0] == '.') {
			if isMatchHelper(s[1:], p, memo) {
				memo[key] = true
				return true
			}
		}
	} else {
		// Normal character or '.' match
		if len(s) > 0 && (p[0] == s[0] || p[0] == '.') {
			if isMatchHelper(s[1:], p[1:], memo) {
				memo[key] = true
				return true
			}
		}
	}
	
	memo[key] = false
	return false
}

func main() {
	// Test cases
	testCases := []struct {
		s string
		p string
	}{
		{"aa", "a"},
		{"aa", "a*"},
		{"ab", ".*"},
		{"aab", "c*a*b"},
		{"", ".*"},
		{"", ""},
		{"a", ""},
		{"", "a"},
		{"ab", ".*c"},
		{"mississippi", "mis*is*p*."},
		{"ab", ".*.."},
		{"aaa", "a*a"},
		{"ab", ".*..a*"},
	}
	
	for i, tc := range testCases {
		result1 := isMatch(tc.s, tc.p)
		result2 := isMatchMemo(tc.s, tc.p)
		
		fmt.Printf("Test Case %d: \"%s\" vs \"%s\" -> DP: %t, Memo: %t\n", 
			i+1, tc.s, tc.p, result1, result2)
	}
}
