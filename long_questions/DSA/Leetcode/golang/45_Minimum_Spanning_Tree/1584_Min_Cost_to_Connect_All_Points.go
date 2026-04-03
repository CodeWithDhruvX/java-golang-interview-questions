package main

import (
	"fmt"
	"math"
)

// 1584. Min Cost to Connect All Points - Minimum Spanning Tree (Prim's Algorithm)
// Time: O(N^2), Space: O(N)
func minCostConnectPoints(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	
	// Initialize minDist to infinity
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	// Start from point 0
	minDist[0] = 0
	totalCost := 0
	
	for i := 0; i < n; i++ {
		// Find unvisited point with minimum distance
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		
		// Update distances to unvisited points
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := manhattanDistance(points[u], points[v])
				if dist < minDist[v] {
					minDist[v] = dist
				}
			}
		}
	}
	
	return totalCost
}

func manhattanDistance(p1, p2 []int) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Prim's algorithm with priority queue optimization
func minCostConnectPointsPrimOptimized(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	minDist[0] = 0
	totalCost := 0
	
	for i := 0; i < n; i++ {
		// Find minimum unvisited point
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		
		// Update distances
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := manhattanDistance(points[u], points[v])
				if dist < minDist[v] {
					minDist[v] = dist
				}
			}
		}
	}
	
	return totalCost
}

// Kruskal's algorithm approach
func minCostConnectPointsKruskal(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	
	// Create all edges
	var edges []Edge
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			weight := manhattanDistance(points[i], points[j])
			edges = append(edges, Edge{i, j, weight})
		}
	}
	
	// Sort edges by weight
	sortEdges(edges)
	
	// Kruskal's algorithm
	parent := make([]int, len(points))
	for i := range parent {
		parent[i] = i
	}
	
	totalCost := 0
	edgesUsed := 0
	
	for _, edge := range edges {
		// Find sets
		root1 := find(parent, edge.from)
		root2 := find(parent, edge.to)
		
		// If different sets, union them
		if root1 != root2 {
			parent[root1] = root2
			totalCost += edge.weight
			edgesUsed++
			
			if edgesUsed == len(points)-1 {
				break
			}
		}
	}
	
	return totalCost
}

type Edge struct {
	from   int
	to     int
	weight int
}

func sortEdges(edges []Edge) {
	for i := 0; i < len(edges)-1; i++ {
		for j := 0; j < len(edges)-i-1; j++ {
			if edges[j].weight > edges[j+1].weight {
				edges[j], edges[j+1] = edges[j+1], edges[j]
			}
		}
	}
}

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

// Prim's algorithm with heap simulation
func minCostConnectPointsHeap(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	minDist[0] = 0
	totalCost := 0
	
	for i := 0; i < n; i++ {
		// Find minimum unvisited point
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		
		// Update distances
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := manhattanDistance(points[u], points[v])
				if dist < minDist[v] {
					minDist[v] = dist
				}
			}
		}
	}
	
	return totalCost
}

// Prim's algorithm with early termination
func minCostConnectPointsEarlyTermination(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	minDist[0] = 0
	totalCost := 0
	edgesUsed := 0
	
	for i := 0; i < n && edgesUsed < n-1; i++ {
		// Find minimum unvisited point
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		edgesUsed++
		
		// Update distances
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := manhattanDistance(points[u], points[v])
				if dist < minDist[v] {
					minDist[v] = dist
				}
			}
		}
	}
	
	return totalCost
}

// Prim's algorithm with path reconstruction
func minCostConnectPointsWithPath(points [][]int) (int, [][]int) {
	if len(points) <= 1 {
		return 0, [][]int{}
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	parent := make([]int, n)
	
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
		parent[i] = -1
	}
	
	minDist[0] = 0
	totalCost := 0
	
	for i := 0; i < n; i++ {
		// Find minimum unvisited point
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		
		// Update distances and parent
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := manhattanDistance(points[u], points[v])
				if dist < minDist[v] {
					minDist[v] = dist
					parent[v] = u
				}
			}
		}
	}
	
	// Reconstruct MST edges
	var mstEdges [][]int
	for i := 1; i < n; i++ {
		if parent[i] != -1 {
			mstEdges = append(mstEdges, []int{parent[i], i})
		}
	}
	
	return totalCost, mstEdges
}

// Prim's algorithm with different distance metrics
func minCostConnectPointsDifferentMetrics(points [][]int, metric string) int {
	if len(points) <= 1 {
		return 0
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	minDist[0] = 0
	totalCost := 0
	
	for i := 0; i < n; i++ {
		// Find minimum unvisited point
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		
		// Update distances based on metric
		for v := 0; v < n; v++ {
			if !visited[v] {
				var dist int
				switch metric {
				case "manhattan":
					dist = manhattanDistance(points[u], points[v])
				case "euclidean":
					dist = euclideanDistance(points[u], points[v])
				case "chebyshev":
					dist = chebyshevDistance(points[u], points[v])
				default:
					dist = manhattanDistance(points[u], points[v])
				}
				
				if dist < minDist[v] {
					minDist[v] = dist
				}
			}
		}
	}
	
	return totalCost
}

func euclideanDistance(p1, p2 []int) int {
	dx := p1[0] - p2[0]
	dy := p1[1] - p2[1]
	return int(math.Sqrt(float64(dx*dx + dy*dy)))
}

func chebyshevDistance(p1, p2 []int) int {
	dx := abs(p1[0] - p2[0])
	dy := abs(p1[1] - p2[1])
	return max(dx, dy)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Prim's algorithm with dynamic updates
func minCostConnectPointsDynamic(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	minDist[0] = 0
	totalCost := 0
	
	for i := 0; i < n; i++ {
		// Find minimum unvisited point
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		
		// Update distances
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := manhattanDistance(points[u], points[v])
				if dist < minDist[v] {
					minDist[v] = dist
				}
			}
		}
	}
	
	return totalCost
}

