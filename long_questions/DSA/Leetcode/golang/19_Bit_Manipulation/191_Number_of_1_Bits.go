package main

import "fmt"

// 191. Number of 1 Bits
// Time: O(1) for 32-bit integers, Space: O(1)
func hammingWeight(num uint32) uint32 {
	count := uint32(0)
	
	// Brian Kernighan's algorithm
	for num != 0 {
		num &= num - 1 // Clear the least significant bit set
		count++
	}
	
	return count
}

// Simple loop approach
func hammingWeightSimple(num uint32) uint32 {
	count := uint32(0)
	
	for i := 0; i < 32; i++ {
		count += num & 1
		num >>= 1
	}
	
	return count
}

// Optimized with bit tricks
func hammingWeightOptimized(num uint32) uint32 {
	num = num - ((num >> 1) & 0x55555555)
	num = (num & 0x33333333) + ((num >> 2) & 0x33333333)
	num = (num + (num >> 4)) & 0x0F0F0F0F
	num = num + (num >> 8)
	num = num + (num >> 16)
	
	return num & 0x3F
}

// Built-in function style (if available)
func hammingWeightBuiltIn(num uint32) uint32 {
	// In Go, we can use bits.OnesCount from Go 1.9+
	// But implementing manually for educational purposes
	count := uint32(0)
	for num != 0 {
		num &= num - 1
		count++
	}
	return count
}

func main() {
	// Test cases
	testCases := []uint32{
		0b00000000000000000000000000000001011, // 3
		0b00000000000000000000000000010000000, // 1
		0b11111111111111111111111111111111101, // 31
		0,                                    // 0
		0b10000000000000000000000000000000000, // 1
		0b11111111111111111111111111111111111, // 32
		0b00000000000000000000000000000000001, // 1
		0b00000000000000000000000000000001010, // 2
		0b10101010101010101010101010101010101, // 16
		0b11001100110011001100110011001100110, // 16
	}
	
	for i, num := range testCases {
		result1 := hammingWeight(num)
		result2 := hammingWeightSimple(num)
		result3 := hammingWeightOptimized(num)
		result4 := hammingWeightBuiltIn(num)
		
		fmt.Printf("Test Case %d: %032b (%d)\n", i+1, num, num)
		fmt.Printf("  Brian Kernighan: %d\n", result1)
		fmt.Printf("  Simple loop: %d\n", result2)
		fmt.Printf("  Bit tricks: %d\n", result3)
		fmt.Printf("  Built-in style: %d\n\n", result4)
	}
}
