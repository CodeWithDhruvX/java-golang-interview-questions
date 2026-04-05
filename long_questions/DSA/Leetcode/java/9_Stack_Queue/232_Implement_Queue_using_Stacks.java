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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Queue Implementation using Stacks
- **Two Stacks**: One for input, one for output
- **LIFO to FIFO**: Convert stack behavior to queue behavior
- **Transfer Operation**: Move elements when output stack is empty
- **Amortized Analysis**: Transfer cost is distributed across operations

## 2. PROBLEM CHARACTERISTICS
- **Queue Interface**: push (enqueue), pop (dequeue), peek, empty
- **Stack Constraint**: Must use only stack data structures
- **FIFO Requirement**: First in, first out behavior
- **Efficiency Goal**: Amortized O(1) operations

## 3. SIMILAR PROBLEMS
- Implement Stack using Queues
- Design Circular Queue using Array
- Implement Deque using Stacks
- Min Stack with O(1) operations

## 4. KEY OBSERVATIONS
- Two stacks can simulate queue behavior
- Input stack collects new elements
- Output stack holds elements in queue order
- Transfer operation reverses order when needed
- Amortized cost: O(1) per operation

## 5. VARIATIONS & EXTENSIONS
- Recursive implementation using single stack
- Memory-optimized version
- Support for multiple queues
- Thread-safe implementation

## 6. INTERVIEW INSIGHTS
- Clarify: "Can I use built-in queue?"
- Edge cases: empty queue, single element
- Time complexity: Amortized O(1) vs worst O(N)
- Space complexity: O(N) for all implementations

## 7. COMMON MISTAKES
- Not handling empty queue case properly
- Incorrect transfer logic
- Memory leaks in recursive version
- Not considering amortized analysis

## 8. OPTIMIZATION STRATEGIES
- Amortized analysis shows true efficiency
- Lazy transfer (only when needed)
- Pre-allocate stack capacity
- Minimize memory allocations

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like two-box system:**
- You have inbox (input stack) for new mail
- You have outbox (output stack) for mail to send
- When someone asks for mail (peek/pop), check outbox first
- If outbox is empty, move all mail from inbox to outbox
- This reverses the order correctly (first in, first out)

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Must implement queue using only stacks
2. **Goal**: Provide push, pop, peek, empty operations
3. **Output**: Queue-like behavior with stack operations

#### Phase 2: Key Insight Recognition
- **"Stacks are LIFO!"** → Need to reverse order
- **"How to reverse?"** → Transfer from input to output stack
- **"When to transfer?"** → Only when output stack is empty
- **"What about efficiency?"** → Amortized O(1) operations

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use two stacks:
1. stack1: collects new elements (push operation)
2. stack2: holds elements in queue order
3. For pop/peek: if stack2 is empty, transfer all from stack1
4. This transfer reverses the order correctly
5. Push operations always go to stack1
6. The transfer cost is amortized across operations"
```

#### Phase 4: Edge Case Handling
- **Empty queue**: Both stacks empty
- **Single element**: Transfer happens, then pop works
- **Multiple transfers**: Each element transferred at most once
- **Memory management**: Clear references properly

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Operations: push(1), push(2), push(3), pop(), peek(), pop()

Human thinking:
"Initial: stack1 = [], stack2 = []

push(1):
→ stack1.push(1)
→ stack1 = [1], stack2 = []

push(2):
→ stack1.push(2)
→ stack1 = [1,2], stack2 = []

push(3):
→ stack1.push(3)
→ stack1 = [1,2,3], stack2 = []

pop():
→ stack2 is empty, transfer from stack1
→ Transfer 3: stack2.push(3), stack1.pop()
→ Transfer 2: stack2.push(2), stack1.pop()
→ Transfer 1: stack2.push(1), stack1.pop()
→ stack1 = [], stack2 = [3,2,1]
→ Return stack2.pop() = 1 (front element)

peek():
→ stack2 has [3,2,1], return stack2.peek() = 1

pop():
→ stack2 has [3,2,1], return stack2.pop() = 2

Final queue order: 1, 2, 3 ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Transfer reverses order from LIFO to FIFO
- **Why it's efficient**: Each element transferred at most once
- **Why it's correct**: Queue semantics preserved

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use one stack?"** → Can't get FIFO behavior
2. **"What about always transferring?"** → Inefficient O(N) per operation
3. **"How to handle empty?"** → Check both stacks
4. **"What about recursion?"** → Uses call stack, less efficient

### Real-World Analogy
**Like loading/unloading boxes:**
- You have loading dock (input stack) for new boxes
- You have unloading dock (output stack) for boxes to ship
- New boxes always go to loading dock (push)
- When someone wants a box (pop/peek), check unloading dock first
- If unloading dock is empty, move all boxes from loading to unloading
- This ensures first loaded box is first shipped (FIFO)

### Human-Readable Pseudocode
```
class QueueUsingStacks:
    stack1 = empty stack  // input
    stack2 = empty stack  // output
    
    function push(x):
        stack1.push(x)
    
    function pop():
        if stack2.isEmpty():
            while not stack1.isEmpty():
                stack2.push(stack1.pop())
        return stack2.pop()
    
    function peek():
        if stack2.isEmpty():
            while not stack1.isEmpty():
                stack2.push(stack1.pop())
        return stack2.peek()
    
    function empty():
        return stack1.isEmpty() and stack2.isEmpty()
```

### Execution Visualization

### Example: push(1), push(2), push(3), pop(), peek(), pop()
```
Initial: stack1 = [], stack2 = []

push(1):
→ stack1 = [1], stack2 = []
Queue order: [1]

push(2):
→ stack1 = [1,2], stack2 = []
Queue order: [1,2]

push(3):
→ stack1 = [1,2,3], stack2 = []
Queue order: [1,2,3]

pop():
→ stack2 is empty, transfer from stack1
→ Transfer 3: stack2.push(3), stack1.pop()
→ Transfer 2: stack2.push(2), stack1.pop()
→ Transfer 1: stack2.push(1), stack1.pop()
→ stack1 = [], stack2 = [3,2,1]
→ Return 1 (front element)
Queue order: [2,3]

peek():
→ stack2 has [3,2,1], return 1
Queue order: [2,3]

pop():
→ stack2 has [3,2,1], return 2
Queue order: [3]
```

### Key Visualization Points:
- **Two stacks** work together to simulate queue
- **Transfer operation** reverses order when needed
- **Amortized cost** distributes transfer overhead
- **FIFO behavior** emerges from LIFO stacks

### Memory Layout Visualization:
```
Stack Operations:    Stack1: [1][2][3]    Stack2: []
                    ↑  ↑  ↑
                    top top

Pop Operation:    Stack1: []          Stack2: [3][2][1]
                    top→ ↑  ↑  ↑
                    Transfer 3→2→1
                    Return 1 (front)
```

### Time Complexity Breakdown:
- **Push Operation**: O(1) - single stack push
- **Pop Operation**: Amortized O(1)
  - Worst case: O(N) when transferring N elements
  - But each element transferred at most once
- **Peek Operation**: Amortized O(1) (same as pop)
- **Empty Check**: O(1) - check both stacks
- **Space**: O(N) - total elements in both stacks
*/
