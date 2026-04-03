package main

import (
	"fmt"
	"math"
)

// 53. Maximum Subarray - Divide and Conquer Approach
// Time: O(N log N), Space: O(log N) for recursion stack
func maxSubArrayDivideAndConquer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	return maxSubArrayHelper(nums, 0, len(nums)-1)
}

func maxSubArrayHelper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	
	mid := left + (right-left)/2
	
	// Maximum subarray in left half
	leftMax := maxSubArrayHelper(nums, left, mid)
	
	// Maximum subarray in right half
	rightMax := maxSubArrayHelper(nums, mid+1, right)
	
	// Maximum subarray crossing the middle
	crossMax := maxCrossingSubArray(nums, left, mid, right)
	
	return max(leftMax, rightMax, crossMax)
}

func maxCrossingSubArray(nums []int, left, mid, right int) int {
	// Maximum sum starting from mid and going left
	leftSum := math.MinInt32
	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}
	
	// Maximum sum starting from mid+1 and going right
	rightSum := math.MinInt32
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}
	
	return leftSum + rightSum
}

func max(a, b, c int) int {
	maxVal := a
	if b > maxVal {
		maxVal = b
	}
	if c > maxVal {
		maxVal = c
	}
	return maxVal
}

// Kadane's algorithm for comparison (O(N))
func maxSubArrayKadane(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	maxSoFar := nums[0]
	maxEndingHere := nums[0]
	
	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(nums[i], maxEndingHere+nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	
	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Divide and Conquer with memoization
func maxSubArrayDivideAndConquerMemo(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	memo := make(map[string]int)
	return maxSubArrayHelperMemo(nums, 0, len(nums)-1, memo)
}

func maxSubArrayHelperMemo(nums []int, left, right int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", left, right)
	if val, exists := memo[key]; exists {
		return val
	}
	
	if left == right {
		memo[key] = nums[left]
		return nums[left]
	}
	
	mid := left + (right-left)/2
	
	leftMax := maxSubArrayHelperMemo(nums, left, mid, memo)
	rightMax := maxSubArrayHelperMemo(nums, mid+1, right, memo)
	crossMax := maxCrossingSubArray(nums, left, mid, right)
	
	result := max(leftMax, rightMax, crossMax)
	memo[key] = result
	return result
}

// Divide and Conquer with tracking indices
func maxSubArrayDivideAndConquerWithIndices(nums []int) (int, int, int) {
	if len(nums) == 0 {
		return 0, -1, -1
	}
	
	return maxSubArrayHelperWithIndices(nums, 0, len(nums)-1)
}

func maxSubArrayHelperWithIndices(nums []int, left, right int) (int, int, int) {
	if left == right {
		return nums[left], left, right
	}
	
	mid := left + (right-left)/2
	
	leftSum, leftStart, leftEnd := maxSubArrayHelperWithIndices(nums, left, mid)
	rightSum, rightStart, rightEnd := maxSubArrayHelperWithIndices(nums, mid+1, right)
	crossSum, crossStart, crossEnd := maxCrossingSubArrayWithIndices(nums, left, mid, right)
	
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftSum, leftStart, leftEnd
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightSum, rightStart, rightEnd
	} else {
		return crossSum, crossStart, crossEnd
	}
}

func maxCrossingSubArrayWithIndices(nums []int, left, mid, right int) (int, int, int) {
	// Maximum sum starting from mid and going left
	leftSum := math.MinInt32
	sum := 0
	leftIdx := mid
	
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
			leftIdx = i
		}
	}
	
	// Maximum sum starting from mid+1 and going right
	rightSum := math.MinInt32
	sum = 0
	rightIdx := mid + 1
	
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
			rightIdx = i
		}
	}
	
	return leftSum + rightSum, leftIdx, rightIdx
}

// Divide and Conquer for maximum product subarray
func maxProductDivideAndConquer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	return maxProductHelper(nums, 0, len(nums)-1)
}

func maxProductHelper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}
	
	mid := left + (right-left)/2
	
	// Maximum product in left half
	leftMax := maxProductHelper(nums, left, mid)
	
	// Maximum product in right half
	rightMax := maxProductHelper(nums, mid+1, right)
	
	// Maximum product crossing the middle
	crossMax := maxCrossingProduct(nums, left, mid, right)
	
	return max(leftMax, rightMax, crossMax)
}

func maxCrossingProduct(nums []int, left, mid, right int) int {
	// For product, we need to consider both positive and negative values
	// Maximum product starting from mid and going left
	maxLeftProd := math.MinInt32
	minLeftProd := math.MaxInt32
	prod := 1
	
	for i := mid; i >= left; i-- {
		prod *= nums[i]
		if prod > maxLeftProd {
			maxLeftProd = prod
		}
		if prod < minLeftProd {
			minLeftProd = prod
		}
	}
	
	// Maximum product starting from mid+1 and going right
	maxRightProd := math.MinInt32
	minRightProd := math.MaxInt32
	prod = 1
	
	for i := mid + 1; i <= right; i++ {
		prod *= nums[i]
		if prod > maxRightProd {
			maxRightProd = prod
		}
		if prod < minRightProd {
			minRightProd = prod
		}
	}
	
	// Consider all combinations
	cross1 := maxLeftProd * maxRightProd
	cross2 := maxLeftProd * minRightProd
	cross3 := minLeftProd * maxRightProd
	cross4 := minLeftProd * minRightProd
	
	return max(cross1, cross2, cross3, cross4)
}

