package main

import (
	"fmt"
	"math"
)

// 743. Network Delay Time - DP on Graphs
// Time: O(N^3) for Floyd-Warshall, O(N^2) for optimized approaches
func networkDelayTimeDP(times [][]int, n int, k int) int {
	if len(times) == 0 {
		return -1
	}
	
	// Build adjacency matrix
	dist := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	
	// Fill direct edges
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		dist[from][to] = weight
	}
	
	// Floyd-Warshall algorithm
	for intermediate := 1; intermediate <= n; intermediate++ {
		for from := 1; from <= n; from++ {
			for to := 1; to <= n; to++ {
				if dist[from][intermediate] != math.MaxInt32 && 
				   dist[intermediate][to] != math.MaxInt32 {
					newDist := dist[from][intermediate] + dist[intermediate][to]
					if newDist < dist[from][to] {
						dist[from][to] = newDist
					}
				}
			}
		}
	}
	
	// Find maximum distance from k
	maxDist := 0
	for i := 1; i <= n; i++ {
		if i != k && dist[k][i] > maxDist {
			maxDist = dist[k][i]
		}
	}
	
	if maxDist == math.MaxInt32 {
		return -1
	}
	
	return maxDist
}

// DP with Dijkstra optimization
func networkDelayTimeDijkstraDP(times [][]int, n int, k int) int {
	if len(times) == 0 {
		return -1
	}
	
	// Build adjacency list
	adj := make([][]Edge, n+1)
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		adj[from] = append(adj[from], Edge{to, weight})
	}
	
	// Dijkstra from source k
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	
	visited := make([]bool, n+1)
	
	for i := 1; i <= n; i++ {
		// Find unvisited node with minimum distance
		u := -1
		minDist := math.MaxInt32
		
		for v := 1; v <= n; v++ {
			if !visited[v] && dist[v] < minDist {
				minDist = dist[v]
				u = v
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		
		// Update distances
		for _, edge := range adj[u] {
			to, weight := edge.to, edge.weight
			if dist[u]+weight < dist[to] {
				dist[to] = dist[u] + weight
			}
		}
	}
	
	// Find maximum distance
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	
	if maxDist == math.MaxInt32 {
		return -1
	}
	
	return maxDist
}

type Edge struct {
	to     int
	weight int
}

// DP with Bellman-Ford
func networkDelayTimeBellmanFord(times [][]int, n int, k int) int {
	if len(times) == 0 {
		return -1
	}
	
	// Initialize distances
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	
	// Relax edges n-1 times
	for i := 1; i < n; i++ {
		updated := false
		
		for _, time := range times {
			from, to, weight := time[0], time[1], time[2]
			if dist[from] != math.MaxInt32 && dist[from]+weight < dist[to] {
				dist[to] = dist[from] + weight
				updated = true
			}
		}
		
		if !updated {
			break
		}
	}
	
	// Find maximum distance
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	
	if maxDist == math.MaxInt32 {
		return -1
	}
	
	return maxDist
}

// DP with multiple sources optimization
func networkDelayTimeMultipleSources(times [][]int, n int, sources []int) int {
	if len(times) == 0 {
		return -1
	}
	
	// Build adjacency matrix
	dist := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	
	// Fill direct edges
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		dist[from][to] = weight
	}
	
	// Floyd-Warshall
	for intermediate := 1; intermediate <= n; intermediate++ {
		for from := 1; from <= n; from++ {
			for to := 1; to <= n; to++ {
				if dist[from][intermediate] != math.MaxInt32 && 
				   dist[intermediate][to] != math.MaxInt32 {
					newDist := dist[from][intermediate] + dist[intermediate][to]
					if newDist < dist[from][to] {
						dist[from][to] = newDist
					}
				}
			}
		}
	}
	
	// Find minimum maximum distance among sources
	minMaxDist := math.MaxInt32
	
	for _, source := range sources {
		maxDist := 0
		for i := 1; i <= n; i++ {
			if i != source && dist[source][i] > maxDist {
				maxDist = dist[source][i]
			}
		}
		
		if maxDist < minMaxDist {
			minMaxDist = maxDist
		}
	}
	
	if minMaxDist == math.MaxInt32 {
		return -1
	}
	
	return minMaxDist
}

// DP with path reconstruction
func networkDelayTimeWithPath(times [][]int, n int, k int) (int, [][]int) {
	if len(times) == 0 {
		return -1, [][]int{}
	}
	
	// Build adjacency matrix with next pointers
	dist := make([][]int, n+1)
	next := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dist[i] = make([]int, n+1)
		next[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
				next[i][j] = -1
			}
		}
	}
	
	// Fill direct edges
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		if weight < dist[from][to] {
			dist[from][to] = weight
			next[from][to] = to
		}
	}
	
	// Floyd-Warshall with path reconstruction
	for intermediate := 1; intermediate <= n; intermediate++ {
		for from := 1; from <= n; from++ {
			for to := 1; to <= n; to++ {
				if dist[from][intermediate] != math.MaxInt32 && 
				   dist[intermediate][to] != math.MaxInt32 {
					newDist := dist[from][intermediate] + dist[intermediate][to]
					if newDist < dist[from][to] {
						dist[from][to] = newDist
						next[from][to] = next[from][intermediate]
					}
				}
			}
		}
	}
	
	// Find maximum distance and reconstruct paths
	maxDist := 0
	var paths [][]int
	
	for i := 1; i <= n; i++ {
		if i != k && dist[k][i] > maxDist {
			maxDist = dist[k][i]
			paths = [][]int{reconstructPath(next, k, i)}
		}
	}
	
	if maxDist == math.MaxInt32 {
		return -1, [][]int{}
	}
	
	return maxDist, paths
}

