package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 150. Evaluate Reverse Polish Notation
// Time: O(N), Space: O(N)
func evalRPN(tokens []string) int {
	stack := []int{}
	
	for _, token := range tokens {
		if isOperator(token) {
			// Pop two operands
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			
			// Perform operation
			result := performOperation(operand1, operand2, token)
			stack = append(stack, result)
		} else {
			// Convert token to integer and push to stack
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	
	return stack[0]
}

// Helper function to check if token is an operator
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

// Helper function to perform arithmetic operation
func performOperation(operand1, operand2 int, operator string) int {
	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2 // Integer division as per problem requirements
	default:
		return 0
	}
}

func main() {
	// Test cases
	testCases := [][]string{
		{"2", "1", "+", "3", "*"},
		{"4", "13", "5", "/", "+"},
		{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
		{"3", "4", "+"},
		{"5", "1", "2", "+", "4", "*", "+", "3", "-"},
		{"2", "3", "11", "+", "*", "5", "-"},
		{"4", "-2", "/", "2", "-3", "*", "-"},
		{"10"},
		{"6", "0", "/"},
		{"-4", "2", "/"},
		{"1", "1", "-", "1", "1", "-"},
	}
	
	for i, tokens := range testCases {
		result := evalRPN(tokens)
		fmt.Printf("Test Case %d: %v -> Result: %d\n", 
			i+1, strings.Join(tokens, " "), result)
	}
}
