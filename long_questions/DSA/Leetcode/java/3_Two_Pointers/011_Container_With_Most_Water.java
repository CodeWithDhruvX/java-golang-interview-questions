import java.util.Arrays;

public class ContainerWithMostWater {
    
    // 11. Container With Most Water
    // Time: O(N), Space: O(1)
    public static int maxArea(int[] height) {
        int left = 0, right = height.length - 1;
        int maxWater = 0;
        
        while (left < right) {
            // Calculate current area
            int width = right - left;
            int height1 = height[left];
            int height2 = height[right];
            int currentHeight = Math.min(height1, height2);
            
            int currentArea = width * currentHeight;
            maxWater = Math.max(maxWater, currentArea);
            
            // Move the pointer with smaller height
            if (height1 < height2) {
                left++;
            } else {
                right--;
            }
        }
        
        return maxWater;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 8, 6, 2, 5, 4, 8, 3, 7},
            {1, 1},
            {4, 3, 2, 1, 4},
            {1, 2, 1},
            {2, 3, 4, 5, 18, 17, 6},
            {1, 2, 3, 4, 5},
            {5, 4, 3, 2, 1},
            {1, 3, 2, 5, 25, 24, 5}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int result = maxArea(testCases[i]);
            System.out.printf("Test Case %d: %s -> Max Area: %d\n", 
                i + 1, Arrays.toString(testCases[i]), result);
        }
    }
}
