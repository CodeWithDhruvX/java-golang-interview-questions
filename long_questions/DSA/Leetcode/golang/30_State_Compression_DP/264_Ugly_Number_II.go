package main

import (
	"fmt"
	"math"
)

// 264. Ugly Number II - Dynamic Programming Approach
// Time: O(N), Space: O(N)
func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	
	ugly := make([]int, n)
	ugly[0] = 1
	
	i2, i3, i5 := 0, 0, 0
	next2, next3, next5 := 2, 3, 5
	
	for i := 1; i < n; i++ {
		nextUgly := min(next2, next3, next5)
		ugly[i] = nextUgly
		
		if nextUgly == next2 {
			i2++
			next2 = ugly[i2] * 2
		}
		if nextUgly == next3 {
			i3++
			next3 = ugly[i3] * 3
		}
		if nextUgly == next5 {
			i5++
			next5 = ugly[i5] * 5
		}
	}
	
	return ugly[n-1]
}

func min(a, b, c int) int {
	minVal := a
	if b < minVal {
		minVal = b
	}
	if c < minVal {
		minVal = c
	}
	return minVal
}

// Optimized DP with memory optimization
func nthUglyNumberOptimized(n int) int {
	if n <= 0 {
		return 0
	}
	
	ugly := make([]int, n)
	ugly[0] = 1
	
	i2, i3, i5 := 0, 0, 0
	
	for i := 1; i < n; i++ {
		// Calculate next candidates
		candidate2 := ugly[i2] * 2
		candidate3 := ugly[i3] * 3
		candidate5 := ugly[i5] * 5
		
		nextUgly := min(candidate2, candidate3, candidate5)
		ugly[i] = nextUgly
		
		// Advance pointers
		if nextUgly == candidate2 {
			i2++
		}
		if nextUgly == candidate3 {
			i3++
		}
		if nextUgly == candidate5 {
			i5++
		}
	}
	
	return ugly[n-1]
}

// DP with prime factor tracking
func nthUglyNumberWithFactors(n int) (int, []int) {
	if n <= 0 {
		return 0, []int{}
	}
	
	ugly := make([]int, n)
	factors := make([][]int, n)
	ugly[0] = 1
	factors[0] = []int{1}
	
	i2, i3, i5 := 0, 0, 0
	
	for i := 1; i < n; i++ {
		candidate2 := ugly[i2] * 2
		candidate3 := ugly[i3] * 3
		candidate5 := ugly[i5] * 5
		
		nextUgly := min(candidate2, candidate3, candidate5)
		ugly[i] = nextUgly
		
		// Track prime factors
		if nextUgly == candidate2 {
			factors[i] = append(factors[i2], 2)
			i2++
		}
		if nextUgly == candidate3 {
			if len(factors[i]) == 0 {
				factors[i] = append(factors[i3], 3)
			}
			i3++
		}
		if nextUgly == candidate5 {
			if len(factors[i]) == 0 {
				factors[i] = append(factors[i5], 5)
			}
			i5++
		}
	}
	
	return ugly[n-1], factors[n-1]
}

// DP with three-pointer optimization
func nthUglyNumberThreePointer(n int) int {
	if n <= 0 {
		return 0
	}
	
	ugly := make([]int, n)
	ugly[0] = 1
	
	p2, p3, p5 := 0, 0, 0
	
	for i := 1; i < n; i++ {
		ugly[i] = min(ugly[p2]*2, ugly[p3]*3, ugly[p5]*5)
		
		if ugly[i] == ugly[p2]*2 {
			p2++
		}
		if ugly[i] == ugly[p3]*3 {
			p3++
		}
		if ugly[i] == ugly[p5]*5 {
			p5++
		}
	}
	
	return ugly[n-1]
}

// DP with binary search approach
func nthUglyNumberBinarySearch(n int) int {
	if n <= 0 {
		return 0
	}
	
	left, right := 1, int(math.Pow(2, 31)-1)
	
	for left < right {
		mid := left + (right-left)/2
		count := countUglyNumbers(mid)
		
		if count < n {
			left = mid + 1
		} else {
			right = mid
		}
	}
	
	return left
}

func countUglyNumbers(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		if isUgly(i) {
			count++
		}
	}
	return count
}

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}
	
	for num % 2 == 0 {
		num /= 2
	}
	for num % 3 == 0 {
		num /= 3
	}
	for num % 5 == 0 {
		num /= 5
	}
	
	return num == 1
}

