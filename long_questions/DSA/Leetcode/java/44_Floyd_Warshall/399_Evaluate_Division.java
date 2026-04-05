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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Floyd-Warshall Algorithm
- **All-Pairs Shortest Path**: Find shortest paths between all pairs
- **Dynamic Programming**: Build solution iteratively
- **Path Reconstruction**: Track intermediate vertices
- **O(N³) Time**: Cubic time complexity for N vertices

## 2. PROBLEM CHARACTERISTICS
- **Division Queries**: Evaluate a/b using given equations
- **Transitive Properties**: Use transitivity of division
- **Graph Representation**: Variables as vertices, equations as edges
- **Multiple Queries**: Answer many division queries efficiently

## 3. SIMILAR PROBLEMS
- Network Delay Time
- Find All-Pairs Shortest Path
- Transitive Closure
- Graph Connectivity

## 4. KEY OBSERVATIONS
- Division a/b can be represented as edge with weight 1/b
- Transitivity: if a/b and b/c, then a/c = (a/b) * (b/c)
- Floyd-Warshall computes all-pairs shortest paths
- Time complexity: O(N³) for preprocessing, O(1) per query
- Space complexity: O(N²) for distance matrix

## 5. VARIATIONS & EXTENSIONS
- Different aggregation functions
- Weighted graphs with different operations
- Dynamic graph updates
- Multiple constraint types

## 6. INTERVIEW INSIGHTS
- Clarify: "Are all variables connected?"
- Edge cases: division by zero, disconnected components
- Time complexity: O(N³ + Q) vs O(Q*N) per query
- Space complexity: O(N²) vs O(N) for adjacency list

## 7. COMMON MISTAKES
- Incorrect distance initialization
- Wrong transitivity implementation
- Not handling division by zero
- Incorrect path reconstruction
- Not checking for disconnected components

## 8. OPTIMIZATION STRATEGIES
- Use adjacency matrix for O(1) access
- Proper initialization of distance matrix
- Efficient path reconstruction
- Handle special cases early

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like building a multiplication table:**
- You have variables and division relationships (a/b = value)
- Need to answer many division queries efficiently
- Build a complete table of all possible divisions
- Use transitivity: a/c = (a/b) * (b/c)
- Floyd-Warshall systematically improves this table
- This is like precomputing all possible results

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Equations like a/b = value, queries like a/c = ?
2. **Goal**: Answer all queries efficiently
3. **Output**: Division results or -1 if impossible

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(Q*N) per query processing
- **"How to optimize?"** → Precompute all possible divisions
- **"Why Floyd-Warshall?"** → Handles transitivity systematically
- **"How to represent?"** → Graph with weighted edges

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use Floyd-Warshall:
1. Map variables to indices 0 to N-1
2. Initialize distance matrix:
   - dist[i][i] = 1 (a/a = 1)
   - dist[i][j] = value for direct a/b
   - dist[j][i] = 1/value for b/a
3. Apply Floyd-Warshall:
   - For each intermediate k:
     - Update all pairs (i,j) using k as intermediate
     - dist[i][j] = min(dist[i][j], dist[i][k] * dist[k][j])
