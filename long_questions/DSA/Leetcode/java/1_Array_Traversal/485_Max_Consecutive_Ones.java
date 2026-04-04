import java.util.Arrays;

public class MaxConsecutiveOnes {
    
    // 485. Max Consecutive Ones
    // Time: O(N), Space: O(1)
    public static int findMaxConsecutiveOnes(int[] nums) {
        int maxCount = 0;
        int currentCount = 0;
        
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 1) {
                currentCount++;
                if (currentCount > maxCount) {
                    maxCount = currentCount;
                }
            } else {
                currentCount = 0;
            }
        }
        
        return maxCount;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 1, 0, 1, 1, 1},
            {1, 0, 1, 1, 0, 1},
            {0, 0, 0},
            {1, 1, 1, 1},
            {1, 0, 1, 0, 1, 0, 1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int result = findMaxConsecutiveOnes(testCases[i]);
            System.out.printf("Test Case %d: %s -> Max Consecutive Ones: %d\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}
