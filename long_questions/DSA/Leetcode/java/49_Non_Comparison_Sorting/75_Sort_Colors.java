import java.util.*;

public class SortColors {
    
    // 75. Sort Colors - Non-Comparison Sorting
    // Time: O(N), Space: O(1) for counting sort
    
    // Counting Sort implementation
    public int[] sortColorsCounting(int[] nums) {
        if (nums.length == 0) {
            return nums;
        }
        
        // Count occurrences of each color
        int[] counts = new int[3];
        for (int num : nums) {
            counts[num]++;
        }
        
        // Reconstruct array
        int index = 0;
        for (int color = 0; color < 3; color++) {
            for (int i = 0; i < counts[color]; i++) {
                nums[index++] = color;
            }
        }
        
        return nums;
    }
    
    // Dutch National Flag algorithm (optimized counting sort)
    public int[] sortColorsDutchNationalFlag(int[] nums) {
        if (nums.length == 0) {
            return nums;
        }
        
        // Three pointers: low, mid, high
        int low = 0, mid = 0, high = nums.length - 1;
        
        while (mid <= high) {
            switch (nums[mid]) {
                case 0: // Red
                    swap(nums, low, mid);
                    low++;
                    mid++;
                    break;
                case 1: // White
                    mid++;
                    break;
                case 2: // Blue
                    swap(nums, mid, high);
                    high--;
                    break;
            }
        }
        
        return nums;
    }
    
    private void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
    
    // Radix Sort implementation
    public int[] sortColorsRadix(int[] nums) {
        if (nums.length == 0) {
            return nums;
        }
        
        // Find maximum number to determine number of digits
        int maxNum = 0;
        for (int num : nums) {
            if (num > maxNum) {
                maxNum = num;
            }
        }
        
        // Perform counting sort for each digit
        for (int exp = 1; maxNum / exp > 0; exp *= 10) {
            nums = countingSortByDigit(nums, exp);
        }
        
        return nums;
    }
    
    private int[] countingSortByDigit(int[] nums, int exp) {
        int n = nums.length;
        int[] output = new int[n];
        int[] count = new int[10];
        
        // Count occurrences of each digit
        for (int i = 0; i < n; i++) {
            count[(nums[i] / exp) % 10]++;
        }
        
        // Calculate cumulative count
        for (int i = 1; i < 10; i++) {
            count[i] += count[i - 1];
        }
        
        // Build output array
        for (int i = n - 1; i >= 0; i--) {
            output[count[(nums[i] / exp) % 10] - 1] = nums[i];
            count[(nums[i] / exp) % 10]--;
        }
        
        return output;
    }
    
    // Bucket Sort implementation
    public int[] sortColorsBucket(int[] nums) {
        if (nums.length == 0) {
            return nums;
        }
        
        // Create buckets for each color
        List<Integer>[] buckets = new ArrayList[3];
        for (int i = 0; i < 3; i++) {
            buckets[i] = new ArrayList<>();
        }
        
        // Distribute elements into buckets
        for (int num : nums) {
            buckets[num].add(num);
        }
        
        // Collect elements from buckets
        int index = 0;
        for (int color = 0; color < 3; color++) {
            for (int num : buckets[color]) {
                nums[index++] = num;
            }
        }
        
        return nums;
    }
    
    // Pigeonhole Sort (for small range)
    public int[] sortColorsPigeonhole(int[] nums) {
        if (nums.length == 0) {
            return nums;
        }
        
        // Find min and max
        int min = nums[0], max = nums[0];
        for (int num : nums) {
            if (num < min) min = num;
            if (num > max) max = num;
        }
        
        int range = max - min + 1;
        if (range > nums.length) {
            // Use counting sort when range is reasonable
            return sortColorsCounting(nums);
        }
        
        // Create holes and place elements
        int[] holes = new int[range];
        boolean[] occupied = new boolean[range];
        
        for (int num : nums) {
            int index = num - min;
            if (!occupied[index]) {
                holes[index] = num;
                occupied[index] = true;
            }
        }
        
        // Collect elements from holes
        int index = 0;
        for (int i = 0; i < range; i++) {
            if (occupied[i]) {
                nums[index++] = holes[i];
            }
        }
        
        return nums;
    }
    
    // Version with detailed explanation
    public class SortColorsResult {
        int[] sortedArray;
        List<String> explanation;
        