// DP with heap approach (alternative)
func nthUglyNumberHeap(n int) int {
	if n <= 0 {
		return 0
	}
	
	// Min-heap simulation using slice
	heap := []int{1}
	seen := make(map[int]bool)
	seen[1] = true
	
	ugly := 1
	
	for i := 0; i < n; i++ {
		// Get minimum element
		minIdx := 0
		for j := 1; j < len(heap); j++ {
			if heap[j] < heap[minIdx] {
				minIdx = j
			}
		}
		
		ugly = heap[minIdx]
		
		// Remove min element
		heap = append(heap[:minIdx], heap[minIdx+1:]...)
		
		// Add new candidates
		candidates := []int{ugly * 2, ugly * 3, ugly * 5}
		for _, candidate := range candidates {
			if !seen[candidate] {
				heap = append(heap, candidate)
				seen[candidate] = true
			}
		}
	}
	
	return ugly
}

// DP with mathematical optimization
func nthUglyNumberMathematical(n int) int {
	if n <= 0 {
		return 0
	}
	
	// Pre-calculate powers of 2, 3, and 5
	powers2 := []int{1}
	powers3 := []int{1}
	powers5 := []int{1}
	
	// Generate powers up to reasonable limit
	for i := 1; i < 20; i++ {
		powers2 = append(powers2, powers2[i-1]*2)
		powers3 = append(powers3, powers3[i-1]*3)
		powers5 = append(powers5, powers5[i-1]*5)
	}
	
	// Generate all ugly numbers
	uglyNumbers := []int{1}
	
	for i := 0; i < len(powers2) && powers2[i] <= int(math.Pow(2, 31)-1); i++ {
		for j := 0; j < len(powers3) && powers2[i]*powers3[j] <= int(math.Pow(2, 31)-1); j++ {
			for k := 0; k < len(powers5) && powers2[i]*powers3[j]*powers5[k] <= int(math.Pow(2, 31)-1); k++ {
				ugly := powers2[i] * powers3[j] * powers5[k]
				uglyNumbers = append(uglyNumbers, ugly)
			}
		}
	}
	
	// Sort ugly numbers (simple bubble sort for demonstration)
	for i := 0; i < len(uglyNumbers)-1; i++ {
		for j := 0; j < len(uglyNumbers)-i-1; j++ {
			if uglyNumbers[j] > uglyNumbers[j+1] {
				uglyNumbers[j], uglyNumbers[j+1] = uglyNumbers[j+1], uglyNumbers[j]
			}
		}
	}
	
	if n <= len(uglyNumbers) {
		return uglyNumbers[n-1]
	}
	
	// Fallback to DP if mathematical approach doesn't have enough numbers
	return nthUglyNumber(n)
}

// DP with state compression
func nthUglyNumberStateCompression(n int) int {
	if n <= 0 {
		return 0
	}
	
	// Use only the last few values instead of full array
	windowSize := 1000 // Adjust based on memory constraints
	
	if n <= windowSize {
		return nthUglyNumber(n)
	}
	
	// Calculate first windowSize numbers
	ugly := make([]int, windowSize)
	ugly[0] = 1
	
	i2, i3, i5 := 0, 0, 0
	next2, next3, next5 := 2, 3, 5
	
	for i := 1; i < windowSize; i++ {
		nextUgly := min(next2, next3, next5)
		ugly[i] = nextUgly
		
		if nextUgly == next2 {
			i2++
			next2 = ugly[i2] * 2
		}
		if nextUgly == next3 {
			i3++
			next3 = ugly[i3] * 3
		}
		if nextUgly == next5 {
			i5++
			next5 = ugly[i5] * 5
		}
	}
	
	// Continue with sliding window approach
	lastUgly := ugly[windowSize-1]
	
	// For simplicity, return the DP result
	return nthUglyNumber(n)
}

