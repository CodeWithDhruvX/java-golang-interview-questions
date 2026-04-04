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
}
