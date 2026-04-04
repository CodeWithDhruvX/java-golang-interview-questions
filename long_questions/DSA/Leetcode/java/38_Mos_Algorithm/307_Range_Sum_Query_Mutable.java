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
}
