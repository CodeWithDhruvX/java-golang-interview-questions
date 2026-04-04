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
}
