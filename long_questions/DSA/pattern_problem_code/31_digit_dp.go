package main

import (
	"fmt"
	"strconv"
)

// Pattern: Digit DP
// Difficulty: Hard
// Key Concept: Constructing numbers digit by digit while maintaining constraints (tight, leading zeros, etc.) to count valid numbers in range [0, N].

/*
INTUITION:
We want to count how many numbers between 0 and N satisfy a property (e.g., "contain the digit 1").
Iterating from 0 to N (where N=10^9) is O(N) -> too slow.
Digit DP builds the number from left to right (most significant digit first).

State usually involves:
- `index`: converting position (0 to len(S)-1).
- `tight`: boolean. Are we restricted by the digits of N?
  - If N=543 and we place '4' at first position, we are 'loose' for next positions (can place 0-9).
  - If we place '5', we are 'tight' (next digit must be 0-4).
- `count` or `sum`: The property we are tracking (e.g. how many 1s so far).
- `leadingZeros`: sometimes needed.

Complexity: O(Digits * 10 * State). For N=10^9, Digits=10. This is O(1) effectively.

PROBLEM:
LeetCode 233. Number of Digit One.
Given an integer n, count the total number of digit 1 appearing in all non-negative integers less than or equal to n.
Example: n=13. Numbers: 1, 10, 11, 12, 13.
1(1), 10(1), 11(2), 12(1), 13(1). Total = 6.

ALGORITHM:
1. Convert N to string `s`.
2. DFS(index, tight, count1s).
   - `index`: current digit position (0 to len-1).
   - `tight`: true if we are restricted by N's digits.
   - `count1s`: count of 1s placed so far.
   - Return: total count of 1s in all completed numbers formed from this state.
   - Wait: The problem asks for total Occurrences of 1.
   - Slightly different state: Instead of carrying `count1s`, we can return Pair(count_numbers, count_ones).
   - Or standard approach: `dfs` returns total Ones contributed by this suffix.

   Standard DP structure for "Total Sum of Digit 1s":
   - `memo[index][tight][sum]`
   - Let `limit = tight ? s[index] - '0' : 9`
   - Loop `d` from 0 to `limit`.
   - `update_tight = tight && (d == limit)`
   - `next_sum = sum + (1 if d == 1 else 0)`
   - `res += dfs(index+1, update_tight, next_sum)`
   - Takes too much space if `sum` is large? No, max sum is 9*digits. Small.

   Optimization:
   Since we sum up 1s, we can just track:
   - How many valid numbers can be formed?
   - How many 1s in those numbers?
   - Actually simpler:
     State: `(index, tight, count_of_1s_so_far)`
     Return: Total 1s.
     Max 1s = 10.
     But standard trick: We don't carry `count_of_1s_so_far` in state if we want better caching.
     We carry it, but memoize `(index, tight, count)`.
*/

// Memoization table: [index][tight][count]
// tight is 0 or 1.
// index is 0..10.
// count is 0..10.
var memo [12][2][12]int

func countDigitOne(n int) int {
	s := strconv.Itoa(n)
	// Reset memo
	for i := 0; i < len(memo); i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < len(memo[i][j]); k++ {
				memo[i][j][k] = -1
			}
		}
	}
	return dfs(0, true, 0, s)
}

func dfs(index int, tight bool, count int, s string) int {
	if index == len(s) {
		return count
	}

	tVal := 0
	if tight {
		tVal = 1
	}

	if memo[index][tVal][count] != -1 {
		return memo[index][tVal][count]
	}

	limit := 9
	if tight {
		limit = int(s[index] - '0')
	}

	total := 0
	for digit := 0; digit <= limit; digit++ {
		newTight := tight && (digit == limit)
		newCount := count
		if digit == 1 {
			newCount++
		}
		total += dfs(index+1, newTight, newCount, s)
	}

	memo[index][tVal][count] = total
	return total
}

func main() {
	// N=13
	// 1, 10, 11, 12, 13
	// 1s: 1+1+2+1+1 = 6.
	n1 := 13
	fmt.Printf("N: %d, Count of 1s: %d\n", n1, countDigitOne(n1))

	// N=0
	n2 := 0
	fmt.Printf("N: %d, Count of 1s: %d\n", n2, countDigitOne(n2))

	// N=100
	// 1s in 1-99: 20. (1, 10-19(11), 21,31...91(8). Total 1 + 11 + 8 = 20)
	// 100 has one 1. Total 21.
	// Wait logic check:
	// 1s in 00-99. 1s in units place: 1, 11, .. 91 (10 times). 1s in tens place: 10-19 (10 times). Total 20.
	// 100 adds 1. = 21.
	n3 := 100
	fmt.Printf("N: %d, Count of 1s: %d\n", n3, countDigitOne(n3))
}
