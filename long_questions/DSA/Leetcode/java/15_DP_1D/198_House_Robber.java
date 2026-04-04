import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class HouseRobber {
    
    // 198. House Robber
    // Time: O(N), Space: O(1)
    public static int rob(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        if (nums.length == 1) {
            return nums[0];
        }
        
        int prev2 = nums[0];
        int prev1 = Math.max(nums[0], nums[1]);
        
        for (int i = 2; i < nums.length; i++) {
            int current = Math.max(prev1, prev2 + nums[i]);
            prev2 = prev1;
            prev1 = current;
        }
        
        return prev1;
    }

    // DP array approach
    public static int robDP(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        if (nums.length == 1) {
            return nums[0];
        }
        
        int[] dp = new int[nums.length];
        dp[0] = nums[0];
        dp[1] = Math.max(nums[0], nums[1]);
        
        for (int i = 2; i < nums.length; i++) {
            dp[i] = Math.max(dp[i - 1], dp[i - 2] + nums[i]);
        }
        
        return dp[nums.length - 1];
    }

    // Recursive with memoization
    public static int robMemo(int[] nums) {
        Map<Integer, Integer> memo = new HashMap<>();
        return robHelper(nums, nums.length - 1, memo);
    }

    private static int robHelper(int[] nums, int index, Map<Integer, Integer> memo) {
        if (index < 0) {
            return 0;
        }
        if (index == 0) {
            return nums[0];
        }
        
        if (memo.containsKey(index)) {
            return memo.get(index);
        }
        
        int result = Math.max(robHelper(nums, index - 1, memo), 
                             robHelper(nums, index - 2, memo) + nums[index]);
        memo.put(index, result);
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 2, 3, 1},
            {2, 7, 9, 3, 1},
            {2, 1, 1, 2},
            {1},
            {1, 2},
            {5, 5, 10, 100, 10, 5},
            {2, 7, 9, 3, 1, 5, 6, 8},
            {100, 1, 1, 100},
            {4, 1, 2, 7, 5, 3, 1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int result1 = rob(testCases[i]);
            int result2 = robDP(testCases[i]);
            int result3 = robMemo(testCases[i]);
            
            System.out.printf("Test Case %d: %s -> Iterative: %d, DP: %d, Memo: %d\n", 
                i + 1, Arrays.toString(testCases[i]), result1, result2, result3);
        }
    }
}
