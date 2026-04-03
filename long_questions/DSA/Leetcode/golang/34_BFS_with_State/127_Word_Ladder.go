package main

import (
	"fmt"
	"math"
)

// 127. Word Ladder - BFS with State
// Time: O(N * L^2), Space: O(N * L^2) where N is word count, L is word length
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(beginWord) != len(endWord) {
		return 0
	}
	
	// Build adjacency list
	adj := buildAdjacencyList(wordList)
	
	// BFS with state: (current word, position)
	queue := [][]string{{beginWord}}
	visited := make(map[string]bool)
	visited[beginWord] = true
	
	steps := 0
	
	for len(queue) > 0 {
		levelSize := len(queue)
		
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			
			if current == endWord {
				return steps
			}
			
			// Generate all possible next states
			for j := 0; j < len(current); j++ {
				for c := 'a'; c <= 'z'; c++ {
					if c != current[j] {
						nextWord := current[:j] + string(c) + current[j+1:]
						
						if !visited[nextWord] && adj[nextWord] {
							visited[nextWord] = true
							queue = append(queue, nextWord)
						}
					}
				}
			}
		}
		
		steps++
	}
	
	return 0
}

func buildAdjacencyList(wordList []string) map[string]bool {
	adj := make(map[string]bool)
	for _, word := range wordList {
		adj[word] = true
	}
	return adj
}

// BFS with state optimization
func ladderLengthOptimized(beginWord string, endWord string, wordList []string) int {
	if len(beginWord) != len(endWord) {
		return 0
	}
	
	// Build adjacency list
	adj := buildAdjacencyList(wordList)
	
	// BFS with state tracking
	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	
	steps := 0
	
	for len(queue) > 0 {
		levelSize := len(queue)
		
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			
			if current == endWord {
				return steps
			}
			
			// Generate next states efficiently
			nextWords := generateNextWords(current, adj)
			
			for _, nextWord := range nextWords {
				if !visited[nextWord] {
					visited[nextWord] = true
					queue = append(queue, nextWord)
				}
			}
		}
		
		steps++
	}
	
	return 0
}

func generateNextWords(word string, adj map[string]bool) []string {
	var nextWords []string
	
	for i := 0; i < len(word); i++ {
		for c := 'a'; c <= 'z'; c++ {
			if c != word[i] {
				nextWord := word[:i] + string(c) + word[i+1:]
				if adj[nextWord] {
					nextWords = append(nextWords, nextWord)
				}
			}
		}
	}
	
	return nextWords
}

// BFS with bidirectional search
func ladderLengthBidirectional(beginWord string, endWord string, wordList []string) int {
	if len(beginWord) != len(endWord) {
		return 0
	}
	
	// Build adjacency list
	adj := buildAdjacencyList(wordList)
	
	// Bidirectional BFS
	beginQueue := []string{beginWord}
	endQueue := []string{endWord}
	beginVisited := make(map[string]bool)
	endVisited := make(map[string]bool)
	beginVisited[beginWord] = true
	endVisited[endWord] = true
	
	steps := 0
	
	for len(beginQueue) > 0 && len(endQueue) > 0 {
		levelSize := len(beginQueue)
		
		// Expand from begin side
		for i := 0; i < levelSize; i++ {
			current := beginQueue[0]
			beginQueue = beginQueue[1:]
			
			// Check if we found connection
			if endVisited[current] {
				return steps + 1
			}
			
			nextWords := generateNextWords(current, adj)
			for _, nextWord := range nextWords {
				if !beginVisited[nextWord] {
					beginVisited[nextWord] = true
					beginQueue = append(beginQueue, nextWord)
				}
			}
		}
		
		// Expand from end side
		levelSize = len(endQueue)
		for i := 0; i < levelSize; i++ {
			current := endQueue[0]
			endQueue = endQueue[1:]
			
			// Check if we found connection
			if beginVisited[current] {
				return steps + 1
			}
			
			nextWords := generateNextWords(current, adj)
			for _, nextWord := range nextWords {
				if !endVisited[nextWord] {
					endVisited[nextWord] = true
					endQueue = append(endQueue, nextWord)
				}
			}
		}
		
		steps++
	}
	
	return 0
}

