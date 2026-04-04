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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Strongly Connected Components for Redundant Edge Detection
- **Union-Find**: Detect cycles by tracking connected components
- **Kosaraju's Algorithm**: Find SCCs using two DFS passes
- **Tarjan's Algorithm**: Single-pass SCC detection with stack
- **Cycle Detection**: Redundant edges create cycles in tree structure

## 2. PROBLEM CHARACTERISTICS
- **Redundant Connection**: Find edge that creates a cycle in tree
- **Tree Structure**: N nodes with N-1 edges, adding one creates cycle
- **Strong Connectivity**: All nodes reachable from each other
- **Edge Removal**: Find edge whose removal maintains connectivity

## 3. SIMILAR PROBLEMS
- Find Critical and Pseudo-Critical Edges (LeetCode 1489) - MST edge analysis
- Network Delay Time (LeetCode 743) - Graph connectivity
- Evaluate Division (LeetCode 399) - Graph relationships
- Minimum Spanning Tree problems - Cycle detection

## 4. KEY OBSERVATIONS
- **Tree Property**: Tree with N nodes has exactly N-1 edges
- **Cycle Creation**: Adding N-th edge creates exactly one cycle
- **Union-Find Natural**: Perfect for detecting cycles during construction
- **SCC Analysis**: Strongly connected components reveal redundant edges

## 5. VARIATIONS & EXTENSIONS
- **Union-Find**: Simple cycle detection during edge processing
- **Kosaraju's**: Two-pass DFS for SCC identification
- **Tarjan's**: Single-pass SCC with low-link values
- **Edge Analysis**: Count edges within and between SCCs

## 6. INTERVIEW INSIGHTS
- Always clarify: "Tree structure? Multiple cycles? Directed vs undirected?"
- Edge cases: single edge, self-loops, multiple edges between nodes
- Time complexity: O(N α(N)) for Union-Find, O(N + E) for DFS
- Space complexity: O(N) for Union-Find, O(N + E) for adjacency
- Key insight: Union-Find is optimal for cycle detection in trees

## 7. COMMON MISTAKES
- Wrong Union-Find initialization (should be 1 to N, not 0 to N-1)
- Missing path compression optimization
- Incorrect cycle detection logic
- Not handling disconnected graphs properly
- Wrong edge removal analysis

## 8. OPTIMIZATION STRATEGIES
- **Union-Find**: O(N α(N)) time, O(N) space - optimal
- **Kosaraju's**: O(N + E) time, O(N + E) space - for SCC analysis
- **Tarjan's**: O(N + E) time, O(N + E) space - single pass
- **Edge Counting**: O(N + E) time, O(N + E) space - for analysis

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like building a road network between cities:**
- You have cities connected by roads forming a tree structure
- Adding one extra road creates a cycle (alternative route between cities)
- You want to find which extra road creates the cycle
- Like a transportation planner identifying redundant roads in a network
- Each redundant road provides an alternative path between cities

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Tree edges plus one extra edge (N edges for N nodes)
2. **Goal**: Find the redundant edge that creates a cycle
3. **Constraints**: Graph becomes tree after removing redundant edge
4. **Output**: Edge that creates the cycle

#### Phase 2: Key Insight Recognition
- **"Tree property"** → N nodes, N-1 edges for tree, N edges = one cycle
- **"Union-Find natural"** → Perfect for detecting cycles during construction
- **"SCC analysis"** → Strongly connected components reveal cycles
- **"Edge removal"** → Need to find edge whose removal maintains connectivity

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the edge that creates a cycle.
Brute force: try removing each edge and check connectivity O(N × (N+E)).

Union-Find Approach:
1. Initialize each node as separate component
2. Process edges in order:
   - If edge connects nodes in same component → cycle found
   - Else: union the components
3. First edge that creates cycle is redundant

This gives O(N α(N)) time, O(N) space!"
```

#### Phase 4: Edge Case Handling
- **Single edge**: No cycle possible, return empty
- **Self-loop**: Always redundant (node connects to itself)
- **Multiple edges**: First redundant edge in order
- **Disconnected graph**: Handle appropriately

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: edges = [[1,2], [1,3], [2,3]]

Human thinking:
"Union-Find Process:
Step 1: Initialize
parent = [1,2,3] (each node is its own parent)

Step 2: Process edge [1,2]
root1 = find(1) = 1
root2 = find(2) = 2
Different roots → no cycle
Union: parent[1] = 2
parent = [2,2,3]

Step 3: Process edge [1,3]
root1 = find(1) = find(2) = 2
root2 = find(3) = 3
Different roots → no cycle
Union: parent[2] = 3
parent = [2,3,3]

Step 4: Process edge [2,3]
root1 = find(2) = find(3) = 3
root2 = find(3) = 3
Same roots → CYCLE FOUND!
Result: [2,3] ✓"
```

#### Phase 6: Intuition Validation
- **Why Union-Find**: Efficiently tracks connected components during construction
- **Why first cycle**: Problem asks for redundant edge in given order
- **Why SCC analysis**: Alternative approach using graph theory
- **Why O(N α(N))**: Near-linear time with path compression

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use DFS?"** → Union-Find is more efficient for this specific problem
2. **"Should I check all edges?"** → No, first edge that creates cycle is answer
3. **"What about directed graphs?"** → Problem uses undirected edges
4. **"Can I use adjacency matrix?"** → Union-Find is more memory efficient
5. **"Why path compression?"** → Essential for near-linear performance

