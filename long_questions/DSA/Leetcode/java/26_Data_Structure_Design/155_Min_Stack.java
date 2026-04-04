import java.util.*;

public class MinStack {
    
    // 155. Min Stack
    // Time: O(1) for all operations, Space: O(N)
    private Stack<Integer> stack;
    private Stack<Integer> minStack;

    // Constructor initializes the stack
    public MinStack() {
        stack = new Stack<>();
        minStack = new Stack<>();
    }

    // Push pushes element onto stack
    public void push(int val) {
        stack.push(val);
        
        // Push to minStack if it's empty or val is <= current min
        if (minStack.isEmpty() || val <= minStack.peek()) {
            minStack.push(val);
        }
    }

    // Pop removes the element on top of the stack
    public void pop() {
        if (stack.isEmpty()) {
            return;
        }
        
        int top = stack.pop();
        
        // Remove from minStack if it matches the popped element
        if (!minStack.isEmpty() && top == minStack.peek()) {
            minStack.pop();
        }
    }

    // Top gets the top element of the stack
    public int top() {
        if (stack.isEmpty()) {
            return -1; // Or handle error appropriately
        }
        return stack.peek();
    }

    // GetMin retrieves the minimum element in the stack
    public int getMin() {
        if (minStack.isEmpty()) {
            return -1; // Or handle error appropriately
        }
        return minStack.peek();
    }

    // Alternative implementation using single stack with min tracking
    public static class MinStackOptimized {
        private Stack<int[]> stack; // Each entry contains {val, min}

        public MinStackOptimized() {
            stack = new Stack<>();
        }

        public void push(int val) {
            int min = val;
            if (!stack.isEmpty()) {
                int currentMin = stack.peek()[1];
                if (val > currentMin) {
                    min = currentMin;
                }
            }
            
            stack.push(new int[]{val, min});
        }

        public void pop() {
            if (!stack.isEmpty()) {
                stack.pop();
            }
        }

        public int top() {
            if (stack.isEmpty()) {
                return -1;
            }
            return stack.peek()[0];
        }

        public int getMin() {
            if (stack.isEmpty()) {
                return -1;
            }
            return stack.peek()[1];
        }
    }

    // Alternative implementation using difference method
    public static class MinStackDiff {
        private Stack<Long> stack;
        private long min;

        public MinStackDiff() {
            stack = new Stack<>();
            min = 0;
        }

        public void push(int val) {
            if (stack.isEmpty()) {
                stack.push(0L);
                min = val;
            } else {
                long diff = (long)val - min;
                stack.push(diff);
                if (val < min) {
                    min = val;
                }
            }
        }

        public void pop() {
            if (stack.isEmpty()) {
                return;
            }
            
            long diff = stack.pop();
            if (diff < 0) {
                min = min - diff;
            }
        }

        public int top() {
            if (stack.isEmpty()) {
                return -1;
            }
            
            long diff = stack.peek();
            if (diff < 0) {
                return (int)min;
            }
            return (int)(min + diff);
        }

        public int getMin() {
            if (stack.isEmpty()) {
                return -1;
            }
            return (int)min;
        }
    }

    public static void main(String[] args) {
        // Test cases
        System.out.println("=== Testing MinStack ===");
        
        // Test 1: Basic operations
        MinStack minStack = new MinStack();
        minStack.push(-2);
        minStack.push(0);
        minStack.push(-3);
        System.out.printf("After pushing -2, 0, -3:\n");
        System.out.printf("  Top: %d, Min: %d\n", minStack.top(), minStack.getMin());
        
        minStack.pop();
        System.out.printf("After pop:\n");
        System.out.printf("  Top: %d, Min: %d\n", minStack.top(), minStack.getMin());
        
        // Test 2: Edge cases
        System.out.println("\n=== Testing Edge Cases ===");
        MinStack emptyStack = new MinStack();
        System.out.printf("Empty stack - Top: %d, Min: %d\n", emptyStack.top(), emptyStack.getMin());
        
        emptyStack.push(5);
        System.out.printf("After pushing 5 - Top: %d, Min: %d\n", emptyStack.top(), emptyStack.getMin());
        
        // Test 3: Duplicate minimums
        System.out.println("\n=== Testing Duplicate Minimums ===");
        MinStack dupStack = new MinStack();
        dupStack.push(1);
        dupStack.push(1);
        dupStack.push(2);
        System.out.printf("After pushing 1, 1, 2 - Top: %d, Min: %d\n", dupStack.top(), dupStack.getMin());
        
        dupStack.pop();
        System.out.printf("After pop - Top: %d, Min: %d\n", dupStack.top(), dupStack.getMin());
        
        // Test optimized versions
        System.out.println("\n=== Testing Optimized Version ===");
        MinStackOptimized optStack = new MinStackOptimized();
        optStack.push(-2);
        optStack.push(0);
        optStack.push(-3);
        System.out.printf("After pushing -2, 0, -3:\n");
        System.out.printf("  Top: %d, Min: %d\n", optStack.top(), optStack.getMin());
        
        optStack.pop();
        System.out.printf("After pop:\n");
        System.out.printf("  Top: %d, Min: %d\n", optStack.top(), optStack.getMin());
        
        // Test difference method
        System.out.println("\n=== Testing Difference Method ===");
        MinStackDiff diffStack = new MinStackDiff();
        diffStack.push(-2);
        diffStack.push(0);
        diffStack.push(-3);
        System.out.printf("After pushing -2, 0, -3:\n");
        System.out.printf("  Top: %d, Min: %d\n", diffStack.top(), diffStack.getMin());
        
        diffStack.pop();
        System.out.printf("After pop:\n");
        System.out.printf("  Top: %d, Min: %d\n", diffStack.top(), diffStack.getMin());
    }
}
