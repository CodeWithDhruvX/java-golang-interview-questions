package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents a weighted edge in the graph
type Edge struct {
	to     int
	weight int
}

// 743. Network Delay Time - Dijkstra's Algorithm
// Time: O((V + E) log V), Space: O(V + E)
func networkDelayTime(times [][]int, n int, k int) int {
	// Build adjacency list
	adj := make(map[int][]Edge)
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		adj[from] = append(adj[from], Edge{to, weight})
	}
	
	// Dijkstra's algorithm
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	
	// Min-heap: {distance, node}
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, Item{0, k})
	
	for minHeap.Len() > 0 {
		current := heap.Pop(minHeap).(Item)
		currentDist, currentNode := current.distance, current.node
		
		// Skip if we've found a better path
		if currentDist > dist[currentNode] {
			continue
		}
		
		// Relax edges
		for _, edge := range adj[currentNode] {
			newDist := currentDist + edge.weight
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				heap.Push(minHeap, Item{newDist, edge.to})
			}
		}
	}
	
	// Find the maximum distance
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1 // Unreachable node
		}
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	
	return maxDist
}

// Min-heap implementation
type Item struct {
	distance int
	node     int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Bellman-Ford algorithm (handles negative weights)
func networkDelayTimeBellmanFord(times [][]int, n int, k int) int {
	// Initialize distances
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	
	// Relax edges n-1 times
	for i := 0; i < n-1; i++ {
		updated := false
		for _, time := range times {
			from, to, weight := time[0], time[1], time[2]
			
			if dist[from] != math.MaxInt32 && dist[from]+weight < dist[to] {
				dist[to] = dist[from] + weight
				updated = true
			}
		}
		
		if !updated {
			break // No more updates needed
		}
	}
	
	// Check for unreachable nodes
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

// SPFA (Shortest Path Faster Algorithm) - optimized Bellman-Ford
func networkDelayTimeSPFA(times [][]int, n int, k int) int {
	// Build adjacency list
	adj := make(map[int][]Edge)
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		adj[from] = append(adj[from], Edge{to, weight})
	}
	
	// Initialize distances
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	
	// Queue for SPFA
	queue := []int{k}
	inQueue := make([]bool, n+1)
	inQueue[k] = true
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		inQueue[current] = false
		
		// Relax edges
		for _, edge := range adj[current] {
			newDist := dist[current] + edge.weight
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				
				if !inQueue[edge.to] {
					queue = append(queue, edge.to)
					inQueue[edge.to] = true
				}
			}
		}
	}
	
	// Find maximum distance
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

// Floyd-Warshall algorithm (all pairs shortest paths)
func networkDelayTimeFloydWarshall(times [][]int, n int, k int) int {
	// Initialize distance matrix
	dist := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	
	// Set direct edges
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		dist[from][to] = weight
	}
	
	// Floyd-Warshall
	for mid := 1; mid <= n; mid++ {
		for from := 1; from <= n; from++ {
			for to := 1; to <= n; to++ {
				if dist[from][mid] != math.MaxInt32 && dist[mid][to] != math.MaxInt32 {
					if dist[from][mid]+dist[mid][to] < dist[from][to] {
						dist[from][to] = dist[from][mid] + dist[mid][to]
					}
				}
			}
		}
	}
	
	// Find maximum distance from source k
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[k][i] == math.MaxInt32 {
			return -1
		}
		if dist[k][i] > maxDist {
			maxDist = dist[k][i]
		}
	}
	
	return maxDist
}

// BFS for unweighted graphs
func networkDelayTimeBFS(times [][]int, n int, k int) int {
	// Build adjacency list (unweighted)
	adj := make(map[int][]int)
	for _, time := range times {
		from, to := time[0], time[1]
		adj[from] = append(adj[from], to)
	}
	
	// BFS from source k
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = -1
	}
	dist[k] = 0
	
	queue := []int{k}
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		for _, neighbor := range adj[current] {
			if dist[neighbor] == -1 {
				dist[neighbor] = dist[current] + 1
				queue = append(queue, neighbor)
			}
		}
	}
	
	// Find maximum distance
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] == -1 {
			return -1
		}
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	
	return maxDist
}

func main() {
	// Test cases
	testCases := []struct {
		times      [][]int
		n          int
		k          int
		description string
	}{
		{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, "Standard case"},
		{[][]int{{1, 2, 1}}, 2, 1, "Simple case"},
		{[][]int{{1, 2, 1}}, 2, 2, "Source at end"},
		{[][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}}, 3, 1, "Multiple paths"},
		{[][]int{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}, {4, 5, 1}}, 5, 1, "Linear chain"},
		{[][]int{{1, 2, 1}, {1, 3, 2}, {2, 4, 1}, {3, 4, 1}}, 4, 1, "Converging paths"},
		{[][]int{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}, {4, 1, 1}}, 4, 1, "Cycle"},
		{[][]int{{1, 2, 1}, {2, 3, 10}, {1, 3, 5}}, 3, 1, "Direct vs indirect"},
		{[][]int{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}, {4, 5, 1}, {5, 6, 1}}, 6, 1, "Long chain"},
		{[][]int{{1, 2, 1}, {1, 3, 1}, {2, 4, 1}, {3, 4, 1}}, 4, 1, "Multiple sources"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Times: %v, n=%d, k=%d\n", tc.times, tc.n, tc.k)
		
		result1 := networkDelayTime(tc.times, tc.n, tc.k)
		result2 := networkDelayTimeBellmanFord(tc.times, tc.n, tc.k)
		result3 := networkDelayTimeSPFA(tc.times, tc.n, tc.k)
		result4 := networkDelayTimeFloydWarshall(tc.times, tc.n, tc.k)
		result5 := networkDelayTimeBFS(tc.times, tc.n, tc.k)
		
		fmt.Printf("  Dijkstra: %d\n", result1)
		fmt.Printf("  Bellman-Ford: %d\n", result2)
		fmt.Printf("  SPFA: %d\n", result3)
		fmt.Printf("  Floyd-Warshall: %d\n", result4)
		fmt.Printf("  BFS (unweighted): %d\n\n", result5)
	}
}
