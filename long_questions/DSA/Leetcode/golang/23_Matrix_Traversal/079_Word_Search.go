package main

import "fmt"

// 79. Word Search
// Time: O(N*3^L) where N is total cells, L is word length
// Space: O(L) for recursion stack
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	
	m, n := len(board), len(board[0])
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(board, i, j, word, 0) {
					return true
				}
			}
		}
	}
	
	return false
}

func dfs(board [][]byte, row, col int, word string, index int) bool {
	// Base case: all characters found
	if index == len(word) {
		return true
	}
	
	// Boundary check
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return false
	}
	
	// Character mismatch
	if board[row][col] != word[index] {
		return false
	}
	
	// Mark as visited
	temp := board[row][col]
	board[row][col] = '#'
	
	// Explore all 4 directions
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	
	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]
		if dfs(board, newRow, newCol, word, index+1) {
			// Restore before returning
			board[row][col] = temp
			return true
		}
	}
	
	// Restore
	board[row][col] = temp
	
	return false
}

// Optimized version with early pruning
func existOptimized(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	
	m, n := len(board), len(board[0])
	
	// Count characters in board and word
	boardCount := make(map[byte]int)
	wordCount := make(map[byte]int)
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			boardCount[board[i][j]]++
		}
	}
	
	for _, char := range word {
		wordCount[byte(char)]++
	}
	
	// Check if board has enough characters
	for char, count := range wordCount {
		if boardCount[char] < count {
			return false
		}
	}
	
	// Start DFS from each matching cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfsOptimized(board, i, j, word, 0) {
					return true
				}
			}
		}
	}
	
	return false
}

func dfsOptimized(board [][]byte, row, col int, word string, index int) bool {
	if index == len(word) {
		return true
	}
	
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return false
	}
	
	if board[row][col] != word[index] {
		return false
	}
	
	// Mark as visited
	temp := board[row][col]
	board[row][col] = '#'
	
	// Explore with optimized direction order
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // Right, Left, Down, Up
	
	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]
		if dfsOptimized(board, newRow, newCol, word, index+1) {
			board[row][col] = temp
			return true
		}
	}
	
	board[row][col] = temp
	return false
}

// Iterative approach using stack
func existIterative(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	
	m, n := len(board), len(board[0])
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfsIterativeHelper(board, i, j, word) {
					return true
				}
			}
		}
	}
	
	return false
}

func dfsIterativeHelper(board [][]byte, startRow, startCol int, word string) bool {
	m, n := len(board), len(board[0])
	
	// Stack for DFS: (row, col, wordIndex, visitedCells)
	stack := []struct {
		row, col, wordIndex int
		visited             [][]bool
	}{
		{startRow, startCol, 0, make([][]bool, m)},
	}
	
	// Initialize visited matrix
	for i := range stack[0].visited {
		stack[0].visited[i] = make([]bool, n)
	}
	
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		row, col, wordIndex := current.row, current.col, current.wordIndex
		visited := current.visited
		
		if wordIndex == len(word) {
			return true
		}
		
		if row < 0 || row >= m || col < 0 || col >= n {
			continue
		}
		
		if visited[row][col] || board[row][col] != word[wordIndex] {
			continue
		}
		
		visited[row][col] = true
		
		// Explore all directions
		directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			
			// Create new visited matrix for each path
			newVisited := make([][]bool, m)
			for i := range newVisited {
				newVisited[i] = make([]bool, n)
				copy(newVisited[i], visited[i])
			}
			
			stack = append(stack, struct {
				row, col, wordIndex int
				visited             [][]bool
			}{newRow, newCol, wordIndex + 1, newVisited})
		}
	}
	
	return false
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

## 1. ALGORITHM PATTERN: DFS Word Search in 2D Grid
- **Grid Traversal**: Search for word in 2D character matrix
- **Backtracking**: Try all 4 directions, backtrack on dead ends
- **Path Building**: Construct path one character at a time
- **Visited Tracking**: Mark cells as visited to prevent reuse

