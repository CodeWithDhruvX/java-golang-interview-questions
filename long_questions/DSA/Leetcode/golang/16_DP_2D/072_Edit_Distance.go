package main

import (
	"fmt"
	"math"
)

// 72. Edit Distance
// Time: O(M*N), Space: O(M*N)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	
	// dp[i][j] = minimum operations to convert word1[0:i] to word2[0:j]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Initialize base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i // delete all characters
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // insert all characters
	}
	
	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // no operation needed
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j],   // delete
					dp[i][j-1],   // insert
					dp[i-1][j-1], // replace
				)
			}
		}
	}
	
	return dp[m][n]
}

// Space optimized version: O(min(M,N)) space
func minDistanceOptimized(word1 string, word2 string) int {
	// Ensure word1 is the shorter string for space optimization
	if len(word1) > len(word2) {
		word1, word2 = word2, word1
	}
	
	m, n := len(word1), len(word2)
	prev := make([]int, m+1)
	current := make([]int, m+1)
	
	// Initialize first row
	for i := 0; i <= m; i++ {
		prev[i] = i
	}
	
	for j := 1; j <= n; j++ {
		current[0] = j
		for i := 1; i <= m; i++ {
			if word1[i-1] == word2[j-1] {
				current[i] = prev[i-1]
			} else {
				current[i] = 1 + min(
					prev[i],     // delete
					current[i-1], // insert
					prev[i-1],   // replace
				)
			}
		}
		// Swap prev and current
		prev, current = current, prev
	}
	
	return prev[m]
}

// Function to reconstruct the actual operations
func minDistanceWithPath(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Initialize base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	
	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j],
					dp[i][j-1],
					dp[i-1][j-1],
				)
			}
		}
	}
	
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: 2D Dynamic Programming for String Transformation
- **2D DP Table**: dp[i][j] represents minimum operations to convert word1[0:i] to word2[0:j]
- **Edit Operations**: Insert, delete, replace operations
- **State Transitions**: Based on character comparison
- **Base Cases**: Converting to/from empty string

## 2. PROBLEM CHARACTERISTICS
- **String Transformation**: Convert one string to another with minimum operations
- **Edit Operations**: Insert, delete, replace (each costs 1)
- **Sequential Processing**: Process characters in order
- **Optimal Substructure**: Solution depends on smaller subproblems

## 3. SIMILAR PROBLEMS
- Regular Expression Matching (LeetCode 10) - Pattern matching DP
- Longest Common Subsequence (LeetCode 1143) - Sequence matching DP
- Delete Operation for Two Strings (LeetCode 583) - String operations DP
- One Edit Distance (LeetCode 161) - Limited operations DP

## 4. KEY OBSERVATIONS
- **Edit Operations**: Three basic operations with unit cost
- **Character Match**: No operation needed if characters same
- **Character Mismatch**: Choose min of insert, delete, replace
- **Base Cases**: Converting to/from empty strings
- **Optimal Substructure**: dp[i][j] depends on dp[i-1][j], dp[i][j-1], dp[i-1][j-1]

## 5. VARIATIONS & EXTENSIONS
- **Different Operation Costs**: Weighted edit operations
- **Limited Operations**: Only certain operations allowed
- **Multiple Strings**: Edit distance between multiple strings
- **Path Reconstruction**: Recover actual edit sequence

## 6. INTERVIEW INSIGHTS
- Always clarify: "What are operation costs? Can I reconstruct path?"
- Edge cases: empty strings, identical strings, single character
- Time complexity: O(M×N) where M=len(word1), N=len(word2)
- Space complexity: O(M×N) for DP table
- Space optimization possible using rolling arrays

## 7. COMMON MISTAKES
- Not handling empty string base cases correctly
- Wrong DP recurrence (missing one of three operations)
- Off-by-one errors in DP table access
- Not using proper min function for three values
- Forgetting to add 1 for operation cost

## 8. OPTIMIZATION STRATEGIES
- **2D DP**: O(M×N) time, O(M×N) space
- **Space Optimization**: O(min(M,N)) space using rolling arrays
- **Early Termination**: Not applicable (need full table)
- **Path Reconstruction**: O(M×N) time, O(M×N) space

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like editing a document with minimal changes:**
- You have two documents: original and target
- You need to transform original to target with minimum edits
- Each edit operation (insert, delete, replace) costs 1
- You build a table showing minimum edits for all prefixes
- The final cell gives the answer for full strings

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Two strings word1, word2
2. **Goal**: Convert word1 to word2 with minimum operations
3. **Operations**: Insert, delete, replace (each costs 1)
4. **Output**: Minimum number of operations

