package main

import "fmt"

// 10. Regular Expression Matching
// Time: O(M*N), Space: O(M*N)
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	
	// dp[i][j] = true if s[0:i] matches p[0:j]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	
	// Empty string matches empty pattern
	dp[0][0] = true
	
	// Empty string matches patterns like a*, a*b*, a*b*c*
	for j := 2; j <= n; j += 2 {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	
	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				// Direct match or wildcard
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// Two cases:
				// 1. '*' matches zero occurrence of preceding character
				dp[i][j] = dp[i][j-2]
				// 2. '*' matches one or more occurrences of preceding character
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}
	
	return dp[m][n]
}

// Recursive with memoization
func isMatchMemo(s string, p string) bool {
	memo := make(map[string]bool)
	return isMatchHelper(s, p, memo)
}

func isMatchHelper(s string, p string, memo map[string]bool) bool {
	key := s + "|" + p
	if val, exists := memo[key]; exists {
		return val
	}
	
	if len(p) == 0 {
		return len(s) == 0
	}
	
	// Check if next character in pattern is '*'
	if len(p) >= 2 && p[1] == '*' {
		// Case 1: '*' matches zero characters
		if isMatchHelper(s, p[2:], memo) {
			memo[key] = true
			return true
		}
		
		// Case 2: '*' matches one or more characters
		if len(s) > 0 && (p[0] == s[0] || p[0] == '.') {
			if isMatchHelper(s[1:], p, memo) {
				memo[key] = true
				return true
			}
		}
	} else {
		// Normal character or '.' match
		if len(s) > 0 && (p[0] == s[0] || p[0] == '.') {
			if isMatchHelper(s[1:], p[1:], memo) {
				memo[key] = true
				return true
			}
		}
	}
	
	memo[key] = false
	return false
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: 2D Dynamic Programming with Pattern Matching
- **2D DP Table**: dp[i][j] represents match between s[0:i] and p[0:j]
- **Pattern Matching**: Handle characters, wildcards, and repetition operators
- **State Transitions**: Based on current characters in both strings
- **Base Cases**: Empty string matches empty pattern and patterns with *

## 2. PROBLEM CHARACTERISTICS
- **Pattern Matching**: Match string against regular expression pattern
- **Wildcards**: '.' matches any character, '*' matches zero or more of preceding char
- **Sequential Processing**: Process characters in order
- **Complex Dependencies**: Current state depends on previous states

## 3. SIMILAR PROBLEMS
- Wildcard Matching (LeetCode 44) - Similar pattern matching with different wildcards
- Edit Distance (LeetCode 72) - String transformation DP
- Longest Common Subsequence (LeetCode 1143) - Sequence matching DP
- Interleaving String (LeetCode 97) - String combination DP

## 4. KEY OBSERVATIONS
- **Pattern Structure**: '*' always applies to preceding character
- **Empty String**: Can match patterns like a*, a*b*, a*b*c*
- **Character Matching**: Direct match or wildcard '.' match
- **Star Handling**: Two cases - zero occurrences or one/more occurrences

## 5. VARIATIONS & EXTENSIONS
- **Different Wildcards**: Support for different regex operators
- **Multiple Stars**: Handle consecutive stars
- **Group Matching**: Support for character groups [abc]
- **Quantifiers**: Support for +, ?, {n,m} quantifiers

## 6. INTERVIEW INSIGHTS
- Always clarify: "What operators are supported? Empty string handling?"
- Edge cases: empty string, empty pattern, patterns with only stars
- Time complexity: O(M×N) where M=len(s), N=len(p)
- Space complexity: O(M×N) for DP table
- Recursive with memoization is also valid approach

## 7. COMMON MISTAKES
- Not handling empty string with stars pattern correctly
- Wrong star handling logic (missing zero occurrence case)
- Incorrect base case initialization
- Not handling consecutive stars properly
- Off-by-one errors in DP table access

## 8. OPTIMIZATION STRATEGIES
- **2D DP**: O(M×N) time, O(M×N) space
- **Space Optimization**: O(N) space using rolling array
- **Memoization**: O(M×N) time, O(M×N) space for recursion
- **Early Termination**: Stop when pattern is exhausted

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like solving a puzzle with rules:**
- You have two strings: the actual string and the pattern
- You need to determine if they can match following specific rules
- '.' is like a joker card - matches anything
- '*' is like a multiplier - applies to the card before it
- You build a table showing which prefixes can match

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String s, pattern p with '.' and '*' operators
2. **Goal**: Determine if s matches p completely
3. **Rules**: '.' matches any char, '*' matches zero or more of preceding char
4. **Output**: Boolean indicating match success

#### Phase 2: Key Insight Recognition
- **"2D DP natural fit"** → Match prefixes of both strings
- **"Star complexity"** → '*' creates two branching possibilities
- **"Pattern structure"** → '*' always modifies preceding character
- **"Base cases"** → Empty string can match certain patterns

#### Phase 3: Strategy Development
```
Human thought process:
"I need to check if string matches pattern.
I'll build a table where dp[i][j] tells if s[0:i] matches p[0:j].
For each position, I consider the current characters:
- If they match directly, inherit from diagonal
- If pattern has '*', consider two cases:
  1. '*' matches zero chars (skip two pattern chars)
  2. '*' matches one or more chars (stay on pattern, move string)
This covers all possibilities systematically."
```

#### Phase 4: Edge Case Handling
- **Empty string & empty pattern**: Match (true)
- **Empty string & star pattern**: Can match if pattern can represent empty
- **Non-empty string & empty pattern**: No match
- **Pattern ending with star**: Can match empty string

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
s = "aab", p = "c*a*b"

Human thinking:
"I'll build a 4×6 table (including empty prefixes):
dp[0][0] = true (empty matches empty)
Handle stars in first row: dp[0][2] = dp[0][0] (c* can be empty)
dp[0][4] = dp[0][2] (c*a* can be empty)

Now fill the table:
Row 1 (s[0:1] = "a"):
- p[0:1] = "c": no match → dp[1][1] = false
- p[0:2] = "c*": 'a' matches 'c' zero times → dp[1][2] = dp[1][0] = false
- p[0:3] = "c*a": 'a' matches 'a' → dp[1][3] = dp[0][2] = true
- p[0:4] = "c*a*": 'a' matches 'a' → dp[1][4] = dp[1][2] || dp[0][4] = true
- p[0:5] = "c*a*b": 'a' matches 'a' → dp[1][5] = dp[0][4] = true

Continue filling all cells...
Final result: dp[3][5] = true ✓ Pattern matches!"
```

#### Phase 6: Intuition Validation
- **Why 2D DP works**: Each prefix combination has unique match status
- **Why star handling**: Two cases cover all possibilities
- **Why base cases**: Empty prefixes are foundation
- **Why O(M×N)**: Each cell computed once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not greedy?"** → Too complex, DP ensures correctness
2. **"Should I use recursion?"** → Works but needs memoization
3. **"What about star precedence?"** → '*' always applies to preceding char
4. **"Can I optimize space?"** → Yes, use rolling array technique

### Real-World Analogy
**Like solving a word puzzle with wildcards:**
- You have a target word and a pattern with special rules
- '.' is like a blank tile - can be any letter
- '*' is like a "repeat" tile - can repeat the letter before it
- You check progressively larger portions of both
- The table shows which combinations work

### Human-Readable Pseudocode
```
function isMatch(s, p):
    m, n = len(s), len(p)
    dp = create 2D array (m+1)×(n+1) filled with false
    dp[0][0] = true
    
    // Handle patterns that can match empty string
    for j from 2 to n step 2:
        if p[j-1] == '*':
            dp[0][j] = dp[0][j-2]
    
    for i from 1 to m:
        for j from 1 to n:
            if p[j-1] == '.' or p[j-1] == s[i-1]:
                dp[i][j] = dp[i-1][j-1]
            else if p[j-1] == '*':
                // Case 1: '*' matches zero characters
                dp[i][j] = dp[i][j-2]
                // Case 2: '*' matches one or more characters
                if p[j-2] == '.' or p[j-2] == s[i-1]:
                    dp[i][j] = dp[i][j] or dp[i-1][j]
    
    return dp[m][n]
```

### Execution Visualization

### Example: s = "aab", p = "c*a*b"
```
DP Table (rows: s prefixes, columns: p prefixes):
      ""  c   c*  c*a  c*a* c*a*b
""    T   F   T    T    T     T
"a"   F   F   F    T    T     T
"aa"  F   F   F    T    T     T
"aab"  F   F   F    F    T     T

Key transitions:
dp[1][3] = dp[0][2] = T (a matches a in c*a)
dp[2][4] = dp[2][2] || dp[1][4] = F || T = T
dp[3][5] = dp[3][3] || dp[2][5] = F || T = T

Final result: dp[3][5] = T ✓
```

### Key Visualization Points:
- **Star handling**: Two cases - zero or more occurrences
- **Wildcards**: '.' matches any character
- **Base cases**: Empty string handling crucial
- **Dependencies**: Each cell depends on previous states

### Memory Layout Visualization:
```
Star Logic Visualization:
Pattern: c*a*b, String: aab

At dp[2][4] (s="aa", p="c*a*"):
- Case 1: '*' matches zero: dp[2][2] = F
- Case 2: '*' matches more: dp[1][4] = T
- Result: dp[2][4] = F || T = T

This shows how '*' creates branching possibilities.
```

### Time Complexity Breakdown:
- **DP Table**: O(M×N) cells to fill
- **Cell Computation**: O(1) time per cell
- **Total Time**: O(M×N)
- **Space Complexity**: O(M×N) for DP table

### Alternative Approaches:

#### 1. Recursive with Memoization (O(M×N) time, O(M×N) space)
```go
func isMatchRecursive(s, p string) bool {
    memo := make(map[string]bool)
    return isMatchHelper(s, p, memo)
}

func isMatchHelper(s, p string, memo map[string]bool) bool {
    key := s + "|" + p
    if val, exists := memo[key]; exists {
        return val
    }
    
    if len(p) == 0 {
        return len(s) == 0
    }
    
    if len(p) >= 2 && p[1] == '*' {
        // Case 1: '*' matches zero characters
        if isMatchHelper(s, p[2:], memo) {
            memo[key] = true
            return true
        }
        
        // Case 2: '*' matches one or more characters
        if len(s) > 0 && (p[0] == s[0] || p[0] == '.') {
            if isMatchHelper(s[1:], p, memo) {
                memo[key] = true
                return true
            }
        }
    } else {
        // Normal character or '.' match
        if len(s) > 0 && (p[0] == s[0] || p[0] == '.') {
            if isMatchHelper(s[1:], p[1:], memo) {
                memo[key] = true
                return true
            }
        }
    }
    
    memo[key] = false
    return false
}
```
- **Pros**: Intuitive, follows problem description
- **Cons**: Recursion overhead, same complexity as DP

#### 2. Iterative with Space Optimization (O(M×N) time, O(N) space)
```go
func isMatchOptimized(s, p string) bool {
    m, n := len(s), len(p)
    prev := make([]bool, n+1)
    curr := make([]bool, n+1)
    
    prev[0] = true
    
    // Handle patterns that can match empty string
    for j := 2; j <= n; j += 2 {
        if p[j-1] == '*' {
            prev[j] = prev[j-2]
        }
    }
    
    for i := 1; i <= m; i++ {
        curr[0] = false
        for j := 1; j <= n; j++ {
            if p[j-1] == '.' || p[j-1] == s[i-1] {
                curr[j] = prev[j-1]
            } else if p[j-1] == '*' {
                curr[j] = curr[j-2]
                if p[j-2] == '.' || p[j-2] == s[i-1] {
                    curr[j] = curr[j] || prev[j]
                }
            }
        }
        prev, curr = curr, make([]bool, n+1)
    }
    
    return prev[n]
}
```
- **Pros**: Reduced space usage
- **Cons**: More complex implementation

#### 3. NFA Simulation (O(M×N) time, O(N) space)
```go
func isMatchNFA(s, p string) bool {
    // Convert pattern to NFA states
    // Simulate NFA traversal
    // More complex but efficient for certain patterns
    return false
}
```
- **Pros**: Efficient for simple patterns
- **Cons**: Very complex to implement

### Extensions for Interviews:
- **Extended Regex**: Support for +, ?, {n,m} quantifiers
- **Character Classes**: Support for [a-z], [0-9] patterns
- **Grouping**: Support for parentheses and grouping
- **Performance Analysis**: Discuss worst-case scenarios
- **Memory Optimization**: Advanced space optimization techniques
*/
func main() {
	// Test cases
	testCases := []struct {
		s string
		p string
	}{
		{"aa", "a"},
		{"aa", "a*"},
		{"ab", ".*"},
		{"aab", "c*a*b"},
		{"", ".*"},
		{"", ""},
		{"a", ""},
		{"", "a"},
		{"ab", ".*c"},
		{"mississippi", "mis*is*p*."},
		{"ab", ".*.."},
		{"aaa", "a*a"},
		{"ab", ".*..a*"},
	}
	
	for i, tc := range testCases {
		result1 := isMatch(tc.s, tc.p)
		result2 := isMatchMemo(tc.s, tc.p)
		
		fmt.Printf("Test Case %d: \"%s\" vs \"%s\" -> DP: %t, Memo: %t\n", 
			i+1, tc.s, tc.p, result1, result2)
	}
}
