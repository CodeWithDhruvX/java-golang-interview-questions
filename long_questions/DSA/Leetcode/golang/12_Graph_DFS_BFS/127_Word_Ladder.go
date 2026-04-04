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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: BFS on Implicit Graph
- **Word Graph**: Each word is a node, edges connect words with one letter difference
- **Shortest Path**: BFS finds shortest path from beginWord to endWord
- **Level Processing**: Process all words at same transformation distance
- **Neighbor Generation**: Generate all possible one-letter transformations

## 2. PROBLEM CHARACTERISTICS
- **Implicit Graph**: Graph not explicitly built, generated on-the-fly
- **Word Transformations**: Each transformation changes exactly one letter
- **Shortest Path Length**: Return number of words in shortest transformation
- **Bidirectional Edges**: If A can transform to B, B can transform to A

## 3. SIMILAR PROBLEMS
- Word Ladder II (LeetCode 126) - Return all shortest paths
- Minimum Genetic Mutation (LeetCode 433) - Similar with gene mutations
- Open the Lock (LeetCode 752) - Similar with lock combinations
- Sliding Puzzle (LeetCode 773) - BFS on puzzle states

## 4. KEY OBSERVATIONS
- **BFS guarantees shortest path**: First time we reach endWord is optimal
- **Word generation**: For each position, try all 26 letters
- **Visited tracking**: Prevent cycles and redundant processing
- **Early termination**: Stop when endWord is found

## 5. VARIATIONS & EXTENSIONS
- **Bidirectional BFS**: Search from both ends simultaneously
- **Pattern optimization**: Pre-compute word patterns for faster neighbor lookup
- **Multiple end words**: Find shortest path to any of multiple targets
- **Word length variations**: Handle words of different lengths

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are all words the same length? Can words be repeated?"
- Edge cases: beginWord equals endWord, endWord not in wordList
- Time complexity: O(M²*N) where M is word length, N is word count
- Space complexity: O(M²*N) for word set and visited map

## 7. COMMON MISTAKES
- Not checking if endWord exists in wordList
- Forgetting to mark words as visited when enqueuing
- Using wrong time complexity assumptions
- Not handling case where beginWord equals endWord
- Generating invalid transformations (same letter)

## 8. OPTIMIZATION STRATEGIES
- **Pattern mapping**: Pre-compute generic patterns for O(1) neighbor lookup
- **Bidirectional BFS**: Search from both ends, reduces search space
- **Early pruning**: Skip words that cannot possibly reach target
- **Level optimization**: Process entire level before incrementing

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like solving a word puzzle where each step changes one letter:**
- You have a dictionary of valid words (wordList)
- You start with a word (beginWord) and want to reach a target (endWord)
- Each move changes exactly one letter to form another valid word
- You want the shortest sequence of changes to reach the target
- Explore all possible changes level by level until you find the target

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: beginWord, endWord, wordList of valid words
2. **Goal**: Find shortest transformation sequence length
3. **Output**: Number of words in shortest sequence, or 0 if impossible
4. **Constraint**: Each step changes exactly one letter to a valid word

#### Phase 2: Key Insight Recognition
- **"Graph problem"** → Words are nodes, one-letter differences are edges
- **"Shortest path"** → BFS is optimal for unweighted graphs
- **"Implicit graph"** → Generate neighbors on-the-fly, don't build full graph
- **"Level processing"** → All words at same transformation distance form a level

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the shortest word transformation sequence.
Each word that differs by one letter is connected.
This forms a graph where I need the shortest path.
BFS will find the shortest path in an unweighted graph.
I'll generate all possible one-letter transformations for each word."
```

#### Phase 4: Edge Case Handling
- **End word not in list**: Return 0 immediately
- **Begin equals end**: Return 1 or handle as special case
- **Empty word list**: Return 0
- **Single letter words**: Handle edge case of minimal transformations

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: "hit" -> "cog", wordList = ["hot","dot","dog","lot","log","cog"]

Human thinking:
"I'll explore level by level:

Level 1: Start with "hit"
Generate all one-letter changes: "ait","bit","cit"...,"zit"
Only "hot" is in wordList
Queue: ["hot"], level = 1

Level 2: Process "hot"  
Generate changes: "dot","lot" are valid
Queue: ["dot","lot"], level = 2

Level 3: Process "dot","lot"
From "dot": "dog" is valid
From "lot": "log" is valid  
Queue: ["dog","log"], level = 3

Level 4: Process "dog","log"
From "dog": "cog" is valid - FOUND!
From "log": "cog" is valid - FOUND!
Return level + 1 = 5

Shortest path: hit -> hot -> dot -> dog -> cog (5 words)"
```

#### Phase 6: Intuition Validation
- **Why BFS works**: Guarantees shortest path in unweighted graph
- **Why word generation**: Need to find all neighbors in implicit graph
- **Why visited tracking**: Prevent cycles and redundant processing
- **Why O(M²*N)**: For each of N words, generate M*26 transformations

