# Advanced Graphs & Union-Find (Product-Based Companies)

While standard BFS and DFS are common, FAANG interviews often push candidates to solve complex mapping, connectivity, and optimization problems using **Union-Find (Disjoint Set)**, **Dijkstra's Algorithm** (Shortest Path with weights), and **Minimum Spanning Trees**.

## Question 1: Redundant Connection (Union-Find)
**Problem Statement:** In this problem, a tree is an undirected graph that is connected and has no cycles. You are given a graph that started as a tree with `n` nodes labeled from `1` to `n`, with one additional edge added. The added edge has two different vertices chosen from `1` to `n`, and was not an edge that already existed. The graph is represented as an array `edges` of length `n` where `edges[i] = [ai, bi]` indicates that there is an edge between nodes `ai` and `bi` in the graph. Return an edge that can be removed so that the resulting graph is a tree of `n` nodes.

### Answer:
This is a textbook application of the **Union-Find (Disjoint Set)** data structure. We start with each node in its own set. As we process each edge, we `Union(u, v)`. If `u` and `v` are already in the same set (i.e., `Find(u) == Find(v)`), it means adding this edge creates a cycle, so this is the redundant edge we must return.

**Code Implementation (Java):**
```java
public class RedundantConnection {
    int[] parent;
    
    public int[] findRedundantConnection(int[][] edges) {
        int n = edges.length;
        parent = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            parent[i] = i; // Initially, every node is its own parent (root)
        }
        
        for (int[] edge : edges) {
            if (!union(edge[0], edge[1])) {
                return edge; // Cycle formed!
            }
        }
        
        return new int[0];
    }
    
    // Find with Path Compression
    private int find(int node) {
        if (parent[node] != node) {
            parent[node] = find(parent[node]); 
        }
        return parent[node];
    }
    
    // Union operation. Returns false if nodes are already in the same set.
    private boolean union(int node1, int node2) {
        int root1 = find(node1);
        int root2 = find(node2);
        
        if (root1 == root2) {
            return false; // Already in the same set
        }
        
        parent[root2] = root1; // Merge sets
        return true;
    }
}
```
**Time Complexity:** O(N * α(N)), where α is the inverse Ackermann function, which is nearly constant O(1). Overall essentially O(N).
**Space Complexity:** O(N) for parent array.

---

## Question 2: Accounts Merge (Union-Find)
**Problem Statement:** Given a list of `accounts` where each element `accounts[i]` is a list of strings, where the first element `accounts[i][0]` is a name, and the rest of the elements are emails representing emails of the account. Merge accounts that share the same email.

### Answer:
We can build a graph where emails are nodes. Emails belonging to the same account are connected. Better yet, we use Union-Find. We map each email to an ID (or just map email -> email root). We union all emails within the same account. Finally, we group the emails by their root.

**Code Implementation (Java):**
```java
import java.util.*;

public class AccountsMerge {
    public List<List<String>> accountsMerge(List<List<String>> accounts) {
        Map<String, String> parent = new HashMap<>(); // email -> root email
        Map<String, String> emailToName = new HashMap<>();
        
        // Initialization & linking each email to itself
        for (List<String> account : accounts) {
            String name = account.get(0);
            for (int i = 1; i < account.size(); i++) {
                String email = account.get(i);
                if (!parent.containsKey(email)) {
                    parent.put(email, email);
                }
                emailToName.put(email, name);
                
                // Union the first email in the account with every other email
                if (i > 1) {
                    union(account.get(1), email, parent);
                }
            }
        }
        
        // Group emails by their root
        Map<String, TreeSet<String>> rootToEmails = new HashMap<>();
        for (String email : parent.keySet()) {
            String root = find(email, parent);
            rootToEmails.putIfAbsent(root, new TreeSet<>());
            rootToEmails.get(root).add(email);
        }
        
        // Format the output
        List<List<String>> result = new ArrayList<>();
        for (String root : rootToEmails.keySet()) {
            List<String> account = new ArrayList<>();
            account.add(emailToName.get(root)); // Add name first
            account.addAll(rootToEmails.get(root)); // Add sorted emails
            result.add(account);
        }
        
        return result;
    }
    
    private String find(String s, Map<String, String> parent) {
        if (!parent.get(s).equals(s)) {
            parent.put(s, find(parent.get(s), parent)); // Path compression
        }
        return parent.get(s);
    }
    
    private void union(String s1, String s2, Map<String, String> parent) {
        String root1 = find(s1, parent);
        String root2 = find(s2, parent);
        if (!root1.equals(root2)) {
            parent.put(root2, root1);
        }
    }
}
```
**Time Complexity:** O(N * K * log(N * K)) where N is number of accounts, K is max emails per account. Sorting at the end dominates. Union-Find is nearly linear.
**Space Complexity:** O(N * K)

