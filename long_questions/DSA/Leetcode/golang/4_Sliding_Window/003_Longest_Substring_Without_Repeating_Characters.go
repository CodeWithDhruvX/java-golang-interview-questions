package main

import "fmt"

// 3. Longest Substring Without Repeating Characters (Variable Size Sliding Window)
// Time: O(N), Space: O(min(N, M)) where M is the size of character set
func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	left := 0
	maxLength := 0
	
	for right := 0; right < len(s); right++ {
		// If character is already in the window, move left pointer
		if index, exists := charMap[s[right]]; exists && index >= left {
			left = index + 1
		}
		
		// Update the character's last seen position
		charMap[s[right]] = right
		
		// Update max length
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	
	return maxLength
}

func main() {
	// Test cases
	testCases := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"",
		"a",
		"au",
		"dvdf",
		"abba",
		"tmmzuxt",
		"abcdefghijklmnopqrstuvwxyz",
		"abccba",
	}
	
	for i, s := range testCases {
		result := lengthOfLongestSubstring(s)
		fmt.Printf("Test Case %d: \"%s\" -> Length of longest substring: %d\n", 
			i+1, s, result)
	}
}
