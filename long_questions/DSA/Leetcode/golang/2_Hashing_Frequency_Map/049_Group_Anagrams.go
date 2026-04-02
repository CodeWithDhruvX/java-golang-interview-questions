package main

import (
	"fmt"
	"sort"
)

// 49. Group Anagrams
// Time: O(N*K*logK), Space: O(N*K)
func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	
	for _, str := range strs {
		// Sort the string to create a key
		sortedStr := sortString(str)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}
	
	// Convert map values to slice
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}
	
	return result
}

// Helper function to sort a string
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// Alternative solution using character count as key (more efficient)
func groupAnagramsOptimized(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	
	for _, str := range strs {
		// Create key based on character count
		key := createCountKey(str)
		anagramMap[key] = append(anagramMap[key], str)
	}
	
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}
	
	return result
}

// Create key based on character count (26 lowercase letters)
func createCountKey(s string) string {
	count := make([]int, 26)
	for _, char := range s {
		count[char-'a']++
	}
	
	key := fmt.Sprintf("%v", count)
	return key
}

func main() {
	// Test cases
	testCases := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
		{""},
		{"a"},
		{"abc", "bca", "cab", "def", "fed", "ghi"},
		{"", "", ""},
		{"a", "b", "c", "a", "b", "c"},
	}
	
	for i, strs := range testCases {
		result := groupAnagrams(strs)
		fmt.Printf("Test Case %d: %v\n", i+1, strs)
		fmt.Printf("Grouped Anagrams: %v\n\n", result)
	}
}
