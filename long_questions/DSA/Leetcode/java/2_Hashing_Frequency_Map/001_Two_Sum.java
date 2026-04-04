import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class TwoSum {
    
    // 1. Two Sum
    // Time: O(N), Space: O(N)
    public static int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> numMap = new HashMap<>();
        
        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];
            if (numMap.containsKey(complement)) {
                return new int[]{numMap.get(complement), i};
            }
            numMap.put(nums[i], i);
        }
        
        return new int[]{};
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testArrays = {
            {2, 7, 11, 15},
            {3, 2, 4},
            {3, 3},
            {1, 2, 3, 4, 5},
            {-1, -2, -3, -4, -5}
        };
        
        int[] targets = {9, 6, 6, 9, -8};
        
        for (int i = 0; i < testArrays.length; i++) {
            int[] result = twoSum(testArrays[i], targets[i]);
            System.out.printf("Test Case %d: nums=%s, target=%d -> Indices: %s\n", 
                i + 1, Arrays.toString(testArrays[i]), targets[i], Arrays.toString(result));
        }
    }
}
