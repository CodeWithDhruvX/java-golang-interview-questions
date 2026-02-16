# 6. Sorting-Based Interview Problems

## 1. Merge Two Sorted Arrays

### Algorithm
1. Initialize `p1` for `nums1` (length `m`), `p2` for `nums2` (length `n`), `p` for the merged end (`m+n-1`).
2. Iterate while `p1 >= 0` and `p2 >= 0`:
   - If `nums1[p1] > nums2[p2]`:
     - `nums1[p] = nums1[p1]`, `p1--`.
   - Else:
     - `nums1[p] = nums2[p2]`, `p2--`.
   - `p--`.
3. If `p2 >= 0`, copy remaining `nums2` elements to `nums1`.

### Code
```go
// MergeSortedArrays merges two sorted arrays into the firt array
// Time Complexity: O(M+N)
// Space Complexity: O(1)
func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1
	
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
```

---

## 2. Find the Median of Two Sorted Arrays

### Algorithm
1. Ensure `nums1` is the smaller array.
2. Use Binary Search on `nums1` (partition method).
3. Find partition index `i` for `nums1`, and `j` for `nums2`.
4. Ensure elements on left partition <= elements on right partition.
5. If valid, calculate median based on max of lefts and min of rights.

### Code
```go
import "math"

// FindMedianSortedArrays finds the median of two sorted arrays
// Time Complexity: O(log(min(M, N)))
// Space Complexity: O(1)
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}
	
	m, n := len(nums1), len(nums2)
	low, high := 0, m
	
	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (m + n + 1) / 2 - partitionX
		
		maxLeftX := math.MinInt
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}
		
		minRightX := math.MaxInt
		if partitionX < m {
			minRightX = nums1[partitionX]
		}
		
		maxLeftY := math.MinInt
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}
		
		minRightY := math.MaxInt
		if partitionY < n {
			minRightY = nums2[partitionY]
		}
		
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (m+n)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2
			}
			return float64(max(maxLeftX, maxLeftY))
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
```

---

## 3. Sort Array by Frequency

### Algorithm
1. Count frequency of each element using a Map.
2. Create a custom sort function.
3. If frequencies are different, sort by frequency (ascending/descending).
4. If frequencies are same, sort by value (ascending/descending as per requirement).

### Code
```go
import "sort"

// FrequencySort sorts array elements by frequency
// Time Complexity: O(N log N)
// Space Complexity: O(N)
func FrequencySort(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	
	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[i] > nums[j] // If freq same, sort by value descending
		}
		return freq[nums[i]] < freq[nums[j]] // Sort by freq ascending
	})
	return nums
}
```

---

## 4. Minimum Number of Swaps to Sort the Array

### Algorithm
1. Create pair of `(value, index)` for each element.
2. Sort the pairs based on value.
3. Use a `visited` array to keep track of visited elements.
4. Iterate through the sorted pairs.
5. If element is visited or already in correct position `(index == i)`, continue.
6. Else, trace the cycle formed by this element and count cycle length `k`.
7. Add `k-1` to total swaps.

### Code
```go
import "sort"

// MinSwaps functions calculates minimum swaps to sort the array
// Time Complexity: O(N log N)
// Space Complexity: O(N)
func MinSwaps(nums []int) int {
	n := len(nums)
	type pair struct {
		val, idx int
	}
	arrPos := make([]pair, n)
	for i, v := range nums {
		arrPos[i] = pair{v, i}
	}
	
	sort.Slice(arrPos, func(i, j int) bool {
		return arrPos[i].val < arrPos[j].val
	})
	
	visited := make([]bool, n)
	swaps := 0
	
	for i := 0; i < n; i++ {
		if visited[i] || arrPos[i].idx == i {
			continue
		}
		
		cycleSize := 0
		j := i
		for !visited[j] {
			visited[j] = true
			j = arrPos[j].idx
			cycleSize++
		}
		if cycleSize > 0 {
			swaps += (cycleSize - 1)
		}
	}
	return swaps
}
```

---

## 5. Find Inversion Count

### Algorithm (Merge Sort based)
1. Split array into two halves.
2. Recursively count inversions in left and right halves.
3. Count inversions where `arr[i] > arr[j]` during merge step.
4. Sum all counts.

### Code
```go
// CountInversions counts pairs where i < j and arr[i] > arr[j]
// Time Complexity: O(N log N)
// Space Complexity: O(N)
func CountInversions(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	
	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)
	copy(left, arr[:mid])
	copy(right, arr[mid:])
	
	count := CountInversions(left) + CountInversions(right)
	
	// Merge and count
	i, j, k := 0, 0, 0
	sortedArr := make([]int, len(arr))
	
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			sortedArr[k] = left[i]
			i++
		} else {
			sortedArr[k] = right[j]
			j++
			count += len(left) - i
		}
		k++
	}
	
	for i < len(left) {
		sortedArr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		sortedArr[k] = right[j]
		j++
		k++
	}
	copy(arr, sortedArr)
	
	return count
}
```

---

## 6. Chocolate Distribution Problem

### Algorithm
1. Sort the array of chocolate packets.
2. Initialize `minDiff = INF`.
3. Iterate from `i = 0` to `n - M` (where M is number of students).
4. `diff = arr[i+M-1] - arr[i]`.
5. Update `minDiff = min(minDiff, diff)`.

### Code
```go
import "sort"
import "math"

// MinDiffChocolate finds min difference between max and min chocolates given to M students
// Time Complexity: O(N log N)
// Space Complexity: O(1)
func MinDiffChocolate(arr []int, m int) int {
	if m == 0 || len(arr) == 0 {
		return 0
	}
	if len(arr) < m {
		return -1
	}
	
	sort.Ints(arr)
	minDiff := math.MaxInt
	
	for i := 0; i+m-1 < len(arr); i++ {
		diff := arr[i+m-1] - arr[i]
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}
```

---

## 7. Arrange Array to Form Largest Number

### Algorithm
1. Convert integers to strings.
2. Sort strings with custom comparator:
   - If `a + b > b + a`, then `a` comes before `b`.
3. Join sorted strings.
4. Handle edge case: if result starts with "0", return "0".

### Code
```go
import (
	"fmt"
	"sort"
	"strings"
)

// LargestNumber forms the largest number from array elements
// Time Complexity: O(N log N)
// Space Complexity: O(N)
func LargestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, v := range nums {
		strs[i] = fmt.Sprintf("%d", v)
	}
	
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	
	if strs[0] == "0" {
		return "0"
	}
	
	return strings.Join(strs, "")
}
```
