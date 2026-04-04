import java.util.*;

public class EvaluateDivision {
    
    // 399. Evaluate Division - Floyd-Warshall for Division Relationships
    // Time: O(N^3 + Q), Space: O(N^2)
    public double[] calcEquation(List<List<String>> equations, double[] values, List<List<String>> queries) {
        // Build variable index map
        Map<String, Integer> varMap = new HashMap<>();
        int idx = 0;
        
        for (List<String> eq : equations) {
            if (!varMap.containsKey(eq.get(0))) {
                varMap.put(eq.get(0), idx++);
            }
            if (!varMap.containsKey(eq.get(1))) {
                varMap.put(eq.get(1), idx++);
            }
        }
        
        int n = varMap.size();
        
        // Initialize distance matrix
        double[][] dist = new double[n][n];
        for (int i = 0; i < n; i++) {
            Arrays.fill(dist[i], -1.0);
            dist[i][i] = 1.0;
        }
        
        // Fill direct relationships
        for (int i = 0; i < equations.size(); i++) {
            List<String> eq = equations.get(i);
            int from = varMap.get(eq.get(0));
            int to = varMap.get(eq.get(1));
            dist[from][to] = values[i];
            dist[to][from] = 1.0 / values[i];
        }
        
        // Floyd-Warshall for division relationships
        for (int k = 0; k < n; k++) {
            for (int i = 0; i < n; i++) {
                for (int j = 0; j < n; j++) {
                    if (dist[i][k] > 0 && dist[k][j] > 0) {
                        double product = dist[i][k] * dist[k][j];
                        if (dist[i][j] < 0 || product < dist[i][j]) {
                            dist[i][j] = product;
                        }
                    }
                }
            }
        }
        
        // Answer queries
        double[] result = new double[queries.size()];
        for (int i = 0; i < queries.size(); i++) {
            List<String> query = queries.get(i);
            if (varMap.containsKey(query.get(0)) && varMap.containsKey(query.get(1))) {
                int fromIdx = varMap.get(query.get(0));
                int toIdx = varMap.get(query.get(1));
                result[i] = dist[fromIdx][toIdx];
            } else {
                result[i] = -1.0;
            }
        }
        
        return result;
    }
    
    // Floyd-Warshall with path tracking
    public class DivisionResult {
        double[] results;
        Map<String, List<String>> paths;
        
        DivisionResult(double[] results, Map<String, List<String>> paths) {
            this.results = results;
            this.paths = paths;
        }
    }
    
    public DivisionResult calcEquationWithPathTracking(List<List<String>> equations, double[] values, List<List<String>> queries) {
        // Build variable index map
        Map<String, Integer> varMap = new HashMap<>();
        int idx = 0;
        
        for (List<String> eq : equations) {
            if (!varMap.containsKey(eq.get(0))) {
                varMap.put(eq.get(0), idx++);
            }
            if (!varMap.containsKey(eq.get(1))) {
                varMap.put(eq.get(1), idx++);
            }
        }
        
        int n = varMap.size();
        
        // Initialize matrices
        double[][] dist = new double[n][n];
        String[][] path = new String[n][n];
        
        for (int i = 0; i < n; i++) {
            Arrays.fill(dist[i], -1.0);
            dist[i][i] = 1.0;
            path[i][i] = (String) varMap.keySet().toArray()[i];
        }
        
        // Fill direct relationships
        for (int i = 0; i < equations.size(); i++) {
            List<String> eq = equations.get(i);
            int from = varMap.get(eq.get(0));
            int to = varMap.get(eq.get(1));
            dist[from][to] = values[i];
            dist[to][from] = 1.0 / values[i];
            path[from][to] = eq.get(1);
            path[to][from] = eq.get(0);
        }
        
        // Floyd-Warshall with path reconstruction
        for (int k = 0; k < n; k++) {
            for (int i = 0; i < n; i++) {
                for (int j = 0; j < n; j++) {
                    if (dist[i][k] > 0 && dist[k][j] > 0) {
                        double product = dist[i][k] * dist[k][j];
                        if (dist[i][j] < 0 || product < dist[i][j]) {
                            dist[i][j] = product;
                            path[i][j] = path[i][k];
                        }
                    }
                }
            }
        }
        
        // Answer queries and build paths
        double[] results = new double[queries.size()];
        Map<String, List<String>> pathsMap = new HashMap<>();
        
        for (int i = 0; i < queries.size(); i++) {
            List<String> query = queries.get(i);
            String queryKey = query.get(0) + "/" + query.get(1);
            
            if (varMap.containsKey(query.get(0)) && varMap.containsKey(query.get(1))) {
                int fromIdx = varMap.get(query.get(0));
                int toIdx = varMap.get(query.get(1));
                results[i] = dist[fromIdx][toIdx];
                
                // Reconstruct path
                List<String> pathList = reconstructPath(path, fromIdx, toIdx, varMap);
                pathsMap.put(queryKey, pathList);
            } else {
                results[i] = -1.0;
                pathsMap.put(queryKey, new ArrayList<>());
            }
        }
        
        return new DivisionResult(results, pathsMap);
    }
    
