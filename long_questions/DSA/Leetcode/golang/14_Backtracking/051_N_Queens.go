package main

import "fmt"

// 51. N-Queens
// Time: O(N!), Space: O(N^2) for board + O(N) for recursion
func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]string, n)
	
	// Initialize empty board
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	
	backtrackNQueens(board, 0, n, &result)
	return result
}

func backtrackNQueens(board [][]string, row, n int, result *[][]string) {
	if row == n {
		// Found a solution
		solution := make([]string, n)
		for i := 0; i < n; i++ {
			solution[i] = ""
			for j := 0; j < n; j++ {
				solution[i] += board[i][j]
			}
		}
		*result = append(*result, solution)
		return
	}
	
	for col := 0; col < n; col++ {
		if isValidNQueens(board, row, col, n) {
			board[row][col] = "Q"
			backtrackNQueens(board, row+1, n, result)
			board[row][col] = "." // Backtrack
		}
	}
}

func isValidNQueens(board [][]string, row, col, n int) bool {
	// Check column
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	
	// Check upper left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	
	// Check upper right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	
	return true
}

// Optimized version using bit manipulation
func solveNQueensBit(n int) [][]string {
	var result [][]string
	queens := make([]int, n)
	
	backtrackNQueensBit(queens, 0, n, 0, 0, 0, &result)
	return result
}

func backtrackNQueensBit(queens []int, row, n, columns, diagonals1, diagonals2 int, result *[][]string) {
	if row == n {
		// Convert queens array to board representation
		board := make([]string, n)
		for i := 0; i < n; i++ {
			rowStr := ""
			for j := 0; j < n; j++ {
				if queens[i] == j {
					rowStr += "Q"
				} else {
					rowStr += "."
				}
			}
			board[i] = rowStr
		}
		*result = append(*result, board)
		return
	}
	
	// Try each column
	for col := 0; col < n; col++ {
		columnMask := 1 << col
		diag1Mask := 1 << (row - col + n - 1)
		diag2Mask := 1 << (row + col)
		
		if (columns&columnMask) == 0 && (diagonals1&diag1Mask) == 0 && (diagonals2&diag2Mask) == 0 {
			queens[row] = col
			backtrackNQueensBit(queens, row+1, n, columns|columnMask, diagonals1|diag1Mask, diagonals2|diag2Mask, result)
		}
	}
}

func main() {
	// Test cases
	testCases := []int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}
	
	for i, n := range testCases {
		result1 := solveNQueens(n)
		result2 := solveNQueensBit(n)
		
		fmt.Printf("Test Case %d: n=%d\n", i+1, n)
		fmt.Printf("  Standard: %d solutions\n", len(result1))
		fmt.Printf("  Bit method: %d solutions\n", len(result2))
		
		if n <= 4 && len(result1) > 0 {
			fmt.Printf("  Sample solution: %v\n", result1[0])
		}
		fmt.Println()
	}
}
