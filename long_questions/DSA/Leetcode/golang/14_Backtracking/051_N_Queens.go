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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Backtracking with Constraint Validation
- **Row-by-Row Placement**: Place queens one row at a time
- **Conflict Detection**: Check columns and diagonals for conflicts
- **Recursive Exploration**: Try all valid positions in each row
- **Backtracking**: Remove queen when exploring different positions

## 2. PROBLEM CHARACTERISTICS
- **N×N Chessboard**: Square board with N rows and N columns
- **Queen Movement**: Queens attack horizontally, vertically, and diagonally
- **Non-attacking**: No two queens can attack each other
- **Complete Solution**: Place exactly N queens on the board

## 3. SIMILAR PROBLEMS
- N-Queens II (LeetCode 52) - Count solutions only
- Sudoku Solver (LeetCode 37) - Similar constraint satisfaction
- Word Search (LeetCode 79) - Grid-based backtracking
- Knight's Tour (Classic) - Chess piece placement problem

## 4. KEY OBSERVATIONS
- **Row constraint**: Exactly one queen per row
- **Column validation**: Track occupied columns
- **Diagonal validation**: Track occupied diagonals (both directions)
- **Backtracking natural fit**: Need to explore all valid configurations

## 5. VARIATIONS & EXTENSIONS
- **Count Only**: Return count of solutions instead of board configurations
- **Bit Manipulation**: Use bit masks for efficient validation
- **Symmetry Optimization**: Use board symmetry to reduce search space
- **Custom Board**: Non-square boards or obstacles

## 6. INTERVIEW INSIGHTS
- Always clarify: "Do you need actual board positions or just count?"
- Edge cases: N=1 (1 solution), N=2,3 (no solution), N=4+ (solutions exist)
- Time complexity: O(N!) in worst case, much less with pruning
- Space complexity: O(N) for recursion + O(N²) for board

## 7. COMMON MISTAKES
- Not checking both diagonal directions
- Using O(N²) validation instead of O(N) with tracking
- Not handling N=2,3 edge cases properly
- Making shallow copies of board solutions
- Not optimizing validation checks

## 8. OPTIMIZATION STRATEGIES
- **Column tracking**: Use array/set for O(1) column validation
- **Diagonal tracking**: Use arrays for diagonal validation
- **Bit manipulation**: Use bit masks for maximum efficiency
- **Early pruning**: Stop exploring when conflicts detected

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like seating N queens at a dinner table with constraints:**
- You have N queens and N×N chessboard
- Each queen must sit in a different row (we place one per row)
- No two queens can see each other (no same column or diagonal)
- Try placing queens row by row, backtracking when conflicts arise
- When all N queens are placed successfully, you have a valid solution

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Integer N (board size)
2. **Goal**: Place N queens on N×N board without conflicts
3. **Output**: All valid board configurations
4. **Constraint**: Queens attack horizontally, vertically, and diagonally

#### Phase 2: Key Insight Recognition
- **"Row-by-row natural fit"** → Place one queen per row to simplify
- **"Constraint validation essential"** → Need efficient conflict detection
- **"Backtracking required"** → Need to explore all valid configurations
- **"Diagonal complexity"** → Two diagonal directions to track

#### Phase 3: Strategy Development
```
Human thought process:
"I need to place N queens without conflicts.
I'll place them row by row, one queen per row.
For each position, I'll check if it's safe (no conflicts).
If safe, I place the queen and move to next row.
If I can't place a queen in a row, I backtrack.
When all N queens are placed, I found a solution."
```

#### Phase 4: Edge Case Handling
- **N=1**: One solution [Q]
- **N=2,3**: No solutions (return empty)
- **N=4+**: Solutions exist, need to find all
- **Large N**: Handle recursion depth and performance

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
N=4 board:

Human thinking:
"I'll place queens row by row:

