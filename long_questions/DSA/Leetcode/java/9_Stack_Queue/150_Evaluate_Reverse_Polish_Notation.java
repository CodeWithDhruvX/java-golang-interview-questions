import java.util.*;

public class EvaluateReversePolishNotation {
    
    // 150. Evaluate Reverse Polish Notation
    // Time: O(N), Space: O(N)
    public static int evalRPN(String[] tokens) {
        if (tokens == null || tokens.length == 0) {
            return 0;
        }
        
        Stack<Integer> stack = new Stack<>();
        Set<String> operators = new HashSet<>(Arrays.asList("+", "-", "*", "/"));
        
        for (String token : tokens) {
            if (operators.contains(token)) {
                int b = stack.pop();
                int a = stack.pop();
                int result = 0;
                
                switch (token) {
                    case "+":
                        result = a + b;
                        break;
                    case "-":
                        result = a - b;
                        break;
                    case "*":
                        result = a * b;
                        break;
                    case "/":
                        result = a / b;
                        break;
                }
                
                stack.push(result);
            } else {
                stack.push(Integer.parseInt(token));
            }
        }
        
        return stack.pop();
    }

    public static void main(String[] args) {
        String[][] testCases = {
            {"2", "1", "+", "3", "*"},
            {"4", "13", "5", "/", "+"},
            {"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
            {"2", "3", "+"},
            {"4", "5", "-"},
            {"6", "7", "*"},
            {"8", "9", "/"},
            {"1"},
            {"2", "0", "/"},
            {"3", "4", "+", "5", "-"},
            {"6", "7", "*", "8", "+"},
            {"9", "10", "/", "11", "*"},
            {"12", "13", "+", "14", "*", "15", "+"},
            {"16", "17", "/", "18", "-", "19", "*"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            String[] tokens = testCases[i];
            int result = evalRPN(tokens);
            
            System.out.printf("Test Case %d: %s -> %d\n", 
                i + 1, Arrays.toString(tokens), result);
        }
    }
}
