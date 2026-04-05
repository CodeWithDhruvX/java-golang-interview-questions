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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Randomized Algorithms
- **Reservoir Sampling**: Random selection from stream of unknown size
- **Monte Carlo**: Probabilistic approximation algorithms
- **Las Vegas**: Always correct but randomized runtime
- **Sherwood**: Randomized pivot selection for sorting

## 2. PROBLEM CHARACTERISTICS
- **Linked List Random Access**: Need O(1) random access in linked list
- **Unknown Size**: List size may not be known in advance
- **Probabilistic Methods**: Use randomness for efficiency
- **Approximation**: Accept probabilistic results for speed

## 3. SIMILAR PROBLEMS
- Random Sampling from Data Stream
- Approximate Counting Algorithms
- Randomized Quick Sort
- Probabilistic Data Structures

## 4. KEY OBSERVATIONS
- Reservoir sampling enables O(1) space random selection
- Monte Carlo provides approximations with bounded error
- Las Vegas algorithms are always correct but randomized
- Randomized algorithms improve average-case performance
- Time complexity varies by algorithm type

## 5. VARIATIONS & EXTENSIONS
- Different sampling probabilities
- Weighted random selection
- Multiple random queries
- Parallel random algorithms

## 6. INTERVIEW INSIGHTS
- Clarify: "Can we modify the list?"
- Edge cases: empty list, single element, large lists
- Time complexity: varies by algorithm type
- Space complexity: O(1) for most algorithms

## 7. COMMON MISTAKES
- Incorrect random number generation
- Wrong reservoir sampling logic
- Biased random selection
- Incorrect probability calculations
- Not handling edge cases properly

## 8. OPTIMIZATION STRATEGIES
- Use proper random number generation
- Efficient reservoir sampling
- Correct probability calculations
- Memory-efficient implementations

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like drawing random names from a hat:**
- You have a linked list and need to select random elements
- Can't access elements by index like an array
- Need algorithms that work with sequential access
- Reservoir sampling maintains uniform probability without knowing size
- Monte Carlo uses random sampling to estimate properties
- Las Vegas algorithms use randomness but guarantee correctness

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Linked list, need random element selection
2. **Goal**: Select random element with uniform probability
3. **Output**: Random element value

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → No O(1) random access in linked list
- **"How to optimize?"** → Use reservoir sampling for unknown size
- **"Why reservoir sampling?"** → Maintains uniform probability
- **"How to handle large lists?"** → Use approximation algorithms

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use multiple randomized algorithms:
1. Reservoir Sampling: O(N) time, O(1) space
   - Process list once, maintain reservoir of size k
   - Each element has 1/k probability of being selected
2. Monte Carlo: O(N*sqrt(N)) time, O(1) space
   - Random walk to estimate list length
   - Use statistical properties for approximation
3. Las Vegas: O(N) time, O(N) space
   - Collect all elements, shuffle, select kth
   - Always correct but uses O(N) space
4. Sherwood: O(N log N) expected time
   - Randomized quicksort with random pivots
   - Improves worst-case behavior"
