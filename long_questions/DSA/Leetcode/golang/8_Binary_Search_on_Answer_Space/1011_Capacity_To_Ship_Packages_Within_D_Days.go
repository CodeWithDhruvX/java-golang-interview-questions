package main

import "fmt"

// 1011. Capacity To Ship Packages Within D Days
// Time: O(N log M), Space: O(1) where M is sum(weights)
func shipWithinDays(weights []int, days int) int {
	left, right := maxWeight(weights), sumWeights(weights)
	result := right
	
	for left <= right {
		mid := left + (right-left)/2
		
		if canShip(weights, days, mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	
	return result
}

func canShip(weights []int, days, capacity int) bool {
	daysNeeded := 1
	currentLoad := 0
	
	for _, weight := range weights {
		if currentLoad+weight <= capacity {
			currentLoad += weight
		} else {
			daysNeeded++
			currentLoad = weight
			if daysNeeded > days {
				return false
			}
		}
	}
	
	return daysNeeded <= days
}

func maxWeight(weights []int) int {
	max := 0
	for _, weight := range weights {
		if weight > max {
			max = weight
		}
	}
	return max
}

func sumWeights(weights []int) int {
	sum := 0
	for _, weight := range weights {
		sum += weight
	}
	return sum
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Binary Search on Answer Space
- **Search Space**: Minimum possible capacity to maximum possible capacity
- **Monotonic Property**: If capacity X works, any capacity > X also works
- **Feasibility Check**: Test if given capacity allows shipping within D days
- **Binary Search**: Narrow down to minimum feasible capacity

## 2. PROBLEM CHARACTERISTICS
- **Optimization Problem**: Find minimum capacity that satisfies constraints
- **Monotonic Feasibility**: Higher capacity always makes shipping easier
- **Partition Constraint**: Must ship all packages within exactly D days
- **Capacity Bounds**: Between max weight and total weight

## 3. SIMILAR PROBLEMS
- Split Array Largest Sum (LeetCode 410) - Similar partitioning problem
- Koko Eating Bananas (LeetCode 875) - Binary search on eating speed
- Minimum Number of Days to Make m Bouquets (LeetCode 1482) - Binary search on days
- Minimum Time to Complete Trips (LeetCode 2187) - Binary search on time

## 4. KEY OBSERVATIONS
- **Monotonic Property**: If capacity C works, any C' > C also works
- **Search Bounds**: Lower bound = max weight, upper bound = sum of all weights
- **Feasibility Logic**: Greedy packing - fill current day until capacity exceeded
- **Optimal Solution**: Binary search finds minimum feasible capacity

## 5. VARIATIONS & EXTENSIONS
- **Exact Days**: Must use exactly D days (not at most)
- **Multiple Constraints**: Add weight limits per package type
- **Package Order**: Allow reordering of packages
- **Variable Days**: Find minimum days for given capacity

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can packages be reordered? Must use exactly D days?"
- Edge cases: single package, D equals number of packages, D equals 1
- Time complexity: O(N log S) where S is sum of weights
- Space complexity: O(1) additional space

## 7. COMMON MISTAKES
- Not setting correct search bounds (left = max weight, right = sum)
- Wrong feasibility logic (not greedy packing)
- Off-by-one errors in binary search
- Not handling edge cases properly
- Using linear search instead of binary search

## 8. OPTIMIZATION STRATEGIES
- **Binary Search**: O(N log S) time, optimal for this problem
- **Greedy Feasibility**: O(N) time to check if capacity works
- **Early Termination**: Stop when days exceed limit
- **Efficient Bounds**: Tighten search space bounds

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the smallest truck that can deliver all packages:**
- You have packages with different weights and a delivery deadline
- You need to find the smallest truck capacity that works
- If a truck of size X works, any larger truck will also work
- You can test different truck sizes and narrow down to the minimum
- This is a perfect binary search problem

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of package weights, number of days D
2. **Goal**: Find minimum ship capacity to deliver all packages in D days
3. **Constraint**: Must ship packages in order, cannot reorder
4. **Output**: Minimum capacity value

#### Phase 2: Key Insight Recognition
- **"Binary search natural fit"** → Monotonic feasibility property
- **"Greedy packing works"** → Fill each day until capacity exceeded
- **"Search bounds"** → Between max weight and total weight
- **"Feasibility check"** → Test if capacity allows D-day shipping

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the minimum capacity that works.
If capacity X works, any larger capacity will also work.
This monotonic property means I can use binary search.
I'll search between max weight (minimum possible) and total weight (maximum possible).
For each capacity, I'll greedily pack packages to see if it fits in D days."
```

#### Phase 4: Edge Case Handling
- **Single package**: Capacity must equal package weight
- **D = number of packages**: Capacity equals max weight
- **D = 1**: Capacity equals total weight
- **All packages same weight**: Capacity = max weight if D ≥ packages

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
weights = [1,2,3,4,5,6,7,8,9,10], D = 5

Human thinking:
"I need to find the minimum capacity.
Search range: max=10 to sum=55

Test capacity 32:
Day 1: 1+2+3+4+5+6+7 = 28 (add 8 would exceed 32)
Day 2: 8+9+10 = 27
Total days = 2 ≤ 5 ✓ Capacity 32 works, try smaller

Test capacity 21:
Day 1: 1+2+3+4+5+6 = 21
Day 2: 7+8 = 15
Day 3: 9+10 = 19
Total days = 3 ≤ 5 ✓ Capacity 21 works, try smaller

Test capacity 15:
Day 1: 1+2+3+4+5 = 15
Day 2: 6+7 = 13
Day 3: 8 = 8
Day 4: 9 = 9
Day 5: 10 = 10
Total days = 5 ≤ 5 ✓ Capacity 15 works, try smaller

Test capacity 14:
Day 1: 1+2+3+4+5 = 14
Day 2: 6+7 = 13
Day 3: 8 = 8
Day 4: 9 = 9
Day 5: 10 = 10
Total days = 5 ≤ 5 ✓ Capacity 14 works, try smaller

Test capacity 13:
Day 1: 1+2+3+4 = 10
Day 2: 5+6 = 11
Day 3: 7 = 7
Day 4: 8 = 8
Day 5: 9 = 9
Day 6: 10 = 10
Total days = 6 > 5 ✗ Capacity 13 doesn't work

Minimum capacity = 15 ✓"
```

#### Phase 6: Intuition Validation
- **Why binary search works**: Monotonic feasibility property
- **Why greedy packing works**: Optimal for given capacity
- **Why search bounds**: Minimum = max weight, maximum = total weight
- **Why O(N log S)**: N for feasibility check, log S for binary search

### Common Human Pitfalls & How to Avoid Them
1. **"Why not try all capacities?"** → Too slow, binary search is optimal
2. **"Should I use DP?"** → No, greedy works for feasibility check
3. **"What about search bounds?"** → Must be max weight to total weight
4. **"Can I optimize further?"** → Binary search is already optimal

### Real-World Analogy
**Like finding the right size moving truck:**
- You have boxes of different sizes to move in a certain number of trips
- You want the smallest truck that can handle all boxes in the required trips
- If a small truck works, a larger one will definitely work
- You can test different truck sizes and find the minimum that works
- This is exactly what binary search does for you

### Human-Readable Pseudocode
```
function shipWithinDays(weights, D):
    left = max(weights)  // Minimum possible capacity
    right = sum(weights) // Maximum possible capacity
    result = right
    
    while left <= right:
        mid = left + (right - left) / 2
        
        if canShip(weights, D, mid):
            result = mid
            right = mid - 1  // Try smaller capacity
        else:
            left = mid + 1   // Need larger capacity
    
    return result

function canShip(weights, D, capacity):
    days = 1
    currentLoad = 0
    
    for weight in weights:
        if currentLoad + weight <= capacity:
            currentLoad += weight
        else:
            days += 1
            currentLoad = weight
            if days > D:
                return false
    
    return days <= D
```

### Execution Visualization

### Example: weights = [3,2,2,4,1,4], D = 3
```
Binary Search Process:
Search range: [4, 16] (max weight to total weight)

Test capacity 10:
Day 1: 3+2+2 = 7
Day 2: 4+1 = 5
Day 3: 4 = 4
Days used = 3 ≤ 3 ✓ Try smaller

Test capacity 7:
Day 1: 3+2 = 5
Day 2: 2+4 = 6
Day 3: 1 = 1
Day 4: 4 = 4
Days used = 4 > 3 ✗ Need larger

Test capacity 8:
Day 1: 3+2+2 = 7
Day 2: 4+1 = 5
Day 3: 4 = 4
Days used = 3 ≤ 3 ✓ Try smaller

Test capacity 6:
Day 1: 3+2 = 5
Day 2: 2+4 = 6
Day 3: 1 = 1
Day 4: 4 = 4
Days used = 4 > 3 ✗ Need larger

Minimum capacity = 8 ✓
```

### Key Visualization Points:
- **Monotonic property**: Higher capacity always works if lower works
- **Greedy packing**: Fill each day optimally for given capacity
- **Binary search**: Narrow down to minimum feasible capacity
- **Search bounds**: Between max weight and total weight

### Memory Layout Visualization:
```
Capacity Test Visualization:
capacity = 8, weights = [3,2,2,4,1,4]

Day 1: [3,2,2] = 7 (next 4 would exceed 8)
Day 2: [4,1] = 5 (next 4 would exceed 8)  
Day 3: [4] = 4

Total days = 3 ≤ D ✓ Capacity 8 works!
```

### Time Complexity Breakdown:
- **Binary Search**: O(log S) iterations where S = sum(weights)
- **Feasibility Check**: O(N) time per iteration
- **Total Time**: O(N log S)
- **Space Complexity**: O(1) additional space

### Alternative Approaches:

#### 1. Linear Search (O(N × S) time, O(1) space)
```go
func shipWithinDaysLinear(weights []int, days int) int {
    for capacity := maxWeight(weights); capacity <= sumWeights(weights); capacity++ {
        if canShip(weights, days, capacity) {
            return capacity
        }
    }
    return sumWeights(weights)
}
```
- **Pros**: Simple to understand
- **Cons**: Too slow for large inputs

#### 2. Dynamic Programming (O(N × D × S) time, O(D × S) space)
```go
func shipWithinDaysDP(weights []int, days int) int {
    n := len(weights)
    dp := make([][]int, days+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
        for j := range dp[i] {
            dp[i][j] = math.MaxInt32
        }
    }
    dp[0][0] = 0
    
    for d := 1; d <= days; d++ {
        for i := 1; i <= n; i++ {
            sum := 0
            for j := i; j >= 1; j-- {
                sum += weights[j-1]
                if dp[d-1][j-1] != math.MaxInt32 {
                    dp[d][i] = min(dp[d][i], max(dp[d-1][j-1], sum))
                }
            }
        }
    }
    
    return dp[days][n]
}
```
- **Pros**: Finds exact optimal solution
- **Cons**: Too complex and slow for this problem

#### 3. Priority Queue Approach (O(N log N) time, O(N) space)
```go
func shipWithinDaysPQ(weights []int, days int) int {
    // This approach doesn't directly solve the problem
    // but could be used for related partitioning problems
    return -1
}
```
- **Pros**: Useful for some partitioning variants
- **Cons**: Not applicable to this specific problem

### Extensions for Interviews:
- **Exact Days Constraint**: Must use exactly D days
- **Package Reordering**: Allow reordering of packages
- **Multiple Constraints**: Add weight limits per day
- **Variable Days**: Find minimum days for given capacity
- **Parallel Shipping**: Multiple ships can work simultaneously
*/
func main() {
	// Test cases
	testCases := []struct {
		weights []int
		days    int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{[]int{3, 2, 2, 4, 1, 4}, 3},
		{[]int{1, 2, 3, 1, 1}, 4},
		{[]int{10, 50, 100, 100, 50, 10}, 5},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{100}, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 9},
		{[]int{5, 5, 5, 5, 5}, 3},
	}
	
	for i, tc := range testCases {
		result := shipWithinDays(tc.weights, tc.days)
		fmt.Printf("Test Case %d: weights=%v, days=%d -> Min capacity: %d\n", 
			i+1, tc.weights, tc.days, result)
	}
}
