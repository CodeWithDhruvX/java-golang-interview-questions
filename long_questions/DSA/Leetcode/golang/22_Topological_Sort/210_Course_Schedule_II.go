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
