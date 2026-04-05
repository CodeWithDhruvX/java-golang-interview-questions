import java.util.*;

public class NetworkDelayTime {
    
    // 743. Network Delay Time - DP on Graphs
    // Time: O(N^3) for Floyd-Warshall, O(N^2) for optimized approaches
    public int networkDelayTimeDP(int[][] times, int n, int k) {
        if (times.length == 0) {
            return -1;
        }
        
        // Build adjacency matrix
        int[][] dist = new int[n + 1][n + 1];
        for (int i = 0; i <= n; i++) {
            Arrays.fill(dist[i], Integer.MAX_VALUE);
            dist[i][i] = 0;
        }
        
        // Fill direct edges
        for (int[] time : times) {
            int from = time[0], to = time[1], weight = time[2];
            dist[from][to] = weight;
        }
        
        // Floyd-Warshall algorithm
        for (int intermediate = 1; intermediate <= n; intermediate++) {
            for (int from = 1; from <= n; from++) {
                for (int to = 1; to <= n; to++) {
                    if (dist[from][intermediate] != Integer.MAX_VALUE && 
                        dist[intermediate][to] != Integer.MAX_VALUE) {
                        int newDist = dist[from][intermediate] + dist[intermediate][to];
                        if (newDist < dist[from][to]) {
                            dist[from][to] = newDist;
                        }
                    }
                }
            }
        }
        
        // Find maximum distance from k
        int maxDist = 0;
        for (int i = 1; i <= n; i++) {
            if (i != k && dist[k][i] > maxDist) {
                maxDist = dist[k][i];
            }
        }
        
        return maxDist == Integer.MAX_VALUE ? -1 : maxDist;
    }
    
    // DP with Dijkstra optimization
    static class Edge {
        int to;
        int weight;
        
        Edge(int to, int weight) {
            this.to = to;
            this.weight = weight;
        }
    }
    
    public int networkDelayTimeDijkstraDP(int[][] times, int n, int k) {
        if (times.length == 0) {
            return -1;
        }
        
        // Build adjacency list
        List<Edge>[] adj = new ArrayList[n + 1];
        for (int i = 0; i <= n; i++) {
            adj[i] = new ArrayList<>();
        }
        
        for (int[] time : times) {
            int from = time[0], to = time[1], weight = time[2];
            adj[from].add(new Edge(to, weight));
        }
        
        // Dijkstra from source k
        int[] dist = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[1] - b[1]);
        pq.offer(new int[]{k, 0});
        dist[k] = 0;
        
        while (!pq.isEmpty()) {
            int[] current = pq.poll();
            int node = current[0];
            int currentDist = current[1];
            
            if (currentDist > dist[node]) {
                continue;
            }
            
            for (Edge edge : adj[node]) {
                int newDist = currentDist + edge.weight;
                if (newDist < dist[edge.to]) {
                    dist[edge.to] = newDist;
                    pq.offer(new int[]{edge.to, newDist});
                }
            }
        }
        
        // Find maximum distance
        int maxDist = 0;
        for (int i = 1; i <= n; i++) {
            if (dist[i] > maxDist) {
                maxDist = dist[i];
            }
        }
        
