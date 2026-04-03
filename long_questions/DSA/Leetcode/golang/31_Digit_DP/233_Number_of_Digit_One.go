package main

import (
	"fmt"
	"math"
)

// 233. Number of Digit One - Digit DP
// Time: O(N * D * 9), Space: O(N * D) where N is number length, D is digit count
func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	
	// Convert to string to get length
	s := fmt.Sprintf("%d", n)
	length := len(s)
	
	// Count digit ones for numbers with less digits
	count := 0
	for i := 1; i < length; i++ {
		count += countDigitOneHelper(i-1)
	}
	
	// Count for numbers with same number of digits
	firstDigit := int(s[0] - '0')
	remaining := n - firstDigit*int(math.Pow(10, float64(length-1)))
	
	if firstDigit == 1 {
		count += remaining + 1
	} else {
		count += int(math.Pow(10, float64(length-1))) + countDigitOneHelper(firstDigit-1)*int(math.Pow(10, float64(length-1)))
	}
	
	return count
}

func countDigitOneHelper(n int) int {
	if n <= 0 {
		return 0
	}
	
	count := 0
	for i := 1; i <= n; i++ {
		count += countOnesInNumber(i)
	}
	return count
}

func countOnesInNumber(num int) int {
	count := 0
	for num > 0 {
		if num%10 == 1 {
			count++
		}
		num /= 10
	}
	return count
}

