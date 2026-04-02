package main

import "fmt"

// 684. Redundant Connection
// Time: O(N α(N)), Space: O(N) where α is inverse Ackermann function
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	
	// Initialize Union-Find
	parent := make([]int, n+1)
	rank := make([]int, n+1)
	
	for i := 1; i <= n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	
	// Union-Find helper functions
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // Path compression
		}
		return parent[x]
	}
	
	var union func(int, int) bool
	union = func(x, y int) bool {
		rootX := find(x)
		rootY := find(y)
		
		if rootX == rootY {
			return false // Already connected, this is the redundant edge
		}
		
		// Union by rank
		if rank[rootX] < rank[rootY] {
			parent[rootX] = rootY
		} else if rank[rootX] > rank[rootY] {
			parent[rootY] = rootX
		} else {
			parent[rootY] = rootX
			rank[rootX]++
		}
		
		return true
	}
	
	// Process edges
	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge // Found redundant connection
		}
	}
	
	return []int{}
}

// Alternative approach without rank
func findRedundantConnectionSimple(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	
	for i := 1; i <= n; i++ {
		parent[i] = i
	}
	
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	
	for _, edge := range edges {
		rootX := find(edge[0])
		rootY := find(edge[1])
		
		if rootX == rootY {
			return edge
		}
		
		parent[rootY] = rootX
	}
	
	return []int{}
}

// Function to detect cycles in undirected graph using DFS
func findRedundantConnectionDFS(edges [][]int) []int {
	n := len(edges)
	adj := make(map[int][]int)
	
	// Build adjacency list
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	
	visited := make(map[int]bool)
	
	var dfs func(int, int) bool
	dfs = func(node, parent int) bool {
		visited[node] = true
		
		for _, neighbor := range adj[node] {
			if neighbor == parent {
				continue
			}
			
			if visited[neighbor] {
				return true // Cycle detected
			}
			
			if dfs(neighbor, node) {
				return true
			}
		}
		
		return false
	}
	
	// Check each edge by temporarily removing it
	for _, edge := range edges {
		// Temporarily remove edge
		adj[edge[0]] = removeNode(adj[edge[0]], edge[1])
		adj[edge[1]] = removeNode(adj[edge[1]], edge[0])
		
		// Clear visited and check for cycles
		visited = make(map[int]bool)
		hasCycle := dfs(1, -1)
		
		// Restore edge
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
		
		if !hasCycle {
			return edge
		}
	}
	
	return []int{}
}

func removeNode(slice []int, node int) []int {
	for i, val := range slice {
		if val == node {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func main() {
	// Test cases
	testCases := [][][]int{
		{{1, 2}, {1, 3}, {2, 3}},
		{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
		{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}},
		{{1, 2}},
		{{1, 2}, {1, 3}},
		{{2, 1}, {3, 2}, {4, 2}, {1, 4}},
		{{1, 2}, {2, 3}, {3, 1}, {1, 4}},
		{{1, 2}, {2, 3}, {4, 1}, {5, 2}, {3, 5}},
	}
	
	for i, edges := range testCases {
		result1 := findRedundantConnection(edges)
		result2 := findRedundantConnectionSimple(edges)
		result3 := findRedundantConnectionDFS(edges)
		
		fmt.Printf("Test Case %d: %v\n", i+1, edges)
		fmt.Printf("  Union-Find: %v\n", result1)
		fmt.Printf("  Simple UF: %v\n", result2)
		fmt.Printf("  DFS: %v\n\n", result3)
	}
}