        return maxDist == Integer.MAX_VALUE ? -1 : maxDist;
    }
    
    // DP with Bellman-Ford
    public int networkDelayTimeBellmanFord(int[][] times, int n, int k) {
        if (times.length == 0) {
            return -1;
        }
        
        int[] dist = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[k] = 0;
        
        // Relax edges N-1 times
        for (int i = 0; i < n; i++) {
            boolean updated = false;
            
            for (int[] time : times) {
                int from = time[0], to = time[1], weight = time[2];
                
                if (dist[from] != Integer.MAX_VALUE && 
                    dist[from] + weight < dist[to]) {
                    dist[to] = dist[from] + weight;
                    updated = true;
                }
            }
            
            if (!updated) {
                break;
            }
        }
        
        // Find maximum distance
        int maxDist = 0;
        for (int i = 1; i <= n; i++) {
            if (dist[i] > maxDist) {
                maxDist = dist[i];
            }
        }
        
        return maxDist == Integer.MAX_VALUE ? -1 : maxDist;
    }
    
    // DP with topological sort
    public int networkDelayTimeTopologicalDP(int[][] times, int n, int k) {
        if (times.length == 0) {
            return -1;
        }
        
        // Build adjacency list and indegree
        List<Edge>[] adj = new ArrayList[n + 1];
        int[] indegree = new int[n + 1];
        
        for (int i = 0; i <= n; i++) {
            adj[i] = new ArrayList<>();
        }
        
        for (int[] time : times) {
            int from = time[0], to = time[1], weight = time[2];
            adj[from].add(new Edge(to, weight));
            indegree[to]++;
        }
        
        // Topological sort
        Queue<Integer> queue = new LinkedList<>();
        int[] topoOrder = new int[n + 1];
        int index = 0;
        
        for (int i = 1; i <= n; i++) {
            if (indegree[i] == 0) {
                queue.offer(i);
            }
        }
        
        while (!queue.isEmpty()) {
            int node = queue.poll();
            topoOrder[index++] = node;
            
            for (Edge edge : adj[node]) {
                indegree[edge.to]--;
                if (indegree[edge.to] == 0) {
                    queue.offer(edge.to);
                }
            }
        }
        
        // DP in topological order
        int[] dist = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[k] = 0;
        
        for (int i = 0; i < n; i++) {
            int node = topoOrder[i];
            
            if (dist[node] != Integer.MAX_VALUE) {
                for (Edge edge : adj[node]) {
                    if (dist[node] + edge.weight < dist[edge.to]) {
                        dist[edge.to] = dist[node] + edge.weight;
                    }
                }
            }
        }
        
        // Find maximum distance
        int maxDist = 0;
        for (int i = 1; i <= n; i++) {
            if (dist[i] > maxDist) {
                maxDist = dist[i];
            }
        }
        
        return maxDist == Integer.MAX_VALUE ? -1 : maxDist;
    }
    
    // Version with detailed explanation
    public class NetworkDelayResult {
        int delay;
        int[] distances;
        java.util.List<String> explanation;
        
        NetworkDelayResult(int delay, int[] distances, java.util.List<String> explanation) {
            this.delay = delay;
            this.distances = distances;
            this.explanation = explanation;
        }
    }
    
    public NetworkDelayResult networkDelayTimeDetailed(int[][] times, int n, int k) {
        java.util.List<String> explanation = new ArrayList<>();
        explanation.add("=== DP on Graphs for Network Delay Time ===");
        explanation.add(String.format("Nodes: %d, Source: %d", n, k));
        explanation.add("Times: " + Arrays.deepToString(times));
        
        if (times.length == 0) {
            explanation.add("No times provided, returning -1");
            return new NetworkDelayResult(-1, new int[0], explanation);
        }
        
        // Build adjacency matrix
        int[][] dist = new int[n + 1][n + 1];
        for (int i = 0; i <= n; i++) {
            Arrays.fill(dist[i], Integer.MAX_VALUE);
            dist[i][i] = 0;
        }
        
        explanation.add("Initialized distance matrix with MAX_VALUE");
        
        // Fill direct edges
        for (int[] time : times) {
            int from = time[0], to = time[1], weight = time[2];
            dist[from][to] = weight;
            explanation.add(String.format("Direct edge: %d -> %d = %d", from, to, weight));
        }
        
        // Floyd-Warshall
        explanation.add("Running Floyd-Warshall algorithm...");
        
        for (int intermediate = 1; intermediate <= n; intermediate++) {
            for (int from = 1; from <= n; from++) {
                for (int to = 1; to <= n; to++) {
                    if (dist[from][intermediate] != Integer.MAX_VALUE && 
                        dist[intermediate][to] != Integer.MAX_VALUE) {
                        int newDist = dist[from][intermediate] + dist[intermediate][to];
                        if (newDist < dist[from][to]) {
                            dist[from][to] = newDist;
                            explanation.add(String.format("  Updated dist[%d][%d] = %d (via %d)", 
                                from, to, newDist, intermediate));
                        }
                    }
                }
            }
        }
        
        // Find maximum distance from k
        int maxDist = 0;
        for (int i = 1; i <= n; i++) {
            if (i != k && dist[k][i] > maxDist) {
                maxDist = dist[k][i];
            }
        }
        
        explanation.add(String.format("Maximum distance from %d: %d", k, maxDist));
        
        int result = maxDist == Integer.MAX_VALUE ? -1 : maxDist;
        explanation.add(String.format("Final result: %d", result));
        
        return new NetworkDelayResult(result, new int[0], explanation);
    }
    
    // Performance comparison
    public void comparePerformance(int[][] times, int n, int k) {
        System.out.println("=== Performance Comparison ===");
        System.out.printf("Nodes: %d, Source: %d, Edges: %d\n", n, k, times.length);
        
        // Floyd-Warshall
        long startTime = System.nanoTime();
        int result1 = networkDelayTimeDP(times, n, k);
        long endTime = System.nanoTime();
        System.out.printf("Floyd-Warshall: %d (took %d ns)\n", result1, endTime - startTime);
        
        // Dijkstra
        startTime = System.nanoTime();
        int result2 = networkDelayTimeDijkstraDP(times, n, k);
        endTime = System.nanoTime();
        System.out.printf("Dijkstra: %d (took %d ns)\n", result2, endTime - startTime);
        
        // Bellman-Ford
        startTime = System.nanoTime();
        int result3 = networkDelayTimeBellmanFord(times, n, k);
        endTime = System.nanoTime();
        System.out.printf("Bellman-Ford: %d (took %d ns)\n", result3, endTime - startTime);
        
        // Topological DP
        startTime = System.nanoTime();
        int result4 = networkDelayTimeTopologicalDP(times, n, k);
        endTime = System.nanoTime();
        System.out.printf("Topological DP: %d (took %d ns)\n", result4, endTime - startTime);
    }
    
    // Check for negative cycles
    public boolean hasNegativeCycle(int[][] times, int n) {
        int[] dist = new int[n + 1];
        Arrays.fill(dist, 0);
        
        for (int i = 0; i < n; i++) {
            boolean updated = false;
            
            for (int[] time : times) {
                int from = time[0], to = time[1], weight = time[2];
                
                if (dist[from] + weight < dist[to]) {
                    dist[to] = dist[from] + weight;
                    updated = true;
                }
            }
            
            if (i == n - 1 && updated) {
                return true; // Negative cycle detected
            }
        }
        
        return false;
    }
    
    // Find all pairs shortest paths
    public int[][] allPairsShortestPaths(int[][] times, int n) {
        int[][] dist = new int[n + 1][n + 1];
        
        // Initialize
        for (int i = 0; i <= n; i++) {
            Arrays.fill(dist[i], Integer.MAX_VALUE);
            dist[i][i] = 0;
        }
        
        // Fill direct edges
        for (int[] time : times) {
            int from = time[0], to = time[1], weight = time[2];
            dist[from][to] = weight;
        }
        
        // Floyd-Warshall
        for (int intermediate = 1; intermediate <= n; intermediate++) {
            for (int from = 1; from <= n; from++) {
                for (int to = 1; to <= n; to++) {
                    if (dist[from][intermediate] != Integer.MAX_VALUE && 
                        dist[intermediate][to] != Integer.MAX_VALUE) {
                        int newDist = dist[from][intermediate] + dist[intermediate][to];
                        if (newDist < dist[from][to]) {
                            dist[from][to] = newDist;
                        }
                    }
                }
            }
        }
        
        return dist;
    }
    
    public static void main(String[] args) {
        NetworkDelayTime ndt = new NetworkDelayTime();
        
        // Test cases
        int[][][] testCases = {
            {{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
            {{2, 1, 1}, {2, 3, 1}, {3, 4, 1}, {4, 5, 1}},
            {{1, 2, 1}},
            {{1, 2, 1}, {1, 3, 1}, {3, 4, 1}, {4, 5}, 1}},
            {{2, 1, 1}, {2, 3, 10}, {3, 4, 1}, {3, 5, 1}},
            {{1, 2, 1}, {2, 3, 2}, {3, 4, 1}, {4, 5, 1}},
            {{1, 2, 5}, {2, 3, 1}, {3, 4, 1}, {4, 5, 10}}
        };
        
        int[] nValues = {4, 5, 2, 5, 4, 5, 5};
        int[] kValues = {2, 2, 2, 1, 3, 3, 1};
        
        String[] descriptions = {
            "Standard case",
            "Longer path",
            "Single edge",
            "Multiple paths",
            "Mixed weights",
            "Increasing weights",
            "Large weight"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("Times: %s, N=%d, K=%d\n", 
                Arrays.deepToString(testCases[i]), nValues[i], kValues[i]);
            
            int result1 = ndt.networkDelayTimeDP(testCases[i], nValues[i], kValues[i]);
            int result2 = ndt.networkDelayTimeDijkstraDP(testCases[i], nValues[i], kValues[i]);
            int result3 = ndt.networkDelayTimeBellmanFord(testCases[i], nValues[i], kValues[i]);
            int result4 = ndt.networkDelayTimeTopologicalDP(testCases[i], nValues[i], kValues[i]);
            
            System.out.printf("Floyd-Warshall: %d\n", result1);
            System.out.printf("Dijkstra: %d\n", result2);
            System.out.printf("Bellman-Ford: %d\n", result3);
            System.out.printf("Topological DP: %d\n", result4);
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        NetworkDelayResult detailedResult = ndt.networkDelayTimeDetailed(
            new int[][]{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2);
        
        System.out.printf("Result: %d\n", detailedResult.delay);
        for (String step : detailedResult.explanation) {
            System.out.println("  " + step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        int[][] performanceTest = new int[100][];
        Random rand = new Random();
        for (int i = 0; i < 100; i++) {
            performanceTest[i] = new int[]{
                rand.nextInt(50) + 1,  // from
                rand.nextInt(50) + 1,  // to
                rand.nextInt(10) + 1    // weight
            };
        }
        
        ndt.comparePerformance(performanceTest, 50, 1);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        System.out.printf("Empty times: %d\n", 
            ndt.networkDelayTimeDP(new int[0][], 0, 0));
        System.out.printf("Single node: %d\n", 
            ndt.networkDelayTimeDP(new int[0][], 1, 1));
        System.out.printf("Unreachable: %d\n", 
            ndt.networkDelayTimeDP(new int[][]{{1, 2, 1}}, 3, 1));
        
        // All pairs shortest paths
        System.out.println("\n=== All Pairs Shortest Paths ===");
        int[][] allPairs = ndt.allPairsShortestPaths(
            new int[][]{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4);
        
        System.out.println("Distance matrix:");
        for (int i = 1; i <= 4; i++) {
            System.out.printf("Node %d: %s\n", i, Arrays.toString(allPairs[i]));
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: DP on Graphs
- **Floyd-Warshall**: All-pairs shortest paths with DP
- **Dijkstra**: Single-source shortest paths with priority queue
- **Bellman-Ford**: Handles negative weights with DP
- **Topological DP**: Uses DAG properties for optimization

## 2. PROBLEM CHARACTERISTICS
- **Weighted Graph**: Directed graph with edge weights
- **Network Delay**: Time for signal to reach all nodes
- **Source Node**: Signal starts from node k
- **Shortest Paths**: Find maximum distance from source

## 3. SIMILAR PROBLEMS
- Shortest Path in Weighted Graph
- Network Delay Time II
- Find the Cheapest Flights Within K Stops
- Minimum Cost to Connect All Points

## 4. KEY OBSERVATIONS
- Floyd-Warshall: O(N³) time, O(N²) space
- Dijkstra: O((V+E) log V) time, O(V+E) space
- Bellman-Ford: O(VE) time, O(V) space
- Topological: O(V+E) time for DAGs
- Multiple approaches for different constraints

## 5. VARIATIONS & EXTENSIONS
- Multiple sources
- Negative cycles detection
- Path reconstruction
- Dynamic updates

## 6. INTERVIEW INSIGHTS
- Clarify: "Are edge weights always positive?"
- Edge cases: disconnected graph, negative cycles
- Time complexity: Choose based on graph properties
- Space complexity: O(N²) vs O(N+E)

## 7. COMMON MISTAKES
- Using Floyd-Warshall for sparse graphs
- Not handling unreachable nodes
- Incorrect distance initialization
- Forgetting negative cycle detection
- Wrong algorithm choice for graph type

## 8. OPTIMIZATION STRATEGIES
- Use Dijkstra for positive weights
- Use Bellman-Ford for negative weights
- Use Floyd-Warshall for dense graphs
- Early termination in Bellman-Ford

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the slowest communication in a network:**
- You have a network (graph) with signal travel times
- Signal starts from source node k
- Each edge represents communication delay
- You want to know the maximum time for signal to reach all nodes
- This is like finding the bottleneck in communication network
- Different algorithms work for different network characteristics

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: times array (u,v,w), number of nodes n, source k
2. **Goal**: Find maximum time for signal to reach all nodes
3. **Output**: Maximum distance from source, or -1 if unreachable

#### Phase 2: Key Insight Recognition
- **"What defines the problem?"** → Weighted directed graph
- **"What do we need?"** → Shortest paths from source
- **"Which algorithm?"** → Depends on edge weights and graph structure
- **"How to handle unreachable?"** → Return -1

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use Floyd-Warshall for all-pairs:
1. Build adjacency matrix with INF distances
2. Set diagonal to 0 (distance to self)
3. Fill direct edges from input
4. For each intermediate node, update all pairs
5. dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
6. After processing all intermediates, find max distance from source"
```

#### Phase 4: Edge Case Handling
- **Empty times**: Return -1 (no edges)
- **Single node**: Return 0 (no delay)
- **Disconnected graph**: Return -1 for unreachable nodes
- **Negative cycles**: Return -1 (infinite delay)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Times: [[2,1,1], [2,3,1], [3,4,1]], N=4, K=2

Human thinking:
"Let's use Floyd-Warshall:

Initialize 5×5 matrix with INF:
[INF, INF, INF, INF, INF]
[INF, INF, INF, INF, INF]
[INF, INF, INF, INF, INF]
[INF, INF, INF, INF, INF]
[INF, INF, INF, INF, INF]

Set diagonal to 0:
[0, INF, INF, INF, INF]
[INF, 0, INF, INF, INF]
[INF, INF, 0, INF, INF]
[INF, INF, INF, 0, INF]
[INF, INF, INF, INF, 0]

Fill direct edges:
dist[2][1] = 1, dist[2][3] = 1, dist[3][4] = 1

Process intermediate k=1:
Update paths through node 1
dist[2][3] = min(1, dist[2][1]+dist[1][3]) = min(1, 1+1) = 1
dist[2][4] = min(INF, dist[2][1]+dist[1][4]) = min(INF, 1+1) = 1

Process intermediate k=2:
Update paths through node 2
dist[2][3] = min(1, dist[2][2]+dist[2][3]) = min(1, 1+1) = 1
dist[2][4] = min(1, dist[2][2]+dist[2][4]) = min(1, INF+1) = 1

Process intermediate k=3:
Update paths through node 3
dist[2][3] = min(1, dist[2][3]+dist[3][3]) = min(1, 1+0) = 1
dist[2][4] = min(1, dist[2][3]+dist[3][4]) = min(1, 1+1) = 1

Final distances from K=2:
To node 1: 1, to node 3: 1, to node 4: 1
Maximum: 1 ✓

Result: 1 ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: DP ensures all intermediate nodes considered
- **Why it's efficient**: O(N³) vs O(N⁴) naive
- **Why it's correct**: Considers all possible paths

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just BFS from source?"** → Only gives single-source, not all pairs
2. **"What about Dijkstra?"** → Better for sparse graphs with positive weights
3. **"How to handle INF?"** → Use value larger than any possible path sum
4. **"What about negative weights?"** → Need Bellman-Ford algorithm

### Real-World Analogy
**Like finding communication delays in a company network:**
- You have offices (nodes) connected by communication lines
- Each line has a delay time (edge weight)
- Message starts from headquarters (source node)
- You want to know the maximum time for message to reach all offices
- Floyd-Warshall finds optimal routes through intermediate offices
- Dijkstra finds fastest routes from headquarters
- This helps identify communication bottlenecks in the network
- Useful in network optimization, logistics planning

### Human-Readable Pseudocode
```
function networkDelayTime(times, n, k):
    if times.isEmpty(): return -1
    
    // Initialize distance matrix
    dist = matrix of size (n+1)×(n+1)
    fill dist with INF
    set dist[i][i] = 0 for all i
    
    // Fill direct edges
    for (u, v, w) in times:
        dist[u][v] = w
    
    // Floyd-Warshall DP
    for intermediate from 1 to n:
        for from from 1 to n:
            for to to 1 to n:
                if dist[from][intermediate] != INF and dist[intermediate][to] != INF:
                    dist[from][to] = min(dist[from][to], 
                                        dist[from][intermediate] + dist[intermediate][to])
    
    // Find maximum distance from source k
    maxDist = 0
    for i from 1 to n:
        if i != k and dist[k][i] > maxDist:
            maxDist = dist[k][i]
    
    return maxDist if maxDist != INF else -1
```

### Execution Visualization

### Example: times=[[2,1,1], [2,3,1], [3,4,1]], n=4, k=2
```
Network Structure:
2 → 1 (weight 1)
2 → 3 (weight 1)
3 → 4 (weight 1)

Floyd-Warshall Process:

Initial: INF everywhere, diagonal = 0

After direct edges:
dist[2][1] = 1, dist[2][3] = INF, dist[3][4] = 1

Intermediate k=1:
dist[2][3] = min(INF, dist[2][1]+dist[1][3]) = min(INF, 1+INF) = 1
dist[2][4] = min(INF, dist[2][1]+dist[1][4]) = min(INF, 1+1) = 1

Intermediate k=2:
dist[2][3] = min(1, dist[2][2]+dist[2][3]) = min(1, INF+1) = 1
dist[2][4] = min(1, dist[2][2]+dist[2][4]) = min(1, 1+1) = 1

Intermediate k=3:
dist[2][3] = min(1, dist[2][3]+dist[3][3]) = min(1, 1+0) = 1
dist[2][4] = min(1, dist[2][3]+dist[3][4]) = min(1, 1+1) = 1

Final distances from K=2:
To 1: 1, To 3: 1, To 4: 1
Maximum: 1 ✓

Visualization:
Step by step: 2→1→3→4 all reachable in time 1
Signal reaches all nodes in maximum time 1
```

### Key Visualization Points:
- **DP transition**: dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])
- **Triple nested loops**: O(N³) time complexity
- **Path reconstruction**: Need additional storage for actual paths
- **Multiple algorithms**: Floyd-Warshall, Dijkstra, Bellman-Ford

### Memory Layout Visualization:
```
Distance Matrix Evolution:
Initial: [INF,INF,INF,INF,INF]
         [INF,0,INF,INF,INF]
         [INF,INF,INF,INF,INF]
         [INF,INF,INF,INF,INF]
         [INF,INF,INF,INF,INF]

After direct edges:
[INF,1,INF,INF,INF]
         [INF,0,INF,INF,INF]
         [INF,INF,INF,INF,INF]
         [INF,INF,INF,1,INF]
         [INF,INF,INF,INF,INF]

After all intermediates:
[INF,1,1,1,INF]
         [INF,0,1,1,INF]
         [INF,INF,0,1,INF]
         [INF,INF,INF,1,INF]
         [INF,INF,INF,INF,0]
         [INF,INF,INF,INF,INF]

Result: max(dist[2][1], dist[2][3], dist[2][4]) = max(1,1,1) = 1
```

### Time Complexity Breakdown:
- **Floyd-Warshall**: O(N³) time, O(N²) space
- **Dijkstra**: O((V+E) log V) time, O(V+E) space
- **Bellman-Ford**: O(VE) time, O(V) space
- **Topological DP**: O(V+E) time for DAGs
- **Algorithm Choice**: Depends on graph density and edge weights
- **Optimal**: Floyd-Warshall for dense graphs, Dijkstra for sparse
*/
