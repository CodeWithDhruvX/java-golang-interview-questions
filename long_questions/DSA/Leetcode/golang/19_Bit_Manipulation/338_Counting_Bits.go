package main

import "fmt"

// 338. Counting Bits
// Time: O(N), Space: O(N)
func countBits(n int) []int {
	result := make([]int, n+1)
	
	for i := 1; i <= n; i++ {
		// Brian Kernighan's algorithm
		count := 0
		num := i
		for num > 0 {
			num &= num - 1 // Clear the least significant bit set
			count++
		}
		result[i] = count
	}
	
	return result
}

// Optimized DP approach: O(N) time, O(N) space
func countBitsDP(n int) []int {
	result := make([]int, n+1)
	
	for i := 1; i <= n; i++ {
		// i >> 1 removes the least significant bit
		// i & 1 gives the least significant bit (0 or 1)
		result[i] = result[i>>1] + (i & 1)
	}
	
	return result
}

// Most optimized DP approach
func countBitsDPOptimized(n int) []int {
	result := make([]int, n+1)
	
	for i := 1; i <= n; i++ {
		// i & (i-1) clears the least significant set bit
		result[i] = result[i&(i-1)] + 1
	}
	
	return result
}

// Function to count bits in a single number
func countBitsInNumber(num int) int {
	count := 0
	for num > 0 {
		count += num & 1
		num >>= 1
	}
	return count
}

// Built-in function approach (Go specific)
func countBitsBuiltIn(n int) []int {
	result := make([]int, n+1)
	
	for i := 0; i <= n; i++ {
		// Go doesn't have built-in popcount like C++, but we can use bit manipulation
		count := 0
		num := i
		for num > 0 {
			num &= num - 1
			count++
		}
		result[i] = count
	}
	
	return result
}

func main() {
	// Test cases
	testCases := []int{
		2, 5, 10, 15, 0, 1, 7, 8, 16, 31,
	}
	
	for i, n := range testCases {
		result1 := countBits(n)
		result2 := countBitsDP(n)
		result3 := countBitsDPOptimized(n)
		result4 := countBitsBuiltIn(n)
		
		fmt.Printf("Test Case %d: n=%d\n", i+1, n)
		fmt.Printf("  Brian Kernighan: %v\n", result1)
		fmt.Printf("  DP (right shift): %v\n", result2)
		fmt.Printf("  DP (optimized): %v\n", result3)
		fmt.Printf("  Built-in style: %v\n\n", result4)
	}
}
