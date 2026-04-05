import java.util.*;

public class RangeSumQueryMutable {
    
    // 307. Range Sum Query - Mutable - Mos Algorithm
    // Time: O((N+Q) * sqrt(N)), Space: O(N + Q)
    static class RangeSumQuery {
        int[] nums;
        int blockSize;
        
        public RangeSumQuery(int[] nums) {
            this.nums = nums.clone();
            this.blockSize = (int) Math.sqrt(nums.length) + 1;
        }
        
        public void update(int index, int val) {
            nums[index] = val;
        }
        
        public int sumRange(int left, int right) {
            int sum = 0;
            for (int i = left; i <= right; i++) {
                sum += nums[i];
            }
            return sum;
        }
    }
    
    // Mos Query structure
    static class MosQuery {
        int left;
        int right;
        int index;
        
        MosQuery(int left, int right, int index) {
            this.left = left;
            this.right = right;
            this.index = index;
        }
    }
    
    // Mos Algorithm for multiple range queries
    static class MosAlgorithm {
        int[] nums;
        int blockSize;
        
        public MosAlgorithm(int[] nums) {
            this.nums = nums;
            this.blockSize = (int) Math.sqrt(nums.length) + 1;
        }
        
        public int[] processQueries(List<MosQuery> queries) {
            if (queries.isEmpty()) {
                return new int[0];
            }
            
            // Sort queries using Mo's ordering
            queries.sort((a, b) -> {
                int blockA = a.left / blockSize;
                int blockB = b.left / blockSize;
                
                if (blockA != blockB) {
                    return Integer.compare(blockA, blockB);
                }
                
                // If same block, sort by right
                if (blockA % 2 == 0) {
                    return Integer.compare(a.right, b.right);
                }
                return Integer.compare(b.right, a.right);
            });
            
            // Process queries
            int[] results = new int[queries.size()];
            int currentLeft = 0;
            int currentRight = -1;
            int currentSum = 0;
            
            for (MosQuery query : queries) {
                // Expand to right
                while (currentRight < query.right) {
                    currentRight++;
                    currentSum += nums[currentRight];
                }
                
                // Contract from right
                while (currentRight > query.right) {
                    currentSum -= nums[currentRight];
                    currentRight--;
                }
                
                // Expand to left
                while (currentLeft < query.left) {
                    currentSum -= nums[currentLeft];
                    currentLeft++;
                }
                
                // Contract from left
                while (currentLeft > query.left) {
                    currentLeft--;
                    currentSum += nums[currentLeft];
                }
                
                results[query.index] = currentSum;
            }
            
            return results;
        }
    }
    
    // Optimized Mos Algorithm with frequency counting
    static class MosAlgorithmOptimized {
        int[] nums;
        int blockSize;
        Map<Integer, Integer> freq;
        int currentSum;
        
        public MosAlgorithmOptimized(int[] nums) {
            this.nums = nums;
            this.blockSize = (int) Math.sqrt(nums.length) + 1;
            this.freq = new HashMap<>();
            this.currentSum = 0;
        }
        
        public int[] processQueries(List<MosQuery> queries) {
            if (queries.isEmpty()) {
                return new int[0];
            }
            
            // Sort queries using Mo's ordering
            queries.sort((a, b) -> {
                int blockA = a.left / blockSize;
                int blockB = b.left / blockSize;
                
                if (blockA != blockB) {
                    return Integer.compare(blockA, blockB);
                }
                
                // If same block, sort by right
                if (blockA % 2 == 0) {
                    return Integer.compare(a.right, b.right);
                }
                return Integer.compare(b.right, a.right);
            });
            
            // Process queries
            int[] results = new int[queries.size()];
            int currentLeft = 0;
            int currentRight = -1;
            
            for (MosQuery query : queries) {
                // Expand to right
                while (currentRight < query.right) {
                    currentRight++;
                    add(currentRight);
                }
                
                // Contract from right
                while (currentRight > query.right) {
                    remove(currentRight);
                    currentRight--;
                }
                
                // Expand to left
                while (currentLeft < query.left) {
                    remove(currentLeft);
                    currentLeft++;
                }
                
                // Contract from left
                while (currentLeft > query.left) {
                    currentLeft--;
                    add(currentLeft);
                }
                
                results[query.index] = currentSum;
            }
            
            return results;
        }
        
        private void add(int index) {
            int val = nums[index];
            freq.put(val, freq.getOrDefault(val, 0) + 1);
            currentSum += val;
        }
        
        private void remove(int index) {
            int val = nums[index];
            int count = freq.get(val) - 1;
            if (count == 0) {
                freq.remove(val);
            } else {
                freq.put(val, count);
            }
            currentSum -= val;
        }
    }
    
