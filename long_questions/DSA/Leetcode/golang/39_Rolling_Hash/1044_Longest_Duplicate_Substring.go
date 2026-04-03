package main

import (
	"fmt"
	"math"
)

// 1044. Longest Duplicate Substring - Rolling Hash with Binary Search
// Time: O(N log N), Space: O(N)
func longestDupSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}
	
	// Binary search for the length of longest duplicate
	left, right := 1, len(s)
	result := ""
	
	for left <= right {
		mid := left + (right-left)/2
		
		if duplicate := findDuplicate(s, mid); duplicate != "" {
			result = duplicate
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return result
}

// Find duplicate substring of given length using rolling hash
func findDuplicate(s string, length int) string {
	if length == 0 {
		return ""
	}
	
	// Rolling hash with base and mod
	base := 256
	mod := 1000000007
	
	// Calculate initial hash
	hash := 0
	for i := 0; i < length; i++ {
		hash = (hash*base + int(s[i])) % mod
	}
	
	// Store seen hashes
	seen := make(map[int]int)
	seen[hash] = 0
	
	// Precompute base^length mod
	pow := 1
	for i := 0; i < length; i++ {
		pow = (pow * base) % mod
	}
	
	// Rolling hash
	for i := 1; i <= len(s)-length; i++ {
		// Remove leftmost character
		hash = (hash - int(s[i-1])*pow%mod + mod) % mod
		
		// Add new character
		hash = (hash*base + int(s[i+length-1])) % mod
		
		// Check for collision
		if prevIndex, exists := seen[hash]; exists {
			// Verify actual substring (handle hash collisions)
			if s[prevIndex:prevIndex+length] == s[i:i+length] {
				return s[i : i+length]
			}
		} else {
			seen[hash] = i
		}
	}
	
	return ""
}

// Optimized version using double hashing
func longestDupSubstringDoubleHash(s string) string {
	if len(s) == 0 {
		return ""
	}
	
	left, right := 1, len(s)
	result := ""
	
	for left <= right {
		mid := left + (right-left)/2
		
		if duplicate := findDuplicateDoubleHash(s, mid); duplicate != "" {
			result = duplicate
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return result
}

func findDuplicateDoubleHash(s string, length int) string {
	if length == 0 {
		return ""
	}
	
	// Two different hash functions
	base1, mod1 := 256, 1000000007
	base2, mod2 := 257, 1000000009
	
	// Calculate initial hashes
	hash1, hash2 := 0, 0
	for i := 0; i < length; i++ {
		hash1 = (hash1*base1 + int(s[i])) % mod1
		hash2 = (hash2*base2 + int(s[i])) % mod2
	}
	
	// Store seen hash pairs
	seen := make(map[[2]int]int)
	seen[[2]int{hash1, hash2}] = 0
	
	// Precompute powers
	pow1, pow2 := 1, 1
	for i := 0; i < length; i++ {
		pow1 = (pow1 * base1) % mod1
		pow2 = (pow2 * base2) % mod2
	}
	
	// Rolling hash
	for i := 1; i <= len(s)-length; i++ {
		// Remove leftmost character
		hash1 = (hash1 - int(s[i-1])*pow1%mod1 + mod1) % mod1
		hash2 = (hash2 - int(s[i-1])*pow2%mod2 + mod2) % mod2
		
		// Add new character
		hash1 = (hash1*base1 + int(s[i+length-1])) % mod1
		hash2 = (hash2*base2 + int(s[i+length-1])) % mod2
		
		hashPair := [2]int{hash1, hash2}
		
		if prevIndex, exists := seen[hashPair]; exists {
			return s[i : i+length]
		} else {
			seen[hashPair] = i
		}
	}
	
	return ""
}

// Version using Rabin-Karp with bit manipulation
func longestDupSubstringRabinKarp(s string) string {
	if len(s) == 0 {
		return ""
	}
	
	left, right := 1, len(s)
	result := ""
	
	for left <= right {
		mid := left + (right-left)/2
		
		if duplicate := findDuplicateRabinKarp(s, mid); duplicate != "" {
			result = duplicate
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return result
}

func findDuplicateRabinKarp(s string, length int) string {
	if length == 0 {
		return ""
	}
	
	// Use base 26 for lowercase letters
	base := 26
	mod := 2 << 61 // Large prime-like number
	
	// Calculate initial hash
	hash := 0
	for i := 0; i < length; i++ {
		hash = (hash*base + int(s[i]-'a')) % mod
	}
	
	// Store seen hashes
	seen := make(map[int]int)
	seen[hash] = 0
	
	// Precompute base^length mod
	pow := 1
	for i := 0; i < length; i++ {
		pow = (pow * base) % mod
	}
	
	// Rolling hash
	for i := 1; i <= len(s)-length; i++ {
		// Remove leftmost character
		hash = (hash - int(s[i-1]-'a')*pow%mod + mod) % mod
		
		// Add new character
		hash = (hash*base + int(s[i+length-1]-'a')) % mod
		
		if prevIndex, exists := seen[hash]; exists {
			if s[prevIndex:prevIndex+length] == s[i:i+length] {
				return s[i : i+length]
			}
		} else {
			seen[hash] = i
		}
	}
	
	return ""
}

// Brute force approach for comparison (O(N^2))
func longestDupSubstringBruteForce(s string) string {
	maxLength := 0
	result := ""
	
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			// Find common prefix length
			k := 0
			for i+k < len(s) && j+k < len(s) && s[i+k] == s[j+k] {
				k++
			}
			
			if k > maxLength {
				maxLength = k
				result = s[i : i+k]
			}
		}
	}
	
	return result
}

// Suffix array approach (more advanced)
func longestDupSubstringSuffixArray(s string) string {
	// Build suffix array (simplified version)
	suffixes := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		suffixes[i] = s[i:]
	}
	
	// Sort suffixes (bubble sort for demonstration)
	for i := 0; i < len(suffixes)-1; i++ {
		for j := 0; j < len(suffixes)-i-1; j++ {
			if suffixes[j] > suffixes[j+1] {
				suffixes[j], suffixes[j+1] = suffixes[j+1], suffixes[j]
			}
		}
	}
	
	// Find longest common prefix between adjacent suffixes
	maxLength := 0
	result := ""
	
	for i := 0; i < len(suffixes)-1; i++ {
		lcp := longestCommonPrefix(suffixes[i], suffixes[i+1])
		if len(lcp) > maxLength {
			maxLength = len(lcp)
			result = lcp
		}
	}
	
	return result
}

func longestCommonPrefix(a, b string) string {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	
	i := 0
	for i < minLen && a[i] == b[i] {
		i++
	}
	
	return a[:i]
}

func main() {
	// Test cases
	fmt.Println("=== Testing Longest Duplicate Substring ===")
	
	testCases := []struct {
		s          string
		description string
	}{
		{"banana", "Standard case"},
		{"abcd", "No duplicates"},
		{"aaaaa", "All same characters"},
		{"abcabcabc", "Repeated pattern"},
		{"", "Empty string"},
		{"a", "Single character"},
		{"abababab", "Alternating pattern"},
		{"abcdeabcdf", "Partial match"},
		{"aaaaaaaaaaaaaaaaaaaaaaaaa", "Long same characters"},
		{"abcdefghijklmnopqrstuvwxyza", "Long with single repeat"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  String: %s\n", tc.s)
		
		result1 := longestDupSubstring(tc.s)
		result2 := longestDupSubstringDoubleHash(tc.s)
		result3 := longestDupSubstringRabinKarp(tc.s)
		result4 := longestDupSubstringBruteForce(tc.s)
		result5 := longestDupSubstringSuffixArray(tc.s)
		
		fmt.Printf("  Standard rolling hash: %s\n", result1)
		fmt.Printf("  Double hash: %s\n", result2)
		fmt.Printf("  Rabin-Karp: %s\n", result3)
		fmt.Printf("  Brute force: %s\n", result4)
		fmt.Printf("  Suffix array: %s\n\n", result5)
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	
	longString := ""
	for i := 0; i < 1000; i++ {
		longString += "abcdefghijklmnopqrstuvwxyz"
	}
	
	fmt.Printf("Testing with string length: %d\n", len(longString))
	
	// Test rolling hash approach
	result := longestDupSubstring(longString)
	fmt.Printf("Rolling hash result length: %d\n", len(result))
	
	// Test edge cases
	fmt.Println("\n=== Edge Case Testing ===")
	
	// Test with very long repeated pattern
	repeatPattern := ""
	for i := 0; i < 100; i++ {
		repeatPattern += "abc"
	}
	
	fmt.Printf("Long repeated pattern: %s\n", longestDupSubstring(repeatPattern))
	
	// Test with no duplicates
	noDup := "abcdefghijklmnopqrstuvwxyz"
	fmt.Printf("No duplicates: %s\n", longestDupSubstring(noDup))
	
	// Test with overlapping duplicates
	overlap := "aaaaabaaaaa"
	fmt.Printf("Overlapping duplicates: %s\n", longestDupSubstring(overlap))
	
	// Test complexity analysis
	fmt.Println("\n=== Complexity Analysis ===")
	
	// Test different string sizes
	sizes := []int{10, 50, 100, 200}
	for _, size := range sizes {
		testString := ""
		for i := 0; i < size; i++ {
			testString += string(rune('a' + (i % 26)))
		}
		
		result = longestDupSubstring(testString)
		fmt.Printf("Size %d: result length %d\n", size, len(result))
	}
}
