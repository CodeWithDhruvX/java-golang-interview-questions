package main

import (
	"fmt"
	"math"
)

// 21. Find largest element in array
func findMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// 22. Find smallest element in array
func findMin(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

// 23. Find second largest element
func secondLargest(arr []int) int {
	first, second := math.MinInt32, math.MinInt32
	for _, num := range arr {
		if num > first {
			second = first
			first = num
		} else if num > second && num != first {
			second = num
		}
	}
	return second
}

// 24. Reverse an array
func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// 25. Rotate array left by K positions
func rotateLeft(arr []int, k int) []int {
	k = k % len(arr)
	reverse(arr, 0, k-1)
	reverse(arr, k, len(arr)-1)
	reverse(arr, 0, len(arr)-1)
	return arr
}

// 26. Rotate array right by K positions
func rotateRight(arr []int, k int) []int {
	k = k % len(arr)
	reverse(arr, 0, len(arr)-k-1)
	reverse(arr, len(arr)-k, len(arr)-1)
	reverse(arr, 0, len(arr)-1)
	return arr
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

// 27. Remove duplicates from array
func removeArrayDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	var result []int
	for _, num := range arr {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

// 28. Find frequency of elements
func arrayFrequency(arr []int) {
	freqMap := make(map[int]int)
	for _, num := range arr {
		freqMap[num]++
	}
	fmt.Println(freqMap)
}

// 29. Find missing number in array (1 to N)
func findMissing(arr []int, N int) int {
	expectedSum := N * (N + 1) / 2
	actualSum := 0
	for _, num := range arr {
		actualSum += num
	}
	return expectedSum - actualSum
}

// 30. Find duplicate number
func findDuplicate(arr []int) int {
	seen := make(map[int]bool)
	for _, num := range arr {
		if seen[num] {
			return num
		}
		seen[num] = true
	}
	return -1
}

// 31. Sort array without built-in methods (Bubble Sort)
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 32. Merge two arrays
func mergeArrays(arr1, arr2 []int) []int {
	return append(arr1, arr2...)
}

// 33. Find common elements in two arrays
func findCommon(arr1, arr2 []int) []int {
	set1 := make(map[int]bool)
	for _, num := range arr1 {
		set1[num] = true
	}
	var common []int
	for _, num := range arr2 {
		if set1[num] {
			common = append(common, num)
			delete(set1, num) // Avoid duplicates
		}
	}
	return common
}

// 34. Move all zeros to end
func moveZeros(arr []int) []int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}
	for count < len(arr) {
		arr[count] = 0
		count++
	}
	return arr
}

// 35. Find sum of array elements
func sumArray(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}
	return total
}

// 36. Find pair with given sum
func findPairWithSum(arr []int, target int) (int, int) {
	seen := make(map[int]bool)
	for _, num := range arr {
		complement := target - num
		if seen[complement] {
			return complement, num
		}
		seen[num] = true
	}
	return -1, -1 // Not found
}

// 37. Find max & min in single loop
func findMaxMin(arr []int) (int, int) {
	max, min := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min
}

// 38. Print array in reverse order
func printReverse(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

// 39. Check array is sorted or not
func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// 40. Count even & odd numbers
func countEvenOdd(arr []int) {
	even, odd := 0, 0
	for _, num := range arr {
		if num%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Printf("Even: %d, Odd: %d\n", even, odd)
}

func main() {
	fmt.Println("21. Find Max:")
	fmt.Println("[1, 5, 3, 9, 2] ->", findMax([]int{1, 5, 3, 9, 2}))

	fmt.Println("\n22. Find Min:")
	fmt.Println("[1, 5, 3, 9, 2] ->", findMin([]int{1, 5, 3, 9, 2}))

	fmt.Println("\n23. Second Largest:")
	fmt.Println("[12, 35, 1, 10, 34, 1] ->", secondLargest([]int{12, 35, 1, 10, 34, 1}))

	fmt.Println("\n24. Reverse Array:")
	fmt.Println("[1, 2, 3, 4] ->", reverseArray([]int{1, 2, 3, 4}))

	fmt.Println("\n25. Rotate Left:")
	fmt.Println("[1, 2, 3, 4, 5], k=2 ->", rotateLeft([]int{1, 2, 3, 4, 5}, 2))

	fmt.Println("\n26. Rotate Right:")
	fmt.Println("[1, 2, 3, 4, 5], k=2 ->", rotateRight([]int{1, 2, 3, 4, 5}, 2))

	fmt.Println("\n27. Remove Duplicates:")
	fmt.Println("[1, 2, 2, 3, 4, 4] ->", removeArrayDuplicates([]int{1, 2, 2, 3, 4, 4}))

	fmt.Println("\n28. Frequency:")
	arrayFrequency([]int{1, 2, 2, 3})

	fmt.Println("\n29. Find Missing:")
	fmt.Println("[1, 2, 4, 5], N=5 ->", findMissing([]int{1, 2, 4, 5}, 5))

	fmt.Println("\n30. Find Duplicate:")
	fmt.Println("[1, 3, 4, 2, 2] ->", findDuplicate([]int{1, 3, 4, 2, 2}))

	fmt.Println("\n31. Bubble Sort:")
	fmt.Println("[5, 1, 4, 2, 8] ->", bubbleSort([]int{5, 1, 4, 2, 8}))

	fmt.Println("\n32. Merge Arrays:")
	fmt.Println("[1, 2], [3, 4] ->", mergeArrays([]int{1, 2}, []int{3, 4}))

	fmt.Println("\n33. Common Elements:")
	fmt.Println("[1, 2, 3], [2, 3, 4] ->", findCommon([]int{1, 2, 3}, []int{2, 3, 4}))

	fmt.Println("\n34. Move Zeros:")
	fmt.Println("[0, 1, 0, 3, 12] ->", moveZeros([]int{0, 1, 0, 3, 12}))

	fmt.Println("\n35. Sum Array:")
	fmt.Println("[1, 2, 3, 4] ->", sumArray([]int{1, 2, 3, 4}))

	fmt.Println("\n36. Find Pair Sum:")
	a, b := findPairWithSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("[2, 7, 11, 15], 9 -> {%d, %d}\n", a, b)

	fmt.Println("\n37. Max & Min:")
	max, min := findMaxMin([]int{3, 5, 1, 2, 4, 8})
	fmt.Printf("[3, 5, 1, 2, 4, 8] -> {%d, %d}\n", max, min)

	fmt.Println("\n38. Print Reverse:")
	printReverse([]int{1, 2, 3, 4})

	fmt.Println("\n39. Is Sorted:")
	fmt.Println("[1, 2, 3, 4, 5] ->", isSorted([]int{1, 2, 3, 4, 5}))

	fmt.Println("\n40. Even Odd:")
	countEvenOdd([]int{1, 2, 3, 4, 5})
}
