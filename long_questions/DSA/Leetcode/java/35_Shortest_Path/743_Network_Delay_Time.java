import java.util.*;

public class NetworkDelayTime {
    
    // Edge represents a weighted edge in the graph
    public static class Edge {
        int to;
        int weight;
        
        public Edge(int to, int weight) {
            this.to = to;
            this.weight = weight;
        }
    }

    // 743. Network Delay Time - Dijkstra's Algorithm
    // Time: O((V + E) log V), Space: O(V + E)
    public static int networkDelayTime(int[][] times, int n, int k) {
        // Build adjacency list
        Map<Integer, List<Edge>> adj = new HashMap<>();
        for (int[] time : times) {
            int from = time[0], to = time[1], weight = time[2];
            adj.computeIfAbsent(from, key -> new ArrayList<>()).add(new Edge(to, weight));
        }
        
        // Dijkstra's algorithm
        int[] dist = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[k] = 0;
        
        // Min-heap: {distance, node}
        PriorityQueue<int[]> minHeap = new PriorityQueue<>((a, b) -> Integer.compare(a[0], b[0]));
        minHeap.offer(new int[]{0, k});
        
        while (!minHeap.isEmpty()) {
            int[] current = minHeap.poll();
            int currentDist = current[0];
            int currentNode = current[1];
            
            // Skip if we've found a better path
            if (currentDist > dist[currentNode]) {
                continue;
            }
            
            // Relax edges
            for (Edge edge : adj.getOrDefault(currentNode, new ArrayList<>())) {
                int newDist = currentDist + edge.weight;
                if (newDist < dist[edge.to]) {
                    dist[edge.to] = newDist;
                    minHeap.offer(new int[]{newDist, edge.to});
                }
            }
        }
        
        // Find the maximum distance
        int maxDist = 0;
        for (int i = 1; i <= n; i++) {
            if (dist[i] == Integer.MAX_VALUE) {
                return -1; // Unreachable node
            }
            maxDist = Math.max(maxDist, dist[i]);
        }
        
        return maxDist;
    }

    // Bellman-Ford algorithm (handles negative weights)
    public static int networkDelayTimeBellmanFord(int[][] times, int n, int k) {
        // Initialize distances
        int[] dist = new int[n + 1];
        Arrays.fill(dist, Integer.MAX_VALUE);
        dist[k] = 0;
        
        // Relax edges V-1 times
        for (int i = 0; i < n - 1; i++) {
            boolean updated = false;
            for (int[] time : times) {
                int from = time[0], to = time[1], weight = time[2];
                
                if (dist[from] != Integer.MAX_VALUE && dist[from] + weight < dist[to]) {
                    dist[to] = dist[from] + weight;
                    updated = true;
                }
            }
            
            // Early termination if no updates
            if (!updated) {
                break;
            }
        }
        
        // Check for negative cycles (not applicable in this problem but good practice)
        for (int[] time : times) {
            int from = time[0], to = time[1], weight = time[2];
            if (dist[from] != Integer.MAX_VALUE && dist[from] + weight < dist[to]) {
                return -1; // Negative cycle detected
            }
        }
        
        // Find the maximum distance
        int maxDist = 0;
        for (int i = 1; i <= n; i++) {
            if (dist[i] == Integer.MAX_VALUE) {
                return -1; // Unreachable node
            }
            maxDist = Math.max(maxDist, dist[i]);
        }
        
        return maxDist;
    }

    // BFS approach for unweighted graphs (not suitable for this problem but educational)
    public static int networkDelayTimeBFS(int[][] times, int n, int k) {
        // Build adjacency list (ignoring weights)
        Map<Integer, List<Integer>> adj = new HashMap<>();
        for (int[] time : times) {
            int from = time[0], to = time[1];
            adj.computeIfAbsent(from, key -> new ArrayList<>()).add(to);
        }
        
        int[] dist = new int[n + 1];
        Arrays.fill(dist, -1);
        dist[k] = 0;
        
        Queue<Integer> queue = new LinkedList<>();
        queue.offer(k);
        
        while (!queue.isEmpty()) {
            int current = queue.poll();
            
            for (int neighbor : adj.getOrDefault(current, new ArrayList<>())) {
                if (dist[neighbor] == -1) {
                    dist[neighbor] = dist[current] + 1;
                    queue.offer(neighbor);
                }
            }
        }
        
        // Find the maximum distance
        int maxDist = 0;
        for (int i = 1; i <= n; i++) {
            if (dist[i] == -1) {
                return -1; // Unreachable node
            }
            maxDist = Math.max(maxDist, dist[i]);
        }
        
        return maxDist;
    }

