package main

import "fmt"

// 210. Course Schedule II
// Time: O(V+E), Space: O(V+E) where V is number of courses, E is prerequisites
func findOrder(numCourses int, prerequisites [][]int) []int {
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
	var result []int
	processed := 0
	
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		
		result = append(result, course)
		processed++
		
		// Update in-degree for dependent courses
		for _, neighbor := range adj[course] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	
	// Check if all courses were processed (no cycle)
	if processed == numCourses {
		return result
	}
	
	return []int{} // Cycle detected
}

// DFS approach
func findOrderDFS(numCourses int, prerequisites [][]int) []int {
	// Build adjacency list
	adj := make(map[int][]int)
	for _, prereq := range prerequisites {
		course := prereq[0]
		prereqCourse := prereq[1]
		adj[prereqCourse] = append(adj[prereqCourse], course)
	}
	
	// States: 0 = unvisited, 1 = visiting, 2 = visited
	state := make([]int, numCourses)
	var result []int
	hasCycle := false
	
	var dfs func(int)
	dfs = func(course int) {
		if hasCycle {
			return
		}
		
		if state[course] == 1 {
			hasCycle = true
			return
		}
		
		if state[course] == 2 {
			return
		}
		
		state[course] = 1 // Mark as visiting
		
		for _, neighbor := range adj[course] {
			dfs(neighbor)
		}
		
		state[course] = 2 // Mark as visited
		result = append(result, course)
	}
	
	// Process all courses
	for i := 0; i < numCourses; i++ {
		if state[i] == 0 {
			dfs(i)
		}
	}
	
	if hasCycle {
		return []int{}
	}
	
	// Reverse the result (post-order)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	
	return result
}

// Alternative approach using topological sort with stack
func findOrderStack(numCourses int, prerequisites [][]int) []int {
	// Build adjacency list and in-degree count
	adj := make(map[int][]int)
	inDegree := make([]int, numCourses)
	
	for _, prereq := range prerequisites {
		course := prereq[0]
		prereqCourse := prereq[1]
		
		adj[prereqCourse] = append(adj[prereqCourse], course)
		inDegree[course]++
	}
	
	// Use stack instead of queue
	stack := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			stack = append(stack, i)
		}
	}
	
	var result []int
	processed := 0
	
	for len(stack) > 0 {
		course := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		result = append(result, course)
		processed++
		
		// Update in-degree for dependent courses
		for _, neighbor := range adj[course] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				stack = append(stack, neighbor)
			}
		}
	}
	
	if processed == numCourses {
		return result
	}
	
	return []int{}
}

