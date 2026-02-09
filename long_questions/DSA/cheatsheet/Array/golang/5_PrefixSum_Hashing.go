package main

// 1. Range Sum Queries (Immutable Array)
// Time: O(N) build, O(1) query, Space: O(N)
type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	prefixSum := make([]int, n)
	if n > 0 {
		prefixSum[0] = nums[0]
		for i := 1; i < n; i++ {
			prefixSum[i] = prefixSum[i-1] + nums[i]
		}
	}
	return NumArray{prefixSum: prefixSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.prefixSum[right]
	}
	return this.prefixSum[right] - this.prefixSum[left-1]
}

// 2. Find Subarray with Sum = 0
// Time: O(N), Space: O(N)
func SubarraySumZero(arr []int) bool {
	sum := 0
	visited := make(map[int]bool)
	visited[0] = true // To handle subarray starting from index 0

	for _, val := range arr {
		sum += val
		if visited[sum] {
			return true
		}
		visited[sum] = true
	}
	return false
}

// 3. Count Subarrays with Sum = K
// Time: O(N), Space: O(N)
func CountSubarraySumK(arr []int, k int) int {
	freq := make(map[int]int)
	freq[0] = 1 // Base case for subarrays starting at index 0
	sum := 0
	count := 0
	for _, val := range arr {
		sum += val
		if c, ok := freq[sum-k]; ok {
			count += c
		}
		freq[sum]++
	}
	return count
}

// 4. Longest Subarray with Distinct Elements
// Time: O(N), Space: O(N)
func LongestDistinctSubarray(arr []int) int {
	lastSeen := make(map[int]int)
	left := 0
	maxLen := 0
	for right, val := range arr {
		if idx, ok := lastSeen[val]; ok {
			if idx+1 > left {
				left = idx + 1
			}
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		lastSeen[val] = right
	}
	return maxLen
}

// 5. Subarray with Given XOR
// Time: O(N), Space: O(N)
func SubarrayXorB(arr []int, b int) int {
	freq := make(map[int]int)
	freq[0] = 1
	xorSum := 0
	count := 0
	for _, val := range arr {
		xorSum ^= val
		if c, ok := freq[xorSum^b]; ok {
			count += c
		}
		freq[xorSum]++
	}
	return count
}

// 6. Count Subarrays with Equal Odd and Even Numbers
// Time: O(N), Space: O(N)
func CountOddEvenSubarrays(arr []int) int {
	freq := make(map[int]int)
	freq[0] = 1
	sum := 0
	count := 0
	for _, val := range arr {
		// Transform: Odd -> 1, Even -> -1
		if val%2 != 0 {
			sum += 1
		} else {
			sum += -1
		}
		if c, ok := freq[sum]; ok {
			count += c
		}
		freq[sum]++
	}
	return count
}
