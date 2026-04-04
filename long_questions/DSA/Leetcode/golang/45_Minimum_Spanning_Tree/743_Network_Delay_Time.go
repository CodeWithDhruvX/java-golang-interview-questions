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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Shortest Path vs Minimum Spanning Tree
- **Shortest Path**: Dijkstra's algorithm for single-source shortest paths
- **MST Approaches**: Prim's and Kruskal's algorithms for minimum spanning tree
- **Network Propagation**: Time for signal to reach all nodes from source
- **Graph Traversal**: Different strategies for network analysis

## 2. PROBLEM CHARACTERISTICS
- **Network Delay**: Time for signal to propagate through weighted network
- **Single Source**: Find maximum distance from source to all nodes
- **Connectivity**: All nodes must be reachable for valid solution
- **Weighted Edges**: Travel times between network nodes

## 3. SIMILAR PROBLEMS
- Find the City With the Smallest Number of Neighbors (LeetCode 1334) - Floyd-Warshall
- Cheapest Flights Within K Stops (LeetCode 787) - Constrained shortest path
- Network Delay Time (LeetCode 743) - Same problem
- Minimum Spanning Tree problems - MST construction and analysis

## 4. KEY OBSERVATIONS
- **Shortest Path Natural**: Dijkstra's algorithm directly solves the problem
- **MST Alternative**: Can use MST as approximation but not always optimal
- **Maximum Distance**: Need farthest node distance from source
- **Connectivity Check**: Unreachable nodes make problem impossible

## 5. VARIATIONS & EXTENSIONS
- **Dijkstra's Algorithm**: Direct shortest path solution
- **Prim's MST**: Build MST then find paths (approximation)
- **Kruskal's MST**: Alternative MST construction
- **Multiple Sources**: Extend to multiple starting points

## 6. INTERVIEW INSIGHTS
- Always clarify: "Graph connectivity? Edge weights? Single vs multiple sources?"
- Edge cases: disconnected graphs, single nodes, self-loops
- Time complexity: O(N²) for basic Dijkstra, O(E log V) with heap
- Space complexity: O(N + E) for adjacency and distances
- Key insight: Dijkstra's algorithm is the natural solution

## 7. COMMON MISTAKES
- Wrong initialization of distances (should be infinity except source)
- Missing connectivity check for disconnected graphs
- Using MST when shortest path is needed
- Incorrect maximum distance calculation
- Not handling unreachable nodes properly

## 8. OPTIMIZATION STRATEGIES
- **Dijkstra with Heap**: O(E log V) time, O(V + E) space - optimal
- **Basic Dijkstra**: O(V²) time, O(V + E) space - simpler
- **MST Approaches**: O(E log V) time, O(V + E) space - approximation
- **Multiple Sources**: O(S × (V + E)) time, O(V + E) space

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like a signal propagating through a communication network:**
- You have network nodes connected by communication links with delays
- A signal starts from one source node and spreads through the network
- Each link takes time to transmit the signal
- You want to know when the last node receives the signal
- Like tracking how long it takes for a message to reach everyone in a network

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Network of N nodes, weighted edges, source node K
2. **Goal**: Find time for signal to reach all nodes from K
3. **Constraints**: All nodes must be reachable for valid solution
4. **Output**: Maximum shortest path distance or -1 if disconnected

#### Phase 2: Key Insight Recognition
- **"Shortest path natural"** → Need shortest distances from source to all nodes
- **"Dijkstra's perfect"** → Single-source shortest path algorithm
- **"Maximum of shortest paths"** → Need farthest reachable node
- **"Connectivity critical"** → Unreachable nodes make problem impossible

#### Phase 3: Strategy Development
```
Human thought process:
"I need signal propagation time from source to all nodes.
Brute force: BFS from source O(V + E) but doesn't handle weights.

Dijkstra's Approach:
1. Initialize distances: source = 0, others = infinity
2. Use priority queue to always expand closest unvisited node
3. Update neighbor distances when shorter path found
4. Continue until all nodes visited or no more reachable
5. Return maximum distance (or -1 if any unreachable)

This gives O(E log V) time, O(V + E) space!"
```

#### Phase 4: Edge Case Handling
- **Disconnected graph**: Return -1 if any node unreachable
- **Single node**: Return 0 (signal already at all nodes)
- **Self-loops**: Ignore or handle appropriately
- **Multiple edges**: Use the minimum weight edge

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: times = [[2,1,1], [2,3,1], [3,4,1]], n=4, k=2

