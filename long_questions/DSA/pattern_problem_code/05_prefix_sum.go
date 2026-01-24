package main

import "fmt"

// Pattern: Prefix Sum
// Difficulty: Easy/Medium
// Key Concept: Pre-calculate cumulative sums to allow O(1) range sum queries.

/*
INTUITION:
If I ask you "How much money did you earn from Day 3 to Day 5?", you normally add: Earnings[3] + Earnings[4] + Earnings[5].
This is O(N) if the range is big.

But if you keep a running ledger:
- Total by Day 1: $10
- Total by Day 2: $30
- Total by Day 3: $40
- Total by Day 4: $60
- Total by Day 5: $100

Then "Day 3 to Day 5" is simply: (Total by Day 5) - (Total by Day 2).
$100 - $30 = $70.
This subtraction is O(1) (Instant), no matter how big the range is.

PROBLEM:
"Range Sum Query - Immutable"
Given an integer array nums, handle multiple queries of the following type:
Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.

ALGORITHM:
1. Build a `prefixSum` array of size N+1.
2. `prefixSum[0]` is roughly 0 (base case).
3. `prefixSum[i+1] = prefixSum[i] + nums[i]`.
4. To get sum from `L` to `R`: Return `prefixSum[R+1] - prefixSum[L]`.
*/

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	// Step 1: Create array of size N+1
	// We use N+1 so that prefixSum[0] = 0, handling the "sum from 0 to X" case easily.
	p := make([]int, len(nums)+1)

	// Step 2: Build cumulative sums
	// nums: [1, 2, 3]
	// p[0] = 0
	// p[1] = 0 + 1 = 1
	// p[2] = 1 + 2 = 3
	// p[3] = 3 + 3 = 6
	for i := 0; i < len(nums); i++ {
		p[i+1] = p[i] + nums[i]
	}

	return NumArray{prefixSum: p}
}

func (this *NumArray) SumRange(left int, right int) int {
	// Formula: P[Right + 1] - P[Left]
	// DRY RUN: Range [1, 2] (values 2, 3) -> Expect 5
	// right=2 -> P[3] (which is 6)
	// left=1  -> P[1] (which is 1)
	// Result = 6 - 1 = 5. Correct.
	return this.prefixSum[right+1] - this.prefixSum[left]
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	fmt.Printf("Input Array: %v\n", nums)

	obj := Constructor(nums)

	// Query 1: Sum Range [0, 2] -> (-2) + 0 + 3 = 1
	fmt.Printf("Sum [0, 2]: %d\n", obj.SumRange(0, 2))

	// Query 2: Sum Range [2, 5] -> 3 + (-5) + 2 + (-1) = -1
	fmt.Printf("Sum [2, 5]: %d\n", obj.SumRange(2, 5))

	/*
		Why is this powerful?
		If we had 1,000,000 queries, doing a "for loop" for each query would be effectively infinite time.
		With Prefix Sum, 1,000,000 queries take just 1,000,000 subtractions. Very fast.
	*/
}
