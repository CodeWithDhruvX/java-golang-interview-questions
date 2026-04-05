import java.util.*;

public class LongestCommonSubsequence {
    
    // 1143. Longest Common Subsequence
    // Time: O(M * N), Space: O(N)
    public static int longestCommonSubsequence(String text1, String text2) {
        if (text1 == null || text2 == null) {
            return 0;
        }
        
        int m = text1.length();
        int n = text2.length();
        
        // Use the shorter string for the DP array to save space
        if (n > m) {
            return longestCommonSubsequence(text2, text1);
        }
        
        int[] dp = new int[n + 1];
        
        for (int i = 1; i <= m; i++) {
            int prev = 0; // dp[j-1] from previous row
            for (int j = 1; j <= n; j++) {
                int temp = dp[j]; // Store current dp[j] for next iteration
                if (text1.charAt(i - 1) == text2.charAt(j - 1)) {
                    dp[j] = prev + 1;
                } else {
                    dp[j] = Math.max(dp[j], dp[j - 1]);
                }
                prev = temp;
            }
        }
        
        return dp[n];
    }

    // 72. Edit Distance
    // Time: O(M * N), Space: O(N)
    public static int minDistance(String word1, String word2) {
        if (word1 == null || word2 == null) {
            return 0;
        }
        
        int m = word1.length();
        int n = word2.length();
        
        // Use the shorter string for the DP array to save space
        if (n > m) {
            return minDistance(word2, word1);
        }
        
        int[] dp = new int[n + 1];
        
        // Initialize first row
        for (int j = 0; j <= n; j++) {
            dp[j] = j;
        }
        
        for (int i = 1; i <= m; i++) {
            int prev = dp[0]; // dp[j-1] from previous row
            dp[0] = i; // Initialize first column
            
            for (int j = 1; j <= n; j++) {
                int temp = dp[j]; // Store current dp[j] for next iteration
                
                if (word1.charAt(i - 1) == word2.charAt(j - 1)) {
                    dp[j] = prev;
                } else {
                    dp[j] = Math.min(prev, Math.min(dp[j], dp[j - 1])) + 1;
                }
                
                prev = temp;
            }
        }
        
        return dp[n];
    }

    public static void main(String[] args) {
        // Test cases for longestCommonSubsequence
        String[][] testCases1 = {
            {"abcde", "ace"},
            {"abc", "abc"},
            {"abc", "def"},
            {"", ""},
            {"", "abc"},
            {"abc", ""},
            {"abacbdab", "bdcaba"},
            {"AGGTAB", "GXTXAYB"},
            {"abcde", "fghij"},
            {"abc", "ac"},
            {"aaaa", "aa"},
            {"abc", "abcabc"},
            {"abc", "cba"},
            {"abc", "bac"},
            {"abcdef", "abcf"}
        };
        
        // Test cases for minDistance
        String[][] testCases2 = {
            {"horse", "ros"},
            {"intention", "execution"},
            {"", ""},
            {"", "abc"},
            {"abc", ""},
            {"abc", "abc"},
            {"abc", "def"},
            {"kitten", "sitting"},
            {"flaw", "lawn"},
            {"ab", "ba"},
            {"abcdef", "azced"},
            {"", "a"},
            {"a", ""},
            {"a", "a"},
            {"ab", "abc"}
        };
        
        System.out.println("Longest Common Subsequence:");
        for (int i = 0; i < testCases1.length; i++) {
            String text1 = testCases1[i][0];
            String text2 = testCases1[i][1];
            int result = longestCommonSubsequence(text1, text2);
            System.out.printf("Test Case %d: \"%s\", \"%s\" -> %d\n", 
                i + 1, text1, text2, result);
        }
        
        System.out.println("\nEdit Distance:");
        for (int i = 0; i < testCases2.length; i++) {
            String word1 = testCases2[i][0];
            String word2 = testCases2[i][1];
            int result = minDistance(word1, word2);
            System.out.printf("Test Case %d: \"%s\", \"%s\" -> %d\n", 
                i + 1, word1, word2, result);
        }
    }
}
