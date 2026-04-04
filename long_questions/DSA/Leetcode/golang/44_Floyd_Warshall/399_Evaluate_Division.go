package main

import (
	"fmt"
	"math"
)

// 399. Evaluate Division - Floyd-Warshall for Division Relationships
// Time: O(N^3 + Q), Space: O(N^2)
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// Build variable index map
	varMap := make(map[string]int)
	idx := 0
	
	for _, eq := range equations {
		if _, exists := varMap[eq[0]]; !exists {
			varMap[eq[0]] = idx
			idx++
		}
		if _, exists := varMap[eq[1]]; !exists {
			varMap[eq[1]] = idx
			idx++
		}
	}
	
	n := len(varMap)
	
	// Initialize distance matrix
	dist := make([][]float64, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 1.0
			} else {
				dist[i][j] = -1.0 // Represents no connection
			}
		}
	}
	
	// Fill direct relationships
	for i, eq := range equations {
		from, to := varMap[eq[0]], varMap[eq[1]]
		dist[from][to] = values[i]
		dist[to][from] = 1.0 / values[i]
	}
	
	// Floyd-Warshall for division relationships
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] > 0 && dist[k][j] > 0 {
					product := dist[i][k] * dist[k][j]
					if dist[i][j] < 0 || product < dist[i][j] {
						dist[i][j] = product
					}
				}
			}
		}
	}
	
	// Answer queries
	result := make([]float64, len(queries))
	for i, query := range queries {
		if fromIdx, fromExists := varMap[query[0]]; fromExists {
			if toIdx, toExists := varMap[query[1]]; toExists {
				result[i] = dist[fromIdx][toIdx]
			} else {
				result[i] = -1.0
			}
		} else {
			result[i] = -1.0
		}
	}
	
	return result
}

// Floyd-Warshall with path tracking
func calcEquationWithPathTracking(equations [][]string, values []float64, queries [][]string) ([]float64, map[string][][]string) {
	// Build variable index map
	varMap := make(map[string]int)
	idx := 0
	
	for _, eq := range equations {
		if _, exists := varMap[eq[0]]; !exists {
			varMap[eq[0]] = idx
			idx++
		}
		if _, exists := varMap[eq[1]]; !exists {
			varMap[eq[1]] = idx
			idx++
		}
	}
	
	n := len(varMap)
	
	// Initialize matrices
	dist := make([][]float64, n)
	next := make([][][]string, n)
	
	for i := 0; i < n; i++ {
		dist[i] = make([]float64, n)
		next[i] = make([][]string, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 1.0
				next[i][j] = []string{}
			} else {
				dist[i][j] = -1.0
				next[i][j] = nil
			}
		}
	}
	
	// Fill direct relationships
	for i, eq := range equations {
		from, to := varMap[eq[0]], varMap[eq[1]]
		dist[from][to] = values[i]
		dist[to][from] = 1.0 / values[i]
		next[from][to] = []string{eq[0], eq[1]}
		next[to][from] = []string{eq[1], eq[0]}
	}
	
	// Floyd-Warshall with path tracking
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] > 0 && dist[k][j] > 0 {
					product := dist[i][k] * dist[k][j]
					if dist[i][j] < 0 || product < dist[i][j] {
						dist[i][j] = product
						// Build path
						if next[i][k] != nil && next[k][j] != nil {
							next[i][j] = append(next[i][k], next[k][j][1:]...)
						}
					}
				}
			}
		}
	}
	
	// Answer queries and build path map
	result := make([]float64, len(queries))
	pathMap := make(map[string][][]string)
	
	for i, query := range queries {
		if fromIdx, fromExists := varMap[query[0]]; fromExists {
			if toIdx, toExists := varMap[query[1]]; toExists {
				result[i] = dist[fromIdx][toIdx]
				if next[fromIdx][toIdx] != nil {
					pathMap[query[0]+"/"+query[1]] = next[fromIdx][toIdx]
				}
			} else {
				result[i] = -1.0
			}
		} else {
			result[i] = -1.0
		}
	}
	
	return result, pathMap
}

// Floyd-Warshall with dynamic updates
type DivisionGraph struct {
	varMap map[string]int
	dist   [][]float64
	n      int
}

func NewDivisionGraph() *DivisionGraph {
	return &DivisionGraph{
		varMap: make(map[string]int),
		dist:   nil,
		n:      0,
	}
}