## 2. PROBLEM CHARACTERISTICS
- **2D Grid Search**: Find word path in character matrix
- **4-Direction Movement**: Can move horizontally/vertically adjacent
- **Path Constraints**: Cannot reuse cells in same word path
- **Early Termination**: Stop when word found or exhausted

## 3. SIMILAR PROBLEMS
- Word Search II (LeetCode 212) - Multiple words with Trie optimization
- Number of Islands (LeetCode 200) - Connected components
- Robot Path Problems - Path finding with constraints
- Maze Solving - Find path through obstacles

## 4. KEY OBSERVATIONS
- **DFS Natural Fit**: Depth-first search explores paths naturally
- **Backtracking Essential**: Need to undo moves when exploring alternatives
- **Character Matching**: Each step must match corresponding word character
- **Visited Management**: Critical to prevent cell reuse in same path

## 5. VARIATIONS & EXTENSIONS
- **Diagonal Movement**: Allow 8-directional movement
- **Multiple Words**: Search for multiple words simultaneously
- **Word Counting**: Count occurrences of word
- **Path Finding**: Return actual path coordinates

## 6. INTERVIEW INSIGHTS
- Always clarify: "Movement directions? Word length limits? Board size?"
- Edge cases: empty board, single cell, word longer than board cells
- Time complexity: O(N*4^L) worst case, where N=cells, L=word length
- Space complexity: O(L) for recursion stack + O(1) for visited marking
- Key insight: DFS with backtracking is natural fit

## 7. COMMON MISTAKES
- Not marking cells as visited (allowing reuse)
- Wrong boundary checking in DFS
- Not restoring visited marks after backtracking
- Incorrect direction order
- Missing base cases (empty board, word found)

## 8. OPTIMIZATION STRATEGIES
- **Character Frequency Check**: O(N+L) time, prune impossible cases
- **Early Termination**: Stop when word found
- **Direction Ordering**: Optimize direction order based on word characteristics
- **Iterative DFS**: O(N*4^L) time, O(N) space using explicit stack

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like solving a word search puzzle:**
- You have a grid of letters like a crossword
- Need to find if a specific word can be traced
- Start from cells matching first letter, explore all paths
- Like being a detective following letter trails
- Backtrack when trail goes cold

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: 2D character board, target word
2. **Goal**: Determine if word exists as path in board
3. **Constraints**: Adjacent moves only, no cell reuse
4. **Output**: Boolean indicating word existence

#### Phase 2: Key Insight Recognition
- **"DFS natural fit"** → Explore paths depth-first
- **"Backtracking essential"** → Need to undo moves for alternatives
- **"Visited tracking"** → Prevent cell reuse in same path
- **"Early pruning"** → Stop when character mismatches

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find if word exists in board.
I can use DFS with backtracking:

1. Find all cells matching first character
2. For each matching cell:
   - Start DFS from that cell
   - Try all 4 directions
   - Mark cells as visited
   - Recursively search for remaining characters
   - Backtrack if path fails
3. If any path succeeds → return true
4. If all paths fail → return false

This explores all possible paths systematically!"
```

#### Phase 4: Edge Case Handling
- **Empty board**: Return false
- **Empty word**: Return true (trivially found)
- **Single cell**: Check if word length is 1 and matches
- **Word longer than cells**: Return false immediately

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Board: [
  [A,B,C,E],
  [S,F,C,S],
  [A,D,E,E]
]
Word: "ABCCED"

Human thinking:
"Find all cells with 'A' (first character):
- (0,0) and (2,0) match 'A'

Try from (0,0) = 'A':
- Mark (0,0) visited
- Try directions: Right→'B', Down→'S', Left/Up out of bounds
- Try Right to (0,1) = 'B' (matches word[1])
  - Mark (0,1) visited
  - Try directions: Right→'C', Down→'F', Left→'A'(visited), Up out of bounds
  - Try Right to (0,2) = 'C' (matches word[2])
    - Mark (0,2) visited
    - Try directions: Right→'E', Down→'C', Left→'B'(visited), Up out of bounds
    - Try Down to (1,2) = 'C' (matches word[3])
      - Mark (1,2) visited
      - Try directions: Right→'S', Down→'E', Left→'F', Up→'C'(visited)
      - Try Down to (2,2) = 'E' (matches word[4])
        - Mark (2,2) visited
        - Try directions: Right out of bounds, Down out of bounds, Left→'D', Up→'C'(visited)
        - Try Left to (2,1) = 'D' (matches word[5])
          - Word complete! Return true ✓

Found word path: (0,0)→(0,1)→(0,2)→(1,2)→(2,2)→(2,1) ✓"
```

