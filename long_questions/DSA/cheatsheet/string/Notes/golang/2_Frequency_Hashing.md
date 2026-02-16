```go
package main

// 1. First Non-Repeating Character
// Time: O(N), Space: O(1) (fixed alphabet)
func FirstNonRepeatingChar(s string) rune {
	freq := make(map[rune]int)
	// Pass 1: Count frequencies
	for _, char := range s {
		freq[char]++
	}
	// Pass 2: Find first with count 1
	for _, char := range s {
		if freq[char] == 1 {
			return char
		}
	}
	return 0 // or some error indicator
}

// 2. Check if Two Strings are Anagrams
// Time: O(N), Space: O(1)
func AreAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	freq := make(map[rune]int)
	for _, char := range s1 {
		freq[char]++
	}
	for _, char := range s2 {
		freq[char]--
		if freq[char] < 0 {
			return false
		}
	}
	return true
}

// 3. Find All Duplicate Characters
// Time: O(N), Space: O(N)
func FindDuplicateChars(s string) []rune {
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}
	var duplicates []rune
	for char, count := range freq {
		if count > 1 {
			duplicates = append(duplicates, char)
		}
	}
	return duplicates
}

// 4. Most Frequent Character
// Time: O(N), Space: O(1)
func MostFrequentChar(s string) rune {
	freq := make(map[rune]int)
	var maxChar rune
	maxCount := 0
	
	for _, char := range s {
		freq[char]++
		if freq[char] > maxCount {
			maxCount = freq[char]
			maxChar = char
		}
	}
	return maxChar
}
```