func reconstructPath(next [][]int, from, to int) []int {
	if next[from][to] == -1 {
		return []int{from, to}
	}
	
	path := []int{from}
	current := from
	
	for current != to {
		current = next[current][to]
		path = append(path, current)
	}
	
	path = append(path, to)
	return path
}

// DP with early termination
func networkDelayTimeEarlyTermination(times [][]int, n int, k int) int {
	if len(times) == 0 {
		return -1
	}
	
	// Build adjacency list
	adj := make([][]Edge, n+1)
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		adj[from] = append(adj[from], Edge{to, weight})
	}
	
	// Dijkstra with early termination when all nodes are visited
	dist := make([]int, n+1)
	visited := make([]bool, n+1)
	
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	
	visitedCount := 0
	
	for visitedCount < n {
		// Find unvisited node with minimum distance
		u := -1
		minDist := math.MaxInt32
		
		for v := 1; v <= n; v++ {
			if !visited[v] && dist[v] < minDist {
				minDist = dist[v]
				u = v
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		visitedCount++
		
		// Update distances
		for _, edge := range adj[u] {
			to, weight := edge.to, edge.weight
			if dist[u]+weight < dist[to] {
				dist[to] = dist[u] + weight
			}
		}
	}
	
	// Find maximum distance
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	
	if maxDist == math.MaxInt32 {
		return -1
	}
	
	return maxDist
}

func main() {
	// Test cases
	fmt.Println("=== Testing Network Delay Time - DP on Graphs ===")
	
	testCases := []struct {
		times      [][]int
		n          int
		k          int
		description string
	}{
		{
			[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			4,
			2,
			"Standard case",
		},
		{
			[][]int{{1, 2, 1}},
			2,
			1,
			"Single edge",
		},
		{
			[][]int{{1, 2, 1}, {1, 3, 2}},
			3,
			1,
			"Disconnected",
		},
		{
			[][]int{{1, 2, 5}, {2, 3, 1}, {3, 4, 1}},
			4,
			2,
			"Long chain",
		},
		{
			[][]int{{1, 2, 1}, {2, 3, 10}, {3, 4, 1}},
			4,
			1,
			"Weight variation",
		},
		{
			[][]int{},
			1,
			1,
			"No edges",
		},
		{
			[][]int{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}, {4, 5, 1}},
			5,
			3,
			"Complete graph",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Times: %v, N: %d, K: %d\n", tc.times, tc.n, tc.k)
		
		result1 := networkDelayTimeDP(tc.times, tc.n, tc.k)
		result2 := networkDelayTimeDijkstraDP(tc.times, tc.n, tc.k)
		result3 := networkDelayTimeBellmanFord(tc.times, tc.n, tc.k)
		
		fmt.Printf("  Floyd-Warshall: %d\n", result1)
		fmt.Printf("  Dijkstra: %d\n", result2)
		fmt.Printf("  Bellman-Ford: %d\n", result3)
		
		// Test path reconstruction
		maxDist, paths := networkDelayTimeWithPath(tc.times, tc.n, tc.k)
		fmt.Printf("  With paths: max_dist=%d, paths=%v\n", maxDist, paths)
		
		fmt.Println()
	}
	
	// Test multiple sources
	fmt.Println("=== Multiple Sources Test ===")
	sources := []int{1, 3}
	result := networkDelayTimeMultipleSources([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, sources)
	fmt.Printf("Sources %v: %d\n", sources, result)
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	// Generate large graph
	largeN := 100
	largeTimes := make([][]int, 0)
	for i := 1; i <= largeN; i++ {
		for j := i + 1; j <= largeN && j <= i+5; j++ {
			weight := (j - i) * 10
			largeTimes = append(largeTimes, []int{i, j, weight})
		}
	}
	
	fmt.Printf("Large test with %d nodes and %d edges\n", largeN, len(largeTimes))
	
	result := networkDelayTimeDP(largeTimes, largeN, 1)
	fmt.Printf("Floyd-Warshall result: %d\n", result)
	
	result = networkDelayTimeDijkstraDP(largeTimes, largeN, 1)
	fmt.Printf("Dijkstra result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Single node
	fmt.Printf("Single node: %d\n", networkDelayTimeDP([][]int{}, 1, 1))
	
	// Self loop
	selfLoop := [][]int{{1, 1, 5}}
	fmt.Printf("Self loop: %d\n", networkDelayTimeDP(selfLoop, 1, 1))
	
	// Multiple edges same direction
	multiEdges := [][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}}
	fmt.Printf("Multiple edges: %d\n", networkDelayTimeDP(multiEdges, 2, 1))
	
	// Very large weights
	largeWeights := [][]int{{1, 2, 1000000}, {2, 3, 2000000}}
	fmt.Printf("Large weights: %d\n", networkDelayTimeDP(largeWeights, 3, 1))
	
	// Test early termination
	fmt.Println("\n=== Early Termination Test ===")
	earlyResult := networkDelayTimeEarlyTermination([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 1)
	fmt.Printf("Early termination: %d\n", earlyResult)
}
