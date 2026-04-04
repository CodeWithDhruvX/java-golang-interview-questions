package main

import "fmt"

// 1295. Find Numbers with Even Number of Digits
// Time: O(N), Space: O(1)
func findNumbers(nums []int) int {
	count := 0
	
	for _, num := range nums {
		if hasEvenDigits(num) {
			count++
		}
	}
	
	return count
}

// Helper function to check if a number has even number of digits
func hasEvenDigits(num int) bool {
	if num == 0 {
		return false
	}
	
	digitCount := 0
	for num > 0 {
		num /= 10
		digitCount++
	}
	
	return digitCount%2 == 0
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Iterative Digit Counting
- **Main Loop**: Iterate through each number in the array
- **Digit Counter**: For each number, count its digits by repeated division
- **Parity Check**: Check if digit count is even using modulo operation
- **Accumulator**: Count numbers that satisfy the even digit condition

## 2. PROBLEM CHARACTERISTICS
- **Digit Analysis**: Need to determine number of digits for each integer
- **Parity Check**: Specifically looking for even digit counts
- **Iterative Processing**: Process each number independently
- **Mathematical Operations**: Division and modulo for digit counting

## 3. SIMILAR PROBLEMS
- Find Numbers with Even Number of Digits (Current problem)
- Number of Digits in a Number (Classic problem)
- Armstrong Numbers (Digit manipulation)
- Palindrome Numbers (Digit reversal and comparison)

## 4. KEY OBSERVATIONS
- **Zero is special**: Has 1 digit, which is odd
- **Negative numbers**: Problem typically assumes positive integers
- **Digit counting**: Repeated division by 10 until number becomes 0
- **Even check**: Simple modulo 2 operation on digit count

## 5. VARIATIONS & EXTENSIONS
- Count numbers with odd number of digits
- Find numbers with specific digit count (e.g., exactly 3 digits)
- Handle negative numbers (ignore sign for digit counting)
- Use logarithmic approach for faster digit counting

## 6. INTERVIEW INSIGHTS
- Always clarify: "Are numbers positive? What about zero?"
- Edge cases: empty array, single element, zero, large numbers
- Space complexity: O(1) - constant extra space
- Time complexity: O(N * D) where D is average digit count

## 7. COMMON MISTAKES
- Forgetting to handle zero as a special case
- Not considering negative numbers (if allowed)
- Using string conversion unnecessarily
- Off-by-one errors in digit counting logic
- Not handling empty array case

## 8. OPTIMIZATION STRATEGIES
- **Logarithmic approach**: Use log10 to count digits in O(1) per number
- **String conversion**: Convert to string and check length (simpler but slower)
- **Precomputed ranges**: Use mathematical ranges for efficiency
- **Bit manipulation**: Not applicable for decimal digit counting

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like counting letters in words:**
- You have a list of numbers (like words)
- For each number, you count how many digits it has (like letters)
- You're only interested in numbers with even digit counts
- You keep a running total of how many numbers meet this criteria
- Zero is a special case with only 1 digit

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers
2. **Goal**: Count numbers with even number of digits
3. **Output**: Total count of such numbers

#### Phase 2: Key Insight Recognition
- **"Digit counting"** → Need a way to count digits in each number
- **"Even check"** → Simple parity test on digit count
- **"Iterative approach"** → Process each number independently
- **"Special case"** → Zero needs special handling

#### Phase 3: Strategy Development
```
Human thought process:
"For each number, I need to count its digits.
I can do this by repeatedly dividing by 10 until the number becomes 0.
Each division removes one digit, so I count how many divisions I need.
After counting digits, I check if the count is even.
If it is, I increment my total count."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 (no numbers to check)
- **Zero**: Has 1 digit (odd), so doesn't count
- **Single digit numbers**: All have 1 digit (odd), don't count
- **Large numbers**: Division approach works for any size

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [12, 345, 2, 6, 7896]

Human thinking:
"Let me check each number one by one:

Number 12:
   12 ÷ 10 = 1 (count: 1)
   1 ÷ 10 = 0 (count: 2)
   Total digits: 2 (even!) → count = 1

Number 345:
   345 ÷ 10 = 34 (count: 1)
   34 ÷ 10 = 3 (count: 2)
   3 ÷ 10 = 0 (count: 3)
   Total digits: 3 (odd) → count stays 1

Number 2:
   2 ÷ 10 = 0 (count: 1)
   Total digits: 1 (odd) → count stays 1

Number 6:
   6 ÷ 10 = 0 (count: 1)
   Total digits: 1 (odd) → count stays 1

Number 7896:
   7896 ÷ 10 = 789 (count: 1)
   789 ÷ 10 = 78 (count: 2)
   78 ÷ 10 = 7 (count: 3)
   7 ÷ 10 = 0 (count: 4)
   Total digits: 4 (even!) → count = 2

Final answer: 2 numbers have even digits"
```

#### Phase 6: Intuition Validation
- **Why it works**: Division by 10 reliably removes one digit each time
- **Why it's correct**: We count digits accurately and check parity correctly
- **Why it's efficient**: Simple arithmetic operations, no complex data structures

### Common Human Pitfalls & How to Avoid Them
1. **"What about zero?"** → Zero has 1 digit, handle as special case
2. **"Should I use strings?"** → Possible but slower than arithmetic
3. **"What about negative numbers?"** → Usually not allowed, or ignore sign
4. **"Can I use logarithms?"** → Yes, but division is more intuitive

### Real-World Analogy
**Like counting pages in books:**
- You have a shelf of books (array of numbers)
- For each book, you count how many pages it has (digits)
- You're only interested in books with even number of pages
- You keep a tally of how many books meet this criteria
- Some books might be very thick (many digits), some very thin

### Human-Readable Pseudocode
```
function findNumbersWithEvenDigits(numberArray):
    evenDigitCount = 0
    
    for each number in numberArray:
        if number is 0:
            digitCount = 1
        else:
            digitCount = 0
            temp = number
            while temp > 0:
                temp = temp / 10
                digitCount = digitCount + 1
        
        if digitCount is even:
            evenDigitCount = evenDigitCount + 1
    
    return evenDigitCount
```

### Execution Visualization

### Example: [12, 345, 2, 6, 7896]
```
Initial: count = 0

=== Processing 12 ===
temp = 12, digitCount = 0
Step 1: temp = 12/10 = 1, digitCount = 1
Step 2: temp = 1/10 = 0, digitCount = 2
digitCount (2) is even → count = 1

=== Processing 345 ===
temp = 345, digitCount = 0
Step 1: temp = 345/10 = 34, digitCount = 1
Step 2: temp = 34/10 = 3, digitCount = 2
Step 3: temp = 3/10 = 0, digitCount = 3
digitCount (3) is odd → count stays 1

=== Processing 2 ===
temp = 2, digitCount = 0
Step 1: temp = 2/10 = 0, digitCount = 1
digitCount (1) is odd → count stays 1

=== Processing 6 ===
temp = 6, digitCount = 0
Step 1: temp = 6/10 = 0, digitCount = 1
digitCount (1) is odd → count stays 1

=== Processing 7896 ===
temp = 7896, digitCount = 0
Step 1: temp = 7896/10 = 789, digitCount = 1
Step 2: temp = 789/10 = 78, digitCount = 2
Step 3: temp = 78/10 = 7, digitCount = 3
Step 4: temp = 7/10 = 0, digitCount = 4
digitCount (4) is even → count = 2

Final: Return count = 2
```

### Key Visualization Points:
- **Division by 10** removes the last digit each iteration
- **Digit counting** continues until number becomes 0
- **Zero handling** is a special case (1 digit)
- **Even check** uses simple modulo 2 operation

### Memory Layout Visualization:
```
Number:   [12][345][2][6][7896]
Digits:   [ 2][  3][1][1][   4]
Even?:    [✓][ ✗][✗][✗][  ✓]
Count:    [1][  1][1][1][   2] (running total)
```

### Time Complexity Breakdown:
- **Outer Loop**: O(N) - iterate through N numbers
- **Inner Loop**: O(D) - D divisions for each number (D = digit count)
- **Total**: O(N * D) - where D is average digit count (typically small)
- **Space**: O(1) - constant extra space

### Alternative Approaches:

#### 1. Logarithmic Approach (O(1) per number)
```go
func hasEvenDigits(num int) bool {
    if num == 0 {
        return false
    }
    digitCount := int(math.Log10(float64(num))) + 1
    return digitCount%2 == 0
}
```
- **Pros**: O(1) digit counting per number
- **Cons**: Requires math library, floating point operations

#### 2. String Conversion Approach
```go
func hasEvenDigits(num int) bool {
    if num == 0 {
        return false
    }
    digitCount := len(strconv.Itoa(num))
    return digitCount%2 == 0
}
```
- **Pros**: Very simple and readable
- **Cons**: String conversion overhead, extra memory

#### 3. Range-Based Approach (O(1) per number)
```go
func hasEvenDigits(num int) bool {
    if num == 0 {
        return false
    }
    // Precomputed ranges for even digit counts
    return (num >= 10 && num <= 99) || 
           (num >= 1000 && num <= 9999) ||
           (num >= 100000 && num <= 999999)
}
```
- **Pros**: Fastest for typical integer ranges
- **Cons**: Limited to specific number ranges

### Extensions for Interviews:
- **Handle negative numbers**: Use absolute value for digit counting
- **Find numbers with specific digit count**: Modify even check to exact count
- **Count by digit ranges**: Group numbers by their digit lengths
- **Optimize for large arrays**: Consider parallel processing
*/

func main() {
	// Test cases
	testCases := [][]int{
		{12, 345, 2, 6, 7896},
		{555, 901, 482, 1771},
		{1, 22, 333, 4444},
		{0, 10, 100, 1000},
		{},
	}
	
	for i, nums := range testCases {
		result := findNumbers(nums)
		fmt.Printf("Test Case %d: %v -> Numbers with even digits: %d\n", i+1, nums, result)
	}
}