    private List<String> reconstructPath(String[][] path, int from, int to, Map<Integer, String> idxToVar) {
        List<String> pathList = new ArrayList<>();
        
        if (path[from][to] == null) {
            return pathList;
        }
        
        // Get variable names from indices
        String[] varNames = new String[path.length];
        for (Map.Entry<Integer, String> entry : idxToVar.entrySet()) {
            varNames[entry.getKey()] = entry.getValue();
        }
        
        pathList.add(varNames[from]);
        int current = from;
        
        while (current != to && path[current][to] != null) {
            String nextVar = path[current][to];
            // Find index of next variable
            int nextIdx = -1;
            for (int i = 0; i < varNames.length; i++) {
                if (varNames[i].equals(nextVar)) {
                    nextIdx = i;
                    break;
                }
            }
            
            if (nextIdx == -1 || nextIdx == current) {
                break;
            }
            
            pathList.add(nextVar);
            current = nextIdx;
        }
        
        return pathList;
    }
    
    // Union-Find approach for comparison
    public double[] calcEquationUnionFind(List<List<String>> equations, double[] values, List<List<String>> queries) {
        // Build variable index map
        Map<String, Integer> varMap = new HashMap<>();
        int idx = 0;
        
        for (List<String> eq : equations) {
            if (!varMap.containsKey(eq.get(0))) {
                varMap.put(eq.get(0), idx++);
            }
            if (!varMap.containsKey(eq.get(1))) {
                varMap.put(eq.get(1), idx++);
            }
        }
        
        int n = varMap.size();
        
        // Union-Find with weight
        UnionFind uf = new UnionFind(n);
        
        // Build relationships
        for (int i = 0; i < equations.size(); i++) {
            List<String> eq = equations.get(i);
            int from = varMap.get(eq.get(0));
            int to = varMap.get(eq.get(1));
            uf.union(from, to, values[i]);
        }
        
        // Answer queries
        double[] result = new double[queries.size()];
        for (int i = 0; i < queries.size(); i++) {
            List<String> query = queries.get(i);
            
            if (!varMap.containsKey(query.get(0)) || !varMap.containsKey(query.get(1))) {
                result[i] = -1.0;
                continue;
            }
            
            int fromIdx = varMap.get(query.get(0));
            int toIdx = varMap.get(query.get(1));
            
            if (uf.isConnected(fromIdx, toIdx)) {
                result[i] = uf.getWeight(fromIdx, toIdx);
            } else {
                result[i] = -1.0;
            }
        }
        
        return result;
    }
    
    // Union-Find with weight
    private static class UnionFind {
        private int[] parent;
        private double[] weight;
        
