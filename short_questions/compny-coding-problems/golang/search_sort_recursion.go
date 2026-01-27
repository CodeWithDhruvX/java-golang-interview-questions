package main

import (
	"fmt"
)

// 71. Linear search
func linearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

// 72. Binary search (Sorted Array)
func binarySearch(arr []int, target int) int {
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

// 73. Bubble sort
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 74. Selection sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// 75. Insertion sort
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 76. Merge sort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// 77. Quick sort
func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// 78. Recursion: Factorial
func factorialRec(N int) int {
	if N <= 1 {
		return 1
	}
	return N * factorialRec(N-1)
}

// 79. Recursion: Fibonacci
func fibonacciRec(N int) int {
	if N <= 1 {
		return N
	}
	return fibonacciRec(N-1) + fibonacciRec(N-2)
}

// 80. Recursion: Sum of digits
func sumDigitsRec(N int) int {
	if N == 0 {
		return 0
	}
	return (N % 10) + sumDigitsRec(N/10)
}

// 81. Recursion: Reverse string
func reverseStringRec(str string) string {
	if len(str) == 0 {
		return ""
	}
	return reverseStringRec(str[1:]) + string(str[0])
}

// 82. Power of number (Recursion)
func powerRec(base, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * powerRec(base, exp-1)
}

// 83. GCD Recursion
func gcdRec(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdRec(b, a%b)
}

// 84. Tower of Hanoi
func towerOfHanoi(n int, from, to, aux string) {
	if n == 1 {
		fmt.Printf("Move disk 1 from %s to %s\n", from, to)
		return
	}
	towerOfHanoi(n-1, from, aux, to)
	fmt.Printf("Move disk %d from %s to %s\n", n, from, to)
	towerOfHanoi(n-1, aux, to, from)
}

// 85. Permutations of a string
func permute(str string, l, r int) {
	if l == r {
		fmt.Println(str)
	} else {
		for i := l; i <= r; i++ {
			runes := []rune(str)
			runes[l], runes[i] = runes[i], runes[l]
			permute(string(runes), l+1, r)
			runes[l], runes[i] = runes[i], runes[l] // backtrack
		}
	}
}

// 86. Nth Fibonacci Number (Recursive/Iterative)
// Already covered in Q79 (Recursive) and Q58 (Iterative)

// 87. Binary Search Recursive
func binarySearchRec(arr []int, low, high, target int) int {
	if low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			return binarySearchRec(arr, low, mid-1, target)
		}
		return binarySearchRec(arr, mid+1, high, target)
	}
	return -1
}

// 88. Check Palindrome String (Recursive)
func isPalindromeRec(str string) bool {
	if len(str) <= 1 {
		return true
	}
	if str[0] != str[len(str)-1] {
		return false
	}
	return isPalindromeRec(str[1 : len(str)-1])
}

// 89. Subset Sum Problem (Recursion)
func isSubsetSum(arr []int, n, sum int) bool {
	if sum == 0 {
		return true
	}
	if n == 0 {
		return false
	}
	if arr[n-1] > sum {
		return isSubsetSum(arr, n-1, sum)
	}
	return isSubsetSum(arr, n-1, sum) || isSubsetSum(arr, n-1, sum-arr[n-1])
}

// 90. Print 1 to N without loops
func printNos(N int) {
	if N > 0 {
		printNos(N - 1)
		fmt.Printf("%d ", N)
	}
}

func main() {
	fmt.Println("71. Linear Search:")
	fmt.Println("[10, 20, 30, 40], 30 ->", linearSearch([]int{10, 20, 30, 40}, 30))

	fmt.Println("\n72. Binary Search:")
	fmt.Println("[10, 20, 30, 40], 30 ->", binarySearch([]int{10, 20, 30, 40}, 30))

	fmt.Println("\n73. Bubble Sort:")
	arr := []int{5, 1, 4, 2, 8}
	bubbleSort(arr)
	fmt.Println("[5, 1, 4, 2, 8] ->", arr)

	fmt.Println("\n74. Selection Sort:")
	arr2 := []int{5, 1, 4, 2, 8}
	selectionSort(arr2)
	fmt.Println("[5, 1, 4, 2, 8] ->", arr2)

	fmt.Println("\n75. Insertion Sort:")
	arr3 := []int{5, 1, 4, 2, 8}
	insertionSort(arr3)
	fmt.Println("[5, 1, 4, 2, 8] ->", arr3)

	fmt.Println("\n76. Merge Sort:")
	fmt.Println("[5, 1, 4, 2, 8] ->", mergeSort([]int{5, 1, 4, 2, 8}))

	fmt.Println("\n77. Quick Sort:")
	arr4 := []int{10, 7, 8, 9, 1, 5}
	quickSort(arr4, 0, len(arr4)-1)
	fmt.Println("[10, 7, 8, 9, 1, 5] ->", arr4)

	fmt.Println("\n78. Factorial Rec:")
	fmt.Println("5 ->", factorialRec(5))

	fmt.Println("\n79. Fibonacci Rec:")
	fmt.Println("5 ->", fibonacciRec(5))

	fmt.Println("\n80. Sum Digits Rec:")
	fmt.Println("123 ->", sumDigitsRec(123))

	fmt.Println("\n81. Reverse String Rec:")
	fmt.Println("hello ->", reverseStringRec("hello"))

	fmt.Println("\n82. Power Rec:")
	fmt.Println("2, 3 ->", powerRec(2, 3))

	fmt.Println("\n83. GCD Rec:")
	fmt.Println("12, 18 ->", gcdRec(12, 18))

	fmt.Println("\n84. Tower of Hanoi (3):")
	towerOfHanoi(3, "A", "C", "B")

	fmt.Println("\n85. Permutations (ABC):")
	permute("ABC", 0, 2)

	fmt.Println("\n87. Binary Search Rec:")
	fmt.Println("[10, 20, 30, 40], 30 ->", binarySearchRec([]int{10, 20, 30, 40}, 0, 3, 30))

	fmt.Println("\n88. Is Palindrome Rec:")
	fmt.Println("madam ->", isPalindromeRec("madam"))
	fmt.Println("hello ->", isPalindromeRec("hello"))

	fmt.Println("\n89. Subset Sum:")
	fmt.Println("[3, 34, 4, 12, 5, 2], 9 ->", isSubsetSum([]int{3, 34, 4, 12, 5, 2}, 6, 9))

	fmt.Println("\n90. Print 1 to N:")
	printNos(10)
	fmt.Println()
}