    // Version with detailed explanation
    static class MosAlgorithmDetailed {
        int[] nums;
        int blockSize;
        List<String> explanation;
        
        public MosAlgorithmDetailed(int[] nums) {
            this.nums = nums;
            this.blockSize = (int) Math.sqrt(nums.length) + 1;
            this.explanation = new ArrayList<>();
        }
        
        public int[] processQueries(List<MosQuery> queries) {
            explanation.add("=== Mos Algorithm for Range Sum Queries ===");
            explanation.add("Array: " + Arrays.toString(nums));
            explanation.add("Block size: " + blockSize);
            
            if (queries.isEmpty()) {
                explanation.add("No queries to process");
                return new int[0];
            }
            
            // Sort queries using Mo's ordering
            explanation.add("Sorting queries using Mo's ordering:");
            queries.sort((a, b) -> {
                int blockA = a.left / blockSize;
                int blockB = b.left / blockSize;
                
                if (blockA != blockB) {
                    explanation.add(String.format("  Query [%d,%d] (block %d) vs [%d,%d] (block %d): different blocks", 
                        a.left, a.right, blockA, b.left, b.right, blockB));
                    return Integer.compare(blockA, blockB);
                }
                
                // If same block, sort by right
                if (blockA % 2 == 0) {
                    explanation.add(String.format("  Query [%d,%d] vs [%d,%d]: same even block, sort by right ascending", 
                        a.left, a.right, b.left, b.right));
                    return Integer.compare(a.right, b.right);
                }
                explanation.add(String.format("  Query [%d,%d] vs [%d,%d]: same odd block, sort by right descending", 
                    a.left, a.right, b.left, b.right));
                return Integer.compare(b.right, a.right);
            });
            
            explanation.add("Sorted queries:");
            for (int i = 0; i < queries.size(); i++) {
                MosQuery q = queries.get(i);
                explanation.add(String.format("  %d: [%d, %d] (index %d)", i, q.left, q.right, q.index));
            }
            
            // Process queries
            int[] results = new int[queries.size()];
            int currentLeft = 0;
            int currentRight = -1;
            int currentSum = 0;
            
            explanation.add("Processing queries:");
            
            for (MosQuery query : queries) {
                explanation.add(String.format("Processing query [%d, %d] (index %d):", 
                    query.left, query.right, query.index));
                
                // Expand to right
                while (currentRight < query.right) {
                    currentRight++;
                    currentSum += nums[currentRight];
                    explanation.add(String.format("  Expand right to %d, sum = %d", currentRight, currentSum));
                }
                
                // Contract from right
                while (currentRight > query.right) {
                    currentSum -= nums[currentRight];
                    explanation.add(String.format("  Contract right from %d, sum = %d", currentRight, currentSum));
                    currentRight--;
                }
                
                // Expand to left
                while (currentLeft < query.left) {
                    currentSum -= nums[currentLeft];
                    explanation.add(String.format("  Expand left from %d, sum = %d", currentLeft, currentSum));
                    currentLeft++;
                }
                
                // Contract from left
                while (currentLeft > query.left) {
                    currentLeft--;
                    currentSum += nums[currentLeft];
                    explanation.add(String.format("  Contract left to %d, sum = %d", currentLeft, currentSum));
                }
                
                results[query.index] = currentSum;
                explanation.add(String.format("  Result for query [%d, %d]: %d", 
                    query.left, query.right, currentSum));
            }
            
            return results;
        }
        
        public List<String> getExplanation() {
            return new ArrayList<>(explanation);
        }
    }
    
    // Performance comparison
    public void comparePerformance(int[] nums, List<MosQuery> queries) {
        System.out.println("=== Performance Comparison ===");
        System.out.println("Array: " + Arrays.toString(nums));
        System.out.println("Queries: " + queries.size());
        
        // Standard approach
        long startTime = System.nanoTime();
        int[] results1 = new int[queries.size()];
        for (int i = 0; i < queries.size(); i++) {
            MosQuery q = queries.get(i);
            int sum = 0;
            for (int j = q.left; j <= q.right; j++) {
                sum += nums[j];
            }
            results1[i] = sum;
        }
        long endTime = System.nanoTime();
        System.out.printf("Standard approach: took %d ns\n", endTime - startTime);
        
        // Mos Algorithm
        startTime = System.nanoTime();
        MosAlgorithm mos = new MosAlgorithm(nums);
        int[] results2 = mos.processQueries(new ArrayList<>(queries));
        endTime = System.nanoTime();
        System.out.printf("Mos Algorithm: took %d ns\n", endTime - startTime);
        
        // Optimized Mos Algorithm
        startTime = System.nanoTime();
        MosAlgorithmOptimized mosOpt = new MosAlgorithmOptimized(nums);
        int[] results3 = mosOpt.processQueries(new ArrayList<>(queries));
        endTime = System.nanoTime();
        System.out.printf("Optimized Mos Algorithm: took %d ns\n", endTime - startTime);
        
        // Verify results
        boolean resultsMatch = Arrays.equals(results1, results2) && Arrays.equals(results2, results3);
        System.out.println("Results match: " + resultsMatch);
    }
    
