package main

import "fmt"

// Pattern: Union-Find (Disjoint Set Union - DSU)
// Difficulty: Medium/Hard
// Key Concept: Managing disjoint sets efficiently using "Parent" arrays and "Path Compression".

/*
INTUITION:
Imagine a party. Everyone is initially alone (their own boss).
- Person 1 shakes hands with Person 2. Now {1, 2} are a group. 1 is the "Representative" (Boss).
- Person 3 shakes hands with Person 4. {3, 4} are a group. 3 is Boss.
- Person 2 shakes hands with Person 4.
  - 2's Boss is 1. 4's boss is 3.
  - They are different groups. We merge them!
  - Boss 1 becomes the boss of Boss 3. Now {1, 2, 3, 4} are one big group.

Operations:
1. `Find(x)`: Who is the ultimate Boss of x? (Follow the chain of command up).
2. `Union(x, y)`: Merge the group of x and the group of y.

Optimization (Path Compression):
When 4 asks "Who is my boss?", he goes 4 -> 3 -> 1.
Next time, 4 remembers "1 is my boss". The path becomes 4 -> 1. This flattens the hierarchy, making future lookups O(1).

PROBLEM:
"Number of Provinces" (Simplified)
Given n cities and edges, count connected components.

ALGORITHM:
1. Init `parent` array where parent[i] = i (Everyone is their own boss).
2. For each edge (u, v):
   - `rootU = Find(u)`
   - `rootV = Find(v)`
   - If `rootU != rootV`: `parent[rootU] = rootV` (Merge). Decrement total components.
*/

type UnionFind struct {
	parent []int
	count  int // Number of disjoint sets (provinces)
}

func NewUnionFind(n int) *UnionFind {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &UnionFind{parent: p, count: n}
}

// Find with Path Compression
// Recursively finds the root and points current node directly to root.
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression logic
	}
	return uf.parent[x]
}

// Union
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		// Merge sets: Make one root the parent of the other
		uf.parent[rootX] = rootY
		uf.count-- // We merged two groups, so total groups decrease by 1
	}
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := NewUnionFind(n)

	// DRY RUN:
	// A -- B
	//      |
	//      C    D (Alone)
	//
	// Init: 4 groups. {A},{B},{C},{D}.
	// Edge A-B: Union(A,B). Parent[A]=B. Count=3. Group {A,B}, {C}, {D}.
	// Edge B-C: Union(B,C). Parent[B]=C. Count=2. Group {A,B,C}, {D}.
	// Result: 2 provinces.

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}

	return uf.count
}

func main() {
	// Matrix:
	// 1 1 0
	// 1 1 0
	// 0 0 1
	// (Node 0 and 1 connected. Node 2 alone).
	input := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}

	fmt.Printf("Input Matrix: %v\n", input)
	result := findCircleNum(input)
	fmt.Printf("Number of Provinces: %d\n", result) // Expected: 2
}
