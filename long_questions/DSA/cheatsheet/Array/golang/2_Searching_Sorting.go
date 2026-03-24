package main

// 1. Linear Search

// Brute Force Approach: Same as linear search (no better brute force exists)
// Time: O(N), Space: O(1)
func LinearSearchBruteForce(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

// Optimized Approach: Linear search is already optimal for unsorted arrays
// Time: O(N), Space: O(1)
func LinearSearchOptimized(arr []int, target int) int {
	return LinearSearchBruteForce(arr, target)
}

// Legacy function for backward compatibility
func LinearSearch(arr []int, target int) int {
	return LinearSearchOptimized(arr, target)
}

// 2. Binary Search (Iterative)

// Brute Force Approach: Linear search on sorted array
// Time: O(N), Space: O(1)
func BinarySearchBruteForce(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

// Optimized Approach: Binary search
// Time: O(log N), Space: O(1)
func BinarySearchOptimized(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Legacy function for backward compatibility
func BinarySearch(arr []int, target int) int {
	return BinarySearchOptimized(arr, target)
}

// 3. Find First and Last Occurrence

// Brute Force Approach: Linear scan to find first and last
// Time: O(N), Space: O(1)
func FirstLastOccurrenceBruteForce(arr []int, target int) (int, int) {
	first, last := -1, -1
	for i, val := range arr {
		if val == target {
			if first == -1 {
				first = i
			}
			last = i
		}
	}
	return first, last
}

// Optimized Approach: Binary search for both bounds
// Time: O(log N), Space: O(1)
func FirstLastOccurrenceOptimized(arr []int, target int) (int, int) {
	first := findBound(arr, target, true)
	if first == -1 {
		return -1, -1
	}
	last := findBound(arr, target, false)
	return first, last
}

// Legacy function for backward compatibility
func FirstLastOccurrence(arr []int, target int) (int, int) {
	return FirstLastOccurrenceOptimized(arr, target)
}

func findBound(arr []int, target int, isFirst bool) int {
	low, high := 0, len(arr)-1
	res := -1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			res = mid
			if isFirst {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}

// 4. Count Occurrences of a Number in Sorted Array

// Brute Force Approach: Linear scan counting
// Time: O(N), Space: O(1)
func CountOccurrencesBruteForce(arr []int, target int) int {
	count := 0
	for _, val := range arr {
		if val == target {
			count++
		}
	}
	return count
}

// Optimized Approach: Binary search for first and last occurrence
// Time: O(log N), Space: O(1)
func CountOccurrencesOptimized(arr []int, target int) int {
	first, last := FirstLastOccurrenceOptimized(arr, target)
	if first == -1 {
		return 0
	}
	return last - first + 1
}

// Legacy function for backward compatibility
func CountOccurrences(arr []int, target int) int {
	return CountOccurrencesOptimized(arr, target)
}

// 5. Find Missing Number (1..N)

// Brute Force Approach: Use hash set to track seen numbers
// Time: O(N), Space: O(N)
func FindMissingNumberBruteForce(arr []int) int {
	n := len(arr)
	seen := make(map[int]bool)
	for _, val := range arr {
		seen[val] = true
	}
	for i := 1; i <= n+1; i++ {
		if !seen[i] {
			return i
		}
	}
	return -1 // Should never reach here
}

// Optimized Approach: XOR method
// Time: O(N), Space: O(1)
func FindMissingNumberOptimized(arr []int) int {
	n := len(arr)
	xorAll := 0
	xorArr := 0

	// xorAll should include 0 to N.
	// We iterate 0 to N.
	for i := 0; i <= n; i++ {
		xorAll ^= i
	}
	// xorArr includes all elements in array.
	for _, val := range arr {
		xorArr ^= val
	}
	return xorAll ^ xorArr
}

// Legacy function for backward compatibility
func FindMissingNumber(arr []int) int {
	return FindMissingNumberOptimized(arr)
}

// 6. Find Element That Appears Only Once

// Brute Force Approach: Use hash map to count frequencies
// Time: O(N), Space: O(N)
func FindSingleNumberBruteForce(arr []int) int {
	freq := make(map[int]int)
	for _, val := range arr {
		freq[val]++
	}
	for val, count := range freq {
		if count == 1 {
			return val
		}
	}
	return -1 // Should never reach here for valid input
}

// Optimized Approach: XOR method
// Time: O(N), Space: O(1)
func FindSingleNumberOptimized(arr []int) int {
	res := 0
	for _, num := range arr {
		res ^= num
	}
	return res
}

// Legacy function for backward compatibility
func FindSingleNumber(arr []int) int {
	return FindSingleNumberOptimized(arr)
}

// 7. Find Peak Element

// Brute Force Approach: Linear scan to find any peak
// Time: O(N), Space: O(1)
func FindPeakElementBruteForce(arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	// Check first element
	if n == 1 || arr[0] >= arr[1] {
		return 0
	}
	// Check middle elements
	for i := 1; i < n-1; i++ {
		if arr[i] >= arr[i-1] && arr[i] >= arr[i+1] {
			return i
		}
	}
	// Check last element
	if arr[n-1] >= arr[n-2] {
		return n - 1
	}
	return -1 // Should never reach here for valid input
}

// Optimized Approach: Binary search
// Time: O(log N), Space: O(1)
func FindPeakElementOptimized(arr []int) int {
	low, high := 0, len(arr)-1
	for low < high {
		mid := low + (high-low)/2
		if arr[mid] < arr[mid+1] {
			// We are on an upward slope, peak is to the right
			low = mid + 1
		} else {
			// We are on a downward slope, peak is here or to the left
			high = mid
		}
	}
	return low
}

// Legacy function for backward compatibility
func FindPeakElement(arr []int) int {
	return FindPeakElementOptimized(arr)
}
