import java.util.*;

public class NumberOfIslands {
    
    // 200. Number of Islands
    // Time: O(M * N), Space: O(M * N) in worst case for recursion stack
    public static int numIslands(char[][] grid) {
        if (grid == null || grid.length == 0 || grid[0].length == 0) {
            return 0;
        }
        
        int m = grid.length;
        int n = grid[0].length;
        int count = 0;
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '1') {
                    count++;
                    dfs(grid, i, j);
                }
            }
        }
        
        return count;
    }
    
    private static void dfs(char[][] grid, int i, int j) {
        int m = grid.length;
        int n = grid[0].length;
        
        if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1') {
            return;
        }
        
        grid[i][j] = '0'; // Mark as visited
        
        dfs(grid, i + 1, j); // Down
        dfs(grid, i - 1, j); // Up
        dfs(grid, i, j + 1); // Right
        dfs(grid, i, j - 1); // Left
    }

    // 133. Clone Graph
    // Time: O(V + E), Space: O(V)
    public static class Node {
        public int val;
        public List<Node> neighbors;
        
        public Node() {
            val = 0;
            neighbors = new ArrayList<Node>();
        }
        
        public Node(int val) {
            this.val = val;
            neighbors = new ArrayList<Node>();
        }
        
        public Node(int val, ArrayList<Node> neighbors) {
            this.val = val;
            this.neighbors = neighbors;
        }
    }
    
    public static Node cloneGraph(Node node) {
        if (node == null) {
            return null;
        }
        
        Map<Node, Node> visited = new HashMap<>();
        Queue<Node> queue = new LinkedList<>();
        
        queue.offer(node);
        visited.put(node, new Node(node.val));
        
        while (!queue.isEmpty()) {
            Node current = queue.poll();
            
            for (Node neighbor : current.neighbors) {
                if (!visited.containsKey(neighbor)) {
                    visited.put(neighbor, new Node(neighbor.val));
                    queue.offer(neighbor);
                }
                visited.get(current).neighbors.add(visited.get(neighbor));
            }
        }
        
        return visited.get(node);
    }

    // Helper method to create graph for testing
    public static Node createGraph(int[][] adjList) {
        if (adjList == null || adjList.length == 0) {
            return null;
        }
        
        Map<Integer, Node> nodes = new HashMap<>();
        
        // Create all nodes
        for (int i = 0; i < adjList.length; i++) {
            nodes.put(i + 1, new Node(i + 1));
        }
        
        // Add neighbors
        for (int i = 0; i < adjList.length; i++) {
            for (int neighbor : adjList[i]) {
                nodes.get(i + 1).neighbors.add(nodes.get(neighbor));
            }
        }
        
        return nodes.get(1);
    }

    public static void main(String[] args) {
        // Test cases for numIslands
        char[][][] testCases1 = {
            {{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}},
            {{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}},
            {{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}},
            {{'0', '0', '0'}, {'0', '0', '0'}, {'0', '0', '0'}},
            {{'1'}, {'1'}, {'0'}, {'1'}},
            {{'1', '0'}, {'0', '1'}},
            {{'1', '1'}, {'1', '1'}},
            {{'1', '0', '1'}, {'0', '1', '0'}, {'1', '0', '1'}},
            {{'1', '1', '1', '1'}, {'1', '0', '0', '1'}, {'1', '0', '0', '1'}, {'1', '1', '1', '1'}},
            {{'1', '0', '0', '0', '1'}, {'0', '1', '0', '1', '0'}, {'0', '0', '1', '0', '0'}, {'0', '1', '0', '1', '0'}, {'1', '0', '0', '0', '1'}}
        };
        
        // Test cases for cloneGraph
        int[][][] testCases2 = {
            {{2, 4}, {1, 3}, {2, 4}, {1, 3}},
            {{}},
            {{2}, {1}},
            {{2, 3}, {1, 4}, {1, 4}, {2, 3}},
            {{2, 3, 4}, {1, 3, 4}, {1, 2, 4}, {1, 2, 3}},
            {{2}, {1, 3}, {2, 4}, {3}},
            {{2, 3}, {1, 4}, {1, 4}, {2, 3}},
            {{2, 3, 4, 5}, {1, 3, 4, 5}, {1, 2, 4, 5}, {1, 2, 3, 5}, {1, 2, 3, 4}},
            {{2}, {1}, {}},
            {{2, 3}, {1, 3}, {1, 2}}
        };
        
        System.out.println("Number of Islands:");
        for (int i = 0; i < testCases1.length; i++) {
            char[][] grid = testCases1[i];
            char[][] gridCopy = new char[grid.length][];
            for (int j = 0; j < grid.length; j++) {
                gridCopy[j] = grid[j].clone();
            }
            
            int result = numIslands(gridCopy);
            System.out.printf("Test Case %d: %s -> %d islands\n", 
                i + 1, Arrays.deepToString(grid), result);
        }
        
        System.out.println("\nClone Graph:");
        for (int i = 0; i < testCases2.length; i++) {
            int[][] adjList = testCases2[i];
            Node original = createGraph(adjList);
            Node cloned = cloneGraph(original);
            
            // Simple validation - check if cloned has same structure
            boolean isValid = true;
            if (original == null && cloned == null) {
                isValid = true;
            } else if (original == null || cloned == null) {
                isValid = false;
            } else if (original.val != cloned.val) {
                isValid = false;
            } else if (original.neighbors.size() != cloned.neighbors.size()) {
                isValid = false;
            }
            
            System.out.printf("Test Case %d: %s -> Clone %s\n", 
                i + 1, Arrays.deepToString(adjList), 
                isValid ? "successful" : "failed");
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Graph Traversal and Cloning
- **Number of Islands**: DFS/BFS flood fill algorithm
- **Clone Graph**: Deep copy with hash map for cycle detection
- **Graph Representation**: Adjacency list with Node objects
- **Visited Tracking**: Prevent infinite loops in DFS

## 2. PROBLEM CHARACTERISTICS
- **Grid Problem**: 2D array representing land/water
- **Connected Components**: Count distinct islands in grid
- **Graph Cloning**: Create independent copy with same structure
- **Cycle Handling**: Hash map prevents infinite recursion

## 3. SIMILAR PROBLEMS
- Max Area of Island
- Surrounded Regions
- Pacific Atlantic Water Flow
- Graph Valid Tree

## 4. KEY OBSERVATIONS
- **Island Detection**: '1' represents land, '0' represents water
- **DFS Traversal**: Mark visited to avoid revisiting
- **Boundary Checking**: Prevent array out of bounds
- **Clone Strategy**: Use map to track original->clone relationships

## 5. VARIATIONS & EXTENSIONS
- BFS instead of DFS for islands
- Union-Find for island counting
- Different island shapes (diagonal connections)
- Graph cloning with random pointers

## 6. INTERVIEW INSIGHTS
- Clarify: "Are diagonal connections considered islands?"
- Edge cases: empty grid, all water, All land
- Time complexity: O(M*N) for grid, O(V+E) for graph
- Space complexity: O(M*N) recursion stack, O(V+E) for clone map

## 7. COMMON MISTAKES
- Not marking visited cells in DFS
- Missing boundary checks
- Not handling empty grid case
- Infinite recursion without visited tracking
- Shallow copying vs deep cloning

## 8. OPTIMIZATION STRATEGIES
- Use iterative stack instead of recursion for DFS
- Union-Find for island counting (more efficient)
- In-place marking in grid
- Pre-allocate visited array

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like exploring a map and making copies:**
- Number of Islands: Counting distinct land masses in water
- Grid is like satellite image where forests (1s) are islands
- Clone Graph: Making exact photocopy of road network
- Need to visit each location and track connections

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding (Number of Islands)
1. **Input**: 2D grid of '1's (land) and '0's (water)
2. **Goal**: Count number of distinct islands
3. **Output**: Integer count of islands
4. **Connected**: Adjacent cells (4-directionally) form same island

#### Phase 2: Key Insight Recognition
- **"How to find islands?"** → DFS/BFS from each unvisited land cell
- **"What defines an island?"** → Connected component of '1's
- **"How to avoid double counting?"** → Mark visited cells
- **"When to stop DFS?"** → When all 4 directions are water or visited

#### Phase 3: Strategy Development (Number of Islands)
```
Human thought process:
"I'll use DFS flood fill:
1. Iterate through every cell in grid
2. When I find unvisited land cell:
   - Start DFS to mark entire island
   - Increment island count
3. DFS explores all connected land cells
4. Mark visited cells to avoid revisiting
5. Continue until all cells processed"
```

#### Phase 4: Algorithm Walkthrough (Number of Islands)
```
Example: Grid:
1 1 0 0
1 1 1 0
0 0 0 1
0 1 1 1

Human thinking:
"Start at (0,0): land, unvisited
DFS from (0,0): marks (0,0), (0,1), (1,0), (1,1)
Island count = 1

Continue scanning:
(0,2): water, skip
(0,3): water, skip
(1,0): land, but already visited
(1,1): land, but already visited
(1,2): land, unvisited → new island
DFS from (1,2): marks (1,2), (2,2)
Island count = 2

Final: 2 islands ✓"
```

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just count '1's?"** → Need to avoid double counting
2. **"DFS vs BFS?"** → Both work, DFS is more natural for recursion
3. **"What about diagonals?"** → Clarify connection rules
4. **"How to handle large grids?"** → Stack overflow, use iterative approach

### Real-World Analogy
**Number of Islands:**
- Like counting countries in a satellite image
- Forests (1s) are distinct countries
- Oceans (0s) separate them
- Use flood fill algorithm to "paint" each country

**Clone Graph:**
- Like making exact copy of a social network
- Each person is a node, friendships are edges
- Need to create new nodes with same relationships
- Prevent infinite loops with visited tracking

### Human-Readable Pseudocode
```
function numIslands(grid):
    if grid is empty: return 0
    
    count = 0
    visited = boolean[rows][cols]
    
    for i from 0 to rows-1:
        for j from 0 to cols-1:
            if grid[i][j] == '1' and not visited[i][j]:
                dfs(grid, i, j, visited)
                count++
    
    return count

function dfs(grid, i, j, visited):
    if out of bounds or grid[i][j] == '0' or visited[i][j]:
        return
    
    visited[i][j] = true
    dfs(grid, i+1, j, visited)
    dfs(grid, i-1, j, visited)
    dfs(grid, i, j+1, visited)
    dfs(grid, i, j-1, visited)

function cloneGraph(node):
    if node is null: return null
    
    visitedMap = new HashMap()
    return cloneNode(node, visitedMap)

function cloneNode(original, visitedMap):
    if original is null: return null
    
    if visitedMap.containsKey(original):
        return visitedMap.get(original)
    
    clone = new Node(original.val)
    visitedMap.put(original, clone)
    
    for neighbor in original.neighbors:
        clone.neighbors.add(cloneNode(neighbor, visitedMap))
    
    return clone
```

### Execution Visualization

### Example: Number of Islands
```
Grid Evolution:
1 1 0 0
1 1 1 0
0 0 0 1
0 1 1 1

DFS from (0,0):
Marks: (0,0) ✓
Explores: (0,1) ✓, (1,0) ✓, (1,1) ✓
Island 1 complete

DFS from (1,2):
Marks: (1,2) ✓
Explores: (2,2) ✓
Island 2 complete

Final: 2 islands ✓
```

### Example: Graph Cloning
```
Original: A↔B↔C (cycle)
Clone Process:
A' → A'' (new node, visitedMap[A]=A')
B' → B'' (new node, visitedMap[B]=B')
C' → C'' (new node, visitedMap[C]=C')

A''.neighbors = [B'']
B''.neighbors = [A'', C'']
C''.neighbors = [B'']

When processing C' again:
visitedMap has C', return existing C''

Final: A''↔B''↔C'' (exact copy) ✓
```

### Key Visualization Points:
- **Flood Fill**: DFS marks all connected land cells
- **Visited Tracking**: Prevents infinite loops and double counting
- **Hash Map Cloning**: Handles cycles and shared references
- **Graph Traversal**: Systematic exploration of connected components

### Time Complexity Breakdown:
- **Number of Islands**: O(M*N) time, O(M*N) space for recursion
- **Clone Graph**: O(V+E) time, O(V+E) space for map
- **Optimal**: Must visit each node/edge at least once
- **Tradeoffs**: Recursion simplicity vs iterative stack safety
*/
