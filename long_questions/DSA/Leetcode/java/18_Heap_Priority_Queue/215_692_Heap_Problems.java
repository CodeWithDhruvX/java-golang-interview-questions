import java.util.*;

public class KthLargestElementInAnArray {
    
    // 215. Kth Largest Element in an Array
    // Time: O(N log K), Space: O(K)
    public static int findKthLargest(int[] nums, int k) {
        if (nums == null || nums.length == 0 || k <= 0 || k > nums.length) {
            return Integer.MIN_VALUE;
        }
        
        // Use min heap to keep k largest elements
        PriorityQueue<Integer> minHeap = new PriorityQueue<>();
        
        for (int num : nums) {
            minHeap.offer(num);
            
            if (minHeap.size() > k) {
                minHeap.poll();
            }
        }
        
        return minHeap.peek();
    }

    // 692. Top K Frequent Words
    // Time: O(N log K), Space: O(N)
    public static List<String> topKFrequent(String[] words, int k) {
        if (words == null || words.length == 0 || k <= 0) {
            return new ArrayList<>();
        }
        
        // Count frequency of each word
        Map<String, Integer> frequencyMap = new HashMap<>();
        for (String word : words) {
            frequencyMap.put(word, frequencyMap.getOrDefault(word, 0) + 1);
        }
        
        // Use max heap with custom comparator (frequency desc, word asc)
        PriorityQueue<Map.Entry<String, Integer>> maxHeap = new PriorityQueue<>(
            (a, b) -> {
                if (!a.getValue().equals(b.getValue())) {
                    return b.getValue() - a.getValue(); // Higher frequency first
                }
                return a.getKey().compareTo(b.getKey()); // Lexicographically smaller first
            }
        );
        
        maxHeap.addAll(frequencyMap.entrySet());
        
        List<String> result = new ArrayList<>();
        for (int i = 0; i < k && !maxHeap.isEmpty(); i++) {
            result.add(maxHeap.poll().getKey());
        }
        
        return result;
    }

