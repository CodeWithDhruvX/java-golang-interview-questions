package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// 1. Merge Two Sorted Arrays
// Time: O(N+M), Space: O(N+M)
func MergeSortedArrays(arr1, arr2 []int) []int {
	n, m := len(arr1), len(arr2)
	res := make([]int, n+m)
	i, j, k := 0, 0, 0

	for i < n && j < m {
		if arr1[i] <= arr2[j] {
			res[k] = arr1[i]
			i++
		} else {
			res[k] = arr2[j]
			j++
		}
		k++
	}

	for i < n {
		res[k] = arr1[i]
		i++
		k++
	}
	for j < m {
		res[k] = arr2[j]
		j++
		k++
	}
	return res
}

// 2. Median of Two Sorted Arrays
// Time: O(log(min(N, M))), Space: O(1)
func MedianSortedArrays(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return MedianSortedArrays(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxLeftX := math.MinInt64
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}

		minRightX := math.MaxInt64
		if partitionX < x {
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt64
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}

		minRightY := math.MaxInt64
		if partitionY < y {
			minRightY = nums2[partitionY]
		}

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x+y)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}
	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 3. Sort Array by Frequency
// Time: O(N log N), Space: O(N)
func SortByFrequency(arr []int) []int {
	freq := make(map[int]int)
	for _, val := range arr {
		freq[val]++
	}

	// We copy arr to a new slice to sort
	res := make([]int, len(arr))
	copy(res, arr)

	sort.Slice(res, func(i, j int) bool {
		if freq[res[i]] == freq[res[j]] {
			// If frequency matches, sort by value (ascending) - Optional rule
			return res[i] < res[j]
		}
		// Sort by frequency descending
		return freq[res[i]] > freq[res[j]]
	})

	return res
}

// 4. Minimum Swaps to Sort Array
// Time: O(N log N), Space: O(N)
func MinSwapsToSort(arr []int) int {
	n := len(arr)
	type pair struct {
		val, idx int
	}
	pairs := make([]pair, n)
	for i, val := range arr {
		pairs[i] = pair{val, i}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val < pairs[j].val
	})

	visited := make([]bool, n)
	swaps := 0

	for i := 0; i < n; i++ {
		if visited[i] || pairs[i].idx == i {
			continue
		}

		cycleLen := 0
		j := i
		for !visited[j] {
			visited[j] = true
			j = pairs[j].idx
			cycleLen++
		}
		if cycleLen > 0 {
			swaps += cycleLen - 1
		}
	}
	return swaps
}

// 5. Count Inversions in Array
// Time: O(N log N), Space: O(N)
func CountInversions(arr []int) int {
	return mergeSortAndCount(arr, 0, len(arr)-1)
}

func mergeSortAndCount(arr []int, l, r int) int {
	count := 0
	if l < r {
		m := (l + r) / 2
		count += mergeSortAndCount(arr, l, m)
		count += mergeSortAndCount(arr, m+1, r)
		count += mergeAndCount(arr, l, m, r)
	}
	return count
}

func mergeAndCount(arr []int, l, m, r int) int {
	left := make([]int, m-l+1)
	right := make([]int, r-m)

	for i := 0; i < len(left); i++ {
		left[i] = arr[l+i]
	}
	for i := 0; i < len(right); i++ {
		right[i] = arr[m+1+i]
	}

	i, j, k := 0, 0, l
	swaps := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
			// Key Insight: left[i] and all subsequent elements in left are > right[j]
			swaps += (len(left) - i)
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
	return swaps
}

// 6. Chocolate Distribution Problem
// Time: O(N log N), Space: O(1)
func ChocolateDistribution(arr []int, m int) int {
	n := len(arr)
	if m == 0 || n == 0 {
		return 0
	}
	if n < m {
		return -1
	}
	sort.Ints(arr)
	minDiff := math.MaxInt64

	for i := 0; i+m-1 < n; i++ {
		diff := arr[i+m-1] - arr[i]
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}

// 7. Largest Number from Array
// Time: O(N log N), Space: O(N)
func LargestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, v := range nums {
		strs[i] = strconv.Itoa(v)
	}

	sort.Slice(strs, func(i, j int) bool {
		// Example: "3" vs "30" -> "330" > "303" ? True
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	if len(strs) > 0 && strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}
