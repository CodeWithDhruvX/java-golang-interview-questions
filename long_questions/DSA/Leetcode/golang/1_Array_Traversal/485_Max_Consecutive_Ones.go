package main

import "fmt"

// 485. Max Consecutive Ones
// Time: O(N), Space: O(1)
func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	currentCount := 0
	
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 0
		}
	}
	
	return maxCount
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Sliding Window with Reset
- **Current Counter**: Tracks length of current consecutive ones sequence
- **Max Counter**: Maintains the maximum length found so far
- **Reset Logic**: When zero encountered, current counter resets to 0

## 2. PROBLEM CHARACTERISTICS
- **Binary Array**: Contains only 0s and 1s
- **Consecutive Pattern**: Looking for maximum streak of 1s
- **Single Pass**: Can be solved in one traversal
- **State Management**: Need to track current streak and maximum

## 3. SIMILAR PROBLEMS
- Max Consecutive Ones II (LeetCode 487) - can flip at most one zero
- Longest Subarray with 1s after at most K zeros (LeetCode 1004)
- Consecutive Characters (LeetCode 1446)
- Maximum Length of Subarray With Positive Product (LeetCode 1567)

## 4. KEY OBSERVATIONS
- **Reset on Zero**: Zero acts as a boundary between sequences
- **Update on One**: Each 1 extends the current sequence
- **Max Tracking**: Need to compare current with maximum after each extension
- **Edge Cases**: Empty array, all zeros, all ones

## 5. VARIATIONS & EXTENSIONS
- Allow K zeros to be flipped: sliding window with count
- Find longest subarray with at most K zeros: more complex
- Find all sequences of maximum length: store starting indices
- Circular array: need to handle wrap-around

## 6. INTERVIEW INSIGHTS
- Always clarify: "Is the array binary (only 0s and 1s)?"
- Edge cases: empty array, single element, all same values
- Space complexity: O(1) - only two integer variables
- Time complexity: O(N) - single pass through array

## 7. COMMON MISTAKES
- Forgetting to reset current counter when zero encountered
- Not updating max counter before resetting current counter
- Using extra space unnecessarily (stack, additional arrays)
- Off-by-one errors in loop boundaries
- Not handling empty array case

## 8. OPTIMIZATION STRATEGIES
- Current solution is already optimal for binary arrays
- For general arrays, could use hash maps for pattern detection
- For very large arrays, consider parallel processing
- Cache-friendly access patterns for better performance

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like counting consecutive wins:**
- You're tracking a winning streak in a game
- Each win (1) extends your current streak
- Each loss (0) breaks the streak and you start over
- You want to know your best streak ever
- You keep two numbers: current streak and best streak

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary array (0s and 1s)
2. **Goal**: Find longest sequence of consecutive 1s
3. **Output**: Maximum length of consecutive 1s

#### Phase 2: Key Insight Recognition
- **"Streak thinking"** → This is about tracking consecutive patterns
- **Reset mechanism** → Zeros act as natural breakpoints
- **Two trackers needed** → Current streak and best streak
- **Single pass sufficient** → No need for complex data structures

#### Phase 3: Strategy Development
```
Human thought process:
"I need to keep track of how many 1s I've seen in a row.
Let's call this 'currentCount'. When I see a 1, I'll increment it.
When I see a 0, I'll reset it to 0 because the streak is broken.
I also need to remember the longest streak I've ever seen,
so I'll keep 'maxCount' and update it whenever currentCount exceeds it."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 (no elements, no streak)
- **All zeros**: Return 0 (no ones at all)
- **All ones**: Return array length (entire array is one streak)
- **Single element**: Return 1 if it's 1, 0 if it's 0

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [1, 1, 0, 1, 1, 1]

Human thinking:
"Okay, I start with current streak = 0, best streak = 0.

Position 0: It's 1 → current streak becomes 1
           Best streak was 0, now it's 1
Position 1: It's 1 → current streak becomes 2
           Best streak was 1, now it's 2
Position 2: It's 0 → streak broken! current streak resets to 0
           Best streak stays 2
Position 3: It's 1 → current streak becomes 1
           Best streak stays 2
Position 4: It's 1 → current streak becomes 2
           Best streak stays 2
Position 5: It's 1 → current streak becomes 3
           Best streak was 2, now it's 3

Done! My best streak was 3 consecutive 1s."
```