    public static void main(String[] args) {
        // Test cases
        Object[][] testCases = {
            {new int[][]{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, "Simple network"},
            {new int[][]{{1, 2, 1}}, 2, 1, "Single edge"},
            {new int[][]{{1, 2, 1}, {2, 1, 1}}, 2, 1, "Bidirectional edge"},
            {new int[][]{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}}, 3, 1, "Multiple paths"},
            {new int[][]{{1, 2, 1}, {2, 3, 2}, {3, 4, 1}, {1, 4, 5}}, 4, 1, "Long path vs short path"},
            {new int[][]{{1, 2, 1}, {1, 3, 2}, {2, 4, 1}, {3, 4, 1}}, 4, 1, "Converging paths"},
            {new int[][]{{1, 2, 1}}, 2, 2, "Source at destination"},
            {new int[][]{{1, 2, 1}, {3, 4, 1}}, 4, 1, "Disconnected components"},
            {new int[][]{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}, {4, 5, 1}}, 5, 1, "Linear chain"},
            {new int[][]{{1, 2, 10}, {2, 3, 10}, {3, 4, 10}, {4, 5, 10}}, 5, 1, "High weights"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[][] times = (int[][]) testCases[i][0];
            int n = (int) testCases[i][1];
            int k = (int) testCases[i][2];
            String description = (String) testCases[i][3];
            
            System.out.printf("Test Case %d: %s\n", i + 1, description);
            System.out.printf("  Times: %s, N=%d, K=%d\n", Arrays.deepToString(times), n, k);
            
            int result1 = networkDelayTime(times, n, k);
            int result2 = networkDelayTimeBellmanFord(times, n, k);
            int result3 = networkDelayTimeBFS(times, n, k);
            
            System.out.printf("  Dijkstra: %d\n", result1);
            System.out.printf("  Bellman-Ford: %d\n", result2);
            System.out.printf("  BFS (unweighted): %d\n\n", result3);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Shortest Path in Weighted Graph
- **Dijkstra's Algorithm**: Single-source shortest paths with priority queue
- **Edge Relaxation**: Update distances when shorter path found
- **Priority Queue**: Always expand node with minimum distance
- **Early Termination**: Stop when target reached

## 2. PROBLEM CHARACTERISTICS
- **Weighted Graph**: Directed graph with positive edge weights
- **Network Delay**: Time for signal to reach all nodes
- **Source Node**: Signal starts from node k
- **Shortest Paths**: Find maximum distance from source
- **Non-negative Weights**: Required for Dijkstra's correctness

## 3. SIMILAR PROBLEMS
- Shortest Path in Weighted Graph
- Network Delay Time II
- Find the Cheapest Flights Within K Stops
- Minimum Cost to Connect All Points

## 4. KEY OBSERVATIONS
- Dijkstra's algorithm finds shortest paths from single source
- Priority queue ensures we always expand closest unvisited node
- Edge relaxation updates distances when shorter path found
- Time complexity: O((V+E) log V) with adjacency list
- Space complexity: O(V+E) for graph storage

## 5. VARIATIONS & EXTENSIONS
- Multiple sources
- Path reconstruction
- K stops constraint
- Dynamic edge weights

## 6. INTERVIEW INSIGHTS
- Clarify: "Are edge weights always positive?"
- Edge cases: disconnected graph, source unreachable
- Time complexity: O((V+E) log V) vs O(VE) Bellman-Ford
- Space complexity: O(V+E) vs O(V²) adjacency matrix

## 7. COMMON MISTAKES
- Using BFS for weighted graph (wrong)
- Not handling unreachable nodes
- Incorrect priority queue implementation
- Forgetting to mark nodes as visited
- Wrong distance initialization

## 8. OPTIMIZATION STRATEGIES
- Use adjacency list instead of matrix for sparse graphs
- Early termination when target reached
- Efficient priority queue implementation
- Use Fibonacci heap for better performance

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the fastest communication routes in a network:**
- You have a network (graph) with communication delays
- Signal starts from source node k
- Each edge represents communication time between nodes
- You want to know the maximum time for signal to reach all nodes
- Dijkstra's algorithm is like expanding from the closest unvisited node
- Priority queue ensures we always process the most promising node

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: times array (u,v,w), number of nodes n, source k
2. **Goal**: Find maximum time for signal to reach all nodes
3. **Output**: Maximum distance from source, or -1 if unreachable

#### Phase 2: Key Insight Recognition
- **"What defines the problem?"** → Weighted directed graph
- **"What do we need?"** → Shortest paths from source
- **"Which algorithm?"** → Dijkstra's for positive weights
- **"How to handle unreachable?"** → Return -1

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use Dijkstra's algorithm:
1. Build adjacency list from edge list
2. Initialize distances to INF, source to 0
3. Use priority queue: {distance, node}
4. Add source to queue
5. While queue not empty:
   - Extract node with minimum distance
   - If distance > current distance, skip
   - For each neighbor:
     - Calculate new distance = current distance + edge weight
     - If new distance < neighbor distance:
       - Update neighbor distance
       - Add to priority queue
6. Find maximum distance from source"
```

#### Phase 4: Edge Case Handling
- **Empty times**: Return -1 (no edges)
- **Single node**: Return 0 (no delay)
- **Disconnected graph**: Return -1 for unreachable nodes
- **Source unreachable**: Return -1

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Times: [[2,1,1], [2,3,1], [3,4,1]], N=4, K=2

Human thinking:
"Let's use Dijkstra's algorithm:

Build adjacency list:
2: [1], 3: [4]
3: [4], 4: []

Initialize distances to INF, dist[2] = 0
Priority queue: [(0,2)]

Step 1:
- Extract (0,2) from queue
- Current node: 2, distance: 0
- Neighbors: [4] with weight 1
- New distance to 4: 0 + 1 = 1
- Update dist[4] = 1, add (1,4) to queue
- Queue: [(1,4)]

Step 2:
- Extract (1,4) from queue
- Current node: 4, distance: 1
- Neighbors: []
- No new updates
- Queue: []

Step 3:
- Queue empty, stop

Final distances: [0, INF, 1, 1]
Maximum from K=2: max(1, INF, 1, 1) = 1 ✓

Result: 1 ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Always expands closest unvisited node
- **Why it's efficient**: O((V+E) log V) vs O(V²) adjacency matrix
- **Why it's correct**: Greedy choice is optimal for non-negative weights

### Common Human Pitfalls & How to Avoid Them
1. **"Why not BFS?"** → BFS doesn't work for weighted graphs
2. **"What about Bellman-Ford?"** → Slower O(VE) vs O((V+E) log V)
3. **"How to handle INF?"** → Use value larger than any possible path sum
4. **"What about visited tracking?"** → Priority queue handles this implicitly

### Real-World Analogy
**Like finding the fastest delivery routes in a logistics network:**
- You have cities (nodes) connected by delivery routes (edges)
- Each route has a delivery time (edge weight)
- Packages start from central hub (source node)
- You want to know the maximum time for packages to reach all cities
- Dijkstra's algorithm always chooses the fastest next delivery route
- Priority queue ensures we always process the most promising delivery
- This helps optimize delivery schedules and identify bottlenecks
- Useful in logistics, network optimization, delivery systems

### Human-Readable Pseudocode
```
function networkDelayTime(times, n, k):
    if times.isEmpty(): return -1
    
    // Build adjacency list
    adj = adjacency list of size n+1
    for (u, v, w) in times:
        adj[u].add((v, w))
    
    // Initialize distances
    dist = array of size n+1
    fill dist with INF
    dist[k] = 0
    
    // Priority queue: {distance, node}
    pq = min-heap
    pq.offer((0, k))
    
    while pq not empty:
        (currentDist, currentNode) = pq.poll()
        
        if currentDist > dist[currentNode]:
            continue
        
        for (neighbor, weight) in adj[currentNode]:
            newDist = currentDist + weight
            if newDist < dist[neighbor]:
                dist[neighbor] = newDist
                pq.offer((newDist, neighbor))
    
    // Find maximum distance from source k
    maxDist = 0
    for i from 1 to n:
        if dist[i] > maxDist:
            maxDist = dist[i]
    
    return maxDist if maxDist != INF else -1
```

### Execution Visualization

### Example: times=[[2,1,1], [2,3,1], [3,4,1]], n=4, k=2
```
Network Structure:
2 → 1 (weight 1)
2 → 3 (weight 1)
3 → 4 (weight 1)

Dijkstra's Process:

Initial: dist=[INF,0,INF,INF], pq=[(0,2)]

Step 1:
- Extract (0,2): current=2, dist=0
- Process edge 2→3: newDist=1, dist[3]=1, pq=[(1,3)]
- Process edge 2→1: already visited

Step 2:
- Extract (1,3): current=3, dist=1
- Process edge 3→4: newDist=2, dist[4]=2, pq=[(2,4)]

Step 3:
- Extract (2,4): current=4, dist=2
- No neighbors from 4
- Queue empty

Final distances: [0, INF, 1, 2]
Maximum from K=2: max(0, INF, 1, 2) = 2 ✓

Visualization:
Step by step: 2→3→4
Distances: [0, INF, 1, 2]
Signal reaches all nodes in maximum time 2
```

### Key Visualization Points:
- **Priority queue** always extracts minimum distance node
- **Edge relaxation** updates when shorter path found
- **Greedy choice** is optimal for non-negative weights
- **Early termination** when target reached

### Memory Layout Visualization:
```
Distance Array Evolution:
Initial: [INF,0,INF,INF]
After step 1: [INF,0,1,INF]
After step 2: [INF,0,1,2]
After step 3: [0,INF,1,2]

Priority Queue Evolution:
Step 1: [(0,2)]
Step 2: [(1,3)]
Step 3: [(2,4)]

Adjacency List:
2: [(3,1)]
3: [(4,1)]
4: []

Result: max(dist[2], dist[3], dist[4]) = 2
```

### Time Complexity Breakdown:
- **Edge Processing**: O(E) to build adjacency list
- **Priority Queue Operations**: O(log V) per operation
- **Node Processing**: O(E) total for all edges
- **Total**: O((V+E) log V) time, O(V+E) space
- **Optimal**: Best for sparse graphs with positive weights
- **vs Bellman-Ford**: O(VE) time, O(V) space
*/
