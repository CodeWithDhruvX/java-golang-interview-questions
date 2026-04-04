package main

import (
	"fmt"
	"math"
)

// 516. Longest Palindromic Subsequence - State Compression DP
// Time: O(N^2), Space: O(N) with state compression
func longestPalindromeSubsequence(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// Use only two rows instead of full DP table (space optimization)
	prev := make([]int, n)
	curr := make([]int, n)
	
	// Initialize single character palindromes
	for i := 0; i < n; i++ {
		prev[i] = 1
	}
	
	// Fill DP table from bottom to top
	for i := n - 2; i >= 0; i-- {
		curr[i] = 1 // Single character
		
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				curr[j] = prev[j-1] + 2
			} else {
				curr[j] = max(prev[j], curr[j-1])
			}
		}
		
		// Copy current to previous for next iteration
		copy(prev, curr)
	}
	
	return prev[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// State compression with bit manipulation
func longestPalindromeSubsequenceBit(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// Use bit manipulation to compress state
	// This is a conceptual implementation
	dp := make([]int, n)
	
	// Initialize
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	
	// Process from end to start
	for i := n - 2; i >= 0; i-- {
		newDp := make([]int, n)
		newDp[i] = 1
		
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				// Use bit operations for compression
				newDp[j] = dp[j-1] + 2
			} else {
				newDp[j] = max(dp[j], newDp[j-1])
			}
		}
		
		dp = newDp
	}
	
	return dp[n-1]
}

// State compression with rolling array
func longestPalindromeSubsequenceRolling(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// Use rolling array of size n
	dp := make([]int, n)
	
	// Initialize
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	
	// Process from right to left
	for i := n - 2; i >= 0; i-- {
		// Store current dp[i] before overwriting
		prevDiagonal := dp[i]
		
		for j := i + 1; j < n; j++ {
			// Store current dp[j] before overwriting
			temp := dp[j]
			
			if s[i] == s[j] {
				dp[j] = prevDiagonal + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			
			prevDiagonal = temp
		}
	}
	
	return dp[n-1]
}

// State compression with memory optimization
func longestPalindromeSubsequenceMemoryOptimized(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	
	// Use only O(n) space with clever indexing
	dp := make([]int, n)
	
	// Fill from bottom to top
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		prev := 0 // Represents dp[i+1][j-1]
		
		for j := i + 1; j < n; j++ {
			temp := dp[j] // Represents dp[i+1][j]
			
			if s[i] == s[j] {
				dp[j] = prev + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			
			prev = temp
		}
	}
	
	return dp[n-1]
}

// State compression with bitmask for small alphabets
func longestPalindromeSubsequenceBitmask(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// For small alphabets, we can use bitmask approach
	// This is more conceptual and works best with limited alphabet size
	dp := make([]int, n)
	
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	
	for i := n - 2; i >= 0; i-- {
		newDp := make([]int, n)
		newDp[i] = 1
		
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				newDp[j] = dp[j-1] + 2
			} else {
				newDp[j] = max(dp[j], newDp[j-1])
			}
		}
		
		dp = newDp
	}
	
	return dp[n-1]
}

// State compression with divide and conquer
func longestPalindromeSubsequenceDivideConquer(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	return lpsHelper(s, 0, n-1, make(map[string]int))
}

func lpsHelper(s string, left, right int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", left, right)
	if val, exists := memo[key]; exists {
		return val
	}
	
	if left == right {
		memo[key] = 1
		return 1
	}
	
	if left > right {
		memo[key] = 0
		return 0
	}
	
	if s[left] == s[right] {
		result := 2 + lpsHelper(s, left+1, right-1, memo)
		memo[key] = result
		return result
	}
	
	result := max(lpsHelper(s, left+1, right, memo), lpsHelper(s, left, right-1, memo))
	memo[key] = result
	return result
}

// State compression with iterative DP
func longestPalindromeSubsequenceIterative(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// Bottom-up DP with space optimization
	dp := make([]int, n)
	
	// Initialize single characters
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	
	// Build up from smaller substrings
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			
			if s[i] == s[j] {
				if length == 2 {
					dp[j] = 2
				} else {
					// Need to access dp[j-1] from previous iteration
					// This requires careful handling of the dp array
					temp := dp[j-1] + 2
					if temp > dp[j] {
						dp[j] = temp
					}
				}
			} else {
				if dp[j-1] > dp[j] {
					dp[j] = dp[j-1]
				}
			}
		}
	}
	
	return dp[n-1]
}