#### Phase 2: Key Insight Recognition
- **"2D DP natural fit"** → Compare prefixes of both strings
- **"Edit operations"** → Three ways to transform
- **"Optimal substructure"** → Current state depends on previous states
- **"Base cases"** → Converting to/from empty strings

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find minimum edits to convert word1 to word2.
I'll build a table where dp[i][j] shows min operations for word1[0:i] → word2[0:j].
For each position, I compare current characters:
- If they match, no operation needed → inherit from diagonal
- If they differ, consider three operations:
  1. Insert: dp[i][j-1] + 1
  2. Delete: dp[i-1][j] + 1  
  3. Replace: dp[i-1][j-1] + 1
Take minimum of these three options.
```

#### Phase 4: Edge Case Handling
- **Empty strings**: dp[0][j] = j (insert all), dp[i][0] = i (delete all)
- **Identical strings**: dp[m][n] = 0
- **Single character**: Simple comparison
- **Large strings**: Use space optimization

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
word1 = "horse", word2 = "ros"

Human thinking:
"I'll build a 6×4 table (including empty prefixes):
Initialize base cases:
dp[0][0] = 0 (empty → empty)
dp[1][0] = 1 (delete 'h'), dp[2][0] = 2 (delete 'ho'), etc.
dp[0][1] = 1 (insert 'r'), dp[0][2] = 2 (insert 'ro'), etc.

Fill dp[1][1] (h → r):
- Characters differ, consider operations:
  Insert: dp[1][0] + 1 = 1 + 1 = 2
  Delete: dp[0][1] + 1 = 1 + 1 = 2
  Replace: dp[0][0] + 1 = 0 + 1 = 1
- Minimum: min(2,2,1) = 1

Continue filling...
Final result: dp[5][3] = 3 ✓ Minimum operations"
```

#### Phase 6: Intuition Validation
- **Why 2D DP works**: Each prefix combination has unique edit distance
- **Why three operations**: Covers all possible single-character transformations
- **Why base cases**: Empty string conversions are straightforward
- **Why O(M×N)**: Each cell computed once with O(1) operations

### Common Human Pitfalls & How to Avoid Them
1. **"Why not greedy?"** → Local optimal ≠ global optimal
2. **"Should I use recursion?"** → Works but needs memoization
3. **"What about operation costs?"** → Clarify if all operations cost 1
4. **"Can I optimize space?"** → Yes, use rolling array technique

### Real-World Analogy
**Like editing a manuscript with minimal changes:**
- You have an original manuscript and target version
- You need to transform original to target with minimum edits
- Each edit (add word, delete word, replace word) costs the same
- You compare progressively larger portions of both texts
- The table shows minimum edits needed for each portion

### Human-Readable Pseudocode
```
function minDistance(word1, word2):
    m, n = len(word1), len(word2)
    dp = create 2D array (m+1)×(n+1)
    
    // Base cases: convert to/from empty strings
    for i from 0 to m:
        dp[i][0] = i  // delete all characters
    for j from 0 to n:
        dp[0][j] = j  // insert all characters
    
    // Fill DP table
    for i from 1 to m:
        for j from 1 to n:
            if word1[i-1] == word2[j-1]:
                dp[i][j] = dp[i-1][j-1]  // no operation needed
            else:
                dp[i][j] = 1 + min(
                    dp[i-1][j],    // delete
                    dp[i][j-1],    // insert
                    dp[i-1][j-1]    // replace
                )
    
    return dp[m][n]
```

### Execution Visualization

### Example: word1 = "horse", word2 = "ros"
```
DP Table (rows: word1 prefixes, columns: word2 prefixes):
      ""  r   ro   ros
""    0   1    2    3
"h"    1   1    2    2
"ho"   2   2    1    2
"hor"  3   3    2    2
"hors" 4   4    3    3
"horse" 5   4    3    3

Key transitions:
dp[1][1] (h → r): min(delete, insert, replace) = min(2,2,1) = 1
dp[2][2] (ho → ro): min(delete, insert, replace) = min(3,3,2) = 2
dp[5][3] (horse → ros): min(delete, insert, replace) = min(4,4,3) = 3

Final result: dp[5][3] = 3 ✓
```

### Key Visualization Points:
- **Three Operations**: Insert, delete, replace
- **Character Match**: No operation if characters same
- **Base Cases**: Empty string conversions
- **DP Recurrence**: Each cell depends on three neighbors

