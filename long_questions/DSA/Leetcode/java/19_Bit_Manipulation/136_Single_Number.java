import java.util.*;

public class SingleNumber {
    
    // 136. Single Number
    // Time: O(N), Space: O(1)
    public static int singleNumber(int[] nums) {
        int result = 0;
        for (int num : nums) {
            result ^= num;
        }
        return result;
    }

    // Alternative approach using hash map (O(N) space)
    public static int singleNumberHash(int[] nums) {
        Map<Integer, Integer> count = new HashMap<>();
        for (int num : nums) {
            count.put(num, count.getOrDefault(num, 0) + 1);
        }
        
        for (Map.Entry<Integer, Integer> entry : count.entrySet()) {
            if (entry.getValue() == 1) {
                return entry.getKey();
            }
        }
        
        return -1; // Should not reach here for valid input
    }

    // XOR properties explanation
    public static int singleNumberWithExplanation(int[] nums) {
        System.out.printf("XOR Properties:\n");
        System.out.printf("1. a ^ a = 0 (XOR of same numbers is 0)\n");
        System.out.printf("2. a ^ 0 = a (XOR with 0 is the number itself)\n");
        System.out.printf("3. XOR is commutative and associative\n\n");
        
        System.out.printf("Processing: ");
        int result = 0;
        for (int i = 0; i < nums.length; i++) {
            if (i > 0) {
                System.out.printf(" ^ ");
            }
            System.out.printf("%d", nums[i]);
            result ^= nums[i];
        }
        System.out.printf(" = %d\n\n", result);
        
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {2, 2, 1},
            {4, 1, 2, 1, 2},
            {1},
            {2, 2, 3, 3, 4},
            {0, 1, 1},
            {-1, -1, -2},
            {99, 99, 100},
            {5, 5, 6, 6, 7, 7, 8},
            {10, 10, 10, 10, 15}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            
            if (i == 0) {
                // Show detailed explanation for first test case
                int result = singleNumberWithExplanation(testCases[i]);
                System.out.printf("Single number: %d\n\n", result);
            } else {
                int result1 = singleNumber(testCases[i]);
                int result2 = singleNumberHash(testCases[i]);
                System.out.printf("XOR: %d, Hash: %d\n\n", result1, result2);
            }
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Bit Manipulation
- **XOR Properties**: a ^ a = 0, a ^ 0 = a, commutative, associative
- **Single Number**: Find element that appears once, others appear twice
- **Pair Cancellation**: XOR of pairs cancels out to 0
- **Linear Scan**: Single pass through array

## 2. PROBLEM CHARACTERISTICS
- **XOR Operation**: Bitwise exclusive OR
- **Frequency Pattern**: One element appears once, others twice
- **Pair Elimination**: Paired elements cancel via XOR
- **O(1) Space**: No additional data structures needed

## 3. SIMILAR PROBLEMS
- Single Number II (one appears thrice)
- Missing Number
- Find Two Unique Numbers
- Reverse Bits

## 4. KEY OBSERVATIONS
- XOR of same numbers is 0: a ^ a = 0
- XOR with 0 is number itself: a ^ 0 = a
- XOR is commutative: a ^ b = b ^ a
- XOR is associative: (a ^ b) ^ c = a ^ (b ^ c)
- Paired elements cancel out in XOR

## 5. VARIATIONS & EXTENSIONS
- One element appears three times
- Find missing number
- Multiple unique elements
- Different frequency patterns

## 6. INTERVIEW INSIGHTS
- Clarify: "Do all other elements appear exactly twice?"
- Edge cases: empty array, single element
- Time complexity: O(N) vs O(N log N) sorting
- Space complexity: O(1) vs O(N) hash map

## 7. COMMON MISTAKES
- Using hash map when XOR suffices
- Not understanding XOR properties
- Integer overflow concerns
- Not handling empty array

## 8. OPTIMIZATION STRATEGIES
- XOR is optimal for this specific pattern
- Early termination for obvious cases
- Bit manipulation for other patterns
- In-place operations

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like pairing up socks:**
- You have socks (array elements) in pairs
- One sock is unpaired (single number)
- When you pair up all matching socks, one remains
- XOR operation acts like pairing: a ^ a = 0
- The unpaired sock remains after all pairings cancel

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array where one element appears once, others twice
2. **Goal**: Find the single element
3. **Output**: The element that appears only once

#### Phase 2: Key Insight Recognition
- **"What's special about twice?"** → Pairs can cancel via XOR
- **"How does XOR help?"** → a ^ a = 0, 0 ^ a = a
- **"What's the pattern?"** → XOR all elements = single element
- **"Why does it work?"** → Pairs cancel, leaving single

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use XOR:
1. Start with result = 0
2. For each element in array:
   - result = result ^ element
3. All paired elements cancel out
4. Only single element remains"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 or handle gracefully
- **Single element**: Return that element
- **All pairs**: Return 0 (no single element)
- **Large numbers**: XOR handles naturally

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
nums = [2, 2, 1]

Human thinking:
"Let's XOR all elements:

Start with result = 0

Process 2:
- result = 0 ^ 2 = 2

Process 2:
- result = 2 ^ 2 = 0
- (First pair cancels out)

Process 1:
- result = 0 ^ 1 = 1
- (Single element remains)

Final result: 1 ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Paired elements cancel via XOR
- **Why it's efficient**: Single pass, O(1) space
- **Why it's correct**: XOR properties guarantee cancellation

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use hash map?"** → O(N) space vs O(1) for XOR
2. **"What about sorting?"** → O(N log N) time vs O(N) for XOR
3. **"How to understand XOR?"** → Remember key properties
4. **"What about other patterns?"** → XOR only works for this specific pattern

### Real-World Analogy
**Like finding the unpaired item in inventory:**
- You have items sold in pairs (two of each)
- One item was sold only once (single)
- You want to find which item is unpaired
- Use XOR to pair up all duplicate items
- The unpaired item remains after all pairings cancel
- This identifies the single sold item

### Human-Readable Pseudocode
```
function singleNumber(nums):
    result = 0
    
    for num in nums:
        result = result ^ num
    
    return result
```

### Execution Visualization

### Example: nums = [2, 2, 1]
```
XOR Process:
Start: result = 0

Step 1: result = 0 ^ 2 = 2
Step 2: result = 2 ^ 2 = 0  (pair cancels)
Step 3: result = 0 ^ 1 = 1  (single remains)

Final: result = 1 ✓

The pairing visualization:
2 ^ 2 = 0 (pair 1 cancels)
0 ^ 1 = 1 (single element remains)
```

### Key Visualization Points:
- **XOR properties** enable pair cancellation
- **Single pass** through array
- **O(1) space** no extra data structures
- **Bitwise operation** is very efficient

### Memory Layout Visualization:
```
XOR Operation Evolution:
result: 0 → 2 → 0 → 1

Process:
0 ^ 2 = 2
2 ^ 2 = 0  (first pair cancels)
0 ^ 1 = 1  (single element remains)

Bit-level view (for 2 and 1):
0: 0000
2: 0010
2: 0010
1: 0001

0 ^ 2 = 0010 (2)
2 ^ 2 = 0000 (0)  ← pair cancels
0 ^ 1 = 0001 (1)  ← single remains
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) iterations
- **Each Iteration**: O(1) XOR operation
- **Total**: O(N) time, O(1) space
- **Optimal**: Cannot do better than O(N) for this problem
- **vs Hash Map**: O(N) time, O(N) space (more memory)
- **vs Sorting**: O(N log N) time, O(1) space (slower)
*/