Row 0:
- Try (0,0): Place Q at (0,0)
  Row 1:
  - Try (1,0): Same column → conflict
  - Try (1,1): Same diagonal → conflict  
  - Try (1,2): Safe! Place Q at (1,2)
    Row 2:
    - Try (2,0): Safe! Place Q at (2,0)
      Row 3:
      - Try (3,0): Same column → conflict
      - Try (3,1): Same diagonal → conflict
      - Try (3,2): Same column → conflict
      - Try (3,3): Same diagonal → conflict
      No valid positions → backtrack
    Backtrack: Remove Q from (2,0)
    - Try (2,1): Same diagonal → conflict
    - Try (2,2): Same column → conflict
    - Try (2,3): Same diagonal → conflict
    No valid positions → backtrack
  Backtrack: Remove Q from (1,2)
  - Try (1,3): Safe! Place Q at (1,3)
    Row 2:
    - Try (2,0): Same diagonal → conflict
    - Try (2,1): Safe! Place Q at (2,1)
      Row 3:
      - Try (3,0): Safe! Place Q at (3,0)
        SUCCESS! All 4 queens placed
        Solution: [Q...,.Q.,..Q.,...Q]
      Backtrack to find other solutions...

Continue exploring all possibilities..."
```

#### Phase 6: Intuition Validation
- **Why row-by-row works**: Simplifies to one queen per row
- **Why backtracking works**: Need to explore all valid configurations
- **Why O(N!) complexity**: Each row has fewer valid positions
- **Why validation optimization**: O(N) validation vs O(N²) naive approach

### Common Human Pitfalls & How to Avoid Them
1. **"Why not place queens randomly?"** → Systematic approach ensures completeness
2. **"Should I use BFS?"** → BFS would use too much memory for board states
3. **"What about very large N?"** → Bit manipulation optimization needed
4. **"Can I optimize further?"** → Bit masks and symmetry are key optimizations

### Real-World Analogy
**Like scheduling N meetings in N time slots with constraints:**
- You have N meetings and N time slots (rows)
- Each meeting needs exactly one time slot
- No two meetings can conflict (same resources or overlapping times)
- Try scheduling meetings one by one, backtracking when conflicts arise
- When all meetings are scheduled successfully, you have a valid schedule

### Human-Readable Pseudocode
```
function solveNQueens(n):
    board = n×n empty board
    result = []
    
    backtrack(board, 0, n, result)
    return result

function backtrack(board, row, n, result):
    if row == n:
        add copy of board to result
        return
    
    for col from 0 to n-1:
        if isValid(board, row, col, n):
            board[row][col] = 'Q'
            backtrack(board, row + 1, n, result)
            board[row][col] = '.'  // Backtrack

function isValid(board, row, col, n):
    // Check column
    for i from 0 to row-1:
        if board[i][col] == 'Q':
            return false
    
    // Check upper left diagonal
    for i, j from row-1, col-1; i >= 0 && j >= 0; i--, j--:
        if board[i][j] == 'Q':
            return false
    
    // Check upper right diagonal
    for i, j from row-1, col+1; i >= 0 && j < n; i--, j++:
        if board[i][j] == 'Q':
            return false
    
    return true
```

### Execution Visualization

### Example: N=4
```
Board Evolution:
Initial:
....
....
....
....

Row 0, Col 0: Place Q
Q...
....
....

Row 1:
Col 0: conflict (same column)
Col 1: conflict (diagonal)
Col 2: safe → Place Q
Q.Q.
....

Row 2:
Col 0: safe → Place Q
Q.Q.
Q...
....

Row 3:
Col 0: conflict (column)
Col 1: conflict (diagonal)
Col 2: conflict (column)  
Col 3: conflict (diagonal)
No valid positions → Backtrack

Continue exploring all possibilities...
Final solutions found:
.Q..
...Q
Q...
..Q.

