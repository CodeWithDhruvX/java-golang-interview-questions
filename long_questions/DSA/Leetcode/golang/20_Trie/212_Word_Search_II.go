package main

import "fmt"

// TrieNode represents a node in the trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// WordSearch represents the word search board with trie optimization
type WordSearch struct {
	trie *TrieNode
}

// Constructor creates a new WordSearch
func ConstructorWordSearch() WordSearch {
	return WordSearch{
		trie: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

// FindWords finds all words on the board that are in the dictionary
// Time: O(M*N*4^L) where M*N is board size, L is max word length
// Space: O(L) for recursion stack + O(total characters in all words) for trie
func (this *WordSearch) FindWords(board [][]byte, words []string) []string {
	// Build trie from words
	for _, word := range words {
		this.insertWord(word)
	}
	
	var result []string
	m, n := len(board), len(board[0])
	
	// Directions: up, down, left, right
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	var dfs func(int, int, *TrieNode, []byte)
	dfs = func(row, col int, node *TrieNode, path []byte) {
		// Check boundaries
		if row < 0 || row >= m || col < 0 || col >= n {
			return
		}
		
		char := rune(board[row][col])
		
		// Check if character exists in trie
		child, exists := node.children[char]
		if !exists {
			return
		}
		
		// Add character to path
		path = append(path, board[row][col])
		
		// Check if we found a word
		if child.isEnd {
			result = append(result, string(path))
			// Mark as visited to avoid duplicates
			child.isEnd = false
		}
		
		// Temporarily mark the cell as visited
		temp := board[row][col]
		board[row][col] = '#'
		
		// Explore all 4 directions
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			dfs(newRow, newCol, child, path)
		}
		
		// Restore the cell
		board[row][col] = temp
	}
	
	// Start DFS from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if _, exists := this.trie.children[rune(board[i][j])]; exists {
				dfs(i, j, this.trie, []byte{})
			}
		}
	}
	
	return result
}

// insertWord inserts a word into the trie
func (this *WordSearch) insertWord(word string) {
	node := this.trie
	
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		node = node.children[char]
	}
	
	node.isEnd = true
}

// Alternative approach without trie (brute force)
func (this *WordSearch) FindWordsBruteForce(board [][]byte, words []string) []string {
	var result []string
	m, n := len(board), len(board[0])
	
	// Convert words to map for O(1) lookup
	wordSet := make(map[string]bool)
	for _, word := range words {
		wordSet[word] = true
	}
	
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	var dfs func(int, int, string, [][]bool)
	dfs = func(row, col int, current string, visited [][]bool) {
		if len(current) > 10 { // Limit word length to prevent excessive recursion
			return
		}
		
		// Check if current string is a valid word
		if wordSet[current] {
			// Add to result if not already present
			found := false
			for _, word := range result {
				if word == current {
					found = true
					break
				}
			}
			if !found {
				result = append(result, current)
			}
		}
		
		// Explore all 4 directions
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && !visited[newRow][newCol] {
				visited[newRow][newCol] = true
				dfs(newRow, newCol, current+string(board[newRow][newCol]), visited)
				visited[newRow][newCol] = false
			}
		}
	}
	
	// Start DFS from each cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited := make([][]bool, m)
			for k := range visited {
				visited[k] = make([]bool, n)
			}
			visited[i][j] = true
			dfs(i, j, string(board[i][j]), visited)
		}
	}
	
	return result
}

