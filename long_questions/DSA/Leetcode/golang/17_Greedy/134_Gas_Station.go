package main

import "fmt"

// 134. Gas Station
// Time: O(N), Space: O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	currentGas := 0
	start := 0
	
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currentGas += gas[i] - cost[i]
		
		// If current gas is negative, we can't start from current start
		if currentGas < 0 {
			start = i + 1
			currentGas = 0
		}
	}
	
	// If total gas is less than total cost, no solution exists
	if totalGas < totalCost {
		return -1
	}
	
	return start
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Greedy with Cumulative Difference Tracking
- **Greedy Strategy**: Track cumulative gas-cost difference
- **Single Pass**: Process stations in order
- **Feasibility Check**: Total gas must be ≥ total cost
- **Start Point**: First station where cumulative difference becomes positive

## 2. PROBLEM CHARACTERISTICS
- **Circular Route**: Must complete full circuit and return to start
- **Resource Management**: Gas gained and cost spent at each station
- **Feasibility Condition**: Need enough total gas for total cost
- **Greedy Validity**: If solution exists, greedy finds optimal start

## 3. SIMILAR PROBLEMS
- Jump Game (LeetCode 55) - Reachability with greedy
- Candy (LeetCode 135) - Distribution with constraints
- Partition Labels (LeetCode 763) - String partitioning
- Can Complete Circuit (LeetCode 157) - Similar circular routing

## 4. KEY OBSERVATIONS
- **Cumulative Difference**: Track gas[i] - cost[i] cumulatively
- **Feasibility**: Total gas must be ≥ total cost
- **Start Point**: First position where cumulative becomes positive
- **Circular Property**: If solution exists, unique starting point
- **Greedy Optimality**: Starting at first positive position works

## 5. VARIATIONS & EXTENSIONS
- **Multiple Circuits**: Find all valid starting points
- **Minimum Gas**: Find minimum starting gas required
- **Capacity Constraints**: Add tank capacity limits
- **Station Costs**: Different cost models or constraints

## 6. INTERVIEW INSIGHTS
- Always clarify: "What if no solution exists? Multiple solutions?"
- Edge cases: single station, equal gas and cost, all gas sufficient
- Time complexity: O(N) time, O(1) space
- Key insight: track cumulative difference, not absolute values
- Greedy works because solution is unique if it exists

## 7. COMMON MISTAKES
- Not checking total gas vs total cost first
- Wrong cumulative difference tracking
- Not handling the circular nature correctly
- Returning wrong index when no solution exists
- Off-by-one errors in array indexing

## 8. OPTIMIZATION STRATEGIES
- **Single Pass**: O(N) time, O(1) space - optimal
- **Early Check**: Verify total gas ≥ total cost first
- **Cumulative Tracking**: Track running difference efficiently
- **Greedy Selection**: First positive cumulative difference is optimal

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like planning a road trip with gas stations:**
- You have a circular route with gas stations
- At each station, you gain gas but pay cost to travel to next
- You need to find a starting point where you never run out of gas
- Track your running balance as you travel
- The first point where balance becomes positive is your answer

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Arrays gas and cost of equal length
2. **Goal**: Find starting station to complete circular route
3. **Constraint**: Must return to starting point with non-negative gas
4. **Output**: Starting station index or -1 if impossible

#### Phase 2: Key Insight Recognition
- **"Circular nature"** → Must complete full circuit
- **"Feasibility first"** → Total gas must be ≥ total cost
- **"Cumulative difference"** → Track gas[i] - cost[i] running sum
- **"Greedy start point"** → First positive cumulative is optimal start

#### Phase 3: Strategy Development
```
Human thought process:
"I need to find a starting gas station.
First, I'll check if it's even possible (total gas ≥ total cost).
Then I'll track the running gas balance as I travel:
balance += gas[i] - cost[i]
If balance becomes negative at station i, I can't start from any previous station.
The first station where balance is positive is my answer!"
```

