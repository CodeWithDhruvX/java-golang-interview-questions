package main

import (
	"fmt"
)

// Pattern: State Compression DP (Bitmask DP)
// Difficulty: Hard
// Key Concept: Using bits to represent a set of items (true/false) when N is small (<= 20), reducing O(N!) or exponential recursions.

/*
INTUITION:
Imagine you need to hire a team to cover a set of skills: {Go, Java, Python}.
Candidate A has {Go}.
Candidate B has {Java, Python}.
Candidate C has {Go, Python}.

You want the SMALLEST team.
Instead of passing a Set<String> around (slow), we use an integer (Bitmask).
skills = {Go: 0, Java: 1, Python: 2} -> Target Mask = 111 (binary) = 7.

Candidate A mask = 001 (1).
Candidate B mask = 110 (6).
Candidate C mask = 101 (5).

DP State: `dp[mask]` = Smallest team indices to achieve `mask`.
Transition: To reach `newMask = oldMask | personMask`, we check if adding `person` to `dp[oldMask]` gives a smaller team than what `dp[newMask]` currently has.

PROBLEM:
LeetCode 1125. Smallest Sufficient Team
Given a list of `req_skills` and a list of `people` (each person has some skills), return the smallest sufficient team.

ALGORITHM:
1. Map each skill to an index 0..M-1.
2. Convert each person's skills into a bitmask.
3. Initialize `dp` array of size `2^M`. `dp[0]` = empty list. All others = infinite/nil.
   - Using map `dp[mask] -> []int` might be easier for non-continuous states, but array is efficient here.
4. Iterate through each `person`:
   - For each existing `mask` in `dp`:
     - `newMask = mask | person_mask`
     - If `len(dp[mask]) + 1 < len(dp[newMask])`:
       - `dp[newMask] = dp[mask] + person`
5. Note: Iterating people then masks ensures we don't use the same person twice for the same state update in one go (similar to Knapsack).
*/

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	m := len(req_skills)
	targetMask := (1 << m) - 1

	// Map skill name to bit index
	skillId := make(map[string]int)
	for i, s := range req_skills {
		skillId[s] = i
	}

	// DP table: mask -> list of person indices
	// We want to minimize length.
	dp := make(map[int][]int)
	dp[0] = []int{}

	for i, pSkills := range people {
		personMask := 0
		for _, s := range pSkills {
			if id, exists := skillId[s]; exists {
				personMask |= (1 << id)
			}
		}

		if personMask == 0 {
			continue // Person has no relevant skills
		}

		// Try to extend all existing teams with this person
		// Note: We copy the map to avoid Concurrent Modification issues logic (reading while writing),
		// but in Go valid to read original map and write to it? No, iterating map is random.
		// Better: collect updates then apply. Or iterate 0 to 2^M.
		// Since M is small (16), 2^16 = 65536.
		// Iterating array is safer and more ordered.

		// Let's use array instead of map for iterating "masks".
		// But array is sparse? No, we build up.
		// Correct Knapsack logic: Iterate over `dp` copy or just iterate mask `target` down to 0?
		// Actually, standard BFS-like approach works well too, or just iterate current reachable masks.

		// Let's grab current reachable masks.
		currentMasks := make([]int, 0, len(dp))
		for mask := range dp {
			currentMasks = append(currentMasks, mask)
		}

		for _, mask := range currentMasks {
			newMask := mask | personMask
			if newMask == mask {
				continue // Person adds nothing new to this specific set
			}

			currentTeam := dp[mask]
			newTeamSize := len(currentTeam) + 1

			existingTeam, exists := dp[newMask]
			if !exists || newTeamSize < len(existingTeam) {
				// Found a better way to reach newMask!
				// Create new slice
				newRes := make([]int, len(currentTeam))
				copy(newRes, currentTeam)
				newRes = append(newRes, i)
				dp[newMask] = newRes
			}
		}
	}

	return dp[targetMask]
}

func main() {
	// Skills: ["java", "nodejs", "react"]
	// 0: "java", "react" (Mask 101)
	// 1: "nodejs" (Mask 010)
	// 2: "nodejs", "react" (Mask 110)
	// Target: 111 (7)
	// Person 0 covers {java, react}. Remaining: {nodejs}. Need person 1 or 2.
	// 0+1 -> {java, react, nodejs} (Size 2)
	// 0+2 -> {java, react, nodejs} (Size 2)
	// Best: [0, 1] or [0, 2]

	req := []string{"java", "nodejs", "react"}
	people := [][]string{
		{"java", "react"},
		{"nodejs"},
		{"nodejs", "react"},
	}

	fmt.Printf("Req: %v\n", req)
	fmt.Printf("Team Indices: %v\n", smallestSufficientTeam(req, people))

	// Example 2
	// req: ["algorithms", "math", "java", "react", "c++", "aws"]
	// ... large example ommitted for brevity
}
