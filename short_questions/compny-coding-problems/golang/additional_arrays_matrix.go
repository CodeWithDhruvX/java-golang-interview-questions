package main

import (
	"fmt"
	"math"
	"sort"
)

// 16. Find leaders in array
func findLeaders(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}
	maxRight := arr[n-1]
	fmt.Printf("%d ", maxRight)
	for i := n - 2; i >= 0; i-- {
		if arr[i] > maxRight {
			maxRight = arr[i]
			fmt.Printf("%d ", maxRight)
		}
	}
	fmt.Println()
}

// 17. Find equilibrium index
func equilibriumPoint(arr []int) int {
	totalSum := 0
	for _, v := range arr {
		totalSum += v
	}
	leftSum := 0
	for i, v := range arr {
		totalSum -= v
		if leftSum == totalSum {
			return i
		}
		leftSum += v
	}
	return -1
}

// 18. Subarray with given sum
func subArraySum(arr []int, sum int) {
	currSum, start := arr[0], 0
	for i := 1; i <= len(arr); i++ {
		for currSum > sum && start < i-1 {
			currSum -= arr[start]
			start++
		}
		if currSum == sum {
			fmt.Printf("Sum found between indexes %d and %d\n", start, i-1)
			return
		}
		if i < len(arr) {
			currSum += arr[i]
		}
	}
	fmt.Println("No subarray found")
}

// 19. Kadane’s algorithm
func kadanes(arr []int) int {
	// If array is possibly empty or all negative logic required
	// But standard Kadane assumes at least one positive or standard handling
	if len(arr) == 0 {
		return 0
	}
	maxSoFar := math.MinInt32
	currMax := 0
	for _, num := range arr {
		currMax += num
		if maxSoFar < currMax {
			maxSoFar = currMax
		}
		if currMax < 0 {
			currMax = 0
		}
	}
	return maxSoFar
}

// 20. Find majority element
func majorityElement(arr []int) int {
	var candidate int
	count := 0
	for _, num := range arr {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// 21. Rearrange array alternately
func rearrangeAlternate(arr []int) {
	sort.Ints(arr)
	i, j := 0, len(arr)-1
	for i < j {
		fmt.Printf("%d %d ", arr[j], arr[i])
		j--
		i++
	}
	if i == j {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

// 22. Rotate array using reversal algorithm (Covered in Array Q25/Q26)

// 23. Find union of two arrays
func unionArrays(arr1, arr2 []int) []int {
	m := make(map[int]bool)
	for _, v := range arr1 {
		m[v] = true
	}
	for _, v := range arr2 {
		m[v] = true
	}
	var res []int
	for k := range m {
		res = append(res, k)
	}
	sort.Ints(res)
	return res
}

// 24. Find intersection of two arrays
func intersectionArrays(arr1, arr2 []int) []int {
	m := make(map[int]bool)
	for _, v := range arr1 {
		m[v] = true
	}
	var res []int
	for _, v := range arr2 {
		if m[v] {
			res = append(res, v)
			delete(m, v) // avoid duplicates
		}
	}
	sort.Ints(res)
	return res
}

// 25. Count pairs with given difference
func countDiffPairs(arr []int, k int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if int(math.Abs(float64(arr[i]-arr[j]))) == k {
				count++
			}
		}
	}
	return count
}

// 26. Find peak element
func findPeak(arr []int) int {
	n := len(arr)
	if n == 1 {
		return 0
	}
	if arr[0] >= arr[1] {
		return 0
	}
	if arr[n-1] >= arr[n-2] {
		return n - 1
	}
	for i := 1; i < n-1; i++ {
		if arr[i] >= arr[i-1] && arr[i] >= arr[i+1] {
			return i
		}
	}
	return -1
}

// 27. Left rotate by 1
func leftRotateOne(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	temp := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = temp
	return arr
}

// 28. Find minimum difference pair
func minDiffPair(arr []int) int {
	sort.Ints(arr)
	minDiff := math.MaxInt32
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < minDiff {
			minDiff = arr[i+1] - arr[i]
		}
	}
	return minDiff
}

// 29. Product of array except self
func productExceptSelf(arr []int) []int {
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	prod := make([]int, n)
	left[0] = 1
	right[n-1] = 1
	for i := 1; i < n; i++ {
		left[i] = arr[i-1] * left[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = arr[i+1] * right[i+1]
	}
	for i := 0; i < n; i++ {
		prod[i] = left[i] * right[i]
	}
	return prod
}

// 30. Find subarray with max product
func maxProductSubarray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxSoFar := arr[0]
	minSoFar := arr[0]
	result := maxSoFar
	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		tempMax := int(math.Max(float64(curr), math.Max(float64(curr*maxSoFar), float64(curr*minSoFar))))
		minSoFar = int(math.Min(float64(curr), math.Min(float64(curr*maxSoFar), float64(curr*minSoFar))))
		maxSoFar = tempMax
		result = int(math.Max(float64(maxSoFar), float64(result)))
	}
	return result
}

