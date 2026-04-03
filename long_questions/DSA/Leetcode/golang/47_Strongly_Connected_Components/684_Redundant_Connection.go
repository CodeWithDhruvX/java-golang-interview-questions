package main

import (
	"fmt"
)

// 684. Redundant Connection - Strongly Connected Components
// Time: O(N^2), Space: O(N)
func findRedundantConnection(edges [][]int) []int {
	if len(edges) == 0 {
		return []int{}
	}
	
	// Build adjacency list
	adj := make([][]int, len(edges)+1)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adj[from] = append(adj[from], to)
		adj[to] = append(adj[to], from)
	}
	
	// Find redundant connection using Union-Find
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}
	
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		
		// Find roots
		rootFrom := find(parent, from)
		rootTo := find(parent, to)
		
		// If already connected, this edge is redundant
		if rootFrom == rootTo {
			return edge
		}
		
		// Union the sets
		parent[rootFrom] = rootTo
	}
	
	return []int{}
}

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

// Strongly Connected Components with Kosaraju's algorithm
func findRedundantConnectionSCC(edges [][]int) []int {
	if len(edges) == 0 {
		return []int{}
	}
	
	// Build adjacency list
	adj := make([][]int, len(edges)+1)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adj[from] = append(adj[from], to)
		adj[to] = append(adj[to], from)
	}
	
	// Kosaraju's algorithm to find SCCs
	visited := make([]bool, len(edges)+1)
	sccs := kosaraju(adj, visited)
	
	// If graph is already strongly connected, no redundant edges
	if len(sccs) == 1 {
		return []int{}
	}
	
	// Find first redundant edge
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		
		// Check if removing this edge makes graph not strongly connected
		if !isStronglyConnectedAfterRemove(adj, from, to) {
			return edge
		}
	}
	
	return []int{}
}

func kosaraju(adj [][]int, visited []bool) [][]int {
	n := len(adj)
	order := []int{}
	
	// First pass: fill stack with vertices in decreasing order of finishing time
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs1(i, adj, visited, &order)
		}
	}
	
	// Second pass: find SCCs
	visited = make([]bool, n)
	sccs := [][]int{}
	
	for i := len(order) - 1; i >= 0; i-- {
		v := order[i]
		if !visited[v] {
			scc := []int{}
			dfs2(v, adj, visited, &scc)
			sccs = append(sccs, scc)
		}
	}
	
	return sccs
}

func dfs1(v int, adj [][]int, visited []bool, order *[]int) {
	visited[v] = true
	
	for _, neighbor := range adj[v] {
		if !visited[neighbor] {
			dfs1(neighbor, adj, visited, order)
		}
	}
	
	*order = append(*order, v)
}

func dfs2(v int, adj [][]int, visited []bool, scc *[]int) {
	visited[v] = true
	*scc = append(*scc, v)
	
	for _, neighbor := range adj[v] {
		if !visited[neighbor] {
			dfs2(neighbor, adj, visited, scc)
		}
	}
}

func isStronglyConnectedAfterRemove(adj [][]int, from, to int) bool {
	n := len(adj)
	
	// Remove edge temporarily
	originalFrom := make([]int, 0)
	originalTo := make([]int, 0)
	
	for i, neighbors := range adj {
		for _, neighbor := range neighbors {
			if i == from && neighbor == to {
				continue
			}
			if i == from {
				originalFrom = append(originalFrom, neighbor)
			} else if i == to {
				originalTo = append(originalTo, neighbor)
			}
		}
	}
	
	adj[from] = originalFrom
	adj[to] = originalTo
	
	// Check if graph is still strongly connected
	visited := make([]bool, n)
	order := []int{}
	
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs1(i, adj, visited, &order)
		}
	}
	
	// Second pass
	visited = make([]bool, n)
	sccs := kosaraju(adj, visited)
	
	// Restore edge
	adj[from] = append(adj[from], to)
	adj[to] = append(adj[to], from)
	
	return len(sccs) == 1
}

// Strongly Connected Components with Tarjan's algorithm
func findRedundantConnectionTarjan(edges [][]int) []int {
	if len(edges) == 0 {
		return []int{}
	}
	
	// Build adjacency list
	adj := make([][]int, len(edges)+1)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adj[from] = append(adj[from], to)
		adj[to] = append(adj[to], from)
	}
	
	// Tarjan's algorithm to find SCCs
	index := 0
	stack := []int{}
	onStack := make([]bool, len(edges)+1)
	indices := make([]int, len(edges)+1)
	lowLink := make([]int, len(edges)+1)
	sccs := [][]int{}
	
	for i := 0; i < len(edges)+1; i++ {
		strongconnect(i, adj, &index, &stack, onStack, &indices, &lowLink, &sccs)
	}
	
	// If graph is already strongly connected, no redundant edges
	if len(sccs) == 1 {
		return []int{}
	}
	
	// Find first redundant edge
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		
		if !isStronglyConnectedAfterRemove(adj, from, to) {
			return edge
		}
	}
	
	return []int{}
}