// Divide and Conquer for circular subarray
func maxSubarrayCircularDivideAndConquer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	// Case 1: Maximum subarray is non-circular (standard Kadane)
	maxKadane := maxSubArrayKadane(nums)
	
	// Case 2: Maximum subarray is circular
	if len(nums) == 1 {
		return maxKadane
	}
	
	// Calculate total sum
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	
	// Find minimum subarray (Kadane on inverted array)
	minKadane := minSubArray(nums)
	
	maxCircular := totalSum - minKadane
	
	if maxCircular == 0 && maxKadane < 0 {
		return maxKadane // All negative numbers
	}
	
	return max(maxKadane, maxCircular)
}

func minSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	minSoFar := nums[0]
	minEndingHere := nums[0]
	
	for i := 1; i < len(nums); i++ {
		minEndingHere = min(nums[i], minEndingHere+nums[i])
		minSoFar = min(minSoFar, minEndingHere)
	}
	
	return minSoFar
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test cases
	fmt.Println("=== Testing Maximum Subarray - Divide and Conquer ===")
	
	testCases := []struct {
		nums       []int
		description string
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, "Standard case"},
		{[]int{1}, "Single element"},
		{[]int{5, 4, -1, 7, 8}, "All positive"},
		{[]int{-1, -2, -3, -4}, "All negative"},
		{[]int{0, 0, 0, 0}, "All zeros"},
		{[]int{-2, -1, -2, -3, -1, -4}, "Mixed negatives"},
		{[]int{2, 3, -2, 4}, "Simple positive"},
		{[]int{-1, 2, 3, -5, 4, 6, -1, 2, -3}, "Complex case"},
		{[]int{100, -1, 100, -1, 100}, "Large positives"},
		{[]int{-2, -3, 4, -1, -2, 1, 5, -3}, "Classic example"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.description)
		fmt.Printf("  Array: %v\n", tc.nums)
		
		result1 := maxSubArrayDivideAndConquer(tc.nums)
		result2 := maxSubArrayKadane(tc.nums)
		result3 := maxSubArrayDivideAndConquerMemo(tc.nums)
		
		sum, start, end := maxSubArrayDivideAndConquerWithIndices(tc.nums)
		
		fmt.Printf("  Divide & Conquer: %d\n", result1)
		fmt.Printf("  Kadane's: %d\n", result2)
		fmt.Printf("  With Memoization: %d\n", result3)
		fmt.Printf("  With Indices: sum=%d, start=%d, end=%d\n", sum, start, end)
		
		// Test circular version
		circularResult := maxSubarrayCircularDivideAndConquer(tc.nums)
		fmt.Printf("  Circular: %d\n\n", circularResult)
	}
	
	// Performance test
	fmt.Println("=== Performance Test ===")
	
	largeArray := make([]int, 10000)
	for i := range largeArray {
		if i%3 == 0 {
			largeArray[i] = -1
		} else {
			largeArray[i] = 1
		}
	}
	
	fmt.Printf("Large array test with %d elements\n", len(largeArray))
	
	result := maxSubArrayDivideAndConquer(largeArray)
	fmt.Printf("Divide & Conquer result: %d\n", result)
	
	result = maxSubArrayKadane(largeArray)
	fmt.Printf("Kadane's result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Empty array
	fmt.Printf("Empty array: %d\n", maxSubArrayDivideAndConquer([]int{}))
	
	// Single negative
	fmt.Printf("Single negative: %d\n", maxSubArrayDivideAndConquer([]int{-5}))
	
	// Single positive
	fmt.Printf("Single positive: %d\n", maxSubArrayDivideAndConquer([]int{5}))
	
	// Alternating pattern
	alternating := []int{1, -1, 1, -1, 1, -1, 1, -1}
	fmt.Printf("Alternating: %d\n", maxSubArrayDivideAndConquer(alternating))
	
	// Large values
	largeVals := []int{1000000, -1000000, 1000000, -1000000, 1000000}
	fmt.Printf("Large values: %d\n", maxSubArrayDivideAndConquer(largeVals))
	
	// Test maximum product
	fmt.Println("\n=== Maximum Product Test ===")
	productTest := []int{2, 3, -2, 4}
	fmt.Printf("Max product: %d\n", maxProductDivideAndConquer(productTest))
	
	productTest2 := []int{-2, 0, -1}
	fmt.Printf("Max product with zeros: %d\n", maxProductDivideAndConquer(productTest2))
}
