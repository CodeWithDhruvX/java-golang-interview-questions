package main

import "fmt"

// 155. Min Stack
// Time: O(1) for all operations, Space: O(N)
type MinStack struct {
	stack    []int
	minStack []int
}

// Constructor initializes the stack
func ConstructorMinStack() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

// Push pushes element onto stack
func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	
	// Push to minStack if it's empty or val is <= current min
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

// Pop removes the element on top of the stack
func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	
	// Remove from minStack if it matches the popped element
	if len(this.minStack) > 0 && top == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

// Top gets the top element of the stack
func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1 // Or handle error appropriately
	}
	return this.stack[len(this.stack)-1]
}

// GetMin retrieves the minimum element in the stack
func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return -1 // Or handle error appropriately
	}
	return this.minStack[len(this.minStack)-1]
}

// Alternative implementation using single stack with min tracking
type MinStackOptimized struct {
	stack []struct {
		val int
		min int
	}
}

func ConstructorMinStackOptimized() MinStackOptimized {
	return MinStackOptimized{
		stack: []struct {
			val int
			min int
		}{},
	}
}

func (this *MinStackOptimized) Push(val int) {
	min := val
	if len(this.stack) > 0 {
		currentMin := this.stack[len(this.stack)-1].min
		if val > currentMin {
			min = currentMin
		}
	}
	
	this.stack = append(this.stack, struct {
		val int
		min int
	}{val, min})
}

func (this *MinStackOptimized) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStackOptimized) Top() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1].val
}

func (this *MinStackOptimized) GetMin() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1].min
}

// Alternative implementation using difference method
type MinStackDiff struct {
	stack []int
	min   int
}

func ConstructorMinStackDiff() MinStackDiff {
	return MinStackDiff{
		stack: []int{},
		min:   0,
	}
}

func (this *MinStackDiff) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.min = val
	} else {
		diff := val - this.min
		this.stack = append(this.stack, diff)
		if val < this.min {
			this.min = val
		}
	}
}

func (this *MinStackDiff) Pop() {
	if len(this.stack) == 0 {
		return
	}
	
	diff := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	
	if diff < 0 {
		this.min = this.min - diff
	}
}

func (this *MinStackDiff) Top() int {
	if len(this.stack) == 0 {
		return -1
	}
	
	diff := this.stack[len(this.stack)-1]
	if diff < 0 {
		return this.min
	}
	return this.min + diff
}

func (this *MinStackDiff) GetMin() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.min
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Stack with Minimum Tracking
- **Dual Stack Approach**: Main stack + min stack for O(1) minimum retrieval
- **Optimized Single Stack**: Store values with running minimums
- **Difference Method**: Track differences from minimum for efficiency
- **O(1) Minimum**: Retrieve minimum in constant time

## 2. PROBLEM CHARACTERISTICS
- **Stack Operations**: Push, pop, top, getMin in O(1) time
- **Minimum Tracking**: Need efficient minimum retrieval
- **Space Complexity**: Trade-off between time and space efficiency
- **Multiple Implementations**: Different approaches for min tracking

## 3. SIMILAR PROBLEMS
- Implement Queue using Stacks (LeetCode 232) - Queue with two stacks
- Max Stack (LeetCode 155) - Track maximum instead of minimum
- Stack with GetMiddle - Implement efficient middle element access
- Min Queue (LeetCode 622) - Queue with minimum tracking

## 4. KEY OBSERVATIONS
- **Min Stack Logic**: Maintain separate stack of minimums
- **Space Trade-off**: Extra space for O(1) minimum operations
- **Optimization Opportunities**: Multiple ways to track minimum
- **Edge Cases**: Empty stack, duplicate minimums, single element

## 5. VARIATIONS & EXTENSIONS
- **Max Stack**: Track maximum instead of minimum
- **Min-Max Stack**: Track both minimum and maximum
- **GetMedian**: Support for median calculation
- **Custom Comparator**: Support for custom comparison functions