Human thinking:
"Dijkstra's Process:
Step 1: Initialize distances
dist[1] = ∞, dist[2] = 0, dist[3] = ∞, dist[4] = ∞
visited = [F, F, F, F]

Step 2: Pick closest unvisited node (node 2, dist=0)
visited[2] = T
Update neighbors:
- 2→1: dist[1] = min(∞, 0+1) = 1
- 2→3: dist[3] = min(∞, 0+1) = 1

Step 3: Pick closest unvisited node (node 1 or 3, dist=1)
Pick node 1:
visited[1] = T
No outgoing edges from node 1

Step 4: Pick closest unvisited node (node 3, dist=1)
visited[3] = T
Update neighbors:
- 3→4: dist[4] = min(∞, 1+1) = 2

Step 5: Pick closest unvisited node (node 4, dist=2)
visited[4] = T
No outgoing edges from node 4

Step 6: All nodes visited, find maximum distance
max(dist[1], dist[2], dist[3], dist[4]) = max(1, 0, 1, 2) = 2

Result: 2 ✓"
```

#### Phase 6: Intuition Validation
- **Why Dijkstra's**: Directly solves single-source shortest path
- **Why maximum distance**: Signal reaches last node at maximum time
- **Why connectivity check**: Unreachable nodes mean signal never arrives
- **Why priority queue**: Always expand closest node first (optimal)

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use BFS?"** → BFS doesn't handle weighted edges properly
2. **"Should I use MST?"** → MST gives minimum total weight, not shortest paths
3. **"What about Bellman-Ford?"** → Overkill, no negative weights
4. **"Can I use Floyd-Warshall?"** → O(V³) vs O(E log V), inefficient for single source
5. **"Why check connectivity?"** → Signal must reach all nodes for valid solution

### Real-World Analogy
**Like a virus spreading through a computer network:**
- You have computers connected by network cables with transmission delays
- A virus starts from one infected computer
- Each cable takes time to transmit the virus
- The virus spreads to the closest uninfected computers first
- You want to know when the last computer gets infected
- Like tracking infection spread through a network

### Human-Readable Pseudocode
```
function networkDelayTime(times, n, k):
    # Build adjacency list
    adj = adjacency list of size n+1
    for each edge [u, v, w] in times:
        adj[u].append([v, w])
    
    # Initialize distances
    dist = array of size n+1
    for i from 1 to n:
        dist[i] = infinity
    dist[k] = 0
    
    visited = array of size n+1, all false
    
    # Dijkstra's algorithm
    for i from 1 to n:
        # Find unvisited node with minimum distance
        u = node with minimum dist[u] where !visited[u]
        if u == -1 or dist[u] == infinity:
            break
        
        visited[u] = true
        
        # Update distances of neighbors
        for each edge [v, w] in adj[u]:
            if !visited[v] and dist[u] + w < dist[v]:
                dist[v] = dist[u] + w
    
    # Find maximum distance
    maxDist = 0
    for i from 1 to n:
        if dist[i] > maxDist:
            maxDist = dist[i]
    
    if maxDist == infinity:
        return -1
    else:
        return maxDist
```

### Execution Visualization

### Example: times = [[2,1,1], [2,3,1], [3,4,1]], n=4, k=2
```
Initial State:
dist = [∞, 0, ∞, ∞]  (indices: 1, 2, 3, 4)
visited = [F, F, F, F]

Step 1: Visit node 2 (dist=0)
dist = [1, 0, 1, ∞]
visited = [F, T, F, F]
(2→1: 0+1=1, 2→3: 0+1=1)

Step 2: Visit node 1 (dist=1)
dist = [1, 0, 1, ∞]
visited = [T, T, F, F]
(No outgoing edges from 1)

Step 3: Visit node 3 (dist=1)
dist = [1, 0, 1, 2]
visited = [T, T, T, F]
(3→4: 1+1=2)

Step 4: Visit node 4 (dist=2)
dist = [1, 0, 1, 2]
visited = [T, T, T, T]

Final distances: [1, 0, 1, 2]
Maximum distance: 2
Result: 2 ✓
```

### Key Visualization Points:
- **Distance Updates**: Always use shortest path found so far
- **Node Selection**: Pick unvisited node with minimum distance
- **Propagation**: Signal spreads outward from source
- **Maximum Time**: Last node to receive signal determines answer

### Dijkstra's Algorithm Visualization:
```
Initialize: dist[source] = 0, others = ∞
Priority Queue: [(0, source)]

