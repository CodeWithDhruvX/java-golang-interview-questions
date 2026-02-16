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
```
