import java.util.*;

public class HouseRobber {
    
    // 198. House Robber
    // Time: O(N), Space: O(1)
    public static int rob(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        if (nums.length == 1) {
            return nums[0];
        }
        
        int prev2 = 0; // rob[i-2]
        int prev1 = nums[0]; // rob[i-1]
        int current = 0;
        
        for (int i = 1; i < nums.length; i++) {
            current = Math.max(prev1, prev2 + nums[i]);
            prev2 = prev1;
            prev1 = current;
        }
        
        return prev1;
    }

    // 213. House Robber II
    // Time: O(N), Space: O(1)
    public static int robII(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        if (nums.length == 1) {
            return nums[0];
        }
        if (nums.length == 2) {
            return Math.max(nums[0], nums[1]);
        }
        
        // Case 1: Rob from first to second-last house
        int case1 = robHelper(nums, 0, nums.length - 2);
        
        // Case 2: Rob from second to last house
        int case2 = robHelper(nums, 1, nums.length - 1);
        
        return Math.max(case1, case2);
    }
    
    private static int robHelper(int[] nums, int start, int end) {
        int prev2 = 0;
        int prev1 = nums[start];
        int current = 0;
        
        for (int i = start + 1; i <= end; i++) {
            current = Math.max(prev1, prev2 + nums[i]);
            prev2 = prev1;
            prev1 = current;
        }
        
        return prev1;
    }

    public static void main(String[] args) {
        // Test cases for House Robber I
        int[][] testCases1 = {
            {1, 2, 3, 1},
            {2, 7, 9, 3, 1},
            {1},
            {1, 2},
            {2, 1},
            {1, 2, 3},
            {2, 3, 2},
            {5, 5, 10, 100, 10, 5},
            {1, 3, 1, 3, 100},
            {2, 1, 1, 2},
            {6, 7, 1, 3, 8, 2, 4},
            {10, 1, 1, 10},
            {1, 2, 3, 4, 5},
            {5, 4, 3, 2, 1},
            {1, 1, 1, 1, 1}
        };
        
        // Test cases for House Robber II
        int[][] testCases2 = {
            {2, 3, 2},
            {1, 2, 3, 1},
            {1, 2, 3},
            {1},
            {1, 2},
            {2, 1},
            {1, 2, 1, 3},
            {6, 7, 1, 30, 8, 2, 4},
            {5, 5, 10, 100, 10, 5},
            {1, 3, 1, 3, 100},
            {2, 1, 1, 2},
            {10, 1, 1, 10},
            {1, 2, 3, 4, 5, 6},
            {100, 1, 1, 100},
            {200, 3, 140, 20, 10}
        };
        
        System.out.println("House Robber I:");
        for (int i = 0; i < testCases1.length; i++) {
            int[] nums = testCases1[i];
            int result = rob(nums);
            System.out.printf("Test Case %d: %s -> %d\n", 
                i + 1, Arrays.toString(nums), result);
        }
        
        System.out.println("\nHouse Robber II:");
        for (int i = 0; i < testCases2.length; i++) {
            int[] nums = testCases2[i];
            int result = robII(nums);
            System.out.printf("Test Case %d: %s -> %d\n", 
                i + 1, Arrays.toString(nums), result);
        }
    }
}
