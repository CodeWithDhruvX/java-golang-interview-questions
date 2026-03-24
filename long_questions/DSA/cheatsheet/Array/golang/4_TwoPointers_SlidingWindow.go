package main

// 1. Move All Zeros to End

// Brute Force Approach: Create new array, copy non-zeros then zeros
// Time: O(N), Space: O(N)
func MoveZerosBruteForce(arr []int) []int {
	n := len(arr)
	result := make([]int, 0, n)
	zeroCount := 0
	
	// Copy non-zero elements
	for _, val := range arr {
		if val != 0 {
			result = append(result, val)
		} else {
			zeroCount++
		}
	}
	
	// Append zeros
	for i := 0; i < zeroCount; i++ {
		result = append(result, 0)
	}
	
	return result
}

// Optimized Approach: Two-pointer technique in-place
// Time: O(N), Space: O(1)
func MoveZerosOptimized(arr []int) {
	write := 0
	// 1. Move non-zeros to front
	for _, val := range arr {
		if val != 0 {
			arr[write] = val
			write++
		}
	}
	// 2. Fill remaining with zeros
	for i := write; i < len(arr); i++ {
		arr[i] = 0
	}
}

// Legacy function for backward compatibility
func MoveZeros(arr []int) {
	MoveZerosOptimized(arr)
}

// 2. Sort Array of 0s, 1s, and 2s (Dutch National Flag)

// Brute Force Approach: Count and fill
// Time: O(N), Space: O(1)
func SortColorsBruteForce(arr []int) {
	count0, count1, count2 := 0, 0, 0
	
	// Count occurrences
	for _, val := range arr {
		switch val {
		case 0:
			count0++
		case 1:
			count1++
		case 2:
			count2++
		}
	}
	
	// Fill array
	index := 0
	for i := 0; i < count0; i++ {
		arr[index] = 0
		index++
	}
	for i := 0; i < count1; i++ {
		arr[index] = 1
		index++
	}
	for i := 0; i < count2; i++ {
		arr[index] = 2
		index++
	}
}

// Optimized Approach: Dutch National Flag algorithm
// Time: O(N), Space: O(1)
func SortColorsOptimized(arr []int) {
	low, mid, high := 0, 0, len(arr)-1
	for mid <= high {
		switch arr[mid] {
		case 0:
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			arr[mid], arr[high] = arr[high], arr[mid]
			high--
		}
	}
}

// Legacy function for backward compatibility
func SortColors(arr []int) {
	SortColorsOptimized(arr)
}

// 3. Find Subarray with Given Sum (positive numbers)