4. Answer queries in O(1) using precomputed matrix
5. Reconstruct paths if needed"
```

#### Phase 4: Edge Case Handling
- **Division by zero**: Return -1 or handle appropriately
- **Disconnected components**: Return -1 for impossible queries
- **Self-division**: a/a = 1 by definition
- **Missing variables**: Handle gracefully

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Equations: a/b = 2, b/c = 3
Query: a/c = ?

Human thinking:
"Let's apply Floyd-Warshall:

Step 1: Map variables
a→0, b→1, c→2

Step 2: Initialize distance matrix
dist[0][0] = 1, dist[0][1] = 2, dist[0][2] = ?
dist[1][0] = 1/2, dist[1][1] = 1, dist[1][2] = 3
dist[2][0] = ?, dist[2][1] = 1/3, dist[2][2] = 1

Step 3: Floyd-Warshall with k=0 (intermediate = a)
Update pairs using a as intermediate:
dist[1][2] = min(dist[1][2], dist[1][0] * dist[0][2])
dist[1][2] = min(?, 0.5 * ?) = ? (keep ?)

Step 4: Floyd-Warshall with k=1 (intermediate = b)
Update pairs using b as intermediate:
dist[0][2] = min(dist[0][2], dist[0][1] * dist[1][2])
dist[0][2] = min(?, 2 * 3) = min(?, 6) = 6 ✓

Step 5: Floyd-Warshall with k=2 (intermediate = c)
Update pairs using c as intermediate:
All pairs already optimal

Final result: a/c = dist[0][2] = 6 ✓

Manual verification:
a/c = (a/b) * (b/c) = 2 * 3 = 6 ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Floyd-Warshall handles all intermediate vertices
- **Why it's efficient**: O(N³) preprocessing, O(1) queries
- **Why it's correct**: Systematically considers all possible paths

### Common Human Pitfalls & How to Avoid Them
1. **"Why not process queries individually?"** → O(Q*N) too slow
2. **"What about matrix initialization?"** → Must set diagonal to 1
3. **"How to handle transitivity?"** → Use min() for path optimization
4. **"What about disconnected graphs?"** → Return -1 for impossible paths

### Real-World Analogy
**Like building a currency conversion table:**
- You have exchange rates between currencies (variables)
- Need to answer many conversion queries efficiently
- Direct rates: a/b = 2, b/c = 3
- Indirect rates: a/c = (a/b) * (b/c) = 2 * 3 = 6
- Floyd-Warshall builds complete conversion table
- This is like precomputing all currency exchange rates
- Useful in expert systems, recommendation engines, graph analysis
- Like having a complete multiplication table for quick lookups

### Human-Readable Pseudocode
```
function evaluateDivisions(equations, values, queries):
    // Map variables to indices
    varMap = {}
    idx = 0
    for eq in equations:
        if eq[0] not in varMap:
            varMap[eq[0]] = idx++
        if eq[1] not in varMap:
            varMap[eq[1]] = idx++
    
    n = varMap.size()
    
    // Initialize distance matrix
    dist = n×n matrix
    for i from 0 to n-1:
        for j from 0 to n-1:
            dist[i][j] = -1  // unreachable
        dist[i][i] = 1  // self-division = 1
    
    // Fill direct relationships
    for i from 0 to equations.length-1:
        from = varMap[equations[i][0]]
        to = varMap[equations[i][1]]
        dist[from][to] = values[i]
        dist[to][from] = 1.0 / values[i]
    
    // Floyd-Warshall
    for k from 0 to n-1:
        for i from 0 to n-1:
            for j from 0 to n-1:
                if dist[i][k] > 0 && dist[k][j] > 0:
                    product = dist[i][k] * dist[k][j]
                    if dist[i][j] < 0 || product < dist[i][j]:
                        dist[i][j] = product
    
    // Answer queries
    results = []
    for query in queries:
        if query[0] in varMap && query[1] in varMap:
            from = varMap[query[0]]
            to = varMap[query[1]]
            results.append(dist[from][to])
        else:
            results.append(-1)
    
    return results
```

### Execution Visualization

### Example: equations=[a/b=2, b/c=3], query=a/c=?
```
Floyd-Warshall Process:

Step 1: Variable mapping
a→0, b→1, c→2

Step 2: Initialize matrix
dist[0][0]=1, dist[0][1]=2, dist[0][2]=?, dist[1][0]=0.5, dist[1][1]=1, dist[1][2]=3, dist[2][0]=?, dist[2][1]=0.33, dist[2][2]=1

Step 3: Floyd-Warshall with k=0 (intermediate=a)
Update using a:
dist[1][2] = min(dist[1][2], dist[1][0] * dist[0][2])
dist[1][2] = min(3, 0.5 * ?) = 3 (keep 3)

Step 4: Floyd-Warshall with k=1 (intermediate=b)
Update using b:
dist[0][2] = min(dist[0][2], dist[0][1] * dist[1][2])
dist[0][2] = min(?, 2 * 3) = min(?, 6) = 6 ✓

Step 5: Floyd-Warshall with k=2 (intermediate=c)
All pairs already optimal

Query a/c: dist[0][2] = 6 ✓

Manual verification:
a/c = (a/b) * (b/c) = 2 * 3 = 6 ✓

Visualization:
Floyd-Warshall systematically considers all intermediate vertices
Each iteration improves the distance matrix
Final matrix contains optimal all-pairs shortest paths
```

### Key Visualization Points:
- **Distance Matrix**: N×N matrix storing division results
- **Floyd-Warshall**: Three nested loops for all-pairs optimization
- **Transitivity**: Uses property a/c = (a/b) * (b/c)
- **Query Answering**: O(1) lookup after O(N³) preprocessing

### Memory Layout Visualization:
```
Variables: a, b, c
Indices: a→0, b→1, c→2

Distance Matrix Evolution:
Initial:
[  1,  2,  ?]
[0.5, 1,  3]
[  ?, 0.33, 1]

After k=0 (intermediate=a):
[  1,  2,  3]
[0.5, 1,  3]
[  ?, 0.33, 1]

After k=1 (intermediate=b):
[  1,  2,  6]
[0.5, 1,  3]
[  ?, 0.33, 1]

Final:
[  1,  2,  6]
[0.5, 1,  3]
[  ?, 0.33, 1]

Query a/c = dist[0][2] = 6 ✓

Matrix stores all possible division results
Floyd-Warshall optimizes all pairs simultaneously
```

### Time Complexity Breakdown:
- **Variable Mapping**: O(E) time, O(V) space
- **Matrix Initialization**: O(V²) time, O(V²) space
- **Floyd-Warshall**: O(V³) time, O(1) space
- **Query Answering**: O(Q) time, O(1) space
- **Total**: O(V³ + Q) time, O(V²) space
- **Optimal**: Best known for all-pairs shortest paths
- **vs Naive**: O(Q*V) per query vs O(1) after preprocessing
*/