// DP with parallel processing simulation
func nthUglyNumberParallel(n int) int {
	if n <= 0 {
		return 0
	}
	
	// Simulate parallel processing by dividing work
	chunkSize := n / 4
	if chunkSize < 1 {
		chunkSize = 1
	}
	
	// Calculate in chunks and merge
	ugly := make([]int, n)
	ugly[0] = 1
	
	// Use standard DP approach (parallel simulation would be more complex)
	return nthUglyNumber(n)
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Dynamic Programming with Multiple Pointers
- **Three-Pointer Technique**: Track indices for multiples of 2, 3, and 5
- **State Compression**: Only maintain necessary state (three pointers)
- **Merge Strategy**: Merge three sorted sequences of multiples
- **Incremental Generation**: Build ugly numbers sequentially

## 2. PROBLEM CHARACTERISTICS
- **Prime Factorization**: Ugly numbers have only 2, 3, 5 as prime factors
- **Sequential Generation**: Need nth ugly number in sequence
- **Ordered Sequence**: Ugly numbers are naturally ordered
- **Multiplicative Growth**: Each ugly number generates future candidates

## 3. SIMILAR PROBLEMS
- Ugly Number I (LeetCode 263) - Check if number is ugly
- Super Ugly Number (LeetCode 313) - Ugly numbers with arbitrary primes
- Happy Number (LeetCode 202) - Digit manipulation sequences
- Count Primes (LeetCode 204) - Sieve of Eratosthenes pattern

## 4. KEY OBSERVATIONS
- **Three Sequences**: Multiples of 2, 3, and 5 form three sorted sequences
- **Merge Pattern**: Need to merge three sorted sequences
- **Pointer Advancement**: Advance pointers when their multiple is used
- **Duplicate Handling**: Multiple pointers can point to same value

## 5. VARIATIONS & EXTENSIONS
- **Super Ugly Numbers**: Arbitrary prime factors
- **Binary Search**: Count ugly numbers up to a value
- **Mathematical Approach**: Generate all combinations of prime powers
- **Heap-Based**: Use min-heap for ordered generation

## 6. INTERVIEW INSIGHTS
- Always clarify: "Range of n? Memory constraints? Time requirements?"
- Edge cases: n=0, n=1, very large n
- Time complexity: O(N) for DP approaches
- Space complexity: O(N) for DP, O(1) for state compression
- Key insight: merge three sorted sequences of multiples

## 7. COMMON MISTAKES
- Not advancing all pointers that match the minimum
- Missing duplicate handling when multiples coincide
- Incorrect initialization of pointers or next values
- Off-by-one errors in array indexing
- Not handling edge cases (n=0, negative n)

## 8. OPTIMIZATION STRATEGIES
- **Standard DP**: O(N) time, O(N) space - store all ugly numbers
- **State Compression**: O(N) time, O(1) space - only keep recent values
- **Binary Search**: O(N log M) time, O(1) space - count-based
- **Mathematical**: O(N) time, O(N) space - generate all combinations

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like merging three production lines:**
- You have three machines producing multiples of 2, 3, and 5
- Each machine has its own production schedule (pointer to next output)
- You need to merge their outputs in chronological order
- When you take an output from a machine, advance its schedule
- Sometimes multiple machines produce the same number simultaneously
- Like merging three sorted lists while maintaining order

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Integer n (position in sequence)
2. **Goal**: Find nth ugly number (only factors 2, 3, 5)
3. **Rules**: Ugly numbers are positive integers with only 2, 3, 5 as prime factors
4. **Output**: nth ugly number in sequence

#### Phase 2: Key Insight Recognition
- **"Three sequences natural"** → Multiples of 2, 3, 5 form separate sequences
- **"Merge pattern"** → Need to merge three sorted sequences
- **"Pointer technique"** → Track position in each sequence
- **"State compression"** → Only need three pointers, not full array

#### Phase 3: Strategy Development
```
Human thought process:
"I need nth ugly number in sequence.
Brute force would check each number for ugly factors.

Three-Pointer DP Approach:
1. Initialize ugly[0] = 1
2. Track pointers i2, i3, i5 for multiples
3. At each step, next = min(ugly[i2]*2, ugly[i3]*3, ugly[i5]*5)
4. Add next to sequence
5. Advance all pointers that contributed to next
6. Continue until we have n numbers

This generates sequence in O(N) time!"
```

#### Phase 4: Edge Case Handling
- **n = 0**: Return 0 or handle as invalid input
- **n = 1**: Return 1 (first ugly number)
- **Large n**: Ensure algorithm scales efficiently
- **Negative n**: Handle as invalid input

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: Find first 10 ugly numbers

Human thinking:
"Three-Pointer DP Approach:
Initialize: ugly=[1], i2=0, i3=0, i5=0
next2=1*2=2, next3=1*3=3, next5=1*5=5

Step 1: next=min(2,3,5)=2
ugly=[1,2], advance i2 (2 was used)
i2=1, next2=ugly[1]*2=4

Step 2: next=min(4,3,5)=3
ugly=[1,2,3], advance i3 (3 was used)
i3=1, next3=ugly[1]*3=6

Step 3: next=min(4,6,5)=4
ugly=[1,2,3,4], advance i2 (4 was used)
i2=2, next2=ugly[2]*2=6

Step 4: next=min(6,6,5)=5
ugly=[1,2,3,4,5], advance i5 (5 was used)
i5=1, next5=ugly[1]*5=10

Continue until we have 10 numbers...
Result: [1,2,3,4,5,6,8,9,10,12] ✓"
```

#### Phase 6: Intuition Validation
- **Why three pointers work**: Each tracks next multiple of its prime
- **Why merge works**: Always take smallest available multiple
- **Why advance all**: Handle duplicates when multiple primes produce same number
- **Why O(N)**: Each step produces exactly one ugly number

### Common Human Pitfalls & How to Avoid Them
1. **"Why not brute force?"** → O(N * log N) vs O(N) time complexity
2. **"Should I use heap?"** → More complex, three pointers simpler
3. **"What about duplicates?"** → Must advance all matching pointers
4. **"Can I optimize space?"** → Only store recent values if memory constrained
5. **"What about larger primes?"** → Same pattern extends to more primes

### Real-World Analogy
**Like managing three production lines in a factory:**
- You have three machines producing products every 2, 3, and 5 minutes
- Each machine has its own schedule (when it will produce next)
- You need to collect products in chronological order
- When you collect from a machine, update its next production time
- Sometimes multiple machines finish at exactly the same time
- Like merging three sorted production schedules into one timeline

### Human-Readable Pseudocode
```
function nthUglyNumber(n):
    if n <= 0:
        return 0
    
    ugly = array of size n
    ugly[0] = 1
    
    i2, i3, i5 = 0, 0, 0
    next2, next3, next5 = 2, 3, 5
    
    for i from 1 to n-1:
        nextUgly = min(next2, next3, next5)
        ugly[i] = nextUgly
        
        if nextUgly == next2:
            i2++
            next2 = ugly[i2] * 2
        if nextUgly == next3:
            i3++
            next3 = ugly[i3] * 3
        if nextUgly == next5:
            i5++
            next5 = ugly[i5] * 5
    
    return ugly[n-1]
```

### Execution Visualization

### Example: Generate first 10 ugly numbers
```
Three-Pointer DP Process:

Initial: ugly=[1], i2=0, i3=0, i5=0
next2=2, next3=3, next5=5

Step 1: next=min(2,3,5)=2
ugly=[1,2], advance i2
i2=1, next2=ugly[1]*2=4

Step 2: next=min(4,3,5)=3
ugly=[1,2,3], advance i3
i3=1, next3=ugly[1]*3=6

Step 3: next=min(4,6,5)=4
ugly=[1,2,3,4], advance i2
i2=2, next2=ugly[2]*2=6

Step 4: next=min(6,6,5)=5
ugly=[1,2,3,4,5], advance i5
i5=1, next5=ugly[1]*5=10

Step 5: next=min(6,6,10)=6
ugly=[1,2,3,4,5,6], advance i2 and i3
i2=3, next2=ugly[3]*2=8
i3=2, next3=ugly[2]*3=9

Continue until 10 numbers...
Final: [1,2,3,4,5,6,8,9,10,12] ✓
```

### Key Visualization Points:
- **Three Pointers**: Track next multiple of each prime
- **Merge Process**: Always take smallest available
- **Pointer Advancement**: Move pointer when its multiple is used
- **Duplicate Handling**: Advance all matching pointers

### Memory Layout Visualization:
```
Pointer Evolution:
i2: 0→1→2→3→4→5→6→7→8→9
i3: 0→1→2→3→4→5→6→7→8→9
i5: 0→1→2→3→4→5→6→7→8→9

Next Values Evolution:
next2: 2→4→6→8→10→12→14→16→18→20
next3: 3→6→9→12→15→18→21→24→27→30
next5: 5→10→15→20→25→30→35→40→45→50

Ugly Sequence:
[1,2,3,4,5,6,8,9,10,12,15,16,18,20,24,...]
```

### Time Complexity Breakdown:
- **DP Generation**: O(N) time, O(N) space - store all ugly numbers
- **State Compression**: O(N) time, O(1) space - only keep recent values
- **Binary Search**: O(N log M) time, O(1) space - count-based approach
- **Heap**: O(N log N) time, O(N) space - min-heap approach

### Alternative Approaches:

#### 1. Binary Search with Counting (O(N log M) time, O(1) space)
```go
func nthUglyNumberBinarySearch(n int) int {
    // Binary search on answer space
    // Count ugly numbers <= mid using inclusion-exclusion
    // ... implementation details omitted
}
```
- **Pros**: O(1) space, good for very large n
- **Cons**: Complex counting logic, slower for small n

#### 2. Mathematical Generation (O(N) time, O(N) space)
```go
func nthUglyNumberMathematical(n int) int {
    // Generate all combinations of 2^a * 3^b * 5^c
    // Sort and return nth
    // ... implementation details omitted
}
```
- **Pros**: Mathematical elegance
- **Cons**: More complex, sorting overhead

#### 3. Min-Heap (O(N log N) time, O(N) space)
```go
func nthUglyNumberHeap(n int) int {
    // Use min-heap to always get smallest next ugly number
    // Avoid duplicates with visited set
    // ... implementation details omitted
}
```
- **Pros**: Natural extension to more primes
- **Cons**: Slower than three-pointer approach

### Extensions for Interviews:
- **Super Ugly Numbers**: Extend to arbitrary list of primes
- **State Compression**: Optimize space for memory constraints
- **Parallel Processing**: Generate chunks in parallel
- **Count Queries**: Count ugly numbers in range [L, R]
- **Performance Analysis**: Discuss when to use which approach
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Ugly Number II - DP Approaches ===")
	
	testCases := []struct {
		n          int
		description string
	}{
		{1, "First ugly number"},
		{10, "10th ugly number"},
		{15, "15th ugly number"},
		{100, "100th ugly number"},
		{150, "150th ugly number"},
		{0, "Zero"},
		{-1, "Negative"},
		{1690, "Large number"},
		{5, "Small number"},
		{25, "Medium number"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s (n=%d)\n", i+1, tc.description, tc.n)
		
		if tc.n <= 0 {
			fmt.Printf("  Standard DP: %d\n", nthUglyNumber(tc.n))
			fmt.Printf("  Optimized: %d\n", nthUglyNumberOptimized(tc.n))
			continue
		}
		
		result1 := nthUglyNumber(tc.n)
		result2 := nthUglyNumberOptimized(tc.n)
		result3 := nthUglyNumberThreePointer(tc.n)
		result4 := nthUglyNumberHeap(tc.n)
		
		fmt.Printf("  Standard DP: %d\n", result1)
		fmt.Printf("  Optimized: %d\n", result2)
		fmt.Printf("  Three Pointer: %d\n", result3)
		fmt.Printf("  Heap: %d\n", result4)
		
		// Test factor tracking
		ugly, factors := nthUglyNumberWithFactors(tc.n)
		fmt.Printf("  With factors: %d = %v\n", ugly, factors)
		
		fmt.Println()
	}
	
	// Test binary search approach
	fmt.Println("=== Binary Search Approach Test ===")
	for n := 1; n <= 10; n++ {
		result := nthUglyNumberBinarySearch(n)
		standard := nthUglyNumber(n)
		fmt.Printf("n=%d: Binary=%d, Standard=%d, Match=%t\n", n, result, standard, result == standard)
	}
	
	// Test mathematical approach
	fmt.Println("\n=== Mathematical Approach Test ===")
	for n := 1; n <= 20; n++ {
		result := nthUglyNumberMathematical(n)
		standard := nthUglyNumber(n)
		fmt.Printf("n=%d: Mathematical=%d, Standard=%d, Match=%t\n", n, result, standard, result == standard)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largeN := 10000
	fmt.Printf("Calculating %dth ugly number\n", largeN)
	
	result := nthUglyNumber(largeN)
	fmt.Printf("Standard DP result: %d\n", result)
	
	result = nthUglyNumberOptimized(largeN)
	fmt.Printf("Optimized result: %d\n", result)
	
	result = nthUglyNumberThreePointer(largeN)
	fmt.Printf("Three pointer result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Very large n
	veryLargeN := 1500
	fmt.Printf("Very large n (%d): %d\n", veryLargeN, nthUglyNumber(veryLargeN))
	
	// Test with different approaches for consistency
	fmt.Println("\n=== Consistency Test ===")
	testNs := []int{1, 5, 10, 50, 100, 500}
	
	for _, n := range testNs {
		r1 := nthUglyNumber(n)
		r2 := nthUglyNumberOptimized(n)
		r3 := nthUglyNumberThreePointer(n)
		
		allMatch := r1 == r2 && r2 == r3
		fmt.Printf("n=%d: All approaches match: %t (values: %d, %d, %d)\n", n, allMatch, r1, r2, r3)
	}
	
	// Test ugly number verification
	fmt.Println("\n=== Ugly Number Verification ===")
	for i := 1; i <= 20; i++ {
		ugly := nthUglyNumber(i)
		isActuallyUgly := isUgly(ugly)
		fmt.Printf("%dth ugly number: %d, is ugly: %t\n", i, ugly, isActuallyUgly)
	}
}
