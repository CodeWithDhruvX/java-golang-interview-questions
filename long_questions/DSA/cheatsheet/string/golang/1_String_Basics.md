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

// Brute Force Reverse String
// Time: O(N^2), Space: O(N)
func ReverseStringBruteForce(s string) string {
	runes := []rune(s)
	result := make([]rune, len(runes))
	for i := 0; i < len(runes); i++ {
		result[len(runes)-1-i] = runes[i]
	}
	return string(result)
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

// Brute Force Check Palindrome
// Time: O(N^2), Space: O(N)
func IsPalindromeBruteForce(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := len(runes) - 1; j > i; j-- {
			if runes[i] != runes[j] {
				return false
			}
		}
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

// Brute Force Count Vowels and Consonants
// Time: O(N^2), Space: O(1)
func CountVowelsConsonantsBruteForce(s string) (int, int) {
	vowels := 0
	consonants := 0
	vowelSet := "aeiouAEIOU"
	
	for i, char := range s {
		if unicode.IsLetter(char) {
			isVowel := false
			for j := 0; j < len(vowelSet); j++ {
				if char == rune(vowelSet[j]) {
					isVowel = true
					break
				}
			}
			if isVowel {
				vowels++
			} else {
				consonants++
			}
		}
	}
	return vowels, consonants
}
```
