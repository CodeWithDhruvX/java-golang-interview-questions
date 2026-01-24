package main

import (
	"container/heap"
	"fmt"
)

// Pattern: Minimum Spanning Tree (Prim's Algorithm)
// Difficulty: Medium/Hard
// Key Concept: Connecting all nodes in a graph with minimum total edge weight using a generic greedy strategy (Priority Queue).

/*
INTUITION:
Imagine building a road network connecting sparse cities. You want the cheapest total pavement.
Start at any city (say City 0).
Look at all roads connected to City 0. Pick the cheapest one (say to City 2).
Now you have {0, 2} connected.
Look at all roads from {0, 2} to unvisited cities. Pick the cheapest.
Repeat until all connected.

Algorithms:
- **Kruskal's**: Sort all edges. Add edge if it doesn't form cycle (Union-Find). Better for sparse graphs E ~ V.
- **Prim's**: Grow from a node. PQ stores edges to unvisited nodes. Better for dense graphs E ~ V^2.

PROBLEM:
LeetCode 1584. Min Cost to Connect All Points.
Given an array points representing integer coordinates of some points on a 2D-plane, return the minimum cost to make all points connected. Cost is Manhattan distance.

ALGORITHM (Prim's):
1. Number of points N.
2. `visited` boolean array.
3. Priority Queue stores `(cost, node)`.
4. Start with node 0. Push `(0, 0)`.
5. Loop while `visited` count < N:
   - Pop `(cost, u)`.
   - If `visited[u]`, continue.
   - Mark `visited[u]`. Add `cost` to total.
   - For every node `v` from 0 to N-1:
     - If not `visited[v]`:
       - `dist = abs(xi-xj) + abs(yi-yj)`.
       - Push `(dist, v)`.
   - Optimization: For Dense Graph (Complete Graph like this, E=N^2), simple array scan for min-dist is O(N^2), while PQ is O(N^2 log N).
   - Simple Prim's (Array version):
     - `minDist` array init to Infinity.
     - Loop N times:
       - Find `u` with Smallest `minDist[u]` not visited.
       - Mark `visited`.
       - Update `minDist[v]` for all neighbors.
     - Time O(N^2). Since N=1000, N^2 = 10^6 (Fast). PQ = 10^6 * 10 (Slower but OK).
     - We will implement PQ version as it's the standard generic Prim's logic. (Note: For strictly dense graphs, Array Prim is better).
*/

// Abs utility
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type EdgeMST struct {
	to, weight int
}

type ItemMST struct {
	to, weight int // Use 'to' as node index
	index      int
}

type PriorityQueueMST []*ItemMST

func (pq PriorityQueueMST) Len() int { return len(pq) }
func (pq PriorityQueueMST) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}
func (pq PriorityQueueMST) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueueMST) Push(x interface{}) {
	n := len(*pq)
	item := x.(*ItemMST)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueueMST) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	visited := make([]bool, n)
	pq := make(PriorityQueueMST, 0)
	heap.Init(&pq)

	// Start from node 0, cost 0
	heap.Push(&pq, &ItemMST{to: 0, weight: 0})

	totalCost := 0
	edgesCount := 0

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*ItemMST)
		u := curr.to
		cost := curr.weight

		if visited[u] {
			continue
		}

		visited[u] = true
		totalCost += cost
		edgesCount++

		if edgesCount == n {
			break
		}

		// Push all edges from u to unvisited v
		// Implicit graph: All connected to All
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := abs(points[u][0]-points[v][0]) + abs(points[u][1]-points[v][1])
				heap.Push(&pq, &ItemMST{to: v, weight: dist})
			}
		}
	}

	return totalCost
}

func main() {
	// Points: (0,0), (2,2), (3,10), (5,2), (7,0)
	// 0-0 connect cost 0.
	// 0(0,0) -> 1(2,2) cost 4.
	// 0(0,0) -> 3(5,2) cost 7.
	// 0(0,0) -> 4(7,0) cost 7.
	// ...
	// Min Cost: 20
	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	fmt.Printf("Min Cost: %d\n", minCostConnectPoints(points))
}
