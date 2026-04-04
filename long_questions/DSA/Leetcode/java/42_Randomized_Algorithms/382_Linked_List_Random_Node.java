import java.util.*;

public class LinkedListRandomNode {
    
    // 382. Linked List Random Node - Randomized Algorithms
    // Time: O(N), Space: O(1)
    static class ListNode {
        int val;
        ListNode next;
        
        ListNode(int val) {
            this.val = val;
        }
    }
    
    static class RandomizedLinkedList {
        ListNode head;
        Random random;
        
        public RandomizedLinkedList(ListNode head) {
            this.head = head;
            this.random = new Random();
        }
        
        // Reservoir sampling for random node selection
        public int getRandom() {
            if (head == null) {
                return -1;
            }
            
            // Reservoir sampling algorithm
            int count = 0;
            int result = head.val;
            ListNode current = head;
            
            while (current != null) {
                count++;
                // With probability 1/count, select current node
                if (random.nextInt(count) == 0) {
                    result = current.val;
                }
                current = current.next;
            }
            
            return result;
        }
    }
    
    // Monte Carlo algorithm for approximate counting
    static class MonteCarloLinkedList {
        ListNode head;
        Random random;
        
        public MonteCarloLinkedList(ListNode head) {
            this.head = head;
            this.random = new Random();
        }
        
        public int approximateCount(int iterations) {
            if (head == null) {
                return 0;
            }
            
            // Monte Carlo sampling to estimate length
            int sampleCount = 0;
            
            for (int i = 0; i < iterations; i++) {
                ListNode current = head;
                int steps = 0;
                
                // Random walk to a random position
                while (current != null) {
                    steps++;
                    if (random.nextDouble() < 0.5) {
                        break;
                    }
                    current = current.next;
                }
                
                sampleCount += steps;
            }
            
            // Estimate total length
            int estimatedLength = sampleCount * 2 / iterations;
            return estimatedLength;
        }
    }
    
    // Las Vegas algorithm for finding kth random element
    static class LasVegasLinkedList {
        ListNode head;
        Random random;
        
        public LasVegasLinkedList(ListNode head) {
            this.head = head;
            this.random = new Random();
        }
        
        public int getKthRandom(int k) {
            if (head == null || k <= 0) {
                return -1;
            }
            
            // Las Vegas algorithm - always correct
            List<Integer> samples = new ArrayList<>();
            ListNode current = head;
            
            while (current != null) {
                samples.add(current.val);
                current = current.next;
            }
            
            // Select kth random element
            Collections.shuffle(samples);
            if (samples.size() >= k) {
                return samples.get(k - 1);
            }
            
            return -1;
        }
    }
    
    // Sherwood algorithm for randomized quicksort
    static class SherwoodLinkedList {
        ListNode head;
        Random random;
        
        public SherwoodLinkedList(ListNode head) {
            this.head = head;
            this.random = new Random();
        }
        
        public void randomizeQuickSort() {
            if (head == null) {
                return;
            }
            
            // Collect all values
            List<Integer> values = new ArrayList<>();
            ListNode current = head;
            while (current != null) {
                values.add(current.val);
                current = current.next;
            }
            
            // Randomized quicksort
            randomizedQuickSort(values, 0, values.size() - 1);
            
            // Rebuild linked list
            this.head = buildLinkedList(values);
        }
        
        private void randomizedQuickSort(List<Integer> arr, int low, int high) {
            if (low < high) {
                int pivotIndex = randomizedPartition(arr, low, high);
                randomizedQuickSort(arr, low, pivotIndex - 1);
                randomizedQuickSort(arr, pivotIndex + 1, high);
            }
        }
        
        private int randomizedPartition(List<Integer> arr, int low, int high) {
            // Random pivot selection
            int pivotIndex = low + random.nextInt(high - low + 1);
            Collections.swap(arr, pivotIndex, high);
            
            int pivot = arr.get(high);
            int i = low - 1;
            
            for (int j = low; j < high; j++) {
                if (arr.get(j) <= pivot) {
                    i++;
                    Collections.swap(arr, i, j);
                }
            }
            
            Collections.swap(arr, i + 1, high);
            return i + 1;
        }
        
        private ListNode buildLinkedList(List<Integer> values) {
            if (values.isEmpty()) {
                return null;
            }
            
            ListNode dummy = new ListNode(0);
            ListNode current = dummy;
            
            for (int val : values) {
                current.next = new ListNode(val);
                current = current.next;
            }
            
            return dummy.next;
        }
    }
    
