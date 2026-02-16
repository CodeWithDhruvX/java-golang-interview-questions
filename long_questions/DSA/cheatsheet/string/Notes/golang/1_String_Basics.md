```go
package main

import (
	"strings"
	"unicode"
)

// 1. Reverse a String
// Time: O(N), Space: O(N)
func ReverseString(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}

// 2. Check if String is Palindrome
// Time: O(N), Space: O(N) (due to rune conversion)
func IsPalindrome(s string) bool {
	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 3. Count Vowels and Consonants
// Time: O(N), Space: O(1)
func CountVowelsConsonants(s string) (int, int) {
	vowels := 0
	consonants := 0
	for _, char := range s {
		if unicode.IsLetter(char) {
			lowerChar := unicode.ToLower(char)
			if strings.ContainsRune("aeiou", lowerChar) {
				vowels++
			} else {
				consonants++
			}
		}
	}
	return vowels, consonants
}
```
