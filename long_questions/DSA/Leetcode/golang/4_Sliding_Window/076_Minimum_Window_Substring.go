package main

import "fmt"

// 76. Minimum Window Substring (Variable Size Sliding Window)
// Time: O(N + M), Space: O(1) for ASCII characters
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	
	// Frequency map for characters in t
	tCount := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tCount[t[i]]++
	}
	
	// Number of unique characters in t that need to be present in the window
	required := len(tCount)
	
	// Left and right pointers
	left, right := 0, 0
	
	// Keep track of how many unique characters in the current window match the required count
	formed := 0
	
	// Frequency map for the current window
	windowCount := make(map[byte]int)
	
	// Result variables: (window length, left, right)
	result := []int{-1, 0, 0}
	
	for right < len(s) {
		char := s[right]
		windowCount[char]++
		
		// If the frequency of the current character matches exactly the required frequency in t
		if count, exists := tCount[char]; exists && windowCount[char] == count {
			formed++
		}
		
		// Try to contract the window till the point it ceases to be 'desirable'
		for left <= right && formed == required {
			// Save the smallest window
			windowSize := right - left + 1
			if result[0] == -1 || windowSize < result[0] {
				result[0] = windowSize
				result[1] = left
				result[2] = right
			}
			
			// The character at the position left is no longer a part of the window
			leftChar := s[left]
			windowCount[leftChar]--
			if count, exists := tCount[leftChar]; exists && windowCount[leftChar] < count {
				formed--
			}
			
			left++
		}
		
		right++
	}
	
	if result[0] == -1 {
		return ""
	}
	
	return s[result[1] : result[2]+1]
}

func main() {
	// Test cases
	testCases := []struct {
		s string
		t string
	}{
		{"ADOBECODEBANC", "ABC"},
		{"a", "a"},
		{"a", "aa"},
		{"ab", "b"},
		{"ab", "c"},
		{"abc", "abc"},
		{"aa", "aa"},
		{"ab", "ab"},
		{"bba", "ab"},
		{"aaaaaaaaaaaabbbbbcdd", "abcdd"},
	}
	
	for i, tc := range testCases {
		result := minWindow(tc.s, tc.t)
		fmt.Printf("Test Case %d: s=\"%s\", t=\"%s\" -> Min window: \"%s\"\n", 
			i+1, tc.s, tc.t, result)
	}
}
