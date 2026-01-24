package main

import "fmt"

// Pattern: Bit Manipulation
// Difficulty: Easy
// Key Concept: Using logical XOR to cancel out duplicates or extract stats without memory.

/*
INTUITION:
"Single Number"
Given a non-empty array of integers where every element appears twice except for one. Find that single one.
Constraint: O(N) time and O(1) space.

Naive approach: Use a Map (O(N) space) or Sort (O(N log N) time).
We need O(N) Time + O(1) Space. Magic?
No, Logic.

Recall XOR (^) properties:
1. x ^ 0 = x (Identity)
2. x ^ x = 0 (Self-Inverse)
3. a ^ b ^ a = (a ^ a) ^ b = 0 ^ b = b (Commutative)

If we XOR *all* the numbers in the array:
[A, B, A, C, B]
= A ^ B ^ A ^ C ^ B
= (A ^ A) ^ (B ^ B) ^ C
= 0 ^ 0 ^ C
= C.

The duplicates cancel each other out! Only the unique number survives.

ALGORITHM:
1. Init `result = 0`.
2. Loop through array, XORing every number with `result`.
3. Return `result`.
*/

func singleNumber(nums []int) int {
	result := 0

	// DRY RUN: [4, 1, 2, 1, 2]
	// Init: 0
	// 1. 0 ^ 4 = 4 (Bin: 100)
	// 2. 4 ^ 1 = 5 (Bin: 101)
	// 3. 5 ^ 2 = 7 (Bin: 111)
	// 4. 7 ^ 1 = 6 (Bin: 110) -- Notice logic: 4^1^2^1. The 1s killed each other. Effectively 4^2.
	// 5. 6 ^ 2 = 4 (Bin: 100) -- The 2s killed each other. Effectively 4.
	// Return 4.

	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Printf("Input: %v\n", nums)
	res := singleNumber(nums)
	fmt.Printf("Single Number: %d\n", res) // Expected: 4
}