        SortColorsResult(int[] sortedArray, List<String> explanation) {
            this.sortedArray = sortedArray;
            this.explanation = explanation;
        }
    }
    
    public SortColorsResult sortColorsDetailed(int[] nums) {
        List<String> explanation = new ArrayList<>();
        explanation.add("=== Dutch National Flag Algorithm ===");
        explanation.add("Input: " + Arrays.toString(nums));
        
        if (nums.length == 0) {
            explanation.add("Empty array, returning as is");
            return new SortColorsResult(nums, explanation);
        }
        
        explanation.add("Using three pointers: low (for 0s), mid (current), high (for 2s)");
        
        int low = 0, mid = 0, high = nums.length - 1;
        int step = 1;
        
        while (mid <= high) {
            explanation.add(String.format("Step %d: low=%d, mid=%d, high=%d, nums[mid]=%d", 
                step++, low, mid, high, nums[mid]));
            
            switch (nums[mid]) {
                case 0: // Red
                    explanation.add(String.format("  nums[%d]=%d (Red) -> swap with low, low++, mid++", 
                        mid, nums[mid]));
                    swap(nums, low, mid);
                    low++;
                    mid++;
                    break;
                case 1: // White
                    explanation.add(String.format("  nums[%d]=%d (White) -> just mid++", mid, nums[mid]));
                    mid++;
                    break;
                case 2: // Blue
                    explanation.add(String.format("  nums[%d]=%d (Blue) -> swap with high, high--", 
                        mid, nums[mid]));
                    swap(nums, mid, high);
                    high--;
                    break;
            }
            
            explanation.add(String.format("  After operation: low=%d, mid=%d, high=%d, array=%s", 
                low, mid, high, Arrays.toString(nums)));
        }
        
        explanation.add("Final sorted array: " + Arrays.toString(nums));
        return new SortColorsResult(nums, explanation);
    }
    
    // Performance comparison
    public void comparePerformance(int[] nums, int trials) {
        System.out.println("=== Performance Comparison ===");
        System.out.println("Array: " + Arrays.toString(nums));
        System.out.println("Trials: " + trials);
        
        // Counting Sort
        long startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            sortColorsCounting(nums.clone());
        }
        long endTime = System.nanoTime();
        System.out.printf("Counting Sort: took %d ns\n", endTime - startTime);
        
