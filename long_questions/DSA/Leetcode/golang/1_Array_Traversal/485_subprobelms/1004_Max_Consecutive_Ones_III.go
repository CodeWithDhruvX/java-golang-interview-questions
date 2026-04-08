package main

import "fmt"

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Sliding Window with Dynamic Boundaries
- **Window Expansion**: Right pointer moves forward, counting zeros
- **Window Contraction**: Left pointer moves when zero count exceeds K
- **Zero Tracking**: Maintain count of zeros within current window
- **Dynamic Adjustment**: Window size adjusts based on constraint satisfaction

## 2. PROBLEM CHARACTERISTICS
- **Binary Array**: Contains only 0s and 1s (can be flipped)
- **Constraint-Based**: Maximum K zeros allowed in window
- **Optimization**: Find maximum window size satisfying constraint
- **Two-Pointer Technique**: Classic sliding window pattern

## 3. SIMILAR PROBLEMS
- Max Consecutive Ones (LeetCode 485) - K=0 case
- Max Consecutive Ones II (LeetCode 487) - K=1 case
- Longest Substring with At Most K Distinct Characters
- Longest Subarray with At Most K Zero Flips
- Fruit Into Baskets (LeetCode 904) - similar sliding window

## 4. KEY OBSERVATIONS
- **Zero as Resource**: Each zero consumes one of K available flips
- **Window Validity**: Window is valid while zeroCount ≤ K
- **Monotonic Expansion**: Right pointer only moves forward
- **Greedy Contraction**: Left pointer moves only when necessary

## 5. VARIATIONS & EXTENSIONS
- **Different Array Types**: General arrays with condition checking
- **Multiple Constraints**: At most K zeros AND at most M ones
- **Circular Arrays**: Need to handle wrap-around logic
- **Return Subarray**: Return actual subarray instead of length

## 6. INTERVIEW INSIGHTS
- **Clarify Constraints**: "Can we flip any zeros or only consecutive ones?"
- **Edge Cases**: K ≥ array length, K = 0, empty array
- **Space Complexity**: O(1) - only pointers and counters
- **Time Complexity**: O(N) - each element visited at most twice

## 7. COMMON MISTAKES
- **Forgetting to update max before shrinking window**
- **Incorrect zero counting logic**
- **Off-by-one errors in window size calculation**
- **Not handling K ≥ array length case**
- **Using nested loops instead of sliding window**

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(N) time, O(1) space
- **Early termination**: If K ≥ total zeros, return array length
- **Cache optimization**: Sequential memory access pattern
- **Parallel processing**: For extremely large arrays

## 9. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like a flexible ruler:**
- You have a ruler (window) that can expand and contract
- You can "ignore" up to K zeros (flip them to ones)
- Expand the ruler as far right as possible
- If you exceed K ignored zeros, slide the left edge forward
- Keep track of the maximum ruler length you achieved

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Binary array + integer K (max zeros to flip)
2. **Goal**: Find longest subarray with at most K zeros
3. **Output**: Maximum length of such subarray
4. **Key Insight**: We're looking for the largest window containing ≤ K zeros

#### Phase 2: Key Insight Recognition
- **"Window thinking"** → This is about finding optimal subarray
- **Constraint satisfaction** → At most K zeros in window
- **Two pointers needed** → Left and right boundaries
- **Dynamic adjustment** → Window expands and contracts based on zeros

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the longest sequence where I can flip at most K zeros.
Let me use two pointers - left and right. I'll expand right as far as possible,
counting zeros. If I have more than K zeros, I'll move left forward until
I'm back to K or fewer zeros. The window size (right-left+1) tells me the
current valid sequence length. I'll track the maximum throughout."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0
- **K = 0**: Same as basic max consecutive ones
- **K ≥ total zeros**: Return array length
- **All zeros**: Return min(K, array length)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Array: [1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0], K = 2

Human thinking:
"Start with left=0, right=0, zeroCount=0, maxLength=0

