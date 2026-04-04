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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Meet in the Middle BFS
- **Bidirectional Search**: Search from both begin and end simultaneously
- **Set-Based BFS**: Use sets instead of queues for level expansion
- **Dynamic Switching**: Always expand the smaller frontier
- **Early Termination**: Stop when frontiers intersect

## 2. PROBLEM CHARACTERISTICS
- **Word Transformations**: Change one character at a time
- **Shortest Path**: Find minimum number of transformations
- **Unweighted Graph**: Each transformation has equal cost
- **BFS Natural**: BFS guarantees shortest path in unweighted graphs

## 3. SIMILAR PROBLEMS
- Word Ladder II (LeetCode 126) - Find all shortest paths
- Open the Lock (LeetCode 752) - Combination lock problem
- Minimum Genetic Mutation (LeetCode 433) - Gene mutations
- Shortest Path in Unweighted Graph - Classic BFS problem

## 4. KEY OBSERVATIONS
- **Bidirectional Efficiency**: Reduces search space from O(b^d) to O(b^(d/2))
- **Set Operations**: More efficient than queue for level-based expansion
- **Heuristic Benefits**: Can use distance-based expansion strategies
- **Preprocessing**: Pattern matching can optimize neighbor generation

## 5. VARIATIONS & EXTENSIONS
- **A* Search**: Use heuristic to guide search direction
- **Pattern Preprocessing**: Generic patterns for faster neighbor lookup
- **Distance-Based**: Expand based on average distance to target
- **Multiple Sources**: Handle multiple begin/end points

## 6. INTERVIEW INSIGHTS
- Always clarify: "Word length constraints? Dictionary size? Return path length vs path?"
- Edge cases: no path exists, begin equals end, missing end word
- Time complexity: O(N * L^2) where N=word count, L=word length
- Space complexity: O(N * L) for word sets and visited tracking
- Key insight: bidirectional search dramatically reduces search space

## 7. COMMON MISTAKES
- Not expanding the smaller frontier first
- Incorrect neighbor generation (missing valid transformations)
- Not handling visited nodes properly
- Wrong termination conditions
- Inefficient set operations causing performance issues

## 8. OPTIMIZATION STRATEGIES
- **Basic Meet in Middle**: O(N * L^2) time, O(N * L) space - standard
- **Pattern Preprocessing**: O(N * L) time, O(N * L) space - faster neighbor lookup
- **A* Search**: O(N * L^2) time, O(N * L) space - heuristic guided
- **Distance-Based**: O(N * L^2) time, O(N * L) space - intelligent expansion

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the shortest meeting point between two people:**
- Two people start at opposite ends of a city (begin and end words)
- Each person explores all possible one-step moves (word transformations)
- They meet in the middle when their exploration areas overlap
- The total distance is the sum of their individual distances
- Always have the person with fewer options explore first (expand smaller set)
- Like two search parties meeting in the middle instead of one searching the entire area

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Begin word, end word, word list dictionary
2. **Goal**: Find shortest transformation sequence length
3. **Rules**: Change one character at a time, intermediate words must be in dictionary
4. **Output**: Minimum number of transformations, or 0 if impossible

#### Phase 2: Key Insight Recognition
- **"Bidirectional search natural"** → Search from both ends simultaneously
- **"Set-based BFS efficient"** → Sets better than queues for level expansion
- **"Expand smaller frontier"** → Always expand the side with fewer nodes
- **"Early termination possible"** → Stop when frontiers intersect

#### Phase 3: Strategy Development
```
Human thought process:
"I need shortest path in unweighted graph.
Standard BFS would explore O(b^d) nodes.

Meet in the Middle Approach:
1. Start BFS from both begin and end
2. Maintain sets for current frontier of each side
3. Always expand the smaller set
4. Generate all one-character transformations
5. Check for intersection between frontiers
6. Return sum of levels when they meet

This reduces search space to O(b^(d/2))!"
```

#### Phase 4: Edge Case Handling
- **No path**: Return 0 if end word not in dictionary
- **Same words**: Return 1 if begin equals end
- **Empty dictionary**: Handle gracefully
- **Single character words**: Special case for efficiency

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: "hit" → "cog", Dictionary: ["hot","dot","dog","lot","log","cog"]

Human thinking:
"Meet in the Middle Approach:
Level 1:
- Begin set: {"hit"}, End set: {"cog"}
- Expand smaller set (begin): generate neighbors of "hit"
- Neighbors: ["hot"] (valid transformation)
- New begin set: {"hot"}, level = 1

Level 2:
- Begin set: {"hot"}, End set: {"cog"}
- Expand smaller set (end): generate neighbors of "cog"
- Neighbors: ["log", "dog"] (valid transformations)
- New end set: {"log", "dog"}, level = 1

Level 3:
- Begin set: {"hot"}, End set: {"log", "dog"}
- Expand smaller set (begin): generate neighbors of "hot"
- Neighbors: ["dot", "lot"]
- Check intersection: "dot" and "lot" intersect with end set!
- Found connection at level 2 + 1 = 3