---

## Question 3: Network Delay Time (Dijkstra's Algorithm)
**Problem Statement:** You are given a network of `n` nodes, labeled from `1` to `n`. You are also given `times`, a list of travel times as directed edges `times[i] = (ui, vi, wi)`, where `ui` is the source node, `vi` is the target node, and `wi` is the time it takes for a signal to travel from source to target. We will send a signal from a given node `k`. Return the minimum time it takes for all the `n` nodes to receive the signal. If it is impossible, return `-1`.

### Answer:
This problem asks for the maximum of the shortest paths from node `k` to all other nodes. This is exactly what **Dijkstra's Algorithm** is designed for (since weights (time) are non-negative).

**Code Implementation (Java):**
```java
import java.util.*;

public class NetworkDelayTime {
    public int networkDelayTime(int[][] times, int n, int k) {
        // Build graph: Map Node -> List of int[] {neighbor, weight}
        Map<Integer, List<int[]>> graph = new HashMap<>();
        for (int[] time : times) {
            graph.putIfAbsent(time[0], new ArrayList<>());
            graph.get(time[0]).add(new int[]{time[1], time[2]});
        }
        
        // Priority Queue (Min-Heap) -> {node, distanceFromSource}
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[1] - b[1]);
        pq.offer(new int[]{k, 0});
        
        Map<Integer, Integer> dist = new HashMap<>();
        
        while (!pq.isEmpty()) {
            int[] info = pq.poll();
            int node = info[0];
            int currentDist = info[1];
            
            if (dist.containsKey(node)) continue; // Already found shortest path to this node
            
            dist.put(node, currentDist);
            
            if (graph.containsKey(node)) {
                for (int[] edge : graph.get(node)) {
                    int neighbor = edge[0];
                    int time = edge[1];
                    if (!dist.containsKey(neighbor)) {
                        pq.offer(new int[]{neighbor, currentDist + time});
                    }
                }
            }
        }
        
        if (dist.size() != n) return -1; // Some nodes are unreachable
        
        int maxTime = 0;
        for (int d : dist.values()) {
            maxTime = Math.max(maxTime, d);
        }
        
        return maxTime;
    }
}
```
**Time Complexity:** O(E log V) where E is number of edges and V is number of vertices.
**Space Complexity:** O(V + E) for Graph and Priority Queue.

---

## Question 4: Min Cost to Connect All Points (Minimum Spanning Tree - Kruskal's/Prim's)
**Problem Statement:** You are given an array `points` representing integer coordinates of some points on a 2D-plane. The cost of connecting two points is the Manhattan distance between them: `|xi - xj| + |yi - yj|`. Return the minimum cost to make all points connected (i.e., exactly one path between any two points).

### Answer:
This requires finding the **Minimum Spanning Tree (MST)**. We can use Prim's Algorithm (best for dense graphs, which this implicitly creates since every point connects to every other point).

**Code Implementation (Java):**
```java
import java.util.PriorityQueue;

public class MinCostConnectPoints {
    public int minCostConnectPoints(int[][] points) {
        int n = points.length;
        // Priority Queue {cost, node_index}
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        boolean[] visited = new boolean[n];
        
        pq.offer(new int[]{0, 0}); // start with cost 0 at point 0
        int cost = 0;
        int edgesUsed = 0;
        
        while (edgesUsed < n) {
            int[] current = pq.poll();
            int currCost = current[0];
            int currNode = current[1];
            
            if (visited[currNode]) continue; // Skip if already included in MST
            
            visited[currNode] = true;
            cost += currCost;
            edgesUsed++;
            
            // Add all edges from current node to priority queue
            for (int nextNode = 0; nextNode < n; nextNode++) {
                if (!visited[nextNode]) {
                    int dist = Math.abs(points[currNode][0] - points[nextNode][0]) + 
                               Math.abs(points[currNode][1] - points[nextNode][1]);
                    pq.offer(new int[]{dist, nextNode});
                }
            }
        }
        
        return cost;
    }
}
```
**Time Complexity:** O(N^2 log N). There are N^2 edges total.
**Space Complexity:** O(N^2) for Priority Queue worst case.
