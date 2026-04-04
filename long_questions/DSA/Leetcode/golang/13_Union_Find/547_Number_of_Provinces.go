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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Union Find for Connected Components
- **Disjoint Set Union**: Track connected components in graph
- **Path Compression**: Optimize find operation with flattening
- **Union by Rank**: Optimize union operation with tree height balancing
- **Component Counting**: Count unique roots to find number of provinces

## 2. PROBLEM CHARACTERISTICS
- **Adjacency Matrix**: Graph represented as N×N connectivity matrix
- **Undirected Graph**: Connections are bidirectional
- **Connected Components**: Each province is a connected component
- **Symmetric Matrix**: Matrix[i][j] = Matrix[j][i] for undirected graph

## 3. SIMILAR PROBLEMS
- Number of Connected Components (LeetCode 323)
- Friend Circles (LeetCode 547 - same problem)
- Graph Valid Tree (LeetCode 261)
- Accounts Merge (LeetCode 721)

## 4. KEY OBSERVATIONS
- **Matrix representation**: isConnected[i][j] = 1 means direct connection
- **Union operations**: Union all directly connected cities
- **Root counting**: Number of unique roots = number of provinces
- **Path compression**: Essential for performance optimization

## 5. VARIATIONS & EXTENSIONS
- **Directed Graph**: Handle one-way connections
- **Weighted Graph**: Consider edge weights
- **Dynamic Updates**: Handle adding/removing connections
- **Multiple Queries**: Answer connectivity queries efficiently

## 6. INTERVIEW INSIGHTS
- Always clarify: "Is the graph directed or undirected?"
- Edge cases: empty matrix, single city, all connected, all disconnected
- Time complexity: O(N²) for processing matrix, O(N α(N)) for Union Find
- Space complexity: O(N) for Union Find data structures

## 7. COMMON MISTAKES
- Not processing entire matrix (only upper triangle)
- Forgetting path compression in find operation
- Not using union by rank optimization
- Counting provinces before all unions processed
- Not handling empty matrix case

## 8. OPTIMIZATION STRATEGIES
- **Upper triangle only**: Process only i<j since matrix is symmetric
- **Path compression**: Essential for near-constant time operations
- **Union by rank**: Balances tree height for optimal performance
- **Early termination**: Not applicable (need to process all connections)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding friend groups in a social network:**
- You have a social network where connections are mutual (undirected)
- Each person is connected to certain others (adjacency matrix)
- People in the same friend group can reach each other through connections
- You want to count how many distinct friend groups exist
- Merge people who are directly connected, then count the groups

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: N×N adjacency matrix representing city connections
2. **Goal**: Count number of connected components (provinces)
3. **Output**: Integer count of provinces
4. **Constraint**: Matrix is symmetric, diagonal entries are 1

#### Phase 2: Key Insight Recognition
- **"Connected components"** → Each province is a connected component
- **"Union Find natural fit"** → Efficiently track and merge connected cities
- **"Matrix processing"** → Need to process all connections in matrix
- **"Root counting"** → Number of unique roots gives province count

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find connected components in a graph.
The adjacency matrix shows which cities are directly connected.
I'll use Union Find to merge connected cities.
For each connection matrix[i][j] = 1, I'll union city i and j.
Finally, I'll count how many unique roots remain."
```

#### Phase 4: Edge Case Handling
- **Empty matrix**: Return 0
- **Single city**: Return 1
- **All connected**: Return 1
- **All disconnected**: Return N

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Matrix:
[1 1 0]
[1 1 0]  
[0 0 1]

Human thinking:
"I'll initialize Union Find with 3 cities: {0}, {1}, {2}

Process connections:
matrix[0][1] = 1 → union(0,1): {0,1}, {2}
matrix[0][2] = 0 → skip
matrix[1][2] = 0 → skip
matrix[1][0] = 1 → union(1,0): already connected
matrix[2][0] = 0 → skip
matrix[2][1] = 0 → skip

Final roots:
find(0) = 0
find(1) = 0  
find(2) = 2

Unique roots: {0, 2} → 2 provinces"
```

#### Phase 6: Intuition Validation
- **Why Union Find works**: Efficiently tracks dynamic connectivity
- **Why path compression**: Prevents O(N²) worst-case behavior
- **Why union by rank**: Keeps trees balanced for optimal performance
- **Why O(N²)**: Need to process N² matrix entries (though could optimize to N²/2)

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use DFS/BFS?"** → Can, but Union Find is more efficient for multiple queries
2. **"Should I process the whole matrix?"** → Yes, but could optimize to upper triangle only
3. **"What about very large N?"** → Union Find with optimizations handles large N well
4. **"Can I use adjacency list instead?"** → Yes, but input is matrix format

