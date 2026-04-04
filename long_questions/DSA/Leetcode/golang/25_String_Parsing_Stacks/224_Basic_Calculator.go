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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Stack-Based Infix Expression Evaluation
- **Stack Usage**: Push results and signs for parentheses handling
- **Number Building**: Parse multi-digit numbers character by character
- **Sign Tracking**: Handle positive/negative number signs
- **Parentheses Processing**: Save state, evaluate sub-expressions

## 2. PROBLEM CHARACTERISTICS
- **Infix Notation**: Standard mathematical expression format
- **Operator Precedence**: Parentheses override normal precedence
- **Left to Right**: Process characters in order
- **State Management**: Track current number, sign, and stack state

## 3. SIMILAR PROBLEMS
- Evaluate Reverse Polish Notation (LeetCode 150) - Postfix evaluation
- Basic Calculator II (LeetCode 227) - Handle all operators and precedence
- Expression Add Operators (LeetCode 282) - Insert operators into expression
- Calculator with Functions - Handle mathematical functions

## 4. KEY OBSERVATIONS
- **Parentheses Handling**: Push current result and sign, start new sub-expression
- **Number Parsing**: Build numbers digit by digit
- **Sign Management**: Track sign for current number being parsed
- **Stack Structure**: Store results and signs separately for parentheses

## 5. VARIATIONS & EXTENSIONS
- **All Operators**: Support for multiplication, division, modulo
- **Operator Precedence**: Handle precedence without parentheses
- **Error Handling**: Report syntax errors and positions
- **Floating Point**: Support for decimal numbers

## 6. INTERVIEW INSIGHTS
- Always clarify: "Operator set? Precedence rules? Error handling?"
- Edge cases: empty string, invalid characters, division by zero
- Time complexity: O(N) where N=string length
- Space complexity: O(N) for stack in worst case
- Key insight: stack saves state for parentheses evaluation

## 7. COMMON MISTAKES
- Not handling multi-digit numbers correctly
- Wrong sign management for negative numbers
- Incorrect parentheses state restoration
- Not handling spaces properly
- Missing operator precedence handling

## 8. OPTIMIZATION STRATEGIES
- **Single Pass**: O(N) time, O(N) space - optimal
- **State Machine**: Use state machine for parsing
- **Early Validation**: Check for invalid characters
- **Memory Optimization**: Minimize stack usage

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like evaluating a mathematical expression step by step:**
- You're reading an expression from left to right
- When you see parentheses, you need to evaluate what's inside first
- Stack helps you remember where you were and what the result was
- Like having a calculator that shows intermediate steps
- Need to track numbers, signs, and parentheses levels

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String containing mathematical expression
2. **Goal**: Evaluate the expression and return result
3. **Constraints**: Handle +, -, parentheses, spaces, multi-digit numbers
4. **Output**: Integer result of evaluation

#### Phase 2: Key Insight Recognition
- **"Parentheses priority"** → Need to evaluate inner parentheses first
- **"Stack saves state"** → Push current result and sign when entering parentheses
- **"Number building"** → Parse multi-digit numbers digit by digit
- **"Sign tracking"** → Handle negative numbers correctly

#### Phase 3: Strategy Development
```
Human thought process:
"I need to evaluate infix expression with parentheses.
This requires handling operator precedence:

Stack-Based Approach:
1. Iterate through characters left to right
2. Parse numbers digit by digit
3. Track current sign (+ or -)
4. When encountering operator:
   - Add current number to result with its sign
   - Update sign for next number
5. When encountering '(':
   - Push current result and sign to stack
   - Reset for new sub-expression
6. When encountering ')':
   - Add current number to result
   - Pop sign and previous result from stack
   - Combine: previousResult + sign*currentResult
7. At end, add final number to result

This handles parentheses correctly!"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return 0 or handle appropriately
- **Spaces**: Skip spaces during parsing
- **Single number**: Return the number directly
- **Invalid characters**: Handle or return error

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Expression: "(1+(4+5+2)-3)+(6+8)" = 23

Human thinking:
"Step 1: '(' → entering parentheses
Push result=0, sign=1 to stack
Reset: result=0, sign=1, num=0

Step 2: '1' → building number
num = 1

Step 3: '+' → operator
result += sign * num = 0 + 1 = 1
sign = 1, num = 0

Step 4: '(' → entering nested parentheses
Push result=1, sign=1 to stack
Reset: result=0, sign=1, num=0

Step 5: '4' → building number
num = 4

Step 6: '+' → operator
result += sign * num = 0 + 4 = 4
sign = 1, num = 0

Continue parsing inner parentheses...
When hitting ')': 
result += sign * num = 4 + 5 + 2 = 11
Pop sign=1, prevResult=1 from stack
result = prevResult + sign*result = 1 + 11 = 12

Continue with outer parentheses...
Final result: 23 ✓"
```

