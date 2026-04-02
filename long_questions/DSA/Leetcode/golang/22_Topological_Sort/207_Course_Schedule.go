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
