package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Pattern: Randomized Algorithms (QuickSelect)
// Difficulty: Medium
// Key Concept: Using randomization to achieve O(N) average efficiency in selection/sorting problems, avoiding O(N^2) worst cases of deterministic pivots.

/*
INTUITION:
Quicksort is O(N log N) on average, but O(N^2) if the pivot is bad (already sorted).
QuickSelect finds the K-th smallest/largest element.
If we sort: O(N log N).
If we just partition:
- Pick pivot. Partition array.
- If pivot index == K, we found it!
- If index < K, recurse Right.
- If index > K, recurse Left.
Average complexity: N + N/2 + N/4... = 2N = O(N).
Worst case: O(N^2).
Random Shuffle or Random Pivot fixes the worst case for almost all practical inputs.

PROBLEM:
LeetCode 215. Kth Largest Element in an Array.
Given an integer array nums and an integer k, return the kth largest element in the array.
Note that it is the kth largest element in the sorted order, not the kth distinct element.

ALGORITHM:
1. `findKthLargest(nums, k)`: Target index `len(nums) - k`.
2. `quickSelect(l, r, target)`:
   - Pick random pivot `p` between l and r.
   - Swap pivot to end.
   - Partition: Put all elements < pivot to left.
   - Place pivot at correct position `finalP`.
   - If `finalP == target`, return `nums[finalP]`.
   - If `finalP < target`, recurse `quickSelect(finalP+1, r, target)`.
   - Else, recurse `quickSelect(l, finalP-1, target)`.
*/

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	// K-th largest is the element that would be at index (N - k) in sorted array
	targetIndex := len(nums) - k
	return quickSelect(nums, 0, len(nums)-1, targetIndex)
}

func quickSelect(nums []int, l, r, k int) int {
	if l == r {
		return nums[l]
	}

	// Random pivot index
	pivotIndex := l + rand.Intn(r-l+1)

	// Move pivot to end
	nums[pivotIndex], nums[r] = nums[r], nums[pivotIndex]

	// Partition
	// Pivot is now at nums[r]
	pivot := nums[r]
	storeIndex := l
	for i := l; i < r; i++ {
		if nums[i] < pivot {
			nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
			storeIndex++
		}
	}
	// Move pivot to its final place
	nums[storeIndex], nums[r] = nums[r], nums[storeIndex]

	if storeIndex == k {
		return nums[storeIndex]
	} else if storeIndex < k {
		return quickSelect(nums, storeIndex+1, r, k)
	} else {
		return quickSelect(nums, l, storeIndex-1, k)
	}
}

func main() {
	// [3,2,1,5,6,4], k=2
	// Sorted: 1,2,3,4,5,6. 2nd largest is 5.
	// Target Index: 6 - 2 = 4 (0-indexed: 4).

	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Printf("Nums: %v, K: %d, Ans: %d\n", nums, k, findKthLargest(nums, k))

	// [3,2,3,1,2,4,5,5,6], k=4
	// Sorted: 1,2,2,3,3,4,5,5,6
	// 4th largest: 4.
	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4
	fmt.Printf("Nums: %v, K: %d, Ans: %d\n", nums2, k2, findKthLargest(nums2, k2))
}