## 6. INTERVIEW INSIGHTS
- Always clarify: "Time complexity requirements? Space constraints?"
- Edge cases: empty stack, single element, all same values
- Time complexity: O(1) for all operations
- Space complexity: O(N) where N=stack size
- Key insight: trade space for O(1) minimum retrieval

## 7. COMMON MISTAKES
- Not handling empty stack correctly
- Wrong min stack update logic
- Inefficient minimum tracking (O(N) instead of O(1))
- Not updating min stack on pop operations
- Off-by-one errors in stack indexing

## 8. OPTIMIZATION STRATEGIES
- **Dual Stack**: O(1) time, O(N) space - standard approach
- **Optimized Single Stack**: O(1) time, O(N) space - space-efficient
- **Difference Method**: O(1) time, O(N) space - mathematical approach
- **Lazy Evaluation**: Compute minimum only when needed

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like a stack with a separate minimum tracker:**
- You have a regular stack for push/pop operations
- You also maintain a separate structure for minimums
- Like having a stack and a separate "minimum so far" display
- When you push, you update both the stack and minimum tracker
- When you pop, you remove from both if necessary

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Stack with push, pop, top, getMin operations
2. **Goal**: Implement stack where getMin() is O(1) time
3. **Constraints**: All operations should be efficient
4. **Output**: Stack with O(1) minimum retrieval

#### Phase 2: Key Insight Recognition
- **"Dual stack natural fit"** → Separate min stack for O(1) access
- **"Space-time trade-off"** → Extra space for faster minimum access
- **"Multiple approaches"** → Several ways to track minimum efficiently
- **"Edge cases critical"** → Empty stack handling is crucial

#### Phase 3: Strategy Development
```
Human thought process:
"I need a stack with O(1) getMin operation.
Regular stack getMin() would be O(N), too slow.

Dual Stack Approach:
1. Main stack: regular push/pop operations
2. Min stack: track minimums seen so far
3. Push: add to main stack, update min stack
4. Pop: remove from main stack, remove from min stack if needed
5. GetMin: return top of min stack

Optimized Single Stack:
1. Store value + current minimum in each stack node
2. Update running minimum on push
3. GetMin: return current minimum directly

Both approaches give O(1) getMin!"
```

#### Phase 4: Edge Case Handling
- **Empty stack**: Return appropriate values (often -1 or error)
- **Single element**: Both stack and min stack have same element
- **Duplicate minimums**: Handle multiple equal minimums correctly
- **Large values**: Handle potential overflow issues

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Operations: Push(5), Push(3), Push(7), GetMin(), Pop(), GetMin()

Human thinking:
"Dual Stack Approach:
1. Push(5): main=[5], min=[5]
2. Push(3): main=[5,3], min=[3] (3 < 5)
3. Push(7): main=[5,3,7], min=[3] (7 > 3)
4. GetMin(): return min.top() = 3 ✓
5. Pop(): remove 7 from main, min unchanged
   main=[5,3], min=[3]
6. GetMin(): return min.top() = 3 ✓

Optimized Single Stack:
1. Push(5): stack=[{val:5, min:5}]
2. Push(3): stack=[{val:5, min:5}, {val:3, min:3}]
3. Push(7): stack=[{val:5, min:5}, {val:3, min:3}, {val:7, min:3}]
4. GetMin(): return stack.top().min = 3 ✓
5. Pop(): remove {val:7, min:3}
   stack=[{val:5, min:5}, {val:3, min:3}]
6. GetMin(): return stack.top().min = 3 ✓

