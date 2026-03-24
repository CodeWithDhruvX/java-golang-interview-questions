package main

import (
	"fmt"
	"math"
)

// 1. Find Largest and Smallest Element

// Brute Force Approach: Sort the array and take first/last elements
// Time: O(N log N), Space: O(1) if using in-place sort
func FindMinMaxBruteForce(arr []int) (int, int, error) {
	if len(arr) == 0 {
		return 0, 0, fmt.Errorf("empty array")
	}
	// Create a copy to avoid modifying original
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	
	// Simple bubble sort for demonstration
	for i := 0; i < len(sortedArr)-1; i++ {
		for j := 0; j < len(sortedArr)-i-1; j++ {
			if sortedArr[j] > sortedArr[j+1] {
				sortedArr[j], sortedArr[j+1] = sortedArr[j+1], sortedArr[j]
			}
		}
	}
	return sortedArr[0], sortedArr[len(sortedArr)-1], nil
}

// Optimized Approach: Single pass
// Time: O(N), Space: O(1)
func FindMinMaxOptimized(arr []int) (int, int, error) {
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

// Legacy function for backward compatibility
func FindMinMax(arr []int) (int, int, error) {
	return FindMinMaxOptimized(arr)
}

// 2. Reverse an Array In-Place

// Brute Force Approach: Create new array and copy in reverse order
// Time: O(N), Space: O(N)
func ReverseArrayBruteForce(arr []int) []int {
	n := len(arr)
	reversed := make([]int, n)
	for i := 0; i < n; i++ {
		reversed[i] = arr[n-1-i]
	}
	return reversed
}

// Optimized Approach: Two-pointer technique
// Time: O(N), Space: O(1)
func ReverseArrayOptimized(arr []int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

// Legacy function for backward compatibility
func ReverseArray(arr []int) {
	ReverseArrayOptimized(arr)
}

// 3. Find Second Largest Element

// Brute Force Approach: Sort array and take second last element
// Time: O(N log N), Space: O(1)
func SecondLargestBruteForce(arr []int) (int, error) {
	if len(arr) < 2 {
		return 0, fmt.Errorf("array must have at least 2 elements")
	}
	// Create a copy to avoid modifying original
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	
	// Simple bubble sort
	for i := 0; i < len(sortedArr)-1; i++ {
		for j := 0; j < len(sortedArr)-i-1; j++ {
			if sortedArr[j] > sortedArr[j+1] {
				sortedArr[j], sortedArr[j+1] = sortedArr[j+1], sortedArr[j]
			}
		}
	}
	
	// Handle duplicates
	for i := len(sortedArr) - 2; i >= 0; i-- {
		if sortedArr[i] != sortedArr[len(sortedArr)-1] {
			return sortedArr[i], nil
		}
	}
	return 0, fmt.Errorf("no second largest element found")
}

// Optimized Approach: Single pass with two variables
// Time: O(N), Space: O(1)
func SecondLargestOptimized(arr []int) (int, error) {
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

// Legacy function for backward compatibility
func SecondLargest(arr []int) (int, error) {
	return SecondLargestOptimized(arr)
}

// 4. Check if Array is Sorted

// Brute Force Approach: Sort and compare with original
// Time: O(N log N), Space: O(N)
func IsSortedBruteForce(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	// Create a copy and sort it
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	
	// Simple bubble sort
	for i := 0; i < len(sortedArr)-1; i++ {
		for j := 0; j < len(sortedArr)-i-1; j++ {
			if sortedArr[j] > sortedArr[j+1] {
				sortedArr[j], sortedArr[j+1] = sortedArr[j+1], sortedArr[j]
			}
		}
	}
	
	// Compare with original
	for i := 0; i < len(arr); i++ {
		if arr[i] != sortedArr[i] {
			return false
		}
	}
	return true
}

// Optimized Approach: Single pass checking adjacent elements
// Time: O(N), Space: O(1)
func IsSortedOptimized(arr []int) bool {
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

// Legacy function for backward compatibility
func IsSorted(arr []int) bool {
	return IsSortedOptimized(arr)
}

// 5. Count Even and Odd Elements

// Brute Force Approach: Use additional storage (not really brute force, but for demonstration)
// Time: O(N), Space: O(N)
func CountEvenOddBruteForce(arr []int) (int, int) {
	even := []int{}
	odd := []int{}
	for _, val := range arr {
		if val%2 == 0 {
			even = append(even, val)
		} else {
			odd = append(odd, val)
		}
	}
	return len(even), len(odd)
}

// Optimized Approach: Simple counters
// Time: O(N), Space: O(1)
func CountEvenOddOptimized(arr []int) (int, int) {
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

// Legacy function for backward compatibility
func CountEvenOdd(arr []int) (int, int) {
	return CountEvenOddOptimized(arr)
}

// 6. Remove Duplicates from Sorted Array

// Brute Force Approach: Use extra array to store unique elements
// Time: O(N), Space: O(N)
func RemoveDuplicatesBruteForce(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	unique := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			unique = append(unique, arr[i])
		}
	}
	return unique
}

// Optimized Approach: In-place modification using two pointers
// Time: O(N), Space: O(1)
func RemoveDuplicatesOptimized(arr []int) int {
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

// Legacy function for backward compatibility
func RemoveDuplicates(arr []int) int {
	return RemoveDuplicatesOptimized(arr)
}

// 7. Left Rotate Array by 1 Position

// Brute Force Approach: Create new array and copy elements
// Time: O(N), Space: O(N)
func RotateLeftByOneBruteForce(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	rotated := make([]int, len(arr))
	for i := 1; i < len(arr); i++ {
		rotated[i-1] = arr[i]
	}
	rotated[len(arr)-1] = arr[0]
	return rotated
}

// Optimized Approach: In-place rotation
// Time: O(N), Space: O(1)
func RotateLeftByOneOptimized(arr []int) {
	if len(arr) == 0 {
		return
	}
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp
}

// Legacy function for backward compatibility
func RotateLeftByOne(arr []int) {
	RotateLeftByOneOptimized(arr)
}

// 8. Left Rotate Array by K Positions

// Brute Force Approach: Rotate K times by 1 position
// Time: O(N*K), Space: O(1)
func RotateLeftByKBruteForce(arr []int, k int) {
	n := len(arr)
	if n == 0 {
		return
	}
	k = k % n
	for i := 0; i < k; i++ {
		RotateLeftByOneOptimized(arr)
	}
}

// Optimized Approach: Reversal algorithm
// Time: O(N), Space: O(1)
func RotateLeftByKOptimized(arr []int, k int) {
	n := len(arr)
	if n == 0 {
		return
	}
	k = k % n
	reverse(arr, 0, k-1)
	reverse(arr, k, n-1)
	reverse(arr, 0, n-1)
}

// Legacy function for backward compatibility
func RotateLeftByK(arr []int, k int) {
	RotateLeftByKOptimized(arr)
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

// 9. Find Sum of All Elements

// Brute Force Approach: Use recursion (stack overhead)
// Time: O(N), Space: O(N) due to recursion stack
func ArraySumBruteForce(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + ArraySumBruteForce(arr[1:])
}

// Optimized Approach: Simple iteration
// Time: O(N), Space: O(1)
func ArraySumOptimized(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

// Legacy function for backward compatibility
func ArraySum(arr []int) int {
	return ArraySumOptimized(arr)
}

// 10. Find Frequency of Each Element

// Brute Force Approach: For each element, count occurrences by scanning array
// Time: O(N^2), Space: O(N)
func FrequencyCountBruteForce(arr []int) map[int]int {
	freq := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		// Check if already counted
		if _, exists := freq[arr[i]]; !exists {
			count := 0
			for j := 0; j < len(arr); j++ {
				if arr[i] == arr[j] {
					count++
				}
			}
			freq[arr[i]] = count
		}
	}
	return freq
}

// Optimized Approach: Single pass with hash map
// Time: O(N), Space: O(N)
func FrequencyCountOptimized(arr []int) map[int]int {
	freq := make(map[int]int)
	for _, val := range arr {
		freq[val]++
	}
	return freq
}

// Legacy function for backward compatibility
func FrequencyCount(arr []int) map[int]int {
	return FrequencyCountOptimized(arr)
}
