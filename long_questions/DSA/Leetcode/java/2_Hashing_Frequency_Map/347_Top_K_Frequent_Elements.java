import java.util.*;

public class TopKFrequentElements {
    
    // 347. Top K Frequent Elements
    // Time: O(N log K), Space: O(N)
    public static int[] topKFrequent(int[] nums, int k) {
        if (nums == null || nums.length == 0 || k <= 0) {
            return new int[0];
        }
        
        // Count frequency of each element
        Map<Integer, Integer> frequencyMap = new HashMap<>();
        for (int num : nums) {
            frequencyMap.put(num, frequencyMap.getOrDefault(num, 0) + 1);
        }
        
        // Use min heap to keep top k frequent elements
        PriorityQueue<int[]> heap = new PriorityQueue<>((a, b) -> a[1] - b[1]);
        
        for (Map.Entry<Integer, Integer> entry : frequencyMap.entrySet()) {
            heap.offer(new int[]{entry.getKey(), entry.getValue()});
            
            if (heap.size() > k) {
                heap.poll();
            }
        }
        
        // Extract elements from heap
        int[] result = new int[k];
        for (int i = k - 1; i >= 0; i--) {
            result[i] = heap.poll()[0];
        }
        
        return result;
    }

    public static void main(String[] args) {
        Object[][] testCases = {
            {new int[]{1, 1, 1, 2, 2, 3}, 2},
            {new int[]{1}, 1},
            {new int[]{1, 2, 3, 4, 5}, 3},
            {new int[]{1, 1, 1, 1, 1}, 1},
            {new int[]{1, 2, 2, 3, 3, 3}, 2},
            {new int[]{4, 4, 4, 6, 6, 5, 5, 5, 5}, 2},
            {new int[]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
            {new int[]{1, 2, 2, 1, 3, 3, 3, 2, 2}, 2},
            {new int[]{5, 5, 5, 5, 5, 5}, 1},
            {new int[]{1, 2, 3, 1, 2, 3, 1, 2, 3}, 3},
            {new int[]{0}, 1},
            {new int[]{1, 2, 3, 4, 5}, 5},
            {new int[]{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}, 4},
            {new int[]{1, 1, 2, 2, 3, 3, 4, 4}, 4},
            {new int[]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] nums = (int[]) testCases[i][0];
            int k = (int) testCases[i][1];
            
            int[] result = topKFrequent(nums, k);
            System.out.printf("Test Case %d: %s, k=%d -> %s\n", 
                i + 1, Arrays.toString(nums), k, Arrays.toString(result));
        }
    }
}
