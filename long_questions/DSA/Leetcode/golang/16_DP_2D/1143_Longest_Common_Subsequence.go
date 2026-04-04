package main

import "fmt"

// 1143. Longest Common Subsequence
// Time: O(M*N), Space: O(M*N)
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	
	// dp[i][j] = length of LCS of text1[0:i] and text2[0:j]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	
	return dp[m][n]
}

// Space optimized version: O(min(M,N)) space
func longestCommonSubsequenceOptimized(text1 string, text2 string) int {
	// Ensure text1 is the shorter string for space optimization
	if len(text1) > len(text2) {
		text1, text2 = text2, text1
	}
	
	m, n := len(text1), len(text2)
	prev := make([]int, m+1)
	current := make([]int, m+1)
	
	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if text1[i-1] == text2[j-1] {
				current[i] = prev[i-1] + 1
			} else {
				current[i] = max(prev[i], current[i-1])
			}
		}
		// Swap prev and current
		prev, current = current, prev
	}
	
	return prev[m]
}

// Function to actually reconstruct the LCS
func longestCommonSubsequenceReconstruct(text1 string, text2 string) string {
	m, n := len(text1), len(text2)
	
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	
	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	
	// Reconstruct the LCS
	var result []byte
	i, j := m, n
	
	for i > 0 && j > 0 {
		if text1[i-1] == text2[j-1] {
			result = append(result, text1[i-1])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}
	
	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	
	return string(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: 2D Dynamic Programming for Sequence Matching
- **2D DP Table**: dp[i][j] represents length of LCS of text1[0:i] and text2[0:j]
- **Sequence Matching**: Find longest common subsequence (not necessarily contiguous)
- **State Transitions**: Based on character comparison
- **Base Cases**: Empty strings have LCS length 0

## 2. PROBLEM CHARACTERISTICS
- **Subsequence Problem**: Find longest sequence appearing in both strings
- **Non-Contiguous**: Characters can be from different positions
- **Order Preservation**: Relative order must be maintained
- **Optimal Substructure**: Solution depends on smaller subproblems

## 3. SIMILAR PROBLEMS
- Edit Distance (LeetCode 72) - String transformation DP
- Regular Expression Matching (LeetCode 10) - Pattern matching DP
- Longest Common Substring (LeetCode 518) - Contiguous sequence matching
- Delete Operation for Two Strings (LeetCode 583) - String operations DP

## 4. KEY OBSERVATIONS
- **Character Match**: If characters match, extend LCS by 1
- **Character Mismatch**: Take maximum of excluding current character from either string
- **DP Recurrence**: dp[i][j] = dp[i-1][j-1] + 1 if match, else max(dp[i-1][j], dp[i][j-1])
- **Base Cases**: Empty strings have LCS length 0

## 5. VARIATIONS & EXTENSIONS
- **Multiple Strings**: LCS of more than two strings
- **Weighted LCS**: Characters have different weights
- **Path Reconstruction**: Recover actual LCS sequence
- **Space Optimization**: Use rolling arrays

## 6. INTERVIEW INSIGHTS
- Always clarify: "Do you need the actual sequence or just length?"
- Edge cases: empty strings, identical strings, no common characters
- Time complexity: O(M×N) where M=len(text1), N=len(text2)
- Space complexity: O(M×N) for DP table
- Space optimization possible using rolling arrays

## 7. COMMON MISTAKES
- Not handling empty string base cases correctly
- Wrong DP recurrence (using +1 instead of max for mismatch)
- Off-by-one errors in DP table access
- Confusing subsequence with substring
- Forgetting to reconstruct LCS when needed

## 8. OPTIMIZATION STRATEGIES
- **2D DP**: O(M×N) time, O(M×N) space
- **Space Optimization**: O(min(M,N)) space using rolling arrays
- **Path Reconstruction**: O(M×N) time, O(M×N) space
- **Early Termination**: Not applicable (need full table)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding common DNA between two species:**
- You have two DNA sequences (text1, text2)
- You want to find the longest common subsequence
- Subsequence means characters in order but not necessarily adjacent
- You build a table showing LCS length for all prefixes
- The final cell gives the maximum LCS length

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Two strings text1, text2
2. **Goal**: Find length of longest common subsequence
3. **Constraint**: Characters must appear in same relative order
4. **Output**: Length of LCS (can reconstruct sequence if needed)

#### Phase 2: Key Insight Recognition
- **"2D DP natural fit"** → Compare prefixes of both strings
- **"Subsequence property"** → Order matters, not contiguity
- **"Character matching"** → Match extends LCS, mismatch takes max of previous
- **"Optimal substructure"** → Current state depends on smaller prefixes

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find longest common subsequence.
I'll build a table where dp[i][j] shows LCS length for text1[0:i] and text2[0:j].
For each position, I compare current characters:
- If they match, extend LCS by 1 → dp[i-1][j-1] + 1
- If they differ, take the best without current char from either string:
  max(dp[i-1][j], dp[i][j-1])
This ensures I always have the optimal LCS length for each prefix pair."
```

#### Phase 4: Edge Case Handling
- **Empty strings**: LCS length is 0
- **Identical strings**: LCS length is string length
- **No common characters**: LCS length is 0
- **Single character**: Simple comparison

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
text1 = "abcde", text2 = "ace"

Human thinking:
"I'll build a 6×4 table (including empty prefixes):
Initialize first row and column with 0s.

Fill dp[1][1] (a vs a):
- Characters match → dp[0][0] + 1 = 1

Fill dp[2][2] (ab vs ac):
- Compare 'b' vs 'c' (no match)
- Take max(dp[1][2], dp[2][1]) = max(1,1) = 1

Fill dp[3][3] (abc vs ace):
- Compare 'c' vs 'e' (no match)
- Take max(dp[2][3], dp[3][2]) = max(1,2) = 2

Continue filling...
Final result: dp[5][3] = 3 ✓ LCS length"
```

#### Phase 6: Intuition Validation
- **Why 2D DP works**: Each prefix combination has unique LCS length
- **Why subsequence logic**: Order preservation is key requirement
- **Why max for mismatch**: Can't extend LCS, so take best previous
- **Why O(M×N)**: Each cell computed once with O(1) operations

### Common Human Pitfalls & How to Avoid Them
1. **"Why not greedy?"** → Local choices don't guarantee global optimum
2. **"Should I use recursion?"** → Works but needs memoization
3. **"What about reconstruction?"** → Need to backtrack from DP table
4. **"Can I optimize space?"** → Yes, use rolling array technique

### Real-World Analogy
**Like finding common ancestry between two family trees:**
- You have two family trees with member names
- You want to find the longest sequence of common ancestors
- The sequence must maintain chronological order
- You compare progressively larger portions of both trees
- The table shows common ancestors for each subtree pair

### Human-Readable Pseudocode
```
function longestCommonSubsequence(text1, text2):
    m, n = len(text1), len(text2)
    dp = create 2D array (m+1)×(n+1) filled with 0
    
    for i from 1 to m:
        for j from 1 to n:
            if text1[i-1] == text2[j-1]:
                dp[i][j] = dp[i-1][j-1] + 1
            else:
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
    
    return dp[m][n]

function reconstructLCS(text1, text2, dp):
    i, j = len(text1), len(text2)
    lcs = empty string
    
    while i > 0 and j > 0:
        if text1[i-1] == text2[j-1]:
            lcs = text1[i-1] + lcs
            i--, j--
        else if dp[i-1][j] > dp[i][j-1]:
            i--
        else:
            j--
    
    return lcs
```

### Execution Visualization

### Example: text1 = "abcde", text2 = "ace"
```
DP Table (rows: text1 prefixes, columns: text2 prefixes):
      ""  a   c   ac  ace
""    0   0   0   0   0
"a"    0   1   1   1   1
"ab"   0   1   1   2   2
"abc"  0   1   2   2   2
"abcd" 0   1   2   3   3
"abcde" 0   1   2   3   3

Key transitions:
dp[1][1] (a vs a): match → dp[0][0] + 1 = 1
dp[2][2] (ab vs ac): max(dp[1][2], dp[2][1]) = max(1,1) = 1
dp[3][3] (abc vs ace): max(dp[2][3], dp[3][2]) = max(2,2) = 2
dp[5][3] (abcde vs ace): max(dp[4][3], dp[5][2]) = max(3,2) = 3

Final result: dp[5][3] = 3 ✓ LCS = "ace"
```

### Key Visualization Points:
- **Character Match**: Extends LCS by 1
- **Character Mismatch**: Takes maximum of previous states
- **Subsequence Property**: Order preserved, not necessarily contiguous
- **DP Recurrence**: Each cell depends on top and left neighbors

### Memory Layout Visualization:
```
LCS Reconstruction Visualization:
text1 = "abcde", text2 = "ace", dp[5][3] = 3

Backtrack from dp[5][3]:
1. 'e' vs 'e' (match) → add 'e', move to dp[4][2]
2. 'd' vs 'c' (no match) → dp[3][2] > dp[4][2], move to dp[3][2]
3. 'c' vs 'c' (match) → add 'c', move to dp[2][1]
4. 'b' vs 'a' (no match) → dp[1][1] > dp[2][1], move to dp[1][1]
5. 'a' vs 'a' (match) → add 'a', move to dp[0][0]

Reconstructed LCS: "ace" (reverse of construction)
```

### Time Complexity Breakdown:
- **DP Table**: O(M×N) cells to fill
- **Cell Computation**: O(1) time per cell
- **Total Time**: O(M×N)
- **Space Complexity**: O(M×N) for DP table

### Alternative Approaches:

#### 1. Recursive with Memoization (O(M×N) time, O(M×N) space)
```go
func longestCommonSubsequenceRecursive(text1, text2 string) int {
    memo := make(map[string]int)
    return lcsHelper(text1, text2, len(text1), len(text2), memo)
}

func lcsHelper(t1, t2 string, i, j int, memo map[string]int) int {
    if i == 0 || j == 0 {
        return 0
    }
    
    key := fmt.Sprintf("%d,%d", i, j)
    if val, exists := memo[key]; exists {
        return val
    }
    
    if t1[i-1] == t2[j-1] {
        result := 1 + lcsHelper(t1, t2, i-1, j-1, memo)
        memo[key] = result
        return result
    }
    
    result := max(
        lcsHelper(t1, t2, i-1, j, memo),
        lcsHelper(t1, t2, i, j-1, memo),
    )
    
    memo[key] = result
    return result
}
```
- **Pros**: Intuitive, follows problem description
- **Cons**: Recursion overhead, same complexity as DP

#### 2. Iterative with Space Optimization (O(M×N) time, O(min(M,N)) space)
```go
func longestCommonSubsequenceOptimized(text1, text2 string) int {
    // Ensure text1 is shorter for space optimization
    if len(text1) > len(text2) {
        text1, text2 = text2, text1
    }
    
    m, n := len(text1), len(text2)
    prev := make([]int, m+1)
    curr := make([]int, m+1)
    
    for j := 1; j <= n; j++ {
        for i := 1; i <= m; i++ {
            if text1[i-1] == text2[j-1] {
                curr[i] = prev[i-1] + 1
            } else {
                curr[i] = max(prev[i], curr[i-1])
            }
        }
        prev, curr = curr, make([]int, m+1)
    }
    
    return prev[m]
}
```
- **Pros**: Reduced space usage
- **Cons**: Slightly more complex implementation

#### 3. Bitset Optimization (O(M×N/wordsize) time, O(N) space)
```go
func longestCommonSubsequenceBitset(text1, text2 string) int {
    // Advanced optimization for character sets
    // Uses bit operations for faster processing
    return -1
}
```
- **Pros**: Very fast for character-limited alphabets
- **Cons**: Complex implementation, limited applicability

### Extensions for Interviews:
- **Multiple Strings**: LCS of more than two strings
- **Weighted LCS**: Characters have different weights
- **Path Reconstruction**: Recover actual LCS sequence
- **Space Optimization**: Advanced space optimization techniques
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []struct {
		text1 string
		text2 string
	}{
		{"abcde", "ace"},
		{"abc", "abc"},
		{"abc", "def"},
		{"", ""},
		{"", "abc"},
		{"abc", ""},
		{"AGGTAB", "GXTXAYB"},
		{"abcdef", "acbcf"},
		{"XMJYAUZ", "MZJAWXU"},
		{"ABCD", "ABDC"},
	}
	
	for i, tc := range testCases {
		result1 := longestCommonSubsequence(tc.text1, tc.text2)
		result2 := longestCommonSubsequenceOptimized(tc.text1, tc.text2)
		result3 := longestCommonSubsequenceReconstruct(tc.text1, tc.text2)
		
		fmt.Printf("Test Case %d: \"%s\", \"%s\"\n", i+1, tc.text1, tc.text2)
		fmt.Printf("  Length: %d, Optimized: %d, LCS: \"%s\"\n\n", 
			result1, result2, result3)
	}
}