#### Phase 6: Intuition Validation
- **Why DFS works**: Natural for path exploration
- **Why backtracking works**: Allows exploring all alternatives
- **Why visited tracking works**: Prevents invalid cell reuse
- **Why O(N*4^L)**: Each cell can start, 4^L paths maximum

### Common Human Pitfalls & How to Avoid Them
1. **"Why not try all paths?"** → DFS with backtracking is systematic
2. **"Should I use BFS?"** → BFS finds shortest path, but need any path
3. **"What about diagonal moves?"** → Clarify movement constraints
4. **"Can I optimize further?"** → Character frequency check helps prune
5. **"What about multiple words?"** → Different problem (Word Search II)

### Real-World Analogy
**Like solving a word search puzzle in a newspaper:**
- You have a grid of letters with hidden words
- Start from letters matching first character of target word
- Trace through adjacent letters, trying to spell the word
- If path goes wrong, backtrack and try different route
- Like being a word detective following letter trails

### Human-Readable Pseudocode
```
function exist(board, word):
    if board is empty or word is empty:
        return false or true
    
    m, n = board dimensions
    
    for i from 0 to m-1:
        for j from 0 to n-1:
            if board[i][j] == word[0]:
                if dfs(board, i, j, word, 0):
                    return true
    
    return false

function dfs(board, row, col, word, index):
    // Base case: all characters found
    if index == len(word):
        return true
    
    // Boundary check
    if row < 0 or row >= m or col < 0 or col >= n:
        return false
    
    // Character mismatch
    if board[row][col] != word[index]:
        return false
    
    // Mark as visited
    temp = board[row][col]
    board[row][col] = '#'
    
    // Explore all 4 directions
    directions = [[-1,0], [1,0], [0,-1], [0,1]] // up, down, left, right
    for dir in directions:
        newRow, newCol = row + dir[0], col + dir[1]
        if dfs(board, newRow, newCol, word, index + 1):
            board[row][col] = temp
            return true
    
    // Restore
    board[row][col] = temp
    return false
```

### Execution Visualization

### Example: Board = [[A,B,C,E],[S,F,C,S],[A,D,E,E]], Word = "ABCCED"
```
Board Visualization:
A B C E
S F C S
A D E E

DFS from (0,0) = 'A':
Step 1: At (0,0), looking for 'B'
- Try Right: (0,1) = 'B' ✓
Step 2: At (0,1), looking for 'C'
- Try Right: (0,2) = 'C' ✓
Step 3: At (0,2), looking for 'C'
- Try Right: (0,3) = 'E' ✗ (need 'C')
- Try Down: (1,2) = 'C' ✓
Step 4: At (1,2), looking for 'E'
- Try Right: (1,3) = 'S' ✗ (need 'E')
- Try Down: (2,2) = 'E' ✓
Step 5: At (2,2), looking for 'D'
- Try Right: out of bounds
- Try Down: out of bounds
- Try Left: (2,1) = 'D' ✓
Step 6: Word complete! Return true ✓

Path found: (0,0)→(0,1)→(0,2)→(1,2)→(2,2)→(2,1)
```

### Key Visualization Points:
- **DFS Exploration**: Systematic path building
- **Backtracking**: Restore state when exploring alternatives
- **Visited Management**: Prevent cell reuse in same path
- **Character Matching**: Each step must match word character