### Real-World Analogy
**Like finding communication networks in a company:**
- You have a company with departments (cities)
- Communication channels exist between certain departments (connections)
- Departments that can communicate (directly or indirectly) form a network
- You want to know how many independent communication networks exist
- Merge departments that can directly communicate, then count distinct networks

### Human-Readable Pseudocode
```
function findCircleNum(isConnected):
    n = length(isConnected)
    if n == 0:
        return 0
    
    // Initialize Union-Find
    parent = [0, 1, 2, ..., n-1]
    rank = [0, 0, 0, ..., 0]
    
    function find(x):
        if parent[x] != x:
            parent[x] = find(parent[x])  // Path compression
        return parent[x]
    
    function union(x, y):
        rootX = find(x)
        rootY = find(y)
        if rootX == rootY:
            return
        
        // Union by rank
        if rank[rootX] < rank[rootY]:
            parent[rootX] = rootY
        else if rank[rootX] > rank[rootY]:
            parent[rootY] = rootX
        else:
            parent[rootY] = rootX
            rank[rootX]++
    
    // Process all connections
    for i from 0 to n-1:
        for j from 0 to n-1:
            if isConnected[i][j] == 1:
                union(i, j)
    
    // Count unique provinces
    provinces = set()
    for i from 0 to n-1:
        provinces.add(find(i))
    
    return length(provinces)
```

### Execution Visualization

### Example Matrix:
```
[1 1 0]
[1 1 0]
[0 0 1]
```

### Union Find Process:
```
Initial: parent=[0,1,2], rank=[0,0,0]

Process (0,1): union(0,1)
- find(0)=0, find(1)=1
- rank[0]=rank[1], parent[1]=0, rank[0]=1
- parent=[0,0,2], rank=[1,0,0]

Process (0,2): matrix[0][2]=0 → skip
Process (1,0): union(1,0)  
- find(1)=0, find(0)=0 → already connected

Process (1,2): matrix[1][2]=0 → skip
Process (2,0): matrix[2][0]=0 → skip
Process (2,1): matrix[2][1]=0 → skip

Final roots:
find(0)=0, find(1)=0, find(2)=2

Unique roots: {0, 2} → 2 provinces
```

### Key Visualization Points:
- **Union operations**: Merge connected cities into same component
- **Path compression**: Flatten tree structure during find operations
- **Root counting**: Number of unique roots equals number of provinces
- **Matrix symmetry**: Could optimize to process only upper triangle

### Memory Layout Visualization:
```
Parent Array Evolution:
[0,1,2] → [0,0,2] (after union 0,1)

Rank Array Evolution:
[0,0,0] → [1,0,0] (after union 0,1)

Final Component Structure:
Component 0: {0, 1}
Component 2: {2}
```

### Time Complexity Breakdown:
- **Matrix processing**: O(N²) to check all entries
- **Union operations**: O(N² α(N)) where α is inverse Ackermann function
- **Find operations**: O(N α(N)) for final root counting
- **Total time**: O(N²) dominated by matrix processing
- **Space**: O(N) for Union Find arrays

### Alternative Approaches:

#### 1. DFS Approach (O(N²) time, O(N) space)
```go
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
```
- **Pros**: Simple to understand, no Union Find complexity
- **Cons**: Not as efficient for multiple connectivity queries

#### 2. BFS Approach (O(N²) time, O(N) space)
```go
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
```
- **Pros**: Iterative, no recursion stack
- **Cons**: Similar performance to DFS, queue management overhead

#### 3. Optimized Matrix Processing (O(N²) time, O(N) space)
```go
func findCircleNumOptimized(isConnected [][]int) int {
    n := len(isConnected)
    if n == 0 {
        return 0
    }
    
    parent := make([]int, n)
    rank := make([]int, n)
    
    for i := 0; i < n; i++ {
        parent[i] = i
        rank[i] = 0
    }
    
    var find func(int) int
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
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
        
        if rank[rootX] < rank[rootY] {
            parent[rootX] = rootY
        } else if rank[rootX] > rank[rootY] {
            parent[rootY] = rootX
        } else {
            parent[rootY] = rootX
            rank[rootX]++
        }
    }
    
    // Process only upper triangle since matrix is symmetric
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if isConnected[i][j] == 1 {
                union(i, j)
            }
        }
    }
    
    provinces := make(map[int]bool)
    for i := 0; i < n; i++ {
        provinces[find(i)] = true
    }
    
    return len(provinces)
}
```
- **Pros**: Reduces matrix processing by half
- **Cons**: Still O(N²) overall complexity

### Extensions for Interviews:
- **Directed Graph**: Handle one-way connections differently
- **Weighted Graph**: Consider edge weights in connectivity
- **Dynamic Updates**: Handle adding/removing connections efficiently
- **Multiple Queries**: Answer connectivity queries without reprocessing
- **Component Size**: Find size of each connected component
*/
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
