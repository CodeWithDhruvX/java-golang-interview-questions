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

// 3. Rotate String
// Time: O(N), Space: O(N)
func RotateString(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	doubled := s + s
	return strings.Contains(doubled, goal)
}
```
