package main

import (
	"fmt"
	"math"
)

// 28. Find the Index of the First Occurrence in a String - KMP Algorithm
// Time: O(N + M), Space: O(M) where N is haystack length, M is needle length
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	
	// Build LPS (Longest Prefix Suffix) array
	lps := buildLPS(needle)
	
	i, j := 0, 0 // i: haystack index, j: needle index
	
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
			
			if j == len(needle) {
				return i - j // Found match
			}
		} else {
			if j != 0 {
				j = lps[j-1] // Use LPS to skip comparisons
			} else {
				i++
			}
		}
	}
	
	return -1 // Not found
}

// Build LPS array for KMP
func buildLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0 // Length of previous longest prefix suffix
	
	i := 1
	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	
	return lps
}

// KMP with detailed tracing
func strStrKMPDetailed(haystack string, needle string) (int, []string) {
	if needle == "" {
		return 0, []string{"Empty needle, returning 0"}
	}
	
	var trace []string
	trace = append(trace, fmt.Sprintf("Building LPS for: %s", needle))
	
	lps := buildLPSWithTrace(needle, &trace)
	trace = append(trace, fmt.Sprintf("LPS array: %v", lps))
	
	i, j := 0, 0
	trace = append(trace, fmt.Sprintf("Starting search: i=%d, j=%d", i, j))
	
	for i < len(haystack) {
		trace = append(trace, fmt.Sprintf("Comparing haystack[%d]='%c' with needle[%d]='%c'", 
			i, haystack[i], j, needle[j]))
		
		if haystack[i] == needle[j] {
			i++
			j++
			trace = append(trace, fmt.Sprintf("Match! i=%d, j=%d", i, j))
			
			if j == len(needle) {
				trace = append(trace, fmt.Sprintf("Found match at index %d", i-j))
				return i - j, trace
			}
		} else {
			if j != 0 {
				trace = append(trace, fmt.Sprintf("Mismatch, using LPS[%d-1]=%d", j, lps[j-1]))
				j = lps[j-1]
			} else {
				trace = append(trace, fmt.Sprintf("Mismatch, moving i forward"))
				i++
			}
		}
	}
	
	trace = append(trace, "No match found")
	return -1, trace
}

func buildLPSWithTrace(pattern string, trace *[]string) []int {
	lps := make([]int, len(pattern))
	length := 0
	
	i := 1
	for i < len(pattern) {
		*trace = append(*trace, fmt.Sprintf("Building LPS: i=%d, j=%d, pattern[i]='%c', pattern[length]='%c'", 
			i, length, pattern[i], pattern[length]))
		
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			*trace = append(*trace, fmt.Sprintf("Match! LPS[%d]=%d", i, length))
			i++
		} else {
			if length != 0 {
				*trace = append(*trace, fmt.Sprintf("Mismatch, using LPS[%d]=%d", length-1, lps[length-1]))
				length = lps[length-1]
			} else {
				lps[i] = 0
				*trace = append(*trace, fmt.Sprintf("No match, LPS[%d]=0", i))
				i++
			}
		}
	}
	
	return lps
}

// Alternative KMP implementation
func strStrKMPAlternative(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	
	// Build prefix table
	prefix := buildPrefixTable(needle)
	
	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for j < len(needle) && haystack[i+j] == needle[j] {
			j++
		}
		
		if j == len(needle) {
			return i
		} else if j > 0 {
			// Skip ahead using prefix table
			i += j - prefix[j-1] - 1
		}
	}
	
	return -1
}

func buildPrefixTable(pattern string) []int {
	prefix := make([]int, len(pattern))
	
	for i := 1; i < len(pattern); i++ {
		j := prefix[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = prefix[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		prefix[i] = j
	}
	
	return prefix
}

// Rabin-Karp approach for comparison
func strStrRabinKarp(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	
	// Calculate hash of needle
	needleHash := 0
	haystackHash := 0
	base := 256
	mod := 1000000007
	
	for i := 0; i < len(needle); i++ {
		needleHash = (needleHash*base + int(needle[i])) % mod
		haystackHash = (haystackHash*base + int(haystack[i])) % mod
	}
	
	// Precompute base^(len(needle)-1) mod
	pow := 1
	for i := 0; i < len(needle)-1; i++ {
		pow = (pow * base) % mod
	}
	
	// Slide window
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystackHash == needleHash {
			// Verify to avoid hash collision
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
		
		if i < len(haystack)-len(needle) {
			// Remove leftmost character
			haystackHash = (haystackHash - int(haystack[i])*pow%mod + mod) % mod
			// Add new character
			haystackHash = (haystackHash*base + int(haystack[i+len(needle)])) % mod
		}
	}
	
	return -1
}

// Brute force approach
func strStrBruteForce(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	
	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for j < len(needle) && haystack[i+j] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i
		}
	}
	
	return -1
}