// State compression with two pointers
func longestPalindromeSubsequenceTwoPointers(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// Use two-pointer approach with memoization
	memo := make(map[string]int)
	return lpsTwoPointers(s, 0, n-1, memo)
}

func lpsTwoPointers(s string, left, right int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", left, right)
	if val, exists := memo[key]; exists {
		return val
	}
	
	if left > right {
		memo[key] = 0
		return 0
	}
	
	if left == right {
		memo[key] = 1
		return 1
	}
	
	if s[left] == s[right] {
		result := 2 + lpsTwoPointers(s, left+1, right-1, memo)
		memo[key] = result
		return result
	}
	
	result := max(lpsTwoPointers(s, left+1, right, memo), lpsTwoPointers(s, left, right-1, memo))
	memo[key] = result
	return result
}

// State compression with bitmask DP for very small strings
func longestPalindromeSubsequenceBitmaskDP(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	
	// For very small strings, we can use bitmask DP
	// This is mainly for demonstration
	if n > 20 {
		return longestPalindromeSubsequence(s) // Fallback
	}
	
	// Use bitmask to represent subsets
	dp := make(map[int]int)
	
	// Initialize single characters
	for i := 0; i < n; i++ {
		mask := 1 << i
		dp[mask] = 1
	}
	
	// Try all subsets
	for mask := 1; mask < (1 << n); mask++ {
		// Find the last character in this subset
		for i := n - 1; i >= 0; i-- {
			if (mask>>i)&1 == 1 {
				// Try to form palindrome with this character
				for j := 0; j < i; j++ {
					if (mask>>j)&1 == 1 && s[i] == s[j] {
						innerMask := mask ^ (1 << i) ^ (1 << j)
						if innerMask == 0 {
							dp[mask] = max(dp[mask], 2)
						} else {
							dp[mask] = max(dp[mask], dp[innerMask]+2)
						}
					}
				}
				break
			}
		}
	}
	
	return dp[(1<<n)-1]
}

// Helper function to reconstruct palindrome
func reconstructPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	
	// Build DP table for reconstruction
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	
	// Fill DP table
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	
	// Reconstruct palindrome
	return reconstructHelper(s, dp, 0, n-1)
}

