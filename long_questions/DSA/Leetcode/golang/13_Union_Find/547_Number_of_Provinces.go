package main

import "fmt"

// 547. Number of Provinces
// Time: O(N²), Space: O(N)
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}
	
	// Initialize Union-Find
	parent := make([]int, n)
	rank := make([]int, n)
	
	for i := 0; i < n; i++ {
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
	
	var union func(int, int)
	union = func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		
		if rootX == rootY {
			return
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
	}
	
	// Process all connections
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}
	
	// Count unique provinces
	provinces := make(map[int]bool)
	for i := 0; i < n; i++ {
		provinces[find(i)] = true
	}
	
	return len(provinces)
}

// DFS approach
func findCircleNumDFS(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}
	
	visited := make([]bool, n)
	provinces := 0
	
	var dfs func(int)
	dfs = func(city int) {
		visited[city] = true
		for neighbor := 0; neighbor < n; neighbor++ {
			if isConnected[city][neighbor] == 1 && !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}
	
	for i := 0; i < n; i++ {
		if !visited[i] {
			provinces++
			dfs(i)
		}
	}
	
	return provinces
}

// BFS approach
func findCircleNumBFS(isConnected [][]int) int {
	n := len(isConnected)
	if n == 0 {
		return 0
	}
	
	visited := make([]bool, n)
	provinces := 0
	
	for i := 0; i < n; i++ {
		if !visited[i] {
			provinces++
			
			// BFS from city i
			queue := []int{i}
			visited[i] = true
			
			for len(queue) > 0 {
				city := queue[0]
				queue = queue[1:]
				
				for neighbor := 0; neighbor < n; neighbor++ {
					if isConnected[city][neighbor] == 1 && !visited[neighbor] {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}
		}
	}
	
	return provinces
}

func main() {
	// Test cases
	testCases := [][][]int{
		{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
		{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 0, 0, 1}},
		{{1}},
		{{}},
		{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}},
		{{1, 1, 0, 0}, {1, 1, 0, 0}, {0, 0, 1, 1}},
	}
	
	for i, matrix := range testCases {
		result1 := findCircleNum(matrix)
		result2 := findCircleNumDFS(matrix)
		result3 := findCircleNumBFS(matrix)
		
		fmt.Printf("Test Case %d: %v\n", i+1, matrix)
		fmt.Printf("  Union-Find: %d, DFS: %d, BFS: %d\n\n", result1, result2, result3)
	}
}
