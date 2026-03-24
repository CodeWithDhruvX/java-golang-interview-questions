```go
package main

import (
	"strconv"
	"strings"
)

// 1. Reverse Words in a Sentence
// Time: O(N), Space: O(N)
func ReverseWords(s string) string {
	parts := strings.Fields(s)
	// Reverse parts
	l, r := 0, len(parts)-1
	for l < r {
		parts[l], parts[r] = parts[r], parts[l]
		l++
		r--
	}
	return strings.Join(parts, " ")
}

// Brute Force Reverse Words in a Sentence
// Time: O(N^2), Space: O(N)
func ReverseWordsBruteForce(s string) string {
	words := []string{}
	currentWord := ""
	
	// Extract words manually
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			currentWord += string(s[i])
		} else if currentWord != "" {
			words = append(words, currentWord)
			currentWord = ""
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}
	
	// Reverse words manually
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	
	// Join manually
	result := ""
	for i := 0; i < len(words); i++ {
		if i > 0 {
			result += " "
		}
		result += words[i]
	}
	return result
}

// 2. String Compression
// Time: O(N), Space: O(1) (in-place modification of byte slice)
func Compress(chars []byte) int {
	write := 0
	anchor := 0
	
	for read := 0; read < len(chars); read++ {
		if read+1 == len(chars) || chars[read+1] != chars[read] {
			chars[write] = chars[anchor]
			write++
			if read > anchor {
				count := strconv.Itoa(read - anchor + 1)
				for i := 0; i < len(count); i++ {
					chars[write] = count[i]
					write++
				}
			}
			anchor = read + 1
		}
	}
	return write
}

// Brute Force String Compression
// Time: O(N^2), Space: O(N)
func CompressBruteForce(chars []byte) []byte {
	if len(chars) == 0 {
		return []byte{}
	}
	
	result := []byte{}
	count := 1
	
	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			count++
		} else {
			result = append(result, chars[i-1])
			if count > 1 {
				countStr := strconv.Itoa(count)
				for j := 0; j < len(countStr); j++ {
					result = append(result, countStr[j])
				}
			}
			count = 1
		}
	}
	
	// Add last character
	result = append(result, chars[len(chars)-1])
	if count > 1 {
		countStr := strconv.Itoa(count)
		for j := 0; j < len(countStr); j++ {
			result = append(result, countStr[j])
		}
	}
	
	return result
}

// 3. Rotate String
// Time: O(N), Space: O(N)
func RotateString(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	doubled := s + s
	return strings.Contains(doubled, goal)
}

// Brute Force Rotate String
// Time: O(N^2), Space: O(1)
func RotateStringBruteForce(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if len(s) == 0 {
		return true
	}
	
	// Try all possible rotations
	for i := 0; i < len(s); i++ {
		rotated := s[i:] + s[:i]
		if rotated == goal {
			return true
		}
	}
	return false
}
```
