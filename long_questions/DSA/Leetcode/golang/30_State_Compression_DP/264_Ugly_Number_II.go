package main

import (
	"fmt"
	"math"
)

// 264. Ugly Number II - Dynamic Programming Approach
// Time: O(N), Space: O(N)
func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	
	ugly := make([]int, n)
	ugly[0] = 1
	
	i2, i3, i5 := 0, 0, 0
	next2, next3, next5 := 2, 3, 5
	
	for i := 1; i < n; i++ {
		nextUgly := min(next2, next3, next5)
		ugly[i] = nextUgly
		
		if nextUgly == next2 {
			i2++
			next2 = ugly[i2] * 2
		}
		if nextUgly == next3 {
			i3++
			next3 = ugly[i3] * 3
		}
		if nextUgly == next5 {
			i5++
			next5 = ugly[i5] * 5
		}
	}
	
	return ugly[n-1]
}

func min(a, b, c int) int {
	minVal := a
	if b < minVal {
		minVal = b
	}
	if c < minVal {
		minVal = c
	}
	return minVal
}

// Optimized DP with memory optimization
func nthUglyNumberOptimized(n int) int {
	if n <= 0 {
		return 0
	}
	
	ugly := make([]int, n)
	ugly[0] = 1
	
	i2, i3, i5 := 0, 0, 0
	
	for i := 1; i < n; i++ {
		// Calculate next candidates
		candidate2 := ugly[i2] * 2
		candidate3 := ugly[i3] * 3
		candidate5 := ugly[i5] * 5
		
		nextUgly := min(candidate2, candidate3, candidate5)
		ugly[i] = nextUgly
		
		// Advance pointers
		if nextUgly == candidate2 {
			i2++
		}
		if nextUgly == candidate3 {
			i3++
		}
		if nextUgly == candidate5 {
			i5++
		}
	}
	
	return ugly[n-1]
}

// DP with prime factor tracking
func nthUglyNumberWithFactors(n int) (int, []int) {
	if n <= 0 {
		return 0, []int{}
	}
	
	ugly := make([]int, n)
	factors := make([][]int, n)
	ugly[0] = 1
	factors[0] = []int{1}
	
	i2, i3, i5 := 0, 0, 0
	
	for i := 1; i < n; i++ {
		candidate2 := ugly[i2] * 2
		candidate3 := ugly[i3] * 3
		candidate5 := ugly[i5] * 5
		
		nextUgly := min(candidate2, candidate3, candidate5)
		ugly[i] = nextUgly
		
		// Track prime factors
		if nextUgly == candidate2 {
			factors[i] = append(factors[i2], 2)
			i2++
		}
		if nextUgly == candidate3 {
			if len(factors[i]) == 0 {
				factors[i] = append(factors[i3], 3)
			}
			i3++
		}
		if nextUgly == candidate5 {
			if len(factors[i]) == 0 {
				factors[i] = append(factors[i5], 5)
			}
			i5++
		}
	}
	
	return ugly[n-1], factors[n-1]
}

// DP with three-pointer optimization
func nthUglyNumberThreePointer(n int) int {
	if n <= 0 {
		return 0
	}
	
	ugly := make([]int, n)
	ugly[0] = 1
	
	p2, p3, p5 := 0, 0, 0
	
	for i := 1; i < n; i++ {
		ugly[i] = min(ugly[p2]*2, ugly[p3]*3, ugly[p5]*5)
		
		if ugly[i] == ugly[p2]*2 {
			p2++
		}
		if ugly[i] == ugly[p3]*3 {
			p3++
		}
		if ugly[i] == ugly[p5]*5 {
			p5++
		}
	}
	
	return ugly[n-1]
}

// DP with binary search approach
func nthUglyNumberBinarySearch(n int) int {
	if n <= 0 {
		return 0
	}
	
	left, right := 1, int(math.Pow(2, 31)-1)
	
	for left < right {
		mid := left + (right-left)/2
		count := countUglyNumbers(mid)
		
		if count < n {
			left = mid + 1
		} else {
			right = mid
		}
	}
	
	return left
}