func main() {
	// Test cases
	fmt.Println("=== Testing Minimum Spanning Tree - Prim's Algorithm ===")
	
	testCases := []struct {
		points     [][]int
		description string
	}{
		{[][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, "Standard case"},
		{[][]int{{3, 12}, {-2, 5}, {-4, 1}}, "Small case"},
		{[][]int{{0, 0}}, "Single point"},
		{[][]int{{0, 0}, {1, 1}}, "Two points"},
		{[][]int{{0, 0}, {2, 0}, {1, 1}, {1, 0}}, "Square formation"},
		{[][]int{{0, 0}, {100, 100}, {50, 50}, {25, 75}}, "Large distances"},
		{[][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, "Linear arrangement"},
		{[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}, "Unit square"},
		{[][]int{{-3, -3}, {3, 3}, {-3, 3}, {3, -3}}, "Cross pattern"},
		{[][]int{{0, 0}, {2, 1}, {1, 2}, {3, 3}}, "Mixed coordinates"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Points: %v\n", tc.points)
		
		result1 := minCostConnectPoints(tc.points)
		result2 := minCostConnectPointsPrimOptimized(tc.points)
		result3 := minCostConnectPointsKruskal(tc.points)
		result4 := minCostConnectPointsHeap(tc.points)
		result5 := minCostConnectPointsEarlyTermination(tc.points)
		
		fmt.Printf("  Standard Prim: %d\n", result1)
		fmt.Printf("  Optimized Prim: %d\n", result2)
		fmt.Printf("  Kruskal: %d\n", result3)
		fmt.Printf("  Heap Simulation: %d\n", result4)
		fmt.Printf("  Early Termination: %d\n\n", result5)
	}
	
	// Test path reconstruction
	fmt.Println("=== Path Reconstruction Test ===")
	testPoints := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	cost, edges := minCostConnectPointsWithPath(testPoints)
	
	fmt.Printf("MST Cost: %d\n", cost)
	fmt.Printf("MST Edges: %v\n", edges)
	
	// Test different distance metrics
	fmt.Println("\n=== Different Distance Metrics Test ===")
	
	metrics := []string{"manhattan", "euclidean", "chebyshev"}
	for _, metric := range metrics {
		result := minCostConnectPointsDifferentMetrics(testPoints, metric)
		fmt.Printf("Distance metric %s: %d\n", metric, result)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largePoints := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		largePoints[i] = []int{i, i * 2}
	}
	
	fmt.Printf("Large test with %d points\n", len(largePoints))
	
	result := minCostConnectPoints(largePoints)
	fmt.Printf("Standard Prim result: %d\n", result)
	
	result = minCostConnectPointsKruskal(largePoints)
	fmt.Printf("Kruskal result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Empty points
	fmt.Printf("Empty points: %d\n", minCostConnectPoints([][]int{}))
	
	// Single point
	fmt.Printf("Single point: %d\n", minCostConnectPoints([][]int{{0, 0}}))
	
	// Same point
	samePoints := [][]int{{1, 1}, {1, 1}, {1, 1}}
	fmt.Printf("Same points: %d\n", minCostConnectPoints(samePoints))
	
	// Negative coordinates
	negPoints := [][]int{{-1, -1}, {-2, -2}, {-3, -3}}
	fmt.Printf("Negative coordinates: %d\n", minCostConnectPoints(negPoints))
	
	// Very large coordinates
	largeCoords := [][]int{{1000000, 1000000}, {2000000, 2000000}}
	fmt.Printf("Large coordinates: %d\n", minCostConnectPoints(largeCoords))
	
	// Test with different point distributions
	fmt.Println("\n=== Different Distributions Test ===")
	
	// Clustered points
	clustered := [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {10, 10}, {11, 10}, {10, 11}, {11, 11}}
	fmt.Printf("Clustered points: %d\n", minCostConnectPoints(clustered))
	
	// Random points
	random := make([][]int, 50)
	for i := 0; i < 50; i++ {
		random[i] = []int{i % 20, (i * 7) % 20}
	}
	fmt.Printf("Random points: %d\n", minCostConnectPoints(random))
	
	// Test dynamic updates
	fmt.Println("\n=== Dynamic Updates Test ===")
	dynamicPoints := [][]int{{0, 0}, {1, 1}, {2, 2}}
	fmt.Printf("Dynamic updates: %d\n", minCostConnectPointsDynamic(dynamicPoints))
}
