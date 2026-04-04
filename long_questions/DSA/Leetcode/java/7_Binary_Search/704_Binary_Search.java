import java.util.Arrays;

public class BinarySearch {
    
    // 704. Binary Search
    // Time: O(log N), Space: O(1)
    public static int search(int[] nums, int target) {
        int left = 0, right = nums.length - 1;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (nums[mid] == target) {
                return mid;
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        
        return -1;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testArrays = {
            {-1, 0, 3, 5, 9, 12},
            {-1, 0, 3, 5, 9, 12},
            {1, 2, 3, 4, 5},
            {1, 2, 3, 4, 5},
            {},
            {1},
            {1},
            {-10, -5, 0, 5, 10},
            {-10, -5, 0, 5, 10},
            {2, 4, 6, 8, 10}
        };
        
        int[] targets = {9, 2, 3, 6, 1, 1, 0, -5, 0, 7};
        
        for (int i = 0; i < testArrays.length; i++) {
            int result = search(testArrays[i], targets[i]);
            System.out.printf("Test Case %d: nums=%s, target=%d -> Index: %d\n", 
                i + 1, Arrays.toString(testArrays[i]), targets[i], result);
        }
    }
}
