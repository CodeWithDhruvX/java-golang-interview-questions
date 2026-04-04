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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: DP on Graphs for Shortest Paths
- **All-Pairs Shortest Path**: Floyd-Warshall algorithm
- **Single-Source Shortest Path**: Dijkstra and Bellman-Ford algorithms
- **Dynamic Programming**: Build optimal solutions incrementally
- **Path Reconstruction**: Track actual shortest paths

## 2. PROBLEM CHARACTERISTICS
- **Weighted Directed Graph**: Edges have travel times
- **Network Delay**: Time for signal to reach all nodes
- **Source Node**: Signal starts from node k
- **Connectivity**: All nodes must be reachable

## 3. SIMILAR PROBLEMS
- Cheapest Flights Within K Stops (LeetCode 787) - Constrained shortest path
- Network Delay Time (LeetCode 743) - Same problem
- Minimum Cost to Connect All Points (LeetCode 1584) - MST problem
- Find the City With the Smallest Number of Neighbors (LeetCode 1334) - Graph analysis

## 4. KEY OBSERVATIONS
- **Shortest Path Natural**: Need shortest paths from source to all nodes
- **All-Pairs vs Single-Source**: Floyd-Warshall vs Dijkstra/Bellman-Ford
- **Negative Weights**: Bellman-Ford handles negative weights
- **Early Termination**: Can stop when all nodes processed

## 5. VARIATIONS & EXTENSIONS
- **Multiple Sources**: Find optimal source among multiple candidates
- **Path Reconstruction**: Track actual shortest paths
- **Early Termination**: Optimize when all nodes visited
- **Constraint Handling**: Add various constraints to paths

## 6. INTERVIEW INSIGHTS
- Always clarify: "Graph size constraints? Edge weights? Negative edges?"
- Edge cases: disconnected graph, single node, no edges
- Time complexity: O(N³) for Floyd-Warshall, O(N²) for Dijkstra
- Space complexity: O(N²) for adjacency matrix, O(N+E) for adjacency list
- Key insight: choose algorithm based on graph characteristics

## 7. COMMON MISTAKES
- Wrong algorithm choice for graph type
- Not handling disconnected graphs properly
- Incorrect initialization of distance matrix
- Missing early termination optimization
- Not handling overflow in distance calculations

## 8. OPTIMIZATION STRATEGIES
- **Floyd-Warshall**: O(N³) time, O(N²) space - all pairs
- **Dijkstra**: O(N²) time, O(N+E) space - single source, positive weights
- **Bellman-Ford**: O(N*E) time, O(N) space - single source, negative weights
- **Early Termination**: O(N²) time, O(N+E) space - optimized Dijkstra

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like planning delivery routes in a city:**
- Each location is a node in the city
- Roads between locations have travel times (edge weights)
- You start deliveries from a central warehouse (source node)
- You need to know the longest delivery time to any location
- You want to minimize the maximum delivery time
- Like logistics planning for package delivery

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Weighted directed graph (times), number of nodes (n), source (k)
2. **Goal**: Find time for signal to reach all nodes
3. **Rules**: Signal travels along edges with given weights
4. **Output**: Maximum shortest path distance from source to any node

#### Phase 2: Key Insight Recognition
- **"Shortest path natural"** → Need shortest paths from source to all nodes
- **"Algorithm choice"** → Floyd-Warshall for all pairs, Dijkstra for single source
- **"Network delay"** → Maximum of shortest path distances
- **"Connectivity check"** → All nodes must be reachable

#### Phase 3: Strategy Development
```
Human thought process:
"I need network delay time from source k.
This is essentially finding shortest paths from k to all nodes.

Algorithm Options:
1. Floyd-Warshall: All pairs shortest paths, O(N³)
2. Dijkstra: Single source, O(N²) for positive weights
3. Bellman-Ford: Single source, handles negative weights

For this problem, Dijkstra is usually best:
- Single source (node k)
- Positive edge weights (time delays)
- Need shortest paths to all nodes
- Take maximum of all shortest paths

This gives O(N²) time!"
```

#### Phase 4: Edge Case Handling
- **Disconnected graph**: Return -1 if any node unreachable
- **Single node**: Return 0 (no travel needed)
- **No edges**: Return -1 unless n=1
- **Negative weights**: Use Bellman-Ford instead of Dijkstra

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: times=[[2,1,1],[2,3,1],[3,4,1]], n=4, k=2

Human thinking:
"Dijkstra Approach from source 2:
Initialize: dist[2]=0, others=∞

Step 1: Visit node 2 (distance 0)
- Update neighbors: dist[1]=1, dist[3]=1

Step 2: Visit node 1 (distance 1) 
- No outgoing edges from 1

Step 3: Visit node 3 (distance 1)
- Update neighbor: dist[4]=2

Step 4: Visit node 4 (distance 2)
- No outgoing edges from 4

All nodes visited, distances: [1,0,1,2]
Maximum distance: 2

