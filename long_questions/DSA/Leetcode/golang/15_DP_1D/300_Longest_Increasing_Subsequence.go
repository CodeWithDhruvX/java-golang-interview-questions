package main

import (
	"fmt"
	"math"
)

// 300. Longest Increasing Subsequence
// Time: O(N log N), Space: O(N)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	// tails[i] will be the smallest ending value of an increasing subsequence of length i+1
	tails := make([]int, 0)
	
	for _, num := range nums {
		// Binary search to find the insertion position
		left, right := 0, len(tails)
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}
		
		// If num is larger than all elements, append it
		if left == len(tails) {
			tails = append(tails, num)
		} else {
			// Replace the existing element
			tails[left] = num
		}
	}
	
	return len(tails)
}

// O(N^2) DP approach
func lengthOfLISDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	dp := make([]int, len(nums))
	maxLength := 1
	
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLength = max(maxLength, dp[i])
	}
	
	return maxLength
}

// Patience sorting approach (same as binary search method but more intuitive)
func lengthOfLISPatience(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	piles := make([][]int, 0)
	
	for _, num := range nums {
		// Find the leftmost pile where top card >= num
		left, right := 0, len(piles)
		for left < right {
			mid := left + (right-left)/2
			if piles[mid][len(piles[mid])-1] >= num {
				right = mid
			} else {
				left = mid + 1
			}
		}
		
		if left == len(piles) {
			// Create new pile
			piles = append(piles, []int{num})
		} else {
			// Place on existing pile
			piles[left] = append(piles[left], num)
		}
	}
	
	return len(piles)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Longest Increasing Subsequence
- **Patience Sorting**: Use binary search for efficient updates
- **Binary Search Optimization**: O(N log N) time complexity
- **Tails Array**: Maintain smallest possible tail for each length
- **DP Alternative**: O(N²) DP approach for understanding

## 2. PROBLEM CHARACTERISTICS
- **Subsequence**: Not necessarily contiguous, maintains order
- **Increasing**: Each element must be greater than previous
- **Longest**: Find maximum length, not actual subsequence
- **Order Preservation**: Original relative order must be maintained

## 3. SIMILAR PROBLEMS
- Longest Common Subsequence (LeetCode 1143) - Two sequences
- Russian Doll Envelopes (LeetCode 354) - 2D version
- Increasing Triplet Subsequence (LeetCode 334) - Fixed length
- Number of Longest Increasing Subsequence (LeetCode 673)

## 4. KEY OBSERVATIONS
- **Binary search natural fit**: Need to find insertion point efficiently
- **Tails array insight**: tails[i] = smallest possible tail of length i+1
- **Patience sorting analogy**: Like playing solitaire with cards
- **O(N log N) optimal**: Can't do better than this for LIS

## 5. VARIATIONS & EXTENSIONS
- **Count LIS**: Number of longest increasing subsequences
- **Reconstruct LIS**: Also return actual subsequence
- **Circular LIS**: Array considered circular
- **2D LIS**: Envelopes problem with width and height

## 6. INTERVIEW INSIGHTS
- Always clarify: "Do you need actual subsequence or just length?"
- Edge cases: empty array (0), single element (1), all decreasing (1)
- Time complexity: O(N log N) optimal, O(N²) DP alternative
- Space complexity: O(N) for tails array

## 7. COMMON MISTAKES
- Confusing subsequence with subarray (contiguous)
- Using O(N²) DP when O(N log N) is expected
- Not understanding tails array purpose
- Implementing binary search incorrectly
- Not handling empty array case

## 8. OPTIMIZATION STRATEGIES
- **Binary search**: Essential for O(N log N) performance
- **Patience sorting**: Intuitive card game analogy
- **Early termination**: Not applicable (need to process all)
- **Memory optimization**: Use single array for tails

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like playing patience solitaire with cards:**
- You have a pile of cards (numbers) face down
- You deal them one by one and place them in piles
- Each pile must be in increasing order from bottom to top
- You can place a card on the leftmost pile where top card is ≥ card
- If no such pile exists, start a new pile on the right
- Number of piles = length of longest increasing subsequence

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of numbers
2. **Goal**: Find length of longest increasing subsequence
3. **Constraint**: Subsequence maintains original order
4. **Output**: Length of LIS (integer)

#### Phase 2: Key Insight Recognition
- **"Patience sorting"** → Card game analogy provides intuition
- **"Binary search"** → Need efficient way to find pile position
- **"Tails array"** → Track smallest possible tail for each length
- **"O(N log N) optimal"** → Can't do better than binary search approach

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the longest increasing subsequence.
For each number, I want to know where it fits in existing subsequences.
I'll maintain tails[i] = smallest tail of subsequence of length i+1.
When I see a new number, I find the leftmost tails[j] ≥ number using binary search.
If I find such j, I replace tails[j] with number (better tail).
If not, I append number (new longer subsequence).
The length of tails array is my answer."
```

#### Phase 4: Edge Case Handling
- **Empty array**: 0 length
- **Single element**: 1 length
- **All decreasing**: 1 length (any single element)
- **All increasing**: N length (entire array)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: [10, 9, 2, 5, 3, 7, 101, 18]

Human thinking:
"I'll process numbers one by one:

10: tails = [10] (new pile)
9:  replace tails[0] with 9 (better tail) → tails = [9]
2:  replace tails[0] with 2 → tails = [2]
5:  append to tails → tails = [2, 5]
3:  replace tails[1] with 3 → tails = [2, 3]
7:  append to tails → tails = [2, 3, 7]
101: append to tails → tails = [2, 3, 7, 101]
18: replace tails[3] with 18 → tails = [2, 3, 7, 18]

Length of tails = 4, so LIS length is 4
(One LIS is: [2, 3, 7, 18])"
```

#### Phase 6: Intuition Validation
- **Why binary search works**: Need to find insertion point efficiently
- **Why tails array works**: Maintains optimal tails for each length
- **Why O(N log N)**: N elements, each binary search takes log N
- **Why not O(N²)**: Binary search optimization is crucial

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use DP?"** → Can, but O(N²) is too slow for large inputs
2. **"Should I track actual subsequence?"** → More complex, usually just need length
3. **"What about duplicates?"** → Handle with ≤ in binary search condition
4. **"Can I use built-in binary search?"** → Yes, but need custom implementation

### Real-World Analogy
**Like organizing books by height on shelves:**
- You have books of different heights (numbers)
- You want to arrange them in increasing order on shelves
- Each shelf must have books in increasing height
- When you get a new book, find the leftmost shelf where it fits
- If no shelf fits, start a new shelf on the right
- Number of shelves = longest possible increasing sequence

### Human-Readable Pseudocode
```
function lengthOfLIS(nums):
    if nums is empty:
        return 0
    
    tails = []  // tails[i] = smallest tail of LIS of length i+1
    
    for num in nums:
        // Binary search to find insertion point
        left, right = 0, length(tails)
        while left < right:
            mid = left + (right-left)/2
            if tails[mid] < num:
                left = mid + 1
            else:
                right = mid
        
        // Insert or replace
        if left == length(tails):
            tails.append(num)
        else:
            tails[left] = num
    
    return length(tails)
```

### Execution Visualization

### Example: nums = [10, 9, 2, 5, 3, 7, 101, 18]
```
Tails Array Evolution:
Start: []

10: [10]
9:  [9] (replace 10 with 9)
2:  [2] (replace 9 with 2)
5:  [2, 5] (append 5)
3:  [2, 3] (replace 5 with 3)
7:  [2, 3, 7] (append 7)
101: [2, 3, 7, 101] (append 101)
18: [2, 3, 7, 18] (replace 101 with 18)

Final: [2, 3, 7, 18]
Length: 4

Binary Search Process for each number:
10: left=0, right=0 → append
9:  left=0, right=1 → replace at 0
2:  left=0, right=1 → replace at 0
5:  left=0, right=1 → append at 1
3:  left=0, right=2 → replace at 1
7:  left=0, right=2 → append at 2
101: left=0, right=3 → append at 3
18: left=0, right=4 → replace at 3
```

### Key Visualization Points:
- **Tails array**: Maintains optimal tails for each length
- **Binary search**: Efficiently finds insertion position
- **Replace strategy**: Always keeps smallest possible tail
- **Length tracking**: Final length is LIS length

### Memory Layout Visualization:
```
Processing Flow:
Input:  [10, 9, 2, 5, 3, 7, 101, 18]
Tails:  [10] → [9] → [2] → [2,5] → [2,3] → [2,3,7] → [2,3,7,101] → [2,3,7,18]

Binary Search Visualization:
For 5: tails=[2], left=0, right=1
- mid=0, tails[0]=2 < 5 → left=1
- left=right=1 → append

For 18: tails=[2,3,7,101], left=0, right=4
- mid=2, tails[2]=7 < 18 → left=3
- mid=3, tails[3]=101 ≥ 18 → right=3
- left=right=3 → replace at position 3
```

### Time Complexity Breakdown:
- **Binary search approach**: O(N log N) time, O(N) space
- **DP O(N²) approach**: O(N²) time, O(N) space
- **Patience sorting**: O(N log N) time, O(N) space
- **Reconstruction**: O(N²) time if actual subsequence needed

### Alternative Approaches:

#### 1. O(N²) DP Approach (O(N²) time, O(N) space)
```go
func lengthOfLISDP(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    dp := make([]int, len(nums))
    maxLength := 1
    
    for i := 0; i < len(nums); i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        maxLength = max(maxLength, dp[i])
    }
    
    return maxLength
}
```
- **Pros**: Simple to understand and implement
- **Cons**: Too slow for large N (N > 1000)

#### 2. Reconstruct LIS (O(N²) time, O(N) space)
```go
func reconstructLIS(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    
    dp := make([]int, len(nums))
    parent := make([]int, len(nums))
    maxLength := 0
    maxIndex := 0
    
    for i := 0; i < len(nums); i++ {
        dp[i] = 1
        parent[i] = -1
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] && dp[j]+1 > dp[i] {
                dp[i] = dp[j] + 1
                parent[i] = j
            }
        }
        
        if dp[i] > maxLength {
            maxLength = dp[i]
            maxIndex = i
        }
    }
    
    // Reconstruct LIS
    lis := make([]int, maxLength)
    current := maxIndex
    for i := maxLength - 1; i >= 0; i-- {
        lis[i] = nums[current]
        current = parent[current]
    }
    
    return lis
}
```
- **Pros**: Returns actual LIS, not just length
- **Cons**: O(N²) time complexity

#### 3. Built-in Binary Search (O(N log N) time, O(N) space)
```go
import "sort"

func lengthOfLISBuiltIn(nums []int) int {
    tails := make([]int, 0)
    
    for _, num := range nums {
        idx := sort.SearchInts(tails, num)
        if idx == len(tails) {
            tails = append(tails, num)
        } else {
            tails[idx] = num
        }
    }
    
    return len(tails)
}
```
- **Pros**: Uses built-in binary search, simpler code
- **Cons**: Slightly less efficient than custom implementation

### Extensions for Interviews:
- **Count LIS**: Number of longest increasing subsequences
- **Circular LIS**: Array considered circular (wrap-around)
- **2D LIS**: Russian Doll Envelopes problem
- **K-th LIS**: Find k-th longest increasing subsequence
- **Online LIS**: Process elements as they arrive
*/
func main() {
	// Test cases
	testCases := [][]int{
		{10, 9, 2, 5, 3, 7, 101, 18},
		{0, 1, 0, 3, 2, 3},
		{7, 7, 7, 7, 7, 7, 7},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 3, 6, 7, 9, 4, 10, 5, 6},
		{2, 2, 2, 2, 2},
		{1},
		{},
		{3, 10, 2, 1, 20},
	}
	
	for i, nums := range testCases {
		result1 := lengthOfLIS(nums)
		result2 := lengthOfLISDP(nums)
		result3 := lengthOfLISPatience(nums)
		
		fmt.Printf("Test Case %d: %v\n", i+1, nums)
		fmt.Printf("  Binary Search: %d, DP O(N²): %d, Patience: %d\n\n", 
			result1, result2, result3)
	}
}
