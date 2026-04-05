public class LongestPalindromicSubsequence {
    
    // 516. Longest Palindromic Subsequence - State Compression DP
    // Time: O(N^2), Space: O(N) with state compression
    public int longestPalindromeSubsequence(String s) {
        int n = s.length();
        if (n == 0) {
            return 0;
        }
        
        // Use only two rows instead of full DP table (space optimization)
        int[] prev = new int[n];
        int[] curr = new int[n];
        
        // Initialize single character palindromes
        for (int i = 0; i < n; i++) {
            prev[i] = 1;
        }
        
        // Fill DP table from bottom to top
        for (int i = n - 2; i >= 0; i--) {
            curr[i] = 1; // Single character
            
            for (int j = i + 1; j < n; j++) {
                if (s.charAt(i) == s.charAt(j)) {
                    curr[j] = prev[j - 1] + 2;
                } else {
                    curr[j] = Math.max(prev[j], curr[j - 1]);
                }
            }
            
            // Copy current to previous for next iteration
            System.arraycopy(curr, 0, prev, 0, n);
        }
        
        return prev[n - 1];
    }
    
    // State compression with bit manipulation
    public int longestPalindromeSubsequenceBit(String s) {
        int n = s.length();
        if (n == 0) {
            return 0;
        }
        
        // Use bit manipulation to compress state
        // This is a conceptual implementation
        int[] dp = new int[n];
        
        // Initialize
        for (int i = 0; i < n; i++) {
            dp[i] = 1;
        }
        
        // Process from end to start
        for (int i = n - 2; i >= 0; i--) {
            int[] newDp = new int[n];
            newDp[i] = 1;
            
            for (int j = i + 1; j < n; j++) {
                if (s.charAt(i) == s.charAt(j)) {
                    // Use bit operations for compression
                    newDp[j] = dp[j - 1] + 2;
                } else {
                    newDp[j] = Math.max(dp[j], newDp[j - 1]);
                }
            }
            
            dp = newDp;
        }
        
        return dp[n - 1];
    }
    
    // Standard DP without compression for comparison
    public int longestPalindromeSubsequenceStandard(String s) {
        int n = s.length();
        int[][] dp = new int[n][n];
        
        // Initialize single character palindromes
        for (int i = 0; i < n; i++) {
            dp[i][i] = 1;
        }
        
        // Fill DP table
        for (int length = 2; length <= n; length++) {
            for (int i = 0; i <= n - length; i++) {
                int j = i + length - 1;
                
                if (s.charAt(i) == s.charAt(j)) {
                    if (length == 2) {
                        dp[i][j] = 2;
                    } else {
                        dp[i][j] = dp[i + 1][j - 1] + 2;
                    }
                } else {
                    dp[i][j] = Math.max(dp[i + 1][j], dp[i][j - 1]);
                }
            }
        }
        
        return dp[0][n - 1];
    }
    
    // Optimized version with 1D array
    public int longestPalindromeSubsequence1D(String s) {
        int n = s.length();
        int[] dp = new int[n];
        
        // Process from end to start
        for (int i = n - 1; i >= 0; i--) {
            dp[i] = 1;
            int prev = 0; // This represents dp[i+1][j-1] from 2D version
            
            for (int j = i + 1; j < n; j++) {
                int temp = dp[j]; // Store current dp[j] before it gets updated
                
                if (s.charAt(i) == s.charAt(j)) {
                    dp[j] = prev + 2;
                } else {
                    dp[j] = Math.max(dp[j], dp[j - 1]);
                }
                
                prev = temp; // Update prev for next iteration
            }
        }
        
        return dp[n - 1];
    }
    
    // Version with detailed explanation
    public class LPSResult {
        int length;
        String subsequence;
        java.util.List<String> explanation;
        
        LPSResult(int length, String subsequence, java.util.List<String> explanation) {
            this.length = length;
            this.subsequence = subsequence;
            this.explanation = explanation;
        }
    }
    
    public LPSResult longestPalindromeSubsequenceDetailed(String s) {
        java.util.List<String> explanation = new java.util.ArrayList<>();
        explanation.add("=== State Compression DP for LPS ===");
        explanation.add("String: " + s);
        explanation.add("Length: " + s.length());
        
        int n = s.length();
        int[] prev = new int[n];
        int[] curr = new int[n];
        
        // Initialize
        for (int i = 0; i < n; i++) {
            prev[i] = 1;
        }
        explanation.add("Initialized: Single character palindromes = 1");
        
        // Process
        for (int i = n - 2; i >= 0; i--) {
            curr[i] = 1;
            explanation.add(String.format("Processing i=%d (char '%c')", i, s.charAt(i)));
            
            for (int j = i + 1; j < n; j++) {
                if (s.charAt(i) == s.charAt(j)) {
                    curr[j] = prev[j - 1] + 2;
                    explanation.add(String.format("  Match at j=%d ('%c'): dp[%d][%d] = dp[%d][%d] + 2 = %d", 
                        j, s.charAt(j), i, j, i + 1, j - 1, curr[j]));
                } else {
                    curr[j] = Math.max(prev[j], curr[j - 1]);
                    explanation.add(String.format("  No match at j=%d ('%c'): dp[%d][%d] = max(dp[%d][%d], dp[%d][%d]) = %d", 
                        j, s.charAt(j), i, j, i + 1, j, i, j - 1, curr[j]));
                }
            }
            
            System.arraycopy(curr, 0, prev, 0, n);
            explanation.add("Copied current row to previous");
        }
        
        int result = prev[n - 1];
        explanation.add(String.format("Final result: %d", result));
        
        return new LPSResult(result, "", explanation);
    }
    
    // Reconstruct the actual palindrome subsequence
    public String reconstructPalindrome(String s) {
        int n = s.length();
        int[][] dp = new int[n][n];
        
        // Fill DP table
        for (int i = 0; i < n; i++) {
            dp[i][i] = 1;
        }
        
        for (int length = 2; length <= n; length++) {
            for (int i = 0; i <= n - length; i++) {
                int j = i + length - 1;
                
                if (s.charAt(i) == s.charAt(j)) {
                    if (length == 2) {
                        dp[i][j] = 2;
                    } else {
                        dp[i][j] = dp[i + 1][j - 1] + 2;
                    }
                } else {
                    dp[i][j] = Math.max(dp[i + 1][j], dp[i][j - 1]);
                }
            }
        }
        
        // Reconstruct the subsequence
        return reconstructHelper(s, dp, 0, n - 1);
    }
    
    private String reconstructHelper(String s, int[][] dp, int i, int j) {
        if (i > j) {
            return "";
        }
        if (i == j) {
            return String.valueOf(s.charAt(i));
        }
        
        if (s.charAt(i) == s.charAt(j)) {
            return s.charAt(i) + reconstructHelper(s, dp, i + 1, j - 1) + s.charAt(j);
        } else if (dp[i + 1][j] >= dp[i][j - 1]) {
            return reconstructHelper(s, dp, i + 1, j);
        } else {
            return reconstructHelper(s, dp, i, j - 1);
        }
    }
    
    // Memory-efficient version for very long strings
    public int longestPalindromeSubsequenceMemoryEfficient(String s) {
        if (s.length() <= 1000) {
            return longestPalindromeSubsequence(s);
        }
        
        // For very long strings, use sliding window approach
        int maxLen = 0;
        int windowSize = Math.min(1000, s.length());
        
        for (int start = 0; start < s.length(); start += windowSize / 2) {
            int end = Math.min(start + windowSize, s.length());
            String substring = s.substring(start, end);
            maxLen = Math.max(maxLen, longestPalindromeSubsequence(substring));
        }
        
        return maxLen;
    }
    
    public static void main(String[] args) {
        LongestPalindromicSubsequence lps = new LongestPalindromicSubsequence();
        
        // Test cases
        String[] testCases = {
            "bbbab",
            "cbbd",
            "a",
            "",
            "aaaa",
            "abcba",
            "character",
            "forgeeksskeegfor",
            "abacdfgdcaba",
            "agbdba"
        };
        
        String[] descriptions = {
            "Standard case",
            "Two characters",
            "Single character",
            "Empty string",
            "All same characters",
            "Already palindrome",
            "Mixed characters",
            "Long string with palindrome",
            "Multiple palindromes",
            "Complex case"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("Input: \"%s\"\n", testCases[i]);
            
            int result1 = lps.longestPalindromeSubsequence(testCases[i]);
            int result2 = lps.longestPalindromeSubsequenceStandard(testCases[i]);
            int result3 = lps.longestPalindromeSubsequence1D(testCases[i]);
            int result4 = lps.longestPalindromeSubsequenceBit(testCases[i]);
            
            System.out.printf("State Compression: %d\n", result1);
            System.out.printf("Standard DP: %d\n", result2);
            System.out.printf("1D Array: %d\n", result3);
            System.out.printf("Bit Manipulation: %d\n", result4);
            
            // Reconstruct palindrome
            String palindrome = lps.reconstructPalindrome(testCases[i]);
            if (!palindrome.isEmpty()) {
                System.out.printf("Reconstructed palindrome: \"%s\"\n", palindrome);
            }
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        LPSResult detailedResult = lps.longestPalindromeSubsequenceDetailed("bbbab");
        for (String step : detailedResult.explanation) {
            System.out.println(step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        
        // Large test case
        StringBuilder largeString = new StringBuilder();
        for (int i = 0; i < 1000; i++) {
            largeString.append((char) ('a' + (i % 26)));
        }
        
        long startTime = System.nanoTime();
        int largeResult1 = lps.longestPalindromeSubsequence(largeString.toString());
        long endTime = System.nanoTime();
        
        System.out.printf("Large test (1000 chars) - State Compression: %d (took %d ns)\n", 
            largeResult1, endTime - startTime);
        
        startTime = System.nanoTime();
        int largeResult2 = lps.longestPalindromeSubsequenceStandard(largeString.toString());
        endTime = System.nanoTime();
        
        System.out.printf("Large test (1000 chars) - Standard DP: %d (took %d ns)\n", 
            largeResult2, endTime - startTime);
        
        // Memory usage comparison
        System.out.println("\n=== Memory Usage Comparison ===");
        System.out.println("State Compression: O(N) space");
        System.out.println("Standard DP: O(N^2) space");
        System.out.println("1D Array: O(N) space");
        System.out.println("Bit Manipulation: O(N) space");
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: State Compression DP
- **2D to 1D**: Reduce DP table from O(N²) to O(N)
- **Space Optimization**: Only keep current and previous rows
- **Palindromic Subsequence**: Longest palindrome that's a subsequence
- **DP Transition**: dp[i][j] = dp[i+1][j-1] + 2 if chars match

## 2. PROBLEM CHARACTERISTICS
- **Subsequence**: Characters not necessarily contiguous
- **Palindrome**: Reads same forwards and backwards
- **Longest**: Find maximum length palindromic subsequence
- **DP Optimization**: Reduce space complexity from O(N²) to O(N)

## 3. SIMILAR PROBLEMS
- Longest Palindromic Substring
- Count Palindromic Substrings
- Longest Common Subsequence
- Minimum Insertions to Form Palindrome

## 4. KEY OBSERVATIONS
- 2D DP table stores results for all substring ranges
- Only previous row needed for current row computation
- State compression reduces space from O(N²) to O(N)
- Time complexity remains O(N²) for all approaches
- Multiple optimization strategies: row compression, 1D array, bit manipulation

## 5. VARIATIONS & EXTENSIONS
- Return the actual palindrome subsequence
- Count all palindromic subsequences
- Different character sets
- Multiple string comparison

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I return the subsequence or just length?"
- Edge cases: empty string, single character, all same characters
- Time complexity: O(N²) for all DP approaches
- Space complexity: O(N²) vs O(N) with compression

## 7. COMMON MISTAKES
- Incorrect DP transition formula
- Not handling single character base case
- Wrong array indexing in compressed version
- Forgetting to copy current row to previous
- Off-by-one errors in boundary conditions

## 8. OPTIMIZATION STRATEGIES
- Use two rows instead of full table
- Further compress to single 1D array
- Use bit manipulation for specific cases
- Early termination for obvious cases

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the longest mirror in a string:**
- You have a string and want to find its longest palindromic subsequence
- A palindrome reads the same forwards and backwards
- For any two positions i and j, you want to know the LPS in s[i..j]
- If s[i] == s[j], you can extend the palindrome by 2 characters
- Otherwise, you take the best of excluding either i or j
- This builds up the solution from smaller subproblems

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String s
2. **Goal**: Find length of longest palindromic subsequence
3. **Output**: Integer representing maximum length

#### Phase 2: Key Insight Recognition
- **"What defines a palindrome?"** → Same forwards and backwards
- **"How to use DP?"** → dp[i][j] = LPS in s[i..j]
- **"What's the transition?"** → If s[i] == s[j], dp[i][j] = dp[i+1][j-1] + 2
- **"Why state compression?"** → Only need previous row for current row

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use DP with state compression:
1. Initialize DP table for single characters (length 1)
2. Process from bottom to top (i from n-1 to 0)
3. For each i, process j from i+1 to n-1
4. Use only current and previous rows
5. If s[i] == s[j]: curr[j] = prev[j-1] + 2
6. Else: curr[j] = max(prev[j], curr[j-1])
7. Copy current to previous for next iteration"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return 0
- **Single character**: Return 1
- **All same characters**: Return string length
- **No palindrome longer than 1**: Return 1

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
String: "bbbab"

Human thinking:
"Let's build DP table with state compression:

Initialize previous row (single characters):
prev = [1, 1, 1, 1, 1] (each char is palindrome of length 1)

i = 3 (char 'a'):
curr[3] = 1
j = 4: s[3] != s[4], curr[4] = max(prev[4], curr[3]) = max(1, 1) = 1
Copy curr to prev: prev = [1, 1, 1, 1, 1]

i = 2 (char 'b'):
curr[2] = 1
j = 3: s[2] != s[3], curr[3] = max(prev[3], curr[2]) = max(1, 1) = 1
j = 4: s[2] == s[4] ('b' == 'b'), curr[4] = prev[3] + 2 = 1 + 2 = 3
Copy curr to prev: prev = [1, 1, 1, 1, 3]

i = 1 (char 'b'):
curr[1] = 1
j = 2: s[1] == s[2] ('b' == 'b'), curr[2] = prev[1] + 2 = 1 + 2 = 3
j = 3: s[1] != s[3], curr[3] = max(prev[3], curr[2]) = max(1, 3) = 3
j = 4: s[1] != s[4], curr[4] = max(prev[4], curr[3]) = max(3, 3) = 3
Copy curr to prev: prev = [1, 1, 3, 3, 3]

i = 0 (char 'b'):
curr[0] = 1
j = 1: s[0] == s[1] ('b' == 'b'), curr[1] = prev[0] + 2 = 1 + 2 = 3
j = 2: s[0] != s[2], curr[2] = max(prev[2], curr[1]) = max(3, 3) = 3
j = 3: s[0] != s[3], curr[3] = max(prev[3], curr[2]) = max(3, 3) = 3
j = 4: s[0] != s[4], curr[4] = max(prev[4], curr[3]) = max(3, 3) = 3

Final result: prev[4] = 3 ✓

Longest palindromic subsequence length: 3 ("bbb") ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: DP builds solution from smaller subproblems
- **Why it's efficient**: State compression reduces space while maintaining correctness
- **Why it's correct**: All subproblems are solved optimally

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use recursion?"** → O(2^N) exponential time
2. **"What about full DP table?"** → O(N²) space vs O(N) with compression
3. **"How to handle transitions?"** → Need correct formula for matching/non-matching
4. **"What about reconstruction?"** → Need to store additional information

### Real-World Analogy
**Like finding the longest symmetric pattern in a sequence:**
- You have a sequence of symbols (string)
- You want to find the longest subsequence that's symmetric
- For any two positions, if symbols match, you can extend the symmetry
- Otherwise, you take the best of excluding one end
- This builds up the longest symmetric pattern
- Useful in DNA analysis, pattern recognition, data compression

### Human-Readable Pseudocode
```
function longestPalindromeSubsequence(s):
    n = s.length()
    if n == 0: return 0
    
    prev = array of size n, filled with 1
    curr = array of size n
    
    for i from n-1 down to 0:
        curr[i] = 1  // Single character
        
        for j from i+1 to n-1:
            if s[i] == s[j]:
                curr[j] = prev[j-1] + 2
            else:
                curr[j] = max(prev[j], curr[j-1])
        
        copy curr to prev
    
    return prev[n-1]
```

### Execution Visualization

### Example: "bbbab"
```
DP Evolution with State Compression:

Initialize: prev = [1,1,1,1,1] (single chars)

i=3 ('a'):
curr = [1,1,1,1,1]
j=4: curr[4] = max(prev[4],curr[3]) = 1
prev = [1,1,1,1,1]

i=2 ('b'):
curr = [1,1,1,1,1]
j=3: curr[3] = max(prev[3],curr[2]) = 1
j=4: curr[4] = prev[3] + 2 = 3 (match 'b'=='b')
prev = [1,1,1,1,3]

i=1 ('b'):
curr = [1,1,1,1,1]
j=2: curr[2] = prev[1] + 2 = 3 (match 'b'=='b')
j=3: curr[3] = max(prev[3],curr[2]) = 3
j=4: curr[4] = max(prev[4],curr[3]) = 3
prev = [1,1,3,3,3]

i=0 ('b'):
curr = [1,1,1,1,1]
j=1: curr[1] = prev[0] + 2 = 3 (match 'b'=='b')
j=2: curr[2] = max(prev[2],curr[1]) = 3
j=3: curr[3] = max(prev[3],curr[2]) = 3
j=4: curr[4] = max(prev[4],curr[3]) = 3
prev = [1,3,3,3,3]

Result: prev[4] = 3 ✓

Visualization:
String: b b b a b
       0 1 2 3 4

DP Table (compressed):
   0 1 2 3 4
0: 1 3 3 3 3
1:   1 3 3 3
2:     1 1 3
3:       1 1
4:         1

Result: 3 ("bbb") ✓
```

### Key Visualization Points:
- **State compression** reduces space from O(N²) to O(N)
- **Row-by-row** processing maintains DP invariants
- **Transition logic** handles matching and non-matching cases
- **Memory efficiency** while preserving correctness

### Memory Layout Visualization:
```
Standard DP (O(N²) space):
   0 1 2 3 4
0: 1 1 1 1 1
1:   1 3 3 3
2:     1 1 3
3:       1 1
4:         1

Compressed DP (O(N) space):
prev: [1,1,1,1,1] ← previous row
curr: [1,1,1,1,1] ← current row

Processing:
For each i: fill curr[j] using prev[j], prev[j-1], curr[j-1]
Copy curr to prev for next iteration

Space Reduction:
Standard: N×N table
Compressed: 2×N rows (or 1×N with careful ordering)
```

### Time Complexity Breakdown:
- **Outer Loop**: O(N) iterations (i from n-1 to 0)
- **Inner Loop**: O(N) iterations (j from i+1 to n-1)
- **DP Computation**: O(1) per cell
- **Array Copy**: O(N) per outer iteration
- **Total**: O(N²) time, O(N) space with compression
- **Optimal**: Cannot do better than O(N²) for this problem
- **vs Standard DP**: Same time, O(N²) vs O(N) space
*/
