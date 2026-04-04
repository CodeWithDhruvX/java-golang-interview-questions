package main

import (
	"fmt"
	"math"
)

// 233. Number of Digit One - Digit DP
// Time: O(N * D * 9), Space: O(N * D) where N is number length, D is digit count
func countDigitOne(n int) int {
	if n <= 0 {
		return 0
	}
	
	// Convert to string to get length
	s := fmt.Sprintf("%d", n)
	length := len(s)
	
	// Count digit ones for numbers with less digits
	count := 0
	for i := 1; i < length; i++ {
		count += countDigitOneHelper(i-1)
	}
	
	// Count for numbers with same number of digits
	firstDigit := int(s[0] - '0')
	remaining := n - firstDigit*int(math.Pow(10, float64(length-1)))
	
	if firstDigit == 1 {
		count += remaining + 1
	} else {
		count += int(math.Pow(10, float64(length-1))) + countDigitOneHelper(firstDigit-1)*int(math.Pow(10, float64(length-1)))
	}
	
	return count
}

func countDigitOneHelper(n int) int {
	if n <= 0 {
		return 0
	}
	
	count := 0
	for i := 1; i <= n; i++ {
		count += countOnesInNumber(i)
	}
	return count
}

func countOnesInNumber(num int) int {
	count := 0
	for num > 0 {
		if num%10 == 1 {
			count++
		}
		num /= 10
	}
	return count
}

