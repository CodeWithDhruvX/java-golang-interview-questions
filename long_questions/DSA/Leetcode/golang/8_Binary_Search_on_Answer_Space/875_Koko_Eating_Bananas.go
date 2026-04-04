package main

import (
	"fmt"
	"math"
)

// 875. Koko Eating Bananas
// Time: O(N log M), Space: O(1) where M is max(piles)
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, maxPiles(piles)
	result := right
	
	for left <= right {
		mid := left + (right-left)/2
		
		if canEatAll(piles, h, mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	
	return result
}

func canEatAll(piles []int, h, speed int) bool {
	hours := 0
	for _, pile := range piles {
		hours += int(math.Ceil(float64(pile) / float64(speed)))
		if hours > h {
			return false
		}
	}
	return hours <= h
}

func maxPiles(piles []int) int {
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	return max
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Binary Search on Answer Space
- **Search Space**: Minimum possible eating speed to maximum possible eating speed
- **Monotonic Property**: If speed X works, any speed > X also works
- **Feasibility Check**: Test if given speed allows eating all bananas in H hours
- **Binary Search**: Narrow down to minimum feasible eating speed

## 2. PROBLEM CHARACTERISTICS
- **Optimization Problem**: Find minimum eating speed that satisfies time constraint
- **Monotonic Feasibility**: Higher speed always makes eating faster
- **Time Constraint**: Must finish all bananas within H hours
- **Speed Bounds**: Between 1 and maximum pile size

## 3. SIMILAR PROBLEMS
- Capacity To Ship Packages Within D Days (LeetCode 1011) - Same pattern
- Split Array Largest Sum (LeetCode 410) - Binary search on max sum
- Minimum Number of Days to Make m Bouquets (LeetCode 1482) - Binary search on days
- Minimum Time to Complete Trips (LeetCode 2187) - Binary search on time

## 4. KEY OBSERVATIONS
- **Monotonic Property**: If speed S works, any S' > S also works
- **Search Bounds**: Lower bound = 1, upper bound = max pile size
- **Feasibility Logic**: Calculate hours needed for given speed
- **Ceiling Division**: Need to round up hours per pile
- **Optimal Solution**: Binary search finds minimum feasible speed

## 5. VARIATIONS & EXTENSIONS
- **Variable Hours**: Find minimum hours for given speed
- **Multiple Monkeys**: Multiple monkeys eating simultaneously
- **Different Pile Types**: Different banana types with different constraints
- **Time-based Eating**: Speed varies by time of day

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can Koko eat fractional bananas? What about H < piles length?"
- Edge cases: single pile, H equals number of piles, H equals 1
- Time complexity: O(N log M) where M is max pile size
- Space complexity: O(1) additional space
- Important: Use ceiling division for hours calculation

## 7. COMMON MISTAKES
- Not setting correct search bounds (left = 1, right = max pile)
- Wrong hours calculation (not using ceiling division)
- Off-by-one errors in binary search
- Not handling edge cases properly
- Using integer division instead of ceiling division

## 8. OPTIMIZATION STRATEGIES
- **Binary Search**: O(N log M) time, optimal for this problem
- **Ceiling Division**: Use math.Ceil or (pile + speed - 1) / speed
- **Early Termination**: Stop when hours exceed H
- **Efficient Bounds**: Tighten search space bounds

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the minimum work rate to meet a deadline:**
- You have tasks (banana piles) that take different amounts of time
- You need to finish all tasks within a deadline (H hours)
- You want to find the minimum work rate that meets the deadline
- If a work rate of X tasks/hour works, any higher rate will also work
- This monotonic property means you can binary search for the minimum

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of banana piles, hours H
2. **Goal**: Find minimum eating speed (bananas per hour) to finish all bananas in H hours
3. **Constraint**: Can only eat from one pile per hour, must finish pile before moving
4. **Output**: Minimum eating speed

#### Phase 2: Key Insight Recognition
- **"Binary search natural fit"** → Monotonic feasibility property
- **"Ceiling division key"** → Need to round up hours per pile
- **"Search bounds"** → Between 1 and max pile size
- **"Feasibility check"** → Test if speed allows H-hour eating

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the minimum eating speed.
If speed X works, any higher speed will also work.
This monotonic property means I can use binary search.
I'll search between 1 (minimum possible) and max pile (maximum possible).
For each speed, I'll calculate hours needed using ceiling division."
```

#### Phase 4: Edge Case Handling
- **Single pile**: Speed = ceil(pile/H)
- **H = number of piles**: Speed = max pile
- **H = 1**: Speed = total bananas
- **All piles same size**: Speed = ceil(pile * count / H)

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
piles = [3,6,7,11], H = 8

Human thinking:
"I need to find the minimum eating speed.
Search range: 1 to 11 (max pile)

Test speed 6:
Hours needed = ceil(3/6) + ceil(6/6) + ceil(7/6) + ceil(11/6)
             = 1 + 1 + 2 + 2 = 6 ≤ 8 ✓ Speed 6 works, try smaller

Test speed 4:
Hours needed = ceil(3/4) + ceil(6/4) + ceil(7/4) + ceil(11/4)
             = 1 + 2 + 2 + 3 = 8 ≤ 8 ✓ Speed 4 works, try smaller

Test speed 3:
Hours needed = ceil(3/3) + ceil(6/3) + ceil(7/3) + ceil(11/3)
             = 1 + 2 + 3 + 4 = 10 > 8 ✗ Speed 3 doesn't work, need larger

Test speed 5:
Hours needed = ceil(3/5) + ceil(6/5) + ceil(7/5) + ceil(11/5)
             = 1 + 2 + 2 + 3 = 8 ≤ 8 ✓ Speed 5 works, try smaller

Test speed 4 (already tested): works, but we need minimum between 4 and 5
Since 3 doesn't work and 4 works, minimum speed = 4 ✓"
```

#### Phase 6: Intuition Validation
- **Why binary search works**: Monotonic feasibility property
- **Why ceiling division**: Must round up hours per pile
- **Why search bounds**: Minimum = 1, maximum = max pile
- **Why O(N log M)**: N for hours calculation, log M for binary search

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use linear search?"** → Too slow, binary search is optimal
2. **"Should I use regular division?"** → No, need ceiling division
3. **"What about search bounds?"** → Must be 1 to max pile
4. **"Can I optimize further?"** → Binary search is already optimal

### Real-World Analogy
**Like finding the minimum typing speed to finish work on time:**
- You have documents of different lengths to type
- You need to finish all documents within a deadline
- You want to find the minimum typing speed that meets the deadline
- If you can type at X words/minute and finish on time, any faster speed will also work
- You can test different speeds and find the minimum that works
- This is exactly what binary search does for you

### Human-Readable Pseudocode
```
function minEatingSpeed(piles, H):
    left = 1           // Minimum possible speed
    right = max(piles) // Maximum possible speed
    result = right
    
    while left <= right:
        mid = left + (right - left) / 2
        
        if canEatAll(piles, H, mid):
            result = mid
            right = mid - 1  // Try smaller speed
        else:
            left = mid + 1   // Need faster speed
    
    return result

function canEatAll(piles, H, speed):
    hours = 0
    for pile in piles:
        hours += ceil(pile / speed)
        if hours > H:
            return false
    return hours <= H
```

### Execution Visualization

### Example: piles = [30,11,23,4,20], H = 6
```
Binary Search Process:
Search range: [1, 30] (1 to max pile)

Test speed 15:
Hours needed = ceil(30/15) + ceil(11/15) + ceil(23/15) + ceil(4/15) + ceil(20/15)
             = 2 + 1 + 2 + 1 + 2 = 8 > 6 ✗ Need faster

Test speed 23:
Hours needed = ceil(30/23) + ceil(11/23) + ceil(23/23) + ceil(4/23) + ceil(20/23)
             = 2 + 1 + 1 + 1 + 1 = 6 ≤ 6 ✓ Try smaller

Test speed 12:
Hours needed = ceil(30/12) + ceil(11/12) + ceil(23/12) + ceil(4/12) + ceil(20/12)
             = 3 + 1 + 2 + 1 + 2 = 9 > 6 ✗ Need faster

Test speed 18:
Hours needed = ceil(30/18) + ceil(11/18) + ceil(23/18) + ceil(4/18) + ceil(20/18)
             = 2 + 1 + 2 + 1 + 2 = 8 > 6 ✗ Need faster

Test speed 21:
Hours needed = ceil(30/21) + ceil(11/21) + ceil(23/21) + ceil(4/21) + ceil(20/21)
             = 2 + 1 + 2 + 1 + 1 = 7 > 6 ✗ Need faster

Minimum speed = 23 ✓
```

### Key Visualization Points:
- **Monotonic property**: Higher speed always works if lower works
- **Ceiling division**: Must round up hours per pile
- **Binary search**: Narrow down to minimum feasible speed
- **Search bounds**: Between 1 and max pile size

### Memory Layout Visualization:
```
Speed Test Visualization:
speed = 23, piles = [30,11,23,4,20]

Hours calculation:
pile 30: ceil(30/23) = 2 hours
pile 11: ceil(11/23) = 1 hour
pile 23: ceil(23/23) = 1 hour
pile 4:  ceil(4/23) = 1 hour
pile 20: ceil(20/23) = 1 hour

Total hours = 6 ≤ H ✓ Speed 23 works!
```

### Time Complexity Breakdown:
- **Binary Search**: O(log M) iterations where M = max(piles)
- **Hours Calculation**: O(N) time per iteration
- **Total Time**: O(N log M)
- **Space Complexity**: O(1) additional space

### Alternative Approaches:

#### 1. Linear Search (O(N × M) time, O(1) space)
```go
func minEatingSpeedLinear(piles []int, h int) int {
    for speed := 1; speed <= maxPiles(piles); speed++ {
        if canEatAll(piles, h, speed) {
            return speed
        }
    }
    return maxPiles(piles)
}
```
- **Pros**: Simple to understand
- **Cons**: Too slow for large inputs

#### 2. Priority Queue Approach (O(N log N) time, O(N) space)
```go
func minEatingSpeedPQ(piles []int, h int) int {
    // This approach doesn't directly solve the problem
    // but could be used for related scheduling problems
    return -1
}
```
- **Pros**: Useful for some scheduling variants
- **Cons**: Not applicable to this specific problem

#### 3. Mathematical Approach (O(N) time, O(1) space)
```go
func minEatingSpeedMath(piles []int, h int) int {
    // For special cases, can use mathematical formulas
    // but general case requires binary search
    return -1
}
```
- **Pros**: Fast for special cases
- **Cons**: Not generally applicable

### Extensions for Interviews:
- **Multiple Monkeys**: Multiple monkeys eating simultaneously
- **Variable Speed**: Speed changes based on time or pile size
- **Different Constraints**: Add constraints on eating patterns
- **Time-based Scheduling**: Speed varies by time of day
- **Resource Allocation**: Generalize to other resource allocation problems
*/
func main() {
	// Test cases
	testCases := []struct {
		piles []int
		h     int
	}{
		{[]int{3, 6, 7, 11}, 8},
		{[]int{30, 11, 23, 4, 20}, 5},
		{[]int{30, 11, 23, 4, 20}, 6},
		{[]int{1}, 1},
		{[]int{100}, 1},
		{[]int{100}, 100},
		{[]int{312884470}, 312884469},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 10},
		{[]int{5, 10, 15, 20, 25}, 15},
	}
	
	for i, tc := range testCases {
		result := minEatingSpeed(tc.piles, tc.h)
		fmt.Printf("Test Case %d: piles=%v, h=%d -> Min speed: %d\n", 
			i+1, tc.piles, tc.h, result)
	}
}
