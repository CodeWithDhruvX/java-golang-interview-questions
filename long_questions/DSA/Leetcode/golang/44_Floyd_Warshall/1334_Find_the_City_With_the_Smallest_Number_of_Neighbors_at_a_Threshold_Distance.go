package main

import (
	"fmt"
	"math"
)

// 1334. Find the City With the Smallest Number of Neighbors at a Threshold Distance - Floyd-Warshall Algorithm
// Time: O(N^3), Space: O(N^2)
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// Initialize distance matrix
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	
	// Fill direct edges
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		if weight < dist[from][to] {
			dist[from][to] = weight
			dist[to][from] = weight // Undirected graph
		}
	}
	
	// Floyd-Warshall algorithm
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					if dist[i][j] > dist[i][k]+dist[k][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
					}
				}
			}
		}
	}
	
	// Find city with minimum neighbors within threshold
	minNeighbors := math.MaxInt32
	result := -1
	
	for i := 0; i < n; i++ {
		neighbors := 0
		for j := 0; j < n; j++ {
			if i != j && dist[i][j] <= distanceThreshold {
				neighbors++
			}
		}
		
		if neighbors < minNeighbors || (neighbors == minNeighbors && i > result) {
			minNeighbors = neighbors
			result = i
		}
	}
	
	return result
}

// Optimized Floyd-Warshall with early termination
func findTheCityOptimized(n int, edges [][]int, distanceThreshold int) int {
	// Initialize distance matrix
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	
	// Fill direct edges
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		if weight < dist[from][to] {
			dist[from][to] = weight
			dist[to][from] = weight
		}
	}
	
	// Floyd-Warshall with optimization
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if dist[i][k] == math.MaxInt32 {
				continue // Skip if no path to k
			}
			for j := 0; j < n; j++ {
				if dist[k][j] == math.MaxInt32 {
					continue // Skip if no path from k
				}
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	
	// Find city with minimum neighbors
	minNeighbors := math.MaxInt32
	result := -1
	
	for i := 0; i < n; i++ {
		neighbors := 0
		for j := 0; j < n; j++ {
			if i != j && dist[i][j] <= distanceThreshold {
				neighbors++
			}
		}
		
		if neighbors < minNeighbors || (neighbors == minNeighbors && i > result) {
			minNeighbors = neighbors
			result = i
		}
	}
	
	return result
}

// Floyd-Warshall with path reconstruction
func findTheCityWithPathReconstruction(n int, edges [][]int, distanceThreshold int) (int, [][]int) {
	// Initialize distance and path matrices
	dist := make([][]int, n)
	next := make([][]int, n)
	
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		next[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
				next[i][j] = j
			} else {
				dist[i][j] = math.MaxInt32
				next[i][j] = -1
			}
		}
	}
	
	// Fill direct edges
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		if weight < dist[from][to] {
			dist[from][to] = weight
			next[from][to] = to
			dist[to][from] = weight
			next[to][from] = from
		}
	}
	
	// Floyd-Warshall with path reconstruction
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					if dist[i][j] > dist[i][k]+dist[k][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
						next[i][j] = next[i][k]
					}
				}
			}
		}
	}
	
	// Find city with minimum neighbors
	minNeighbors := math.MaxInt32
	result := -1
	
	for i := 0; i < n; i++ {
		neighbors := 0
		for j := 0; j < n; j++ {
			if i != j && dist[i][j] <= distanceThreshold {
				neighbors++
			}
		}
		
		if neighbors < minNeighbors || (neighbors == minNeighbors && i > result) {
			minNeighbors = neighbors
			result = i
		}
	}
	
	return result, next
}

// Reconstruct path between two cities
func reconstructPath(next [][]int, from, to int) []int {
	if next[from][to] == -1 {
		return []int{}
	}
	
	path := []int{from}
	for from != to {
		from = next[from][to]
		path = append(path, from)
	}
	
	return path
}