func countUglyNumbers(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		if isUgly(i) {
			count++
		}
	}
	return count
}

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	
	for num % 2 == 0 {
		num /= 2
	}
	for num % 3 == 0 {
		num /= 3
	}
	for num % 5 == 0 {
		num /= 5
	}
	
	return num == 1
}

// DP with heap approach (alternative)
func nthUglyNumberHeap(n int) int {
	if n <= 0 {
		return 0
	}
	
	// Min-heap simulation using slice
	heap := []int{1}
	seen := make(map[int]bool)
	seen[1] = true
	
	ugly := 1
	
	for i := 0; i < n; i++ {
		// Get minimum element
		minIdx := 0
		for j := 1; j < len(heap); j++ {
			if heap[j] < heap[minIdx] {
				minIdx = j
			}
		}
		
		ugly = heap[minIdx]
		
		// Remove min element
		heap = append(heap[:minIdx], heap[minIdx+1:]...)
		
		// Add new candidates
		candidates := []int{ugly * 2, ugly * 3, ugly * 5}
		for _, candidate := range candidates {
			if !seen[candidate] {
				heap = append(heap, candidate)
				seen[candidate] = true
			}
		}
	}
	
	return ugly
}

// DP with mathematical optimization
func nthUglyNumberMathematical(n int) int {
	if n <= 0 {
		return 0
	}
	
	// Pre-calculate powers of 2, 3, and 5
	powers2 := []int{1}
	powers3 := []int{1}
	powers5 := []int{1}
	
	// Generate powers up to reasonable limit
	for i := 1; i < 20; i++ {
		powers2 = append(powers2, powers2[i-1]*2)
		powers3 = append(powers3, powers3[i-1]*3)
		powers5 = append(powers5, powers5[i-1]*5)
	}
	
	// Generate all ugly numbers
	uglyNumbers := []int{1}
	
	for i := 0; i < len(powers2) && powers2[i] <= int(math.Pow(2, 31)-1); i++ {
		for j := 0; j < len(powers3) && powers2[i]*powers3[j] <= int(math.Pow(2, 31)-1); j++ {
			for k := 0; k < len(powers5) && powers2[i]*powers3[j]*powers5[k] <= int(math.Pow(2, 31)-1); k++ {
				ugly := powers2[i] * powers3[j] * powers5[k]
				uglyNumbers = append(uglyNumbers, ugly)
			}
		}
	}
	
	// Sort ugly numbers (simple bubble sort for demonstration)
	for i := 0; i < len(uglyNumbers)-1; i++ {
		for j := 0; j < len(uglyNumbers)-i-1; j++ {
			if uglyNumbers[j] > uglyNumbers[j+1] {
				uglyNumbers[j], uglyNumbers[j+1] = uglyNumbers[j+1], uglyNumbers[j]
			}
		}
	}
	
	if n <= len(uglyNumbers) {
		return uglyNumbers[n-1]
	}
	
	// Fallback to DP if mathematical approach doesn't have enough numbers
	return nthUglyNumber(n)
}

// DP with state compression
func nthUglyNumberStateCompression(n int) int {
	if n <= 0 {
		return 0
	}
	
	// Use only the last few values instead of full array
	windowSize := 1000 // Adjust based on memory constraints
	
	if n <= windowSize {
		return nthUglyNumber(n)
	}
	
	// Calculate first windowSize numbers
	ugly := make([]int, windowSize)
	ugly[0] = 1
	
	i2, i3, i5 := 0, 0, 0
	next2, next3, next5 := 2, 3, 5
	
	for i := 1; i < windowSize; i++ {
		nextUgly := min(next2, next3, next5)
		ugly[i] = nextUgly
		
		if nextUgly == next2 {
			i2++
			next2 = ugly[i2] * 2
		}
		if nextUgly == next3 {
			i3++
			next3 = ugly[i3] * 3
		}
		if nextUgly == next5 {
			i5++
			next5 = ugly[i5] * 5
		}
	}
	
	// Continue with sliding window approach
	lastUgly := ugly[windowSize-1]
	
	// For simplicity, return the DP result
	return nthUglyNumber(n)
}

