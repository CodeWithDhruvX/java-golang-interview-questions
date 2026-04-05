import java.util.*;

public class RedundantConnection {
    
    // 684. Redundant Connection - Strongly Connected Components
    // Time: O(N^2), Space: O(N)
    public int[] findRedundantConnection(int[][] edges) {
        if (edges.length == 0) {
            return new int[0];
        }
        
        // Build adjacency list
        List<Integer>[] adj = new ArrayList[edges.length + 1];
        for (int i = 0; i <= edges.length; i++) {
            adj[i] = new ArrayList<>();
        }
        
        for (int[] edge : edges) {
            int from = edge[0], to = edge[1];
            adj[from].add(to);
            adj[to].add(from);
        }
        
        // Find redundant connection using Union-Find
        int[] parent = new int[edges.length + 1];
        for (int i = 0; i <= edges.length; i++) {
            parent[i] = i;
        }
        
        for (int[] edge : edges) {
            int from = edge[0], to = edge[1];
            
            // Find roots
            int rootFrom = find(parent, from);
            int rootTo = find(parent, to);
            
            // If already connected, this edge is redundant
            if (rootFrom == rootTo) {
                return edge;
            }
            
            // Union sets
            parent[rootFrom] = rootTo;
        }
        
        return new int[0];
    }
    
    private int find(int[] parent, int x) {
        if (parent[x] != x) {
            parent[x] = find(parent, parent[x]);
        }
        return parent[x];
    }
    
    // Strongly Connected Components with Kosaraju's algorithm
    public int[] findRedundantConnectionSCC(int[][] edges) {
        if (edges.length == 0) {
            return new int[0];
        }
        
        // Build adjacency list
        List<Integer>[] adj = new ArrayList[edges.length + 1];
        List<Integer>[] reverseAdj = new ArrayList[edges.length + 1];
        
        for (int i = 0; i <= edges.length; i++) {
            adj[i] = new ArrayList<>();
            reverseAdj[i] = new ArrayList<>();
        }
        
        for (int[] edge : edges) {
            int from = edge[0], to = edge[1];
            adj[from].add(to);
            reverseAdj[to].add(from);
        }
        
        // Kosaraju's algorithm to find SCCs
        boolean[] visited = new boolean[edges.length + 1];
        List<List<Integer>> sccs = kosaraju(adj, reverseAdj, visited);
        
        // If graph is already strongly connected, no redundant edges
        if (sccs.size() == 1) {
            return new int[0];
        }
        
        // Find the first edge that creates a cycle
        for (int[] edge : edges) {
            if (createsCycle(adj, edge[0], edge[1], new boolean[edges.length + 1])) {
                return edge;
            }
        }
        
        return new int[0];
    }
    
    private List<List<Integer>> kosaraju(List<Integer>[] adj, List<Integer>[] reverseAdj, boolean[] visited) {
        int n = adj.length;
        List<List<Integer>> sccs = new ArrayList<>();
        
        // First pass: order vertices by finish time
        List<Integer> order = new ArrayList<>();
        for (int i = 1; i < n; i++) {
            if (!visited[i]) {
                dfs1(i, adj, visited, order);
            }
        }
        
        // Second pass: process vertices in reverse order
        Arrays.fill(visited, false);
        for (int i = order.size() - 1; i >= 0; i--) {
            int vertex = order.get(i);
            if (!visited[vertex]) {
                List<Integer> scc = new ArrayList<>();
                dfs2(vertex, reverseAdj, visited, scc);
                sccs.add(scc);
            }
        }
        
        return sccs;
    }
    
    private void dfs1(int vertex, List<Integer>[] adj, boolean[] visited, List<Integer> order) {
        visited[vertex] = true;
        for (int neighbor : adj[vertex]) {
            if (!visited[neighbor]) {
                dfs1(neighbor, adj, visited, order);
            }
        }
        order.add(vertex);
    }
    
    private void dfs2(int vertex, List<Integer>[] adj, boolean[] visited, List<Integer> scc) {
        visited[vertex] = true;
        for (int neighbor : adj[vertex]) {
            if (!visited[neighbor]) {
                dfs2(neighbor, adj, visited, scc);
            }
        }
        scc.add(vertex);
    }
    
    private boolean createsCycle(List<Integer>[] adj, int from, int to, boolean[] visited) {
        // Check if adding edge from->to creates a cycle
        Arrays.fill(visited, false);
        return hasPathDFS(adj, to, from, visited);
    }
    
