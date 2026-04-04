import java.util.Arrays;

public class FindFirstAndLastPositionOfElementInSortedArray {
    
    // 34. Find First and Last Position of Element in Sorted Array
    // Time: O(log N), Space: O(1)
    public static int[] searchRange(int[] nums, int target) {
        return new int[]{findFirstOccurrence(nums, target), findLastOccurrence(nums, target)};
    }

    private static int findFirstOccurrence(int[] nums, int target) {
        int left = 0, right = nums.length - 1;
        int result = -1;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (nums[mid] == target) {
                result = mid;
                right = mid - 1; // Continue searching left half
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        
        return result;
    }

    private static int findLastOccurrence(int[] nums, int target) {
        int left = 0, right = nums.length - 1;
        int result = -1;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (nums[mid] == target) {
                result = mid;
                left = mid + 1; // Continue searching right half
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        
        return result;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testArrays = {
            {5, 7, 7, 8, 8, 10},
            {5, 7, 7, 8, 8, 10},
            {},
            {1},
            {1},
            {2, 2, 2, 2, 2},
            {1, 2, 3, 4, 5},
            {1, 2, 3, 4, 5},
            {-3, -2, -1, 0, 1},
            {1, 3, 5, 7, 9}
        };
        
        int[] targets = {8, 6, 0, 1, 0, 2, 3, 6, -1, 4};
        
        for (int i = 0; i < testArrays.length; i++) {
            int[] result = searchRange(testArrays[i], targets[i]);
            System.out.printf("Test Case %d: nums=%s, target=%d -> Range: %s\n", 
                i + 1, Arrays.toString(testArrays[i]), targets[i], Arrays.toString(result));
        }
    }
}
