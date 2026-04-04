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
}
