import java.util.*;

public class ValidParentheses {
    
    // 20. Valid Parentheses
    // Time: O(N), Space: O(N)
    public static boolean isValid(String s) {
        if (s == null || s.length() == 0) {
            return true;
        }
        
        Stack<Character> stack = new Stack<>();
        Map<Character, Character> mapping = new HashMap<>();
        mapping.put(')', '(');
        mapping.put('}', '{');
        mapping.put(']', '[');
        
        for (char c : s.toCharArray()) {
            if (mapping.containsValue(c)) {
                stack.push(c);
            } else if (mapping.containsKey(c)) {
                if (stack.isEmpty() || stack.pop() != mapping.get(c)) {
                    return false;
                }
            }
        }
        
        return stack.isEmpty();
    }

    public static void main(String[] args) {
        String[] testCases = {
            "()",
            "()[]{}",
            "(]",
            "([)]",
            "{[]}",
            "",
            "(",
            ")",
            "((()))",
            "({[]})",
            "([{}])",
            "({[)]}",
            "({[]})()[]{}",
            "({[({[]})]})",
            "({[({[({[]})]})]})"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            String s = testCases[i];
            boolean result = isValid(s);
            
            System.out.printf("Test Case %d: \"%s\" -> %b\n", 
                i + 1, s, result);
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Stack-Based Parentheses Matching
- **Stack Usage**: LIFO structure for opening brackets
- **Mapping Strategy**: Map closing to opening brackets
- **Validation**: Check matching pairs and proper nesting
- **Early Termination**: Return false on first mismatch

## 2. PROBLEM CHARACTERISTICS
- **Bracket Types**: Three types: (), {}, []
- **Nesting Rules**: Opening must come before closing
- **Perfect Matching**: Every opening has corresponding closing
- **Order Matters**: Last opened must be first closed (LIFO)

## 3. SIMILAR PROBLEMS
- Evaluate Reverse Polish Notation
- Min Stack with additional operations
- Longest Valid Parentheses
- Generate Parentheses

## 4. KEY OBSERVATIONS
- **Opening Brackets**: Push to stack
- **Closing Brackets**: Must match top of stack
- **Invalid Cases**: Empty stack for closing, mismatched pair
- **Final Check**: Stack must be empty for valid string

## 5. VARIATIONS & EXTENSIONS
- Different bracket types (angle brackets, quotes)
- Multiple types of brackets in same string
- Return position of first error
- Generate all valid combinations

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I consider empty string as valid?"
- Edge cases: single bracket, all opening, all closing
- Time complexity: O(N) - single pass through string
- Space complexity: O(N) in worst case (all opening brackets)

## 7. COMMON MISTAKES
- Not handling empty string case
- Incorrect mapping of closing to opening brackets
- Forgetting to check stack emptiness before popping
- Not handling characters that aren't brackets

## 8. OPTIMIZATION STRATEGIES
- Use array instead of HashMap for bracket mapping
- Early return on invalid character
- Use switch statement for bracket types
- Pre-allocate stack capacity if possible

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like matching socks and shoes:**
- Opening brackets are like putting on socks (left foot, right foot)
- Closing brackets are like putting on matching shoes
- You must put on socks before shoes (LIFO order)
- Wrong order or mismatched pairs are invalid

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String containing brackets and other characters
2. **Goal**: Determine if brackets are properly matched
3. **Output**: Boolean indicating validity
4. **Rules**: Every opening must have corresponding closing in correct order

#### Phase 2: Key Insight Recognition
- **"How to track bracket pairs?"** → Use stack for LIFO order
- **"How to validate closing?"** → Check against stack top
- **"What about other characters?"** → Ignore them or handle separately
- **"When is it invalid?"** → Mismatch or stack empty when closing

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use a stack to track opening brackets:
1. Create mapping of closing to opening brackets
2. Iterate through string character by character
3. If opening bracket: push to stack
4. If closing bracket: check stack top
   - If stack empty or doesn't match: invalid
   - If matches: pop and continue
5. At end: valid if stack is empty"
```

#### Phase 4: Algorithm Walkthrough
```
Example: "({[]})"

Human thinking:
"Stack: [], result = true so far

c='(' (opening): push '('
Stack: ['('], result = true

c='{' (opening): push '{'
Stack: ['(', '{'], result = true

c='[' (opening): push '['
Stack: ['(', '{', '['], result = true

c=']' (closing): top is '[', matches '['
Stack: ['(', '{'], result = true

c='}' (closing): top is '{', matches '{'
Stack: ['('], result = true

c=')' (closing): top is '(', matches '('
Stack: [], result = true

End of string: stack is empty, so valid ✓"
```

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just count brackets?"** → Order matters, need LIFO
2. **"What about nested structures?"** → Stack handles nesting naturally
3. **"How to handle invalid characters?"** → Ignore or return false
4. **"When to stop early?"** → Return false on first mismatch

### Real-World Analogy
**Like checking proper code formatting:**
- Opening brackets are like opening code blocks
- Closing brackets are like ending code blocks
- Proper nesting ensures code structure is valid
- Mismatched brackets are like syntax errors
- Stack ensures inner blocks close before outer blocks

### Human-Readable Pseudocode
```
function isValid(s):
    if s is empty: return true
    
    stack = empty stack
    mapping = {')':'(', '}':'{', ']':'['}
    
    for each character c in s:
        if c is opening bracket:
            stack.push(c)
        else if c is closing bracket:
            if stack is empty: return false
            top = stack.pop()
            if top != mapping[c]: return false
    
    return stack.isEmpty()
```

### Execution Visualization

### Example: "([)]"
```
Step-by-step:
c='(' → push: [ ]
c='[' → push: [(, [ ]
c=')' → top='[', expected='(', mismatch! → return false

Stack state when invalid: [(, [ ]
Result: FALSE (mismatch at position 2)
```

### Example: "({[]})"
```
Step-by-step:
c='(' → push: [ ]
c='{' → push: [(, { ]
c='[' → push: [(, {, [ ]
c=']' → top='[', matches, pop: [(, { ]
c='}' → top='{', matches, pop: [(, ]
c=')' → top='(', matches, pop: []
End: stack empty → valid ✓

Stack Evolution:
[ ] → [(, [ ] → [(, {, [ ] → [(, { ] → [(, ] → [ ] → []
Result: TRUE ✓
```

### Key Visualization Points:
- **Stack LIFO**: Last opened must be first closed
- **Mapping Validation**: Closing must match expected opening
- **Early Termination**: Return false on first mismatch
- **Final Check**: Empty stack means all brackets matched

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each character visited once
- **Stack Operations**: O(1) push/pop amortized
- **Space**: O(N) worst case (all opening brackets)
- **Optimal**: Cannot do better than O(N) for this problem
*/
