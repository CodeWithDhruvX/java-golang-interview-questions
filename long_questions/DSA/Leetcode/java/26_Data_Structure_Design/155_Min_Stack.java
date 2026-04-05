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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Data Structure Design
- **Min Stack**: Stack with O(1) minimum retrieval
- **Auxiliary Structure**: Additional stack to track minimums
- **Space-Time Tradeoff**: Extra space for constant time minimum
- **Multiple Implementations**: Different optimization strategies

## 2. PROBLEM CHARACTERISTICS
- **Stack Operations**: push, pop, top, getMin in O(1)
- **Minimum Tracking**: Need efficient minimum element retrieval
- **Constraint Handling**: Handle empty stack gracefully
- **Performance Requirements**: All operations should be O(1)

## 3. SIMILAR PROBLEMS
- Max Stack (with O(1) maximum)
- Queue with O(1) minimum
- Stack with O(1) middle element
- Design Min Queue

## 4. KEY OBSERVATIONS
- Regular stack getMin() requires O(N) traversal
- Auxiliary stack stores minimums efficiently
- Different strategies: difference tracking, pair storage, single stack
- Space complexity: O(N) for auxiliary structure
- Time complexity: O(1) for all operations

## 5. VARIATIONS & EXTENSIONS
- Support for duplicate minimums
- Stack with range queries
- Thread-safe implementation
- Memory-constrained versions

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I handle negative values?"
- Edge cases: empty stack, single element, duplicates
- Time complexity: O(1) vs O(N) regular stack
- Space complexity: O(N) auxiliary vs O(1) regular

## 7. COMMON MISTAKES
- Not handling empty stack properly
- Incorrect minimum updates in auxiliary stack
- Integer overflow with long differences
- Not maintaining O(1) for all operations

## 8. OPTIMIZATION STRATEGIES
- Lazy minimum updates
- In-place calculations where possible
- Early termination for obvious cases
- Choose appropriate strategy based on usage pattern

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like a smart filing cabinet:**
- You have a regular stack of papers (main stack)
- You also have a separate index card showing the minimum value
- When you add new papers, you update the index card
- When you remove papers, you check if the removed paper was the minimum
- If yes, you need to find the new minimum from remaining papers
- This gives you O(1) access to the minimum value

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Need stack with O(1) minimum operations
2. **Goal**: Implement push, pop, top, getMin efficiently
3. **Output**: Stack data structure with constant-time minimum

#### Phase 2: Key Insight Recognition
- **"Why regular stack insufficient?"** → getMin() requires O(N) traversal
- **"How to track minimum efficiently?"** → Auxiliary structure
- **"What's the tradeoff?"** → Extra space for O(1) minimum access
- **"Why multiple approaches?"** → Different optimization strategies

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use auxiliary stack:
1. Main stack stores all elements
2. Min stack stores minimum values efficiently
3. Push: update both stacks appropriately
4. Pop: remove from main, update min stack if needed
5. getMin: return top of min stack
6. This gives O(1) for all operations"
```

#### Phase 4: Edge Case Handling
- **Empty stack**: Return error or sentinel value
- **Single element**: Both stacks have same value
- **Duplicates**: Handle multiple minimums correctly
- **Large values**: Use long to avoid overflow

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Operations: push(5), push(2), push(7), getMin(), pop(), getMin()

Human thinking:
"Let's track both stacks:

Initial: main=[], min=[]

Push 5:
- main: [5], min: [5]
- min = 5

Push 2:
- main: [5,2], min: [5,2]
- min = 2

Push 7:
- main: [5,2,7], min: [5,2,7]
- min = 2

getMin(): return 2 ✓

Pop():
- Remove 7 from main: [5,2]
- 7 was not the minimum, min stack unchanged
- min = 2

getMin(): return 2 ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Auxiliary stack always has current minimum
- **Why it's efficient**: All operations are O(1)
- **Why it's correct**: Minimum tracking is maintained accurately

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just scan main stack?"** → O(N) time for getMin()
2. **"What about recalculation?"** → O(N²) if done naively
3. **"How to handle duplicates?"** → Need proper logic in auxiliary stack
4. **"What about memory?"** → Trade space for time efficiency

### Real-World Analogy
**Like a priority checkout line:**
- You have a regular line of customers (main stack)
- You also have an express lane for VIP customers (min stack)
- When new customers arrive, you decide if they're VIP enough
- VIP customers go to express lane (min stack)
- Regular customers go to regular lane (main stack)
- When a VIP customer leaves, you check if they were the minimum VIP
- If yes, find the next minimum VIP from the express lane
- This gives you O(1) access to the minimum priority customer

### Human-Readable Pseudocode
```
class MinStack:
    mainStack = empty stack
    minStack = empty stack
    
    function push(val):
        mainStack.push(val)
        if minStack.isEmpty() or val <= minStack.peek():
            minStack.push(val)
        else:
            diff = val - minStack.peek()
            minStack.push(diff)
    
    function pop():
        val = mainStack.pop()
        if !minStack.isEmpty() and val == minStack.peek():
            minStack.pop()
        else if !minStack.isEmpty():
            diff = minStack.pop()
            min = min - diff
            minStack.push(diff)
        return val
    
    function getMin():
        if minStack.isEmpty():
            return -1  // or error
        return minStack.peek()
```

### Execution Visualization

### Example: push(5), push(2), push(7), getMin(), pop(), getMin()
```
Stack Evolution:
Initial: main=[], min=[]

Push 5:
main: [5], min: [5], min=5

Push 2:
main: [5,2], min: [5,2], min=2

Push 7:
main: [5,2,7], min: [5,2,7], min=2

getMin(): return 2 ✓

Pop():
- Remove 7 from main: [5,2]
- 7 ≠ min(2), min stack unchanged
- min = 2

getMin(): return 2 ✓

Visualization:
Main Stack: [5,2,7] ← 7 removed
Min Stack: [5,2] ← minimum remains 2
```

### Key Visualization Points:
- **Dual stack** approach for O(1) minimum
- **Difference tracking** maintains relative values
- **Lazy updates** only when necessary
- **Space-time tradeoff** for constant-time operations

### Memory Layout Visualization:
```
Two-Stack Approach:
Main Stack: [5,2,7] (bottom)
Min Stack: [5,2] (top)

Operation Flow:
Push: main.push(), min.push() if needed
Pop: main.pop(), min.pop() if popped was min
getMin: min.peek()

State Tracking:
- Main stack: actual values
- Min stack: minimum values or differences
- Relationship: min.peek() = actual minimum in main stack
```

### Time Complexity Breakdown:
- **Push**: O(1) - push to main, possibly push to min
- **Pop**: O(1) - pop from main, possibly pop from min
- **getMin**: O(1) - peek from min stack
- **Space**: O(N) for auxiliary stack
- **Optimal**: O(1) time for all operations with O(N) space
- **vs Regular Stack**: O(N) time for getMin(), O(N) space
*/
