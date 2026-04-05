import java.util.*;

public class CourseSchedule {
    
    // 207. Course Schedule
    // Time: O(V+E), Space: O(V+E) where V is number of courses, E is prerequisites
    public static boolean canFinish(int numCourses, int[][] prerequisites) {
        // Build adjacency list and in-degree count
        Map<Integer, List<Integer>> adj = new HashMap<>();
        int[] inDegree = new int[numCourses];
        
        for (int[] prereq : prerequisites) {
            int course = prereq[0];
            int prereqCourse = prereq[1];
            
            adj.computeIfAbsent(prereqCourse, k -> new ArrayList<>()).add(course);
            inDegree[course]++;
        }
        
        // Find courses with no prerequisites (in-degree = 0)
        Queue<Integer> queue = new LinkedList<>();
        for (int i = 0; i < numCourses; i++) {
            if (inDegree[i] == 0) {
                queue.offer(i);
            }
        }
        
        // Process courses
        int processed = 0;
        while (!queue.isEmpty()) {
            int course = queue.poll();
            processed++;
            
            // Update in-degree for dependent courses
            for (int neighbor : adj.getOrDefault(course, new ArrayList<>())) {
                inDegree[neighbor]--;
                if (inDegree[neighbor] == 0) {
                    queue.offer(neighbor);
                }
            }
        }
        
        return processed == numCourses;
    }

    // DFS approach to detect cycles
    public static boolean canFinishDFS(int numCourses, int[][] prerequisites) {
        // Build adjacency list
        Map<Integer, List<Integer>> adj = new HashMap<>();
        for (int[] prereq : prerequisites) {
            int course = prereq[0];
            int prereqCourse = prereq[1];
            adj.computeIfAbsent(prereqCourse, k -> new ArrayList<>()).add(course);
        }
        
        // States: 0 = unvisited, 1 = visiting, 2 = visited
        int[] state = new int[numCourses];
        
        for (int i = 0; i < numCourses; i++) {
            if (state[i] == 0 && hasCycleDFS(i, adj, state)) {
                return false; // Cycle detected
            }
        }
        
        return true;
    }
    
    private static boolean hasCycleDFS(int course, Map<Integer, List<Integer>> adj, int[] state) {
        if (state[course] == 1) {
            return true; // Cycle detected
        }
        if (state[course] == 2) {
            return false; // Already processed
        }
        
        state[course] = 1; // Mark as visiting
        
        for (int neighbor : adj.getOrDefault(course, new ArrayList<>())) {
            if (hasCycleDFS(neighbor, adj, state)) {
                return true;
            }
        }
        
        state[course] = 2; // Mark as visited
        return false;
    }

    // Alternative approach using Union-Find to detect cycles
    public static boolean canFinishUnionFind(int numCourses, int[][] prerequisites) {
        int[] parent = new int[numCourses];
        for (int i = 0; i < numCourses; i++) {
            parent[i] = i;
        }
        
        for (int[] prereq : prerequisites) {
            int course = prereq[0];
            int prereqCourse = prereq[1];
            
            if (!union(prereqCourse, course, parent)) {
                return false; // Cycle detected
            }
        }
        
        return true;
    }
    
    private static int find(int x, int[] parent) {
        if (parent[x] != x) {
            parent[x] = find(parent[x], parent);
        }
        return parent[x];
    }
    
    private static boolean union(int x, int y, int[] parent) {
        int rootX = find(x, parent);
        int rootY = find(y, parent);
        
        if (rootX == rootY) {
            return false; // Cycle detected
        }
        
        parent[rootY] = rootX;
        return true;
    }

