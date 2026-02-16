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

// 2. Longest Common Prefix
// Time: O(N*L*logN), Space: O(L)
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	
	// Sort the strings
	sort.Strings(strs)
	
	s1 := strs[0]
	s2 := strs[len-1] // Most different from s1
	
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
```
