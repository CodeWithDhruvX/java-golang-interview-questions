```go
package main

import (
	"sort"
	"strings"
)

// 1. Implement strStr() (Find Substring)
// Time: O(N*M) - Using built-in which is optimized, potentially O(N+M)
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	return strings.Index(haystack, needle)
}

// Naive implementation for educational purposes
// Time: O(N*M), Space: O(1)
func StrStrNaive(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	
	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for j < len(needle) {
			if haystack[i+j] != needle[j] {
				break
			}
			j++
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

// Brute Force strStr() - Even more naive approach
// Time: O(N^3), Space: O(1)
func StrStrBruteForce(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	
	// Try every possible starting position
	for i := 0; i <= len(haystack)-len(needle); i++ {
		matched := true
		// Check each character manually
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				matched = false
				break
			}
		}
		if matched {
			return i
		}
	}
	return -1
}

// 2. Longest Common Prefix
// Time: O(N*L*logN), Space: O(L)
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	
	// Sort the strings
	sort.Strings(strs)
	
	s1 := strs[0]
	s2 := strs[len(strs)-1] // Most different from s1
	
	idx := 0
	for idx < len(s1) && idx < len(s2) {
		if s1[idx] == s2[idx] {
			idx++
		} else {
			break
		}
	}
	return s1[:idx]
}

// Brute Force Longest Common Prefix
// Time: O(N^2 * L), Space: O(L)
func LongestCommonPrefixBruteForce(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	
	// Start with the first string as prefix
	prefix := strs[0]
	
	// Compare with every other string
	for i := 1; i < len(strs); i++ {
		prefix = commonPrefix(prefix, strs[i])
		if prefix == "" {
			break
		}
	}
	return prefix
}

func commonPrefix(s1, s2 string) string {
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}
	
	for i := 0; i < minLen; i++ {
		if s1[i] != s2[i] {
			return s1[:i]
		}
	}
	return s1[:minLen]
}
