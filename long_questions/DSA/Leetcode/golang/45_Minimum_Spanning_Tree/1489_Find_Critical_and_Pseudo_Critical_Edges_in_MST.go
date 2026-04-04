package main

import (
	"fmt"
	"math"
)

// 1489. Find Critical and Pseudo-Critical Edges in MST - Minimum Spanning Tree
// Time: O(N^2) for MST, Space: O(N^2)
func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	if n <= 1 {
		return [][]int{}
	}
	
	// Find MST weight using Kruskal's algorithm
	mstWeight := findMSTWeight(n, edges)
	
	// Find critical and pseudo-critical edges
	var result [][]int
	
	for i, edge := range edges {
		// Remove this edge and find new MST weight
		newWeight := findMSTWeightWithoutEdge(n, edges, i)
		
		if newWeight > mstWeight {
			// Critical edge
			result = append(result, []int{edge[0], edge[1], edge[2], 1})
		} else if newWeight == mstWeight {
			// Pseudo-critical edge
			result = append(result, []int{edge[0], edge[1], edge[2], 2})
		}
	}
	
	return result
}

func findMSTWeight(n int, edges [][]int) int {
	if len(edges) < n-1 {
		return math.MaxInt32
	}
	
	// Sort edges by weight
	sortedEdges := make([][]int, len(edges))
	copy(sortedEdges, edges)
	
	for i := 0; i < len(sortedEdges)-1; i++ {
		for j := 0; j < len(sortedEdges)-i-1; j++ {
			if sortedEdges[j][2] > sortedEdges[j+1][2] {
				sortedEdges[j], sortedEdges[j+1] = sortedEdges[j+1], sortedEdges[j]
			}
		}
	}
	
	// Kruskal's algorithm
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	
	totalWeight := 0
	edgesUsed := 0
	
	for _, edge := range sortedEdges {
		from, to, weight := edge[0], edge[1], edge[2]
		root1 := find(parent, from)
		root2 := find(parent, to)
		
		if root1 != root2 {
			parent[root1] = root2
			totalWeight += weight
			edgesUsed++
			
			if edgesUsed == n-1 {
				break
			}
		}
	}
	
	if edgesUsed < n-1 {
		return math.MaxInt32
	}
	
	return totalWeight
}

func findMSTWeightWithoutEdge(n int, edges [][]int, excludeIdx int) int {
	if len(edges) <= n-1 {
		return math.MaxInt32
	}
	
	// Create new edges array excluding the specified edge
	newEdges := make([][]int, 0, len(edges)-1)
	for i, edge := range edges {
		if i != excludeIdx {
			newEdges = append(newEdges, edge)
		}
	}
	
	return findMSTWeight(n, newEdges)
}

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

// Optimized version with early termination
func findCriticalAndPseudoCriticalEdgesOptimized(n int, edges [][]int) [][]int {
	if n <= 1 {
		return [][]int{}
	}
	
	// Find MST weight
	mstWeight := findMSTWeight(n, edges)
	
	// Build MST adjacency list
	mstAdj := buildMSTAdjacency(n, edges)
	
	var result [][]int
	
	for i, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		
		// Check if edge is in MST
		inMST := isInMST(mstAdj, from, to, weight)
		
		if inMST {
			// Check if it's critical by removing it
			newWeight := findMSTWeightWithoutEdge(n, edges, i)
			if newWeight > mstWeight {
				result = append(result, []int{from, to, weight, 1})
			} else {
				result = append(result, []int{from, to, weight, 2})
			}
		} else {
			// Edge not in MST, check if it can create alternative MST
			alternativeWeight := findAlternativeMSTWeight(n, edges, from, to, weight)
			if alternativeWeight == mstWeight {
				result = append(result, []int{from, to, weight, 2})
			}
		}
	}
	
	return result
}

