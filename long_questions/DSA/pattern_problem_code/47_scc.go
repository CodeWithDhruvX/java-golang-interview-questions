package main

import (
	"fmt"
)

// Pattern: Strongly Connected Components (Tarjan's - Critical Connections)
// Difficulty: Hard
// Key Concept: Using DFS with `discovery_time` and `low_link` values to identify cycles and bridges (edges that, if removed, disconnect the graph).

/*
INTUITION:
In a network of servers, a "Critical Connection" is a cable that, if cut, separates the network into two isolated pieces.
We want to find all such cables.
Using DFS, we can traverse the graph.
- `disc[u]`: Time when we first visited `u`.
- `low[u]`: The lowest `disc` value reachable from `u` (including via back-edges in DFS tree).
If we go from `u` to `v` and find that `low[v] > disc[u]`, it means there is NO back-edge from `v` (or its descendants) to `u` or `u`'s ancestors.
This implies `u-v` is the ONLY path to `v`. It is a BRIDGE.

PROBLEM:
LeetCode 1192. Critical Connections in a Network.
There are `n` servers numbered 0 to `n-1` connected by undirected pairs.
Return all critical connections in the network in any order.

ALGORITHM:
1. Build adjacency graph.
2. Initialize `disc` and `low` arrays with -1. `time` = 0.
3. Call `dfs(u, parent)`:
   - `disc[u] = low[u] = time++`.
   - For neighbor `v` in `adj[u]`:
     - If `v == parent`, continue.
     - If `v` visited (`disc[v] != -1`):
       - `low[u] = min(low[u], disc[v])` (Back-edge).
     - Else:
       - `dfs(v, u)`
       - `low[u] = min(low[u], low[v])` (Tree-edge).
       - If `low[v] > disc[u]`, then `(u, v)` is a Bridge.
*/

var timeCount int
var ans [][]int

func criticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	for _, c := range connections {
		u, v := c[0], c[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	disc := make([]int, n)
	low := make([]int, n)
	for i := range disc {
		disc[i] = -1
		low[i] = -1
	}

	timeCount = 0
	ans = [][]int{}

	// Graph is connected, so one DFS from 0 covers all.
	dfs(0, -1, graph, disc, low)

	return ans
}

func dfs(u, p int, graph [][]int, disc, low []int) {
	disc[u] = timeCount
	low[u] = timeCount
	timeCount++

	for _, v := range graph[u] {
		if v == p {
			continue
		}
		if disc[v] != -1 {
			// Visited, Backedge
			if disc[v] < low[u] {
				low[u] = disc[v]
			}
		} else {
			// Not visited, Tree edge
			dfs(v, u, graph, disc, low)
			if low[v] < low[u] {
				low[u] = low[v]
			}
			// Bridge Condition
			if low[v] > disc[u] {
				ans = append(ans, []int{u, v})
			}
		}
	}
}

func main() {
	// 4 nodes. 0-1, 1-2, 2-0 (Cycle 0-1-2), 1-3 (Edge out).
	// Edge 1-3 is critical. others are in cycle.
	n := 4
	connections := [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}

	fmt.Printf("Critical: %v\n", criticalConnections(n, connections))
}
