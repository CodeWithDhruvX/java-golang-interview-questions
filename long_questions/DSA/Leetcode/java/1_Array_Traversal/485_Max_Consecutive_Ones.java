import java.util.Arrays;

public class MaxConsecutiveOnes {
    
    // 485. Max Consecutive Ones
    // Time: O(N), Space: O(1)
    public static int findMaxConsecutiveOnes(int[] nums) {
        int maxCount = 0;
        int currentCount = 0;
        
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 1) {
                currentCount++;
                if (currentCount > maxCount) {
                    maxCount = currentCount;
                }
            } else {
                currentCount = 0;
            }
        }
        
        return maxCount;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 1, 0, 1, 1, 1},
            {1, 0, 1, 1, 0, 1},
            {0, 0, 0},
            {1, 1, 1, 1},
            {1, 0, 1, 0, 1, 0, 1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int result = findMaxConsecutiveOnes(testCases[i]);
            System.out.printf("Test Case %d: %s -> Max Consecutive Ones: %d\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Sliding Window with Reset
- **Current Counter**: Tracks length of current consecutive ones sequence
- **Max Counter**: Tracks maximum length found so far
- **Reset Logic**: Reset current counter when zero encountered

## 2. PROBLEM CHARACTERISTICS
- **Binary Array**: Contains only 0s and 1s
- **Consecutive Elements**: Looking for sequences of 1s
- **Return Value**: Maximum length of consecutive ones
- **Single Pass**: Can be solved in one traversal

## 3. SIMILAR PROBLEMS
- Max Consecutive Ones II (allow flipping one zero)
- Longest Subarray with All 1's
- Max Consecutive Characters in String
- Find Longest Increasing Subsequence

## 4. KEY OBSERVATIONS
- Only need to track current streak and maximum streak
- Zero acts as a natural separator
- No need for complex data structures
- Simple state machine: counting or reset

## 5. VARIATIONS & EXTENSIONS
- Allow K zeros to be flipped
- Find longest subarray with at most K zeros
- Return the starting and ending indices
- Handle different data types (strings)

## 6. INTERVIEW INSIGHTS
- Clarify: "Is the array binary?" (contains only 0s and 1s)
- Edge cases: empty array, all zeros, all ones
- Space complexity: O(1) - only two counters
- Time complexity: O(N) - single pass

## 7. COMMON MISTAKES
- Forgetting to update max counter before reset
- Not handling empty array case
- Using extra space unnecessarily
- Off-by-one errors in counting

## 8. OPTIMIZATION STRATEGIES
- Current solution is already optimal
- For multiple queries, consider preprocessing
- For streaming data, maintain running state
- For very large arrays, consider parallel processing

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like counting consecutive sunny days:**
- You're tracking weather patterns (0=rainy, 1=sunny)
- You want to find the longest sunny streak
- When it rains (0), you reset your current streak counter
- You keep track of the longest streak you've seen so far

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary array (0s and 1s)
2. **Goal**: Find longest sequence of consecutive 1s
3. **Output**: Maximum length of consecutive ones

#### Phase 2: Key Insight Recognition
- **"I just need two counters!"** → Current and maximum
- **Zero means reset** → Natural separator
- **One means increment** → Continue current streak
- **Compare and update** → Track the best so far

#### Phase 3: Strategy Development
```
Human thought process:
"I'll keep track of two numbers:
1. currentCount - how many 1s I've seen in a row
2. maxCount - the longest streak I've found so far

When I see a 1, I increment currentCount.
When I see a 0, I reset currentCount to 0.
After each increment, I check if currentCount is the new maximum."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0
- **All zeros**: Return 0
- **All ones**: Return array length
- **Single element**: Return 1 if it's 1, 0 if it's 0

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [1, 1, 0, 1, 1, 1]

Human thinking:
"Okay, currentCount = 0, maxCount = 0.
Position 0: It's 1 → currentCount = 1, maxCount = 1.
Position 1: It's 1 → currentCount = 2, maxCount = 2.
Position 2: It's 0 → reset currentCount = 0, maxCount stays 2.
Position 3: It's 1 → currentCount = 1, maxCount stays 2.
Position 4: It's 1 → currentCount = 2, maxCount stays 2.
Position 5: It's 1 → currentCount = 3, maxCount = 3.

Done! Maximum consecutive ones = 3."
```

#### Phase 6: Intuition Validation
- **Why it works**: We accurately track streaks and reset appropriately
- **Why it's efficient**: Single pass, constant space
- **Why it's correct**: We never miss a potential maximum

### Common Human Pitfalls & How to Avoid Them
1. **"Should I use a sliding window?"** → Simple counters are sufficient
2. **"Do I need to track positions?"** → Only length is required
3. **"When should I update max?"** → After incrementing currentCount
4. **"What about the last streak?"** → Handled by continuous checking

### Real-World Analogy
**Like tracking winning streaks in sports:**
- You're following a team's game results (0=loss, 1=win)
- You want to know their longest winning streak
- Each win extends the current streak
- A loss resets the streak to zero
- You always remember the best streak ever

### Human-Readable Pseudocode
```
function findMaxConsecutiveOnes(binaryArray):
    currentStreak = 0
    maxStreak = 0
    
    for each element in array:
        if element is 1:
            currentStreak++
            maxStreak = max(maxStreak, currentStreak)
        else:
            currentStreak = 0
    
    return maxStreak
```

### Execution Visualization

### Example: [1, 1, 0, 1, 1, 1]
```
Initial: nums = [1, 1, 0, 1, 1, 1], currentCount = 0, maxCount = 0

Step 1: i=0, nums[0]=1
→ currentCount = 1, maxCount = 1
State: currentCount=1, maxCount=1

Step 2: i=1, nums[1]=1
→ currentCount = 2, maxCount = 2
State: currentCount=2, maxCount=2

Step 3: i=2, nums[2]=0
→ currentCount = 0 (reset), maxCount = 2
State: currentCount=0, maxCount=2

Step 4: i=3, nums[3]=1
→ currentCount = 1, maxCount = 2
State: currentCount=1, maxCount=2

Step 5: i=4, nums[4]=1
→ currentCount = 2, maxCount = 2
State: currentCount=2, maxCount=2

Step 6: i=5, nums[5]=1
→ currentCount = 3, maxCount = 3
State: currentCount=3, maxCount=3

Final: Return maxCount = 3
```

### Key Visualization Points:
- **currentCount** resets to 0 on every zero
- **maxCount** only increases when currentCount exceeds it
- **Zero acts** as natural separator between streaks
- **Final result** is the maximum streak length found

### Memory Layout Visualization:
```
Array:        [1][1][0][1][1][1]
currentCount: [1][2][0][1][2][3]
maxCount:     [1][2][2][2][2][3]
                         ^
                         Maximum found
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each element visited once
- **Constant Space**: O(1) - only two integer variables
- **No Additional Data Structures**: Pure counting approach
- **Optimal**: Cannot be improved for single query
*/