func buildMSTAdjacency(n int, edges [][]int) [][]int {
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, n)
	}
	
	// Build MST and get adjacency
	sortedEdges := make([][]int, len(edges))
	copy(sortedEdges, edges)
	
	// Sort by weight
	for i := 0; i < len(sortedEdges)-1; i++ {
		for j := 0; j < len(sortedEdges)-i-1; j++ {
			if sortedEdges[j][2] > sortedEdges[j+1][2] {
				sortedEdges[j], sortedEdges[j+1] = sortedEdges[j+1], sortedEdges[j]
			}
		}
	}
	
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	
	for _, edge := range sortedEdges {
		from, to := edge[0], edge[1]
		root1 := find(parent, from)
		root2 := find(parent, to)
		
		if root1 != root2 {
			parent[root1] = root2
			adj[from][to] = edge[2]
			adj[to][from] = edge[2]
		}
	}
	
	return adj
}

func isInMST(adj [][]int, from, to, weight int) bool {
	return adj[from][to] == weight
}

func findAlternativeMSTWeight(n int, edges [][]int, newFrom, newTo, newWeight int) int {
	// Add new edge and find MST weight
	newEdges := make([][]int, len(edges)+1)
	copy(newEdges, edges)
	newEdges[len(edges)] = []int{newFrom, newTo, newWeight}
	
	return findMSTWeight(n, newEdges)
}

// Version with Union-Find optimization
func findCriticalAndPseudoCriticalEdgesUnionFind(n int, edges [][]int) [][]int {
	if n <= 1 {
		return [][]int{}
	}
	
	// Sort edges by weight
	sortedEdges := make([][]int, len(edges))
	copy(sortedEdges, edges)
	
	for i := 0; i < len(sortedEdges)-1; i++ {
		for j := 0; j < len(sortedEdges)-i-1; j++ {
			if sortedEdges[j][2] > sortedEdges[j+1][2] {
				sortedEdges[j], sortedEdges[j+1] = sortedEdges[j+1], sortedEdges[j]
			}
		}
	}
	
	// Find MST weight and track used edges
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	
	mstWeight := 0
	usedEdges := make([]bool, len(edges))
	
	for i, edge := range sortedEdges {
		from, to, weight := edge[0], edge[1], edge[2]
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
			
			mstWeight += weight
			usedEdges[i] = true
		}
	}
	
	// Check all edges
	var result [][]int
	
	for i, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		
		if usedEdges[i] {
			// Edge in MST, check if critical
			newWeight := findMSTWeightWithoutEdge(n, edges, i)
			if newWeight > mstWeight {
				result = append(result, []int{from, to, weight, 1})
			} else {
				result = append(result, []int{from, to, weight, 2})
			}
		} else {
			// Edge not in MST, check if pseudo-critical
			alternativeWeight := findAlternativeMSTWeight(n, edges, from, to, weight)
			if alternativeWeight == mstWeight {
				result = append(result, []int{from, to, weight, 2})
			}
		}
	}
	
	return result
}

func findWithCompression(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = findWithCompression(parent, parent[x])
	}
	return parent[x]
}

// Version with detailed analysis
func findCriticalAndPseudoCriticalEdgesDetailed(n int, edges [][]int) ([][]int, map[string]int) {
	if n <= 1 {
		return [][]int{}, map[string]int{}
	}
	
	analysis := make(map[string]int)
	
	// Find MST weight
	mstWeight := findMSTWeight(n, edges)
	analysis["mst_weight"] = mstWeight
	
	// Count total edges
	totalEdges := len(edges)
	analysis["total_edges"] = totalEdges
	
	// Find critical and pseudo-critical edges
	var result [][]int
	criticalCount := 0
	pseudoCriticalCount := 0
	
	for i, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		
		newWeight := findMSTWeightWithoutEdge(n, edges, i)
		
		if newWeight > mstWeight {
			result = append(result, []int{from, to, weight, 1})
			criticalCount++
		} else if newWeight == mstWeight {
			result = append(result, []int{from, to, weight, 2})
			pseudoCriticalCount++
		}
	}
	
	analysis["critical_count"] = criticalCount
	analysis["pseudo_critical_count"] = pseudoCriticalCount
	analysis["non_critical_count"] = totalEdges - criticalCount - pseudoCriticalCount
	
	return result, analysis
}

