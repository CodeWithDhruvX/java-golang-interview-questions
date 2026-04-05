import java.util.*;

public class UnionFindProblems {
    
    // 547. Number of Provinces
    // Time: O(N²), Space: O(N)
    public static int findCircleNum(int[][] isConnected) {
        if (isConnected == null || isConnected.length == 0) {
            return 0;
        }
        
        int n = isConnected.length;
        UnionFind uf = new UnionFind(n);
        
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                if (isConnected[i][j] == 1) {
                    uf.union(i, j);
                }
            }
        }
        
        return uf.getCount();
    }

    // 684. Redundant Connection
    // Time: O(N α(N)), Space: O(N)
    public static int[] findRedundantConnection(int[][] edges) {
        if (edges == null || edges.length == 0) {
            return new int[0];
        }
        
        int n = edges.length;
        UnionFind uf = new UnionFind(n + 1);
        
        for (int[] edge : edges) {
            int u = edge[0];
            int v = edge[1];
            
            if (!uf.union(u, v)) {
                return new int[]{u, v};
            }
        }
        
        return new int[0];
    }
    
    // Union Find Data Structure
    private static class UnionFind {
        private int[] parent;
        private int[] rank;
        private int count;
        
        public UnionFind(int size) {
            parent = new int[size];
            rank = new int[size];
            count = size;
            
            for (int i = 0; i < size; i++) {
                parent[i] = i;
                rank[i] = 1;
            }
        }
        
        public int find(int x) {
            if (parent[x] != x) {
                parent[x] = find(parent[x]); // Path compression
            }
            return parent[x];
        }
        
        public boolean union(int x, int y) {
            int rootX = find(x);
            int rootY = find(y);
            
            if (rootX == rootY) {
                return false; // Already connected
            }
            
            // Union by rank
            if (rank[rootX] < rank[rootY]) {
                parent[rootX] = rootY;
            } else if (rank[rootX] > rank[rootY]) {
                parent[rootY] = rootX;
            } else {
                parent[rootY] = rootX;
                rank[rootX]++;
            }
            
            count--;
            return true;
        }
        
        public int getCount() {
            return count;
        }
    }

    public static void main(String[] args) {
        // Test cases for findCircleNum
        int[][][] testCases1 = {
            {{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
            {{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
            {{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
            {{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}},
            {{1, 1, 0, 0}, {1, 1, 0, 0}, {0, 0, 1, 1}, {0, 0, 1, 1}},
            {{1, 0, 1, 0}, {0, 1, 0, 1}, {1, 0, 1, 0}, {0, 1, 0, 1}},
            {{1, 1}, {1, 1}},
            {{1, 0}, {0, 1}},
            {{1, 1, 0}, {1, 1, 1}, {0, 1, 1}},
            {{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}
        };
        
        // Test cases for findRedundantConnection
        int[][][] testCases2 = {
            {{1, 2}, {1, 3}, {2, 3}},
            {{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
            {{1, 2}, {2, 3}, {3, 1}},
            {{1, 2}},
            {{1, 2}, {2, 3}, {3, 4}, {4, 1}},
            {{1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}},
            {{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 1}},
            {{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}},
            {{1, 2}, {1, 3}, {2, 3}},
            {{1, 2}, {1, 3}, {3, 4}, {1, 4}}
        };
        
        System.out.println("Number of Provinces:");
        for (int i = 0; i < testCases1.length; i++) {
            int[][] isConnected = testCases1[i];
            int result = findCircleNum(isConnected);
            System.out.printf("Test Case %d: %s -> %d provinces\n", 
                i + 1, Arrays.deepToString(isConnected), result);
        }
        
        System.out.println("\nRedundant Connection:");
        for (int i = 0; i < testCases2.length; i++) {
            int[][] edges = testCases2[i];
            int[] result = findRedundantConnection(edges);
            System.out.printf("Test Case %d: %s -> %s\n", 
                i + 1, Arrays.deepToString(edges), 
                result.length == 0 ? "[]" : Arrays.toString(result));
        }
    }
}