// BFS with state compression
func ladderLengthStateCompression(beginWord string, endWord string, wordList []string) int {
	if len(beginWord) != len(endWord) {
		return 0
	}
	
	// Build adjacency list with pattern matching
	adj := buildAdjacencyList(wordList)
	
	// Use state: (current word, position)
	type State struct {
		word   string
		pos    int
	}
	
	queue := []State{{beginWord, 0}}
	visited := make(map[string]bool)
	visited[beginWord] = true
	
	steps := 0
	
	for len(queue) > 0 {
		levelSize := len(queue)
		
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			
			if current.word == endWord {
				return steps
			}
			
			// Generate next states with position tracking
			for pos := current.pos; pos < len(current.word); pos++ {
				for c := 'a'; c <= 'z'; c++ {
					if c != current.word[pos] {
						nextWord := current.word[:pos] + string(c) + current.word[pos+1:]
						
						if !visited[nextWord] && adj[nextWord] {
							visited[nextWord] = true
							queue = append(queue, State{nextWord, pos + 1})
						}
					}
				}
			}
		}
		
		steps++
	}
	
	return 0
}

// BFS with heuristic
func ladderLengthHeuristic(beginWord string, endWord string, wordList []string) int {
	if len(beginWord) != len(endWord) {
		return 0
	}
	
	// Build adjacency list
	adj := buildAdjacencyList(wordList)
	
	// Priority queue simulation with heuristic (distance to target)
	type Item struct {
		word  string
		steps int
		heuristic int
	}
	
	pq := []Item{{beginWord, 0, calculateHeuristic(beginWord, endWord)}}
	visited := make(map[string]bool)
	visited[beginWord] = true
	
	for len(pq) > 0 {
		// Find item with minimum priority
		minIdx := 0
		for i := 1; i < len(pq); i++ {
			if pq[i].steps+pq[i].heuristic < pq[minIdx].steps+pq[minIdx].heuristic {
				minIdx = i
			}
		}
		
		current := pq[minIdx]
		pq = append(pq[:minIdx], pq[minIdx+1:]...)
		
		if current.word == endWord {
			return current.steps
		}
		
		// Generate next words
		nextWords := generateNextWords(current.word, adj)
		for _, nextWord := range nextWords {
			if !visited[nextWord] {
				visited[nextWord] = true
				pq = append(pq, Item{nextWord, current.steps + 1, calculateHeuristic(nextWord, endWord)})
			}
		}
	}
	
	return 0
}

func calculateHeuristic(word1, word2 string) int {
	diff := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			diff++
		}
	}
	return diff
}

// BFS with path reconstruction
func ladderLengthWithPath(beginWord string, endWord string, wordList []string) (int, []string) {
	if len(beginWord) != len(endWord) {
		return 0, []string{}
	}
	
	// Build adjacency list
	adj := buildAdjacencyList(wordList)
	
	// BFS with parent tracking
	type State struct {
		word   string
		parent string
	}
	
	queue := []State{{beginWord, ""}}
	visited := make(map[string]bool)
	visited[beginWord] = true
	
	steps := 0
	
	for len(queue) > 0 {
		levelSize := len(queue)
		
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			
			if current.word == endWord {
				// Reconstruct path
				path := []string{endWord}
				for current.parent != "" {
					path = append([]string{current.parent}, path...)
					// Find parent in visited
					for state := range queue {
						if state.word == current.parent {
							current.parent = state.parent
							break
						}
					}
				}
				
				return steps, path
			}
			
			// Generate next words
			nextWords := generateNextWords(current.word, adj)
			for _, nextWord := range nextWords {
				if !visited[nextWord] {
					visited[nextWord] = true
					queue = append(queue, State{nextWord, current.word})
				}
			}
		}
		
		steps++
	}
	
	return 0, []string{}
}

