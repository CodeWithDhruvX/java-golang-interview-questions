package main

import "fmt"

// 371. Sum of Two Integers
// Time: O(1), Space: O(1)
func getSum(a int, b int) int {
	for b != 0 {
		// Calculate carry bits
		carry := a & b
		
		// Calculate sum bits without carry
		a = a ^ b
		
		// Shift carry to the left
		b = carry << 1
	}
	
	return a
}

// Recursive version
func getSumRecursive(a int, b int) int {
	if b == 0 {
		return a
	}
	
	// Calculate carry and sum
	carry := a & b
	sum := a ^ b
	
	// Recursively add carry to sum
	return getSumRecursive(sum, carry<<1)
}

// Step-by-step explanation version
func getSumWithExplanation(a int, b int) int {
	fmt.Printf("Adding %d and %d using bit manipulation:\n", a, b)
	fmt.Printf("Initial: a=%d (%032b), b=%d (%032b)\n\n", a, b, a, b)
	
	step := 1
	for b != 0 {
		fmt.Printf("Step %d:\n", step)
		
		// Calculate carry (where both bits are 1)
		carry := a & b
		fmt.Printf("  Carry (a & b): %d (%032b)\n", carry, carry)
		
		// Calculate sum without carry (XOR)
		a = a ^ b
		fmt.Printf("  Sum without carry (a ^ b): %d (%032b)\n", a, a)
		
		// Shift carry left to add in next higher bit position
		b = carry << 1
		fmt.Printf("  Shifted carry (carry << 1): %d (%032b)\n", b, b)
		fmt.Printf("  New a: %d, New b: %d\n\n", a, b)
		
		step++
	}
	
	fmt.Printf("Final result: %d\n", a)
	return a
}

func main() {
	// Test cases
	testCases := [][2]int{
		{1, 2},
		{2, 3},
		{0, 1},
		{-1, 1},
		{-2, -3},
		{10, 20},
		{100, 200},
		{-10, 20},
		{123, 456},
		{999, 1},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		
		if i == 0 {
			// Show detailed explanation for first test case
			result := getSumWithExplanation(tc[0], tc[1])
			fmt.Printf("Iterative result: %d\n\n", result)
		} else {
			result1 := getSum(tc[0], tc[1])
			result2 := getSumRecursive(tc[0], tc[1])
			fmt.Printf("Adding %d + %d\n", tc[0], tc[1])
			fmt.Printf("Iterative: %d, Recursive: %d\n\n", result1, result2)
		}
	}
}
