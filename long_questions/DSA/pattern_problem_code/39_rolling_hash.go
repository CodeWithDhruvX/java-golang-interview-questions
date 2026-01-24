package main

import (
	"fmt"
)

// Pattern: Rolling Hash (Rabin-Karp)
// Difficulty: Hard
// Key Concept: Computing hashes of substrings in O(1) by reusing the hash of the previous substring (Rolling).

/*
INTUITION:
We want to find if a substring "abc" exists in a long text.
Comparing "abc" at every position is O(N*M).
If we convert "abc" to a number (Hash), we can compare numbers in O(1).
Hash("abc") = 'a'*26^2 + 'b'*26^1 + 'c'*26^0.

When moving window from "abc" to "bcd":
- Remove 'a' (leading term).
- Shift left (multiply by 26).
- Add 'd'.
NewHash = ((OldHash - Val('a') * 26^(L-1)) * 26 + Val('d')) % Mod.

PROBLEM:
LeetCode 1044. Longest Duplicate Substring. (Or simpler: Repeated DNA Sequences).
We will implement "Repeated DNA Sequences" (Find all 10-letter substrings that occur more than once) as it's cleaner.
For Longest Duplicate Substring, we would use Binary Search on Length + Rolling Hash.

ALGORITHM (Repeated DNA Sequences):
1. Base = 26 (or 4 for DNA), Mod = 10^9 + 7 (or 2^64 using unsigned int).
2. Calculate hash of first L=10 chars.
3. Store in map.
4. Loop from i=1 to N-L:
   - Update hash using rolling formula.
   - Check if hash in map. If yes, found duplicate.
*/

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	seen := make(map[int]int) // Hash -> Count
	res := make([]string, 0)
	added := make(map[string]bool)

	// Rolling Hash Parameters
	base := 4
	mod := 1_000_000_007

	// Map chars to 0-3
	toInt := func(c byte) int {
		switch c {
		case 'A':
			return 0
		case 'C':
			return 1
		case 'G':
			return 2
		case 'T':
			return 3
		}
		return 0
	}

	// Calculate 4^9 for removing leading char
	// L = 10. Max power is 4^9.
	pow := 1
	L := 10
	for i := 0; i < L-1; i++ {
		pow = (pow * base) % mod
	}

	// Initial Hash
	currHash := 0
	for i := 0; i < L; i++ {
		currHash = (currHash*base + toInt(s[i])) % mod
	}
	seen[currHash] = 1

	for i := 1; i <= len(s)-L; i++ {
		// Remove leading char at s[i-1]
		prevCharVal := toInt(s[i-1])
		currHash = (currHash - (prevCharVal*pow)%mod + mod) % mod

		// Shift and Add new char at s[i+L-1]
		currHash = (currHash * base) % mod
		newCharVal := toInt(s[i+L-1])
		currHash = (currHash + newCharVal) % mod

		if seen[currHash] > 0 {
			sub := s[i : i+L]
			if !added[sub] {
				res = append(res, sub)
				added[sub] = true
			}
		}
		seen[currHash]++
	}

	return res
}

func main() {
	// AAAAACCCCC AAAAACCCCC CAAAAAGGGT
	// Repeated: AAAAACCCCC
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGT"
	fmt.Printf("String: %s\n", s)
	fmt.Printf("Repeated: %v\n", findRepeatedDnaSequences(s))
}
