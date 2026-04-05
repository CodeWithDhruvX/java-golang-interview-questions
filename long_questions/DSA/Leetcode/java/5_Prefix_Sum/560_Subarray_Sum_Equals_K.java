import java.util.*;

public class SubarraySumEqualsK {
    
    // 560. Subarray Sum Equals K
    // Time: O(N), Space: O(N)
    public static int subarraySum(int[] nums, int k) {
        int count = 0;
        Map<Integer, Integer> prefixSum = new HashMap<>();
        prefixSum.put(0, 1); // Initialize with sum 0 occurring once
        
        int currentSum = 0;
        
        for (int num : nums) {
            currentSum += num;
            
            // Check if (currentSum - k) exists in prefixSum
            if (prefixSum.containsKey(currentSum - k)) {
                count += prefixSum.get(currentSum - k);
            }
            
            // Update prefixSum with currentSum
            prefixSum.put(currentSum, prefixSum.getOrDefault(currentSum, 0) + 1);
        }
        
        return count;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testArrays = {
            {1, 1, 1},
            {1, 2, 3},
            {1, -1, 0},
            {0, 0, 0, 0},
            {-1, -1, 1},
            {3, 4, 7, -2, 2, 1, 4, 2},
            {1},
            {1},
            {},
            {1, 2, 1, 2, 1}
        };
        
        int[] kValues = {2, 3, 0, 0, 0, 7, 1, 0, 0, 3};
        
        for (int i = 0; i < testArrays.length; i++) {
            int result = subarraySum(testArrays[i], kValues[i]);
            System.out.printf("Test Case %d: nums=%s, k=%d -> Subarrays with sum k: %d\n", 
                i + 1, Arrays.toString(testArrays[i]), kValues[i], result);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Prefix Sum with Hash Map
- **Prefix Sum**: Running sum of array elements
- **Hash Map**: Stores frequency of prefix sums
- **Target Check**: Look for (currentSum - target) in map
- **Frequency Count**: Multiple occurrences of same prefix sum

## 2. PROBLEM CHARACTERISTICS
- **Subarray Sum**: Need continuous subarrays that sum to K
- **Negative Numbers**: Allowed, makes sliding window difficult
- **Count Problem**: Need number of such subarrays
- **Efficiency Required**: O(N²) brute force too slow

## 3. SIMILAR PROBLEMS
- Subarray Sum Equals K (variations)
- Maximum Size Subarray Sum Equals K
- Number of Subarrays with Sum Less Than K
- Continuous Subarray Sum

## 4. KEY OBSERVATIONS
- Prefix sum enables O(1) subarray sum calculation
- sum[i,j] = prefix[j] - prefix[i-1]
- Need subarray sum = K → prefix[j] - prefix[i-1] = K
- Therefore: prefix[i-1] = prefix[j] - K

## 5. VARIATIONS & EXTENSIONS
- Find actual subarrays (not just count)
- Handle large numbers with modulo
- Multiple queries on same array
- Streaming version with limited memory

## 6. INTERVIEW INSIGHTS
- Clarify: "Can array contain negative numbers?"
- Edge cases: empty array, all zeros, single element
- Why sliding window doesn't work with negatives
- Hash map vs array for prefix sums

## 7. COMMON MISTAKES
- Not initializing prefix sum with 0
- Off-by-one errors in prefix sum indexing
- Using sliding window (doesn't work with negatives)
- Not handling multiple occurrences properly

## 8. OPTIMIZATION STRATEGIES
- Current solution is optimal for single pass
- For multiple queries, consider prefix sum array
- For memory constraints, use limited-size hash map
- For streaming, maintain running prefix sum

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like tracking running balance:**
- You're tracking daily bank balance (prefix sum)
- You want to know when balance difference equals target amount
- If current balance - previous balance = target, you found a period!
- You keep track of all previous balances for quick lookup

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers and target sum K
2. **Goal**: Count continuous subarrays that sum to K
3. **Output**: Number of such subarrays

#### Phase 2: Key Insight Recognition
- **"Subarray sum = prefix[j] - prefix[i-1]!"** → O(1) calculation
- **"Need prefix[j] - prefix[i-1] = K"** → Rearrange
- **"Therefore prefix[i-1] = prefix[j] - K"** → Hash map lookup!
- **"Initialize with sum 0!"** → Handle subarrays starting at index 0

#### Phase 3: Strategy Development
```
Human thought process:
"I'll maintain a running sum as I traverse the array.
For each position, I'll check if (currentSum - K)
has been seen before as a prefix sum.
If yes, every occurrence represents a valid subarray ending here.
I'll store the current prefix sum for future lookups."
```

#### Phase 4: Edge Case Handling
- **Empty array**: No subarrays, return 0
- **All zeros**: Every subarray sums to 0, handle K=0
- **Single element**: Either 0 or 1 subarray depending on K
- **Negative numbers**: Handled correctly by prefix sum approach

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: nums = [1, 1, 1], K = 2

Human thinking:
"Initialize prefixSum = {0:1} (sum 0 seen once)
currentSum = 0

Position 0: num = 1, currentSum = 1
           Need to find (1 - 2) = -1 in prefixSum
           -1 not found, no subarrays ending here
           Store currentSum: prefixSum = {0:1, 1:1}

Position 1: num = 1, currentSum = 2
           Need to find (2 - 2) = 0 in prefixSum
           0 found once! One subarray ending here: [0,1]
           Store currentSum: prefixSum = {0:1, 1:1, 2:1}

Position 2: num = 1, currentSum = 3
           Need to find (3 - 2) = 1 in prefixSum
           1 found once! One subarray ending here: [2]
           Store currentSum: prefixSum = {0:1, 1:2, 2:1, 3:1}

Total subarrays found: 2"
```

#### Phase 6: Intuition Validation
- **Why it works**: Prefix sum difference equals subarray sum
- **Why it's efficient**: O(1) lookup for each position
- **Why it's correct**: Every valid subarray will be counted exactly once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not sliding window?"** → Doesn't work with negative numbers
2. **"What about initialization?"** → Must include sum 0 occurring once
3. **"Can I use array instead of hash map?"** → Hash map handles large ranges
4. **"What about multiple occurrences?"** → Use frequency count

### Real-World Analogy
**Like finding profitable trading periods:**
- You have daily stock price changes (array)
- You want to find periods where total change equals target profit
- You track cumulative profit (prefix sum)
- When current cumulative - previous cumulative = target profit, you found a period!
- You remember all cumulative profits for quick comparison

### Human-Readable Pseudocode
```
function countSubarraysWithSumK(numbers, target):
    prefixSums = {0: 1}  // sum 0 seen once
    currentSum = 0
    count = 0
    
    for each number in numbers:
        currentSum += number
        
        needed = currentSum - target
        if needed exists in prefixSums:
            count += prefixSums[needed]
        
        prefixSums[currentSum] = prefixSums.getOrDefault(currentSum, 0) + 1
    
    return count
```

### Execution Visualization

### Example: [1, 1, 1], K = 2
```
Initial: nums = [1, 1, 1], K = 2
prefixSum = {0: 1}, currentSum = 0, count = 0

Step 1: i=0, nums[0]=1
→ currentSum = 1
→ needed = 1 - 2 = -1
→ -1 not in prefixSum
→ prefixSum = {0: 1, 1: 1}
State: count = 0

Step 2: i=1, nums[1]=1
→ currentSum = 2
→ needed = 2 - 2 = 0
→ 0 found in prefixSum with count 1
→ count = 0 + 1 = 1 (subarray [0,1] found)
→ prefixSum = {0: 1, 1: 1, 2: 1}
State: count = 1

Step 3: i=2, nums[2]=1
→ currentSum = 3
→ needed = 3 - 2 = 1
→ 1 found in prefixSum with count 1
→ count = 1 + 1 = 2 (subarray [2] found)
→ prefixSum = {0: 1, 1: 2, 2: 1, 3: 1}
State: count = 2

Final: Return 2
```

### Key Visualization Points:
- **Prefix sum** tracks running total
- **Hash map lookup** finds previous sums that create valid subarrays
- **Frequency count** handles multiple occurrences
- **Initialization** with sum 0 handles subarrays starting at index 0

### Memory Layout Visualization:
```
Array:    [1][1][1]
Index:     0  1  2
Prefix:    0  1  2  3
Needed:   -1  0  1
Found:     ✗  ✓  ✓
Subarrays:     [0,1] [2]
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each element visited once
- **Hash Operations**: O(1) average for lookup and insert
- **Total**: O(N) time, O(N) space
- **Optimal**: Cannot do better for arbitrary arrays with negatives
*/