    // Version with detailed explanation
    static class RandomizedLinkedListDetailed {
        ListNode head;
        Random random;
        List<String> explanation;
        
        public RandomizedLinkedListDetailed(ListNode head) {
            this.head = head;
            this.random = new Random();
            this.explanation = new ArrayList<>();
        }
        
        public int getRandom() {
            explanation.add("=== Reservoir Sampling Algorithm ===");
            
            if (head == null) {
                explanation.add("Empty list, returning -1");
                return -1;
            }
            
            int count = 0;
            int result = head.val;
            ListNode current = head;
            
            explanation.add("Starting reservoir sampling:");
            
            while (current != null) {
                count++;
                explanation.add(String.format("  Node %d: value = %d, count = %d", count, current.val, count));
                
                // With probability 1/count, select current node
                if (random.nextInt(count) == 0) {
                    result = current.val;
                    explanation.add(String.format("    Selected node %d (1/count probability)", count));
                } else {
                    explanation.add(String.format("    Skipping node %d", count));
                }
                
                current = current.next;
            }
            
            explanation.add(String.format("Final result: %d", result));
            explanation.add(String.format("Total nodes processed: %d", count));
            
            return result;
        }
        
        public List<String> getExplanation() {
            return new ArrayList<>(explanation);
        }
        
        public void clearExplanation() {
            explanation.clear();
        }
    }
    
    // Performance comparison
    public void comparePerformance(ListNode head, int trials) {
        System.out.println("=== Performance Comparison ===");
        System.out.println("Trials: " + trials);
        
        // Reservoir sampling
        RandomizedLinkedList rll = new RandomizedLinkedList(head);
        long startTime = System.nanoTime();
        
        for (int i = 0; i < trials; i++) {
            rll.getRandom();
        }
        
        long endTime = System.nanoTime();
        System.out.printf("Reservoir sampling: took %d ns\n", endTime - startTime);
        
        // Monte Carlo approximation
        MonteCarloLinkedList mcll = new MonteCarloLinkedList(head);
        startTime = System.nanoTime();
        
        for (int i = 0; i < trials; i++) {
            mcll.approximateCount(100);
        }
        
        endTime = System.nanoTime();
        System.out.printf("Monte Carlo approximation: took %d ns\n", endTime - startTime);
        
        // Las Vegas algorithm
        LasVegasLinkedList lvll = new LasVegasLinkedList(head);
        startTime = System.nanoTime();
        
        for (int i = 0; i < trials; i++) {
            lvll.getKthRandom(1);
        }
        
        endTime = System.nanoTime();
        System.out.printf("Las Vegas algorithm: took %d ns\n", endTime - startTime);
    }
    
    // Statistical analysis
    public void statisticalAnalysis(ListNode head, int samples) {
        System.out.println("=== Statistical Analysis ===");
        System.out.println("Samples: " + samples);
        
        RandomizedLinkedList rll = new RandomizedLinkedList(head);
        Map<Integer, Integer> frequency = new HashMap<>();
        
        for (int i = 0; i < samples; i++) {
            int value = rll.getRandom();
            frequency.put(value, frequency.getOrDefault(value, 0) + 1);
        }
        
        System.out.println("Frequency distribution:");
        for (Map.Entry<Integer, Integer> entry : frequency.entrySet()) {
            double percentage = (double) entry.getValue() / samples * 100;
            System.out.printf("  Value %d: %d times (%.2f%%)\n", 
                entry.getKey(), entry.getValue(), percentage);
        }
        
        // Calculate actual distribution
        Map<Integer, Integer> actualFrequency = new HashMap<>();
        ListNode current = head;
        int totalNodes = 0;
        
        while (current != null) {
            actualFrequency.put(current.val, actualFrequency.getOrDefault(current.val, 0) + 1);
            totalNodes++;
            current = current.next;
        }
        
        System.out.println("\nActual distribution:");
        for (Map.Entry<Integer, Integer> entry : actualFrequency.entrySet()) {
            double percentage = (double) entry.getValue() / totalNodes * 100;
            System.out.printf("  Value %d: %d times (%.2f%%)\n", 
                entry.getKey(), entry.getValue(), percentage);
        }
    }
    
    // Helper method to create linked list
    public ListNode createLinkedList(int[] values) {
        if (values.length == 0) {
            return null;
        }
        
        ListNode dummy = new ListNode(0);
        ListNode current = dummy;
        
        for (int val : values) {
            current.next = new ListNode(val);
            current = current.next;
        }
        
        return dummy.next;
    }
    
