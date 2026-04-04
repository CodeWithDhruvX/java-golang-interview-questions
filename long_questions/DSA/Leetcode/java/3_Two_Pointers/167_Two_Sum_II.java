import java.util.Arrays;

public class TwoSumII {
    
    // 167. Two Sum II - Input Array Is Sorted
    // Time: O(N), Space: O(1)
    public static int[] twoSumII(int[] numbers, int target) {
        int left = 0, right = numbers.length - 1;
        
        while (left < right) {
            int sum = numbers[left] + numbers[right];
            
            if (sum == target) {
                return new int[]{left + 1, right + 1}; // 1-indexed
            } else if (sum < target) {
                left++;
            } else {
                right--;
            }
        }
        
        return new int[]{};
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testArrays = {
            {2, 7, 11, 15},
            {2, 3, 4},
            {-1, 0},
            {1, 2, 3, 4, 4, 9, 56, 90},
            {-10, -5, -3, 0, 1, 3, 5, 10},
            {1, 3, 5, 7, 9}
        };
        
        int[] targets = {9, 6, -1, 8, 0, 12};
        
        for (int i = 0; i < testArrays.length; i++) {
            int[] result = twoSumII(testArrays[i], targets[i]);
            System.out.printf("Test Case %d: numbers=%s, target=%d -> Indices: %s\n", 
                i + 1, Arrays.toString(testArrays[i]), targets[i], Arrays.toString(result));
        }
    }
}
