package main

import "fmt"

// 303. Range Sum Query - Immutable
// Time: O(1) query, Space: O(N) for preprocessing
type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	return NumArray{prefixSum: prefixSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Prefix Sum Preprocessing
- **Preprocessing Phase**: Build prefix sum array in O(N) time
- **Query Phase**: Answer range sum queries in O(1) time
- **Immutable Data**: Array doesn't change after preprocessing
- **Space-Time Tradeoff**: Use O(N) space for O(1) queries

## 2. PROBLEM CHARACTERISTICS
- **Range Queries**: Multiple sum queries on subarrays
- **Immutable Array**: No updates to original array
- **Preprocessing**: One-time setup cost for fast queries
- **Difference Calculation**: Use prefix sums to compute range sums

## 3. SIMILAR PROBLEMS
- Range Sum Query 2D - Immutable (LeetCode 304) - 2D version
- Range Sum Query - Mutable (LeetCode 307) - With updates
- Subarray Sum Equals K (LeetCode 560) - Prefix sum with hash map
- Product of Array Except Self (LeetCode 238) - Prefix/suffix products

## 4. KEY OBSERVATIONS
- **Prefix sum definition**: prefix[i] = sum of elements[0..i-1]
- **Range sum formula**: sum[left..right] = prefix[right+1] - prefix[left]
- **One-time cost**: Preprocessing enables O(1) queries
- **Immutable constraint**: No need to handle updates

## 5. VARIATIONS & EXTENSIONS
- **Mutable array**: Need segment tree or BIT for updates
- **2D array**: 2D prefix sums for rectangular regions
- **Range updates**: Need difference array with BIT
- **Multiple operations**: Combine sum, min, max queries

## 6. INTERVIEW INSIGHTS
- Always clarify: "Is the array mutable? How many queries?"
- Edge cases: empty array, single element, invalid ranges
- Time complexity: O(N) preprocessing, O(1) per query
- Space complexity: O(N) for prefix sum array

## 7. COMMON MISTAKES
- Not handling empty array case
- Using 0-based vs 1-based indexing incorrectly
- Not validating query bounds
- Recomputing sums for each query (O(N²) total)
- Off-by-one errors in prefix sum calculation

## 8. OPTIMIZATION STRATEGIES
- **Preprocessing**: Build prefix sum once
- **Constant time queries**: Use difference formula
- **Memory optimization**: Use minimal extra space
- **Batch processing**: Handle multiple queries efficiently

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like cumulative mileage tracking:**
- You're tracking daily mileage on a road trip
- You want to know total mileage between any two days
- Instead of recalculating each time, you keep a running total
- For any day range, you subtract the earlier total from the later total
- This gives you the mileage for that specific period

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers, multiple range sum queries
2. **Goal**: Answer sum queries efficiently
3. **Constraint**: Array is immutable (no updates)
4. **Output**: Sum of elements in specified range

#### Phase 2: Key Insight Recognition
- **"Preprocessing natural fit"** → Build prefix sums once
- **"Difference formula"** → Range sum = prefix[right+1] - prefix[left]
- **"Immutable advantage"** → No need to handle updates
- **"Query efficiency"** → O(1) per query after preprocessing

#### Phase 3: Strategy Development
```
Human thought process:
"I need to answer many range sum queries efficiently.
Instead of calculating each sum from scratch, I'll preprocess once.
I'll build a prefix sum array where each entry is the sum of all previous elements.
For any range [left, right], I can compute the sum as:
prefix[right+1] - prefix[left]
This gives me O(1) queries after O(N) preprocessing."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 for any query
- **Single element**: Return element value for [0,0] query
- **Invalid ranges**: Handle bounds checking
- **Negative numbers**: Works correctly with prefix sums

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: nums = [-2, 0, 3, -5, 2, -1]
Query: sumRange(2, 5)

Human thinking:
"First, I'll build the prefix sum array:
prefix[0] = 0 (sum of nothing)
prefix[1] = -2 (sum of [-2])
prefix[2] = -2 + 0 = -2 (sum of [-2,0])
prefix[3] = -2 + 0 + 3 = 1 (sum of [-2,0,3])
prefix[4] = 1 + (-5) = -4 (sum of [-2,0,3,-5])
prefix[5] = -4 + 2 = -2 (sum of [-2,0,3,-5,2])
prefix[6] = -2 + (-1) = -3 (sum of [-2,0,3,-5,2,-1])

Now for query sumRange(2,5):
sum = prefix[5+1] - prefix[2] = prefix[6] - prefix[2]
sum = -3 - (-2) = -1

Let me verify: nums[2] + nums[3] + nums[4] + nums[5] = 3 + (-5) + 2 + (-1) = -1 ✓"
```

#### Phase 6: Intuition Validation
- **Why prefix sums work**: Cumulative sums enable range calculations
- **Why O(1) queries**: Simple arithmetic operation
- **Why preprocessing needed**: One-time cost for many queries
- **Why difference formula works**: Subtraction cancels prefix contributions

### Common Human Pitfalls & How to Avoid Them
1. **"Why not compute each query directly?"** → O(N²) total time for N queries
2. **"Should I use segment tree?"** → Overkill for immutable array
3. **"What about updates?"** → Different problem (mutable version)
4. **"Can I optimize space?"** → O(N) space is minimal for this approach

### Real-World Analogy
**Like tracking cumulative sales data:**
- You have daily sales figures for a store
- You want to know total sales for any time period
- Instead of adding up days each time, you keep running totals
- For any date range, you subtract the earlier total from the later total
- This gives you sales for that specific period instantly

### Human-Readable Pseudocode
```
class NumArray:
    prefixSum = []
    
    constructor(nums):
        prefixSum[0] = 0
        for i from 0 to len(nums)-1:
            prefixSum[i+1] = prefixSum[i] + nums[i]
    
    sumRange(left, right):
        return prefixSum[right+1] - prefixSum[left]
```

### Execution Visualization

### Example: nums = [-2, 0, 3, -5, 2, -1]
```
Prefix Sum Construction:
prefix[0] = 0
prefix[1] = 0 + (-2) = -2
prefix[2] = -2 + 0 = -2
prefix[3] = -2 + 3 = 1
prefix[4] = 1 + (-5) = -4
prefix[5] = -4 + 2 = -2
prefix[6] = -2 + (-1) = -3

Final prefix array: [0, -2, -2, 1, -4, -2, -3]

Query Examples:
sumRange(0, 2) = prefix[3] - prefix[0] = 1 - 0 = 1
sumRange(2, 5) = prefix[6] - prefix[2] = -3 - (-2) = -1
sumRange(0, 5) = prefix[6] - prefix[0] = -3 - 0 = -3
```

### Key Visualization Points:
- **Prefix accumulation**: Building cumulative sums
- **Difference calculation**: Subtracting to get range sums
- **Constant time queries**: Simple arithmetic operation
- **Preprocessing benefit**: One-time cost for many queries

### Memory Layout Visualization:
```
Original Array: [-2, 0, 3, -5, 2, -1]
Index:          0   1  2   3  4   5

Prefix Array:   [0, -2, -2, 1, -4, -2, -3]
Index:          0   1   2   3   4   5   6

Query sumRange(2,5):
prefix[6] - prefix[2] = -3 - (-2) = -1
                    ↑         ↑
                 sum(0..5)  sum(0..1)
                 = sum(2..5)
```

### Time Complexity Breakdown:
- **Preprocessing**: O(N) time, O(N) space
- **Each query**: O(1) time, O(1) space
- **Total for Q queries**: O(N + Q) time, O(N) space
- **Naive approach**: O(N × Q) time, O(1) space

### Alternative Approaches:

#### 1. Naive Approach (O(N × Q) time, O(1) space)
```go
func sumRangeNaive(nums []int, left, right int) int {
    sum := 0
    for i := left; i <= right; i++ {
        sum += nums[i]
    }
    return sum
}
```
- **Pros**: No preprocessing, simple to implement
- **Cons**: O(N) per query, inefficient for many queries

#### 2. Segment Tree (O(N) preprocessing, O(log N) query) - For Mutable Arrays
```go
type SegmentNode struct {
    sum   int
    left  *SegmentNode
    right *SegmentNode
}

func (node *SegmentNode) query(left, right int) int {
    // Segment tree query implementation
    return 0 // Simplified for brevity
}
```
- **Pros**: Handles updates, O(log N) queries
- **Cons**: Complex implementation, overkill for immutable arrays

#### 3. Binary Indexed Tree (O(N) preprocessing, O(log N) query) - For Mutable Arrays
```go
type BIT struct {
    tree []int
    n    int
}

func (bit *BIT) query(index int) int {
    sum := 0
    for index > 0 {
        sum += bit.tree[index]
        index -= index & -index
    }
    return sum
}
```
- **Pros**: Handles updates, simpler than segment tree
- **Cons**: O(log N) queries, more complex than prefix sums

### Extensions for Interviews:
- **2D Prefix Sums**: Range sum queries in 2D matrices
- **Mutable Arrays**: Use segment trees or BIT for updates
- **Range Updates**: Use difference arrays with BIT
- **Multiple Operations**: Combine sum, min, max queries
- **Streaming Data**: Handle data that arrives over time
*/
func main() {
	// Test cases
	testCases := []struct {
		nums []int
		queries []struct {
			left  int
			right int
		}
	}{
		{
			[]int{-2, 0, 3, -5, 2, -1},
			[]struct {
				left  int
				right int
			}{{0, 2}, {2, 5}, {0, 5}},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]struct {
				left  int
				right int
			}{{0, 0}, {0, 4}, {1, 3}, {2, 2}},
		},
		{
			[]int{},
			[]struct {
				left  int
				right int
			}{},
		},
		{
			[]int{-1, -1, -1, -1},
			[]struct {
				left  int
				right int
			}{{0, 3}, {1, 2}, {0, 1}},
		},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: nums=%v\n", i+1, tc.nums)
		numArray := Constructor(tc.nums)
		
		for j, query := range tc.queries {
			result := numArray.SumRange(query.left, query.right)
			fmt.Printf("  Query %d: sumRange(%d, %d) = %d\n", 
				j+1, query.left, query.right, result)
		}
		fmt.Println()
	}
}