### Common Human Pitfalls & How to Avoid Them
1. **"Why not build the entire graph?"** → Too memory intensive, generate on-the-fly
2. **"Should I use DFS?"** → No, DFS doesn't guarantee shortest path
3. **"What about duplicate words?"** → Use visited set to handle duplicates
4. **"Can I optimize further?"** → Use pattern mapping or bidirectional BFS

### Real-World Analogy
**Like solving a word chain puzzle in a newspaper:**
- You have a starting word and target word
- Each step must change exactly one letter to form a valid English word
- You want the shortest possible chain to reach the target
- Try all possible one-letter changes at each step
- Keep track of words you've already used to avoid going in circles
- Stop as soon as you reach the target word

### Human-Readable Pseudocode
```
function ladderLength(beginWord, endWord, wordList):
    if endWord not in wordList:
        return 0
    
    wordSet = set(wordList)
    queue = [beginWord]
    visited = set(beginWord)
    level = 1
    
    while queue not empty:
        levelSize = length(queue)
        
        for i from 0 to levelSize-1:
            current = queue.dequeue()
            
            if current == endWord:
                return level
            
            for j from 0 to length(current)-1:
                for c from 'a' to 'z':
                    if c == current[j]:
                        continue
                    
                    nextWord = current[0:j] + c + current[j+1:]
                    
                    if nextWord in wordSet and nextWord not in visited:
                        visited.add(nextWord)
                        queue.enqueue(nextWord)
        
        level++
    
    return 0
```

### Execution Visualization

### Example: "hit" -> "cog"
```
Word List: ["hot","dot","dog","lot","log","cog"]

Level 1: queue=["hit"]
- Generate: "hot" (valid)
- queue=["hot"], visited={"hit","hot"}

Level 2: queue=["hot"] 
- Generate: "dot","lot" (valid)
- queue=["dot","lot"], visited={"hit","hot","dot","lot"}

Level 3: queue=["dot","lot"]
- From "dot": generate "dog" (valid)
- From "lot": generate "log" (valid)
- queue=["dog","log"], visited={"hit","hot","dot","lot","dog","log"}

Level 4: queue=["dog","log"]
- From "dog": generate "cog" (FOUND!)
- Return level = 4
```

### Key Visualization Points:
- **Level expansion**: Each level represents one transformation step
- **Word generation**: Try all 26 letters at each position
- **Visited tracking**: Prevent revisiting same words
- **Early termination**: Stop when endWord found

### Memory Layout Visualization:
```
Queue Evolution:
["hit"] → ["hot"] → ["dot","lot"] → ["dog","log"] → []
Level:    1      2        3           4         5

Visited Set Growth:
{"hit"} → {"hit","hot"} → {"hit","hot","dot","lot"} → {"hit","hot","dot","lot","dog","log"}
```

### Time Complexity Breakdown:
- **Each word processed**: O(M*26) word generations
- **Total words**: N in worst case
- **Total time**: O(M²*N) where M is word length
- **Space**: O(N) for word set and visited map

### Alternative Approaches:

#### 1. Pattern Mapping Optimization (O(M*N) time, O(M*N) space)
```go
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
```
- **Pros**: O(1) neighbor lookup instead of O(M*26)
- **Cons**: O(M*N) extra space for pattern map

#### 2. Bidirectional BFS (O(M²*√N) time, O(N) space)
```go
func ladderLengthBidirectional(beginWord string, endWord string, wordList []string) int {
    wordSet := make(map[string]bool)
    for _, word := range wordList {
        wordSet[word] = true
    }
    
    if _, exists := wordSet[endWord]; !exists {
        return 0
    }
    
    beginQueue := []string{beginWord}
    endQueue := []string{endWord}
    beginVisited := make(map[string]bool)
    endVisited := make(map[string]bool)
    beginVisited[beginWord] = true
    endVisited[endWord] = true
    
    level := 1
    
    for len(beginQueue) > 0 && len(endQueue) > 0 {
        // Always expand the smaller queue
        if len(beginQueue) > len(endQueue) {
            beginQueue, endQueue = endQueue, beginQueue
            beginVisited, endVisited = endVisited, beginVisited
        }
        
        levelSize := len(beginQueue)
        
        for i := 0; i < levelSize; i++ {
            current := beginQueue[0]
            beginQueue = beginQueue[1:]
            
            for j := 0; j < len(current); j++ {
                for c := 'a'; c <= 'z'; c++ {
                    if byte(c) == current[j] {
                        continue
                    }
                    
                    nextWord := current[:j] + string(c) + current[j+1:]
                    
                    if endVisited[nextWord] {
                        return level + 1
                    }
                    
                    if wordSet[nextWord] && !beginVisited[nextWord] {
                        beginVisited[nextWord] = true
                        beginQueue = append(beginQueue, nextWord)
                    }
                }
            }
        }
        
        level++
    }
    
    return 0
}
```
- **Pros**: Reduces search space significantly
- **Cons**: More complex implementation

### Extensions for Interviews:
- **Word Ladder II**: Return all shortest paths
- **Multiple Targets**: Find shortest path to any of multiple end words
- **Custom Alphabet**: Handle different character sets
- **Word Length Variations**: Handle words of different lengths
- **Minimum Genetic Mutation**: Similar problem with gene sequences
*/
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
