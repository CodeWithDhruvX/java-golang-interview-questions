package main

import (
	"container/heap"
	"fmt"
	"math"
)

// 1514. Path with Maximum Probability
// Time: O((V + E) log V), Space: O(V + E)
func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Build adjacency list
	adj := make(map[int][]Edge)
	for i, edge := range edges {
		from, to := edge[0], edge[1]
		prob := succProb[i]
		adj[from] = append(adj[from], Edge{to, prob})
		adj[to] = append(adj[to], Edge{from, prob}) // Undirected graph
	}
	
	// Dijkstra's algorithm with max-heap (probability)
	prob := make([]float64, n)
	for i := 0; i < n; i++ {
		prob[i] = 0
	}
	prob[start] = 1
	
	// Max-heap: {probability, node}
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, Item{1.0, start})
	
	for maxHeap.Len() > 0 {
		current := heap.Pop(maxHeap).(Item)
		currentProb, currentNode := current.probability, current.node
		
		// Skip if we've found a better path
		if currentProb < prob[currentNode] {
			continue
		}
		
		// Early exit if we reached the target
		if currentNode == end {
			return currentProb
		}
		
		// Relax edges
		for _, edge := range adj[currentNode] {
			newProb := currentProb * edge.weight
			if newProb > prob[edge.to] {
				prob[edge.to] = newProb
				heap.Push(maxHeap, Item{newProb, edge.to})
			}
		}
	}
	
	return prob[end]
}

// Edge represents a weighted edge (probability)
type Edge struct {
	to     int
	weight float64
}

// Max-heap implementation for probability
type Item struct {
	probability float64
	node        int
}

type MaxHeap []Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].probability > h[j].probability }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Alternative implementation using BFS with priority queue
func maxProbabilityBFS(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Build adjacency list
	adj := make(map[int][]Edge)
	for i, edge := range edges {
		from, to := edge[0], edge[1]
		prob := succProb[i]
		adj[from] = append(adj[from], Edge{to, prob})
		adj[to] = append(adj[to], Edge{from, prob})
	}
	
	// Priority queue (max-heap)
	pq := &MaxHeap{}
	heap.Init(pq)
	heap.Push(pq, Item{1.0, start})
	
	// Track maximum probabilities
	maxProb := make([]float64, n)
	for i := 0; i < n; i++ {
		maxProb[i] = 0
	}
	maxProb[start] = 1
	
	for pq.Len() > 0 {
		current := heap.Pop(pq).(Item)
		prob, node := current.probability, current.node
		
		// Skip if we've found a better path
		if prob < maxProb[node] {
			continue
		}
		
		// Early exit
		if node == end {
			return prob
		}
		
		// Explore neighbors
		for _, edge := range adj[node] {
			newProb := prob * edge.weight
			if newProb > maxProb[edge.to] {
				maxProb[edge.to] = newProb
				heap.Push(pq, Item{newProb, edge.to})
			}
		}
	}
	
	return maxProb[end]
}

// Modified Dijkstra's algorithm with custom comparison
func maxProbabilityModifiedDijkstra(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Build adjacency list
	adj := make(map[int][]Edge)
	for i, edge := range edges {
		from, to := edge[0], edge[1]
		prob := succProb[i]
		adj[from] = append(adj[from], Edge{to, prob})
		adj[to] = append(adj[to], Edge{from, prob})
	}
	
	// Use negative probabilities with min-heap (convert to minimization)
	negProb := make([]float64, n)
	for i := 0; i < n; i++ {
		negProb[i] = math.Inf(1) // Infinity
	}
	negProb[start] = -1 // Negative of 1
	
	// Min-heap: {negative probability, node}
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, MinItem{-1.0, start})
	
	for minHeap.Len() > 0 {
		current := heap.Pop(minHeap).(MinItem)
		currentNegProb, currentNode := current.negProb, current.node
		
		// Skip if we've found a better path
		if currentNegProb > negProb[currentNode] {
			continue
		}
		
		// Early exit
		if currentNode == end {
			return -currentNegProb
		}
		
		// Relax edges
		for _, edge := range adj[currentNode] {
			newNegProb := currentNegProb * edge.weight
			if newNegProb > negProb[edge.to] {
				negProb[edge.to] = newNegProb
				heap.Push(minHeap, MinItem{newNegProb, edge.to})
			}
		}
	}
	
	return -negProb[end]
}

