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
