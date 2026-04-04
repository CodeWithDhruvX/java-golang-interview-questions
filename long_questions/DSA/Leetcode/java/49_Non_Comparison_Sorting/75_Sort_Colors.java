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
}
