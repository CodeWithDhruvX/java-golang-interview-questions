package main

import "fmt"

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

// 133. Clone Graph
// Time: O(V+E), Space: O(V)
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	
	visited := make(map[*Node]*Node)
	return cloneNode(node, visited)
}

func cloneNode(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	
	// If node is already cloned, return the clone
	if clone, exists := visited[node]; exists {
		return clone
	}
	
	// Create a clone for the current node
	clone := &Node{Val: node.Val}
	visited[node] = clone
	
	// Clone all neighbors
	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, cloneNode(neighbor, visited))
	}
	
	return clone
}

// Helper function to create a graph from adjacency list
func createGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}
	
	nodes := make([]*Node, len(adjList))
	for i := range nodes {
		nodes[i] = &Node{Val: i + 1}
	}
	
	for i, neighbors := range adjList {
		for _, neighbor := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighbor-1])
		}
	}
	
	return nodes[0]
}

// Helper function to convert graph to adjacency list for verification
func graphToAdjList(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}
	
	visited := make(map[*Node]bool)
	adjList := make(map[int][]int)
	
	var dfs func(*Node)
	dfs = func(n *Node) {
		if n == nil || visited[n] {
			return
		}
		
		visited[n] = true
		for _, neighbor := range n.Neighbors {
			adjList[n.Val] = append(adjList[n.Val], neighbor.Val)
			dfs(neighbor)
		}
	}
	
	dfs(node)
	
	// Convert to slice format
	result := make([][]int, len(visited))
	for val, neighbors := range adjList {
		result[val-1] = neighbors
	}
	
	return result
}

func main() {
	// Test cases
	testCases := [][][]int{
		{{2, 4}, {1, 3}, {2, 4}, {1, 3}}, // 4-node graph
		{{}}, // Single node
		{}, // Empty graph
		{{2}, {1}}, // 2-node graph
		{{2, 3}, {1, 3}, {1, 2}}, // 3-node complete graph
		{{2}, {3}, {4}, {5}, {1}}, // 5-node cycle
	}
	
	for i, adjList := range testCases {
		node := createGraph(adjList)
		cloned := cloneGraph(node)
		clonedAdjList := graphToAdjList(cloned)
		fmt.Printf("Test Case %d: %v -> Cloned: %v\n", i+1, adjList, clonedAdjList)
	}
}
