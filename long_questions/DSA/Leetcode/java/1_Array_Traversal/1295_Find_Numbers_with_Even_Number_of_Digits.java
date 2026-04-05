import java.util.Arrays;

public class FindNumbersWithEvenNumberOfDigits {
    
    // 1295. Find Numbers with Even Number of Digits
    // Time: O(N), Space: O(1)
    public static int findNumbers(int[] nums) {
        int count = 0;
        
        for (int num : nums) {
            if (hasEvenDigits(num)) {
                count++;
            }
        }
        
        return count;
    }

    // Helper function to check if a number has even number of digits
    private static boolean hasEvenDigits(int num) {
        if (num == 0) {
            return false;
        }
        
        int digitCount = 0;
        while (num > 0) {
            num /= 10;
            digitCount++;
        }
        
        return digitCount % 2 == 0;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {12, 345, 2, 6, 7896},
            {555, 901, 482, 1771},
            {1, 22, 333, 4444},
            {0, 10, 100, 1000},
            {}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int result = findNumbers(testCases[i]);
            System.out.printf("Test Case %d: %s -> Numbers with even digits: %d\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Digit Counting with Modulo Check
- **Digit Counting**: Count digits by repeated division by 10
- **Even Check**: Use modulo operation to check if count is even
- **Helper Function**: Separate digit counting logic for clarity

## 2. PROBLEM CHARACTERISTICS
- **Integer Array**: Contains positive integers (including zero)
- **Digit Analysis**: Need to count digits for each number
- **Even Check**: Determine if digit count is even
- **Counting Result**: Return total numbers with even digits

## 3. SIMILAR PROBLEMS
- Find Numbers with Odd Number of Digits
- Count Numbers with Specific Digit Properties
- Self Dividing Numbers
- Armstrong Numbers

## 4. KEY OBSERVATIONS
- Zero has 1 digit (odd), so it's not counted
- Positive numbers: digits = floor(log10(n)) + 1
- Can use string conversion as alternative approach
- Division by 10 removes last digit each time

## 5. VARIATIONS & EXTENSIONS
- Use logarithmic formula for digit counting
- Handle negative numbers (consider sign)
- Count numbers with specific digit counts
- Find numbers with digit sum properties

## 6. INTERVIEW INSIGHTS
- Clarify: "Is zero included?" (has 1 digit)
- Clarify: "Are numbers positive?" (affects digit counting)
- Edge cases: empty array, single numbers, large numbers
- Alternative approaches: string conversion vs mathematical

## 7. COMMON MISTAKES
- Mishandling zero (special case with 1 digit)
- Off-by-one errors in digit counting
- Using extra space unnecessarily
- Not considering negative numbers

## 8. OPTIMIZATION STRATEGIES
- Use logarithmic formula: digits = (int)Math.log10(n) + 1
- String conversion: Integer.toString(n).length()
- Precompute digit counts for ranges
- Use mathematical properties for optimization

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like counting letters in words:**
- You have a list of numbers (like words)
- You want to find words with even number of letters
- For each number, you count its digits (like letters)
- You keep a tally of how many have even digit counts

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers
2. **Goal**: Count numbers with even number of digits
3. **Output**: Total count of such numbers

#### Phase 2: Key Insight Recognition
- **"I need to count digits!"** → Core subproblem
- **"How to count digits?"** → Divide by 10 repeatedly
- **"How to check even?"** → Modulo 2 operation
- **"Zero is special!"** → Has 1 digit, not 0

#### Phase 3: Strategy Development
```
Human thought process:
"I'll create a helper function to count digits:
1. Handle zero as special case (1 digit)
2. For positive numbers, divide by 10 until it becomes 0
3. Count how many divisions I performed
4. Check if this count is even using modulo 2

Then I'll iterate through the array and count how many
numbers have even digit counts."
```

#### Phase 4: Edge Case Handling
- **Zero**: Has 1 digit (odd), not counted
- **Empty array**: Return 0
- **Single digit numbers**: 1-9 have 1 digit (odd)
- **Large numbers**: Handle with loop, no overflow issues

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [12, 345, 2, 6, 7896]

Human thinking:
"Let me check each number:
12: 12 → 1 → 0 (2 divisions, even) → count it!
345: 345 → 34 → 3 → 0 (3 divisions, odd) → don't count
2: 2 → 0 (1 division, odd) → don't count
6: 6 → 0 (1 division, odd) → don't count
7896: 7896 → 789 → 78 → 7 → 0 (4 divisions, even) → count it!

Total: 2 numbers have even digits."
```

#### Phase 6: Intuition Validation
- **Why it works**: Division by 10 removes one digit each time
- **Why it's efficient**: Simple arithmetic operations
- **Why it's correct**: Handles all positive integers correctly

### Common Human Pitfalls & How to Avoid Them
1. **"What about zero?"** → Special case with 1 digit
2. **"Should I use strings?"** → Mathematical approach is more efficient
3. **"How about negatives?"** → Clarify input constraints
4. **"Can I optimize more?"** → Logarithmic formula exists

### Real-World Analogy
**Like counting pages in books:**
- You have a shelf of books (numbers)
- You want to find books with even number of pages
- For each book, you count its pages
- You keep track of how many books meet your criteria
- Some books are thin (few pages), others are thick

### Human-Readable Pseudocode
```
function countEvenDigitNumbers(numbers):
    count = 0
    
    for each number in numbers:
        if hasEvenDigits(number):
            count++
    
    return count

function hasEvenDigits(number):
    if number == 0:
        return false  // 0 has 1 digit (odd)
    
    digitCount = 0
    while number > 0:
        number = number / 10
        digitCount++
    
    return digitCount % 2 == 0
```

### Execution Visualization

### Example: [12, 345, 2, 6, 7896]
```
Initial: nums = [12, 345, 2, 6, 7896], count = 0

Number 1: 12
→ 12 > 0: digitCount = 1, num = 1
→ 1 > 0: digitCount = 2, num = 0
→ 2 % 2 == 0: count = 1

Number 2: 345
→ 345 > 0: digitCount = 1, num = 34
→ 34 > 0: digitCount = 2, num = 3
→ 3 > 0: digitCount = 3, num = 0
→ 3 % 2 != 0: count remains 1

Number 3: 2
→ 2 > 0: digitCount = 1, num = 0
→ 1 % 2 != 0: count remains 1

Number 4: 6
→ 6 > 0: digitCount = 1, num = 0
→ 1 % 2 != 0: count remains 1

Number 5: 7896
→ 7896 > 0: digitCount = 1, num = 789
→ 789 > 0: digitCount = 2, num = 78
→ 78 > 0: digitCount = 3, num = 7
→ 7 > 0: digitCount = 4, num = 0
→ 4 % 2 == 0: count = 2

Final: Return count = 2
```

### Key Visualization Points:
- **Division by 10** removes one digit each iteration
- **Digit count** equals number of divisions until zero
- **Zero is special** - handled as edge case
- **Modulo 2** determines if count is even

### Memory Layout Visualization:
```
Number:   12     345     2      6     7896
Digits:   [1,2]  [3,4,5] [2]    [6]    [7,8,9,6]
Count:    2      3       1      1      4
Even?:    ✓      ✗       ✗      ✗      ✓
Result:   count=1 count=1 count=1 count=1 count=2
```

### Alternative Approaches:
1. **String Conversion**: `Integer.toString(num).length() % 2 == 0`
2. **Logarithmic**: `(int)Math.log10(num) + 1` for positive numbers
3. **Precomputed Ranges**: Count numbers in [10-99], [1000-9999], etc.

### Time Complexity Breakdown:
- **Main Loop**: O(N) - iterate through N numbers
- **Digit Counting**: O(log10(maxNum)) - divisions per number
- **Total**: O(N * log10(maxNum)) - worst case
- **Space**: O(1) - constant extra space
*/