right=0: nums[0]=1 → zeroCount=0, window=[1], length=1, maxLength=1
right=1: nums[1]=1 → zeroCount=0, window=[1,1], length=2, maxLength=2
right=2: nums[2]=1 → zeroCount=0, window=[1,1,1], length=3, maxLength=3
right=3: nums[3]=0 → zeroCount=1, window=[1,1,1,0], length=4, maxLength=4
right=4: nums[4]=0 → zeroCount=2, window=[1,1,1,0,0], length=5, maxLength=5
right=5: nums[5]=0 → zeroCount=3 > K! Need to shrink:
         left=0: nums[0]=1 → zeroCount=3, left=1
         left=1: nums[1]=1 → zeroCount=3, left=2
         left=2: nums[2]=1 → zeroCount=3, left=3
         left=3: nums[3]=0 → zeroCount=2, left=4
         Now window=[0,0,1], length=3

right=6: nums[6]=1 → zeroCount=2, window=[0,0,1,1], length=4
right=7: nums[7]=1 → zeroCount=2, window=[0,0,1,1,1], length=5
right=8: nums[8]=1 → zeroCount=2, window=[0,0,1,1,1,1], length=6, maxLength=6
right=9: nums[9]=1 → zeroCount=2, window=[0,0,1,1,1,1,1], length=7, maxLength=7
right=10: nums[10]=0 → zeroCount=3 > K! Need to shrink:
          left=4: nums[4]=0 → zeroCount=2, left=5
          Now window=[0,1,1,1,1,1,0], length=6

Done! Maximum length was 7."
```

#### Phase 6: Intuition Validation
- **Why it works**: We maintain invariant of ≤ K zeros in window
- **Why it's efficient**: Each element processed at most twice
- **Why it's correct**: We explore all possible valid windows

### Common Human Pitfalls & How to Avoid Them
1. **"Should I store all zeros positions?"** → No, just count them
2. **"When do I update max length?"** → After ensuring window validity
3. **"What if K is very large?"** → Handle as special case
4. **"Do I need to actually flip zeros?"** → No, just count them

### Real-World Analogy
**Like managing a budget:**
- You have a budget of K "zero flips" you can spend
- Each zero you encounter costs 1 from your budget
- Expand your shopping cart (window) as much as possible
- If you exceed budget, remove items from the left until within budget
- Track the maximum cart size you achieved

### Human-Readable Pseudocode
```
function maxConsecutiveOnesIII(binaryArray, K):
    left = 0
    zeroCount = 0
    maxLength = 0
    
    for right from 0 to binaryArray.length-1:
        if binaryArray[right] == 0:
            zeroCount = zeroCount + 1
        
        while zeroCount > K:
            if binaryArray[left] == 0:
                zeroCount = zeroCount - 1
            left = left + 1
        
        currentLength = right - left + 1
        if currentLength > maxLength:
            maxLength = currentLength
    
    return maxLength
```

### Execution Visualization

### Example: [1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0], K = 2
```
Initial: left=0, zeroCount=0, maxLength=0

Step 1: right=0, nums[0]=1
→ zeroCount=0, window=[1], length=1 → maxLength=1

Step 2: right=1, nums[1]=1
→ zeroCount=0, window=[1,1], length=2 → maxLength=2

Step 3: right=2, nums[2]=1
→ zeroCount=0, window=[1,1,1], length=3 → maxLength=3

Step 4: right=3, nums[3]=0
→ zeroCount=1, window=[1,1,1,0], length=4 → maxLength=4

Step 5: right=4, nums[4]=0
→ zeroCount=2, window=[1,1,1,0,0], length=5 → maxLength=5

Step 6: right=5, nums[5]=0
→ zeroCount=3 > K=2 → Shrink window:
   left=0: nums[0]=1 → zeroCount=3, left=1
   left=1: nums[1]=1 → zeroCount=3, left=2
   left=2: nums[2]=1 → zeroCount=3, left=3
   left=3: nums[3]=0 → zeroCount=2, left=4
→ window=[0,0,1], length=3 → maxLength=5

Step 7: right=6, nums[6]=1
→ zeroCount=2, window=[0,0,1,1], length=4 → maxLength=5

Step 8: right=7, nums[7]=1
→ zeroCount=2, window=[0,0,1,1,1], length=5 → maxLength=5

Step 9: right=8, nums[8]=1
→ zeroCount=2, window=[0,0,1,1,1,1], length=6 → maxLength=6

Step 10: right=9, nums[9]=1
→ zeroCount=2, window=[0,0,1,1,1,1,1], length=7 → maxLength=7