Result: 3 transformations ✓"
```

#### Phase 6: Intuition Validation
- **Why bidirectional works**: Reduces exponential search space
- **Why set expansion works**: More efficient than queue for level-based BFS
- **Why smaller first**: Minimizes total number of expansions
- **Why early termination**: Intersection guarantees shortest path

### Common Human Pitfalls & How to Avoid Them
1. **"Why not standard BFS?"** → O(b^d) vs O(b^(d/2)) search space
2. **"Should I use queue?"** → Sets more efficient for level expansion
3. **"What about path reconstruction?"** → Need parent tracking for actual path
4. **"Can I optimize further?"** → Pattern preprocessing for neighbor lookup
5. **"What about heuristics?"** → A* search with Hamming distance heuristic

### Real-World Analogy
**Like coordinating two search teams in a maze:**
- Two teams start at opposite entrances of a maze (begin and end words)
- Each team explores all possible one-step moves from their current position
- Teams communicate when they find paths that connect
- Always have the team with fewer options explore first
- When teams meet, they've found the shortest path through the maze
- Like finding shortest meeting point in a social network between two people

### Human-Readable Pseudocode
```
function wordLadder(beginWord, endWord, wordList):
    if endWord not in wordList:
        return 0
    
    beginSet = {beginWord}
    endSet = {endWord}
    visited = set()
    level = 1
    
    while beginSet and endSet:
        # Always expand the smaller set
        if len(beginSet) > len(endSet):
            swap(beginSet, endSet)
        
        nextSet = set()
        
        for word in beginSet:
            for i from 0 to len(word)-1:
                for c from 'a' to 'z':
                    if c == word[i]:
                        continue
                    
                    newWord = word[:i] + c + word[i+1:]
                    
                    if newWord in endSet:
                        return level + 1
                    
                    if newWord in wordList and newWord not in visited:
                        visited.add(newWord)
                        nextSet.add(newWord)
        
        beginSet = nextSet
        level += 1
    
    return 0
```

### Execution Visualization

### Example: "hit" → "cog" with Dictionary ["hot","dot","dog","lot","log","cog"]
```
Meet in the Middle Process:

Level 1:
Begin: {"hit"}, End: {"cog"}
Expand smaller set (begin):
- Generate neighbors of "hit": ["hot"]
- Intersection? No
New Begin: {"hot"}, End: {"cog"}

Level 2:
Begin: {"hot"}, End: {"cog"}
Expand smaller set (end):
- Generate neighbors of "cog": ["log", "dog"]
- Intersection? No
New Begin: {"hot"}, End: {"log", "dog"}

Level 3:
Begin: {"hot"}, End: {"log", "dog"}
Expand smaller set (begin):
- Generate neighbors of "hot": ["dot", "lot"]
- Intersection? Yes! "dot" and "lot" are in end set
- Return level + 1 = 3

Result: 3 transformations ✓
```

### Key Visualization Points:
- **Bidirectional Search**: Search from both ends simultaneously
- **Set Expansion**: Use sets for efficient level-based expansion
- **Dynamic Switching**: Always expand the smaller frontier
- **Early Termination**: Stop when frontiers intersect

### Memory Layout Visualization:
```
Search Frontiers Evolution:
Level 1: Begin={"hit"}, End={"cog"}
Level 2: Begin={"hot"}, End={"log","dog"}
Level 3: Begin={"dot","lot"}, End={"log","dog"}
Intersection found!

Neighbor Generation:
For "hit":
- Change position 0: "ait", "bit", "cit", ..., "zit"
- Change position 1: "hat", "hbt", "hct", ..., "hzt"
- Change position 2: "hia", "hib", "hic", ..., "hiz"
- Valid in dictionary: "hot"

Search Space Reduction:
Standard BFS: explores all nodes within distance d
Bidirectional: explores nodes within distance d/2 from both ends
```

### Time Complexity Breakdown:
- **Neighbor Generation**: O(L * 26) per word, L=word length
- **Set Operations**: O(1) average for insertion/lookup
- **Total**: O(N * L^2) time, O(N * L) space
- **Optimization**: Reduces from O(b^d) to O(b^(d/2)) where b=branching factor, d=depth

### Alternative Approaches:

#### 1. Standard BFS (O(N * L^2) time, O(N * L) space)
```go
func ladderLengthBFS(beginWord, endWord string, wordList []string) int {
    // Standard single-direction BFS
    // Explores O(b^d) nodes instead of O(b^(d/2))
    // ... implementation details omitted
}
```
- **Pros**: Simpler to implement
- **Cons**: Much larger search space

#### 2. A* Search (O(N * L^2) time, O(N * L) space)
```go
func ladderLengthAStar(beginWord, endWord string, wordList []string) int {
    // Use Hamming distance as heuristic
    // Priority queue based on f = g + h
    // ... implementation details omitted
}
```
- **Pros**: Can be faster with good heuristics
- **Cons**: More complex, priority queue overhead

#### 3. Pattern Preprocessing (O(N * L) time, O(N * L) space)
```go
func ladderLengthPattern(beginWord, endWord string, wordList []string) int {
    // Preprocess generic patterns like "*ot", "h*t", "ho*"
    // Faster neighbor lookup using pattern matching
    // ... implementation details omitted
}
```
- **Pros**: Much faster neighbor generation
- **Cons**: Extra preprocessing time and space

### Extensions for Interviews:
- **Path Reconstruction**: Track parent pointers to return actual transformation sequence
- **Multiple Paths**: Find all shortest transformation sequences
- **Weighted Transformations**: Handle different costs for different character changes
- **Bidirectional A***: Combine bidirectional search with heuristics
- **Performance Analysis**: Discuss when to use which approach based on input characteristics
*/
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
