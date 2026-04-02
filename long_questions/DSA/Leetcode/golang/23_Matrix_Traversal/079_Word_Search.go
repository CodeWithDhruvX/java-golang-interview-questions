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