// 32. Separate positive and negative numbers
func separatePosNeg(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			j++
		}
	}
	return arr
}

// 33. Count distinct elements
func countDistinct(arr []int) int {
	s := make(map[int]bool)
	for _, num := range arr {
		s[num] = true
	}
	return len(s)
}

// 34. Replace element with next greatest
func replaceNextGreatest(arr []int) []int {
	maxFromRight := -1
	for i := len(arr) - 1; i >= 0; i-- {
		temp := arr[i]
		arr[i] = maxFromRight
		if temp > maxFromRight {
			maxFromRight = temp
		}
	}
	return arr
}

// 35. Find smallest subarray with sum > X
func smallestSubWithSum(arr []int, x int) int {
	minLen := len(arr) + 1
	currSum, start, end := 0, 0, 0
	for end < len(arr) {
		currSum += arr[end]
		end++
		for currSum > x && start < end {
			if end-start < minLen {
				minLen = end - start
			}
			currSum -= arr[start]
			start++
		}
	}
	return minLen
}

// --- MATRIX ---

// 36. Matrix addition
func addMatrices(A, B [][]int) [][]int {
	rows := len(A)
	cols := len(A[0])
	C := make([][]int, rows)
	for i := range C {
		C[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

// 37. Matrix multiplication
func multiplyMatrices(A, B [][]int) [][]int {
	rowsA := len(A)
	colsA := len(A[0])
	colsB := len(B[0])
	C := make([][]int, rowsA)
	for i := range C {
		C[i] = make([]int, colsB)
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

// 38. Transpose of matrix
func transpose(A [][]int) [][]int {
	rows := len(A)
	cols := len(A[0])
	T := make([][]int, cols)
	for i := range T {
		T[i] = make([]int, rows)
		for j := 0; j < rows; j++ {
			T[i][j] = A[j][i]
		}
	}
	return T
}

// 39. Rotate matrix 90 degrees (clockwise)
func rotateMatrix(A [][]int) [][]int {
	T := transpose(A) // returns new matrix
	// Reverse each row
	for i := 0; i < len(T); i++ {
		left, right := 0, len(T[i])-1
		for left < right {
			T[i][left], T[i][right] = T[i][right], T[i][left]
			left++
			right--
		}
	}
	return T
}

// 40. Print matrix in spiral order
func spiralOrder(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			fmt.Printf("%d ", matrix[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			fmt.Printf("%d ", matrix[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				fmt.Printf("%d ", matrix[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				fmt.Printf("%d ", matrix[i][left])
			}
			left++
		}
	}
	fmt.Println()
}

// 41. Search element in sorted matrix
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	row, col := 0, cols-1
	for row < rows && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}

// 42. Diagonal sum (Primary & Secondary)
func diagonalSum(matrix [][]int) int {
	n := len(matrix)
	sum := 0
	for i := 0; i < n; i++ {
		sum += matrix[i][i]
		sum += matrix[i][n-i-1]
	}
	if n%2 != 0 {
		sum -= matrix[n/2][n/2]
	}
	return sum
}

// 43. Print boundary elements
func printBoundary(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	for col := 0; col < cols; col++ {
		fmt.Printf("%d ", matrix[0][col])
	}
	for row := 1; row < rows; row++ {
		fmt.Printf("%d ", matrix[row][cols-1])
	}
	for col := cols - 2; col >= 0; col-- {
		fmt.Printf("%d ", matrix[rows-1][col])
	}
	for row := rows - 2; row > 0; row-- {
		fmt.Printf("%d ", matrix[row][0])
	}
	fmt.Println()
}

// 44. Check symmetric matrix
func isSymmetric(matrix [][]int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	if rows != cols {
		return false
	}
	// Check transpose logic
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] != matrix[j][i] {
				return false
			}
		}
	}
	return true
}

// 46. Count zeros and ones
func countZeroOne(matrix [][]int) {
	zeros, ones := 0, 0
	for _, row := range matrix {
		for _, cell := range row {
			if cell == 0 {
				zeros++
			} else if cell == 1 {
				ones++
			}
		}
	}
	fmt.Printf("Zeros: %d, Ones: %d\n", zeros, ones)
}

// 47. Row with maximum 1s
func rowMaxOnes(matrix [][]int) int {
	maxOnes := 0
	rowIndex := -1
	for i, row := range matrix {
		count := 0
		for _, val := range row {
			if val == 1 {
				count++
			}
		}
		if count > maxOnes {
			maxOnes = count
			rowIndex = i
		}
	}
	return rowIndex
}

// 49. Snake pattern printing
func snakePattern(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		if i%2 == 0 {
			for j := 0; j < cols; j++ {
				fmt.Printf("%d ", matrix[i][j])
			}
		} else {
			for j := cols - 1; j >= 0; j-- {
				fmt.Printf("%d ", matrix[i][j])
			}
		}
	}
	fmt.Println()
}

// 50. Identity matrix check
func isIdentity(matrix [][]int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == j && matrix[i][j] != 1 {
				return false
			}
			if i != j && matrix[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println("16. Find Leaders:")
	findLeaders([]int{16, 17, 4, 3, 5, 2})

	fmt.Println("\n17. Equilibrium Point:")
	fmt.Println("Index:", equilibriumPoint([]int{1, 3, 5, 2, 2}))

	fmt.Println("\n18. Subarray Sum:")
	subArraySum([]int{1, 2, 3, 7, 5}, 12)

	fmt.Println("\n19. Kadanes:")
	fmt.Println("Max Sum:", kadanes([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	fmt.Println("\n20. Majority Element:")
	fmt.Println("Majority:", majorityElement([]int{3, 2, 3}))

	fmt.Println("\n21. Rearrange Alternate:")
	rearrangeAlternate([]int{1, 2, 3, 4, 5, 6})

	fmt.Println("\n23. Union Arrays:")
	fmt.Println("Union:", unionArrays([]int{1, 2, 3}, []int{2, 3, 4}))

	fmt.Println("\n24. Intersection Arrays:")
	fmt.Println("Intersection:", intersectionArrays([]int{1, 2, 3}, []int{2, 3, 4}))

	fmt.Println("\n25. Count Diff Pairs:")
	fmt.Println("Count:", countDiffPairs([]int{1, 5, 3, 4, 2}, 3))

	fmt.Println("\n26. Find Peak:")
	fmt.Println("Peak Index:", findPeak([]int{1, 2, 3, 1}))

	fmt.Println("\n27. Left Rotate By 1:")
	fmt.Println("Rotated:", leftRotateOne([]int{1, 2, 3, 4, 5}))

	fmt.Println("\n28. Min Diff Pair:")
	fmt.Println("Min Diff:", minDiffPair([]int{2, 4, 5, 9, 7}))

	fmt.Println("\n29. Product Except Self:")
	fmt.Println("Product:", productExceptSelf([]int{1, 2, 3, 4}))

	fmt.Println("\n30. Max Product Subarray:")
	fmt.Println("Max Product:", maxProductSubarray([]int{2, 3, -2, 4}))

	fmt.Println("\n32. Separate Pos Neg:")
	fmt.Println("Separated:", separatePosNeg([]int{-1, 2, -3, 4, 5, 6, -7, 8, 9}))

	fmt.Println("\n33. Count Distinct:")
	fmt.Println("Distinct Count:", countDistinct([]int{10, 20, 20, 10, 30}))

	fmt.Println("\n34. Replace Next Greatest:")
	fmt.Println("Replaced:", replaceNextGreatest([]int{16, 17, 4, 3, 5, 2}))

	fmt.Println("\n35. Smallest Sub Sum > X:")
	fmt.Println("Min Len:", smallestSubWithSum([]int{1, 4, 45, 6, 0, 19}, 51))

	fmt.Println("\n--- MATRIX ---")
	A := [][]int{{1, 2}, {3, 4}}
	B := [][]int{{5, 6}, {7, 8}}
	fmt.Println("36. Add Matrices:", addMatrices(A, B))
	fmt.Println("37. Multiply Matrices:", multiplyMatrices(A, [][]int{{1, 0}, {0, 1}}))
	fmt.Println("38. Transpose:", transpose(A))
	fmt.Println("39. Rotate 90:", rotateMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println("40. Spiral Order:")
	spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println("41. Search Matrix:", searchMatrix([][]int{{1, 3, 5}, {7, 9, 11}}, 9))
	fmt.Println("42. Diagonal Sum:", diagonalSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println("43. Print Boundary:")
	printBoundary([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println("44. Is Symmetric:", isSymmetric([][]int{{1, 2}, {2, 1}}))
	fmt.Println("46. Count Zero One:")
	countZeroOne([][]int{{0, 1}, {1, 0}})
	fmt.Println("47. Row Max Ones:", rowMaxOnes([][]int{{0, 1, 1}, {0, 0, 1}}))
	fmt.Println("49. Snake Pattern:")
	snakePattern([][]int{{1, 2}, {3, 4}})
	fmt.Println("50. Is Identity:", isIdentity([][]int{{1, 0}, {0, 1}}))
}