Both approaches work correctly!"
```

#### Phase 6: Intuition Validation
- **Why dual stack works**: Separate min stack tracks history of minimums
- **Why optimized works**: Each node carries its own minimum information
- **Why O(1)**: Direct access to minimum without traversal
- **Why space trade-off**: Extra space enables constant time minimum

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just traverse stack?"** → O(N) time is too slow
2. **"Should I use heap?"** → Heap doesn't maintain LIFO order
3. **"What about getMedian?"** → Different problem, needs different approach
4. **"Can I optimize space?"** → Trade-offs between time and space
5. **"What about duplicates?"** → Handle equal minimums correctly

### Real-World Analogy
**Like a stack with a separate minimum display:**
- You have a stack of items (main stack)
- You also have a separate display showing current minimum
- When you add items, you update both the stack and minimum display
- When you remove items, you update both if needed
- Like a warehouse inventory system with current minimum price display
- Minimum display updates automatically as items are added/removed

### Human-Readable Pseudocode
```
class MinStack:
    stack = []
    minStack = []
    
    function push(val):
        stack.push(val)
        if minStack is empty or val <= minStack.top():
            minStack.push(val)
    
    function pop():
        if stack is empty:
            return error
        val = stack.pop()
        if val == minStack.top():
            minStack.pop()
        return val
    
    function top():
        if stack is empty:
            return error
        return stack.top()
    
    function getMin():
        if minStack is empty:
            return error
        return minStack.top()

// Optimized single stack version
class MinStackOptimized:
    stack = []
    
    function push(val):
        min = val if stack is empty else min(val, stack.top().min)
        stack.push({val: val, min: min})
    
    function pop():
        if stack is empty:
            return error
        return stack.pop().val
    
    function top():
        if stack is empty:
            return error
        return stack.top().val
    
    function getMin():
        if stack is empty:
            return error
        return stack.top().min
```

### Execution Visualization

### Example: Operations: Push(5), Push(3), Push(7), GetMin(), Pop(), GetMin()
```
Dual Stack Approach:
Initial: stack=[], minStack=[]

Push(5): stack=[5], minStack=[5]
Push(3): stack=[5,3], minStack=[3] (3 < 5)
Push(7): stack=[5,3,7], minStack=[3] (7 > 3)
GetMin(): return minStack.top() = 3 ✓
Pop(): remove 7 from stack, minStack unchanged
stack=[5,3], minStack=[3]
GetMin(): return minStack.top() = 3 ✓

Optimized Single Stack:
Initial: stack=[]

Push(5): stack=[{val:5, min:5}]
Push(3): stack=[{val:5, min:5}, {val:3, min:3}]
Push(7): stack=[{val:5, min:5}, {val:3, min:3}, {val:7, min:3}]
GetMin(): return stack.top().min = 3 ✓
Pop(): remove {val:7, min:3}
stack=[{val:5, min:5}, {val:3, min:3}]
GetMin(): return stack.top().min = 3 ✓
```

### Key Visualization Points:
- **Dual Stack**: Main stack + separate min stack
- **Min Stack Logic**: Only push when new value is ≤ current minimum
- **Optimized Stack**: Each node carries its own minimum
- **O(1) Access**: Direct access to minimum without traversal

### Memory Layout Visualization:
```
Dual Stack State Evolution:
Initial: stack=[], minStack=[]
After Push(5): stack=[5], minStack=[5]
After Push(3): stack=[5,3], minStack=[3]
After Push(7): stack=[5,3,7], minStack=[3]
After Pop(): stack=[5,3], minStack=[3]

Optimized Stack State Evolution:
Initial: stack=[]
After Push(5): stack=[{val:5, min:5}]
After Push(3): stack=[{val:5, min:5}, {val:3, min:3}]
After Push(7): stack=[{val:5, min:5}, {val:3, min:3}, {val:7, min:3}]
After Pop(): stack=[{val:5, min:5}, {val:3, min:3}]

Operation Complexity:
Push: O(1) time, O(1) space
Pop: O(1) time, O(1) space
Top: O(1) time, O(1) space
GetMin: O(1) time, O(1) space
```

### Time Complexity Breakdown:
- **Push**: O(1) time, O(1) space for all approaches
- **Pop**: O(1) time, O(1) space for all approaches
- **Top**: O(1) time, O(1) space for all approaches
- **GetMin**: O(1) time, O(1) space (key optimization)
- **Space**: O(N) where N=stack size

### Alternative Approaches:

#### 1. Lazy Minimum Calculation (O(1) amortized time)
```go
type MinStackLazy struct {
    stack []int
    min   int
    dirty bool
}

