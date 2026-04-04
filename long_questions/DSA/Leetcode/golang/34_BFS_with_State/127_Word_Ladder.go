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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: BFS with State Tracking
- **State-Based BFS**: Track complex state beyond just position
- **Word Transformation**: Generate all possible next states
- **Bidirectional Search**: Search from both ends simultaneously
- **Heuristic Guidance**: Use distance-to-goal heuristics

## 2. PROBLEM CHARACTERISTICS
- **Word Graph**: Words connected by one-letter differences
- **Shortest Path**: Find minimum transformation sequence
- **Unweighted Graph**: Each transformation has equal cost
- **State Generation**: Dynamically generate neighbor states

## 3. SIMILAR PROBLEMS
- Open the Lock (LeetCode 752) - Combination lock BFS
- Minimum Genetic Mutation (LeetCode 433) - Gene sequence BFS
- Word Ladder II (LeetCode 126) - Find all shortest paths
- Sliding Puzzle (LeetCode 773) - 2D puzzle BFS

## 4. KEY OBSERVATIONS
- **Implicit Graph**: Graph edges not explicitly stored
- **State Generation**: Neighbors generated on-the-fly
- **Level Processing**: BFS guarantees shortest path in unweighted graphs
- **Visited Tracking**: Essential to avoid cycles

## 5. VARIATIONS & EXTENSIONS
- **Bidirectional Search**: Reduce search space by half
- **Heuristic Search**: Use A* algorithm with distance heuristics
- **Path Reconstruction**: Track parent pointers for actual paths
- **Multiple Targets**: Find distances to multiple destinations

## 6. INTERVIEW INSIGHTS
- Always clarify: "Word length constraints? Dictionary size? Need actual path?"
- Edge cases: no path exists, begin equals end, missing end word
- Time complexity: O(N * L²) where N=word count, L=word length
- Space complexity: O(N * L) for visited tracking
- Key insight: generate neighbors dynamically instead of storing full graph

## 7. COMMON MISTAKES
- Not handling visited tracking properly
- Generating invalid neighbors (not in dictionary)
- Incorrect level counting in BFS
- Missing bidirectional optimization opportunities
- Not handling edge cases (empty dictionary, same words)

## 8. OPTIMIZATION STRATEGIES
- **Standard BFS**: O(N * L²) time, O(N * L) space - basic approach
- **Bidirectional**: O(N * L²) time, O(N * L) space - faster in practice
- **Heuristic**: O(N * L²) time, O(N * L) space - A* search
- **State Compression**: O(N * L²) time, O(N * L) space - optimized state

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like solving a word puzzle:**
- Each word is a location in a word city
- You can move between locations that differ by one letter
- You start at one location and want to reach another
- You want the shortest route (fewest moves)
- You explore all possible moves level by level
- Like playing word ladder game systematically

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Begin word, end word, dictionary of valid words
2. **Goal**: Find shortest transformation sequence
3. **Rules**: Change one letter at a time, intermediate words must be valid
4. **Output**: Length of shortest sequence (or 0 if impossible)

#### Phase 2: Key Insight Recognition
- **"BFS natural"** → Need shortest path in unweighted graph
- **"Implicit graph"** → Generate neighbors on-the-fly
- **"State generation"** → All one-letter transformations
- **"Bidirectional opportunity"** → Search from both ends

#### Phase 3: Strategy Development
```
Human thought process:
"I need shortest word transformation.
This is shortest path in unweighted graph.

BFS with State Approach:
1. Start BFS from begin word
2. For each word, generate all one-letter transformations
3. Only keep transformations that are in dictionary
4. Track visited words to avoid cycles
5. Process level by level for shortest path
6. Stop when we reach end word

This gives O(N * L²) time!"
```

#### Phase 4: Edge Case Handling
- **No path**: Return 0 if end word unreachable
- **Same words**: Return 0 or 1 depending on interpretation
- **Empty dictionary**: Handle gracefully
- **Missing end word**: Return 0 immediately

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: begin="hit", end="cog", dict=["hot","dot","dog","lot","log","cog"]

Human thinking:
"BFS with State Approach:
Level 0: ["hit"] (start)
Generate neighbors of "hit":
- Change position 0: "ait", "bit", "cit", ..., "zit" → only "hot" valid
- Change position 1: "hat", "hbt", "hct", ..., "hzt" → only "hot" valid  
- Change position 2: "hia", "hib", "hic", ..., "hiz" → only "hot" valid
Level 1: ["hot"]

Level 1: ["hot"]
Generate neighbors of "hot":
- Change position 0: "aot", "bot", "cot", ..., "zot" → "dot", "lot" valid
- Change position 1: "hat", "hbt", "hct", ..., "hzt" → none valid
- Change position 2: "hoa", "hob", "hoc", ..., "hoz" → none valid
Level 2: ["dot", "lot"]

