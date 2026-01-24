package main

import (
	"fmt"
	"math"
	"sort"
)

// Pattern: Mo's Algorithm
// Difficulty: Hard
// Key Concept: Offline query processing by reordering queries to minimize pointer movement (Square Root Decomposition).

/*
INTUITION:
We have an array `A` and 100,000 queries `(L, R)`.
Standard approach: Loop L to R for each query. O(N * Q) -> TLE.
Mo's Algorithm:
If we answer query [1, 10], and the next query is [2, 11], we don't restart. We just remove index 1 and add index 11.
We sort the queries in a specific order so that the total movement of L and R is minimized.
Sorting Order:
- Divide array into blocks of size `sqrt(N)`.
- Sort by `L / BlockSize`.
- If blocks are same, sort by `R`.

Complexity: O((N + Q) * sqrt(N)).

PROBLEM:
Count Distinct Integers in Range (D-Query type).
Given `nums` and list of queries `[L, R]`, return count of unique numbers in each range.

ALGORITHM:
1. Define `Query` struct {L, R, ID}.
2. Sort queries using Mo's Comparator.
3. Maintain global `count[]` (frequency of each number) and `distinct` (current distinct count).
4. `add(x)`: if count[x] == 0, distinct++. count[x]++.
5. `remove(x)`: count[x]--. if count[x] == 0, distinct--.
6. Use `currL`, `currR`. For each query:
   - Move `currL` to `q.L` (Add/Remove).
   - Move `currR` to `q.R` (Add/Remove).
   - Store answer[q.ID] = distinct.
*/

type Query struct {
	L, R, ID, Block int
}

func countDistinct(nums []int, queries [][]int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	// Block size
	blockSize := int(math.Sqrt(float64(n)))

	qs := make([]Query, len(queries))
	for i, q := range queries {
		qs[i] = Query{
			L:     q[0],
			R:     q[1],
			ID:    i,
			Block: q[0] / blockSize,
		}
	}

	// Sort Queries
	sort.Slice(qs, func(i, j int) bool {
		if qs[i].Block != qs[j].Block {
			return qs[i].Block < qs[j].Block
		}
		// Optimization: if block is odd, sort R ascending; if even, descending (Hilbert curve effect).
		// Standard ascending R is fine.
		return qs[i].R < qs[j].R
	})

	// Current state
	currL, currR := 0, -1
	currentDistinct := 0
	freq := make(map[int]int) // Or array if values are small (0-10^5)

	add := func(idx int) {
		val := nums[idx]
		if freq[val] == 0 {
			currentDistinct++
		}
		freq[val]++
	}

	remove := func(idx int) {
		val := nums[idx]
		freq[val]--
		if freq[val] == 0 {
			currentDistinct--
		}
	}

	ans := make([]int, len(queries))

	for _, q := range qs {
		// Adjust R
		for currR < q.R {
			currR++
			add(currR)
		}
		for currR > q.R {
			remove(currR)
			currR--
		}

		// Adjust L
		for currL < q.L {
			remove(currL)
			currL++
		}
		for currL > q.L {
			currL--
			add(currL)
		}

		ans[q.ID] = currentDistinct
	}

	return ans
}

func main() {
	// Nums: [1, 1, 2, 1, 3]
	// queries: [0, 4], [1, 3], [2, 4]
	// q1 [0,4]: 1,1,2,1,3. Distinct: {1,2,3} -> 3
	// q2 [1,3]: 1,2,1. Distinct: {1,2} -> 2
	// q3 [2,4]: 2,1,3. Distinct: {1,2,3} -> 3

	nums := []int{1, 1, 2, 1, 3}
	queries := [][]int{
		{0, 4},
		{1, 3},
		{2, 4},
	}

	fmt.Printf("Nums: %v\n", nums)
	fmt.Printf("Answers: %v\n", countDistinct(nums, queries))
}