### Memory Layout Visualization:
```
Board State During DFS:
A B C E
S F C S
A D E E

Current Path: A→B→C→C→E→D
Visited Cells: (0,0), (0,1), (0,2), (1,2), (2,2), (2,1)
Current Position: (2,1) = 'D'
Next Character: Looking for final character
Word Index: 5/6 (found 5 characters, need 1 more)

Visited Marking:
# B C #
# F # S
# D # #

Path Reconstruction:
A(0,0) → B(0,1) → C(0,2) → C(1,2) → E(2,2) → D(2,1)
```

### Time Complexity Breakdown:
- **Worst Case**: O(N*4^L) where N=total cells, L=word length
- **Best Case**: O(N) when word is found quickly
- **Space**: O(L) for recursion stack + O(1) for visited marking
- **Optimization**: Character frequency check can prune impossible cases

### Alternative Approaches:

#### 1. Character Frequency Pruning (O(N+L) time, O(L) space)
```go
func existOptimized(board [][]byte, word string) bool {
    if len(board) == 0 || len(board[0]) == 0 {
        return false
    }
    
    // Count characters in board and word
    boardCount := make(map[byte]int)
    wordCount := make(map[byte]int)
    
    for i := range board {
        for j := range board[i] {
            boardCount[board[i][j]]++
        }
    }
    
    for _, char := range word {
        wordCount[byte(char)]++
    }
    
    // Check if board has enough characters
    for char, count := range wordCount {
        if boardCount[char] < count {
            return false
        }
    }
    
    // Proceed with DFS only if characters exist
    return exist(board, word)
}
```
- **Pros**: Prunes impossible cases early
- **Cons**: Still O(N*4^L) in worst case

#### 2. Iterative DFS with Stack (O(N*4^L) time, O(N) space)
```go
func existIterative(board [][]byte, word string) bool {
    // Use explicit stack instead of recursion
    // ... implementation details omitted
    // More complex but avoids recursion depth limits
}
```
- **Pros**: No recursion depth issues
- **Cons**: More complex implementation

#### 3. BFS for Shortest Path (O(N*4^L) time, O(N) space)
```go
func existBFS(board [][]byte, word string) bool {
    // BFS to find shortest path for word
    // More complex, finds shortest path but any path works
    // ... implementation details omitted
}
```
- **Pros**: Finds shortest path if needed
- **Cons**: More complex, overkill for existence check

### Extensions for Interviews:
- **Diagonal Movement**: Allow 8-directional movement
- **Multiple Words**: Search for multiple words simultaneously
- **Word Counting**: Count occurrences of word
- **Path Finding**: Return actual path coordinates
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []struct {
		board []string
		word  string
	}{
		{[]string{"ABCE", "SFCS", "ADEE"}, "ABCCED"},
		{[]string{"ABCE", "SFCS", "ADEE"}, "SEE"},
		{[]string{"ABCE", "SFCS", "ADEE"}, "ABCB"},
		{[]string{"a"}, "a"},
		{[]string{"a"}, "b"},
		{[]string{"AB"}, "AB"},
		{[]string{"AB"}, "BA"},
		{[]string{"ABC", "DEF", "GHI"}, "ABEDCFIHG"},
		{[]string{"CAA", "AAA", "BCD"}, "AAB"},
		{[]string{"ABCE", "SFES", "ADEE"}, "ABCESEEEFS"},
	}
	
	for i, tc := range testCases {
		board := createBoard(tc.board)
		
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("  Board: %v\n", tc.board)
		fmt.Printf("  Word: %s\n", tc.word)
		
		// Make copies for different approaches
		board1 := createBoard(tc.board)
		board2 := createBoard(tc.board)
		board3 := createBoard(tc.board)
		
		result1 := exist(board1, tc.word)
		result2 := existOptimized(board2, tc.word)
		result3 := existIterative(board3, tc.word)
		
		fmt.Printf("  Standard DFS: %t\n", result1)
		fmt.Printf("  Optimized DFS: %t\n", result2)
		fmt.Printf("  Iterative: %t\n\n", result3)
	}
}
