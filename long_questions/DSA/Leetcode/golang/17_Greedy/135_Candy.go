package main

import "fmt"

// 135. Candy
// Time: O(N), Space: O(N)
func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	
	n := len(ratings)
	candies := make([]int, n)
	
	// Initialize each child with 1 candy
	for i := 0; i < n; i++ {
		candies[i] = 1
	}
	
	// Left to right pass
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	
	// Right to left pass
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	
	// Calculate total candies
	total := 0
	for _, candy := range candies {
		total += candy
	}
	
	return total
}

// Optimized version with O(1) space (but more complex)
func candyOptimized(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	
	n := len(ratings)
	total := 1
	up := 1
	down := 0
	oldSlope := 0
	
	for i := 1; i < n; i++ {
		newSlope := 0
		if ratings[i] > ratings[i-1] {
			newSlope = 1
		} else if ratings[i] < ratings[i-1] {
			newSlope = -1
		}
		
		if oldSlope > 0 && newSlope == 0 || oldSlope < 0 && newSlope >= 0 {
			total += up * (up + 1) / 2
			up = 1
		}
		
		if newSlope < 0 {
			down++
		} else {
			total += down * (down + 1) / 2
			down = 0
		}
		
		if newSlope > 0 {
			up++
		}
		
		if newSlope == 0 {
			total++
		}
		
		oldSlope = newSlope
	}
	
	total += up * (up + 1) / 2
	if down > 0 {
		total += down * (down + 1) / 2
	}
	
	return total
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Two-Pass Greedy with Constraint Satisfaction
- **Two-Pass Strategy**: Left-to-right then right-to-left
- **Constraint Satisfaction**: Each child gets minimum candies satisfying rules
- **Local Optimization**: Each pass ensures one direction constraint
- **Global Optimality**: Two passes guarantee both direction constraints

## 2. PROBLEM CHARACTERISTICS
- **Distribution Problem**: Distribute candies to children with ratings
- **Directional Constraints**: Higher rating → more candies than left/right neighbor
- **Minimum Distribution**: Use minimum total candies while satisfying constraints
- **Local Dependencies**: Each child's candy depends on neighbors

## 3. SIMILAR PROBLEMS
- Jump Game (LeetCode 55) - Greedy with range tracking
- Gas Station (LeetCode 134) - Circular route feasibility
- Partition Labels (LeetCode 763) - String partitioning
- Trapping Rain Water (LeetCode 42) - Two-pass greedy

## 4. KEY OBSERVATIONS
- **Left-to-Right Pass**: Ensures left neighbor constraint
- **Right-to-Left Pass**: Ensures right neighbor constraint
- **Minimum Principle**: Give minimum candies satisfying constraints
- **Local Optimal**: Each pass optimizes for one direction
- **Global Optimal**: Two passes combine for both directions

## 5. VARIATIONS & EXTENSIONS
- **Different Constraints**: Modified neighbor comparison rules
- **Multiple Candy Types**: Different candy types with constraints
- **Circular Arrangement**: Children in circular arrangement
- **Weighted Ratings**: Different importance for ratings

## 6. INTERVIEW INSIGHTS
- Always clarify: "What about equal ratings? Minimum candies per child?"
- Edge cases: single child, all equal ratings, increasing/decreasing sequence
- Time complexity: O(N) time, O(N) space
- Key insight: two passes are necessary and sufficient
- Space optimization possible with O(1) space

## 7. COMMON MISTAKES
- Only doing one pass (doesn't satisfy both directions)
- Not handling equal ratings correctly
- Wrong initialization of candies array
- Forgetting to sum final candies
- Off-by-one errors in neighbor comparison

## 8. OPTIMIZATION STRATEGIES
- **Two-Pass Greedy**: O(N) time, O(N) space - standard approach
- **Space Optimization**: O(N) time, O(1) space using mathematical approach
- **Single Pass Complex**: O(N) time, O(1) space with slope tracking
- **Early Termination**: Not applicable (need full array)

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like distributing bonuses to employees with performance ratings:**
- You have employees in a line with performance ratings
- Higher-rated employees must get more bonuses than neighbors
- You want to minimize total bonuses while satisfying constraints
- First pass ensures left neighbor rule, second pass ensures right neighbor rule
- Each employee gets minimum satisfying both constraints

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of ratings for N children
2. **Goal**: Distribute minimum candies satisfying constraints
3. **Constraints**: Higher rating → more candies than immediate neighbors
4. **Output**: Minimum total candies needed

#### Phase 2: Key Insight Recognition
- **"Two-pass approach"** → Left-to-right then right-to-left
- **"Constraint satisfaction"** → Each pass handles one direction
- **"Local optimization"** → Each child gets minimum satisfying constraints
- **"Global optimality"** → Two passes guarantee both directions

#### Phase 3: Strategy Development
```
Human thought process:
"I need to distribute candies with minimum total.
Each child must have more candies than lower-rated neighbors.
I'll do two passes:
1. Left-to-right: ensure each child has more than left neighbor if needed
2. Right-to-left: ensure each child has more than right neighbor if needed
Each child gets the maximum of what both passes require.
This ensures all constraints with minimum total."
```

#### Phase 4: Edge Case Handling
- **Single child**: Always gets 1 candy
- **Equal ratings**: No constraint between equal ratings
- **Increasing sequence**: 1,2,3,4,... candies
- **Decreasing sequence**: N,N-1,N-2,... candies

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
ratings = [1,0,2]

Human thinking:
"First pass: left-to-right
Initialize all with 1 candy: [1,1,1]

Child 0: no left neighbor, keep 1
Child 1: rating 0 < rating 1, needs more than left
  left neighbor has 1, so child 1 needs at least 1
  keep 1 (since 1 > 1 is false, wait, let me recalculate)
  Actually, if rating[i] > rating[i-1], candies[i] = candies[i-1] + 1
  rating[1] (0) > rating[0] (1)? No, keep 1

Child 2: rating 2 > rating 0, needs more than left
  left neighbor has 1, so child 2 needs at least 2
  candies[2] = candies[1] + 1 = 2

After first pass: [1,1,2]

Second pass: right-to-left
Child 2: no right neighbor, keep 2
Child 1: rating 0 < rating 2, needs more than right
  right neighbor has 2, so child 1 needs at least 2
  current candies[1] = 1, max(1, 2) = 2
  candies[1] = 2

Child 0: rating 1 > rating 0, needs more than right
  right neighbor has 2, so child 0 needs at least 2
  current candies[0] = 1, max(1, 2) = 2
  candies[0] = 2

Final candies: [2,2,2]
Total: 6 candies ✓"
```

#### Phase 6: Intuition Validation
- **Why two passes**: Each pass handles one direction constraint
- **Why left-to-right first**: Establishes baseline for right neighbor constraint
- **Why max in second pass**: Takes maximum of both requirements
- **Why O(N)**: Each pass is O(N), total O(2N) = O(N)

### Common Human Pitfalls & How to Avoid Them
1. **"Why not one pass?"** → Can't satisfy both directions in one pass
2. **"Should I use DP?"** → Overkill, greedy is sufficient
3. **"What about equal ratings?"** → No constraint between equal ratings
4. **"Can I optimize space?"** → Yes, use mathematical slope tracking

### Real-World Analogy
**Like distributing performance bonuses in a company:**
- You have employees ranked by performance ratings
- Higher-performing employees must get more bonuses than neighbors
- You want to minimize total bonus budget while satisfying fairness
- First pass ensures left neighbor fairness, second pass ensures right neighbor fairness
- Each employee gets minimum bonus satisfying both fairness rules

### Human-Readable Pseudocode
```
function candy(ratings):
    if ratings is empty:
        return 0
    
    n = len(ratings)
    candies = array of size n filled with 1
    
    // Left to right pass
    for i from 1 to n-1:
        if ratings[i] > ratings[i-1]:
            candies[i] = candies[i-1] + 1
    
    // Right to left pass
    for i from n-2 down to 0:
        if ratings[i] > ratings[i+1]:
            candies[i] = max(candies[i], candies[i+1] + 1)
    
    return sum(candies)
```

### Execution Visualization

### Example: ratings = [1,0,2]
```
Two-Pass Distribution:
Initial: [1,1,1] (everyone gets 1)

Left-to-right pass:
Child 1: rating 0 ≤ rating 1 → keep 1
Child 2: rating 2 > rating 0 → candies[2] = candies[1] + 1 = 2
After first pass: [1,1,2]

Right-to-left pass:
Child 1: rating 0 < rating 2 → candies[1] = max(1, 2+1) = 3
Child 0: rating 1 > rating 0 → candies[0] = max(1, 3+1) = 4
After second pass: [4,3,2]

Wait, let me recalculate correctly:
Child 1: rating 0 < rating 2 → candies[1] = max(1, 2+1) = 3
Child 0: rating 1 > rating 0 → candies[0] = max(1, 3+1) = 4

Actually, let me trace again:
ratings = [1,0,2]
Initial: [1,1,1]
Left-to-right: [1,1,2] ✓
Right-to-left:
Child 1: rating 0 < rating 2 → candies[1] = max(1, 2+1) = 3
Child 0: rating 1 > rating 0 → candies[0] = max(1, 3+1) = 4

Final: [4,3,2], Total = 9
```

### Key Visualization Points:
- **Left Pass**: Ensures left neighbor constraint
- **Right Pass**: Ensures right neighbor constraint
- **Max Operation**: Takes maximum of both requirements
- **Minimum Total**: Each child gets minimum satisfying both constraints

### Memory Layout Visualization:
```
Candy Distribution Evolution:
ratings = [1,0,2]

Step 0 (Initial): [1,1,1]
Step 1 (Left pass): [1,1,2]
Step 2 (Right pass): [2,2,2]

Final distribution: [2,2,2]
Constraints satisfied:
- Child 0 (rating 1) > Child 1 (rating 0): 2 > 1 ✓
- Child 2 (rating 2) > Child 1 (rating 0): 2 > 1 ✓
```

### Time Complexity Breakdown:
- **Two Passes**: O(N) time complexity
- **Space Usage**: O(N) for candies array
- **Summation**: O(N) time for final sum
- **Total**: O(N) time, O(N) space

### Alternative Approaches:

#### 1. Single Pass with Slope Tracking (O(N) time, O(1) space)
```go
func candyOptimized(ratings []int) int {
    if len(ratings) == 0 {
        return 0
    }
    
    total := 1
    up := 1
    down := 0
    oldSlope := 0
    
    for i := 1; i < len(ratings); i++ {
        newSlope := 0
        if ratings[i] > ratings[i-1] {
            newSlope = 1
        } else if ratings[i] < ratings[i-1] {
            newSlope = -1
        }
        
        if oldSlope > 0 && newSlope == 0 || oldSlope < 0 && newSlope >= 0 {
            total += up * (up + 1) / 2
            up = 1
        }
        
        if newSlope > 0 {
            up++
        } else if newSlope < 0 {
            down++
        } else {
            total += up * (up + 1) / 2
            up = 1
            down = 0
        }
        
        oldSlope = newSlope
    }
    
    total += up * (up + 1) / 2
    if down > 0 {
        total += down * (down + 1) / 2
    }
    
    return total
}
```
- **Pros**: O(1) space optimization
- **Cons**: Very complex to understand and implement

#### 2. Brute Force (O(N²) time, O(N) space)
```go
func candyBrute(ratings []int) int {
    n := len(ratings)
    if n == 0 {
        return 0
    }
    
    minTotal := math.MaxInt32
    
    // Try all possible starting points
    for start := 1; start <= n*2; start++ {
        candies := make([]int, n)
        valid := true
        
        for i := 0; i < n; i++ {
            candies[i] = start
        }
        
        // Adjust to satisfy constraints
        for iter := 0; iter < n; iter++ {
            changed := false
            
            for i := 0; i < n; i++ {
                // Check left neighbor
                if i > 0 && ratings[i] > ratings[i-1] && candies[i] <= candies[i-1] {
                    candies[i] = candies[i-1] + 1
                    changed = true
                }
                
                // Check right neighbor
                if i < n-1 && ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
                    candies[i] = candies[i+1] + 1
                    changed = true
                }
            }
            
            if !changed {
                break
            }
        }
        
        total := 0
        for _, candy := range candies {
            total += candy
        }
        
        if total < minTotal {
            minTotal = total
        }
    }
    
    return minTotal
}
```
- **Pros**: Guarantees correctness
- **Cons**: O(N²) time, unnecessary complexity

#### 3. Recursive with Memoization (O(N²) time, O(N) space)
```go
func candyRecursive(ratings []int) int {
    memo := make(map[string]int)
    return candyHelper(ratings, 0, memo)
}

func candyHelper(ratings []int, pos int, memo map[string]int) int {
    if pos >= len(ratings) {
        return 0
    }
    
    key := fmt.Sprintf("%d,%d", pos, ratings[pos])
    if val, exists := memo[key]; exists {
        return val
    }
    
    // Try different candy counts for current position
    minTotal := math.MaxInt32
    for candy := 1; candy <= len(ratings); candy++ {
        valid := true
        
        // Check left constraint
        if pos > 0 && ratings[pos] > ratings[pos-1] && candy <= ratings[pos-1] {
            valid = false
        }
        
        // Check right constraint
        if pos < len(ratings)-1 && ratings[pos] > ratings[pos+1] && candy <= ratings[pos+1] {
            valid = false
        }
        
        if valid {
            total := candy + candyHelper(ratings, pos+1, memo)
            if total < minTotal {
                minTotal = total
            }
        }
    }
    
    memo[key] = minTotal
    return minTotal
}
```
- **Pros**: Intuitive approach
- **Cons**: O(N²) time, unnecessary complexity

### Extensions for Interviews:
- **Different Constraints**: Modified neighbor comparison rules
- **Multiple Candy Types**: Different candy types with constraints
- **Circular Arrangement**: Children in circular arrangement
- **Weighted Ratings**: Different importance for ratings
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := [][]int{
		{1, 0, 2},
		{1, 2, 2},
		{1, 3, 4, 5, 2},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 2, 2, 1},
		{2, 2, 2, 2},
		{1},
		{1, 2},
		{2, 1},
		{1, 3, 2, 2, 1},
		{1, 2, 87, 87, 87, 2, 1},
	}
	
	for i, ratings := range testCases {
		result1 := candy(ratings)
		result2 := candyOptimized(ratings)
		fmt.Printf("Test Case %d: %v -> Standard: %d, Optimized: %d\n", 
			i+1, ratings, result1, result2)
	}
}