    // Helper method to print linked list
    public void printLinkedList(ListNode head) {
        List<Integer> values = new ArrayList<>();
        ListNode current = head;
        
        while (current != null) {
            values.add(current.val);
            current = current.next;
        }
        
        System.out.println(values);
    }
    
    public static void main(String[] args) {
        LinkedListRandomNode lln = new LinkedListRandomNode();
        
        // Test cases
        int[][] testCases = {
            {1, 2, 3, 4, 5},
            {1, 1, 1, 1, 1},
            {5, 4, 3, 2, 1},
            {10, 20, 30},
            {100},
            {1, 100, 50, 25, 75},
            {0, 0, 0, 0},
            {-1, -2, -3, -4, -5},
            {7, 7, 7, 7, 7, 7}
        };
        
        String[] descriptions = {
            "Sequential values",
            "All same values",
            "Descending values",
            "Large values",
            "Single element",
            "Mixed values",
            "All zeros",
            "All negative",
            "Many duplicates"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.println("Array: " + Arrays.toString(testCases[i]));
            
            ListNode head = lln.createLinkedList(testCases[i]);
            lln.printLinkedList(head);
            
            // Test random selection
            RandomizedLinkedList rll = new RandomizedLinkedList(head);
            System.out.println("Random selections:");
            for (int j = 0; j < 5; j++) {
                System.out.printf("  Selection %d: %d\n", j + 1, rll.getRandom());
            }
            
            // Test Monte Carlo approximation
            MonteCarloLinkedList mcll = new MonteCarloLinkedList(head);
            int actualLength = testCases[i].length;
            int estimatedLength = mcll.approximateCount(1000);
            System.out.printf("Actual length: %d, Estimated: %d, Error: %.2f%%\n", 
                actualLength, estimatedLength, Math.abs(estimatedLength - actualLength) * 100.0 / actualLength);
            
            // Test Las Vegas algorithm
            LasVegasLinkedList lvll = new LasVegasLinkedList(head);
            System.out.printf("Kth random (k=3): %d\n", lvll.getKthRandom(3));
            
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        ListNode detailedHead = lln.createLinkedList(new int[]{1, 2, 3, 4, 5});
        RandomizedLinkedListDetailed rllDetailed = new RandomizedLinkedListDetailed(detailedHead);
        
        int randomResult = rllDetailed.getRandom();
        System.out.printf("Random result: %d\n", randomResult);
        
        for (String step : rllDetailed.getExplanation()) {
            System.out.println(step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        ListNode performanceHead = lln.createLinkedList(new int[1000]);
        for (int i = 0; i < 1000; i++) {
            performanceHead = new ListNode(i % 100);
        }
        
        lln.comparePerformance(performanceHead, 10000);
        
        // Statistical analysis
        System.out.println("\n=== Statistical Analysis ===");
        ListNode analysisHead = lln.createLinkedList(new int[]{1, 2, 3, 4, 5, 1, 2, 3});
        lln.statisticalAnalysis(analysisHead, 10000);
        
        // Sherwood algorithm test
        System.out.println("\n=== Sherwood Algorithm Test ===");
        SherwoodLinkedList sll = new SherwoodLinkedList(lln.createLinkedList(new int[]{5, 3, 8, 1, 9, 2}));
        System.out.print("Before sorting: ");
        lln.printLinkedList(sll.head);
        
        sll.randomizeQuickSort();
        System.out.print("After sorting: ");
        lln.printLinkedList(sll.head);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        
        // Empty list
        RandomizedLinkedList emptyRll = new RandomizedLinkedList(null);
        System.out.printf("Empty list random: %d\n", emptyRll.getRandom());
        
        // Single element
        ListNode singleNode = new ListNode(42);
        RandomizedLinkedList singleRll = new RandomizedLinkedList(singleNode);
        System.out.printf("Single element random: %d\n", singleRll.getRandom());
        
        // Large list
        int[] largeArray = new int[10000];
        for (int i = 0; i < 10000; i++) {
            largeArray[i] = i % 100;
        }
        ListNode largeHead = lln.createLinkedList(largeArray);
        RandomizedLinkedList largeRll = new RandomizedLinkedList(largeHead);
        
        long startTime = System.nanoTime();
        for (int i = 0; i < 1000; i++) {
            largeRll.getRandom();
        }
        long endTime = System.nanoTime();
        System.out.printf("Large list (10000 elements, 1000 selections): took %d ns\n", endTime - startTime);
    }
}
