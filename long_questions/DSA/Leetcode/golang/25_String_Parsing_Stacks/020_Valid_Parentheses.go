package main

import "fmt"

// 20. Valid Parentheses
// Time: O(N), Space: O(N)
func isValid(s string) bool {
	stack := []rune{}
	
	// Map closing brackets to opening brackets
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	
	for _, char := range s {
		// If it's an opening bracket, push to stack
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			// If stack is empty or top doesn't match, return false
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			// Invalid character
			return false
		}
	}
	
	// Stack should be empty for valid string
	return len(stack) == 0
}

func main() {
	// Test cases
	testCases := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"",
		"(",
		")",
		"({[]})",
		"({[)]})",
		"((()))",
		"{{{{}}}}",
		"({})",
		"[({})]",
		"([{}])",
		"([{}]))",
	}
	
	for i, s := range testCases {
		result := isValid(s)
		fmt.Printf("Test Case %d: \"%s\" -> Valid: %t\n", i+1, s, result)
	}
}
