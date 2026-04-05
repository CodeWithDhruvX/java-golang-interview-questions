import java.util.*;

public class FourSum {
    
    // 18. 4Sum
    // Time: O(N³), Space: O(N) for result (excluding output)
    public static List<List<Integer>> fourSum(int[] nums, int target) {
        List<List<Integer>> result = new ArrayList<>();
        if (nums == null || nums.length < 4) {
            return result;
        }
        
        Arrays.sort(nums);
        int n = nums.length;
        
        for (int i = 0; i < n - 3; i++) {
            // Skip duplicates for first element
            if (i > 0 && nums[i] == nums[i - 1]) {
                continue;
            }
            
            // Early termination optimization
            if (nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3] > target) {
                break;
            }
            
            for (int j = i + 1; j < n - 2; j++) {
                // Skip duplicates for second element
                if (j > i + 1 && nums[j] == nums[j - 1]) {
                    continue;
                }
                
                int left = j + 1;
                int right = n - 1;
                int newTarget = target - nums[i] - nums[j];
                
                while (left < right) {
                    int sum = nums[left] + nums[right];
                    
                    if (sum == newTarget) {
                        result.add(Arrays.asList(nums[i], nums[j], nums[left], nums[right]));
                        
                        // Skip duplicates for third element
                        while (left < right && nums[left] == nums[left + 1]) {
                            left++;
                        }
                        
                        // Skip duplicates for fourth element
                        while (left < right && nums[right] == nums[right - 1]) {
                            right--;
                        }
                        
                        left++;
                        right--;
                    } else if (sum < newTarget) {
                        left++;
                    } else {
                        right--;
                    }
                }
            }
        }
        