#### Phase 4: Edge Case Handling
- **No solution**: Return -1 if total gas < total cost
- **Single station**: If gas[0] ≥ cost[0], return 0
- **All sufficient**: Any starting point works, return 0
- **Equal totals**: If gas equals cost, any starting point works

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
gas = [1,2,3,4,5], cost = [3,4,5,1,2]

Human thinking:
"First check feasibility:
Total gas = 1+2+3+4+5 = 15
Total cost = 3+4+5+1+2 = 15
Feasible ✓

Now track cumulative balance:
Station 0: balance = 1-3 = -2 (can't start here)
Station 1: balance = -2 + 2-4 = -4 (can't start here)
Station 2: balance = -4 + 3-5 = -6 (can't start here)
Station 3: balance = -6 + 4-1 = -3 (can't start here)
Station 4: balance = -3 + 5-2 = 0 (can't start here)

Wait, let me recalculate:
Station 0: balance = 1-3 = -2
Station 1: balance = -2 + 2-4 = -4
Station 2: balance = -4 + 3-5 = -6
Station 3: balance = -6 + 4-1 = -3
Station 4: balance = -3 + 5-2 = 0

All balances are ≤ 0, so no solution exists.
Let me try another example:
gas = [5,1,2,3,4], cost = [4,4,1,5,2]

Total gas = 15, Total cost = 16
Not feasible, return -1 ✓"

Let me try a feasible example:
gas = [1,2,3,4,5], cost = [3,4,5,1,2]

Total gas = 15, Total cost = 15
Feasible ✓

Station 0: balance = 1-3 = -2
Station 1: balance = -2 + 2-4 = -4
Station 2: balance = -4 + 3-5 = -6
Station 3: balance = -6 + 4-1 = -3
Station 4: balance = -3 + 5-2 = 0

All balances ≤ 0, return 0 (can start anywhere) ✓"
```

#### Phase 6: Intuition Validation
- **Why greedy works**: If solution exists, it's unique and greedy finds it
- **Why cumulative tracking**: Running balance shows feasibility of each start
- **Why first positive**: First station where you can have positive balance
- **Why O(N)**: Single pass through arrays

### Common Human Pitfalls & How to Avoid Them
1. **"Why not try all starts?"** → O(N²), greedy is O(N)
2. **"Should I use DP?"** → Overkill, greedy is sufficient
3. **"What about multiple solutions?"** → Greedy finds one, discuss if multiple exist
4. **"Can I optimize further?"** → Greedy is already optimal

### Real-World Analogy
**Like planning a delivery route with fuel stations:**
- You have a circular delivery route with fuel stations
- At each station, you get fuel but pay to travel to next
- You need to find a starting depot where you never run out of fuel
- Track your running fuel balance as you make deliveries
- The first station where balance becomes positive is your optimal depot

### Human-Readable Pseudocode
```
function canCompleteCircuit(gas, cost):
    totalGas = sum(gas)
    totalCost = sum(cost)
    
    if totalGas < totalCost:
        return -1  // Not enough gas overall
    
    balance = 0
    start = 0
    
    for i from 0 to len(gas)-1:
        balance += gas[i] - cost[i]
        
        if balance < 0:
            start = i + 1
            balance = 0
    
    return start
```

### Execution Visualization

### Example: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
```
Circuit Analysis:
Total gas = 15, Total cost = 15 ✓ Feasible

Station-by-station analysis:
Station 0: balance = 1-3 = -2
Station 1: balance = -2 + 2-4 = -4  
Station 2: balance = -4 + 3-5 = -6
Station 3: balance = -6 + 4-1 = -3
Station 4: balance = -3 + 5-2 = 0

All balances ≤ 0, can start anywhere ✓
Return 0 (or any valid starting point)
```

### Key Visualization Points:
- **Feasibility Check**: Total gas must be ≥ total cost
- **Cumulative Balance**: Track running gas - cost difference
- **Start Point**: First station where balance becomes positive
- **Circular Property**: Solution wraps around to complete circuit

### Memory Layout Visualization:
```
Balance Tracking Visualization:
gas = [1,2,3,4,5], cost = [3,4,5,1,2]

Step 0: balance = 0
After station 0: balance = 0 + (1-3) = -2
After station 1: balance = -2 + (2-4) = -4
After station 2: balance = -4 + (3-5) = -6
After station 3: balance = -6 + (4-1) = -3
After station 4: balance = -3 + (5-2) = 0

Negative balances mean can't start from those stations.
```

### Time Complexity Breakdown:
- **Feasibility Check**: O(N) time for sum calculation
- **Single Pass**: O(N) time for balance tracking
- **Total Time**: O(N) time complexity
- **Space Complexity**: O(1) additional space

### Alternative Approaches:

#### 1. Brute Force (O(N²) time, O(1) space)
```go
func canCompleteCircuitBrute(gas, cost []int) int {
    n := len(gas)
    for start := 0; start < n; start++ {
        tank := 0
        feasible := true
        
        for i := 0; i < n; i++ {
            station := (start + i) % n
            tank += gas[station] - cost[station]
            if tank < 0 {
                feasible = false
                break
            }
        }
        
        if feasible {
            return start
        }
    }
    
    return -1
}
```
- **Pros**: Guarantees correctness
- **Cons**: O(N²) time, unnecessary complexity

#### 2. Prefix Sum with Early Reset (O(N) time, O(1) space)
```go
func canCompleteCircuitPrefix(gas, cost []int) int {
    n := len(gas)
    if sum(gas) < sum(cost) {
        return -1
    }
    
    balance := 0
    start := 0
    
    for i := 0; i < n; i++ {
        balance += gas[i] - cost[i]
        
        if balance < 0 {
            start = i + 1
            balance = 0
        }
    }
    
    return start
}
```
- **Pros**: Optimal O(N) time
- **Cons**: Same as greedy approach

#### 3. Two-Pass Approach (O(N) time, O(1) space)
```go
func canCompleteCircuitTwoPass(gas, cost []int) int {
    n := len(gas)
    if sum(gas) < sum(cost) {
        return -1
    }
    
    // First pass: find minimum balance point
    minBalance := 0
    balance := 0
    minIndex := 0
    
    for i := 0; i < n; i++ {
        balance += gas[i] - cost[i]
        if balance < minBalance {
            minBalance = balance
            minIndex = i + 1
        }
    }
    
    // Second pass: find first index with balance >= minBalance
    balance = 0
    for i := 0; i < n; i++ {
        balance += gas[i] - cost[i]
        if balance >= minBalance && i >= minIndex {
            return i
        }
    }
    
    return -1
}
```
- **Pros**: Alternative perspective on same algorithm
- **Cons**: More complex than single pass

### Extensions for Interviews:
- **Multiple Solutions**: Find all valid starting points
- **Capacity Constraints**: Add tank capacity limits
- **Station Costs**: Different cost models or constraints
- **Route Optimization**: Find optimal route visiting all stations
- **Performance Analysis**: Discuss worst-case scenarios
*/
func main() {
	// Test cases
	testCases := []struct {
		gas  []int
		cost []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}},
		{[]int{2, 3, 4}, []int{3, 4, 3}},
		{[]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}},
		{[]int{3, 3, 4}, []int{3, 4, 4}},
		{[]int{1, 2, 3, 4, 5, 5, 70}, []int{2, 3, 4, 3, 4, 5, 50}},
		{[]int{2}, []int{2}},
		{[]int{1}, []int{2}},
		{[]int{5, 8, 2, 8}, []int{6, 5, 6, 6}},
		{[]int{4, 5, 3, 1, 4}, []int{5, 4, 3, 4, 2}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
	}
	
	for i, tc := range testCases {
		result := canCompleteCircuit(tc.gas, tc.cost)
		fmt.Printf("Test Case %d: gas=%v, cost=%v -> Start: %d\n", 
			i+1, tc.gas, tc.cost, result)
	}
}
