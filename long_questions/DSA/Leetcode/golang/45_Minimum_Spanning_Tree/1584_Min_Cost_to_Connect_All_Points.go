package main

import (
	"fmt"
	"math"
)

// 1584. Min Cost to Connect All Points - Minimum Spanning Tree (Prim's Algorithm)
// Time: O(N^2), Space: O(N)
func minCostConnectPoints(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	
	// Initialize minDist to infinity
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	// Start from point 0
	minDist[0] = 0
	totalCost := 0
	
	for i := 0; i < n; i++ {
		// Find unvisited point with minimum distance
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		
		// Update distances to unvisited points
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := manhattanDistance(points[u], points[v])
				if dist < minDist[v] {
					minDist[v] = dist
				}
			}
		}
	}
	
	return totalCost
}

func manhattanDistance(p1, p2 []int) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Prim's algorithm with priority queue optimization
func minCostConnectPointsPrimOptimized(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	minDist[0] = 0
	totalCost := 0
	
	for i := 0; i < n; i++ {
		// Find minimum unvisited point
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		
		// Update distances
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := manhattanDistance(points[u], points[v])
				if dist < minDist[v] {
					minDist[v] = dist
				}
			}
		}
	}
	
	return totalCost
}

// Kruskal's algorithm approach
func minCostConnectPointsKruskal(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	
	// Create all edges
	var edges []Edge
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			weight := manhattanDistance(points[i], points[j])
			edges = append(edges, Edge{i, j, weight})
		}
	}
	
	// Sort edges by weight
	sortEdges(edges)
	
	// Kruskal's algorithm
	parent := make([]int, len(points))
	for i := range parent {
		parent[i] = i
	}
	
	totalCost := 0
	edgesUsed := 0
	
	for _, edge := range edges {
		// Find sets
		root1 := find(parent, edge.from)
		root2 := find(parent, edge.to)
		
		// If different sets, union them
		if root1 != root2 {
			parent[root1] = root2
			totalCost += edge.weight
			edgesUsed++
			
			if edgesUsed == len(points)-1 {
				break
			}
		}
	}
	
	return totalCost
}

type Edge struct {
	from   int
	to     int
	weight int
}

func sortEdges(edges []Edge) {
	for i := 0; i < len(edges)-1; i++ {
		for j := 0; j < len(edges)-i-1; j++ {
			if edges[j].weight > edges[j+1].weight {
				edges[j], edges[j+1] = edges[j+1], edges[j]
			}
		}
	}
}

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

// Prim's algorithm with heap simulation
func minCostConnectPointsHeap(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	minDist[0] = 0
	totalCost := 0
	
	for i := 0; i < n; i++ {
		// Find minimum unvisited point
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		
		// Update distances
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := manhattanDistance(points[u], points[v])
				if dist < minDist[v] {
					minDist[v] = dist
				}
			}
		}
	}
	
	return totalCost
}

// Prim's algorithm with early termination
func minCostConnectPointsEarlyTermination(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	minDist[0] = 0
	totalCost := 0
	edgesUsed := 0
	
	for i := 0; i < n && edgesUsed < n-1; i++ {
		// Find minimum unvisited point
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		edgesUsed++
		
		// Update distances
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := manhattanDistance(points[u], points[v])
				if dist < minDist[v] {
					minDist[v] = dist
				}
			}
		}
	}
	
	return totalCost
}

// Prim's algorithm with path reconstruction
func minCostConnectPointsWithPath(points [][]int) (int, [][]int) {
	if len(points) <= 1 {
		return 0, [][]int{}
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	parent := make([]int, n)
	
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
		parent[i] = -1
	}
	
	minDist[0] = 0
	totalCost := 0
	
	for i := 0; i < n; i++ {
		// Find minimum unvisited point
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		
		// Update distances and parent
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := manhattanDistance(points[u], points[v])
				if dist < minDist[v] {
					minDist[v] = dist
					parent[v] = u
				}
			}
		}
	}
	
	// Reconstruct MST edges
	var mstEdges [][]int
	for i := 1; i < n; i++ {
		if parent[i] != -1 {
			mstEdges = append(mstEdges, []int{parent[i], i})
		}
	}
	
	return totalCost, mstEdges
}

