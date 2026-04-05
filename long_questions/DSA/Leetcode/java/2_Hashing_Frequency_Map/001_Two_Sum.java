import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class TwoSum {
    
    // 1. Two Sum
    // Time: O(N), Space: O(N)
    public static int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> numMap = new HashMap<>();
        
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            if (numMap.containsKey(complement)) {
                return new int[]{numMap.get(complement), i};
            }
            numMap.put(nums[i], i);
        }
        
        return new int[]{};
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testArrays = {
            {2, 7, 11, 15},
            {3, 2, 4},
            {3, 3},
            {1, 2, 3, 4, 5},
            {-1, -2, -3, -4, -5}
        };
        
        int[] targets = {9, 6, 6, 9, -8};
        
        for (int i = 0; i < testArrays.length; i++) {
            int[] result = twoSum(testArrays[i], targets[i]);
            System.out.printf("Test Case %d: nums=%s, target=%d -> Indices: %s\n", 
                i + 1, Arrays.toString(testArrays[i]), targets[i], Arrays.toString(result));
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Hash Map for Complement Lookup
- **Hash Map**: Stores previously seen numbers and their indices
- **Complement Calculation**: target - current number
- **Single Pass**: Find solution in one traversal
- **Early Return**: Exit as soon as solution found

## 2. PROBLEM CHARACTERISTICS
- **Array Input**: Unsorted array of integers
- **Target Sum**: Need to find two numbers that sum to target
- **Unique Solution**: Guaranteed exactly one solution
- **Index Return**: Return indices (not values)

## 3. SIMILAR PROBLEMS
- Two Sum II (sorted array)
- Three Sum
- Four Sum
- Subarray Sum Equals K

## 4. KEY OBSERVATIONS
- Need to find complement (target - current number)
- Hash map provides O(1) lookup for complement
- Store numbers as we process them
- Can return immediately when complement found

## 5. VARIATIONS & EXTENSIONS
- Multiple solutions possible
- Return values instead of indices
- Handle duplicates differently
- Find all pairs that sum to target

## 6. INTERVIEW INSIGHTS
- Clarify: "Is there exactly one solution?"
- Clarify: "Can we use the same element twice?"
- Edge cases: empty array, no solution, negative numbers
- Space-time tradeoff: O(N) space vs O(N²) brute force

## 7. COMMON MISTAKES
- Using same element twice
- Not handling negative numbers
- Returning wrong indices order
- Forgetting to store current number before checking

## 8. OPTIMIZATION STRATEGIES
- Current solution is optimal for single query
- For multiple queries, consider sorting + two pointers
- For memory constraints, use bit manipulation tricks
- For streaming data, use limited-size hash map

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding a dance partner:**
- You have a list of people with numbers
- You need to find two people whose numbers add up to target
- As you meet each person, you remember their number
- When meeting someone new, you check if their "partner number" is already known
- If yes, you've found your pair!

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of numbers and target sum
2. **Goal**: Find two numbers that add to target
3. **Output**: Their indices in the array

#### Phase 2: Key Insight Recognition
- **"I need to find complement!"** → target - current number
- **"How to find complement quickly?"** → Hash map lookup
- **"What to store?"** → number and its index
- **"When to check?"** → Before storing current number

#### Phase 3: Strategy Development
```
Human thought process:
"I'll keep a map of numbers I've seen so far.
For each number, I'll calculate what I need (complement).
If I've already seen the complement, I found my pair!
If not, I'll add current number to the map and continue."
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty array
- **No solution**: Return empty array (though problem guarantees solution)
- **Single element**: Cannot form pair
- **Negative numbers**: Handled normally by complement calculation

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: nums = [2, 7, 11, 15], target = 9

Human thinking:
"Let me create my empty map first.
Position 0: number = 2, complement = 9-2 = 7
           Is 7 in my map? No.
           Store 2 at position 0.
           Map: {2:0}

Position 1: number = 7, complement = 9-7 = 2
           Is 2 in my map? Yes! At position 0.
           Found my pair: [0, 1]
           I'm done!"
```

#### Phase 6: Intuition Validation
- **Why it works**: We check complement before storing, avoiding self-pairing
- **Why it's efficient**: O(1) hash map lookups, single pass
- **Why it's correct**: We'll find the pair when we encounter the second element

### Common Human Pitfalls & How to Avoid Them
1. **"Should I store first then check?"** → No, you might use same element twice
2. **"What about order of indices?"** → Complement index first, then current
3. **"Can I optimize space?"** → Not without changing time complexity
4. **"What if no solution exists?"** → Return empty array

### Real-World Analogy
**Like finding matching socks:**
- You have a pile of socks (numbers)
- You're looking for a matching pair that sums to target
- As you pick up each sock, you remember its color
- When you pick up a new sock, you check if you've seen its match
- If yes, you've found your pair!

### Human-Readable Pseudocode
```
function twoSum(numbers, target):
    seenNumbers = empty map
    
    for i from 0 to numbers.length-1:
        current = numbers[i]
        complement = target - current
        
        if complement exists in seenNumbers:
            return [seenNumbers[complement], i]
        
        store current with its index in seenNumbers
    
    return empty array
```

### Execution Visualization

### Example: [2, 7, 11, 15], target = 9
```
Initial: nums = [2, 7, 11, 15], target = 9, map = {}

Step 1: i=0, nums[0]=2
→ complement = 9-2 = 7
→ 7 not in map
→ Store 2: map = {2:0}
State: map={2:0}, result=[]

Step 2: i=1, nums[1]=7
→ complement = 9-7 = 2
→ 2 is in map at position 0
→ Found pair: [0, 1]
State: map={2:0, 7:1}, result=[0, 1]

Final: Return [0, 1]
```

### Key Visualization Points:
- **Hash map** stores number → index mapping
- **Complement calculation** happens for each element
- **Early return** when solution found
- **Single pass** through the array

### Memory Layout Visualization:
```
Array: [2][7][11][15]
       ^  ^
       |  |
       |  +-- Current: 7, Complement: 2 (found!)
       +----- Stored: 2 at index 0

Map Evolution:
Step 0: {}
Step 1: {2:0}
Step 2: {2:0, 7:1} ✓ Solution found!
```

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each element visited once
- **Hash Operations**: O(1) average for insert and lookup
- **Total**: O(N) time, O(N) space
- **Optimal**: Cannot do better than O(N) for unsorted array
*/