Continue until we reach "cog"...
Result: 5 transformations ✓"
```

#### Phase 6: Intuition Validation
- **Why BFS works**: Unweighted graph, shortest path guaranteed
- **Why state generation**: Implicit graph, don't store all edges
- **Why visited tracking**: Prevent infinite loops and redundant work
- **Why O(N * L²)**: Each word generates L*26 neighbors, N words total

### Common Human Pitfalls & How to Avoid Them
1. **"Why not DFS?"** → DFS doesn't guarantee shortest path
2. **"Should I store full graph?"** → Generate neighbors on-the-fly is better
3. **"What about bidirectional?"** → Can reduce search space significantly
4. **"Can I optimize further?"** → Use heuristics or bidirectional search
5. **"What about path reconstruction?"** → Need parent tracking

### Real-World Analogy
**Like solving a crossword puzzle step by step:**
- Each word is a puzzle piece
- You can change one letter at a time to get new words
- You start with one word and want to reach another
- You explore all possible one-letter changes systematically
- You keep track of words you've already tried
- Like playing word transformation games methodically

### Human-Readable Pseudocode
```
function wordLadder(beginWord, endWord, wordList):
    if endWord not in wordList:
        return 0
    
    wordSet = set(wordList)
    queue = [beginWord]
    visited = {beginWord}
    steps = 0
    
    while queue is not empty:
        levelSize = len(queue)
        
        for i from 0 to levelSize-1:
            current = queue.pop_front()
            
            if current == endWord:
                return steps
            
            # Generate all one-letter transformations
            for pos from 0 to len(current)-1:
                for c from 'a' to 'z':
                    if c != current[pos]:
                        nextWord = current[:pos] + c + current[pos+1:]
                        
                        if nextWord in wordSet and nextWord not in visited:
                            visited.add(nextWord)
                            queue.append(nextWord)
        
        steps += 1
    
    return 0
```

### Execution Visualization

### Example: begin="hit", end="cog", dict=["hot","dot","dog","lot","log","cog"]
```
BFS with State Process:

Level 0: ["hit"]
Generate neighbors of "hit":
- h* t: "ait", "bit", ..., "zit" → "hot" ✓
- h * t: "hat", "hbt", ..., "hzt" → "hot" ✓  
- hi * : "hia", "hib", ..., "hiz" → "hot" ✓
Queue: ["hot"], Visited: {"hit", "hot"}

Level 1: ["hot"]
Generate neighbors of "hot":
- h* t: "aot", "bot", ..., "zot" → "dot", "lot" ✓
- h * t: "hat", "hbt", ..., "hzt" → none
- ho * : "hoa", "hob", ..., "hoz" → none
Queue: ["dot", "lot"], Visited: {"hit", "hot", "dot", "lot"}

Level 2: ["dot", "lot"]
Generate neighbors of "dot":
- d* t: "aot", "bot", ..., "zot" → "got", "lot" ✓
- d * t: "dat", "dbt", ..., "dzt" → "dog" ✓
- do * : "doa", "dob", ..., "doz" → none
Queue: ["lot", "got", "dog"], Visited: {..., "got", "dog"}

Continue until "cog" found at Level 5
Result: 5 ✓
```

### Key Visualization Points:
- **State Generation**: Generate all one-letter transformations
- **Level Processing**: BFS ensures shortest path
- **Visited Tracking**: Prevent cycles and redundant work
- **Early Termination**: Stop when target found

### Graph Visualization:
```
Word Graph Structure:
    hit → hot → dot → dog → cog
          ↓    ↓
         lot → log

Shortest Path: hit → hot → dot → dog → cog (5 steps)
Alternative: hit → hot → lot → log → cog (5 steps)
```

### Time Complexity Breakdown:
- **Standard BFS**: O(N * L²) time, O(N * L) space - basic approach
- **Bidirectional**: O(N * L²) time, O(N * L) space - faster in practice
- **Heuristic**: O(N * L²) time, O(N * L) space - A* search
- **State Compression**: O(N * L²) time, O(N * L) space - optimized state

### Alternative Approaches:

#### 1. Bidirectional BFS (O(N * L²) time, O(N * L) space)
```go
func wordLadderBidirectional(beginWord, endWord, wordList []string) int {
    // Search from both ends simultaneously
    // Reduce search space by half
    // ... implementation details omitted
}
```
- **Pros**: Much faster in practice, reduces search space
- **Cons**: More complex implementation

#### 2. A* Search (O(N * L²) time, O(N * L) space)
```go
func wordLadderAStar(beginWord, endWord, wordList []string) int {
    // Use priority queue with heuristic
    // Heuristic: Hamming distance to target
    // ... implementation details omitted
}
```
- **Pros**: Can be faster with good heuristics
- **Cons**: More complex, priority queue overhead

#### 3. Preprocessing (O(N * L²) time, O(N * L²) space)
```go
func wordLadderPreprocessed(beginWord, endWord, wordList []string) int {
    // Build explicit graph with generic patterns
    // Pattern: "*ot", "h*t", "ho*"
    // Faster neighbor lookup
    // ... implementation details omitted
}
```
- **Pros**: Faster neighbor generation
- **Cons**: Higher memory usage, preprocessing time

### Extensions for Interviews:
- **Path Reconstruction**: Track parent pointers to return actual sequence
- **Multiple Paths**: Find all shortest transformation sequences
- **Bidirectional Optimization**: Discuss when and how to implement
- **Heuristic Design**: Design effective distance heuristics
- **Performance Analysis**: Compare different approaches for various input sizes
*/
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
