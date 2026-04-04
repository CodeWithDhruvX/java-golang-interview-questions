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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Union Find for Cycle Detection
- **Incremental Union**: Process edges one by one
- **Cycle Detection**: Edge connecting already connected nodes creates cycle
- **Early Termination**: Return first edge that creates cycle
- **Path Compression**: Optimize find operation for efficiency

## 2. PROBLEM CHARACTERISTICS
- **Tree Structure**: N nodes, N-1 edges for valid tree
- **Redundant Edge**: Adding one edge creates exactly one cycle
- **Undirected Graph**: Edges are bidirectional
- **First Occurrence**: Return the first edge that creates a cycle

## 3. SIMILAR PROBLEMS
- Redundant Connection II (LeetCode 685) - Directed graph
- Graph Valid Tree (LeetCode 261)
- Find Critical Connections (LeetCode 1192)
- Number of Operations to Make Network Connected (LeetCode 1319)

## 4. KEY OBSERVATIONS
- **Tree property**: Valid tree has N-1 edges and is acyclic
- **Cycle detection**: Union nodes, if already connected → cycle
- **First redundant**: Process edges in given order, return first cycle-causing edge
- **Union Find efficiency**: Near-constant time operations with optimizations

## 5. VARIATIONS & EXTENSIONS
- **Directed Graph**: Handle directed edges (LeetCode 685)
- **Multiple Cycles**: Find all redundant edges
- **Edge Removal**: Remove edge to make graph a tree
- **Dynamic Updates**: Handle adding/removing edges

## 6. INTERVIEW INSIGHTS
- Always clarify: "Is the graph directed or undirected?"
- Edge cases: single edge, already tree, multiple cycles
- Time complexity: O(N α(N)) where α is inverse Ackermann function
- Space complexity: O(N) for Union Find data structures

## 7. COMMON MISTAKES
- Not using path compression (causes poor performance)
- Not using union by rank optimization
- Processing edges out of given order
- Not handling 1-indexed node numbering correctly
- Forgetting to return first redundant edge

## 8. OPTIMIZATION STRATEGIES
- **Path compression**: Essential for near-constant time operations
- **Union by rank**: Balances tree height for optimal performance
- **Early termination**: Stop as soon as redundant edge found
- **Simple Union**: Can skip rank for simplicity (still O(N α(N)))

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like building a network and detecting the first redundant connection:**
- You have N computers that need to be connected in a tree structure
- You're given N connections (edges) one by one
- A valid tree needs exactly N-1 connections
- The extra connection will create a cycle (redundant)
- You want to find the first connection that creates a cycle

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: N edges forming a graph with one redundant edge
2. **Goal**: Find the edge that creates a cycle
3. **Output**: The first edge that makes the graph cyclic
4. **Constraint**: Return edge in the order it was given

#### Phase 2: Key Insight Recognition
- **"Tree property"** → Valid tree has no cycles, N nodes need N-1 edges
- **"Union Find natural fit"** → Efficiently track connectivity as edges added
- **"Cycle detection"** → Edge connecting already connected nodes creates cycle
- **"First occurrence"** → Process edges in order, return first cycle-causing edge

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the edge that creates a cycle.
I'll use Union Find to track connectivity as I add edges.
For each edge, I'll check if the nodes are already connected.
If they are, this edge creates a cycle and is redundant.
If not, I'll union the nodes and continue.
The first edge that creates a cycle is my answer."
```

#### Phase 4: Edge Case Handling
- **Single edge**: No cycle possible, but problem guarantees solution
- **Already tree**: No redundant edge, but problem guarantees one
- **Multiple cycles**: Return first edge that creates any cycle
- **Disconnected components**: Union handles this naturally

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Edges: [[1,2], [1,3], [2,3]]

Human thinking:
"I'll start with 3 separate components: {1}, {2}, {3}

Process edge [1,2]:
- find(1) = 1, find(2) = 2 (different components)
- union(1,2): {1,2}, {3}
- No cycle, continue

Process edge [1,3]:
- find(1) = 1, find(3) = 3 (different components)  
- union(1,3): {1,2,3}
- No cycle, continue

Process edge [2,3]:
- find(2) = 1, find(3) = 1 (same component!)
- This edge creates a cycle → REDUNDANT!
- Return [2,3]"
```

#### Phase 6: Intuition Validation
- **Why Union Find works**: Efficiently tracks dynamic connectivity
- **Why cycle detection works**: Edge between already connected nodes creates cycle
- **Why O(N α(N))**: Near-constant time operations with path compression
- **Why first edge**: Process in given order, return first cycle-causing edge

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use DFS?"** → Could, but Union Find is more efficient for incremental building
2. **"Should I check all edges?"** → No, stop at first redundant edge
3. **"What about directed edges?"** → Different problem (LeetCode 685)
4. **"Can I skip optimizations?"** → Path compression is essential for performance

