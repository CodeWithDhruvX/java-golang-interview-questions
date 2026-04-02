package main

import (
	"fmt"
	"strconv"
)

// 224. Basic Calculator
// Time: O(N), Space: O(N)
func calculate(s string) int {
	stack := []int{}
	result := 0
	sign := 1
	num := 0
	
	for i := 0; i < len(s); i++ {
		char := s[i]
		
		switch {
		case char >= '0' && char <= '9':
			// Build the number
			num = num*10 + int(char-'0')
		case char == '+':
			// Add previous number to result
			result += sign * num
			sign = 1
			num = 0
		case char == '-':
			// Add previous number to result
			result += sign * num
			sign = -1
			num = 0
		case char == '(':
			// Push result and sign to stack
			stack = append(stack, result)
			stack = append(stack, sign)
			// Reset for new sub-expression
			result = 0
			sign = 1
			num = 0
		case char == ')':
			// Add previous number to result
			result += sign * num
			// Pop sign and previous result
			prevSign := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prevResult := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// Calculate new result
			result = prevResult + prevSign*result
			num = 0
		// Ignore spaces
		case char == ' ':
			// Continue
		}
	}
	
	// Add the last number
	result += sign * num
	return result
}

func main() {
	// Test cases
	testCases := []string{
		"1 + 1",
		" 2-1 + 2 ",
		"(1+(4+5+2)-3)+(6+8)",
		"1-(5)",
		"(7)-(0)+(4)",
		"42",
		"    ",
		"1+2-3",
		"(1)",
		"((2+3)*4)",
		"1-(1+1)",
		"0",
		"- (3 + (2 - 1))",
		"3+2*2", // This won't work as this solution doesn't handle multiplication
	}
	
	for i, s := range testCases {
		result := calculate(s)
		fmt.Printf("Test Case %d: \"%s\" -> Result: %d\n", i+1, s, result)
	}
}