        // Dutch National Flag
        startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            sortColorsDutchNationalFlag(nums.clone());
        }
        endTime = System.nanoTime();
        System.out.printf("Dutch National Flag: took %d ns\n", endTime - startTime);
        
        // Bucket Sort
        startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            sortColorsBucket(nums.clone());
        }
        endTime = System.nanoTime();
        System.out.printf("Bucket Sort: took %d ns\n", endTime - startTime);
    }
    
    // Stability analysis
    public void stabilityAnalysis(int[] nums) {
        System.out.println("=== Stability Analysis ===");
        System.out.println("Original: " + Arrays.toString(nums));
        
        // Create array with indices to track stability
        int[][] indexedArray = new int[nums.length][2];
        for (int i = 0; i < nums.length; i++) {
            indexedArray[i][0] = nums[i];
            indexedArray[i][1] = i;
        }
        
        // Sort using Dutch National Flag
        sortColorsDutchNationalFlag(nums);
        
        System.out.println("Sorted: " + Arrays.toString(nums));
        System.out.println("Note: Dutch National Flag is not stable for equal elements");
    }
    
    // Memory usage analysis
    public void memoryUsageAnalysis(int[] nums) {
        System.out.println("=== Memory Usage Analysis ===");
        System.out.println("Array size: " + nums.length);
        
        System.out.println("Counting Sort:");
        System.out.println("  Space: O(1) - only 3 counters needed");
        System.out.println("  Time: O(N) - single pass to count, single pass to reconstruct");
        
        System.out.println("Dutch National Flag:");
        System.out.println("  Space: O(1) - only 3 pointers");
        System.out.println("  Time: O(N) - single pass");
        
        System.out.println("Bucket Sort:");
        System.out.println("  Space: O(N+K) - buckets + output array");
        System.out.println("  Time: O(N+K) - distribution + collection");
        
        System.out.println("Radix Sort:");
        System.out.println("  Space: O(N+B) - counting array + output array");
        System.out.println("  Time: O(N*B) - B passes for B digits");
    }
    
    // Generalized counting sort for any range
    public int[] countingSortGeneral(int[] nums, int min, int max) {
        if (nums.length == 0) {
            return nums;
        }
        
        int range = max - min + 1;
        int[] count = new int[range];
        
        // Count occurrences
        for (int num : nums) {
            count[num - min]++;
        }
        
        // Reconstruct array
        int index = 0;
        for (int i = 0; i < range; i++) {
            while (count[i] > 0) {
                nums[index++] = i + min;
                count[i]--;
            }
        }
        
        return nums;
    }
    
    public static void main(String[] args) {
        SortColors sc = new SortColors();
        
        // Test cases
        int[][] testCases = {
            {2, 0, 2, 1, 1, 0},
            {2, 0, 1},
            {0},
            {1, 1, 1, 1, 1},
            {0, 1, 2, 0, 1, 2, 0, 1, 2},
            {2, 2, 2, 2, 2},
            {1, 0, 2, 1, 0, 2, 1, 0, 2, 1},
            {0, 0, 0, 0, 0},
            {2, 1, 0, 2, 0, 1, 2, 1, 0}
        };
        
        String[] descriptions = {
            "Mixed colors",
            "Simple case",
            "Single element",
            "All red",
            "Alternating pattern",
            "All blue",
            "Random pattern",
            "All white",
            "Complex pattern"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.println("Input: " + Arrays.toString(testCases[i]));
            
            int[] result1 = sc.sortColorsCounting(testCases[i].clone());
            int[] result2 = sc.sortColorsDutchNationalFlag(testCases[i].clone());
            int[] result3 = sc.sortColorsBucket(testCases[i].clone());
            int[] result4 = sc.sortColorsRadix(testCases[i].clone());
            
            System.out.println("Counting Sort: " + Arrays.toString(result1));
            System.out.println("Dutch National Flag: " + Arrays.toString(result2));
            System.out.println("Bucket Sort: " + Arrays.toString(result3));
            System.out.println("Radix Sort: " + Arrays.toString(result4));
            
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        SortColorsResult detailedResult = sc.sortColorsDetailed(new int[]{2, 0, 2, 1, 1, 0});
        
        System.out.println("Result: " + Arrays.toString(detailedResult.sortedArray));
        for (String step : detailedResult.explanation) {
            System.out.println(step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        int[] performanceArray = new int[1000];
        Random rand = new Random();
        for (int i = 0; i < 1000; i++) {
            performanceArray[i] = rand.nextInt(3);
        }
        
        sc.comparePerformance(performanceArray, 10000);
        
        // Stability analysis
        System.out.println("\n=== Stability Analysis ===");
        int[] stabilityArray = {2, 0, 2, 1, 1, 0};
        sc.stabilityAnalysis(stabilityArray);
        
        // Memory usage analysis
        System.out.println("\n=== Memory Usage Analysis ===");
        sc.memoryUsageAnalysis(new int[100]);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        
        // Empty array
        int[] emptyResult = sc.sortColorsCounting(new int[0]);
        System.out.println("Empty array: " + Arrays.toString(emptyResult));
        
        // Single element
        int[] singleResult = sc.sortColorsCounting(new int[]{1});
        System.out.println("Single element: " + Arrays.toString(singleResult));
        
        // Large array
        int[] largeArray = new int[10000];
        for (int i = 0; i < 10000; i++) {
            largeArray[i] = i % 3;
        }
        
        long startTime = System.nanoTime();
        sc.sortColorsDutchNationalFlag(largeArray);
        long endTime = System.nanoTime();
        System.out.printf("Large array (10000 elements): took %d ns\n", endTime - startTime);
        
        // Generalized counting sort
        System.out.println("\n=== Generalized Counting Sort ===");
        int[] generalArray = {5, 3, 7, 5, 3, 7};
        int[] generalResult = sc.countingSortGeneral(generalArray, 3, 7);
        System.out.println("General array: " + Arrays.toString(generalArray));
        System.out.println("Generalized result: " + Arrays.toString(generalResult));
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Non-Comparison Sorting
- **Counting Sort**: Count occurrences and reconstruct
- **Dutch National Flag**: Three-way partitioning
- **Radix Sort**: Digit-by-digit sorting
- **Bucket Sort**: Distribute into buckets

## 2. PROBLEM CHARACTERISTICS
- **Color Sorting**: Sort array with 0s, 1s, 2s
- **Limited Range**: Small range of values (0-2)
- **In-Place Sorting**: Sort without extra space
- **Linear Time**: Achieve O(N) time complexity

## 3. SIMILAR PROBLEMS
- Sort Characters by Frequency
- Sort Array By Parity
- Sort Binary Array
- Sort Colors II (extended range)

## 4. KEY OBSERVATIONS
- Counting sort works well for small range
- Dutch National Flag uses three pointers
- Radix sort extends counting to larger ranges
- Time complexity: O(N) vs O(N log N) comparison sort
- Space complexity: O(1) vs O(N) for counting sort

## 5. VARIATIONS & EXTENSIONS
- Extended color ranges
- Stability considerations
- Multiple passes for larger ranges
- Custom comparison functions

## 6. INTERVIEW INSIGHTS
- Clarify: "What is the value range?"
- Edge cases: empty array, single element, all same color
- Time complexity: O(N) vs O(N log N) comparison sort
- Space complexity: O(1) vs O(N) counting sort

## 7. COMMON MISTAKES
- Incorrect pointer management in Dutch National Flag
- Wrong counting sort range handling
- Incorrect radix sort digit processing
- Not handling edge cases properly
- Stability issues in counting sort

## 8. OPTIMIZATION STRATEGIES
- Use Dutch National Flag for O(1) space
- Implement counting sort for small ranges
- Use radix sort for larger ranges
- Efficient bucket distribution

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like sorting colored balls:**
- You have balls of three colors (red, white, blue)
- Need to sort them by color efficiently
- Counting sort: count each color, then place them
- Dutch National Flag: three pointers to partition in one pass
- Radix sort: sort by digits for larger ranges
- This is like organizing items by their properties

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array with values 0, 1, 2
2. **Goal**: Sort array in ascending order
3. **Output**: Sorted array

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(N log N) comparison sort
- **"How to optimize?"** → Use counting sort for small range
- **"Why Dutch National Flag?"** → Three-way partitioning in one pass
- **"Why counting sort?"** → Count occurrences, reconstruct

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use Dutch National Flag algorithm:
1. Three pointers: low (for 0s), mid (current), high (for 2s)
2. While mid <= high:
   - If nums[mid] == 0: swap with low, low++, mid++
   - If nums[mid] == 1: mid++ (already in correct position)
   - If nums[mid] == 2: swap with high, high--
3. This partitions array in one pass
4. For counting sort:
   - Count occurrences of 0, 1, 2
   - Place 0s, then 1s, then 2s"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return as is
- **Single element**: Already sorted
- **All same color**: Handle correctly
- **Large arrays**: Ensure O(N) time complexity

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Array: [2, 0, 2, 1, 1, 0]

Human thinking:
"Let's apply Dutch National Flag:

Initialize: low=0, mid=0, high=5

Step 1: mid=0, nums[0]=2 (Blue)
- nums[0] == 2: swap with high
- Swap nums[0] and nums[5]: [0, 0, 2, 1, 1, 2]
- high=4, mid stays 0

Step 2: mid=0, nums[0]=0 (Red)
- nums[0] == 0: swap with low
- Swap nums[0] and nums[0]: [0, 0, 2, 1, 1, 2] (no change)
- low=1, mid=1

Step 3: mid=1, nums[1]=0 (Red)
- nums[1] == 0: swap with low
- Swap nums[1] and nums[1]: [0, 0, 2, 1, 1, 2] (no change)
- low=2, mid=2

Step 4: mid=2, nums[2]=2 (Blue)
- nums[2] == 2: swap with high
- Swap nums[2] and nums[4]: [0, 0, 1, 1, 2, 2]
- high=3, mid stays 2

Step 5: mid=2, nums[2]=1 (White)
- nums[2] == 1: mid++ (already correct)
- mid=3

Step 6: mid=3, nums[3]=1 (White)
- nums[3] == 1: mid++ (already correct)
- mid=4

Step 7: mid=4, nums[4]=2 (Blue)
- nums[4] == 2: swap with high
- Swap nums[4] and nums[3]: [0, 0, 1, 2, 1, 2]
- high=2, mid stays 4

Now mid=4 > high=2, algorithm terminates

Final result: [0, 0, 1, 1, 2, 2] ✓

Manual verification:
All 0s at beginning, 1s in middle, 2s at end ✓
Sorted in ascending order ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Three pointers maintain correct partitions
- **Why it's efficient**: O(N) time vs O(N log N) comparison sort
- **Why it's correct**: Invariant maintains sorted partitions

### Common Human Pitfalls & How to Avoid Them
1. **"Why not use quicksort?"** → O(N log N) slower than O(N)
2. **"What about counting sort?"** → Uses O(N) space for range array
3. **"How to handle three colors?"** → Use three-way partitioning
4. **"What about stability?"** → Dutch National Flag is not stable

### Real-World Analogy
**Like sorting laundry by color:**
- You have a pile of clothes in three colors (red, white, blue)
- Need to sort them efficiently
- Dutch National Flag: three baskets, one for each color
- Pick each item and place in correct basket
- Counting sort: count each color, then place them
- This is like organizing items by their properties
- Useful in data processing, image processing, statistics
- Like sorting items into labeled bins

### Human-Readable Pseudocode
```
function dutchNationalFlag(nums):
    low = 0, mid = 0, high = nums.length - 1
    
    while mid <= high:
        if nums[mid] == 0:  // Red
            swap(nums, low, mid)
            low++
            mid++
        elif nums[mid] == 1:  // White
            mid++  // Already in correct position
        else:  // nums[mid] == 2, Blue
            swap(nums, mid, high)
            high--
    
    return nums

function countingSort(nums):
    count = [0, 0, 0]  // For colors 0, 1, 2
    
    // Count occurrences
    for num in nums:
        count[num]++
    
    // Reconstruct array
    index = 0
    for color in [0, 1, 2]:
        for i from 0 to count[color]-1:
            nums[index++] = color
    
    return nums
```

### Execution Visualization

### Example: nums=[2,0,2,1,1,0]
```
Dutch National Flag Process:

Initialize: low=0, mid=0, high=5
Array: [2,0,2,1,1,0]

Step 1: mid=0, nums[0]=2 (Blue)
- Swap with high: swap(0,5)
- Array: [0,0,2,1,1,2]
- high=4

Step 2: mid=0, nums[0]=0 (Red)
- Swap with low: swap(0,0)
- Array: [0,0,2,1,1,2]
- low=1, mid=1

Step 3: mid=1, nums[1]=0 (Red)
- Swap with low: swap(1,1)
- Array: [0,0,2,1,1,2]
- low=2, mid=2

Step 4: mid=2, nums[2]=2 (Blue)
- Swap with high: swap(2,4)
- Array: [0,0,1,1,2,2]
- high=3

Step 5: mid=2, nums[2]=1 (White)
- Already correct: mid=3
- Array: [0,0,1,1,2,2]

Continue...

Final result: [0,0,1,1,2,2] ✓

Visualization:
Three pointers maintain partitions
Reds at left, whites in middle, blues at right
Single pass achieves sorting ✓
```

### Key Visualization Points:
- **Three Pointers**: Low (0s), Mid (current), High (2s)
- **Partition Invariant**: [0...low-1] are 0s, [low...mid-1] are 1s, [mid+1...high] are 2s
- **Single Pass**: O(N) time complexity
- **In-Place**: O(1) extra space

### Memory Layout Visualization:
```
Dutch National Flag Evolution:

Initial: [2,0,2,1,1,0]
low=0, mid=0, high=5

After Step 1: [0,0,2,1,1,2]
low=1, mid=0, high=4
Partition: [0] | [0,2,1,1] | [2]

After Step 2: [0,0,2,1,1,2]
low=2, mid=1, high=4
Partition: [0,0] | [2,1,1] | [2]

After Step 3: [0,0,2,1,1,2]
low=2, mid=2, high=4
Partition: [0,0] | [2,1,1] | [2]

After Step 4: [0,0,1,1,2,2]
low=2, mid=3, high=3
Partition: [0,0] | [1,1] | [2,2]

Final: [0,0,1,1,2,2]
All 0s at left, 1s in middle, 2s at right ✓
```

### Time Complexity Breakdown:
- **Dutch National Flag**: O(N) time, O(1) space
- **Counting Sort**: O(N) time, O(1) space (for fixed range)
- **Radix Sort**: O(N*B) time, O(N+B) space
- **Bucket Sort**: O(N+K) time, O(N+K) space
- **Optimal**: Best possible for this problem
- **vs Comparison Sort**: O(N) vs O(N log N) time
*/