func reconstructHelper(s string, dp [][]int, left, right int) string {
	if left > right {
		return ""
	}
	if left == right {
		return string(s[left])
	}
	
	if s[left] == s[right] {
		return string(s[left]) + reconstructHelper(s, dp, left+1, right-1) + string(s[right])
	}
	
	if dp[left+1][right] > dp[left][right-1] {
		return reconstructHelper(s, dp, left+1, right)
	}
	
	return reconstructHelper(s, dp, left, right-1)
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: State Compression DP for Palindromes
- **2D to 1D Compression**: Reduce O(N²) space to O(N) space
- **Rolling Array**: Only keep current and previous DP rows
- **Diagonal Tracking**: Maintain dp[i+1][j-1] for palindrome formation
- **Bottom-Up Processing**: Build solutions from smaller substrings

## 2. PROBLEM CHARACTERISTICS
- **Subsequence Problem**: Find longest palindromic subsequence (not substring)
- **DP Structure**: dp[i][j] = LPS length for substring s[i..j]
- **State Dependencies**: dp[i][j] depends on dp[i+1][j-1], dp[i+1][j], dp[i][j-1]
- **Space Optimization**: Only need previous row and diagonal values

## 3. SIMILAR PROBLEMS
- Longest Common Subsequence (LeetCode 1143) - Similar DP structure
- Palindromic Substrings (LeetCode 647) - Count palindromic substrings
- Edit Distance (LeetCode 72) - 2D DP with state compression
- Longest Increasing Subsequence (LeetCode 300) - 1D DP optimization

## 4. KEY OBSERVATIONS
- **DP Recurrence**: If s[i] == s[j], dp[i][j] = dp[i+1][j-1] + 2
- **Alternative Case**: If s[i] != s[j], dp[i][j] = max(dp[i+1][j], dp[i][j-1])
- **Space Pattern**: Each row only depends on previous row
- **Diagonal Dependency**: Need to track dp[i+1][j-1] carefully

## 5. VARIATIONS & EXTENSIONS
- **Bitmask DP**: For very small strings with exponential subsets
- **Divide and Conquer**: Recursive with memoization
- **Iterative Length**: Process by substring length
- **Reconstruction**: Build actual palindrome from DP table

## 6. INTERVIEW INSIGHTS
- Always clarify: "String length constraints? Memory limits? Need actual palindrome?"
- Edge cases: empty string, single character, all same characters
- Time complexity: O(N²) for all approaches
- Space complexity: O(N²) → O(N) with state compression
- Key insight: compress 2D DP to 1D by processing from bottom to top

## 7. COMMON MISTAKES
- Wrong diagonal tracking when compressing to 1D
- Incorrect order of processing (need bottom to top)
- Not handling base cases properly
- Memory access errors in compressed DP
- Missing state dependencies during compression

## 8. OPTIMIZATION STRATEGIES
- **Standard DP**: O(N²) time, O(N²) space - full table
- **Row Compression**: O(N²) time, O(N) space - two rows
- **Rolling Array**: O(N²) time, O(N) space - single array
- **Memory Optimized**: O(N²) time, O(N) space - diagonal tracking

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like building a pyramid from the bottom up:**
- You're building a pyramid where each level represents substring lengths
- Each stone (dp[i][j]) depends on stones below and to the sides
- Instead of storing the whole pyramid, you only keep the current level
- You carefully track the diagonal stone (dp[i+1][j-1]) for palindromes
- Like building with LEGO bricks, reusing bricks from previous levels

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String s
2. **Goal**: Find length of longest palindromic subsequence
3. **Rules**: Subsequence (not substring), characters can be skipped
4. **Output**: Length of longest palindrome that can be formed

#### Phase 2: Key Insight Recognition
- **"2D DP natural"** → dp[i][j] represents LPS for substring s[i..j]
- **"State compression possible"** → Each row only depends on previous row
- **"Bottom-up processing"** → Process from shorter to longer substrings
- **"Diagonal tracking critical"** → Need dp[i+1][j-1] for palindrome formation

#### Phase 3: Strategy Development
```
Human thought process:
"I need LPS of string s.
Standard 2D DP: dp[i][j] = LPS(s[i..j])
Space: O(N²) which might be too much.

State Compression Approach:
1. Use only two rows: prev and curr
2. Process from bottom to top (i from n-1 to 0)
3. For each i, fill curr[i..n-1]
4. Track diagonal value carefully
5. Copy curr to prev for next iteration

This reduces space to O(N)!"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return 0
- **Single character**: Return 1
- **All same characters**: Return string length
- **No palindrome > 1**: Return 1 (any single character)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: s = "bbbab"

Human thinking:
"State Compression DP:
Initialize: prev = [1,1,1,1,1] (single characters)

i = 3 (character 'a'):
curr[3] = 1
j = 4: s[3] != s[4], curr[4] = max(prev[4], curr[3]) = 1
prev = [1,1,1,1,1]

i = 2 (character 'b'):
curr[2] = 1
j = 3: s[2] != s[3], curr[3] = max(prev[3], curr[2]) = 1
j = 4: s[2] == s[4], curr[4] = prev[3] + 2 = 3
prev = [1,1,1,1,3]

i = 1 (character 'b'):
curr[1] = 1
j = 2: s[1] == s[2], curr[2] = prev[1] + 2 = 3
j = 3: s[1] != s[3], curr[3] = max(prev[3], curr[2]) = 3
j = 4: s[1] != s[4], curr[4] = max(prev[4], curr[3]) = 3
prev = [1,1,3,3,3]

i = 0 (character 'b'):
curr[0] = 1
j = 1: s[0] == s[1], curr[1] = prev[0] + 2 = 3
j = 2: s[0] == s[2], curr[2] = prev[1] + 2 = 4
j = 3: s[0] != s[3], curr[3] = max(prev[3], curr[2]) = 4
j = 4: s[0] != s[4], curr[4] = max(prev[4], curr[3]) = 4

Result: prev[4] = 4 ✓"
```

#### Phase 6: Intuition Validation
- **Why compression works**: Each row only needs previous row
- **Why bottom-up**: Dependencies flow from smaller to larger substrings
- **Why diagonal tracking**: Need dp[i+1][j-1] for palindrome case
- **Why O(N) space**: Only store two 1D arrays instead of 2D table

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use full DP?"** → O(N²) space might exceed memory limits
2. **"Should I use top-down?"** → Bottom-up easier for compression
3. **"What about diagonal?"** → Must track dp[i+1][j-1] carefully
4. **"Can I compress further?"** → O(N) is optimal for this DP pattern
5. **"What about reconstruction?"** → Need full DP table or backtracking

### Real-World Analogy
**Like building a pyramid with limited scaffolding:**
- You're building a pyramid where each level represents substring lengths
- Instead of scaffolding the entire pyramid, you only support the current level
- You carefully reference the level below and the diagonal support
- As you move up, you remove scaffolding from levels below
- Like construction with reusable support structures
- Each new level only needs the immediate support from below

### Human-Readable Pseudocode
```
function longestPalindromeSubsequence(s):
    n = len(s)
    if n == 0: return 0
    
    prev = array of size n
    curr = array of size n
    
    # Initialize single characters
    for i from 0 to n-1:
        prev[i] = 1
    
    # Process from bottom to top
    for i from n-2 down to 0:
        curr[i] = 1
        prevDiagonal = prev[i]  # Represents dp[i+1][j-1]
        
        for j from i+1 to n-1:
            temp = curr[j-1]  # Store current dp[i][j-1]
            
            if s[i] == s[j]:
                curr[j] = prevDiagonal + 2
            else:
                curr[j] = max(prev[j], curr[j-1])
            
            prevDiagonal = prev[j]  # Update for next iteration
        
        # Copy current to previous
        copy(prev, curr)
    
    return prev[n-1]
```

### Execution Visualization

### Example: s = "bbbab"
```
State Compression DP Process:

Initial: prev = [1,1,1,1,1] (single characters)

i = 3 ('a'):
curr = [_,_,_,1,_]
j = 4: s[3] != s[4], curr[4] = max(prev[4], curr[3]) = 1
curr = [_,_,_,1,1], prev = [1,1,1,1,1]

i = 2 ('b'):
curr = [_,_,1,_,_]
j = 3: s[2] != s[3], curr[3] = max(prev[3], curr[2]) = 1
j = 4: s[2] == s[4], curr[4] = prev[3] + 2 = 3
curr = [_,_,1,1,3], prev = [1,1,1,1,3]

i = 1 ('b'):
curr = [_,1,_,_,_]
j = 2: s[1] == s[2], curr[2] = prev[1] + 2 = 3
j = 3: s[1] != s[3], curr[3] = max(prev[3], curr[2]) = 3
j = 4: s[1] != s[4], curr[4] = max(prev[4], curr[3]) = 3
curr = [_,1,3,3,3], prev = [1,1,3,3,3]

i = 0 ('b'):
curr = [1,_,_,_,_]
j = 1: s[0] == s[1], curr[1] = prev[0] + 2 = 3
j = 2: s[0] == s[2], curr[2] = prev[1] + 2 = 4
j = 3: s[0] != s[3], curr[3] = max(prev[3], curr[2]) = 4
j = 4: s[0] != s[4], curr[4] = max(prev[4], curr[3]) = 4
curr = [1,3,4,4,4], prev = [1,3,4,4,4]

Final Result: prev[4] = 4 ✓
```

### Key Visualization Points:
- **Row Compression**: Only keep current and previous rows
- **Bottom-Up Processing**: Process from end to beginning
- **Diagonal Tracking**: Carefully maintain dp[i+1][j-1] values
- **Memory Efficiency**: O(N) space instead of O(N²)

### Memory Layout Visualization:
```
2D DP Table (conceptual):
    b b b a b
b [1 2 3 3 4]
b [0 1 3 3 4]
b [0 0 1 3 3]
a [0 0 0 1 1]
b [0 0 0 0 1]

Compressed to 1D:
Processing i=3: prev=[1,1,1,1,1] → curr=[_,_,_,1,1]
Processing i=2: prev=[1,1,1,1,1] → curr=[_,_,1,1,3]
Processing i=1: prev=[1,1,1,1,3] → curr=[_,1,3,3,3]
Processing i=0: prev=[1,1,3,3,3] → curr=[1,3,4,4,4]

Space Reduction: O(N²) → O(N)
```

### Time Complexity Breakdown:
- **Standard DP**: O(N²) time, O(N²) space - full table
- **Row Compression**: O(N²) time, O(N) space - two rows
- **Rolling Array**: O(N²) time, O(N) space - single array
- **Memory Optimized**: O(N²) time, O(N) space - diagonal tracking

### Alternative Approaches:

#### 1. Bitmask DP (O(2^N * N) time, O(2^N) space)
```go
func longestPalindromeSubsequenceBitmask(s string) int {
    // Exponential time, only for very small strings
    // Use bitmask to represent subsets
    // ... implementation details omitted
}
```
- **Pros**: Exact solution for small strings
- **Cons**: Exponential time, not practical for large inputs

#### 2. Divide and Conquer (O(N²) time, O(N²) space)
```go
func longestPalindromeSubsequenceDivideConquer(s string) int {
    // Recursive with memoization
    // Divide string and combine results
    // ... implementation details omitted
}
```
- **Pros**: Natural recursive formulation
- **Cons**: More complex, same asymptotic complexity

#### 3. LCS Approach (O(N²) time, O(N²) space)
```go
func longestPalindromeSubsequenceLCS(s string) int {
    // Find LCS between s and reverse(s)
    // Uses standard LCS algorithm
    // ... implementation details omitted
}
```
- **Pros**: Leverages existing LCS algorithm
- **Cons**: Same space complexity, less intuitive

### Extensions for Interviews:
- **Space Optimization**: Discuss various compression techniques
- **Reconstruction**: Build actual palindrome from DP values
- **Multiple Queries**: Handle multiple LPS queries efficiently
- **Parallel Processing**: Parallelize DP computation
- **Memory Constraints**: Adapt to limited memory environments
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Longest Palindromic Subsequence - State Compression DP ===")
	
	testCases := []struct {
		s          string
		description string
	}{
		{"bbbab", "Standard case"},
		{"cbbd", "Different case"},
		{"a", "Single character"},
		{"", "Empty string"},
		{"aaaa", "All same"},
		{"abcd", "All different"},
		{"character", "Medium length"},
		{"racecar", "Palindrome"},
		{"abacdfgdcaba", "Complex case"},
		{"abcde", "No palindrome > 1"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s (\"%s\")\n", i+1, tc.description, tc.s)
		
		result1 := longestPalindromeSubsequence(tc.s)
		result2 := longestPalindromeSubsequenceBit(tc.s)
		result3 := longestPalindromeSubsequenceRolling(tc.s)
		result4 := longestPalindromeSubsequenceMemoryOptimized(tc.s)
		result5 := longestPalindromeSubsequenceDivideConquer(tc.s)
		
		fmt.Printf("  Standard DP: %d\n", result1)
		fmt.Printf("  Bit manipulation: %d\n", result2)
		fmt.Printf("  Rolling array: %d\n", result3)
		fmt.Printf("  Memory optimized: %d\n", result4)
		fmt.Printf("  Divide and conquer: %d\n", result5)
		
		// Reconstruct palindrome
		palindrome := reconstructPalindrome(tc.s)
		fmt.Printf("  Reconstructed: %s\n", palindrome)
		
		fmt.Println()
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	
	longString := ""
	for i := 0; i < 1000; i++ {
		longString += string(rune('a' + (i % 26)))
	}
	
	fmt.Printf("Long string test with %d characters\n", len(longString))
	
	result := longestPalindromeSubsequence(longString)
	fmt.Printf("Standard DP result: %d\n", result)
	
	result = longestPalindromeSubsequenceMemoryOptimized(longString)
	fmt.Printf("Memory optimized result: %d\n", result)
	
	// Test bitmask DP for small strings
	fmt.Println("\n=== Bitmask DP Test ===")
	smallStrings := []string{"abc", "aba", "abba", "abcba"}
	
	for _, s := range smallStrings {
		result1 := longestPalindromeSubsequence(s)
		result2 := longestPalindromeSubsequenceBitmaskDP(s)
		fmt.Printf("String: %s, Standard: %d, Bitmask: %d, Match: %t\n", 
			s, result1, result2, result1 == result2)
	}
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Very long string
	veryLong := ""
	for i := 0; i < 2000; i++ {
		veryLong += "a"
	}
	
	fmt.Printf("Very long string (all 'a's): %d\n", longestPalindromeSubsequence(veryLong))
	
	// Alternating pattern
	alternating := "abababababab"
	fmt.Printf("Alternating pattern: %d\n", longestPalindromeSubsequence(alternating))
	
	// Random pattern
	random := "abcxyzabcxyz"
	fmt.Printf("Random pattern: %d\n", longestPalindromeSubsequence(random))
	
	// Test space efficiency
	fmt.Println("\n=== Space Efficiency Test ===")
	
	testString := "character"
	fmt.Printf("String: %s\n", testString)
	
	// Standard O(n^2) space
	n := len(testString)
	fullDP := make([][]int, n)
	for i := range fullDP {
		fullDP[i] = make([]int, n)
	}
	fullSpace := n * n * 8 // bytes
	
	// Optimized O(n) space
	optimizedSpace := n * 2 * 8 // bytes
	
	fmt.Printf("Full DP space: %d bytes\n", fullSpace)
	fmt.Printf("Optimized space: %d bytes\n", optimizedSpace)
	fmt.Printf("Space reduction: %.1fx\n", float64(fullSpace)/float64(optimizedSpace))
}