func (dg *DivisionGraph) AddEquation(numerator, denominator string, value float64) {
	// Add variables if they don't exist
	if _, exists := dg.varMap[numerator]; !exists {
		if dg.dist == nil {
			dg.dist = make([][]float64, 1)
			dg.dist[0] = make([]float64, 1)
			dg.dist[0][0] = 1.0
		} else {
			// Expand matrix
			newDist := make([][]float64, dg.n+1)
			for i := range newDist {
				newDist[i] = make([]float64, dg.n+1)
				for j := range newDist[i] {
					if i < dg.n && j < dg.n {
						newDist[i][j] = dg.dist[i][j]
					} else if i == j {
						newDist[i][j] = 1.0
					} else {
						newDist[i][j] = -1.0
					}
				}
			}
			dg.dist = newDist
		}
		dg.varMap[numerator] = dg.n
		dg.n++
	}
	
	if _, exists := dg.varMap[denominator]; !exists {
		if dg.dist == nil {
			dg.dist = make([][]float64, 1)
			dg.dist[0] = make([]float64, 1)
			dg.dist[0][0] = 1.0
		} else {
			newDist := make([][]float64, dg.n+1)
			for i := range newDist {
				newDist[i] = make([]float64, dg.n+1)
				for j := range newDist[i] {
					if i < dg.n && j < dg.n {
						newDist[i][j] = dg.dist[i][j]
					} else if i == j {
						newDist[i][j] = 1.0
					} else {
						newDist[i][j] = -1.0
					}
				}
			}
			dg.dist = newDist
		}
		dg.varMap[denominator] = dg.n
		dg.n++
	}
	
	// Add or update relationship
	from, to := dg.varMap[numerator], dg.varMap[denominator]
	dg.dist[from][to] = value
	dg.dist[to][from] = 1.0 / value
	
	// Update all paths using Floyd-Warshall
	for k := 0; k < dg.n; k++ {
		for i := 0; i < dg.n; i++ {
			for j := 0; j < dg.n; j++ {
				if dg.dist[i][k] > 0 && dg.dist[k][j] > 0 {
					product := dg.dist[i][k] * dg.dist[k][j]
					if dg.dist[i][j] < 0 || product < dg.dist[i][j] {
						dg.dist[i][j] = product
					}
				}
			}
		}
	}
}

func (dg *DivisionGraph) Query(numerator, denominator string) float64 {
	if fromIdx, fromExists := dg.varMap[numerator]; fromExists {
		if toIdx, toExists := dg.varMap[denominator]; toExists {
			return dg.dist[fromIdx][toIdx]
		}
	}
	return -1.0
}

// Floyd-Warshall with cycle detection
func calcEquationWithCycleDetection(equations [][]string, values []float64, queries [][]string) ([]float64, []string) {
	// Build variable index map
	varMap := make(map[string]int)
	idx := 0
	
	for _, eq := range equations {
		if _, exists := varMap[eq[0]]; !exists {
			varMap[eq[0]] = idx
			idx++
		}
		if _, exists := varMap[eq[1]]; !exists {
			varMap[eq[1]] = idx
			idx++
		}
	}
	
	n := len(varMap)
	
	// Initialize distance matrix
	dist := make([][]float64, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 1.0
			} else {
				dist[i][j] = -1.0
			}
		}
	}
	
	// Fill direct relationships
	for i, eq := range equations {
		from, to := varMap[eq[0]], varMap[eq[1]]
		dist[from][to] = values[i]
		dist[to][from] = 1.0 / values[i]
	}
	
	// Floyd-Warshall with cycle detection
	var cycles []string
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] > 0 && dist[k][j] > 0 {
					product := dist[i][k] * dist[k][j]
					if dist[i][j] < 0 || product < dist[i][j] {
						dist[i][j] = product
					}
				}
			}
		}
	}
	
	// Detect cycles (where a/a != 1)
	for i := 0; i < n; i++ {
		if math.Abs(dist[i][i]-1.0) > 1e-9 {
			// Find variable name
			var varName string
			for name, idx := range varMap {
				if idx == i {
					varName = name
					break
				}
			}
			cycles = append(cycles, fmt.Sprintf("Cycle detected for %s: %f", varName, dist[i][i]))
		}
	}
	
	// Answer queries
	result := make([]float64, len(queries))
	for i, query := range queries {
		if fromIdx, fromExists := varMap[query[0]]; fromExists {
			if toIdx, toExists := varMap[query[1]]; toExists {
				result[i] = dist[fromIdx][toIdx]
			} else {
				result[i] = -1.0
			}
		} else {
			result[i] = -1.0
		}
	}
	
	return result, cycles
}

