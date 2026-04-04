import java.util.*;

public class FindTheDuplicateNumber {
    
    // 287. Find the Duplicate Number
    // Time: O(N), Space: O(1) - Floyd's Tortoise and Hare Algorithm
    public static int findDuplicate(int[] nums) {
        // Phase 1: Find the intersection point
        int slow = nums[0];
        int fast = nums[0];
        
        while (true) {
            slow = nums[slow];
            fast = nums[nums[fast]];
            if (slow == fast) {
                break;
            }
        }
        
        // Phase 2: Find the entrance to the cycle
        slow = nums[0];
        while (slow != fast) {
            slow = nums[slow];
            fast = nums[fast];
        }
        
        return slow;
    }

    // Alternative solution using cyclic sort (modifies the array)
    public static int findDuplicateCyclicSort(int[] nums) {
        int i = 0;
        int n = nums.length;
        
        while (i < n) {
            int correctPos = nums[i] - 1;
            if (nums[i] != nums[correctPos]) {
                swap(nums, i, correctPos);
            } else {
                i++;
            }
        }
        
        // The duplicate will be at the position where the number doesn't match index+1
        for (i = 0; i < n; i++) {
            if (nums[i] != i + 1) {
                return nums[i];
            }
        }
        
        return -1; // Should never reach here for valid input
    }
    
    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

    // Binary search approach
    public static int findDuplicateBinarySearch(int[] nums) {
        int left = 1;
        int right = nums.length - 1;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            int count = 0;
            
            // Count numbers less than or equal to mid
            for (int num : nums) {
                if (num <= mid) {
                    count++;
                }
            }
            
            if (count > mid) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        
        return left;
    }

    // Hash set approach (uses extra space)
    public static int findDuplicateHashSet(int[] nums) {
        Set<Integer> seen = new HashSet<>();
        
        for (int num : nums) {
            if (seen.contains(num)) {
                return num;
            }
            seen.add(num);
        }
        
        return -1; // Should never reach here for valid input
    }

    // Array marking approach (modifies array)
    public static int findDuplicateArrayMarking(int[] nums) {
        for (int num : nums) {
            int index = Math.abs(num);
            if (nums[index] < 0) {
                return index;
            }
            nums[index] = -nums[index];
        }
        
        return -1; // Should never reach here for valid input
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testCases = {
            {1, 3, 4, 2, 2},
            {3, 1, 3, 4, 2},
            {1, 1},
            {2, 2, 2, 2, 2},
            {1, 4, 4, 3, 2},
            {5, 4, 3, 2, 1, 5},
            {3, 1, 2, 3, 4, 5},
            {2, 5, 9, 6, 9, 3, 8, 9, 7, 1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            // Make copies for different approaches
            int[] nums1 = testCases[i].clone();
            int[] nums2 = testCases[i].clone();
            int[] nums3 = testCases[i].clone();
            int[] nums4 = testCases[i].clone();
            int[] nums5 = testCases[i].clone();
            
            int result1 = findDuplicate(nums1);
            int result2 = findDuplicateCyclicSort(nums2);
            int result3 = findDuplicateBinarySearch(nums3);
            int result4 = findDuplicateHashSet(nums4);
            int result5 = findDuplicateArrayMarking(nums5);
            
            System.out.printf("Test Case %d: %s\n", i + 1, Arrays.toString(testCases[i]));
            System.out.printf("  Floyd's: %d\n", result1);
            System.out.printf("  Cyclic: %d\n", result2);
            System.out.printf("  Binary:  %d\n", result3);
            System.out.printf("  HashSet: %d\n", result4);
            System.out.printf("  Marking: %d\n\n", result5);
        }
    }
}