// Floyd-Warshall with multiple sources optimization
func findTheCityMultipleSources(n int, edges [][]int, distanceThreshold int, sources []int) int {
	// Initialize distance matrix
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	
	// Fill direct edges
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		if weight < dist[from][to] {
			dist[from][to] = weight
			dist[to][from] = weight
		}
	}
	
	// Floyd-Warshall
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					if dist[i][j] > dist[i][k]+dist[k][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
					}
				}
			}
		}
	}
	
	// Find best source city
	minNeighbors := math.MaxInt32
	result := -1
	
	for _, source := range sources {
		neighbors := 0
		for j := 0; j < n; j++ {
			if source != j && dist[source][j] <= distanceThreshold {
				neighbors++
			}
		}
		
		if neighbors < minNeighbors || (neighbors == minNeighbors && source > result) {
			minNeighbors = neighbors
			result = source
		}
	}
	
	return result
}

// Floyd-Warshall with dynamic threshold
func findTheCityDynamicThreshold(n int, edges [][]int, distanceThreshold int) (int, map[int]int) {
	// Initialize distance matrix
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	
	// Fill direct edges
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		if weight < dist[from][to] {
			dist[from][to] = weight
			dist[to][from] = weight
		}
	}
	
	// Floyd-Warshall
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 {
					if dist[i][j] > dist[i][k]+dist[k][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
					}
				}
			}
		}
	}
	
	// Calculate neighbors for each threshold
	thresholdToCity := make(map[int]int)
	
	for threshold := 1; threshold <= distanceThreshold; threshold++ {
		minNeighbors := math.MaxInt32
		bestCity := -1
		
		for i := 0; i < n; i++ {
			neighbors := 0
			for j := 0; j < n; j++ {
				if i != j && dist[i][j] <= threshold {
					neighbors++
				}
			}
			
			if neighbors < minNeighbors || (neighbors == minNeighbors && i > bestCity) {
				minNeighbors = neighbors
				bestCity = i
			}
		}
		
		thresholdToCity[threshold] = bestCity
	}
	
	// Return result for original threshold
	return thresholdToCity[distanceThreshold], thresholdToCity
}

// Alternative approach using Dijkstra from each city
func findTheCityDijkstra(n int, edges [][]int, distanceThreshold int) int {
	// Build adjacency list
	adj := make([][][]int, n)
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		adj[from] = append(adj[from], []int{to, weight})
		adj[to] = append(adj[to], []int{from, weight})
	}
	
	minNeighbors := math.MaxInt32
	result := -1
	
	// Run Dijkstra from each city
	for i := 0; i < n; i++ {
		distances := dijkstra(n, adj, i)
		
		neighbors := 0
		for j := 0; j < n; j++ {
			if i != j && distances[j] <= distanceThreshold {
				neighbors++
			}
		}
		
		if neighbors < minNeighbors || (neighbors == minNeighbors && i > result) {
			minNeighbors = neighbors
			result = i
		}
	}
	
	return result
}

