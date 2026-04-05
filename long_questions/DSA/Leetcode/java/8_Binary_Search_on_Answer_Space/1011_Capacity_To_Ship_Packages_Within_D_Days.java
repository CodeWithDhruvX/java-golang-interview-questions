public class CapacityToShipPackagesWithinDDays {
    
    // 1011. Capacity To Ship Packages Within D Days
    // Time: O(N log M), Space: O(1) where M is sum(weights)
    public int shipWithinDays(int[] weights, int days) {
        int left = maxWeight(weights);
        int right = sumWeights(weights);
        int result = right;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (canShip(weights, days, mid)) {
                result = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        
        return result;
    }
    
    private boolean canShip(int[] weights, int days, int capacity) {
        int daysNeeded = 1;
        int currentLoad = 0;
        
        for (int weight : weights) {
            if (currentLoad + weight <= capacity) {
                currentLoad += weight;
            } else {
                daysNeeded++;
                currentLoad = weight;
                if (daysNeeded > days) {
                    return false;
                }
            }
        }
        
        return daysNeeded <= days;
    }
    
    private int maxWeight(int[] weights) {
        int max = 0;
        for (int weight : weights) {
            if (weight > max) {
                max = weight;
            }
        }
        return max;
    }
    
    private int sumWeights(int[] weights) {
        int sum = 0;
        for (int weight : weights) {
            sum += weight;
        }
        return sum;
    }
    
    // Alternative approach with detailed explanation
    public int shipWithinDaysDetailed(int[] weights, int days) {
        System.out.println("=== Binary Search on Answer Space ===");
        System.out.printf("Weights: %s, Days: %d\n", java.util.Arrays.toString(weights), days);
        
        int left = maxWeight(weights);
        int right = sumWeights(weights);
        int result = right;
        
        System.out.printf("Search space: [%d, %d]\n", left, right);
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            System.out.printf("Trying capacity: %d\n", mid);
            
            if (canShipDetailed(weights, days, mid)) {
                result = mid;
                right = mid - 1;
                System.out.printf("  ✓ Can ship in %d days, searching lower\n", days);
            } else {
                left = mid + 1;
                System.out.printf("  ✗ Cannot ship, need higher capacity\n");
            }
        }
        
        System.out.printf("Final result: %d\n", result);
        return result;
    }
    
    private boolean canShipDetailed(int[] weights, int days, int capacity) {
        int daysNeeded = 1;
        int currentLoad = 0;
        StringBuilder dayLoads = new StringBuilder();
        
        for (int i = 0; i < weights.length; i++) {
            int weight = weights[i];
            
            if (currentLoad + weight <= capacity) {
                currentLoad += weight;
                dayLoads.append(weight).append(" ");
            } else {
                System.out.printf("    Day %d: %s (load: %d)\n", daysNeeded, dayLoads.toString().trim(), currentLoad);
                daysNeeded++;
                currentLoad = weight;
                dayLoads = new StringBuilder();
                dayLoads.append(weight).append(" ");
                
                if (daysNeeded > days) {
                    return false;
                }
            }
        }
        
        System.out.printf("    Day %d: %s (load: %d)\n", daysNeeded, dayLoads.toString().trim(), currentLoad);
        return daysNeeded <= days;
    }
    
    // Brute force approach for comparison
    public int shipWithinDaysBruteForce(int[] weights, int days) {
        int minCapacity = maxWeight(weights);
        int maxCapacity = sumWeights(weights);
        
        for (int capacity = minCapacity; capacity <= maxCapacity; capacity++) {
            if (canShip(weights, days, capacity)) {
                return capacity;
            }
        }
        
        return maxCapacity;
    }
    
    // Optimized version with early termination
    public int shipWithinDaysOptimized(int[] weights, int days) {
        int left = maxWeight(weights);
        int right = sumWeights(weights);
        int result = right;
        
        // Early optimization: if days == weights.length, answer is maxWeight
        if (days == weights.length) {
            return left;
        }
        
        // Early optimization: if days == 1, answer is sumWeights
        if (days == 1) {
            return right;
        }
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (canShipOptimized(weights, days, mid)) {
                result = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        
        return result;
    }
    
    private boolean canShipOptimized(int[] weights, int days, int capacity) {
        int daysNeeded = 1;
        int currentLoad = 0;
        
        for (int weight : weights) {
            if (weight > capacity) {
                return false; // Early termination
            }
            
            if (currentLoad + weight <= capacity) {
                currentLoad += weight;
            } else {
                daysNeeded++;
                currentLoad = weight;
                
                if (daysNeeded > days) {
                    return false;
                }
            }
        }
        
        return true;
    }
    
    // Version that returns the actual shipping arrangement
    public class ShippingPlan {
        int minCapacity;
        java.util.List<java.util.List<Integer>> dailyLoads;
        
        ShippingPlan(int minCapacity, java.util.List<java.util.List<Integer>> dailyLoads) {
            this.minCapacity = minCapacity;
            this.dailyLoads = dailyLoads;
        }
    }
    
    public ShippingPlan getShippingPlan(int[] weights, int days) {
        int minCapacity = shipWithinDays(weights, days);
        java.util.List<java.util.List<Integer>> dailyLoads = new java.util.ArrayList<>();
        
        int currentLoad = 0;
        java.util.List<Integer> currentDay = new java.util.ArrayList<>();
        
        for (int weight : weights) {
            if (currentLoad + weight <= minCapacity) {
                currentLoad += weight;
                currentDay.add(weight);
            } else {
                dailyLoads.add(new java.util.ArrayList<>(currentDay));
                currentDay.clear();
                currentDay.add(weight);
                currentLoad = weight;
            }
        }
        
        if (!currentDay.isEmpty()) {
            dailyLoads.add(currentDay);
        }
        
        return new ShippingPlan(minCapacity, dailyLoads);
    }
    
    public static void main(String[] args) {
        CapacityToShipPackagesWithinDDays solver = new CapacityToShipPackagesWithinDDays();
        
        // Test cases
        int[][][] testCases = {
            {{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {5}},
            {{3, 2, 2, 4, 1, 4}, {3}},
            {{1, 2, 3, 1, 1}, {4}},
            {{10, 50, 100, 100, 50, 10}, {5}},
            {{1, 2, 3, 4, 5}, {5}},
            {{1, 2, 3, 4, 5}, {1}},
            {{100}, {1}},
            {{1, 1, 1, 1, 1, 1, 1, 1, 1}, {9}},
            {{5, 5, 5, 5, 5}, {3}}
        };
        
        String[] descriptions = {
            "Standard case",
            "Mixed weights",
            "Small weights",
            "Large weights",
            "Equal weights and days",
            "Single day",
            "Single package",
            "Many small packages",
            "Uniform weights"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            int[] weights = testCases[i][0];
            int days = testCases[i][1][0];
            
            int result1 = solver.shipWithinDays(weights, days);
            int result2 = solver.shipWithinDaysBruteForce(weights, days);
            int result3 = solver.shipWithinDaysOptimized(weights, days);
            
            System.out.printf("  Binary Search: %d\n", result1);
            System.out.printf("  Brute Force: %d\n", result2);
            System.out.printf("  Optimized: %d\n", result3);
            
            // Get shipping plan
            ShippingPlan plan = solver.getShippingPlan(weights, days);
            System.out.printf("  Shipping Plan (capacity %d):\n", plan.minCapacity);
            for (int j = 0; j < plan.dailyLoads.size(); j++) {
                System.out.printf("    Day %d: %s\n", j + 1, plan.dailyLoads.get(j));
            }
            System.out.println();
        }
        
        // Detailed explanation for one case
        System.out.println("=== Detailed Explanation ===");
        int[] detailedWeights = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
        int detailedDays = 5;
        solver.shipWithinDaysDetailed(detailedWeights, detailedDays);
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        
        // Large test case
        int[] largeWeights = new int[1000];
        for (int i = 0; i < 1000; i++) {
            largeWeights[i] = (i % 100) + 1;
        }
        
        long startTime = System.nanoTime();
        int largeResult = solver.shipWithinDays(largeWeights, 50);
        long endTime = System.nanoTime();
        
        System.out.printf("Large test (1000 weights, 50 days): %d (took %d ns)\n", 
            largeResult, endTime - startTime);
        
        startTime = System.nanoTime();
        int bruteResult = solver.shipWithinDaysBruteForce(largeWeights, 50);
        endTime = System.nanoTime();
        
        System.out.printf("Brute force result: %d (took %d ns)\n", 
            bruteResult, endTime - startTime);
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Binary Search on Answer Space
- **Binary Search**: Search for minimum feasible capacity
- **Feasibility Check**: Can all packages be shipped within D days
- **Search Space**: [1, maxWeight] to [sum(weights)]
- **Monotonic**: If capacity works, all smaller capacities work

## 2. PROBLEM CHARACTERISTICS
- **Package Weights**: Array of package weights
- **Shipping Constraint**: D days to ship all packages
- **Capacity Question**: Find minimum capacity to meet constraint
- **Feasibility**: Check if capacity can ship within time limit

## 3. SIMILAR PROBLEMS
- Koko Eating Bananas
- Split Array Largest Sum
- Minimum Number of Days to Make M Bouquets
- Allocate Minimum Number of Books

## 4. KEY OBSERVATIONS
- Binary search works because feasibility is monotonic
- If capacity C works, any capacity > C also works
- Search space: [1, maxWeight] to [sum(weights)]
- Feasibility check requires simulating shipping process

## 5. VARIATIONS & EXTENSIONS
- Return actual shipping arrangement
- Multiple queries on same weights
- Different constraints (e.g., at most K packages per day)
- Real-time streaming version

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I return the shipping arrangement?"
- Edge cases: empty weights, single package, D=1
- Time complexity: O(N log M) vs O(NM) brute force
- Space complexity: O(1) vs O(N) for brute force

## 7. COMMON MISTAKES
- Using linear search instead of binary search
- Incorrect feasibility check implementation
- Off-by-one errors in binary search bounds
- Not handling edge cases properly

## 8. OPTIMIZATION STRATEGIES
- Binary search is optimal for single query
- Early termination for obvious cases
- Mathematical bounds for search space
- Pre-sorting if weights are unsorted

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the right truck size:**
- You have packages with different weights (weights)
- You want to ship them within D days (deadline)
- You need to find the smallest truck that can do this job
- If a truck size works, any larger truck also works

- This monotonic property allows binary search

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of weights and integer D
2. **Goal**: Find minimum capacity to ship all packages within D days
3. **Output**: Minimum feasible capacity

#### Phase 2: Key Insight Recognition
- **"Can I test a capacity?"** → Simulate shipping process
- **"Is the relationship monotonic?"** → Yes! If capacity C works, C+1 works
- **"What's the search range?"** → [1, maxWeight] to [sum(weights)]
- **"Can I use binary search?"** → Yes, monotonic feasibility

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use binary search to find the minimum capacity:
1. Define search space: [1, maxWeight] to [sum(weights)]
2. For each capacity mid:
   - Check if I can ship all packages within D days
   - If yes, try smaller capacity
   - If no, try larger capacity
3. The smallest 'yes' capacity is my answer"
```

#### Phase 4: Edge Case Handling
- **Empty weights**: Return 0 (no capacity needed)
- **D >= weights.length**: Return maxWeight (ship each separately)
- **Single package**: Return its weight
- **Large D**: Binary search helps avoid O(ND) brute force

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: weights = [1,2,3,4,5,6,7,8,9,10], D = 5

Human thinking:
"Search space: [1, 10] to [55]

Step 1: mid = (1+10)/2 = 5
Can I ship with capacity 5?
- Day 1: [1,2,3,4,5] (weight 15) ✓
- Day 2: [6,7,8,9,10] (weight 40) ✓
- Day 3: [11] (weight 11) ✓
- Day 4: [12] (weight 12) ✓
- Day 5: [13] (weight 13) ✓
- Yes! Try smaller capacity

Step 2: mid = (1+4)/2 = 2
Can I ship with capacity 2?
- Day 1: [1,2] (weight 3) ✓
- Day 2: [3,4] (weight 7) ✓
- Day 3: [5,6] (weight 11) ✓
- Day 4: [7,8] (weight 15) ✓
- Day 5: [9,10] (weight 19) > 5 days ✗
- No! Try larger capacity

Step 3: mid = (2+4)/2 = 3
Can I ship with capacity 3?
- Yes! Answer is 3"
```

#### Phase 6: Intuition Validation
- **Why it works**: Monotonic feasibility enables binary search
- **Why it's efficient**: O(N log M) vs O(ND) brute force
- **Why it's correct**: Binary search finds minimal feasible capacity

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just try all capacities?"** → Too slow O(ND)
2. **"What's wrong with linear search?"** → Misses logarithmic improvement
3. **"How to check feasibility?"** → Must simulate actual shipping process
4. **"What about the bounds?"** → Lower bound is 1, upper is sum(weights)

### Real-World Analogy
**Like finding the right delivery truck:**
- You have packages of different sizes (weights)
- You need to deliver all within D days (deadline)
- You test different truck capacities
- If a small truck works, a bigger truck will also work
- You want the smallest truck that meets the deadline
- This is exactly the binary search problem!

### Human-Readable Pseudocode
```
function shipWithinDays(weights, D):
    left = 1
    right = sum(weights)
    result = right
    
    while left <= right:
        mid = left + (right - left) / 2
        
        if canShip(weights, D, mid):
            result = mid
            right = mid - 1
        else:
            left = mid + 1
    
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
    
    return true
```

### Execution Visualization

### Example: weights = [1,2,3,4,5,6,7,8,9,10], D = 5
```
Initial: weights = [1,2,3,4,5,6,7,8,9,10], D = 5
Search space: [1, 10] to [55]

Step 1: mid = 5
Testing capacity 5:
Day 1: [1,2,3,4,5] = 15 ≤ 5 ✓
Day 2: [6,7,8,9,10] = 40 > 5 ✗
Day 3: [11] = 11 > 5 ✗
Day 4: [12] = 12 > 5 ✗
Day 5: [13] = 13 > 5 ✗
Result: Can ship, search lower

Step 2: mid = 2
Testing capacity 2:
Day 1: [1,2] = 3 ≤ 2 ✓
Day 2: [3,4] = 7 > 2 ✗
Day 3: [5,6] = 11 > 2 ✗
Day 4: [7,8] = 15 > 2 ✗
Day 5: [9,10] = 19 > 2 ✗
Result: Cannot ship, search higher

Step 3: mid = 3
Testing capacity 3:
Day 1: [1,2,3] = 6 ≤ 3 ✓
Day 2: [4,5] = 9 > 3 ✗
Day 3: [6,7] = 13 > 3 ✗
Day 4: [8,9] = 17 > 3 ✗
Day 5: [10] = 10 > 3 ✗
Result: Can ship, answer = 3 ✓
```

### Key Visualization Points:
- **Binary search** on capacity values
- **Feasibility check** simulates actual shipping
- **Monotonic property**: If capacity works, larger works
- **Search space reduction**: Narrow down to minimum feasible capacity

### Memory Layout Visualization:
```
Weights: [1][2][3][4][5][6][7][8][9][10]
Capacity 5: [1,2,3]|[4,5]|[6]|[7]|[8]|[9]|[10]
           Day 1  Day 2  Day 3  Day 4  Day 5
           ✓     ✗     ✗     ✗     ✗     ✗

Capacity 3: [1,2]|[3,4]|[5,6]|[7]|[8]|[9]|[10]
           Day 1  Day 2  Day 3  Day 4  Day 5
           ✓     ✗     ✗     ✗     ✗     ✗

Final Answer: 3 ✓
```

### Time Complexity Breakdown:
- **Binary Search**: O(log M) iterations where M = sum(weights)
- **Feasibility Check**: O(N) per iteration
- **Total**: O(N log M) time, O(1) space
- **Optimal**: Best possible for single query
- **vs Brute Force**: O(ND) where D is max weight
*/
