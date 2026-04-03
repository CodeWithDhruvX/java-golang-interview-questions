package main

import (
	"fmt"
)

// 214. Shortest Palindrome - Z Algorithm
// Time: O(N), Space: O(N)
func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	// Create pattern: s + '#' + reverse(s)
	pattern := s + "#" + reverseString(s)
	
	// Compute Z-array
	zArray := computeZArray(pattern)
	
	// Find longest palindrome prefix
	maxPrefix := 0
	for i := len(s) + 1; i < len(zArray); i++ {
		if zArray[i] > maxPrefix {
			maxPrefix = zArray[i]
		}
	}
	
	// Add reverse of suffix to front
	suffix := s[maxPrefix:]
	return reverseString(suffix) + s
}

func computeZArray(s string) []int {
	n := len(s)
	z := make([]int, n)
	z[0] = n
	
	l, r := 0, 0
	
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}
	
	return z
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Z Algorithm for pattern matching
func zAlgorithmPatternSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}
	
	// Create pattern: pattern + '$' + text
	combined := pattern + "$" + text
	z := computeZArray(combined)
	
	// Find pattern occurrences
	occurrences := []int{}
	patternLen := len(pattern)
	
	for i := patternLen + 1; i < len(z); i++ {
		if z[i] == patternLen {
			occurrences = append(occurrences, i-(patternLen+1))
		}
	}
	
	return occurrences
}

// Z Algorithm for longest common prefix
func longestCommonPrefix(s1, s2 string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	
	// Create pattern: s1 + '#' + s2
	combined := s1 + "#" + s2
	z := computeZArray(combined)
	
	// Find longest prefix
	maxPrefix := 0
	for i := len(s1) + 1; i < len(z); i++ {
		if z[i] > maxPrefix {
			maxPrefix = z[i]
		}
	}
	
	return maxPrefix
}

// Z Algorithm for string compression
func zAlgorithmCompression(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	z := computeZArray(s)
	
	// Find smallest pattern that can generate the string
	for i := 1; i < len(z); i++ {
		if z[i] == len(s)-i && len(s)%i == 0 {
			pattern := s[:i]
			repeated := ""
			for j := 0; j < len(s)/i; j++ {
				repeated += pattern
			}
			if repeated == s {
				return fmt.Sprintf("%s{%d}", pattern, len(s)/i)
			}
		}
	}
	
	return s
}

// Z Algorithm for string rotation
func zAlgorithmRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	
	// Create pattern: s2 + '$' + s1 + s1
	combined := s2 + "$" + s1 + s1
	z := computeZArray(combined)
	
	// Check if s1 is rotation of s2
	for i := len(s2) + 1; i < len(z); i++ {
		if z[i] == len(s1) {
			return true
		}
	}
	
	return false
}

// Z Algorithm for palindrome detection
func zAlgorithmPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	
	// Create pattern: s + '#' + reverse(s)
	pattern := s + "#" + reverseString(s)
	z := computeZArray(pattern)
	
	// Check if entire string is palindrome
	for i := len(s) + 1; i < len(z); i++ {
		if z[i] == len(s) {
			return true
		}
	}
	
	return false
}

// Z Algorithm for substring search with wildcards
func zAlgorithmWildcardSearch(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}
	
	// Replace wildcards with unique characters
	wildcardPattern := make([]byte, len(pattern))
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '?' {
			wildcardPattern[i] = 0 // Use 0 as wildcard
		} else {
			wildcardPattern[i] = pattern[i]
		}
	}
	
	// Create combined pattern
	combined := string(wildcardPattern) + "$" + text
	z := computeZArray(combined)
	
	// Find matches
	occurrences := []int{}
	patternLen := len(pattern)
	
	for i := patternLen + 1; i < len(z); i++ {
		if z[i] >= patternLen {
			occurrences = append(occurrences, i-(patternLen+1))
		}
	}
	
	return occurrences
}

