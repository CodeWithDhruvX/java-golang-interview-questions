package main

import "fmt"

// 739. Daily Temperatures
// Time: O(N), Space: O(N)
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{} // Store indices
	
	for i := 0; i < n; i++ {
		// While current temperature is greater than temperature at stack top
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	
	// Remaining indices in stack have no warmer day (result is already 0)
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Monotonic Stack for Next Greater Element
- **Monotonic Decreasing Stack**: Maintain stack of decreasing temperatures
- **Index Tracking**: Store indices instead of temperature values
- **Distance Calculation**: Calculate days until warmer temperature
- **Immediate Resolution**: Resolve when warmer temperature found

## 2. PROBLEM CHARACTERISTICS
- **Next Warmer Day**: Find days until warmer temperature
- **Forward Looking**: Only consider future days
- **Distance Calculation**: Need number of days, not the warmer day itself
- **No Warmer Day**: Return 0 if no warmer day exists

## 3. SIMILAR PROBLEMS
- Next Greater Element I (LeetCode 496) - Similar next greater pattern
- Next Greater Element II (LeetCode 503) - Circular array version
- Largest Rectangle in Histogram (LeetCode 84) - Monotonic stack application
- Trapping Rain Water (LeetCode 42) - Similar bar-based problem

## 4. KEY OBSERVATIONS
- **Stack property**: Stack maintains decreasing temperatures
- **Index storage**: Store indices to calculate distances
- **Immediate resolution**: Calculate distance when warmer day found
- **Zero initialization**: Default to 0 for days with no warmer temperature

## 5. VARIATIONS & EXTENSIONS
- **Circular year**: Wrap around to beginning of year
- **Multiple queries**: Query for specific date ranges
- **Temperature threshold**: Days until temperature exceeds threshold
- **Cooling days**: Days until cooler temperature

## 6. INTERVIEW INSIGHTS
- Always clarify: "What about equal temperatures? Can temperatures be negative?"
- Edge cases: empty array, single day, monotonic temperatures
- Time complexity: O(N) time, O(N) space
- Stack invariant: Always maintain decreasing temperatures

## 7. COMMON MISTAKES
- Not initializing result array with zeros
- Using nested loops (O(N²) time)
- Storing temperature values instead of indices
- Not calculating distance correctly
- Forgetting to process remaining stack elements

## 8. OPTIMIZATION STRATEGIES
- **Monotonic stack**: O(N) time, O(N) space
- **Index tracking**: Store indices for distance calculation
- **Single pass**: Process each temperature once
- **Early resolution**: Calculate distances immediately when possible

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like waiting for a warmer day:**
- You have daily temperature readings for a period
- For each day, you want to know how many days until the next warmer day
- You scan forward in time, keeping track of days waiting for warmer weather
- When you find a warmer day, you can tell all the previous cooler days how long they waited
- If you reach the end without finding a warmer day, those days wait forever (0 days)

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of daily temperatures
2. **Goal**: For each day, find days until next warmer temperature
3. **Constraint**: Only consider future days
4. **Output**: Array of days until warmer temperature (0 if none)

#### Phase 2: Key Insight Recognition
- **"Monotonic stack natural fit"** → Maintain decreasing temperatures
- **"Index tracking"** → Store indices to calculate distances
- **"Immediate resolution"** → Calculate distance when warmer day found
- **"Forward processing"** → Process from left to right

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find how many days until warmer weather for each day.
I'll scan from left to right, keeping track of days waiting for warmer weather.
I'll use a stack to store indices of days with decreasing temperatures.
When I find a warmer day, I can resolve all the cooler days behind it.
The distance is simply the difference in indices."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty array
- **Single day**: Return [0] (no future days)
- **Monotonic decreasing**: All zeros (no warmer days)
- **Equal temperatures**: Treat as not warmer

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: [73, 74, 75, 71, 69, 72, 76, 73]

Human thinking:
"I'll process with a stack of indices:

Initialize: result = [0,0,0,0,0,0,0,0], stack = []

Day 0 (73°F): stack = [0]
Day 1 (74°F): 74 > 73, pop 0, result[0] = 1-0 = 1, stack = []
           push 1, stack = [1]
Day 2 (75°F): 75 > 74, pop 1, result[1] = 2-1 = 1, stack = []
           push 2, stack = [2]
Day 3 (71°F): 71 < 75, push 3, stack = [2, 3]
Day 4 (69°F): 69 < 71, push 4, stack = [2, 3, 4]
Day 5 (72°F): 72 > 69, pop 4, result[4] = 5-4 = 1, stack = [2, 3]
           72 > 71, pop 3, result[3] = 5-3 = 2, stack = [2]
           72 < 75, push 5, stack = [2, 5]
Day 6 (76°F): 76 > 72, pop 5, result[5] = 6-5 = 1, stack = [2]
           76 > 75, pop 2, result[2] = 6-2 = 4, stack = []
           push 6, stack = [6]
Day 7 (73°F): 73 < 76, push 7, stack = [6, 7]

Remaining stack elements have no warmer day:
result[6] = 0, result[7] = 0 (already 0)

Final result: [1, 1, 4, 2, 1, 1, 0, 0] ✓"
```

#### Phase 6: Intuition Validation
- **Why monotonic stack works**: Maintains days waiting for warmer weather
- **Why index tracking works**: Enables distance calculation
- **Why immediate resolution works**: Calculate distances when warmer day found
- **Why O(N) time**: Each day pushed and popped at most once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use nested loops?"** → O(N²) time, too slow for large inputs
2. **"Should I store temperatures?"** → No, store indices for distance calculation
3. **"What about equal temperatures?"** → Not considered warmer, keep in stack
4. **"Can I optimize further?"** → Monotonic stack is already optimal

### Real-World Analogy
**Like waiting for a warmer day for outdoor activities:**
- You have daily weather forecasts for the coming month
- For each day, you want to know how many days until it's warm enough for your outdoor activity
- You look ahead day by day, noting which days are still waiting for warmer weather
- When you find a warmer day, you can tell all the previous cooler days how long they need to wait
- If you reach the end of the month without finding a warmer day, those days are out of luck

### Human-Readable Pseudocode
```
function dailyTemperatures(temperatures):
    n = len(temperatures)
    result = array of n zeros
    stack = []  // store indices
    
    for i from 0 to n-1:
        while stack not empty and temperatures[i] > temperatures[stack.top]:
            prevIndex = stack.pop()
            result[prevIndex] = i - prevIndex
        stack.push(i)
    
    return result
```

### Execution Visualization

### Example: temperatures = [73, 74, 75, 71, 69, 72, 76, 73]
```
Stack Evolution and Result Updates:
Initialize: result = [0,0,0,0,0,0,0,0], stack = []

Day 0 (73): stack = [0]
Day 1 (74): 74 > 73, pop 0, result[0] = 1, stack = [], push 1 → stack = [1]
Day 2 (75): 75 > 74, pop 1, result[1] = 1, stack = [], push 2 → stack = [2]
Day 3 (71): 71 < 75, push 3 → stack = [2, 3]
Day 4 (69): 69 < 71, push 4 → stack = [2, 3, 4]
Day 5 (72): 72 > 69, pop 4, result[4] = 1, stack = [2, 3]
           72 > 71, pop 3, result[3] = 2, stack = [2]
           72 < 75, push 5 → stack = [2, 5]
Day 6 (76): 76 > 72, pop 5, result[5] = 1, stack = [2]
           76 > 75, pop 2, result[2] = 4, stack = []
           push 6 → stack = [6]
Day 7 (73): 73 < 76, push 7 → stack = [6, 7]

Final result: [1, 1, 4, 2, 1, 1, 0, 0]
```

### Key Visualization Points:
- **Stack invariant**: Always maintains decreasing temperatures
- **Index storage**: Store indices for distance calculation
- **Immediate resolution**: Calculate distance when warmer day found
- **Zero default**: Days with no warmer day remain 0

### Memory Layout Visualization:
```
Temperatures: [73, 74, 75, 71, 69, 72, 76, 73]
Indices:       0   1   2   3   4   5   6   7

Stack at day 5: [2, 3, 4, 5]
                ↑  ↑  ↑  ↑
                2  3  4  5 (indices)
                75 71 69 72 (temperatures)

When processing day 6 (76°F):
76 > 72, pop index 5, result[5] = 6-5 = 1
76 > 71, pop index 3, result[3] = 6-3 = 2
76 > 75, pop index 2, result[2] = 6-2 = 4
Stack becomes empty
```

### Time Complexity Breakdown:
- **Monotonic stack**: O(N) time, O(N) space
- **Nested loops**: O(N²) time, O(1) space (too slow)
- **Binary search**: O(N log N) time, O(N) space (more complex)
- **Brute force**: O(N²) time, O(1) space

### Alternative Approaches:

#### 1. Brute Force (O(N²) time, O(1) space)
```go
func dailyTemperaturesBrute(temperatures []int) []int {
    n := len(temperatures)
    result := make([]int, n)
    
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if temperatures[j] > temperatures[i] {
                result[i] = j - i
                break
            }
        }
    }
    
    return result
}
```
- **Pros**: Simple to understand
- **Cons**: O(N²) time, too slow for large inputs

#### 2. Reverse Processing (O(N) time, O(N) space)
```go
func dailyTemperaturesReverse(temperatures []int) []int {
    n := len(temperatures)
    result := make([]int, n)
    stack := []int{} // store indices
    
    for i := n - 1; i >= 0; i-- {
        // Remove indices with temperature <= current
        for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
            stack = stack[:len(stack)-1]
        }
        
        if len(stack) > 0 {
            result[i] = stack[len(stack)-1] - i
        }
        
        stack = append(stack, i)
    }
    
    return result
}
```
- **Pros**: Same complexity, different processing direction
- **Cons**: Slightly more complex logic

#### 3. Bucket Sort Approach (O(N × W) time, O(W) space where W is temperature range)
```go
func dailyTemperaturesBucket(temperatures []int) []int {
    n := len(temperatures)
    result := make([]int, n)
    
    for i := 0; i < n; i++ {
        for temp := temperatures[i] + 1; temp <= 100; temp++ {
            for j := i + 1; j < n; j++ {
                if temperatures[j] == temp {
                    result[i] = j - i
                    break
                }
            }
            if result[i] > 0 {
                break
            }
        }
    }
    
    return result
}
```
- **Pros**: Works for limited temperature ranges
- **Cons**: Still O(N²) in worst case, complex

### Extensions for Interviews:
- **Circular Year**: Next warmer day considering year wrap-around
- **Temperature Threshold**: Days until temperature exceeds specific threshold
- **Cooling Days**: Days until cooler temperature
- **Multiple Queries**: Query for specific date ranges
- **Temperature Groups**: Group days by temperature ranges
*/
func main() {
	// Test cases
	testCases := [][]int{
		{73, 74, 75, 71, 69, 72, 76, 73},
		{30, 40, 50, 60},
		{30, 60, 90},
		{90, 80, 70, 60},
		{55, 60, 65, 70, 75},
		{65, 70, 65, 60, 65},
		{50},
		{},
		{73, 73, 73, 73},
		{30, 40, 30, 50, 30, 60, 30},
	}
	
	for i, temperatures := range testCases {
		result := dailyTemperatures(temperatures)
		fmt.Printf("Test Case %d: %v -> Days until warmer: %v\n", 
			i+1, temperatures, result)
	}
}