// Brute Force Approach: Check all possible subarrays
// Time: O(N^2), Space: O(1)
func SubarraySumBruteForce(arr []int, target int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		currentSum := 0
		for j := i; j < n; j++ {
			currentSum += arr[j]
			if currentSum == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Optimized Approach: Sliding window
// Time: O(N), Space: O(1)
func SubarraySumOptimized(arr []int, target int) []int {
	left, currentSum := 0, 0
	for right, val := range arr {
		currentSum += val
		// Shrink window while sum is too large
		for currentSum > target && left <= right {
			currentSum -= arr[left]
			left++
		}
		if currentSum == target {
			return []int{left, right}
		}
	}
	return nil
}

// Legacy function for backward compatibility
func SubarraySum(arr []int, target int) []int {
	return SubarraySumOptimized(arr, target)
}

// 4. Find Maximum Sum Subarray (Kadane's Algorithm)

// Brute Force Approach: Check all possible subarrays
// Time: O(N^2), Space: O(1)
func MaxSumSubarrayBruteForce(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxSum := arr[0]
	n := len(arr)
	for i := 0; i < n; i++ {
		currentSum := 0
		for j := i; j < n; j++ {
			currentSum += arr[j]
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}
	return maxSum
}

// Optimized Approach: Kadane's Algorithm
// Time: O(N), Space: O(1)
func MaxSumSubarrayOptimized(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxSum := arr[0]
	currentSum := 0
	for _, val := range arr {
		currentSum += val
		if currentSum > maxSum {
			maxSum = currentSum
		}
		if currentSum < 0 {
			currentSum = 0
		}
	}
	return maxSum
}

// Legacy function for backward compatibility
func MaxSumSubarray(arr []int) int {
	return MaxSumSubarrayOptimized(arr)
}

// 5. Find Longest Subarray with Sum = K (Positives & Negatives)

// Brute Force Approach: Check all possible subarrays
// Time: O(N^2), Space: O(1)
func LongestSubarraySumKBruteForce(arr []int, k int) int {
	maxLen := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		currentSum := 0
		for j := i; j < n; j++ {
			currentSum += arr[j]
			if currentSum == k && j-i+1 > maxLen {
				maxLen = j - i + 1
			}
		}
	}
	return maxLen
}

// Optimized Approach: Hash map with prefix sums
// Time: O(N), Space: O(N)
func LongestSubarraySumKOptimized(arr []int, k int) int {
	m := make(map[int]int)
	m[0] = -1 // Base case: Sum 0 at index -1
	sum := 0
	maxLen := 0
	for i, val := range arr {
		sum += val
		if idx, ok := m[sum-k]; ok {
			if i-idx > maxLen {
				maxLen = i - idx
			}
		}
		// Only add sum to map if not present to keep earliest index (longest subarray)
		if _, ok := m[sum]; !ok {
			m[sum] = i
		}
	}
	return maxLen
}

// Legacy function for backward compatibility
func LongestSubarraySumK(arr []int, k int) int {
	return LongestSubarraySumKOptimized(arr, k)
}

// 6. Smallest Subarray with Sum > X

// Brute Force Approach: Check all possible subarrays
// Time: O(N^2), Space: O(1)
func SmallestSubarraySumXBruteForce(arr []int, x int) int {
	n := len(arr)
	minLen := n + 1
	for i := 0; i < n; i++ {
		currentSum := 0
		for j := i; j < n; j++ {
			currentSum += arr[j]
			if currentSum > x && j-i+1 < minLen {
				minLen = j - i + 1
				break // No need to check longer subarrays starting from i
			}
		}
	}
	if minLen > n {
		return 0
	}
	return minLen
}

// Optimized Approach: Sliding window
// Time: O(N), Space: O(1)
func SmallestSubarraySumXOptimized(arr []int, x int) int {
	n := len(arr)
	minLen := n + 1
	left, sum := 0, 0
	for right, val := range arr {
		sum += val
		for sum > x {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= arr[left]
			left++
		}
	}
	if minLen > n {
		return 0
	}
	return minLen
}

// Legacy function for backward compatibility
func SmallestSubarraySumX(arr []int, x int) int {
	return SmallestSubarraySumXOptimized(arr, x)
}

// 7. Longest Subarray with Equal 0s and 1s

// Brute Force Approach: Check all possible subarrays
// Time: O(N^2), Space: O(1)
func LongestEqual01BruteForce(arr []int) int {
	maxLen := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		count0, count1 := 0, 0
		for j := i; j < n; j++ {
			if arr[j] == 0 {
				count0++
			} else {
				count1++
			}
			if count0 == count1 && j-i+1 > maxLen {
				maxLen = j - i + 1
			}
		}
	}
	return maxLen
}

// Optimized Approach: Hash map with prefix sums (treat 0 as -1)
// Time: O(N), Space: O(N)
func LongestEqual01Optimized(arr []int) int {
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	maxLen := 0
	for i, val := range arr {
		// Treat 0 as -1
		if val == 0 {
			sum += -1
		} else {
			sum += 1
		}

		if idx, ok := m[sum]; ok {
			if i-idx > maxLen {
				maxLen = i - idx
			}
		} else {
			m[sum] = i
		}
	}
	return maxLen
}

// Legacy function for backward compatibility
func LongestEqual01(arr []int) int {
	return LongestEqual01Optimized(arr)
}
