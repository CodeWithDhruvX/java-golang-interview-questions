import java.util.*;

public class ValidParentheses {
    
    // 20. Valid Parentheses
    // Time: O(N), Space: O(N)
    public static boolean isValid(String s) {
        Deque<Character> stack = new ArrayDeque<>();
        
        // Map closing brackets to opening brackets
        Map<Character, Character> bracketMap = new HashMap<>();
        bracketMap.put(')', '(');
        bracketMap.put('}', '{');
        bracketMap.put(']', '[');
        
        for (char c : s.toCharArray()) {
            // If it's an opening bracket, push to stack
            switch (c) {
                case '(', '{', '[':
                    stack.push(c);
                    break;
                case ')', '}', ']':
                    // If stack is empty or top doesn't match, return false
                    if (stack.isEmpty() || stack.pop() != bracketMap.get(c)) {
                        return false;
                    }
                    break;
                default:
                    // Invalid character
                    return false;
            }
        }
        
        // Stack should be empty for valid string
        return stack.isEmpty();
    }

    public static void main(String[] args) {
        // Test cases
        String[] testCases = {
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
            "([{}]))"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            boolean result = isValid(testCases[i]);
            System.out.printf("Test Case %d: \"%s\" -> Valid: %b\n", 
                i + 1, testCases[i], result);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: String Parsing with Stacks
- **Bracket Matching**: Validate balanced parentheses/brackets
- **Stack Operations**: Push opening, pop for closing
- **Character Mapping**: Map closing brackets to opening ones
- **LIFO Processing**: Last opened must be first closed

## 2. PROBLEM CHARACTERISTICS
- **String Validation**: Check if string has valid bracket sequence
- **Multiple Bracket Types**: Parentheses, braces, brackets
- **Balanced Property**: Every opening has matching closing
- **Order Constraint**: Closing must match most recent opening

## 3. SIMILAR PROBLEMS
- Valid Palindrome
- Remove Invalid Parentheses
- Longest Valid Parentheses
- Minimum Add to Make Valid

## 4. KEY OBSERVATIONS
- Stack naturally handles LIFO bracket matching
- Map closing brackets to opening for validation
- Stack empty at end means all brackets matched
- Any mismatched or extra closing makes string invalid
- Time complexity: O(N) where N is string length

## 5. VARIATIONS & EXTENSIONS
- Different bracket types
- Minimum insertions to balance
- Maximum nesting depth
- Bracket coloring problems
- Expression evaluation

## 6. INTERVIEW INSIGHTS
- Clarify: "What characters should I consider?"
- Edge cases: empty string, single bracket, mismatched types
- Time complexity: O(N) vs O(N²) naive
- Space complexity: O(N) worst case

## 7. COMMON MISTAKES
- Not handling all bracket types
- Incorrect mapping of closing to opening
- Not checking stack emptiness properly
- Forgetting to handle invalid characters
- Off-by-one errors in string processing

## 8. OPTIMIZATION STRATEGIES
- Early termination on obvious mismatches
- Use switch statement for character handling
- Handle invalid characters gracefully
- Optimize for specific character sets

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like matching socks and shoes:**
- You have a pile of socks (opening brackets)
- Each time you encounter a sock, put it on the pile
- When you encounter a shoe (closing bracket), it must match the sock
- Take the top sock from the pile and pair with the shoe
- If shoe doesn't match top sock, or pile is empty when shoe arrives → invalid
- At the end, pile should be empty (all socks paired)

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: String containing brackets
2. **Goal**: Check if brackets are balanced
3. **Output**: Boolean indicating validity

#### Phase 2: Key Insight Recognition
- **"What makes brackets balanced?"** → Each opening has matching closing
- **"How to track openings?"** → Stack (LIFO) data structure
- **"What's the matching rule?"** → Closing must match most recent opening
- **"Why stack?"** → Natural LIFO behavior for bracket matching

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use a stack:
1. Map closing brackets to opening ones
2. For each character in string:
   - If opening: push to stack
   - If closing: check stack top matches
   - If matches: pop from stack
   - If doesn't match: return false
3. At end, stack should be empty for valid string"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Return true (no brackets to match)
- **Single bracket**: Return false (unmatched)
- **Mismatched types**: Return false immediately
- **Extra closing**: Return false when stack empty

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
String: "({[]})"

Human thinking:
"Let's process character by character:

Stack: [], Position: 0
Char '(': opening → push '(', Stack: ['(']
Char '{': opening → push '{', Stack: ['(', '{']
Char '[': opening → push '[', Stack: ['(', '{', '[']
Char ']': closing → top is '[', matches! → pop, Stack: ['(', '{']
Char ']': closing → top is '{', doesn't match! → return false ✗

Found mismatch: string is invalid ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Stack ensures LIFO matching property
- **Why it's efficient**: Each character processed once
- **Why it's correct**: All bracket pairs validated

### Common Human Pitfalls & How to Avoid Them
1. **"Why not count brackets?"** → Can't handle order, only count
2. **"What about recursion?"** → Stack is more natural for this problem
3. **"How to handle multiple types?"** → Need proper mapping
4. **"What about invalid characters?"** → Should return false

### Real-World Analogy
**Like matching containers and lids:**
- You have containers (opening brackets) and lids (closing brackets)
- Each container must match its specific lid type
- You stack containers as you encounter them
- When you find a lid, it must match the top container
- If lid doesn't match container type, you can't seal it
- At the end, all containers should be sealed (stack empty)
- This ensures proper container-lid matching

### Human-Readable Pseudocode
```
function isValid(s):
    stack = empty stack
    bracketMap = {')': '(', '}': '{', ']': '['}
    
    for character in s:
        if character is opening bracket:
            stack.push(character)
        else if character is closing bracket:
            if stack is empty or stack.pop() != bracketMap[character]:
                return false
    
    return stack.isEmpty()
```

### Execution Visualization

### Example: "({[]})"
```
Character Processing:
Stack Evolution:
Position 0: '(', push → ['(']
Position 1: '{', push → ['(', '{']
Position 2: '[', push → ['(', '{', '[']
Position 3: ']', top='[' matches → ['(', '{']
Position 4: ']', top='{' matches → ['(']
Position 5: '}', top='(' matches → []
Position 6: ')', stack empty → VIOLATION!

Result: Invalid (unmatched ')') ✓

Visualization:
( { [ ] } )
↑ ↑ ↑ ↑ ↑
| | | | |
| | | | └─ Mismatch!
```

### Key Visualization Points:
- **Stack behavior** naturally handles LIFO matching
- **Bracket mapping** ensures correct pairing
- **Early termination** on mismatches
- **Final check** ensures all brackets matched

### Memory Layout Visualization:
```
Stack Evolution:
Step 1: push '(' → ['(']
Step 2: push '{' → ['(', '{']
Step 3: push '[' → ['(', '{', '[']
Step 4: pop '[' (matches ']') → ['(', '{']
Step 5: pop '{' (matches ']') → ['(']
Step 6: pop '(' (matches ')') → []
Step 7: encounter ')' with empty stack → ERROR

Final State: Empty stack would be valid
Actual: Stack empty when encountering ')' → Invalid
```

### Time Complexity Breakdown:
- **Each Character**: O(1) stack operations
- **String Processing**: O(N) where N is string length
- **Total**: O(N) time, O(N) worst-case space
- **Optimal**: Cannot do better than O(N) for this problem
- **vs Counting**: O(N) time but can't handle order
*/
