package main

import (
	"container/heap"
	"fmt"
	"math"
)

// 1514. Path with Maximum Probability
// Time: O((V + E) log V), Space: O(V + E)
func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Build adjacency list
	adj := make(map[int][]Edge)
	for i, edge := range edges {
		from, to := edge[0], edge[1]
		prob := succProb[i]
		adj[from] = append(adj[from], Edge{to, prob})
		adj[to] = append(adj[to], Edge{from, prob}) // Undirected graph
	}
	
	// Dijkstra's algorithm with max-heap (probability)
	prob := make([]float64, n)
	for i := 0; i < n; i++ {
		prob[i] = 0
	}
	prob[start] = 1
	
	// Max-heap: {probability, node}
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, Item{1.0, start})
	
	for maxHeap.Len() > 0 {
		current := heap.Pop(maxHeap).(Item)
		currentProb, currentNode := current.probability, current.node
		
		// Skip if we've found a better path
		if currentProb < prob[currentNode] {
			continue
		}
		
		// Early exit if we reached the target
		if currentNode == end {
			return currentProb
		}
		
		// Relax edges
		for _, edge := range adj[currentNode] {
			newProb := currentProb * edge.weight
			if newProb > prob[edge.to] {
				prob[edge.to] = newProb
				heap.Push(maxHeap, Item{newProb, edge.to})
			}
		}
	}
	
	return prob[end]
}

// Edge represents a weighted edge (probability)
type Edge struct {
	to     int
	weight float64
}

// Max-heap implementation for probability
type Item struct {
	probability float64
	node        int
}

type MaxHeap []Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].probability > h[j].probability }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Alternative implementation using BFS with priority queue
func maxProbabilityBFS(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Build adjacency list
	adj := make(map[int][]Edge)
	for i, edge := range edges {
		from, to := edge[0], edge[1]
		prob := succProb[i]
		adj[from] = append(adj[from], Edge{to, prob})
		adj[to] = append(adj[to], Edge{from, prob})
	}
	
	// Priority queue (max-heap)
	pq := &MaxHeap{}
	heap.Init(pq)
	heap.Push(pq, Item{1.0, start})
	
	// Track maximum probabilities
	maxProb := make([]float64, n)
	for i := 0; i < n; i++ {
		maxProb[i] = 0
	}
	maxProb[start] = 1
	
	for pq.Len() > 0 {
		current := heap.Pop(pq).(Item)
		prob, node := current.probability, current.node
		
		// Skip if we've found a better path
		if prob < maxProb[node] {
			continue
		}
		
		// Early exit
		if node == end {
			return prob
		}
		
		// Explore neighbors
		for _, edge := range adj[node] {
			newProb := prob * edge.weight
			if newProb > maxProb[edge.to] {
				maxProb[edge.to] = newProb
				heap.Push(pq, Item{newProb, edge.to})
			}
		}
	}
	
	return maxProb[end]
}

// Modified Dijkstra's algorithm with custom comparison
func maxProbabilityModifiedDijkstra(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Build adjacency list
	adj := make(map[int][]Edge)
	for i, edge := range edges {
		from, to := edge[0], edge[1]
		prob := succProb[i]
		adj[from] = append(adj[from], Edge{to, prob})
		adj[to] = append(adj[to], Edge{from, prob})
	}
	
	// Use negative probabilities with min-heap (convert to minimization)
	negProb := make([]float64, n)
	for i := 0; i < n; i++ {
		negProb[i] = math.Inf(1) // Infinity
	}
	negProb[start] = -1 // Negative of 1
	
	// Min-heap: {negative probability, node}
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, MinItem{-1.0, start})
	
	for minHeap.Len() > 0 {
		current := heap.Pop(minHeap).(MinItem)
		currentNegProb, currentNode := current.negProb, current.node
		
		// Skip if we've found a better path
		if currentNegProb > negProb[currentNode] {
			continue
		}
		
		// Early exit
		if currentNode == end {
			return -currentNegProb
		}
		
		// Relax edges
		for _, edge := range adj[currentNode] {
			newNegProb := currentNegProb * edge.weight
			if newNegProb > negProb[edge.to] {
				negProb[edge.to] = newNegProb
				heap.Push(minHeap, MinItem{newNegProb, edge.to})
			}
		}
	}
	
	return -negProb[end]
}

// Min-heap implementation for negative probabilities
type MinItem struct {
	negProb float64
	node    int
}

