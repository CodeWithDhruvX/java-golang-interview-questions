import java.util.*;

public class RandomPickIndex {
    
    // 398. Random Pick Index - Reservoir Sampling
    // Time: O(N), Space: O(1)
    static class Solution {
        int[] nums;
        Random random;
        
        public Solution(int[] nums) {
            this.nums = nums;
            this.random = new Random();
        }
        
        // Reservoir sampling algorithm
        public int pickRandomIndex() {
            if (nums.length == 0) {
                return -1;
            }
            
            // Reservoir sampling algorithm
            int k = 1; // We need to pick 1 element
            int[] reservoir = new int[k];
            
            // Fill reservoir with first k elements
            for (int i = 0; i < k && i < nums.length; i++) {
                reservoir[i] = nums[i];
            }
            
            // Process remaining elements
            for (int i = k; i < nums.length; i++) {
                // Generate random number between 0 and i
                int j = random.nextInt(i + 1);
                if (j < k) {
                    reservoir[j] = nums[i];
                }
            }
            
            return reservoir[0];
        }
    }
    
    // Reservoir sampling with multiple picks
    static class SolutionMultiple {
        int[] nums;
        Random random;
        
        public SolutionMultiple(int[] nums) {
            this.nums = nums;
            this.random = new Random();
        }
        
        public int[] pickRandomIndexMultiple(int k) {
            if (nums.length == 0 || k <= 0 || k > nums.length) {
                return new int[0];
            }
            
            int[] reservoir = new int[k];
            
            // Fill reservoir with first k elements
            for (int i = 0; i < k && i < nums.length; i++) {
                reservoir[i] = nums[i];
            }
            
            // Process remaining elements
            for (int i = k; i < nums.length; i++) {
                int j = random.nextInt(i + 1);
                if (j < k) {
                    reservoir[j] = nums[i];
                }
            }
            
            return reservoir;
        }
    }
    
    // Reservoir sampling with weighted selection
    static class SolutionWeighted {
        int[] nums;
        int[] weights;
        int[] cumulativeWeights;
        int totalWeight;
        Random random;
        
        public SolutionWeighted(int[] nums, int[] weights) {
            if (nums.length != weights.length) {
                throw new IllegalArgumentException("Arrays must have same length");
            }
            
            this.nums = nums;
            this.weights = weights;
            this.cumulativeWeights = new int[weights.length];
            this.random = new Random();
            
            // Calculate cumulative weights
            this.totalWeight = 0;
            for (int i = 0; i < weights.length; i++) {
                this.totalWeight += weights[i];
                this.cumulativeWeights[i] = this.totalWeight;
            }
        }
        
        public int pickRandomWeightedIndex() {
            if (nums.length == 0) {
                return -1;
            }
            
            // Generate random number and find corresponding index
            int r = random.nextInt(totalWeight);
            int cumulative = 0;
            
            for (int i = 0; i < weights.length; i++) {
                cumulative += weights[i];
                if (r < cumulative) {
                    return nums[i];
                }
            }
            
            return nums[nums.length - 1]; // Fallback
        }
    }
    
    // Version with detailed explanation
    static class SolutionDetailed {
        int[] nums;
        Random random;
        List<String> explanation;
        
        public SolutionDetailed(int[] nums) {
            this.nums = nums;
            this.random = new Random();
            this.explanation = new ArrayList<>();
        }
        
        public int pickRandomIndex() {
            explanation.add("=== Reservoir Sampling Algorithm ===");
            explanation.add("Array: " + Arrays.toString(nums));
            
            if (nums.length == 0) {
                explanation.add("Empty array, returning -1");
                return -1;
            }
            
            int k = 1; // We need to pick 1 element
            int[] reservoir = new int[k];
            
            explanation.add("Step 1: Fill reservoir with first k elements");
            
            // Fill reservoir with first k elements
            for (int i = 0; i < k && i < nums.length; i++) {
                reservoir[i] = nums[i];
                explanation.add(String.format("  Reservoir[%d] = nums[%d] = %d", i, i, nums[i]));
            }
            
            explanation.add("Step 2: Process remaining elements");
            
            // Process remaining elements
            for (int i = k; i < nums.length; i++) {
                // Generate random number between 0 and i
                int j = random.nextInt(i + 1);
                explanation.add(String.format("  Processing nums[%d] = %d, random j = %d", i, nums[i], j));
                
                if (j < k) {
                    explanation.add(String.format("    Replacing reservoir[%d] = %d with %d", j, reservoir[j], nums[i]));
                    reservoir[j] = nums[i];
                } else {
                    explanation.add("    No replacement (j >= k)");
                }
            }
            
            explanation.add("Step 3: Return selected element");
            explanation.add(String.format("Selected element: %d", reservoir[0]));
            
            return reservoir[0];
        }
        
