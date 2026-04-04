package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents a weighted edge in the graph
type Edge struct {
	to     int
	weight int
}

// 743. Network Delay Time - Dijkstra's Algorithm
// Time: O((V + E) log V), Space: O(V + E)
func networkDelayTime(times [][]int, n int, k int) int {
	// Build adjacency list
	adj := make(map[int][]Edge)
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		adj[from] = append(adj[from], Edge{to, weight})
	}
	
	// Dijkstra's algorithm
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	
	// Min-heap: {distance, node}
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, Item{0, k})
	
	for minHeap.Len() > 0 {
		current := heap.Pop(minHeap).(Item)
		currentDist, currentNode := current.distance, current.node
		
		// Skip if we've found a better path
		if currentDist > dist[currentNode] {
			continue
		}
		
		// Relax edges
		for _, edge := range adj[currentNode] {
			newDist := currentDist + edge.weight
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				heap.Push(minHeap, Item{newDist, edge.to})
			}
		}
	}
	
	// Find the maximum distance
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1 // Unreachable node
		}
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	
	return maxDist
}

// Min-heap implementation
type Item struct {
	distance int
	node     int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Bellman-Ford algorithm (handles negative weights)
func networkDelayTimeBellmanFord(times [][]int, n int, k int) int {
	// Initialize distances
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	
	// Relax edges n-1 times
	for i := 0; i < n-1; i++ {
		updated := false
		for _, time := range times {
			from, to, weight := time[0], time[1], time[2]
			
			if dist[from] != math.MaxInt32 && dist[from]+weight < dist[to] {
				dist[to] = dist[from] + weight
				updated = true
			}
		}
		
		if !updated {
			break // No more updates needed
		}
	}
	
	// Check for unreachable nodes
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1
		}
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	
	return maxDist
}

// SPFA (Shortest Path Faster Algorithm) - optimized Bellman-Ford
func networkDelayTimeSPFA(times [][]int, n int, k int) int {
	// Build adjacency list
	adj := make(map[int][]Edge)
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		adj[from] = append(adj[from], Edge{to, weight})
	}
	
	// Initialize distances
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0
	
	// Queue for SPFA
	queue := []int{k}
	inQueue := make([]bool, n+1)
	inQueue[k] = true
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		inQueue[current] = false
		
		// Relax edges
		for _, edge := range adj[current] {
			newDist := dist[current] + edge.weight
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				
				if !inQueue[edge.to] {
					queue = append(queue, edge.to)
					inQueue[edge.to] = true
				}
			}
		}
	}
	
	// Find maximum distance
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1
		}
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	
	return maxDist
}

// Floyd-Warshall algorithm (all pairs shortest paths)
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
	
	// Set direct edges
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		dist[from][to] = weight
	}
	
	// Floyd-Warshall
	for mid := 1; mid <= n; mid++ {
		for from := 1; from <= n; from++ {
			for to := 1; to <= n; to++ {
				if dist[from][mid] != math.MaxInt32 && dist[mid][to] != math.MaxInt32 {
					if dist[from][mid]+dist[mid][to] < dist[from][to] {
						dist[from][to] = dist[from][mid] + dist[mid][to]
					}
				}
			}
		}
	}
	
	// Find maximum distance from source k
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[k][i] == math.MaxInt32 {
			return -1
		}
		if dist[k][i] > maxDist {
			maxDist = dist[k][i]
		}
	}
	
	return maxDist
}

// BFS for unweighted graphs
func networkDelayTimeBFS(times [][]int, n int, k int) int {
	// Build adjacency list (unweighted)
	adj := make(map[int][]int)
	for _, time := range times {
		from, to := time[0], time[1]
		adj[from] = append(adj[from], to)
	}
	
	// BFS from source k
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = -1
	}
	dist[k] = 0
	
	queue := []int{k}
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		for _, neighbor := range adj[current] {
			if dist[neighbor] == -1 {
				dist[neighbor] = dist[current] + 1
				queue = append(queue, neighbor)
			}
		}
	}
	
	// Find maximum distance
	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] == -1 {
			return -1
		}
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

