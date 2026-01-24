package main

import (
	"fmt"
	"math"
)

// Pattern: Floyd-Warshall Algorithm
// Difficulty: Medium
// Key Concept: Computing All-Pairs Shortest Path in O(V^3) using Dynamic Programming.

/*
INTUITION:
Dijkstra finds shortest path from ONE source to ALL others.
Floyd-Warshall finds shortest path from ALL sources to ALL others.
It works by considering intermediate nodes.
`dist[i][j]` = shortest distance from i to j.
We iteratively improve this by asking: "Is it faster to go from i to j through k?"
`dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])`.
We try this for every possible intermediate node k (from node 0 to N-1).

PROBLEM:
LeetCode 1334. Find the City With the Smallest Number of Neighbors at a Threshold Distance.
There are n cities numbered from 0 to n-1. Given array `edges` where `edges[i] = [from, to, weight]`, and `distanceThreshold`.
Return the city that can reach the fewest number of cities within the threshold. If there are multiple such cities, return the city with the greatest ID.

ALGORITHM:
1. Initialize `dist` matrix (n x n) with Infinity.
2. `dist[i][i] = 0`.
3. For each edge `[u, v, w]`, `dist[u][v] = dist[v][u] = w`.
4. Run Floyd-Warshall:
   - Loop `k` from 0 to n-1.
   - Loop `i` from 0 to n-1.
   - Loop `j` from 0 to n-1.
   - `dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])`.
5. Count reachable cities for each city `i` (where `dist[i][target] <= threshold`).
6. Find min count. Favor larger `i` on tie.
*/

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// Initialize Dist Matrix
	dist := make([][]int, n)
	inf := math.MaxInt32 / 2 // Safe infinity to avoid overflow on addition
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = inf
			}
		}
	}

	// Fill initial edges
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		dist[u][v] = w
		dist[v][u] = w
	}

	// Floyd-Warshall
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	minReachable := n + 1 // Max possible is n
	bestCity := -1

	for i := 0; i < n; i++ {
		reachableCount := 0
		for j := 0; j < n; j++ {
			if i != j && dist[i][j] <= distanceThreshold {
				reachableCount++
			}
		}

		// Problem asks for "city with smallest number".
		// If multiple, "city with greatest ID".
		if reachableCount <= minReachable {
			minReachable = reachableCount
			bestCity = i
		}
	}

	return bestCity
}

func main() {
	// 4 cities. Threshold 4.
	// 0-1 (3), 1-2 (1), 1-3 (4), 2-3 (1).
	// Distances:
	// 0: to 1(3), to 2(3+1=4), to 3(3+4=7 or 3+1+1=5). Reachable: {1, 2}. Count 2.
	// 1: to 0(3), to 2(1), to 3(1+1=2 or 4). Reachable: {0, 2, 3}. Count 3.
	// 2: to 0(4), to 1(1), to 3(1). Reachable: {0, 1, 3}. Count 3.
	// 3: to 0(5), to 1(2), to 2(1). Reachable: {1, 2}. Count 2.
	// Min count is 2 (Cities 0 and 3). Return 3.

	n := 4
	edges := [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
	thresh := 4
	fmt.Printf("Best City: %d\n", findTheCity(n, edges, thresh))
}
