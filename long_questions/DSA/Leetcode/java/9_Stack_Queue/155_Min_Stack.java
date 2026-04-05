import java.util.*;

public class MinStack {
    
    // 155. Min Stack
    // Time: O(1) for all operations, Space: O(N)
    public static class MinStackImplementation {
        private Stack<Integer> stack;
        private Stack<Integer> minStack;
        
        public MinStackImplementation() {
            stack = new Stack<>();
            minStack = new Stack<>();
        }
        
        public void push(int val) {
            stack.push(val);
            
            if (minStack.isEmpty() || val <= minStack.peek()) {
                minStack.push(val);
            }
        }
        
        public void pop() {
            if (stack.peek().equals(minStack.peek())) {
                minStack.pop();
            }
            stack.pop();
        }
        
        public int top() {
            return stack.peek();
        }
        
        public int getMin() {
            return minStack.peek();
        }
    }

    public static void main(String[] args) {
        // Test Case 1: Basic operations
        System.out.println("Test Case 1: Basic operations");
        MinStackImplementation minStack1 = new MinStackImplementation();
        minStack1.push(-2);
        minStack1.push(0);
        minStack1.push(-3);
        System.out.printf("Min after pushing -2, 0, -3: %d\n", minStack1.getMin());
        minStack1.pop();
        System.out.printf("Top after pop: %d\n", minStack1.top());
        System.out.printf("Min after pop: %d\n", minStack1.getMin());
        
        // Test Case 2: Duplicate minimums
        System.out.println("\nTest Case 2: Duplicate minimums");
        MinStackImplementation minStack2 = new MinStackImplementation();
        minStack2.push(1);
        minStack2.push(1);
        minStack2.push(2);
        System.out.printf("Min: %d\n", minStack2.getMin());
        minStack2.pop();
        System.out.printf("Min after pop 2: %d\n", minStack2.getMin());
        minStack2.pop();
        System.out.printf("Min after pop 1: %d\n", minStack2.getMin());
        
        // Test Case 3: Negative numbers
        System.out.println("\nTest Case 3: Negative numbers");
        MinStackImplementation minStack3 = new MinStackImplementation();
        minStack3.push(-5);
        minStack3.push(-3);
        minStack3.push(-10);
        System.out.printf("Min: %d\n", minStack3.getMin());
        minStack3.pop();
        System.out.printf("Min after pop: %d\n", minStack3.getMin());
        
        // Test Case 4: Single element
        System.out.println("\nTest Case 4: Single element");
        MinStackImplementation minStack4 = new MinStackImplementation();
        minStack4.push(42);
        System.out.printf("Min: %d, Top: %d\n", minStack4.getMin(), minStack4.top());
        
        // Test Case 5: Increasing sequence
        System.out.println("\nTest Case 5: Increasing sequence");
        MinStackImplementation minStack5 = new MinStackImplementation();
        minStack5.push(1);
        minStack5.push(2);
        minStack5.push(3);
        minStack5.push(4);
        System.out.printf("Min: %d\n", minStack5.getMin());
        minStack5.pop();
        minStack5.pop();
        System.out.printf("Min after 2 pops: %d\n", minStack5.getMin());
    }
}