    public static void main(String[] args) {
        // Test cases
        Object[][] testCases = {
            {2, new int[][]{{1, 0}}, "Simple linear dependency"},
            {2, new int[][]{{1, 0}, {0, 1}}, "Direct cycle"},
            {4, new int[][]{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, "Complex DAG"},
            {1, new int[][]{}, "Single course no prerequisites"},
            {3, new int[][]{{0, 1}, {0, 2}, {1, 2}}, "Multiple dependencies"},
            {3, new int[][]{{0, 1}, {1, 2}, {2, 0}}, "3-course cycle"},
            {5, new int[][]{{1, 0}, {2, 0}, {3, 1}, {4, 2}, {4, 3}}, "Complex structure"},
            {0, new int[][]{}, "No courses"},
            {4, new int[][]{{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3}}, "Complex with cycle"},
            {4, new int[][]{{1, 0}, {2, 1}, {3, 2}}, "Linear chain"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int numCourses = (int) testCases[i][0];
            int[][] prereqs = (int[][]) testCases[i][1];
            String description = (String) testCases[i][2];
            
            System.out.printf("Test Case %d: %s\n", i + 1, description);
            System.out.printf("  Courses: %d, Prerequisites: %s\n", numCourses, Arrays.deepToString(prereqs));
            
            boolean result1 = canFinish(numCourses, prereqs);
            boolean result2 = canFinishDFS(numCourses, prereqs);
            boolean result3 = canFinishUnionFind(numCourses, prereqs);
            
            System.out.printf("  BFS (Kahn's): %b\n", result1);
            System.out.printf("  DFS: %b\n", result2);
            System.out.printf("  Union-Find: %b\n\n", result3);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Topological Sort
- **Course Dependencies**: Directed acyclic graph of prerequisites
- **Cycle Detection**: Check if course schedule is possible
- **Topological Order**: Linear ordering respecting dependencies
- **Graph Processing**: BFS with in-degree counting

## 2. PROBLEM CHARACTERISTICS
- **Directed Graph**: Courses as nodes, prerequisites as edges
- **Cycle Detection**: Cycles make scheduling impossible
- **Topological Sort**: Find valid course ordering
- **Prerequisite Chain**: Course A must be taken before B if A→B

## 3. SIMILAR PROBLEMS
- Course Schedule II (find order)
- Alien Dictionary
- Minimum Height of Buildings
- Build Order from Projects

## 4. KEY OBSERVATIONS
- Course schedule possible iff graph has no cycles
- Kahn's algorithm uses in-degree counting
- DFS can detect cycles using visitation states
- Union-Find can detect cycles during edge processing
- Time complexity: O(V+E) where V=courses, E=prerequisites

## 5. VARIATIONS & EXTENSIONS
- Find actual course ordering
- Multiple semesters/parallel courses
- Weighted prerequisites
- Different constraint types

## 6. INTERVIEW INSIGHTS
- Clarify: "Are prerequisites guaranteed to be valid?"
- Edge cases: no courses, single course, self-loops
- Time complexity: O(V+E) vs O(V×E) naive
- Space complexity: O(V+E) for graph storage

## 7. COMMON MISTAKES
- Not handling self-loops properly
- Confusing course and prerequisite order
- Missing cycle detection
- Incorrect in-degree initialization

## 8. OPTIMIZATION STRATEGIES
- Use adjacency list for sparse graphs
- Early cycle detection in Union-Find
- In-place processing when possible
- Choose appropriate algorithm based on constraints

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like planning course registration:**
- You have courses with prerequisites (must take A before B)
- This creates a dependency graph (directed edges)
- You can only take a course if all prerequisites are completed
- If there's a cycle (A→B→C→A), no valid schedule exists
- Need to check if the dependency graph is acyclic

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Number of courses and prerequisite pairs
2. **Goal**: Can you finish all courses?
3. **Output**: Boolean indicating if schedule is possible

#### Phase 2: Key Insight Recognition
- **"What makes schedule impossible?"** → Cycles in dependencies
- **"How to detect cycles?"** → Topological sort or DFS
- **"What's the graph structure?"** → Courses as nodes, prerequisites as edges
- **"Why topological sort?"** → Linear ordering respects dependencies

#### Phase 3: Strategy Development
```
Human thought process:
"I'll check for cycles:
1. Build adjacency list and in-degree counts
2. Find courses with no prerequisites (in-degree = 0)
3. Process these courses, reduce in-degree of dependents
4. If we can process all courses, no cycles exist
5. If some courses remain unprocessed, there's a cycle"
```

#### Phase 4: Edge Case Handling
- **No courses**: Return true (nothing to schedule)
- **Single course**: Return true unless it has self-prerequisite
- **Self-loop**: Course requires itself → return false
- **Disconnected components**: Handle independently

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Courses: 4, Prerequisites: [[1,0], [2,0], [3,1], [3,2]]

Human thinking:
"Let's build the dependency graph:
Course 0 ← (no prerequisites)
Course 1 ← Course 0
Course 2 ← Course 0
Course 3 ← Courses 1 and 2

In-degree counts:
Course 0: 0 (no prerequisites)
Course 1: 1 (needs course 0)
Course 2: 1 (needs course 0)
Course 3: 2 (needs courses 1 and 2)

Process courses with in-degree = 0:
Queue: [0], processed = 0

Process course 0:
- Remove from queue, processed = 1
- Update dependents:
  - Course 1: in-degree becomes 0 → add to queue
  - Course 2: in-degree becomes 0 → add to queue
Queue: [1, 2]

Process course 1:
- Remove from queue, processed = 2
- Update dependent:
  - Course 3: in-degree becomes 1 (was 2)
Queue: [2]

Process course 2:
- Remove from queue, processed = 3
- Update dependent:
  - Course 3: in-degree becomes 0 → add to queue
Queue: [3]

Process course 3:
- Remove from queue, processed = 4
- No dependents to update
Queue: []

All 4 courses processed → No cycles ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Topological sort exists iff graph is acyclic
- **Why it's efficient**: Each edge processed once
- **Why it's correct**: Kahn's algorithm is proven cycle detection

### Common Human Pitfalls & How to Avoid Them
1. **"Why not try all orders?"** → Exponential O(V!)
2. **"What about DFS?"** → Works but more complex
3. **"How to handle multiple prerequisites?"** → Sum in-degree counts
4. **"What about disconnected components?"** → Process independently

### Real-World Analogy
**Like planning a multi-step project:**
- You have tasks (courses) that depend on each other
- Task B cannot start until Task A is complete
- Some tasks have multiple prerequisites
- You want to know if the project can be completed
- If there's a circular dependency (A→B→C→A), project is impossible
- Topological sort gives you a valid execution order

### Human-Readable Pseudocode
```
function canFinish(numCourses, prerequisites):
    // Build graph
    adj = adjacency list
    inDegree = array of size numCourses, initialized to 0
    
    for [course, prereq] in prerequisites:
        adj[prereq].add(course)
        inDegree[course]++
    
    // Find courses with no prerequisites
    queue = all courses with inDegree[i] == 0
    processed = 0
    
    while queue is not empty:
        course = queue.dequeue()
        processed++
        
        for dependent in adj[course]:
            inDegree[dependent]--
            if inDegree[dependent] == 0:
                queue.enqueue(dependent)
    
    return processed == numCourses
```

### Execution Visualization

### Example: 4 courses, [[1,0], [2,0], [3,1], [3,2]]
```
Dependency Graph:
0 ← (no prerequisites)
1 ← 0
2 ← 0
3 ← 1, 2

In-degree Evolution:
Initial: [0, 1, 1, 2]
After processing 0: [0, 0, 0, 1]
After processing 1: [0, 0, 0, 1]
After processing 2: [0, 0, 0, 0]
After processing 3: [0, 0, 0, 0]

Queue Processing:
Start: [0] → Process 0 → Queue: [1,2]
Process 1 → Queue: [2] → Process 2 → Queue: [3]
Process 3 → Queue: [] → Done

Result: All 4 courses processed ✓

Visualization:
0 → 1,2 (depend on 0)
   ↓
   3 (depends on 1,2)
```

### Key Visualization Points:
- **Graph building** from prerequisites
- **In-degree counting** for each course
- **Queue processing** of courses with no unmet prerequisites
- **Cycle detection** when some courses remain unprocessed
- **Topological order** emerges from processing sequence

### Memory Layout Visualization:
```
Graph Representation:
Adjacency List:
0: [1, 2]
1: [3]
2: [3]
3: []

In-degree Array:
Index: 0  1  2  3
Value: 0  1  1  2

Processing Evolution:
Step 1: Queue=[0], Processed=0
Step 2: Queue=[1,2], Processed=1
Step 3: Queue=[2], Processed=2
Step 4: Queue=[3], Processed=3
Step 5: Queue=[], Processed=4

Final: Processed=4, Total=4 → Success ✓
```

### Time Complexity Breakdown:
- **Graph Building**: O(E) where E is number of prerequisites
- **In-degree Calculation**: O(E) during graph building
- **Queue Processing**: O(V+E) where V is number of courses
- **Total**: O(V+E) time, O(V+E) space
- **Optimal**: Linear time for this problem
- **vs DFS**: O(V+E) time but recursion overhead
*/