func strongconnect(v int, adj [][]int, index *int, stack *[]int, onStack *[]bool, indices []int, lowLink []int, sccs *[][]int) {
	indices[v] = *index
	lowLink[v] = *index
	*index++
	
	*stack = append(*stack, v)
	onStack[v] = true
	
	for _, neighbor := range adj[v] {
		if indices[neighbor] == -1 {
			strongconnect(neighbor, adj, index, stack, onStack, indices, lowLink, sccs)
		}
	}
	
	// Check if v is a root of an SCC
	w := *stack
	for len(*stack) > 0 && indices[w] != lowLink[v] {
		w = (*stack)[len(*stack)-1]
	}
	
	if indices[w] == lowLink[v] {
		// Start new SCC
		scc := []int{}
		for len(*stack) > 0 && indices[(*stack)[len(*stack)-1]] != lowLink[v] {
			w = (*stack)[len(*stack)-1]
			scc = append(scc, w)
		}
		*sccs = append(*sccs, scc)
	} else {
		// Pop from stack
		*stack = (*stack)[:len(*stack)-1]
	}
}

// Strongly Connected Components with Union-Find
func findRedundantConnectionUnionFind(edges [][]int) []int {
	if len(edges) == 0 {
		return []int{}
	}
	
	// Build adjacency list
	adj := make([][]int, len(edges)+1)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adj[from] = append(adj[from], to)
		adj[to] = append(adj[to], from)
	}
	
	// Use Union-Find to detect cycles
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}
	
	// Process edges in reverse order to detect cycles
	for i := len(edges) - 1; i >= 0; i-- {
		edge := edges[i]
		from, to := edge[0], edge[1]
		
		// If adding this edge creates a cycle, it's redundant
		if find(parent, from) == find(parent, to) {
			return edge
		}
		
		// Union the sets
		rootFrom := find(parent, from)
		rootTo := find(parent, to)
		parent[rootFrom] = rootTo
	}
	
	return []int{}
}

// Strongly Connected Components with edge counting
func findRedundantConnectionWithCount(edges [][]int) ([]int, int) {
	if len(edges) == 0 {
		return []int{}, 0
	}
	
	// Build adjacency list
	adj := make([][]int, len(edges)+1)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adj[from] = append(adj[from], to)
		adj[to] = append(adj[to], from)
	}
	
	// Find SCCs
	visited := make([]bool, len(edges)+1)
	sccs := kosaraju(adj, visited)
	
	// Count edges within each SCC
	edgeCount := 0
	for _, scc := range sccs {
		edgeCount += countEdgesInSCC(scc, adj)
	}
	
	// If all edges are within SCCs, no redundant edges
	if edgeCount == len(edges) {
		return []int{}, 0
	}
	
	// Find first redundant edge
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		
		// Check if this edge connects different SCCs
		fromSCC := findSCCForVertex(from, sccs)
		toSCC := findSCCForVertex(to, sccs)
		
		if fromSCC != toSCC {
			return edge
		}
	}
	
	return []int{}, 0
}

func countEdgesInSCC(scc []int, adj [][]int) int {
	count := 0
	visited := make(map[int]bool)
	
	for _, v := range scc {
		visited[v] = true
		for _, neighbor := range adj[v] {
			if visited[neighbor] {
				count++
			}
		}
	}
	
	return count
}

func findSCCForVertex(v int, sccs [][]int) int {
	for i, scc := range sccs {
		for _, vertex := range scc {
			if vertex == v {
				return i
			}
		}
	}
	return -1
}

// Strongly Connected Components with multiple test cases
func findRedundantConnectionMultipleTests(edges [][]int) [][]int {
	if len(edges) == 0 {
		return [][]int{}
	}
	
	// Test different approaches
	result1 := findRedundantConnection(edges)
	result2 := findRedundantConnectionSCC(edges)
	result3 := findRedundantConnectionTarjan(edges)
	result4 := findRedundantConnectionUnionFind(edges)
	result5, count := findRedundantConnectionWithCount(edges)
	
	return [][]int{result1, result2, result3, result4, result5}
}

