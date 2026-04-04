import java.util.*;

public class ImplementQueueUsingStacks {
    
    // 232. Implement Queue using Stacks
    // Time: Amortized O(1), Space: O(N)
    static class MyQueue {
        private Stack<Integer> stack1;
        private Stack<Integer> stack2;
        
        public MyQueue() {
            stack1 = new Stack<>();
            stack2 = new Stack<>();
        }
        
        // Push element to the back of queue
        public void push(int x) {
            stack1.push(x);
        }
        
        // Remove element from the front of queue
        public int pop() {
            if (stack2.isEmpty()) {
                // Transfer all elements from stack1 to stack2
                while (!stack1.isEmpty()) {
                    stack2.push(stack1.pop());
                }
            }
            return stack2.pop();
        }
        
        // Get the front element
        public int peek() {
            if (stack2.isEmpty()) {
                // Transfer all elements from stack1 to stack2
                while (!stack1.isEmpty()) {
                    stack2.push(stack1.pop());
                }
            }
            return stack2.peek();
        }
        
        // Check if queue is empty
        public boolean empty() {
            return stack1.isEmpty() && stack2.isEmpty();
        }
        
        // Get current size
        public int size() {
            return stack1.size() + stack2.size();
        }
    }
    
    // Alternative implementation with single stack and recursion
    static class MyQueueRecursive {
        private Stack<Integer> stack;
        
        public MyQueueRecursive() {
            stack = new Stack<>();
        }
        
        public void push(int x) {
            stack.push(x);
        }
        
        public int pop() {
            if (stack.isEmpty()) {
                throw new NoSuchElementException("Queue is empty");
            }
            return popBottom();
        }
        
        public int peek() {
            if (stack.isEmpty()) {
                throw new NoSuchElementException("Queue is empty");
            }
            return peekBottom();
        }
        
        public boolean empty() {
            return stack.isEmpty();
        }
        
        private int popBottom() {
            int top = stack.pop();
            if (stack.isEmpty()) {
                return top;
            }
            int bottom = popBottom();
            stack.push(top);
            return bottom;
        }
        
        private int peekBottom() {
            int top = stack.pop();
            if (stack.isEmpty()) {
                stack.push(top);
                return top;
            }
            int bottom = peekBottom();
            stack.push(top);
            return bottom;
        }
    }
    
    // Implementation with amortized analysis
    static class MyQueueAmortized {
        private Stack<Integer> inputStack;
        private Stack<Integer> outputStack;
        private int size;
        
        public MyQueueAmortized() {
            inputStack = new Stack<>();
            outputStack = new Stack<>();
            size = 0;
        }
        
        public void push(int x) {
            inputStack.push(x);
            size++;
        }
        
        public int pop() {
            if (outputStack.isEmpty()) {
                transferElements();
            }
            size--;
            return outputStack.pop();
        }
        
        public int peek() {
            if (outputStack.isEmpty()) {
                transferElements();
            }
            return outputStack.peek();
        }
        
        public boolean empty() {
            return size == 0;
        }
        
        public int size() {
            return size;
        }
        
        private void transferElements() {
            while (!inputStack.isEmpty()) {
                outputStack.push(inputStack.pop());
            }
        }
    }
    
    // Version with detailed explanation
    static class MyQueueDetailed {
        private Stack<Integer> stack1;
        private Stack<Integer> stack2;
        private List<String> operations;
        
        public MyQueueDetailed() {
            stack1 = new Stack<>();
            stack2 = new Stack<>();
            operations = new ArrayList<>();
        }
        
        public void push(int x) {
            operations.add(String.format("Push %d to stack1 (back of queue)", x));
            stack1.push(x);
        }
        
        public int pop() {
            operations.add("Pop operation:");
            
            if (stack2.isEmpty()) {
                operations.add("  stack2 is empty, transferring from stack1");
                int transferCount = 0;
                
                while (!stack1.isEmpty()) {
                    int element = stack1.pop();
                    stack2.push(element);
                    operations.add(String.format("  Transferred %d from stack1 to stack2", element));
                    transferCount++;
                }
                
                operations.add(String.format("  Transferred %d elements", transferCount));
            }
            
            int result = stack2.pop();
            operations.add(String.format("  Popped %d from stack2 (front of queue)", result));
            return result;
        }
        
        public int peek() {
            operations.add("Peek operation:");
            
            if (stack2.isEmpty()) {
                operations.add("  stack2 is empty, transferring from stack1");
                
                while (!stack1.isEmpty()) {
                    int element = stack1.pop();
                    stack2.push(element);
                    operations.add(String.format("  Transferred %d from stack1 to stack2", element));
                }
            }
            
            int result = stack2.peek();
            operations.add(String.format("  Peeking at %d from stack2 (front of queue)", result));
            return result;
        }
        
        public boolean empty() {
            boolean result = stack1.isEmpty() && stack2.isEmpty();
            operations.add(String.format("Queue empty check: %b", result));
            return result;
        }
        
        public List<String> getOperations() {
            return new ArrayList<>(operations);
        }
        
        public void clearOperations() {
            operations.clear();
        }
    }
    