func dijkstra(n int, adj [][]int, start int) []int {
	dist := make([]int, n)
	visited := make([]bool, n)
	
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0
	
	for count := 0; count < n; count++ {
		// Find unvisited vertex with minimum distance
		minDist := math.MaxInt32
		u := -1
		
		for v := 0; v < n; v++ {
			if !visited[v] && dist[v] < minDist {
				minDist = dist[v]
				u = v
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		
		// Update distances of adjacent vertices
		for _, edge := range adj[u] {
			v, weight := edge[0], edge[1]
			if !visited[v] && dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
			}
		}
	}
	
	return dist
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Floyd-Warshall All-Pairs Shortest Path
- **Dynamic Programming**: Build shortest paths incrementally
- **All-Pairs Shortest Path**: Compute distances between all city pairs
- **Intermediate Vertices**: Consider each vertex as potential intermediate
- **Path Relaxation**: Continuously improve distance estimates

## 2. PROBLEM CHARACTERISTICS
- **Graph Analysis**: Weighted undirected graph with cities and edges
- **Threshold Filtering**: Count neighbors within distance threshold
- **Optimization**: Find city with minimum reachable neighbors
- **Complete Graph**: Need all-pairs distances for analysis

## 3. SIMILAR PROBLEMS
- Evaluate Division (LeetCode 399) - Floyd-Warshall for equation solving
- Network Delay Time (LeetCode 743) - Single-source shortest path
- Find the City With the Smallest Number of Neighbors (LeetCode 1334) - Same problem
- Cheapest Flights Within K Stops (LeetCode 787) - Bellman-Ford variant

## 4. KEY OBSERVATIONS
- **Complete Distances**: Need all-pairs shortest paths for neighbor counting
- **Floyd-Warshall Ideal**: Perfect for dense graphs with all-pairs queries
- **Threshold Logic**: Count neighbors within distance limit
- **Tie Breaking**: Choose highest numbered city on ties

## 5. VARIATIONS & EXTENSIONS
- **Standard Floyd-Warshall**: Basic all-pairs shortest path
- **Optimized Version**: Early termination for unreachable nodes
- **Path Reconstruction**: Store next matrix for path retrieval
- **Multiple Sources**: Query from specific source cities

## 6. INTERVIEW INSIGHTS
- Always clarify: "Graph density? Edge weights? Multiple queries?"
- Edge cases: disconnected graphs, single cities, extreme thresholds
- Time complexity: O(N³) for Floyd-Warshall, O(N×E×logN) for N×Dijkstra
- Space complexity: O(N²) for distance matrix
- Key insight: Floyd-Warshall optimal for dense graphs with all-pairs queries

## 7. COMMON MISTAKES
- Wrong initialization of distance matrix
- Missing edge case for unreachable nodes
- Incorrect tie-breaking logic
- Integer overflow in distance calculations
- Wrong graph representation (directed vs undirected)

## 8. OPTIMIZATION STRATEGIES
- **Floyd-Warshall**: O(N³) time, O(N²) space - optimal for dense graphs
- **N×Dijkstra**: O(N×E×logN) time, O(N²) space - better for sparse graphs
- **Optimized Floyd-Warshall**: Early termination for unreachable nodes
- **Path Reconstruction**: O(N³) time, O(N²) space - with path tracking

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like calculating travel times between all cities:**
- You want to know the shortest travel time between every pair of cities
- You consider each city as a potential layover/stopover point
- You progressively improve your travel time estimates
- For each threshold distance, you count how many cities are reachable
- You find the city with the fewest reachable neighbors
- Like a travel planner computing all possible routes and finding the most isolated city

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: N cities, weighted edges, distance threshold
2. **Goal**: Find city with minimum neighbors within threshold
3. **Constraints**: Need all-pairs distances for neighbor counting
4. **Output**: City index with fewest reachable neighbors

#### Phase 2: Key Insight Recognition
- **"All-pairs needed"** → Need distances between every city pair
- **"Floyd-Warshall natural"** → Perfect for all-pairs shortest paths
- **"Threshold counting"** → Count neighbors within distance limit
- **"Tie breaking"** → Choose highest numbered city

#### Phase 3: Strategy Development
```
Human thought process:
"I need to count neighbors within threshold for each city.
Brute force: run Dijkstra from each city O(N×E×logN).

Floyd-Warshall Approach:
1. Initialize distance matrix with direct edges
2. For each intermediate city k:
   - Update all pairs (i,j) through k
   - dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
3. Count neighbors within threshold for each city
4. Find city with minimum neighbors (highest index on ties)

This gives O(N³) time, O(N²) space!"
```

#### Phase 4: Edge Case Handling
- **Single city**: Return 0 (only city available)
- **No edges**: All cities isolated, return highest index
- **Disconnected graph**: Handle unreachable nodes properly
- **Zero threshold**: Only count direct neighbors

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: 4 cities, edges: [[0,1,3], [1,2,1], [1,3,4], [2,3,1]], threshold=4

Human thinking:
"Floyd-Warshall Process:
Step 1: Initialize distance matrix
  0  3  ∞  ∞
  3  0  1  4
  ∞  1  0  1
  ∞  4  1  0

Step 2: k=0 (city 0 as intermediate)
Update paths through city 0
  0  3  ∞  ∞
  3  0  1  4
  ∞  1  0  1
  ∞  4  1  0
(No improvements through city 0)

Step 3: k=1 (city 1 as intermediate)
Update paths through city 1
  0  3  4  7
  3  0  1  4
  4  1  0  1
  7  4  1  0
(0→2: 3+1=4, 0→3: 3+4=7, etc.)

Continue for k=2, k=3...
Final distance matrix:
  0  3  4  5
  3  0  1  2
  4  1  0  1
  5  2  1  0

Step 4: Count neighbors within threshold=4
City 0: neighbors to 1(3), 2(4), 3(5) → 2 neighbors
City 1: neighbors to 0(3), 2(1), 3(2) → 3 neighbors
City 2: neighbors to 0(4), 1(1), 3(1) → 3 neighbors
City 3: neighbors to 1(2), 2(1) → 2 neighbors

Step 5: Find minimum (ties go to higher index)
Cities 0 and 3 both have 2 neighbors → choose 3

Result: 3 ✓"
```

#### Phase 6: Intuition Validation
- **Why Floyd-Warshall**: Computes all-pairs distances efficiently
- **Why O(N³)**: Triple nested loop for all intermediate vertices
- **Why threshold counting**: Direct application of computed distances
- **Why tie breaking**: Problem requires highest index on ties

### Common Human Pitfalls & How to Avoid Them
1. **"Why not Dijkstra from each city?"** → O(N×E×logN) vs O(N³), depends on graph density
2. **"Should I use Bellman-Ford?"** → No, no negative weights
3. **"What about sparse graphs?"** → N×Dijkstra might be better
4. **"Can I optimize space?"** → Yes, but need all distances for counting
5. **"Why initialize with infinity?"** → Represents unreachable initially

### Real-World Analogy
**Like planning a transportation network:**
- You have cities connected by roads with distances
- You want to know the shortest route between every pair of cities
- You consider each city as a potential transfer point
- You build a complete distance table
- For each city, you count how many other cities are within a certain distance
- You find the most isolated city (fewest nearby cities)
- Like a logistics company planning delivery routes and finding the most remote location

### Human-Readable Pseudocode
```
function findTheCity(n, edges, distanceThreshold):
    # Initialize distance matrix
    dist = n×n matrix
    for i from 0 to n-1:
        for j from 0 to n-1:
            if i == j: dist[i][j] = 0
            else: dist[i][j] = infinity
    
    # Fill direct edges
    for each edge [from, to, weight]:
        dist[from][to] = weight
        dist[to][from] = weight  # Undirected
    
    # Floyd-Warshall algorithm
    for k from 0 to n-1:
        for i from 0 to n-1:
            for j from 0 to n-1:
                if dist[i][k] + dist[k][j] < dist[i][j]:
                    dist[i][j] = dist[i][k] + dist[k][j]
    
    # Count neighbors within threshold
    minNeighbors = infinity
    result = -1
    
    for i from 0 to n-1:
        neighbors = 0
        for j from 0 to n-1:
            if i != j and dist[i][j] <= distanceThreshold:
                neighbors += 1
        
        if neighbors < minNeighbors or (neighbors == minNeighbors and i > result):
            minNeighbors = neighbors
            result = i
    
    return result
```

### Execution Visualization

### Example: 4 cities, threshold=4
```
Initial Distance Matrix:
    0   1   2   3
0 [ 0,  3,  ∞,  ∞ ]
1 [ 3,  0,  1,  4 ]
2 [ ∞,  1,  0,  1 ]
3 [ ∞,  4,  1,  0 ]

After Floyd-Warshall:
    0   1   2   3
0 [ 0,  3,  4,  5 ]
1 [ 3,  0,  1,  2 ]
2 [ 4,  1,  0,  1 ]
3 [ 5,  2,  1,  0 ]

Neighbor Counting (threshold=4):
City 0: 1(3), 2(4), 3(5) → 2 neighbors
City 1: 0(3), 2(1), 3(2) → 3 neighbors  
City 2: 0(4), 1(1), 3(1) → 3 neighbors
City 3: 1(2), 2(1) → 2 neighbors

Result: City 3 (highest index among minimums) ✓
```

### Key Visualization Points:
- **Distance Matrix**: Complete shortest paths between all pairs
- **Intermediate Updates**: Each city considered as transfer point
- **Threshold Filtering**: Count only neighbors within distance limit
- **Tie Breaking**: Choose highest numbered city on equal neighbors

### Floyd-Warshall Process Visualization:
```
For each intermediate city k:
  For each source city i:
    For each destination city j:
      dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])

