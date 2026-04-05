import java.util.Arrays;

public class SortColors {
    
    // 75. Sort Colors
    // Time: O(N), Space: O(1)
    public static void sortColors(int[] nums) {
        int low = 0, mid = 0, high = nums.length - 1;
        
        while (mid <= high) {
            if (nums[mid] == 0) {
                swap(nums, low, mid);
                low++;
                mid++;
            } else if (nums[mid] == 1) {
                mid++;
            } else {
                swap(nums, mid, high);
                high--;
            }
        }
    }
    
    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

    public static void main(String[] args) {
        int[][] testCases = {
            {2, 0, 2, 1, 1, 0},
            {0, 1, 2, 0, 1, 2},
            {2, 2, 1, 1, 0, 0},
            {0, 0, 0},
            {1, 1, 1},
            {2, 2, 2},
            {0},
            {1},
            {2},
            {1, 0, 2, 1, 0, 2, 1, 0}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] nums = testCases[i].clone();
            int[] original = nums.clone();
            
            sortColors(nums);
            System.out.printf("Test Case %d: %s -> After sorting: %s\n", 
                i + 1, Arrays.toString(original), Arrays.toString(nums));
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Dutch National Flag
- **Three Pointers**: low, mid, high partition array
- **Single Pass**: Process all elements in one traversal
- **In-Place Sorting**: No extra space needed
- **Color Partitioning**: 0s, 1s, 2s in order

## 2. PROBLEM CHARACTERISTICS
- **Three Categories**: Exactly three distinct values (0, 1, 2)
- **In-Place Requirement**: Must modify original array
- **Order Preservation**: All 0s first, then 1s, then 2s
- **Single Pass Constraint**: Optimal solution uses O(N) time

## 3. SIMILAR PROBLEMS
- Partition array around pivot
- Sort binary array (0s and 1s)
- Move zeros to end
- Separate even and odd numbers

## 4. KEY OBSERVATIONS
- **low pointer**: Tracks boundary for 0s
- **mid pointer**: Current element being processed
- **high pointer**: Tracks boundary for 2s
- **Invariant**: Elements before low are 0, after high are 2

## 5. VARIATIONS & EXTENSIONS
- Generalize to K colors (requires counting sort)
- Four-way partitioning
- Stable partitioning (preserve relative order)
- Partition with custom comparison

## 6. INTERVIEW INSIGHTS
- Clarify: "Are there exactly three colors, or could be more?"
- Edge cases: empty array, single element, all same color
- Time complexity: O(N) is optimal
- Space complexity: O(1) - in-place requirement

## 7. COMMON MISTAKES
- Not handling empty array case
- Off-by-one errors in pointer management
- Breaking invariants during swaps
- Not processing all elements

## 8. OPTIMIZATION STRATEGIES
- Dutch National Flag is optimal for 3 colors
- Counting sort works but uses O(N) space
- Single pass is key requirement
- Minimize swaps for better cache performance

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like sorting colored balls in buckets:**
- You have red (0), white (1), and blue (2) balls mixed together
- You want to sort them: reds first, then whites, then blues
- Use three markers to divide the array into four regions
- Process each ball once and move it to correct region

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array with only 0s, 1s, and 2s
2. **Goal**: Sort in-place with 0s first, 1s middle, 2s last
3. **Output**: Modified array (void return)

#### Phase 2: Key Insight Recognition
- **"How to sort three colors efficiently?"** → Three pointers approach
- **"What should each pointer do?"** → low for 0s, mid for current, high for 2s
- **"How to maintain invariants?"** → Careful pointer movement and swapping
- **"Single pass possible?"** → Yes, with Dutch National Flag algorithm

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use three pointers:
1. low: boundary for next 0
2. mid: current element being examined
3. high: boundary for next 2

While mid <= high:
- If nums[mid] is 0: swap with low, move both pointers
- If nums[mid] is 1: just move mid (already in right place)
- If nums[mid] is 2: swap with high, move high only
This maintains the invariant that:
- Elements < low are 0s
- Elements low to mid-1 are 1s
- Elements mid to high are unprocessed
- Elements > high are 2s"
```

#### Phase 4: Algorithm Walkthrough
```
Example: [2, 0, 2, 1, 1, 0]

Initial: low=0, mid=0, high=5
Array: [2, 0, 2, 1, 1, 0]

Step 1: mid=0, nums[0]=2 (blue)
Swap nums[0] and nums[5]: [0, 0, 2, 1, 1, 2]
high--: high=4

Step 2: mid=0, nums[0]=0 (red)
Swap nums[0] and nums[0]: [0, 0, 2, 1, 1, 2]
low++, mid++: low=1, mid=1

Step 3: mid=1, nums[1]=0 (red)
Swap nums[1] and nums[1]: [0, 0, 2, 1, 1, 2]
low++, mid++: low=2, mid=2

Step 4: mid=2, nums[2]=2 (blue)
Swap nums[2] and nums[4]: [0, 0, 1, 1, 2, 2]
high--: high=3

Step 5: mid=2, nums[2]=1 (white)
mid++: mid=3

Step 6: mid=3, nums[3]=1 (white)
mid++: mid=4

Now mid=4 > high=3, done!
Final: [0, 0, 1, 1, 2, 2] ✓
```

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sort?"** → Need O(N) time, O(1) space
2. **"What about counting sort?"** → Uses extra space, not in-place
3. **"How to handle invariants?"** → Careful pointer movement logic
4. **"When to stop?"** → When mid crosses high

### Real-World Analogy
**Like organizing laundry by color:**
- You have a pile of mixed clothes: reds, whites, blues
- You want to sort them into three piles
- Use three dividers to create four regions
- Process each item once and move to correct pile
- No extra space needed - just rearrange in place

### Human-Readable Pseudocode
```
function sortColors(nums):
    low = 0
    mid = 0
    high = nums.length - 1
    
    while mid <= high:
        if nums[mid] == 0:
            swap(nums, low, mid)
            low++
            mid++
        else if nums[mid] == 1:
            mid++
        else: // nums[mid] == 2
            swap(nums, mid, high)
            high--
```

### Execution Visualization

### Example: [2, 0, 2, 1, 1, 0]
```
Initial: low=0, mid=0, high=5
Array:   [2, 0, 2, 1, 1, 0]
Regions: [ |unprocessed| ]

Step 1: nums[mid]=2 (blue)
Swap with high: [0, 0, 2, 1, 1, 2]
high=4, mid=0
Regions: [ |unprocessed|2]

Step 2: nums[mid]=0 (red)
Swap with low: [0, 0, 2, 1, 1, 2]
low=1, mid=1
Regions: [0|unprocessed|2]

Step 3: nums[mid]=0 (red)
Swap with low: [0, 0, 2, 1, 1, 2]
low=2, mid=2
Regions: [0,0|unprocessed|2]

Step 4: nums[mid]=2 (blue)
Swap with high: [0, 0, 1, 1, 2, 2]
high=3, mid=2
Regions: [0,0|unprocessed|2,2]

Final: [0, 0, 1, 1, 2, 2]
Regions: [0,0,1,1| |2,2]
```

### Key Visualization Points:
- **Three Pointers**: Create four distinct regions
- **Invariant Maintenance**: Always keep colors in correct regions
- **Single Pass**: Each element processed exactly once
- **In-Place**: No additional space required

### Time Complexity Breakdown:
- **Single Pass**: O(N) - each element visited once
- **Constant Space**: O(1) - only three pointers
- **Optimal**: Cannot do better than O(N) for this problem
- **Cache Friendly**: Sequential access pattern
*/