#### Phase 6: Intuition Validation
- **Why stack works**: Saves state when entering parentheses
- **Why sign tracking**: Handles negative numbers correctly
- **Why number building**: Parses multi-digit numbers
- **Why O(N)**: Each character processed once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use two stacks?"** → Single stack with result/sign works
2. **"Should I handle multiplication?"** → Problem statement limits to +,-
3. **"What about operator precedence?"** → Only parentheses needed for this problem
4. **"Can I optimize further?"** → O(N) is already optimal
5. **"What about negative numbers?"** → Sign tracking handles this

### Real-World Analogy
**Like using a basic calculator step by step:**
- You're typing an expression into a calculator
- Calculator shows intermediate results
- When you hit parentheses, calculator evaluates what's inside
- Like solving math problems step by step
- Stack memory helps remember previous calculations

### Human-Readable Pseudocode
```
function calculate(s):
    stack = empty stack
    result = 0
    sign = 1
    num = 0
    
    for i from 0 to len(s)-1:
        char = s[i]
        
        if char is digit:
            num = num * 10 + (char - '0')
        else if char is '+':
            result += sign * num
            sign = 1
            num = 0
        else if char is '-':
            result += sign * num
            sign = -1
            num = 0
        else if char is '(':
            stack.push(result)
            stack.push(sign)
            result = 0
            sign = 1
            num = 0
        else if char is ')':
            result += sign * num
            prevSign = stack.pop()
            prevResult = stack.pop()
            result = prevResult + prevSign * result
            num = 0
        // skip spaces
    
    result += sign * num
    return result
```

### Execution Visualization

### Example: Expression = "(1+(4+5+2)-3)" = 9
```
Character Processing:
Index: 0  1  2  3  4  5  6  7  8  9 10 11
Char:  (  1  +  (  4  +  5  +  2  )  -  3  )
Stack: [] [0,1] [0,1,1] [0,1,1] [0,1,1] [0,1,1] [0,1] [0,1] [0,1]

Step-by-step:
1. '(' → push result=0, sign=1, reset
   Stack: [0,1], result=0, sign=1, num=0

2. '1' → num=1
   num=1

3. '+' → result+=1*1=1, sign=1, num=0
   result=1

4. '(' → push result=1, sign=1, reset
   Stack: [0,1,1,1], result=0, sign=1, num=0

5. '4' → num=4
   num=4

6. '+' → result+=1*4=4, sign=1, num=0
   result=4

7. '5' → num=5
   num=5

8. '+' → result+=1*5=9, sign=1, num=0
   result=9

9. '2' → num=2
   num=2

10. ')' → result+=1*2=11
   pop sign=1, prevResult=1
   result = 1 + 1*11 = 12
   Stack: [0,1]

11. '-' → result+=1*3=9, sign=-1, num=0
   result=9, sign=-1

Final Result: 9 ✓
```

### Key Visualization Points:
- **Number Building**: Parse multi-digit numbers digit by digit
- **Operator Processing**: Apply sign and add to result
- **Parentheses Handling**: Save state on '(', restore on ')'
- **Stack Evolution**: Push result/sign pairs, pop on closing

### Memory Layout Visualization:
```
Stack State Evolution:
Initial:  []
After '(':  [0,1]      (result=0, sign=1)
After second '(': [0,1,1,1]  (result=1, sign=1)
After ')':  [0,1]        (popped sign=1, result=1)

Variable Tracking:
result: 0→1→4→9→11→12→9
sign: 1→1→1→1→1→-1
num: 0→1→0→4→0→5→0→2→0→3→0

Evaluation Flow:
1. Parse numbers digit by digit
2. On operators: apply previous number to result
3. On '(': save current state, start fresh
4. On ')': finish sub-expression, restore state
5. Combine results using saved signs
```

### Time Complexity Breakdown:
- **Character Processing**: O(N) time where N=string length
- **Stack Operations**: O(1) per push/pop operation
- **Number Building**: O(1) per digit
- **Total**: O(N) time, O(N) space in worst case
- **Optimal**: Cannot do better than processing each character

### Alternative Approaches:

#### 1. Two-Stack Approach (O(N) time, O(N) space)
```go
func calculateTwoStacks(s string) int {
    // One stack for numbers, one for operators
    // More complex but handles all operators
    // ... implementation details omitted
}
```
- **Pros**: Can handle all operators and precedence
- **Cons**: More complex implementation

#### 2. Recursive Evaluation (O(N) time, O(N) space)
```go
func calculateRecursive(s string) int {
    // Recursively evaluate innermost parentheses first
    // Less efficient than stack approach
    // ... implementation details omitted
}
```
- **Pros**: Elegant for parentheses handling
- **Cons**: Less efficient, potential stack overflow

#### 3. Shunting Yard Algorithm (O(N) time, O(N) space)
```go
func calculateShuntingYard(s string) int {
    // Convert infix to postfix, then evaluate
    // Overkill for simple calculator
    // ... implementation details omitted
}
```
- **Pros**: Handles all operators and precedence
- **Cons**: Overkill for this specific problem

### Extensions for Interviews:
- **All Operators**: Support for *, /, %, ^, etc.
- **Operator Precedence**: Handle precedence without parentheses
- **Error Handling**: Report syntax errors and positions
- **Floating Point**: Support for decimal numbers
- **Performance Analysis**: Discuss worst-case scenarios
*/
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
