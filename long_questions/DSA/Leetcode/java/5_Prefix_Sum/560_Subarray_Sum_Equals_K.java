import java.util.*;

public class SubarraySumEqualsK {
    
    // 560. Subarray Sum Equals K
    // Time: O(N), Space: O(N)
    public static int subarraySum(int[] nums, int k) {
        int count = 0;
        Map<Integer, Integer> prefixSum = new HashMap<>();
        prefixSum.put(0, 1); // Initialize with sum 0 occurring once
        
        int currentSum = 0;
        
        for (int num : nums) {
            currentSum += num;
            
            // Check if (currentSum - k) exists in prefixSum
            if (prefixSum.containsKey(currentSum - k)) {
                count += prefixSum.get(currentSum - k);
            }
            
            // Update prefixSum with currentSum
            prefixSum.put(currentSum, prefixSum.getOrDefault(currentSum, 0) + 1);
        }
        
        return count;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testArrays = {
            {1, 1, 1},
            {1, 2, 3},
            {1, -1, 0},
            {0, 0, 0, 0},
            {-1, -1, 1},
            {3, 4, 7, -2, 2, 1, 4, 2},
            {1},
            {1},
            {},
            {1, 2, 1, 2, 1}
        };
        
        int[] kValues = {2, 3, 0, 0, 0, 7, 1, 0, 0, 3};
        
        for (int i = 0; i < testArrays.length; i++) {
            int result = subarraySum(testArrays[i], kValues[i]);
            System.out.printf("Test Case %d: nums=%s, k=%d -> Subarrays with sum k: %d\n", 
                i + 1, Arrays.toString(testArrays[i]), kValues[i], result);
        }
    }
}
