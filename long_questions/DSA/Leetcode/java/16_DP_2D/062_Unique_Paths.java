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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: 2D Dynamic Programming
- **Grid Path Counting**: Count paths in 2D grid with obstacles
- **DP State**: dp[i][j] = number of ways to reach cell (i,j)
- **Recurrence**: dp[i][j] = dp[i-1][j] + dp[i][j-1]
- **Space Optimization**: Use 1D array instead of 2D array

## 2. PROBLEM CHARACTERISTICS
- **Grid Movement**: Can only move right or down
- **Start/End**: From top-left (0,0) to bottom-right (m-1,n-1)
- **No Obstacles**: All cells are traversable
- **Unique Paths**: Count all possible paths

## 3. SIMILAR PROBLEMS
- Unique Paths II (with obstacles)
- Minimum Path Sum
- Grid Traveler
- Dungeon Game

## 4. KEY OBSERVATIONS
- Each cell can be reached from above or left only
- This creates binomial coefficient pattern
- dp[i][j] = C(i+j, i) where C is combination
- Space can be optimized to 1D using rolling array
- Result fits in integer for reasonable grid sizes

## 5. VARIATIONS & EXTENSIONS
- Grid with obstacles
- Different movement constraints
- Modulo arithmetic for large results
- Multiple start/end points

## 6. INTERVIEW INSIGHTS
- Clarify: "Are there obstacles in the grid?"
- Edge cases: empty grid, single cell, large grid
- Time complexity: O(M×N) vs O(2^(M+N)) brute force
- Space complexity: O(N) vs O(M×N) 2D DP

## 7. COMMON MISTAKES
- Using recursion without memoization
- Integer overflow for large grids
- Off-by-one errors in DP indexing
- Not handling edge cases properly

## 8. OPTIMIZATION STRATEGIES
- 1D DP array reduces space from O(M×N) to O(N)
- Mathematical formula using combinations
- Early termination for single row/column
- Modulo arithmetic for large numbers

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like counting routes on a city grid:**
- You have a grid of streets (m rows, n columns)
- You start at top-left corner (origin)
- You can only move right (east) or down (south)
- You want to reach bottom-right corner (destination)
- How many different routes can you take?

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Grid dimensions m and n
2. **Goal**: Number of unique paths from (0,0) to (m-1,n-1)
3. **Output**: Count of possible paths

#### Phase 2: Key Insight Recognition
- **"How to reach a cell?"** → From above OR from left
- **"What's the recurrence?"** → dp[i][j] = dp[i-1][j] + dp[i][j-1]
- **"What are base cases?"** → First row and first column are all 1s
- **"Can I optimize?"** → Use 1D array, only need previous row

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use DP:
1. Create dp table where dp[i][j] = paths to (i,j)
2. Base cases: first row and first column are all 1s
3. For each cell (i,j):
   - dp[i][j] = dp[i-1][j] + dp[i][j-1]
4. This counts all paths ending at (i,j)
5. Optimize space to 1D array"
```

#### Phase 4: Edge Case Handling
- **Empty grid**: Return 0
- **Single cell**: Return 1
- **Single row**: Only one way (all rights)
- **Single column**: Only one way (all downs)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Grid: 3×7 (m=3, n=7)

Human thinking:
"Let's build the DP table:

Base cases:
- First row: [1, 1, 1, 1, 1, 1, 1] (only right moves)
- First column: [1, 1, 1] (only down moves)

Cell (1,1):
- From above: dp[0][1] = 1
- From left: dp[1][0] = 1
- dp[1][1] = 1 + 1 = 2

Cell (1,2):
- From above: dp[0][2] = 1
- From left: dp[1][1] = 2
- dp[1][2] = 1 + 2 = 3

Continue filling...
Final dp[2][6] = number of paths to bottom-right ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Each path is sequence of rights and downs
- **Why it's efficient**: Each cell computed once using previous results
- **Why it's correct**: DP ensures all paths counted exactly once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not try all paths?"** → Exponential O(2^(M+N))
2. **"What about recursion?"** → Stack overflow and recomputation
3. **"How to handle large grids?"** → Use 1D DP and modulo
4. **"What about obstacles?"** → Different problem (Unique Paths II)

### Real-World Analogy
**Like counting delivery routes in a city grid:**
- You have a city divided into a grid of blocks
- You start at northwest corner, want to reach southeast
- You can only travel east or south on each block
- How many different routes can the delivery person take?
- Each intersection (cell) counts routes from northwest
- Total routes equals number of unique paths to destination

### Human-Readable Pseudocode
```
function uniquePaths(m, n):
    if m == 0 or n == 0:
        return 0
    
    // Use 1D DP for space optimization
    dp = [1] * n
    
    // Fill first row
    for j from 0 to n-1:
        dp[j] = 1
    
    // Fill remaining rows
    for i from 1 to m-1:
        for j from 0 to n-1:
            dp[j] = dp[j] + (j == 0 ? dp[j] : dp[j-1])
    
    return dp[n-1]
```

### Execution Visualization

### Example: m = 3, n = 7
```
DP Table Evolution:
   0 1 2 3 4 5 6
0  1 1 1 1 1 1 1
1  1 2 3 4 5 6 7
2  1 3 6 10 15 21 28

Step-by-step:
Row 0: [1, 1, 1, 1, 1, 1, 1, 1]
Row 1: [1, 2, 3, 4, 5, 6, 7]
Row 2: [1, 3, 6, 10, 15, 21, 28]

Final answer: dp[2][6] = 28 paths ✓

Physical meaning:
28 different routes from top-left to bottom-right ✓
```

### Key Visualization Points:
- **DP recurrence** builds from previous row and column
- **Space optimization** uses rolling 1D array
- **Binomial pattern**: C(i+j, i) for each cell
- **Base cases**: First row and column are all 1s

### Memory Layout Visualization:
```
DP Array Evolution (1D optimization):
Initial: [1, 1, 1, 1, 1, 1, 1]

After row 1: [1, 2, 3, 4, 5, 6, 7]
After row 2: [1, 3, 6, 10, 15, 21, 28]

Final: [1, 3, 6, 10, 15, 21, 28]
Answer: 28 (dp[7] or dp[n-1]) ✓
```

### Time Complexity Breakdown:
- **DP Table**: O(M×N) cells
- **Each Cell**: O(1) computation (addition)
- **Total**: O(M×N) time, O(N) space (1D optimization)
- **Optimal**: Cannot do better than O(M×N) for this problem
- **vs Brute Force**: O(2^(M+N)) trying all paths
*/