// Version with multiple MSTs
func findCriticalAndPseudoCriticalEdgesMultipleMST(n int, edges [][]int) [][]int {
	if n <= 1 {
		return [][]int{}
	}
	
	// Find all possible MST weights
	mstWeights := findAllMSTWeights(n, edges)
	
	if len(mstWeights) == 0 {
		return [][]int{}
	}
	
	minWeight := mstWeights[0]
	
	// Analyze each edge
	var result [][]int
	
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		
		// Check if edge appears in any MST
		inAnyMST := false
		for _, mstWeight := range mstWeights {
			if edgeInMST(n, edges, from, to, weight, mstWeight) {
				inAnyMST = true
				break
			}
		}
		
		if inAnyMST {
			// Check if critical
			newWeight := findMSTWeightWithoutEdge(n, edges, findEdgeIndex(edges, edge))
			if newWeight > minWeight {
				result = append(result, []int{from, to, weight, 1})
			} else {
				result = append(result, []int{from, to, weight, 2})
			}
		}
	}
	
	return result
}

func findAllMSTWeights(n int, edges [][]int) []int {
	// This is a simplified version - in practice, finding all MSTs is complex
	// For demonstration, we'll find one MST weight
	weight := findMSTWeight(n, edges)
	if weight != math.MaxInt32 {
		return []int{weight}
	}
	return []int{}
}

func edgeInMST(n int, edges [][]int, from, to, weight int, targetWeight int) bool {
	// Simplified check - in practice, this would be more complex
	testEdges := make([][]int, len(edges))
	copy(testEdges, edges)
	testEdges = append(testEdges, []int{from, to, weight})
	
	return findMSTWeight(n, testEdges) == targetWeight
}

func findEdgeIndex(edges [][]int, target []int) int {
	for i, edge := range edges {
		if edge[0] == target[0] && edge[1] == target[1] && edge[2] == target[2] {
			return i
		}
	}
	return -1
}

// Version with cycle detection
func findCriticalAndPseudoCriticalEdgesWithCycleDetection(n int, edges [][]int) [][]int {
	if n <= 1 {
		return [][]int{}
	}
	
	// Build adjacency list
	adj := make([][][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adj[from] = append(adj[from], []int{to, edge[2]})
		adj[to] = append(adj[to], []int{from, edge[2]})
	}
	
	// Find MST weight
	mstWeight := findMSTWeight(n, edges)
	
	var result [][]int
	
	for i, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		
		// Check if removing this edge disconnects the graph
		if createsCycle(adj, from, to) {
			// Edge creates cycle, check if it's pseudo-critical
			newWeight := findMSTWeightWithoutEdge(n, edges, i)
			if newWeight == mstWeight {
				result = append(result, []int{from, to, weight, 2})
			}
		} else {
			// Edge is a bridge, check if critical
			newWeight := findMSTWeightWithoutEdge(n, edges, i)
			if newWeight > mstWeight {
				result = append(result, []int{from, to, weight, 1})
			}
		}
	}
	
	return result
}

