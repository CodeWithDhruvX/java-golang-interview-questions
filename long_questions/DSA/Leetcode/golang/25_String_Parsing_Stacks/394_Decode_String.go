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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Stack-Based String Decoding
- **Stack Usage**: Push strings and numbers for pattern decoding
- **Pattern Recognition**: [number[string]] means repeat string number times
- **Nested Processing**: Handle nested patterns using stack
- **String Building**: Construct result by repeating substrings

## 2. PROBLEM CHARACTERISTICS
- **Encoded Format**: Numbers followed by strings in brackets
- **Repetition Logic**: [k[encoded]] means repeat encoded string k times
- **Nesting Support**: Patterns can be nested within other patterns
- **Sequential Processing**: Process characters left to right

## 3. SIMILAR PROBLEMS
- Encode String (LeetCode 394) - Reverse operation
- String Compression - Various encoding/decoding schemes
- Parser Design - Build parsers for different formats
- Regular Expression Matching - Pattern matching in strings

## 4. KEY OBSERVATIONS
- **Stack Natural Fit**: LIFO handles nested patterns perfectly
- **Number Building**: Parse multi-digit numbers character by character
- **String Accumulation**: Build current string character by character
- **Pattern Completion**: Process completed when closing bracket found

## 5. VARIATIONS & EXTENSIONS
- **Different Delimiters**: Support for {}, (), or custom delimiters
- **Multiple Brackets**: Handle different bracket types simultaneously
- **Error Handling**: Report invalid encoding formats
- **Large Numbers**: Handle very large repetition counts

## 6. INTERVIEW INSIGHTS
- Always clarify: "Bracket types? Number format? Error handling?"
- Edge cases: empty string, invalid format, large numbers
- Time complexity: O(N) where N=string length
- Space complexity: O(N) for stack in worst case
- Key insight: stack naturally handles nested patterns

## 7. COMMON MISTAKES
- Not handling multi-digit numbers correctly
- Wrong stack push/pop order
- Not handling empty strings in patterns
- Incorrect string repetition logic
- Not handling nested patterns properly

## 8. OPTIMIZATION STRATEGIES
- **Single Pass**: O(N) time, O(N) space - optimal
- **Early Validation**: Check for invalid format characters
- **String Builder**: Use efficient string building
- **Memory Pool**: Reuse stack for multiple decodings

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like decoding compressed text with nested instructions:**
- You have encoded text with repetition instructions
- [3[abc]] means repeat "abc" three times
- Instructions can be nested: [2[a[bc]]]
- Stack helps remember what you were working on when entering nested patterns
- Like following assembly instructions step by step

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Encoded string with pattern [number[string]]
2. **Goal**: Decode the string to original form
3. **Rules**: Repeat string inside brackets number times
4. **Output**: Decoded string

#### Phase 2: Key Insight Recognition
- **"Stack natural fit"** → LIFO handles nested patterns
- **"Number building"** → Parse multi-digit numbers
- **"String accumulation"** → Build current string character by character
- **"Pattern completion"** → Process when closing bracket found

#### Phase 3: Strategy Development
```
Human thought process:
"I need to decode nested repetition patterns.
This requires handling nested structures:

Stack-Based Approach:
1. Iterate through characters left to right
2. Build current number digit by digit
3. Build current string character by character
4. When encountering '[':
   - Push current string and number to stack
   - Reset for new pattern
5. When encountering ']':
   - Pop number and previous string from stack
   - Repeat current string number times
   - Concatenate with previous string
6. Normal characters: add to current string

This handles nesting correctly!"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return empty string
- **No brackets**: Return original string
- **Invalid format**: Handle or return error
- **Large numbers**: Handle potential memory issues

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
String: "3[a2[c]]" = "accaccacc"

Human thinking:
"Step 1: '3' → building number
currentNum = 3

Step 2: '[' → entering pattern
Push currentStr="", currentNum=3 to stack
Reset: currentStr="", currentNum=0

Step 3: 'a' → building string
currentStr = "a"

Step 4: '2' → building number
currentNum = 2

Step 5: '[' → entering nested pattern
Push currentStr="a", currentNum=2 to stack
Reset: currentStr="", currentNum=0

Step 6: 'c' → building string
currentStr = "c"

Step 7: ']' → closing nested pattern
Pop num=2, prevStr="a" from stack
currentStr = prevStr + repeat("c", 2) = "acc"

Step 8: ']' → closing outer pattern
Pop num=3, prevStr="" from stack
currentStr = prevStr + repeat("acc", 3) = "accaccacc"

Final result: "accaccacc" ✓"
```

