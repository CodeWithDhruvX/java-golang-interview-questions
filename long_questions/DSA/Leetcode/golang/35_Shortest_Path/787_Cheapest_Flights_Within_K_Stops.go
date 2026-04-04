package main

import (
	"fmt"
	"math"
)

// 787. Cheapest Flights Within K Stops - Bellman-Ford with K stops constraint
// Time: O(K*E), Space: O(V)
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// Initialize distances
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[src] = 0
	
	// Relax edges up to K+1 times (K stops = K+1 edges)
	for i := 0; i <= k; i++ {
		// Copy current distances to avoid using updated distances in same iteration
		tempDist := make([]int, n)
		copy(tempDist, dist)
		
		updated := false
		for _, flight := range flights {
			from, to, price := flight[0], flight[1], flight[2]
			
			if dist[from] != math.MaxInt32 && dist[from]+price < tempDist[to] {
				tempDist[to] = dist[from] + price
				updated = true
			}
		}
		
		dist = tempDist
		
		if !updated {
			break // No more updates needed
		}
	}
	
	if dist[dst] == math.MaxInt32 {
		return -1
	}
	
	return dist[dst]
}

// Modified Bellman-Ford with early termination
func findCheapestPriceOptimized(n int, flights [][]int, src int, dst int, k int) int {
	// Initialize distances
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt32
	}
	dist[src] = 0
	
	// Relax edges exactly K+1 times
	for i := 0; i <= k; i++ {
		updated := false
		
		// Use previous iteration's distances
		prevDist := make([]int, n)
		copy(prevDist, dist)
		
		for _, flight := range flights {
			from, to, price := flight[0], flight[1], flight[2]
			
			if prevDist[from] != math.MaxInt32 && prevDist[from]+price < dist[to] {
				dist[to] = prevDist[from] + price
				updated = true
			}
		}
		
		if !updated {
			break
		}
	}
	
	if dist[dst] == math.MaxInt32 {
		return -1
	}
	
	return dist[dst]
}

// BFS approach with priority queue (Dijkstra-like)
func findCheapestPriceBFS(n int, flights [][]int, src int, dst int, k int) int {
	// Build adjacency list
	adj := make(map[int][]Flight)
	for _, flight := range flights {
		from, to, price := flight[0], flight[1], flight[2]
		adj[from] = append(adj[from], Flight{to, price})
	}
	
	// Priority queue: {price, node, stops}
	type State struct {
		price int
		node  int
		stops int
	}
	
	// Use a simple priority queue implementation
	pq := []State{{0, src, 0}}
	
	// Track minimum price to reach each node with given stops
	minPrice := make(map[[2]int]int) // {node, stops} -> price
	
	for len(pq) > 0 {
		// Find minimum price state
		minIdx := 0
		for i := 1; i < len(pq); i++ {
			if pq[i].price < pq[minIdx].price {
				minIdx = i
			}
		}
		
		current := pq[minIdx]
		pq = append(pq[:minIdx], pq[minIdx+1:]...)
		
		price, node, stops := current.price, current.node, current.stops
		
		// Early exit
		if node == dst {
			return price
		}
		
		// Check if we've exceeded stops
		if stops > k {
			continue
		}
		
		// Skip if we've found a better path to this node with fewer stops
		if prevPrice, exists := minPrice[[2]int{node, stops}]; exists && prevPrice <= price {
			continue
		}
		minPrice[[2]int{node, stops}] = price
		
		// Explore neighbors
		for _, flight := range adj[node] {
			newPrice := price + flight.price
			newStops := stops + 1
			
			// Check if this path is promising
			if prevPrice, exists := minPrice[[2]int{flight.to, newStops}]; !exists || newPrice < prevPrice {
				pq = append(pq, State{newPrice, flight.to, newStops})
			}
		}
	}
	
	return -1
}

