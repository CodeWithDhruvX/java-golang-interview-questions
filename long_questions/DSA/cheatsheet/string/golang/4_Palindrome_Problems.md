```go
package main

// 1. Count Palindromic Substrings
// Time: O(N^2), Space: O(1)
func CountSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		// Odd length palindromes
		count += countExpand(s, i, i)
		// Even length palindromes
		count += countExpand(s, i, i+1)
	}
	return count
}

func countExpand(s string, left, right int) int {
	res := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		res++
		left--
		right++
	}
	return res
}

// Brute Force Count Palindromic Substrings
// Time: O(N^3), Space: O(1)
func CountSubstringsBruteForce(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindromeBruteForce(s, i, j) {
				count++
			}
		}
	}
	return count
}

func isPalindromeBruteForce(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 2. Valid Palindrome II (One Deletion)
// Time: O(N), Space: O(1)
func ValidPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalindromeRange(s, left+1, right) || isPalindromeRange(s, left, right-1)
		}
		left++
		right--
	}
	return true
}

func isPalindromeRange(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Brute Force Valid Palindrome II
// Time: O(N^2), Space: O(1)
func ValidPalindromeBruteForce(s string) bool {
	// First check if it's already a palindrome
	if isPalindromeRange(s, 0, len(s)-1) {
		return true
	}
	
	// Try removing each character
	for i := 0; i < len(s); i++ {
		// Create string without character at position i
		newStr := s[:i] + s[i+1:]
		if isPalindromeRange(newStr, 0, len(newStr)-1) {
			return true
		}
	}
	return false
}
```
