import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

public class ContainsDuplicate {
    
    // 217. Contains Duplicate
    // Time: O(N), Space: O(N)
    public static boolean containsDuplicate(int[] nums) {
        Set<Integer> numSet = new HashSet<>();
        
        for (int num : nums) {
            if (!numSet.add(num)) {
                return true;
            }
        }
        
        return false;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 2, 3, 1},
            {1, 2, 3, 4},
            {1, 1, 1, 3, 2, 2, 2},
            {},
            {0},
            {-1, -2, -3, -1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            boolean result = containsDuplicate(testCases[i]);
            System.out.printf("Test Case %d: %s -> Contains duplicate: %b\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Hash Set for Duplicate Detection
- **Hash Set**: Stores unique elements encountered
- **Early Detection**: Return true as soon as duplicate found
- **Single Pass**: Process array once
- **Set Properties**: add() returns false for duplicates

## 2. PROBLEM CHARACTERISTICS
- **Array Input**: Collection of integers
- **Duplicate Check**: Determine if any element appears twice
- **Boolean Result**: Simple true/false answer
- **Early Exit**: Can stop as soon as duplicate found

## 3. SIMILAR PROBLEMS
- Contains Duplicate II (within k distance)
- Contains Duplicate III (within t value)
- Find All Duplicates in Array
- Single Number

## 4. KEY OBSERVATIONS
- Hash set provides O(1) duplicate detection
- Set.add() returns false if element already exists
- Can return immediately when duplicate found
- No need to process entire array if duplicate found early

## 5. VARIATIONS & EXTENSIONS
- Find all duplicates
- Count occurrences of each element
- Handle distance constraints (k apart)
- Handle value difference constraints (t apart)

## 6. INTERVIEW INSIGHTS
- Clarify: "What counts as duplicate?" (exact same value)
- Edge cases: empty array, single element, all unique
- Space-time tradeoff: O(N) space vs O(N²) time
- Alternative: Sort first, then check adjacent

## 7. COMMON MISTAKES
- Using array instead of set (inefficient)
- Not returning early when duplicate found
- Mishandling empty array case
- Forgetting about hash set properties

## 8. OPTIMIZATION STRATEGIES
- Current solution is optimal for time
- For memory constraints, consider sorting approach
- For streaming data, use bounded-size set
- For range-limited data, use boolean array

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like tracking party guests:**
- You're checking IDs at a party entrance
- You keep a list of everyone who has entered
- When someone new arrives, you check if they're already on your list
- If yes, they're a duplicate - party crasher!
- If no, you add them to your list

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of integers
2. **Goal**: Check if any number appears more than once
3. **Output**: True if duplicate exists, false otherwise

#### Phase 2: Key Insight Recognition
- **"I need to remember what I've seen!"** → Use a set
- **"How to check quickly?"** → Hash set O(1) lookup
- **"When to stop?"** → As soon as duplicate found
- **"Set.add() is perfect!"** → Returns false for duplicates

#### Phase 3: Strategy Development
```
Human thought process:
"I'll create an empty set to track unique numbers.
For each number, I'll try to add it to the set.
If add() returns false, it means I've seen it before - duplicate!
If I make it through the whole array, no duplicates found."
```

#### Phase 4: Edge Case Handling
- **Empty array**: No duplicates, return false
- **Single element**: Cannot have duplicate, return false
- **All same elements**: First duplicate at index 1
- **All unique**: Process entire array, return false

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Starting with: [1, 2, 3, 1]

Human thinking:
"Let me create my empty set.
Position 0: number = 1, set = {} → add 1 successfully
Position 1: number = 2, set = {1} → add 2 successfully  
Position 2: number = 3, set = {1,2} → add 3 successfully
Position 3: number = 1, set = {1,2,3} → try to add 1
           1 is already in set! Duplicate found!
           I can stop right here and return true."
```

#### Phase 6: Intuition Validation
- **Why it works**: Set automatically handles uniqueness
- **Why it's efficient**: O(1) operations, early exit possible
- **Why it's correct**: We check every element against all previous ones

### Common Human Pitfalls & How to Avoid Them
1. **"Should I use an array?"** → No, set is much faster
2. **"Can I optimize space?"** → Only with sorting (slower)
3. **"What about early exit?"** → Always return as soon as duplicate found
4. **"How to handle negatives?"** → Set handles them fine

### Real-World Analogy
**Like checking for duplicate library books:**
- You're scanning books into a library system
- You keep a record of all book IDs you've scanned
- When you scan a new book, you check if it's already in your system
- If yes, someone's trying to return the same book twice!
- If no, you add it to your system and continue

### Human-Readable Pseudocode
```
function containsDuplicate(numbers):
    seenNumbers = empty set
    
    for each number in numbers:
        if number cannot be added to seenNumbers:
            return true  // duplicate found
        add number to seenNumbers
    
    return false  // no duplicates found
```

### Execution Visualization

### Example: [1, 2, 3, 1]
```
Initial: nums = [1, 2, 3, 1], set = {}

Step 1: i=0, nums[0]=1
→ add(1) succeeds
→ set = {1}
State: set={1}, result=false

Step 2: i=1, nums[1]=2
→ add(2) succeeds
→ set = {1, 2}
State: set={1,2}, result=false

Step 3: i=2, nums[2]=3
→ add(3) succeeds
→ set = {1, 2, 3}
State: set={1,2,3}, result=false

Step 4: i=3, nums[3]=1
→ add(1) fails (already exists)
→ Return true immediately
State: set={1,2,3}, result=true
```

### Key Visualization Points:
- **Hash set** stores unique elements only
- **add() operation** fails for duplicates
- **Early return** when duplicate detected
- **Single pass** through array

### Memory Layout Visualization:
```
Array: [1][2][3][1]
       ^
       |
       +-- Current: 1 (duplicate!)
       
Set Evolution:
Step 0: {}
Step 1: {1}
Step 2: {1,2}
Step 3: {1,2,3}
Step 4: {1,2,3} ✓ Duplicate found!
```

### Time Complexity Breakdown:
- **Best Case**: O(1) - duplicate found at index 1
- **Worst Case**: O(N) - no duplicates, process all elements
- **Average Case**: O(N) - typical scenario
- **Space**: O(N) - storing unique elements
*/
