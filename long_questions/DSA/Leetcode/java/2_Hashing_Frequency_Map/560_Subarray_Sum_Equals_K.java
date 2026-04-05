import java.util.*;

public class SubarraySumEqualsK {
    
    // 560. Subarray Sum Equals K
    // Time: O(N), Space: O(N)
    public static int subarraySum(int[] nums, int k) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        
        Map<Integer, Integer> prefixSumCount = new HashMap<>();
        prefixSumCount.put(0, 1); // Empty prefix sum
        
        int count = 0;
        int currentSum = 0;
        
        for (int num : nums) {
            currentSum += num;
            
            // Check if (currentSum - k) exists in map
            if (prefixSumCount.containsKey(currentSum - k)) {
                count += prefixSumCount.get(currentSum - k);
            }
            
            // Update prefix sum count
            prefixSumCount.put(currentSum, prefixSumCount.getOrDefault(currentSum, 0) + 1);
        }
        
        return count;
    }

    public static void main(String[] args) {
        Object[][] testCases = {
            {new int[]{1, 1, 1}, 2},
            {new int[]{1, 2, 3}, 3},
            {new int[]{1, -1, 0}, 0},
            {new int[]{1, 2, 1, 2, 1}, 3},
            {new int[]{0, 0, 0, 0}, 0},
            {new int[]{-1, -1, 1}, 0},
            {new int[]{1}, 1},
            {new int[]{1}, 2},
            {new int[]{}, 1},
            {new int[]{1, 2, 3, 4, 5}, 9},
            {new int[]{1, 2, 3, -2, 5}, 5},
            {new int[]{1, 2, 3, 4, 5, 6}, 7},
            {new int[]{1, 1, 1, 1, 1}, 2},
            {new int[]{2, 2, 2, 2, 2}, 4}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] nums = (int[]) testCases[i][0];
            int k = (int) testCases[i][1];
            
            int result = subarraySum(nums, k);
            System.out.printf("Test Case %d: %s, k=%d -> %d\n", 
                i + 1, Arrays.toString(nums), k, result);
        }
    }
}
