package main

import (
	"fmt"
	"math"
)

// 1489. Find Critical and Pseudo-Critical Edges in MST - Minimum Spanning Tree
// Time: O(N^2) for MST, Space: O(N^2)
func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	if n <= 1 {
		return [][]int{}
	}
	
	// Find MST weight using Kruskal's algorithm
	mstWeight := findMSTWeight(n, edges)
	
	// Find critical and pseudo-critical edges
	var result [][]int
	
	for i, edge := range edges {
		// Remove this edge and find new MST weight
		newWeight := findMSTWeightWithoutEdge(n, edges, i)
		
		if newWeight > mstWeight {
			// Critical edge
			result = append(result, []int{edge[0], edge[1], edge[2], 1})
		} else if newWeight == mstWeight {
			// Pseudo-critical edge
			result = append(result, []int{edge[0], edge[1], edge[2], 2})
		}
	}
	
	return result
}

func findMSTWeight(n int, edges [][]int) int {
	if len(edges) < n-1 {
		return math.MaxInt32
	}
	
	// Sort edges by weight
	sortedEdges := make([][]int, len(edges))
	copy(sortedEdges, edges)
	
	for i := 0; i < len(sortedEdges)-1; i++ {
		for j := 0; j < len(sortedEdges)-i-1; j++ {
			if sortedEdges[j][2] > sortedEdges[j+1][2] {
				sortedEdges[j], sortedEdges[j+1] = sortedEdges[j+1], sortedEdges[j]
			}
		}
	}
	
	// Kruskal's algorithm
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	
	totalWeight := 0
	edgesUsed := 0
	
	for _, edge := range sortedEdges {
		from, to, weight := edge[0], edge[1], edge[2]
		root1 := find(parent, from)
		root2 := find(parent, to)
		
		if root1 != root2 {
			parent[root1] = root2
			totalWeight += weight
			edgesUsed++
			
			if edgesUsed == n-1 {
				break
			}
		}
	}
	
	if edgesUsed < n-1 {
		return math.MaxInt32
	}
	
	return totalWeight
}

func findMSTWeightWithoutEdge(n int, edges [][]int, excludeIdx int) int {
	if len(edges) <= n-1 {
		return math.MaxInt32
	}
	
	// Create new edges array excluding the specified edge
	newEdges := make([][]int, 0, len(edges)-1)
	for i, edge := range edges {
		if i != excludeIdx {
			newEdges = append(newEdges, edge)
		}
	}
	
	return findMSTWeight(n, newEdges)
}

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

// Optimized version with early termination
func findCriticalAndPseudoCriticalEdgesOptimized(n int, edges [][]int) [][]int {
	if n <= 1 {
		return [][]int{}
	}
	
	// Find MST weight
	mstWeight := findMSTWeight(n, edges)
	
	// Build MST adjacency list
	mstAdj := buildMSTAdjacency(n, edges)
	
	var result [][]int
	
	for i, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		
		// Check if edge is in MST
		inMST := isInMST(mstAdj, from, to, weight)
		
		if inMST {
			// Check if it's critical by removing it
			newWeight := findMSTWeightWithoutEdge(n, edges, i)
			if newWeight > mstWeight {
				result = append(result, []int{from, to, weight, 1})
			} else {
				result = append(result, []int{from, to, weight, 2})
			}
		} else {
			// Edge not in MST, check if it can create alternative MST
			alternativeWeight := findAlternativeMSTWeight(n, edges, from, to, weight)
			if alternativeWeight == mstWeight {
				result = append(result, []int{from, to, weight, 2})
			}
		}
	}
	
	return result
}

func buildMSTAdjacency(n int, edges [][]int) [][]int {
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, n)
	}
	
	// Build MST and get adjacency
	sortedEdges := make([][]int, len(edges))
	copy(sortedEdges, edges)
	
	// Sort by weight
	for i := 0; i < len(sortedEdges)-1; i++ {
		for j := 0; j < len(sortedEdges)-i-1; j++ {
			if sortedEdges[j][2] > sortedEdges[j+1][2] {
				sortedEdges[j], sortedEdges[j+1] = sortedEdges[j+1], sortedEdges[j]
			}
		}
	}
	
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	
	for _, edge := range sortedEdges {
		from, to := edge[0], edge[1]
		root1 := find(parent, from)
		root2 := find(parent, to)
		
		if root1 != root2 {
			parent[root1] = root2
			adj[from][to] = edge[2]
			adj[to][from] = edge[2]
		}
	}
	
	return adj
}

func isInMST(adj [][]int, from, to, weight int) bool {
	return adj[from][to] == weight
}

func findAlternativeMSTWeight(n int, edges [][]int, newFrom, newTo, newWeight int) int {
	// Add new edge and find MST weight
	newEdges := make([][]int, len(edges)+1)
	copy(newEdges, edges)
	newEdges[len(edges)] = []int{newFrom, newTo, newWeight}
	
	return findMSTWeight(n, newEdges)
}