// Built-in string search for comparison
func strStrBuiltIn(haystack string, needle string) int {
	return findSubstring(haystack, needle)
}

func findSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func main() {
	// Test cases
	fmt.Println("=== Testing String Search Algorithms ===")
	
	testCases := []struct {
		haystack   string
		needle     string
		description string
	}{
		{"sadbutsad", "sad", "Standard case - multiple matches"},
		{"leetcode", "leeto", "No match"},
		{"", "", "Both empty"},
		{"", "a", "Empty haystack"},
		{"a", "", "Empty needle"},
		{"a", "a", "Single character match"},
		{"a", "b", "Single character no match"},
		{"hello", "ll", "Partial match"},
		{"mississippi", "issi", "Overlapping patterns"},
		{"abcabcabc", "abcabc", "Repeated pattern"},
		{"aaaaa", "aaa", "All same characters"},
		{"abcdefgh", "efg", "Middle match"},
		{"abcdefgh", "h", "End match"},
		{"abcdefgh", "a", "Start match"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Haystack: '%s', Needle: '%s'\n", tc.haystack, tc.needle)
		
		result1 := strStr(tc.haystack, tc.needle)
		result2 := strStrKMPAlternative(tc.haystack, tc.needle)
		result3 := strStrRabinKarp(tc.haystack, tc.needle)
		result4 := strStrBruteForce(tc.haystack, tc.needle)
		result5 := strStrBuiltIn(tc.haystack, tc.needle)
		
		fmt.Printf("  KMP: %d\n", result1)
		fmt.Printf("  KMP Alternative: %d\n", result2)
		fmt.Printf("  Rabin-Karp: %d\n", result3)
		fmt.Printf("  Brute Force: %d\n", result4)
		fmt.Printf("  Built-in: %d\n\n", result5)
	}
	
	// Detailed KMP tracing
	fmt.Println("=== KMP Detailed Tracing ===")
	_, trace := strStrKMPDetailed("ababcabcabababd", "ababd")
	for _, step := range trace {
		fmt.Printf("  %s\n", step)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	longHaystack := ""
	longNeedle := ""
	
	// Build long strings
	for i := 0; i < 1000; i++ {
		longHaystack += "abcdefghijklmnopqrstuvwxyz"
	}
	for i := 0; i < 100; i++ {
		longNeedle += "xyz"
	}
	
	fmt.Printf("Long haystack length: %d\n", len(longHaystack))
	fmt.Printf("Long needle length: %d\n", len(longNeedle))
	
	result := strStr(longHaystack, longNeedle)
	fmt.Printf("KMP result: %d\n", result)
	
	result = strStrRabinKarp(longHaystack, longNeedle)
	fmt.Printf("Rabin-Karp result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Very long needle
	veryLongNeedle := ""
	for i := 0; i < 10000; i++ {
		veryLongNeedle += "a"
	}
	
	fmt.Printf("Very long needle: %d\n", strStr("aaaaa", veryLongNeedle))
	
	// Unicode characters
	unicodeHaystack := "你好世界你好世界"
	unicodeNeedle := "世界"
	fmt.Printf("Unicode search: %d\n", strStr(unicodeHaystack, unicodeNeedle))
	
	// Case sensitivity
	caseHaystack := "HelloWorld"
	caseNeedle := "hello"
	fmt.Printf("Case sensitive: %d\n", strStr(caseHaystack, caseNeedle))
	
	// Multiple matches
	multiHaystack := "abcabcabcabc"
	multiNeedle := "abc"
	fmt.Printf("Multiple matches: %d\n", strStr(multiHaystack, multiNeedle))
}
