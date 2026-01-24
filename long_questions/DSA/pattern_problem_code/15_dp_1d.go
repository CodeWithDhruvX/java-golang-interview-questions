package main

import "fmt"

// Pattern: Dynamic Programming (1D)
// Difficulty: Easy
// Key Concept: Solving a problem by breaking it down into smaller overlapping subproblems and storing the results.

/*
INTUITION:
"Climbing Stairs"
You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps.
In how many distinct ways can you climb to the top?

Think backwards:
To reach Step 10, you MUST come from either Step 9 (taking 1 step) or Step 8 (taking 2 steps).
So, Ways(10) = Ways(9) + Ways(8).
This looks like Fibonacci!

Naive Recursion is O(2^N) (Terrible).
DP Memoization makes it O(N).
We can remember the answer for each step in an array `dp`.

ALGORITHM:
1. `dp[0] = 0` (or 1 depending on def), `dp[1] = 1`, `dp[2] = 2`.
2. Loop `i` from 3 to N.
3. `dp[i] = dp[i-1] + dp[i-2]`.
4. Return `dp[n]`.

Optimization:
We only need the last two numbers. We don't need the whole array.
`curr = prev1 + prev2`.
*/

func climbStairs(n int) int {
	if n <= 2 {
		return n // If 1 step -> 1 way. If 2 steps -> 2 ways (1+1, 2).
	}

	// We only store the last two values to save space (O(1) space complexity).
	oneStepBefore := 2  // Ways to reach step 2
	twoStepsBefore := 1 // Ways to reach step 1
	currentWays := 0

	// DRY RUN: n = 4
	// i=3: curr = 2 + 1 = 3 ways. (TwoSteps -> 2, OneStep -> 3)
	// i=4: curr = 3 + 2 = 5 ways.
	// Result: 5.

	for i := 3; i <= n; i++ {
		currentWays = oneStepBefore + twoStepsBefore

		// Shift our "frame" forward
		twoStepsBefore = oneStepBefore
		oneStepBefore = currentWays
	}

	return currentWays
}

func main() {
	n := 5
	/*
		Ways to climb 5:
		1. 1+1+1+1+1
		2. 1+1+1+2
		3. 1+1+2+1
		4. 1+2+1+1
		5. 2+1+1+1
		6. 1+2+2
		7. 2+1+2
		8. 2+2+1
		Total = 8.
	*/

	fmt.Printf("Stairs: %d\n", n)
	result := climbStairs(n)
	fmt.Printf("Ways to Climb: %d\n", result) // Expected: 8
}
