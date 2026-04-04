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
}
