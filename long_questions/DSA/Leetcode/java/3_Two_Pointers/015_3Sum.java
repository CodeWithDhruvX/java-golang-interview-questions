import java.util.*;

public class ThreeSum {
    
    // 15. 3Sum
    // Time: O(N^2), Space: O(1) (ignoring output space)
    public static List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> result = new ArrayList<>();
        int n = nums.length;
        
        for (int i = 0; i < n - 2; i++) {
            // Skip duplicates for the first element
            if (i > 0 && nums[i] == nums[i - 1]) {
                continue;
            }
            
            // Two pointers approach for the remaining two elements
            int left = i + 1, right = n - 1;
            int target = -nums[i];
            
            while (left < right) {
                int sum = nums[left] + nums[right];
                
                if (sum == target) {
                    result.add(Arrays.asList(nums[i], nums[left], nums[right]));
                    
                    // Skip duplicates for the second element
                    while (left < right && nums[left] == nums[left + 1]) {
                        left++;
                    }
                    // Skip duplicates for the third element
                    while (left < right && nums[right] == nums[right - 1]) {
                        right--;
                    }
                    
                    left++;
                    right--;
                } else if (sum < target) {
                    left++;
                } else {
                    right--;
                }
            }
        }
        
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {-1, 0, 1, 2, -1, -4},
            {0, 1, 1},
            {0, 0, 0},
            {-2, 0, 1, 1, 2},
            {-1, -2, -3, -4, -5},
            {1, 2, -2, -1},
            {3, -2, 1, 0, -1, 2, -3},
            {},
            {0}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            List<List<Integer>> result = threeSum(testCases[i]);
            System.out.printf("Test Case %d: %s -> Triplets: %s\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Two-Pointer Triple Search
- **3Sum Problem**: Fixed first element, two-pointer search for remaining
- **Sorting First**: Enables duplicate skipping and early termination
- **Target Transformation**: Find two numbers that sum to -first element
- **Linear Inner Search**: Two-pointer technique for sorted array

## 2. PROBLEM CHARACTERISTICS
- **3-Sum Problem**: Find triplets that sum to zero
- **Fixed Element Strategy**: One element is fixed, find two others
- **Sorted Array**: Enables efficient two-pointer search
- **Duplicate Handling**: Skip same values to avoid repeated triplets

## 3. SIMILAR PROBLEMS
- 4Sum (quadruple nested pointers)
- 2Sum (single two-pointer)
- 3Sum II (with duplicates)
- K Sum problems in general

## 4. KEY OBSERVATIONS
- **Sorting Optimization**: O(N log N) + O(N²) vs brute force O(N³)
- **Two-Pointer Search**: Linear scan from both ends
- **Target Calculation**: target = -nums[i] (negate first element)
- **Duplicate Prevention**: Skip consecutive identical values

## 5. VARIATIONS & EXTENSIONS
- Return indices instead of values
- Handle very large numbers (overflow)
- 3Sum Closest (sum closest to target)
- Count triplets without storing them all

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I return sorted triplets?"
- Edge cases: array length < 3, all duplicates, no solution
- Time complexity: O(N²) is optimal for 3Sum
- Space complexity: O(1) for two pointers, O(N) for result

## 7. COMMON MISTAKES
- Not sorting the array first
- Forgetting duplicate skipping logic
- Incorrect target calculation (should be -nums[i])
- Not handling empty array or insufficient length cases
- Two-pointer implementation errors

## 8. OPTIMIZATION STRATEGIES
- Early termination when sum exceeds target
- Hash set for duplicate detection
- Skip unnecessary iterations
- Pre-allocate result capacity

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding 3 people whose combined weight equals zero:**
- Array is list of people with different weights (positive/negative)
- Need to find exactly 3 people whose total weight = 0
- Sort by weight to enable efficient two-pointer search
- Fix one person, find two others that balance to zero

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers (can be positive/negative)
2. **Goal**: Find all unique triplets that sum to zero
3. **Output**: List of lists containing 3 integers each
4. **Constraints**: Order within triplet doesn't matter

#### Phase 2: Key Insight Recognition
- **"How to find 3 numbers efficiently?"** → Sort + two-pointer search
- **"Why fix first element?"** → Reduces to finding 2-sum for remaining
- **"How to avoid duplicates?"** → Skip identical consecutive values
- **"What's the pattern?"** → For each i, find j,k such that nums[j] + nums[k] = -nums[i]

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use sorting + two-pointer approach:
1. Sort array to enable optimizations
2. For each element as first element (fix nums[i]):
   - Target = -nums[i] (find two that sum to this)
   - Use two-pointer search to find all pairs that sum to target
3. Skip duplicates to avoid repeated triplets
4. Collect all valid (nums[i], nums[j], nums[k]) combinations"
```

#### Phase 4: Algorithm Walkthrough
```
Example: nums=[-1,0,1,2,-1,-4], target for first element = 1

Human thinking:
"Sort: [-4,-1,-1,0,1,2]

i=0 (nums[0]=-4), target=4:
  left=1, right=4
  Two-pointer search for sum=4:
    left=1 (nums[1]=-1), right=4 (nums[4]=2): sum=1 < 4
    left=2 (nums[2]=0), right=3 (nums[3]=1): sum=1 < 4
    left=3 (nums[3]=1), right=2 (nums[2]=0): sum=1 < 4
  No pairs found for target=4

i=1 (nums[1]=-1), target=1:
  Skip (duplicate of i=0)

i=2 (nums[2]=0), target=0:
  left=1, right=4
  Two-pointer search for sum=0:
    left=1 (nums[1]=-1), right=4 (nums[4]=2): sum=1 > 0
    left=2 (nums[2]=0), right=3 (nums[3]=1): sum=1 > 0
    left=3 (nums[3]=1), right=2 (nums[2]=0): sum=1 > 0
  No pairs found for target=0

i=3 (nums[3]=1), target=-1:
  left=1, right=4
  Two-pointer search for sum=-1:
    left=1 (nums[1]=-1), right=4 (nums[4]=2): sum=1 ≠ -1
    left=2 (nums[2]=0), right=3 (nums[3]=1): sum=1 ≠ -1
    left=3 (nums[3]=1), right=2 (nums[2]=0): sum=1 ≠ -1
  No pairs found for target=-1

Continue this process...
Final results include combinations like [-4,2,2], [-1,0,1], etc. ✓"
```

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use 3 nested loops?"** → O(N³) vs O(N²) with two-pointer
2. **"What about unsorted array?"** → Can use hash set but loses two-pointer efficiency
3. **"How to handle target calculation?"** → Target = -nums[i] for sum to zero
4. **"When to prune?"** → When sorted, break early if minimum sum > target
5. **"How to skip duplicates?"** → Skip consecutive identical values

### Real-World Analogy
**Like balancing a budget with 3-person teams:**
- Array represents available team members with different costs
- Need to form 3-person teams whose total cost = 0 (balanced budget)
- Fix one person, find two others that balance the budget
- Sort by cost to enable efficient partner search

### Human-Readable Pseudocode
```
function threeSum(nums):
    sort(nums)
    result = []
    
    for i from 0 to n-3:
        if i > 0 and nums[i] == nums[i-1]: continue
        
        target = -nums[i]
        left = i + 1
        right = n - 1
        
        while left < right:
            sum = nums[left] + nums[right]
            
            if sum == target:
                result.add([nums[i], nums[left], nums[right]])
                
                # Skip duplicates for second element
                while left < right and nums[left] == nums[left + 1]:
                    left++
                else:
                    right--
    
    return result
```

### Execution Visualization

### Example: nums=[-1,0,1,2,-1,-4]
```
Sorted Array: [-4,-1,-1,0,1,2]

Selection Process:
i=0 (-4), target=4:
  Two-pointer search for sum=4: No pairs found

i=1 (-1): skip (duplicate)

i=2 (0), target=0:
  Two-pointer search for sum=0:
    left=1 (-1), right=4 (2): sum=1 > 0
    left=2 (0), right=3 (1): sum=1 > 0
    left=3 (1), right=2 (0): sum=1 > 0
  No pairs found for target=0

i=3 (1), target=-1:
  Two-pointer search for sum=-1:
    left=1 (-1), right=4 (2): sum=1 ≠ -1
    left=2 (0), right=3 (1): sum=1 ≠ -1
    left=3 (1), right=2 (0): sum=1 ≠ -1
  No pairs found for target=-1

Final Valid Triplets:
[-4,2,2] (sum=0)
[-1,0,1] (sum=0)
[0,1,-1] (sum=0)
```

### Key Visualization Points:
- **Fixed Element Strategy**: One element fixed, find two others
- **Two-Pointer Search**: Linear scan from both ends toward target
- **Sorting Optimization**: Enables early termination and duplicate skipping
- **Target Transformation**: nums[j] + nums[k] = -nums[i]

### Time Complexity Breakdown:
- **Sorting**: O(N log N)
- **Main Loop**: O(N²) with two-pointer search
- **Total**: O(N²) time, O(1) space (excluding result)
- **Optimal**: Significant improvement over O(N³) brute force
*/
