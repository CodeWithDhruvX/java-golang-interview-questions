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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Stack-Based Expression Evaluation
- **Stack Usage**: Push operands, pop on operators for evaluation
- **Postfix Processing**: Process tokens left to right
- **Operator Application**: Pop two operands, apply operator, push result
- **Final Result**: Single value remains on stack

## 2. PROBLEM CHARACTERISTICS
- **Postfix Notation**: Operators follow operands (Reverse Polish Notation)
- **Stack Evaluation**: Natural fit for postfix expressions
- **Binary Operations**: Each operator takes two operands
- **Integer Arithmetic**: Handle integer division and overflow

## 3. SIMILAR PROBLEMS
- Basic Calculator (LeetCode 224) - Infix expression evaluation
- Evaluate Division (LeetCode 399) - Division with constraints
- Calculator with Parentheses - Handle precedence with stacks
- Expression Add Operators - Insert operators into expression

## 4. KEY OBSERVATIONS
- **Stack Natural Fit**: Postfix evaluation is inherently stack-based
- **Operand Order**: Second popped operand is left operand
- **Operator Precedence**: Not needed in postfix (already encoded)
- **Error Handling**: Need to handle invalid expressions

## 5. VARIATIONS & EXTENSIONS
- **Different Operators**: Support for exponent, modulo, etc.
- **Floating Point**: Handle decimal numbers and operations
- **Variable Length**: Support for multi-digit numbers and variables
- **Error Reporting**: Return specific error types/positions

## 6. INTERVIEW INSIGHTS
- Always clarify: "Operator set? Integer division? Error handling?"
- Edge cases: single number, invalid tokens, division by zero
- Time complexity: O(N) where N=number of tokens
- Space complexity: O(N) for operand stack
- Key insight: stack naturally models postfix evaluation

## 7. COMMON MISTAKES
- Wrong operand order (operand2, operand1 vs operand1, operand2)
- Not handling insufficient operands
- Integer division issues (rounding, negative numbers)
- Not validating operator tokens
- Stack underflow/overflow errors

## 8. OPTIMIZATION STRATEGIES
- **Basic Stack**: O(N) time, O(N) space - optimal
- **Early Validation**: Pre-validate token sequence
- **Operator Optimization**: Use switch for fast operator dispatch
- **Memory Pool**: Reuse stack for multiple evaluations

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like calculating with a calculator that uses stacks:**
- You have a stack-based calculator (like old HP calculators)
- Numbers go on the stack, operations use stack numbers
- Postfix means "number number operation" format
- Like saying "3 4 +" instead of "3 + 4"
- Stack naturally tracks intermediate results

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of tokens in Reverse Polish Notation
2. **Goal**: Evaluate the expression and return result
3. **Format**: Postfix notation (operands before operators)
4. **Output**: Integer result of evaluation

#### Phase 2: Key Insight Recognition
- **"Stack natural fit"** → Postfix evaluation is inherently stack-based
- **"Operand order"** → Need to maintain correct operand order
- **"No precedence needed"** → Postfix already encodes precedence
- **"Error handling"** → Need to validate expression structure

#### Phase 3: Strategy Development
```
Human thought process:
"I need to evaluate postfix expression.
This is perfect for stack usage:

Stack Evaluation Approach:
1. Iterate through tokens left to right
2. If token is number: push to stack
3. If token is operator:
   - Pop two operands (operand2, then operand1)
   - Apply operator: operand1 op operand2
   - Push result back to stack
4. Final result is single value on stack

This naturally handles operator precedence!"
```

#### Phase 4: Edge Case Handling
- **Single number**: Push and return directly
- **Invalid expression**: Handle insufficient operands
- **Division by zero**: Handle or return error
- **Empty input**: Handle gracefully

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Tokens: ["2", "1", "+", "3", "*"] = (2+1)*3 = 9

Human thinking:
"Step 1: '2' → number, push to stack
Stack: [2]

Step 2: '1' → number, push to stack
Stack: [2, 1]

Step 3: '+' → operator, pop two operands
operand2 = 1 (top), operand1 = 2 (next)
result = 2 + 1 = 3
Push result: [3]
Stack: [3]

Step 4: '3' → number, push to stack
Stack: [3, 3]

Step 5: '*' → operator, pop two operands
operand2 = 3 (top), operand1 = 3 (next)
result = 3 * 3 = 9
Push result: [9]
Stack: [9]