    private boolean hasPathDFS(List<Integer>[] adj, int start, int target, boolean[] visited) {
        if (start == target) {
            return true;
        }
        
        visited[start] = true;
        for (int neighbor : adj[start]) {
            if (!visited[neighbor]) {
                if (hasPathDFS(adj, neighbor, target, visited)) {
                    return true;
                }
            }
        }
        
        return false;
    }
    
    // Tarjan's algorithm for SCC
    public int[] findRedundantConnectionTarjan(int[][] edges) {
        if (edges.length == 0) {
            return new int[0];
        }
        
        // Build adjacency list
        List<Integer>[] adj = new ArrayList[edges.length + 1];
        for (int i = 0; i <= edges.length; i++) {
            adj[i] = new ArrayList<>();
        }
        
        for (int[] edge : edges) {
            int from = edge[0], to = edge[1];
            adj[from].add(to);
        }
        
        // Tarjan's algorithm
        int[] ids = new int[edges.length + 1];
        int[] low = new int[edges.length + 1];
        boolean[] onStack = new boolean[edges.length + 1];
        Deque<Integer> stack = new ArrayDeque<>();
        
        int id = 0;
        for (int i = 1; i <= edges.length; i++) {
            if (ids[i] == 0) {
                if (tarjanDFS(i, adj, ids, low, onStack, stack, id) > 1) {
                    // Found multiple SCCs, check for redundant edge
                    for (int[] edge : edges) {
                        if (isRedundantEdge(adj, edge[0], edge[1], ids)) {
                            return edge;
                        }
                    }
                }
            }
        }
        
        return new int[0];
    }
    
    private int tarjanDFS(int at, List<Integer>[] adj, int[] ids, int[] low, 
                          boolean[] onStack, Deque<Integer> stack, int id) {
        stack.push(at);
        onStack[at] = true;
        ids[at] = low[at] = ++id;
        
        int sccCount = 0;
        
        for (int to : adj[at]) {
            if (ids[to] == 0) {
                sccCount += tarjanDFS(to, adj, ids, low, onStack, stack, id);
                low[at] = Math.min(low[at], low[to]);
            } else if (onStack[to]) {
                low[at] = Math.min(low[at], ids[to]);
            }
        }
        
        // If at is root of SCC
        if (low[at] == ids[at]) {
            while (!stack.isEmpty() && stack.peek() != at) {
                onStack[stack.pop()] = false;
            }
            if (!stack.isEmpty()) {
                stack.pop();
                onStack[at] = false;
            }
            sccCount++;
        }
        
        return sccCount;
    }
    
    private boolean isRedundantEdge(List<Integer>[] adj, int from, int to, int[] ids) {
        // Check if edge creates a cycle in existing SCC
        return hasPathDFS(adj, to, from, new boolean[adj.length]);
    }
    
    // Version with detailed explanation
    public class RedundantConnectionResult {
        int[] redundantEdge;
        List<String> explanation;
        List<List<Integer>> sccs;
        
        RedundantConnectionResult(int[] redundantEdge, List<String> explanation, List<List<Integer>> sccs) {
            this.redundantEdge = redundantEdge;
            this.explanation = explanation;
            this.sccs = sccs;
        }
    }
    
    public RedundantConnectionResult findRedundantConnectionDetailed(int[][] edges) {
        List<String> explanation = new ArrayList<>();
        explanation.add("=== Redundant Connection using SCC ===");
        explanation.add("Edges: " + Arrays.deepToString(edges));
        
        if (edges.length == 0) {
            explanation.add("No edges provided, returning empty array");
            return new RedundantConnectionResult(new int[0], explanation, new ArrayList<>());
        }
        
        // Build adjacency list
        List<Integer>[] adj = new ArrayList[edges.length + 1];
        List<Integer>[] reverseAdj = new ArrayList[edges.length + 1];
        
        for (int i = 0; i <= edges.length; i++) {
            adj[i] = new ArrayList<>();
            reverseAdj[i] = new ArrayList<>();
        }
        
        for (int[] edge : edges) {
            int from = edge[0], to = edge[1];
            adj[from].add(to);
            reverseAdj[to].add(from);
            explanation.add(String.format("Added edge: %d -> %d", from, to));
        }
        
        // Kosaraju's algorithm
        boolean[] visited = new boolean[edges.length + 1];
        List<List<Integer>> sccs = kosaraju(adj, reverseAdj, visited);
        
        explanation.add("Strongly Connected Components found:");
        for (int i = 0; i < sccs.size(); i++) {
            explanation.add(String.format("  SCC %d: %s", i + 1, sccs.get(i)));
        }
        
        // Check for redundant edge
        for (int[] edge : edges) {
            explanation.add(String.format("Checking edge [%d, %d]:", edge[0], edge[1]));
            
            if (createsCycle(adj, edge[0], edge[1], new boolean[edges.length + 1])) {
                explanation.add(String.format("  Edge [%d, %d] creates cycle - REDUNDANT!", edge[0], edge[1]));
                return new RedundantConnectionResult(edge, explanation, sccs);
            } else {
                explanation.add(String.format("  Edge [%d, %d] does not create cycle", edge[0], edge[1]));
            }
        }
        
        explanation.add("No redundant edges found");
        return new RedundantConnectionResult(new int[0], explanation, sccs);
    }
    
