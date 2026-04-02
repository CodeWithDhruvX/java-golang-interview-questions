package main

import (
	"fmt"
	"unicode"
)

// 125. Valid Palindrome
// Time: O(N), Space: O(1)
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	
	for left < right {
		// Skip non-alphanumeric characters
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}
		
		// Compare characters (case-insensitive)
		if left < right && unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		
		left++
		right--
	}
	
	return true
}

// Helper function to check if character is alphanumeric
func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func main() {
	// Test cases
	testCases := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		" ",
		"",
		"madam",
		"Able was I ere I saw Elba",
		"No lemon, no melon",
		"12321",
		"12345",
		".,",
	}
	
	for i, s := range testCases {
		result := isPalindrome(s)
		fmt.Printf("Test Case %d: \"%s\" -> Palindrome: %t\n", i+1, s, result)
	}
}
