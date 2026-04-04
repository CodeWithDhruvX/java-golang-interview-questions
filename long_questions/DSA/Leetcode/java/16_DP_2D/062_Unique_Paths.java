public class UniquePaths {
    
    // 62. Unique Paths
    // Time: O(M*N), Space: O(M*N)
    public static int uniquePaths(int m, int n) {
        if (m <= 0 || n <= 0) {
            return 0;
        }
        
        // dp[i][j] = number of ways to reach cell (i,j)
        int[][] dp = new int[m][n];
        
        // Initialize first row and first column
        for (int i = 0; i < m; i++) {
            dp[i][0] = 1;
        }
        for (int j = 0; j < n; j++) {
            dp[0][j] = 1;
        }
        
        // Fill the dp table
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
            }
        }
        
        return dp[m - 1][n - 1];
    }

    // Space optimized version: O(N) space
    public static int uniquePathsOptimized(int m, int n) {
        if (m <= 0 || n <= 0) {
            return 0;
        }
        
        // Use 1D array representing current row
        int[] dp = new int[n];
        Arrays.fill(dp, 1);
        
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                dp[j] = dp[j] + dp[j - 1];
            }
        }
        
        return dp[n - 1];
    }

    // Mathematical approach using combinatorics
    public static int uniquePathsMath(int m, int n) {
        // Total steps = (m-1) + (n-1)
        // Choose (m-1) down steps or (n-1) right steps
        // C(total, m-1) or C(total, n-1)
        
        int total = m + n - 2;
        int down = m - 1;
        int right = n - 1;
        
        // Use min(down, right) to reduce computation
        if (down > right) {
            int temp = down;
            down = right;
            right = temp;
        }
        
        long result = 1;
        for (int i = 1; i <= down; i++) {
            result = result * (total - down + i) / i;
        }
        
        return (int) result;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {3, 7},
            {3, 2},
            {1, 1},
            {1, 5},
            {5, 1},
            {2, 2},
            {10, 10},
            {7, 3},
            {3, 3},
            {100, 1},
            {1, 100}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int m = testCases[i][0];
            int n = testCases[i][1];
            
            int result1 = uniquePaths(m, n);
            int result2 = uniquePathsOptimized(m, n);
            int result3 = uniquePathsMath(m, n);
            
            System.out.printf("Test Case %d: m=%d, n=%d -> DP: %d, Optimized: %d, Math: %d\n", 
                i + 1, m, n, result1, result2, result3);
        }
    }
}