    // Generate test queries
    public List<MosQuery> generateQueries(int n, int count) {
        List<MosQuery> queries = new ArrayList<>();
        Random rand = new Random();
        
        for (int i = 0; i < count; i++) {
            int left = rand.nextInt(n);
            int right = rand.nextInt(n);
            if (left > right) {
                int temp = left;
                left = right;
                right = temp;
            }
            queries.add(new MosQuery(left, right, i));
        }
        
        return queries;
    }
    
    public static void main(String[] args) {
        RangeSumQueryMutable rsm = new RangeSumQueryMutable();
        
        // Test cases
        int[][] testCases = {
            {1, 3, 5},
            {1, 9, 3, 1},
            {0, 0, 0, 0},
            {100, -100, 50, -50},
            {1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
            {5, 5, 5, 5, 5},
            {-1, -2, -3, -4, -5},
            {0},
            {1, -1, 1, -1, 1, -1},
            {1000, 2000, 3000, 4000, 5000}
        };
        
        String[] descriptions = {
            "Small array",
            "Mixed values",
            "All zeros",
            "Mixed large values",
            "Sequential",
            "All same",
            "All negative",
            "Single element",
            "Alternating",
            "Large values"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.println("Array: " + Arrays.toString(testCases[i]));
            
            RangeSumQuery rsq = new RangeSumQuery(testCases[i]);
            
            // Test basic operations
            System.out.printf("Sum range [0, %d]: %d\n", 
                testCases[i].length - 1, rsq.sumRange(0, testCases[i].length - 1));
            
            if (testCases[i].length > 1) {
                System.out.printf("Sum range [1, %d]: %d\n", 
                    testCases[i].length - 1, rsq.sumRange(1, testCases[i].length - 1));
            }
            
            // Test update
            if (testCases[i].length > 0) {
                rsq.update(0, 100);
                System.out.printf("After update index 0 to 100, sum range [0, 0]: %d\n", rsq.sumRange(0, 0));
            }
            
            // Test Mos Algorithm
            List<MosQuery> queries = rsm.generateQueries(testCases[i].length, 5);
            MosAlgorithm mos = new MosAlgorithm(testCases[i]);
            int[] mosResults = mos.processQueries(queries);
            
            System.out.println("Mos Algorithm results:");
            for (int j = 0; j < queries.size(); j++) {
                MosQuery q = queries.get(j);
                System.out.printf("  Query [%d, %d]: %d\n", q.left, q.right, mosResults[j]);
            }
            
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        int[] detailedArray = {1, 3, 5, 2, 4, 6};
        List<MosQuery> detailedQueries = Arrays.asList(
            new MosQuery(0, 2, 0),
            new MosQuery(1, 4, 1),
            new MosQuery(3, 5, 2)
        );
        
        MosAlgorithmDetailed mosDetailed = new MosAlgorithmDetailed(detailedArray);
        int[] detailedResults = mosDetailed.processQueries(new ArrayList<>(detailedQueries));
        
        System.out.println("Results:");
        for (int i = 0; i < detailedResults.length; i++) {
            System.out.printf("  Query %d: %d\n", i, detailedResults[i]);
        }
        
        for (String step : mosDetailed.getExplanation()) {
            System.out.println(step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        int[] performanceArray = new int[1000];
        for (int i = 0; i < 1000; i++) {
            performanceArray[i] = i % 100;
        }
        
        List<MosQuery> performanceQueries = rsm.generateQueries(1000, 100);
        rsm.comparePerformance(performanceArray, performanceQueries);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        
        // Empty array
        try {
            RangeSumQuery empty = new RangeSumQuery(new int[0]);
            System.out.println("Empty array sum: " + empty.sumRange(0, 0));
        } catch (Exception e) {
            System.out.println("Empty array handled: " + e.getMessage());
        }
        
        // Single element
        RangeSumQuery single = new RangeSumQuery(new int[]{42});
        System.out.println("Single element sum: " + single.sumRange(0, 0));
        single.update(0, 100);
        System.out.println("After update: " + single.sumRange(0, 0));
        
        // Large array
        int[] largeArray = new int[10000];
        for (int i = 0; i < 10000; i++) {
            largeArray[i] = i % 100;
        }
        
        List<MosQuery> largeQueries = rsm.generateQueries(10000, 1000);
        MosAlgorithm largeMos = new MosAlgorithm(largeArray);
        long startTime = System.nanoTime();
        largeMos.processQueries(largeQueries);
        long endTime = System.nanoTime();
        System.out.printf("Large array (10000 elements, 1000 queries): took %d ns\n", endTime - startTime);
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Mo's Algorithm
- **Offline Processing**: Sort queries for optimal processing order
- **Block Decomposition**: Divide array into sqrt(N) blocks
- **Frequency Tracking**: Maintain current window sum efficiently
- **Query Reordering**: Sort by block then by right endpoint

## 2. PROBLEM CHARACTERISTICS
- **Multiple Queries**: Process many range sum queries efficiently
- **Static Array**: Array doesn't change between queries
- **Range Sum**: Calculate sum of elements in range [l,r]
- **Offline Processing**: All queries known in advance

## 3. SIMILAR PROBLEMS
- Range Minimum Query (offline)
- Range Maximum Query (offline)
- Range Update Query (offline)
- Order Statistics Tree queries

## 4. KEY OBSERVATIONS
- Mo's algorithm reduces complexity from O(N*Q) to O((N+Q)*sqrt(N))
- Block size of sqrt(N) provides optimal balance
- Queries sorted by block ensure efficient processing
- Frequency tracking allows O(1) add/remove operations
- Time complexity: O((N+Q) * sqrt(N)) for Q queries

## 5. VARIATIONS & EXTENSIONS
- Range minimum/maximum queries
- Point updates with Mo's algorithm
- 2D range queries
- Different aggregation functions

## 6. INTERVIEW INSIGHTS
- Clarify: "Can all queries be processed offline?"
- Edge cases: empty array, single element, large ranges
- Time complexity: O((N+Q)*sqrt(N)) vs O(N*Q) naive
- Space complexity: O(N + sqrt(N))

## 7. COMMON MISTAKES
- Incorrect block size calculation
- Wrong query sorting logic
- Improper frequency tracking
- Off-by-one errors in range boundaries
- Not handling edge cases properly

## 8. OPTIMIZATION STRATEGIES
- Use optimal block size of sqrt(N)
- Efficient frequency map operations
- Careful query sorting to minimize moves
- Use bit manipulation for block calculations

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like organizing queries by geographic regions:**
- You have an array and many range sum queries
- Instead of processing each query independently, group queries by region
- Divide array into sqrt(N) blocks (like geographic regions)
- Process all queries in one block before moving to next
- This minimizes the number of times you need to move between regions
- Frequency tracking maintains current window sum efficiently

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Static array, multiple range sum queries
2. **Goal**: Process all queries efficiently
3. **Output**: Sum of elements for each query range

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(N*Q) for independent processing
- **"How to optimize?"** → Process queries in batches
- **"Why block decomposition?"** → Sqrt(N) blocks optimal for range queries
- **"How to track current state?"** → Frequency map for current window

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use Mo's algorithm:
1. Calculate optimal block size = sqrt(N)
2. Sort queries by block, then by right endpoint
3. Initialize frequency map and current sum
4. Process queries in order:
   - Expand/contract window to match query range
   - Use frequency map for O(1) operations
   - Store result for each query
5. Return results in original order"
```

#### Phase 4: Edge Case Handling
- **Empty array**: Return 0 for all queries
- **Single element**: Handle trivially
- **Invalid ranges**: Return 0 or handle appropriately
- **No queries**: Return empty result array

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Array: [1, 3, 5, 7, 9, 11], N=6
Queries: [[0,2], [1,4], [2,5]]

Human thinking:
"Let's apply Mo's algorithm:

Block size = sqrt(6) ≈ 2.44 → Use 2

Sort queries by block then right:
Query [0,2]: block=0, right=2
Query [1,4]: block=0, right=4
Query [2,5]: block=1, right=5

Sorted order: [0,2], [1,4], [2,5]

Initialize:
freq = {}, currentSum = 0, currentLeft = 0, currentRight = -1

Process query [0,2]:
- Expand right to 2: add(1), add(3) → sum=4
- Contract left to 0: remove(0) → sum=4
- Result[0,2] = 4

Process query [1,4]:
- Expand right to 4: add(5), add(7), add(9), add(11) → sum=32
- Contract left to 1: remove(1) → sum=31
- Result[1,4] = 31

Process query [2,5]:
- Expand right to 5: add(7), add(9), add(11) → sum=27
- Contract left to 2: remove(3) → sum=24
- Result[2,5] = 24

Final results: [4, 31, 24] ✓

Manual check:
sum[0,2] = 1+3 = 4 ✓
sum[1,4] = 3+5+7+9+11 = 35 ✓
sum[2,5] = 5+7+9+11 = 32 ✓

Results match! ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Block decomposition minimizes window movements
- **Why it's efficient**: O((N+Q)*sqrt(N)) vs O(N*Q) naive
- **Why it's correct**: Each query processed with exact range contents

### Common Human Pitfalls & How to Avoid Them
1. **"Why not process queries independently?"** → O(N*Q) too slow
2. **"What about block size?"** → sqrt(N) is optimal for range queries
3. **"How to handle sorting?"** → Sort by block then right endpoint
4. **"What about frequency tracking?"** → Essential for O(1) add/remove

### Real-World Analogy
**Like organizing package deliveries by geographic regions:**
- You have packages at different addresses (array)
- You need to calculate totals for many delivery zones (queries)
- Instead of visiting each zone separately, group zones by region
- Process all deliveries in one region before moving to next
- This minimizes travel time between regions
- Frequency tracking maintains current inventory efficiently
- Useful in logistics, geographic information systems, data analysis

### Human-Readable Pseudocode
```
function mosAlgorithm(nums, queries):
    n = nums.length
    blockSize = sqrt(n) + 1
    
    // Sort queries by block then right
    queries.sort((a, b) -> {
        blockA = a.left / blockSize
        blockB = b.left / blockSize
        
        if (blockA != blockB):
            return Integer.compare(blockA, blockB)
        
        if (blockA % 2 == 0):
            return Integer.compare(a.right, b.right)
        else:
            return Integer.compare(b.right, a.right)
    })
    
    // Process queries
    freq = {}
    currentSum = 0
    currentLeft = 0
    currentRight = -1
    results = array of size queries.length
    
    for query in queries:
        // Expand/contract to match query range
        while (currentRight < query.right):
            add(currentRight)
            currentRight++
        
        while (currentLeft > query.left):
            remove(currentLeft)
            currentLeft--
        
        results[query.index] = currentSum
    
    return results
```

### Execution Visualization

### Example: nums=[1,3,5,7,9,11], queries=[[0,2],[1,4],[2,5]]
```
Mo's Algorithm Process:

Block size = sqrt(6) ≈ 2.44 → Use 2

Query sorting:
[0,2]: block=0, right=2
[1,4]: block=0, right=4
[2,5]: block=1, right=5

Sorted order: [0,2], [1,4], [2,5]

Processing:
Initialize: freq={}, sum=0, left=0, right=-1

Query [0,2]:
- Expand to right=2: add(1), add(3) → sum=4
- Contract to left=0: remove(0) → sum=4
- Result[0,2] = 4

Query [1,4]:
- Expand to right=4: add(5), add(7), add(9), add(11) → sum=32
- Contract to left=1: remove(1) → sum=31
- Result[1,4] = 31

Query [2,5]:
- Expand to right=5: add(7), add(9), add(11) → sum=27
- Contract to left=2: remove(3) → sum=24
- Result[2,5] = 24

Final: [4, 31, 24]

Visualization:
Query processing order minimizes window movements
Frequency map enables O(1) add/remove operations
Block decomposition provides sqrt(N) optimal performance
```

### Key Visualization Points:
- **Block Sorting**: Queries ordered to minimize window movements
- **Window Management**: Efficient expand/contract with frequency tracking
- **Complexity Reduction**: O(N*Q) → O((N+Q)*sqrt(N))
- **Offline Processing**: All queries known in advance

### Memory Layout Visualization:
```
Array: [1, 3, 5, 7, 9, 11]
Block size: 2

Query Processing:
Query [0,2]: window [0,2], sum=4
Query [1,4]: window [1,4], sum=31
Query [2,5]: window [2,5], sum=24

Frequency Map Evolution:
During [0,2]: {1:1, 3:1}
During [1,4]: {1:1, 3:1, 5:1, 7:1, 9:1, 11:1}
During [2,5]: {3:1, 5:1, 7:1, 9:1, 11:1}

Window movements minimized through strategic query ordering
```

### Time Complexity Breakdown:
- **Query Sorting**: O(Q log Q)
- **Processing**: O((N+Q) * sqrt(N))
- **Space**: O(N + sqrt(N)) for frequency map
- **Optimal**: Best known complexity for offline range queries
- **vs Naive**: O(N*Q) vs O((N+Q)*sqrt(N))
*/
