package main

import "fmt"

// 338. Counting Bits
// Time: O(N), Space: O(N)
func countBits(n int) []int {
	result := make([]int, n+1)
	
	for i := 1; i <= n; i++ {
		// Brian Kernighan's algorithm
		count := 0
		num := i
		for num > 0 {
			num &= num - 1 // Clear the least significant bit set
			count++
		}
		result[i] = count
	}
	
	return result
}

// Optimized DP approach: O(N) time, O(N) space
func countBitsDP(n int) []int {
	result := make([]int, n+1)
	
	for i := 1; i <= n; i++ {
		// i >> 1 removes the least significant bit
		// i & 1 gives the least significant bit (0 or 1)
		result[i] = result[i>>1] + (i & 1)
	}
	
	return result
}

// Most optimized DP approach
func countBitsDPOptimized(n int) []int {
	result := make([]int, n+1)
	
	for i := 1; i <= n; i++ {
		// i & (i-1) clears the least significant set bit
		result[i] = result[i&(i-1)] + 1
	}
	
	return result
}

// Function to count bits in a single number
func countBitsInNumber(num int) int {
	count := 0
	for num > 0 {
		count += num & 1
		num >>= 1
	}
	return count
}

// Built-in function approach (Go specific)
func countBitsBuiltIn(n int) []int {
	result := make([]int, n+1)
	
	for i := 0; i <= n; i++ {
		// Go doesn't have built-in popcount like C++, but we can use bit manipulation
		count := 0
		num := i
		for num > 0 {
			num &= num - 1
			count++
		}
		result[i] = count
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Dynamic Programming for Bit Counting
- **DP State**: dp[i] = number of 1s in binary representation of i
- **State Transition**: dp[i] = dp[i>>1] + (i & 1)
- **Bit Analysis**: i & 1 extracts least significant bit, i >> 1 removes it
- **Precomputation**: Build table for all numbers from 0 to n

## 2. PROBLEM CHARACTERISTICS
- **Range Query**: Count bits for all numbers from 0 to n
- **Dynamic Programming**: Use previously computed results
- **Efficient Computation**: O(n) time for all queries
- **Memory Tradeoff**: O(n) space for O(1) query time

## 3. SIMILAR PROBLEMS
- Number of 1 Bits (LeetCode 191) - Single number optimization
- Reverse Bits (LeetCode 190) - Bit manipulation
- Single Number (LeetCode 136) - XOR properties
- Bit Manipulation problems using DP

## 4. KEY OBSERVATIONS
- **DP Recurrence**: dp[i] = dp[i>>1] + (i & 1)
- **Bit Isolation**: i & 1 gives current LSB, i >> 1 shifts right
- **Base Case**: dp[0] = 0 (no bits set in 0)
- **Range Queries**: Precompute enables O(1) per query

## 5. VARIATIONS & EXTENSIONS
- **Multiple Queries**: Answer many bit count queries efficiently
- **Different Ranges**: Count bits for arbitrary ranges
- **Bit Position Analysis**: Count specific bit positions
- **Streaming Updates**: Handle dynamic bit updates

## 6. INTERVIEW INSIGHTS
- Always clarify: "Single query or multiple? Range of n?"
- Edge cases: n = 0, negative numbers, large values
- Time complexity: O(n) preprocessing, O(1) per query
- Space complexity: O(n) for DP table
- Key insight: dp[i] = dp[i>>1] + (i & 1) recurrence

## 7. COMMON MISTAKES
- Wrong DP recurrence (off-by-one errors)
- Not handling base case dp[0] = 0 correctly
- Integer overflow for large n values
- Not understanding bit operations (>> and &)
- Confusing dp[i] with dp[i-1] relationships

## 8. OPTIMIZATION STRATEGIES
- **DP Approach**: O(n) preprocessing, O(1) per query
- **Space Optimization**: Use dp[i>>1] instead of separate arrays
- **Bit Tricks**: Use bit manipulation for efficiency
- **Early Termination**: Not applicable for range queries

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like building a bit count pyramid:**
- You need to count bits for all numbers from 0 to n
- Each number's bit count depends on the number without its last bit
- Build the pyramid bottom-up: dp[0], dp[1], dp[2], ..., dp[n]
- Each level uses the result from the previous level

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Integer n
2. **Goal**: Return array where result[i] = count of 1s in i
3. **Constraint**: Need result for all numbers from 0 to n
4. **Output**: Array of bit counts

#### Phase 2: Key Insight Recognition
- **"DP natural fit"** → dp[i] depends on dp[i>>1]
- **"Bit analysis"** → i & 1 gives current bit, i >> 1 removes it
- **"Recurrence relation"** → dp[i] = dp[i>>1] + (i & 1)
- **"Base case"** → dp[0] = 0 (no bits in 0)

#### Phase 3: Strategy Development
```
Human thought process:
"I need bit counts for all numbers from 0 to n.
I can use DP where dp[i] = count of 1s in i.
The key insight: dp[i] = dp[i>>1] + (i & 1)
Why? i>>1 removes the last bit, i & 1 checks if last bit was 1.
So dp[i] = bits in i without last bit + (1 if last bit was set).
Build this bottom-up from 0 to n!"
```

#### Phase 4: Edge Case Handling
- **n = 0**: Return [0]
- **Negative numbers**: Need to handle two's complement
- **Large n**: Ensure array size is manageable
- **Single query**: Return dp[n] directly

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
n = 5

Human thinking:
"I'll build dp array from 0 to 5:
dp[0] = 0 (no bits in 0)

dp[1]: dp[1>>1] + (1 & 1) = dp[0] + 1 = 0 + 1 = 1
dp[2]: dp[2>>1] + (2 & 1) = dp[1] + 0 = 1 + 0 = 1
dp[3]: dp[3>>1] + (3 & 1) = dp[1] + 1 = 1 + 1 = 2
dp[4]: dp[4>>1] + (4 & 1) = dp[2] + 0 = 1 + 0 = 1
dp[5]: dp[5>>1] + (5 & 1) = dp[2] + 1 = 1 + 1 = 2

Final result: [0,1,1,2,1,2] ✓"

Let me verify:
0: 0 (0 bits) ✓
1: 1 (1 bit) ✓  
2: 10 (2 bits) ✓
3: 11 (3 bits) ✓
4: 100 (3 bits) ✓
5: 101 (3 bits) ✓
```

#### Phase 6: Intuition Validation
- **Why DP works**: Each number's bit count builds on previous results
- **Why recurrence works**: Separates last bit from remaining bits
- **Why O(n)**: Each dp[i] computed in O(1) time
- **Why O(n) space**: Need to store all intermediate results

### Common Human Pitfalls & How to Avoid Them
1. **"Why not compute each individually?"** → O(n log n) vs O(n) total
2. **"Should I use recursion?"** → Works but may have stack overflow
3. **"What about multiple queries?"** → DP enables O(1) per query
4. **"Can I optimize space?"** → Need O(n) space for recurrence
5. **"What about bit tricks?"** → DP already uses bit manipulation optimally

### Real-World Analogy
**Like building a population count for each generation:**
- You have generations numbered 0 to n
- Each generation's population count depends on previous generation
- The last bit represents "new births" in current generation
- Previous generation's count + current births = current population
- Build the complete demographic table

### Human-Readable Pseudocode
```
function countBits(n):
    if n < 0:
        return []
    
    dp = array of size n+1
    dp[0] = 0
    
    for i from 1 to n:
        dp[i] = dp[i>>1] + (i & 1)
    
    return dp
```

### Execution Visualization

### Example: n = 5
```
DP Table Construction:
dp[0] = 0

dp[1]: dp[0] + (1 & 1) = 0 + 1 = 1
dp[2]: dp[1] + (2 & 1) = 1 + 0 = 1  
dp[3]: dp[1] + (3 & 1) = 1 + 1 = 2
dp[4]: dp[2] + (4 & 1) = 1 + 0 = 1
dp[5]: dp[2] + (5 & 1) = 1 + 1 = 2

Final result: [0,1,1,2,1,2] ✓
```

### Key Visualization Points:
- **DP Recurrence**: dp[i] = dp[i>>1] + (i & 1)
- **Bit Analysis**: i & 1 extracts LSB, i >> 1 removes it
- **Base Case**: dp[0] = 0
- **Range Coverage**: Results for all numbers 0 to n

### Memory Layout Visualization:
```
DP State Evolution:
n = 5

Step 0: dp[0] = 0
Step 1: dp[1] = dp[0] + (1 & 1) = 0 + 1 = 1
Step 2: dp[2] = dp[1] + (2 & 1) = 1 + 0 = 1
Step 3: dp[3] = dp[1] + (3 & 1) = 1 + 1 = 2
Step 4: dp[4] = dp[2] + (4 & 1) = 1 + 0 = 1
Step 5: dp[5] = dp[2] + (5 & 1) = 1 + 1 = 2

Binary representations:
0: 0000 (0 bits)
1: 0001 (1 bit)
2: 0010 (1 bit)
3: 0011 (2 bits)
4: 0100 (1 bit)
5: 0101 (2 bits)

DP matches bit counts ✓
```

### Time Complexity Breakdown:
- **DP Construction**: O(n) time, O(n) space
- **Single Query**: O(1) time after preprocessing
- **Bit Operations**: Each dp[i] uses O(1) bit operations
- **Total**: O(n) preprocessing, O(1) per query

### Alternative Approaches:

#### 1. Individual Computation (O(n log n) time, O(1) space)
```go
func countBitsIndividual(n int) []int {
    result := make([]int, n+1)
    for i := 0; i <= n; i++ {
        count := 0
        num := i
        for num > 0 {
            count += num & 1
            num >>= 1
        }
        result[i] = count
    }
    return result
}
```
- **Pros**: Simple to understand
- **Cons**: O(n log n) time, recomputes for each query

#### 2. Optimized DP (O(n) time, O(n) space)
```go
func countBitsDPOptimized(n int) []int {
    result := make([]int, n+1)
    result[0] = 0
    
    for i := 1; i <= n; i++ {
        // dp[i] = dp[i>>1] + (i & 1) = result[i>>1] + (i & 1)
        result[i] = result[i>>1] + (i & 1)
    }
    
    return result
}
```
- **Pros**: Same complexity, more optimized implementation
- **Cons**: Still O(n) space

#### 3. Most Optimized DP (O(n) time, O(n) space)
```go
func countBitsDPMostOptimized(n int) []int {
    result := make([]int, n+1)
    result[0] = 0
    
    for i := 1; i <= n; i++ {
        // Use dp[i] = result[i & (i-1)] + 1
        result[i] = result[i&(i-1)] + 1
    }
    
    return result
}
```
- **Pros**: More bit-efficient recurrence
- **Cons**: Slightly more complex recurrence

### Extensions for Interviews:
- **Multiple Queries**: Answer many bit count queries efficiently
- **Different Ranges**: Count bits for arbitrary ranges
- **Bit Position Analysis**: Count specific bit positions
- **Streaming Updates**: Handle dynamic bit updates
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []int{
		2, 5, 10, 15, 0, 1, 7, 8, 16, 31,
	}
	
	for i, n := range testCases {
		result1 := countBits(n)
		result2 := countBitsDP(n)
		result3 := countBitsDPOptimized(n)
		result4 := countBitsBuiltIn(n)
		
		fmt.Printf("Test Case %d: n=%d\n", i+1, n)
		fmt.Printf("  Brian Kernighan: %v\n", result1)
		fmt.Printf("  DP (right shift): %v\n", result2)
		fmt.Printf("  DP (optimized): %v\n", result3)
		fmt.Printf("  Built-in style: %v\n\n", result4)
	}
}