        public List<String> getExplanation() {
            return new ArrayList<>(explanation);
        }
        
        public void clearExplanation() {
            explanation.clear();
        }
    }
    
    // Performance comparison
    public void comparePerformance(int[] nums, int trials) {
        System.out.println("=== Performance Comparison ===");
        System.out.println("Array: " + Arrays.toString(nums));
        System.out.println("Trials: " + trials);
        
        // Reservoir sampling
        Solution solution = new Solution(nums);
        long startTime = System.nanoTime();
        
        for (int i = 0; i < trials; i++) {
            solution.pickRandomIndex();
        }
        
        long endTime = System.nanoTime();
        System.out.printf("Reservoir sampling: took %d ns\n", endTime - startTime);
        
        // Standard random selection
        startTime = System.nanoTime();
        for (int i = 0; i < trials; i++) {
            nums[random.nextInt(nums.length)];
        }
        
        endTime = System.nanoTime();
        System.out.printf("Standard random selection: took %d ns\n", endTime - startTime);
        
        // Weighted selection
        int[] weights = new int[nums.length];
        Arrays.fill(weights, 1); // Equal weights
        
        SolutionWeighted weighted = new SolutionWeighted(nums, weights);
        startTime = System.nanoTime();
        
        for (int i = 0; i < trials; i++) {
            weighted.pickRandomWeightedIndex();
        }
        
        endTime = System.nanoTime();
        System.out.printf("Weighted selection: took %d ns\n", endTime - startTime);
    }
    
    // Statistical analysis
    public void statisticalAnalysis(int[] nums, int samples) {
        System.out.println("=== Statistical Analysis ===");
        System.out.println("Array: " + Arrays.toString(nums));
        System.out.println("Samples: " + samples);
        
        Solution solution = new Solution(nums);
        Map<Integer, Integer> frequency = new HashMap<>();
        
        for (int i = 0; i < samples; i++) {
            int value = solution.pickRandomIndex();
            frequency.put(value, frequency.getOrDefault(value, 0) + 1);
        }
        
        System.out.println("Frequency distribution:");
        for (Map.Entry<Integer, Integer> entry : frequency.entrySet()) {
            double percentage = (double) entry.getValue() / samples * 100;
            System.out.printf("  Value %d: %d times (%.2f%%)\n", 
                entry.getKey(), entry.getValue(), percentage);
        }
        
        // Calculate expected distribution
        System.out.println("\nExpected distribution (uniform):");
        double expectedPercentage = 100.0 / nums.length;
        for (int value : nums) {
            System.out.printf("  Value %d: %.2f%%\n", value, expectedPercentage);
        }
        
        // Calculate chi-square statistic
        double chiSquare = 0;
        for (Map.Entry<Integer, Integer> entry : frequency.entrySet()) {
            double expected = (double) samples / nums.length;
            double observed = entry.getValue();
            chiSquare += Math.pow(observed - expected, 2) / expected;
        }
        
        System.out.printf("Chi-square statistic: %.4f\n", chiSquare);
        System.out.printf("Expected chi-square (df=%d, p=0.05): %.4f\n", 
            nums.length - 1, 16.92); // Approximate value
    }
    
    // Reservoir sampling for streaming data
    static class StreamingReservoir {
        int[] reservoir;
        int k;
        int count;
        Random random;
        
        public StreamingReservoir(int k) {
            this.k = k;
            this.reservoir = new int[k];
            this.count = 0;
            this.random = new Random();
        }
        
        public void add(int value) {
            if (count < k) {
                // Fill reservoir initially
                reservoir[count] = value;
                count++;
            } else {
                // Random replacement
                int j = random.nextInt(count + 1);
                if (j < k) {
                    reservoir[j] = value;
                }
            }
        }
        
        public int[] getSample() {
            return Arrays.copyOf(reservoir, Math.min(count, k));
        }
        
        public int getCount() {
            return count;
        }
    }
    