// Helper function to create board from strings
func createBoard(boardStr []string) [][]byte {
	board := make([][]byte, len(boardStr))
	for i, row := range boardStr {
		board[i] = []byte(row)
	}
	return board
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Trie-Optimized Word Search on 2D Grid
- **Trie Pruning**: Use trie to guide search and prune invalid paths
- **DFS Traversal**: Depth-first search on 2D grid
- **Path Tracking**: Maintain current path during traversal
- **Early Termination**: Stop when path doesn't match any word prefix

## 2. PROBLEM CHARACTERISTICS
- **2D Grid Search**: Find words in character matrix
- **Multiple Words**: Search for dictionary words simultaneously
- **Path Constraints**: Can move horizontally/vertically, no cell reuse
- **Optimization Need**: Brute force is too expensive (O(M*N*4^L))

## 3. SIMILAR PROBLEMS
- Word Search (LeetCode 79) - Single word search on grid
- Design Add and Search Words (LeetCode 211) - Trie with wildcards
- Implement Trie (LeetCode 208) - Basic trie implementation
- Boggle Solver - Classic word search game

## 4. KEY OBSERVATIONS
- **Trie Guidance**: Trie tells if current path can lead to valid word
- **Pruning Power**: Stop exploring paths not in any word prefix
- **DFS Nature**: Natural fit for path exploration on grid
- **Cell Reuse**: Each cell can be used only once per word

## 5. VARIATIONS & EXTENSIONS
- **Diagonal Movement**: Allow 8-directional movement
- **Word Length Limits**: Restrict minimum/maximum word lengths
- **Multiple Boards**: Search across multiple connected boards
- **Real-time Updates**: Handle dynamic board changes

## 6. INTERVIEW INSIGHTS
- Always clarify: "Movement directions? Word length limits? Board size?"
- Edge cases: empty board, no words, single cell boards
- Time complexity: O(M*N*4^L) worst case, O(M*N*total chars) with trie
- Space complexity: O(total characters) for trie + O(L) for recursion
- Key insight: trie prunes impossible paths early

## 7. COMMON MISTAKES
- Not marking cells as visited (allowing reuse)
- Wrong boundary checking in DFS
- Not restoring visited marks after backtracking
- Missing diagonal directions if required
- Inefficient duplicate detection

## 8. OPTIMIZATION STRATEGIES
- **Trie Pruning**: O(M*N*total chars) vs O(M*N*4^L) brute force
- **Early Termination**: Stop when path not in trie
- **Duplicate Prevention**: Mark found words to avoid duplicates
- **Direction Optimization**: Check only valid directions

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like solving a word search puzzle with a smart dictionary:**
- You have a grid of letters and need to find hidden words
- Instead of trying all possible paths blindly, use a trie as guide
- The trie tells you "this path could lead to a valid word" or "stop here"
- Like having a GPS that only lets you explore roads that lead to destinations

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: 2D character board, list of words to find
2. **Goal**: Find all words from list that exist on board
3. **Constraint**: Words formed by adjacent cells, no cell reuse
4. **Movement**: Typically 4-directional (up, down, left, right)

#### Phase 2: Key Insight Recognition
- **"Brute force too slow"** → O(M*N*4^L) is exponential
- **"Trie natural fit"** → Guides search and prunes invalid paths
- **"DFS traversal"** → Natural for path exploration
- **"Prefix checking"** → Stop when current path not in any word prefix

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find words on a 2D grid efficiently.
Brute force tries all paths from each cell - too slow!
Instead, build a trie from all words, then:
1. Start DFS from each cell
2. At each step, check if current path is in trie
3. If not in trie, stop exploring this path (pruning!)
4. If path leads to word end, add to results
5. Mark cells as visited to prevent reuse
6. Backtrack and restore visited marks

This way, we only explore paths that could lead to valid words!"
```

#### Phase 4: Edge Case Handling
- **Empty board**: Return empty result
- **Empty word list**: Return empty result
- **Single cell**: Handle 1x1 board correctly
- **No matches**: Return empty array, not null

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Board: ["oath", "pea", "eat", "rain"]
Words: ["oath", "pea", "eat", "rain"]

Human thinking:
"I'll build trie from words, then search:

Build trie:
- "oath": o->a->t->h (END)
- "pea": p->e->a (END)
- "eat": e->a->t (END)
- "rain": r->a->i->n (END)

Search from cell (0,0) = 'o':
- Path "o" is in trie, continue
- Try directions: right to 'a', down to 'p'
- Right to 'a': path "oa" in trie, continue
- Right to 't': path "oat" not in trie, stop
- Down to 'p': path "op" not in trie, stop
- Backtrack to 'o', try other directions...

Search from cell (0,1) = 'a':
- Path "a" in trie, continue
- Try directions: many possibilities
- Eventually find "oath" path: o->a->t->h ✓
- Found word, add to results

Continue from all cells...
Final results: ["oath", "pea", "eat", "rain"] ✓"
```

#### Phase 6: Intuition Validation
- **Why trie works**: Prunes impossible paths early
- **Why DFS works**: Natural for path exploration
- **Why visited marking needed**: Prevents cell reuse in same word
- **Why significant speedup**: Reduces exponential to linear in total characters

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just brute force?"** → O(4^L) per starting cell is too slow
2. **"Should I use BFS?"** → DFS is more natural for path building
3. **"What about diagonal moves?"** → Clarify movement constraints
4. **"Can I optimize further?"** → Trie is already optimal for this problem
5. **"What about duplicates?"** → Need to mark found words to avoid duplicates

### Real-World Analogy
**Like solving a word search puzzle with a smart assistant:**
- You have a crossword-style grid and a dictionary
- Instead of trying random letter combinations, your assistant checks the dictionary
- As you trace a path, the assistant says "keep going" or "stop, no words start like this"
- This prevents wasting time on dead-end paths
- You only explore paths that could actually lead to valid words

### Human-Readable Pseudocode
```
function findWords(board, words):
    // Build trie from words
    trie = buildTrie(words)
    result = []
    
    for each cell (i,j) in board:
        if board[i][j] is in trie.children:
            dfs(i, j, trie, "", result)
    
    return result

function dfs(row, col, node, path, result):
    // Check boundaries and visited
    if outOfBounds or visited:
        return
    
    char = board[row][col]
    if char not in node.children:
        return
    
    // Mark as visited
    visited[row][col] = true
    path += char
    
    // Check if we found a word
    if node.children[char].isEnd:
        result.add(path)
        // Mark to avoid duplicates
        node.children[char].isEnd = false
    
    // Explore all 4 directions
    for each direction (up, down, left, right):
        newRow, newCol = row + dir.row, col + dir.col
        dfs(newRow, newCol, node.children[char], path, result)
    
    // Backtrack
    visited[row][col] = false
```

### Execution Visualization

### Example: Board = ["oath","pea","eat","rain"], Words = ["oath","pea","eat","rain"]
```
Trie Construction:
o -> a -> t -> h (END)
p -> e -> a (END)
e -> a -> t (END)
r -> a -> i -> n (END)

Search Process:
Starting from (0,0) = 'o':
Path: "o" ✓ (in trie)
From 'o', try right to 'a':
Path: "oa" ✓ (in trie)
From 'a', try right to 't':
Path: "oat" ✗ (not in trie) → backtrack
From 'a', try down to 'p':
Path: "op" ✗ (not in trie) → backtrack

Starting from (0,1) = 'a':
Path: "a" ✓ (in trie)
From 'a', try left to 'o':
Path: "ao" ✗ (not in trie) → backtrack
From 'a', try down to 'e':
Path: "ae" ✗ (not in trie) → backtrack
From 'a', try right to 't':
Path: "at" ✓ (in trie)
From 't', try down to 'h':
Path: "ath" ✓ (in trie)
From 'h', try left to 'o':
Path: "oath" ✓ (found word!) → add to results

Continue exploring all cells...
Final results: ["oath", "pea", "eat", "rain"] ✓
```

### Key Visualization Points:
- **Trie Pruning**: Stop when path not in trie
- **DFS Exploration**: Systematic path building
- **Visited Tracking**: Prevent cell reuse
- **Word Detection**: Check isEnd at each step

### Memory Layout Visualization:
```
Board State During Search:
o a t h
p e a t
e a t n
r a i n

Current search path: o->a->t->h
Visited cells: (0,0), (0,1), (0,2), (0,3)
Path string: "oath"
Trie navigation: root->o->a->t->h (END) ✓
Found word: "oath" → add to results
```

### Time Complexity Breakdown:
- **Trie Building**: O(total characters) time, O(total characters) space
- **Search**: O(M*N*4^L) worst case, O(M*N*total chars) with trie pruning
- **Space**: O(total characters) for trie + O(L) for recursion stack
- **Optimization**: Trie reduces exponential to linear in many cases

### Alternative Approaches:

#### 1. Brute Force DFS (O(M*N*4^L) time, O(L) space)
```go
func findWordsBruteForce(board [][]byte, words []string) []string {
    wordSet := make(map[string]bool)
    for _, word := range words {
        wordSet[word] = true
    }
    
    var result []string
    directions := [][2]int{{-1,0},{1,0},{0,-1},{0,1}}
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            visited := make([][]bool, len(board))
            dfsBruteForce(i, j, "", board, visited, wordSet, &result, directions)
        }
    }
    
    return result
}
```
- **Pros**: Simple to implement
- **Cons**: Exponential time, explores all paths blindly

#### 2. Hash Set with Prefix Pruning (O(M*N*4^L) time, O(L) space)
```go
func findWordsWithPrefix(board [][]byte, words []string) []string {
    // Build prefix set for pruning
    prefixes := make(map[string]bool)
    for _, word := range words {
        for i := 1; i <= len(word); i++ {
            prefixes[word[:i]] = true
        }
    }
    
    // Use prefix set to prune DFS
    // Still exponential but better than pure brute force
}
```
- **Pros**: Some pruning without full trie
- **Cons**: Still exponential, less efficient than trie

#### 3. BFS from Words (O(W*L^2) time, O(M*N) space)
```go
func findWordsBFS(board [][]byte, words []string) []string {
    // For each word, try to find it on board
    // More efficient when few words, large board
    for _, word := range words {
        if findWordOnBoard(board, word) {
            result = append(result, word)
        }
    }
    return result
}
```
- **Pros**: Linear in number of words
- **Cons**: Inefficient when many words, small board

### Extensions for Interviews:
- **Diagonal Movement**: Allow 8-directional movement
- **Word Length Limits**: Restrict minimum/maximum word lengths
- **Multiple Boards**: Search across multiple connected boards
- **Real-time Updates**: Handle dynamic board changes
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	ws := ConstructorWordSearch()
	
	fmt.Println("=== Test Case 1 ===")
	board1 := createBoard([]string{"oath", "pea", "eat", "rain"})
	words1 := []string{"oath", "pea", "eat", "rain", "h"}
	
	result1 := ws.FindWords(board1, words1)
	fmt.Printf("Board: %v\n", board1)
	fmt.Printf("Words: %v\n", words1)
	fmt.Printf("Found: %v\n\n", result1)
	
	fmt.Println("=== Test Case 2 ===")
	board2 := createBoard([]string{"a", "b"})
	words2 := []string{"ab", "ba", "a", "b"}
	
	result2 := ws.FindWords(board2, words2)
	fmt.Printf("Board: %v\n", board2)
	fmt.Printf("Words: %v\n", words2)
	fmt.Printf("Found: %v\n\n", result2)
	
	fmt.Println("=== Test Case 3 ===")
	board3 := createBoard([]string{"abc", "def", "ghi"})
	words3 := []string{"abc", "abd", "abf", "ade", "aei", "def", "ghi", "cfi"}
	
	result3 := ws.FindWords(board3, words3)
	fmt.Printf("Board: %v\n", board3)
	fmt.Printf("Words: %v\n", words3)
	fmt.Printf("Found: %v\n\n", result3)
	
	fmt.Println("=== Performance Comparison ===")
	// Compare trie vs brute force approach
	ws2 := ConstructorWordSearch()
	
	testBoard := createBoard([]string{"abcd", "efgh", "ijkl", "mnop"})
	testWords := []string{"aeim", "bfjn", "cgko", "dhlp", "abcd", "mnop", "ae", "im", "xyz"}
	
	// Trie approach
	resultTrie := ws2.FindWords(testBoard, testWords)
	fmt.Printf("Trie approach found: %v\n", resultTrie)
	
	// Brute force approach
	resultBrute := ws2.FindWordsBruteForce(testBoard, testWords)
	fmt.Printf("Brute force found: %v\n", resultBrute)
}