        public UnionFind(int n) {
            parent = new int[n];
            weight = new double[n];
            for (int i = 0; i < n; i++) {
                parent[i] = i;
                weight[i] = 1.0;
            }
        }
        
        public int find(int x) {
            if (parent[x] != x) {
                int origParent = parent[x];
                parent[x] = find(parent[x]);
                weight[x] *= weight[origParent];
            }
            return parent[x];
        }
        
        public void union(int x, int y, double value) {
            int rootX = find(x);
            int rootY = find(y);
            
            if (rootX != rootY) {
                parent[rootX] = rootY;
                weight[rootX] = value * weight[y] / weight[x];
            }
        }
        
        public boolean isConnected(int x, int y) {
            return find(x) == find(y);
        }
        
        public double getWeight(int x, int y) {
            int rootX = find(x);
            int rootY = find(y);
            
            if (rootX != rootY) {
                return -1.0;
            }
            
            return weight[x] / weight[y];
        }
    }
    
    // Version with detailed explanation
    public class DivisionAnalysis {
        double[] results;
        List<String> explanation;
        Map<String, Double> variableValues;
        
        DivisionAnalysis(double[] results, List<String> explanation, Map<String, Double> variableValues) {
            this.results = results;
            this.explanation = explanation;
            this.variableValues = variableValues;
        }
    }
    
    public DivisionAnalysis analyzeEquations(List<List<String>> equations, double[] values, List<List<String>> queries) {
        List<String> explanation = new ArrayList<>();
        explanation.add("=== Division Equation Analysis ===");
        
        // Build variable index map
        Map<String, Integer> varMap = new HashMap<>();
        Map<Integer, String> idxToVar = new HashMap<>();
        int idx = 0;
        
        for (List<String> eq : equations) {
            if (!varMap.containsKey(eq.get(0))) {
                varMap.put(eq.get(0), idx);
                idxToVar.put(idx, eq.get(0));
                idx++;
            }
            if (!varMap.containsKey(eq.get(1))) {
                varMap.put(eq.get(1), idx);
                idxToVar.put(idx, eq.get(1));
                idx++;
            }
            explanation.add(String.format("Equation: %s / %s = %.2f", eq.get(0), eq.get(1), values[equations.indexOf(eq)]));
        }
        
        int n = varMap.size();
        explanation.add(String.format("Total variables: %d", n));
        
        // Initialize distance matrix
        double[][] dist = new double[n][n];
        for (int i = 0; i < n; i++) {
            Arrays.fill(dist[i], -1.0);
            dist[i][i] = 1.0;
        }
        
        explanation.add("Initialized distance matrix with diagonal = 1.0");
        
        // Fill direct relationships
        for (int i = 0; i < equations.size(); i++) {
            List<String> eq = equations.get(i);
            int from = varMap.get(eq.get(0));
            int to = varMap.get(eq.get(1));
            dist[from][to] = values[i];
            dist[to][from] = 1.0 / values[i];
            
            explanation.add(String.format("Direct relationship: %s -> %s = %.2f, %s -> %s = %.2f", 
                eq.get(0), eq.get(1), values[i], eq.get(1), eq.get(0), 1.0 / values[i]));
        }
        
        // Floyd-Warshall
        explanation.add("Running Floyd-Warshall algorithm...");
        for (int k = 0; k < n; k++) {
            for (int i = 0; i < n; i++) {
                for (int j = 0; j < n; j++) {
                    if (dist[i][k] > 0 && dist[k][j] > 0) {
                        double product = dist[i][k] * dist[k][j];
                        if (dist[i][j] < 0 || product < dist[i][j]) {
                            dist[i][j] = product;
                            if (i < n && j < n && k < n) {
                                explanation.add(String.format("  Updated: %s -> %s = %.2f (via %s)", 
                                    idxToVar.get(i), idxToVar.get(j), product, idxToVar.get(k)));
                            }
                        }
                    }
                }
            }
        }
        
        // Answer queries
        double[] results = new double[queries.size()];
        explanation.add("Answering queries:");
        
        for (int i = 0; i < queries.size(); i++) {
            List<String> query = queries.get(i);
            explanation.add(String.format("Query %d: %s / %s", i + 1, query.get(0), query.get(1)));
            
            if (varMap.containsKey(query.get(0)) && varMap.containsKey(query.get(1))) {
                int fromIdx = varMap.get(query.get(0));
                int toIdx = varMap.get(query.get(1));
                results[i] = dist[fromIdx][toIdx];
                explanation.add(String.format("  Result: %.2f", results[i]));
            } else {
                results[i] = -1.0;
                explanation.add("  Result: -1.0 (variable not found)");
            }
        }
        
        return new DivisionAnalysis(results, explanation, new HashMap<>());
    }
    