#### Phase 6: Intuition Validation
- **Why it works**: We accurately track each streak and compare with best
- **Why it's efficient**: Single pass, minimal state tracking
- **Why it's correct**: We consider every possible streak exactly once

### Common Human Pitfalls & How to Avoid Them
1. **"Should I store all streaks?"** → No, just need the maximum
2. **"When should I update max?"** → After extending current streak, not before reset
3. **"What about negative numbers?"** → Problem specifies binary array
4. **"Do I need to remember positions?"** → No, just need the length

### Real-World Analogy
**Like monitoring a production line:**
- You're tracking consecutive successful products (1s)
- Each defective product (0) breaks the success streak
- You want to know the longest run of perfect products
- You keep a counter for current run and best run ever
- When defects appear, you restart counting

### Human-Readable Pseudocode
```
function findMaxConsecutiveOnes(binaryArray):
    maxStreak = 0
    currentStreak = 0
    
    for each element in binaryArray:
        if element is 1:
            currentStreak = currentStreak + 1
            if currentStreak > maxStreak:
                maxStreak = currentStreak
        else:  // element is 0
            currentStreak = 0
    
    return maxStreak
```

### Execution Visualization

### Example: [1, 1, 0, 1, 1, 1]
```
Initial: maxCount = 0, currentCount = 0

Step 1: i=0, nums[0]=1
→ currentCount++ (1)
→ currentCount (1) > maxCount (0) → maxCount = 1
State: maxCount=1, currentCount=1

Step 2: i=1, nums[1]=1
→ currentCount++ (2)
→ currentCount (2) > maxCount (1) → maxCount = 2
State: maxCount=2, currentCount=2

Step 3: i=2, nums[2]=0
→ Reset currentCount = 0
→ maxCount remains 2
State: maxCount=2, currentCount=0

Step 4: i=3, nums[3]=1
→ currentCount++ (1)
→ currentCount (1) <= maxCount (2) → no change to max
State: maxCount=2, currentCount=1

Step 5: i=4, nums[4]=1
→ currentCount++ (2)
→ currentCount (2) <= maxCount (2) → no change to max
State: maxCount=2, currentCount=2

Step 6: i=5, nums[5]=1
→ currentCount++ (3)
→ currentCount (3) > maxCount (2) → maxCount = 3
State: maxCount=3, currentCount=3

Final: Return maxCount = 3
```

### Key Visualization Points:
- **currentCount** resets to 0 on every zero
- **maxCount** only increases when currentCount exceeds it
- **Streak boundaries** are clearly defined by zeros
- **Final answer** is the maximum streak length found

### Memory Layout Visualization:
```
Array:    [1][1][0][1][1][1]
Position:  0  1  2  3  4  5
Current:   1  2  0  1  2  3
Max:       1  2  2  2  2  3
                    ^
                    Maximum streak found here
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each element visited once
- **Constant Space**: O(1) - only two integer variables
- **No Additional Data Structures**: Pure counting approach
- **Cache Friendly**: Sequential memory access pattern

### Alternative Approaches:
1. **Sliding Window with K zeros**: More complex, allows flipping zeros
2. **Dynamic Programming**: Overkill for this simple problem
3. **Stack-based**: Unnecessary complexity for binary arrays

### Extensions for Interviews:
- **What if array contains other numbers?** → Filter first or add conditions
- **Find all maximum streak positions?** → Store starting indices when updating max
- **Allow K zeros to be ignored?** → Use sliding window with zero counter
*/

func main() {
	// Test cases
	testCases := [][]int{
		{1, 1, 0, 1, 1, 1},
		{1, 0, 1, 1, 0, 1},
		{0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 1, 0, 1, 0, 1},
	}
	
	for i, nums := range testCases {
		result := findMaxConsecutiveOnes(nums)
		fmt.Printf("Test Case %d: %v -> Max Consecutive Ones: %d\n", i+1, nums, result)
	}
}
