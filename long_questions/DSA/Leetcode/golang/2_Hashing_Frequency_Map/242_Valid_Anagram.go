package main

import "fmt"

// 242. Valid Anagram
// Time: O(N), Space: O(1) for ASCII characters (26 for lowercase letters)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	// Assuming only lowercase English letters
	count := make([]int, 26)
	
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	
	return true
}

// Alternative solution using frequency map for general characters
func isAnagramGeneral(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	count := make(map[rune]int)
	
	for _, char := range s {
		count[char]++
	}
	
	for _, char := range t {
		count[char]--
		if count[char] < 0 {
			return false
		}
	}
	
	return true
}

func main() {
	// Test cases
	testCases := []struct {
		s string
		t string
	}{
		{"anagram", "nagaram"},
		{"rat", "car"},
		{"a", "a"},
		{"ab", "ba"},
		{"", ""},
		{"abc", "ab"},
		{"Hello", "olleH"},
	}
	
	for i, tc := range testCases {
		result := isAnagram(tc.s, tc.t)
		resultGeneral := isAnagramGeneral(tc.s, tc.t)
		fmt.Printf("Test Case %d: \"%s\" & \"%s\" -> Anagram: %t (General: %t)\n", 
			i+1, tc.s, tc.t, result, resultGeneral)
	}
}
