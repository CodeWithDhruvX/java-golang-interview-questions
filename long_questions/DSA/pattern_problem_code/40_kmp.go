package main

import (
	"fmt"
)

// Pattern: KMP Algorithm
// Difficulty: Hard
// Key Concept: String matching in O(N+M) by preprocessing the pattern to know how much to shift upon mismatch (LPS Array).

/*
INTUITION:
Naive string matching is O(N*M). If we match "AAAA" but fail at "B", we normally shift by 1.
But if we know "A" matches "A", maybe we can shift more?

KMP Concepts:
- **Prefix**: substring starting at 0.
- **Suffix**: substring ending at end.
- **LPS[i]**: Length of longest proper prefix of substring `pat[0...i]` that is also a suffix of `pat[0...i]`.

Example: "ABABC"
i=0 ("A"): LPS=0
i=1 ("AB"): LPS=0
i=2 ("ABA"): Prefix "A", Suffix "A" -> LPS=1.
i=3 ("ABAB"): Prefix "AB", Suffix "AB" -> LPS=2.
i=4 ("ABABC"): LPS=0.

When matching `txt` and `pat`:
- If `txt[i] == pat[j]`: i++, j++.
- If match: return start index. continue with `j = lps[j-1]`.
- If mismatch after `j` matches:
  - Don't reset `i`!
  - `j = lps[j-1]`. (Jump back to the matching prefix).
  - If j=0, just i++.

PROBLEM:
LeetCode 28. Find the Index of the First Occurrence in a String (Implement strStr).

ALGORITHM:
1. `computeLPS(pat)`:
   - `len = 0`. `i = 1`.
   - Loop `i < M`:
     - If `pat[i] == pat[len]`: len++. lps[i] = len. i++.
     - Else:
       - If `len != 0`: `len = lps[len-1]`.
       - Else: `lps[i] = 0`. i++.
2. `search(txt, pat)`:
   - `i = 0` (txt), `j = 0` (pat).
   - Loop `i < N`:
     - If `pat[j] == txt[i]`: i++, j++.
     - If `j == M`: FOUND at `i-j`. `j = lps[j-1]`.
     - Else if `i < N` and `pat[j] != txt[i]`:
       - If `j != 0`: `j = lps[j-1]`.
       - Else: i++.
*/

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	n, m := len(haystack), len(needle)
	lps := computeLPS(needle)

	i, j := 0, 0
	for i < n {
		if haystack[i] == needle[j] {
			i++
			j++
		}

		if j == m {
			return i - j
		} else if i < n && haystack[i] != needle[j] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}

func computeLPS(pat string) []int {
	m := len(pat)
	lps := make([]int, m)
	lenVal := 0
	i := 1

	for i < m {
		if pat[i] == pat[lenVal] {
			lenVal++
			lps[i] = lenVal
			i++
		} else {
			if lenVal != 0 {
				lenVal = lps[lenVal-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func main() {
	// txt: ABABDABACDABABCABAB
	// pat: ABABCABAB
	// Match at end?

	txt := "ABABDABACDABABCABAB"
	pat := "ABABCABAB"
	fmt.Printf("Txt: %s, Pat: %s, Index: %d\n", txt, pat, strStr(txt, pat))

	// Simple
	fmt.Printf("sadbutsad, sad -> %d\n", strStr("sadbutsad", "sad"))
	fmt.Printf("leetcode, leeto -> %d\n", strStr("leetcode", "leeto"))
}
