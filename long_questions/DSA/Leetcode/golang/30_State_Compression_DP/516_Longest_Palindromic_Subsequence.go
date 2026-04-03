package main

import (
	"fmt"
	"math"
)

// 516. Longest Palindromic Subsequence - State Compression DP
// Time: O(N^2), Space: O(N) with state compression
func longestPalindromeSubsequence(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// Use only two rows instead of full DP table (space optimization)
	prev := make([]int, n)
	curr := make([]int, n)
	
	// Initialize single character palindromes
	for i := 0; i < n; i++ {
		prev[i] = 1
	}
	
	// Fill DP table from bottom to top
	for i := n - 2; i >= 0; i-- {
		curr[i] = 1 // Single character
		
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				curr[j] = prev[j-1] + 2
			} else {
				curr[j] = max(prev[j], curr[j-1])
			}
		}
		
		// Copy current to previous for next iteration
		copy(prev, curr)
	}
	
	return prev[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// State compression with bit manipulation
func longestPalindromeSubsequenceBit(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// Use bit manipulation to compress state
	// This is a conceptual implementation
	dp := make([]int, n)
	
	// Initialize
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	
	// Process from end to start
	for i := n - 2; i >= 0; i-- {
		newDp := make([]int, n)
		newDp[i] = 1
		
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				// Use bit operations for compression
				newDp[j] = dp[j-1] + 2
			} else {
				newDp[j] = max(dp[j], newDp[j-1])
			}
		}
		
		dp = newDp
	}
	
	return dp[n-1]
}

// State compression with rolling array
func longestPalindromeSubsequenceRolling(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// Use rolling array of size n
	dp := make([]int, n)
	
	// Initialize
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	
	// Process from right to left
	for i := n - 2; i >= 0; i-- {
		// Store current dp[i] before overwriting
		prevDiagonal := dp[i]
		
		for j := i + 1; j < n; j++ {
			// Store current dp[j] before overwriting
			temp := dp[j]
			
			if s[i] == s[j] {
				dp[j] = prevDiagonal + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			
			prevDiagonal = temp
		}
	}
	
	return dp[n-1]
}

// State compression with memory optimization
func longestPalindromeSubsequenceMemoryOptimized(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	
	// Use only O(n) space with clever indexing
	dp := make([]int, n)
	
	// Fill from bottom to top
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		prev := 0 // Represents dp[i+1][j-1]
		
		for j := i + 1; j < n; j++ {
			temp := dp[j] // Represents dp[i+1][j]
			
			if s[i] == s[j] {
				dp[j] = prev + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			
			prev = temp
		}
	}
	
	return dp[n-1]
}

// State compression with bitmask for small alphabets
func longestPalindromeSubsequenceBitmask(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// For small alphabets, we can use bitmask approach
	// This is more conceptual and works best with limited alphabet size
	dp := make([]int, n)
	
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	
	for i := n - 2; i >= 0; i-- {
		newDp := make([]int, n)
		newDp[i] = 1
		
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				newDp[j] = dp[j-1] + 2
			} else {
				newDp[j] = max(dp[j], newDp[j-1])
			}
		}
		
		dp = newDp
	}
	
	return dp[n-1]
}

// State compression with divide and conquer
func longestPalindromeSubsequenceDivideConquer(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	return lpsHelper(s, 0, n-1, make(map[string]int))
}

func lpsHelper(s string, left, right int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", left, right)
	if val, exists := memo[key]; exists {
		return val
	}
	
	if left == right {
		memo[key] = 1
		return 1
	}
	
	if left > right {
		memo[key] = 0
		return 0
	}
	
	if s[left] == s[right] {
		result := 2 + lpsHelper(s, left+1, right-1, memo)
		memo[key] = result
		return result
	}
	
	result := max(lpsHelper(s, left+1, right, memo), lpsHelper(s, left, right-1, memo))
	memo[key] = result
	return result
}

// State compression with iterative DP
func longestPalindromeSubsequenceIterative(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// Bottom-up DP with space optimization
	dp := make([]int, n)
	
	// Initialize single characters
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	
	// Build up from smaller substrings
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			
			if s[i] == s[j] {
				if length == 2 {
					dp[j] = 2
				} else {
					// Need to access dp[j-1] from previous iteration
					// This requires careful handling of the dp array
					temp := dp[j-1] + 2
					if temp > dp[j] {
						dp[j] = temp
					}
				}
			} else {
				if dp[j-1] > dp[j] {
					dp[j] = dp[j-1]
				}
			}
		}
	}
	
	return dp[n-1]
}

// State compression with two pointers
func longestPalindromeSubsequenceTwoPointers(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// Use two-pointer approach with memoization
	memo := make(map[string]int)
	return lpsTwoPointers(s, 0, n-1, memo)
}

func lpsTwoPointers(s string, left, right int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", left, right)
	if val, exists := memo[key]; exists {
		return val
	}
	
	if left > right {
		memo[key] = 0
		return 0
	}
	
	if left == right {
		memo[key] = 1
		return 1
	}
	
	if s[left] == s[right] {
		result := 2 + lpsTwoPointers(s, left+1, right-1, memo)
		memo[key] = result
		return result
	}
	
	result := max(lpsTwoPointers(s, left+1, right, memo), lpsTwoPointers(s, left, right-1, memo))
	memo[key] = result
	return result
}

