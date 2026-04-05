import java.util.*;

public class FindMedianFromDataStream {
    
    // 295. Find Median from Data Stream
    // Time: O(log N) for addNum, O(1) for findMedian, Space: O(N)
    public static class MedianFinder {
        private PriorityQueue<Integer> maxHeap; // Lower half
        private PriorityQueue<Integer> minHeap; // Higher half
        
        public MedianFinder() {
            maxHeap = new PriorityQueue<>(Collections.reverseOrder());
            minHeap = new PriorityQueue<>();
        }
        
        public void addNum(int num) {
            // Add to maxHeap first
            maxHeap.offer(num);
            
            // Balance heaps
            minHeap.offer(maxHeap.poll());
            
            // Ensure maxHeap has equal or one more element than minHeap
            if (maxHeap.size() < minHeap.size()) {
                maxHeap.offer(minHeap.poll());
            }
        }
        
        public double findMedian() {
            if (maxHeap.size() == minHeap.size()) {
                return (maxHeap.peek() + minHeap.peek()) / 2.0;
            } else {
                return maxHeap.peek();
            }
        }
    }

    public static void main(String[] args) {
        // Test Case 1: Basic operations
        System.out.println("Test Case 1: Basic operations");
        MedianFinder mf1 = new MedianFinder();
        mf1.addNum(1);
        System.out.printf("Median after adding 1: %.1f\n", mf1.findMedian());
        mf1.addNum(2);
        System.out.printf("Median after adding 2: %.1f\n", mf1.findMedian());
        mf1.addNum(3);
        System.out.printf("Median after adding 3: %.1f\n", mf1.findMedian());
        
        // Test Case 2: Even number of elements
        System.out.println("\nTest Case 2: Even number of elements");
        MedianFinder mf2 = new MedianFinder();
        mf2.addNum(5);
        System.out.printf("Median after adding 5: %.1f\n", mf2.findMedian());
        mf2.addNum(15);
        System.out.printf("Median after adding 15: %.1f\n", mf2.findMedian());
        mf2.addNum(1);
        System.out.printf("Median after adding 1: %.1f\n", mf2.findMedian());
        mf2.addNum(3);
        System.out.printf("Median after adding 3: %.1f\n", mf2.findMedian());
        
        // Test Case 3: Same numbers
        System.out.println("\nTest Case 3: Same numbers");
        MedianFinder mf3 = new MedianFinder();
        mf3.addNum(2);
        mf3.addNum(2);
        mf3.addNum(2);
        mf3.addNum(2);
        System.out.printf("Median after adding four 2s: %.1f\n", mf3.findMedian());
        
        // Test Case 4: Negative numbers
        System.out.println("\nTest Case 4: Negative numbers");
        MedianFinder mf4 = new MedianFinder();
        mf4.addNum(-1);
        System.out.printf("Median after adding -1: %.1f\n", mf4.findMedian());
        mf4.addNum(-2);
        System.out.printf("Median after adding -2: %.1f\n", mf4.findMedian());
        mf4.addNum(-3);
        System.out.printf("Median after adding -3: %.1f\n", mf4.findMedian());
        
        // Test Case 5: Large numbers
        System.out.println("\nTest Case 5: Large numbers");
        MedianFinder mf5 = new MedianFinder();
        mf5.addNum(1000);
        mf5.addNum(2000);
        mf5.addNum(3000);
        System.out.printf("Median after adding 1000, 2000, 3000: %.1f\n", mf5.findMedian());
        
        // Test Case 6: Random sequence
        System.out.println("\nTest Case 6: Random sequence");
        MedianFinder mf6 = new MedianFinder();
        int[] sequence = {6, 10, 2, 8, 4, 12, 1};
        for (int num : sequence) {
            mf6.addNum(num);
            System.out.printf("After adding %d, median: %.1f\n", num, mf6.findMedian());
        }
    }
}