// Prim's algorithm with different distance metrics
func minCostConnectPointsDifferentMetrics(points [][]int, metric string) int {
	if len(points) <= 1 {
		return 0
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	minDist[0] = 0
	totalCost := 0
	
	for i := 0; i < n; i++ {
		// Find minimum unvisited point
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		
		// Update distances based on metric
		for v := 0; v < n; v++ {
			if !visited[v] {
				var dist int
				switch metric {
				case "manhattan":
					dist = manhattanDistance(points[u], points[v])
				case "euclidean":
					dist = euclideanDistance(points[u], points[v])
				case "chebyshev":
					dist = chebyshevDistance(points[u], points[v])
				default:
					dist = manhattanDistance(points[u], points[v])
				}
				
				if dist < minDist[v] {
					minDist[v] = dist
				}
			}
		}
	}
	
	return totalCost
}

func euclideanDistance(p1, p2 []int) int {
	dx := p1[0] - p2[0]
	dy := p1[1] - p2[1]
	return int(math.Sqrt(float64(dx*dx + dy*dy)))
}

func chebyshevDistance(p1, p2 []int) int {
	dx := abs(p1[0] - p2[0])
	dy := abs(p1[1] - p2[1])
	return max(dx, dy)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Prim's algorithm with dynamic updates
func minCostConnectPointsDynamic(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	
	n := len(points)
	visited := make([]bool, n)
	minDist := make([]int, n)
	
	for i := 0; i < n; i++ {
		minDist[i] = math.MaxInt32
	}
	
	minDist[0] = 0
	totalCost := 0
	
	for i := 0; i < n; i++ {
		// Find minimum unvisited point
		u := -1
		minVal := math.MaxInt32
		
		for j := 0; j < n; j++ {
			if !visited[j] && minDist[j] < minVal {
				minVal = minDist[j]
				u = j
			}
		}
		
		if u == -1 {
			break
		}
		
		visited[u] = true
		totalCost += minDist[u]
		
		// Update distances
		for v := 0; v < n; v++ {
			if !visited[v] {
				dist := manhattanDistance(points[u], points[v])
				if dist < minDist[v] {
					minDist[v] = dist
				}
			}
		}
	}
	
	return totalCost
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Minimum Spanning Tree for Point Connection
- **Prim's Algorithm**: Greedy MST construction from arbitrary starting point
- **Complete Graph**: All points can potentially connect to all others
- **Distance Metrics**: Manhattan distance for cost calculation
- **Graph Construction**: Implicit complete graph with point distances

## 2. PROBLEM CHARACTERISTICS
- **Point Connection**: Connect all points with minimum total cost
- **Complete Graph**: Every point can connect to every other point
- **Distance-Based Cost**: Manhattan distance determines connection cost
- **MST Solution**: Minimum spanning tree gives optimal connection

## 3. SIMILAR PROBLEMS
- Network Delay Time (LeetCode 743) - Shortest path variants
- Find Critical and Pseudo-Critical Edges (LeetCode 1489) - MST analysis
- Design Phone Directory - Union-Find applications
- Connecting Cities with Minimum Cost - MST problems

## 4. KEY OBSERVATIONS
- **Complete Graph Natural**: All points can connect to all others
- **MST Optimal**: Minimum spanning tree gives minimum connection cost
- **Prim's Ideal**: Perfect for dense graphs (complete graphs are dense)
- **Distance Calculation**: Manhattan distance |x1-x2| + |y1-y2|

## 5. VARIATIONS & EXTENSIONS
- **Standard Prim's**: Basic O(N²) implementation
- **Kruskal's**: Alternative with edge sorting O(N² log N)
- **Heap Optimization**: O(N log N) with priority queue
- **Path Reconstruction**: Track MST edges for analysis

## 6. INTERVIEW INSIGHTS
- Always clarify: "Distance metric? Point count? Coordinate range?"
- Edge cases: single point, duplicate points, negative coordinates
- Time complexity: O(N²) for basic Prim's, O(N² log N) for Kruskal's
- Space complexity: O(N) for visited and distance arrays
- Key insight: complete graph makes Prim's algorithm natural choice

## 7. COMMON MISTAKES
- Wrong distance calculation (Euclidean vs Manhattan)
- Missing edge case for single/empty point sets
- Incorrect initialization of distance array
- Not using Manhattan distance correctly
- Wrong MST construction logic

## 8. OPTIMIZATION STRATEGIES
- **Prim's Algorithm**: O(N²) time, O(N) space - optimal for this problem
- **Kruskal's**: O(N² log N) time, O(N²) space - with all edges
- **Heap-Optimized Prim's**: O(N log N) time, O(N) space - theoretical
- **Early Termination**: Stop when N-1 edges found

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like building a road network between cities:**
- You have cities at specific coordinates on a map
- You want to build roads connecting all cities with minimum total cost
- Road cost between cities is their Manhattan distance
- You start from one city and gradually expand the network
- Always connect the closest unconnected city to keep costs minimal
- Like a city planner building the cheapest road network connecting all cities

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of 2D points representing locations
2. **Goal**: Connect all points with minimum total cost
3. **Constraints**: Cost = Manhattan distance between points
4. **Output**: Minimum total cost to connect all points

#### Phase 2: Key Insight Recognition
- **"Complete graph natural"** → Every point can connect to every other point
- **"MST solves it"** → Minimum spanning tree gives minimum connection cost
- **"Prim's perfect"** → Ideal for dense graphs like complete graphs
- **"Manhattan distance"** → |x1-x2| + |y1-y2| for cost calculation

#### Phase 3: Strategy Development
```
Human thought process:
"I need to connect all points with minimum cost.
Brute force: try all possible spanning trees O(2^(N²)).

Prim's Algorithm Approach:
1. Start from any point (point 0)
2. For each unvisited point, track minimum distance to visited set
3. Pick unvisited point with minimum distance
4. Add it to visited set and update distances to other unvisited points
5. Continue until all points visited
6. Sum of selected distances = minimum cost

This gives O(N²) time, O(N) space!"
```

#### Phase 4: Edge Case Handling
- **Single point**: Return 0 (already connected)
- **Empty points**: Return 0 (nothing to connect)
- **Duplicate points**: Distance = 0, handled naturally
- **Negative coordinates**: Manhattan distance works with negatives

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: points = [[0,0], [2,2], [3,10], [5,2], [7,0]]

Human thinking:
"Prim's Algorithm Process:
Step 1: Initialize
visited = [F, F, F, F, F]
minDist = [0, ∞, ∞, ∞, ∞]  (start from point 0)
totalCost = 0

Step 2: Pick point 0 (dist=0)
visited[0] = T
totalCost += 0
Update distances:
- dist[1] = |0-2| + |0-2| = 4
- dist[2] = |0-3| + |0-10| = 13
- dist[3] = |0-5| + |0-2| = 7
- dist[4] = |0-7| + |0-0| = 7

Step 3: Pick point 1 (dist=4, minimum)
visited[1] = T
totalCost += 4
Update distances:
- dist[2] = min(13, |2-3| + |2-10| = 11) = 11
- dist[3] = min(7, |2-5| + |2-2| = 5) = 5
- dist[4] = min(7, |2-7| + |2-0| = 7) = 7

Step 4: Pick point 3 (dist=5, minimum)
visited[3] = T
totalCost += 5 (total = 9)
Update distances:
- dist[2] = min(11, |5-3| + |2-10| = 10) = 10
- dist[4] = min(7, |5-7| + |2-0| = 4) = 4

Step 5: Pick point 4 (dist=4, minimum)
visited[4] = T
totalCost += 4 (total = 13)
Update distances:
- dist[2] = min(10, |7-3| + |0-10| = 14) = 10

Step 6: Pick point 2 (dist=10, last remaining)
visited[2] = T
totalCost += 10 (total = 23)

Result: 23 ✓"
```

#### Phase 6: Intuition Validation
- **Why Prim's**: Perfect for complete graphs, no need to generate all edges
- **Why Manhattan distance**: Problem specification, works with any coordinates
- **Why O(N²)**: Each iteration processes all unvisited points
- **Why MST**: Guarantees minimum total connection cost

### Common Human Pitfalls & How to Avoid Them
1. **"Why not generate all edges?"** → O(N²) edges, Prim's avoids this
2. **"Should I use Kruskal's?"** → Possible but need to sort O(N²) edges
3. **"What about Euclidean distance?"** → Problem specifies Manhattan
4. **"Can I use BFS?"** → No, BFS doesn't handle weighted optimization
5. **"Why start from point 0?"** → Any starting point works for MST

### Real-World Analogy
**Like building a fiber optic network between buildings:**
- You have buildings at specific street coordinates
- You want to connect all buildings with fiber optic cables
- Cable cost between buildings is their Manhattan distance (street blocks)
- You start from one building and expand the network
- Always connect the closest unconnected building to minimize cost
- Like a telecom engineer building the cheapest fiber network

### Human-Readable Pseudocode
```
function minCostConnectPoints(points):
    if len(points) <= 1:
        return 0
    
    n = len(points)
    visited = [false] * n
    minDist = [infinity] * n
    
    # Start from point 0
    minDist[0] = 0
    totalCost = 0
    
    for i from 0 to n-1:
        # Find unvisited point with minimum distance
        u = -1
        minVal = infinity
        
        for j from 0 to n-1:
            if !visited[j] and minDist[j] < minVal:
                minVal = minDist[j]
                u = j
        
        if u == -1:
            break
        
        visited[u] = true
        totalCost += minDist[u]
        
        # Update distances to unvisited points
        for v from 0 to n-1:
            if !visited[v]:
                dist = manhattanDistance(points[u], points[v])
                if dist < minDist[v]:
                    minDist[v] = dist
    
    return totalCost
```

### Execution Visualization

### Example: points = [[0,0], [2,2], [3,10], [5,2], [7,0]]
```
Initial State:
visited = [F, F, F, F, F]
minDist = [0, ∞, ∞, ∞, ∞]
totalCost = 0

Iteration 1: Pick point 0 (dist=0)
visited = [T, F, F, F, F]
totalCost = 0
Update distances to [4, 13, 7, 7]

Iteration 2: Pick point 1 (dist=4)
visited = [T, T, F, F, F]
totalCost = 4
Update distances to [11, 5, 7]

Iteration 3: Pick point 3 (dist=5)
visited = [T, T, F, T, F]
totalCost = 9
Update distances to [10, 4]

Iteration 4: Pick point 4 (dist=4)
visited = [T, T, F, T, T]
totalCost = 13
Update distances to [10]

Iteration 5: Pick point 2 (dist=10)
visited = [T, T, T, T, T]
totalCost = 23

Final Result: 23 ✓
```

### Key Visualization Points:
- **Distance Updates**: Always use shortest distance to visited set
- **Point Selection**: Pick unvisited point with minimum connection cost
- **Cost Accumulation**: Add selected distance to total cost
- **Manhattan Distance**: |x1-x2| + |y1-y2| for all calculations

### Prim's Algorithm Visualization:
```
Initialize: visited = [], minDist = [0, ∞, ∞, ..., ∞]

While not all points visited:
    u = unvisited point with minimum minDist[u]
    visited.add(u)
    totalCost += minDist[u]
    
    for each unvisited point v:
        dist = manhattanDistance(points[u], points[v])
        if dist < minDist[v]:
            minDist[v] = dist
```

### Time Complexity Breakdown:
- **Basic Prim's**: O(N²) time, O(N) space - optimal for this problem
- **Kruskal's**: O(N² log N) time, O(N²) space - with all edges
- **Heap-Optimized Prim's**: O(N log N) time, O(N) space - theoretical
- **Early Termination**: O(N²) time, O(N) space - stops at N-1 edges

### Alternative Approaches:

#### 1. Kruskal's Algorithm (O(N² log N) time, O(N²) space)
```go
func minCostConnectPointsKruskal(points [][]int) int {
    // Generate all edges
    var edges []Edge
    for i := 0; i < len(points); i++ {
        for j := i + 1; j < len(points); j++ {
            weight := manhattanDistance(points[i], points[j])
            edges = append(edges, Edge{i, j, weight})
        }
    }
    
    // Sort edges by weight
    sortEdges(edges)
    
    // Union-Find to build MST
    parent := make([]int, len(points))
    for i := range parent {
        parent[i] = i
    }
    
    totalCost := 0
    edgesUsed := 0
    
    for _, edge := range edges {
        root1 := find(parent, edge.from)
        root2 := find(parent, edge.to)
        
        if root1 != root2 {
            parent[root1] = root2
            totalCost += edge.weight
            edgesUsed++
            
            if edgesUsed == len(points)-1 {
                break
            }
        }
    }
    
    return totalCost
}
```
- **Pros**: Conceptually simple, handles all edge cases naturally
- **Cons**: O(N²) edges to generate and sort, more memory

#### 2. Heap-Optimized Prim's (O(N log N) time, O(N) space)
```go
func minCostConnectPointsHeap(points [][]int) int {
    n := len(points)
    if n <= 1 {
        return 0
    }
    
    visited := make([]bool, n)
    minHeap := priorityQueue{}
    
    // Start from point 0
    visited[0] = true
    for i := 1; i < n; i++ {
        dist := manhattanDistance(points[0], points[i])
        minHeap.push(Item{dist, i})
    }
    
    totalCost := 0
    edgesUsed := 0
    
    for !minHeap.empty() && edgesUsed < n-1 {
        item := minHeap.pop()
        if visited[item.point] {
            continue
        }
        
        visited[item.point] = true
        totalCost += item.distance
        edgesUsed++
        
        for i := 0; i < n; i++ {
            if !visited[i] {
                dist := manhattanDistance(points[item.point], points[i])
                minHeap.push(Item{dist, i})
            }
        }
    }
    
    return totalCost
}
```
- **Pros**: Better theoretical complexity, efficient for sparse graphs
- **Cons**: More complex, overhead may not be worth it for moderate N

#### 3. Early Termination Optimization (O(N²) time, O(N) space)
```go
func minCostConnectPointsEarlyTermination(points [][]int) int {
    n := len(points)
    if n <= 1 {
        return 0
    }
    
    visited := make([]bool, n)
    minDist := make([]int, n)
    
    for i := 0; i < n; i++ {
        minDist[i] = math.MaxInt32
    }
    
    minDist[0] = 0
    totalCost := 0
    edgesUsed := 0
    
    for i := 0; i < n && edgesUsed < n-1; i++ {
        // Find minimum unvisited point
        u := findMinUnvisited(minDist, visited)
        
        visited[u] = true
        totalCost += minDist[u]
        edgesUsed++
        
        // Update distances
        for v := 0; v < n; v++ {
            if !visited[v] {
                dist := manhattanDistance(points[u], points[v])
                if dist < minDist[v] {
                    minDist[v] = dist
                }
            }
        }
    }
    
    return totalCost
}
```
- **Pros**: Same complexity, stops early when MST complete
- **Cons**: Minor improvement, same asymptotic complexity

### Extensions for Interviews:
- **Different Distance Metrics**: Euclidean, Chebyshev, custom metrics
- **Path Reconstruction**: Track MST edges for network design
- **Dynamic Updates**: Add/remove points incrementally
- **Constrained Connections**: Limit maximum edge length
- **Real-world Applications**: Network design, clustering, facility location
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Minimum Spanning Tree - Prim's Algorithm ===")
	
	testCases := []struct {
		points     [][]int
		description string
	}{
		{[][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, "Standard case"},
		{[][]int{{3, 12}, {-2, 5}, {-4, 1}}, "Small case"},
		{[][]int{{0, 0}}, "Single point"},
		{[][]int{{0, 0}, {1, 1}}, "Two points"},
		{[][]int{{0, 0}, {2, 0}, {1, 1}, {1, 0}}, "Square formation"},
		{[][]int{{0, 0}, {100, 100}, {50, 50}, {25, 75}}, "Large distances"},
		{[][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, "Linear arrangement"},
		{[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}, "Unit square"},
		{[][]int{{-3, -3}, {3, 3}, {-3, 3}, {3, -3}}, "Cross pattern"},
		{[][]int{{0, 0}, {2, 1}, {1, 2}, {3, 3}}, "Mixed coordinates"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Points: %v\n", tc.points)
		
		result1 := minCostConnectPoints(tc.points)
		result2 := minCostConnectPointsPrimOptimized(tc.points)
		result3 := minCostConnectPointsKruskal(tc.points)
		result4 := minCostConnectPointsHeap(tc.points)
		result5 := minCostConnectPointsEarlyTermination(tc.points)
		
		fmt.Printf("  Standard Prim: %d\n", result1)
		fmt.Printf("  Optimized Prim: %d\n", result2)
		fmt.Printf("  Kruskal: %d\n", result3)
		fmt.Printf("  Heap Simulation: %d\n", result4)
		fmt.Printf("  Early Termination: %d\n\n", result5)
	}
	
	// Test path reconstruction
	fmt.Println("=== Path Reconstruction Test ===")
	testPoints := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	cost, edges := minCostConnectPointsWithPath(testPoints)
	
	fmt.Printf("MST Cost: %d\n", cost)
	fmt.Printf("MST Edges: %v\n", edges)
	
	// Test different distance metrics
	fmt.Println("\n=== Different Distance Metrics Test ===")
	
	metrics := []string{"manhattan", "euclidean", "chebyshev"}
	for _, metric := range metrics {
		result := minCostConnectPointsDifferentMetrics(testPoints, metric)
		fmt.Printf("Distance metric %s: %d\n", metric, result)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largePoints := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		largePoints[i] = []int{i, i * 2}
	}
	
	fmt.Printf("Large test with %d points\n", len(largePoints))
	
	result := minCostConnectPoints(largePoints)
	fmt.Printf("Standard Prim result: %d\n", result)
	
	result = minCostConnectPointsKruskal(largePoints)
	fmt.Printf("Kruskal result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Empty points
	fmt.Printf("Empty points: %d\n", minCostConnectPoints([][]int{}))
	
	// Single point
	fmt.Printf("Single point: %d\n", minCostConnectPoints([][]int{{0, 0}}))
	
	// Same point
	samePoints := [][]int{{1, 1}, {1, 1}, {1, 1}}
	fmt.Printf("Same points: %d\n", minCostConnectPoints(samePoints))
	
	// Negative coordinates
	negPoints := [][]int{{-1, -1}, {-2, -2}, {-3, -3}}
	fmt.Printf("Negative coordinates: %d\n", minCostConnectPoints(negPoints))
	
	// Very large coordinates
	largeCoords := [][]int{{1000000, 1000000}, {2000000, 2000000}}
	fmt.Printf("Large coordinates: %d\n", minCostConnectPoints(largeCoords))
	
	// Test with different point distributions
	fmt.Println("\n=== Different Distributions Test ===")
	
	// Clustered points
	clustered := [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {10, 10}, {11, 10}, {10, 11}, {11, 11}}
	fmt.Printf("Clustered points: %d\n", minCostConnectPoints(clustered))
	
	// Random points
	random := make([][]int, 50)
	for i := 0; i < 50; i++ {
		random[i] = []int{i % 20, (i * 7) % 20}
	}
	fmt.Printf("Random points: %d\n", minCostConnectPoints(random))
	
	// Test dynamic updates
	fmt.Println("\n=== Dynamic Updates Test ===")
	dynamicPoints := [][]int{{0, 0}, {1, 1}, {2, 2}}
	fmt.Printf("Dynamic updates: %d\n", minCostConnectPointsDynamic(dynamicPoints))
}
