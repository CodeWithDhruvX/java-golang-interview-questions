package main

import (
	"fmt"
	"strings"
)

// 127. Word Ladder
// Time: O(M²*N), Space: O(M²*N) where M is word length, N is word list size
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	
	if _, exists := wordSet[endWord]; !exists {
		return 0
	}
	
	// BFS setup
	queue := []string{beginWord}
	level := 1
	visited := make(map[string]bool)
	visited[beginWord] = true
	
	for len(queue) > 0 {
		levelSize := len(queue)
		
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			
			if current == endWord {
				return level
			}
			
			// Generate all possible next words
			for j := 0; j < len(current); j++ {
				for c := 'a'; c <= 'z'; c++ {
					if byte(c) == current[j] {
						continue
					}
					
					nextWord := current[:j] + string(c) + current[j+1:]
					
					if wordSet[nextWord] && !visited[nextWord] {
						visited[nextWord] = true
						queue = append(queue, nextWord)
					}
				}
			}
		}
		level++
	}
	
	return 0
}

// Optimized version using pattern mapping
func ladderLengthOptimized(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	
	if _, exists := wordSet[endWord]; !exists {
		return 0
	}
	
	// Create pattern mapping
	patternMap := make(map[string][]string)
	
	// Add beginWord to wordList for pattern creation
	allWords := append([]string{beginWord}, wordList...)
	for _, word := range allWords {
		for i := 0; i < len(word); i++ {
			pattern := word[:i] + "*" + word[i+1:]
			patternMap[pattern] = append(patternMap[pattern], word)
		}
	}
	
	// BFS setup
	queue := []string{beginWord}
	level := 1
	visited := make(map[string]bool)
	visited[beginWord] = true
	
	for len(queue) > 0 {
		levelSize := len(queue)
		
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			
			if current == endWord {
				return level
			}
			
			// Generate all possible patterns
			for j := 0; j < len(current); j++ {
				pattern := current[:j] + "*" + current[j+1:]
				
				for _, nextWord := range patternMap[pattern] {
					if !visited[nextWord] {
						visited[nextWord] = true
						queue = append(queue, nextWord)
					}
				}
			}
		}
		level++
	}
	
	return 0
}

func main() {
	// Test cases
	testCases := []struct {
		beginWord string
		endWord   string
		wordList  []string
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}},
		{"a", "c", []string{"a", "b", "c"}},
		{"hot", "dog", []string{"hot", "dog", "dot"}},
		{"hit", "cog", []string{"hig", "hog", "hog", "cog"}},
		{"game", "thee", []string{"fame", "gane", "gate", "gaze", "tame", "tape", "tale", "gale", "hale", "hate", "haze", "haze", "hate", "hale", "gale", "tale", "tape", "tame", "gaze", "gate", "gane", "fame", "thee"}},
	}
	
	for i, tc := range testCases {
		result1 := ladderLength(tc.beginWord, tc.endWord, tc.wordList)
		result2 := ladderLengthOptimized(tc.beginWord, tc.endWord, tc.wordList)
		fmt.Printf("Test Case %d: %s -> %s, words=[%s]\n", 
			i+1, tc.beginWord, tc.endWord, strings.Join(tc.wordList, ", "))
		fmt.Printf("  BFS: %d, Optimized: %d\n\n", result1, result2)
	}
}
