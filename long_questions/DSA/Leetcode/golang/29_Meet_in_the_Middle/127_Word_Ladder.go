package main

import (
	"fmt"
	"sort"
)

// 127. Word Ladder - Meet in the Middle Approach
// Time: O(N * L^2), Space: O(N * L) where N is word count, L is word length
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// Check if endWord exists in wordList
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	
	if !wordSet[endWord] {
		return 0
	}
	
	// Meet in the middle BFS
	beginSet := map[string]bool{beginWord: true}
	endSet := map[string]bool{endWord: true}
	visited := make(map[string]bool)
	level := 1
	
	for len(beginSet) > 0 && len(endSet) > 0 {
		// Always expand the smaller set
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}
		
		nextSet := make(map[string]bool)
		
		for word := range beginSet {
			// Generate all possible next words
			for i := 0; i < len(word); i++ {
				for c := 'a'; c <= 'z'; c++ {
					if byte(c) == word[i] {
						continue
					}
					
					newWord := word[:i] + string(c) + word[i+1:]
					
					if endSet[newWord] {
						return level + 1
					}
					
					if wordSet[newWord] && !visited[newWord] {
						visited[newWord] = true
						nextSet[newWord] = true
					}
				}
			}
		}
		
		beginSet = nextSet
		level++
	}
	
	return 0
}

// Standard BFS approach for comparison
func ladderLengthBFS(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	
	if !wordSet[endWord] {
		return 0
	}
	
	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	level := 1
	
	for len(queue) > 0 {
		levelSize := len(queue)
		
		for i := 0; i < levelSize; i++ {
			word := queue[0]
			queue = queue[1:]
			
			if word == endWord {
				return level
			}
			
			// Generate all possible next words
			for j := 0; j < len(word); j++ {
				for c := 'a'; c <= 'z'; c++ {
					if byte(c) == word[j] {
						continue
					}
					
					newWord := word[:j] + string(c) + word[j+1:]
					
					if wordSet[newWord] && !visited[newWord] {
						visited[newWord] = true
						queue = append(queue, newWord)
					}
				}
			}
		}
		
		level++
	}
	
	return 0
}

// Bidirectional BFS with preprocessing
func ladderLengthBidirectional(beginWord string, endWord string, wordList []string) int {
	// Check if endWord exists
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	
	if !wordSet[endWord] {
		return 0
	}
	
	// Preprocess: create generic patterns
	patternMap := make(map[string][]string)
	
	for word := range wordSet {
		for i := 0; i < len(word); i++ {
			pattern := word[:i] + "*" + word[i+1:]
			patternMap[pattern] = append(patternMap[pattern], word)
		}
	}
	
	// Bidirectional BFS
	beginSet := map[string]bool{beginWord: true}
	endSet := map[string]bool{endWord: true}
	visited := make(map[string]bool)
	visited[beginWord] = true
	visited[endWord] = true
	level := 1
	
	for len(beginSet) > 0 && len(endSet) > 0 {
		// Always expand the smaller set
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}
		
		nextSet := make(map[string]bool)
		
		for word := range beginSet {
			// Generate all patterns for current word
			for i := 0; i < len(word); i++ {
				pattern := word[:i] + "*" + word[i+1:]
				
				// Check all words with this pattern
				for _, nextWord := range patternMap[pattern] {
					if endSet[nextWord] {
						return level + 1
					}
					
					if !visited[nextWord] {
						visited[nextWord] = true
						nextSet[nextWord] = true
					}
				}
			}
		}
		
		beginSet = nextSet
		level++
	}
	
	return 0
}

// Optimized meet in the middle with early termination
func ladderLengthOptimized(beginWord string, endWord string, wordList []string) int {
	// Quick check
	if beginWord == endWord {
		return 1
	}
	
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	
	if !wordSet[endWord] {
		return 0
	}
	
	// If beginWord is not in wordSet, add it
	if !wordSet[beginWord] {
		wordSet[beginWord] = true
	}
	
	// Meet in the middle
	beginSet := map[string]bool{beginWord: true}
	endSet := map[string]bool{endWord: true}
	visited := make(map[string]bool)
	level := 1
	
	for len(beginSet) > 0 && len(endSet) > 0 {
		// Expand smaller set
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}
		
		if len(beginSet)*len(endSet) > len(wordSet) {
			// If sets are too large, switch to standard BFS
			return ladderLengthBFS(beginWord, endWord, wordList)
		}
		
		nextSet := make(map[string]bool)
		
		for word := range beginSet {
			// Generate neighbors
			for i := 0; i < len(word); i++ {
				for c := 'a'; c <= 'z'; c++ {
					if byte(c) == word[i] {
						continue
					}
					
					newWord := word[:i] + string(c) + word[i+1:]
					
					if endSet[newWord] {
						return level + 1
					}
					
					if wordSet[newWord] && !visited[newWord] {
						visited[newWord] = true
						nextSet[newWord] = true
					}
				}
			}
		}
		
		beginSet = nextSet
		level++
	}
	
	return 0
}

