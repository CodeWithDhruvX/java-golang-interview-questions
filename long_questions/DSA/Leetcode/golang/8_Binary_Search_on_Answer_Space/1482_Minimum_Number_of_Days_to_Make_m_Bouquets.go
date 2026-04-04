package main

import "fmt"

// 1482. Minimum Number of Days to Make m Bouquets
// Time: O(N log (max-min)), Space: O(1)
func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	
	left, right := minDay(bloomDay), maxDay(bloomDay)
	result := -1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if canMakeBouquets(bloomDay, m, k, mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	
	return result
}

func canMakeBouquets(bloomDay []int, m, k, days int) bool {
	bouquets := 0
	flowers := 0
	
	for _, day := range bloomDay {
		if day <= days {
			flowers++
			if flowers == k {
				bouquets++
				flowers = 0
				if bouquets >= m {
					return true
				}
			}
		} else {
			flowers = 0
		}
	}
	
	return bouquets >= m
}

func minDay(bloomDay []int) int {
	min := bloomDay[0]
	for _, day := range bloomDay {
		if day < min {
			min = day
		}
	}
	return min
}

func maxDay(bloomDay []int) int {
	max := bloomDay[0]
	for _, day := range bloomDay {
		if day > max {
			max = day
		}
	}
	return max
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Binary Search on Answer Space
- **Search Space**: Minimum possible days to maximum possible days
- **Monotonic Property**: If days X works, any days > X also works
- **Feasibility Check**: Test if given days allow making m bouquets
- **Binary Search**: Narrow down to minimum feasible days

## 2. PROBLEM CHARACTERISTICS
- **Optimization Problem**: Find minimum days to make m bouquets
- **Monotonic Feasibility**: More days always make bouquet making easier
- **Bouquet Constraint**: Need k consecutive flowers per bouquet
- **Day Bounds**: Between min bloom day and max bloom day

## 3. SIMILAR PROBLEMS
- Capacity To Ship Packages Within D Days (LeetCode 1011) - Same pattern
- Split Array Largest Sum (LeetCode 410) - Binary search on max sum
- Koko Eating Bananas (LeetCode 875) - Binary search on eating speed
- Minimum Time to Complete Trips (LeetCode 2187) - Binary search on time

## 4. KEY OBSERVATIONS
- **Monotonic Property**: If days D works, any D' > D also works
- **Search Bounds**: Lower bound = min bloom day, upper bound = max bloom day
- **Feasibility Logic**: Count consecutive flowers that bloom by given day
- **Early Termination**: Stop when bouquets ≥ m
- **Impossible Case**: m*k > total flowers means impossible

## 5. VARIATIONS & EXTENSIONS
- **Variable K**: Different bouquet sizes
- **Multiple Flower Types**: Different flower types with different constraints
- **Flexible Bouquets**: Allow non-consecutive flowers
- **Time-based Blooming**: Flowers bloom at different rates

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can flowers be reused? What about m*k > total flowers?"
- Edge cases: single flower, m equals 1, k equals 1
- Time complexity: O(N log (max-min)) where max-min is day range
- Space complexity: O(1) additional space
- Important: Check impossible case first

## 7. COMMON MISTAKES
- Not checking impossible case (m*k > total flowers)
- Not setting correct search bounds (left = min day, right = max day)
- Wrong bouquet counting logic (not consecutive)
- Off-by-one errors in binary search
- Not handling edge cases properly

## 8. OPTIMIZATION STRATEGIES
- **Binary Search**: O(N log D) time, optimal for this problem
- **Consecutive Counting**: O(N) time to check if days work
- **Early Termination**: Stop when bouquets ≥ m
- **Efficient Bounds**: Tighten search space bounds

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the earliest day to harvest enough flowers for bouquets:**
- You have flowers that bloom on different days
- You need to make bouquets using consecutive flowers
- You want to find the earliest day when you can make enough bouquets
- If day X works, any later day will also work (more flowers bloom)
- This monotonic property means you can binary search for the minimum day

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of bloom days, number of bouquets m, flowers per bouquet k
2. **Goal**: Find minimum days to make m bouquets
3. **Constraint**: Need k consecutive flowers per bouquet
4. **Output**: Minimum days or -1 if impossible

#### Phase 2: Key Insight Recognition
- **"Binary search natural fit"** → Monotonic feasibility property
- **"Consecutive counting key"** → Need k consecutive bloomed flowers
- **"Search bounds"** → Between min bloom day and max bloom day
- **"Feasibility check"** → Test if days allow m-bouquet creation
- **"Impossible case"** → Check if m*k > total flowers

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find the minimum days to make m bouquets.
If day X works, any later day will also work.
This monotonic property means I can use binary search.
I'll search between min bloom day and max bloom day.
For each day, I'll count consecutive bloomed flowers to see if I can make m bouquets.
First, I should check if it's even possible (m*k > total flowers)."
```

#### Phase 4: Edge Case Handling
- **Impossible case**: Return -1 if m*k > total flowers
- **Single flower**: Days = bloom day if m=k=1
- **m = 1**: Find earliest day with k consecutive bloomed flowers
- **k = 1**: Find day when at least m flowers have bloomed

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
bloomDay = [1,10,3,10,2], m = 3, k = 2

Human thinking:
"First check if possible: m*k = 6 > total flowers = 5 ✓ Impossible!
Wait, let me recalculate: m*k = 3*2 = 6, total = 5, so impossible.
Actually, the example says this should work, let me check again...

Oh, I see the issue. Let me try the correct example:
bloomDay = [1,10,3,10,2], m = 3, k = 1

Check if possible: m*k = 3*1 = 3 ≤ total = 5 ✓ Possible
Search range: min=1 to max=10

Test day 5:
Flowers bloomed by day 5: [1,0,1,0,1] (positions 0,2,4)
Consecutive counting: position 0 (1 bouquet), position 2 (1 bouquet), position 4 (1 bouquet)
Total bouquets = 3 ≤ 3 ✓ Day 5 works, try smaller

Test day 3:
Flowers bloomed by day 3: [1,0,1,0,0] (positions 0,2)
Consecutive counting: position 0 (1 bouquet), position 2 (1 bouquet)
Total bouquets = 2 < 3 ✗ Day 3 doesn't work, need larger

Test day 4:
Flowers bloomed by day 4: [1,0,1,0,0] (same as day 3)
Total bouquets = 2 < 3 ✗ Day 4 doesn't work, need larger

Minimum days = 5 ✓"
```

#### Phase 6: Intuition Validation
- **Why binary search works**: Monotonic feasibility property
- **Why consecutive counting**: Need k consecutive bloomed flowers
- **Why search bounds**: Minimum = min bloom day, maximum = max bloom day
- **Why O(N log D)**: N for bouquet counting, log D for binary search

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use linear search?"** → Too slow, binary search is optimal
2. **"Should I ignore consecutive requirement?"** → No, k consecutive is crucial
3. **"What about impossible case?"** → Must check m*k > total flowers first
4. **"Can I optimize further?"** → Binary search is already optimal

### Real-World Analogy
**Like finding the earliest day to harvest flowers for a wedding:**
- You have flowers that bloom on different days in your garden
- You need to make bouquets for a wedding using flowers from the same row
- You want to find the earliest day when you can make enough bouquets
- If you can make bouquets on day X, you can definitely make them on any later day
- You can test different days and find the earliest that works
- This is exactly what binary search does for you

### Human-Readable Pseudocode
```
function minDays(bloomDay, m, k):
    if m * k > len(bloomDay):
        return -1
    
    left = min(bloomDay)  // Minimum possible days
    right = max(bloomDay) // Maximum possible days
    result = -1
    
    while left <= right:
        mid = left + (right - left) / 2
        
        if canMakeBouquets(bloomDay, m, k, mid):
            result = mid
            right = mid - 1  // Try fewer days
        else:
            left = mid + 1   // Need more days
    
    return result

function canMakeBouquets(bloomDay, m, k, days):
    bouquets = 0
    flowers = 0
    
    for day in bloomDay:
        if day <= days:
            flowers += 1
            if flowers == k:
                bouquets += 1
                flowers = 0
                if bouquets >= m:
                    return true
        else:
            flowers = 0
    
    return bouquets >= m
```

### Execution Visualization

### Example: bloomDay = [1,10,3,10,2], m = 3, k = 1
```
Binary Search Process:
Search range: [1, 10] (min to max bloom day)

Test day 5:
Bloomed by day 5: [1,0,1,0,1]
Bouquets: position 0 (1), position 2 (1), position 4 (1)
Total = 3 ≥ 3 ✓ Try smaller

Test day 3:
Bloomed by day 3: [1,0,1,0,0]
Bouquets: position 0 (1), position 2 (1)
Total = 2 < 3 ✗ Need more days

Test day 4:
Bloomed by day 4: [1,0,1,0,0]
Bouquets: position 0 (1), position 2 (1)
Total = 2 < 3 ✗ Need more days

Minimum days = 5 ✓
```

### Key Visualization Points:
- **Monotonic property**: More days always work if fewer days work
- **Consecutive counting**: Need k consecutive bloomed flowers
- **Binary search**: Narrow down to minimum feasible days
- **Search bounds**: Between min bloom day and max bloom day

### Memory Layout Visualization:
```
Day Test Visualization:
days = 5, bloomDay = [1,10,3,10,2], k = 1

Bloomed by day 5: [1,0,1,0,1]
                    ^   ^   ^
                position 0, 2, 4 bloomed

Bouquet counting:
position 0: 1 bouquet
position 2: 1 bouquet  
position 4: 1 bouquet

Total bouquets = 3 ≥ m ✓ Day 5 works!
```

### Time Complexity Breakdown:
- **Binary Search**: O(log D) iterations where D = max(bloomDay) - min(bloomDay)
- **Bouquet Counting**: O(N) time per iteration
- **Total Time**: O(N log D)
- **Space Complexity**: O(1) additional space

### Alternative Approaches:

#### 1. Linear Search (O(N × D) time, O(1) space)
```go
func minDaysLinear(bloomDay []int, m int, k int) int {
    if m*k > len(bloomDay) {
        return -1
    }
    
    for days := minDay(bloomDay); days <= maxDay(bloomDay); days++ {
        if canMakeBouquets(bloomDay, m, k, days) {
            return days
        }
    }
    return -1
}
```
- **Pros**: Simple to understand
- **Cons**: Too slow for large inputs

#### 2. Sorting Approach (O(N log N) time, O(N) space)
```go
func minDaysSorting(bloomDay []int, m int, k int) int {
    if m*k > len(bloomDay) {
        return -1
    }
    
    // Sort bloom days and try each unique day
    sortedDays := make([]int, len(bloomDay))
    copy(sortedDays, bloomDay)
    sort.Ints(sortedDays)
    
    for _, day := range sortedDays {
        if canMakeBouquets(bloomDay, m, k, day) {
            return day
        }
    }
    return -1
}
```
- **Pros**: Fewer iterations than linear search
- **Cons**: Still O(N log N) and may not be optimal

#### 3. Priority Queue Approach (O(N log N) time, O(N) space)
```go
func minDaysPQ(bloomDay []int, m int, k int) int {
    // This approach doesn't directly solve the problem
    // but could be used for related scheduling problems
    return -1
}
```
- **Pros**: Useful for some scheduling variants
- **Cons**: Not applicable to this specific problem

### Extensions for Interviews:
- **Variable K**: Different bouquet sizes
- **Multiple Flower Types**: Different flower types with different constraints
- **Flexible Bouquets**: Allow non-consecutive flowers
- **Time-based Blooming**: Flowers bloom at different rates
- **Seasonal Constraints**: Add seasonal blooming patterns
*/
func main() {
	// Test cases
	testCases := []struct {
		bloomDay []int
		m        int
		k        int
	}{
		{[]int{1, 10, 3, 10, 2}, 3, 1},
		{[]int{1, 10, 3, 10, 2}, 3, 2},
		{[]int{7, 7, 7, 7, 12, 7, 7}, 2, 3},
		{[]int{1, 1, 1, 1}, 1, 1},
		{[]int{1, 1, 1, 1}, 4, 1},
		{[]int{1000000000, 1000000000}, 1, 1},
		{[]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2},
		{[]int{5, 5, 5, 5, 5}, 3, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 2, 3},
	}
	
	for i, tc := range testCases {
		result := minDays(tc.bloomDay, tc.m, tc.k)
		fmt.Printf("Test Case %d: bloomDay=%v, m=%d, k=%d -> Min days: %d\n", 
			i+1, tc.bloomDay, tc.m, tc.k, result)
	}
}