## 1. ALGORITHM PATTERN: Single-Source Shortest Path in Weighted Graph
- **Dijkstra's Algorithm**: Priority queue-based shortest path from source
- **Bellman-Ford**: Edge relaxation for graphs with negative weights
- **SPFA**: Queue-based optimization of Bellman-Ford
- **Floyd-Warshall**: All-pairs shortest paths with dynamic programming

## 2. PROBLEM CHARACTERISTICS
- **Weighted Directed Graph**: Network edges with travel times
- **Single Source**: Find shortest paths from one source node
- **Network Delay**: Maximum time for signal to reach all nodes
- **Connectivity**: All nodes must be reachable from source

## 3. SIMILAR PROBLEMS
- Cheapest Flights Within K Stops (LeetCode 787) - Constrained shortest path
- Path with Maximum Probability (LeetCode 1514) - Max probability path
- Network Delay Time (LeetCode 743) - Same problem
- Minimum Spanning Tree (LeetCode 1584) - Different optimization goal

## 4. KEY OBSERVATIONS
- **Dijkstra Optimal**: Works for non-negative weights
- **Bellman-Ford General**: Handles negative weights, detects cycles
- **SPFA Efficient**: Optimized Bellman-Ford for sparse graphs
- **Maximum Distance**: Need max of all shortest paths from source

## 5. VARIATIONS & EXTENSIONS
- **Priority Queue**: Dijkstra with min-heap for efficiency
- **Edge Relaxation**: Bellman-Ford with early termination
- **Queue Optimization**: SPFA for better average performance
- **All Pairs**: Floyd-Warshall for complete distance matrix

## 6. INTERVIEW INSIGHTS
- Always clarify: "Edge weights negative? Graph size? Multiple queries?"
- Edge cases: disconnected graph, single node, source unreachable
- Time complexity: O((V+E) log V) for Dijkstra, O(VE) for Bellman-Ford
- Space complexity: O(V+E) for adjacency list, O(V²) for matrix
- Key insight: different algorithms for different weight constraints

## 7. COMMON MISTAKES
- Using Dijkstra with negative weights
- Not handling disconnected nodes properly
- Missing early termination in Bellman-Ford
- Wrong priority queue implementation
- Not checking for unreachable nodes

## 8. OPTIMIZATION STRATEGIES
- **Dijkstra**: O((V+E) log V) time, O(V+E) space - non-negative weights
- **Bellman-Ford**: O(VE) time, O(V) space - handles negative weights
- **SPFA**: O(VE) worst, O(E) average time, O(V) space - optimized
- **Floyd-Warshall**: O(V³) time, O(V²) space - all pairs

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like calculating delivery times in a network:**
- Each node is a city, edges are delivery routes with travel times
- You start from one source city and want fastest delivery to all others
- The network delay is the time until the last city receives its package
- You need the shortest path from source to every other city
- Like calculating the maximum delivery time in a logistics network

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Network edges (from, to, weight), number of nodes, source node
2. **Goal**: Find shortest paths from source to all nodes
3. **Rules**: Return maximum shortest path time, -1 if any node unreachable
4. **Output**: Network delay time or -1 for disconnected graph

#### Phase 2: Key Insight Recognition
- **"Shortest path natural"** → Need shortest path algorithm from source
- **"Weight constraints matter"** → Choose algorithm based on edge weights
- **"Maximum of minima"** → Network delay = max of all shortest paths
- **"Connectivity check"** → Must verify all nodes reachable

#### Phase 3: Strategy Development
```
Human thought process:
"I need shortest paths from source to all nodes.
Different algorithms for different weight constraints.

Dijkstra Approach (non-negative weights):
1. Build adjacency list from edges
2. Use min-heap: {distance, node}
3. Initialize distances: source=0, others=∞
4. Extract min, relax edges, update distances
5. Return max distance or -1 if unreachable

This gives optimal O((V+E) log V) for non-negative weights!"
```

#### Phase 4: Edge Case Handling
- **Disconnected graph**: Return -1 if any node unreachable
- **Single node**: Return 0 (no travel needed)
- **Negative weights**: Use Bellman-Ford instead of Dijkstra
- **Empty graph**: Handle gracefully based on constraints

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: times=[[2,1,1],[2,3,1],[3,4,1]], n=4, k=2