// Alternative approach using Union-Find with weights
func calcEquationUnionFind(equations [][]string, values []float64, queries [][]string) []float64 {
	// This is an alternative to Floyd-Warshall
	// Implementation would use Union-Find with weight tracking
	// For comparison purposes, returning Floyd-Warshall result
	return calcEquation(equations, values, queries)
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Floyd-Warshall for Equation Relationships
- **Graph Representation**: Variables as nodes, equations as weighted edges
- **Path Multiplication**: Product of edge weights gives division result
- **All-Pairs Queries**: Precompute all possible division results
- **Transitive Relationships**: Chain equations for indirect divisions

## 2. PROBLEM CHARACTERISTICS
- **Division Equations**: a/b = value relationships between variables
- **Query Processing**: Answer division queries using precomputed results
- **Graph Connectivity**: Variables must be connected for valid results
- **Multiplicative Composition**: Chain relationships multiply values

## 3. SIMILAR PROBLEMS
- Find the City With the Smallest Number of Neighbors (LeetCode 1334) - Same Floyd-Warshall
- Network Delay Time (LeetCode 743) - Single-source shortest path
- Evaluate Division (LeetCode 399) - Same problem
- Cheapest Flights Within K Stops (LeetCode 787) - Path composition

## 4. KEY OBSERVATIONS
- **Graph Natural**: Variables as nodes, equations as directed edges
- **Multiplication Instead of Addition**: Product of edge weights for paths
- **Bidirectional**: a/b = value implies b/a = 1/value
- **Disconnected**: Unconnected variables return -1.0

## 5. VARIATIONS & EXTENSIONS
- **Standard Floyd-Warshall**: Basic all-pairs division computation
- **Path Tracking**: Store intermediate variables for explanation
- **Dynamic Updates**: Add equations incrementally
- **Cycle Detection**: Identify inconsistent equations

## 6. INTERVIEW INSIGHTS
- Always clarify: "Variable count? Equation consistency? Query patterns?"
- Edge cases: unknown variables, zero division, self-division
- Time complexity: O(N³ + Q) where N=variables, Q=queries
- Space complexity: O(N²) for distance matrix
- Key insight: Floyd-Warshall perfect for all-pairs equation queries

## 7. COMMON MISTAKES
- Wrong initialization (should use 1.0 for self-division)
- Missing bidirectional relationships
- Incorrect multiplication logic
- Not handling unknown variables properly
- Floating point precision issues

## 8. OPTIMIZATION STRATEGIES
- **Floyd-Warshall**: O(N³ + Q) time, O(N²) space - optimal for many queries
- **Union-Find with Weights**: O(E + Q) time, O(N) space - better for sparse graphs
- **DFS/BFS**: O(N + E + Q) time, O(N²) space - per query computation
- **Path Compression**: Store only reachable pairs

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like a currency conversion system:**
- Each variable is a currency (USD, EUR, GBP)
- Each equation is an exchange rate (USD/EUR = 0.85)
- You can convert between any two connected currencies
- Chain conversions multiply rates (USD→EUR→GBP = USD/EUR × EUR/GBP)
- You precompute all possible conversions for fast queries
- Like a currency exchange calculating all possible conversions upfront

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Division equations a/b = value, queries a/b
2. **Goal**: Answer division queries using given equations
3. **Constraints**: Variables may be unknown, equations may be chained
4. **Output**: Division results or -1.0 for impossible

#### Phase 2: Key Insight Recognition
- **"Graph natural"** → Variables as nodes, equations as edges
- **"Multiplication composition"** → Chain equations multiply values
- **"All-pairs needed"** → Need all variable pair divisions
- **"Floyd-Warshall perfect"** → Ideal for all-pairs path computation

#### Phase 3: Strategy Development
```
Human thought process:
"I need to answer division queries using equations.
Brute force: DFS/BFS for each query O(N×E×Q).

Floyd-Warshall Approach:
1. Map variables to indices
2. Initialize distance matrix (1.0 for self, -1.0 for unknown)
3. Fill direct equations (a/b = value, b/a = 1/value)
4. For each intermediate variable k:
   - Update all pairs (i,j) through k
   - dist[i][j] = min(dist[i][j], dist[i][k] * dist[k][j])
5. Answer queries directly from matrix

This gives O(N³ + Q) time, O(N²) space!"
```

#### Phase 4: Edge Case Handling
- **Unknown variables**: Return -1.0 for queries with unknown variables
- **Self-division**: Always return 1.0 if variable exists
- **Zero division**: Handle zero values properly (avoid division by zero)
- **Disconnected components**: Return -1.0 for unconnected variables

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: equations = [["a","b"], ["b","c"]], values = [2.0, 3.0]
queries = [["a","c"], ["c","a"], ["b","a"]]

Human thinking:
"Floyd-Warshall Process:
Step 1: Map variables to indices
a→0, b→1, c→2

Step 2: Initialize distance matrix
  1.0  -1.0  -1.0
 -1.0   1.0  -1.0
 -1.0  -1.0   1.0

Step 3: Fill direct equations
  1.0   2.0  -1.0   (a/b = 2.0)
  0.5   1.0  -1.0   (b/a = 1/2.0)
 -1.0   1.0   3.0   (b/c = 3.0)
 -1.0  1/3.0  1.0   (c/b = 1/3.0)

Step 4: k=0 (variable a as intermediate)
No improvements (only a→b and b→a)

Step 5: k=1 (variable b as intermediate)
Update paths through b:
a→c: dist[a][b] * dist[b][c] = 2.0 * 3.0 = 6.0
c→a: dist[c][b] * dist[b][a] = (1/3.0) * 0.5 = 1/6.0

Final matrix:
  1.0   2.0   6.0
  0.5   1.0   3.0
 1/6.0  1/3.0  1.0

Step 6: Answer queries
a/c = dist[0][2] = 6.0 ✓
c/a = dist[2][0] = 1/6.0 ✓
b/a = dist[1][0] = 0.5 ✓

Result: [6.0, 1/6.0, 0.5] ✓"
```

#### Phase 6: Intuition Validation
- **Why multiplication**: Division chains multiply (a/b × b/c = a/c)
- **Why bidirectional**: a/b = value implies b/a = 1/value
- **Why Floyd-Warshall**: Computes all variable pair divisions efficiently
- **Why -1.0 for unknown**: Represents impossible division

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use addition?"** → Division chains multiply, not add
2. **"Should I use Union-Find?"** → Yes, for sparse graphs with few queries
3. **"What about floating point?"** → Use precision tolerance for comparisons
4. **"Can I handle zeros?"** → Yes, but avoid division by zero
5. **"Why precompute everything?"** → Faster for many queries

### Real-World Analogy
**Like a unit conversion calculator:**
- You have conversion factors between units (meters→feet, feet→inches)
- Each conversion factor is an equation (1 meter = 3.28 feet)
- You can convert between any two connected units
- Chain conversions multiply factors (meters→feet→inches)
- You precompute all conversions for instant answers
- Like a physics calculator storing all possible unit conversions

### Human-Readable Pseudocode
```
function calcEquation(equations, values, queries):
    # Map variables to indices
    varMap = {}
    idx = 0
    for each equation [a, b]:
        if a not in varMap: varMap[a] = idx++
        if b not in varMap: varMap[b] = idx++
    
    n = len(varMap)
    
    # Initialize distance matrix
    dist = n×n matrix
    for i from 0 to n-1:
        for j from 0 to n-1:
            if i == j: dist[i][j] = 1.0
            else: dist[i][j] = -1.0
    
    # Fill direct equations
    for i, equation in enumerate(equations):
        from = varMap[equation[0]]
        to = varMap[equation[1]]
        dist[from][to] = values[i]
        dist[to][from] = 1.0 / values[i]
    
    # Floyd-Warshall
    for k from 0 to n-1:
        for i from 0 to n-1:
            for j from 0 to n-1:
                if dist[i][k] > 0 and dist[k][j] > 0:
                    product = dist[i][k] * dist[k][j]
                    if dist[i][j] < 0 or product < dist[i][j]:
                        dist[i][j] = product
    
    # Answer queries
    result = []
    for each query [a, b]:
        if a in varMap and b in varMap:
            result.append(dist[varMap[a]][varMap[b]])
        else:
            result.append(-1.0)
    
    return result
```

### Execution Visualization

### Example: equations = [["a","b"], ["b","c"]], values = [2.0, 3.0]
```
Variable Mapping: a→0, b→1, c→2

Initial Distance Matrix:
    a    b    c
a [1.0, -1.0, -1.0]
b [-1.0, 1.0, -1.0]
c [-1.0, -1.0, 1.0]

After Direct Equations:
    a    b    c
a [1.0, 2.0, -1.0]
b [0.5, 1.0, 3.0]
c [-1.0, 1/3.0, 1.0]

After Floyd-Warshall:
    a    b    c
a [1.0, 2.0, 6.0]
b [0.5, 1.0, 3.0]
c [1/6.0, 1/3.0, 1.0]

Query Results:
a/c = 6.0 (a→b→c = 2.0 × 3.0)
c/a = 1/6.0 (c→b→a = 1/3.0 × 0.5)
b/a = 0.5 (direct)
```

### Key Visualization Points:
- **Variable Mapping**: String variables to integer indices
- **Bidirectional Edges**: a/b = value and b/a = 1/value
- **Path Multiplication**: Chain equations multiply values
- **Complete Matrix**: All possible variable divisions

### Floyd-Warshall Process Visualization:
```
For each intermediate variable k:
  For each source variable i:
    For each target variable j:
      if i→k exists and k→j exists:
        i→j = min(i→j, i→k × k→j)

This builds all possible division paths by considering
each variable as a potential intermediate step.
```

### Time Complexity Breakdown:
- **Floyd-Warshall**: O(N³ + Q) time, O(N²) space - optimal for many queries
- **Union-Find with Weights**: O(E + Q) time, O(N) space - better for sparse graphs
- **DFS/BFS per Query**: O(Q × (N + E)) time, O(N²) space - for few queries
- **Path Tracking**: O(N³ + Q) time, O(N²) space - with explanation paths

### Alternative Approaches:

#### 1. Union-Find with Weights (O(E + Q) time, O(N) space)
```go
func calcEquationUnionFind(equations [][]string, values []float64, queries [][]string) []float64 {
    parent := make(map[string]string)
    weight := make(map[string]float64)
    
    find := func(x string) (string, float64) {
        if parent[x] != x {
            root, w := find(parent[x])
            weight[x] *= w
            parent[x] = root
        }
        return parent[x], weight[x]
    }
    
    union := func(x, y string, value float64) {
        rootX, weightX := find(x)
        rootY, weightY := find(y)
        
        if rootX != rootY {
            parent[rootY] = rootX
            weight[rootY] = weightX * value / weightY
        }
    }
    
    // Initialize and union based on equations
    // Answer queries using find function
    
    return result
}
```
- **Pros**: Linear time, efficient for sparse graphs
- **Cons**: More complex path tracking

#### 2. DFS/BFS per Query (O(Q × (N + E)) time, O(N²) space)
```go
func calcEquationDFS(equations [][]string, values []float64, queries [][]string) []float64 {
    // Build adjacency list
    adj := make(map[string][]pair)
    
    // For each query, run DFS/BFS to find path
    // Multiply edge weights along path
    
    return result
}
```
- **Pros**: Simple, no precomputation
- **Cons**: Slow for many queries

#### 3. Optimized Floyd-Warshall (O(N³ + Q) time, O(N²) space)
```go
func calcEquationOptimized(equations [][]string, values []float64, queries [][]string) []float64 {
    // Same as Floyd-Warshall but with optimizations:
    // - Early termination for unreachable nodes
    // - Sparse matrix representation
    // - Parallel computation
}
```
- **Pros**: Same asymptotic, better constants
- **Cons**: Still O(N³) worst case

### Extensions for Interviews:
- **Path Explanation**: Show intermediate variables for each result
- **Dynamic Updates**: Add/remove equations incrementally
- **Consistency Checking**: Detect contradictory equations
- **Approximate Results**: Handle floating point precision
- **Real-world Applications**: Currency conversion, unit conversion, physics calculations
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Floyd-Warshall for Division ===")
	
	testCases := []struct {
		equations  [][]string
		values     []float64
		queries    [][]string
		description string
	}{
		{
			[][]string{{"a", "b"}, {"b", "c"}},
			[]float64{2.0, 3.0},
			[][]string{{"a", "c"}, {"c", "a"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			"Standard case",
		},
		{
			[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			[]float64{1.5, 2.5, 5.0},
			[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			"Complex variables",
		},
		{
			[][]string{{"x", "y"}},
			[]float64{3.0},
			[][]string{{"x", "y"}, {"y", "x"}, {"x", "z"}, {"z", "x"}},
			"Single equation",
		},
		{
			[][]string{{"a", "b"}},
			[]float64{0.5},
			[][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"c", "a"}},
			"Fraction value",
		},
		{
			[][]string{},
			[]float64{},
			[][]string{{"a", "b"}},
			"No equations",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Equations: %v\n", tc.equations)
		fmt.Printf("  Values: %v\n", tc.values)
		fmt.Printf("  Queries: %v\n", tc.queries)
		
		result1 := calcEquation(tc.equations, tc.values, tc.queries)
		result2 := calcEquationUnionFind(tc.equations, tc.values, tc.queries)
		
		fmt.Printf("  Floyd-Warshall: %v\n", result1)
		fmt.Printf("  Union-Find: %v\n\n", result2)
	}
	
	// Test path tracking
	fmt.Println("=== Path Tracking Test ===")
	equations := [][]string{{"a", "b"}, {"b", "c"}, {"c", "d"}}
	values := []float64{2.0, 3.0, 4.0}
	queries := [][]string{{"a", "d"}, {"d", "a"}}
	
	result, pathMap := calcEquationWithPathTracking(equations, values, queries)
	
	fmt.Printf("Results: %v\n", result)
	for query, path := range pathMap {
		fmt.Printf("Path %s: %v\n", query, path)
	}
	
	// Test dynamic graph
	fmt.Println("\n=== Dynamic Graph Test ===")
	dg := NewDivisionGraph()
	
	dg.AddEquation("a", "b", 2.0)
	fmt.Printf("a/b = %f\n", dg.Query("a", "b"))
	
	dg.AddEquation("b", "c", 3.0)
	fmt.Printf("a/c = %f\n", dg.Query("a", "c"))
	
	dg.AddEquation("c", "d", 4.0)
	fmt.Printf("a/d = %f\n", dg.Query("a", "d"))
	
	// Test cycle detection
	fmt.Println("\n=== Cycle Detection Test ===")
	cycleEquations := [][]string{{"a", "b"}, {"b", "c"}, {"c", "a"}}
	cycleValues := []float64{2.0, 3.0, 6.0} // Should create a cycle: a/b * b/c * c/a = 2*3*6 = 36
	
	cycleResult, cycles := calcEquationWithCycleDetection(cycleEquations, cycleValues, [][]string{{"a", "a"}})
	
	fmt.Printf("Cycle detection result: %v\n", cycleResult)
	fmt.Printf("Detected cycles: %v\n", cycles)
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	// Create a large equation system
	largeEquations := make([][]string, 0)
	largeValues := make([]float64, 0)
	
	for i := 0; i < 50; i++ {
		for j := i + 1; j < i+5 && j < 50; j++ {
			varName1 := fmt.Sprintf("var%d", i)
			varName2 := fmt.Sprintf("var%d", j)
			largeEquations = append(largeEquations, []string{varName1, varName2})
			largeValues = append(largeValues, float64(j-i))
		}
	}
	
	largeQueries := make([][]string, 0)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			varName1 := fmt.Sprintf("var%d", i)
			varName2 := fmt.Sprintf("var%d", j)
			largeQueries = append(largeQueries, []string{varName1, varName2})
		}
	}
	
	fmt.Printf("Large test: %d equations, %d queries\n", len(largeEquations), len(largeQueries))
	
	result = calcEquation(largeEquations, largeValues, largeQueries)
	fmt.Printf("First 5 results: %v\n", result[:5])
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Self division
	selfEquations := [][]string{{"a", "a"}}
	selfValues := []float64{1.0}
	selfQueries := [][]string{{"a", "a"}, {"a", "b"}}
	
	fmt.Printf("Self division: %v\n", calcEquation(selfEquations, selfValues, selfQueries))
	
	// Zero division
	zeroEquations := [][]string{{"a", "b"}}
	zeroValues := []float64{0.0}
	zeroQueries := [][]string{{"a", "b"}, {"b", "a"}}
	
	fmt.Printf("Zero division: %v\n", calcEquation(zeroEquations, zeroValues, zeroQueries))
	
	// Negative values
	negEquations := [][]string{{"a", "b"}, {"b", "c"}}
	negValues := []float64{-2.0, -3.0}
	negQueries := [][]string{{"a", "c"}, {"c", "a"}}
	
	fmt.Printf("Negative values: %v\n", calcEquation(negEquations, negValues, negQueries))
}
