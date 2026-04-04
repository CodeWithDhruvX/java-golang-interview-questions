package main

import "fmt"

// 191. Number of 1 Bits
// Time: O(1) for 32-bit integers, Space: O(1)
func hammingWeight(num uint32) uint32 {
	count := uint32(0)
	
	// Brian Kernighan's algorithm
	for num != 0 {
		num &= num - 1 // Clear the least significant bit set
		count++
	}
	
	return count
}

// Simple loop approach
func hammingWeightSimple(num uint32) uint32 {
	count := uint32(0)
	
	for i := 0; i < 32; i++ {
		count += num & 1
		num >>= 1
	}
	
	return count
}

// Optimized with bit tricks
func hammingWeightOptimized(num uint32) uint32 {
	num = num - ((num >> 1) & 0x55555555)
	num = (num & 0x33333333) + ((num >> 2) & 0x33333333)
	num = (num + (num >> 4)) & 0x0F0F0F0F
	num = num + (num >> 8)
	num = num + (num >> 16)
	
	return num & 0x3F
}

// Built-in function style (if available)
func hammingWeightBuiltIn(num uint32) uint32 {
	// In Go, we can use bits.OnesCount from Go 1.9+
	// But implementing manually for educational purposes
	count := uint32(0)
	for num != 0 {
		num &= num - 1
		count++
	}
	return count
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Bit Counting with Brian Kernighan's Algorithm
- **Brian Kernighan's Algorithm**: n &= (n-1) clears least significant set bit
- **Bit Counting**: Count set bits by repeatedly clearing LSb
- **Efficient Loop**: Runs only for set bits, not all 32 bits
- **Bit Manipulation**: Leverage n & (n-1) property

## 2. PROBLEM CHARACTERISTICS
- **Population Count**: Count number of 1s in binary representation
- **32-bit Integer**: Standard integer size optimization
- **Efficient Counting**: Need faster than checking all 32 bits
- **Bit-Level Operations**: Direct bit manipulation required

## 3. SIMILAR PROBLEMS
- Counting Bits (LeetCode 338) - Same core problem
- Reverse Bits (LeetCode 190) - Bit manipulation
- Single Number (LeetCode 136) - XOR properties
- Bit Manipulation problems using bit tricks

## 4. KEY OBSERVATIONS
- **LSB Clearing**: n & (n-1) clears the least significant set bit
- **Iteration Count**: Each iteration removes one set bit
- **Termination**: Loop ends when n becomes 0
- **Efficiency**: O(number of set bits) instead of O(32)

## 5. VARIATIONS & EXTENSIONS
- **Different Integer Sizes**: 64-bit, 16-bit integers
- **Population Count**: Built-in hardware instructions
- **Parallel Counting**: Count multiple bits simultaneously
- **Streaming Bits**: Handle bit streams efficiently

## 6. INTERVIEW INSIGHTS
- Always clarify: "What integer size? Signed/unsigned?"
- Edge cases: 0, negative numbers, maximum values
- Time complexity: O(set bits) time, O(1) space
- Key insight: n & (n-1) removes LSb efficiently
- Alternative: Built-in popcount functions

## 7. COMMON MISTAKES
- Using wrong bit operation (n >> 1 instead of n & (n-1))
- Not handling negative numbers correctly
- Infinite loop with wrong termination condition
- Not understanding two's complement for negative numbers
- Off-by-one errors in bit counting

## 8. OPTIMIZATION STRATEGIES
- **Brian Kernighan**: O(set bits) time, O(1) space - optimal
- **Lookup Table**: O(1) time, O(256) space for 8-bit chunks
- **Parallel Counting**: O(log bits) time, O(1) space
- **Built-in Functions**: O(1) time using hardware popcount

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like counting set bits by removing them one by one:**
- You have a number and want to count its 1s
- Each iteration, you remove the rightmost 1
- Count how many times you can remove a 1
- Brian Kernighan's trick efficiently finds and removes the rightmost 1

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: 32-bit integer n
2. **Goal**: Count number of 1s in binary representation
3. **Constraint**: Use bit manipulation efficiently
4. **Output**: Count of set bits

#### Phase 2: Key Insight Recognition
- **"Bit clearing natural fit"** → n & (n-1) clears LSb
- **"Iteration efficiency"** → Only iterate for set bits
- **"Brian Kernighan's algorithm"** → Classic bit counting technique
- **"Termination condition"** → Loop ends when n becomes 0

#### Phase 3: Strategy Development
```
Human thought process:
"I need to count set bits in n.
Instead of checking all 32 bits, I can use Brian Kernighan's trick:
1. While n is not 0:
   - Count++ (found a set bit)
   - n = n & (n-1) (clear the rightmost set bit)
2. Return count

This only loops for set bits, not all bits!"
```

#### Phase 4: Edge Case Handling
- **n = 0**: Return 0 (no set bits)
- **Negative numbers**: Handle two's complement correctly
- **Large numbers**: Ensure 32-bit handling
- **Maximum value**: All 32 bits set

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
n = 13 (1101 in binary)

Human thinking:
"I'll use Brian Kernighan's algorithm:

Iteration 1: n = 13 (1101)
- Count = 1
- n = 13 & 12 = 1101 & 1100 = 1100 = 12

Iteration 2: n = 12 (1100)
- Count = 2
- n = 12 & 11 = 1100 & 1011 = 1000 = 8

Iteration 3: n = 8 (1000)
- Count = 3
- n = 8 & 7 = 1000 & 0111 = 0000 = 0

Iteration 4: n = 0, terminate
Final count: 3 ✓"

Let me trace n = 0:
Iteration 1: n = 0, terminate immediately
Final count: 0 ✓
```

#### Phase 6: Intuition Validation
- **Why n & (n-1) works**: Clears rightmost set bit efficiently
- **Why O(set bits)**: Each iteration removes one set bit
- **Why efficient**: Better than checking all 32 bits
- **Why optimal**: Can't do better than O(set bits)

### Common Human Pitfalls & How to Avoid Them
1. **"Why not check all bits?"** → O(32) vs O(set bits), Brian Kernighan is better
2. **"Should I use n >> 1?"** → Doesn't clear the bit, just shifts
3. **"What about negative numbers?"** → Need to handle two's complement
4. **"Can I optimize further?"** → Built-in popcount is best if available
5. **"What about 64-bit?"** → Same principle, just more bits

### Real-World Analogy
**Like counting lit bulbs in a row by turning them off:**
- You have a row of light bulbs (bits), some are on (1s)
- You want to count how many are on
- Each time you find a lit bulb, you turn it off
- Count how many bulbs you turn off
- Brian Kernighan's trick efficiently finds and turns off the rightmost lit bulb

### Human-Readable Pseudocode
```
function hammingWeight(n):
    if n == 0:
        return 0
    
    count = 0
    while n != 0:
        count++
        n = n & (n-1)  // Clear rightmost set bit
    
    return count
```

### Execution Visualization

### Example: n = 13 (1101 in binary)
```
Bit Evolution During Processing:
Initial: n = 13 (1101), count = 0

Iteration 1: n = 13 & 12 = 1101 & 1100 = 1100 = 12, count = 1
Iteration 2: n = 12 & 11 = 1100 & 1011 = 1000 = 8, count = 2
Iteration 3: n = 8 & 7 = 1000 & 0111 = 0000 = 0, count = 3

Final count: 3 ✓
```

### Key Visualization Points:
- **Bit Clearing**: n & (n-1) clears rightmost set bit
- **Iteration Count**: Each iteration removes one set bit
- **Termination**: Loop ends when all bits cleared
- **Efficiency**: Only processes set bits

### Memory Layout Visualization:
```
Bit State During Processing:
n = 13 (1101)

Step 1: n = 1101, count = 0
Operation: n = 1101 & (1101-1) = 1101 & 1100 = 1100
Result: n = 1100, count = 1

Step 2: n = 1100, count = 1
Operation: n = 1100 & (1100-1) = 1100 & 1011 = 1000
Result: n = 1000, count = 2

Step 3: n = 1000, count = 2
Operation: n = 1000 & (1000-1) = 1000 & 0111 = 0000
Result: n = 0000, count = 3

Final: count = 3 ✓
```

### Time Complexity Breakdown:
- **Brian Kernighan**: O(set bits) time, O(1) space
- **Simple Loop**: O(32) time, O(1) space
- **Lookup Table**: O(1) time, O(256) space
- **Built-in Popcount**: O(1) time, O(1) space

### Alternative Approaches:

#### 1. Simple Loop (O(32) time, O(1) space)
```go
func hammingWeightSimple(num uint32) uint32 {
    count := uint32(0)
    for i := 0; i < 32; i++ {
        count += num & 1
        num >>= 1
    }
    return count
}
```
- **Pros**: Simple to understand
- **Cons**: Always 32 iterations, even for small numbers

#### 2. Lookup Table (O(1) time, O(256) space)
```go
func hammingWeightLookup(num uint32) uint32 {
    // Precomputed table for 8-bit chunks
    lookupTable := [256]uint32{
        0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
        // ... (complete 256 entries)
    }
    
    return lookupTable[num&0xFF] + 
           lookupTable[(num>>8)&0xFF] + 
           lookupTable[(num>>16)&0xFF] + 
           lookupTable[(num>>24)&0xFF]
}
```
- **Pros**: O(1) time
- **Cons**: Requires 256-entry table

#### 3. Parallel Counting (O(log bits) time, O(1) space)
```go
func hammingWeightParallel(num uint32) uint32 {
    // Count 2 bits at a time using lookup
    lookupTable := [16]uint32{
        0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
    }
    
    num = num - ((num >> 1) & 0x55555555)
    num = (num & 0x33333333) + ((num >> 2) & 0x33333333)
    num = (num + (num >> 4)) & 0x0F0F0F0F
    num = num + (num >> 8)
    
    return lookupTable[num&0xF] + 
           lookupTable[(num>>4)&0xF] + 
           lookupTable[(num>>8)&0xF] + 
           lookupTable[(num>>12)&0xF] + 
           lookupTable[(num>>16)&0xF] + 
           lookupTable[(num>>20)&0xF] + 
           lookupTable[(num>>24)&0xF] + 
           lookupTable[(num>>28)&0xF]
}
```
- **Pros**: Faster than Brian Kernighan for dense bit patterns
- **Cons**: More complex implementation

### Extensions for Interviews:
- **Different Integer Sizes**: 64-bit, 16-bit integers
- **Population Count**: Built-in hardware instructions
- **Parallel Counting**: Count multiple bits simultaneously
- **Streaming Bits**: Handle bit streams efficiently
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []uint32{
		0b00000000000000000000000000000001011, // 3
		0b00000000000000000000000000010000000, // 1
		0b11111111111111111111111111111111101, // 31
		0,                                    // 0
		0b10000000000000000000000000000000000, // 1
		0b11111111111111111111111111111111111, // 32
		0b00000000000000000000000000000000001, // 1
		0b00000000000000000000000000000001010, // 2
		0b10101010101010101010101010101010101, // 16
		0b11001100110011001100110011001100110, // 16
	}
	
	for i, num := range testCases {
		result1 := hammingWeight(num)
		result2 := hammingWeightSimple(num)
		result3 := hammingWeightOptimized(num)
		result4 := hammingWeightBuiltIn(num)
		
		fmt.Printf("Test Case %d: %032b (%d)\n", i+1, num, num)
		fmt.Printf("  Brian Kernighan: %d\n", result1)
		fmt.Printf("  Simple loop: %d\n", result2)
		fmt.Printf("  Bit tricks: %d\n", result3)
		fmt.Printf("  Built-in style: %d\n\n", result4)
	}
}