### Real-World Analogy
**Like identifying redundant connections in a computer network:**
- You have computers connected by network cables forming a tree
- Adding an extra cable creates a redundant connection (loop)
- You want to find which cable is redundant
- Like a network administrator finding extra cables that create loops
- Each redundant cable provides an alternative data path

### Human-Readable Pseudocode
```
function findRedundantConnection(edges):
    parent = array of size N+1
    
    # Initialize each node as its own parent
    for i from 1 to N:
        parent[i] = i
    
    # Process each edge
    for each edge [u, v]:
        rootU = find(parent, u)
        rootV = find(parent, v)
        
        # If same component, this edge creates cycle
        if rootU == rootV:
            return edge
        
        # Union the components
        parent[rootU] = rootV
    
    return []

function find(parent, x):
    if parent[x] != x:
        parent[x] = find(parent, parent[x])  # Path compression
    return parent[x]
```

### Execution Visualization

### Example: edges = [[1,2], [2,3], [3,4], [4,1]]
```
Initial: parent = [1,2,3,4]

Process [1,2]:
find(1) = 1, find(2) = 2 → different
parent[1] = 2
parent = [2,2,3,4]

Process [2,3]:
find(2) = 2, find(3) = 3 → different
parent[2] = 3
parent = [2,3,3,4]

Process [3,4]:
find(3) = 3, find(4) = 4 → different
parent[3] = 4
parent = [2,3,4,4]

Process [4,1]:
find(4) = 4, find(1) = find(2) = find(3) = find(4) = 4 → SAME!
CYCLE FOUND: [4,1] ✓
```

### Key Visualization Points:
- **Component Tracking**: Each node tracks its root parent
- **Path Compression**: Find operation flattens tree structure
- **Cycle Detection**: Same root indicates existing connection
- **Union Operation**: Merges different components

### Union-Find Process Visualization:
```
Initialize: parent[i] = i for all i

For each edge [u, v]:
    rootU = find(parent, u)
    rootV = find(parent, v)
    
    if rootU == rootV:
        return [u, v]  # Cycle found
    else:
        parent[rootU] = rootV  # Union
```

### Time Complexity Breakdown:
- **Union-Find**: O(N α(N)) time, O(N) space - optimal
- **Kosaraju's**: O(N + E) time, O(N + E) space - for SCC analysis
- **Tarjan's**: O(N + E) time, O(N + E) space - single pass
- **Edge Counting**: O(N + E) time, O(N + E) space - for analysis

### Alternative Approaches:

#### 1. Kosaraju's Algorithm (O(N + E) time, O(N + E) space)
```go
func findRedundantConnectionKosaraju(edges [][]int) []int {
    // Build adjacency list
    adj := buildAdjacencyList(edges)
    
    // First DFS pass to get finishing order
    visited := make([]bool, len(edges)+1)
    order := []int{}
    
    for i := 1; i <= len(edges); i++ {
        if !visited[i] {
            dfs1(i, adj, visited, &order)
        }
    }
    
    // Second DFS pass on reversed graph to find SCCs
    visited = make([]bool, len(edges)+1)
    sccs := [][]int{}
    
    for i := len(order) - 1; i >= 0; i-- {
        v := order[i]
        if !visited[v] {
            scc := []int{}
            dfs2(v, reversedAdj, visited, &scc)
            sccs = append(sccs, scc)
        }
    }
    
    // Find edge connecting nodes in same SCC
    for _, edge := range edges {
        if nodesInSameSCC(edge[0], edge[1], sccs) {
            return edge
        }
    }
    
    return []int{}
}
```
- **Pros**: Provides full SCC analysis
- **Cons**: More complex than Union-Find for this specific problem

#### 2. Tarjan's Algorithm (O(N + E) time, O(N + E) space)
```go
func findRedundantConnectionTarjan(edges [][]int) []int {
    adj := buildAdjacencyList(edges)
    
    index := 0
    stack := []int{}
    onStack := make([]bool, len(edges)+1)
    indices := make([]int, len(edges)+1)
    lowLink := make([]int, len(edges)+1)
    
    for i := 1; i <= len(edges); i++ {
        if indices[i] == -1 {
            strongconnect(i, adj, &index, &stack, onStack, indices, lowLink)
        }
    }
    
    // Find redundant edge from SCC analysis
    return findRedundantFromSCCs(sccs, edges)
}
```
- **Pros**: Single pass, efficient for large graphs
- **Cons**: Complex implementation with low-link values

#### 3. DFS Cycle Detection (O(N + E) time, O(N + E) space)
```go
func findRedundantConnectionDFS(edges [][]int) []int {
    adj := buildAdjacencyList(edges)
    
    for _, edge := range edges {
        // Temporarily remove edge
        from, to := edge[0], edge[1]
        removeEdge(adj, from, to)
        
        // Check if graph is still connected
        if !isConnected(adj) {
            return edge
        }
        
        // Restore edge
        addEdge(adj, from, to)
    }
    
    return []int{}
}
```
- **Pros**: Simple concept
- **Cons**: O(E × (N + E)) time, inefficient

### Extensions for Interviews:
- **Multiple Redundant Edges**: Find all edges that create cycles
- **Directed Graphs**: Handle strongly connected components in directed graphs
- **Edge Weight**: Find redundant edge with minimum/maximum weight
- **Dynamic Updates**: Add/remove edges incrementally
- **Real-world Applications**: Network design, circuit analysis, dependency resolution
*/
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
