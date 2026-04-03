package main

import "fmt"

// 249. Group Shifted Strings - Hashing with String Processing
// Time: O(N * L), Space: O(N * L) where N is number of strings, L is average length
func groupStrings(strings []string) [][]string {
	if len(strings) == 0 {
		return [][]string{}
	}
	
	// Map to store groups by key
	groups := make(map[string][]string)
	
	for _, s := range strings {
		key := getShiftKey(s)
		groups[key] = append(groups[key], s)
	}
	
	// Convert map to slice
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	
	return result
}

// Get shift key for a string
func getShiftKey(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	// Calculate shift from first character
	shift := s[0] - 'a'
	
	key := make([]byte, len(s))
	key[0] = 'a' // Normalize first character to 'a'
	
	for i := 1; i < len(s); i++ {
		// Calculate normalized character
		normalized := (s[i]-'a'-shift+26)%26 + 'a'
		key[i] = normalized
	}
	
	return string(key)
}

// Alternative approach using differences
func groupStringsDifferences(strings []string) [][]string {
	if len(strings) == 0 {
		return [][]string{}
	}
	
	groups := make(map[string][]string)
	
	for _, s := range strings {
		key := getDifferenceKey(s)
		groups[key] = append(groups[key], s)
	}
	
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	
	return result
}

func getDifferenceKey(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	key := make([]byte, len(s)-1)
	
	for i := 1; i < len(s); i++ {
		diff := (s[i] - s[i-1] + 26) % 26
		key[i-1] = diff + 'a' // Convert to character
	}
	
	return string(key)
}

// Optimized version with sorting
func groupStringsOptimized(strings []string) [][]string {
	if len(strings) == 0 {
		return [][]string{}
	}
	
	groups := make(map[string][]string)
	
	for _, s := range strings {
		key := getShiftKeyOptimized(s)
		groups[key] = append(groups[key], s)
	}
	
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	
	return result
}

func getShiftKeyOptimized(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	shift := s[0] - 'a'
	key := make([]byte, len(s))
	
	for i := range s {
		key[i] = (s[i]-'a'-shift+26)%26 + 'a'
	}
	
	return string(key)
}

// Version that handles all characters (not just lowercase)
func groupStringsAllChars(strings []string) [][]string {
	if len(strings) == 0 {
		return [][]string{}
	}
	
	groups := make(map[string][]string)
	
	for _, s := range strings {
		key := getUniversalShiftKey(s)
		groups[key] = append(groups[key], s)
	}
	
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	
	return result
}

func getUniversalShiftKey(s string) string {
	if len(s) <= 1 {
		return s
	}
	
	// Find the minimum character to normalize
	minChar := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < minChar {
			minChar = s[i]
		}
	}
	
	shift := minChar - 'a'
	key := make([]byte, len(s))
	
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			key[i] = (s[i]-'a'-shift+26)%26 + 'a'
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			key[i] = (s[i]-'A'-shift+26)%26 + 'a'
		} else {
			key[i] = s[i] // Keep non-alphabetic characters as is
		}
	}
	
	return string(key)
}

// Brute force approach for comparison
func groupStringsBruteForce(strings []string) [][]string {
	if len(strings) == 0 {
		return [][]string{}
	}
	
	used := make([]bool, len(strings))
	result := [][]string{}
	
	for i := 0; i < len(strings); i++ {
		if used[i] {
			continue
		}
		
		group := []string{strings[i]}
		used[i] = true
		
		// Find all strings that can be shifted to match strings[i]
		for j := i + 1; j < len(strings); j++ {
			if !used[j] && canShift(strings[i], strings[j]) {
				group = append(group, strings[j])
				used[j] = true
			}
		}
		
		result = append(result, group)
	}
	
	return result
}

func canShift(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	
	if len(s1) == 1 {
		return true
	}
	
	shift := (s2[0] - s1[0] + 26) % 26
	
	for i := 1; i < len(s1); i++ {
		if (s2[i]-s1[i]+26)%26 != shift {
			return false
		}
	}
	
	return true
}

// Version with detailed explanation
func groupStringsWithExplanation(strings []string) ([][]string, []string) {
	var explanation []string
	
	if len(strings) == 0 {
		explanation = append(explanation, "Empty input, returning empty result")
		return [][]string{}, explanation
	}
	
	explanation = append(explanation, fmt.Sprintf("Processing %d strings", len(strings)))
	
	groups := make(map[string][]string)
	
	for i, s := range strings {
		explanation = append(explanation, fmt.Sprintf("String %d: '%s'", i, s))
		
		key := getShiftKey(s)
		explanation = append(explanation, fmt.Sprintf("  Shift key: '%s'", key))
		
		groups[key] = append(groups[key], s)
		explanation = append(explanation, fmt.Sprintf("  Added to group with key '%s'", key))
	}
	
	explanation = append(explanation, fmt.Sprintf("Found %d groups", len(groups)))
	
	result := make([][]string, 0, len(groups))
	for key, group := range groups {
		result = append(result, group)
		explanation = append(explanation, fmt.Sprintf("Group '%s': %v", key, group))
	}
	
	return result, explanation
}

