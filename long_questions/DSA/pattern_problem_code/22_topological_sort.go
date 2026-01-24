package main

import "fmt"

// Pattern: Topological Sort
// Difficulty: Medium/Hard
// Key Concept: Ordering tasks with dependencies. If A -> B, A must come before B.

/*
INTUITION:
"Course Schedule II"
To take Course B, you must take Course A.
This is a Graph problem. Nodes = Courses. Edges = Dependencies.
A -> B.

How do we solve it? (Kahn's Algorithm)
1. "Indegree": Count how many requirements each course has.
   - A: 0 reqs.
   - B: 1 req (A).
2. Start with courses that have 0 reqs (Indegree 0). (We can take them now!).
   - Take A.
3. When we take A, we "fulfill" the requirement for B.
   - B's reqs go from 1 -> 0.
   - Now B has 0 reqs. Add B to the queue.

If we finish, and we took all courses, success!
If we can't take all (Cycle exists, e.g., A->B->A), fail.

PROBLEM:
Given numCourses and prerequisites [A, B] (B -> A), return valid ordering.

ALGORITHM:
1. Build Graph (Adjacency List) + Indegree Array.
2. Push all nodes with Indegree 0 to Queue.
3. While Queue not empty:
   - Pop node. Add to Order.
   - For each neighbor: decremenet Indegree.
   - If Indegree becomes 0, push neighbor to Queue.
4. Check if Order length == numCourses.
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	indegree := make([]int, numCourses)

	// Step 1: Build Graph
	for _, req := range prerequisites {
		course, pre := req[0], req[1]
		// Edge: pre -> course
		graph[pre] = append(graph[pre], course)
		indegree[course]++
	}

	// Step 2: Init Queue with 0-indegree nodes
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := []int{}

	// Step 3: Process Queue
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		result = append(result, current)

		if neighbors, exists := graph[current]; exists {
			for _, neighbor := range neighbors {
				indegree[neighbor]--
				if indegree[neighbor] == 0 {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	// Step 4: Check for cycle
	if len(result) != numCourses {
		return []int{} // Cycle detected, impossible
	}

	return result
}

func main() {
	// 4 Courses.
	// 1 -> 0
	// 2 -> 0
	// 3 -> 1
	// 3 -> 2
	// Dependencies: 0 must be done before 1 and 2. 1 & 2 before 3.
	numCourses := 4
	prereqs := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}

	fmt.Printf("Courses: %d, Prereqs: %v\n", numCourses, prereqs)
	res := findOrder(numCourses, prereqs)

	// Expected: [0, 1, 2, 3] OR [0, 2, 1, 3]
	fmt.Printf("Order: %v\n", res)
}