// Digit DP with memoization
func countDigitOneDP(n int) int {
	if n <= 0 {
		return 0
	}
	
	s := fmt.Sprintf("%d", n)
	length := len(s)
	
	// DP table: dp[pos][tight][hasOne]
	dp := make([][][]int, length+1)
	for i := range dp {
		dp[i] = make([]int, 2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	
	// Initialize
	for i := 0; i < length; i++ {
		for tight := 0; tight < 2; tight++ {
			for hasOne := 0; hasOne < 2; hasOne++ {
				dp[i][tight][hasOne] = -1
			}
		}
	}
	
	// Count from 1 to n-1
	for i := 1; i < n; i++ {
		count += countOnesInNumberDP(fmt.Sprintf("%d", i), dp, length)
	}
	
	// Count for n
	count += countOnesInNumberDP(s, dp, length)
	
	return count
}

func countOnesInNumberDP(s string, dp [][][]int, length int) int {
	if len(s) == 0 {
		return 0
	}
	
	return countOnesInNumberDPHelper(s, 0, 0, false, dp, length)
}

func countOnesInNumberDPHelper(s string, pos int, tight int, hasOne int, dp [][][]int, length int) int {
	if pos == length {
		return hasOne
	}
	
	if dp[pos][tight][hasOne] != -1 {
		return dp[pos][tight][hasOne]
	}
	
	limit := 9
	if tight == 1 {
		limit = int(s[pos] - '0')
	}
	
	total := 0
	for digit := 1; digit <= limit; digit++ {
		newTight := 0
		if digit == limit && tight == 1 {
			newTight = 1
		}
		
		newHasOne := hasOne
		if digit == 1 {
			newHasOne = 1
		}
		
		total += countOnesInNumberDPHelper(s, pos+1, newTight, newHasOne, dp, length)
	}
	
	dp[pos][tight][hasOne] = total
	return total
}

// Digit DP with different base
func countDigitOneBase(n, base int) int {
	if n <= 0 {
		return 0
	}
	
	s := fmt.Sprintf("%d", n)
	length := len(s)
	
	count := 0
	for i := 1; i < length; i++ {
		count += countDigitOneHelperBase(i-1, base)
	}
	
	firstDigit := int(s[0] - '0')
	remaining := n - firstDigit*int(math.Pow(float64(base), float64(length-1)))
	
	if firstDigit == 1 {
		count += remaining + 1
	} else {
		count += int(math.Pow(float64(base), float64(length-1))) + countDigitOneHelperBase(firstDigit-1, base)*int(math.Pow(float64(base), float64(length-1)))
	}
	
	return count
}

func countDigitOneHelperBase(n, base int) int {
	count := 0
	for n > 0 {
		if n%base == 1 {
			count++
		}
		n /= base
	}
	return count
}

// Digit DP for range queries
func countDigitOneInRange(low, high int) int {
	return countDigitOne(high) - countDigitOne(low-1)
}

// Digit DP with preprocessing
func countDigitOnePreprocessed(n int) int {
	if n <= 0 {
		return 0
	}
	
	// Precompute counts for all numbers up to n
	onesCount := make([]int, n+1)
	onesCount[0] = 0
	
	for i := 1; i <= n; i++ {
		onesCount[i] = onesCount[i/10] + countOnesInNumber(i%10)
	}
	
	// Calculate prefix sums
	prefixOnes := make([]int, n+1)
	prefixOnes[0] = 0
	
	for i := 1; i <= n; i++ {
		prefixOnes[i] = prefixOnes[i-1] + onesCount[i]
	}
	
	return prefixOnes[n]
}

// Digit DP with mathematical formula
func countDigitOneMathematical(n int) int {
	if n <= 0 {
		return 0
	}
	
	count := 0
	length := 1
	temp := n
	
	for temp >= 10 {
		length++
		temp /= 10
	}
	
	// Count for numbers with less digits
	for i := 1; i < length; i++ {
		count += int(math.Pow(10, float64(i-1))) * int(math.Pow(8, float64(i-1)))
	}
	
	// Count for numbers with same number of digits
	s := fmt.Sprintf("%d", n)
	firstDigit := int(s[0] - '0')
	remaining := n - firstDigit*int(math.Pow(10, float64(length-1)))
	
	if firstDigit == 1 {
		count += remaining + 1
	} else {
		count += int(math.Pow(10, float64(length-1))) + 
			firstDigit*int(math.Pow(8, float64(length-1))) +
			int(math.Pow(10, float64(length-1))) * (length-1)
	}
	
	return count
}

// Digit DP for multiple ranges
func countDigitOneMultipleRanges(ranges [][2]int) []int {
	results := make([]int, len(ranges))
	
	for i, r := range ranges {
		results[i] = countDigitOneInRange(r[0], r[1])
	}
	
	return results
}

// Digit DP with constraint
func countDigitOneWithConstraint(n, maxOnes int) int {
	if n <= 0 {
		return 0
	}
	
	count := 0
	for i := 1; i <= n; i++ {
		ones := countOnesInNumber(i)
		if ones <= maxOnes {
			count++
		}
	}
	
	return count
}

// Digit DP with binary representation
func countDigitOneBinary(n int) int {
	if n <= 0 {
		return 0
	}
	
	count := 0
	for i := 1; i <= n; i++ {
		binary := fmt.Sprintf("%b", i)
		for _, bit := range binary {
			if bit == '1' {
				count++
				break
			}
		}
	}
	
	return count
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Digit DP with Position-Based Counting
- **Digit Position Analysis**: Count ones at each digit position
- **Mathematical Formula**: Use place value patterns for efficient counting
- **Tight Constraint**: Track if current prefix matches target number
- **State Compression**: Only track essential DP states

## 2. PROBLEM CHARACTERISTICS
- **Digit Counting**: Count occurrences of digit '1' in range [1, n]
- **Positional Analysis**: Each digit position contributes independently
- **Range Queries**: Support for counting in arbitrary ranges
- **Base Independence**: Pattern extends to different number bases

## 3. SIMILAR PROBLEMS
- Numbers With Same Consecutive Differences (LeetCode 967) - Digit DP generation
- Count Numbers with Unique Digits (LeetCode 357) - Digit DP with constraints
- Sum of Digits in Range (LeetCode 1067) - Digit DP for digit sums
- Numbers at Distance N (LeetCode 967) - Digit DP with distance constraints

## 4. KEY OBSERVATIONS
- **Position Independence**: Each digit position can be analyzed separately
- **Pattern Recognition**: 10^k numbers have predictable '1' counts
- **First Digit Special**: First digit requires special handling
- **Range Decomposition**: Range queries become prefix sum differences

## 5. VARIATIONS & EXTENSIONS
- **Different Bases**: Extend to binary, octal, hexadecimal
- **Multiple Digits**: Count all digits, not just '1'
- **Constraints**: Limit number of '1's in each number
- **Range Queries**: Support multiple range queries efficiently

## 6. INTERVIEW INSIGHTS
- Always clarify: "Number range constraints? Base system? Multiple queries?"
- Edge cases: n=0, n=1, single digit numbers
- Time complexity: O(log N) for mathematical approach
- Space complexity: O(1) for mathematical, O(log N) for DP
- Key insight: analyze each digit position independently

## 7. COMMON MISTAKES
- Off-by-one errors in range calculations
- Incorrect handling of first digit position
- Wrong tight constraint propagation
- Missing edge cases for small numbers
- Inefficient brute force approaches

## 8. OPTIMIZATION STRATEGIES
- **Mathematical Formula**: O(log N) time, O(1) space - optimal
- **Digit DP**: O(D * 10) time, O(D) space - general approach
- **Preprocessing**: O(N) time, O(N) space - for multiple queries
- **Range Optimization**: O(log N) per query with prefix sums

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like counting votes in different voting districts:**
- Each digit position is like a voting district (ones, tens, hundreds)
- In each district, '1' gets a predictable number of votes
- The first district (most significant digit) has special rules
- You count votes district by district and sum them up
- Like analyzing election results by precinct

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Integer n
2. **Goal**: Count total occurrences of digit '1' in numbers 1 to n
3. **Rules**: Count all '1's in all positions (not just leading)
4. **Output**: Total count of digit '1'

#### Phase 2: Key Insight Recognition
- **"Position analysis natural"** → Each digit position contributes independently
- **"Pattern recognition"** → 10^k numbers have predictable '1' counts
- **"First digit special"** → Most significant digit needs special handling
- **"Range decomposition"** → Range queries become prefix differences

#### Phase 3: Strategy Development
```
Human thought process:
"I need to count '1's in 1..n.
Brute force would be O(N * log N).

Mathematical Approach:
1. For each digit position (ones, tens, hundreds...):
   - Count complete cycles of 0-9
   - Handle partial cycle at the end
   - Special case for most significant digit
2. Sum contributions from all positions
3. Handle edge cases for small numbers

This gives O(log N) time!"
```

#### Phase 4: Edge Case Handling
- **n = 0**: Return 0 (no positive numbers)
- **n = 1**: Return 1 (only number 1)
- **Single digit**: Simple counting
- **Large n**: Ensure mathematical precision

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Example: n = 13

Human thinking:
"Mathematical Approach:
Positions: ones (10^0), tens (10^1)

Ones position:
- Complete cycles: 13/10 = 1 complete cycle (0-9)
- Each complete cycle has 1 '1' at ones position → 1
- Partial cycle: 13%10 = 3 (numbers 10,11,12,13)
- '1's in partial: 1 (number 11)
- Total ones at ones position: 1 + 1 = 2

Tens position:
- Complete cycles: 13/100 = 0 complete cycles
- First digit: 13/10 = 1
- Since first digit = 1: remaining + 1 = (13-10) + 1 = 4
- '1's at tens position: 4 (numbers 10,11,12,13)

Total '1's: 2 (ones) + 4 (tens) = 6

Verification: 1,10,11,12,13 → '1's: 1+1+2+1+1 = 6 ✓"
```

#### Phase 6: Intuition Validation
- **Why position analysis works**: Each position cycles independently
- **Why mathematical formula works**: Patterns repeat every power of 10
- **Why first digit special**: It doesn't complete full cycles
- **Why O(log N)**: Only process each digit position once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not brute force?"** → O(N * log N) vs O(log N) time complexity
2. **"Should I use DP?"** → Mathematical approach simpler and faster
3. **"What about other digits?"** → Same pattern works for any digit
4. **"Can I handle ranges?"** → Use prefix sum technique
5. **"What about different bases?"** → Pattern extends with base changes

### Real-World Analogy
**Like analyzing sales data by product categories:**
- Each digit position is like a product category
- You analyze sales patterns for each category separately
- Some categories have complete sales cycles, others have partial
- The main category (most significant digit) has special reporting
- You combine all category analyses for total sales report
- Like business intelligence analysis by dimensions

### Human-Readable Pseudocode
```
function countDigitOne(n):
    if n <= 0: return 0
    
    count = 0
    position = 1  # 10^0 for ones place
    
    while position <= n:
        # Complete cycles
        complete = n // (position * 10)
        count += complete * position
        
        # Partial cycle
        remainder = n % (position * 10)
        if remainder >= position:
            count += min(remainder - position + 1, position)
        
        position *= 10
    
    return count
```

### Execution Visualization

### Example: n = 13
```
Position Analysis:

Position 1 (ones place, position = 1):
Complete cycles: 13 // 10 = 1
Contribution: 1 * 1 = 1 (from cycle 0-9)
Partial: 13 % 10 = 3 >= 1
Partial contribution: min(3-1+1, 1) = 1 (from number 11)
Total at ones: 1 + 1 = 2

Position 2 (tens place, position = 10):
Complete cycles: 13 // 100 = 0
Contribution: 0 * 10 = 0
Partial: 13 % 100 = 13 >= 10
Partial contribution: min(13-10+1, 10) = 4 (from 10,11,12,13)
Total at tens: 0 + 4 = 4

Position 3 (hundreds place, position = 100):
position > n, stop

Final Result: 2 + 4 = 6 ✓
```

### Key Visualization Points:
- **Position Independence**: Each digit position analyzed separately
- **Cycle Pattern**: 0-9 cycles repeat predictably
- **Partial Handling**: Careful handling of incomplete final cycle
- **Summation**: Total is sum of all position contributions

### Mathematical Pattern Visualization:
```
For position = 10^k:
- Complete cycles: floor(n / 10^(k+1)) * 10^k
- Partial cycle: min(max(n % 10^(k+1) - 10^k + 1, 0), 10^k)
- Total at position: complete + partial

Example: n = 345
Position 1 (ones): floor(345/10)*1 + min(345%10-1+1,1) = 34*1 + min(5,1) = 39
Position 2 (tens): floor(345/100)*10 + min(345%100-10+1,10) = 3*10 + min(45,10) = 40
Position 3 (hundreds): floor(345/1000)*100 + min(345%1000-100+1,100) = 0*100 + min(245,100) = 100
Total: 39 + 40 + 100 = 179
```

### Time Complexity Breakdown:
- **Mathematical**: O(log N) time, O(1) space - optimal
- **Digit DP**: O(D * 10) time, O(D) space - general approach
- **Preprocessing**: O(N) time, O(N) space - for multiple queries
- **Range Queries**: O(log N) per query with prefix sums

### Alternative Approaches:

#### 1. Brute Force (O(N * log N) time, O(1) space)
```go
func countDigitOneBruteForce(n int) int {
    count := 0
    for i := 1; i <= n; i++ {
        for num := i; num > 0; num /= 10 {
            if num%10 == 1 {
                count++
            }
        }
    }
    return count
}
```
- **Pros**: Simple to understand
- **Cons**: Too slow for large n

#### 2. Digit DP (O(D * 10) time, O(D) space)
```go
func countDigitOneDP(n int) int {
    // DP with states: position, tight, hasOne
    // More general but complex
    // ... implementation details omitted
}
```
- **Pros**: Handles arbitrary constraints
- **Cons**: Overkill for simple counting

#### 3. Preprocessing (O(N) time, O(N) space)
```go
func countDigitOnePreprocessed(n int) int {
    // Precompute counts for all numbers up to n
    // Use prefix sums for range queries
    // ... implementation details omitted
}
```
- **Pros**: Fast for multiple queries
- **Cons**: High memory usage

### Extensions for Interviews:
- **Different Digits**: Count any digit (0-9) with same pattern
- **Multiple Digits**: Count all digits simultaneously
- **Range Queries**: Efficient multiple range queries
- **Base Conversion**: Extend to binary, octal, hexadecimal
- **Constraints**: Count numbers with limited '1's
*/
func main() {
	// Test cases
	fmt.Println("=== Testing Number of Digit One - Digit DP ===")
	
	testCases := []struct {
		n          int
		description string
	}{
		{13, "Standard case"},
		{0, "Zero"},
		{1, "Single digit"},
		{9, "Single digit max"},
		{10, "Two digits"},
		{99, "Two digits max"},
		{100, "Three digits"},
		{999, "Three digits max"},
		{1000, "Four digits"},
		{1111, "All ones"},
		{1234, "Mixed digits"},
		{10000, "Five digits"},
		{100000, "Six digits"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s (n=%d)\n", i+1, tc.description, tc.n)
		
		result1 := countDigitOne(tc.n)
		result2 := countDigitOneDP(tc.n)
		result3 := countDigitOneMathematical(tc.n)
		result4 := countDigitOnePreprocessed(tc.n)
		
		fmt.Printf("  Standard: %d\n", result1)
		fmt.Printf("  DP: %d\n", result2)
		fmt.Printf("  Mathematical: %d\n", result3)
		fmt.Printf("  Preprocessed: %d\n\n", result4)
	}
	
	// Test range queries
	fmt.Println("=== Range Queries Test ===")
	ranges := [][2]int{
		{1, 13},
		{10, 100},
		{100, 1000},
		{1000, 10000},
	}
	
	for i, r := range ranges {
		result := countDigitOneInRange(r[0], r[1])
		fmt.Printf("Range [%d, %d]: %d\n", r[0], r[1], result)
	}
	
	// Test different bases
	fmt.Println("\n=== Different Base Test ===")
	bases := []int{2, 8, 16}
	testN := 100
	
	for _, base := range bases {
		result := countDigitOneBase(testN, base)
		fmt.Printf("Base %d (n=%d): %d\n", base, testN, result)
	}
	
	// Test constraints
	fmt.Println("\n=== Constraint Test ===")
	constraints := []int{1, 2, 3, 5}
	
	for _, maxOnes := range constraints {
		result := countDigitOneWithConstraint(100, maxOnes)
		fmt.Printf("Max ones %d: %d\n", maxOnes, result)
	}
	
	// Test binary representation
	fmt.Println("\n=== Binary Representation Test ===")
	binaryTest := []int{5, 10, 15, 31, 63}
	
	for _, n := range binaryTest {
		result := countDigitOneBinary(n)
		fmt.Printf("n=%d (binary=%b): %d\n", n, n, result)
	}
	
	// Performance test
	fmt.Println("\n=== Performance Test ===")
	
	largeN := 1000000
	fmt.Printf("Large test with n=%d\n", largeN)
	
	result := countDigitOneDP(largeN)
	fmt.Printf("DP result: %d\n", result)
	
	result = countDigitOneMathematical(largeN)
	fmt.Printf("Mathematical result: %d\n", result)
	
	// Edge cases
	fmt.Println("\n=== Edge Cases ===")
	
	// Negative number
	fmt.Printf("Negative number: %d\n", countDigitOne(-5))
	
	// Very large number
	veryLarge := 1000000000
	fmt.Printf("Very large number: %d\n", countDigitOneMathematical(veryLarge))
	
	// Test with multiple ranges
	fmt.Println("\n=== Multiple Ranges Test ===")
	multipleRanges := [][2]int{
		{1, 10},
		{11, 20},
		{21, 30},
	}
	
	results := countDigitOneMultipleRanges(multipleRanges)
	for i, result := range results {
		fmt.Printf("Range %d: %d\n", i, result)
	}
}
