package main

import "fmt"

// 438. Find All Anagrams in a String (Fixed Size Sliding Window)
// Time: O(N), Space: O(1) for ASCII characters
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	
	result := []int{}
	pCount := make([]int, 26)
	sCount := make([]int, 26)
	
	// Initialize frequency count for pattern and first window
	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}
	
	// Check if first window is an anagram
	if matches(pCount, sCount) {
		result = append(result, 0)
	}
	
	// Slide the window through the string
	for i := len(p); i < len(s); i++ {
		// Remove the leftmost character
		sCount[s[i-len(p)]-'a']--
		// Add the new character
		sCount[s[i]-'a']++
		
		// Check if current window is an anagram
		if matches(pCount, sCount) {
			result = append(result, i-len(p)+1)
		}
	}
	
	return result
}

// Helper function to check if two frequency arrays match
func matches(pCount, sCount []int) bool {
	for i := 0; i < 26; i++ {
		if pCount[i] != sCount[i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	testCases := []struct {
		s string
		p string
	}{
		{"cbaebabacd", "abc"},
		{"abab", "ab"},
		{"aaaaaaaaaa", "aaaa"},
		{"abacbabc", "abc"},
		{"", "a"},
		{"a", ""},
		{"abc", "def"},
		{"abababab", "ab"},
	}
	
	for i, tc := range testCases {
		result := findAnagrams(tc.s, tc.p)
		fmt.Printf("Test Case %d: s=\"%s\", p=\"%s\" -> Anagram indices: %v\n", 
			i+1, tc.s, tc.p, result)
	}
}