// BFS with multiple targets
func ladderLengthMultipleTargets(beginWord string, targets []string, wordList []string) map[string]int {
	if len(targets) == 0 {
		return map[string]int{}
	}
	
	// Build adjacency list
	adj := buildAdjacencyList(wordList)
	
	// BFS from begin word
	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	
	steps := 0
	results := make(map[string]int)
	
	for len(queue) > 0 {
		levelSize := len(queue)
		
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			
			// Check if current is a target
			for _, target := range targets {
				if current == target && results[target] == 0 {
					results[target] = steps
				}
			}
			
			// Generate next words
			nextWords := generateNextWords(current, adj)
			for _, nextWord := range nextWords {
				if !visited[nextWord] {
					visited[nextWord] = true
					queue = append(queue, nextWord)
				}
			}
		}
		
		steps++
	}
	
	return results
}

func main() {
	// Test cases
	fmt.Println("=== Testing Word Ladder - BFS with State ===")
	
	testCases := []struct {
		beginWord  string
		endWord    string
		wordList   []string
		description string
	}{
		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			"Standard case",
		},
		{
			"hit",
			"hate",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			"No path",
		},
		{
			"a",
			"c",
			[]string{"a", "b", "c"},
			"Simple path",
		},
		{
			"abc",
			"abd",
			[]string{"abc", "abd", "abf", "bef", "bcd"},
			"Multiple paths",
		},
		{
			"game",
			"thee",
			[]string{"fame", "game", "gain", "made", "name", "tame"},
			"Complex case",
		},
		{
			"lead",
			"gold",
			[]string{"load", "goad", "lead", "gold"},
			"Short words",
		},
		{
			"same",
			"same",
			[]string{"same"},
			"Same begin and end",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Begin: %s, End: %s\n", tc.beginWord, tc.endWord)
		fmt.Printf("  Word List: %v\n", tc.wordList)
		
		result1 := ladderLength(tc.beginWord, tc.endWord, tc.wordList)
		result2 := ladderLengthOptimized(tc.beginWord, tc.endWord, tc.wordList)
		result3 := ladderLengthBidirectional(tc.beginWord, tc.endWord, tc.wordList)
		result4 := ladderLengthHeuristic(tc.beginWord, tc.endWord, tc.wordList)
		
		fmt.Printf("  Standard BFS: %d\n", result1)
		fmt.Printf("  Optimized BFS: %d\n", result2)
		fmt.Printf("  Bidirectional BFS: %d\n", result3)
		fmt.Printf("  Heuristic BFS: %d\n", result4)
		
		// Test path reconstruction
		steps, path := ladderLengthWithPath(tc.beginWord, tc.endWord, tc.wordList)
		fmt.Printf("  With path: steps=%d, path=%v\n", steps, path)
		
		fmt.Println()
	}
	
	// Test multiple targets
	fmt.Println("=== Multiple Targets Test ===")
	targets := []string{"cog", "hot", "lot"}
	results := ladderLengthMultipleTargets("hit", targets, []string{"hot", "dot", "dog", "lot", "log", "cog"})
	fmt.Printf("Targets %v: %v\n", targets, results)
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	// Large word list
	largeWordList := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		word := ""
		for j := 0; j < 3; j++ {
			word += string('a' + (i+j)%26)
		}
		largeWordList[i] = word
	}
	
	fmt.Printf("Large test with %d words\n", len(largeWordList))
	
	result := ladderLengthOptimized("aaa", "zzz", largeWordList)
	fmt.Printf("Optimized BFS result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Empty word list
	fmt.Printf("Empty word list: %d\n", ladderLength("hit", "cog", []string{}))
	
	// Single letter words
	fmt.Printf("Single letters: %d\n", ladderLength("a", "c", []string{"a", "b", "c"}))
	
	// No path
	fmt.Printf("No path: %d\n", ladderLength("hit", "xyz", []string{"hit", "hot", "dot", "dog", "lot", "log", "cog"}))
	
	// Same begin and end
	fmt.Printf("Same begin and end: %d\n", ladderLength("same", "same", []string{"same"}))
	
	// Very long words
	fmt.Printf("Long words: %d\n", ladderLengthOptimized("aaaaaaaa", "zzzzzzzz", []string{"aaaaaaaa", "bbbbbbbb", "zzzzzzzz"}))
	
	// Test state compression
	fmt.Println("\n=== State Compression Test ===")
	compressionResult := ladderLengthStateCompression("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	fmt.Printf("State compression result: %d\n", compressionResult)
}