### Real-World Analogy
**Like building a computer network and detecting redundant cables:**
- You have N computers that need to be connected
- You're adding network cables one by one
- A good network connects all computers with no redundant paths
- When you add a cable between already connected computers, it's redundant
- You want to identify the first redundant cable added

### Human-Readable Pseudocode
```
function findRedundantConnection(edges):
    n = length(edges)
    
    // Initialize Union-Find (1-indexed)
    parent = [0, 1, 2, ..., n]
    rank = [0, 0, 0, ..., 0]
    
    function find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])  // Path compression
        return parent[x]
    
    function union(x, y):
        rootX = find(x)
        rootY = find(y)
        
        if rootX == rootY:
            return false  // Already connected, creates cycle
        else:
            // Union by rank
            if rank[rootX] < rank[rootY]:
                parent[rootX] = rootY
            else if rank[rootX] > rank[rootY]:
                parent[rootY] = rootX
            else:
                parent[rootY] = rootX
                rank[rootX]++
            return true
    
    for edge in edges:
        if not union(edge[0], edge[1]):
            return edge  // Found redundant connection
    
    return []  // Should never reach here given problem constraints
```

### Execution Visualization

### Example Edges: [[1,2], [1,3], [2,3]]
```
Initial: parent=[0,1,2,3], rank=[0,0,0,0]

Process [1,2]: union(1,2)
- find(1)=1, find(2)=2
- rank[1]=rank[2], parent[2]=1, rank[1]=1
- parent=[0,1,1,3], rank=[0,1,0,0]
- Success, continue

Process [1,3]: union(1,3)  
- find(1)=1, find(3)=3
- rank[1]=rank[3], parent[3]=1, rank[1]=2
- parent=[0,1,1,1], rank=[0,2,0,0]
- Success, continue

Process [2,3]: union(2,3)
- find(2)=1, find(3)=1
- Same root! Creates cycle → REDUNDANT
- Return [2,3]
```

### Key Visualization Points:
- **Incremental building**: Add edges one by one, track connectivity
- **Cycle detection**: Edge between nodes in same component creates cycle
- **Path compression**: Flatten tree structure during find operations
- **Early termination**: Stop at first redundant edge

### Memory Layout Visualization:
```
Component Evolution:
{1}, {2}, {3} → {1,2}, {3} → {1,2,3}

Parent Array Evolution:
[0,1,2,3] → [0,1,1,3] → [0,1,1,1]

Rank Array Evolution:
[0,0,0,0] → [0,1,0,0] → [0,2,0,0]

Cycle Detection:
Edge [2,3]: find(2)=1, find(3)=1 → Same root → Cycle!
```

### Time Complexity Breakdown:
- **Union operations**: O(N α(N)) where α is inverse Ackermann function
- **Find operations**: O(N α(N)) with path compression
- **Total time**: O(N α(N)) ≈ O(N) for practical purposes
- **Space**: O(N) for Union Find arrays

### Alternative Approaches:

#### 1. Simple Union Find (O(N α(N)) time, O(N) space)
```go
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
```
- **Pros**: Simpler implementation, still efficient
- **Cons**: No rank optimization, slightly less balanced trees

#### 2. DFS Cycle Detection (O(N²) time, O(N) space)
```go
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
```
- **Pros**: Conceptually simple cycle detection
- **Cons**: O(N²) time, complex edge removal logic

#### 3. Union Find with Size Tracking (O(N α(N)) time, O(N) space)
```go
func findRedundantConnectionSize(edges [][]int) []int {
    n := len(edges)
    parent := make([]int, n+1)
    size := make([]int, n+1)
    
    for i := 1; i <= n; i++ {
        parent[i] = i
        size[i] = 1
    }
    
    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    
    var union func(int, int) bool
    union = func(x, y int) bool {
        rootX := find(x)
        rootY := find(y)
        
        if rootX == rootY {
            return false
        }
        
        // Union by size
        if size[rootX] < size[rootY] {
            parent[rootX] = rootY
            size[rootY] += size[rootX]
        } else {
            parent[rootY] = rootX
            size[rootX] += size[rootY]
        }
        
        return true
    }
    
    for _, edge := range edges {
        if !union(edge[0], edge[1]) {
            return edge
        }
    }
    
    return []int{}
}
```
- **Pros**: Size-based union, alternative to rank
- **Cons**: Similar performance to rank-based approach

### Extensions for Interviews:
- **Directed Graph**: Handle directed edges (LeetCode 685)
- **Multiple Redundant Edges**: Find all redundant edges
- **Edge Removal**: Remove minimal edges to make graph a tree
- **Dynamic Updates**: Handle adding/removing edges efficiently
- **Cycle Count**: Count number of cycles in graph
*/
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