Final result: 9 ✓"
```

#### Phase 6: Intuition Validation
- **Why stack works**: Postfix evaluation is inherently stack-based
- **Why operand order matters**: operand1 op operand2 vs operand2 op operand1
- **Why no precedence needed**: Postfix already encodes evaluation order
- **Why O(N)**: Each token processed exactly once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not convert to infix?"** → More complex, loses postfix advantage
2. **"Should I use recursion?"** → Stack is more efficient and natural
3. **"What about operator precedence?"** → Not needed in postfix notation
4. **"Can I optimize further?"** → O(N) is already optimal
5. **"What about different operators?"** → Extend operator set and handling

### Real-World Analogy
**Like using a stack-based calculator or programming language:**
- Many calculators and programming languages use stack evaluation
- Forth and PostScript languages use postfix notation
- Stack naturally tracks intermediate results
- Like evaluating compiled bytecode or assembly
- Each operation consumes inputs and produces output

### Human-Readable Pseudocode
```
function evalRPN(tokens):
    stack = empty stack
    
    for token in tokens:
        if token is number:
            stack.push(parseInt(token))
        else if token is operator:
            if stack.size() < 2:
                return error // insufficient operands
            
            operand2 = stack.pop()
            operand1 = stack.pop()
            
            result = applyOperator(operand1, operand2, token)
            stack.push(result)
        else:
            return error // invalid token
    
    if stack.size() != 1:
        return error // invalid expression
    
    return stack.pop()

function applyOperator(a, b, op):
    switch op:
        case "+": return a + b
        case "-": return a - b
        case "*": return a * b
        case "/": return a / b // handle division by zero
```

### Execution Visualization

### Example: Tokens = ["2", "1", "+", "3", "*"] = (2+1)*3 = 9
```
Token Processing:
Index: 0    1    2      3    4
Token:  2    1    +      3    *
Stack:  [2] [2,1] [3] [3,3] [9]

Step-by-step:
1. '2' → number → push: [2]
2. '1' → number → push: [2,1]
3. '+' → operator:
   - operand2 = 1, operand1 = 2
   - result = 2 + 1 = 3
   - push result: [3]
4. '3' → number → push: [3,3]
5. '*' → operator:
   - operand2 = 3, operand1 = 3
   - result = 3 * 3 = 9
   - push result: [9]

Final Result: 9 ✓
```

### Key Visualization Points:
- **Number Processing**: Push operands directly to stack
- **Operator Processing**: Pop two, apply, push result
- **Stack Evolution**: Grows with numbers, shrinks with operations
- **Final State**: Single result remains

### Memory Layout Visualization:
```
Stack State Evolution:
Initial:  []
After '2':  [2]
After '1':  [2, 1]
After '+':  [3]        (2+1=3, popped 2,1)
After '3':  [3, 3]
After '*':  [9]        (3*3=9, popped 3,3)

Operator Application:
+ : operand1=2, operand2=1, result=3
* : operand1=3, operand2=3, result=9

Evaluation Order:
1. Process tokens left to right
2. Numbers: push to stack
3. Operators: pop two, apply, push result
4. Final: single value on stack
```

### Time Complexity Breakdown:
- **Token Processing**: O(N) time where N=number of tokens
- **Stack Operations**: O(1) per push/pop operation
- **Operator Application**: O(1) per operation
- **Total**: O(N) time, O(N) space in worst case
- **Optimal**: Cannot do better than processing each token

### Alternative Approaches:

#### 1. Recursive Evaluation (O(N) time, O(N) space)
```go
func evalRPNRecursive(tokens []string) int {
    // Recursively evaluate from end
    // More complex than stack approach
    // ... implementation details omitted
}
```
- **Pros**: Elegant for some variations
- **Cons**: More complex, same asymptotic complexity

#### 2. Two-Pass Evaluation (O(N) time, O(N) space)
```go
func evalRPNTwoPass(tokens []string) int {
    // First pass: identify operation boundaries
    // Second pass: evaluate in chunks
    // More complex than stack approach
    // ... implementation details omitted
}
```
- **Pros**: Can provide more detailed error info
- **Cons**: More complex, same complexity

#### 3. Tree-Based Evaluation (O(N) time, O(N) space)
```go
func evalRPNTree(tokens []string) int {
    // Build expression tree, then evaluate
    // Overkill for simple evaluation
    // ... implementation details omitted
}
```
- **Pros**: Good for complex expressions with precedence
- **Cons**: Overkill for simple postfix evaluation

### Extensions for Interviews:
- **Different Operators**: Support for exponent, modulo, etc.
- **Floating Point**: Handle decimal numbers and operations
- **Error Handling**: Return specific error types/positions
- **Variable Support**: Handle variables in expressions
- **Performance Analysis**: Discuss worst-case scenarios
*/
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
