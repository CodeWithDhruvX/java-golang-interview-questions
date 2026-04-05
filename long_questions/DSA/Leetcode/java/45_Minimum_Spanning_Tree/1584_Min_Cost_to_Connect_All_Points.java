import java.util.*;

public class MinCostToConnectAllPoints {
    
    // 1584. Min Cost to Connect All Points - Minimum Spanning Tree (Prim's Algorithm)
    // Time: O(N^2), Space: O(N)
    public int minCostConnectPoints(int[][] points) {
        if (points.length <= 1) {
            return 0;
        }
        
        int n = points.length;
        boolean[] visited = new boolean[n];
        int[] minDist = new int[n];
        
        // Initialize minDist to infinity
        Arrays.fill(minDist, Integer.MAX_VALUE);
        
        // Start from point 0
        minDist[0] = 0;
        int totalCost = 0;
        
        for (int i = 0; i < n; i++) {
            // Find unvisited point with minimum distance
            int u = -1;
            int minVal = Integer.MAX_VALUE;
            
            for (int j = 0; j < n; j++) {
                if (!visited[j] && minDist[j] < minVal) {
                    minVal = minDist[j];
                    u = j;
                }
            }
            
            if (u == -1) {
                break;
            }
            
            visited[u] = true;
            totalCost += minDist[u];
            
            // Update distances to unvisited points
            for (int v = 0; v < n; v++) {
                if (!visited[v]) {
                    int dist = manhattanDistance(points[u], points[v]);
                    if (dist < minDist[v]) {
                        minDist[v] = dist;
                    }
                }
            }
        }
        
        return totalCost;
    }
    
    private int manhattanDistance(int[] p1, int[] p2) {
        return Math.abs(p1[0] - p2[0]) + Math.abs(p1[1] - p2[1]);
    }
    
    // Prim's algorithm with priority queue optimization
    public int minCostConnectPointsPrimOptimized(int[][] points) {
        if (points.length <= 1) {
            return 0;
        }
        
        int n = points.length;
        boolean[] visited = new boolean[n];
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[1] - b[1]);
        
        // Start from point 0
        pq.offer(new int[]{0, 0});
        int totalCost = 0;
        int edgesUsed = 0;
        
        while (!pq.isEmpty() && edgesUsed < n) {
            int[] current = pq.poll();
            int u = current[0];
            int cost = current[1];
            
            if (visited[u]) {
                continue;
            }
            
            visited[u] = true;
            totalCost += cost;
            edgesUsed++;
            
            // Add all edges from u to unvisited points
            for (int v = 0; v < n; v++) {
                if (!visited[v]) {
                    int dist = manhattanDistance(points[u], points[v]);
                    pq.offer(new int[]{v, dist});
                }
            }
        }
        
