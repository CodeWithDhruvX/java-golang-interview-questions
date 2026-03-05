# Advanced Dynamic Programming (Product-Based Companies)

While standard 1D and 2D DP problems (like Coin Change or Longest Common Subsequence) are common, L5+ FAANG interviews often test **DP on Trees**, **DP on Matrix/Graphs**, and **Partition DP** (where you divide a problem into smaller overlapping intervals).

## Question 1: Binary Tree Maximum Path Sum (DP on Trees)
**Problem Statement:** A path in a binary tree is a sequence of nodes where each pair of adjacent nodes has an edge connecting them. Given the `root` of a binary tree, return the maximum path sum of any non-empty path.

### Answer:
This is a classic Post-Order Traversal (DFS) problem that acts like DP. For any given node, a max path could go straight down through it, or it could "arch" over the node (including both its left and right max paths). However, a node can only return *one* straight path (either left or right) to its parent.

**Code Implementation (Java):**
```java
class TreeNodeDP {
    int val;
    TreeNodeDP left;
    TreeNodeDP right;
    TreeNodeDP(int x) { val = x; }
}

public class MaxPathSumBinaryTree {
    private int globalMax;

    public int maxPathSum(TreeNodeDP root) {
        globalMax = Integer.MIN_VALUE;
        findMaxBranch(root);
        return globalMax;
    }
    
    // Returns the max path sum of a single branch stretching downwards from the current node
    private int findMaxBranch(TreeNodeDP node) {
        if (node == null) return 0;
        
        // Compute max paths from the children. If negative, ignore them (0)
        int leftMax = Math.max(0, findMaxBranch(node.left));
        int rightMax = Math.max(0, findMaxBranch(node.right));
        
        // Case 1: The path arches over the current node (includes both children)
        // We update the global maximum if this arch path is the highest seen so far
        int archSum = node.val + leftMax + rightMax;
        globalMax = Math.max(globalMax, archSum);
        
        // Case 2: Return the max single branch path to the parent
        return node.val + Math.max(leftMax, rightMax);
    }
}
```
**Time Complexity:** O(N) where N is the number of nodes.
**Space Complexity:** O(H) where H is the height of the tree (call stack).

---

## Question 2: House Robber III (DP on Trees with Status)
**Problem Statement:** The thief has found himself a new place for his thievery again. There is only one entrance to this area, called `root`. Except for the `root`, each house has one and only one parent house. After a tour, the smart thief realized that all houses in this place form a binary tree. It will automatically contact the police if two directly-linked houses were broken into on the same night. Given the `root` of the binary tree, return the maximum amount of money the thief can rob without alerting the police.

### Answer:
For every node, we have a state: we either **rob** it or we **skip** it.
- If we rob a node, we cannot rob its children.
- If we skip a node, we are free to rob or skip its children (whichever yields more money).
Our DFS function will return an array of size 2: `[max_money_if_skipped, max_money_if_robbed]`.

**Code Implementation (Java):**
```java
public class HouseRobberIII {
    public int rob(TreeNodeDP root) {
        int[] result = robStatus(root);
        return Math.max(result[0], result[1]);
    }
    
    // Returns an array: index 0 = max if skip node, index 1 = max if rob node
    private int[] robStatus(TreeNodeDP node) {
        if (node == null) {
            return new int[]{0, 0};
        }
        
        int[] leftStatus = robStatus(node.left);
        int[] rightStatus = robStatus(node.right);
        
        int[] currStatus = new int[2];
        
        // If we skip current node, we can either rob or skip the children (whichever is bigger)
        currStatus[0] = Math.max(leftStatus[0], leftStatus[1]) + 
                        Math.max(rightStatus[0], rightStatus[1]);
                        
        // If we rob current node, we must definitively skip the children
        currStatus[1] = node.val + leftStatus[0] + rightStatus[0];
        
        return currStatus;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(H)

---

## Question 3: Longest Increasing Path in a Matrix (DP on Matrix / Memoization)
**Problem Statement:** Given an `m x n` integers `matrix`, return the length of the longest increasing path in `matrix`. From each cell, you can either move in four directions.

### Answer:
A pure DFS will recalculate the longest path from a cell many times. We use a **Memoization Table** (`dp[i][j]`) to cache the longest path starting from cell `(i, j)`. This converts an exponential time brute force into a linear scan of the grid.

**Code Implementation (Java):**
```java
public class LongestIncreasingPath {
    private int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
    private int m, n;
    
