package main

import (
	"fmt"
	"math"
)

// 1334. Find the City With the Smallest Number of Neighbors at a Threshold Distance - Floyd-Warshall Algorithm
// Time: O(N^3), Space: O(N^2)
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// Initialize distance matrix
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	
	// Fill direct edges
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		if weight < dist[from][to] {
			dist[from][to] = weight
			dist[to][from] = weight // Undirected graph
		}
	}
	
	// Floyd-Warshall algorithm
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					if dist[i][j] > dist[i][k]+dist[k][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
					}
				}
			}
		}
	}
	
	// Find city with minimum neighbors within threshold
	minNeighbors := math.MaxInt32
	result := -1
	
	for i := 0; i < n; i++ {
		neighbors := 0
		for j := 0; j < n; j++ {
			if i != j && dist[i][j] <= distanceThreshold {
				neighbors++
			}
		}
		
		if neighbors < minNeighbors || (neighbors == minNeighbors && i > result) {
			minNeighbors = neighbors
			result = i
		}
	}
	
	return result
}

// Optimized Floyd-Warshall with early termination
func findTheCityOptimized(n int, edges [][]int, distanceThreshold int) int {
	// Initialize distance matrix
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	
	// Fill direct edges
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		if weight < dist[from][to] {
			dist[from][to] = weight
			dist[to][from] = weight
		}
	}
	
	// Floyd-Warshall with optimization
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if dist[i][k] == math.MaxInt32 {
				continue // Skip if no path to k
			}
			for j := 0; j < n; j++ {
				if dist[k][j] == math.MaxInt32 {
					continue // Skip if no path from k
				}
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	
	// Find city with minimum neighbors
	minNeighbors := math.MaxInt32
	result := -1
	
	for i := 0; i < n; i++ {
		neighbors := 0
		for j := 0; j < n; j++ {
			if i != j && dist[i][j] <= distanceThreshold {
				neighbors++
			}
		}
		
		if neighbors < minNeighbors || (neighbors == minNeighbors && i > result) {
			minNeighbors = neighbors
			result = i
		}
	}
	
	return result
}

// Floyd-Warshall with path reconstruction
func findTheCityWithPathReconstruction(n int, edges [][]int, distanceThreshold int) (int, [][]int) {
	// Initialize distance and path matrices
	dist := make([][]int, n)
	next := make([][]int, n)
	
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		next[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
				next[i][j] = j
			} else {
				dist[i][j] = math.MaxInt32
				next[i][j] = -1
			}
		}
	}
	
	// Fill direct edges
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		if weight < dist[from][to] {
			dist[from][to] = weight
			next[from][to] = to
			dist[to][from] = weight
			next[to][from] = from
		}
	}
	
	// Floyd-Warshall with path reconstruction
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					if dist[i][j] > dist[i][k]+dist[k][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
						next[i][j] = next[i][k]
					}
				}
			}
		}
	}
	
	// Find city with minimum neighbors
	minNeighbors := math.MaxInt32
	result := -1
	
	for i := 0; i < n; i++ {
		neighbors := 0
		for j := 0; j < n; j++ {
			if i != j && dist[i][j] <= distanceThreshold {
				neighbors++
			}
		}
		
		if neighbors < minNeighbors || (neighbors == minNeighbors && i > result) {
			minNeighbors = neighbors
			result = i
		}
	}
	
	return result, next
}

// Reconstruct path between two cities
func reconstructPath(next [][]int, from, to int) []int {
	if next[from][to] == -1 {
		return []int{}
	}
	
	path := []int{from}
	for from != to {
		from = next[from][to]
		path = append(path, from)
	}
	
	return path
}

// Floyd-Warshall with multiple sources optimization
func findTheCityMultipleSources(n int, edges [][]int, distanceThreshold int, sources []int) int {
	// Initialize distance matrix
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	
	// Fill direct edges
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		if weight < dist[from][to] {
			dist[from][to] = weight
			dist[to][from] = weight
		}
	}
	
	// Floyd-Warshall
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					if dist[i][j] > dist[i][k]+dist[k][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
					}
				}
			}
		}
	}
	
	// Find best source city
	minNeighbors := math.MaxInt32
	result := -1
	
	for _, source := range sources {
		neighbors := 0
		for j := 0; j < n; j++ {
			if source != j && dist[source][j] <= distanceThreshold {
				neighbors++
			}
		}
		
		if neighbors < minNeighbors || (neighbors == minNeighbors && source > result) {
			minNeighbors = neighbors
			result = source
		}
	}
	
	return result
}

// Floyd-Warshall with dynamic threshold
func findTheCityDynamicThreshold(n int, edges [][]int, distanceThreshold int) (int, map[int]int) {
	// Initialize distance matrix
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	
	// Fill direct edges
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		if weight < dist[from][to] {
			dist[from][to] = weight
			dist[to][from] = weight
		}
	}
	
	// Floyd-Warshall
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					if dist[i][j] > dist[i][k]+dist[k][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
					}
				}
			}
		}
	}
	
	// Calculate neighbors for each threshold
	thresholdToCity := make(map[int]int)
	
	for threshold := 1; threshold <= distanceThreshold; threshold++ {
		minNeighbors := math.MaxInt32
		bestCity := -1
		
		for i := 0; i < n; i++ {
			neighbors := 0
			for j := 0; j < n; j++ {
				if i != j && dist[i][j] <= threshold {
					neighbors++
				}
			}
			
			if neighbors < minNeighbors || (neighbors == minNeighbors && i > bestCity) {
				minNeighbors = neighbors
				bestCity = i
			}
		}
		
		thresholdToCity[threshold] = bestCity
	}
	
	// Return result for original threshold
	return thresholdToCity[distanceThreshold], thresholdToCity
}

