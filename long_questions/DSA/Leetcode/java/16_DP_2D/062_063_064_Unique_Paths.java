import java.util.*;

public class UniquePaths {
    
    // 62. Unique Paths
    // Time: O(M * N), Space: O(N)
    public static int uniquePaths(int m, int n) {
        if (m <= 0 || n <= 0) {
            return 0;
        }
        
        int[] dp = new int[n];
        Arrays.fill(dp, 1);
        
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                dp[j] += dp[j - 1];
            }
        }
        
        return dp[n - 1];
    }

    // 63. Unique Paths II
    // Time: O(M * N), Space: O(N)
    public static int uniquePathsWithObstacles(int[][] obstacleGrid) {
        if (obstacleGrid == null || obstacleGrid.length == 0 || obstacleGrid[0].length == 0) {
            return 0;
        }
        
        int m = obstacleGrid.length;
        int n = obstacleGrid[0].length;
        int[] dp = new int[n];
        
        // Initialize first row
        dp[0] = obstacleGrid[0][0] == 0 ? 1 : 0;
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (obstacleGrid[i][j] == 1) {
                    dp[j] = 0;
                } else if (j > 0) {
                    dp[j] += dp[j - 1];
                }
            }
        }
        
        return dp[n - 1];
    }

    // 64. Minimum Path Sum
    // Time: O(M * N), Space: O(N)
    public static int minPathSum(int[][] grid) {
        if (grid == null || grid.length == 0 || grid[0].length == 0) {
            return 0;
        }
        
        int m = grid.length;
        int n = grid[0].length;
        int[] dp = new int[n];
        
        // Initialize first row
        dp[0] = grid[0][0];
        for (int j = 1; j < n; j++) {
            dp[j] = dp[j - 1] + grid[0][j];
        }
        
        for (int i = 1; i < m; i++) {
            dp[0] += grid[i][0];
            for (int j = 1; j < n; j++) {
                dp[j] = Math.min(dp[j - 1], dp[j]) + grid[i][j];
            }
        }
        
        return dp[n - 1];
    }

    public static void main(String[] args) {
        // Test cases for uniquePaths
        Object[][] testCases1 = {
            {3, 7},
            {3, 2},
            {1, 1},
            {1, 5},
            {5, 1},
            {2, 2},
            {4, 4},
            {10, 10},
            {7, 3},
            {5, 5}
        };
        
        // Test cases for uniquePathsWithObstacles
        int[][][] testCases2 = {
            {{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
            {{0, 1}, {0, 0}},
            {{1, 0}, {0, 0}},
            {{0, 0}, {0, 0}},
            {{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
            {{0}},
            {{1}},
            {{0, 1, 0}, {0, 0, 0}, {1, 0, 0}},
            {{0, 0, 0}, {1, 1, 0}, {0, 0, 0}},
            {{0, 0}, {0, 0}, {0, 0}}
        };
        
        // Test cases for minPathSum
        int[][][] testCases3 = {
            {{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
            {{1, 2, 3}, {4, 5, 6}},
            {{1, 2}, {1, 1}},
            {{1}},
            {{1, 1}, {1, 1}},
            {{5, 9, 2}, {4, 6, 5}, {10, 3, 1}},
            {{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
            {{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
            {{2, 3, 1}, {1, 4, 8}, {5, 6, 7}},
            {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
        };
        
        System.out.println("Unique Paths:");
        for (int i = 0; i < testCases1.length; i++) {
            int m = (int) testCases1[i][0];
            int n = (int) testCases1[i][1];
            int result = uniquePaths(m, n);
            System.out.printf("Test Case %d: m=%d, n=%d -> %d\n", 
                i + 1, m, n, result);
        }
        
        System.out.println("\nUnique Paths II (with obstacles):");
        for (int i = 0; i < testCases2.length; i++) {
            int[][] grid = testCases2[i];
            int result = uniquePathsWithObstacles(grid);
            System.out.printf("Test Case %d: %s -> %d\n", 
                i + 1, Arrays.deepToString(grid), result);
        }
        
        System.out.println("\nMinimum Path Sum:");
        for (int i = 0; i < testCases3.length; i++) {
            int[][] grid = testCases3[i];
            int result = minPathSum(grid);
            System.out.printf("Test Case %d: %s -> %d\n", 
                i + 1, Arrays.deepToString(grid), result);
        }
    }
}