        return result;
    }

    public static void main(String[] args) {
        Object[][] testCases = {
            {new int[]{1, 0, -1, 0, -2, 2}, 0},
            {new int[]{2, 2, 2, 2, 2}, 8},
            {new int[]{1, 1, 1, 1, 1}, 4},
            {new int[]{-3, -2, -1, 0, 0, 1, 2, 3}, 0},
            {new int[]{0, 0, 0, 0}, 0},
            {new int[]{1, 2, 3, 4, 5}, 10},
            {new int[]{-1, -2, -3, -4, -5}, -10},
            {new int[]{1, 0, -1, 0, -2, 2, -1, 1}, 0},
            {new int[]{1000000000, 1000000000, 1000000000, 1000000000}, 0},
            {new int[]{0, 1, -1, 2, -2, 3, -3}, 0}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] nums = (int[]) testCases[i][0];
            int target = (int) testCases[i][1];
            List<List<Integer>> result = fourSum(nums, target);
            
            System.out.printf("Test Case %d: %s, target=%d -> %s\n", 
                i + 1, Arrays.toString(nums), target, result);
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Multi-Pointer Combination
- **4Sum Problem**: Nested two-pointer technique
- **Sorting First**: Enables early termination and duplicate skipping
- **Two Sum Reduction**: Inner problem reduces to known 2Sum
- **Quadratic Complexity**: Optimal for this problem class

## 2. PROBLEM CHARACTERISTICS
- **K-Sum Problem**: Find K numbers that sum to target
- **Combination Search**: Choose K elements from array
- **Duplicate Handling**: Skip same values to avoid duplicate results
- **Early Termination**: Stop when sum exceeds target (sorted array)

## 3. SIMILAR PROBLEMS
- 3Sum (triple nested pointers)
- 2Sum (single two-pointer)
- 4Sum II (with duplicates allowed)
- K Sum problems in general

## 4. KEY OBSERVATIONS
- **Sorted Array**: Enables early termination optimization
- **Nested Loops**: i < j < k < l pattern
- **Duplicate Skipping**: Essential to avoid duplicate quadruplets
- **Target Achievement**: nums[i] + nums[j] + nums[k] + nums[l] = target

## 5. VARIATIONS & EXTENSIONS
- Return indices instead of values
- Handle very large numbers (overflow)
- Different K values (3Sum, 5Sum)
- Unsorted input with hash map approach

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I return indices or values?"
- Edge cases: array length < K, all duplicates, no solution
- Time complexity: O(N³) is optimal for 4Sum
- Space complexity: O(N) for result storage

## 7. COMMON MISTAKES
- Not sorting the array first
- Forgetting duplicate skipping logic
- Incorrect loop boundaries (n-3, n-2, n-1)
- Not handling empty array or insufficient length cases

## 8. OPTIMIZATION STRATEGIES
- Early termination when sum > target
- Hash set for duplicate detection
- Pruning impossible combinations
- Parallel processing for large arrays

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding 4 people whose total weight equals target:**
- Array is list of people with different weights
- Need to find exactly 4 people whose combined weight = target
- Sort by weight to enable efficient search
- Use 3-level selection: pick first person, then find 2 more, then find 1 more

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers, target sum
2. **Goal**: Find all unique quadruplets that sum to target
3. **Output**: List of lists containing 4 integers each
4. **Constraints**: Order within quadruplet doesn't matter

#### Phase 2: Key Insight Recognition
- **"How to find 4 numbers efficiently?"** → Reduce to 2Sum problem
- **"Why sort first?"** → Enables early termination and duplicate handling
- **"How to avoid duplicates?"** → Skip same values in appropriate positions
- **"What's the pattern?"** → Nested loops: i < j < k < l

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use nested two-pointers:
1. Sort array to enable optimizations
2. Fix first element (i), then find 3 more (j,k,l)
3. For each combination, check if sum equals target
4. Skip duplicates to avoid repeated results
5. Use early termination when minimum possible sum exceeds target
6. Collect all valid quadruplets"
```

#### Phase 4: Algorithm Walkthrough
```
Example: nums=[1,0,-1,0,-2,2], target=0

Human thinking:
"Sort: [-2,-1,0,0,1,2]

i=0 (nums[0]=-2):
  j=1 (nums[1]=-1): skip duplicate
  j=2 (nums[2]=0): 
    k=3 (nums[3]=0): 
      l=4 (nums[4]=1): sum=-2+-1+0+1=-2 ≠ 0
    k=4 (nums[4]=2): sum=-2+-1+0+2=-1 ≠ 0

i=1 (nums[1]=-1):
  j=2 (nums[2]=0): 
    k=3 (nums[3]=0): 
      l=4 (nums[4]=1): sum=-1+0+0+1=0 ✓ Found: [-1,0,0,1]

Continue this process...
Final result: [[-2,-1,1,2], [-2,0,0,2], [-1,0,1,1]] ✓"
```

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use 4 nested loops?"** → O(N⁴) vs O(N³) with two-pointer
2. **"What about unsorted array?"** → Can use hash map but loses early termination
3. **"How to handle duplicates?"** → Skip in appropriate positions
4. **"When to prune?"** → When sorted, break early if sum > target

### Real-World Analogy
**Like forming a project team with specific budget:**
- Array represents available team members with different costs
- Need exactly 4 people whose total cost = budget
- Sort by cost to find optimal combinations
- Skip people with same skills to avoid duplicate roles
- Early termination when minimum cost exceeds budget

### Human-Readable Pseudocode
```
function fourSum(nums, target):
    if nums.length < 4: return []
    
    sort(nums)
    result = []
    
    for i from 0 to n-4:
        if i > 0 and nums[i] == nums[i-1]: continue
        
        if nums[i] + nums[i+1] + nums[i+2] + nums[i+3] > target:
            break
            
        for j from i+1 to n-3:
            if j > i+1 and nums[j] == nums[j-1]: continue
            
            for k from j+1 to n-2:
                if nums[i] + nums[j] + nums[k] + nums[k+1] > target:
                    break
                    
                for l from k+1 to n-1:
                    if nums[i] + nums[j] + nums[k] + nums[l] == target:
                        result.add([nums[i], nums[j], nums[k], nums[l]])
    
    return result
```

### Execution Visualization

### Example: nums=[1,0,-1,0,-2,2], target=0
```
Sorted Array: [-2,-1,0,0,1,2]

Selection Process:
i=0 (-2):
  j=1 (0): k=2 (0): l=3 (1): sum=-2+0+0+1=-1
  j=1 (0): k=2 (0): l=4 (2): sum=-2+0+0+2=0 ✓ Found: [-2,0,0,1]

i=1 (-1): skip (duplicate of i=0)

i=2 (0): j=3 (0): k=4 (1): l=5 (2): sum=0+0+0+1+2=3 > 0, break

Final Valid Combinations:
[-2,-1,0,1]  (sum=-3)
[-2,0,0,2]  (sum=-2)
[-1,0,0,1]   (sum=0) ✓
```

### Key Visualization Points:
- **Nested Pointers**: i < j < k < l pattern for 4-sum selection
- **Early Termination**: Break when minimum possible sum exceeds target
- **Duplicate Skipping**: Essential in sorted array to avoid repeats
- **2Sum Reduction**: Inner problem becomes finding 2 numbers for remaining sum

### Time Complexity Breakdown:
- **Sorting**: O(N log N)
- **Main Loops**: O(N³) in worst case
- **Early Termination**: Can significantly reduce average case
- **Space**: O(N) for result storage (excluding output)
- **Optimal**: O(N³) is best possible for 4Sum problem
*/