    public int longestIncreasingPath(int[][] matrix) {
        if (matrix == null || matrix.length == 0) return 0;
        
        m = matrix.length;
        n = matrix[0].length;
        int[][] cache = new int[m][n];
        int maxPath = 1;
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                maxPath = Math.max(maxPath, dfsMemo(matrix, i, j, cache));
            }
        }
        return maxPath;
    }
    
    private int dfsMemo(int[][] matrix, int i, int j, int[][] cache) {
        if (cache[i][j] != 0) return cache[i][j]; // Return computed result
        
        int longest = 1; // Minimum path is the cell itself
        
        for (int[] dir : dirs) {
            int row = i + dir[0];
            int col = j + dir[1];
            
            if (row >= 0 && row < m && col >= 0 && col < n && matrix[row][col] > matrix[i][j]) {
                longest = Math.max(longest, 1 + dfsMemo(matrix, row, col, cache));
            }
        }
        
        cache[i][j] = longest;
        return longest;
    }
}
```
**Time Complexity:** O(M * N). Every cell is evaluated and cached exactly once.
**Space Complexity:** O(M * N) for the recursion stack and the cache matrix.

---

## Question 4: Burst Balloons (Hard Partition DP)
**Problem Statement:** You are given `n` balloons, indexed from `0` to `n - 1`. Each balloon is painted with a number on it represented by an array `nums`. You are asked to burst all the balloons. If you burst the `i-th` balloon, you will get `nums[i - 1] * nums[i] * nums[i + 1]` coins. If `i - 1` or `i + 1` goes out of bounds, treat it as if there is a balloon with a `1` painted on it. Return the maximum coins you can collect by bursting the balloons wisely.

### Answer:
This is a classic "Matrix Chain Multiplication" style DP problem. It's very difficult to think "Top Down" because bursting a balloon changes the neighbors of the remaining balloons. The insight requires thinking in reverse: **Which balloon should be burst LAST within the window `[left, right]`?**

**Code Implementation (Java):**
```java
public class BurstBalloons {
    public int maxCoins(int[] nums) {
        // Create a new array wrapping original nums with 1s at ends
        int n = nums.length;
        int[] paddedNums = new int[n + 2];
        paddedNums[0] = 1;
        paddedNums[n + 1] = 1;
        for (int i = 0; i < n; i++) {
            paddedNums[i + 1] = nums[i];
        }
        
        // Cache for memoization
        int[][] memo = new int[n + 2][n + 2];
        
        // Calculate max coins we can get between indices 1 and n
        return dp(paddedNums, memo, 1, n);
    }
    
    // dp(left, right) returns max coins obtained by bursting balloons strictly INclusive [left, right]
    private int dp(int[] nums, int[][] memo, int left, int right) {
        if (left > right) return 0;
        if (memo[left][right] != 0) return memo[left][right];
        
        int max = 0;
        // Assume 'i' is the VERY LAST balloon we burst in the range [left, right].
        for (int i = left; i <= right; i++) {
            // Coins generated by bursting 'i' last = 
            // the coins from bursting the left subproblem 
            // + coins from right subproblem 
            // + coins from bursting 'i' itself (using the outer boundaries left-1 and right+1 which are still unburst)
            int coins = nums[left - 1] * nums[i] * nums[right + 1];
            coins += dp(nums, memo, left, i - 1) + dp(nums, memo, i + 1, right);
            
            max = Math.max(max, coins);
        }
        
        memo[left][right] = max;
        return max;
    }
}
```
**Time Complexity:** O(N^3). N^2 states in the memo table, and an O(N) loop to try every 'last' burst within the state.
**Space Complexity:** O(N^2) for the memoization table.