    public static void main(String[] args) {
        // Test cases for findKthLargest
        Object[][] testCases1 = {
            {new int[]{3, 2, 1, 5, 6, 4}, 2},
            {new int[]{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
            {new int[]{1}, 1},
            {new int[]{1, 2}, 2},
            {new int[]{2, 1}, 2},
            {new int[]{1, 2, 3, 4, 5}, 1},
            {new int[]{1, 2, 3, 4, 5}, 5},
            {new int[]{5, 4, 3, 2, 1}, 3},
            {new int[]{1, 1, 1, 1, 1}, 3},
            {new int[]{1, 2, 2, 3, 3, 3}, 2}
        };
        
        // Test cases for topKFrequent
        Object[][] testCases2 = {
            {new String[]{"i", "love", "leetcode", "i", "love", "coding"}, 2},
            {new String[]{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4},
            {new String[]{"a", "a", "a", "a", "a"}, 1},
            {new String[]{"a", "b", "c", "d", "e"}, 3},
            {new String[]{"apple", "banana", "apple", "orange", "banana", "apple"}, 2},
            {new String[]{"hello", "world", "hello", "leetcode"}, 1},
            {new String[]{"dog", "cat", "dog", "cat", "dog"}, 2},
            {new String[]{"a", "b", "a", "b", "c", "c", "c"}, 1},
            {new String[]{"x", "y", "z", "x", "y", "x"}, 2},
            {new String[]{"abc", "def", "ghi", "abc", "def", "abc"}, 3}
        };
        
        System.out.println("Kth Largest Element:");
        for (int i = 0; i < testCases1.length; i++) {
            int[] nums = (int[]) testCases1[i][0];
            int k = (int) testCases1[i][1];
            int result = findKthLargest(nums, k);
            System.out.printf("Test Case %d: %s, k=%d -> %d\n", 
                i + 1, Arrays.toString(nums), k, result);
        }
        
        System.out.println("\nTop K Frequent Words:");
        for (int i = 0; i < testCases2.length; i++) {
            String[] words = (String[]) testCases2[i][0];
            int k = (int) testCases2[i][1];
            List<String> result = topKFrequent(words, k);
            System.out.printf("Test Case %d: %s, k=%d -> %s\n", 
                i + 1, Arrays.toString(words), k, result);
        }
    }
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Heap-Based Selection and Frequency Analysis
- **Kth Largest**: Min-heap of size K to track top K elements
- **Top K Frequent**: Max-heap with custom comparator for frequency ordering
- **Heap Property**: Maintains partial ordering for efficient selection
- **Frequency Counting**: Map to count occurrences before heap selection

## 2. PROBLEM CHARACTERISTICS
- **Selection Problem**: Find Kth largest element in unsorted array
- **Frequency Analysis**: Find most frequent K items
- **Partial Ordering**: Don't need full sort, just top K elements
- **Efficiency Requirement**: Better than O(N log N) full sort for large N

## 3. SIMILAR PROBLEMS
- Find Median from Data Stream
- Merge K Sorted Lists
- Sliding Window Maximum
- Priority Queue scheduling

## 4. KEY OBSERVATIONS
- **Min-Heap for Kth Largest**: Keeps only K largest elements
- **Max-Heap for Top K Frequent**: Orders by frequency then lexicographically
- **Space-Time Tradeoff**: O(N log K) vs O(N log N) for full sort
- **Custom Comparator**: Needed for tie-breaking in frequency problems

## 5. VARIATIONS & EXTENSIONS
- Find Kth smallest element
- Dynamic K (changes over time)
- Range queries on heap
- Distributed heap implementations

## 6. INTERVIEW INSIGHTS
- Clarify: "Should I handle duplicates in Kth largest?"
- Edge cases: K > array length, empty array, K = 0
- Time complexity: O(N log K) is optimal for selection
- Space complexity: O(K) for heap, O(N) for frequency map

## 7. COMMON MISTAKES
- Using max-heap for Kth largest (inefficient)
- Not handling edge cases (empty array, invalid K)
- Incorrect comparator for frequency problems
- Forgetting to handle tie-breaking rules

## 8. OPTIMIZATION STRATEGIES
- Use min-heap for Kth largest instead of max-heap
- Quickselect algorithm for O(N) average time
- Pre-allocate heap capacity when possible
- Use primitive arrays for better cache performance

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding top performers in a competition:**
- For Kth largest: Keep only the best K performers seen so far
- For top K frequent: Count how many times each performer scored
- Use a priority system (heap) to always know who's on top
- Eliminate lower performers when better ones appear

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding (Kth Largest)
1. **Input**: Unsorted array, integer K
2. **Goal**: Find Kth largest element (1-indexed)
3. **Output**: The Kth largest value
4. **Constraint**: K is valid (1 ≤ K ≤ array length)

#### Phase 2: Key Insight Recognition
- **"How to find Kth largest efficiently?"** → Don't sort entire array
- **"What data structure helps?"** → Min-heap of size K
- **"How does heap help?"** → Always keeps K largest seen so far
- **"When to update heap?"** → When size exceeds K, remove smallest

#### Phase 3: Strategy Development (Kth Largest)
```
Human thought process:
"I'll use a min-heap to track top K elements:
1. Create empty min-heap
2. For each number in array:
   - Add to heap
   - If heap size > K: remove smallest (maintains K largest)
3. At end: heap.peek() is Kth largest
4. This is O(N log K) instead of O(N log N) for full sort"
```

#### Phase 4: Algorithm Walkthrough (Kth Largest)
```
Example: nums=[3,2,1,5,6,4], K=3

Human thinking:
"Start with empty heap

Process 3: heap=[3], size=1
Process 2: heap=[2,3], size=2  
Process 1: heap=[1,2,3], size=3
Process 5: heap=[1,2,3,5], size=4 > K=3
Remove smallest (1): heap=[2,3,5], size=3
Process 6: heap=[2,3,5,6], size=4 > K=3
Remove smallest (2): heap=[3,5,6], size=3
Process 4: heap=[3,4,5,6], size=4 > K=3
Remove smallest (3): heap=[4,5,6], size=3

Final heap: [4,5,6], peek()=4
So 4th largest (K=3) is 4 ✓"
```

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just sort?"** → O(N log N) vs O(N log K), K << N typically
2. **"Min-heap vs Max-heap?"** → Min-heap keeps K largest efficiently
3. **"What about duplicates?"** → Handle according to problem requirements
4. **"How to handle invalid K?"** → Return appropriate value or throw exception

### Real-World Analogy
**Like managing a leaderboard in a game:**
- Players have scores (array elements)
- You want to know who's ranked Kth
- Keep a top-K board (min-heap) showing best players
- When new player scores high, add to board
- If board is full, remove the lowest score
- Top of board always shows Kth best score

### Human-Readable Pseudocode
```
function findKthLargest(nums, k):
    if nums is empty or k <= 0: return error
    
    minHeap = empty min-heap
    
    for each num in nums:
        minHeap.offer(num)
        if minHeap.size() > k:
            minHeap.poll()  // remove smallest
    
    return minHeap.peek()  // Kth largest

function topKFrequent(words, k):
    if words is empty or k <= 0: return []
    
    frequencyMap = count word frequencies
    maxHeap = max-heap with custom comparator:
        - Higher frequency first
        - If same frequency: lexicographically smaller first
    
    for each word, count in frequencyMap:
        maxHeap.offer([word, count])
    
    result = []
    repeat k times:
        if maxHeap.isEmpty(): break
        result.add(maxHeap.poll()[0])
    
    return result
```

### Execution Visualization

### Example: nums=[3,2,1,5,6,4], K=3
```
Heap Evolution:
Step 1: [3]
Step 2: [2,3]
Step 3: [1,2,3]  (size=3=K)
Step 4: [1,2,3,5] → remove 1 → [2,3,5]
Step 5: [2,3,5,6] → remove 2 → [3,5,6]
Step 6: [3,4,5,6] → remove 3 → [4,5,6]

Final heap: [4,5,6]
Kth largest (3rd) = heap.peek() = 4 ✓
```

### Example: words=["i","love","leetcode","i","love","coding"], K=2
```
Frequency Map: {i:2, love:2, leetcode:1, coding:1}

Max-Heap (frequency, then lexicographic):
["i",2], ["love",2], ["coding",1], ["leetcode",1]

Extract K=2:
1. ["i",2] → result=["i"]
2. ["love",2] → result=["i","love"]

Final: ["i","love"] ✓
```

### Key Visualization Points:
- **Heap Property**: Maintains partial ordering efficiently
- **Size Management**: Keep exactly K elements for selection
- **Custom Comparator**: Handles tie-breaking rules
- **Frequency Analysis**: Count first, then select top K

### Time Complexity Breakdown:
- **Kth Largest**: O(N log K) time, O(K) space
- **Top K Frequent**: O(N) for frequency + O(N log K) for heap, O(N) space
- **Better than Sorting**: When K << N, O(N log K) < O(N log N)
- **Optimal**: Cannot do better than O(N) for selection without preprocessing
*/
