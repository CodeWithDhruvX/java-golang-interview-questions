package main

import (
	"fmt"
)

// Pattern: Z-Algorithm
// Difficulty: Hard
// Key Concept: String matching using Z-array. Z[i] is the length of the longest substring starting from S[i] which is also a prefix of S.

/*
INTUITION:
Like KMP, Z-algorithm avoids re-scanning characters.
It maintains a "Z-box" `[L, R]` which is the segment of `S` matching the prefix.
If current index `i` is inside the box (`i <= R`):
- We can copy value from the corresponding prefix index `k = i - L`.
- `Z[i] = min(R - i + 1, Z[k])`.
- If the copied value touches the boundary `R`, we might be able to extend the box further (Naive scan).
If `i > R`, we reset the box to `i` and Naive scan.

PROBLEM:
Same as KMP: LeetCode 28. Find the Index of First Occurrence.
We concatenate `P + "$" + T`.
We compute Z-array.
If `Z[i] == len(P)`, then pattern occurs at that position.

ALGORITHM:
1. Construct string `S = pat + "$" + text`.
2. Compute `Z` array for `S`.
3. Loop through `Z` array starting after `$`.
4. If `Z[i] == len(pat)`, return index converted to text coordinates.
*/

func strStrZ(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	// Separator should be a char not present in text/pattern usually, but for LeetCode simple strings '$' is fine.
	concat := needle + "$" + haystack
	z := calculateZ(concat)

	m := len(needle)
	n := len(haystack)

	for i := 0; i < n; i++ {
		// Index in Z array corresponding to haystack[i] is i + m + 1
		if z[i+m+1] == m {
			return i
		}
	}

	return -1
}

func calculateZ(s string) []int {
	n := len(s)
	z := make([]int, n)
	L, R := 0, 0

	for i := 1; i < n; i++ {
		if i <= R {
			// Inside the Z-box
			// Corresponding index in prefix
			k := i - L
			if z[k] < R-i+1 {
				z[i] = z[k]
			} else {
				// Touches boundary, try to extend
				z[i] = R - i + 1
				for i+z[i] < n && s[z[i]] == s[i+z[i]] {
					z[i]++
				}
				// Update box
				if i+z[i]-1 > R {
					L = i
					R = i + z[i] - 1
				}
			}
		} else {
			// Outside the Z-box, naive scan
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				z[i]++
			}
			if z[i] > 0 {
				L = i
				R = i + z[i] - 1
			}
		}
	}
	return z
}

func main() {
	// txt: hello
	// pat: ll
	// S: ll$hello.
	// Z: [0(l), 1(l), 0($), 0(h), 0(e), 2(l), 0(l), 0(o)]
	// Z[5] is 2 == len(pat). Index in txt is 5 - (2+1) = 2.

	txt := "hello"
	pat := "ll"
	fmt.Printf("Index: %d\n", strStrZ(txt, pat))

	txt2 := "aaaaa"
	pat2 := "bba"
	fmt.Printf("Index: %d\n", strStrZ(txt2, pat2)) // -1
}
