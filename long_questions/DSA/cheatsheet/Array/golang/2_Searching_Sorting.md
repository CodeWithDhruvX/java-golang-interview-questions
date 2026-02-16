```go
package main

// 1. Linear Search
// Time: O(N), Space: O(1)
func LinearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

// 2. Binary Search (Iterative)
// Time: O(log N), Space: O(1)
func BinarySearch(arr []int, target int) int {
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

// 3. Find First and Last Occurrence
// Time: O(log N), Space: O(1)
func FirstLastOccurrence(arr []int, target int) (int, int) {
	first := findBound(arr, target, true)
	if first == -1 {
		return -1, -1
	}
	last := findBound(arr, target, false)
	return first, last
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
// Time: O(log N), Space: O(1)
func CountOccurrences(arr []int, target int) int {
	first, last := FirstLastOccurrence(arr, target)
	if first == -1 {
		return 0
	}
	return last - first + 1
}

// 5. Find Missing Number (1..N) using XOR
// Time: O(N), Space: O(1)
func FindMissingNumber(arr []int) int {
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

// 6. Find Element That Appears Only Once
// Time: O(N), Space: O(1)
func FindSingleNumber(arr []int) int {
	res := 0
	for _, num := range arr {
		res ^= num
	}
	return res
}

// 7. Find Peak Element
// Time: O(log N), Space: O(1)
func FindPeakElement(arr []int) int {
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
```