// Alternative approach using Dijkstra from each city
func findTheCityDijkstra(n int, edges [][]int, distanceThreshold int) int {
	// Build adjacency list
	adj := make([][][]int, n)
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		adj[from] = append(adj[from], []int{to, weight})
		adj[to] = append(adj[to], []int{from, weight})
	}
	
	minNeighbors := math.MaxInt32
	result := -1
	
	// Run Dijkstra from each city
	for i := 0; i < n; i++ {
		distances := dijkstra(n, adj, i)
		
		neighbors := 0
		for j := 0; j < n; j++ {
			if i != j && distances[j] <= distanceThreshold {
				neighbors++
			}
		}
		
		if neighbors < minNeighbors || (neighbors == minNeighbors && i > result) {
			minNeighbors = neighbors
			result = i
		}
	}
	
	return result
}

func dijkstra(n int, adj [][]int, start int) []int {
	dist := make([]int, n)
	visited := make([]bool, n)
	
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0
	
	for count := 0; count < n; count++ {
		// Find unvisited vertex with minimum distance
		minDist := math.MaxInt32
		u := -1
		
		for v := 0; v < n; v++ {
			if !visited[v] && dist[v] < minDist {
				minDist = dist[v]
				u = v
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		
		// Update distances of adjacent vertices
		for _, edge := range adj[u] {
			v, weight := edge[0], edge[1]
			if !visited[v] && dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
			}
		}
	}
	
	return dist
}

func main() {
	// Test cases
	fmt.Println("=== Testing Floyd-Warshall Algorithm ===")
	
	testCases := []struct {
		n                 int
		edges             [][]int
		distanceThreshold int
		description       string
	}{
		{
			4,
			[][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}},
			4,
			"Standard case",
		},
		{
			5,
			[][]int{{0, 1, 2}, {0, 2, 4}, {1, 2, 1}, {1, 3, 2}, {2, 3, 3}, {3, 4, 1}},
			5,
			"Larger graph",
		},
		{
			3,
			[][]int{{0, 1, 1}, {1, 2, 1}, {2, 0, 1}},
			2,
			"Triangle graph",
		},
		{
			6,
			[][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}, {2, 3, 1}, {3, 4, 1}, {4, 5, 1}, {5, 0, 10}},
			20,
			"Complex graph",
		},
		{
			2,
			[][]int{{0, 1, 1}},
			1,
			"Two cities",
		},
		{
			1,
			[][]int{},
			0,
			"Single city",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Cities: %d, Threshold: %d\n", tc.n, tc.distanceThreshold)
		fmt.Printf("  Edges: %v\n", tc.edges)
		
		result1 := findTheCity(tc.n, tc.edges, tc.distanceThreshold)
		result2 := findTheCityOptimized(tc.n, tc.edges, tc.distanceThreshold)
		result3 := findTheCityDijkstra(tc.n, tc.edges, tc.distanceThreshold)
		
		fmt.Printf("  Floyd-Warshall: %d\n", result1)
		fmt.Printf("  Optimized: %d\n", result2)
		fmt.Printf("  Dijkstra: %d\n\n", result3)
	}
	
	// Test path reconstruction
	fmt.Println("=== Path Reconstruction Test ===")
	n, edges, threshold := 4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4
	city, next := findTheCityWithPathReconstruction(n, edges, threshold)
	
	fmt.Printf("Best city: %d\n", city)
	fmt.Printf("Path from 0 to 3: %v\n", reconstructPath(next, 0, 3))
	fmt.Printf("Path from 0 to 2: %v\n", reconstructPath(next, 0, 2))
	
	// Test multiple sources
	fmt.Println("\n=== Multiple Sources Test ===")
	sources := []int{0, 2}
	result := findTheCityMultipleSources(n, edges, threshold, sources)
	fmt.Printf("Best city from sources %v: %d\n", sources, result)
	
	// Test dynamic threshold
	fmt.Println("\n=== Dynamic Threshold Test ===")
	city, thresholdMap := findTheCityDynamicThreshold(n, edges, threshold)
	fmt.Printf("Best city for threshold %d: %d\n", threshold, city)
	
	for t, bestCity := range thresholdMap {
		fmt.Printf("  Threshold %d: City %d\n", t, bestCity)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largeN := 100
	largeEdges := make([][]int, 0)
	for i := 0; i < largeN; i++ {
		for j := i + 1; j < largeN && j < i+5; j++ {
			weight := (j-i) * 10
			largeEdges = append(largeEdges, []int{i, j, weight})
		}
	}
	
	fmt.Printf("Large test with %d cities and %d edges\n", largeN, len(largeEdges))
	
	result = findTheCity(largeN, largeEdges, 50)
	fmt.Printf("Floyd-Warshall result: %d\n", result)
	
	result = findTheCityOptimized(largeN, largeEdges, 50)
	fmt.Printf("Optimized result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// No edges
	noEdges := [][]int{}
	fmt.Printf("No edges: %d\n", findTheCity(3, noEdges, 10))
	
	// Very large threshold
	veryLargeThreshold := 1000
	fmt.Printf("Very large threshold: %d\n", findTheCity(4, edges, veryLargeThreshold))
	
	// Very small threshold
	verySmallThreshold := 0
	fmt.Printf("Very small threshold: %d\n", findTheCity(4, edges, verySmallThreshold))
	
	// Disconnected graph
	disconnectedEdges := [][]int{{0, 1, 1}, {2, 3, 1}}
	fmt.Printf("Disconnected graph: %d\n", findTheCity(4, disconnectedEdges, 5))
}
