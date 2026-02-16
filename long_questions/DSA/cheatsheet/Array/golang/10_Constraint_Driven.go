package main

import (
	"math"
)

// 1. Find Duplicate Number (Floyd's Cycle Detection)
// Time: O(N), Space: O(1)
// Constraint: Array contains n+1 integers where each integer is in range [1, n].
func FindDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	// Phase 1: Detect cycle
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// Phase 2: Find entrance
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// 2. Find Missing and Repeating Number
// Time: O(N), Space: O(1)
// Using equation method (sum and sum of squares) to avoid modifying array
// Or using index marking if modification is allowed.
// Let's use Index Marking for O(N) simple logic, assuming modification allowed or copy.
// If purely "Read-Only", Math approach is better but risks overflow.
// Let's use Math approach with int64 to be safe.
func FindMissingAndRepeating(arr []int) (int, int) {
	n := len(arr)
	s := int64(n * (n + 1) / 2)
	sSq := int64(n * (n + 1) * (2*n + 1) / 6)

	actualSum := int64(0)
	actualSumSq := int64(0)

	for _, val := range arr {
		actualSum += int64(val)
		actualSumSq += int64(val) * int64(val)
	}

	// val1 = missing, val2 = repeating
	// s - actualSum = val1 - val2
	diff := s - actualSum // val1 - val2

	// sSq - actualSumSq = val1^2 - val2^2 = (val1 - val2)(val1 + val2)
	diffSq := sSq - actualSumSq

	sum := diffSq / diff // val1 + val2

	missing := (diff + sum) / 2
	repeating := sum - missing

	return int(missing), int(repeating)
}

// 3. Maximum Sum Subarray with at least K elements
// Time: O(N), Space: O(N)
func MaxSumSubarrayAtLeastK(arr []int, k int) int {
	n := len(arr)
	if n < k {
		return 0 // or error
	}

	// Kadane's window array
	maxSum := make([]int, n)
	currentSum := 0
	// Standard Kadane's from left
	for i := 0; i < n; i++ {
		currentSum += arr[i]
		maxSum[i] = currentSum
		if currentSum < 0 {
			currentSum = 0
		}
	}
	// Actually we need a sliding window approach combined.
	// 1. Compute prefix sum
	prefixSum := make([]int, n)
	prefixSum[0] = arr[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}

	// Sum of first k elements
	res := prefixSum[k-1]

	// Sliding window
	// using a variable to store min prefix sum seen so far before the window
	// window sum ending at i of length >= k is prefixSum[i] - minPrefixSum[i-k]

	// Better approach:
	// Answer is max(prefixSum[i] - min_prefix_sum_before_i_minus_k)

	minPrefix := 0
	// We iterate i from k-1 to n-1
	// The window ends at i. It starts at j <= i - k + 1.
	// Sum = prefixSum[i] - prefixSum[j-1].
	// To maximize sum, minimize prefixSum[j-1].
	// j-1 ranges from -1 to i - k.

	currentMinPrefix := 0 // represents min prefix sum for indices < i-k
	// At i=k-1, we consider window [0...k-1]. Sum is prefixSum[k-1] - prefix(negative 1 -> 0).

	for i := k; i < n; i++ {
		// update min prefix for the index just leaving the mandated window gap
		// index leaving is i-k-1. If i=k, leaving is -1.
		// Wait, logic:
		// We want max(prefixSum[i] - prefixSum[j]) where j <= i-k.
		// For i=k, max is prefixSum[k] - min(prefixSum[0], prefixSum[-1]=0)

		// Let's re-verify loop
		// Initialize res with first K window
		res = int(math.Max(float64(res), float64(prefixSum[i]-minPrefix)))

		// Update minPrefix for next iteration
		// The value available to subtract becomes prefixSum[i-k+1-1] i.e. prefixSum[i-k]
		if i-k >= 0 {
			if prefixSum[i-k] < minPrefix {
				minPrefix = prefixSum[i-k]
			}
		}

		val := prefixSum[i] - minPrefix
		if val > res {
			res = val
		}
	}
	return res
}

// 4. Subarray with Equal Number of 0s, 1s, and 2s
// Time: O(N), Space: O(N)
func LongestEqual012(arr []int) int {
	// Treat as differences.
	// let c0, c1, c2 be counts.
	// We want c0=c1=c2.
	// This implies c0-c1 = 0 AND c1-c2 = 0.
	// Store pair (c0-c1, c1-c2) in map.

	m := make(map[string]int)
	c0, c1, c2 := 0, 0, 0
	m["0,0"] = -1
	maxLen := 0

	for i, val := range arr {
		if val == 0 {
			c0++
		}
		if val == 1 {
			c1++
		}
		if val == 2 {
			c2++
		}

		diff1 := c0 - c1
		diff2 := c1 - c2
		key := string(diff1) + "," + string(diff2) // Simplified key gen
		// Better key generation to avoid collision/unicode issues:
		// fmt.Sprintf("%d,%d", diff1, diff2) is safer but slower.
		// Since diffs can be negative, standard string conversion is safe.

		// Optimizing key gen:
		// Just using a simplistic struct if Go allowed map keys as structs (it does!)
		type Key struct{ d1, d2 int }
		k := Key{diff1, diff2}

		// But let's stick to string map for simplicity in snippet if struct definition is verbose
		// key := fmt.Sprintf("%d,%d", diff1, diff2) // Needs fmt

		// Let's use the struct approach inside function
	}
	// Re-writing with struct
	return longestEqual012Helper(arr)
}

func longestEqual012Helper(arr []int) int {
	type Key struct {
		d1, d2 int
	}
	m := make(map[Key]int)
	m[Key{0, 0}] = -1

	c0, c1, c2 := 0, 0, 0
	maxLen := 0

	for i, val := range arr {
		if val == 0 {
			c0++
		}
		if val == 1 {
			c1++
		}
		if val == 2 {
			c2++
		}

		key := Key{c0 - c1, c1 - c2}

		if idx, ok := m[key]; ok {
			if i-idx > maxLen {
				maxLen = i - idx
			}
		} else {
			m[key] = i
		}
	}
	return maxLen
}

// 5. Count Subarrays with Product < K
// Time: O(N), Space: O(1)
func NumSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	prod := 1
	ans := 0
	left := 0

	for right, val := range nums {
		prod *= val
		for prod >= k && left <= right {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}
