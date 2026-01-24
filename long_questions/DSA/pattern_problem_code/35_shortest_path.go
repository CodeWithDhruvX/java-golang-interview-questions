package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Pattern: Shortest Path (Dijkstra's Algorithm)
// Difficulty: Medium/Hard
// Key Concept: Finding the shortest path in a weighted graph with non-negative weights using a priority queue.

/*
INTUITION:
BFS explores layer by layer (Edge weight = 1).
Dijkstra is "BFS with a brain". It explores the promising (cheapest) paths first.
We maintain a "known shortest distance" to every node. Initially Infinity.
We put (0, StartNode) in a Min-PriorityQueue.
While Queue is not empty:
- Pop the node `u` with the smallest distance.
- For every neighbor `v` of `u`:
  - `newDist = dist[u] + weight(u, v)`
  - If `newDist < dist[v]`:
    - `dist[v] = newDist`
    - Push `v` to Queue.

PROBLEM:
LeetCode 743. Network Delay Time.
You are given a network of `n` nodes, labeled from 1 to `n`. You are also given `times`, a list of travel times as directed edges `times[i] = (u, v, w)`, where `u` is the source node, `v` is the target node, and `w` is the time it takes for a signal to travel from source to target.
We send a signal from a given node `k`. Return the minimum time it takes for all the `n` nodes to receive the signal. If it is impossible for all the `n` nodes to receive the signal, return -1.

ALGORITHM:
1. Build Adjacency List: `graph[u] -> [(v, w), ...]`.
2. Initialize `dist` map/array with Infinity. `dist[k] = 0`.
3. PriorityQueue `pq` holds `Item{node, distance}`. Push `Item{k, 0}`.
4. While `pq` not empty:
   - Pop `curr`.
   - If `curr.dist > dist[curr.node]`, continue (stale entry).
   - For neighbor `next, weight` in `graph[curr.node]`:
     - If `dist[curr.node] + weight < dist[next]`:
       - `dist[next] = dist[curr.node] + weight`
       - Push `Item{next, dist[next]}`
5. Result is `max(dist)`. If any node is Infinity, return -1.
*/

type Edge struct {
	to, weight int
}

// Priority Queue Implementation
type Item struct {
	node, dist int
	index      int // for heap update if needed (not needed for simple Dijkstra)
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist // Min Heap based on distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func networkDelayTime(times [][]int, n int, k int) int {
	// Build Graph
	graph := make(map[int][]Edge)
	for _, t := range times {
		u, v, w := t[0], t[1], t[2]
		graph[u] = append(graph[u], Edge{to: v, weight: w})
	}

	// Dist array
	dist := make(map[int]int)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0

	// PQ
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: k, dist: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Item)
		u := curr.node
		d := curr.dist

		if d > dist[u] {
			continue
		}

		for _, e := range graph[u] {
			v := e.to
			weight := e.weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(&pq, &Item{node: v, dist: dist[v]})
			}
		}
	}

	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1
		}
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}

	return maxDist
}

func main() {
	// 2 -> (1, 1ms)
	// 2 -> (3, 1ms)
	// 3 -> (4, 1ms)
	// Send from 2.
	// 2->1 (1ms). 2->3 (1ms) -> 3->4 (1+1=2ms).
	// Max: 2ms.
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n := 4
	k := 2
	fmt.Printf("Max Delay: %d\n", networkDelayTime(times, n, k))

	// Disconnected
	times2 := [][]int{{1, 2, 1}}
	n2 := 2
	k2 := 2
	fmt.Printf("Max Delay: %d\n", networkDelayTime(times2, n2, k2)) // Expect -1
}