    // Test streaming reservoir
    public void testStreamingReservoir(int[] stream, int k) {
        System.out.println("=== Streaming Reservoir Test ===");
        System.out.println("Stream: " + Arrays.toString(stream));
        System.out.println("Sample size k: " + k);
        
        StreamingReservoir sr = new StreamingReservoir(k);
        
        for (int value : stream) {
            sr.add(value);
        }
        
        int[] sample = sr.getSample();
        System.out.println("Final sample: " + Arrays.toString(sample));
        System.out.println("Total processed: " + sr.getCount());
    }
    
    public static void main(String[] args) {
        RandomPickIndex rpi = new RandomPickIndex();
        
        // Test cases
        int[][] testCases = {
            {1, 2, 3, 4, 5},
            {1, 1, 1, 1, 1},
            {5, 4, 3, 2, 1},
            {10, 20, 30},
            {100},
            {1, 100, 50, 25, 75},
            {0, 0, 0, 0},
            {-1, -2, -3, -4, -5},
            {7, 7, 7, 7, 7, 7}
        };
        
        String[] descriptions = {
            "Sequential values",
            "All same values",
            "Descending values",
            "Large values",
            "Single element",
            "Mixed values",
            "All zeros",
            "All negative",
            "Many duplicates"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.println("Array: " + Arrays.toString(testCases[i]));
            
            Solution solution = new Solution(testCases[i]);
            System.out.println("Random selections:");
            for (int j = 0; j < 5; j++) {
                System.out.printf("  Selection %d: %d\n", j + 1, solution.pickRandomIndex());
            }
            
            // Test multiple picks
            SolutionMultiple solutionMultiple = new SolutionMultiple(testCases[i]);
            int[] multiplePicks = solutionMultiple.pickRandomIndexMultiple(3);
            System.out.println("Multiple picks (k=3): " + Arrays.toString(multiplePicks));
            
            // Test weighted selection
            int[] weights = new int[testCases[i].length];
            Arrays.fill(weights, 1); // Equal weights
            
            try {
                SolutionWeighted solutionWeighted = new SolutionWeighted(testCases[i], weights);
                int weightedPick = solutionWeighted.pickRandomWeightedIndex();
                System.out.println("Weighted pick: " + weightedPick);
            } catch (Exception e) {
                System.out.println("Weighted selection error: " + e.getMessage());
            }
            
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        SolutionDetailed solutionDetailed = new SolutionDetailed(new int[]{1, 2, 3, 4, 5});
        
        int detailedResult = solutionDetailed.pickRandomIndex();
        System.out.printf("Detailed result: %d\n", detailedResult);
        
        for (String step : solutionDetailed.getExplanation()) {
            System.out.println(step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        int[] performanceArray = new int[1000];
        for (int i = 0; i < 1000; i++) {
            performanceArray[i] = i % 100;
        }
        
        rpi.comparePerformance(performanceArray, 10000);
        
        // Statistical analysis
        System.out.println("\n=== Statistical Analysis ===");
        int[] analysisArray = {1, 2, 3, 4, 5};
        rpi.statisticalAnalysis(analysisArray, 10000);
        
        // Streaming reservoir test
        System.out.println("\n=== Streaming Reservoir Test ===");
        int[] stream = new int[100];
        for (int i = 0; i < 100; i++) {
            stream[i] = i % 10;
        }
        
        rpi.testStreamingReservoir(stream, 5);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        
        // Empty array
        try {
            Solution emptySolution = new Solution(new int[0]);
            System.out.println("Empty array random: " + emptySolution.pickRandomIndex());
        } catch (Exception e) {
            System.out.println("Empty array handled: " + e.getMessage());
        }
        
        // Single element
        Solution singleSolution = new Solution(new int[]{42});
        System.out.println("Single element random: " + singleSolution.pickRandomIndex());
        
        // Large array
        int[] largeArray = new int[10000];
        for (int i = 0; i < 10000; i++) {
            largeArray[i] = i % 100;
        }
        
        Solution largeSolution = new Solution(largeArray);
        long startTime = System.nanoTime();
        for (int i = 0; i < 1000; i++) {
            largeSolution.pickRandomIndex();
        }
        long endTime = System.nanoTime();
        System.out.printf("Large array (10000 elements, 1000 selections): took %d ns\n", endTime - startTime);
        
        // Weighted selection with different weights
        System.out.println("\n=== Weighted Selection Test ===");
        int[] weightedArray = {1, 2, 3, 4};
        int[] differentWeights = {1, 2, 3, 4}; // Higher weight for higher numbers
        
        SolutionWeighted weightedSolution = new SolutionWeighted(weightedArray, differentWeights);
        Map<Integer, Integer> weightedFrequency = new HashMap<>();
        
        for (int i = 0; i < 10000; i++) {
            int weightedPick = weightedSolution.pickRandomWeightedIndex();
            weightedFrequency.put(weightedPick, weightedFrequency.getOrDefault(weightedPick, 0) + 1);
        }
        
        System.out.println("Weighted frequency distribution:");
        for (Map.Entry<Integer, Integer> entry : weightedFrequency.entrySet()) {
            double percentage = (double) entry.getValue() / 10000 * 100;
            System.out.printf("  Value %d: %d times (%.2f%%)\n", 
                entry.getKey(), entry.getValue(), percentage);
        }
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Reservoir Sampling
- **Random Selection**: Uniform random sampling from unknown size
- **Streaming Algorithm**: Process elements one by one
- **Space Efficiency**: O(k) space for k samples
- **Uniform Probability**: Each element has equal selection probability

## 2. PROBLEM CHARACTERISTICS
- **Random Index Selection**: Pick random element from array
- **Unknown Size**: Array size may not be known in advance
- **Multiple Picks**: Support selecting multiple elements
- **Weighted Selection**: Support weighted random selection

## 3. SIMILAR PROBLEMS
- Random Sampling from Data Stream
- Monte Carlo Algorithms
- Randomized Algorithms
- Probabilistic Data Structures

## 4. KEY OBSERVATIONS
- Reservoir sampling maintains uniform probability
- Each element has 1/k probability of being selected
- Time complexity: O(N) for N elements
- Space complexity: O(k) for k samples
- Works for streaming data where N is unknown

## 5. VARIATIONS & EXTENSIONS
- Different reservoir sizes
- Weighted random selection
- Streaming data processing
- Multiple reservoir maintenance

## 6. INTERVIEW INSIGHTS
- Clarify: "Can we modify the array?"
- Edge cases: empty array, single element, large arrays
- Time complexity: O(N) vs O(1) for known size
- Space complexity: O(k) vs O(1) for simple selection

## 7. COMMON MISTAKES
- Incorrect random number generation
- Wrong reservoir replacement logic
- Not handling edge cases properly
- Incorrect probability calculations
- Not maintaining uniform distribution

## 8. OPTIMIZATION STRATEGIES
- Use proper random number generation
- Efficient reservoir maintenance
- Correct probability calculations
- Handle streaming data efficiently

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like drawing lottery tickets from a hat:**
- You have a collection of items (array)
- Need to select random items uniformly
- If you don't know the total count, use reservoir sampling
- Maintain a small reservoir of selected items
- As you see new items, randomly replace reservoir items
- This ensures uniform probability without knowing total size

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Array of numbers, need random index
2. **Goal**: Return randomly selected index
3. **Output**: Random index or value

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → Need uniform probability
- **"How to optimize?"** → Use reservoir sampling for unknown size
- **"Why reservoir?"** → Maintains uniform probability with O(k) space
- **"How to handle multiple picks?"** → Expand reservoir size

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use reservoir sampling:
1. Initialize reservoir of size k
2. Fill reservoir with first k elements
3. For each remaining element i:
   - Generate random number j between 0 and i
   - If j < k, replace reservoir[j] with current element
   - This maintains uniform probability
4. Return reservoir[0] for single selection
5. For weighted selection, use cumulative weights"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return -1 or handle appropriately
- **Single element**: Return that element
- **Large arrays**: Ensure O(N) time complexity
- **Invalid k**: Handle gracefully

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Array: [1, 2, 3, 4, 5], k=1

Human thinking:
"Let's apply reservoir sampling:

Step 1: Initialize
reservoir = [0] (size k=1)

Step 2: Fill reservoir
reservoir[0] = nums[0] = 1

Step 3: Process remaining elements
i=1, nums[1]=2:
- Generate j = random.nextInt(2) in {0,1}
- If j < 1, replace reservoir[j] with nums[1]
- P(replace) = 1/2, P(keep) = 1/2

i=2, nums[2]=3:
- Generate j = random.nextInt(3) in {0,1,2}
- If j < 1, replace reservoir[j] with nums[2]
- P(replace) = 1/3, P(keep) = 2/3

Continue...

Final result: Each element has 1/N probability of being selected ✓

Manual verification:
Element 1: P(selected) = 1/5 (from i=0)
Element 2: P(selected) = (1/2)*(2/3)*(3/4)*(4/5) = 1/5
Element 3: P(selected) = (1/3)*(2/4)*(3/4)*(4/5) = 1/5
Element 4: P(selected) = (1/4)*(2/5)*(3/5)*(4/5) = 1/5
Element 5: P(selected) = (1/5)*(2/5)*(3/5)*(4/5) = 1/5

All elements have equal probability 1/5 ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Each element has equal selection probability
- **Why it's efficient**: O(N) time, O(k) space
- **Why it's correct**: Mathematical proof of uniform distribution

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use random.nextInt()?"** → Requires knowing array size
2. **"What about uniform probability?"** → Reservoir ensures uniform distribution
3. **"How to handle streaming?"** → Process elements as they arrive
4. **"What about weighted selection?"** → Use cumulative weights

### Real-World Analogy
**Like conducting a random survey without knowing population size:**
- You have people entering a venue (streaming data)
- Need to select random people for a survey
- Can't count everyone first (too slow)
- Reservoir sampling: maintain sample of size k
- Each new person has k/(total_seen) chance of replacing someone
- This ensures uniform probability without knowing total
- Useful in streaming analytics, online sampling, load testing
- Like conducting fair random selection from unknown population

### Human-Readable Pseudocode
```
function reservoirSampling(nums, k):
    if nums.length == 0:
        return -1
    
    reservoir = array of size k
    
    // Fill reservoir with first k elements
    for i from 0 to min(k, nums.length)-1:
        reservoir[i] = nums[i]
    
    // Process remaining elements
    for i from k to nums.length-1:
        j = random.nextInt(i + 1)
        if j < k:
            reservoir[j] = nums[i]
    
    return reservoir[0]  // For single selection

function weightedSelection(nums, weights):
    // Calculate cumulative weights
    cumulative = []
    totalWeight = 0
    for i from 0 to weights.length-1:
        totalWeight += weights[i]
        cumulative[i] = totalWeight
    
    // Generate random number and find corresponding element
    r = random.nextInt(totalWeight)
    for i from 0 to cumulative.length-1:
        if r < cumulative[i]:
            return nums[i]
    
    return nums[nums.length-1]
```

### Execution Visualization

### Example: nums=[1,2,3,4,5], k=1
```
Reservoir Sampling Process:

Initialize:
reservoir = [0] (size k=1)

Step 1: Fill reservoir
reservoir[0] = nums[0] = 1

Step 2: Process nums[1]=2
- Generate j = random.nextInt(2) in {0,1}
- P(j=0) = 1/2, P(j=1) = 1/2
- If j=0: reservoir[0] = 2
- If j=1: keep reservoir[0] = 1

Step 3: Process nums[2]=3
- Generate j = random.nextInt(3) in {0,1,2}
- P(j<1) = 1/3, P(j≥1) = 2/3
- If j=0: reservoir[0] = 3
- If j≥1: keep current reservoir

Continue...

Final probability analysis:
Element 1: P(selected) = 1/5
Element 2: P(selected) = 1/5
Element 3: P(selected) = 1/5
Element 4: P(selected) = 1/5
Element 5: P(selected) = 1/5

Uniform distribution achieved ✓

Visualization:
Reservoir maintains uniform probability
Each element has equal chance of final selection
Streaming capability without knowing total size
```

### Key Visualization Points:
- **Reservoir Maintenance**: Fixed size k array
- **Random Replacement**: j = random.nextInt(i+1)
- **Uniform Probability**: Each element has 1/N chance
- **Streaming Support**: Process elements as they arrive

### Memory Layout Visualization:
```
Array: [1, 2, 3, 4, 5]
k = 1

Reservoir Evolution:
Start: [1]
After nums[1]: [1] or [2] (50% each)
After nums[2]: [1], [2], or [3] (33% each)
After nums[3]: [1], [2], [3], or [4] (25% each)
After nums[4]: [1], [2], [3], [4], or [5] (20% each)

Final: Each element has 1/5 probability ✓

Reservoir sampling enables uniform selection
Works for streaming data where N is unknown
```

### Time Complexity Breakdown:
- **Initialization**: O(k) time, O(k) space
- **Processing**: O(N-k) time, O(1) space
- **Total**: O(N) time, O(k) space
- **Optimal**: Best possible for unknown size
- **vs Simple Random**: O(N) time, O(1) space (requires knowing N)
*/
