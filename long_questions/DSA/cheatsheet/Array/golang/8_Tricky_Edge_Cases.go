package main

import (
	"math"
	"sort"
)

// 1. Shortest Unsorted Continuous Subarray
// Time: O(N), Space: O(1)
func ShortestUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	end := -1
	maxVal := math.MinInt64

	for i := 0; i < n; i++ {
		if nums[i] >= maxVal {
			maxVal = nums[i]
		} else {
			end = i
		}
	}

	start := 0
	minVal := math.MaxInt64
	for i := n - 1; i >= 0; i-- {
		if nums[i] <= minVal {
			minVal = nums[i]
		} else {
			start = i
		}
	}

	if end == -1 {
		return 0
	}
	return end - start + 1
}

// 2. Check if Sorted by One Swap
// Time: O(N), Space: O(1)
func CheckSortedOneSwap(arr []int) bool {
	n := len(arr)
	type pair struct {
		first, second int
	}
	var inversions []int // storing indices i where arr[i] > arr[i+1]

	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			inversions = append(inversions, i)
		}
	}

	if len(inversions) == 0 {
		return true
	}
	if len(inversions) > 2 {
		return false
	}

	// If 2 inversions: swap first element of first inv with second element of second inv
	// If 1 inversion: swap the two involved elements

	// Create a copy to swap and check
	// (Actual swap logic slightly complex to generalize, verifying "sorted after swap" is easier)
	// We need to actually swap and check. To avoid modifying input, we copy.
	temp := make([]int, n)
	copy(temp, arr)

	if len(inversions) == 1 {
		i := inversions[0]
		temp[i], temp[i+1] = temp[i+1], temp[i]
	} else {
		// 2 inversions
		i := inversions[0]
		j := inversions[1]
		// Swap arr[i] and arr[j+1]
		temp[i], temp[j+1] = temp[j+1], temp[i]
	}

	for i := 0; i < n-1; i++ {
		if temp[i] > temp[i+1] {
			return false
		}
	}
	return true
}

// 3. Check if Sorted by Reversing One Subarray
// Time: O(N), Space: O(1) (ignoring copy)
func CheckSortedReverseSubarray(arr []int) bool {
	n := len(arr)
	// Find first decreasing part
	start := -1
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			start = i
			break
		}
	}
	if start == -1 {
		return true // Already sorted
	}

	end := -1
	for i := n - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			end = i
			break
		}
	}

	// Reverse start...end
	temp := make([]int, n)
	copy(temp, arr)
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]
	}

	// Check if sorted
	for i := 0; i < n-1; i++ {
		if temp[i] > temp[i+1] {
			return false
		}
	}
	return true
}

// 4. First Missing Positive Integer
// Time: O(N), Space: O(1)
func FirstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// Cyclic sort: Place nums[i] at index nums[i]-1
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// swap nums[i] and nums[nums[i]-1]
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

// 5. Elements Appearing > N/K Times
// Time: O(N * K), Space: O(K)
func ElementMoreThanNK(arr []int, k int) []int {
	n := len(arr)
	if k < 2 {
		return arr // Or specific logic
	}

	// Candidates map: val -> count
	candidates := make(map[int]int)

	for _, val := range arr {
		if _, exists := candidates[val]; exists {
			candidates[val]++
		} else {
			if len(candidates) < k-1 {
				candidates[val] = 1
			} else {
				// Decrement all
				toDelete := []int{}
				for c := range candidates {
					candidates[c]--
					if candidates[c] == 0 {
						toDelete = append(toDelete, c)
					}
				}
				for _, c := range toDelete {
					delete(candidates, c)
				}
			}
		}
	}

	// Verification
	res := []int{}
	for c := range candidates {
		count := 0
		for _, val := range arr {
			if val == c {
				count++
			}
		}
		if count > n/k {
			res = append(res, c)
		}
	}
	return res
}

// 6. Max Sum Such That No Two Elements Are Adjacent (House Robber)
// Time: O(N), Space: O(1)
func MaxSumNoAdjacent(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	prev2 := 0
	prev1 := arr[0] // Assuming positive numbers for simplicity, or max(0, arr[0])

	for i := 1; i < len(arr); i++ {
		pick := arr[i] + prev2
		skip := prev1
		curr := pick
		if skip > pick {
			curr = skip
		}
		prev2 = prev1
		prev1 = curr
	}
	return prev1
}

// 7. Subarray Sum = 0 (Space Optimized? O(N log N))
// Time: O(N log N), Space: O(N) (for prefix array copy)
func SubarraySumZeroOptimized(arr []int) bool {
	n := len(arr)
	prefix := make([]int, n)
	sum := 0
	for i, v := range arr {
		sum += v
		if sum == 0 {
			return true
		}
		prefix[i] = sum
	}

	sort.Ints(prefix)
	for i := 0; i < n-1; i++ {
		if prefix[i] == prefix[i+1] {
			return true
		}
	}
	return false
}

// 8. Rearrange such that arr[i] = i
// Time: O(N), Space: O(1)
func IndexMapping(arr []int) {
	for i := 0; i < len(arr); i++ {
		for arr[i] != -1 && arr[i] != i {
			target := arr[i]
			// Swap arr[i] with arr[target]
			arr[i], arr[target] = arr[target], arr[i]
		}
	}
}
