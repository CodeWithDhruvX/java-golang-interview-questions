package main

import (
	"fmt"
	"sort"
)

// 269. Alien Dictionary
// Time: O(C + E) where C is number of unique characters, E is number of edges
// Space: O(C + E)
func alienOrder(words []string) string {
	// Build adjacency list and in-degree count
	adj := make(map[byte][]byte)
	inDegree := make(map[byte]int)
	
	// Initialize all characters
	for _, word := range words {
		for _, char := range word {
			inDegree[byte(char)] = 0
		}
	}
	
	// Build graph from adjacent words
	for i := 0; i < len(words)-1; i++ {
		word1, word2 := words[i], words[i+1]
		minLen := min(len(word1), len(word2))
		
		// Check for invalid case (word2 is prefix of word1)
		if len(word1) > len(word2) && word1[:minLen] == word2[:minLen] {
			return ""
		}
		
		// Find first different character
		for j := 0; j < minLen; j++ {
			if word1[j] != word2[j] {
				char1, char2 := byte(word1[j]), byte(word2[j])
				adj[char1] = append(adj[char1], char2)
				inDegree[char2]++
				break
			}
		}
	}
	
	// Topological sort using BFS (Kahn's algorithm)
	queue := []byte{}
	
	// Find characters with no prerequisites
	for char, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, char)
		}
	}
	
	var result []byte
	
	for len(queue) > 0 {
		// Sort queue to ensure deterministic order
		sort.Slice(queue, func(i, j int) bool {
			return queue[i] < queue[j]
		})
		
		char := queue[0]
		queue = queue[1:]
		
		result = append(result, char)
		
		// Update in-degree for dependent characters
		for _, neighbor := range adj[char] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	
	// Check for cycle
	if len(result) != len(inDegree) {
		return ""
	}
	
	return string(result)
}

// DFS approach for topological sort
func alienOrderDFS(words []string) string {
	// Build adjacency list
	adj := make(map[byte][]byte)
	chars := make(map[byte]bool)
	
	// Initialize all characters
	for _, word := range words {
		for _, char := range word {
			chars[byte(char)] = true
		}
	}
	
	// Build graph from adjacent words
	for i := 0; i < len(words)-1; i++ {
		word1, word2 := words[i], words[i+1]
		minLen := min(len(word1), len(word2))
		
		// Check for invalid case
		if len(word1) > len(word2) && word1[:minLen] == word2[:minLen] {
			return ""
		}
		
		for j := 0; j < minLen; j++ {
			if word1[j] != word2[j] {
				char1, char2 := byte(word1[j]), byte(word2[j])
				adj[char1] = append(adj[char1], char2)
				break
			}
		}
	}
	
	// DFS for topological sort
	visited := make(map[byte]int) // 0 = unvisited, 1 = visiting, 2 = visited
	var result []byte
	hasCycle := false
	
	var dfs func(byte)
	dfs = func(char byte) {
		if visited[char] == 1 {
			hasCycle = true
			return
		}
		if visited[char] == 2 {
			return
		}
		
		visited[char] = 1
		for _, neighbor := range adj[char] {
			dfs(neighbor)
		}
		visited[char] = 2
		result = append(result, char)
	}
	
	// Process all characters
	for char := range chars {
		if visited[char] == 0 {
			dfs(char)
		}
	}
	
	if hasCycle {
		return ""
	}
	
	// Reverse result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	
	return string(result)
}

// Helper function to validate alien dictionary order
func validateAlienOrder(words []string, order string) bool {
	if order == "" {
		return false
	}
	
	// Create character order map
	orderMap := make(map[byte]int)
	for i, char := range order {
		orderMap[byte(char)] = i
	}
	
	// Check if all words follow the order
	for i := 0; i < len(words)-1; i++ {
		word1, word2 := words[i], words[i+1]
		minLen := min(len(word1), len(word2))
		
		for j := 0; j < minLen; j++ {
			char1, char2 := byte(word1[j]), byte(word2[j])
			if char1 != char2 {
				if orderMap[char1] > orderMap[char2] {
					return false
				}
				break
			}
		}
		
		// Check if word2 is prefix of word1 (invalid case)
		if len(word1) > len(word2) && word1[:minLen] == word2[:minLen] {
			return false
		}
	}
	
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test cases
	testCases := []struct {
		words      []string
		description string
	}{
		{[]string{"wrt", "wrf", "er", "ett", "rftt"}, "Standard case"},
		{[]string{"z", "x"}, "Simple two words"},
		{[]string{"z", "x", "z"}, "Cycle case"},
		{[]string{"abc", "ab"}, "Invalid prefix case"},
		{[]string{"a", "b", "c"}, "No constraints"},
		{[]string{"caa", "aaa", "aab"}, "Complex case"},
		{[]string{"ab", "adc"}, "Simple constraint"},
		{[]string{"abc", "ab"}, "Prefix conflict"},
		{[]string{"ab", "abc", "ac"}, "Multiple constraints"},
		{[]string{"a", "a", "b", "b"}, "Duplicate words"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Words: %v\n", tc.words)
		
		result1 := alienOrder(tc.words)
		result2 := alienOrderDFS(tc.words)
		
		fmt.Printf("  BFS Order: '%s'\n", result1)
		fmt.Printf("  DFS Order: '%s'\n", result2)
		
		if result1 != "" {
			valid1 := validateAlienOrder(tc.words, result1)
			fmt.Printf("  BFS Valid: %t\n", valid1)
		}
		if result2 != "" {
			valid2 := validateAlienOrder(tc.words, result2)
			fmt.Printf("  DFS Valid: %t\n", valid2)
		}
		
		fmt.Println()
	}
}