Human thinking:
"Dijkstra Approach:
Step 1: Build adjacency list
2: [1(1), 3(1)]
3: [4(1)]
1: []
4: []

Step 2: Initialize distances
dist=[∞,∞,∞,∞] (1-based)
dist[2]=0

Step 3: Priority queue: {(0,2)}
Extract (0,2):
- Relax 2→1: dist[1]=1, push (1,1)
- Relax 2→3: dist[3]=1, push (1,3)

Step 4: Extract (1,1):
- No outgoing edges

Step 5: Extract (1,3):
- Relax 3→4: dist[4]=2, push (2,4)

Step 6: Extract (2,4):
- No outgoing edges

Final distances: [∞,1,0,1,2]
Max distance = 2 ✓"
```

#### Phase 6: Intuition Validation
- **Why Dijkstra works**: Greedy choice optimal for non-negative weights
- **Why priority queue**: Always expand shortest unprocessed path first
- **Why max distance**: Network delay = time for farthest node
- **Why O((V+E) log V)**: Each edge processed once, heap operations dominate

### Common Human Pitfalls & How to Avoid Them
1. **"Why not BFS?"** → BFS only for unweighted graphs
2. **"Should I use Bellman-Ford?"** → Only if negative weights possible
3. **"What about Floyd-Warshall?"** → Overkill for single-source problem
4. **"Can I use adjacency matrix?"** → Less efficient for sparse graphs
5. **"Why check unreachable?"** → Must return -1 if any node unreachable

### Real-World Analogy
**Like a delivery network planning system:**
- Each city is a node, delivery routes are weighted edges
- You start from one warehouse and want fastest delivery to all cities
- The network delay is when the last city gets its package
- You calculate shortest paths using fastest routes
- Like optimizing a logistics network for minimum maximum delivery time

### Human-Readable Pseudocode
```
function networkDelayTime(times, n, k):
    # Build adjacency list
    adj = adjacency list of size n+1
    for from, to, weight in times:
        adj[from].append((to, weight))
    
    # Dijkstra's algorithm
    dist = array of size n+1, initialized to infinity
    dist[k] = 0
    
    # Min-heap: {distance, node}
    minHeap = priority queue
    minHeap.push((0, k))
    
    while minHeap is not empty:
        currentDist, currentNode = minHeap.pop()
        
        # Skip outdated entries
        if currentDist > dist[currentNode]:
            continue
        
        # Relax edges
        for neighbor, weight in adj[currentNode]:
            newDist = currentDist + weight
            if newDist < dist[neighbor]:
                dist[neighbor] = newDist
                minHeap.push((newDist, neighbor))
    
    # Find maximum distance
    maxDist = 0
    for i from 1 to n:
        if dist[i] == infinity:
            return -1  # Unreachable node
        maxDist = max(maxDist, dist[i])
    
    return maxDist
```

### Execution Visualization

### Example: times=[[2,1,1],[2,3,1],[3,4,1]], n=4, k=2
```
Dijkstra Process:

Step 1: Build adjacency list
2: [(1,1), (3,1)]
3: [(4,1)]
1: []
4: []

Step 2: Initialize distances
dist = [∞,∞,0,∞,∞] (1-based indexing)
Heap: [(0,2)]

Step 3: Extract (0,2)
- Relax 2→1: dist[1] = 1, push (1,1)
- Relax 2→3: dist[3] = 1, push (1,3)
Heap: [(1,1), (1,3)]

Step 4: Extract (1,1)
- No outgoing edges from 1
Heap: [(1,3)]

Step 5: Extract (1,3)
- Relax 3→4: dist[4] = 2, push (2,4)
Heap: [(2,4)]

Step 6: Extract (2,4)
- No outgoing edges from 4
Heap: []

Final distances: [∞,1,0,1,2]
Max distance = 2 ✓
```

### Key Visualization Points:
- **Priority Queue**: Always expand shortest unprocessed path
- **Edge Relaxation**: Update distances when shorter path found
- **Greedy Choice**: Optimal for non-negative weights
- **Network Delay**: Maximum of all shortest paths

### Graph Visualization:
```
Network Graph:
    1 (1) ← 2 (0)
    ↓
    3 (1) → 4 (2)