// Digit DP with memoization
func countDigitOneDP(n int) int {
	if n <= 0 {
		return 0
	}
	
	s := fmt.Sprintf("%d", n)
	length := len(s)
	
	// DP table: dp[pos][tight][hasOne]
	dp := make([][][]int, length+1)
	for i := range dp {
		dp[i] = make([]int, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	
	// Initialize
	for i := 0; i < length; i++ {
		for tight := 0; tight < 2; tight++ {
			for hasOne := 0; hasOne < 2; hasOne++ {
				dp[i][tight][hasOne] = -1
			}
		}
	}
	
	// Count from 1 to n-1
	for i := 1; i < n; i++ {
		count += countOnesInNumberDP(fmt.Sprintf("%d", i), dp, length)
	}
	
	// Count for n
	count += countOnesInNumberDP(s, dp, length)
	
	return count
}

func countOnesInNumberDP(s string, dp [][][]int, length int) int {
	if len(s) == 0 {
		return 0
	}
	
	return countOnesInNumberDPHelper(s, 0, 0, false, dp, length)
}

func countOnesInNumberDPHelper(s string, pos int, tight int, hasOne int, dp [][][]int, length int) int {
	if pos == length {
		return hasOne
	}
	
	if dp[pos][tight][hasOne] != -1 {
		return dp[pos][tight][hasOne]
	}
	
	limit := 9
	if tight == 1 {
		limit = int(s[pos] - '0')
	}
	
	total := 0
	for digit := 1; digit <= limit; digit++ {
		newTight := 0
		if digit == limit && tight == 1 {
			newTight = 1
		}
		
		newHasOne := hasOne
		if digit == 1 {
			newHasOne = 1
		}
		
		total += countOnesInNumberDPHelper(s, pos+1, newTight, newHasOne, dp, length)
	}
	
	dp[pos][tight][hasOne] = total
	return total
}

// Digit DP with different base
func countDigitOneBase(n, base int) int {
	if n <= 0 {
		return 0
	}
	
	s := fmt.Sprintf("%d", n)
	length := len(s)
	
	count := 0
	for i := 1; i < length; i++ {
		count += countDigitOneHelperBase(i-1, base)
	}
	
	firstDigit := int(s[0] - '0')
	remaining := n - firstDigit*int(math.Pow(float64(base), float64(length-1)))
	
	if firstDigit == 1 {
		count += remaining + 1
	} else {
		count += int(math.Pow(float64(base), float64(length-1))) + countDigitOneHelperBase(firstDigit-1, base)*int(math.Pow(float64(base), float64(length-1)))
	}
	
	return count
}

func countDigitOneHelperBase(n, base int) int {
	count := 0
	for n > 0 {
		if n%base == 1 {
			count++
		}
		n /= base
	}
	return count
}

// Digit DP for range queries
func countDigitOneInRange(low, high int) int {
	return countDigitOne(high) - countDigitOne(low-1)
}

// Digit DP with preprocessing
func countDigitOnePreprocessed(n int) int {
	if n <= 0 {
		return 0
	}
	
	// Precompute counts for all numbers up to n
	onesCount := make([]int, n+1)
	onesCount[0] = 0
	
	for i := 1; i <= n; i++ {
		onesCount[i] = onesCount[i/10] + countOnesInNumber(i%10)
	}
	
	// Calculate prefix sums
	prefixOnes := make([]int, n+1)
	prefixOnes[0] = 0
	
	for i := 1; i <= n; i++ {
		prefixOnes[i] = prefixOnes[i-1] + onesCount[i]
	}
	
	return prefixOnes[n]
}

// Digit DP with mathematical formula
func countDigitOneMathematical(n int) int {
	if n <= 0 {
		return 0
	}
	
	count := 0
	length := 1
	temp := n
	
	for temp >= 10 {
		length++
		temp /= 10
	}
	
	// Count for numbers with less digits
	for i := 1; i < length; i++ {
		count += int(math.Pow(10, float64(i-1))) * int(math.Pow(8, float64(i-1)))
	}
	
	// Count for numbers with same number of digits
	s := fmt.Sprintf("%d", n)
	firstDigit := int(s[0] - '0')
	remaining := n - firstDigit*int(math.Pow(10, float64(length-1)))
	
	if firstDigit == 1 {
		count += remaining + 1
	} else {
		count += int(math.Pow(10, float64(length-1))) + 
			firstDigit*int(math.Pow(8, float64(length-1))) +
			int(math.Pow(10, float64(length-1))) * (length-1)
	}
	
	return count
}

// Digit DP for multiple ranges
func countDigitOneMultipleRanges(ranges [][2]int) []int {
	results := make([]int, len(ranges))
	
	for i, r := range ranges {
		results[i] = countDigitOneInRange(r[0], r[1])
	}
	
	return results
}

// Digit DP with constraint
func countDigitOneWithConstraint(n, maxOnes int) int {
	if n <= 0 {
		return 0
	}
	
	count := 0
	for i := 1; i <= n; i++ {
		ones := countOnesInNumber(i)
		if ones <= maxOnes {
			count++
		}
	}
	
	return count
}

// Digit DP with binary representation
func countDigitOneBinary(n int) int {
	if n <= 0 {
		return 0
	}
	
	count := 0
	for i := 1; i <= n; i++ {
		binary := fmt.Sprintf("%b", i)
		for _, bit := range binary {
			if bit == '1' {
				count++
				break
			}
		}
	}
	
	return count
}

func main() {
	// Test cases
	fmt.Println("=== Testing Number of Digit One - Digit DP ===")
	
	testCases := []struct {
		n          int
		description string
	}{
		{13, "Standard case"},
		{0, "Zero"},
		{1, "Single digit"},
		{9, "Single digit max"},
		{10, "Two digits"},
		{99, "Two digits max"},
		{100, "Three digits"},
		{999, "Three digits max"},
		{1000, "Four digits"},
		{1111, "All ones"},
		{1234, "Mixed digits"},
		{10000, "Five digits"},
		{100000, "Six digits"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s (n=%d)\n", i+1, tc.description, tc.n)
		
		result1 := countDigitOne(tc.n)
		result2 := countDigitOneDP(tc.n)
		result3 := countDigitOneMathematical(tc.n)
		result4 := countDigitOnePreprocessed(tc.n)
		
		fmt.Printf("  Standard: %d\n", result1)
		fmt.Printf("  DP: %d\n", result2)
		fmt.Printf("  Mathematical: %d\n", result3)
		fmt.Printf("  Preprocessed: %d\n\n", result4)
	}
	
	// Test range queries
	fmt.Println("=== Range Queries Test ===")
	ranges := [][2]int{
		{1, 13},
		{10, 100},
		{100, 1000},
		{1000, 10000},
	}
	
	for i, r := range ranges {
		result := countDigitOneInRange(r[0], r[1])
		fmt.Printf("Range [%d, %d]: %d\n", r[0], r[1], result)
	}
	
	// Test different bases
	fmt.Println("\n=== Different Base Test ===")
	bases := []int{2, 8, 16}
	testN := 100
	
	for _, base := range bases {
		result := countDigitOneBase(testN, base)
		fmt.Printf("Base %d (n=%d): %d\n", base, testN, result)
	}
	
	// Test constraints
	fmt.Println("\n=== Constraint Test ===")
	constraints := []int{1, 2, 3, 5}
	
	for _, maxOnes := range constraints {
		result := countDigitOneWithConstraint(100, maxOnes)
		fmt.Printf("Max ones %d: %d\n", maxOnes, result)
	}
	
	// Test binary representation
	fmt.Println("\n=== Binary Representation Test ===")
	binaryTest := []int{5, 10, 15, 31, 63}
	
	for _, n := range binaryTest {
		result := countDigitOneBinary(n)
		fmt.Printf("n=%d (binary=%b): %d\n", n, n, result)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largeN := 1000000
	fmt.Printf("Large test with n=%d\n", largeN)
	
	result := countDigitOneDP(largeN)
	fmt.Printf("DP result: %d\n", result)
	
	result = countDigitOneMathematical(largeN)
	fmt.Printf("Mathematical result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Negative number
	fmt.Printf("Negative number: %d\n", countDigitOne(-5))
	
	// Very large number
	veryLarge := 1000000000
	fmt.Printf("Very large number: %d\n", countDigitOneMathematical(veryLarge))
	
	// Test with multiple ranges
	fmt.Println("\n=== Multiple Ranges Test ===")
	multipleRanges := [][2]int{
		{1, 10},
		{11, 20},
		{21, 30},
	}
	
	results := countDigitOneMultipleRanges(multipleRanges)
	for i, result := range results {
		fmt.Printf("Range %d: %d\n", i, result)
	}
}
