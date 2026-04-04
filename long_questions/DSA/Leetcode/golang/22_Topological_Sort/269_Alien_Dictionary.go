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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Topological Sort for Alien Dictionary Order
- **Graph Construction**: Build directed graph from word order constraints
- **Edge Detection**: Character A → Character B if A comes before B
- **Topological Sorting**: Find valid character ordering
- **Cycle Detection**: Invalid input if graph has cycles

## 2. PROBLEM CHARACTERISTICS
- **Partial Order**: Words provide relative ordering constraints
- **Character Graph**: Characters as nodes, order relations as edges
- **Prefix Validation**: Longer word cannot be prefix of shorter
- **Unique Order**: Need to find one valid character ordering

## 3. SIMILAR PROBLEMS
- Course Schedule (LeetCode 207) - Cycle detection in dependencies
- Course Schedule II (LeetCode 210) - Find valid course order
- Sequence Reconstruction (LeetCode 444) - Reconstruct from sequence
- Minimum Height Trees (LeetCode 310) - Find tree roots

## 4. KEY OBSERVATIONS
- **Word Pairs**: Each adjacent word pair gives ordering constraint
- **First Difference**: First differing character determines order
- **Prefix Rule**: Longer word cannot be prefix of shorter word
- **Graph Theory**: Problem reduces to topological sort of character DAG

## 5. VARIATIONS & EXTENSIONS
- **Multiple Valid Orders**: Multiple possible character orderings
- **Partial Information**: Not enough constraints to determine complete order
- **Character Sets**: Different alphabets or character sets
- **Dynamic Updates**: Adding new words dynamically

## 6. INTERVIEW INSIGHTS
- Always clarify: "Character set? Multiple valid orders? Prefix validation?"
- Edge cases: single word, duplicate words, insufficient constraints
- Time complexity: O(C + E) where C=unique characters, E=constraints
- Space complexity: O(C + E) for graph storage
- Key insight: topological sort of character dependency graph

## 7. COMMON MISTAKES
- Not handling prefix validation (longer word prefix of shorter)
- Wrong graph direction (A before B vs B before A)
- Not processing all connected components
- Missing cycle detection
- Incorrect character extraction from words