// A* search approach (heuristic)
func ladderLengthAStar(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	
	if !wordSet[endWord] {
		return 0
	}
	
	type State struct {
		word  string
		level int
		h     int // heuristic: Hamming distance to endWord
	}
	
	// Priority queue: f = g + h
	pq := []State{{beginWord, 1, hammingDistance(beginWord, endWord)}}
	visited := make(map[string]int)
	visited[beginWord] = 1
	
	for len(pq) > 0 {
		// Get state with minimum f value
		minIdx := 0
		for i := 1; i < len(pq); i++ {
			if pq[i].level+pq[i].h < pq[minIdx].level+pq[minIdx].h {
				minIdx = i
			}
		}
		
		current := pq[minIdx]
		pq = append(pq[:minIdx], pq[minIdx+1:]...)
		
		if current.word == endWord {
			return current.level
		}
		
		// Generate neighbors
		for i := 0; i < len(current.word); i++ {
			for c := 'a'; c <= 'z'; c++ {
				if byte(c) == current.word[i] {
					continue
				}
				
				newWord := current.word[:i] + string(c) + current.word[i+1:]
				
				if wordSet[newWord] {
					if visited[newWord] == 0 || current.level+1 < visited[newWord] {
						visited[newWord] = current.level + 1
						pq = append(pq, State{newWord, current.level + 1, hammingDistance(newWord, endWord)})
					}
				}
			}
		}
	}
	
	return 0
}

func hammingDistance(a, b string) int {
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance
}

// Meet in the middle with distance-based expansion
func ladderLengthDistanceBased(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	
	if !wordSet[endWord] {
		return 0
	}
	
	beginSet := map[string]bool{beginWord: true}
	endSet := map[string]bool{endWord: true}
	visited := make(map[string]bool)
	level := 1
	
	for len(beginSet) > 0 && len(endSet) > 0 {
		// Expand set closer to target
		beginDist := averageDistance(beginSet, endWord)
		endDist := averageDistance(endSet, beginWord)
		
		if beginDist > endDist {
			beginSet, endSet = endSet, beginSet
		}
		
		nextSet := make(map[string]bool)
		
		for word := range beginSet {
			for i := 0; i < len(word); i++ {
				for c := 'a'; c <= 'z'; c++ {
					if byte(c) == word[i] {
						continue
					}
					
					newWord := word[:i] + string(c) + word[i+1:]
					
					if endSet[newWord] {
						return level + 1
					}
					
					if wordSet[newWord] && !visited[newWord] {
						visited[newWord] = true
						nextSet[newWord] = true
					}
				}
			}
		}
		
		beginSet = nextSet
		level++
	}
	
	return 0
}

func averageDistance(words map[string]bool, target string) float64 {
	total := 0
	count := 0
	
	for word := range words {
		total += hammingDistance(word, target)
		count++
	}
	
	if count == 0 {
		return 0
	}
	
	return float64(total) / float64(count)
}

func main() {
	// Test cases
	fmt.Println("=== Testing Word Ladder - Meet in the Middle ===")
	
	testCases := []struct {
		beginWord  string
		endWord    string
		wordList   []string
		description string
	}{
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, "Standard case"},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, "No path"},
		{"a", "c", []string{"a", "b", "c"}, "Single character"},
		{"hit", "hot", []string{"hot"}, "Direct connection"},
		{"hit", "cog", []string{"hot", "dot", "dog", "cog"}, "Short path"},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog", "hog", "hig", "hig"}, "With extra words"},
		{"hit", "hit", []string{"hit"}, "Same word"},
		{"hit", "hot", []string{"hit", "hot"}, "Two words"},
		{"hit", "cog", []string{"hot", "hig", "hig", "hog", "cog"}, "Alternative path"},
		{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog", "cig", "cit", "cot", "cog"}, "Many paths"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Begin: '%s', End: '%s'\n", tc.beginWord, tc.endWord)
		fmt.Printf("  Word List: %v\n", tc.wordList)
		
		result1 := ladderLength(tc.beginWord, tc.endWord, tc.wordList)
		result2 := ladderLengthBFS(tc.beginWord, tc.endWord, tc.wordList)
		result3 := ladderLengthBidirectional(tc.beginWord, tc.endWord, tc.wordList)
		result4 := ladderLengthOptimized(tc.beginWord, tc.endWord, tc.wordList)
		result5 := ladderLengthAStar(tc.beginWord, tc.endWord, tc.wordList)
		result6 := ladderLengthDistanceBased(tc.beginWord, tc.endWord, tc.wordList)
		
		fmt.Printf("  Meet in Middle: %d\n", result1)
		fmt.Printf("  BFS: %d\n", result2)
		fmt.Printf("  Bidirectional: %d\n", result3)
		fmt.Printf("  Optimized: %d\n", result4)
		fmt.Printf("  A*: %d\n", result5)
		fmt.Printf("  Distance-based: %d\n\n", result6)
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	
	// Generate large word list
	largeWordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	
	// Add many variations
	for i := 0; i < 1000; i++ {
		word := "hot"
		for j := 0; j < 3; j++ {
			if i%2 == 0 {
				word = word[:j] + string('a'+(i%26)) + word[j+1:]
			}
		}
		largeWordList = append(largeWordList, word)
	}
	
	fmt.Printf("Large word list with %d words\n", len(largeWordList))
	
	result := ladderLength("hit", "cog", largeWordList)
	fmt.Printf("Meet in middle result: %d\n", result)
	
	result = ladderLengthBFS("hit", "cog", largeWordList)
	fmt.Printf("BFS result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Empty word list
	fmt.Printf("Empty word list: %d\n", ladderLength("hit", "cog", []string{}))
	
	// Same begin and end
	fmt.Printf("Same begin and end: %d\n", ladderLength("hit", "hit", []string{"hit"}))
	
	// Very long words
	longBegin := "aaaaaaaa"
	longEnd := "bbbbbbbb"
	longWordList := []string{"aaaaaaaa", "aaaaaaab", "aaaaabbb", "aaabbbbb", "abbbbbbb", "bbbbbbbb"}
	fmt.Printf("Long words: %d\n", ladderLength(longBegin, longEnd, longWordList))
	
	// Words with no possible transformations
	fmt.Printf("No transformations: %d\n", ladderLength("abc", "xyz", []string{"abc", "def", "ghi"}))
}
