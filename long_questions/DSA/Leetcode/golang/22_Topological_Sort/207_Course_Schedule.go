package main

import "fmt"

// 207. Course Schedule
// Time: O(V+E), Space: O(V+E) where V is number of courses, E is prerequisites
func canFinish(numCourses int, prerequisites [][]int) bool {
	// Build adjacency list and in-degree count
	adj := make(map[int][]int)
	inDegree := make([]int, numCourses)
	
	for _, prereq := range prerequisites {
		course := prereq[0]
		prereqCourse := prereq[1]
		
		adj[prereqCourse] = append(adj[prereqCourse], course)
		inDegree[course]++
	}
	
	// Find courses with no prerequisites (in-degree = 0)
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	
	// Process courses
	processed := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		processed++
		
		// Update in-degree for dependent courses
		for _, neighbor := range adj[course] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	
	return processed == numCourses
}

// DFS approach to detect cycles
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	// Build adjacency list
	adj := make(map[int][]int)
	for _, prereq := range prerequisites {
		course := prereq[0]
		prereqCourse := prereq[1]
		adj[prereqCourse] = append(adj[prereqCourse], course)
	}
	
	// States: 0 = unvisited, 1 = visiting, 2 = visited
	state := make([]int, numCourses)
	
	var dfs func(int) bool
	dfs = func(course int) bool {
		if state[course] == 1 {
			return true // Cycle detected
		}
		if state[course] == 2 {
			return false // Already processed
		}
		
		state[course] = 1 // Mark as visiting
		
		for _, neighbor := range adj[course] {
			if dfs(neighbor) {
				return true
			}
		}
		
		state[course] = 2 // Mark as visited
		return false
	}
	
	// Check all courses for cycles
	for i := 0; i < numCourses; i++ {
		if state[i] == 0 {
			if dfs(i) {
				return false // Cycle detected
			}
		}
	}
	
	return true
}