// Dynamic Programming approach
func findCheapestPriceDP(n int, flights [][]int, src int, dst int, k int) int {
	// dp[i][j] = minimum cost to reach node j using exactly i flights
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][src] = 0
	
	for i := 1; i <= k+1; i++ {
		// Copy previous row
		for j := range dp[i] {
			dp[i][j] = dp[i-1][j]
		}
		
		// Try all flights
		for _, flight := range flights {
			from, to, price := flight[0], flight[1], flight[2]
			
			if dp[i-1][from] != math.MaxInt32 {
				if dp[i-1][from]+price < dp[i][to] {
					dp[i][to] = dp[i-1][from] + price
				}
			}
		}
	}
	
	// Find minimum cost across all flights up to K+1
	minCost := math.MaxInt32
	for i := 0; i <= k+1; i++ {
		if dp[i][dst] < minCost {
			minCost = dp[i][dst]
		}
	}
	
	if minCost == math.MaxInt32 {
		return -1
	}
	
	return minCost
}

// Flight represents a flight edge
type Flight struct {
	to    int
	price int
}

// Floyd-Warshall with K stops constraint
func findCheapestPriceFloydWarshall(n int, flights [][]int, src int, dst int, k int) int {
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
	
	// Set direct flights
	for _, flight := range flights {
		from, to, price := flight[0], flight[1], flight[2]
		dist[from][to] = price
	}
	
	// Floyd-Warshall with intermediate nodes limited to K+1
	for mid := 0; mid < n; mid++ {
		for from := 0; from < n; from++ {
			for to := 0; to < n; to++ {
				if dist[from][mid] != math.MaxInt32 && dist[mid][to] != math.MaxInt32 {
					if dist[from][mid]+dist[mid][to] < dist[from][to] {
						dist[from][to] = dist[from][mid] + dist[mid][to]
					}
				}
			}
		}
		
		// Check if we've used K+1 intermediate nodes
		if mid == k {
			break
		}
	}
	
	if dist[src][dst] == math.MaxInt32 {
		return -1
	}
	
	return dist[src][dst]
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Constrained Shortest Path with K Stops
- **Bellman-Ford Modification**: Limit edge relaxations to K+1 iterations
- **State Tracking**: Track (cost, node, stops) in priority queue
- **Dynamic Programming**: DP[i][j] = min cost to reach node j with i flights
- **Early Termination**: Stop when no more improvements possible

## 2. PROBLEM CHARACTERISTICS
- **Weighted Directed Graph**: Flight routes with costs
- **Path Constraint**: Maximum K intermediate stops allowed
- **Cost Minimization**: Find cheapest path within constraint
- **Edge Relaxation**: Standard shortest path with additional constraint

## 3. SIMILAR PROBLEMS
- Network Delay Time (LeetCode 743) - Unconstrained shortest path
- Path with Maximum Probability (LeetCode 1514) - Max probability path
- Find the Cheapest Flight (LeetCode 787) - Same problem
- Minimum Cost to Connect All Points (LeetCode 1584) - MST problem

## 4. KEY OBSERVATIONS
- **Bellman-Ford Natural**: K stops = K+1 edges = K+1 relaxations
- **State Complexity**: Need to track both cost and stops used
- **DP Formulation**: Each iteration allows one more flight
- **Priority Queue**: Can use modified Dijkstra with stops constraint

## 5. VARIATIONS & EXTENSIONS
- **Priority Queue Approach**: Modified Dijkstra with stops tracking
- **Dynamic Programming**: 2D DP table for flights and destinations
- **Floyd-Warshall**: Limited intermediate nodes
- **Bidirectional Search**: Search from both ends with constraints

## 6. INTERVIEW INSIGHTS
- Always clarify: "Flight constraints? Negative costs? Multiple queries?"
- Edge cases: no path within K stops, direct flight, K=0
- Time complexity: O(K*E) for Bellman-Ford, O((V+E) log V) for PQ
- Space complexity: O(V) for distances, O(V*E) for DP
- Key insight: K stops = exactly K+1 edge relaxations

## 7. COMMON MISTAKES
- Wrong number of iterations (K vs K+1)
- Not copying distances properly between iterations
- Using updated distances in same iteration
- Missing early termination optimization
- Not handling unreachable destinations

## 8. OPTIMIZATION STRATEGIES
- **Bellman-Ford**: O(K*E) time, O(V) space - standard approach
- **Priority Queue**: O((V+E) log V) time, O(V) space - faster in practice
- **Dynamic Programming**: O(K*E) time, O(V*K) space - clear formulation
- **Floyd-Warshall**: O(V³) time, O(V²) space - all pairs

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like planning a multi-leg flight trip with budget constraints:**
- Each flight has a cost and connects two cities
- You can make at most K+1 flights (K stops means K+1 legs)
- You want the cheapest way to reach your destination
- You can't use more flights than allowed
- Like finding the cheapest multi-city itinerary within flight limits

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Flight routes (from, to, price), source, destination, max stops K
2. **Goal**: Find cheapest price from source to destination with ≤K stops
3. **Rules**: K stops = at most K+1 flights, can't exceed limit
4. **Output**: Minimum price or -1 if impossible

#### Phase 2: Key Insight Recognition
- **"Bellman-Ford natural"** → K stops = K+1 edge relaxations
- **"State complexity"** → Need to track both cost and number of flights
- **"DP formulation"** → Each iteration allows one more flight
- **"Priority queue possible"** → Modified Dijkstra with stops constraint

#### Phase 3: Strategy Development
```
Human thought process:
"I need cheapest flight with at most K stops.
Standard shortest path doesn't consider stops constraint.

Bellman-Ford Approach:
1. Initialize distances: src=0, others=∞
2. For each iteration i from 0 to K:
   - Copy current distances (don't use updates in same iteration)
   - Relax all flights using previous iteration distances
   - This represents paths with exactly i+1 flights
3. After K+1 iterations, we have paths with ≤K+1 flights
4. Return distance to destination

This handles the K stops constraint naturally!"
```

#### Phase 4: Edge Case Handling
- **No path**: Return -1 if destination unreachable within K stops
- **Direct flight**: K=0 means only direct flights allowed
- **Large K**: If K ≥ V-1, becomes standard shortest path
- **Same source/destination**: Return 0

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: n=3, flights=[[0,1,100],[1,2,100],[0,2,500]], src=0, dst=2, k=1

Human thinking:
"Bellman-Ford Approach:
Iteration 0 (0 flights, only source reachable):
dist = [0, ∞, ∞]

Iteration 1 (1 flight, direct flights only):
- Flight 0→1: dist[1] = min(∞, 0+100) = 100
- Flight 1→2: dist[2] = min(∞, ∞+100) = ∞ (dist[1] was ∞ in this iteration)
- Flight 0→2: dist[2] = min(∞, 0+500) = 500
dist = [0, 100, 500]

Iteration 2 (2 flights, one stop allowed):
Copy prevDist = [0, 100, 500]
- Flight 0→1: dist[1] = min(100, 0+100) = 100 (no change)
- Flight 1→2: dist[2] = min(500, 100+100) = 200 (improvement!)
- Flight 0→2: dist[2] = min(200, 0+500) = 200 (no change)
dist = [0, 100, 200]

Result: 200 (path 0→1→2) ✓"
```

#### Phase 6: Intuition Validation
- **Why Bellman-Ford works**: Each iteration adds one more flight to paths
- **Why copy distances**: Prevent using same iteration updates
- **Why K+1 iterations**: K stops = K+1 flights maximum
- **Why O(K*E)**: K+1 iterations, each processes all E edges

### Common Human Pitfalls & How to Avoid Them
1. **"Why not standard Dijkstra?"** → Doesn't handle stops constraint
2. **"Should I use K or K+1 iterations?"** → K stops = K+1 flights
3. **"What about copying distances?"** → Must prevent same-iteration updates
4. **"Can I use DP?"** → Yes, 2D DP formulation works well
5. **"What about priority queue?"** → Modified Dijkstra with stops tracking

### Real-World Analogy
**Like booking a multi-city flight itinerary:**
- Each flight leg has a price and connects two cities
- You have a budget and want to minimize total cost
- You can only make a limited number of stops (layovers)
- You want the cheapest way to reach your final destination
- Like planning a business trip with stopover limits

### Human-Readable Pseudocode
```
function findCheapestPrice(n, flights, src, dst, k):
    # Initialize distances
    dist = array of size n, initialized to infinity
    dist[src] = 0
    
    # Bellman-Ford with K stops constraint
    for i from 0 to k:
        # Copy previous distances
        tempDist = copy(dist)
        updated = false
        
        # Relax all flights
        for each flight (from, to, price):
            if dist[from] != infinity and dist[from] + price < tempDist[to]:
                tempDist[to] = dist[from] + price
                updated = true
        
        dist = tempDist
        
        if not updated:
            break  # No more improvements
    
    return dist[dst] if dist[dst] != infinity else -1
```

### Execution Visualization

### Example: n=3, flights=[[0,1,100],[1,2,100],[0,2,500]], src=0, dst=2, k=1
```
Bellman-Ford Process:

Initial: dist = [0, ∞, ∞]

Iteration 0 (0 flights):
- No flights processed (only source reachable)
dist = [0, ∞, ∞]

Iteration 1 (1 flight allowed):
Copy tempDist = [0, ∞, ∞]
- Flight 0→1: tempDist[1] = 0 + 100 = 100
- Flight 1→2: tempDist[2] = ∞ (dist[1] was ∞)
- Flight 0→2: tempDist[2] = 0 + 500 = 500
dist = [0, 100, 500]

Iteration 2 (2 flights allowed, 1 stop):
Copy tempDist = [0, 100, 500]
- Flight 0→1: tempDist[1] = min(100, 0+100) = 100
- Flight 1→2: tempDist[2] = min(500, 100+100) = 200 ✓
- Flight 0→2: tempDist[2] = min(200, 0+500) = 200
dist = [0, 100, 200]

Result: 200 (path 0→1→2 with 1 stop) ✓
```

### Key Visualization Points:
- **Iteration Control**: Each iteration allows one more flight
- **Distance Copying**: Prevent using same-iteration updates
- **Path Building**: Build paths incrementally by flight count
- **Constraint Enforcement**: K stops = K+1 iterations maximum

### Flight Path Visualization:
```
Flight Network:
    100      100
0 -------> 1 -------> 2
 \                ^
  \500           /
   \            /
    ----------> (direct)

Path Options:
- Direct: 0→2, cost=500, 0 stops
- One stop: 0→1→2, cost=200, 1 stop
- Constraint: k=1 (allow 1 stop)
Cheapest: 200 via 0→1→2 ✓
```

### Time Complexity Breakdown:
- **Bellman-Ford**: O(K*E) time, O(V) space - standard approach
- **Priority Queue**: O((V+E) log V) time, O(V) space - modified Dijkstra
- **Dynamic Programming**: O(K*E) time, O(V*K) space - 2D DP
- **Floyd-Warshall**: O(V³) time, O(V²) space - all pairs

### Alternative Approaches:

#### 1. Priority Queue with Stops (O((V+E) log V) time, O(V) space)
```go
func findCheapestPricePQ(n int, flights [][]int, src int, dst int, k int) int {
    // Modified Dijkstra with stops tracking
    // State: {cost, node, stops}
    // Priority queue by cost
    // ... implementation details omitted
}
```
- **Pros**: Often faster in practice, early termination
- **Cons**: More complex state management

#### 2. Dynamic Programming (O(K*E) time, O(V*K) space)
```go
func findCheapestPriceDP(n int, flights [][]int, src int, dst int, k int) int {
    // dp[i][j] = min cost to reach node j with exactly i flights
    // Build solution incrementally
    // ... implementation details omitted
}
```
- **Pros**: Clear formulation, easy to understand
- **Cons**: Higher space complexity

#### 3. Floyd-Warshall with Constraints (O(V³) time, O(V²) space)
```go
func findCheapestPriceFW(n int, flights [][]int, src int, dst int, k int) int {
    // Limited intermediate nodes
    // Not optimal for single-source queries
    // ... implementation details omitted
}
```
- **Pros**: Gives all pairs shortest paths
- **Cons**: Overkill for single-source problem

### Extensions for Interviews:
- **Path Reconstruction**: Track actual flight route
- **Multiple Queries**: Handle multiple (src,dst,k) queries efficiently
- **Negative Costs**: Handle potential negative flight costs
- **Time Constraints**: Add time dimension to flights
- **Real-world Variations**: Flight schedules, layover times
*/
func main() {
	// Test cases
	testCases := []struct {
		n          int
		flights    [][]int
		src        int
		dst        int
		k          int
		description string
	}{
		{
			3,
			[][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			0, 2, 1,
			"Standard case",
		},
		{
			3,
			[][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			0, 2, 0,
			"Direct flight only",
		},
		{
			4,
			[][]int{{0, 1, 100}, {1, 2, 100}, {2, 3, 100}, {0, 3, 500}},
			0, 3, 1,
			"Within K stops",
		},
		{
			4,
			[][]int{{0, 1, 100}, {1, 2, 100}, {2, 3, 100}, {0, 3, 500}},
			0, 3, 0,
			"Direct flight expensive",
		},
		{
			5,
			[][]int{{0, 1, 100}, {1, 2, 100}, {2, 3, 100}, {3, 4, 100}, {0, 4, 500}},
			0, 4, 2,
			"Multiple stops",
		},
		{
			2,
			[][]int{{0, 1, 100}},
			0, 1, 0,
			"Simple case",
		},
		{
			2,
			[][]int{{0, 1, 100}},
			0, 1, 1,
			"Extra stops allowed",
		},
		{
			3,
			[][]int{{0, 1, 100}, {1, 2, 200}, {0, 2, 500}},
			0, 2, 1,
			"Indirect cheaper",
		},
		{
			3,
			[][]int{{0, 1, 100}, {1, 2, 200}, {0, 2, 50}},
			0, 2, 1,
			"Direct cheaper",
		},
		{
			5,
			[][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {3, 4, 1}, {0, 4, 10}},
			0, 4, 3,
			"Long chain vs direct",
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  n=%d, src=%d, dst=%d, k=%d\n", tc.n, tc.src, tc.dst, tc.k)
		fmt.Printf("  Flights: %v\n", tc.flights)
		
		result1 := findCheapestPrice(tc.n, tc.flights, tc.src, tc.dst, tc.k)
		result2 := findCheapestPriceOptimized(tc.n, tc.flights, tc.src, tc.dst, tc.k)
		result3 := findCheapestPriceBFS(tc.n, tc.flights, tc.src, tc.dst, tc.k)
		result4 := findCheapestPriceDP(tc.n, tc.flights, tc.src, tc.dst, tc.k)
		result5 := findCheapestPriceFloydWarshall(tc.n, tc.flights, tc.src, tc.dst, tc.k)
		
		fmt.Printf("  Bellman-Ford: %d\n", result1)
		fmt.Printf("  Optimized BF: %d\n", result2)
		fmt.Printf("  BFS Approach: %d\n", result3)
		fmt.Printf("  DP Approach: %d\n", result4)
		fmt.Printf("  Floyd-Warshall: %d\n\n", result5)
	}
}