..Q.
Q...
...Q
.Q..
```

### Key Visualization Points:
- **Row-by-row placement**: Each level represents one row
- **Conflict validation**: Check column and both diagonals
- **Backtracking**: Remove queen when no valid positions in next row
- **Complete solution**: Base case when all N queens placed
- **Multiple solutions**: Continue exploring after finding each solution

### Memory Layout Visualization:
```
Board State Evolution:
[.,.,.,.] → [Q,.,.,.] → [Q,.,Q,.] → [Q,.,Q,.] → [Q,.,Q,Q] (conflict)
          ↘ [Q,.,.,Q] → [Q,Q,.,Q] (conflict)
          ↘ [Q,.,.,.] → [Q,.,.,.] → [Q,.,.,.]
     ↘ [.,Q,.,.] → [.,Q,.,.] → [.,Q,Q,.] → [.,Q,Q,Q] (conflict)
          ↘ [.,Q,.,Q] → [.,Q,.,Q] → [Q,Q,.,Q] (conflict)
          ↘ [.,Q,.,.] → [.,Q,.,.] → [.,Q,.,.]
... (continues exploring)

Valid solutions found at complete board states
```

### Time Complexity Breakdown:
- **Worst case**: O(N!) - each row has fewer valid positions
- **Pruning effect**: Significantly reduces actual explored paths
- **Space complexity**: O(N) for recursion + O(N²) for board
- **Validation cost**: O(N) per placement with optimization

### Alternative Approaches:

#### 1. Bit Manipulation (O(N!) time, O(N) space)
```go
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
```
- **Pros**: Very fast validation using bit operations
- **Cons**: More complex bit manipulation logic

#### 2. Column Tracking with Arrays (O(N!) time, O(N) space)
```go
func solveNQueensOptimized(n int) [][]string {
    var result [][]string
    board := make([][]string, n)
    columns := make([]bool, n)
    diag1 := make([]bool, 2*n-1) // row-col+n-1
    diag2 := make([]bool, 2*n-1) // row+col
    
    // Initialize board
    for i := 0; i < n; i++ {
        board[i] = make([]string, n)
        for j := 0; j < n; j++ {
            board[i][j] = "."
        }
    }
    
    backtrackOptimized(board, columns, diag1, diag2, 0, n, &result)
    return result
}

func backtrackOptimized(board [][]string, columns, diag1, diag2 []bool, row, n int, result *[][]int) {
    if row == n {
        // Add solution
        return
    }
    
    for col := 0; col < n; col++ {
        if !columns[col] && !diag1[row-col+n-1] && !diag2[row+col] {
            columns[col] = true
            diag1[row-col+n-1] = true
            diag2[row+col] = true
            board[row][col] = "Q"
            
            backtrackOptimized(board, columns, diag1, diag2, row+1, n, result)
            
            columns[col] = false
            diag1[row-col+n-1] = false
            diag2[row+col] = false
            board[row][col] = "."
        }
    }
}
```
- **Pros**: O(1) validation with arrays
- **Cons**: Still uses O(N²) space for board

#### 3. Symmetry Optimization (O(N!/2) time, O(N) space)
```go
func solveNQueensSymmetry(n int) [][]string {
    var result [][]string
    
    // Handle even N: only need to place first N/2 queens in first row
    // Handle odd N: need to handle middle column separately
    // This is a complex optimization that reduces search space by ~2x
    
    // Simplified version - just concept
    if n%2 == 0 {
        // Symmetry for even N
        backtrackSymmetry(n, 0, n/2, &result)
    } else {
        // Symmetry for odd N
        backtrackSymmetry(n, 0, n/2, &result)
        backtrackMiddle(n, n/2, &result)
    }
    
    return result
}
```
- **Pros**: Reduces search space by about half
- **Cons**: Complex implementation, need to handle symmetry carefully

### Extensions for Interviews:
- **N-Queens II**: Count solutions only (no board storage needed)
- **Custom Board**: Board with obstacles or non-square shape
- **Knight's Tour**: Similar backtracking with different movement rules
- **Sudoku Solver**: General constraint satisfaction problem
- **Parallel Processing**: Solve multiple configurations simultaneously
*/
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
