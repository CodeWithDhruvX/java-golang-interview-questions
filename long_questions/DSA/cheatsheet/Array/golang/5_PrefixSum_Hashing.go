package main

// 1. Range Sum Queries (Immutable Array)

// Brute Force Approach: Calculate sum for each query
// Time: O(N) per query, Space: O(1)
type NumArrayBruteForce struct {
	nums []int
}

func ConstructorBruteForce(nums []int) NumArrayBruteForce {
	return NumArrayBruteForce{nums: nums}
}

func (this *NumArrayBruteForce) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.nums[i]
	}
	return sum
}

// Optimized Approach: Prefix sum array
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

// Brute Force Approach: Check all possible subarrays
// Time: O(N^2), Space: O(1)
func SubarraySumZeroBruteForce(arr []int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		currentSum := 0
		for j := i; j < n; j++ {
			currentSum += arr[j]
			if currentSum == 0 {
				return true
			}
		}
	}
	return false
}

// Optimized Approach: Hash map with prefix sums
// Time: O(N), Space: O(N)
func SubarraySumZeroOptimized(arr []int) bool {
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

// Legacy function for backward compatibility
func SubarraySumZero(arr []int) bool {
	return SubarraySumZeroOptimized(arr)
}

// 3. Count Subarrays with Sum = K

// Brute Force Approach: Check all possible subarrays
// Time: O(N^2), Space: O(1)
func CountSubarraySumKBruteForce(arr []int, k int) int {
	count := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		currentSum := 0
		for j := i; j < n; j++ {
			currentSum += arr[j]
			if currentSum == k {
				count++
			}
		}
	}
	return count
}

// Optimized Approach: Hash map with prefix sum frequency
// Time: O(N), Space: O(N)
func CountSubarraySumKOptimized(arr []int, k int) int {
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

// Legacy function for backward compatibility
func CountSubarraySumK(arr []int, k int) int {
	return CountSubarraySumKOptimized(arr, k)
}

// 4. Longest Subarray with Distinct Elements

// Brute Force Approach: Check all possible subarrays
// Time: O(N^2), Space: O(N)
func LongestDistinctSubarrayBruteForce(arr []int) int {
	maxLen := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		seen := make(map[int]bool)
		for j := i; j < n; j++ {
			if seen[arr[j]] {
				break
			}
			seen[arr[j]] = true
			if j-i+1 > maxLen {
				maxLen = j - i + 1
			}
		}
	}
	return maxLen
}

// Optimized Approach: Sliding window with hash map
// Time: O(N), Space: O(N)
func LongestDistinctSubarrayOptimized(arr []int) int {
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

// Legacy function for backward compatibility
func LongestDistinctSubarray(arr []int) int {
	return LongestDistinctSubarrayOptimized(arr)
}

// 5. Subarray with Given XOR

// Brute Force Approach: Check all possible subarrays
// Time: O(N^2), Space: O(1)
func SubarrayXorBBruteForce(arr []int, b int) int {
	count := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		currentXor := 0
		for j := i; j < n; j++ {
			currentXor ^= arr[j]
			if currentXor == b {
				count++
			}
		}
	}
	return count
}

// Optimized Approach: Hash map with prefix XOR frequency
// Time: O(N), Space: O(N)
func SubarrayXorBOptimized(arr []int, b int) int {
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

// Legacy function for backward compatibility
func SubarrayXorB(arr []int, b int) int {
	return SubarrayXorBOptimized(arr, b)
}

// 6. Count Subarrays with Equal Odd and Even Numbers

// Brute Force Approach: Check all possible subarrays
// Time: O(N^2), Space: O(1)
func CountOddEvenSubarraysBruteForce(arr []int) int {
	count := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		evenCount, oddCount := 0, 0
		for j := i; j < n; j++ {
			if arr[j]%2 == 0 {
				evenCount++
			} else {
				oddCount++
			}
			if evenCount == oddCount {
				count++
			}
		}
	}
	return count
}

// Optimized Approach: Hash map with prefix sum frequency (transform: Odd->1, Even->-1)
// Time: O(N), Space: O(N)
func CountOddEvenSubarraysOptimized(arr []int) int {
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

// Legacy function for backward compatibility
func CountOddEvenSubarrays(arr []int) int {
	return CountOddEvenSubarraysOptimized(arr)
}