Step 11: right=10, nums[10]=0
→ zeroCount=3 > K=2 → Shrink window:
    left=4: nums[4]=0 → zeroCount=2, left=5
→ window=[0,1,1,1,1,1,0], length=6 → maxLength=7

Final: Return maxLength = 7
```

### Key Visualization Points:
- **Window boundaries** move monotonically forward
- **Zero count** never exceeds K after adjustment
- **Maximum length** updated when window is valid
- **Shrinking happens** only when constraint violated

### Memory Layout Visualization:
```
Array:    [1][1][1][0][0][0][1][1][1][1][0]
Position:  0  1  2  3  4  5  6  7  8  9 10
Right:     ^  ^  ^  ^  ^  ^  ^  ^  ^  ^  ^
Left:      ^              ^           ^
Window:    [1,1,1,0,0]    [0,0,1,1,1,1,1]
Zeros:     0  0  0  1  2  3→2  2  2  2  2→3→2
MaxLen:    1  2  3  4  5  5  5  5  6  7  7
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each element visited at most twice
- **Constant Space**: O(1) - only pointers and counters
- **No Additional Data Structures**: Pure sliding window
- **Cache Friendly**: Sequential memory access pattern

### Alternative Approaches:
1. **Prefix Sum + Binary Search**: O(N log N) time, O(N) space
2. **Store Zero Positions**: O(N) time, O(N) space
3. **Dynamic Programming**: Overkill for this problem
4. **Two Pass with Preprocessing**: More complex than needed

### Extensions for Interviews:
- **Return the actual subarray**: Store start/end indices
- **Handle multiple test cases efficiently**: Reuse logic
- **Find all maximum windows**: Collect all start positions
- **Generalize to other conditions**: Replace zero check with any predicate
*/

// LeetCode 1004: Max Consecutive Ones III
// Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array
// if you can flip at most k 0's.

// Time Complexity: O(n)
// Space Complexity: O(1)

// MaxConsecutiveOnesIII uses sliding window approach
func MaxConsecutiveOnesIII(nums []int, k int) int {
	left := 0
	maxLength := 0
	zeroCount := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		// If we have more than k zeros, shrink the window from left
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// Update maximum length
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

// BruteForceApproach - O(n^2) time complexity
func BruteForceApproach(nums []int, k int) int {
	maxLength := 0

	for i := 0; i < len(nums); i++ {
		zeroCount := 0
		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				zeroCount++
			}
			if zeroCount > k {
				break
			}
			currentLength := j - i + 1
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}

// OptimizedApproach - Same as sliding window approach
func OptimizedApproach(nums []int, k int) int {
	return MaxConsecutiveOnesIII(nums, k)
}

func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 1, 0, 0}, 0, 3},
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 0, 1}, 4, 4},
		{[]int{1, 0, 1, 0, 1}, 1, 3},
	}

	fmt.Println("LeetCode 1004: Max Consecutive Ones III")
	fmt.Println("=========================================")

	for i, tc := range testCases {
		result := MaxConsecutiveOnesIII(tc.nums, tc.k)
		bruteForce := BruteForceApproach(tc.nums, tc.k)
		optimized := OptimizedApproach(tc.nums, tc.k)

		fmt.Printf("\nTest Case %d:\n", i+1)
		fmt.Printf("Input: nums = %v, k = %d\n", tc.nums, tc.k)
		fmt.Printf("Expected: %d\n", tc.want)
		fmt.Printf("Sliding Window: %d\n", result)
		fmt.Printf("Brute Force: %d\n", bruteForce)
		fmt.Printf("Optimized: %d\n", optimized)

		if result == tc.want && bruteForce == tc.want && optimized == tc.want {
			fmt.Printf("✓ All approaches passed!\n")
		} else {
			fmt.Printf("✗ Some approaches failed!\n")
		}
	}

	// Demonstrate the sliding window process
	fmt.Println("\n\nSliding Window Process Demonstration:")
	fmt.Println("=====================================")
	demoNums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	demoK := 2
	fmt.Printf("Array: %v\n", demoNums)
	fmt.Printf("K: %d\n", demoK)
	fmt.Printf("Result: %d\n", MaxConsecutiveOnesIII(demoNums, demoK))
}
