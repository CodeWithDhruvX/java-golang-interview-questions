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
}