#### Phase 6: Intuition Validation
- **Why stack works**: LIFO handles nested patterns perfectly
- **Why number building**: Parses multi-digit repetition counts
- **Why string accumulation**: Builds strings to be repeated
- **Why O(N)**: Each character processed once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use recursion?"** → Stack is more efficient and natural
2. **"Should I handle different brackets?"** → Clarify bracket types
3. **"What about empty patterns?"** → Handle [ ] or [0[]] correctly
4. **"Can I optimize further?"** → O(N) is already optimal
5. **"What about invalid formats?"** → Add validation logic

### Real-World Analogy
**Like decompressing a compressed file with nested instructions:**
- You have a compressed text file
- Contains instructions like "repeat 3 times: abc"
- Instructions can be nested within other instructions
- Stack memory helps remember where you were in the nesting
- Like following a recipe with sub-recipes
- Each instruction tells you how to build the final text

### Human-Readable Pseudocode
```
function decodeString(s):
    stack = empty stack
    currentNum = 0
    currentStr = ""
    
    for char in s:
        if char is digit:
            currentNum = currentNum * 10 + (char - '0')
        else if char == '[':
            stack.push(currentStr)
            stack.push(currentNum)
            currentStr = ""
            currentNum = 0
        else if char == ']':
            repeatTimes = stack.pop()
            prevStr = stack.pop()
            currentStr = prevStr + repeat(currentStr, repeatTimes)
        else:
            currentStr += char
    
    return currentStr

function repeat(s, times):
    result = ""
    for i from 0 to times-1:
        result += s
    return result
```

### Execution Visualization

### Example: String = "3[a2[c]]" = "accaccacc"
```
Character Processing:
Index: 0  1  2  3  4  5  6  7  8  9
Char:  3  [  a  2  [  c  ]  ]
Stack: [] ["",3] ["",3,"a",2] ["",3,"a",2] ["",3] ["",3]

Step-by-step:
1. '3' → currentNum = 3
2. '[' → push "",3, reset
   Stack: ["",3], currentStr="", currentNum=0
3. 'a' → currentStr = "a"
4. '2' → currentNum = 2
5. '[' → push "a",2, reset
   Stack: ["",3,"a",2], currentStr="", currentNum=0
6. 'c' → currentStr = "c"
7. ']' → pop 2,"a", currentStr = "a" + repeat("c",2) = "acc"
   Stack: ["",3]
8. ']' → pop 3,"", currentStr = "" + repeat("acc",3) = "accaccacc"
   Stack: []

Final Result: "accaccacc" ✓
```

### Key Visualization Points:
- **Number Building**: Parse multi-digit repetition counts
- **String Building**: Accumulate characters to be repeated
- **Stack Operations**: Push string/number pairs, pop on closing
- **Pattern Completion**: Process completed when closing bracket found

### Memory Layout Visualization:
```
Stack State Evolution:
Initial:  []
After '3[':  ["",3]      (currentStr="", currentNum=3)
After 'a2[': ["",3,"a",2]  (currentStr="", currentNum=2)
After first ']':  ["",3]        (popped "a",2, currentStr="acc")

Variable Tracking:
currentNum: 3→0→2→0→0→0
currentStr: ""→""→"a"→""→"c"→"acc"→"accaccacc"

Pattern Processing:
1. Parse numbers digit by digit
2. On '[': save current state, start new pattern
3. On ']': complete current pattern, restore previous state
4. Normal chars: add to current string
5. Repeat logic: string * number
```

### Time Complexity Breakdown:
- **Character Processing**: O(N) time where N=string length
- **Stack Operations**: O(1) per push/pop operation
- **String Repeating**: O(L) where L=length of string to repeat
- **Total**: O(N + total output length) time, O(N) space
- **Optimal**: Cannot do better than processing each character

### Alternative Approaches:

#### 1. Recursive Decoding (O(N) time, O(N) space)
```go
func decodeStringRecursive(s string) string {
    // Recursively decode nested patterns
    // Less efficient than stack approach
    // ... implementation details omitted
}
```
- **Pros**: Elegant for nested structures
- **Cons**: Less efficient, potential stack overflow

#### 2. Two-Pass Processing (O(N) time, O(N) space)
```go
func decodeStringTwoPass(s string) string {
    // First pass: identify pattern boundaries
    // Second pass: decode in chunks
    // More complex than stack approach
    // ... implementation details omitted
}
```
- **Pros**: Can provide more detailed error info
- **Cons**: More complex, same asymptotic complexity

#### 3. Iterator-Based (O(N) time, O(N) space)
```go
func decodeStringIterator(s string) string {
    // Use iterator to traverse string
    // Similar to stack approach with different data structure
    // ... implementation details omitted
}
```
- **Pros**: Different data structure approach
- **Cons**: Same complexity, more complex

### Extensions for Interviews:
- **Different Brackets**: Support for {}, (), or custom delimiters
- **Error Handling**: Report invalid format positions
- **Large Numbers**: Handle very large repetition counts efficiently
- **Multiple Formats**: Support different encoding schemes
- **Performance Analysis**: Discuss worst-case scenarios
*/
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