    public static void main(String[] args) {
        EvaluateDivision ed = new EvaluateDivision();
        
        // Test cases
        List<List<String>> equations1 = Arrays.asList(
            Arrays.asList("a", "b"),
            Arrays.asList("b", "c")
        );
        double[] values1 = {2.0, 3.0};
        List<List<String>> queries1 = Arrays.asList(
            Arrays.asList("a", "c"),
            Arrays.asList("b", "a"),
            Arrays.asList("a", "e"),
            Arrays.asList("a", "a"),
            Arrays.asList("x", "x")
        );
        
        List<List<String>> equations2 = Arrays.asList(
            Arrays.asList("a", "b"),
            Arrays.asList("e", "f"),
            Arrays.asList("b", "e")
        );
        double[] values2 = {3.4, 1.4, 2.3};
        List<List<String>> queries2 = Arrays.asList(
            Arrays.asList("b", "a"),
            Arrays.asList("a", "f"),
            Arrays.asList("f", "f"),
            Arrays.asList("e", "e"),
            Arrays.asList("c", "c"),
            Arrays.asList("a", "c"),
            Arrays.asList("f", "e")
        );
        
        Object[][] testCases = {
            {equations1, values1, queries1, "Standard case"},
            {equations2, values2, queries2, "Complex case"},
            {Arrays.asList(), new double[]{}, Arrays.asList(Arrays.asList("a", "b")), "Empty equations"},
            {Arrays.asList(Arrays.asList("x", "x")), new double[]{1.0}, Arrays.asList(Arrays.asList("x", "x")), "Self division"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, testCases[i][3]);
            
            @SuppressWarnings("unchecked")
            List<List<String>> eqs = (List<List<String>>) testCases[i][0];
            double[] vals = (double[]) testCases[i][1];
            @SuppressWarnings("unchecked")
            List<List<String>> qs = (List<List<String>>) testCases[i][2];
            
            double[] result1 = ed.calcEquation(eqs, vals, qs);
            double[] result2 = ed.calcEquationUnionFind(eqs, vals, qs);
            
            System.out.printf("Floyd-Warshall: %s\n", Arrays.toString(result1));
            System.out.printf("Union-Find: %s\n", Arrays.toString(result2));
            System.out.println();
        }
        
        // Detailed analysis
        System.out.println("=== Detailed Analysis ===");
        DivisionAnalysis analysis = ed.analyzeEquations(equations1, values1, queries1);
        System.out.printf("Results: %s\n", Arrays.toString(analysis.results));
        for (String step : analysis.explanation) {
            System.out.println(step);
        }
        
        // Path tracking
        System.out.println("\n=== Path Tracking ===");
        DivisionResult pathResult = ed.calcEquationWithPathTracking(equations1, values1, queries1);
        System.out.printf("Results: %s\n", Arrays.toString(pathResult.results));
        for (Map.Entry<String, List<String>> entry : pathResult.paths.entrySet()) {
            System.out.printf("Path %s: %s\n", entry.getKey(), entry.getValue());
        }
    }
}