type MinHeap []MinItem

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].negProb < h[j].negProb }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(MinItem))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// BFS with logarithmic transformation
func maxProbabilityLogTransform(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Build adjacency list
	adj := make(map[int][]Edge)
	for i, edge := range edges {
		from, to := edge[0], edge[1]
		prob := succProb[i]
		adj[from] = append(adj[from], Edge{to, prob})
		adj[to] = append(adj[to], Edge{from, prob})
	}
	
	// Use log probabilities (convert multiplication to addition)
	logProb := make([]float64, n)
	for i := 0; i < n; i++ {
		logProb[i] = math.Inf(-1) // -Infinity
	}
	logProb[start] = 0 // log(1) = 0
	
	// Max-heap: {log probability, node}
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, Item{0, start})
	
	for maxHeap.Len() > 0 {
		current := heap.Pop(maxHeap).(Item)
		currentLogProb, currentNode := current.probability, current.node
		
		// Skip if we've found a better path
		if currentLogProb < logProb[currentNode] {
			continue
		}
		
		// Early exit
		if currentNode == end {
			return math.Exp(currentLogProb)
		}
		
		// Relax edges
		for _, edge := range adj[currentNode] {
			newLogProb := currentLogProb + math.Log(edge.weight)
			if newLogProb > logProb[edge.to] {
				logProb[edge.to] = newLogProb
				heap.Push(maxHeap, Item{newLogProb, edge.to})
			}
		}
	}
	
	return math.Exp(logProb[end])
}

func main() {
	// Test cases
	testCases := []struct {
		n          int
		edges      [][]int
		succProb   []float64
		start      int
		end        int
		description string
	}{
		{
			3,
			[][]int{{0, 1}, {1, 2}, {0, 2}},
			[]float64{0.5, 0.5, 0.2},
			0, 2,
			"Standard case",
		},
		{
			3,
			[][]int{{0, 1}, {1, 2}, {0, 2}},
			[]float64{0.5, 0.5, 0.3},
			0, 2,
			"Higher direct probability",
		},
		{
			3,
			[][]int{{0, 1}},
			[]float64{0.5},
			0, 2,
			"No path",
		},
		{
			5,
			[][]int{{1, 0}, {2, 0}, {3, 2}, {4, 1}},
			[]float64{0.2, 0.1, 0.7, 0.9},
			0, 3,
			"Complex graph",
		},
		{
			4,
			[][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}},
			[]float64{0.5, 0.5, 0.5, 0.1},
			0, 3,
			"Indirect vs direct",
		},
		{
			2,
			[][]int{{0, 1}},
			[]float64{1.0},
			0, 1,
			"Maximum probability",
		},
		{
			4,
			[][]int{{0, 1}, {1, 2}, {2, 3}},
			[]float64{0.1, 0.1, 0.1},
			0, 3,
			"Low probabilities",
		},
		{
			6,
			[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {0, 5}},
			[]float64{0.5, 0.5, 0.5, 0.5, 0.5, 0.1},
			0, 5,
			"Long path vs direct",
		},
		{
			3,
			[][]int{{0, 1}, {0, 2}, {1, 2}},
			[]float64{0.9, 0.1, 0.9},
			0, 2,
			"Multiple high probability paths",
		},
		{
			1,
			[][]int{},
			[]float64{},
			0, 0,
			"Single node",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  n=%d, start=%d, end=%d\n", tc.n, tc.start, tc.end)
		fmt.Printf("  Edges: %v\n", tc.edges)
		fmt.Printf("  Probabilities: %v\n", tc.succProb)
		
		result1 := maxProbability(tc.n, tc.edges, tc.succProb, tc.start, tc.end)
		result2 := maxProbabilityBFS(tc.n, tc.edges, tc.succProb, tc.start, tc.end)
		result3 := maxProbabilityModifiedDijkstra(tc.n, tc.edges, tc.succProb, tc.start, tc.end)
		result4 := maxProbabilityLogTransform(tc.n, tc.edges, tc.succProb, tc.start, tc.end)
		
		fmt.Printf("  Standard Dijkstra: %.6f\n", result1)
		fmt.Printf("  BFS Approach: %.6f\n", result2)
		fmt.Printf("  Modified Dijkstra: %.6f\n", result3)
		fmt.Printf("  Log Transform: %.6f\n\n", result4)
	}
}