// Helper function to validate the course order
func validateOrder(order []int, prerequisites [][]int) bool {
	// Create position map
	pos := make(map[int]int)
	for i, course := range order {
		pos[course] = i
	}
	
	// Check all prerequisites
	for _, prereq := range prerequisites {
		course := prereq[0]
		prereqCourse := prereq[1]
		
		if pos[course] <= pos[prereqCourse] {
			return false // Prerequisite comes after course
		}
	}
	
	return true
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Topological Sort for Course Ordering
- **Graph Construction**: Courses as nodes, prerequisites as directed edges
- **Topological Ordering**: Find valid course sequence
- **Multiple Solutions**: BFS (Kahn's), DFS, Stack-based approaches
- **Cycle Detection**: Return empty array if no valid ordering exists

## 2. PROBLEM CHARACTERISTICS
- **Course Dependencies**: Must complete prerequisites before courses
- **Valid Ordering**: Find sequence satisfying all prerequisites
- **Multiple Valid Orders**: May have multiple possible orderings
- **Cycle Handling**: Return empty array if cycles exist

## 3. SIMILAR PROBLEMS
- Course Schedule (LeetCode 207) - Cycle detection only
- Alien Dictionary (LeetCode 269) - Character ordering from words
- Sequence Reconstruction (LeetCode 444) - Reconstruct from subsequence
- Minimum Height Trees (LeetCode 310) - Find tree roots

## 4. KEY OBSERVATIONS
- **Graph Theory**: Problem reduces to topological sort of course DAG
- **Multiple Solutions**: Different valid orderings may exist
- **Cycle Detection**: No valid ordering if graph has cycles
- **Deterministic Output**: Sort queue for consistent results

## 5. VARIATIONS & EXTENSIONS
- **Multiple Valid Orders**: Return all possible orderings
- **Parallel Courses**: Find minimum semesters needed
- **Prerequisite Updates**: Dynamic course additions
- **Weighted Courses**: Course priorities or credits

## 6. INTERVIEW INSIGHTS
- Always clarify: "Multiple valid orders? Course numbering? Cycle handling?"
- Edge cases: empty courses, no prerequisites, cycles, disconnected components
- Time complexity: O(V+E) where V=courses, E=prerequisites
- Space complexity: O(V+E) for graph storage
- Key insight: topological sort finds valid ordering in DAG

## 7. COMMON MISTAKES
- Not handling cycles correctly (returning partial order)
- Wrong graph direction (prereq → course vs course → prereq)
- Not processing all connected components
- Incorrect cycle detection in DFS
- Not handling empty result case properly

## 8. OPTIMIZATION STRATEGIES
- **BFS (Kahn's)**: O(V+E) time, O(V+E) space
- **DFS Topological Sort**: O(V+E) time, O(V) space
- **Stack-based**: O(V+E) time, O(V) space
- **Early Termination**: Stop when cycle detected

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like planning course sequence with prerequisites:**
- You have courses with prerequisite requirements
- Need to find order that satisfies all dependencies
- Like building a study plan where you can't take advanced courses first
- Multiple valid orders may exist (different study sequences)
- If there's a cycle (A→B→A), impossible to complete

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Number of courses, prerequisite pairs [course, prereq]
2. **Goal**: Find valid course ordering sequence
3. **Constraint**: Must complete prerequisites before courses
4. **Output**: Array of courses in valid order (empty if impossible)

#### Phase 2: Key Insight Recognition
- **"Graph natural fit"** → Courses as nodes, prereqs as edges
- **"Topological sort"** → Standard algorithm for this problem
- **"Multiple solutions"** → Different valid orderings may exist
- **"Cycle detection"** → No valid ordering if graph has cycles

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find course order that satisfies prerequisites.
This is classic topological sort:

BFS Approach (Kahn's Algorithm):
1. Build graph and count in-degrees
2. Start with courses having 0 in-degree (no prerequisites)
3. Remove them and update in-degrees of dependent courses
4. Continue until all courses processed or stuck (cycle)
5. If all processed → return order, else return empty

DFS Approach:
1. Build adjacency list
2. DFS with cycle detection
3. Add courses to result in post-order
4. Reverse result for correct ordering
5. Return empty if cycle detected"
```

#### Phase 4: Edge Case Handling
- **No courses**: Return empty array
- **No prerequisites**: Return any order of all courses
- **Single course**: Return [course]
- **Cycle**: Return empty array
- **Disconnected components**: Handle all components

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Courses: 4, Prerequisites: [[1,0], [2,0], [3,1], [3,2]]

Human thinking:
"Build graph:
Course 0: enables courses 1, 2
Course 1: enables course 3
Course 2: enables course 3
Course 3: enables nothing

In-degree counts:
Course 0: 0 (no prerequisites)
Course 1: 1 (requires course 0)
Course 2: 1 (requires course 0)
Course 3: 2 (requires courses 1, 2)

BFS Processing:
Initial queue: [0] (courses with in-degree 0)

Step 1: Process course 0
- Remove 0 from queue
- result = [0]
- Update course 1 in-degree: 1→0, add to queue
- Update course 2 in-degree: 1→0, add to queue
- queue = [1, 2]

Step 2: Process course 1 (sorted order: 1 before 2)
- Remove 1 from queue
- result = [0, 1]
- Update course 3 in-degree: 2→1
- queue = [2]

Step 3: Process course 2
- Remove 2 from queue
- result = [0, 1, 2]
- Update course 3 in-degree: 1→0, add to queue
- queue = [3]

Step 4: Process course 3
- Remove 3 from queue
- result = [0, 1, 2, 3]
- No dependents to update
- queue = []

Final result: [0, 1, 2, 3] ✓ All courses processed!"
```

#### Phase 6: Intuition Validation
- **Why graph works**: Natural representation of course dependencies
- **Why topological sort works**: Finds ordering satisfying all constraints
- **Why multiple orders possible**: Different valid sequences may exist
- **Why cycle detection works**: Cycles make completion impossible

### Common Human Pitfalls & How to Avoid Them
1. **"Why not try all permutations?"** → Exponential vs O(V+E)
2. **"Should I use BFS or DFS?"** → Both work, choose based on preference
3. **"What about multiple valid orders?"** → Return any valid one
4. **"Can I optimize further?"** → O(V+E) is already optimal
5. **"What about course priorities?"** → Different problem (weighted topological sort)

### Real-World Analogy
**Like planning a university curriculum with prerequisites:**
- You have courses that require other courses first
- Need to find study sequence that satisfies all requirements
- Multiple valid sequences may exist (different study plans)
- If there's a circular dependency, impossible to graduate
- Build dependency graph and find valid ordering

### Human-Readable Pseudocode
```
function findCourseOrder(numCourses, prerequisites):
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
    result = []
    while queue is not empty:
        course = queue.pop_front()
        result.append(course)
        
        for each neighbor in adj[course]:
            inDegree[neighbor]--
            if inDegree[neighbor] == 0:
                queue.append(neighbor)
    
    // Check if all courses were processed
    if len(result) == numCourses:
        return result
    else:
        return [] // Cycle detected
```

### Execution Visualization

### Example: numCourses = 4, prerequisites = [[1,0], [2,0], [3,1], [3,2]]
```
Graph Construction:
Course 0: [1, 2]
Course 1: [3]
Course 2: [3]
Course 3: []

In-degree Array:
[0, 1, 1, 2] (courses 0,1,2,3)

BFS Processing:
Initial queue: [0] (courses with in-degree 0)

Step 1: Process course 0
- Remove 0 from queue
- result = [0]
- Update course 1 in-degree: 1→0, queue=[1]
- Update course 2 in-degree: 1→0, queue=[1,2]

Step 2: Process course 1 (queue sorted: 1 before 2)
- Remove 1 from queue
- result = [0, 1]
- Update course 3 in-degree: 2→1, queue=[2]

Step 3: Process course 2
- Remove 2 from queue
- result = [0, 1, 2]
- Update course 3 in-degree: 1→0, queue=[3]

Step 4: Process course 3
- Remove 3 from queue
- result = [0, 1, 2, 3]
- No neighbors to update
- queue = []

Final result: [0, 1, 2, 3] ✓ All courses processed!
```

### Key Visualization Points:
- **Graph Building**: Course → dependent courses
- **In-degree Tracking**: Number of prerequisites per course
- **Queue Processing**: Courses with no prerequisites first
- **Cycle Detection**: Stuck queue indicates cycle

### Memory Layout Visualization:
```
Graph State During Processing:
Courses: 4, Prerequisites: [[1,0], [2,0], [3,1], [3,2]]

Adjacency List:
0: [1, 2]  (course 0 enables courses 1, 2)
1: [3]     (course 1 enables course 3)
2: [3]     (course 2 enables course 3)
3: []      (course 3 enables nothing)

In-degree Array:
[0, 1, 1, 2] (courses 0,1,2,3)

Queue Evolution:
Initial: [0] (courses with in-degree 0)
After 0: [1, 2] (courses 1,2 now have in-degree 0)
After 1: [2] (course 3 in-degree becomes 1)
After 2: [3] (course 3 in-degree becomes 0)
After 3: [] (all courses processed)

Result: [0, 1, 2, 3] (all 4 courses) ✓ No cycles!
```

### Time Complexity Breakdown:
- **Graph Building**: O(V+E) time, O(V+E) space
- **BFS Processing**: O(V+E) time, O(V) space for queue
- **DFS Processing**: O(V+E) time, O(V) space for recursion
- **Total**: O(V+E) time, O(V+E) space

### Alternative Approaches:

#### 1. DFS Topological Sort (O(V+E) time, O(V) space)
```go
func findOrderDFS(numCourses int, prerequisites [][]int) []int {
    adj := make(map[int][]int)
    for _, prereq := range prerequisites {
        course, prereqCourse := prereq[0], prereq[1]
        adj[prereqCourse] = append(adj[prereqCourse], course)
    }
    
    state := make([]int, numCourses) // 0=unvisited, 1=visiting, 2=visited
    var result []int
    
    var dfs func(int)
    dfs = func(course int) {
        if state[course] == 1 {
            return // Cycle detected
        }
        if state[course] == 2 {
            return // Already processed
        }
        
        state[course] = 1 // Visiting
        for _, neighbor := range adj[course] {
            dfs(neighbor)
        }
        state[course] = 2 // Visited
        result = append(result, course)
    }
    
    for i := 0; i < numCourses; i++ {
        if state[i] == 0 {
            dfs(i)
        }
    }
    
    // Check for cycles
    for i := 0; i < numCourses; i++ {
        if state[i] == 1 {
            return [] // Cycle detected
        }
    }
    
    // Reverse result for correct order
    reverse(result)
    return result
}
```
- **Pros**: Natural recursive formulation
- **Cons**: Recursion depth, more complex cycle handling

#### 2. Stack-based Topological Sort (O(V+E) time, O(V) space)
```go
func findOrderStack(numCourses int, prerequisites [][]int) []int {
    // Build same graph as BFS approach
    adj := make(map[int][]int)
    inDegree := make([]int, numCourses)
    
    for _, prereq := range prerequisites {
        course, prereqCourse := prereq[0], prereq[1]
        adj[prereqCourse] = append(adj[prereqCourse], course)
        inDegree[course]++
    }
    
    // Use stack instead of queue
    stack := []int{}
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            stack = append(stack, i)
        }
    }
    
    var result []int
    for len(stack) > 0 {
        course := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, course)
        
        for _, neighbor := range adj[course] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                stack = append(stack, neighbor)
            }
        }
    }
    
    if len(result) == numCourses {
        return result
    }
    return [] // Cycle detected
}
```
- **Pros**: No queue needed, stack operations
- **Cons**: Less intuitive than BFS for many

#### 3. Multiple Valid Orders (O(V!*E) time, O(V+E) space)
```go
func findAllOrders(numCourses int, prerequisites [][]int) [][]int {
    // Find all possible topological orders
    // This is complex and typically not required
    // Use backtracking to explore all valid orders
    // ... implementation details omitted
}
```
- **Pros**: Finds all possible valid orders
- **Cons**: Exponential complexity, usually overkill

### Extensions for Interviews:
- **Multiple Valid Orders**: Discuss when multiple orderings are possible
- **Parallel Courses**: Find minimum semesters needed
- **Prerequisite Updates**: Dynamic course additions
- **Weighted Courses**: Course priorities or credits
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
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, "Complex DAG"},
		{1, [][]int{}, "Single course no prerequisites"},
		{3, [][]int{{0, 1}, {0, 2}, {1, 2}}, "Multiple dependencies"},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, "Complex structure"},
		{2, [][]int{{0, 1}, {1, 0}}, "Direct cycle"},
		{3, [][]int{{0, 1}, {1, 2}, {2, 0}}, "3-course cycle"},
		{0, [][]int{}, "No courses"},
		{4, [][]int{{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3}}, "Complex with cycle"},
		{4, [][]int{{1, 0}, {2, 1}, {3, 2}}, "Linear chain"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Courses: %d, Prerequisites: %v\n", tc.numCourses, tc.prereqs)
		
		result1 := findOrder(tc.numCourses, tc.prereqs)
		result2 := findOrderDFS(tc.numCourses, tc.prereqs)
		result3 := findOrderStack(tc.numCourses, tc.prereqs)
		
		fmt.Printf("  BFS (Kahn's): %v\n", result1)
		fmt.Printf("  DFS: %v\n", result2)
		fmt.Printf("  Stack: %v\n", result3)
		
		// Validate results
		if len(result1) > 0 {
			valid1 := validateOrder(result1, tc.prereqs)
			fmt.Printf("  BFS Valid: %t\n", valid1)
		}
		if len(result2) > 0 {
			valid2 := validateOrder(result2, tc.prereqs)
			fmt.Printf("  DFS Valid: %t\n", valid2)
		}
		if len(result3) > 0 {
			valid3 := validateOrder(result3, tc.prereqs)
			fmt.Printf("  Stack Valid: %t\n", valid3)
		}
		
		fmt.Println()
	}
}