// Version with Union-Find optimization
func findCriticalAndPseudoCriticalEdgesUnionFind(n int, edges [][]int) [][]int {
	if n <= 1 {
		return [][]int{}
	}
	
	// Sort edges by weight
	sortedEdges := make([][]int, len(edges))
	copy(sortedEdges, edges)
	
	for i := 0; i < len(sortedEdges)-1; i++ {
		for j := 0; j < len(sortedEdges)-i-1; j++ {
			if sortedEdges[j][2] > sortedEdges[j+1][2] {
				sortedEdges[j], sortedEdges[j+1] = sortedEdges[j+1], sortedEdges[j]
			}
		}
	}
	
	// Find MST weight and track used edges
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	
	mstWeight := 0
	usedEdges := make([]bool, len(edges))
	
	for i, edge := range sortedEdges {
		from, to, weight := edge[0], edge[1], edge[2]
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
			
			mstWeight += weight
			usedEdges[i] = true
		}
	}
	
	// Check all edges
	var result [][]int
	
	for i, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		
		if usedEdges[i] {
			// Edge in MST, check if critical
			newWeight := findMSTWeightWithoutEdge(n, edges, i)
			if newWeight > mstWeight {
				result = append(result, []int{from, to, weight, 1})
			} else {
				result = append(result, []int{from, to, weight, 2})
			}
		} else {
			// Edge not in MST, check if pseudo-critical
			alternativeWeight := findAlternativeMSTWeight(n, edges, from, to, weight)
			if alternativeWeight == mstWeight {
				result = append(result, []int{from, to, weight, 2})
			}
		}
	}
	
	return result
}

func findWithCompression(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = findWithCompression(parent, parent[x])
	}
	return parent[x]
}

// Version with detailed analysis
func findCriticalAndPseudoCriticalEdgesDetailed(n int, edges [][]int) ([][]int, map[string]int) {
	if n <= 1 {
		return [][]int{}, map[string]int{}
	}
	
	analysis := make(map[string]int)
	
	// Find MST weight
	mstWeight := findMSTWeight(n, edges)
	analysis["mst_weight"] = mstWeight
	
	// Count total edges
	totalEdges := len(edges)
	analysis["total_edges"] = totalEdges
	
	// Find critical and pseudo-critical edges
	var result [][]int
	criticalCount := 0
	pseudoCriticalCount := 0
	
	for i, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		
		newWeight := findMSTWeightWithoutEdge(n, edges, i)
		
		if newWeight > mstWeight {
			result = append(result, []int{from, to, weight, 1})
			criticalCount++
		} else if newWeight == mstWeight {
			result = append(result, []int{from, to, weight, 2})
			pseudoCriticalCount++
		}
	}
	
	analysis["critical_count"] = criticalCount
	analysis["pseudo_critical_count"] = pseudoCriticalCount
	analysis["non_critical_count"] = totalEdges - criticalCount - pseudoCriticalCount
	
	return result, analysis
}

// Version with multiple MSTs
func findCriticalAndPseudoCriticalEdgesMultipleMST(n int, edges [][]int) [][]int {
	if n <= 1 {
		return [][]int{}
	}
	
	// Find all possible MST weights
	mstWeights := findAllMSTWeights(n, edges)
	
	if len(mstWeights) == 0 {
		return [][]int{}
	}
	
	minWeight := mstWeights[0]
	
	// Analyze each edge
	var result [][]int
	
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		
		// Check if edge appears in any MST
		inAnyMST := false
		for _, mstWeight := range mstWeights {
			if edgeInMST(n, edges, from, to, weight, mstWeight) {
				inAnyMST = true
				break
			}
		}
		
		if inAnyMST {
			// Check if critical
			newWeight := findMSTWeightWithoutEdge(n, edges, findEdgeIndex(edges, edge))
			if newWeight > minWeight {
				result = append(result, []int{from, to, weight, 1})
			} else {
				result = append(result, []int{from, to, weight, 2})
			}
		}
	}
	
	return result
}

func findAllMSTWeights(n int, edges [][]int) []int {
	// This is a simplified version - in practice, finding all MSTs is complex
	// For demonstration, we'll find one MST weight
	weight := findMSTWeight(n, edges)
	if weight != math.MaxInt32 {
		return []int{weight}
	}
	return []int{}
}

func edgeInMST(n int, edges [][]int, from, to, weight int, targetWeight int) bool {
	// Simplified check - in practice, this would be more complex
	testEdges := make([][]int, len(edges))
	copy(testEdges, edges)
	testEdges = append(testEdges, []int{from, to, weight})
	
	return findMSTWeight(n, testEdges) == targetWeight
}

func findEdgeIndex(edges [][]int, target []int) int {
	for i, edge := range edges {
		if edge[0] == target[0] && edge[1] == target[1] && edge[2] == target[2] {
			return i
		}
	}
	return -1
}