### Memory Layout Visualization:
```
Edit Operations Visualization:
word1 = "horse", word2 = "ros"

At dp[3][2] (hor → ro):
- Characters: 'o' vs 'o' (match) → dp[2][1] = 2
- But wait, let me check dp[3][3] (hor → ros):
- Characters: 'r' vs 's' (different)
  Insert: dp[3][2] + 1 = 2 + 1 = 3
  Delete: dp[2][3] + 1 = 2 + 1 = 3
  Replace: dp[2][2] + 1 = 1 + 1 = 2
- Result: min(3,3,2) = 2

This shows how each operation contributes to the minimum.
```

### Time Complexity Breakdown:
- **DP Table**: O(M×N) cells to fill
- **Cell Computation**: O(1) time per cell (min of 3 values)
- **Total Time**: O(M×N)
- **Space Complexity**: O(M×N) for DP table

### Alternative Approaches:

#### 1. Recursive with Memoization (O(M×N) time, O(M×N) space)
```go
func minDistanceRecursive(word1, word2 string) int {
    memo := make(map[string]int)
    return minDistanceHelper(word1, word2, len(word1), len(word2), memo)
}

func minDistanceHelper(w1, w2 string, i, j int, memo map[string]int) int {
    if i == 0 {
        return j  // insert all remaining
    }
    if j == 0 {
        return i  // delete all remaining
    }
    
    key := fmt.Sprintf("%d,%d", i, j)
    if val, exists := memo[key]; exists {
        return val
    }
    
    if w1[i-1] == w2[j-1] {
        result := minDistanceHelper(w1, w2, i-1, j-1, memo)
        memo[key] = result
        return result
    }
    
    insert := minDistanceHelper(w1, w2, i, j-1, memo) + 1
    delete := minDistanceHelper(w1, w2, i-1, j, memo) + 1
    replace := minDistanceHelper(w1, w2, i-1, j-1, memo) + 1
    
    result := min(insert, delete, replace)
    memo[key] = result
    return result
}
```
- **Pros**: Intuitive, follows problem description
- **Cons**: Recursion overhead, same complexity as DP

#### 2. Iterative with Space Optimization (O(M×N) time, O(min(M,N)) space)
```go
func minDistanceOptimized(word1, word2 string) int {
    // Ensure word1 is shorter for space optimization
    if len(word1) > len(word2) {
        word1, word2 = word2, word1
    }
    
    m, n := len(word1), len(word2)
    prev := make([]int, m+1)
    curr := make([]int, m+1)
    
    // Initialize first row
    for i := 0; i <= m; i++ {
        prev[i] = i  // delete all characters
    }
    
    for j := 1; j <= n; j++ {
        curr[0] = j  // insert all characters
        for i := 1; i <= m; i++ {
            if word1[i-1] == word2[j-1] {
                curr[i] = prev[i-1]
            } else {
                curr[i] = 1 + min(
                    prev[i],     // delete
                    curr[i-1],   // insert
                    prev[i-1],   // replace
                )
            }
        }
        prev, curr = curr, make([]int, m+1)
    }
    
    return prev[m]
}
```
- **Pros**: Reduced space usage
- **Cons**: Slightly more complex implementation

#### 3. Wagner-Fischer Algorithm (O(M×N) time, O(N) space)
```go
func minDistanceWagnerFischer(word1, word2 string) int {
    // Advanced algorithm with space optimization
    // Uses only O(N) space instead of O(M×N)
    return -1
}
```
- **Pros**: Optimal space usage
- **Cons**: Very complex to implement

### Extensions for Interviews:
- **Operation Costs**: Different costs for insert, delete, replace
- **Path Reconstruction**: Recover actual edit sequence
- **Multiple Strings**: Edit distance between multiple strings
- **Limited Operations**: Only certain operations allowed
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []struct {
		word1 string
		word2 string
	}{
		{"horse", "ros"},
		{"intention", "execution"},
		{"", ""},
		{"", "abc"},
		{"abc", ""},
		{"abc", "abc"},
		{"kitten", "sitting"},
		{"flaw", "lawn"},
		{"algorithm", "altruistic"},
		{"sunday", "saturday"},
	}
	
	for i, tc := range testCases {
		result1 := minDistance(tc.word1, tc.word2)
		result2 := minDistanceOptimized(tc.word1, tc.word2)
		result3 := minDistanceWithPath(tc.word1, tc.word2)
		
		fmt.Printf("Test Case %d: \"%s\" -> \"%s\"\n", i+1, tc.word1, tc.word2)
		fmt.Printf("  Standard: %d, Optimized: %d, With path: %d\n\n", 
			result1, result2, result3)
	}
}