```

#### Phase 4: Edge Case Handling
- **Empty list**: Return -1 or handle appropriately
- **Single element**: Return the only element
- **Large lists**: Use efficient algorithms
- **Invalid k**: Handle edge cases

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Linked List: 1 → 2 → 3 → 4 → 5

Human thinking:
"Let's apply reservoir sampling for k=1:

Initialize: count=0, result=head.val (1), current=head

Process node 1 (value 1):
- count=1
- random.nextInt(1) == 0 → select node 1
- result=1

Process node 2 (value 2):
- count=2
- random.nextInt(2) == 0 with probability 1/2
- If selected: result=2, else keep result=1

Process node 3 (value 3):
- count=3
- random.nextInt(3) == 0 with probability 1/3
- If selected: result=3, else keep previous

Continue until end...

Final result: Each node has 1/N probability of being selected ✓

Monte Carlo approximation:
Random walk with 50% probability of stopping at each step
Average steps ≈ N/2
Estimated length ≈ 2 * average steps

Las Vegas algorithm:
Collect all elements: [1,2,3,4,5]
Shuffle: [3,1,5,2,4]
Select kth element: guaranteed uniform ✓

Sherwood algorithm:
Randomized quicksort with random pivots
Expected O(N log N) time vs O(N²) worst case ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Each algorithm maintains uniform probability
- **Why it's efficient**: Avoids O(N) random access simulation
- **Why it's correct**: Mathematical proof of uniform distribution

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just count and use random?"** → Requires O(N) traversal first
2. **"What about bias?"** → Must ensure uniform probability
3. **"How to handle reservoir size?"** → Maintain proper replacement logic
4. **"What about approximation error?"** → Bound error mathematically

### Real-World Analogy
**Like conducting a random survey without knowing population size:**
- You have people entering a venue (linked list)
- Need to select random people for a survey
- Can't count everyone first (too slow)
- Reservoir sampling: maintain sample, replace with probability
- Monte Carlo: estimate total from random sample
- Las Vegas: collect everyone, then randomize
- Useful in streaming data, network sampling, load balancing
- Like conducting fair random selection from unknown population

### Human-Readable Pseudocode
```
function reservoirSampling(head, k):
    if head == null: return -1
    
    count = 0
    result = head.val
    current = head
    
    while current != null:
        count++
        // With probability 1/count, select current
        if random.nextInt(count) == 0:
            result = current.val
        current = current.next
    
    return result

function monteCarloApproximation(head, iterations):
    if head == null: return 0
    
    totalSteps = 0
    for i from 1 to iterations:
        steps = randomWalk(head)
        totalSteps += steps
    
    return totalSteps * 2 / iterations

function lasVegasSelection(head, k):
    values = []
    current = head
    
    while current != null:
        values.add(current.val)
        current = current.next
    
    shuffle(values)
    return values[k-1] if k <= values.length else -1
```

### Execution Visualization

### Example: Linked List 1→2→3→4→5, k=1
```
Reservoir Sampling Process:

Initialize: count=0, result=1, current=1

Step 1: Node 1 (value 1)
- count=1, random.nextInt(1)=0 → select
- result=1

Step 2: Node 2 (value 2)
- count=2, random.nextInt(2) in {0,1}
- P(select) = 1/2, P(keep) = 1/2

Step 3: Node 3 (value 3)
- count=3, random.nextInt(3) in {0,1,2}
- P(select) = 1/3, P(keep) = 2/3

Continue...

Final: Each node has 1/N selection probability ✓

Monte Carlo Approximation:
Random walk with 50% stopping probability:
Expected steps per walk = N/2 = 2.5
Estimated length = 2 * 2.5 = 5
Actual length = 5 ✓

Las Vegas Algorithm:
Collect: [1,2,3,4,5]
Shuffle: [3,1,5,2,4]
Select 1st: 3 (uniform probability 1/5) ✓

Sherwood Algorithm:
Randomized quicksort with random pivots:
Expected O(N log N) vs O(N²) worst case ✓
```

### Key Visualization Points:
- **Reservoir Sampling**: Maintains uniform probability with O(1) space
- **Monte Carlo**: Probabilistic approximation with bounded error
- **Las Vegas**: Always correct but uses O(N) space
- **Sherwood**: Randomized to improve worst-case behavior

### Memory Layout Visualization:
```
Linked List: 1 → 2 → 3 → 4 → 5

Reservoir Sampling:
count=0, result=1
count=1, result=1 or 2 (50% each)
count=2, result=1,2, or 3 (33% each)
count=3, result=1,2,3, or 4 (25% each)
count=4, result=1,2,3,4, or 5 (20% each)

Monte Carlo:
Random walk: average 2.5 steps
Estimated length: 5

Las Vegas:
Collected: [1,2,3,4,5]
Shuffled: [3,1,5,2,4]
Selected: 3

Each algorithm provides different trade-offs between time, space, and accuracy
```

### Time Complexity Breakdown:
- **Reservoir Sampling**: O(N) time, O(1) space
- **Monte Carlo**: O(N*sqrt(N)) time, O(1) space
- **Las Vegas**: O(N) time, O(N) space
- **Sherwood**: O(N log N) expected time, O(N) space
- **Trade-offs**: Time vs space vs accuracy
- **Optimal**: Choose based on specific requirements
*/