// Alternative approach using Union-Find to detect cycles
func canFinishUnionFind(numCourses int, prerequisites [][]int) bool {
	parent := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		parent[i] = i
	}
	
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	
	var union func(int, int) bool
	union = func(x, y int) bool {
		rootX := find(x)
		rootY := find(y)
		
		if rootX == rootY {
			return false // Cycle detected
		}
		
		parent[rootY] = rootX
		return true
	}
	
	for _, prereq := range prerequisites {
		course := prereq[0]
		prereqCourse := prereq[1]
		
		if !union(prereqCourse, course) {
			return false // Cycle detected
		}
	}
	
	return true
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Topological Sort for Cycle Detection
- **Graph Representation**: Courses as nodes, prerequisites as directed edges
- **Cycle Detection**: Check if graph has cycles (impossible schedule)
- **Topological Order**: Valid order exists if no cycles
- **Multiple Approaches**: BFS (Kahn's), DFS, Union-Find

## 2. PROBLEM CHARACTERISTICS
- **Directed Graph**: Courses with prerequisite relationships
- **Cycle Detection**: Determine if valid ordering exists
- **Course Dependencies**: Must complete prerequisites before courses
- **Feasibility Check**: Return true if all courses can be completed

## 3. SIMILAR PROBLEMS
- Course Schedule II (LeetCode 210) - Find valid course order
- Minimum Height Trees (LeetCode 310) - Find tree roots
- Alien Dictionary (LeetCode 269) - Determine dictionary order
- Sequence Reconstruction (LeetCode 444) - Reconstruct sequence

## 4. KEY OBSERVATIONS
- **Graph Theory**: Problem reduces to cycle detection in directed graph
- **Topological Sort**: Valid order exists iff no cycles
- **Multiple Solutions**: BFS, DFS, Union-Find all work
- **Edge Cases**: Empty courses, no prerequisites, single course

## 5. VARIATIONS & EXTENSIONS
- **Course Ordering**: Find actual valid order (LeetCode 210)
- **Multiple Prerequisites**: Courses can have multiple requirements
- **Parallel Courses**: Maximum courses per semester
- **Prerequisite Updates**: Dynamic prerequisite changes

## 6. INTERVIEW INSIGHTS
- Always clarify: "Course numbering? Duplicate prerequisites? Self-loops?"
- Edge cases: empty courses, no prerequisites, cycles
- Time complexity: O(V+E) where V=courses, E=prerequisites
- Space complexity: O(V+E) for graph storage
- Key insight: cycle detection in directed graph

## 7. COMMON MISTAKES
- Not handling self-cycles (course depends on itself)
- Wrong graph direction (prereq → course vs course → prereq)
- Not processing all connected components
- Incorrect cycle detection in DFS
- Union-Find implementation errors

## 8. OPTIMIZATION STRATEGIES
- **BFS (Kahn's)**: O(V+E) time, O(V+E) space
- **DFS Cycle Detection**: O(V+E) time, O(V) space
- **Union-Find**: O(V+E) time, O(V) space
- **Early Termination**: Stop when cycle detected

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like course planning with dependencies:**
- You have courses with prerequisite requirements
- Need to determine if all courses can be completed
- If there's a cycle (A→B→A), impossible to complete
- Like a dependency graph where cycles break the system
- Topological sort finds valid ordering or detects cycles

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Number of courses, prerequisite pairs [course, prereq]
2. **Goal**: Determine if all courses can be completed
3. **Constraint**: Must complete prerequisites before courses
4. **Output**: Boolean indicating feasibility

#### Phase 2: Key Insight Recognition
- **"Graph natural fit"** → Courses as nodes, prereqs as edges
- **"Cycle detection"** → Valid schedule exists iff no cycles
- **"Topological sort"** → Standard algorithm for this problem
- **"Multiple approaches"** → BFS, DFS, Union-Find all work

#### Phase 3: Strategy Development
```
Human thought process:
"I need to check if course dependencies have cycles.
This is a classic topological sort problem:

BFS Approach (Kahn's Algorithm):
1. Build graph and count in-degrees
2. Start with courses having 0 in-degree (no prerequisites)
3. Remove them and update in-degrees of dependent courses
4. If we can remove all courses → no cycles
5. If stuck → cycle exists

DFS Approach:
1. Build adjacency list
2. DFS with cycle detection (visiting state)
3. If we encounter a node being visited → cycle
4. If DFS completes without cycles → valid schedule

Union-Find Approach:
1. For each prerequisite [course, prereq]
2. If course and prereq are already connected → cycle
3. Otherwise, union them
4. If no cycles detected → valid schedule"
```

#### Phase 4: Edge Case Handling
- **No courses**: Return true (nothing to complete)
- **No prerequisites**: Return true (all courses independent)
- **Single course**: Return true unless self-dependent
- **Self-cycle**: Course depends on itself → return false

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Courses: 2, Prerequisites: [[1,0]]

Human thinking:
"Build graph:
Course 0: no prerequisites
Course 1: requires course 0

BFS Approach:
1. Count in-degrees: course 0 has 0, course 1 has 1
2. Queue starts with [0] (no prerequisites)
3. Process course 0:
   - Remove from queue
   - Update course 1 in-degree: 1 → 0
   - Add course 1 to queue
4. Process course 1:
   - Remove from queue
   - No dependents to update
5. All courses processed → no cycles ✓

Result: true (can complete both courses)"
```

#### Phase 6: Intuition Validation
- **Why graph works**: Natural representation of dependencies
- **Why cycle detection works**: Cycles make completion impossible
- **Why topological sort works**: Standard algorithm for DAG validation
- **Why multiple approaches work**: Different ways to detect cycles

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just try all orders?"** → Exponential permutations
2. **"Should I use BFS or DFS?"** → Both work, choose based on preference
3. **"What about Union-Find?"** → Works but less intuitive for this problem
4. **"Can I optimize further?"** → O(V+E) is already optimal
5. **"What about course ordering?"** → Different problem (LeetCode 210)

### Real-World Analogy
**Like planning a curriculum with course prerequisites:**
- You have courses that require other courses first
- Need to determine if the curriculum is feasible
- If there's a circular dependency (A needs B, B needs A), impossible
- Build dependency graph and check for cycles
- Valid curriculum has no prerequisite cycles

### Human-Readable Pseudocode
```
function canFinish(numCourses, prerequisites):
    // Build graph and in-degree count
    adj = adjacency list
    inDegree = array of size numCourses
    
    for each [course, prereq] in prerequisites:
        adj[prereq].append(course)
        inDegree[course]++
    
    // Initialize queue with courses having no prerequisites
    queue = []
    for i from 0 to numCourses-1:
        if inDegree[i] == 0:
            queue.append(i)
    
    // Process courses
    processed = 0
    while queue is not empty:
        course = queue.pop_front()
        processed++
        
        for each neighbor in adj[course]:
            inDegree[neighbor]--
            if inDegree[neighbor] == 0:
                queue.append(neighbor)
    
    return processed == numCourses
```

### Execution Visualization

### Example: numCourses = 2, prerequisites = [[1,0]]
```
Graph Construction:
Course 0: []
Course 1: [0]

In-degree calculation:
inDegree[0] = 0 (no prerequisites)
inDegree[1] = 1 (requires course 0)

BFS Processing:
Initial queue: [0] (courses with in-degree 0)

Step 1: Process course 0
- Remove 0 from queue
- processed = 1
- Update neighbors: course 1 in-degree: 1 → 0
- Add course 1 to queue: [1]

Step 2: Process course 1
- Remove 1 from queue
- processed = 2
- No neighbors to update

Final: processed = 2, numCourses = 2 ✓
All courses can be completed!
```

### Key Visualization Points:
- **Graph Building**: Course → dependent courses
- **In-degree Tracking**: Number of prerequisites per course
- **Queue Processing**: Courses with no prerequisites first
- **Cycle Detection**: Stuck queue indicates cycle

### Memory Layout Visualization:
```
Graph State During Processing:
Courses: 2, Prerequisites: [[1,0]]

Adjacency List:
0: [1]  (course 0 enables course 1)
1: []   (course 1 enables nothing)

In-degree Array:
[0, 1]  (course 0: 0 prereqs, course 1: 1 prereq)

Queue Evolution:
Initial: [0]  (courses with 0 in-degree)
After processing 0: [1]  (course 1 now has 0 in-degree)
After processing 1: []  (all courses processed)

Result: processed=2, total=2 ✓ No cycles!
```

### Time Complexity Breakdown:
- **Graph Building**: O(V+E) time, O(V+E) space
- **BFS Processing**: O(V+E) time, O(V) space for queue
- **DFS Processing**: O(V+E) time, O(V) space for recursion
- **Union-Find**: O(V+E) time, O(V) space
- **Total**: O(V+E) time, O(V+E) space

### Alternative Approaches:

#### 1. DFS Cycle Detection (O(V+E) time, O(V) space)
```go
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
    adj := make(map[int][]int)
    for _, prereq := range prerequisites {
        course, prereqCourse := prereq[0], prereq[1]
        adj[prereqCourse] = append(adj[prereqCourse], course)
    }
    
    state := make([]int, numCourses) // 0=unvisited, 1=visiting, 2=visited
    
    var dfs func(int) bool
    dfs = func(course int) bool {
        if state[course] == 1 {
            return true // Cycle detected
        }
        if state[course] == 2 {
            return false // Already processed
        }
        
        state[course] = 1 // Visiting
        for _, neighbor := range adj[course] {
            if dfs(neighbor) {
                return true
            }
        }
        state[course] = 2 // Visited
        return false
    }
    
    for i := 0; i < numCourses; i++ {
        if state[i] == 0 && dfs(i) {
            return false
        }
    }
    return true
}
```
- **Pros**: Intuitive, works for cycle detection
- **Cons**: Recursion depth, more complex than BFS

#### 2. Union-Find Cycle Detection (O(V+E) time, O(V) space)
```go
func canFinishUnionFind(numCourses int, prerequisites [][]int) bool {
    parent := make([]int, numCourses)
    for i := 0; i < numCourses; i++ {
        parent[i] = i
    }
    
    find := func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    
    union := func(x, y int) bool {
        rootX, rootY := find(x), find(y)
        if rootX == rootY {
            return false // Cycle detected
        }
        parent[rootY] = rootX
        return true
    }
    
    for _, prereq := range prerequisites {
        course, prereqCourse := prereq[0], prereq[1]
        if !union(prereqCourse, course) {
            return false // Cycle detected
        }
    }
    return true
}
```
- **Pros**: Simple implementation, no recursion
- **Cons**: Less intuitive, path compression needed for efficiency

#### 3. Matrix Representation (O(V²) time, O(V²) space)
```go
func canFinishMatrix(numCourses int, prerequisites [][]int) bool {
    // Use adjacency matrix for small number of courses
    graph := make([][]bool, numCourses)
    for i := 0; i < numCourses; i++ {
        graph[i] = make([]bool, numCourses)
    }
    
    for _, prereq := range prerequisites {
        course, prereqCourse := prereq[0], prereq[1]
        graph[prereqCourse][course] = true
    }
    
    // Apply topological sort using matrix
    // ... (implementation details omitted)
}
```
- **Pros**: Simple for small number of courses
- **Cons**: O(V²) space, inefficient for large V

### Extensions for Interviews:
- **Course Ordering**: Find actual valid order (LeetCode 210)
- **Multiple Prerequisites**: Courses can have multiple requirements
- **Parallel Courses**: Maximum courses per semester
- **Prerequisite Updates**: Dynamic prerequisite changes
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []struct {
		numCourses  int
		prereqs     [][]int
		description string
	}{
		{2, [][]int{{1, 0}}, "Simple linear dependency"},
		{2, [][]int{{1, 0}, {0, 1}}, "Direct cycle"},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, "Complex DAG"},
		{1, [][]int{}, "Single course no prerequisites"},
		{3, [][]int{{0, 1}, {0, 2}, {1, 2}}, "Multiple dependencies"},
		{3, [][]int{{0, 1}, {1, 2}, {2, 0}}, "3-course cycle"},
		{5, [][]int{{1, 0}, {2, 0}, {3, 1}, {4, 2}, {4, 3}}, "Complex structure"},
		{0, [][]int{}, "No courses"},
		{4, [][]int{{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3}}, "Complex with cycle"},
		{4, [][]int{{1, 0}, {2, 1}, {3, 2}}, "Linear chain"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Courses: %d, Prerequisites: %v\n", tc.numCourses, tc.prereqs)
		
		result1 := canFinish(tc.numCourses, tc.prereqs)
		result2 := canFinishDFS(tc.numCourses, tc.prereqs)
		result3 := canFinishUnionFind(tc.numCourses, tc.prereqs)
		
		fmt.Printf("  BFS (Kahn's): %t\n", result1)
		fmt.Printf("  DFS: %t\n", result2)
		fmt.Printf("  Union-Find: %t\n\n", result3)
	}
}
