package main

import (
	"fmt"
	"math"
	"sort"
)

// Pattern: Meet-in-the-Middle
// Difficulty: Hard
// Key Concept: Reduce time complexity from O(2^N) to O(2^(N/2)) by splitting the problem, solving halves, and merging.

/*
INTUITION:
You have 30 coins. You want to divide them into two piles with equal count (15 each) such that the difference in value is minimized.

Brute Force: Try every combination of picking 15 from 30.
Combinations C(30, 15) â‰ˆ 155,000,000. Too slow for a 1-second limit.

Meet-in-the-Middle trick:
1. Split the coins into two halves: Left (15 coins) and Right (15 coins).
2. Generate ALL possible sums for the Left half. Store them grouped by "how many coins used".
   - e.g., LeftSums[3] = list of all sums using exactly 3 coins from Left.
   - This takes 2^15 â‰ˆ 32,000 operations. Super fast.
3. Do the same for Right half.
4. Now, to make a full pile of 15 coins:
   - If we take `k` coins from Left, we MUST take `15-k` coins from Right.
   - We want (SumLeft + SumRight) to be as close to TotalSum/2 as possible.
5. For each `k`, iterate through LeftSums[k]. Use Binary Search on RightSums[15-k] to find the best match.

PROBLEM:
LeetCode 2035. Partition Array Into Two Arrays to Minimize Sum Difference.
Given an integer array nums of 2n integers, divide nums into two arrays of length n to minimize the absolute difference of their sums.

ALGORITHM:
1. Split `nums` into `leftPart` and `rightPart`.
2. Generate `leftDist[k]` = sorted list of sums using `k` elements from `leftPart`.
3. Generate `rightDist[k]` = sorted list of sums using `k` elements from `rightPart`.
4. Target sum for one partition is `TotalSum / 2`.
5. Iterate `k` from 0 to n:
   - For each `sumL` in `leftDist[k]`:
     - Need `sumR` from `rightDist[n-k]` such that `sumL + sumR` is close to Target.
     - Use Binary Search (lower_bound) on `rightDist[n-k]` to find `target - sumL`.
     - Update min difference.
*/

func minimumDifference(nums []int) int {
	n := len(nums) / 2
	sumTotal := 0
	for _, x := range nums {
		sumTotal += x
	}
	target := sumTotal / 2 // We want one partition sum to be close to this

	leftPart := nums[:n]
	rightPart := nums[n:]

	// Helper to generate all subset sums grouped by count
	generateSums := func(arr []int) map[int][]int {
		// Map: count -> list of sums
		sums := make(map[int][]int)
		sums[0] = []int{0}

		for _, val := range arr {
			// Iterate backwards to avoid using same element twice for same count in one pass (DP style)
			// Or just build a new map layer by layer
			// Since N is small (15), we can just iterate.
			// Better: Recursive/Bitmask approach is cleaner for "all subsets".
			// Given N=15, simple recursion is fine.
			// But let's use the iterative "add to existing" approach.
			newSums := make(map[int][]int)
			for count, existingSums := range sums {
				for _, s := range existingSums {
					newSums[count] = append(newSums[count], s)         // Don't include val
					newSums[count+1] = append(newSums[count+1], s+val) // Include val
				}
			}
			// Does this logic duplicate?
			// No, the standard way is: Start with {0: [0]}. For x in arr: for count in desc: add x.
			// Let's stick to true generation logic.
		}
		return nil // Placeholder, implemented properly below
	}

	// Real generator
	getSums := func(arr []int) map[int][]int {
		res := make(map[int][]int)
		res[0] = []int{0}

		for _, num := range arr {
			for count := len(res) - 1; count >= 0; count-- {
				for _, s := range res[count] {
					res[count+1] = append(res[count+1], s+num)
				}
			}
		}
		// Sort each list for binary search
		for k := range res {
			sort.Ints(res[k])
		}
		return res
	}

	leftSums := getSums(leftPart)
	rightSums := getSums(rightPart)

	minDiff := math.MaxInt32

	// Iterate k (count taken from left)
	for k := 0; k <= n; k++ {
		// We need n-k from right
		if _, ok := rightSums[n-k]; !ok {
			continue
		}

		rSums := rightSums[n-k]

		for _, lSum := range leftSums[k] {
			needed := target - lSum

			// Binary Search in rSums for closest to 'needed'
			idx := sort.SearchInts(rSums, needed)

			// Check idx and idx-1
			candidates := []int{}
			if idx < len(rSums) {
				candidates = append(candidates, rSums[idx])
			}
			if idx > 0 {
				candidates = append(candidates, rSums[idx-1])
			}

			for _, rSum := range candidates {
				totalPartSum := lSum + rSum
				otherPartSum := sumTotal - totalPartSum
				diff := abs(totalPartSum - otherPartSum)
				if diff < minDiff {
					minDiff = diff
				}
			}
		}
	}

	return minDiff
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	// N=3. Total 6. Split into 2 arrays of size 3.
	// [3, 9, 7, 3] -> n=2. Total 22. Target 11.
	// Split: [3, 9] and [7, 3].
	// L: 0->[0], 1->[3,9], 2->[12]
	// R: 0->[0], 1->[7,3], 2->[10]
	// k=0 (L=0, R=10): Sum=10. Diff=|10-12|=2.
	// k=1 (L=[3,9], R=[7,3]): 3+7=10, 3+3=6, 9+7=16, 9+3=12. Closest to 11 is 12 (Diff |12-10|=2) or 10.
	// k=2 (L=12, R=0): Sum=12. Diff=2.
	// Ans: 2.
	nums := []int{3, 9, 7, 3}
	fmt.Printf("Nums: %v\n", nums)
	fmt.Printf("Min Diff: %d\n", minimumDifference(nums))

	// [2,-1,0,4,-2,-9] -> [2,-1,0] and [4,-2,-9]
	// Min diff 0? (Sum -6). Part1: [2,4,-9]=-3. Part2: [-1,0,-2]=-3. Diff 0.
	nums2 := []int{2, -1, 0, 4, -2, -9}
	fmt.Printf("Nums: %v\n", nums2)
	fmt.Printf("Min Diff: %d\n", minimumDifference(nums2))
}
