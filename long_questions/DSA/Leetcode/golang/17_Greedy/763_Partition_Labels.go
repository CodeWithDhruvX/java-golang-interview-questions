package main

import "fmt"

// 763. Partition Labels
// Time: O(N), Space: O(1) for 26 letters
func partitionLabels(s string) []int {
	// Record the last occurrence of each character
	last := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		last[s[i]] = i
	}
	
	var result []int
	start := 0
	end := 0
	
	for i := 0; i < len(s); i++ {
		end = max(end, last[s[i]])
		
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}
	
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test cases
	testCases := []string{
		"ababcbacadefegdehijhklij",
		"eccbbbbdec",
		"abac",
		"a",
		"aaaaa",
		"abcde",
		"abacaba",
		"abcdefghijklmnopqrstuvwxyz",
		"zzzzzzzzzz",
		"ababababab",
	}
	
	for i, s := range testCases {
		result := partitionLabels(s)
		fmt.Printf("Test Case %d: \"%s\" -> Partitions: %v\n", i+1, s, result)
	}
}
