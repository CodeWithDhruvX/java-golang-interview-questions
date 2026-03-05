# Dynamic Programming and Backtracking (Product-Based Companies)

Dynamic Programming (DP) and Backtracking are essential for solving complex optimization and enumeration problems. Interviews at Amazon, Google, and Microsoft feature these heavily.

## Question 1: Longest Palindromic Substring
**Problem Statement:** Given a string `s`, return the longest palindromic substring in `s`.

### Answer:
The most intuitive and optimally space-efficient approach is "Expand Around Center". A palindrome mirrors around its center. Therefore, a palindrome can be expanded from its center, and there are only `2n - 1` such centers (since even-length palindromes have their center between two letters).

**Code Implementation (Java):**
```java
public class LongestPalindromicSubstring {
    public String longestPalindrome(String s) {
        if (s == null || s.length() < 1) return "";
        int start = 0, end = 0;
        
        for (int i = 0; i < s.length(); i++) {
            int len1 = expandAroundCenter(s, i, i);       // Odd length
            int len2 = expandAroundCenter(s, i, i + 1);   // Even length
            int len = Math.max(len1, len2);
            
            if (len > end - start) {
                start = i - (len - 1) / 2;
                end = i + len / 2;
            }
        }
        return s.substring(start, end + 1);
    }
    
    private int expandAroundCenter(String s, int left, int right) {
        while (left >= 0 && right < s.length() && s.charAt(left) == s.charAt(right)) {
            left--;
            right++;
        }
        return right - left - 1;
    }
}
```
**Time Complexity:** O(N^2)
**Space Complexity:** O(1)

---

## Question 2: Coin Change
**Problem Statement:** You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money. Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return `-1`.

### Answer:
This is a classic DP problem (Unbounded Knapsack). We use a 1D DP array where `dp[i]` stores the minimum number of coins needed for amount `i`. We iterate through all amounts from 1 to `amount` and for each coin check if picking that coin yields a smaller total coin count.

**Code Implementation (Java):**
```java
import java.util.Arrays;

public class CoinChange {
    public int coinChange(int[] coins, int amount) {
        int[] dp = new int[amount + 1];
        Arrays.fill(dp, amount + 1); // Fill with max possible value
        dp[0] = 0; // 0 coins needed for amount 0
        
        for (int i = 1; i <= amount; i++) {
            for (int coin : coins) {
                if (coin <= i) {
                    dp[i] = Math.min(dp[i], dp[i - coin] + 1);
                }
            }
        }
        
        return dp[amount] > amount ? -1 : dp[amount];
    }
}
```
**Time Complexity:** O(amount * len(coins))
**Space Complexity:** O(amount)

---

## Question 3: Word Break
**Problem Statement:** Given a string `s` and a dictionary of strings `wordDict`, return `true` if `s` can be segmented into a space-separated sequence of one or more dictionary words.

### Answer:
We use a boolean DP array `dp` where `dp[i]` represents if the prefix of `s` of length `i` can be broken into dictionary words. For every length `i`, we check if any partition `j` allows `dp[j]` to be true AND the remaining substring `s.substring(j, i)` is in the dictionary.

**Code Implementation (Java):**
```java
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class WordBreak {
    public boolean wordBreak(String s, List<String> wordDict) {
        Set<String> wordSet = new HashSet<>(wordDict);
        boolean[] dp = new boolean[s.length() + 1];
        dp[0] = true; // Empty string is always a valid break
        
        for (int i = 1; i <= s.length(); i++) {
            for (int j = 0; j < i; j++) {
                if (dp[j] && wordSet.contains(s.substring(j, i))) {
                    dp[i] = true;
                    break; // Move to next i since we only need to know if it's breakable
                }
            }
        }
        
        return dp[s.length()];
    }
}
```
**Time Complexity:** O(N^3) (due to substring matching up to length N, can be optimized with Tries)
**Space Complexity:** O(N)

---

## Question 4: N-Queens
**Problem Statement:** The n-queens puzzle is the problem of placing `n` queens on an `n x n` chessboard such that no two queens attack each other. Given an integer `n`, return all distinct solutions to the n-queens puzzle.

### Answer:
This requires a Backtracking approach. We place a queen row by row. At each row, we try placing a queen in each column, checking if it is safe (no conflicts in column, positive diagonal, or negative diagonal). If it's safe, we recursively place a queen in the next row. If we reach row `n`, we found a solution.

**Code Implementation (Java):**
```java
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class NQueens {
    public List<List<String>> solveNQueens(int n) {
        List<List<String>> res = new ArrayList<>();
        char[][] board = new char[n][n];
        for (int i = 0; i < n; i++) {
            Arrays.fill(board[i], '.');
        }
        backtrack(res, board, 0, n);
        return res;
    }
    
    private void backtrack(List<List<String>> res, char[][] board, int row, int n) {
        if (row == n) {
            List<String> validBoard = new ArrayList<>();
            for (int i = 0; i < n; i++) {
                validBoard.add(new String(board[i]));
            }
            res.add(validBoard);
            return;
        }
        
        for (int col = 0; col < n; col++) {
            if (isValid(board, row, col, n)) {
                board[row][col] = 'Q';
                backtrack(res, board, row + 1, n);
                board[row][col] = '.'; // Backtrack
            }
        }
    }
    
    private boolean isValid(char[][] board, int row, int col, int n) {
        // Check column
        for (int i = 0; i < row; i++) {
            if (board[i][col] == 'Q') return false;
        }
        // Check upper-left diagonal
        for (int i = row - 1, j = col - 1; i >= 0 && j >= 0; i--, j--) {
            if (board[i][j] == 'Q') return false;
        }
        // Check upper-right diagonal
        for (int i = row - 1, j = col + 1; i >= 0 && j < n; i--, j++) {
            if (board[i][j] == 'Q') return false;
        }
        return true;
    }
}
```
**Time Complexity:** O(N!)
**Space Complexity:** O(N^2)