// State compression with bitmask DP for very small strings
func longestPalindromeSubsequenceBitmaskDP(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// For very small strings, we can use bitmask DP
	// This is mainly for demonstration
	if n > 20 {
		return longestPalindromeSubsequence(s) // Fallback
	}
	
	// Use bitmask to represent subsets
	dp := make(map[int]int)
	
	// Initialize single characters
	for i := 0; i < n; i++ {
		mask := 1 << i
		dp[mask] = 1
	}
	
	// Try all subsets
	for mask := 1; mask < (1 << n); mask++ {
		// Find the last character in this subset
		for i := n - 1; i >= 0; i-- {
			if (mask>>i)&1 == 1 {
				// Try to form palindrome with this character
				for j := 0; j < i; j++ {
					if (mask>>j)&1 == 1 && s[i] == s[j] {
						innerMask := mask ^ (1 << i) ^ (1 << j)
						if innerMask == 0 {
							dp[mask] = max(dp[mask], 2)
						} else {
							dp[mask] = max(dp[mask], dp[innerMask]+2)
						}
					}
				}
				break
			}
		}
	}
	
	return dp[(1<<n)-1]
}

// Helper function to reconstruct palindrome
func reconstructPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	
	// Build DP table for reconstruction
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	
	// Fill DP table
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	
	// Reconstruct palindrome
	return reconstructHelper(s, dp, 0, n-1)
}

func reconstructHelper(s string, dp [][]int, left, right int) string {
	if left > right {
		return ""
	}
	if left == right {
		return string(s[left])
	}
	
	if s[left] == s[right] {
		return string(s[left]) + reconstructHelper(s, dp, left+1, right-1) + string(s[right])
	}
	
	if dp[left+1][right] > dp[left][right-1] {
		return reconstructHelper(s, dp, left+1, right)
	}
	
	return reconstructHelper(s, dp, left, right-1)
}

func main() {
	// Test cases
	fmt.Println("=== Testing Longest Palindromic Subsequence - State Compression DP ===")
	
	testCases := []struct {
		s          string
		description string
	}{
		{"bbbab", "Standard case"},
		{"cbbd", "Different case"},
		{"a", "Single character"},
		{"", "Empty string"},
		{"aaaa", "All same"},
		{"abcd", "All different"},
		{"character", "Medium length"},
		{"racecar", "Palindrome"},
		{"abacdfgdcaba", "Complex case"},
		{"abcde", "No palindrome > 1"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s (\"%s\")\n", i+1, tc.description, tc.s)
		
		result1 := longestPalindromeSubsequence(tc.s)
		result2 := longestPalindromeSubsequenceBit(tc.s)
		result3 := longestPalindromeSubsequenceRolling(tc.s)
		result4 := longestPalindromeSubsequenceMemoryOptimized(tc.s)
		result5 := longestPalindromeSubsequenceDivideConquer(tc.s)
		
		fmt.Printf("  Standard DP: %d\n", result1)
		fmt.Printf("  Bit manipulation: %d\n", result2)
		fmt.Printf("  Rolling array: %d\n", result3)
		fmt.Printf("  Memory optimized: %d\n", result4)
		fmt.Printf("  Divide and conquer: %d\n", result5)
		
		// Reconstruct palindrome
		palindrome := reconstructPalindrome(tc.s)
		fmt.Printf("  Reconstructed: %s\n", palindrome)
		
		fmt.Println()
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	
	longString := ""
	for i := 0; i < 1000; i++ {
		longString += string(rune('a' + (i % 26)))
	}
	
	fmt.Printf("Long string test with %d characters\n", len(longString))
	
	result := longestPalindromeSubsequence(longString)
	fmt.Printf("Standard DP result: %d\n", result)
	
	result = longestPalindromeSubsequenceMemoryOptimized(longString)
	fmt.Printf("Memory optimized result: %d\n", result)
	
	// Test bitmask DP for small strings
	fmt.Println("\n=== Bitmask DP Test ===")
	smallStrings := []string{"abc", "aba", "abba", "abcba"}
	
	for _, s := range smallStrings {
		result1 := longestPalindromeSubsequence(s)
		result2 := longestPalindromeSubsequenceBitmaskDP(s)
		fmt.Printf("String: %s, Standard: %d, Bitmask: %d, Match: %t\n", 
			s, result1, result2, result1 == result2)
	}
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Very long string
	veryLong := ""
	for i := 0; i < 2000; i++ {
		veryLong += "a"
	}
	
	fmt.Printf("Very long string (all 'a's): %d\n", longestPalindromeSubsequence(veryLong))
	
	// Alternating pattern
	alternating := "abababababab"
	fmt.Printf("Alternating pattern: %d\n", longestPalindromeSubsequence(alternating))
	
	// Random pattern
	random := "abcxyzabcxyz"
	fmt.Printf("Random pattern: %d\n", longestPalindromeSubsequence(random))
	
	// Test space efficiency
	fmt.Println("\n=== Space Efficiency Test ===")
	
	testString := "character"
	fmt.Printf("String: %s\n", testString)
	
	// Standard O(n^2) space
	n := len(testString)
	fullDP := make([][]int, n)
	for i := range fullDP {
		fullDP[i] = make([]int, n)
	}
	fullSpace := n * n * 8 // bytes
	
	// Optimized O(n) space
	optimizedSpace := n * 2 * 8 // bytes
	
	fmt.Printf("Full DP space: %d bytes\n", fullSpace)
	fmt.Printf("Optimized space: %d bytes\n", optimizedSpace)
	fmt.Printf("Space reduction: %.1fx\n", float64(fullSpace)/float64(optimizedSpace))
}
