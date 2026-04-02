package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 394. Decode String
// Time: O(N), Space: O(N)
func decodeString(s string) string {
	stack := []string{}
	currentNum := 0
	currentStr := ""
	
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		
		switch {
		case char >= "0" && char <= "9":
			// Build the number
			num, _ := strconv.Atoi(char)
			currentNum = currentNum*10 + num
		case char == "[":
			// Push current string and number to stack
			stack = append(stack, currentStr)
			stack = append(stack, fmt.Sprintf("%d", currentNum))
			currentStr = ""
			currentNum = 0
		case char == "]":
			// Pop number and previous string
			numStr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prevStr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			
			num, _ := strconv.Atoi(numStr)
			currentStr = prevStr + strings.Repeat(currentStr, num)
		default:
			// Add character to current string
			currentStr += char
		}
	}
	
	return currentStr
}

func main() {
	// Test cases
	testCases := []string{
		"3[a]2[bc]",
		"3[a2[c]]",
		"2[abc]3[cd]ef",
		"abc3[cd]xyz",
		"10[a]",
		"2[3[a]b]",
		"3[z]2[2[y]pq4[2[jk]e1[f]]]ef",
		"",
		"a",
		"1[a]",
		"100[b]",
		"2[2[2[b]]]",
		"3[a]2[b]1[c]",
	}
	
	for i, s := range testCases {
		result := decodeString(s)
		fmt.Printf("Test Case %d: \"%s\" -> Decoded: \"%s\"\n", i+1, s, result)
	}
}