// Version with cycle detection
func findCriticalAndPseudoCriticalEdgesWithCycleDetection(n int, edges [][]int) [][]int {
	if n <= 1 {
		return [][]int{}
	}
	
	// Build adjacency list
	adj := make([][][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adj[from] = append(adj[from], []int{to, edge[2]})
		adj[to] = append(adj[to], []int{from, edge[2]})
	}
	
	// Find MST weight
	mstWeight := findMSTWeight(n, edges)
	
	var result [][]int
	
	for i, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		
		// Check if removing this edge disconnects the graph
		if createsCycle(adj, from, to) {
			// Edge creates cycle, check if it's pseudo-critical
			newWeight := findMSTWeightWithoutEdge(n, edges, i)
			if newWeight == mstWeight {
				result = append(result, []int{from, to, weight, 2})
			}
		} else {
			// Edge is a bridge, check if critical
			newWeight := findMSTWeightWithoutEdge(n, edges, i)
			if newWeight > mstWeight {
				result = append(result, []int{from, to, weight, 1})
			}
		}
	}
	
	return result
}

func createsCycle(adj [][]int, from, to int) bool {
	// Simplified cycle detection
	// In practice, this would use DFS or Union-Find
	return false // Simplified for demonstration
}

func main() {
	// Test cases
	fmt.Println("=== Testing Critical and Pseudo-Critical Edges in MST ===")
	
	testCases := []struct {
		n          int
		edges      [][]int
		description string
	}{
		{
			4,
			[][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}},
			"Simple cycle",
		},
		{
			5,
			[][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {3, 4, 1}, {4, 0, 1}},
			"Complex graph",
		},
		{
			3,
			[][]int{{0, 1, 1}, {1, 2, 2}, {0, 2, 3}},
			"Triangle",
		},
		{
			2,
			[][]int{{0, 1, 1}},
			"Single edge",
		},
		{
			4,
			[][]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 3}, {0, 3, 4}},
			"Path graph",
		},
		{
			4,
			[][]int{{0, 1, 1}, {0, 2, 2}, {0, 3, 3}, {1, 2, 4}},
			"Star graph",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  N: %d, Edges: %v\n", tc.n, tc.edges)
		
		result1 := findCriticalAndPseudoCriticalEdges(tc.n, tc.edges)
		result2 := findCriticalAndPseudoCriticalEdgesOptimized(tc.n, tc.edges)
		result3 := findCriticalAndPseudoCriticalEdgesUnionFind(tc.n, tc.edges)
		
		fmt.Printf("  Standard: %v\n", result1)
		fmt.Printf("  Optimized: %v\n", result2)
		fmt.Printf("  Union-Find: %v\n\n", result3)
	}
	
	// Test detailed analysis
	fmt.Println("=== Detailed Analysis Test ===")
	testN, testEdges := 4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}}
	result, analysis := findCriticalAndPseudoCriticalEdgesDetailed(testN, testEdges)
	
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Analysis: %v\n", analysis)
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	// Generate large graph
	largeN := 50
	largeEdges := make([][]int, 0)
	for i := 0; i < largeN; i++ {
		for j := i + 1; j < largeN && j < i+10; j++ {
			weight := (j - i) * 5
			largeEdges = append(largeEdges, []int{i, j, weight})
		}
	}
	
	fmt.Printf("Large test with %d nodes and %d edges\n", largeN, len(largeEdges))
	
	result := findCriticalAndPseudoCriticalEdgesOptimized(largeN, largeEdges)
	fmt.Printf("Result length: %d\n", len(result))
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Single node
	fmt.Printf("Single node: %v\n", findCriticalAndPseudoCriticalEdges(1, [][]int{}))
	
	// No edges
	fmt.Printf("No edges: %v\n", findCriticalAndPseudoCriticalEdges(3, [][]int{}))
	
	// Disconnected graph
	disconnected := [][]int{{0, 1, 1}, {2, 3, 1}}
	fmt.Printf("Disconnected: %v\n", findCriticalAndPseudoCriticalEdges(4, disconnected))
	
	// Multiple edges between same nodes
	multiEdges := [][]int{{0, 1, 1}, {0, 1, 2}, {0, 1, 3}, {1, 2, 1}}
	fmt.Printf("Multiple edges: %v\n", findCriticalAndPseudoCriticalEdges(3, multiEdges))
	
	// Test with different edge weights
	fmt.Println("\n=== Different Edge Weights Test ===")
	
	weightTest := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 10}, {0, 3, 5}}
	result = findCriticalAndPseudoCriticalEdges(4, weightTest)
	fmt.Printf("Mixed weights: %v\n", result)
	
	// Test with uniform weights
	uniform := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}}
	result = findCriticalAndPseudoCriticalEdges(4, uniform)
	fmt.Printf("Uniform weights: %v\n", result)
}