Shortest paths from 2:
- 2→1: distance 1
- 2→3: distance 1  
- 2→3→4: distance 2

Network delay = max(1,1,2) = 2 ✓
```

### Time Complexity Breakdown:
- **Dijkstra**: O((V+E) log V) time, O(V+E) space - non-negative weights
- **Bellman-Ford**: O(VE) time, O(V) space - handles negative weights
- **SPFA**: O(VE) worst, O(E) average time, O(V) space - optimized
- **Floyd-Warshall**: O(V³) time, O(V²) space - all pairs

### Alternative Approaches:

#### 1. Bellman-Ford (O(VE) time, O(V) space)
```go
func networkDelayTimeBellmanFord(times [][]int, n int, k int) int {
    // Edge relaxation V-1 times
    // Handles negative weights
    // Detects negative cycles
    // ... implementation details omitted
}
```
- **Pros**: Handles negative weights, simple implementation
- **Cons**: Slower than Dijkstra for non-negative weights

#### 2. SPFA (O(VE) worst, O(E) average time, O(V) space)
```go
func networkDelayTimeSPFA(times [][]int, n int, k int) int {
    // Queue-based optimization of Bellman-Ford
    // Only relax nodes whose distances changed
    // Better average performance
    // ... implementation details omitted
}
```
- **Pros**: Faster than Bellman-Ford on average
- **Cons**: Still O(VE) worst case

#### 3. Floyd-Warshall (O(V³) time, O(V²) space)
```go
func networkDelayTimeFloydWarshall(times [][]int, n int, k int) int {
    // All pairs shortest paths
    // Dynamic programming approach
    // Overkill for single source
    // ... implementation details omitted
}
```
- **Pros**: Gives all pairs distances
- **Cons**: Unnecessary for single-source problem

### Extensions for Interviews:
- **Multiple Sources**: Find shortest paths from multiple sources
- **Path Reconstruction**: Track actual shortest paths
- **Negative Cycles**: Detect and handle negative weight cycles
- **Dynamic Updates**: Handle edge weight changes efficiently
- **Real-world Applications**: Network routing, logistics optimization
*/
func main() {
	// Test cases
	testCases := []struct {
		times      [][]int
		n          int
		k          int
		description string
	}{
		{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, "Standard case"},
		{[][]int{{1, 2, 1}}, 2, 1, "Simple case"},
		{[][]int{{1, 2, 1}}, 2, 2, "Source at end"},
		{[][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}}, 3, 1, "Multiple paths"},
		{[][]int{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}, {4, 5, 1}}, 5, 1, "Linear chain"},
		{[][]int{{1, 2, 1}, {1, 3, 2}, {2, 4, 1}, {3, 4, 1}}, 4, 1, "Converging paths"},
		{[][]int{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}, {4, 1, 1}}, 4, 1, "Cycle"},
		{[][]int{{1, 2, 1}, {2, 3, 10}, {1, 3, 5}}, 3, 1, "Direct vs indirect"},
		{[][]int{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}, {4, 5, 1}, {5, 6, 1}}, 6, 1, "Long chain"},
		{[][]int{{1, 2, 1}, {1, 3, 1}, {2, 4, 1}, {3, 4, 1}}, 4, 1, "Multiple sources"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Times: %v, n=%d, k=%d\n", tc.times, tc.n, tc.k)
		
		result1 := networkDelayTime(tc.times, tc.n, tc.k)
		result2 := networkDelayTimeBellmanFord(tc.times, tc.n, tc.k)
		result3 := networkDelayTimeSPFA(tc.times, tc.n, tc.k)
		result4 := networkDelayTimeFloydWarshall(tc.times, tc.n, tc.k)
		result5 := networkDelayTimeBFS(tc.times, tc.n, tc.k)
		
		fmt.Printf("  Dijkstra: %d\n", result1)
		fmt.Printf("  Bellman-Ford: %d\n", result2)
		fmt.Printf("  SPFA: %d\n", result3)
		fmt.Printf("  Floyd-Warshall: %d\n", result4)
		fmt.Printf("  BFS (unweighted): %d\n\n", result5)
	}
}
