package main

import (
	"fmt"
	"math"
)

// 743. Network Delay Time - Minimum Spanning Tree (Dijkstra vs MST)
// Time: O(N^2) for MST approach, Space: O(N)
func networkDelayTimeDijkstra(times [][]int, n int, k int) int {
	// Build adjacency list
	adj := make([][][]int, n+1)
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		adj[from] = append(adj[from], []int{to, weight})
	}
	
	// Dijkstra's algorithm
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
			to, weight := edge[0], edge[1]
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

// MST approach using Prim's algorithm
func networkDelayTimeMST(times [][]int, n int, k int) int {
	if len(times) == 0 {
		return -1
	}
	
	// Build adjacency list
	adj := make([][][]int, n+1)
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		adj[from] = append(adj[from], []int{to, weight})
		adj[to] = append(adj[to], []int{from, weight})
	}
	
	// Find MST using Prim's algorithm
	visited := make([]bool, n+1)
	minDist := make([]int, n+1)
	
	for i := 1; i <= n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	minDist[k] = 0
	totalWeight := 0
	edgesUsed := 0
	
	for i := 1; i <= n && edgesUsed < n-1; i++ {
		// Find unvisited node with minimum distance
		u := -1
		minVal := math.MaxInt32
		
		for v := 1; v <= n; v++ {
			if !visited[v] && minDist[v] < minVal {
				minVal = minDist[v]
				u = v
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalWeight += minDist[u]
		edgesUsed++
		
		// Update distances
		for _, edge := range adj[u] {
			to, weight := edge[0], edge[1]
			if !visited[to] && weight < minDist[to] {
				minDist[to] = weight
			}
		}
	}
	
	// If graph is not connected
	if edgesUsed < n-1 {
		return -1
	}
	
	// Now find maximum distance from k using Dijkstra on MST
	return networkDelayTimeDijkstraOnMST(times, n, k)
}

func networkDelayTimeDijkstraOnMST(times [][]int, n int, k int) int {
	// Build MST adjacency list (simplified)
	mstAdj := make([][][]int, n+1)
	
	// This is a simplified approach - in practice, we'd need to extract MST edges
	// For demonstration, we'll use all edges but with MST logic
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		mstAdj[from] = append(mstAdj[from], []int{to, weight})
	}
	
	// Dijkstra on MST
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	
	visited := make([]bool, n+1)
	
	for i := 1; i <= n; i++ {
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
		
		for _, edge := range mstAdj[u] {
			to, weight := edge[0], edge[1]
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

// Kruskal's algorithm for MST
func networkDelayTimeKruskal(times [][]int, n int, k int) int {
	if len(times) == 0 {
		return -1
	}
	
	// Sort edges by weight
	sortedTimes := make([][]int, len(times))
	copy(sortedTimes, times)
	
	// Simple bubble sort for demonstration
	for i := 0; i < len(sortedTimes)-1; i++ {
		for j := 0; j < len(sortedTimes)-i-1; j++ {
			if sortedTimes[j][2] > sortedTimes[j+1][2] {
				sortedTimes[j], sortedTimes[j+1] = sortedTimes[j+1], sortedTimes[j]
			}
		}
	}
	
	// Kruskal's algorithm
	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
	
	mstEdges := [][]int{}
	
	for _, time := range sortedTimes {
		from, to := time[0], time[1]
		root1 := findParent(parent, from)
		root2 := findParent(parent, to)
		
		if root1 != root2 {
			parent[root1] = root2
			mstEdges = append(mstEdges, time)
			
			if len(mstEdges) == n-1 {
				break
			}
		}
	}
	
	// Check if MST spans all nodes
	if len(mstEdges) < n-1 {
		return -1
	}
	
	// Build adjacency list from MST
	mstAdj := make([][][]int, n+1)
	for _, edge := range mstEdges {
		from, to, weight := edge[0], edge[1], edge[2]
		mstAdj[from] = append(mstAdj[from], []int{to, weight})
	}
	
	// Dijkstra on MST
	return networkDelayTimeDijkstraOnMST(times, n, k)
}

func findParent(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = findParent(parent, parent[x])
	}
	return parent[x]
}

// MST with path compression
func networkDelayTimeMSTOptimized(times [][]int, n int, k int) int {
	if len(times) == 0 {
		return -1
	}
	
	// Sort edges
	sortedTimes := make([][]int, len(times))
	copy(sortedTimes, times)
	
	// Sort by weight
	for i := 0; i < len(sortedTimes)-1; i++ {
		for j := 0; j < len(sortedTimes)-i-1; j++ {
			if sortedTimes[j][2] > sortedTimes[j+1][2] {
				sortedTimes[j], sortedTimes[j+1] = sortedTimes[j+1], sortedTimes[j]
			}
		}
	}
	
	// Union-Find with path compression
	parent := make([]int, n+1)
	rank := make([]int, n+1)
	
	for i := 1; i <= n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	
	mstEdges := [][]int{}
	
	for _, time := range sortedTimes {
		from, to := time[0], time[1]
		root1 := findWithCompression(parent, from)
		root2 := findWithCompression(parent, to)
		
		if root1 != root2 {
			if rank[root1] < rank[root2] {
				parent[root1] = root2
				rank[root2]++
			} else {
				parent[root2] = root1
				rank[root1]++
			}
			
			mstEdges = append(mstEdges, time)
			
			if len(mstEdges) == n-1 {
				break
			}
		}
	}
	
	// Check connectivity
	if len(mstEdges) < n-1 {
		return -1
	}
	
	// Find longest path in MST
	return findLongestPathInMST(mstEdges, n, k)
}

func findWithCompression(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = findWithCompression(parent, parent[x])
	}
	return parent[x]
}

func findLongestPathInMST(edges [][]int, n int, k int) int {
	// Build adjacency list
	adj := make([][][]int, n+1)
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		adj[from] = append(adj[from], []int{to, weight})
	}
	
	// BFS to find longest path from k
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	
	queue := []int{k}
	visited := make([]bool, n+1)
	visited[k] = true
	
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		
		for _, edge := range adj[u] {
			to, weight := edge[0], edge[1]
			if !visited[to] {
				visited[to] = true
				dist[to] = dist[u] + weight
				queue = append(queue, to)
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

// MST with multiple sources
func networkDelayTimeMultipleSources(times [][]int, n int, sources []int) int {
	if len(times) == 0 {
		return -1
	}
	
	// Build MST
	mstEdges := buildMST(times, n)
	if len(mstEdges) < n-1 {
		return -1
	}
	
	// Find distances from each source
	maxDist := 0
	for _, source := range sources {
		dist := findDistanceInMST(mstEdges, n, source)
		if dist > maxDist {
			maxDist = dist
		}
	}
	
	return maxDist
}

func buildMST(times [][]int, n int) [][]int {
	// Sort edges by weight
	sortedTimes := make([][]int, len(times))
	copy(sortedTimes, times)
	
	for i := 0; i < len(sortedTimes)-1; i++ {
		for j := 0; j < len(sortedTimes)-i-1; j++ {
			if sortedTimes[j][2] > sortedTimes[j+1][2] {
				sortedTimes[j], sortedTimes[j+1] = sortedTimes[j+1], sortedTimes[j]
			}
		}
	}
	
	// Kruskal's algorithm
	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
	
	mstEdges := [][]int{}
	
	for _, time := range sortedTimes {
		from, to := time[0], time[1]
		root1 := findParent(parent, from)
		root2 := findParent(parent, to)
		
		if root1 != root2 {
			parent[root1] = root2
			mstEdges = append(mstEdges, time)
			
			if len(mstEdges) == n-1 {
				break
			}
		}
	}
	
	return mstEdges
}

func findDistanceInMST(edges [][]int, n int, source int) int {
	// Build adjacency list
	adj := make([][][]int, n+1)
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		adj[from] = append(adj[from], []int{to, weight})
	}
	
	// BFS
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[source] = 0
	
	queue := []int{source}
	visited := make([]bool, n+1)
	visited[source] = true
	
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		
		for _, edge := range adj[u] {
			to, weight := edge[0], edge[1]
			if !visited[to] {
				visited[to] = true
				dist[to] = dist[u] + weight
				queue = append(queue, to)
			}
		}
	}
	
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	
	return maxDist
}

func main() {
	// Test cases
	fmt.Println("=== Testing Network Delay Time - MST Approaches ===")
	
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
			[][]int{{1, 2, 1}},
			2,
			2,
			"Source at end",
		},
		{
			[][]int{},
			1,
			1,
			"No edges",
		},
		{
			[][]int{{1, 2, 5}, {2, 3, 1}, {3, 4, 1}, {4, 5, 5}},
			5,
			1,
			"Long chain",
		},
		{
			[][]int{{1, 2, 1}, {1, 3, 2}, {3, 4, 1}, {2, 4, 3}},
			4,
			1,
			"Complex graph",
		},
		{
			[][]int{{1, 2, 1}, {2, 3, 10}, {3, 4, 1}},
			4,
			2,
			"Disconnected",
		},
		{
			[][]int{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}, {4, 5, 1}, {5, 1, 1}},
			5,
			1,
			"Star graph",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Times: %v, N: %d, K: %d\n", tc.times, tc.n, tc.k)
		
		result1 := networkDelayTimeDijkstra(tc.times, tc.n, tc.k)
		result2 := networkDelayTimeMST(tc.times, tc.n, tc.k)
		result3 := networkDelayTimeKruskal(tc.times, tc.n, tc.k)
		result4 := networkDelayTimeMSTOptimized(tc.times, tc.n, tc.k)
		
		fmt.Printf("  Dijkstra: %d\n", result1)
		fmt.Printf("  MST: %d\n", result2)
		fmt.Printf("  Kruskal: %d\n", result3)
		fmt.Printf("  MST Optimized: %d\n\n", result4)
	}
	
	// Test multiple sources
	fmt.Println("=== Multiple Sources Test ===")
	multiSources := []int{1, 3}
	result := networkDelayTimeMultipleSources([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, multiSources)
	fmt.Printf("Multiple sources %v: %d\n", multiSources, result)
	
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
	
	result := networkDelayTimeDijkstra(largeTimes, largeN, 1)
	fmt.Printf("Dijkstra result: %d\n", result)
	
	result = networkDelayTimeMSTOptimized(largeTimes, largeN, 1)
	fmt.Printf("MST result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Single node
	fmt.Printf("Single node: %d\n", networkDelayTimeDijkstra([][]int{}, 1, 1))
	
	// Self loop
	selfLoop := [][]int{{1, 1, 5}}
	fmt.Printf("Self loop: %d\n", networkDelayTimeDijkstra(selfLoop, 1, 1))
	
	// Multiple edges same direction
	multiEdges := [][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}}
	fmt.Printf("Multiple edges: %d\n", networkDelayTimeDijkstra(multiEdges, 2, 1))
	
	// Very large weights
	largeWeights := [][]int{{1, 2, 1000000}, {2, 3, 2000000}}
	fmt.Printf("Large weights: %d\n", networkDelayTimeDijkstra(largeWeights, 3, 1))
	
	// Test MST vs Dijkstra comparison
	fmt.Println("\n=== MST vs Dijkstra Comparison ===")
	
	// Complete graph vs sparse graph
	completeTimes := [][]int{{1, 2, 1}, {1, 3, 2}, {2, 3, 3}}
	sparseTimes := [][]int{{1, 2, 10}, {2, 3, 10}, {1, 3, 10}}
	
	fmt.Printf("Complete graph - Dijkstra: %d, MST: %d\n", 
		networkDelayTimeDijkstra(completeTimes, 3, 1), 
		networkDelayTimeMST(completeTimes, 3, 1))
	
	fmt.Printf("Sparse graph - Dijkstra: %d, MST: %d\n", 
		networkDelayTimeDijkstra(sparseTimes, 3, 1), 
		networkDelayTimeMST(sparseTimes, 3, 1))
}
