```go
package main

import (
	"fmt"
	"math"
)

// 1. Find Largest and Smallest Element
// Time: O(N), Space: O(1)
func FindMinMax(arr []int) (int, int, error) {
	if len(arr) == 0 {
		return 0, 0, fmt.Errorf("empty array")
	}
	minVal, maxVal := arr[0], arr[0]
	for _, val := range arr[1:] {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
	}
	return minVal, maxVal, nil
}

// 2. Reverse an Array In-Place
// Time: O(N), Space: O(1)
func ReverseArray(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

// 3. Find Second Largest Element
// Time: O(N), Space: O(1)
func SecondLargest(arr []int) (int, error) {
	if len(arr) < 2 {
		return 0, fmt.Errorf("array must have at least 2 elements")
	}
	largest, second := math.MinInt, math.MinInt

	for _, val := range arr {
		if val > largest {
			second = largest
			largest = val
		} else if val > second && val != largest {
			second = val
		}
	}

	if second == math.MinInt {
		return 0, fmt.Errorf("no second largest element found")
	}
	return second, nil
}

// 4. Check if Array is Sorted
// Time: O(N), Space: O(1)
func IsSorted(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// 5. Count Even and Odd Elements
// Time: O(N), Space: O(1)
func CountEvenOdd(arr []int) (int, int) {
	even, odd := 0, 0
	for _, val := range arr {
		if val%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return even, odd
}

// 6. Remove Duplicates from Sorted Array
// Time: O(N), Space: O(1)
func RemoveDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[j] != arr[i] {
			i++
			arr[i] = arr[j]
		}
	}
	return i + 1
}

// 7. Left Rotate Array by 1 Position
// Time: O(N), Space: O(1)
func RotateLeftByOne(arr []int) {
	if len(arr) == 0 {
		return
	}
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp
}

// 8. Left Rotate Array by K Positions
// Time: O(N), Space: O(1)
func RotateLeftByK(arr []int, k int) {
	n := len(arr)
	if n == 0 {
		return
	}
	k = k % n
	reverse(arr, 0, k-1)
	reverse(arr, k, n-1)
	reverse(arr, 0, n-1)
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

// 9. Find Sum of All Elements
// Time: O(N), Space: O(1)
func ArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

// 10. Find Frequency of Each Element
// Time: O(N), Space: O(N)
func FrequencyCount(arr []int) map[int]int {
	freq := make(map[int]int)
	for _, val := range arr {
		freq[val]++
	}
	return freq
}
```
