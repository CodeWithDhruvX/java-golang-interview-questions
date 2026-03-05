# Matrix Traversal & BFS with State (Product-Based Companies)

Graph traversal on 2D grids (Matrices) is one of the most frequently tested concepts at FAANG companies. It often requires careful bounds checking, visited set management, and sometimes tracking "state" alongside your BFS queue.

## Question 1: Number of Islands
**Problem Statement:** Given an `m x n` 2D binary grid `grid` which represents a map of `'1'`s (land) and `'0'`s (water), return the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.

### Answer:
This is the quintessential matrix traversal problem. We iterate through every cell in the grid. When we find a `'1'`, we increment our island count, and then launch a DFS (or BFS) to "sink" (turn to `'0'`) the entire connected island. This prevents us from counting the same island multiple times.

**Code Implementation (Java):**
```java
public class NumberOfIslands {
    public int numIslands(char[][] grid) {
        if (grid == null || grid.length == 0) return 0;
        
        int count = 0;
        int rows = grid.length;
        int cols = grid[0].length;
        
        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < cols; j++) {
                if (grid[i][j] == '1') {
                    count++;
                    dfsSink(grid, i, j);
                }
            }
        }
        return count;
    }
    
    private void dfsSink(char[][] grid, int r, int c) {
        // Bounds check and water check
        if (r < 0 || c < 0 || r >= grid.length || c >= grid[0].length || grid[r][c] == '0') {
            return;
        }
        
        // "Sink" the island piece
        grid[r][c] = '0';
        
        // Traverse 4 directions
        dfsSink(grid, r + 1, c); // Down
        dfsSink(grid, r - 1, c); // Up
        dfsSink(grid, r, c + 1); // Right
        dfsSink(grid, r, c - 1); // Left
    }
}
```
**Time Complexity:** O(M * N) where M is rows, N is columns
**Space Complexity:** O(M * N) worst case for recursion stack if the entire grid is one island.

---

## Question 2: Rotting Oranges (Multi-source BFS)
**Problem Statement:** You are given an `m x n` grid where each cell can have one of three values: `0` representing an empty cell, `1` representing a fresh orange, or `2` representing a rotten orange. Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten. Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return `-1`.

### Answer:
Because rotting spreads outward simultaneously from multiple sources level by level, this requires **Multi-source Breadth-First Search (BFS)**. We first scan the grid to add all initially rotten oranges to a Queue and count all fresh oranges. Then we process the Queue minute by minute.

**Code Implementation (Java):**
```java
import java.util.LinkedList;
import java.util.Queue;

public class RottingOranges {
    public int orangesRotting(int[][] grid) {
        if (grid == null || grid.length == 0) return 0;
        
        Queue<int[]> queue = new LinkedList<>();
        int freshCount = 0;
        
        // 1. Initialize Queue with all rotten oranges and count fresh ones
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] == 2) {
                    queue.offer(new int[]{i, j});
                } else if (grid[i][j] == 1) {
                    freshCount++;
                }
            }
        }
        
        if (freshCount == 0) return 0; // Already zero fresh oranges
        
        int minutes = 0;
        int[][] dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        
        // 2. BFS
        while (!queue.isEmpty()) {
            minutes++;
            int size = queue.size();
            
            for (int i = 0; i < size; i++) {
                int[] curr = queue.poll();
                
                for (int[] dir : dirs) { // 4 directions
                    int r = curr[0] + dir[0];
                    int c = curr[1] + dir[1];
                    
                    // If out of bounds or not a fresh orange, skip
                    if (r < 0 || c < 0 || r >= grid.length || c >= grid[0].length || grid[r][c] != 1) {
                        continue;
                    }
                    
                    // Rot the orange, decrement freshcount, push to queue
                    grid[r][c] = 2;
                    freshCount--;
                    queue.offer(new int[]{r, c});
                }
            }
        }
        
        return freshCount == 0 ? minutes - 1 : -1;
    }
}
```
**Time Complexity:** O(M * N)
**Space Complexity:** O(M * N) worst case for the queue.

---

## Question 3: Shortest Path in a Grid with Obstacles Elimination (BFS with State)
**Problem Statement:** You are given an `m x n` integer matrix grid where each cell is either `0` (empty) or `1` (obstacle). You can move up, down, left, or right. You can eliminate at most `k` obstacles. Return the minimum number of steps to walk from the upper left corner `(0, 0)` to the lower right corner `(m - 1, n - 1)`. If no such path exists, return `-1`.

### Answer:
Because we want the *shortest path* combined with variable obstacle breaking, we use BFS. However, a traditional `visited[r][c]` boolean array isn't enough, because returning to the same cell with *more* unused obstacle breaks remaining is a different, valid state. Our visited state must be `visited[r][c] = max_k_remaining`. 

**Code Implementation (Java):**
```java
import java.util.LinkedList;
import java.util.Queue;

public class ShortestPathObstacles {
    public int shortestPath(int[][] grid, int k) {
        int m = grid.length, n = grid[0].length;
        if (m == 1 && n == 1) return 0;
        
        // visited[r][c] stores the maximum 'k' remaining at that cell
        int[][] visited = new int[m][n];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                visited[i][j] = -1;
            }
        }
        
        // Queue stores int[]: {row, col, steps, k_remaining}
        Queue<int[]> queue = new LinkedList<>();
        queue.offer(new int[]{0, 0, 0, k});
        visited[0][0] = k;
        
        int[][] dirs = {{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
        
        while (!queue.isEmpty()) {
            int[] curr = queue.poll();
            int r = curr[0], c = curr[1], steps = curr[2], currK = curr[3];
            
            // Reached destination
            if (r == m - 1 && c == n - 1) return steps;
            
            for (int[] dir : dirs) {
                int nr = r + dir[0];
                int nc = c + dir[1];
                
                if (nr >= 0 && nr < m && nc >= 0 && nc < n) {
                    int nextK = currK - grid[nr][nc]; // Decrease k if obstacle
                    
                    // If we have valid breaks and this path brings us to (nr, nc) 
                    // with more remaining breaks than before
                    if (nextK >= 0 && nextK > visited[nr][nc]) {
                        visited[nr][nc] = nextK;
                        queue.offer(new int[]{nr, nc, steps + 1, nextK});
                    }
                }
            }
        }
        
        return -1;
    }
}
```
**Time Complexity:** O(M * N * K). At worst, we visit each cell `K` times.
**Space Complexity:** O(M * N) for the visited array and queue.
