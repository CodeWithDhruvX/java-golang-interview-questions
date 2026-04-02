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
