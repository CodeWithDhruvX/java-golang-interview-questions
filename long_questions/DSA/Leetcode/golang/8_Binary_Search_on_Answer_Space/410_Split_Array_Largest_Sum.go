package main

import "fmt"

// 410. Split Array Largest Sum
// Time: O(N log S), Space: O(1) where S is sum(nums)
func splitArray(nums []int, k int) int {
	left, right := maxNum(nums), sumNums(nums)
	result := right
	
	for left <= right {
		mid := left + (right-left)/2
		
		if canSplit(nums, k, mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	
	return result
}

func canSplit(nums []int, k, maxSum int) bool {
	subarrays := 1
	currentSum := 0
	
	for _, num := range nums {
		if currentSum+num <= maxSum {
			currentSum += num
		} else {
			subarrays++
			currentSum = num
			if subarrays > k {
				return false
			}
		}
	}
	
	return subarrays <= k
}

func maxNum(nums []int) int {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func sumNums(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Binary Search on Answer Space
- **Search Space**: Minimum possible largest sum to maximum possible largest sum
- **Monotonic Property**: If max sum X works, any max sum > X also works
- **Feasibility Check**: Test if given max sum allows splitting into k subarrays
- **Binary Search**: Narrow down to minimum feasible largest sum

## 2. PROBLEM CHARACTERISTICS
- **Optimization Problem**: Find minimum possible largest subarray sum
- **Partition Constraint**: Must split array into exactly k subarrays
- **Monotonic Feasibility**: Higher max sum always makes splitting easier
- **Sum Bounds**: Between max element and total sum

## 3. SIMILAR PROBLEMS
- Capacity To Ship Packages Within D Days (LeetCode 1011) - Same pattern
- Koko Eating Bananas (LeetCode 875) - Binary search on eating speed
- Minimum Number of Days to Make m Bouquets (LeetCode 1482) - Binary search on days
- Minimum Time to Complete Trips (LeetCode 2187) - Binary search on time

## 4. KEY OBSERVATIONS
- **Monotonic Property**: If max sum S works, any S' > S also works
- **Search Bounds**: Lower bound = max element, upper bound = total sum
- **Feasibility Logic**: Greedy splitting - accumulate until sum exceeded
- **Optimal Solution**: Binary search finds minimum feasible max sum

## 5. VARIATIONS & EXTENSIONS
- **Exact K Subarrays**: Must use exactly k subarrays (not at most)
- **Variable K**: Find minimum k for given max sum
- **Empty Subarrays**: Allow empty subarrays in partition
- **Weighted Sums**: Consider weighted sum calculations

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can subarrays be empty? Must use exactly k subarrays?"
- Edge cases: single element, k equals array length, k equals 1
- Time complexity: O(N log S) where S is sum of elements
- Space complexity: O(1) additional space

## 7. COMMON MISTAKES
- Not setting correct search bounds (left = max element, right = sum)
- Wrong feasibility logic (not greedy accumulation)
- Off-by-one errors in binary search
- Not handling edge cases properly
- Using DP instead of binary search (overkill)

## 8. OPTIMIZATION STRATEGIES
- **Binary Search**: O(N log S) time, optimal for this problem
- **Greedy Feasibility**: O(N) time to check if max sum works
- **Early Termination**: Stop when subarrays exceed k
- **Efficient Bounds**: Tighten search space bounds

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the fairest way to split candy among kids:**
- You have a line of candy bags with different amounts
- You need to split them among k kids
- You want to minimize the maximum candy any kid gets
- If you allow a maximum of X candies per kid and it works, allowing more will also work
- This monotonic property means you can binary search for the minimum

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of numbers, number of subarrays k
2. **Goal**: Split array into k subarrays with minimum possible largest sum
3. **Constraint**: Must maintain order of elements
4. **Output**: Minimum possible largest subarray sum

#### Phase 2: Key Insight Recognition
- **"Binary search natural fit"** → Monotonic feasibility property
- **"Greedy splitting works"** → Accumulate until sum exceeded
- **"Search bounds"** → Between max element and total sum
- **"Feasibility check"** → Test if max sum allows k-subarray splitting

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the minimum possible largest sum.
If max sum X works, any larger max sum will also work.
This monotonic property means I can use binary search.
I'll search between max element (minimum possible) and total sum (maximum possible).
For each max sum, I'll greedily accumulate elements to see if it fits in k subarrays."
```

#### Phase 4: Edge Case Handling
- **Single element**: Result is the element itself
- **k = array length**: Result is max element
- **k = 1**: Result is total sum
- **All elements equal**: Result = ceil(total/k) or max element

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
nums = [7,2,5,10,8], k = 2

Human thinking:
"I need to find the minimum largest sum.
Search range: max=10 to sum=32

Test max sum 21:
Subarray 1: 7+2+5 = 14 (add 10 would exceed 21)
Subarray 2: 10+8 = 18
Total subarrays = 2 ≤ 2 ✓ Max sum 21 works, try smaller

Test max sum 15:
Subarray 1: 7+2+5 = 14 (add 10 would exceed 15)
Subarray 2: 10 (add 8 would exceed 15)
Subarray 3: 8 = 8
Total subarrays = 3 > 2 ✗ Max sum 15 doesn't work, need larger

Test max sum 18:
Subarray 1: 7+2+5 = 14 (add 10 would exceed 18)
Subarray 2: 10+8 = 18
Total subarrays = 2 ≤ 2 ✓ Max sum 18 works, try smaller

Test max sum 16:
Subarray 1: 7+2+5 = 14 (add 10 would exceed 16)
Subarray 2: 10 (add 8 would exceed 16)
Subarray 3: 8 = 8
Total subarrays = 3 > 2 ✗ Max sum 16 doesn't work, need larger

Minimum largest sum = 18 ✓"
```

#### Phase 6: Intuition Validation
- **Why binary search works**: Monotonic feasibility property
- **Why greedy splitting works**: Optimal for given max sum
- **Why search bounds**: Minimum = max element, maximum = total sum
- **Why O(N log S)**: N for feasibility check, log S for binary search

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use DP?"** → Overkill, binary search is optimal
2. **"Should I try all splits?"** → Too slow, binary search is efficient
3. **"What about search bounds?"** → Must be max element to total sum
4. **"Can I optimize further?"** → Binary search is already optimal

### Real-World Analogy
**Like splitting a bill fairly among friends:**
- You have a list of expenses that need to be split among k friends
- You want to minimize the maximum amount any friend pays
- If you allow a maximum of $X per person and it works, allowing more will also work
- You can test different maximum amounts and find the minimum that works
- This is exactly what binary search does for you

### Human-Readable Pseudocode
```
function splitArray(nums, k):
    left = max(nums)  // Minimum possible largest sum
    right = sum(nums) // Maximum possible largest sum
    result = right
    
    while left <= right:
        mid = left + (right - left) / 2
        
        if canSplit(nums, k, mid):
            result = mid
            right = mid - 1  // Try smaller max sum
        else:
            left = mid + 1   // Need larger max sum
    
    return result

function canSplit(nums, k, maxSum):
    subarrays = 1
    currentSum = 0
    
    for num in nums:
        if currentSum + num <= maxSum:
            currentSum += num
        else:
            subarrays += 1
            currentSum = num
            if subarrays > k:
                return false
    
    return subarrays <= k
```

### Execution Visualization

### Example: nums = [7,2,5,10,8], k = 2
```
Binary Search Process:
Search range: [10, 32] (max element to total sum)

Test max sum 21:
Subarray 1: [7,2,5] = 14
Subarray 2: [10,8] = 18
Subarrays used = 2 ≤ 2 ✓ Try smaller

Test max sum 15:
Subarray 1: [7,2,5] = 14
Subarray 2: [10] = 10
Subarray 3: [8] = 8
Subarrays used = 3 > 2 ✗ Need larger

Test max sum 18:
Subarray 1: [7,2,5] = 14
Subarray 2: [10,8] = 18
Subarrays used = 2 ≤ 2 ✓ Try smaller

Test max sum 16:
Subarray 1: [7,2,5] = 14
Subarray 2: [10] = 10
Subarray 3: [8] = 8
Subarrays used = 3 > 2 ✗ Need larger

Minimum largest sum = 18 ✓
```

### Key Visualization Points:
- **Monotonic property**: Higher max sum always works if lower works
- **Greedy splitting**: Accumulate until max sum exceeded
- **Binary search**: Narrow down to minimum feasible max sum
- **Search bounds**: Between max element and total sum

### Memory Layout Visualization:
```
Max Sum Test Visualization:
maxSum = 18, nums = [7,2,5,10,8]

Subarray 1: [7,2,5] = 14 (next 10 would exceed 18)
Subarray 2: [10,8] = 18

Total subarrays = 2 ≤ k ✓ Max sum 18 works!
```

### Time Complexity Breakdown:
- **Binary Search**: O(log S) iterations where S = sum(nums)
- **Feasibility Check**: O(N) time per iteration
- **Total Time**: O(N log S)
- **Space Complexity**: O(1) additional space

### Alternative Approaches:

#### 1. Dynamic Programming (O(N² × k) time, O(N × k) space)
```go
func splitArrayDP(nums []int, k int) int {
    n := len(nums)
    dp := make([][]int, k+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
        for j := range dp[i] {
            dp[i][j] = math.MaxInt32
        }
    }
    dp[0][0] = 0
    
    // Prefix sums for quick sum calculation
    prefix := make([]int, n+1)
    for i := 1; i <= n; i++ {
        prefix[i] = prefix[i-1] + nums[i-1]
    }
    
    for i := 1; i <= k; i++ {
        for j := 1; j <= n; j++ {
            for l := 0; l < j; l++ {
                currentSum := prefix[j] - prefix[l]
                dp[i][j] = min(dp[i][j], max(dp[i-1][l], currentSum))
            }
        }
    }
    
    return dp[k][n]
}
```
- **Pros**: Finds exact optimal solution
- **Cons**: Too slow for large inputs

#### 2. Linear Search (O(N × S) time, O(1) space)
```go
func splitArrayLinear(nums []int, k int) int {
    for maxSum := maxNum(nums); maxSum <= sumNums(nums); maxSum++ {
        if canSplit(nums, k, maxSum) {
            return maxSum
        }
    }
    return sumNums(nums)
}
```
- **Pros**: Simple to understand
- **Cons**: Too slow for large inputs

#### 3. Priority Queue Approach (O(N log N) time, O(N) space)
```go
func splitArrayPQ(nums []int, k int) int {
    // This approach doesn't directly solve the problem
    // but could be used for related partitioning problems
    return -1
}
```
- **Pros**: Useful for some partitioning variants
- **Cons**: Not applicable to this specific problem

### Extensions for Interviews:
- **Exact K Subarrays**: Must use exactly k subarrays
- **Variable K**: Find minimum k for given max sum
- **Empty Subarrays**: Allow empty subarrays in partition
- **Weighted Sums**: Consider weighted sum calculations
- **Multiple Constraints**: Add constraints on subarray lengths
*/
func main() {
	// Test cases
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{7, 2, 5, 10, 8}, 2},
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{1, 4, 4}, 3},
		{[]int{10, 5, 2, 7, 8, 9}, 3},
		{[]int{2, 3, 1, 2, 4, 3}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{[]int{100}, 1},
		{[]int{1, 1, 1, 1, 1}, 5},
		{[]int{5, 5, 5, 5}, 2},
	}
	
	for i, tc := range testCases {
		result := splitArray(tc.nums, tc.k)
		fmt.Printf("Test Case %d: nums=%v, k=%d -> Min largest sum: %d\n", 
			i+1, tc.nums, tc.k, result)
	}
}
