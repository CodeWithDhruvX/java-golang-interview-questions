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

// Brute Force First Non-Repeating Character
// Time: O(N^2), Space: O(1)
func FirstNonRepeatingCharBruteForce(s string) rune {
	runes := []rune(s)
	for i, char := range runes {
		isUnique := true
		for j := 0; j < len(runes); j++ {
			if i != j && char == runes[j] {
				isUnique = false
				break
			}
		}
		if isUnique {
			return char
		}
	}
	return 0
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

// Brute Force Check Anagrams
// Time: O(N^2), Space: O(1)
func AreAnagramsBruteForce(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	used := make([]bool, len(runes2))
	
	for i := 0; i < len(runes1); i++ {
		found := false
		for j := 0; j < len(runes2); j++ {
			if !used[j] && runes1[i] == runes2[j] {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
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

// Brute Force Find Duplicate Characters
// Time: O(N^2), Space: O(N)
func FindDuplicateCharsBruteForce(s string) []rune {
	runes := []rune(s)
	var duplicates []rune
	added := make(map[rune]bool)
	
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] == runes[j] && !added[runes[i]] {
				duplicates = append(duplicates, runes[i])
				added[runes[i]] = true
				break
			}
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

// Brute Force Most Frequent Character
// Time: O(N^2), Space: O(1)
func MostFrequentCharBruteForce(s string) rune {
	runes := []rune(s)
	if len(runes) == 0 {
		return 0
	}
	
	maxChar := runes[0]
	maxCount := 0
	
	for i := 0; i < len(runes); i++ {
		count := 0
		for j := 0; j < len(runes); j++ {
			if runes[i] == runes[j] {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			maxChar = runes[i]
		}
	}
	return maxChar
}
```