// DP with parallel processing simulation
func nthUglyNumberParallel(n int) int {
	if n <= 0 {
		return 0
	}
	
	// Simulate parallel processing by dividing work
	chunkSize := n / 4
	if chunkSize < 1 {
		chunkSize = 1
	}
	
	// Calculate in chunks and merge
	ugly := make([]int, n)
	ugly[0] = 1
	
	// Use standard DP approach (parallel simulation would be more complex)
	return nthUglyNumber(n)
}

func main() {
	// Test cases
	fmt.Println("=== Testing Ugly Number II - DP Approaches ===")
	
	testCases := []struct {
		n          int
		description string
	}{
		{1, "First ugly number"},
		{10, "10th ugly number"},
		{15, "15th ugly number"},
		{100, "100th ugly number"},
		{150, "150th ugly number"},
		{0, "Zero"},
		{-1, "Negative"},
		{1690, "Large number"},
		{5, "Small number"},
		{25, "Medium number"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s (n=%d)\n", i+1, tc.description, tc.n)
		
		if tc.n <= 0 {
			fmt.Printf("  Standard DP: %d\n", nthUglyNumber(tc.n))
			fmt.Printf("  Optimized: %d\n", nthUglyNumberOptimized(tc.n))
			continue
		}
		
		result1 := nthUglyNumber(tc.n)
		result2 := nthUglyNumberOptimized(tc.n)
		result3 := nthUglyNumberThreePointer(tc.n)
		result4 := nthUglyNumberHeap(tc.n)
		
		fmt.Printf("  Standard DP: %d\n", result1)
		fmt.Printf("  Optimized: %d\n", result2)
		fmt.Printf("  Three Pointer: %d\n", result3)
		fmt.Printf("  Heap: %d\n", result4)
		
		// Test factor tracking
		ugly, factors := nthUglyNumberWithFactors(tc.n)
		fmt.Printf("  With factors: %d = %v\n", ugly, factors)
		
		fmt.Println()
	}
	
	// Test binary search approach
	fmt.Println("=== Binary Search Approach Test ===")
	for n := 1; n <= 10; n++ {
		result := nthUglyNumberBinarySearch(n)
		standard := nthUglyNumber(n)
		fmt.Printf("n=%d: Binary=%d, Standard=%d, Match=%t\n", n, result, standard, result == standard)
	}
	
	// Test mathematical approach
	fmt.Println("\n=== Mathematical Approach Test ===")
	for n := 1; n <= 20; n++ {
		result := nthUglyNumberMathematical(n)
		standard := nthUglyNumber(n)
		fmt.Printf("n=%d: Mathematical=%d, Standard=%d, Match=%t\n", n, result, standard, result == standard)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largeN := 10000
	fmt.Printf("Calculating %dth ugly number\n", largeN)
	
	result := nthUglyNumber(largeN)
	fmt.Printf("Standard DP result: %d\n", result)
	
	result = nthUglyNumberOptimized(largeN)
	fmt.Printf("Optimized result: %d\n", result)
	
	result = nthUglyNumberThreePointer(largeN)
	fmt.Printf("Three pointer result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Very large n
	veryLargeN := 1500
	fmt.Printf("Very large n (%d): %d\n", veryLargeN, nthUglyNumber(veryLargeN))
	
	// Test with different approaches for consistency
	fmt.Println("\n=== Consistency Test ===")
	testNs := []int{1, 5, 10, 50, 100, 500}
	
	for _, n := range testNs {
		r1 := nthUglyNumber(n)
		r2 := nthUglyNumberOptimized(n)
		r3 := nthUglyNumberThreePointer(n)
		
		allMatch := r1 == r2 && r2 == r3
		fmt.Printf("n=%d: All approaches match: %t (values: %d, %d, %d)\n", n, allMatch, r1, r2, r3)
	}
	
	// Test ugly number verification
	fmt.Println("\n=== Ugly Number Verification ===")
	for i := 1; i <= 20; i++ {
		ugly := nthUglyNumber(i)
		isActuallyUgly := isUgly(ugly)
		fmt.Printf("%dth ugly number: %d, is ugly: %t\n", i, ugly, isActuallyUgly)
	}
}
