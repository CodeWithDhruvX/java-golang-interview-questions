package main

import "fmt"

// Pattern: Backtracking
// Difficulty: Medium
// Key Concept: Exploring all potential solutions by building them piece by piece, and undoing ("backtracking") when a path fails.

/*
INTUITION:
Imagine a maze. You stand at a junction.
- You choose to go Left. You mark "I went Left".
- You walk until you hit a dead end.
- You walk back ("Backtrack") to the junction.
- You un-mark "I went Left".
- Now you choose to go Right.

This "Choose -> Explore -> Unchoose" structure is Backtracking.
It explores the state space tree (DFS).

PROBLEM:
"Permutations"
Given an array nums of distinct integers, return all the possible permutations.

ALGORITHM:
1. `backtrack(path, userMap)`:
   - Base Case: If `len(path) == len(nums)`, we have a full permutation! Add copy of `path` to results. Return.
   - Recursive Step: Loop through all numbers in `nums`.
     - If number is NOT used:
       - 1. **Choose**: Add to `path`, Mark as used.
       - 2. **Explore**: Call `backtrack`.
       - 3. **Unchoose**: Remove from `path`, Mark as unused. (This restores state for the next loop iteration).
*/

func permute(nums []int) [][]int {
	results := [][]int{}
	used := make([]bool, len(nums)) // Track if number is already in current path

	// Helper closure
	var backtrack func(path []int)
	backtrack = func(path []int) {
		// Base Case
		if len(path) == len(nums) {
			// Must make a deep copy, otherwise 'path' will change later
			temp := make([]int, len(path))
			copy(temp, path)
			results = append(results, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if !used[i] {
				// Choose
				used[i] = true
				path = append(path, nums[i])

				// Explore
				backtrack(path)

				// Unchoose (Backtrack)
				// Remove last element
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}

	backtrack([]int{})
	return results
}

func main() {
	input := []int{1, 2, 3}
	fmt.Printf("Input: %v\n", input)

	/*
		DRY RUN (Partial):
		- Start: []
		  - Pick 1. Path: [1]. Used: {T, F, F}
		    - Pick 2. Path: [1, 2]. Used: {T, T, F}
		      - Pick 3. Path: [1, 2, 3]. Complete! Add to result.
		      - Backtrack 3. Used: {T, T, F}. Path: [1, 2].
		    - Pick 3. Path: [1, 3]. Used: {T, F, T}
		      - Pick 2. Path: [1, 3, 2]. Complete! Add to result.
		      - Backtrack 2. Path: [1, 3].
		  - Backtrack 1. Path: [].
		  - Pick 2... (and so on)
	*/

	res := permute(input)
	fmt.Printf("Permutations:\n%v\n", res)
}
