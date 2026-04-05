import java.util.*;

public class Permutations {
    
    // 46. Permutations
    // Time: O(N * N!), Space: O(N!) for result + O(N) for recursion
    public static List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        boolean[] used = new boolean[nums.length];
        int[] current = new int[nums.length];
        
        backtrackPermute(nums, used, current, 0, result);
        return result;
    }

    private static void backtrackPermute(int[] nums, boolean[] used, int[] current, 
                                   int pos, List<List<Integer>> result) {
        if (pos == nums.length) {
            // Make a copy of current permutation
            List<Integer> temp = new ArrayList<>();
            for (int num : current) {
                temp.add(num);
            }
            result.add(temp);
            return;
        }
        
        for (int i = 0; i < nums.length; i++) {
            if (!used[i]) {
                used[i] = true;
                current[pos] = nums[i];
                
                backtrackPermute(nums, used, current, pos + 1, result);
                
                used[i] = false;
            }
        }
    }

    // Alternative approach using swapping
    public static List<List<Integer>> permuteSwap(int[] nums) {
        List<List<Integer>> result = new ArrayList<>();
        backtrackSwap(nums, 0, result);
        return result;
    }

    private static void backtrackSwap(int[] nums, int start, List<List<Integer>> result) {
        if (start == nums.length) {
            // Make a copy of current permutation
            List<Integer> temp = new ArrayList<>();
            for (int num : nums) {
                temp.add(num);
            }
            result.add(temp);
            return;
        }
        
        for (int i = start; i < nums.length; i++) {
            swap(nums, start, i);
            backtrackSwap(nums, start + 1, result);
            swap(nums, start, i);
        }
    }

    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 2, 3},
            {0, 1},
            {1},
            {},
            {1, 2, 3, 4},
            {1, 1, 2},
            {1, 2, 2},
            {5, 6, 7},
            {1, 2, 3, 4, 5}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            // Make copies for both methods
            int[] nums1 = testCases[i].clone();
            int[] nums2 = testCases[i].clone();
            
            List<List<Integer>> result1 = permute(nums1);
            List<List<Integer>> result2 = permuteSwap(nums2);
            
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            System.out.printf("  Used array: %d permutations\n", result1.size());
            System.out.printf("  Swap method: %d permutations\n\n", result2.size());
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Backtracking
- **Permutation Generation**: Generate all possible arrangements
- **Recursive Exploration**: Try each unused element at each position
- **Backtracking**: Undo choices when exploring dead ends
- **Complete Solutions**: When all positions are filled

## 2. PROBLEM CHARACTERISTICS
- **Permutations**: All possible orderings of array elements
- **No Duplicates**: Each element appears exactly once
- **Complete Search**: Must explore all possibilities
- **Result Collection**: Store all valid permutations

## 3. SIMILAR PROBLEMS
- Combinations
- Subsets
- Letter Case Permutation
- N-Queens Problem

## 4. KEY OBSERVATIONS
- Number of permutations = N! for N elements
- Backtracking explores each position systematically
- Used array tracks which elements are placed
- Each recursive call fixes one position
- Time complexity is O(N!) by necessity

## 5. VARIATIONS & EXTENSIONS
- Permutations with duplicates allowed
- K-permutations (select K from N)
- Next permutation (lexicographic order)
- Permutation with constraints

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I handle duplicate elements?"
- Edge cases: empty array, single element
- Time complexity: O(N!) vs O(N^N) naive
- Space complexity: O(N!) for result storage

## 7. COMMON MISTAKES
- Not handling used array correctly
- Forgetting to backtrack (undo choices)
- Creating result copies incorrectly
- Stack overflow for large N

## 8. OPTIMIZATION STRATEGIES
- Swapping approach reduces memory usage
- Iterative generation with next_permutation
- Early pruning for constrained permutations
- In-place modifications save space

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like arranging books on a shelf:**
- You have N different books (array elements)
- You want to try all possible arrangements
- For each position on shelf, choose an unused book
- If arrangement works, record it
- Backtrack and try different book for that position
- Continue until all positions are filled

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of distinct integers
2. **Goal**: Generate all possible permutations
3. **Output**: List of all permutations

#### Phase 2: Key Insight Recognition
- **"How many permutations?"** → N! for N elements
- **"How to generate systematically?"** → Backtracking recursion
- **"What to track?"** → Used elements and current position
- **"When to stop?"** → When all positions filled

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use backtracking:
1. For each position from 0 to N-1:
   - Try each unused element at this position
   - Mark element as used
   - Recursively fill next positions
   - Backtrack: unmark element when returning
