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
}
