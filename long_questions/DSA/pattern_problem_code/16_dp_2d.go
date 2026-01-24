package main

import "fmt"

// Pattern: Dynamic Programming (2D)
// Difficulty: Medium
// Key Concept: Solving problems with *two* changing variables (e.g., indices of two strings) using a Grid table.

/*
INTUITION:
"Longest Common Subsequence" (LCS)
Text1: "abcde"
Text2: "ace"
Common: "a", "c", "e". Length = 3.

Imagine a Grid where rows are chars of Text1 and cols are chars of Text2.
- Cell (i, j) asks: "What is the LCS length using `text1[0..i]` and `text2[0..j]`?"
- If chars match (`text1[i] == text2[j]`):
  - We found a match! Add 1 to whatever the result was BEFORE these two characters (Diagonal top-left).
  - `dp[i][j] = 1 + dp[i-1][j-1]`
- If chars DON'T match:
  - We throw away either char i or char j and take the best result.
  - `dp[i][j] = max(dp[i-1][j], dp[i][j-1])`

PROBLEM:
Given two strings text1 and text2, return the length of their longest common subsequence.

ALGORITHM:
1. Make a 2D array `dp` of size `(len1+1) x (len2+1)`. (Use +1 for empty string base case).
2. Loop `i` from 1 to len1.
3. Loop `j` from 1 to len2.
4. Apply logic above.
5. Return `dp[len1][len2]`.
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)

	// Create 2D DP Array
	// Go initializes int arrays with 0, which is our correct Base Case (Empty string LCS is 0).
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	// Double Loop
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			// Careful with string indexing: i=1 means char at index 0
			char1 := text1[i-1]
			char2 := text2[j-1]

			if char1 == char2 {
				// Characters match!
				// Take result from [i-1][j-1] and add 1
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				// No match. Inherit the best from top or left.
				top := dp[i-1][j]
				left := dp[i][j-1]
				if top > left {
					dp[i][j] = top
				} else {
					dp[i][j] = left
				}
			}
		}
	}

	// The bottom-right cell contains the answer for the full strings
	return dp[len1][len2]
}

func main() {
	t1 := "abcde"
	t2 := "ace"

	fmt.Printf("Strings: %s, %s\n", t1, t2)
	res := longestCommonSubsequence(t1, t2)
	fmt.Printf("LCS Length: %d\n", res) // Expected: 3

	/*
		Dry Run Grid Visualization:
		      ""  a  c  e
		""    0  0  0  0
		a     0  1  1  1  (Matches 'a')
		b     0  1  1  1  (Inherits 1)
		c     0  1  2  2  (Matches 'c')
		d     0  1  2  2  (Inherits 2)
		e     0  1  2  3  (Matches 'e')
	*/
}
