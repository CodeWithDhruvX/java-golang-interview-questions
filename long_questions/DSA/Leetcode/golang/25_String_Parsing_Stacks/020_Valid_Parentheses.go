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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Stack-Based Bracket Matching
- **Stack Usage**: Push opening brackets, pop on matching closing
- **Bracket Mapping**: Map closing brackets to opening brackets
- **LIFO Principle**: Last opened bracket must be closed first
- **Validation**: String valid if stack empty at end

## 2. PROBLEM CHARACTERISTICS
- **Nested Structure**: Brackets can be nested within each other
- **Matching Pairs**: Each opening has specific closing bracket
- **Order Validation**: Brackets must close in correct order
- **Complete Matching**: All brackets must be properly paired

## 3. SIMILAR PROBLEMS
- Generate Parentheses (LeetCode 22) - Generate valid combinations
- Longest Valid Parentheses (LeetCode 32) - Find longest valid substring
- Remove Invalid Parentheses (LeetCode 301) - Remove minimum to make valid
- Score of Parentheses (LeetCode 856) - Calculate score based on nesting

## 4. KEY OBSERVATIONS
- **Stack Natural Fit**: LIFO property matches bracket nesting
- **Early Validation**: Can detect invalid immediately on mismatch
- **Empty Stack Check**: Final validation requires empty stack
- **Character Validation**: Need to handle invalid characters

## 5. VARIATIONS & EXTENSIONS
- **Different Bracket Types**: Additional bracket types (angle brackets)
- **Multiple Strings**: Validate multiple strings simultaneously
- **Error Reporting**: Return specific error types/positions
- **Partial Validation**: Find longest valid substring

## 6. INTERVIEW INSIGHTS
- Always clarify: "Bracket types? Invalid characters? Empty string?"
- Edge cases: empty string, single bracket, invalid characters
- Time complexity: O(N) where N=string length
- Space complexity: O(N) in worst case (all opening brackets)
- Key insight: stack perfectly models nested structure

## 7. COMMON MISTAKES
- Not handling empty stack before popping
- Wrong bracket mapping (closing to opening)
- Not checking final stack emptiness
- Missing invalid character handling
- Off-by-one errors in stack indexing

## 8. OPTIMIZATION STRATEGIES
- **Basic Stack**: O(N) time, O(N) space - optimal
- **Early Termination**: Stop on first invalid character
- **Character Filtering**: Pre-filter valid characters
- **Memory Pool**: Reuse stack for multiple validations

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like matching nested containers:**
- You have containers that can contain other containers
- Each container type has specific lid (bracket pair)
- Must close containers in reverse order of opening
- Like Russian nesting dolls - smallest opened last
- Stack naturally models this nesting behavior

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String containing brackets
2. **Goal**: Determine if brackets are properly matched
3. **Rules**: Each opening must have matching closing in correct order
4. **Output**: Boolean indicating validity

#### Phase 2: Key Insight Recognition
- **"Stack natural fit"** → LIFO matches bracket nesting
- **"Bracket mapping"** → Need to map closing to opening
- **"Early validation"** → Can detect errors immediately
- **"Final check"** → Stack must be empty at end

#### Phase 3: Strategy Development
```
Human thought process:
"I need to validate bracket matching.
This is perfect for stack usage:

Stack Approach:
1. Iterate through each character
2. If opening bracket: push to stack
3. If closing bracket:
   - Check if stack empty (invalid)
   - Pop top and check if matches
   - If no match: invalid
4. After processing all characters:
   - Stack must be empty for valid string

This naturally handles nesting and order!"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return true (no brackets to validate)
- **Single bracket**: Return false (cannot be matched)
- **Invalid characters**: Return false or clarify handling
- **All opening**: Return false (stack not empty at end)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
String: "({[]})"

Human thinking:
"Step 1: '(' → opening, push to stack
Stack: [ ( ]

Step 2: '{' → opening, push to stack  
Stack: [ (, { ]

Step 3: '[' → opening, push to stack
Stack: [ (, {, [ ]

Step 4: ']' → closing, check top
Top is '[', matches ']' ✓ Pop
Stack: [ (, { ]

Step 5: '}' → closing, check top
Top is '{', matches '}' ✓ Pop
Stack: [ ( ]

Step 6: ')' → closing, check top
Top is '(', matches ')' ✓ Pop
Stack: [ ] (empty)

Step 7: End of string, stack empty ✓ Valid!"
```

