```go
package main

// 1. Move All Zeros to End
// Time: O(N), Space: O(1)
func MoveZeros(arr []int) {
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

// 2. Sort Array of 0s, 1s, and 2s (Dutch National Flag)
// Time: O(N), Space: O(1)
func SortColors(arr []int) {
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

// 3. Find Subarray with Given Sum (positive numbers)
// Time: O(N), Space: O(1)
func SubarraySum(arr []int, target int) []int {
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

// 4. Find Maximum Sum Subarray (Kadane's Algorithm)
// Time: O(N), Space: O(1)
func MaxSumSubarray(arr []int) int {
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

// 5. Find Longest Subarray with Sum = K (Positives & Negatives)
// Time: O(N), Space: O(N)
func LongestSubarraySumK(arr []int, k int) int {
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

// 6. Smallest Subarray with Sum > X
// Time: O(N), Space: O(1)
func SmallestSubarraySumX(arr []int, x int) int {
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

// 7. Longest Subarray with Equal 0s and 1s
// Time: O(N), Space: O(N)
func LongestEqual01(arr []int) int {
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
```