2. When all positions filled, record permutation
3. This explores all N! possibilities"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return empty list
- **Single element**: Return [[element]]
- **Large N**: May cause stack overflow
- **Memory constraints**: N! grows very quickly

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Input: [1, 2, 3]

Human thinking:
"Start with empty permutation:
position = 0, used = [false, false, false], current = []

Position 0:
- Try 1: used[0] = false, place 1
  → current = [1], used = [true, false, false]
  → Recurse to position 1

Position 1:
- Try 2: used[1] = false, place 2
  → current = [1, 2], used = [true, true, false]
  → Recurse to position 2

Position 2:
- Try 3: used[2] = false, place 3
  → current = [1, 2, 3], used = [true, true, true]
  → All positions filled, record [1, 2, 3]
  → Backtrack: remove 3, used[2] = false

Position 1 (backtrack):
- Try 3: used[2] = false, place 3
  → current = [1, 3], used = [true, false, true]
  → Recurse to position 2

Position 2:
- Try 2: used[1] = true, skip
- Try remaining options...
Continue exploring all possibilities...

Final result: [[1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], [3,2,1]] ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Systematic exploration of all possibilities
- **Why it's complete**: Each position tries every unused element
- **Why it's correct**: Backtracking ensures no duplicates

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just generate randomly?"** → Won't get all permutations
2. **"What about iterative approach?"** → Backtracking is more natural
3. **"How to handle duplicates?"** → Need counting and skipping logic
4. **"What about memory?"** → N! grows very fast

### Real-World Analogy
**Like creating all possible schedules:**
- You have N tasks (array elements)
- You want to try all possible orderings
- For each time slot, choose an unscheduled task
- Schedule it and mark as used
- Continue until all time slots are filled
- Backtrack when a schedule doesn't work
- Each complete schedule is one permutation

### Human-Readable Pseudocode
```
function permute(nums):
    result = []
    used = [false] * nums.length
    current = [] * nums.length
    
    backtrack(nums, used, current, 0, result)
    return result

function backtrack(nums, used, current, pos, result):
    if pos == nums.length:
        result.add(copy(current))
        return
    
    for i from 0 to nums.length-1:
        if not used[i]:
            used[i] = true
            current[pos] = nums[i]
            backtrack(nums, used, current, pos + 1, result)
            used[i] = false  // backtrack
```

### Execution Visualization

### Example: nums = [1, 2, 3]
```
Backtracking Tree:
                    pos=0: try 1
                   /         pos=1: try 2
                  /           pos=2: try 3
                 /             [1,2,3] ✓
                pos=2: try 2 (used)
               pos=2: try 1 (used)
              
              pos=1: try 3
             /         pos=2: try 2
            /          [1,3,2] ✓
           pos=2: try 1
          /           [1,3,1] ✓
         
        pos=0: try 2
       /         pos=1: try 1
      /          pos=2: try 3
     /           [2,1,3] ✓
    pos=1: try 3
   /         pos=2: try 2
  /          [2,3,1] ✓
 pos=2: try 1
/           [2,3,2] ✓

All 6 permutations generated ✓
```

### Key Visualization Points:
- **Backtracking tree** explores all choice combinations
- **Used array** prevents duplicate selections
- **Position tracking** builds permutations incrementally
- **Complete solutions** when reaching leaf nodes

### Memory Layout Visualization:
```
Recursive Stack for [1,2,3]:
Level 0: permute([1,2,3], pos=0)
├─ Try 1: permute([_,2,3], pos=1)
│  ├─ Try 2: permute([1,_,3], pos=2)
│  │  └─ Try 3: [1,2,3] ✓
│  └─ Try 3: permute([1,_,2], pos=2)
│     └─ Try 2: [1,3,2] ✓
└─ Try 2: permute([1,_,3], pos=1)
   ├─ Try 1: permute([2,_,3], pos=2)
   │  └─ Try 3: [2,1,3] ✓
   └─ Try 3: permute([2,_,1], pos=2)
      └─ Try 1: [2,3,1] ✓

Used Array Evolution:
[1,2,3] → [true,true,true] → [true,true,true] → backtrack...
```

### Time Complexity Breakdown:
- **Each Level**: N choices, then N-1, then N-2, etc.
- **Total Permutations**: N! possibilities
- **Time**: O(N!) for generating all permutations
- **Space**: O(N!) for result storage + O(N) for recursion
- **Optimal**: Cannot do better than O(N!) for this problem
*/