func main() {
	// Test cases
	fmt.Println("=== Testing Redundant Connection - Strongly Connected Components ===")
	
	testCases := []struct {
		edges      [][]int
		description string
	}{
		{
			[][]int{{1, 2}, {1, 3}, {2, 3}},
			"Triangle with redundant edge",
		},
		{
			[][]int{{1, 2}, {2, 3}},
			"Triangle without redundant edge",
		},
		{
			[][]int{{1, 2}, {1, 3}, {2, 3}, {3, 1}},
			"Triangle with cycle",
		},
		{
			[][]int{{1, 2}, {2, 4}, {3, 1}},
			"Simple graph",
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}},
			"Square",
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 1}},
			"Pentagon",
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 1}},
			"Hexagon",
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {7, 1}},
			"Heptagon",
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {7, 8}, {8, 1}},
			"Octagon",
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}, {9, 1}},
			"Nonagon",
		},
		{
			[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}, {9, 10}, {10, 1}},
			"Decagon",
		},
		{
			[][]int{},
			"Empty graph",
		},
		{
			[][]int{{1, 2}},
			"Single edge",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Edges: %v\n", tc.edges)
		
		result1 := findRedundantConnection(tc.edges)
		result2 := findRedundantConnectionSCC(tc.edges)
		result3 := findRedundantConnectionTarjan(tc.edges)
		result4 := findRedundantConnectionUnionFind(tc.edges)
		result5, count := findRedundantConnectionWithCount(tc.edges)
		
		fmt.Printf("  Union-Find: %v\n", result1)
		fmt.Printf("  Kosaraju: %v\n", result2)
		fmt.Printf("  Tarjan: %v\n", result3)
		fmt.Printf("  Union-Find: %v\n", result4)
		fmt.Printf("  With count: %v, count: %d\n\n", result5, count)
		
		fmt.Println()
	}
	
	// Test multiple approaches
	fmt.Println("=== Multiple Approaches Test ===")
	complexEdges := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 1}}
	multipleResults := findRedundantConnectionMultipleTests(complexEdges)
	
	for i, result := range multipleResults {
		fmt.Printf("Approach %d: %v\n", i+1, result)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	// Generate large graph
	largeEdges := make([][]int, 0)
	for i := 1; i <= 1000; i++ {
		for j := i + 1; j <= 1000 && j <= i+10; j++ {
			largeEdges = append(largeEdges, []int{i, j})
		}
	}
	
	fmt.Printf("Large test with %d edges\n", len(largeEdges))
	
	start := time.Now()
	result := findRedundantConnection(largeEdges)
	duration := time.Since(start)
	
	fmt.Printf("Large graph result: %v\n", result)
	fmt.Printf("Time taken: %v\n", duration)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Single vertex
	fmt.Printf("Single vertex: %v\n", findRedundantConnection([][]int{}))
	
	// Self loop
	selfLoop := [][]int{{1, 1}}
	fmt.Printf("Self loop: %v\n", findRedundantConnection(selfLoop))
	
	// Multiple edges between same vertices
	multiEdges := [][]int{{1, 2}, {1, 2}, {1, 2}}
	fmt.Printf("Multiple edges: %v\n", findRedundantConnection(multiEdges))
	
	// Disconnected graph
	disconnected := [][]int{{1, 2}, {3, 4}}
	fmt.Printf("Disconnected: %v\n", findRedundantConnection(disconnected))
	
	// Complete graph
	complete := [][]int{{1, 2}, {1, 3}, {2, 3}}
	fmt.Printf("Complete graph: %v\n", findRedundantConnection(complete))
	
	// Test with different SCC sizes
	fmt.Println("\n=== SCC Size Analysis ===")
	
	// Create graph with two SCCs connected by one edge
	twoSCCEdges := [][]int{{1, 2}, {2, 3}, {3, 1}, {4, 5}, {5, 4}}
	twoSCCResult := findRedundantConnectionSCC(twoSCCEdges)
	fmt.Printf("Two SCCs connected by one edge: %v\n", twoSCCResult)
	
	// Test Tarjan's algorithm specifically
	fmt.Println("\n=== Tarjan Algorithm Test ===")
	tarjanResult := findRedundantConnectionTarjan([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}})
	fmt.Printf("Tarjan result: %v\n", tarjanResult)
	
	// Test edge counting approach
	fmt.Println("\n=== Edge Counting Test ===")
	edgeCountResult, count := findRedundantConnectionWithCount([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 1}})
	fmt.Printf("Edge counting result: %v, count: %d\n", edgeCountResult, count)
}