        return totalCost;
    }
    
    // Kruskal's algorithm approach
    public int minCostConnectPointsKruskal(int[][] points) {
        if (points.length <= 1) {
            return 0;
        }
        
        int n = points.length;
        List<int[]> edges = new ArrayList<>();
        
        // Generate all possible edges
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                int dist = manhattanDistance(points[i], points[j]);
                edges.add(new int[]{dist, i, j});
            }
        }
        
        // Sort edges by weight
        edges.sort((a, b) -> a[0] - b[0]);
        
        // Union-Find to detect cycles
        UnionFind uf = new UnionFind(n);
        int totalCost = 0;
        int edgesUsed = 0;
        
        for (int[] edge : edges) {
            int dist = edge[0];
            int u = edge[1];
            int v = edge[2];
            
            if (uf.find(u) != uf.find(v)) {
                uf.union(u, v);
                totalCost += dist;
                edgesUsed++;
                
                if (edgesUsed == n - 1) {
                    break;
                }
            }
        }
        
        return totalCost;
    }
    
    // Union-Find data structure
    private static class UnionFind {
        private int[] parent;
        private int[] rank;
        
        public UnionFind(int n) {
            parent = new int[n];
            rank = new int[n];
            for (int i = 0; i < n; i++) {
                parent[i] = i;
                rank[i] = 0;
            }
        }
        
        public int find(int x) {
            if (parent[x] != x) {
                parent[x] = find(parent[x]);
            }
            return parent[x];
        }
        
        public void union(int x, int y) {
            int rootX = find(x);
            int rootY = find(y);
            
            if (rootX != rootY) {
                if (rank[rootX] < rank[rootY]) {
                    parent[rootX] = rootY;
                } else if (rank[rootX] > rank[rootY]) {
                    parent[rootY] = rootX;
                } else {
                    parent[rootY] = rootX;
                    rank[rootX]++;
                }
            }
        }
    }
    
    // Version with detailed explanation
    public class MSTResult {
        int totalCost;
        List<int[]> edges;
        List<String> explanation;
        
        MSTResult(int totalCost, List<int[]> edges, List<String> explanation) {
            this.totalCost = totalCost;
            this.edges = edges;
            this.explanation = explanation;
        }
    }
    
    public MSTResult minCostConnectPointsDetailed(int[][] points) {
        List<String> explanation = new ArrayList<>();
        explanation.add("=== Prim's Algorithm for MST ===");
        explanation.add("Points: " + Arrays.deepToString(points));
        
        int n = points.length;
        boolean[] visited = new boolean[n];
        int[] minDist = new int[n];
        List<int[]> mstEdges = new ArrayList<>();
        
        Arrays.fill(minDist, Integer.MAX_VALUE);
        minDist[0] = 0;
        
        explanation.add("Starting from point 0");
        explanation.add("Initial minDist: " + Arrays.toString(minDist));
        
        int totalCost = 0;
        
        for (int i = 0; i < n; i++) {
            // Find unvisited point with minimum distance
            int u = -1;
            int minVal = Integer.MAX_VALUE;
            
            for (int j = 0; j < n; j++) {
                if (!visited[j] && minDist[j] < minVal) {
                    minVal = minDist[j];
                    u = j;
                }
            }
            
            if (u == -1) {
                break;
            }
            
            visited[u] = true;
            totalCost += minDist[u];
            
            if (i > 0) {
                // Find which point connected to u with minDist[u]
                for (int v = 0; v < n; v++) {
                    if (visited[v] && v != u) {
                        int dist = manhattanDistance(points[u], points[v]);
                        if (dist == minDist[u]) {
                            mstEdges.add(new int[]{v, u, dist});
                            explanation.add(String.format("Added edge: %d -> %d (cost: %d)", v, u, dist));
                            break;
                        }
                    }
                }
            }
            
            explanation.add(String.format("Selected point %d (cost: %d), total: %d", u, minDist[u], totalCost));
            
            // Update distances
            for (int v = 0; v < n; v++) {
                if (!visited[v]) {
                    int dist = manhattanDistance(points[u], points[v]);
                    if (dist < minDist[v]) {
                        minDist[v] = dist;
                        explanation.add(String.format("  Updated distance to point %d: %d", v, dist));
                    }
                }
            }
        }
        
        explanation.add("Final MST cost: " + totalCost);
        explanation.add("MST edges: " + mstEdges.size());
        
        return new MSTResult(totalCost, mstEdges, explanation);
    }
    
    // Boruvka's algorithm (alternative MST approach)
    public int minCostConnectPointsBoruvka(int[][] points) {
        if (points.length <= 1) {
            return 0;
        }
        
        int n = points.length;
        UnionFind uf = new UnionFind(n);
        int totalCost = 0;
        int components = n;
        
        while (components > 1) {
            // Find cheapest edge for each component
            int[] cheapestEdge = new int[n];
            Arrays.fill(cheapestEdge, -1);
            
            for (int i = 0; i < n; i++) {
                for (int j = i + 1; j < n; j++) {
                    if (uf.find(i) != uf.find(j)) {
                        int dist = manhattanDistance(points[i], points[j]);
                        int rootI = uf.find(i);
                        int rootJ = uf.find(j);
                        
                        if (cheapestEdge[rootI] == -1 || dist < cheapestEdge[rootI]) {
                            cheapestEdge[rootI] = dist;
                        }
                        if (cheapestEdge[rootJ] == -1 || dist < cheapestEdge[rootJ]) {
                            cheapestEdge[rootJ] = dist;
                        }
                    }
                }
            }
            
            // Add cheapest edges
            for (int i = 0; i < n; i++) {
                if (cheapestEdge[i] != -1) {
                    totalCost += cheapestEdge[i];
                    components--;
                }
            }
            
            // Merge components
            for (int i = 0; i < n; i++) {
                for (int j = i + 1; j < n; j++) {
                    if (uf.find(i) != uf.find(j)) {
                        int dist = manhattanDistance(points[i], points[j]);
                        if (dist == cheapestEdge[uf.find(i)] || dist == cheapestEdge[uf.find(j)]) {
                            uf.union(i, j);
                        }
                    }
                }
            }
        }
        
        return totalCost;
    }
    
    // Performance comparison
    public void comparePerformance(int[][] points) {
        System.out.println("=== Performance Comparison ===");
        System.out.println("Points: " + points.length);
        
        // Prim's basic
        long startTime = System.nanoTime();
        int result1 = minCostConnectPoints(points);
        long endTime = System.nanoTime();
        System.out.printf("Prim's basic: %d (took %d ns)\n", result1, endTime - startTime);
        
        // Prim's optimized
        startTime = System.nanoTime();
        int result2 = minCostConnectPointsPrimOptimized(points);
        endTime = System.nanoTime();
        System.out.printf("Prim's optimized: %d (took %d ns)\n", result2, endTime - startTime);
        
        // Kruskal's
        startTime = System.nanoTime();
        int result3 = minCostConnectPointsKruskal(points);
        endTime = System.nanoTime();
        System.out.printf("Kruskal's: %d (took %d ns)\n", result3, endTime - startTime);
    }
    
    public static void main(String[] args) {
        MinCostToConnectAllPoints mst = new MinCostToConnectAllPoints();
        
        // Test cases
        int[][][] testCases = {
            {{0,0},{2,2},{3,10},{5,2},{7,0}},
            {{3,12},{-2,5},{-4,1}},
            {{0,0},{1,1},{1,0},{-1,1}},
            {{0,0}},
            {{0,0},{1,0}},
            {{0,0},{2,0},{4,0},{6,0}},
            {{0,0},{0,2},{0,4},{0,6}},
            {{-1000,-1000},{1000,1000},{-1000,1000},{1000,-1000}}
        };
        
        String[] descriptions = {
            "Standard case",
            "Triangle case",
            "Square case",
            "Single point",
            "Two points",
            "Line horizontal",
            "Line vertical",
            "Large coordinates"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("Points: %s\n", Arrays.deepToString(testCases[i]));
            
            int result1 = mst.minCostConnectPoints(testCases[i]);
            int result2 = mst.minCostConnectPointsPrimOptimized(testCases[i]);
            int result3 = mst.minCostConnectPointsKruskal(testCases[i]);
            
            System.out.printf("Prim's basic: %d\n", result1);
            System.out.printf("Prim's optimized: %d\n", result2);
            System.out.printf("Kruskal's: %d\n", result3);
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        MSTResult detailedResult = mst.minCostConnectPointsDetailed(testCases[0]);
        System.out.printf("Total cost: %d\n", detailedResult.totalCost);
        System.out.println("MST edges:");
        for (int[] edge : detailedResult.edges) {
            System.out.printf("  %d -> %d (cost: %d)\n", edge[0], edge[1], edge[2]);
        }
        System.out.println("Explanation:");
        for (String step : detailedResult.explanation) {
            System.out.println("  " + step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        int[][] performanceTest = new int[100][];
        Random rand = new Random();
        for (int i = 0; i < 100; i++) {
            performanceTest[i] = new int[]{rand.nextInt(1000), rand.nextInt(1000)};
        }
        
        mst.comparePerformance(performanceTest);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        System.out.printf("All same point: %d\n", 
            mst.minCostConnectPoints(new int[][]{{0,0},{0,0},{0,0}}));
        System.out.printf("Already connected: %d\n", 
            mst.minCostConnectPoints(new int[][]{{0,0},{1,0},{2,0}}));
        System.out.printf("Large distances: %d\n", 
            mst.minCostConnectPoints(new int[][]{{-1000,-1000},{1000,1000}}));
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Minimum Spanning Tree
- **Graph Connectivity**: Connect all points with minimum total cost
- **Prim's Algorithm**: Greedy approach with priority queue
- **Kruskal's Algorithm**: Edge sorting with Union-Find
- **Boruvka's Algorithm**: Optimized for dense graphs

## 2. PROBLEM CHARACTERISTICS
- **Point Connection**: Connect all points with minimum total distance
- **Manhattan Distance**: Distance metric for grid-based points
- **Complete Graph**: Need to connect N points with N-1 edges
- **Optimization**: Find minimum cost spanning tree

## 3. SIMILAR PROBLEMS
- Network Design
- Minimum Cost to Connect Cities
- Cable Installation Problem
- Road Network Construction

## 4. KEY OBSERVATIONS
- MST connects all vertices with minimum total weight
- Manhattan distance works for grid-based coordinates
- Prim's algorithm: O(N²) time, O(N) space
- Kruskal's algorithm: O(E log E) time, O(E) space
- Boruvka's algorithm: O(E log V) for dense graphs

## 5. VARIATIONS & EXTENSIONS
- Different distance metrics (Euclidean)
- Weighted edges with different constraints
- Dynamic point addition
- Multiple MST variants

## 6. INTERVIEW INSIGHTS
- Clarify: "What distance metric?"
- Edge cases: single point, duplicate points, large coordinates
- Time complexity: O(N²) vs O(E log E) vs O(E log V)
- Space complexity: O(N) vs O(E) for different algorithms

## 7. COMMON MISTAKES
- Incorrect distance calculation
- Wrong priority queue implementation
- Union-Find path compression issues
- Not handling edge cases properly
- Incorrect edge weight assignment

## 8. OPTIMIZATION STRATEGIES
- Use priority queue for Prim's algorithm
- Implement Union-Find with path compression
- Early termination when MST is complete
- Efficient distance calculations

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like building a road network:**
- You have cities (points) that need to be connected
- Want to minimize total road length (cost)
- Each road connects two cities directly
- MST finds cheapest way to connect all cities
- This is like finding the most efficient network topology

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of 2D points
2. **Goal**: Connect all points with minimum total distance
3. **Output**: Minimum total cost

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(N³) to try all possible connections
- **"How to optimize?"** → Use greedy MST algorithms
- **"Why Prim's algorithm?"** → Grows tree from arbitrary starting point
- **"Why Kruskal's algorithm?"** → Sorts edges and adds them greedily

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use Prim's algorithm:
1. Start from any point (point 0)
2. Maintain minDist array and visited set
3. At each step:
   - Find unvisited point with minimum distance to tree
   - Add it to tree (update minDist for its neighbors)
   - Mark as visited
4. Continue until all points are visited
5. Sum all distances added to tree
6. Return total cost"
```

#### Phase 4: Edge Case Handling
- **Single point**: Return 0 (no cost to connect)
- **Empty array**: Return 0
- **Duplicate points**: Handle appropriately
- **Large coordinates**: Use efficient distance calculations

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Points: [[0,0], [2,2], [3,1], [4,0]]

Human thinking:
"Let's apply Prim's algorithm:

Step 1: Initialize
visited = [false, false, false, false, false]
minDist = [∞, ∞, ∞, ∞]
minDist[0] = 0 (start from point 0)
totalCost = 0

Step 2: First iteration
Find unvisited point with minimum minDist:
- Point 1: minDist[1] = |0-2| + |0-2| = 4
- Point 2: minDist[2] = |0-2| + |0-2| = 4
- Point 3: minDist[3] = |0-1| + |2-1| = 3 (minimum)
Select point 3, cost = 3
Mark visited[3] = true
totalCost = 3

Update minDist for neighbors of point 3:
- Point 0: already visited
- Point 1: minDist[1] = min(4, |3-2|+|2-2|) = 4
- Point 2: minDist[2] = min(4, |3-2|+|2-2|) = 4
- Point 4: minDist[4] = min(4, |3-0|+|4-0|) = 7

Step 3: Second iteration
Find unvisited point with minimum minDist:
- Point 1: minDist[1] = 4 (minimum)
Select point 1, cost = 4
Mark visited[1] = true
totalCost = 3 + 4 = 7

Update minDist for neighbors of point 1:
- Point 0: already visited
- Point 2: minDist[2] = min(4, |1-2|+|2-2|) = 4
- Point 3: already visited
- Point 4: minDist[4] = min(7, |1-0|+|4-0|) = 7

Continue until all points visited...

Final MST edges: (0,3), (1,3), (2,4)
Total cost: 3 + 4 + 7 = 14 ✓

Manual verification:
All points connected with minimum total distance ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Greedy choice is optimal for MST
- **Why it's efficient**: O(N²) vs O(N³) brute force
- **Why it's correct**: Cut property ensures optimality

### Common Human Pitfalls & How to Avoid Them
1. **"Why not try all combinations?"** → O(N³) too slow
2. **"What about distance metric?"** → Use Manhattan for grid coordinates
3. **"How to implement priority queue?"** → Use min-heap for efficiency
4. **"What about cycles?"** → MST is always a tree (no cycles)

### Real-World Analogy
**Like designing an optimal road network:**
- You have cities that need to be connected by roads
- Want to minimize total road construction cost
- Each road connects two cities with cost equal to distance
- MST finds cheapest way to connect all cities
- This is used in network design, infrastructure planning
- Useful in telecommunications, transportation, utility networks
- Like finding the most efficient way to build a network

### Human-Readable Pseudocode
```
function primMST(points):
    if points.length <= 1:
        return 0
    
    n = points.length
    visited = [false] * n
    minDist = [∞] * n
    minDist[0] = 0
    totalCost = 0
    
    for i from 1 to n-1:
        // Find unvisited point with minimum distance
        u = -1
        minVal = ∞
        for j from 0 to n-1:
            if !visited[j] and minDist[j] < minVal:
                minVal = minDist[j]
                u = j
        
        if u == -1:
            break
        
        visited[u] = true
        totalCost += minDist[u]
        
        // Update distances for neighbors of u
        for v from 0 to n-1:
            if !visited[v]:
                dist = manhattanDistance(points[u], points[v])
                if dist < minDist[v]:
                    minDist[v] = dist
    
    return totalCost
```

### Execution Visualization

### Example: points=[[0,0],[2,2],[3,1],[4,0]]
```
Prim's Algorithm Process:

Step 1: Initialize
visited = [F,F,F,F,F]
minDist = [0,∞,∞,∞]
Start from point 0, cost = 0

Step 2: First iteration
Find minimum unvisited point:
- Point 1: dist = 4, Point 2: dist = 4, Point 3: dist = 3 (min)
Select point 3, cost = 3
visited = [F,F,T,F,F]
Update neighbors of point 3

Step 3: Second iteration
Find minimum unvisited point:
- Point 1: dist = 4 (min)
Select point 1, cost = 4
visited = [T,T,F,F]
Update neighbors of point 1

Continue...

Final MST: edges (0,3), (1,3), (2,4)
Total cost: 3 + 4 + 7 = 14 ✓

Visualization:
Greedy choice of minimum edge at each step
Tree grows by adding cheapest connections
All points connected with minimum total distance
```

### Key Visualization Points:
- **Greedy Choice**: Always select cheapest available edge
- **Tree Growth**: MST grows from starting point
- **Cut Property**: No cycles in final tree
- **Optimality**: Local optimal choices lead to global optimum

### Memory Layout Visualization:
```
Points: P0(0,0), P1(2,2), P2(3,1), P3(4,0)

MST Construction:
Start: P0, cost=0
Add P3: cost=3, edges=[(0,3)]
Add P1: cost=4, edges=[(0,3),(1,3)]
Add P2: cost=7, edges=[(0,3),(1,3),(2,4)]

Final tree connects all points
Total cost: 14 (minimum possible)

Prim's algorithm builds MST incrementally
Each step maintains tree property and optimality
```

### Time Complexity Breakdown:
- **Prim's Basic**: O(N²) time, O(N) space
- **Prim's Optimized**: O(N log N) time, O(N) space
- **Kruskal's**: O(E log E) time, O(E) space
- **Boruvka's**: O(E log V) expected for dense graphs
- **Optimal**: Choose based on graph density and requirements
- **vs Brute Force**: O(N³) vs O(N²) with MST algorithms
*/
