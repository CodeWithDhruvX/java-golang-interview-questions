package main

import (
	"fmt"
	"math"
)

// 459. Repeated Substring Pattern - KMP Algorithm
// Time: O(N), Space: O(N)
func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}
	
	// Build LPS array
	lps := buildLPS(s)
	
	// Check if the string can be constructed by repeating a substring
	lpsValue := lps[len(s)-1]
	length := len(s)
	
	// If there's a proper prefix which is also suffix
	// and the length of the string is divisible by (length - lpsValue)
	if lpsValue > 0 && length%(length-lpsValue) == 0 {
		return true
	}
	
	return false
}

// Build LPS array
func buildLPS(s string) []int {
	lps := make([]int, len(s))
	length := 0
	
	i := 1
	for i < len(s) {
		if s[i] == s[length] {
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

// Alternative approach using string concatenation
func repeatedSubstringPatternConcat(s string) bool {
	if len(s) <= 1 {
		return false
	}
	
	// Concatenate string with itself and check if original appears in middle
	concatenated := s + s
	
	// Search for original string starting from position 1
	for i := 1; i < len(s); i++ {
		if i+len(s) <= len(concatenated) && concatenated[i:i+len(s)] == s {
			return true
		}
	}
	
	return false
}

// Optimized concatenation approach
func repeatedSubstringPatternConcatOptimized(s string) bool {
	if len(s) <= 1 {
		return false
	}
	
	concatenated := s + s
	
	// Check if original string appears in concatenated string (excluding first and last positions)
	for i := 1; i < len(s); i++ {
		if concatenated[i:i+len(s)] == s {
			return true
		}
	}
	
	return false
}

// Brute force approach
func repeatedSubstringPatternBruteForce(s string) bool {
	n := len(s)
	
	// Try all possible substring lengths
	for length := 1; length <= n/2; length++ {
		if n%length != 0 {
			continue // Skip if length doesn't divide n
		}
		
		// Check if repeating pattern works
		pattern := s[0:length]
		valid := true
		
		for i := length; i < n; i += length {
			if s[i:i+length] != pattern {
				valid = false
				break
			}
		}
		
		if valid {
			return true
		}
	}
	
	return false
}

// Mathematical approach using divisors
func repeatedSubstringPatternMath(s string) bool {
	n := len(s)
	
	// Find all divisors of n (except n itself)
	for length := 1; length <= n/2; length++ {
		if n%length == 0 {
			// Check if this length forms a repeating pattern
			pattern := s[0:length]
			valid := true
			
			for i := length; i < n; i += length {
				if s[i:i+length] != pattern {
					valid = false
					break
				}
			}
			
			if valid {
				return true
			}
		}
	}
	
	return false
}

// Rolling hash approach
func repeatedSubstringPatternRollingHash(s string) bool {
	n := len(s)
	
	if n <= 1 {
		return false
	}
	
	// Try all possible substring lengths
	for length := 1; length <= n/2; length++ {
		if n%length != 0 {
			continue
		}
		
		// Calculate hash of first substring
		hash := 0
		base := 26
		mod := 1000000007
		
		for i := 0; i < length; i++ {
			hash = (hash*base + int(s[i]-'a')) % mod
		}
		
		// Check if all substrings have same hash
		valid := true
		currentHash := hash
		
		for i := length; i < n; i += length {
			// Calculate hash of current substring
			subHash := 0
			for j := i; j < i+length; j++ {
				subHash = (subHash*base + int(s[j]-'a')) % mod
			}
			
			if subHash != currentHash {
				valid = false
				break
			}
		}
		
		if valid {
			// Verify to avoid hash collision
			pattern := s[0:length]
			for i := length; i < n; i += length {
				if s[i:i+length] != pattern {
					valid = false
					break
				}
			}
			
			if valid {
				return true
			}
		}
	}
	
	return false
}

// KMP with detailed explanation
func repeatedSubstringPatternKMPDetailed(s string) (bool, []string) {
	var explanation []string
	
	if len(s) <= 1 {
		explanation = append(explanation, "String length <= 1, cannot have repeated pattern")
		return false, explanation
	}
	
	explanation = append(explanation, fmt.Sprintf("Building LPS array for: %s", s))
	
	lps := buildLPSWithExplanation(s, &explanation)
	explanation = append(explanation, fmt.Sprintf("LPS array: %v", lps))
	
	lpsValue := lps[len(s)-1]
	length := len(s)
	
	explanation = append(explanation, fmt.Sprintf("LPS[-1] = %d, Length = %d", lpsValue, length))
	
	if lpsValue > 0 {
		explanation = append(explanation, "There is a proper prefix which is also suffix")
		
		if length%(length-lpsValue) == 0 {
			explanation = append(explanation, fmt.Sprintf("Length %% (Length - LPS[-1]) = %d %% %d = 0", length, length-lpsValue))
			explanation = append(explanation, "String can be constructed by repeating a substring")
			return true, explanation
		} else {
			explanation = append(explanation, fmt.Sprintf("Length %% (Length - LPS[-1]) = %d %% %d != 0", length, length-lpsValue))
			explanation = append(explanation, "Cannot construct by repeating a substring")
		}
	} else {
		explanation = append(explanation, "No proper prefix which is also suffix")
	}
	
	return false, explanation
}

func buildLPSWithExplanation(s string, explanation *[]string) []int {
	lps := make([]int, len(s))
	length := 0
	
	i := 1
	for i < len(s) {
		*explanation = append(*explanation, fmt.Sprintf("i=%d, length=%d, s[i]='%c', s[length]='%c'", 
			i, length, s[i], s[length]))
		
		if s[i] == s[length] {
			length++
			lps[i] = length
			*explanation = append(*explanation, fmt.Sprintf("Match! LPS[%d]=%d", i, length))
			i++
		} else {
			if length != 0 {
				*explanation = append(*explanation, fmt.Sprintf("Mismatch, using LPS[%d]=%d", length-1, lps[length-1]))
				length = lps[length-1]
			} else {
				lps[i] = 0
				*explanation = append(*explanation, fmt.Sprintf("No match, LPS[%d]=0", i))
				i++
			}
		}
	}
	
	return lps
}

// Find the actual repeating substring
func findRepeatingSubstring(s string) string {
	n := len(s)
	
	if n <= 1 {
		return ""
	}
	
	lps := buildLPS(s)
	lpsValue := lps[len(s)-1]
	
	if lpsValue > 0 && n%(n-lpsValue) == 0 {
		return s[0 : n-lpsValue]
	}
	
	return ""
}

func main() {
	// Test cases
	fmt.Println("=== Testing Repeated Substring Pattern ===")
	
	testCases := []struct {
		s          string
		expected   bool
		description string
	}{
		{"abab", true, "Simple repeating pattern"},
		{"aba", false, "Not a repeating pattern"},
		{"abcabcabcabc", true, "Long repeating pattern"},
		{"aaaa", true, "All same characters"},
		{"abcabc", true, "Two repetitions"},
		{"abac", false, "No pattern"},
		{"a", false, "Single character"},
		{"", false, "Empty string"},
		{"xyzxyzxyz", true, "Three repetitions"},
		{"abcd", false, "All unique"},
		{"zzzzzz", true, "Many same characters"},
		{"abcab", false, "Partial match"},
		{"abcabcabcabcabc", true, "Five repetitions"},
		{"ababa", false, "Overlapping but not repeating"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  String: '%s'\n", tc.s)
		
		result1 := repeatedSubstringPattern(tc.s)
		result2 := repeatedSubstringPatternConcat(tc.s)
		result3 := repeatedSubstringPatternBruteForce(tc.s)
		result4 := repeatedSubstringPatternMath(tc.s)
		result5 := repeatedSubstringPatternRollingHash(tc.s)
		
		fmt.Printf("  KMP: %t\n", result1)
		fmt.Printf("  Concatenation: %t\n", result2)
		fmt.Printf("  Brute Force: %t\n", result3)
		fmt.Printf("  Mathematical: %t\n", result4)
		fmt.Printf("  Rolling Hash: %t\n", result5)
		
		// Check if results match expected
		if result1 == tc.expected {
			fmt.Printf("  ✓ Expected: %t\n", tc.expected)
		} else {
			fmt.Printf("  ✗ Expected: %t, Got: %t\n", tc.expected, result1)
		}
		
		// Show repeating substring if found
		if result1 {
			pattern := findRepeatingSubstring(tc.s)
			fmt.Printf("  Repeating pattern: '%s'\n", pattern)
		}
		
		fmt.Println()
	}
	
	// KMP detailed explanation
	fmt.Println("=== KMP Detailed Explanation ===")
	_, explanation := repeatedSubstringPatternKMPDetailed("abcabcabc")
	for _, step := range explanation {
		fmt.Printf("  %s\n", step)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	longString := ""
	for i := 0; i < 1000; i++ {
		longString += "abc"
	}
	
	fmt.Printf("Long string length: %d\n", len(longString))
	fmt.Printf("KMP result: %t\n", repeatedSubstringPattern(longString))
	fmt.Printf("Concatenation result: %t\n", repeatedSubstringPatternConcat(longString))
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Very long string with no pattern
	noPattern := ""
	for i := 0; i < 1000; i++ {
		noPattern += string(rune('a' + (i % 26)))
	}
	fmt.Printf("No pattern string: %t\n", repeatedSubstringPattern(noPattern))
	
	// All same characters
	allSame := ""
	for i := 0; i < 10000; i++ {
		allSame += "a"
	}
	fmt.Printf("All same characters: %t\n", repeatedSubstringPattern(allSame))
	
	// Complex pattern
	complex := "ababcababcababc"
	fmt.Printf("Complex pattern: %t\n", repeatedSubstringPattern(complex))
	
	// Find all repeating substrings
	fmt.Println("\n=== Finding Repeating Substrings ===")
	testStrings := []string{"abcabcabc", "ababab", "aaaaaa", "xyzxyzxyzxyz"}
	
	for _, s := range testStrings {
		pattern := findRepeatingSubstring(s)
		if pattern != "" {
			fmt.Printf("'%s' -> pattern: '%s'\n", s, pattern)
		} else {
			fmt.Printf("'%s' -> no repeating pattern\n", s)
		}
	}
}