func (this *MinStackLazy) GetMin() int {
    if this.dirty {
        this.recalculateMin()
        this.dirty = false
    }
    return this.min
}
```
- **Pros**: Potentially more space efficient
- **Cons**: More complex implementation

#### 2. Segment Tree (O(log N) time, O(N) space)
```go
type MinStackSegment struct {
    tree *SegmentTree
    stack []int
}

func (this *MinStackSegment) GetMin() int {
    return this.tree.Query(0, len(this.stack)-1)
}
```
- **Pros**: Efficient for range queries
- **Cons**: Overkill for simple getMin operation

#### 3. Binary Indexed Tree (O(log N) time, O(N) space)
```go
type MinStackBIT struct {
    bit *BinaryIndexedTree
    stack []int
}

func (this *MinStackBIT) GetMin() int {
    return this.bit.QueryMin()
}
```
- **Pros**: Efficient for dynamic updates
- **Cons**: Complex implementation

### Extensions for Interviews:
- **Max Stack**: Track maximum instead of minimum
- **Min-Max Stack**: Track both minimum and maximum
- **GetMedian**: Support for median calculation
- **Custom Comparator**: Support for custom comparison functions
- **Performance Analysis**: Discuss space-time trade-offs
*/
func main() {
	// Test cases
	fmt.Println("=== Testing MinStack ===")
	
	// Test 1: Basic operations
	minStack := ConstructorMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Printf("After pushing -2, 0, -3:\n")
	fmt.Printf("  Top: %d, Min: %d\n", minStack.Top(), minStack.GetMin())
	
	minStack.Pop()
	fmt.Printf("After pop:\n")
	fmt.Printf("  Top: %d, Min: %d\n", minStack.Top(), minStack.GetMin())
	
	// Test 2: Edge cases
	fmt.Println("\n=== Testing Edge Cases ===")
	emptyStack := ConstructorMinStack()
	fmt.Printf("Empty stack - Top: %d, Min: %d\n", emptyStack.Top(), emptyStack.GetMin())
	
	emptyStack.Push(5)
	fmt.Printf("After pushing 5 - Top: %d, Min: %d\n", emptyStack.Top(), emptyStack.GetMin())
	
	// Test 3: Duplicate minimums
	fmt.Println("\n=== Testing Duplicate Minimums ===")
	dupStack := ConstructorMinStack()
	dupStack.Push(1)
	dupStack.Push(1)
	dupStack.Push(2)
	fmt.Printf("After pushing 1, 1, 2 - Top: %d, Min: %d\n", dupStack.Top(), dupStack.GetMin())
	
	dupStack.Pop()
	fmt.Printf("After pop - Top: %d, Min: %d\n", dupStack.Top(), dupStack.GetMin())
	
	// Test optimized versions
	fmt.Println("\n=== Testing Optimized Version ===")
	optStack := ConstructorMinStackOptimized()
	optStack.Push(-2)
	optStack.Push(0)
	optStack.Push(-3)
	fmt.Printf("After pushing -2, 0, -3:\n")
	fmt.Printf("  Top: %d, Min: %d\n", optStack.Top(), optStack.GetMin())
	
	optStack.Pop()
	fmt.Printf("After pop:\n")
	fmt.Printf("  Top: %d, Min: %d\n", optStack.Top(), optStack.GetMin())
	
	// Test difference method
	fmt.Println("\n=== Testing Difference Method ===")
	diffStack := ConstructorMinStackDiff()
	diffStack.Push(-2)
	diffStack.Push(0)
	diffStack.Push(-3)
	fmt.Printf("After pushing -2, 0, -3:\n")
	fmt.Printf("  Top: %d, Min: %d\n", diffStack.Top(), diffStack.GetMin())
	
	diffStack.Pop()
	fmt.Printf("After pop:\n")
	fmt.Printf("  Top: %d, Min: %d\n", diffStack.Top(), diffStack.GetMin())
}