Result: 2 ✓"
```

#### Phase 6: Intuition Validation
- **Why Dijkstra works**: Positive edge weights, single source
- **Why maximum delay**: Need time for farthest node
- **Why O(N²)**: Each step processes one node, updates all edges
- **Why early termination**: Can stop when all nodes visited

### Common Human Pitfalls & How to Avoid Them
1. **"Why not BFS?"** → Weighted edges need Dijkstra, not BFS
2. **"Should I use Floyd-Warshall?"** → Only if need all pairs
3. **"What about negative weights?"** → Use Bellman-Ford instead
4. **"Can I optimize further?"** → Use priority queue for O(E log V)
5. **"What about disconnected graphs?"** → Return -1 if unreachable

### Real-World Analogy
**Like planning emergency response times:**
- Each location is a neighborhood in a city
- Roads between neighborhoods have travel times
- Emergency services start from a central station
- You need to know the longest response time to any neighborhood
- You want to ensure all neighborhoods are reachable within reasonable time
- Like urban planning for emergency services coverage

### Human-Readable Pseudocode
```
function networkDelayTime(times, n, k):
    # Build adjacency list
    adj = adjacency list of size n+1
    for each (from, to, weight) in times:
        adj[from].append((to, weight))
    
    # Dijkstra's algorithm
    dist = array of size n+1, initialized to infinity
    dist[k] = 0
    visited = array of size n+1, initialized to false
    
    for i from 1 to n:
        # Find unvisited node with minimum distance
        u = node with minimum dist[u] where !visited[u]
        if u == -1: break
        
        visited[u] = true
        
        # Update distances to neighbors
        for each (to, weight) in adj[u]:
            if dist[u] + weight < dist[to]:
                dist[to] = dist[u] + weight
    
    # Find maximum distance
    maxDist = 0
    for i from 1 to n:
        if dist[i] == infinity:
            return -1  # disconnected
        maxDist = max(maxDist, dist[i])
    
    return maxDist
```

### Execution Visualization

### Example: times=[[2,1,1],[2,3,1],[3,4,1]], n=4, k=2
```
Dijkstra Process:

Initial: dist=[∞,0,∞,∞], visited=[false,false,false,false]

Step 1: Visit node 2 (dist=0)
- Update: dist[1]=1, dist[3]=1
- visited[2]=true

Step 2: Visit node 1 (dist=1)
- No updates (no outgoing edges)
- visited[1]=true

Step 3: Visit node 3 (dist=1)
- Update: dist[4]=2
- visited[3]=true

Step 4: Visit node 4 (dist=2)
- No updates (no outgoing edges)
- visited[4]=true

Final distances: dist=[∞,1,0,1,2]
Maximum distance: 2

Result: 2 ✓
```

### Key Visualization Points:
- **Priority Processing**: Always process closest unvisited node
- **Distance Updates**: Relax edges from current node
- **Early Termination**: Stop when all nodes visited
- **Maximum Delay**: Take maximum of all shortest paths

### Graph Visualization:
```
Graph Structure:
    1 ← 2 → 3 → 4
        1   1   1

Source: Node 2
Shortest paths:
2 → 1: distance 1
2 → 3: distance 1  
2 → 4: distance 2 (via 3)

Network delay: max(1,1,2) = 2
```

### Time Complexity Breakdown:
- **Floyd-Warshall**: O(N³) time, O(N²) space - all pairs
- **Dijkstra**: O(N²) time, O(N+E) space - single source
- **Bellman-Ford**: O(N*E) time, O(N) space - handles negatives
- **Priority Queue**: O(E log V) time, O(V) space - optimized Dijkstra

### Alternative Approaches:

#### 1. Floyd-Warshall (O(N³) time, O(N²) space)
```go
func networkDelayTimeFloydWarshall(times [][]int, n int, k int) int {
    // Build adjacency matrix
    // Run Floyd-Warshall for all pairs
    // Extract distances from source k
    // ... implementation details omitted
}
```
- **Pros**: Gives all pairs shortest paths
- **Cons**: Slower for single source problem

#### 2. Bellman-Ford (O(N*E) time, O(N) space)
```go
func networkDelayTimeBellmanFord(times [][]int, n int, k int) int {
    // Handle negative edge weights
    // Relax edges N-1 times
    // Check for negative cycles
    // ... implementation details omitted
}
```
- **Pros**: Handles negative weights
- **Cons**: Slower than Dijkstra for positive weights

#### 3. BFS for Unweighted (O(N+E) time, O(N) space)
```go
func networkDelayTimeBFS(times [][]int, n int, k int) int {
    // Only works for unweighted graphs
    // Use queue for level-order traversal
    // ... implementation details omitted
}
```
- **Pros**: Fast for unweighted graphs
- **Cons**: Doesn't work with weighted edges

### Extensions for Interviews:
- **Multiple Sources**: Find optimal source among multiple candidates
- **Path Reconstruction**: Track actual shortest paths
- **Negative Weights**: Handle graphs with negative edge weights
- **Priority Queue**: Optimize Dijkstra with heap
- **Connectivity Analysis**: Check if graph is fully connected
*/
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
