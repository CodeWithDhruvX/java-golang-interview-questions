package main

import "fmt"

// 371. Sum of Two Integers
// Time: O(1), Space: O(1)
func getSum(a int, b int) int {
	for b != 0 {
		// Calculate carry bits
		carry := a & b
		
		// Calculate sum bits without carry
		a = a ^ b
		
		// Shift carry to the left
		b = carry << 1
	}
	
	return a
}

// Recursive version
func getSumRecursive(a int, b int) int {
	if b == 0 {
		return a
	}
	
	// Calculate carry and sum
	carry := a & b
	sum := a ^ b
	
	// Recursively add carry to sum
	return getSumRecursive(sum, carry<<1)
}

// Step-by-step explanation version
func getSumWithExplanation(a int, b int) int {
	fmt.Printf("Adding %d and %d using bit manipulation:\n", a, b)
	fmt.Printf("Initial: a=%d (%032b), b=%d (%032b)\n\n", a, b, a, b)
	
	step := 1
	for b != 0 {
		fmt.Printf("Step %d:\n", step)
		
		// Calculate carry (where both bits are 1)
		carry := a & b
		fmt.Printf("  Carry (a & b): %d (%032b)\n", carry, carry)
		
		// Calculate sum without carry (XOR)
		a = a ^ b
		fmt.Printf("  Sum without carry (a ^ b): %d (%032b)\n", a, a)
		
		// Shift carry left to add in next higher bit position
		b = carry << 1
		fmt.Printf("  Shifted carry (carry << 1): %d (%032b)\n", b, b)
		fmt.Printf("  New a: %d, New b: %d\n\n", a, b)
		
		step++
	}
	
	fmt.Printf("Final result: %d\n", a)
	return a
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Bitwise Addition without Arithmetic Operators
- **Bitwise Addition**: Use XOR for sum without carry, AND for carry
- **Carry Propagation**: Shift carry left to add in next position
- **Iterative Process**: Repeat until no carry remains
- **Binary Arithmetic**: Simulate addition at bit level

## 2. PROBLEM CHARACTERISTICS
- **Arithmetic Operation**: Add two integers without + operator
- **Bit Manipulation**: Use only bitwise operators
- **Carry Handling**: Manage overflow between bit positions
- **Iterative Solution**: Process until carry is eliminated

## 3. SIMILAR PROBLEMS
- Subtract Two Numbers (LeetCode 371 variant) - Bitwise subtraction
- Multiply Two Numbers (LeetCode 43) - Bitwise multiplication
- Divide Two Numbers (LeetCode 29) - Bitwise division
- Bit Manipulation arithmetic problems

## 4. KEY OBSERVATIONS
- **XOR Property**: a ^ b gives sum without carry
- **AND Property**: a & b gives positions where both bits are 1 (carry)
- **Carry Shift**: carry << 1 moves carry to next position
- **Termination**: Loop ends when no carry remains

## 5. VARIATIONS & EXTENSIONS
- **Subtraction**: Bitwise subtraction using two's complement
- **Multiplication**: Bitwise multiplication with shift and add
- **Division**: Bitwise division with shift and subtract
- **Different Bases**: Addition in other number bases

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can use arithmetic operators? Handle overflow?"
- Edge cases: zero, negative numbers, overflow
- Time complexity: O(1) for fixed-size integers, O(log max) for arbitrary
- Space complexity: O(1)
- Key insight: simulate binary addition process

## 7. COMMON MISTAKES
- Not handling carry correctly
- Wrong termination condition (infinite loop)
- Not understanding XOR vs AND roles
- Mishandling negative numbers
- Forgetting to shift carry properly

## 8. OPTIMIZATION STRATEGIES
- **Iterative Addition**: O(log max(a,b)) time, O(1) space
- **Recursive Addition**: Same complexity, stack overhead
- **Early Termination**: Stop when carry becomes 0
- **Fixed Optimization**: O(1) for 32-bit integers

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like doing binary addition by hand:**
- You have two binary numbers to add
- At each position, you calculate sum and carry
- Sum without carry: XOR (1+0=1, 1+1=0, 0+1=1, 0+0=0)
- Carry: AND (only 1+1 produces carry)
- Move carry to next position by shifting left
- Repeat until no carry remains

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Two integers a and b
2. **Goal**: Compute sum without using + operator
3. **Constraint**: Use only bitwise operators
4. **Output**: Integer sum of a and b

#### Phase 2: Key Insight Recognition
- **"Binary addition simulation"** → Simulate manual binary addition
- **"XOR for sum"** → a ^ b gives sum without carry
- **"AND for carry"** → a & b gives carry positions
- **"Iterative process"** → Need to propagate carry until done

#### Phase 3: Strategy Development
```
Human thought process:
"I need to add a and b using bitwise operators.
In binary addition:
- Sum bit = a XOR b (without carry)
- Carry bit = a AND b (where both are 1)
- Next carry = carry << 1 (shift to next position)
- Repeat until no carry remains

Algorithm:
while b != 0:
    carry = a & b
    a = a ^ b
    b = carry << 1
return a"
```

#### Phase 4: Edge Case Handling
- **a = 0**: Return b
- **b = 0**: Return a
- **Both zero**: Return 0
- **Negative numbers**: Handle two's complement correctly
- **Overflow**: Depends on integer size

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
a = 5 (0101), b = 3 (0011)

Human thinking:
"I'll simulate binary addition:

Iteration 1:
- carry = a & b = 0101 & 0011 = 0001 = 1
- a = a ^ b = 0101 ^ 0011 = 0110 = 6
- b = carry << 1 = 0001 << 1 = 0010 = 2

Iteration 2:
- carry = a & b = 0110 & 0010 = 0010 = 2
- a = a ^ b = 0110 ^ 0010 = 0100 = 4
- b = carry << 1 = 0010 << 1 = 0100 = 4

Iteration 3:
- carry = a & b = 0100 & 0100 = 0100 = 4
- a = a ^ b = 0100 ^ 0100 = 0000 = 0
- b = carry << 1 = 0100 << 1 = 1000 = 8

Iteration 4:
- carry = a & b = 0000 & 1000 = 0000 = 0
- Loop terminates (b = 0)

Final result: a = 0 ✓ (5 + 3 = 8)"
```

#### Phase 6: Intuition Validation
- **Why XOR works**: Gives sum without considering carry
- **Why AND works**: Identifies positions where carry occurs
- **Why shift left**: Moves carry to next higher position
- **Why loop terminates**: b becomes 0 when no more carry

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use a + b?"** → Problem constraint forbids arithmetic operators
2. **"Should I use OR instead of XOR?"** → OR doesn't give correct sum without carry
3. **"What about negative numbers?"** → Need to handle two's complement
4. **"Can I optimize further?"** → Already optimal for arbitrary precision
5. **"What about overflow?"** → Depends on language/integer size

### Real-World Analogy
**Like adding numbers using only basic logic gates:**
- You have two numbers in binary form
- You can only use XOR, AND, and SHIFT operations
- XOR acts like addition without carry
- AND acts like carry detection
- SHIFT moves carry to next position
- Repeat until no more carries

### Human-Readable Pseudocode
```
function getSum(a, b):
    while b != 0:
        carry = a & b
        a = a ^ b
        b = carry << 1
    
    return a
```

### Execution Visualization

### Example: a = 5 (0101), b = 3 (0011)
```
Binary Addition Simulation:
Initial: a = 0101, b = 0011

Iteration 1:
carry = a & b = 0101 & 0011 = 0001 = 1
a = a ^ b = 0101 ^ 0011 = 0110 = 6
b = carry << 1 = 0001 << 1 = 0010 = 2

Iteration 2:
carry = a & b = 0110 & 0010 = 0010 = 2
a = a ^ b = 0110 ^ 0010 = 0100 = 4
b = carry << 1 = 0010 << 1 = 0100 = 4

Iteration 3:
carry = a & b = 0100 & 0100 = 0100 = 4
a = a ^ b = 0100 ^ 0100 = 0000 = 0
b = carry << 1 = 0100 << 1 = 1000 = 8

Iteration 4:
carry = a & b = 0000 & 1000 = 0000 = 0
Loop terminates (b = 0)

Final result: a = 0 ✓ (5 + 3 = 8)
```

### Key Visualization Points:
- **XOR Operation**: Sum without carry
- **AND Operation**: Carry detection
- **Shift Operation**: Carry propagation
- **Termination**: Loop ends when no carry remains

### Memory Layout Visualization:
```
Bit State During Addition:
a = 5 (0101), b = 3 (0011)

Step-by-step evolution:
Step 1: carry = 0001, a = 0110, b = 0010
Step 2: carry = 0010, a = 0100, b = 0100
Step 3: carry = 0100, a = 0000, b = 1000
Step 4: carry = 0000, a = 0000, b = 0000

Final: a = 0000 = 8 ✓
```

### Time Complexity Breakdown:
- **Fixed-size Integers**: O(1) time (constant number of bits)
- **Arbitrary Precision**: O(log max(a,b)) time
- **Space Complexity**: O(1) additional space
- **Bit Operations**: Each iteration uses constant-time operations

### Alternative Approaches:

#### 1. Recursive Addition (O(log n) time, O(log n) space)
```go
func getSumRecursive(a, b int) int {
    if b == 0 {
        return a
    }
    
    carry := a & b
    sum := a ^ b
    return getSumRecursive(sum, carry<<1)
}
```
- **Pros**: Elegant recursive formulation
- **Cons**: Stack overhead, same time complexity

#### 2. Lookup Table Addition (O(1) time, O(2ⁿ) space)
```go
func getSumLookup(a, b int) int {
    // Precompute all possible sums for small numbers
    // Not practical for general case
    return a + b // Fallback to normal addition
}
```
- **Pros**: O(1) time for small ranges
- **Cons**: Exponential space, not practical

#### 3. Parallel Bit Addition (O(log n) time, O(1) space)
```go
func getSumParallel(a, b int) int {
    // Use parallel prefix techniques
    // Complex implementation for marginal gains
    return getSum(a, b) // Use standard approach
}
```
- **Pros**: Theoretical parallel speedup
- **Cons**: Much more complex, minimal practical benefit

### Extensions for Interviews:
- **Subtraction**: Bitwise subtraction using two's complement
- **Multiplication**: Bitwise multiplication algorithms
- **Division**: Bitwise division algorithms
- **Different Bases**: Addition in other number systems
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := [][2]int{
		{1, 2},
		{2, 3},
		{0, 1},
		{-1, 1},
		{-2, -3},
		{10, 20},
		{100, 200},
		{-10, 20},
		{123, 456},
		{999, 1},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		
		if i == 0 {
			// Show detailed explanation for first test case
			result := getSumWithExplanation(tc[0], tc[1])
			fmt.Printf("Iterative result: %d\n\n", result)
		} else {
			result1 := getSum(tc[0], tc[1])
			result2 := getSumRecursive(tc[0], tc[1])
			fmt.Printf("Adding %d + %d\n", tc[0], tc[1])
			fmt.Printf("Iterative: %d, Recursive: %d\n\n", result1, result2)
		}
	}
}