// Z Algorithm for multiple pattern search
func zAlgorithmMultiplePatterns(text string, patterns []string) map[string][]int {
	results := make(map[string][]int)
	
	for _, pattern := range patterns {
		occurrences := zAlgorithmPatternSearch(text, pattern)
		results[pattern] = occurrences
	}
	
	return results
}

// Z Algorithm for string similarity
func zAlgorithmSimilarity(s1, s2 string) float64 {
	if len(s1) == 0 && len(s2) == 0 {
		return 1.0
	}
	
	if len(s1) == 0 || len(s2) == 0 {
		return 0.0
	}
	
	// Find longest common prefix
	lcp := longestCommonPrefix(s1, s2)
	
	// Calculate similarity as ratio of LCP to shorter string
	shorter := len(s1)
	if len(s2) < shorter {
		shorter = len(s2)
	}
	
	return float64(lcp) / float64(shorter)
}

// Z Algorithm for string period
func zAlgorithmPeriod(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	
	z := computeZArray(s)
	
	// Find smallest period
	for i := 1; i < len(z); i++ {
		if z[i] == len(s)-i && len(s)%i == 0 {
			return i
		}
	}
	
	return len(s)
}

// Z Algorithm for border detection
func zAlgorithmBorders(s string) []int {
	if len(s) <= 1 {
		return []int{}
	}
	
	z := computeZArray(s)
	borders := []int{}
	
	// Find all borders (proper prefixes that are also suffixes)
	for i := 1; i < len(z); i++ {
		if z[i] == len(s)-i {
			borders = append(borders, i)
		}
	}
	
	return borders
}

