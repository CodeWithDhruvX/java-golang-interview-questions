public class MaximumSubarray {
    
    // 53. Maximum Subarray - Divide and Conquer Approach
    // Time: O(N log N), Space: O(log N) for recursion stack
    public int maxSubArrayDivideAndConquer(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        return maxSubArrayHelper(nums, 0, nums.length - 1);
    }
    
    private int maxSubArrayHelper(int[] nums, int left, int right) {
        if (left == right) {
            return nums[left];
        }
        
        int mid = left + (right - left) / 2;
        
        // Maximum subarray in left half
        int leftMax = maxSubArrayHelper(nums, left, mid);
        
        // Maximum subarray in right half
        int rightMax = maxSubArrayHelper(nums, mid + 1, right);
        
        // Maximum subarray crossing the middle
        int crossMax = maxCrossingSubArray(nums, left, mid, right);
        
        return Math.max(leftMax, Math.max(rightMax, crossMax));
    }
    
    private int maxCrossingSubArray(int[] nums, int left, int mid, int right) {
        // Maximum sum starting from mid and going left
        int leftSum = Integer.MIN_VALUE;
        int sum = 0;
        
        for (int i = mid; i >= left; i--) {
            sum += nums[i];
            if (sum > leftSum) {
                leftSum = sum;
            }
        }
        
        // Maximum sum starting from mid+1 and going right
        int rightSum = Integer.MIN_VALUE;
        sum = 0;
        
        for (int i = mid + 1; i <= right; i++) {
            sum += nums[i];
            if (sum > rightSum) {
                rightSum = sum;
            }
        }
        
        return leftSum + rightSum;
    }
    
    // Kadane's algorithm for comparison (O(N))
    public int maxSubArrayKadane(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        int maxSoFar = nums[0];
        int maxEndingHere = nums[0];
        
        for (int i = 1; i < nums.length; i++) {
            maxEndingHere = Math.max(nums[i], maxEndingHere + nums[i]);
            maxSoFar = Math.max(maxSoFar, maxEndingHere);
        }
        
        return maxSoFar;
    }
    
    // Version with detailed explanation
    public class MaxSubarrayResult {
        int maxSum;
        int startIndex;
        int endIndex;
        java.util.List<String> explanation;
        
        MaxSubarrayResult(int maxSum, int startIndex, int endIndex, java.util.List<String> explanation) {
            this.maxSum = maxSum;
            this.startIndex = startIndex;
            this.endIndex = endIndex;
            this.explanation = explanation;
        }
    }
    
    public MaxSubarrayResult maxSubArrayDetailed(int[] nums) {
        java.util.List<String> explanation = new java.util.ArrayList<>();
        explanation.add("=== Divide and Conquer for Maximum Subarray ===");
        explanation.add("Array: " + java.util.Arrays.toString(nums));
        
        if (nums.length == 0) {
            explanation.add("Empty array, returning 0");
            return new MaxSubarrayResult(0, -1, -1, explanation);
        }
        
        MaxSubarrayInfo result = maxSubArrayHelperDetailed(nums, 0, nums.length - 1, explanation);
        
        explanation.add(String.format("Final result: sum=%d, start=%d, end=%d", 
            result.maxSum, result.startIndex, result.endIndex));
        
        return new MaxSubarrayResult(result.maxSum, result.startIndex, result.endIndex, explanation);
    }
    
    private static class MaxSubarrayInfo {
        int maxSum;
        int startIndex;
        int endIndex;
        
        MaxSubarrayInfo(int maxSum, int startIndex, int endIndex) {
            this.maxSum = maxSum;
            this.startIndex = startIndex;
            this.endIndex = endIndex;
        }
    }
    
    private MaxSubarrayInfo maxSubArrayHelperDetailed(int[] nums, int left, int right, java.util.List<String> explanation) {
        if (left == right) {
            explanation.add(String.format("Base case: single element nums[%d] = %d", left, nums[left]));
            return new MaxSubarrayInfo(nums[left], left, right);
        }
        
        int mid = left + (right - left) / 2;
        explanation.add(String.format("Dividing: left=%d, mid=%d, right=%d", left, mid, right));
        
        // Maximum subarray in left half
        explanation.add("Finding maximum in left half...");
        MaxSubarrayInfo leftMax = maxSubArrayHelperDetailed(nums, left, mid, explanation);
        explanation.add(String.format("Left max: sum=%d, range=[%d,%d]", leftMax.maxSum, leftMax.startIndex, leftMax.endIndex));
        
        // Maximum subarray in right half
        explanation.add("Finding maximum in right half...");
        MaxSubarrayInfo rightMax = maxSubArrayHelperDetailed(nums, mid + 1, right, explanation);
        explanation.add(String.format("Right max: sum=%d, range=[%d,%d]", rightMax.maxSum, rightMax.startIndex, rightMax.endIndex));
        
        // Maximum subarray crossing the middle
        explanation.add("Finding maximum crossing middle...");
        MaxSubarrayInfo crossMax = maxCrossingSubArrayDetailed(nums, left, mid, right, explanation);
        explanation.add(String.format("Cross max: sum=%d, range=[%d,%d]", crossMax.maxSum, crossMax.startIndex, crossMax.endIndex));
        
        // Find the maximum of three
        MaxSubarrayInfo result;
        if (leftMax.maxSum >= rightMax.maxSum && leftMax.maxSum >= crossMax.maxSum) {
            result = leftMax;
            explanation.add("Left half has the maximum");
        } else if (rightMax.maxSum >= leftMax.maxSum && rightMax.maxSum >= crossMax.maxSum) {
            result = rightMax;
            explanation.add("Right half has the maximum");
        } else {
            result = crossMax;
            explanation.add("Crossing subarray has the maximum");
        }
        
        return result;
    }
    
    private MaxSubarrayInfo maxCrossingSubArrayDetailed(int[] nums, int left, int mid, int right, java.util.List<String> explanation) {
        // Maximum sum starting from mid and going left
        explanation.add(String.format("Finding left crossing sum from %d to %d", mid, left));
        
        int leftSum = Integer.MIN_VALUE;
        int sum = 0;
        int leftIndex = mid;
        
        for (int i = mid; i >= left; i--) {
            sum += nums[i];
            explanation.add(String.format("  nums[%d] = %d, running sum = %d", i, nums[i], sum));
            
            if (sum > leftSum) {
                leftSum = sum;
                leftIndex = i;
                explanation.add(String.format("  New left max: %d at index %d", leftSum, i));
            }
        }
        
        // Maximum sum starting from mid+1 and going right
        explanation.add(String.format("Finding right crossing sum from %d to %d", mid + 1, right));
        
        int rightSum = Integer.MIN_VALUE;
        sum = 0;
        int rightIndex = mid + 1;
        
        for (int i = mid + 1; i <= right; i++) {
            sum += nums[i];
            explanation.add(String.format("  nums[%d] = %d, running sum = %d", i, nums[i], sum));
            
            if (sum > rightSum) {
                rightSum = sum;
                rightIndex = i;
                explanation.add(String.format("  New right max: %d at index %d", rightSum, i));
            }
        }
        
        int totalSum = leftSum + rightSum;
        explanation.add(String.format("Crossing sum: left=%d + right=%d = %d, range=[%d,%d]", 
            leftSum, rightSum, totalSum, leftIndex, rightIndex));
        
        return new MaxSubarrayInfo(totalSum, leftIndex, rightIndex);
    }
    
    // Performance comparison
    public void comparePerformance(int[] nums) {
        System.out.println("=== Performance Comparison ===");
        System.out.println("Array: " + java.util.Arrays.toString(nums));
        
        // Divide and Conquer
        long startTime = System.nanoTime();
        int result1 = maxSubArrayDivideAndConquer(nums);
        long endTime = System.nanoTime();
        System.out.printf("Divide and Conquer: %d (took %d ns)\n", result1, endTime - startTime);
        
        // Kadane's algorithm
        startTime = System.nanoTime();
        int result2 = maxSubArrayKadane(nums);
        endTime = System.nanoTime();
        System.out.printf("Kadane's algorithm: %d (took %d ns)\n", result2, endTime - startTime);
    }
    
    // Find maximum subarray in each quadrant
    public int[] maxSubArrayQuadrants(int[] nums) {
        if (nums.length == 0) {
            return new int[]{0, -1, -1};
        }
        
        int n = nums.length;
        int[] results = new int[4];
        
        // Divide into 4 quadrants
        int quarter = n / 4;
        
        results[0] = maxSubArrayHelper(nums, 0, quarter - 1);
        results[1] = maxSubArrayHelper(nums, quarter, n/2 - 1);
        results[2] = maxSubArrayHelper(nums, n/2, 3*quarter - 1);
        results[3] = maxSubArrayHelper(nums, 3*quarter, n - 1);
        
        return results;
    }
    
    // Parallel divide and conquer (conceptual)
    public int maxSubArrayParallel(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        
        // For demonstration, we'll use sequential approach
        // In practice, this would use multiple threads
        return maxSubArrayDivideAndConquer(nums);
    }
    
    // Maximum subarray with constraints
    public int maxSubArrayWithConstraints(int[] nums, int minLength, int maxLength) {
        if (nums.length == 0) {
            return 0;
        }
        
        int maxSum = Integer.MIN_VALUE;
        
        for (int i = 0; i < nums.length; i++) {
            for (int j = i; j < nums.length; j++) {
                int length = j - i + 1;
                
                if (length >= minLength && length <= maxLength) {
                    int sum = 0;
                    for (int k = i; k <= j; k++) {
                        sum += nums[k];
                    }
                    maxSum = Math.max(maxSum, sum);
                }
            }
        }
        
        return maxSum;
    }
    
    public static void main(String[] args) {
        MaximumSubarray ms = new MaximumSubarray();
        
        // Test cases
        int[][] testCases = {
            {-2,1,-3,4,-1,2,1,-5,4},
            {1},
            {5,4,-1,7,8},
            {-2,-3,-1,-5},
            {1,2,3,4,5},
            {-1,2,-3,4,-5,6},
            {100, -200, 300, -400, 500},
            {-2, -1, -3, -4},
            {0, 0, 0, 0},
            {8, -19, 5, -4, 20}
        };
        
        String[] descriptions = {
            "Standard case",
            "Single element",
            "All positive",
            "All negative",
            "Increasing sequence",
            "Alternating signs",
            "Mixed large numbers",
            "All negative small",
            "All zeros",
            "Complex case"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("Array: %s\n", java.util.Arrays.toString(testCases[i]));
            
            int result1 = ms.maxSubArrayDivideAndConquer(testCases[i]);
            int result2 = ms.maxSubArrayKadane(testCases[i]);
            
            System.out.printf("Divide and Conquer: %d\n", result1);
            System.out.printf("Kadane's algorithm: %d\n", result2);
            
            // Detailed analysis for first case
            if (i == 0) {
                MaxSubarrayResult detailed = ms.maxSubArrayDetailed(testCases[i]);
                System.out.println("Detailed explanation:");
                for (String step : detailed.explanation) {
                    System.out.println("  " + step);
                }
            }
            
            System.out.println();
        }
        
        // Performance comparison
        System.out.println("=== Performance Comparison ===");
        int[] largeArray = new int[10000];
        java.util.Random rand = new java.util.Random();
        for (int i = 0; i < largeArray.length; i++) {
            largeArray[i] = rand.nextInt(200) - 100; // Random between -100 and 100
        }
        
        ms.comparePerformance(largeArray);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        System.out.printf("Empty array: %d\n", ms.maxSubArrayDivideAndConquer(new int[]{}));
        System.out.printf("Single negative: %d\n", ms.maxSubArrayDivideAndConquer(new int[]{-5}));
        System.out.printf("Single positive: %d\n", ms.maxSubArrayDivideAndConquer(new int[]{5}));
        System.out.printf("Two elements: %d\n", ms.maxSubArrayDivideAndConquer(new int[]{1, -2}));
        
        // Maximum subarray with constraints
        System.out.println("\n=== With Constraints ===");
        int[] constrainedArray = {1, -2, 3, 4, -5, 8};
        int constrainedResult = ms.maxSubArrayWithConstraints(constrainedArray, 2, 3);
        System.out.printf("Array: %s\n", java.util.Arrays.toString(constrainedArray));
        System.out.printf("Max subarray with length 2-3: %d\n", constrainedResult);
    }
}
