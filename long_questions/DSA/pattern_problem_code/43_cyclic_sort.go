package main

import (
	"fmt"
)

// Pattern: Cyclic Sort
// Difficulty: Hard
// Key Concept: Sorting numbers in the range [1, N] in O(N) time and O(1) space by placing element `x` at index `x-1`.

/*
INTUITION:
We have an unsorted array of numbers. Some might be negative, some huge.
We want to find the "First Missing Positive" (smallest positive integer not present).
Ideally, if we had `[1, 2, 3, 4]`, the answer is 5.
If we had `[3, 4, -1, 1]`:
- Ideally, index 0 should have 1.
- Index 1 should have 2.
- Index 2 should have 3.
- Index 3 should have 4.

Crucial Observation: The answer must be between 1 and N+1.
So we only care about putting numbers `1 to N` in their correct spots (`index 0 to N-1`).
We swap `nums[i]` to `nums[nums[i]-1]` if it's in the valid range.
If `nums[i] == 3`, we swap it to index 2.
We repeat this until `nums[i]` is either correct or out of range.

PROBLEM:
LeetCode 41. First Missing Positive.
Given an unsorted integer array nums, return the smallest missing positive integer.
You must implement an algorithm that runs in O(n) time and uses constant extra space.

ALGORITHM:
1. Iterate `i` from 0 to N-1.
2. While `nums[i]` is in range [1, N] AND `nums[i]` is not at correct position (`nums[nums[i]-1] != nums[i]`):
   - Swap `nums[i]` with `nums[nums[i]-1]`.
3. Iterate `i` from 0 to N-1 again.
   - If `nums[i] != i + 1`: Return `i + 1`.
4. If all match, return `N + 1`.
*/

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// Valid range: [1, n]
		// Correct index for val x is x-1
		// Check duplicates: nums[i] should not duplicate what's already at target
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// Swap nums[i] with target
			targetIdx := nums[i] - 1
			nums[i], nums[targetIdx] = nums[targetIdx], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	// [1, 2, 0] -> Sorts to [1, 2, 0].
	// i=0: val 1. Correct.
	// i=1: val 2. Correct.
	// i=2: val 0. Ignore.
	// Check: idx 0 has 1? Yes. idx 1 has 2? Yes. idx 2 has 3? No (has 0). Return 3.
	nums1 := []int{1, 2, 0}
	fmt.Printf("Nums: %v, First Missing: %d\n", nums1, firstMissingPositive(nums1))

	// [3, 4, -1, 1]
	// i=0: val 3. Swap to idx 2. -> [-1, 4, 3, 1]
	// i=0: val -1. Ignore.
	// i=1: val 4. Swap to idx 3. -> [-1, 1, 3, 4]
	// i=1: val 1. Swap to idx 0. -> [1, -1, 3, 4]
	// i=1: val -1. Ignore. (Wait, after swap, i stays at 1? The 'while' loop handles current i. Yes.)
	// i=2: val 3. Correct.
	// i=3: val 4. Correct.
	// Final: [1, -1, 3, 4].
	// Check:
	// idx 0 -> 1. OK.
	// idx 1 -> 2? No (has -1). Return 2.
	nums2 := []int{3, 4, -1, 1}
	nums2Copy := make([]int, len(nums2))
	copy(nums2Copy, nums2)
	fmt.Printf("Nums: %v, First Missing: %d\n", nums2Copy, firstMissingPositive(nums2))

	// [7,8,9,11,12] -> Miss 1.
	nums3 := []int{7, 8, 9, 11, 12}
	fmt.Printf("Nums: %v, First Missing: %d\n", nums3, firstMissingPositive(nums3))
}