// Min-heap implementation for negative probabilities
type MinItem struct {
	negProb float64
	node    int
}

type MinHeap []MinItem

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].negProb < h[j].negProb }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(MinItem))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// BFS with logarithmic transformation
func maxProbabilityLogTransform(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// Build adjacency list
	adj := make(map[int][]Edge)
	for i, edge := range edges {
		from, to := edge[0], edge[1]
		prob := succProb[i]
		adj[from] = append(adj[from], Edge{to, prob})
		adj[to] = append(adj[to], Edge{from, prob})
	}
	
	// Use log probabilities (convert multiplication to addition)
	logProb := make([]float64, n)
	for i := 0; i < n; i++ {
		logProb[i] = math.Inf(-1) // -Infinity
	}
	logProb[start] = 0 // log(1) = 0
	
	// Max-heap: {log probability, node}
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	heap.Push(maxHeap, Item{0, start})
	
	for maxHeap.Len() > 0 {
		current := heap.Pop(maxHeap).(Item)
		currentLogProb, currentNode := current.probability, current.node
		
		// Skip if we've found a better path
		if currentLogProb < logProb[currentNode] {
			continue
		}
		
		// Early exit
		if currentNode == end {
			return math.Exp(currentLogProb)
		}
		
		// Relax edges
		for _, edge := range adj[currentNode] {
			newLogProb := currentLogProb + math.Log(edge.weight)
			if newLogProb > logProb[edge.to] {
				logProb[edge.to] = newLogProb
				heap.Push(maxHeap, Item{newLogProb, edge.to})
			}
		}
	}
	
	return math.Exp(logProb[end])
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Modified Dijkstra's Algorithm for Maximum Probability
- **Priority Queue**: Use max-heap instead of min-heap
- **Probability Multiplication**: Path probability = product of edge probabilities
- **Early Termination**: Stop when target node is reached
- **Logarithmic Transformation**: Convert multiplication to addition for numerical stability

## 2. PROBLEM CHARACTERISTICS
- **Weighted Graph**: Edges have success probabilities (0, 1]
- **Path Probability**: Product of probabilities along path
- **Maximization**: Find path with maximum probability
- **Undirected Graph**: Edges work in both directions

## 3. SIMILAR PROBLEMS
- Network Delay Time (LeetCode 743) - Standard shortest path
- Cheapest Flights Within K Stops (LeetCode 787) - Constrained shortest path
- Path with Minimum Effort (LeetCode 1631) - Minimax path
- Find the Maximum Probability Path (LeetCode 1514) - Same problem

## 4. KEY OBSERVATIONS
- **Dijkstra Adaptation**: Standard Dijkstra works for maximization with max-heap
- **Probability Properties**: Product of probabilities decreases with path length
- **Log Transformation**: log(a*b) = log(a) + log(b) for numerical stability
- **Early Exit**: First time we pop target from heap gives optimal solution

## 5. VARIATIONS & EXTENSIONS
- **Logarithmic Approach**: Convert to additive problem
- **Negative Probabilities**: Use min-heap with negative values
- **Path Counting**: Count number of maximum probability paths
- **Multiple Targets**: Find maximum probability to multiple destinations

## 6. INTERVIEW INSIGHTS
- Always clarify: "Graph size constraints? Probability precision requirements?"
- Edge cases: no path exists, start equals end, disconnected graph
- Time complexity: O((V + E) log V) with heap
- Space complexity: O(V + E) for adjacency and heap
- Key insight: Dijkstra works for any monotonic objective function

## 7. COMMON MISTAKES
- Using min-heap instead of max-heap
- Not handling floating-point precision issues
- Missing early termination optimization
- Incorrect probability multiplication logic
- Not handling disconnected graphs properly

## 8. OPTIMIZATION STRATEGIES
- **Standard Max-Heap**: O((V + E) log V) time, O(V + E) space - basic
- **Log Transform**: O((V + E) log V) time, O(V + E) space - numerical stability
- **Negative Values**: O((V + E) log V) time, O(V + E) space - min-heap trick
- **Early Termination**: O((V + E) log V) time, O(V + E) space - faster in practice

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the most reliable route in a network:**
- Each connection has a reliability score (probability of success)
- You want the route with highest overall reliability
- Overall reliability = product of individual connection reliabilities
- You explore routes in order of decreasing reliability
- Like finding the most dependable delivery route

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Graph edges with success probabilities, start and end nodes
2. **Goal**: Find path with maximum success probability
3. **Rules**: Path probability = product of edge probabilities
4. **Output**: Maximum probability value (0 if no path)

#### Phase 2: Key Insight Recognition
- **"Dijkstra adaptable"** → Standard Dijkstra works for any monotonic objective
- **"Max-heap natural"** → Want maximum probability, so use max-heap
- **"Product multiplication"** → Path probability = product of edge probabilities
- **"Log transformation possible"** → Convert multiplication to addition

#### Phase 3: Strategy Development
```
Human thought process:
"I need path with maximum probability.
Standard Dijkstra finds minimum distance.

Modified Dijkstra Approach:
1. Use max-heap instead of min-heap
2. Path probability = product of edge probabilities
3. Start with probability 1 at source
4. For each edge: newProb = currentProb * edgeProb
5. Keep maximum probability for each node
6. Early exit when target reached

This gives optimal max probability path!"
```

#### Phase 4: Edge Case Handling
- **No path**: Return 0 if end unreachable
- **Same node**: Return 1 (probability of staying at same node)
- **Disconnected graph**: Handle gracefully with 0 probability
- **Zero probabilities**: Handle edges with 0 probability

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: 3 nodes, edges: (0,1):0.5, (1,2):0.5, (0,2):0.2, start=0, end=2

Human thinking:
"Modified Dijkstra Approach:
Initialize: prob[0]=1, prob[1]=0, prob[2]=0
Max-heap: [(1.0, 0)]

Step 1: Pop (1.0, 0)
- Edge 0→1: newProb = 1.0 * 0.5 = 0.5, prob[1]=0.5, push (0.5, 1)
- Edge 0→2: newProb = 1.0 * 0.2 = 0.2, prob[2]=0.2, push (0.2, 2)
Heap: [(0.5, 1), (0.2, 2)]

Step 2: Pop (0.5, 1)
- Edge 1→2: newProb = 0.5 * 0.5 = 0.25 > prob[2]=0.2
- Update prob[2]=0.25, push (0.25, 2)
Heap: [(0.25, 2), (0.2, 2)]

Step 3: Pop (0.25, 2)
- Reached target! Return 0.25

Result: 0.25 (path 0→1→2) ✓"
```

#### Phase 6: Intuition Validation
- **Why max-heap works**: Always expand highest probability path first
- **Why Dijkstra works**: Optimal substructure property holds
- **Why early exit**: First time target popped from heap is optimal
- **Why O((V+E) log V)**: Each edge processed once, heap operations dominate

### Common Human Pitfalls & How to Avoid Them
1. **"Why not standard Dijkstra?"** → Need max-heap for maximization
2. **"Should I use BFS?"** → Need priority queue for weighted edges
3. **"What about precision?"** → Use log transformation for stability
4. **"Can I use min-heap?"** → Yes, with negative probabilities
5. **"What about floating point?"** → Handle precision carefully

### Real-World Analogy
**Like finding the most reliable communication route:**
- Each connection has a success rate (probability message gets through)
- You want the route with highest overall success rate
- Overall success rate = product of individual connection success rates
- You explore routes in order of decreasing reliability
- Like finding the most dependable way to send important data

### Human-Readable Pseudocode
```
function maxProbability(n, edges, succProb, start, end):
    # Build adjacency list
    adj = adjacency list of size n
    for each edge (u, v) with probability p:
        adj[u].append((v, p))
        adj[v].append((u, p))  # Undirected
    
    # Initialize probabilities
    prob = array of size n, initialized to 0
    prob[start] = 1
    
    # Max-heap: (probability, node)
    maxHeap = priority queue with max-heap property
    maxHeap.push((1.0, start))
    
    while maxHeap is not empty:
        currentProb, currentNode = maxHeap.pop()
        
        # Skip if outdated
        if currentProb < prob[currentNode]:
            continue
        
        # Early exit if reached target
        if currentNode == end:
            return currentProb
        
        # Explore neighbors
        for each (neighbor, edgeProb) in adj[currentNode]:
            newProb = currentProb * edgeProb
            if newProb > prob[neighbor]:
                prob[neighbor] = newProb
                maxHeap.push((newProb, neighbor))
    
    return prob[end]  # 0 if unreachable
```

### Execution Visualization

### Example: 3 nodes, edges: (0,1):0.5, (1,2):0.5, (0,2):0.2
```
Modified Dijkstra Process:

Initial: prob=[1.0, 0.0, 0.0], heap=[(1.0, 0)]

Step 1: Pop (1.0, 0)
- Edge 0→1: newProb = 1.0 × 0.5 = 0.5
  prob[1] = 0.5, heap.push((0.5, 1))
- Edge 0→2: newProb = 1.0 × 0.2 = 0.2
  prob[2] = 0.2, heap.push((0.2, 2))
Heap: [(0.5, 1), (0.2, 2)]

Step 2: Pop (0.5, 1)
- Edge 1→2: newProb = 0.5 × 0.5 = 0.25
  0.25 > prob[2]=0.2, so prob[2] = 0.25
  heap.push((0.25, 2))
Heap: [(0.25, 2), (0.2, 2)]

Step 3: Pop (0.25, 2)
- Reached target node 2!
- Return 0.25

Result: 0.25 (path 0→1→2) ✓
```

### Key Visualization Points:
- **Max-Heap Priority**: Always expand highest probability path
- **Probability Multiplication**: Path probability = product of edges
- **Early Termination**: First target pop gives optimal solution
- **Probability Update**: Keep maximum probability for each node

### Graph Visualization:
```
Probability Graph:
    0.5      0.5
0 ------- 1 ------- 2
 \        /
  \ 0.2  /
   \    /
     (direct)

Path 0→1→2: 0.5 × 0.5 = 0.25
Path 0→2: 0.2
Maximum: 0.25 via 0→1→2
```

### Time Complexity Breakdown:
- **Standard Max-Heap**: O((V + E) log V) time, O(V + E) space
- **Log Transform**: O((V + E) log V) time, O(V + E) space
- **Negative Values**: O((V + E) log V) time, O(V + E) space
- **Early Termination**: Often faster in practice

### Alternative Approaches:

#### 1. Logarithmic Transformation (O((V + E) log V) time, O(V + E) space)
```go
func maxProbabilityLogTransform(n int, edges [][]int, succProb []float64, start int, end int) float64 {
    // Use log probabilities: log(a*b) = log(a) + log(b)
    // Convert multiplication to addition for numerical stability
    // ... implementation details omitted
}
```
- **Pros**: Better numerical stability, additive operations
- **Cons**: Requires log/exp operations, small overhead

#### 2. Min-Heap with Negatives (O((V + E) log V) time, O(V + E) space)
```go
func maxProbabilityNegative(n int, edges [][]int, succProb []float64, start int, end int) float64 {
    // Use negative probabilities with min-heap
    // Minimize negative probability = maximize positive probability
    // ... implementation details omitted
}
```
- **Pros**: Can use standard min-heap implementation
- **Cons**: Less intuitive, negative values

#### 3. BFS with Priority (O(V + E) time for special cases)
```go
func maxProbabilityBFS(n int, edges [][]int, succProb []float64, start int, end int) float64 {
    // Works for unweighted probabilities or special cases
    // Not general solution
    // ... implementation details omitted
}
```
- **Pros**: Simpler implementation
- **Cons**: Doesn't work for general weighted case

### Extensions for Interviews:
- **Path Reconstruction**: Track parent pointers to return actual path
- **Multiple Targets**: Find maximum probability to multiple destinations
- **Precision Handling**: Discuss floating-point precision issues
- **Alternative Objectives**: Minimize expected cost, maximize expected utility
- **Real-world Applications**: Network reliability, communication systems
*/
func main() {
	// Test cases
	testCases := []struct {
		n          int
		edges      [][]int
		succProb   []float64
		start      int
		end        int
		description string
	}{
		{
			3,
			[][]int{{0, 1}, {1, 2}, {0, 2}},
			[]float64{0.5, 0.5, 0.2},
			0, 2,
			"Standard case",
		},
		{
			3,
			[][]int{{0, 1}, {1, 2}, {0, 2}},
			[]float64{0.5, 0.5, 0.3},
			0, 2,
			"Higher direct probability",
		},
		{
			3,
			[][]int{{0, 1}},
			[]float64{0.5},
			0, 2,
			"No path",
		},
		{
			5,
			[][]int{{1, 0}, {2, 0}, {3, 2}, {4, 1}},
			[]float64{0.2, 0.1, 0.7, 0.9},
			0, 3,
			"Complex graph",
		},
		{
			4,
			[][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}},
			[]float64{0.5, 0.5, 0.5, 0.1},
			0, 3,
			"Indirect vs direct",
		},
		{
			2,
			[][]int{{0, 1}},
			[]float64{1.0},
			0, 1,
			"Maximum probability",
		},
		{
			4,
			[][]int{{0, 1}, {1, 2}, {2, 3}},
			[]float64{0.1, 0.1, 0.1},
			0, 3,
			"Low probabilities",
		},
		{
			6,
			[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {0, 5}},
			[]float64{0.5, 0.5, 0.5, 0.5, 0.5, 0.1},
			0, 5,
			"Long path vs direct",
		},
		{
			3,
			[][]int{{0, 1}, {0, 2}, {1, 2}},
			[]float64{0.9, 0.1, 0.9},
			0, 2,
			"Multiple high probability paths",
		},
		{
			1,
			[][]int{},
			[]float64{},
			0, 0,
			"Single node",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  n=%d, start=%d, end=%d\n", tc.n, tc.start, tc.end)
		fmt.Printf("  Edges: %v\n", tc.edges)
		fmt.Printf("  Probabilities: %v\n", tc.succProb)
		
		result1 := maxProbability(tc.n, tc.edges, tc.succProb, tc.start, tc.end)
		result2 := maxProbabilityBFS(tc.n, tc.edges, tc.succProb, tc.start, tc.end)
		result3 := maxProbabilityModifiedDijkstra(tc.n, tc.edges, tc.succProb, tc.start, tc.end)
		result4 := maxProbabilityLogTransform(tc.n, tc.edges, tc.succProb, tc.start, tc.end)
		
		fmt.Printf("  Standard Dijkstra: %.6f\n", result1)
		fmt.Printf("  BFS Approach: %.6f\n", result2)
		fmt.Printf("  Modified Dijkstra: %.6f\n", result3)
		fmt.Printf("  Log Transform: %.6f\n\n", result4)
	}
}
