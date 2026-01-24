package main

import (
	"container/list"
	"fmt"
)

// Pattern: BFS with State (Shortest Path in Grid with Obstacles Elimination)
// Difficulty: Hard
// Key Concept: Standard BFS tracks `visited[row][col]`. Here we track `visited[row][col][remaining_eliminations]` because arriving at (r,c) with 3 eliminations left is different from arriving with 0 left.

/*
INTUITION:
You are in a maze with walls (1) and empty spaces (0). You want to go from Top-Left to Bottom-Right.
You have a superpower: you can walk through at most K walls.
What is the shortest path?

Standard BFS: Queue stores (r, c).
Problem: If I hit a wall at (2,2) and use my superpower, I am at (2,2) with K-1 charges.
Another path might reach (2,2) without using any charges (K charges). The second path is "better" even if it took same steps, because it has more potential for future walls.
But if the first path is shorter (fewer steps), we want that.

State: `(row, col, remaining_k)`
Level: `steps`

We strictly move level by level. If we reach a state `(r, c, k)` that we have seen before (with same or more k?), we skip.
Actually, if we reach `(r, c)` with LESS k than before, we skip. If we reach with MORE k, it might be useful.
Simplest approach: Treat `(r, c, k)` as a unique node. `visited[r][c][k] = true`.

PROBLEM:
LeetCode 1293. Shortest Path in a Grid with Obstacles Elimination.
Given an m x n integer matrix grid where each cell is either 0 (empty) or 1 (obstacle), return the minimum number of steps to walk from (0, 0) to (m - 1, n - 1) such that you eliminate at most k obstacles.

ALGORITHM:
1. Queue stores struct `{r, c, k, steps}`.
2. `visited` array `[m][n][k+1]` boolean.
3. Push start `{0, 0, k, 0}`.
4. Loop while Queue is not empty:
   - Pop `current`.
   - If `current.r == m-1 && current.c == n-1`, return `current.steps`.
   - For each neighbor `(nr, nc)`:
     - If cell is 0 (Empty):
       - If not visited `[nr][nc][current.k]`:
         - Add `{nr, nc, current.k, steps+1}`.
         - Mark visited.
     - If cell is 1 (Obstacle):
       - If `current.k > 0` and not visited `[nr][nc][current.k-1]`:
         - Add `{nr, nc, current.k-1, steps+1}`,
         - Mark visited.
5. If queue empty, return -1.

Optimization:
Manhattan distance pruning? `steps + dist > best`?
Greedy A*?
For simple BFS, standard queue is fine.
*/

type State struct {
	r, c, k, steps int
}

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])

	// Optimization: If k is large enough to walk Manhattan distance (straight line basically), return Manhattan.
	// If k >= m + n - 3 (approx), we can just go shortest path.
	if k >= m+n-2 {
		return m + n - 2
	}

	queue := list.New()
	queue.PushBack(State{0, 0, k, 0})

	// Visited map: visited[r][c][k]
	// Using map for sparse, or 3D array?
	// k is small (usually <= m*n, but logically limited to m+n).
	// Let's use a flattened map or 3D array if k is small.
	// Leetcode constraints: k <= m*n. If k is huge, we used the optimization above.
	// So k is effectively limited to m+n.
	visited := make([][][]bool, m)
	for i := range visited {
		visited[i] = make([][]bool, n)
		for j := range visited[i] {
			visited[i][j] = make([]bool, k+1)
		}
	}

	visited[0][0][k] = true

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(State)

		if curr.r == m-1 && curr.c == n-1 {
			return curr.steps
		}

		for _, d := range directions {
			nr, nc := curr.r+d[0], curr.c+d[1]

			if nr >= 0 && nr < m && nc >= 0 && nc < n {
				newK := curr.k
				if grid[nr][nc] == 1 {
					newK--
				}

				if newK >= 0 && !visited[nr][nc][newK] {
					visited[nr][nc][newK] = true
					queue.PushBack(State{nr, nc, newK, curr.steps + 1})
				}
			}
		}
	}

	return -1
}

func main() {
	// Grid:
	// 0 0 0
	// 1 1 0
	// 0 0 0
	// 0 1 1
	// 0 0 0
	// k=1
	// Path (0,0)->(0,1)->(0,2)->(1,2)->(2,2)->(3,2)->(4,2). Steps: 6?
	// Obstacles at (1,0) (1,1) (3,1) (3,2).
	// Let's check efficient path.
	// (0,0)->(0,1)->(0,2)->(1,2)...

	grid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 0},
	}
	k := 1
	fmt.Printf("Steps: %d\n", shortestPath(grid, k)) // Expected 6

	// k=0 -> blocked?
	// (0,0)->..->(0,2)->(1,2)(0)->(2,2)(0)->(3,2)(1-wall!) -> Blocked?
	// Actually (3,2) is 1. (3,1) is 1.
	// To get to (4,2), need to pass row 3.
	// (3,0) is 0. (3,1) is 1. (3,2) is 1.
	// Can go (2,0)->(3,0)->(4,0)->(4,1)->(4,2). Steps: 10?
	k2 := 0
	fmt.Printf("Steps (k=0): %d\n", shortestPath(grid, k2))
}