## 8. OPTIMIZATION STRATEGIES
- **BFS (Kahn's)**: O(C + E) time, O(C + E) space
- **DFS Topological Sort**: O(C + E) time, O(C) space
- **Early Validation**: Check prefix constraints during graph building
- **Deterministic Order**: Sort queue for consistent output

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like deducing alphabet order from dictionary:**
- You have words in an unknown alien language
- Adjacent words give clues about character order
- Like seeing "cat" then "car" - you know 't' comes before 'r'
- Build up character constraints and find consistent ordering
- Like solving a puzzle with partial information

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Sorted list of words in alien dictionary
2. **Goal**: Find character order that makes words sorted
3. **Constraint**: Valid order must satisfy all word orderings
4. **Output**: String representing character order

#### Phase 2: Key Insight Recognition
- **"Graph natural fit"** → Characters as nodes, order constraints as edges
- **"Topological sort"** → Find ordering that satisfies all constraints
- **"Prefix validation"** → Longer word cannot be prefix of shorter
- **"First difference"** → First differing characters give order constraint

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find character order from sorted words.
Each adjacent word pair gives me ordering information:

1. For each adjacent word pair:
   - Find first differing character
   - That gives me A → B constraint
   - Also check if longer word is prefix of shorter (invalid!)

2. Build graph with these constraints
3. Apply topological sort to find valid order
4. If cycle exists → no valid ordering

This is classic topological sort on character dependencies!"
```

#### Phase 4: Edge Case Handling
- **Single word**: Return any character order
- **Duplicate words**: Handle gracefully
- **Prefix violation**: Return empty string (invalid input)
- **Insufficient constraints**: May have multiple valid orders

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Words: ["wrt", "wrf", "er", "ett", "rftt"]

Human thinking:
"Let me analyze adjacent pairs:

1. "wrt" vs "wrf":
   - First difference: 't' vs 'f' at position 2
   - Constraint: 't' → 'f'

2. "wrf" vs "er":
   - First difference: 'w' vs 'e' at position 0
   - Constraint: 'w' → 'e'

3. "er" vs "ett":
   - First difference: 'r' vs 't' at position 1
   - Constraint: 'r' → 't'

4. "ett" vs "rftt":
   - First difference: 'e' vs 'r' at position 0
   - Constraint: 'e' → 'r'

Graph constraints:
t → f, w → e, r → t, e → r

Topological sort gives: w, e, r, t, f
Final order: "wertf" ✓"
```

#### Phase 6: Intuition Validation
- **Why graph works**: Natural representation of character dependencies
- **Why topological sort works**: Finds ordering satisfying all constraints
- **Why prefix validation needed**: Prevents impossible cases
- **Why first difference works**: Only relevant character for ordering

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use word order?"** → Need to extract character relationships
2. **"Should I use all characters?"** → Only characters in words matter
3. **"What about duplicate words?"** → Handle gracefully, ignore duplicates
4. **"Can there be multiple valid orders?"** → Yes, return any valid one
5. **"What about prefix cases?"** → Must detect and reject invalid input

### Real-World Analogy
**Like deducing alphabet order from a partially sorted dictionary:**
- You find a dictionary sorted by unknown alphabet
- By comparing adjacent words, you learn character relationships
- "apple" before "apply" tells you 'l' comes before 'p'
- Build up these constraints to deduce the alphabet
- Like being a detective solving a letter-ordering puzzle

### Human-Readable Pseudocode
```
function alienOrder(words):
    if words is empty:
        return ""
    
    // Build graph and in-degree
    adj = adjacency list
    inDegree = character count map
    
    // Process adjacent word pairs
    for i from 0 to len(words)-2:
        word1, word2 = words[i], words[i+1]
        
        // Check prefix constraint
        if len(word1) > len(word2) and word2 is prefix of word1:
            return ""
        
        // Find first difference
        for j from 0 to min(len(word1), len(word2))-1:
            if word1[j] != word2[j]:
                char1, char2 = word1[j], word2[j]
                adj[char1].append(char2)
                inDegree[char2]++
                break
    
    // Topological sort
    queue = characters with in-degree = 0
    result = ""
    
    while queue is not empty:
        char = queue.pop_front()
        result += char
        
        for each neighbor in adj[char]:
            inDegree[neighbor]--
            if inDegree[neighbor] == 0:
                queue.append(neighbor)
    
    if len(result) != number of unique characters:
        return "" // Cycle detected
    
    return result
```

### Execution Visualization

### Example: Words = ["wrt", "wrf", "er", "ett", "rftt"]
```
Graph Construction:
1. "wrt" vs "wrf": first diff at pos 2: 't' vs 'f'
   Constraint: t → f

2. "wrf" vs "er": first diff at pos 0: 'w' vs 'e'
   Constraint: w → e

3. "er" vs "ett": first diff at pos 1: 'r' vs 't'
   Constraint: r → t

4. "ett" vs "rftt": first diff at pos 0: 'e' vs 'r'
   Constraint: e → r

Graph Edges: t→f, w→e, r→t, e→r

Topological Sort:
Initial queue (in-degree 0): [w]
Process w: result="w", neighbors=[e], in-degree[e]=0→-1, queue=[e]
Process e: result="we", neighbors=[r], in-degree[r]=1→0, queue=[r]
Process r: result="wer", neighbors=[t], in-degree[t]=1→0, queue=[t]
Process t: result="wert", neighbors=[f], in-degree[f]=1→0, queue=[f]
Process f: result="wertf", no neighbors, queue=[]

Final order: "wertf" ✓
```

### Key Visualization Points:
- **Edge Detection**: First differing character gives order constraint
- **Prefix Validation**: Longer word cannot be prefix of shorter
- **Graph Building**: Characters as nodes, constraints as directed edges
- **Topological Sort**: BFS (Kahn's) for valid ordering

### Memory Layout Visualization:
```
Graph State During Processing:
Characters: w, r, t, f, e
Edges: t→f, w→e, r→t, e→r

In-degree counts:
w: 0, r: 1, t: 1, f: 1, e: 1

Queue Evolution:
Initial: [w] (characters with in-degree 0)
After w: [e] (e's in-degree becomes 0)
After e: [r] (r's in-degree becomes 0)
After r: [t] (t's in-degree becomes 0)
After t: [f] (f's in-degree becomes 0)
After f: [] (all processed)

Result: "wertf" (all 5 characters) ✓ No cycles!
```

### Time Complexity Breakdown:
- **Graph Building**: O(N * L) where N=words, L=average word length
- **Topological Sort**: O(C + E) where C=unique characters, E=constraints
- **Space**: O(C + E) for graph storage
- **Total**: O(N * L + C + E) time, O(C + E) space

### Alternative Approaches:

#### 1. DFS Topological Sort (O(C+E) time, O(C) space)
```go
func alienOrderDFS(words []string) string {
    // Build same graph as BFS approach
    adj := make(map[byte][]byte)
    inDegree := make(map[byte]int)
    
    // ... graph building code ...
    
    // DFS topological sort
    visited := make(map[byte]int) // 0=unvisited, 1=visiting, 2=visited
    var result []byte
    
    var dfs func(byte)
    dfs = func(char byte) {
        if visited[char] == 1 {
            return // Cycle detected
        }
        if visited[char] == 2 {
            return // Already processed
        }
        
        visited[char] = 1
        for _, neighbor := range adj[char] {
            dfs(neighbor)
        }
        visited[char] = 2
        result = append(result, char)
    }
    
    // Process all characters
    for char := range inDegree {
        if visited[char] == 0 {
            dfs(char)
        }
    }
    
    // Reverse result for correct order
    reverse(result)
    return string(result)
}
```
- **Pros**: No queue needed, natural recursive formulation
- **Cons**: Recursion depth, more complex cycle handling

#### 2. Multiple Valid Orders Detection (O(C+E) time, O(C+E) space)
```go
func alienOrderMultiple(words []string) []string {
    // Build graph and find all possible topological orders
    // This is more complex and typically not required
    // ... implementation details omitted
}
```
- **Pros**: Finds all possible valid orders
- **Cons**: Exponential complexity, usually overkill

#### 3. Constraint Propagation (O(C+E) time, O(C+E) space)
```go
func alienOrderConstraintPropagation(words []string) string {
    // Use constraint satisfaction approach
    // Build partial order constraints and propagate
    // More complex but can detect inconsistencies earlier
    // ... implementation details omitted
}
```
- **Pros**: Early detection of impossible cases
- **Cons**: More complex implementation

### Extensions for Interviews:
- **Multiple Valid Orders**: Discuss when multiple character orders are possible
- **Partial Information**: Handle cases with insufficient constraints
- **Character Sets**: Different alphabets or character sets
- **Dynamic Updates**: Adding new words dynamically
- **Performance Analysis**: Discuss worst-case scenarios
*/
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
