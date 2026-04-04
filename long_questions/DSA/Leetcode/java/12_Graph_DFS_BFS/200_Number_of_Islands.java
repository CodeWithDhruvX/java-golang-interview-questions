import java.util.Arrays;

public class NumberOfIslands {
    
    // 200. Number of Islands
    // Time: O(M*N), Space: O(M*N) for recursion stack
    public static int numIslands(char[][] grid) {
        if (grid.length == 0 || grid[0].length == 0) {
            return 0;
        }
        
        int m = grid.length, n = grid[0].length;
        int islands = 0;
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '1') {
                    islands++;
                    dfs(grid, i, j, m, n);
                }
            }
        }
        
        return islands;
    }

    private static void dfs(char[][] grid, int i, int j, int m, int n) {
        if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1') {
            return;
        }
        
        // Mark as visited
        grid[i][j] = '0';
        
        // Explore all 4 directions
        dfs(grid, i + 1, j, m, n);
        dfs(grid, i - 1, j, m, n);
        dfs(grid, i, j + 1, m, n);
        dfs(grid, i, j - 1, m, n);
    }

    // Helper function to create grid from string array
    public static char[][] createGrid(String[] str) {
        char[][] grid = new char[str.length][];
        for (int i = 0; i < str.length; i++) {
            grid[i] = str[i].toCharArray();
        }
        return grid;
    }

    public static void main(String[] args) {
        // Test cases
        String[][] testCases = {
            {"11110", "11010", "11000", "00000"},
            {"11000", "11000", "00100", "00011"},
            {"1"},
            {"0"},
            {"111", "111", "111"},
            {"000", "000", "000"},
            {"10101", "01010", "10101"},
            {"10000", "01000", "00100", "00010", "00001"},
            {"110", "011", "001"},
            {}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            char[][] grid = createGrid(testCases[i]);
            int result = numIslands(grid);
            System.out.printf("Test Case %d: %s -> Islands: %d\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}