#### Phase 6: Intuition Validation
- **Why stack works**: LIFO matches nested bracket behavior
- **Why mapping needed**: Need to know which opening matches closing
- **Why early check works**: Can detect invalid immediately
- **Why final check needed**: All openings must be closed

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just count?"** → Counting doesn't capture order
2. **"Should I use recursion?"** → Stack is more efficient
3. **"What about other characters?"** → Clarify handling policy
4. **"Can I optimize further?"** → O(N) is already optimal
5. **"What about different bracket types?"** → Extend mapping approach

### Real-World Analogy
**Like checking if a document has properly nested tags:**
- HTML/XML tags must be properly nested and closed
- Each opening tag has specific closing tag
- Tags can be nested within other tags
- Like validating HTML structure
- Stack naturally tracks which tag we're currently inside

### Human-Readable Pseudocode
```
function isValidParentheses(s):
    stack = empty stack
    bracketMap = {
        ')': '(',
        '}': '{',
        ']': '['
    }
    
    for char in s:
        if char is opening bracket:
            stack.push(char)
        else if char is closing bracket:
            if stack is empty:
                return false
            top = stack.pop()
            if top != bracketMap[char]:
                return false
        else:
            return false // invalid character
    
    return stack.isEmpty()
```

### Execution Visualization

### Example: String = "({[]})"
```
Character Processing:
Index: 0  1  2  3  4  5  6
Char:  (  {  [  ]  }  )
Stack: [ (] [({] [({[ [({] [({] [({] [({] [] 

Step-by-step:
1. '(' → opening → push: [ ( ]
2. '{' → opening → push: [ (, { ]
3. '[' → opening → push: [ (, {, [ ]
4. ']' → closing → top is '[', matches ✓ pop: [ (, { ]
5. '}' → closing → top is '{', matches ✓ pop: [ ( ]
6. ')' → closing → top is '(', matches ✓ pop: [ ]
7. End → stack empty ✓ Valid!

Final Result: true ✓
```

### Key Visualization Points:
- **Stack Growth**: Push all opening brackets
- **Stack Shrink**: Pop on successful matches
- **Order Validation**: LIFO ensures correct closing order
- **Final Check**: Empty stack means all matched

### Memory Layout Visualization:
```
Stack State Evolution:
Initial:  []
After '(':  [ ( ]
After '{':  [ (, { ]
After '[':  [ (, {, [ ]
After ']':  [ (, { ]    (popped '[')
After '}':  [ ( ]        (popped '{')
After ')':  [ ]            (popped '(')
Final:    []            (empty ✓)

Bracket Mapping:
')' → '('
'}' → '{'
']' → '['

Validation Steps:
1. Check character type
2. Push opening brackets
3. Match closing brackets with stack top
4. Validate final stack emptiness
```

### Time Complexity Breakdown:
- **Character Processing**: O(N) time where N=string length
- **Stack Operations**: O(1) per push/pop operation
- **Total**: O(N) time, O(N) space in worst case
- **Optimal**: Cannot do better than processing each character

### Alternative Approaches:

#### 1. Counter Replacement (O(N) time, O(1) space)
```go
func isValidCounters(s string) bool {
    // Only works for single bracket type
    // Not suitable for multiple bracket types
    // ... implementation details omitted
}
```
- **Pros**: Constant space
- **Cons**: Only works for single bracket type

#### 2. Two-Pass Validation (O(N) time, O(N) space)
```go
func isValidTwoPass(s string) bool {
    // First pass: check character counts
    // Second pass: validate order
    // More complex than stack approach
    // ... implementation details omitted
}
```
- **Pros**: Can provide more detailed error info
- **Cons**: More complex, same asymptotic complexity

#### 3. Recursive Approach (O(N) time, O(N) space)
```go
func isValidRecursive(s string) bool {
    // Recursively find and remove matched pairs
    // Less efficient than stack
    // ... implementation details omitted
}
```
- **Pros**: Elegant for some variations
- **Cons**: Less efficient, potential stack overflow

### Extensions for Interviews:
- **Different Bracket Types**: Handle angle brackets, quotes
- **Error Reporting**: Return specific error positions/types
- **Multiple Validation**: Validate multiple strings efficiently
- **Performance Analysis**: Discuss worst-case scenarios
- **Real-world Applications**: HTML/XML validation, expression parsing
*/
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