This progressively builds optimal paths by considering all possible
intermediate cities as transfer points.
```

### Time Complexity Breakdown:
- **Floyd-Warshall**: O(N³) time, O(N²) space - optimal for dense graphs
- **N×Dijkstra**: O(N×E×logN) time, O(N²) space - better for sparse graphs
- **Optimized Floyd-Warshall**: O(N³) time, O(N²) space - with early termination
- **Path Reconstruction**: O(N³) time, O(N²) space - with path tracking

### Alternative Approaches:

#### 1. N×Dijkstra (O(N×E×logN) time, O(N²) space)
```go
func findTheCityDijkstra(n int, edges [][]int, distanceThreshold int) int {
    // Build adjacency list
    adj := make([][][]int, n)
    for _, edge := range edges {
        from, to, weight := edge[0], edge[1], edge[2]
        adj[from] = append(adj[from], []int{to, weight})
        adj[to] = append(adj[to], []int{from, weight})
    }
    
    // Run Dijkstra from each city
    minNeighbors := math.MaxInt32
    result := -1
    
    for i := 0; i < n; i++ {
        distances := dijkstra(n, adj, i)
        neighbors := 0
        for j := 0; j < n; j++ {
            if i != j && distances[j] <= distanceThreshold {
                neighbors++
            }
        }
        
        if neighbors < minNeighbors || (neighbors == minNeighbors && i > result) {
            minNeighbors = neighbors
            result = i
        }
    }
    
    return result
}
```
- **Pros**: Better for sparse graphs, no O(N³) complexity
- **Cons**: More complex, multiple priority queue operations

#### 2. Optimized Floyd-Warshall (O(N³) time, O(N²) space)
```go
func findTheCityOptimized(n int, edges [][]int, distanceThreshold int) int {
    // Same as Floyd-Warshall but with early termination
    // Skip updates when intermediate node is unreachable
    // Reduces constant factors in practice
}
```
- **Pros**: Same complexity, better practical performance
- **Cons**: Still O(N³) worst case

#### 3. Bellman-Ford (O(N×E) time, O(N²) space)
```go
func findTheCityBellmanFord(n int, edges [][]int, distanceThreshold int) int {
    // Run Bellman-Ford from each city
    // Not optimal since no negative weights
    // More complex than necessary
}
```
- **Pros**: Handles negative weights (if needed)
- **Cons**: Overkill for positive weights, slower than Dijkstra

### Extensions for Interviews:
- **Path Reconstruction**: Store intermediate nodes for path retrieval
- **Dynamic Thresholds**: Query different thresholds efficiently
- **Multiple Sources**: Find best city from specific source set
- **Graph Updates**: Handle dynamic edge additions/removals
- **Real-world Applications**: Network analysis, transportation planning, social networks
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Floyd-Warshall Algorithm ===")
	
	testCases := []struct {
		n                 int
		edges             [][]int
		distanceThreshold int
		description       string
	}{
		{
			4,
			[][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}},
			4,
			"Standard case",
		},
		{
			5,
			[][]int{{0, 1, 2}, {0, 2, 4}, {1, 2, 1}, {1, 3, 2}, {2, 3, 3}, {3, 4, 1}},
			5,
			"Larger graph",
		},
		{
			3,
			[][]int{{0, 1, 1}, {1, 2, 1}, {2, 0, 1}},
			2,
			"Triangle graph",
		},
		{
			6,
			[][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}, {2, 3, 1}, {3, 4, 1}, {4, 5, 1}, {5, 0, 10}},
			20,
			"Complex graph",
		},
		{
			2,
			[][]int{{0, 1, 1}},
			1,
			"Two cities",
		},
		{
			1,
			[][]int{},
			0,
			"Single city",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Cities: %d, Threshold: %d\n", tc.n, tc.distanceThreshold)
		fmt.Printf("  Edges: %v\n", tc.edges)
		
		result1 := findTheCity(tc.n, tc.edges, tc.distanceThreshold)
		result2 := findTheCityOptimized(tc.n, tc.edges, tc.distanceThreshold)
		result3 := findTheCityDijkstra(tc.n, tc.edges, tc.distanceThreshold)
		
		fmt.Printf("  Floyd-Warshall: %d\n", result1)
		fmt.Printf("  Optimized: %d\n", result2)
		fmt.Printf("  Dijkstra: %d\n\n", result3)
	}
	
	// Test path reconstruction
	fmt.Println("=== Path Reconstruction Test ===")
	n, edges, threshold := 4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4
	city, next := findTheCityWithPathReconstruction(n, edges, threshold)
	
	fmt.Printf("Best city: %d\n", city)
	fmt.Printf("Path from 0 to 3: %v\n", reconstructPath(next, 0, 3))
	fmt.Printf("Path from 0 to 2: %v\n", reconstructPath(next, 0, 2))
	
	// Test multiple sources
	fmt.Println("\n=== Multiple Sources Test ===")
	sources := []int{0, 2}
	result := findTheCityMultipleSources(n, edges, threshold, sources)
	fmt.Printf("Best city from sources %v: %d\n", sources, result)
	
	// Test dynamic threshold
	fmt.Println("\n=== Dynamic Threshold Test ===")
	city, thresholdMap := findTheCityDynamicThreshold(n, edges, threshold)
	fmt.Printf("Best city for threshold %d: %d\n", threshold, city)
	
	for t, bestCity := range thresholdMap {
		fmt.Printf("  Threshold %d: City %d\n", t, bestCity)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largeN := 100
	largeEdges := make([][]int, 0)
	for i := 0; i < largeN; i++ {
		for j := i + 1; j < largeN && j < i+5; j++ {
			weight := (j-i) * 10
			largeEdges = append(largeEdges, []int{i, j, weight})
		}
	}
	
	fmt.Printf("Large test with %d cities and %d edges\n", largeN, len(largeEdges))
	
	result = findTheCity(largeN, largeEdges, 50)
	fmt.Printf("Floyd-Warshall result: %d\n", result)
	
	result = findTheCityOptimized(largeN, largeEdges, 50)
	fmt.Printf("Optimized result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// No edges
	noEdges := [][]int{}
	fmt.Printf("No edges: %d\n", findTheCity(3, noEdges, 10))
	
	// Very large threshold
	veryLargeThreshold := 1000
	fmt.Printf("Very large threshold: %d\n", findTheCity(4, edges, veryLargeThreshold))
	
	// Very small threshold
	verySmallThreshold := 0
	fmt.Printf("Very small threshold: %d\n", findTheCity(4, edges, verySmallThreshold))
	
	// Disconnected graph
	disconnectedEdges := [][]int{{0, 1, 1}, {2, 3, 1}}
	fmt.Printf("Disconnected graph: %d\n", findTheCity(4, disconnectedEdges, 5))
}
