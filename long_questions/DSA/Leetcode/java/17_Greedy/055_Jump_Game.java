import java.util.Arrays;

public class JumpGame {
    
    // 55. Jump Game
    // Time: O(N), Space: O(1)
    public static boolean canJump(int[] nums) {
        int maxReach = 0;
        
        for (int i = 0; i < nums.length; i++) {
            if (i > maxReach) {
                return false;
            }
            maxReach = Math.max(maxReach, i + nums[i]);
            if (maxReach >= nums.length - 1) {
                return true;
            }
        }
        
        return true;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {2, 3, 1, 1, 4},
            {3, 2, 1, 0, 4},
            {0},
            {1},
            {2, 0, 0},
            {1, 1, 1, 1, 1},
            {3, 2, 1, 0, 4, 5},
            {2, 0, 0, 0, 1},
            {1, 2, 3},
            {100, 0, 0, 0, 0}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            boolean result = canJump(testCases[i]);
            System.out.printf("Test Case %d: %s -> Can jump: %b\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}
