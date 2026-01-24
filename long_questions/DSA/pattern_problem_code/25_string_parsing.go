package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Pattern: String Parsing / Decoding
// Difficulty: Medium
// Key Concept: Using a Stack (or Recursion) to handle nested structures like brackets [].

/*
INTUITION:
"Decode String"
Input: "3[a]2[bc]" -> "aaabcbc"
Input: "3[a2[c]]" -> "accaccacc"

When we see "3[", we know we need to repeat something 3 times. But we don't know *what* yet until we find the closing "]".
This "Last In, First Out" processing of nested brackets screams STACK.

We need two stacks (or one stack of structs):
1. `countStack`: To remember "Repeat 3 times".
2. `stringStack`: To remember "What came before this bracket" (e.g. the "a" in "a2[c]").

ALGORITHM:
Iterate char by char:
1. If Digit: Build the number (it might be multi-digit like "12").
2. If '[':
   - Push current number to `countStack`.
   - Push current built string to `stringStack`.
   - Reset number and current string. (Enter new context).
3. If ']':
   - Pop count `k` from `countStack`.
   - Pop previous string `prev` from `stringStack`.
   - `currentString = prev + (currentString * k)`.
4. If Letter: Append to `currentString`.
*/

func decodeString(s string) string {
	countStack := []int{}
	stringStack := []string{}
	currentString := ""
	currentNum := 0

	// DRY RUN: "3[a2[c]]"
	//
	// '3': currNum = 3.
	// '[': Push 3 (Count), Push "" (Str). Reset.
	//        CS: [3], SS: [""]
	// 'a': currStr = "a".
	// '2': currNum = 2.
	// '[': Push 2, Push "a". Reset.
	//        CS: [3, 2], SS: ["", "a"]
	// 'c': currStr = "c".
	// ']': Pop k=2. Pop prev="a".
	//        currStr = "a" + "c"*2 = "acc".
	// ']': Pop k=3. Pop prev="".
	//        currStr = "" + "acc"*3 = "accaccacc".
	// Result: "accaccacc".

	for i := 0; i < len(s); i++ {
		char := s[i]

		if unicode.IsDigit(rune(char)) {
			num, _ := strconv.Atoi(string(char))
			currentNum = currentNum*10 + num
		} else if char == '[' {
			countStack = append(countStack, currentNum)
			stringStack = append(stringStack, currentString)
			currentNum = 0
			currentString = ""
		} else if char == ']' {
			// Pop Count
			k := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			// Pop Previous String
			prevStr := stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]

			// Decode: Prev + (Curr * k)
			decodedPart := strings.Repeat(currentString, k)
			currentString = prevStr + decodedPart
		} else {
			// Regular character
			currentString += string(char)
		}
	}

	return currentString
}

func main() {
	input := "3[a2[c]]"
	fmt.Printf("Input: %s\n", input)

	res := decodeString(input)
	fmt.Printf("Decoded: %s\n", res) // Expected: accaccacc
}