func createsCycle(adj [][]int, from, to int) bool {
	// Simplified cycle detection
	// In practice, this would use DFS or Union-Find
	return false // Simplified for demonstration
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Minimum Spanning Tree Edge Analysis
- **Kruskal's Algorithm**: Greedy edge selection with Union-Find
- **Critical Edges**: Edges that must be in all MSTs (removal increases weight)
- **Pseudo-Critical Edges**: Edges that can be in some MSTs (alternative paths exist)
- **Edge Classification**: Analyze edge importance in MST construction

## 2. PROBLEM CHARACTERISTICS
- **MST Analysis**: Find critical and non-critical edges in minimum spanning tree
- **Edge Importance**: Determine which edges are essential vs optional
- **Multiple MSTs**: Handle graphs with multiple minimum spanning trees
- **Union-Find**: Efficient cycle detection and component management

## 3. SIMILAR PROBLEMS
- Network Delay Time (LeetCode 743) - Shortest path variants
- Min Cost to Connect All Points (LeetCode 1584) - MST construction
- Find the City With the Smallest Number of Neighbors (LeetCode 1334) - Graph analysis
- Evaluate Division (LeetCode 399) - Graph connectivity

## 4. KEY OBSERVATIONS
- **Critical Edge Test**: Remove edge and check if MST weight increases
- **Pseudo-Critical Test**: Edge can appear in some MST with same weight
- **Union-Find Essential**: Efficiently detect cycles and manage components
- **Edge Classification**: Each edge falls into one of three categories

## 5. VARIATIONS & EXTENSIONS
- **Standard Analysis**: Basic critical/pseudo-critical classification
- **Optimized Version**: Early termination and adjacency tracking
- **Union-Find Optimization**: Path compression and union by rank
- **Multiple MST Analysis**: Handle graphs with multiple optimal trees

## 6. INTERVIEW INSIGHTS
- Always clarify: "Graph connectivity? Edge weights? Multiple MSTs?"
- Edge cases: disconnected graphs, single nodes, uniform weights
- Time complexity: O(E² × log V) for analysis, O(E log V) for MST
- Space complexity: O(V + E) for Union-Find and adjacency
- Key insight: edge importance determined by MST weight comparison

## 7. COMMON MISTAKES
- Wrong MST weight calculation
- Incorrect edge classification logic
- Missing Union-Find path compression
- Not handling disconnected graphs properly
- Wrong edge removal/reconstruction

## 8. OPTIMIZATION STRATEGIES
- **Kruskal + Union-Find**: O(E log V) time, O(V + E) space - optimal
- **Prim's Algorithm**: O(E log V) time, O(V + E) space - alternative
- **Optimized Analysis**: O(E² log V) time, O(V + E) space - with caching
- **Multiple MST Handling**: O(E × MST_count) time, O(V + E) space

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like analyzing a road network's critical roads:**
- You have cities connected by roads with construction costs
- You want to build the cheapest network connecting all cities
- Some roads are essential (must be built), others are optional
- Critical roads: if removed, network becomes more expensive
- Pseudo-critical roads: can be used in some optimal networks
- Like a transportation planner identifying essential vs optional routes

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Graph with N nodes and weighted edges
2. **Goal**: Classify edges as critical, pseudo-critical, or neither
3. **Constraints**: Need to analyze edge importance in MST context
4. **Output**: Edge classification with weights and types

#### Phase 2: Key Insight Recognition
- **"MST natural"** → Need minimum spanning tree as baseline
- **"Edge removal test"** → Critical if removal increases MST weight
- **"Alternative paths"** → Pseudo-critical if alternative MSTs exist
- **"Union-Find essential"** → Efficient cycle detection for MST

#### Phase 3: Strategy Development
```
Human thought process:
"I need to classify edges by their importance in MST.
Brute force: try all edge combinations O(2^E).

Analysis Approach:
1. Find MST weight using Kruskal's algorithm
2. For each edge:
   - Remove it and find new MST weight
   - If weight increases: critical edge
   - If weight same: pseudo-critical edge
   - If weight infinite (disconnected): critical edge
3. Use Union-Find for efficient cycle detection
4. Return classification results

This gives O(E² log V) time, O(V + E) space!"
```

#### Phase 4: Edge Case Handling
- **Disconnected graph**: No MST possible, handle appropriately
- **Single node**: No edges, return empty classification
- **Uniform weights**: All edges potentially pseudo-critical
- **Multiple edges**: Handle parallel edges correctly

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: 4 nodes, edges = [[0,1,1], [1,2,1], [2,3,1], [0,3,1]]

Human thinking:
"MST Analysis Process:
Step 1: Find MST weight using Kruskal's
Sort edges by weight: all weight 1
Pick edges that don't create cycles:
- Edge (0,1): include, components: {0,1}, {2}, {3}
- Edge (1,2): include, components: {0,1,2}, {3}
- Edge (2,3): include, components: {0,1,2,3}
MST complete, weight = 3

Step 2: Analyze each edge
Edge (0,1):
- Remove it, find new MST
- Can use (0,3), (1,2), (2,3) = weight 3
- Same weight → pseudo-critical

Edge (1,2):
- Remove it, find new MST
- Can use (0,1), (0,3), (2,3) = weight 3
- Same weight → pseudo-critical

Edge (2,3):
- Remove it, find new MST
- Can use (0,1), (1,2), (0,3) = weight 3
- Same weight → pseudo-critical

Edge (0,3):
- Remove it, find new MST
- Can use (0,1), (1,2), (2,3) = weight 3
- Same weight → pseudo-critical

Result: All edges are pseudo-critical ✓"
```

#### Phase 6: Intuition Validation
- **Why MST weight**: Need baseline for comparison
- **Why edge removal**: Tests criticality by impact on optimal solution
- **Why Union-Find**: Efficiently manages connected components
- **Why classification**: Three categories cover all possibilities

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use Prim's?"** → Kruskal better for edge analysis
2. **"Should I check all subsets?"** → No, too expensive
3. **"What about multiple MSTs?"** → Handle by weight comparison
4. **"Can I optimize further?"** → Yes, with caching and early termination
5. **"Why Union-Find?"** → Essential for efficient cycle detection

### Real-World Analogy
**Like analyzing a power grid's critical transmission lines:**
- You have power plants and cities connected by transmission lines
- Each line has a construction/maintenance cost
- You want the cheapest network that powers all cities
- Critical lines: if removed, network becomes more expensive or impossible
- Optional lines: can be used in some optimal configurations
- Like an electrical engineer identifying essential vs backup power lines

### Human-Readable Pseudocode
```
function findCriticalAndPseudoCriticalEdges(n, edges):
    if n <= 1:
        return []
    
    # Find MST weight using Kruskal's algorithm
    mstWeight = findMSTWeight(n, edges)
    
    result = []
    
    for each edge i:
        # Remove this edge and find new MST weight
        newWeight = findMSTWeightWithoutEdge(n, edges, i)
        
        if newWeight > mstWeight:
            # Critical edge - removal increases weight
            result.append([edge[0], edge[1], edge[2], 1])
        elif newWeight == mstWeight:
            # Pseudo-critical edge - alternative MST exists
            result.append([edge[0], edge[1], edge[2], 2])
        # else: edge not in any MST
    
    return result

function findMSTWeight(n, edges):
    sort edges by weight
    parent = [0, 1, ..., n-1]
    
    totalWeight = 0
    edgesUsed = 0
    
    for each edge (from, to, weight):
        root1 = find(parent, from)
        root2 = find(parent, to)
        
        if root1 != root2:
            union(parent, root1, root2)
            totalWeight += weight
            edgesUsed++
            
            if edgesUsed == n-1:
                break
    
    return totalWeight if edgesUsed == n-1 else infinity
```

### Execution Visualization

### Example: 4 nodes, edges = [[0,1,1], [1,2,1], [2,3,1], [0,3,1]]
```
MST Construction (Kruskal's):
1. Sort edges: all weight 1
2. Edge (0,1): include ✓
3. Edge (1,2): include ✓  
4. Edge (2,3): include ✓
5. Edge (0,3): skip (creates cycle)

MST Weight = 3

Edge Analysis:
Edge (0,1):
- Remove, remaining edges: [(1,2), (2,3), (0,3)]
- New MST: (0,3), (1,2), (2,3) = weight 3
- Same weight → Pseudo-critical

Edge (1,2):
- Remove, remaining edges: [(0,1), (2,3), (0,3)]
- New MST: (0,1), (0,3), (2,3) = weight 3
- Same weight → Pseudo-critical

Edge (2,3):
- Remove, remaining edges: [(0,1), (1,2), (0,3)]
- New MST: (0,1), (1,2), (0,3) = weight 3
- Same weight → Pseudo-critical

Edge (0,3):
- Remove, remaining edges: [(0,1), (1,2), (2,3)]
- New MST: (0,1), (1,2), (2,3) = weight 3
- Same weight → Pseudo-critical

Result: All edges are pseudo-critical
```

### Key Visualization Points:
- **MST Baseline**: Need optimal weight for comparison
- **Edge Removal**: Tests criticality by impact on solution
- **Union-Find**: Manages connected components efficiently
- **Classification Logic**: Three categories based on weight impact

### Kruskal's Algorithm Visualization:
```
Sort edges by weight: [e1, e2, e3, e4, ...]
Initialize each node as separate component

For each edge:
    if edge connects different components:
        include edge in MST
        merge components
    else:
        skip edge (creates cycle)

Stop when MST has n-1 edges
```

### Time Complexity Breakdown:
- **Kruskal's MST**: O(E log V) time, O(V + E) space
- **Edge Analysis**: O(E × (E log V)) time, O(V + E) space
- **Union-Find**: Nearly O(1) amortized per operation
- **Optimized Version**: O(E² log V) time, O(V + E) space

### Alternative Approaches:

#### 1. Prim's Algorithm (O(E log V) time, O(V + E) space)
```go
func findMSTWeightPrim(n int, edges [][]int) int {
    // Build adjacency list
    adj := buildAdjacencyList(n, edges)
    
    // Prim's algorithm with priority queue
    visited := make([]bool, n)
    minHeap := priorityQueue{}
    
    // Start from node 0
    visited[0] = true
    for _, edge := range adj[0] {
        minHeap.push(edge)
    }
    
    totalWeight := 0
    edgesUsed := 0
    
    for !minHeap.empty() && edgesUsed < n-1 {
        edge := minHeap.pop()
        if !visited[edge.to] {
            visited[edge.to] = true
            totalWeight += edge.weight
            edgesUsed++
            
            for _, nextEdge := range adj[edge.to] {
                if !visited[nextEdge.to] {
                    minHeap.push(nextEdge)
                }
            }
        }
    }
    
    return totalWeight
}
```
- **Pros**: Good for dense graphs, no need to sort all edges
- **Cons**: More complex, requires priority queue

#### 2. Optimized Edge Analysis (O(E² log V) time, O(V + E) space)
```go
func findCriticalAndPseudoCriticalEdgesOptimized(n int, edges [][]int) [][]int {
    // Pre-compute MST weight
    mstWeight := findMSTWeight(n, edges)
    
    // Build MST adjacency for quick edge lookup
    mstAdj := buildMSTAdjacency(n, edges)
    
    result := make([][]int, 0)
    
    for i, edge := range edges {
        // Quick check if edge is in MST
        inMST := isInMST(mstAdj, edge)
        
        if inMST {
            // Only test edges that are in MST
            newWeight := findMSTWeightWithoutEdge(n, edges, i)
            if newWeight > mstWeight {
                result = append(result, []int{edge[0], edge[1], edge[2], 1})
            } else {
                result = append(result, []int{edge[0], edge[1], edge[2], 2})
            }
        }
    }
    
    return result
}
```
- **Pros**: Faster in practice, skips unnecessary computations
- **Cons**: Still O(E²) worst case, more complex

#### 3. Multiple MST Handling (O(E × MST_count) time, O(V + E) space)
```go
func findAllMSTClassifications(n int, edges [][]int) [][]int {
    // Find all possible MST weights
    mstWeights := findAllMSTWeights(n, edges)
    
    if len(mstWeights) == 0 {
        return []
    }
    
    minWeight := mstWeights[0]
    
    // Classify edges based on appearance in any MST
    result := make([][]int, 0)
    
    for _, edge := range edges {
        appearsInAnyMST := false
        for _, weight := range mstWeights {
            if edgeInMST(n, edges, edge, weight) {
                appearsInAnyMST = true
                break
            }
        }
        
        if appearsInAnyMST {
            // Check criticality by removal test
            newWeight := findMSTWeightWithoutEdge(n, edges, findEdgeIndex(edges, edge))
            if newWeight > minWeight {
                result = append(result, []int{edge[0], edge[1], edge[2], 1})
            } else {
                result = append(result, []int{edge[0], edge[1], edge[2], 2})
            }
        }
    }
    
    return result
}
```
- **Pros**: Handles multiple optimal solutions correctly
- **Cons**: Complex, finding all MSTs is computationally expensive

### Extensions for Interviews:
- **Edge Weight Changes**: How do classifications change with weight updates?
- **Dynamic Graphs**: Handle edge additions/removals incrementally
- **Constrained MST**: MST with specific edge requirements
- **Network Design**: Apply to real-world network optimization
- **Real-world Applications**: Network design, circuit design, resource allocation
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Critical and Pseudo-Critical Edges in MST ===")
	
	testCases := []struct {
		n          int
		edges      [][]int
		description string
	}{
		{
			4,
			[][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}},
			"Simple cycle",
		},
		{
			5,
			[][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {3, 4, 1}, {4, 0, 1}},
			"Complex graph",
		},
		{
			3,
			[][]int{{0, 1, 1}, {1, 2, 2}, {0, 2, 3}},
			"Triangle",
		},
		{
			2,
			[][]int{{0, 1, 1}},
			"Single edge",
		},
		{
			4,
			[][]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 3}, {0, 3, 4}},
			"Path graph",
		},
		{
			4,
			[][]int{{0, 1, 1}, {0, 2, 2}, {0, 3, 3}, {1, 2, 4}},
			"Star graph",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  N: %d, Edges: %v\n", tc.n, tc.edges)
		
		result1 := findCriticalAndPseudoCriticalEdges(tc.n, tc.edges)
		result2 := findCriticalAndPseudoCriticalEdgesOptimized(tc.n, tc.edges)
		result3 := findCriticalAndPseudoCriticalEdgesUnionFind(tc.n, tc.edges)
		
		fmt.Printf("  Standard: %v\n", result1)
		fmt.Printf("  Optimized: %v\n", result2)
		fmt.Printf("  Union-Find: %v\n\n", result3)
	}
	
	// Test detailed analysis
	fmt.Println("=== Detailed Analysis Test ===")
	testN, testEdges := 4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}}
	result, analysis := findCriticalAndPseudoCriticalEdgesDetailed(testN, testEdges)
	
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Analysis: %v\n", analysis)
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	// Generate large graph
	largeN := 50
	largeEdges := make([][]int, 0)
	for i := 0; i < largeN; i++ {
		for j := i + 1; j < largeN && j < i+10; j++ {
			weight := (j - i) * 5
			largeEdges = append(largeEdges, []int{i, j, weight})
		}
	}
	
	fmt.Printf("Large test with %d nodes and %d edges\n", largeN, len(largeEdges))
	
	result := findCriticalAndPseudoCriticalEdgesOptimized(largeN, largeEdges)
	fmt.Printf("Result length: %d\n", len(result))
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Single node
	fmt.Printf("Single node: %v\n", findCriticalAndPseudoCriticalEdges(1, [][]int{}))
	
	// No edges
	fmt.Printf("No edges: %v\n", findCriticalAndPseudoCriticalEdges(3, [][]int{}))
	
	// Disconnected graph
	disconnected := [][]int{{0, 1, 1}, {2, 3, 1}}
	fmt.Printf("Disconnected: %v\n", findCriticalAndPseudoCriticalEdges(4, disconnected))
	
	// Multiple edges between same nodes
	multiEdges := [][]int{{0, 1, 1}, {0, 1, 2}, {0, 1, 3}, {1, 2, 1}}
	fmt.Printf("Multiple edges: %v\n", findCriticalAndPseudoCriticalEdges(3, multiEdges))
	
	// Test with different edge weights
	fmt.Println("\n=== Different Edge Weights Test ===")
	
	weightTest := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 10}, {0, 3, 5}}
	result = findCriticalAndPseudoCriticalEdges(4, weightTest)
	fmt.Printf("Mixed weights: %v\n", result)
	
	// Test with uniform weights
	uniform := [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}}
	result = findCriticalAndPseudoCriticalEdges(4, uniform)
	fmt.Printf("Uniform weights: %v\n", result)
}
