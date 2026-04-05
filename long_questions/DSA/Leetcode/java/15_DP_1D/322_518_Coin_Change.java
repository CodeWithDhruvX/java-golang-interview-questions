import java.util.*;

public class CoinChange {
    
    // 322. Coin Change
    // Time: O(N * S), Space: O(S) where S is the target amount
    public static int coinChange(int[] coins, int amount) {
        if (coins == null || coins.length == 0 || amount < 0) {
            return -1;
        }
        
        int[] dp = new int[amount + 1];
        Arrays.fill(dp, amount + 1); // Initialize with a value larger than any possible solution
        dp[0] = 0;
        
        for (int i = 1; i <= amount; i++) {
            for (int coin : coins) {
                if (coin <= i) {
                    dp[i] = Math.min(dp[i], dp[i - coin] + 1);
                }
            }
        }
        
        return dp[amount] > amount ? -1 : dp[amount];
    }

    // 518. Coin Change II
    // Time: O(N * S), Space: O(S) where S is the target amount
    public static int change(int amount, int[] coins) {
        if (coins == null || coins.length == 0 || amount < 0) {
            return 0;
        }
        
        int[] dp = new int[amount + 1];
        dp[0] = 1;
        
        for (int coin : coins) {
            for (int i = coin; i <= amount; i++) {
                dp[i] += dp[i - coin];
            }
        }
        
        return dp[amount];
    }

    public static void main(String[] args) {
        // Test cases for coinChange
        Object[][] testCases1 = {
            {new int[]{1, 2, 5}, 11},
            {new int[]{2}, 3},
            {new int[]{1}, 0},
            {new int[]{1}, 1},
            {new int[]{1, 3, 4, 5}, 7},
            {new int[]{2, 5, 10, 1}, 27},
            {new int[]{186, 419, 83, 408}, 6249},
            {new int[]{1, 2, 5, 10}, 12},
            {new int[]{2, 5, 10}, 3},
            {new int[]{1, 2, 5}, 100}
        };
        
        // Test cases for change
        Object[][] testCases2 = {
            {5, new int[]{1, 2, 5}},
            {3, new int[]{2}},
            {10, new int[]{2, 5, 3, 6}},
            {0, new int[]{1, 2, 5}},
            {1, new int[]{1}},
            {100, new int[]{1, 2, 5, 10}},
            {7, new int[]{1, 3, 4, 5}},
            {27, new int[]{2, 5, 10, 1}},
            {500, new int[]{1, 2, 5, 10, 25, 50}},
            {8, new int[]{2, 3, 5}}
        };
        
        System.out.println("Coin Change I:");
        for (int i = 0; i < testCases1.length; i++) {
            int[] coins = (int[]) testCases1[i][0];
            int amount = (int) testCases1[i][1];
            int result = coinChange(coins, amount);
            System.out.printf("Test Case %d: %s, amount=%d -> %d\n", 
                i + 1, Arrays.toString(coins), amount, result);
        }
        
        System.out.println("\nCoin Change II:");
        for (int i = 0; i < testCases2.length; i++) {
            int amount = (int) testCases2[i][0];
            int[] coins = (int[]) testCases2[i][1];
            int result = change(amount, coins);
            System.out.printf("Test Case %d: amount=%d, %s -> %d\n", 
                i + 1, amount, Arrays.toString(coins), result);
        }
    }
}
