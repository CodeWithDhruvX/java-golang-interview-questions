import java.util.Arrays;
import java.util.PriorityQueue;

public class KthLargestElementInAnArray {
    
    // 215. Kth Largest Element in an Array
    // Time: O(N log K), Space: O(K)
    public static int findKthLargest(int[] nums, int k) {
        // Use a min-heap of size k
        PriorityQueue<Integer> minHeap = new PriorityQueue<>();
        
        for (int num : nums) {
            minHeap.offer(num);
            if (minHeap.size() > k) {
                minHeap.poll();
            }
        }
        
        return minHeap.peek();
    }

    // Alternative solution using QuickSelect (O(N) average case)
    public static int findKthLargestQuickSelect(int[] nums, int k) {
        return quickSelect(nums, 0, nums.length - 1, nums.length - k);
    }

    private static int quickSelect(int[] nums, int left, int right, int kthSmallest) {
        if (left == right) {
            return nums[left];
        }
        
        int pivotIndex = partition(nums, left, right);
        
        if (kthSmallest == pivotIndex) {
            return nums[pivotIndex];
        } else if (kthSmallest < pivotIndex) {
            return quickSelect(nums, left, pivotIndex - 1, kthSmallest);
        } else {
            return quickSelect(nums, pivotIndex + 1, right, kthSmallest);
        }
    }

    private static int partition(int[] nums, int left, int right) {
        int pivot = nums[right];
        int i = left;
        
        for (int j = left; j < right; j++) {
            if (nums[j] <= pivot) {
                swap(nums, i, j);
                i++;
            }
        }
        
        swap(nums, i, right);
        return i;
    }

    private static void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

    public static void main(String[] args) {
        // Test cases
        int[][] testArrays = {
            {3, 2, 1, 5, 6, 4},
            {3, 2, 3, 1, 2, 4, 5, 5, 6},
            {1},
            {2, 1},
            {1, 2, 3, 4, 5},
            {5, 4, 3, 2, 1},
            {3, 2, 3, 1, 2, 4, 5, 5, 6},
            {7, 10, 4, 3, 20, 15},
            {-1, -2, -3, -4, -5},
            {100, 200, 300, 400, 500}
        };
        
        int[] kValues = {2, 4, 1, 1, 5, 1, 1, 3, 2, 4};
        
        for (int i = 0; i < testArrays.length; i++) {
            // Make copies for both methods
            int[] nums1 = testArrays[i].clone();
            int[] nums2 = testArrays[i].clone();
            
            int result1 = findKthLargest(nums1, kValues[i]);
            int result2 = findKthLargestQuickSelect(nums2, kValues[i]);
            
            System.out.printf("Test Case %d: nums=%s, k=%d -> Heap: %d, QuickSelect: %d\n", 
                i + 1, Arrays.toString(testArrays[i]), kValues[i], result1, result2);
        }
    }
}
