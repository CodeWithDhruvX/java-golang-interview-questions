import java.util.HashMap;
import java.util.Map;

public class ClimbingStairs {
    
    // 70. Climbing Stairs
    // Time: O(N), Space: O(1)
    public static int climbStairs(int n) {
        if (n <= 2) {
            return n;
        }
        
        int prev2 = 1, prev1 = 2;
        
        for (int i = 3; i <= n; i++) {
            int current = prev1 + prev2;
            prev2 = prev1;
            prev1 = current;
        }
        
        return prev1;
    }

    // DP array approach
    public static int climbStairsDP(int n) {
        if (n <= 2) {
            return n;
        }
        
        int[] dp = new int[n + 1];
        dp[1] = 1;
        dp[2] = 2;
        
        for (int i = 3; i <= n; i++) {
            dp[i] = dp[i - 1] + dp[i - 2];
        }
        
        return dp[n];
    }

    // Recursive with memoization
    public static int climbStairsMemo(int n) {
        Map<Integer, Integer> memo = new HashMap<>();
        return climbStairsHelper(n, memo);
    }

    private static int climbStairsHelper(int n, Map<Integer, Integer> memo) {
        if (n <= 2) {
            return n;
        }
        
        if (memo.containsKey(n)) {
            return memo.get(n);
        }
        
        int result = climbStairsHelper(n - 1, memo) + climbStairsHelper(n - 2, memo);
        memo.put(n, result);
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[] testCases = {1, 2, 3, 4, 5, 10, 20, 30, 45, 50};
        
        for (int i = 0; i < testCases.length; i++) {
            int n = testCases[i];
            int result1 = climbStairs(n);
            int result2 = climbStairsDP(n);
            int result3 = climbStairsMemo(n);
            
            System.out.printf("Test Case %d: n=%d -> Iterative: %d, DP: %d, Memo: %d\n", 
                i + 1, n, result1, result2, result3);
        }
    }
}
