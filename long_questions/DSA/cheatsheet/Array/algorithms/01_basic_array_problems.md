# 1. Basic Array Problems (Warm-up)

## 1. Find Largest and Smallest Element

### Algorithm
1. Initialize `minVal` and `maxVal` with the first element of the array.
2. Iterate through the array starting from the second element.
3. For each element `x`:
   - If `x > maxVal`, update `maxVal = x`.
   - If `x < minVal`, update `minVal = x`.
4. Return `minVal` and `maxVal`.

### Code
```go
package main

import (
	"fmt"
	"math"
)

// FindMinMax returns the minimum and maximum elements in an array
// Time Complexity: O(N)
// Space Complexity: O(1)
func FindMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0 // Handle empty array case appropriately
	}
	
	minVal := arr[0]
	maxVal := arr[0]
	
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
	}
	return minVal, maxVal
}
```

---

## 2. Reverse an Array In-Place

### Algorithm
1. Initialize two pointers: `left = 0` and `right = n-1`.
2. Loop while `left < right`:
   - Swap `arr[left]` and `arr[right]`.
   - Increment `left`.
   - Decrement `right`.
3. The array is reversed in place.

### Code
```go
// ReverseArray reverses the array in-place
// Time Complexity: O(N)
// Space Complexity: O(1)
func ReverseArray(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
```

---

## 3. Find Second Largest Element

### Algorithm
1. Initialize `largest` and `secondLargest` to minimum integer value (`math.MinInt`).
2. Iterate through the array:
   - If `current > largest`:
     - Update `secondLargest = largest`.
     - Update `largest = current`.
   - Else if `current > secondLargest` AND `current != largest`:
     - Update `secondLargest = current`.
3. If `secondLargest` is still `math.MinInt`, it means no second largest exists (all elements are same).
4. Return `secondLargest`.

### Code
```go
// FindSecondLargest returns the second largest element in the array
// Time Complexity: O(N)
// Space Complexity: O(1)
func FindSecondLargest(arr []int) int {
	if len(arr) < 2 {
		return -1 // Not enough elements
	}
	
	largest := math.MinInt
	secondLargest := math.MinInt
	
	for _, val := range arr {
		if val > largest {
			secondLargest = largest
			largest = val
		} else if val > secondLargest && val != largest {
			secondLargest = val
		}
	}
	
	if secondLargest == math.MinInt {
		return -1 // No second largest distinct element found
	}
	
	return secondLargest
}
```

---

## 4. Check if Array is Sorted

### Algorithm
1. Iterate from index `0` to `n-2`.
2. Check if `arr[i] > arr[i+1]`.
3. If true, the array is not sorted; return `false`.
4. If the loop completes without returning `false`, return `true`.

### Code
```go
// IsSorted checks if the array is sorted in non-decreasing order
// Time Complexity: O(N)
// Space Complexity: O(1)
func IsSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
```

---

## 5. Count Even and Odd Elements

### Algorithm
1. Initialize `evenCount = 0` and `oddCount = 0`.
2. Iterate through each element `x` in the array.
3. If `x % 2 == 0`, increment `evenCount`.
4. Else, increment `oddCount`.
5. Return both counts.

### Code
```go
// CountEvenOdd returns the count of even and odd numbers
// Time Complexity: O(N)
// Space Complexity: O(1)
func CountEvenOdd(arr []int) (int, int) {
	evenCount := 0
	oddCount := 0
	
	for _, val := range arr {
		if val%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	return evenCount, oddCount
}
```

---

## 6. Remove Duplicate Elements from a Sorted Array

### Algorithm
1. If the array is empty, return 0.
2. Initialize `i = 0` (slow pointer).
3. Iterate with `j` (fast pointer) from `1` to `n-1`.
4. If `arr[j] != arr[i]`:
   - Increment `i`.
   - Update `arr[i] = arr[j]`.
5. Return `i + 1` (the length of unique elements).

### Code
```go
// RemoveDuplicates removes duplicates in-place from a sorted array
// Returns the new length of the array
// Time Complexity: O(N)
// Space Complexity: O(1)
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
```

---

## 7. Left Rotate an Array by 1 Position

### Algorithm
1. Store the first element in a variable `temp = arr[0]`.
2. Iterate from `i = 1` to `n-1`.
3. Shift elements left: `arr[i-1] = arr[i]`.
4. Place `temp` at the last position: `arr[n-1] = temp`.

### Code
```go
// RotateLeftByOne rotates the array left by 1 position
// Time Complexity: O(N)
// Space Complexity: O(1)
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
```

---

## 8. Left Rotate an Array by K Positions

### Algorithm
1. `K = K % N` to handle cases where `K > N`.
2. Reverse the first `K` elements (`0` to `K-1`).
3. Reverse the remaining `N-K` elements (`K` to `N-1`).
4. Reverse the entire array (`0` to `N-1`).

### Code
```go
// RotateLeftByK rotates the array left by K positions
// Time Complexity: O(N)
// Space Complexity: O(1)
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
```

---

## 9. Find the Sum of All Elements

### Algorithm
1. Initialize `sum = 0`.
2. Iterate through each element `x` of the array.
3. Add `x` to `sum`.
4. Return `sum`.

### Code
```go
// ArraySum calculates the sum of all elements
// Time Complexity: O(N)
// Space Complexity: O(1)
func ArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
```

---

## 10. Find the Frequency of Each Element

### Algorithm
1. Create a map `freq` to store counts.
2. Iterate through the array.
3. For each element `x`, increment `freq[x]`.
4. Return the map.

### Code
```go
// CountFrequency counts the frequency of each element
// Time Complexity: O(N)
// Space Complexity: O(N)
func CountFrequency(arr []int) map[int]int {
	freq := make(map[int]int)
	for _, val := range arr {
		freq[val]++
	}
	return freq
}
```