func main() {
	// Test cases
	fmt.Println("=== Testing Z Algorithm ===")
	
	testCases := []struct {
		s          string
		description string
	}{
		{"aacecaaa", "Standard case"},
		{"abcd", "No palindrome prefix"},
		{"a", "Single character"},
		{"", "Empty string"},
		{"racecar", "Palindrome"},
		{"abacdfgdcaba", "Multiple palindrome prefixes"},
		{"aaaaa", "All same characters"},
		{"abcba", "Odd length palindrome"},
		{"abccba", "Even length palindrome"},
		{"xyz", "No palindrome"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s (\"%s\")\n", i+1, tc.description, tc.s)
		
		result := shortestPalindrome(tc.s)
		fmt.Printf("  Shortest palindrome: %s\n", result)
		
		fmt.Println()
	}
	
	// Test pattern matching
	fmt.Println("=== Pattern Matching Test ===")
	text := "ababcabcababc"
	pattern := "abc"
	occurrences := zAlgorithmPatternSearch(text, pattern)
	fmt.Printf("Text: %s, Pattern: %s\n", text, pattern)
	fmt.Printf("Occurrences: %v\n", occurrences)
	
	// Test longest common prefix
	fmt.Println("\n=== Longest Common Prefix Test ===")
	lcpTests := [][]string{
		{"abcdef", "abcxyz"},
		{"hello", "world"},
		{"same", "same"},
		{"prefix", "pre"},
	}
	
	for _, test := range lcpTests {
		lcp := longestCommonPrefix(test[0], test[1])
		fmt.Printf("LCP of \"%s\" and \"%s\": %d\n", test[0], test[1], lcp)
	}
	
	// Test string compression
	fmt.Println("\n=== String Compression Test ===")
	compressionTests := []string{
		"abababab",
		"abcabcabc",
		"aaaaaa",
		"abcdef",
		"xyzxyzxyz",
	}
	
	for _, test := range compressionTests {
		compressed := zAlgorithmCompression(test)
		fmt.Printf("Original: %s, Compressed: %s\n", test, compressed)
	}
	
	// Test string rotation
	fmt.Println("\n=== String Rotation Test ===")
	rotationTests := [][2]string{
		{"abcde", "cdeab"},
		{"hello", "llohe"},
		{"test", "sett"},
		{"abc", "def"},
	}
	
	for _, test := range rotationTests {
		isRotation := zAlgorithmRotation(test[0], test[1])
		fmt.Printf("Is \"%s\" rotation of \"%s\": %t\n", test[1], test[0], isRotation)
	}
	
	// Test palindrome detection
	fmt.Println("\n=== Palindrome Detection Test ===")
	palindromeTests := []string{
		"racecar",
		"hello",
		"madam",
		"test",
		"level",
	}
	
	for _, test := range palindromeTests {
		isPalindrome := zAlgorithmPalindrome(test)
		fmt.Printf("Is \"%s\" palindrome: %t\n", test, isPalindrome)
	}
	
	// Test wildcard search
	fmt.Println("\n=== Wildcard Search Test ===")
	wildcardTests := []struct {
		text    string
		pattern string
	}{
		{"abcde", "a?cde"},
		{"hello", "h?llo"},
		{"test", "t??t"},
		{"world", "?????"},
	}
	
	for _, test := range wildcardTests {
		wildcardOccurrences := zAlgorithmWildcardSearch(test.text, test.pattern)
		fmt.Printf("Text: %s, Pattern: %s, Occurrences: %v\n", 
			test.text, test.pattern, wildcardOccurrences)
	}
	
	// Test multiple patterns
	fmt.Println("\n=== Multiple Patterns Test ===")
	text := "ababcabcababc"
	patterns := []string{"abc", "ab", "bc"}
	multipleResults := zAlgorithmMultiplePatterns(text, patterns)
	
	fmt.Printf("Text: %s\n", text)
	for pattern, occurrences := range multipleResults {
		fmt.Printf("Pattern \"%s\": %v\n", pattern, occurrences)
	}
	
	// Test string similarity
	fmt.Println("\n=== String Similarity Test ===")
	similarityTests := [][2]string{
		{"hello", "help"},
		{"test", "testing"},
		{"abc", "abc"},
		{"xyz", "abc"},
	}
	
	for _, test := range similarityTests {
		similarity := zAlgorithmSimilarity(test[0], test[1])
		fmt.Printf("Similarity between \"%s\" and \"%s\": %.2f\n", 
			test[0], test[1], similarity)
	}
	
	// Test string period
	fmt.Println("\n=== String Period Test ===")
	periodTests := []string{
		"abababab",
		"abcabcabc",
		"aaaaaa",
		"abcdef",
		"xyzxyzxyz",
	}
	
	for _, test := range periodTests {
		period := zAlgorithmPeriod(test)
		fmt.Printf("Period of \"%s\": %d\n", test, period)
	}
	
	// Test border detection
	fmt.Println("\n=== Border Detection Test ===")
	borderTests := []string{
		"ababcab",
		"aaaaa",
		"abcab",
		"test",
		"ababa",
	}
	
	for _, test := range borderTests {
		borders := zAlgorithmBorders(test)
		fmt.Printf("Borders of \"%s\": %v\n", test, borders)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	// Large string
	largeString := ""
	for i := 0; i < 10000; i++ {
		largeString += string('a' + (i % 26))
	}
	
	fmt.Printf("Large string test with %d characters\n", len(largeString))
	
	start := time.Now()
	result := shortestPalindrome(largeString)
	duration := time.Since(start)
	
	fmt.Printf("Large string result length: %d\n", len(result))
	fmt.Printf("Time taken: %v\n", duration)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Very long palindrome
	longPalindrome := ""
	for i := 0; i < 1000; i++ {
		longPalindrome += string('a' + (i % 26))
	}
	longPalindrome += reverseString(longPalindrome)
	
	fmt.Printf("Long palindrome test: length=%d\n", len(longPalindrome))
	result = shortestPalindrome(longPalindrome)
	fmt.Printf("Result length: %d\n", len(result))
	
	// Empty string
	fmt.Printf("Empty string: %s\n", shortestPalindrome(""))
	
	// Single character
	fmt.Printf("Single character: %s\n", shortestPalindrome("a"))
}