    // Performance comparison
    public void comparePerformance(int[][] edges, int trials) {
        System.out.println("=== Performance Comparison ===");
        System.out.println("Edges: " + Arrays.deepToString(edges));
        System.out.println("Trials: " + trials);
        
        // Union-Find approach
        long startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            findRedundantConnection(edges);
        }
        long endTime = System.nanoTime();
        System.out.printf("Union-Find approach: took %d ns\n", endTime - startTime);
        
        // Kosaraju's algorithm
        startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            findRedundantConnectionSCC(edges);
        }
        endTime = System.nanoTime();
        System.out.printf("Kosaraju's algorithm: took %d ns\n", endTime - startTime);
        
        // Tarjan's algorithm
        startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            findRedundantConnectionTarjan(edges);
        }
        endTime = System.nanoTime();
        System.out.printf("Tarjan's algorithm: took %d ns\n", endTime - startTime);
    }
    
    // Check if graph is strongly connected
    public boolean isStronglyConnected(int[][] edges) {
        if (edges.length == 0) {
            return false;
        }
        
        int n = edges.length + 1;
        List<Integer>[] adj = new ArrayList[n];
        List<Integer>[] reverseAdj = new ArrayList[n];
        
        for (int i = 0; i < n; i++) {
            adj[i] = new ArrayList<>();
            reverseAdj[i] = new ArrayList<>();
        }
        
        for (int[] edge : edges) {
            int from = edge[0], to = edge[1];
            adj[from].add(to);
            reverseAdj[to].add(from);
        }
        
        // Check if all nodes are reachable from any starting node
        boolean[] visited = new boolean[n];
        dfs1(1, adj, visited, new ArrayList<>());
        
        for (int i = 1; i < n; i++) {
            if (!visited[i]) {
                return false;
            }
        }
        
        return true;
    }
    
    public static void main(String[] args) {
        RedundantConnection rc = new RedundantConnection();
        
        // Test cases
        int[][][] testCases = {
            {{1, 2}, {1, 3}, {2, 3}},
            {{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
            {{1, 2}, {2, 3}, {3, 1}},
            {{1, 2}},
            {{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 1}},
            {{1, 2}, {2, 3}, {3, 4}, {4, 1}},
            {{1, 2}, {2, 3}, {3, 1}, {1, 4}},
            {{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 1}}
        };
        
        String[] descriptions = {
            "Triangle with redundant edge",
            "Complete graph minus one edge",
            "Already strongly connected",
            "Single edge",
            "Cycle graph",
            "Square graph",
            "Triangle with extra edge",
            "Hexagon graph"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.println("Edges: " + Arrays.deepToString(testCases[i]));
            
            int[] result1 = rc.findRedundantConnection(testCases[i]);
            int[] result2 = rc.findRedundantConnectionSCC(testCases[i]);
            int[] result3 = rc.findRedundantConnectionTarjan(testCases[i]);
            
            System.out.printf("Union-Find: %s\n", Arrays.toString(result1));
            System.out.printf("Kosaraju: %s\n", Arrays.toString(result2));
            System.out.printf("Tarjan: %s\n", Arrays.toString(result3));
            
            boolean isSC = rc.isStronglyConnected(testCases[i]);
            System.out.printf("Strongly connected: %b\n", isSC);
            
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        RedundantConnectionResult detailedResult = rc.findRedundantConnectionDetailed(
            new int[][]{{1, 2}, {1, 3}, {2, 3}});
        
        System.out.println("Redundant edge: " + Arrays.toString(detailedResult.redundantEdge));
        System.out.println("SCCs: " + detailedResult.sccs);
        
        for (String step : detailedResult.explanation) {
            System.out.println(step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        int[][] performanceEdges = new int[100][];
        for (int i = 0; i < 100; i++) {
            performanceEdges[i] = new int[]{(i % 50) + 1, ((i + 1) % 50) + 1};
        }
        
        rc.comparePerformance(performanceEdges, 1000);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        
        // Empty edges
        int[] emptyResult = rc.findRedundantConnection(new int[0][0]);
        System.out.println("Empty edges: " + Arrays.toString(emptyResult));
        
        // Single edge
        int[] singleResult = rc.findRedundantConnection(new int[][]{{1, 2}});
        System.out.println("Single edge: " + Arrays.toString(singleResult));
        
        // Large graph
        int[][] largeEdges = new int[1000][];
        for (int i = 0; i < 1000; i++) {
            largeEdges[i] = new int[]{(i % 500) + 1, ((i + 1) % 500) + 1};
        }
        
        long startTime = System.nanoTime();
        rc.findRedundantConnection(largeEdges);
        long endTime = System.nanoTime();
        System.out.printf("Large graph (1000 edges): took %d ns\n", endTime - startTime);
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Strongly Connected Components
- **SCC Detection**: Find strongly connected components in directed graph
- **Redundant Edge**: Edge that creates a cycle in existing SCC
- **Kosaraju's Algorithm**: Two-pass DFS for SCC detection
- **Tarjan's Algorithm**: Single-pass DFS with low-link values

## 2. PROBLEM CHARACTERISTICS
- **Directed Graph**: Find redundant connection in directed graph
- **Cycle Detection**: Edge that creates cycle in existing connectivity
- **Strong Connectivity**: All nodes reachable from each other
- **Multiple Algorithms**: Union-Find, Kosaraju, Tarjan

## 3. SIMILAR PROBLEMS
- Critical Connections in Network
- Network Delay Time
- Find All SCCs
- Graph Connectivity Analysis

## 4. KEY OBSERVATIONS
- Redundant edge creates cycle in existing SCC
- Union-Find can detect redundant edges efficiently
- Kosaraju's algorithm uses two DFS passes
- Tarjan's algorithm uses single DFS pass
- Time complexity: O(V+E) for all algorithms

## 5. VARIATIONS & EXTENSIONS
- Different SCC algorithms
- Multiple redundant edges
- Dynamic graph updates
- Weighted directed graphs

## 6. INTERVIEW INSIGHTS
- Clarify: "Is graph directed?"
- Edge cases: empty graph, single node, already connected
- Time complexity: O(V+E) vs O(V*(V+E)) naive
- Space complexity: O(V+E) vs O(V²) adjacency matrix

## 7. COMMON MISTAKES
- Incorrect cycle detection
- Wrong SCC algorithm implementation
- Not handling directed edges properly
- Incorrect Union-Find path compression
- Wrong edge direction handling

## 8. OPTIMIZATION STRATEGIES
- Use Union-Find with path compression
- Implement Kosaraju's or Tarjan's algorithm
- Efficient adjacency list representation
- Early termination when redundant edge found

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like analyzing a computer network:**
- You have computers connected by one-way connections
- Some connections might be redundant (create cycles)
- Need to find which connection is already redundant
- Strongly connected components mean all computers can reach each other
- Redundant edge creates cycle within existing connectivity

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Directed edges representing connections
2. **Goal**: Find redundant connection
3. **Output**: Edge that creates cycle in existing SCC

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(V*(V+E)) to check each edge
- **"How to optimize?"** → Use Union-Find or SCC algorithms
- **"Why Union-Find?"** → Detects cycles efficiently
- **"Why SCC algorithms?"** → Find strongly connected components

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use Union-Find approach:
1. Initialize each node as its own set
2. For each edge (u,v):
   - Find roots of u and v
   - If roots are same, edge creates cycle (redundant)
   - If different, union the sets
3. Return first redundant edge found
4. For SCC detection:
   - Use Kosaraju's or Tarjan's algorithm
   - Find all strongly connected components
   - Check edges that create cycles within SCCs"
```

#### Phase 4: Edge Case Handling
- **Empty graph**: Return empty array
- **Single node**: No redundant edges possible
- **Already connected**: All edges might be redundant
- **Disconnected graph**: Handle multiple components

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Edges: [[1,2], [2,3], [3,1]]

Human thinking:
"Let's apply Union-Find:

Step 1: Initialize
parent = [0,1,2,3] (each node is its own parent)

Step 2: Process edge [1,2]
- Find root of 1: find(1) = 1
- Find root of 2: find(2) = 2
- Roots are different: union(1,2)
- parent[1] = 2

Step 3: Process edge [2,3]
- Find root of 2: find(2) = 2 (parent[2]=2)
- Find root of 3: find(3) = 3
- Roots are different: union(2,3)
- parent[2] = 3

Step 4: Process edge [3,1]
- Find root of 3: find(3) = 3 (parent[3]=3)
- Find root of 1: find(1) = 3 (parent[1]=2, parent[2]=3)
- Roots are same: edge [3,1] creates cycle!
- Return [3,1] as redundant edge ✓

Manual verification:
1→2→3→1 forms cycle ✓
Edge [3,1] is redundant ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Union-Find detects cycles efficiently
- **Why it's efficient**: O(V+E) vs O(V*(V+E)) naive
- **Why it's correct**: Cycle detection proves redundancy

### Common Human Pitfalls & How to Avoid Them
1. **"Why not check all paths?"** → O(V*(V+E)) too slow
2. **"What about directed edges?"** → Must handle direction properly
3. **"How to detect cycles?"** → Use Union-Find or DFS
4. **"What about SCCs?"** → Use Kosaraju's or Tarjan's algorithm

### Real-World Analogy
**Like finding redundant network connections:**
- You have a computer network with one-way connections
- Some connections might create unnecessary loops
- Need to find which connection is already redundant
- Strongly connected components are fully interconnected sub-networks
- Redundant edge creates loop within existing connectivity
- Useful in network design, circuit analysis, dependency graphs
- Like finding unnecessary backup connections in a network

### Human-Readable Pseudocode
```
function findRedundantEdge(edges):
    if edges.length == 0:
        return []
    
    n = max vertex number + 1
    parent = [0,1,2,...,n-1]
    
    for edge in edges:
        u = edge[0], v = edge[1]
        rootU = find(parent, u)
        rootV = find(parent, v)
        
        if rootU == rootV:
            return edge  // Redundant edge found
        
        union(parent, rootU, rootV)
    
    return []  // No redundant edge

function find(parent, x):
    if parent[x] != x:
        parent[x] = find(parent, parent[x])
    return parent[x]

function union(parent, x, y):
    parent[x] = y

function kosarajuSCC(adj):
    // First pass: order vertices by finish time
    visited = [false] * n
    order = []
    
    for v from 0 to n-1:
        if !visited[v]:
            dfs1(v, adj, visited, order)
    
    // Second pass: process vertices in reverse order
    visited = [false] * n
    sccs = []
    
    for v in reversed(order):
        if !visited[v]:
            scc = []
            dfs2(v, adj, visited, scc)
            sccs.add(scc)
    
    return sccs
```

### Execution Visualization

### Example: edges=[[1,2],[2,3],[3,1]]
```
Union-Find Process:

Initialize:
parent = [0,1,2,3]

Process edge [1,2]:
- root(1) = 1, root(2) = 2
- Different: union(1,2)
- parent[1] = 2

Process edge [2,3]:
- root(2) = 2, root(3) = 3
- Different: union(2,3)
- parent[2] = 3

Process edge [3,1]:
- root(3) = 3, root(1) = 3 (via path 1→2→3)
- Same: edge [3,1] creates cycle!
- Return [3,1] as redundant edge ✓

Cycle verification:
1→2→3→1 forms directed cycle ✓
Edge [3,1] is redundant ✓

Visualization:
Union-Find efficiently detects cycles
Redundant edge creates cycle in existing connectivity
Strongly connected component is {1,2,3}
```

### Key Visualization Points:
- **Union-Find**: Detects cycles efficiently
- **Path Compression**: Optimizes find operations
- **Cycle Detection**: Same root indicates cycle
- **SCC Algorithms**: Find strongly connected components

### Memory Layout Visualization:
```
Graph: 1→2→3→1

Union-Find Evolution:
Start: parent=[0,1,2,3]
After [1,2]: parent=[2,1,2,3]
After [2,3]: parent=[2,3,3,3]
After [3,1]: parent=[2,3,3,3] (cycle detected)

SCC: {1,2,3} (all nodes reachable from each other)
Redundant edge: [3,1] (creates cycle)

Union-Find efficiently detects redundant connections
SCC analysis reveals strongly connected components
```

### Time Complexity Breakdown:
- **Union-Find**: O(V+E) time, O(V) space
- **Kosaraju's**: O(V+E) time, O(V+E) space
- **Tarjan's**: O(V+E) time, O(V+E) space
- **Optimal**: Best possible for this problem
- **vs Naive**: O(V*(V+E)) vs O(V+E) with Union-Find
*/