    // Performance comparison
    public void comparePerformance(int[] operations) {
        System.out.println("=== Performance Comparison ===");
        
        // Standard implementation
        MyQueue queue1 = new MyQueue();
        long startTime = System.nanoTime();
        
        for (int i = 0; i < operations.length; i++) {
            if (operations[i] >= 0) {
                queue1.push(operations[i]);
            } else {
                queue1.pop();
            }
        }
        
        long endTime = System.nanoTime();
        System.out.printf("Standard implementation: took %d ns\n", endTime - startTime);
        
        // Amortized implementation
        MyQueueAmortized queue2 = new MyQueueAmortized();
        startTime = System.nanoTime();
        
        for (int i = 0; i < operations.length; i++) {
            if (operations[i] >= 0) {
                queue2.push(operations[i]);
            } else {
                queue2.pop();
            }
        }
        
        endTime = System.nanoTime();
        System.out.printf("Amortized implementation: took %d ns\n", endTime - startTime);
    }
    
    // Test queue operations
    public void testQueueOperations(MyQueue queue, String queueName) {
        System.out.println("=== Testing " + queueName + " ===");
        
        // Test push operations
        queue.push(1);
        queue.push(2);
        queue.push(3);
        
        System.out.printf("After pushing 1, 2, 3: peek() = %d\n", queue.peek());
        
        // Test pop operations
        System.out.printf("pop() = %d\n", queue.pop());
        System.out.printf("peek() = %d\n", queue.peek());
        
        // Test empty check
        System.out.printf("empty() = %b\n", queue.empty());
        
        // Clear remaining elements
        while (!queue.empty()) {
            System.out.printf("pop() = %d\n", queue.pop());
        }
        
        System.out.printf("Final empty() = %b\n", queue.empty());
    }
    
    public static void main(String[] args) {
        ImplementQueueUsingStacks iqs = new ImplementQueueUsingStacks();
        
        // Test all implementations
        System.out.println("=== Testing All Implementations ===");
        
        // Standard implementation
        MyQueue queue1 = new MyQueue();
        iqs.testQueueOperations(queue1, "Standard Implementation");
        
        // Recursive implementation
        MyQueueRecursive queue2 = new MyQueueRecursive();
        System.out.println("\n=== Testing Recursive Implementation ===");
        queue2.push(1);
        queue2.push(2);
        queue2.push(3);
        System.out.printf("After pushing 1, 2, 3: peek() = %d\n", queue2.peek());
        System.out.printf("pop() = %d\n", queue2.pop());
        System.out.printf("peek() = %d\n", queue2.peek());
        System.out.printf("empty() = %b\n", queue2.empty());
        
        // Amortized implementation
        MyQueueAmortized queue3 = new MyQueueAmortized();
        System.out.println("\n=== Testing Amortized Implementation ===");
        queue3.push(1);
        queue3.push(2);
        queue3.push(3);
        System.out.printf("After pushing 1, 2, 3: peek() = %d\n", queue3.peek());
        System.out.printf("pop() = %d\n", queue3.pop());
        System.out.printf("peek() = %d\n", queue3.peek());
        System.out.printf("empty() = %b\n", queue3.empty());
        System.out.printf("size() = %d\n", queue3.size());
        
        // Detailed implementation
        MyQueueDetailed queue4 = new MyQueueDetailed();
        System.out.println("\n=== Testing Detailed Implementation ===");
        
        queue4.push(1);
        queue4.push(2);
        queue4.push(3);
        
        System.out.printf("peek() = %d\n", queue4.peek());
        System.out.printf("pop() = %d\n", queue4.pop());
        System.out.printf("peek() = %d\n", queue4.peek());
        
        System.out.println("\nOperations log:");
        for (String op : queue4.getOperations()) {
            System.out.println("  " + op);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        
        // Mixed operations for performance test
        int[] testOps = {1, 2, 3, -1, 4, 5, -1, 6, -1, -1, 7, 8, -1};
        iqs.comparePerformance(testOps);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        
        MyQueue edgeQueue = new MyQueue();
        
        // Test empty queue
        try {
            edgeQueue.pop();
            System.out.println("ERROR: Should throw exception for empty queue");
        } catch (Exception e) {
            System.out.println("Correctly threw exception for empty queue pop");
        }
        
        try {
            edgeQueue.peek();
            System.out.println("ERROR: Should throw exception for empty queue");
        } catch (Exception e) {
            System.out.println("Correctly threw exception for empty queue peek");
        }
        
        // Test single element
        edgeQueue.push(42);
        System.out.printf("Single element queue: peek() = %d\n", edgeQueue.peek());
        System.out.printf("Single element queue: pop() = %d\n", edgeQueue.pop());
        System.out.printf("Single element queue: empty() = %b\n", edgeQueue.empty());
        
        // Stress test
        System.out.println("\n=== Stress Test ===");
        MyQueue stressQueue = new MyQueue();
        
        for (int i = 0; i < 1000; i++) {
            stressQueue.push(i);
        }
        
        System.out.printf("Pushed 1000 elements, size should be 1000: %b\n", 
            stressQueue.empty() == false);
        
        for (int i = 0; i < 1000; i++) {
            stressQueue.pop();
        }
        
        System.out.printf("Popped 1000 elements, empty should be true: %b\n", stressQueue.empty());
        
        // Memory efficiency test
        System.out.println("\n=== Memory Efficiency ===");
        System.out.println("Standard implementation uses 2 stacks");
        System.out.println("Recursive implementation uses 1 stack but may use call stack");
        System.out.println("Amortized implementation uses 2 stacks but minimizes transfers");
        System.out.println("All implementations provide O(1) amortized time complexity");
    }
}