// Helper function to check if two strings are in the same group
func areInSameGroup(s1, s2 string) bool {
	return getShiftKey(s1) == getShiftKey(s2)
}

// Find all possible shifts of a string
func getAllShifts(s string) []string {
	if len(s) <= 1 {
		return []string{s}
	}
	
	shifts := make([]string, 26)
	
	for shift := 0; shift < 26; shift++ {
		shifted := make([]byte, len(s))
		
		for i := range s {
			shifted[i] = (s[i]-'a'+byte(shift))%26 + 'a'
		}
		
		shifts[shift] = string(shifted)
	}
	
	return shifts
}

func main() {
	// Test cases
	fmt.Println("=== Testing Group Shifted Strings ===")
	
	testCases := []struct {
		strings    []string
		description string
	}{
		{[]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}, "Standard case"},
		{[]string{"a"}, "Single character"},
		{[]string{"abc", "def", "ghi"}, "All same length"},
		{[]string{"a", "b", "c", "d", "e"}, "All single characters"},
		{[]string{"abc", "bcd", "cde", "def"}, "Sequential shifts"},
		{[]string{"az", "ba", "ab", "bc"}, "Mixed shifts"},
		{[]string{}, "Empty array"},
		{[]string{"abc", "abc", "abc"}, "Duplicate strings"},
		{[]string{"abcdefghijklmnopqrstuvwxyz", "bcdefghijklmnopqrstuvwxyza"}, "Full alphabet"},
		{[]string{"az", "by", "cx", "dw"}, "Different patterns"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Input: %v\n", tc.strings)
		
		result1 := groupStrings(tc.strings)
		result2 := groupStringsDifferences(tc.strings)
		result3 := groupStringsOptimized(tc.strings)
		result4 := groupStringsBruteForce(tc.strings)
		
		fmt.Printf("  Standard: %v\n", result1)
		fmt.Printf("  Differences: %v\n", result2)
		fmt.Printf("  Optimized: %v\n", result3)
		fmt.Printf("  Brute Force: %v\n\n", result4)
	}
	
	// Detailed explanation
	fmt.Println("=== Detailed Explanation ===")
	testStrings := []string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}
	result, explanation := groupStringsWithExplanation(testStrings)
	
	fmt.Printf("Result: %v\n", result)
	for _, step := range explanation {
		fmt.Printf("  %s\n", step)
	}
	
	// Test helper functions
	fmt.Println("\n=== Helper Functions Test ===")
	
	fmt.Printf("Are 'abc' and 'bcd' in same group? %t\n", areInSameGroup("abc", "bcd"))
	fmt.Printf("Are 'abc' and 'xyz' in same group? %t\n", areInSameGroup("abc", "xyz"))
	
	fmt.Printf("All shifts of 'abc': %v\n", getAllShifts("abc"))
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largeStrings := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		base := "abc"
		shift := i % 26
		shifted := make([]byte, len(base))
		for j := range base {
			shifted[j] = (base[j]-'a'+byte(shift))%26 + 'a'
		}
		largeStrings[i] = string(shifted)
	}
	
	fmt.Printf("Large test with %d strings\n", len(largeStrings))
	
	result = groupStrings(largeStrings)
	fmt.Printf("Found %d groups\n", len(result))
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Mixed case strings
	mixedCase := []string{"ABC", "BCD", "CDE", "abc", "bcd"}
	fmt.Printf("Mixed case: %v\n", groupStringsAllChars(mixedCase))
	
	// Strings with numbers
	withNumbers := []string{"a1b", "b1c", "c1d"}
	fmt.Printf("With numbers: %v\n", groupStrings(withNumbers))
	
	// Very long strings
	longStrings := []string{
		"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
		"bcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyza",
	}
	fmt.Printf("Long strings: %v\n", groupStrings(longStrings))
	
	// Test with all possible single characters
	fmt.Println("\n=== All Single Characters ===")
	var allChars []string
	for i := 0; i < 26; i++ {
		allChars = append(allChars, string('a'+i))
	}
	
	result = groupStrings(allChars)
	fmt.Printf("All single characters grouped: %v\n", result)
}