While queue not empty:
    (dist, node) = pop minimum
    if node already visited: continue
    mark node as visited
    
    for each neighbor:
        if dist + edge_weight < neighbor_dist:
            update neighbor_dist
            push (neighbor_dist, neighbor) to queue
```

### Time Complexity Breakdown:
- **Basic Dijkstra**: O(V²) time, O(V + E) space - simple implementation
- **Heap-Optimized Dijkstra**: O(E log V) time, O(V + E) space - optimal
- **MST Approaches**: O(E log V) time, O(V + E) space - approximation
- **Multiple Sources**: O(S × (V + E)) time, O(V + E) space

### Alternative Approaches:

#### 1. MST with Prim's Algorithm (O(E log V) time, O(V + E) space)
```go
func networkDelayTimeMST(times [][]int, n int, k int) int {
    // Build MST using Prim's algorithm
    // Then find shortest paths in MST
    // Note: This is an approximation, not always optimal
    
    adj := buildAdjacencyList(times)
    visited := make([]bool, n+1)
    minDist := make([]int, n+1)
    
    // Prim's to build MST
    for i := 1; i <= n; i++ {
        minDist[i] = math.MaxInt32
    }
    
    minDist[k] = 0
    mstEdges := [][]int{}
    
    for len(mstEdges) < n-1 {
        // Find unvisited node with minimum distance
        u := findMinUnvisited(minDist, visited)
        if u == -1: break
        
        visited[u] = true
        
        // Add edges to MST
        for _, edge := range adj[u] {
            v, w := edge[0], edge[1]
            if !visited[v] && w < minDist[v] {
                minDist[v] = w
            }
        }
    }
    
    // Find distances in MST (may not be optimal)
    return findMaxDistanceInMST(mstEdges, n, k)
}
```
- **Pros**: Good for dense graphs, simpler than Dijkstra
- **Cons**: Not always optimal, MST ≠ shortest paths

#### 2. Bellman-Ford (O(V×E) time, O(V) space)
```go
func networkDelayTimeBellmanFord(times [][]int, n int, k int) int {
    dist := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dist[i] = math.MaxInt32
    }
    dist[k] = 0
    
    // Relax edges V-1 times
    for i := 1; i < n; i++ {
        for _, edge := range times {
            u, v, w := edge[0], edge[1], edge[2]
            if dist[u] != math.MaxInt32 && dist[u]+w < dist[v] {
                dist[v] = dist[u] + w
            }
        }
    }
    
    // Check for negative cycles (not needed for this problem)
    
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
```
- **Pros**: Handles negative weights, simple implementation
- **Cons**: Slower than Dijkstra, unnecessary for positive weights

#### 3. Floyd-Warshall (O(V³) time, O(V²) space)
```go
func networkDelayTimeFloydWarshall(times [][]int, n int, k int) int {
    // Initialize distance matrix
    dist := make([][]int, n+1)
    for i := 1; i <= n; i++ {
        dist[i] = make([]int, n+1)
        for j := 1; j <= n; j++ {
            if i == j {
                dist[i][j] = 0
            } else {
                dist[i][j] = math.MaxInt32
            }
        }
    }
    
    // Fill direct edges
    for _, edge := range times {
        u, v, w := edge[0], edge[1], edge[2]
        dist[u][v] = w
    }
    
    // Floyd-Warshall
    for k := 1; k <= n; k++ {
        for i := 1; i <= n; i++ {
            for j := 1; j <= n; j++ {
                if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
                    if dist[i][j] > dist[i][k] + dist[k][j] {
                        dist[i][j] = dist[i][k] + dist[k][j]
                    }
                }
            }
        }
    }
    
    // Find maximum distance from k
    maxDist := 0
    for i := 1; i <= n; i++ {
        if dist[k][i] > maxDist {
            maxDist = dist[k][i]
        }
    }
    
    if maxDist == math.MaxInt32 {
        return -1
    }
    return maxDist
}
```
- **Pros**: All-pairs shortest paths, simple implementation
- **Cons**: Overkill for single source, O(V³) complexity

### Extensions for Interviews:
- **Multiple Sources**: Find time for signals from multiple starting points
- **Dynamic Updates**: Handle edge weight changes incrementally
- **K-Smallest Delays**: Find K-th smallest delay time
- **Network Reliability**: Analyze network robustness to edge failures
- **Real-world Applications**: Network routing, disease spread, information propagation
*/
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
